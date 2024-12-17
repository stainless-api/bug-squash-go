// File generated from our OpenAPI spec by Stainless.

package acme_test

import (
	"context"
	"os"
	"testing"

	"github.com/acme/acme-go"
	"github.com/acme/acme-go/internal/testutil"
	"github.com/acme/acme-go/option"
)

func TestUsage(t *testing.T) {
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
	account, err := client.Accounts.New(context.TODO(), acme.AccountNewParams{
		Name: acme.F("My First Acme Account"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", account.ID)
}
