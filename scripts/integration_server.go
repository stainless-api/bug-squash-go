package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/shared"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func sendError(w http.ResponseWriter, error *increase.Error) {
	errorB, err := json.Marshal(error)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(int(error.Status))
	w.Write(errorB)
}

func rootFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /\n")
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, "{\"status\": \"ok\"}\n")
}

func checkAuthorization(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("Authorization") != "Bearer sk_test_1234567890" {
		error := &increase.Error{
			Detail: "See https://increase.com/documentation/api#authorization-and-testing",
			Status: 403,
			Title:  "Incorrect or missing Authorization",
			Type:   "insufficient_permissions_error",
		}
		sendError(w, error)
		return true
	}
	return false
}

type CachedRequest struct {
	Method         string
	URL            *url.URL
	Body           []byte
	ResponseStatus int
	ResponseBody   []byte
}

var cachedRequests map[string]*CachedRequest

func init() {
	cachedRequests = make(map[string]*CachedRequest)
}

func checkIdempotency(w http.ResponseWriter, r *http.Request) bool {
	ikey := r.Header.Get("Idempotency-Key")
	// No idempotency in play if no key given.
	if ikey == "" {
		return false
	}
	cr, cached := cachedRequests[ikey]
	// No cached idempotent request to return.
	if !cached {
		return false
	}
	// We're seeing an idempotency key again; make sure it isn't
	// using a different request shape than last time, before
	// returning the cached response.
	rBody, err := io.ReadAll(r.Body)
	check(err)
	r.Body = io.NopCloser(bytes.NewReader(rBody))
	mismatch := false
	if cr.Method != r.Method {
		mismatch = true
	}
	if cr.URL.Path != r.URL.Path {
		mismatch = true
	}
	if cr.URL.RawQuery != r.URL.RawQuery {
		mismatch = true
	}
	if !bytes.Equal(cr.Body, rBody) {
		mismatch = true
	}
	if mismatch {
		error := &increase.Error{
			Detail: "Idempotency token reused with dissimilar request",
			Status: 422,
			Title:  "Invalid idempotency configuration",
			Type:   "idempotency_unprocessable_error",
		}
		sendError(w, error)
	} else {
		w.Header().Add("Idempotent-Replayed", "true")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(cr.ResponseStatus)
		if cr.ResponseBody != nil {
			w.Write(cr.ResponseBody)
		}
	}
	return true
}

var checkDepositNextID = 0
var allCheckDeposits []*increase.CheckDeposit

func init() {
	allCheckDeposits = make([]*increase.CheckDeposit, 5)
	for i := 0; i < 5; i++ {
		allCheckDeposits[i] = &increase.CheckDeposit{
			ID:               fmt.Sprintf("check_deposit_%08d", checkDepositNextID),
			CreatedAt:        time.Now(),
			AccountID:        fmt.Sprintf("account_%08d", rand.Int()),
			Amount:           42,
			Currency:         "USD",
			Status:           "submitted",
			FrontImageFileID: fmt.Sprintf("file_%08d", rand.Int()),
			BackImageFileID:  fmt.Sprintf("file_%08d", rand.Int()),
			Type:             "check_deposit",
		}
		checkDepositNextID += 1
	}
}

func getCheckDepositsFunc(w http.ResponseWriter, r *http.Request) {
	// Simulate processing a list request, doing basic limit and cursor
	// processing. Filter params not yet supported and will error.
	query := r.URL.Query()
	limit := 100
	cursor := ""
	var err error
	for queryKey, queryVals := range query {
		if queryKey == "limit" {
			limit, err = strconv.Atoi(queryVals[0])
			check(err)
		} else if queryKey == "cursor" {
			cursor = queryVals[0]
		} else {
			panic(fmt.Sprintf("Unexpected query param '%s'", queryKey))
		}
	}
	start := 0
	for i := 0; i < len(allCheckDeposits); i++ {
		if allCheckDeposits[i].ID > cursor {
			start = i
			break
		}
	}
	end := start + limit - 1
	if end > len(allCheckDeposits)-1 {
		end = len(allCheckDeposits) - 1
	}
	nextCursor := ""
	if end < len(allCheckDeposits)-1 {
		nextCursor = (allCheckDeposits[end]).ID
	}
	data := allCheckDeposits[start : end+1]
	page := &shared.Page[*increase.CheckDeposit]{
		Data:       data,
		NextCursor: nextCursor,
	}
	pageB, err := json.Marshal(page)
	check(err)
	w.Header().Add("Content-Type", "application/json")
	w.Write(pageB)

	// Cache requests for future idempotent attempts.
	// TODO: Discuss dropping idempotency for GET requests.
	if ikey := r.Header.Get("Idempotency-Key"); ikey != "" {
		cachedRequests[ikey] = &CachedRequest{
			Method:         "GET",
			URL:            r.URL,
			Body:           nil,
			ResponseStatus: 200,
			ResponseBody:   pageB,
		}
	}
}

var seenBodies map[string]bool

func init() {
	seenBodies = make(map[string]bool)
}

func postCheckDepositsFunc(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	check(err)

	// Simulate processing create request.
	cd := &increase.CheckDeposit{}
	err = json.Unmarshal(body, cd)
	cd.ID = fmt.Sprintf("check_deposit_%08d", checkDepositNextID)
	checkDepositNextID += 1
	cd.CreatedAt = time.Now()
	cd.Status = "submitted"
	cd.Type = "check_deposit"
	cdB, err := json.Marshal(cd)

	// Store a rolling window of recent deposits so that the list stays small.
	allCheckDeposits = append(allCheckDeposits[1:], cd)

	// Cache requests for future idempotent attempts.
	if ikey := r.Header.Get("Idempotency-Key"); ikey != "" {
		fmt.Printf("Caching: %s\n", ikey)
		cachedRequests[ikey] = &CachedRequest{
			Method:         "POST",
			URL:            r.URL,
			Body:           body,
			ResponseStatus: 200,
			ResponseBody:   cdB,
		}
	}

	// Always fail the first time we see a unique body, to exercise retries.
	if !seenBodies[string(body)] {
		seenBodies[string(body)] = true
		w.WriteHeader(500)
		return
	}

	// Otherwise OK to respond with success.
	check(err)
	w.Header().Add("Content-Type", "application/json")
	w.Write(cdB)
}

func checkDepositsFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s /check_deposits\n", r.Method)
	if checkAuthorization(w, r) {
		return
	}
	if checkIdempotency(w, r) {
		return
	}
	if r.Method == "GET" {
		getCheckDepositsFunc(w, r)
	} else if r.Method == "POST" {
		postCheckDepositsFunc(w, r)
	} else {
		w.WriteHeader(404)
		return
	}
}

func main() {
	fmt.Printf("Listening on :8077\n")
	http.HandleFunc("/", rootFunc)
	http.HandleFunc("/check_deposits", checkDepositsFunc)
	err := http.ListenAndServe(":8077", nil)
	check(err)
}
