// File generated from our OpenAPI spec by Stainless.

package acme_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/acme/acme-go"
	"github.com/acme/acme-go/internal/testutil"
	"github.com/acme/acme-go/option"
)

func TestBookkeepingEntrySetNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := acme.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.BookkeepingEntrySets.New(context.TODO(), acme.BookkeepingEntrySetNewParams{
		Entries: acme.F([]acme.BookkeepingEntrySetNewParamsEntry{{
			AccountID: acme.F("bookkeeping_account_9husfpw68pzmve9dvvc7"),
			Amount:    acme.F(int64(100)),
		}, {
			AccountID: acme.F("bookkeeping_account_t2obldz1rcu15zr54umg"),
			Amount:    acme.F(int64(-100)),
		}}),
		Date:          acme.F(time.Now()),
		TransactionID: acme.F("transaction_uyrp7fld2ium70oa7oi"),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBookkeepingEntrySetGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := acme.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.BookkeepingEntrySets.Get(context.TODO(), "bookkeeping_entry_set_n80c6wr2p8gtc6p4ingf")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBookkeepingEntrySetListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := acme.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.BookkeepingEntrySets.List(context.TODO(), acme.BookkeepingEntrySetListParams{
		Cursor:        acme.F("string"),
		Limit:         acme.F(int64(1)),
		TransactionID: acme.F("string"),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
