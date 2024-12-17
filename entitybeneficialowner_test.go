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

func TestEntityBeneficialOwnerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.BeneficialOwners.New(context.TODO(), acme.EntityBeneficialOwnerNewParams{
		BeneficialOwner: acme.F(acme.EntityBeneficialOwnerNewParamsBeneficialOwner{
			Individual: acme.F(acme.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividual{
				Name:        acme.F("Ian Crease"),
				DateOfBirth: acme.F(time.Now()),
				Address: acme.F(acme.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualAddress{
					Line1: acme.F("33 Liberty Street"),
					Line2: acme.F("x"),
					City:  acme.F("New York"),
					State: acme.F("NY"),
					Zip:   acme.F("10045"),
				}),
				ConfirmedNoUsTaxID: acme.F(true),
				Identification: acme.F(acme.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentification{
					Method: acme.F(acme.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber),
					Number: acme.F("078051120"),
					Passport: acme.F(acme.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationPassport{
						FileID:         acme.F("string"),
						ExpirationDate: acme.F(time.Now()),
						Country:        acme.F("x"),
					}),
					DriversLicense: acme.F(acme.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationDriversLicense{
						FileID:         acme.F("string"),
						BackFileID:     acme.F("string"),
						ExpirationDate: acme.F(time.Now()),
						State:          acme.F("x"),
					}),
					Other: acme.F(acme.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationOther{
						Country:        acme.F("x"),
						Description:    acme.F("x"),
						ExpirationDate: acme.F(time.Now()),
						FileID:         acme.F("string"),
						BackFileID:     acme.F("string"),
					}),
				}),
			}),
			CompanyTitle: acme.F("CEO"),
			Prongs:       acme.F([]acme.EntityBeneficialOwnerNewParamsBeneficialOwnerProng{acme.EntityBeneficialOwnerNewParamsBeneficialOwnerProngControl}),
		}),
		EntityID: acme.F("entity_n8y8tnk2p9339ti393yi"),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityBeneficialOwnerArchive(t *testing.T) {
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
	_, err := client.Entities.BeneficialOwners.Archive(context.TODO(), acme.EntityBeneficialOwnerArchiveParams{
		BeneficialOwnerID: acme.F("entity_setup_beneficial_owner_submission_vgkyk7dj5eb4sfhdbkx7"),
		EntityID:          acme.F("entity_n8y8tnk2p9339ti393yi"),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityBeneficialOwnerUpdateAddressWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.BeneficialOwners.UpdateAddress(context.TODO(), acme.EntityBeneficialOwnerUpdateAddressParams{
		Address: acme.F(acme.EntityBeneficialOwnerUpdateAddressParamsAddress{
			Line1: acme.F("33 Liberty Street"),
			Line2: acme.F("Unit 2"),
			City:  acme.F("New York"),
			State: acme.F("NY"),
			Zip:   acme.F("10045"),
		}),
		BeneficialOwnerID: acme.F("entity_setup_beneficial_owner_submission_vgkyk7dj5eb4sfhdbkx7"),
		EntityID:          acme.F("entity_n8y8tnk2p9339ti393yi"),
	})
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
