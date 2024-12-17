package main

import (
	"context"
	"fmt"

	"github.com/acme/acme-go"
)

func main() {
	client := acme.NewClient()
	page, err := client.Cards.List(context.TODO(), acme.CardListParams{
		AccountID: acme.F("account_in71c4amph0vgo2qllky"),
	})
	for page != nil {
		fmt.Printf("%+#v %+#v\n", page, err)
		for _, item := range page.Data {
			println(item.ID)
		}
		page, err = page.GetNextPage()
	}
	if err != nil {
		panic(err.Error())
	}
}
