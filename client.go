// File generated from our OpenAPI spec by Stainless.

package acme

import (
	"os"

	"github.com/acme/acme-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the acme API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options        []option.RequestOption
	Accounts       *AccountService
	Transactions   *TransactionService
	ACHTransfers   *ACHTransferService
	Documents      *DocumentService
	WireTransfers  *WireTransferService
	CheckTransfers *CheckTransferService
	Files          *FileService
	Groups         *GroupService
	CheckDeposits  *CheckDepositService
	RoutingNumbers *RoutingNumberService
}

// NewClient generates a new client with the default option read from the
// environment (ACME_API_KEY). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("ACME_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Accounts = NewAccountService(opts...)
	r.Transactions = NewTransactionService(opts...)
	r.ACHTransfers = NewACHTransferService(opts...)
	r.Documents = NewDocumentService(opts...)
	r.WireTransfers = NewWireTransferService(opts...)
	r.CheckTransfers = NewCheckTransferService(opts...)
	r.Files = NewFileService(opts...)
	r.Groups = NewGroupService(opts...)
	r.CheckDeposits = NewCheckDepositService(opts...)
	r.RoutingNumbers = NewRoutingNumberService(opts...)

	return
}
