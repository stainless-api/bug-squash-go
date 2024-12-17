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

func TestCardProfileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CardProfiles.New(context.TODO(), acme.CardProfileNewParams{
		Description: acme.F("My Card Profile"),
		DigitalWallets: acme.F(acme.CardProfileNewParamsDigitalWallets{
			TextColor: acme.F(acme.CardProfileNewParamsDigitalWalletsTextColor{
				Red:   acme.F(int64(26)),
				Green: acme.F(int64(43)),
				Blue:  acme.F(int64(59)),
			}),
			IssuerName:            acme.F("MyBank"),
			CardDescription:       acme.F("MyBank Signature Card"),
			ContactWebsite:        acme.F("https://example.com"),
			ContactEmail:          acme.F("user@example.com"),
			ContactPhone:          acme.F("+18885551212"),
			BackgroundImageFileID: acme.F("file_1ai913suu1zfn1pdetru"),
			AppIconFileID:         acme.F("file_8zxqkwlh43wo144u8yec"),
		}),
		PhysicalCards: acme.F(acme.CardProfileNewParamsPhysicalCards{
			ContactPhone:       acme.F("x"),
			FrontImageFileID:   acme.F("string"),
			CarrierImageFileID: acme.F("string"),
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

func TestCardProfileGet(t *testing.T) {
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
	_, err := client.CardProfiles.Get(context.TODO(), "card_profile_cox5y73lob2eqly18piy")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardProfileListWithOptionalParams(t *testing.T) {
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
	_, err := client.CardProfiles.List(context.TODO(), acme.CardProfileListParams{
		Cursor: acme.F("string"),
		Limit:  acme.F(int64(1)),
		PhysicalCardsStatus: acme.F(acme.CardProfileListParamsPhysicalCardsStatus{
			In: acme.F([]acme.CardProfileListParamsPhysicalCardsStatusIn{acme.CardProfileListParamsPhysicalCardsStatusInNotEligible, acme.CardProfileListParamsPhysicalCardsStatusInRejected, acme.CardProfileListParamsPhysicalCardsStatusInPendingCreating}),
		}),
		Status: acme.F(acme.CardProfileListParamsStatus{
			In: acme.F([]acme.CardProfileListParamsStatusIn{acme.CardProfileListParamsStatusInPending, acme.CardProfileListParamsStatusInRejected, acme.CardProfileListParamsStatusInActive}),
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

func TestCardProfileArchive(t *testing.T) {
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
	_, err := client.CardProfiles.Archive(context.TODO(), "card_profile_cox5y73lob2eqly18piy")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
