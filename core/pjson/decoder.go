package pjson

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"github.com/tidwall/gjson"
)

var decoders sync.Map // map[reflect.Type]decoderFunc

func Unmarshal(raw []byte, to any) error {
	d := &decoder{}
	value := reflect.ValueOf(to).Elem()
	result := gjson.ParseBytes(raw)
	if !value.IsValid() {
		return fmt.Errorf("pjson: cannot marshal into invalid value")
	}
	return d.typeDecoder(value.Type())(result, value)
}

type decoder struct{}

type decoderFunc func(node gjson.Result, value reflect.Value) error

type decoderField struct {
	tag parsedStructTag
	fn  decoderFunc
	idx []int
}

func (d *decoder) typeDecoder(t reflect.Type) decoderFunc {
	if fi, ok := decoders.Load(t); ok {
		return fi.(decoderFunc)
	}

	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  decoderFunc
	)
	wg.Add(1)
	fi, loaded := decoders.LoadOrStore(t, decoderFunc(func(node gjson.Result, v reflect.Value) error {
		wg.Wait()
		return f(node, v)
	}))
	if loaded {
		return fi.(decoderFunc)
	}

	// Compute the real decoder and replace the indirect func with it.
	f = d.newTypeDecoder(t)
	wg.Done()
	decoders.Store(t, f)
	return f
}

func (d *decoder) newTypeDecoder(t reflect.Type) decoderFunc {
	switch t.Kind() {
	case reflect.Pointer:
		inner := t.Elem()
		innerDecoder := d.typeDecoder(inner)

		return func(n gjson.Result, v reflect.Value) error {
			if !v.IsValid() {
				return fmt.Errorf("pjson: unexpected invalid reflection value %+#v", v)
			}

			newValue := reflect.New(inner).Elem()
			err := innerDecoder(n, newValue)
			if err != nil {
				return err
			}

			v.Set(newValue.Addr())
			return nil
		}
	case reflect.Struct:
		return d.newStructTypeDecoder(t)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return d.newArrayTypeDecoder(t)
	case reflect.Map:
		return d.newMapDecoder(t)
	case reflect.Interface:
		return func(node gjson.Result, value reflect.Value) error {
			if !value.IsValid() {
				return fmt.Errorf("pjson: unexpected invalid value %+#v", value)
			}
			if node.Value() != nil {
				value.Set(reflect.ValueOf(node.Value()))
			}
			return nil
		}
	default:
		return d.newPrimitiveTypeDecoder(t)
	}
}

func (d *decoder) newMapDecoder(t reflect.Type) decoderFunc {
	keyType := t.Key()
	itemType := t.Elem()
	itemDecoder := d.typeDecoder(itemType)

	return func(node gjson.Result, value reflect.Value) (err error) {
		mapValue := reflect.MakeMapWithSize(t, len(node.Map()))

		node.ForEach(func(key, value gjson.Result) bool {
			// It's fine for us to just use `ValueOf` here because the key types will
			// always be primitive types so we don't need to decode it using the standard pattern
			keyValue := reflect.ValueOf(key.Value())
			if keyValue.Type() != keyType {
				err = fmt.Errorf("pjson: expected key type %v but got %v", keyType, keyValue.Type())
				return false
			}

			itemValue := reflect.New(itemType).Elem()
			err = itemDecoder(value, itemValue)
			if err != nil {
				return false
			}

			mapValue.SetMapIndex(keyValue, itemValue)
			return true
		})

		if err != nil {
			return err
		}
		value.Set(mapValue)
		return nil
	}
}

func (d *decoder) newArrayTypeDecoder(t reflect.Type) decoderFunc {
	itemDecoder := d.typeDecoder(t.Elem())

	return func(node gjson.Result, value reflect.Value) (err error) {
		arrayNode := node.Array()

		arrayValue := reflect.MakeSlice(reflect.SliceOf(t.Elem()), len(arrayNode), len(arrayNode))
		for i, itemNode := range arrayNode {
			err = itemDecoder(itemNode, arrayValue.Index(i))
			if err != nil {
				return err
			}
		}

		value.Set(arrayValue)
		return nil
	}
}

func (d *decoder) newStructTypeDecoder(t reflect.Type) decoderFunc {
	// map of json field name to struct field decoders
	decoderFields := map[string]decoderField{}
	extraDecoder := (*decoderField)(nil)

	// This helper allows us to recursively collect field encoders into a flat
	// array. The parameter `index` keeps track of the access patterns necessary
	// to get to some field.
	var collectFieldDecoders func(r reflect.Type, index []int)
	collectFieldDecoders = func(r reflect.Type, index []int) {
		for i := 0; i < r.NumField(); i++ {
			idx := append(index, i)
			field := t.FieldByIndex(idx)
			// If this is an embedded struct, traverse one level deeper to extract
			// the fields and get their encoders as well.
			if field.Anonymous {
				collectFieldDecoders(field.Type, idx)
				continue
			}
			tag, ok := field.Tag.Lookup(pjsonStructTag)
			if !ok {
				continue
			}
			ptag := parseStructTag(tag)
			// We only want to support unexported fields if they're tagged with
			// `extras` because that field shouldn't be part of the public API. We
			// also want to only keep the top level extras
			if ptag.storeExtraProperties && len(index) == 0 {
				extraDecoder = &decoderField{ptag, nil, idx}
				continue
			}
			if ptag.storeExtraProperties || !field.IsExported() {
				continue
			}
			decoderFields[ptag.name] = decoderField{ptag, d.typeDecoder(field.Type), idx}
		}
	}
	collectFieldDecoders(t, []int{})

	return func(node gjson.Result, value reflect.Value) (err error) {
		extraFields := JSONExtras{}

		for fieldName, itemNode := range node.Map() {
			df, ok := decoderFields[fieldName]
			if !ok {
				extraFields[fieldName] = itemNode.Value()
				continue
			}
			if df.tag.storeExtraProperties {
				panic("unexpected field decoder tagged with `extra`")
			}
			err = df.fn(itemNode, value.FieldByIndex(df.idx))
			if err != nil {
				return err
			}
		}

		if len(extraFields) == 0 {
			return nil
		}

		if extraDecoder != nil {
			makeSetable(value.FieldByIndex(extraDecoder.idx)).Set(reflect.ValueOf(extraFields))
		}

		return nil
	}
}

func (d *decoder) newPrimitiveTypeDecoder(t reflect.Type) decoderFunc {
	switch t.Kind() {
	case reflect.String:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetString(n.String())
			return nil
		}
	case reflect.Bool:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetBool(n.Bool())
			return nil
		}
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetInt(n.Int())
			return nil
		}
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetUint(n.Uint())
			return nil
		}
	case reflect.Float32, reflect.Float64:
		return func(n gjson.Result, v reflect.Value) error {
			v.SetFloat(n.Float())
			return nil
		}
	default:
		return func(node gjson.Result, v reflect.Value) error {
			return fmt.Errorf("unknown type received at primitive decoder: %s", t.String())
		}
	}
}

// Given a reflect.Value that points to an unexported struct field, returns a new reflect.Value
// that will not panic when v.Set() is called.
//
// https://stackoverflow.com/questions/17981651/is-there-any-way-to-access-private-fields-of-a-struct-from-another-package#comment74658463_17982725
func makeSetable(v reflect.Value) reflect.Value {
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}
