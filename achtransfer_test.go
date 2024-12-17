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

func TestACHTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ACHTransfers.New(context.TODO(), acme.ACHTransferNewParams{
		AccountID:                acme.F("account_in71c4amph0vgo2qllky"),
		Amount:                   acme.F(int64(100)),
		StatementDescriptor:      acme.F("New ACH transfer"),
		AccountNumber:            acme.F("987654321"),
		Addendum:                 acme.F("x"),
		CompanyDescriptiveDate:   acme.F("x"),
		CompanyDiscretionaryData: acme.F("x"),
		CompanyEntryDescription:  acme.F("x"),
		CompanyName:              acme.F("x"),
		EffectiveDate:            acme.F(time.Now()),
		ExternalAccountID:        acme.F("string"),
		Funding:                  acme.F(acme.ACHTransferNewParamsFundingChecking),
		IndividualID:             acme.F("x"),
		IndividualName:           acme.F("x"),
		RequireApproval:          acme.F(true),
		RoutingNumber:            acme.F("101050001"),
		StandardEntryClassCode:   acme.F(acme.ACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit),
		UniqueIdentifier:         acme.F("x"),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferGet(t *testing.T) {
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
	_, err := client.ACHTransfers.Get(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.ACHTransfers.List(context.TODO(), acme.ACHTransferListParams{
		AccountID: acme.F("string"),
		CreatedAt: acme.F(acme.ACHTransferListParamsCreatedAt{
			After:      acme.F(time.Now()),
			Before:     acme.F(time.Now()),
			OnOrAfter:  acme.F(time.Now()),
			OnOrBefore: acme.F(time.Now()),
		}),
		Cursor:            acme.F("string"),
		ExternalAccountID: acme.F("string"),
		Limit:             acme.F(int64(1)),
		UniqueIdentifier:  acme.F("x"),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferApprove(t *testing.T) {
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
	_, err := client.ACHTransfers.Approve(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferCancel(t *testing.T) {
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
	_, err := client.ACHTransfers.Cancel(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
