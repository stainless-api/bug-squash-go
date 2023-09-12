// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestCancel(t *testing.T) {
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	res, err := client.Accounts.New(cancelCtx, increase.AccountNewParams{
		Name: increase.F("My First Increase Account"),
	})
	if err == nil && res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
}

func TestCancelDelay(t *testing.T) {
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Millisecond * time.Duration(2))
		cancel()
	}()
	res, err := client.Accounts.New(cancelCtx, increase.AccountNewParams{
		Name: increase.F("My First Increase Account"),
	})
	if err == nil && res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
}
