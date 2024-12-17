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

func TestExportNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Exports.New(context.TODO(), acme.ExportNewParams{
		Category: acme.F(acme.ExportNewParamsCategoryTransactionCsv),
		AccountStatementOfx: acme.F(acme.ExportNewParamsAccountStatementOfx{
			AccountID: acme.F("string"),
			CreatedAt: acme.F(acme.ExportNewParamsAccountStatementOfxCreatedAt{
				After:      acme.F(time.Now()),
				Before:     acme.F(time.Now()),
				OnOrAfter:  acme.F(time.Now()),
				OnOrBefore: acme.F(time.Now()),
			}),
		}),
		BalanceCsv: acme.F(acme.ExportNewParamsBalanceCsv{
			AccountID: acme.F("string"),
			CreatedAt: acme.F(acme.ExportNewParamsBalanceCsvCreatedAt{
				After:      acme.F(time.Now()),
				Before:     acme.F(time.Now()),
				OnOrAfter:  acme.F(time.Now()),
				OnOrBefore: acme.F(time.Now()),
			}),
		}),
		BookkeepingAccountBalanceCsv: acme.F(acme.ExportNewParamsBookkeepingAccountBalanceCsv{
			BookkeepingAccountID: acme.F("string"),
			CreatedAt: acme.F(acme.ExportNewParamsBookkeepingAccountBalanceCsvCreatedAt{
				After:      acme.F(time.Now()),
				Before:     acme.F(time.Now()),
				OnOrAfter:  acme.F(time.Now()),
				OnOrBefore: acme.F(time.Now()),
			}),
		}),
		EntityCsv: acme.F(acme.ExportNewParamsEntityCsv{
			Status: acme.F(acme.ExportNewParamsEntityCsvStatus{
				In: acme.F([]acme.ExportNewParamsEntityCsvStatusIn{acme.ExportNewParamsEntityCsvStatusInActive, acme.ExportNewParamsEntityCsvStatusInArchived, acme.ExportNewParamsEntityCsvStatusInDisabled}),
			}),
		}),
		TransactionCsv: acme.F(acme.ExportNewParamsTransactionCsv{
			AccountID: acme.F("account_in71c4amph0vgo2qllky"),
			CreatedAt: acme.F(acme.ExportNewParamsTransactionCsvCreatedAt{
				After:      acme.F(time.Now()),
				Before:     acme.F(time.Now()),
				OnOrAfter:  acme.F(time.Now()),
				OnOrBefore: acme.F(time.Now()),
			}),
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

func TestExportGet(t *testing.T) {
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
	_, err := client.Exports.Get(context.TODO(), "export_8s4m48qz3bclzje0zwh9")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExportListWithOptionalParams(t *testing.T) {
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
	_, err := client.Exports.List(context.TODO(), acme.ExportListParams{
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
