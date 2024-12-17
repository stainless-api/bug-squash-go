// File generated from our OpenAPI spec by Stainless.

package acme_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/acme/acme-go"
	"github.com/acme/acme-go/internal/testutil"
	"github.com/acme/acme-go/option"
)

func TestSimulationCardAuthorizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.Cards.Authorize(context.TODO(), acme.SimulationCardAuthorizeParams{
		Amount:               acme.F(int64(1000)),
		CardID:               acme.F("card_oubs0hwk5rn6knuecxg2"),
		DigitalWalletTokenID: acme.F("string"),
		EventSubscriptionID:  acme.F("event_subscription_001dzz0r20rcdxgb013zqb8m04g"),
		PhysicalCardID:       acme.F("string"),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSimulationCardSettlementWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.Cards.Settlement(context.TODO(), acme.SimulationCardSettlementParams{
		CardID:               acme.F("card_oubs0hwk5rn6knuecxg2"),
		PendingTransactionID: acme.F("pending_transaction_k1sfetcau2qbvjbzgju4"),
		Amount:               acme.F(int64(1)),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
