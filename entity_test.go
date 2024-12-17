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

func TestEntityNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.New(context.TODO(), acme.EntityNewParams{
		Structure: acme.F(acme.EntityNewParamsStructureCorporation),
		Corporation: acme.F(acme.EntityNewParamsCorporation{
			Name:               acme.F("National Phonograph Company"),
			Website:            acme.F("https://example.com"),
			TaxIdentifier:      acme.F("602214076"),
			IncorporationState: acme.F("NY"),
			Address: acme.F(acme.EntityNewParamsCorporationAddress{
				Line1: acme.F("33 Liberty Street"),
				Line2: acme.F("x"),
				City:  acme.F("New York"),
				State: acme.F("NY"),
				Zip:   acme.F("10045"),
			}),
			BeneficialOwners: acme.F([]acme.EntityNewParamsCorporationBeneficialOwner{{
				Individual: acme.F(acme.EntityNewParamsCorporationBeneficialOwnersIndividual{
					Name:        acme.F("Ian Crease"),
					DateOfBirth: acme.F(time.Now()),
					Address: acme.F(acme.EntityNewParamsCorporationBeneficialOwnersIndividualAddress{
						Line1: acme.F("33 Liberty Street"),
						Line2: acme.F("x"),
						City:  acme.F("New York"),
						State: acme.F("NY"),
						Zip:   acme.F("10045"),
					}),
					ConfirmedNoUsTaxID: acme.F(true),
					Identification: acme.F(acme.EntityNewParamsCorporationBeneficialOwnersIndividualIdentification{
						Method: acme.F(acme.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber),
						Number: acme.F("078051120"),
						Passport: acme.F(acme.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport{
							FileID:         acme.F("string"),
							ExpirationDate: acme.F(time.Now()),
							Country:        acme.F("x"),
						}),
						DriversLicense: acme.F(acme.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense{
							FileID:         acme.F("string"),
							BackFileID:     acme.F("string"),
							ExpirationDate: acme.F(time.Now()),
							State:          acme.F("x"),
						}),
						Other: acme.F(acme.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther{
							Country:        acme.F("x"),
							Description:    acme.F("x"),
							ExpirationDate: acme.F(time.Now()),
							FileID:         acme.F("string"),
							BackFileID:     acme.F("string"),
						}),
					}),
				}),
				CompanyTitle: acme.F("CEO"),
				Prongs:       acme.F([]acme.EntityNewParamsCorporationBeneficialOwnersProng{acme.EntityNewParamsCorporationBeneficialOwnersProngControl}),
			}}),
		}),
		Description: acme.F("x"),
		Joint: acme.F(acme.EntityNewParamsJoint{
			Name: acme.F("x"),
			Individuals: acme.F([]acme.EntityNewParamsJointIndividual{{
				Name:        acme.F("x"),
				DateOfBirth: acme.F(time.Now()),
				Address: acme.F(acme.EntityNewParamsJointIndividualsAddress{
					Line1: acme.F("x"),
					Line2: acme.F("x"),
					City:  acme.F("x"),
					State: acme.F("x"),
					Zip:   acme.F("x"),
				}),
				ConfirmedNoUsTaxID: acme.F(true),
				Identification: acme.F(acme.EntityNewParamsJointIndividualsIdentification{
					Method: acme.F(acme.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber),
					Number: acme.F("xxxx"),
					Passport: acme.F(acme.EntityNewParamsJointIndividualsIdentificationPassport{
						FileID:         acme.F("string"),
						ExpirationDate: acme.F(time.Now()),
						Country:        acme.F("x"),
					}),
					DriversLicense: acme.F(acme.EntityNewParamsJointIndividualsIdentificationDriversLicense{
						FileID:         acme.F("string"),
						BackFileID:     acme.F("string"),
						ExpirationDate: acme.F(time.Now()),
						State:          acme.F("x"),
					}),
					Other: acme.F(acme.EntityNewParamsJointIndividualsIdentificationOther{
						Country:        acme.F("x"),
						Description:    acme.F("x"),
						ExpirationDate: acme.F(time.Now()),
						FileID:         acme.F("string"),
						BackFileID:     acme.F("string"),
					}),
				}),
			}, {
				Name:        acme.F("x"),
				DateOfBirth: acme.F(time.Now()),
				Address: acme.F(acme.EntityNewParamsJointIndividualsAddress{
					Line1: acme.F("x"),
					Line2: acme.F("x"),
					City:  acme.F("x"),
					State: acme.F("x"),
					Zip:   acme.F("x"),
				}),
				ConfirmedNoUsTaxID: acme.F(true),
				Identification: acme.F(acme.EntityNewParamsJointIndividualsIdentification{
					Method: acme.F(acme.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber),
					Number: acme.F("xxxx"),
					Passport: acme.F(acme.EntityNewParamsJointIndividualsIdentificationPassport{
						FileID:         acme.F("string"),
						ExpirationDate: acme.F(time.Now()),
						Country:        acme.F("x"),
					}),
					DriversLicense: acme.F(acme.EntityNewParamsJointIndividualsIdentificationDriversLicense{
						FileID:         acme.F("string"),
						BackFileID:     acme.F("string"),
						ExpirationDate: acme.F(time.Now()),
						State:          acme.F("x"),
					}),
					Other: acme.F(acme.EntityNewParamsJointIndividualsIdentificationOther{
						Country:        acme.F("x"),
						Description:    acme.F("x"),
						ExpirationDate: acme.F(time.Now()),
						FileID:         acme.F("string"),
						BackFileID:     acme.F("string"),
					}),
				}),
			}, {
				Name:        acme.F("x"),
				DateOfBirth: acme.F(time.Now()),
				Address: acme.F(acme.EntityNewParamsJointIndividualsAddress{
					Line1: acme.F("x"),
					Line2: acme.F("x"),
					City:  acme.F("x"),
					State: acme.F("x"),
					Zip:   acme.F("x"),
				}),
				ConfirmedNoUsTaxID: acme.F(true),
				Identification: acme.F(acme.EntityNewParamsJointIndividualsIdentification{
					Method: acme.F(acme.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber),
					Number: acme.F("xxxx"),
					Passport: acme.F(acme.EntityNewParamsJointIndividualsIdentificationPassport{
						FileID:         acme.F("string"),
						ExpirationDate: acme.F(time.Now()),
						Country:        acme.F("x"),
					}),
					DriversLicense: acme.F(acme.EntityNewParamsJointIndividualsIdentificationDriversLicense{
						FileID:         acme.F("string"),
						BackFileID:     acme.F("string"),
						ExpirationDate: acme.F(time.Now()),
						State:          acme.F("x"),
					}),
					Other: acme.F(acme.EntityNewParamsJointIndividualsIdentificationOther{
						Country:        acme.F("x"),
						Description:    acme.F("x"),
						ExpirationDate: acme.F(time.Now()),
						FileID:         acme.F("string"),
						BackFileID:     acme.F("string"),
					}),
				}),
			}}),
		}),
		NaturalPerson: acme.F(acme.EntityNewParamsNaturalPerson{
			Name:        acme.F("x"),
			DateOfBirth: acme.F(time.Now()),
			Address: acme.F(acme.EntityNewParamsNaturalPersonAddress{
				Line1: acme.F("x"),
				Line2: acme.F("x"),
				City:  acme.F("x"),
				State: acme.F("x"),
				Zip:   acme.F("x"),
			}),
			ConfirmedNoUsTaxID: acme.F(true),
			Identification: acme.F(acme.EntityNewParamsNaturalPersonIdentification{
				Method: acme.F(acme.EntityNewParamsNaturalPersonIdentificationMethodSocialSecurityNumber),
				Number: acme.F("xxxx"),
				Passport: acme.F(acme.EntityNewParamsNaturalPersonIdentificationPassport{
					FileID:         acme.F("string"),
					ExpirationDate: acme.F(time.Now()),
					Country:        acme.F("x"),
				}),
				DriversLicense: acme.F(acme.EntityNewParamsNaturalPersonIdentificationDriversLicense{
					FileID:         acme.F("string"),
					BackFileID:     acme.F("string"),
					ExpirationDate: acme.F(time.Now()),
					State:          acme.F("x"),
				}),
				Other: acme.F(acme.EntityNewParamsNaturalPersonIdentificationOther{
					Country:        acme.F("x"),
					Description:    acme.F("x"),
					ExpirationDate: acme.F(time.Now()),
					FileID:         acme.F("string"),
					BackFileID:     acme.F("string"),
				}),
			}),
		}),
		Relationship: acme.F(acme.EntityNewParamsRelationshipAffiliated),
		SupplementalDocuments: acme.F([]acme.EntityNewParamsSupplementalDocument{{
			FileID: acme.F("file_makxrc67oh9l6sg7w9yc"),
		}}),
		Trust: acme.F(acme.EntityNewParamsTrust{
			Name:           acme.F("x"),
			Category:       acme.F(acme.EntityNewParamsTrustCategoryRevocable),
			TaxIdentifier:  acme.F("x"),
			FormationState: acme.F("x"),
			Address: acme.F(acme.EntityNewParamsTrustAddress{
				Line1: acme.F("x"),
				Line2: acme.F("x"),
				City:  acme.F("x"),
				State: acme.F("x"),
				Zip:   acme.F("x"),
			}),
			FormationDocumentFileID: acme.F("string"),
			Trustees: acme.F([]acme.EntityNewParamsTrustTrustee{{
				Structure: acme.F(acme.EntityNewParamsTrustTrusteesStructureIndividual),
				Individual: acme.F(acme.EntityNewParamsTrustTrusteesIndividual{
					Name:        acme.F("x"),
					DateOfBirth: acme.F(time.Now()),
					Address: acme.F(acme.EntityNewParamsTrustTrusteesIndividualAddress{
						Line1: acme.F("x"),
						Line2: acme.F("x"),
						City:  acme.F("x"),
						State: acme.F("x"),
						Zip:   acme.F("x"),
					}),
					ConfirmedNoUsTaxID: acme.F(true),
					Identification: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentification{
						Method: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber),
						Number: acme.F("xxxx"),
						Passport: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{
							FileID:         acme.F("string"),
							ExpirationDate: acme.F(time.Now()),
							Country:        acme.F("x"),
						}),
						DriversLicense: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{
							FileID:         acme.F("string"),
							BackFileID:     acme.F("string"),
							ExpirationDate: acme.F(time.Now()),
							State:          acme.F("x"),
						}),
						Other: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							Country:        acme.F("x"),
							Description:    acme.F("x"),
							ExpirationDate: acme.F(time.Now()),
							FileID:         acme.F("string"),
							BackFileID:     acme.F("string"),
						}),
					}),
				}),
			}, {
				Structure: acme.F(acme.EntityNewParamsTrustTrusteesStructureIndividual),
				Individual: acme.F(acme.EntityNewParamsTrustTrusteesIndividual{
					Name:        acme.F("x"),
					DateOfBirth: acme.F(time.Now()),
					Address: acme.F(acme.EntityNewParamsTrustTrusteesIndividualAddress{
						Line1: acme.F("x"),
						Line2: acme.F("x"),
						City:  acme.F("x"),
						State: acme.F("x"),
						Zip:   acme.F("x"),
					}),
					ConfirmedNoUsTaxID: acme.F(true),
					Identification: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentification{
						Method: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber),
						Number: acme.F("xxxx"),
						Passport: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{
							FileID:         acme.F("string"),
							ExpirationDate: acme.F(time.Now()),
							Country:        acme.F("x"),
						}),
						DriversLicense: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{
							FileID:         acme.F("string"),
							BackFileID:     acme.F("string"),
							ExpirationDate: acme.F(time.Now()),
							State:          acme.F("x"),
						}),
						Other: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							Country:        acme.F("x"),
							Description:    acme.F("x"),
							ExpirationDate: acme.F(time.Now()),
							FileID:         acme.F("string"),
							BackFileID:     acme.F("string"),
						}),
					}),
				}),
			}, {
				Structure: acme.F(acme.EntityNewParamsTrustTrusteesStructureIndividual),
				Individual: acme.F(acme.EntityNewParamsTrustTrusteesIndividual{
					Name:        acme.F("x"),
					DateOfBirth: acme.F(time.Now()),
					Address: acme.F(acme.EntityNewParamsTrustTrusteesIndividualAddress{
						Line1: acme.F("x"),
						Line2: acme.F("x"),
						City:  acme.F("x"),
						State: acme.F("x"),
						Zip:   acme.F("x"),
					}),
					ConfirmedNoUsTaxID: acme.F(true),
					Identification: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentification{
						Method: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber),
						Number: acme.F("xxxx"),
						Passport: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{
							FileID:         acme.F("string"),
							ExpirationDate: acme.F(time.Now()),
							Country:        acme.F("x"),
						}),
						DriversLicense: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{
							FileID:         acme.F("string"),
							BackFileID:     acme.F("string"),
							ExpirationDate: acme.F(time.Now()),
							State:          acme.F("x"),
						}),
						Other: acme.F(acme.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							Country:        acme.F("x"),
							Description:    acme.F("x"),
							ExpirationDate: acme.F(time.Now()),
							FileID:         acme.F("string"),
							BackFileID:     acme.F("string"),
						}),
					}),
				}),
			}}),
			Grantor: acme.F(acme.EntityNewParamsTrustGrantor{
				Name:        acme.F("x"),
				DateOfBirth: acme.F(time.Now()),
				Address: acme.F(acme.EntityNewParamsTrustGrantorAddress{
					Line1: acme.F("x"),
					Line2: acme.F("x"),
					City:  acme.F("x"),
					State: acme.F("x"),
					Zip:   acme.F("x"),
				}),
				ConfirmedNoUsTaxID: acme.F(true),
				Identification: acme.F(acme.EntityNewParamsTrustGrantorIdentification{
					Method: acme.F(acme.EntityNewParamsTrustGrantorIdentificationMethodSocialSecurityNumber),
					Number: acme.F("xxxx"),
					Passport: acme.F(acme.EntityNewParamsTrustGrantorIdentificationPassport{
						FileID:         acme.F("string"),
						ExpirationDate: acme.F(time.Now()),
						Country:        acme.F("x"),
					}),
					DriversLicense: acme.F(acme.EntityNewParamsTrustGrantorIdentificationDriversLicense{
						FileID:         acme.F("string"),
						BackFileID:     acme.F("string"),
						ExpirationDate: acme.F(time.Now()),
						State:          acme.F("x"),
					}),
					Other: acme.F(acme.EntityNewParamsTrustGrantorIdentificationOther{
						Country:        acme.F("x"),
						Description:    acme.F("x"),
						ExpirationDate: acme.F(time.Now()),
						FileID:         acme.F("string"),
						BackFileID:     acme.F("string"),
					}),
				}),
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

func TestEntityGet(t *testing.T) {
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
	_, err := client.Entities.Get(context.TODO(), "entity_n8y8tnk2p9339ti393yi")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityListWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.List(context.TODO(), acme.EntityListParams{
		CreatedAt: acme.F(acme.EntityListParamsCreatedAt{
			After:      acme.F(time.Now()),
			Before:     acme.F(time.Now()),
			OnOrAfter:  acme.F(time.Now()),
			OnOrBefore: acme.F(time.Now()),
		}),
		Cursor: acme.F("string"),
		Limit:  acme.F(int64(1)),
		Status: acme.F(acme.EntityListParamsStatus{
			In: acme.F([]acme.EntityListParamsStatusIn{acme.EntityListParamsStatusInActive, acme.EntityListParamsStatusInArchived, acme.EntityListParamsStatusInDisabled}),
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

func TestEntityArchive(t *testing.T) {
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
	_, err := client.Entities.Archive(context.TODO(), "entity_n8y8tnk2p9339ti393yi")
	if err != nil {
		var apierr *acme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityUpdateAddressWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.UpdateAddress(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		acme.EntityUpdateAddressParams{
			Address: acme.F(acme.EntityUpdateAddressParamsAddress{
				Line1: acme.F("33 Liberty Street"),
				Line2: acme.F("Unit 2"),
				City:  acme.F("New York"),
				State: acme.F("NY"),
				Zip:   acme.F("10045"),
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
