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

func TestInboundACHTransferGet(t *testing.T) {
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
	_, err := client.InboundACHTransfers.Get(context.TODO(), "inbound_ach_transfer_tdrwqr3fq9gnnq49odev")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboundACHTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.InboundACHTransfers.List(context.TODO(), acme.InboundACHTransferListParams{
		AccountID: acme.F("string"),
		CreatedAt: acme.F(acme.InboundACHTransferListParamsCreatedAt{
			After:      acme.F(time.Now()),
			Before:     acme.F(time.Now()),
			OnOrAfter:  acme.F(time.Now()),
			OnOrBefore: acme.F(time.Now()),
		}),
		Cursor: acme.F("string"),
		Limit:  acme.F(int64(1)),
		Status: acme.F(acme.InboundACHTransferListParamsStatusPending),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboundACHTransferDecline(t *testing.T) {
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
	_, err := client.InboundACHTransfers.Decline(context.TODO(), "inbound_ach_transfer_tdrwqr3fq9gnnq49odev")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboundACHTransferNotificationOfChangeWithOptionalParams(t *testing.T) {
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
	_, err := client.InboundACHTransfers.NotificationOfChange(
		context.TODO(),
		"inbound_ach_transfer_tdrwqr3fq9gnnq49odev",
		acme.InboundACHTransferNotificationOfChangeParams{
			UpdatedAccountNumber: acme.F("987654321"),
			UpdatedRoutingNumber: acme.F("101050001"),
		},
	)
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboundACHTransferTransferReturn(t *testing.T) {
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
	_, err := client.InboundACHTransfers.TransferReturn(
		context.TODO(),
		"inbound_ach_transfer_tdrwqr3fq9gnnq49odev",
		acme.InboundACHTransferTransferReturnParams{
			Reason: acme.F(acme.InboundACHTransferTransferReturnParamsReasonPaymentStopped),
		},
	)
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
