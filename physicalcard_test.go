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

func TestPhysicalCardNew(t *testing.T) {
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
	_, err := client.PhysicalCards.New(context.TODO(), acme.PhysicalCardNewParams{
		CardID:        acme.F("card_oubs0hwk5rn6knuecxg2"),
		CardProfileID: acme.F("card_profile_cox5y73lob2eqly18piy"),
		Cardholder: acme.F(acme.PhysicalCardNewParamsCardholder{
			FirstName: acme.F("Ian"),
			LastName:  acme.F("Crease"),
		}),
		Shipment: acme.F(acme.PhysicalCardNewParamsShipment{
			Method: acme.F(acme.PhysicalCardNewParamsShipmentMethodUsps),
			Address: acme.F(acme.PhysicalCardNewParamsShipmentAddress{
				Name:        acme.F("Ian Crease"),
				Line1:       acme.F("33 Liberty Street"),
				Line2:       acme.F("Unit 2"),
				Line3:       acme.F("x"),
				PhoneNumber: acme.F("x"),
				City:        acme.F("New York"),
				State:       acme.F("NY"),
				PostalCode:  acme.F("10045"),
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

func TestPhysicalCardGet(t *testing.T) {
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
	_, err := client.PhysicalCards.Get(context.TODO(), "physical_card_ode8duyq5v2ynhjoharl")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCardUpdate(t *testing.T) {
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
	_, err := client.PhysicalCards.Update(
		context.TODO(),
		"physical_card_ode8duyq5v2ynhjoharl",
		acme.PhysicalCardUpdateParams{
			Status: acme.F(acme.PhysicalCardUpdateParamsStatusDisabled),
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

func TestPhysicalCardListWithOptionalParams(t *testing.T) {
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
	_, err := client.PhysicalCards.List(context.TODO(), acme.PhysicalCardListParams{
		CardID: acme.F("string"),
		CreatedAt: acme.F(acme.PhysicalCardListParamsCreatedAt{
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
