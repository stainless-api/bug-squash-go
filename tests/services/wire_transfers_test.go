package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestWireTransfersCreateInboundWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.WireTransfers.Simulations.CreateInbound(context.TODO(), &types.SimulateAWireTransferToYourAccountParameters{AccountNumberID: increase.P("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.P(1000)})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
