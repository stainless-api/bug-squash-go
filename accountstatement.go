package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

type AccountStatementService struct {
	Options []option.RequestOption
}

func NewAccountStatementService(opts ...option.RequestOption) (r *AccountStatementService) {
	r = &AccountStatementService{}
	r.Options = opts
	return
}

// Retrieve an Account Statement
func (r *AccountStatementService) Get(ctx context.Context, account_statement_id string, opts ...option.RequestOption) (res *AccountStatement, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("account_statements/%s", account_statement_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Account Statements
func (r *AccountStatementService) List(ctx context.Context, query AccountStatementListParams, opts ...option.RequestOption) (res *shared.Page[AccountStatement], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "account_statements"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Account Statements
func (r *AccountStatementService) ListAutoPaging(ctx context.Context, query AccountStatementListParams, opts ...option.RequestOption) *shared.PageAutoPager[AccountStatement] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type AccountStatement struct {
	// The Account Statement identifier.
	ID string `json:"id,required"`
	// The identifier for the Account this Account Statement belongs to.
	AccountID string `json:"account_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Statement was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the File containing a PDF of the statement.
	FileID string `json:"file_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the
	// start of the period the Account Statement covers.
	StatementPeriodStart time.Time `json:"statement_period_start,required" format:"date-time"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the end
	// of the period the Account Statement covers.
	StatementPeriodEnd time.Time `json:"statement_period_end,required" format:"date-time"`
	// The Account's balance at the start of its statement period.
	StartingBalance int64 `json:"starting_balance,required"`
	// The Account's balance at the start of its statement period.
	EndingBalance int64 `json:"ending_balance,required"`
	// A constant representing the object's type. For this resource it will always be
	// `account_statement`.
	Type AccountStatementType `json:"type,required"`
	JSON AccountStatementJSON
}

type AccountStatementJSON struct {
	ID                   apijson.Metadata
	AccountID            apijson.Metadata
	CreatedAt            apijson.Metadata
	FileID               apijson.Metadata
	StatementPeriodStart apijson.Metadata
	StatementPeriodEnd   apijson.Metadata
	StartingBalance      apijson.Metadata
	EndingBalance        apijson.Metadata
	Type                 apijson.Metadata
	raw                  string
	Extras               map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountStatement using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountStatement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStatementType string

const (
	AccountStatementTypeAccountStatement AccountStatementType = "account_statement"
)

type AccountStatementListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Account Statements to those belonging to the specified Account.
	AccountID            field.Field[string]                                         `query:"account_id"`
	StatementPeriodStart field.Field[AccountStatementListParamsStatementPeriodStart] `query:"statement_period_start"`
}

// URLQuery serializes AccountStatementListParams into a url.Values of the query
// parameters associated with this value
func (r AccountStatementListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type AccountStatementListParamsStatementPeriodStart struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes AccountStatementListParamsStatementPeriodStart into a
// url.Values of the query parameters associated with this value
func (r AccountStatementListParamsStatementPeriodStart) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type AccountStatementListResponse struct {
	// The contents of the list.
	Data []AccountStatement `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       AccountStatementListResponseJSON
}

type AccountStatementListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountStatementListResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AccountStatementListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}