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

func TestRealTimeDecisionGet(t *testing.T) {
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
	_, err := client.RealTimeDecisions.Get(context.TODO(), "real_time_decision_j76n2e810ezcg3zh5qtn")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimeDecisionActionWithOptionalParams(t *testing.T) {
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
	_, err := client.RealTimeDecisions.Action(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
		acme.RealTimeDecisionActionParams{
			CardAuthorization: acme.F(acme.RealTimeDecisionActionParamsCardAuthorization{
				Decision: acme.F(acme.RealTimeDecisionActionParamsCardAuthorizationDecisionApprove),
			}),
			DigitalWalletAuthentication: acme.F(acme.RealTimeDecisionActionParamsDigitalWalletAuthentication{
				Result: acme.F(acme.RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess),
			}),
			DigitalWalletToken: acme.F(acme.RealTimeDecisionActionParamsDigitalWalletToken{
				Approval: acme.F(acme.RealTimeDecisionActionParamsDigitalWalletTokenApproval{
					CardProfileID: acme.F("string"),
					Phone:         acme.F("x"),
					Email:         acme.F("x"),
				}),
				Decline: acme.F(acme.RealTimeDecisionActionParamsDigitalWalletTokenDecline{
					Reason: acme.F("x"),
				}),
			}),
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
