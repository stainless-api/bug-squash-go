package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/acme/acme-go"
	"github.com/acme/acme-go/option"
)

func main() {
	client := acme.NewClient(option.WithAPIKey("bad_api_key"))
	_, err := client.Cards.List(context.TODO(), acme.CardListParams{
		AccountID: acme.F("account_in71c4amph0vgo2qllky"),
	})
	if err == nil {
		panic("expected an error with an invalid API key")
	}

	var apierr *acme.Error
	if errors.As(err, &apierr) {
		fmt.Printf("%+#v\n", apierr)
		println(apierr.Error())
		println(apierr.Detail)
	}
}
