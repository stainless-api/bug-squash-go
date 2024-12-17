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

func TestSimulationWireTransferNewInboundWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.WireTransfers.NewInbound(context.TODO(), acme.SimulationWireTransferNewInboundParams{
		AccountNumberID:                         acme.F("account_number_v18nkfqm6afpsrvy82b2"),
		Amount:                                  acme.F(int64(1000)),
		BeneficiaryAddressLine1:                 acme.F("x"),
		BeneficiaryAddressLine2:                 acme.F("x"),
		BeneficiaryAddressLine3:                 acme.F("x"),
		BeneficiaryName:                         acme.F("x"),
		BeneficiaryReference:                    acme.F("x"),
		OriginatorAddressLine1:                  acme.F("x"),
		OriginatorAddressLine2:                  acme.F("x"),
		OriginatorAddressLine3:                  acme.F("x"),
		OriginatorName:                          acme.F("x"),
		OriginatorRoutingNumber:                 acme.F("x"),
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
