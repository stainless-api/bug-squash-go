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

func TestPendingTransactionGet(t *testing.T) {
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
	_, err := client.PendingTransactions.Get(context.TODO(), "pending_transaction_k1sfetcau2qbvjbzgju4")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPendingTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.PendingTransactions.List(context.TODO(), acme.PendingTransactionListParams{
		AccountID: acme.F("string"),
		Category: acme.F(acme.PendingTransactionListParamsCategory{
			In: acme.F([]acme.PendingTransactionListParamsCategoryIn{acme.PendingTransactionListParamsCategoryInAccountTransferInstruction, acme.PendingTransactionListParamsCategoryInACHTransferInstruction, acme.PendingTransactionListParamsCategoryInCardAuthorization}),
		}),
		CreatedAt: acme.F(acme.PendingTransactionListParamsCreatedAt{
			After:      acme.F(time.Now()),
			Before:     acme.F(time.Now()),
			OnOrAfter:  acme.F(time.Now()),
			OnOrBefore: acme.F(time.Now()),
		}),
		Cursor:   acme.F("string"),
		Limit:    acme.F(int64(1)),
		RouteID:  acme.F("string"),
		SourceID: acme.F("string"),
		Status: acme.F(acme.PendingTransactionListParamsStatus{
			In: acme.F([]acme.PendingTransactionListParamsStatusIn{acme.PendingTransactionListParamsStatusInPending, acme.PendingTransactionListParamsStatusInComplete}),
		}),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
