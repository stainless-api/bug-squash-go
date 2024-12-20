// File generated from our OpenAPI spec by Stainless.

package acme

import (
	"context"
	"fmt"
	"net/http"

	"github.com/acme/acme-go/internal/requestconfig"
	"github.com/acme/acme-go/option"
)

// SimulationAccountTransferService contains methods and other services that help
// with interacting with the acme API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSimulationAccountTransferService] method instead.
type SimulationAccountTransferService struct {
	Options []option.RequestOption
}

// NewSimulationAccountTransferService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimulationAccountTransferService(opts ...option.RequestOption) (r *SimulationAccountTransferService) {
	r = &SimulationAccountTransferService{}
	r.Options = opts
	return
}

// If your account is configured to require approval for each transfer, this
// endpoint simulates the approval of an [Account Transfer](#account-transfers).
// You can also approve sandbox Account Transfers in the dashboard. This transfer
// must first have a `status` of `pending_approval`.
func (r *SimulationAccountTransferService) Complete(ctx context.Context, accountTransferID string, opts ...option.RequestOption) (res *AccountTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/account_transfers/%s/complete", accountTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
