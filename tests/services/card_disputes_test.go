package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestCardDisputesActionWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.CardDisputes.Action(
		context.TODO(),
		"card_dispute_h9sc95nbl1cgltpp7men",
		&types.SimulatesAdvancingTheStateOfACardDisputeParameters{Status: increase.P(types.SimulatesAdvancingTheStateOfACardDisputeParametersStatusAccepted), Explanation: increase.P("This was a valid recurring transaction")},
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
