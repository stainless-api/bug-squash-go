package services

import "context"
import "increase/core"
import "increase/types"
import "fmt"
import "increase/pagination"

type AccountStatementService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewAccountStatementService(requester core.Requester) (r *AccountStatementService) {
	r = &AccountStatementService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *AccountStatementService) Retrieve(ctx context.Context, account_statement_id string, opts ...*core.RequestOpts) (res *types.AccountStatement, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/account_statements/%s", account_statement_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

func (r *AccountStatementService) List(ctx context.Context, query *types.ListAccountStatementsQuery, opts ...*core.RequestOpts) (res *types.AccountStatementsPage, err error) {
	page := &types.AccountStatementsPage{
		Page: &pagination.Page[types.AccountStatement]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/account_statements",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
