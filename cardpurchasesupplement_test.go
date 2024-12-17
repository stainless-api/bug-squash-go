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

func TestCardPurchaseSupplementGet(t *testing.T) {
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
	_, err := client.CardPurchaseSupplements.Get(context.TODO(), "card_purchase_supplement_ijuc45iym4jchnh2sfk3")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardPurchaseSupplementListWithOptionalParams(t *testing.T) {
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
	_, err := client.CardPurchaseSupplements.List(context.TODO(), acme.CardPurchaseSupplementListParams{
		CardPaymentID: acme.F("string"),
		CreatedAt: acme.F(acme.CardPurchaseSupplementListParamsCreatedAt{
			After:      acme.F(time.Now()),
			Before:     acme.F(time.Now()),
			OnOrAfter:  acme.F(time.Now()),
			OnOrBefore: acme.F(time.Now()),
		}),
		Cursor: acme.F("string"),
		Limit:  acme.F(int64(1)),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
