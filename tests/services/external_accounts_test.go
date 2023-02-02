package services

import (
	"context"
	"increase"
	"testing"

	client "increase"
	"increase/types"
)

func TestExternalAccountsNewWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.ExternalAccounts.New(context.TODO(), &types.CreateAnExternalAccountParameters{RoutingNumber: increase.P("101050001"), AccountNumber: increase.P("987654321"), Description: increase.P("Landlord")})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestExternalAccountsGet(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.ExternalAccounts.Get(context.TODO(), "external_account_ukk55lr923a3ac0pp7iv")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestExternalAccountsUpdateWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.ExternalAccounts.Update(
		context.TODO(),
		"external_account_ukk55lr923a3ac0pp7iv",
		&types.UpdateAnExternalAccountParameters{},
	)
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestExternalAccountsListWithOptionalParams(t *testing.T) {
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.ExternalAccounts.List(context.TODO(), &types.ExternalAccountListParams{})
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}