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

func TestSimulationInboundWireDrawdownRequestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.InboundWireDrawdownRequests.New(context.TODO(), acme.SimulationInboundWireDrawdownRequestNewParams{
		Amount:                                  acme.F(int64(10000)),
		BeneficiaryAccountNumber:                acme.F("987654321"),
		BeneficiaryRoutingNumber:                acme.F("101050001"),
		Currency:                                acme.F("USD"),
		MessageToRecipient:                      acme.F("Invoice 29582"),
		OriginatorAccountNumber:                 acme.F("987654321"),
		OriginatorRoutingNumber:                 acme.F("101050001"),
		RecipientAccountNumberID:                acme.F("account_number_v18nkfqm6afpsrvy82b2"),
		BeneficiaryAddressLine1:                 acme.F("33 Liberty Street"),
		BeneficiaryAddressLine2:                 acme.F("New York, NY, 10045"),
		BeneficiaryAddressLine3:                 acme.F("x"),
		BeneficiaryName:                         acme.F("Ian Crease"),
		OriginatorAddressLine1:                  acme.F("33 Liberty Street"),
		OriginatorAddressLine2:                  acme.F("New York, NY, 10045"),
		OriginatorAddressLine3:                  acme.F("x"),
		OriginatorName:                          acme.F("Ian Crease"),
		OriginatorToBeneficiaryInformationLine1: acme.F("x"),
		OriginatorToBeneficiaryInformationLine2: acme.F("x"),
		OriginatorToBeneficiaryInformationLine3: acme.F("x"),
		OriginatorToBeneficiaryInformationLine4: acme.F("x"),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
