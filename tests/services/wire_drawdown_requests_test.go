package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestWireDrawdownRequestsNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are broken")
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.WireDrawdownRequests.New(context.TODO(), &types.CreateAWireDrawdownRequestParameters{AccountNumberID: increase.P("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.P(10000), MessageToRecipient: increase.P("Invoice 29582"), RecipientAccountNumber: increase.P("987654321"), RecipientRoutingNumber: increase.P("101050001"), RecipientName: increase.P("Ian Crease")})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestWireDrawdownRequestsGet(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.WireDrawdownRequests.Get(context.TODO(), "wire_drawdown_request_q6lmocus3glo0lr2bfv3")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestWireDrawdownRequestsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.WireDrawdownRequests.List(context.TODO(), &types.WireDrawdownRequestListParams{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}