// The Prism mock server doesn't support a few things, like real pagination
// and intermittent errors. We have a few tests here to better exercise
// those code paths in the client.

package acme_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/acme/acme-go"
	"github.com/acme/acme-go/internal/testutil"
	"github.com/acme/acme-go/option"
)

// Test that we can succesfully paginate through a list endpoint.
// Note that the integration server is known to have at least 5 items to list,
// so we'll limit our requests to 2 per page to force getting at least the
// cases of: a first page, a middle page (i.e. not first or last), and a last
// page.
func TestAutoPaginationIntegration(t *testing.T) {
	baseURL := "http://localhost:8077"
	apiKey := "sk_test_1234567890"
	if !testutil.CheckIntegrationServer(t, baseURL) {
		return
	}
	client := acme.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey(apiKey),
	)
	cds := make(map[string]bool)

	iter := client.CheckDeposits.ListAutoPaging(context.TODO(), acme.CheckDepositListParams{
		Limit: acme.F(int64(2)),
	})
	for iter.Next() {
		cd := iter.Current()
		cds[cd.ID] = true
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}

	if len(cds) < 5 {
		t.Errorf("Check Deposits listed = %d, want >= 5", len(cds))
	}
}

// Test that we're sending idempotency keys correctly. We want to make sure
// that if we retry a request we're sending the same idempotency key on each retry
// and, thus, that the server doesn't try to process the same request twice and
// cause unexpected, duplicative behavior.
// Note that our integration server is known to apply the first instance of
// the create call to its internal DB on the first request, but then to respond
// with a temporary error, so the client will be forced to retry, at which point
// the request can succeed.
func TestNewIdempotentlyIntegration(t *testing.T) {
	baseURL := "http://localhost:8077"
	apiKey := "sk_test_1234567890"
	if !testutil.CheckIntegrationServer(t, baseURL) {
		return
	}
	client := acme.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey(apiKey),
	)
	accountID := fmt.Sprintf("account_%08d", rand.Int())
	// part 1: create a check deposit. This method should trigger the idempotency behavior.
	// we can't test that directly, so see part 2 below
	_, err := client.CheckDeposits.New(context.TODO(), acme.CheckDepositNewParams{
		AccountID:        acme.F(accountID),
		Amount:           acme.F(int64(42)),
		Currency:         acme.F("USD"),
		FrontImageFileID: acme.F(fmt.Sprintf("file_%08d", rand.Int())),
		BackImageFileID:  acme.F(fmt.Sprintf("file_%08d", rand.Int())),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}

	// part 2: to check if we've done part 1 correctly, we'll count the number
	// of check deposits that get created. we should only see 1!
	iter := client.CheckDeposits.ListAutoPaging(context.TODO(), acme.CheckDepositListParams{})
	checkDepositsInAccount := 0
	for iter.Next() {
		// use this to filter down to only the check deposits on the account for this test
		if accountID == iter.Current().AccountID {
			checkDepositsInAccount += 1
		}
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if checkDepositsInAccount != 1 {
		t.Errorf("Saw %d deposits in account, but we want exactly 1 time", checkDepositsInAccount)
	}
}
