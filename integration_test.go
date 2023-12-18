// The Prism mock server doesn't support a few things, like real pagination
// and intermittent errors. We have a few tests here to better exercise
// those code paths in the client.

package increase_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
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
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey(apiKey),
	)
	cds := make(map[string]bool)

	iter := client.CheckDeposits.ListAutoPaging(context.TODO(), increase.CheckDepositListParams{
		Limit: increase.F(int64(2)),
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

// Test that when making requests against an API endpoint that has side
// effects, and we lose the response from the server after those effects
// happen, that subsequent client retries do not replay those side effects,
// and instead correctly use an idempotency key to ensure the side effect only
// happens once.
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
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey(apiKey),
	)
	accountID := fmt.Sprintf("account_%08d", rand.Int())
	_, err := client.CheckDeposits.New(context.TODO(), increase.CheckDepositNewParams{
		AccountID:        increase.F(accountID),
		Amount:           increase.F(int64(42)),
		Currency:         increase.F("USD"),
		FrontImageFileID: increase.F(fmt.Sprintf("file_%08d", rand.Int())),
		BackImageFileID:  increase.F(fmt.Sprintf("file_%08d", rand.Int())),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}

	iter := client.CheckDeposits.ListAutoPaging(context.TODO(), increase.CheckDepositListParams{})
	accountIDSeen := 0
	for iter.Next() {
		if accountID == iter.Current().AccountID {
			accountIDSeen += 1
		}
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if accountIDSeen != 1 {
		t.Errorf("Unique account ID seen in created records %d times, want exactly 1 time", accountIDSeen)
	}
}
