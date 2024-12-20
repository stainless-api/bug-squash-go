// File generated from our OpenAPI spec by Stainless.

package acme

import (
	"context"
	"fmt"
	"net/http"

	"github.com/acme/acme-go/internal/requestconfig"
	"github.com/acme/acme-go/option"
)

// SimulationCheckTransferService contains methods and other services that help
// with interacting with the acme API. Note, unlike clients, this service does
// not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSimulationCheckTransferService] method instead.
type SimulationCheckTransferService struct {
	Options []option.RequestOption
}

// NewSimulationCheckTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimulationCheckTransferService(opts ...option.RequestOption) (r *SimulationCheckTransferService) {
	r = &SimulationCheckTransferService{}
	r.Options = opts
	return
}

// Simulates a [Check Transfer](#check-transfers) being deposited at a bank. This
// transfer must first have a `status` of `mailed`.
func (r *SimulationCheckTransferService) Deposit(ctx context.Context, checkTransferID string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_transfers/%s/deposit", checkTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simulates the mailing of a [Check Transfer](#check-transfers), which happens
// once per weekday in production but can be sped up in sandbox. This transfer must
// first have a `status` of `pending_approval` or `pending_submission`.
func (r *SimulationCheckTransferService) Mail(ctx context.Context, checkTransferID string, opts ...option.RequestOption) (res *CheckTransfer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("simulations/check_transfers/%s/mail", checkTransferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
