# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BalanceLookup">BalanceLookup</a>

Methods:

- <code title="post /accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNewParams">AccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}">client.Accounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}">client.Accounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountUpdateParams">AccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountListParams">AccountListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Account">Account</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/balance">client.Accounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountService.Balance">Balance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountBalanceParams">AccountBalanceParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BalanceLookup">BalanceLookup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/{account_id}/close">client.Accounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccountNumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumber">AccountNumber</a>

Methods:

- <code title="post /account_numbers">client.AccountNumbers.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumberService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumberNewParams">AccountNumberNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumber">AccountNumber</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account_numbers/{account_number_id}">client.AccountNumbers.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumberService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountNumberID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumber">AccountNumber</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /account_numbers/{account_number_id}">client.AccountNumbers.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumberService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountNumberID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumberUpdateParams">AccountNumberUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumber">AccountNumber</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account_numbers">client.AccountNumbers.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumberListParams">AccountNumberListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountNumber">AccountNumber</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BookkeepingAccounts

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccount">BookkeepingAccount</a>
- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingBalanceLookup">BookkeepingBalanceLookup</a>

Methods:

- <code title="post /bookkeeping_accounts">client.BookkeepingAccounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccountNewParams">BookkeepingAccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccount">BookkeepingAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /bookkeeping_accounts/{bookkeeping_account_id}">client.BookkeepingAccounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bookkeepingAccountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccountUpdateParams">BookkeepingAccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccount">BookkeepingAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bookkeeping_accounts">client.BookkeepingAccounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccountListParams">BookkeepingAccountListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccount">BookkeepingAccount</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bookkeeping_accounts/{bookkeeping_account_id}/balance">client.BookkeepingAccounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccountService.Balance">Balance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bookkeepingAccountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingAccountBalanceParams">BookkeepingAccountBalanceParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingBalanceLookup">BookkeepingBalanceLookup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BookkeepingEntrySets

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntrySet">BookkeepingEntrySet</a>

Methods:

- <code title="post /bookkeeping_entry_sets">client.BookkeepingEntrySets.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntrySetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntrySetNewParams">BookkeepingEntrySetNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntrySet">BookkeepingEntrySet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bookkeeping_entry_sets/{bookkeeping_entry_set_id}">client.BookkeepingEntrySets.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntrySetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bookkeepingEntrySetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntrySet">BookkeepingEntrySet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bookkeeping_entry_sets">client.BookkeepingEntrySets.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntrySetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntrySetListParams">BookkeepingEntrySetListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntrySet">BookkeepingEntrySet</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# BookkeepingEntries

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntry">BookkeepingEntry</a>

Methods:

- <code title="get /bookkeeping_entries/{bookkeeping_entry_id}">client.BookkeepingEntries.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bookkeepingEntryID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntry">BookkeepingEntry</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bookkeeping_entries">client.BookkeepingEntries.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntryListParams">BookkeepingEntryListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#BookkeepingEntry">BookkeepingEntry</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RealTimeDecisions

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimeDecision">RealTimeDecision</a>

Methods:

- <code title="get /real_time_decisions/{real_time_decision_id}">client.RealTimeDecisions.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimeDecisionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, realTimeDecisionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimeDecision">RealTimeDecision</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /real_time_decisions/{real_time_decision_id}/action">client.RealTimeDecisions.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimeDecisionService.Action">Action</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, realTimeDecisionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimeDecisionActionParams">RealTimeDecisionActionParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimeDecision">RealTimeDecision</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RealTimePaymentsTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimePaymentsTransfer">RealTimePaymentsTransfer</a>

Methods:

- <code title="post /real_time_payments_transfers">client.RealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimePaymentsTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimePaymentsTransferNewParams">RealTimePaymentsTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimePaymentsTransfer">RealTimePaymentsTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /real_time_payments_transfers/{real_time_payments_transfer_id}">client.RealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimePaymentsTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, realTimePaymentsTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimePaymentsTransfer">RealTimePaymentsTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /real_time_payments_transfers">client.RealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimePaymentsTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimePaymentsTransferListParams">RealTimePaymentsTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimePaymentsTransfer">RealTimePaymentsTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cards

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDetails">CardDetails</a>

Methods:

- <code title="post /cards">client.Cards.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardNewParams">CardNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cards/{card_id}">client.Cards.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /cards/{card_id}">client.Cards.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardUpdateParams">CardUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cards">client.Cards.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardListParams">CardListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Card">Card</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cards/{card_id}/details">client.Cards.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardService.GetSensitiveDetails">GetSensitiveDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDetails">CardDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CardDisputes

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDispute">CardDispute</a>

Methods:

- <code title="post /card_disputes">client.CardDisputes.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDisputeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDisputeNewParams">CardDisputeNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDispute">CardDispute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card_disputes/{card_dispute_id}">client.CardDisputes.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDisputeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardDisputeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDispute">CardDispute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card_disputes">client.CardDisputes.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDisputeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDisputeListParams">CardDisputeListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDispute">CardDispute</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CardProfiles

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfile">CardProfile</a>

Methods:

- <code title="post /card_profiles">client.CardProfiles.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfileNewParams">CardProfileNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfile">CardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card_profiles/{card_profile_id}">client.CardProfiles.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardProfileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfile">CardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card_profiles">client.CardProfiles.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfileListParams">CardProfileListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfile">CardProfile</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /card_profiles/{card_profile_id}/archive">client.CardProfiles.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfileService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardProfileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfile">CardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CardPurchaseSupplements

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPurchaseSupplement">CardPurchaseSupplement</a>

Methods:

- <code title="get /card_purchase_supplements/{card_purchase_supplement_id}">client.CardPurchaseSupplements.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPurchaseSupplementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardPurchaseSupplementID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPurchaseSupplement">CardPurchaseSupplement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card_purchase_supplements">client.CardPurchaseSupplements.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPurchaseSupplementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPurchaseSupplementListParams">CardPurchaseSupplementListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPurchaseSupplement">CardPurchaseSupplement</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ExternalAccounts

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccount">ExternalAccount</a>

Methods:

- <code title="post /external_accounts">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccountNewParams">ExternalAccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccount">ExternalAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /external_accounts/{external_account_id}">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, externalAccountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccount">ExternalAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /external_accounts/{external_account_id}">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, externalAccountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccountUpdateParams">ExternalAccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccount">ExternalAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /external_accounts">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccountListParams">ExternalAccountListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExternalAccount">ExternalAccount</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Exports

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Export">Export</a>

Methods:

- <code title="post /exports">client.Exports.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExportService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExportNewParams">ExportNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Export">Export</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /exports/{export_id}">client.Exports.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExportService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, exportID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Export">Export</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /exports">client.Exports.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExportService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ExportListParams">ExportListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Export">Export</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DigitalWalletTokens

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#DigitalWalletToken">DigitalWalletToken</a>

Methods:

- <code title="get /digital_wallet_tokens/{digital_wallet_token_id}">client.DigitalWalletTokens.<a href="https://pkg.go.dev/github.com/acme/acme-go#DigitalWalletTokenService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, digitalWalletTokenID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#DigitalWalletToken">DigitalWalletToken</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /digital_wallet_tokens">client.DigitalWalletTokens.<a href="https://pkg.go.dev/github.com/acme/acme-go#DigitalWalletTokenService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#DigitalWalletTokenListParams">DigitalWalletTokenListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#DigitalWalletToken">DigitalWalletToken</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Transaction">Transaction</a>

Methods:

- <code title="get /transactions/{transaction_id}">client.Transactions.<a href="https://pkg.go.dev/github.com/acme/acme-go#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/acme/acme-go#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#TransactionListParams">TransactionListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Transaction">Transaction</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PendingTransactions

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PendingTransaction">PendingTransaction</a>

Methods:

- <code title="get /pending_transactions/{pending_transaction_id}">client.PendingTransactions.<a href="https://pkg.go.dev/github.com/acme/acme-go#PendingTransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pendingTransactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PendingTransaction">PendingTransaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pending_transactions">client.PendingTransactions.<a href="https://pkg.go.dev/github.com/acme/acme-go#PendingTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PendingTransactionListParams">PendingTransactionListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PendingTransaction">PendingTransaction</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Programs

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Program">Program</a>

Methods:

- <code title="get /programs/{program_id}">client.Programs.<a href="https://pkg.go.dev/github.com/acme/acme-go#ProgramService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, programID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Program">Program</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /programs">client.Programs.<a href="https://pkg.go.dev/github.com/acme/acme-go#ProgramService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ProgramListParams">ProgramListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Program">Program</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DeclinedTransactions

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#DeclinedTransaction">DeclinedTransaction</a>

Methods:

- <code title="get /declined_transactions/{declined_transaction_id}">client.DeclinedTransactions.<a href="https://pkg.go.dev/github.com/acme/acme-go#DeclinedTransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, declinedTransactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#DeclinedTransaction">DeclinedTransaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /declined_transactions">client.DeclinedTransactions.<a href="https://pkg.go.dev/github.com/acme/acme-go#DeclinedTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#DeclinedTransactionListParams">DeclinedTransactionListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#DeclinedTransaction">DeclinedTransaction</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccountTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransfer">AccountTransfer</a>

Methods:

- <code title="post /account_transfers">client.AccountTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransferNewParams">AccountTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransfer">AccountTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account_transfers/{account_transfer_id}">client.AccountTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransfer">AccountTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account_transfers">client.AccountTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransferListParams">AccountTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransfer">AccountTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /account_transfers/{account_transfer_id}/approve">client.AccountTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransferService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransfer">AccountTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /account_transfers/{account_transfer_id}/cancel">client.AccountTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransferService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransfer">AccountTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ACHTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransfer">ACHTransfer</a>

Methods:

- <code title="post /ach_transfers">client.ACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransferNewParams">ACHTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ach_transfers/{ach_transfer_id}">client.ACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ach_transfers">client.ACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransferListParams">ACHTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransfer">ACHTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ach_transfers/{ach_transfer_id}/approve">client.ACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransferService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ach_transfers/{ach_transfer_id}/cancel">client.ACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransferService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ACHPrenotifications

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHPrenotification">ACHPrenotification</a>

Methods:

- <code title="post /ach_prenotifications">client.ACHPrenotifications.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHPrenotificationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHPrenotificationNewParams">ACHPrenotificationNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHPrenotification">ACHPrenotification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ach_prenotifications/{ach_prenotification_id}">client.ACHPrenotifications.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHPrenotificationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achPrenotificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHPrenotification">ACHPrenotification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ach_prenotifications">client.ACHPrenotifications.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHPrenotificationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHPrenotificationListParams">ACHPrenotificationListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHPrenotification">ACHPrenotification</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Documents

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Document">Document</a>

Methods:

- <code title="get /documents/{document_id}">client.Documents.<a href="https://pkg.go.dev/github.com/acme/acme-go#DocumentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, documentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Document">Document</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /documents">client.Documents.<a href="https://pkg.go.dev/github.com/acme/acme-go#DocumentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#DocumentListParams">DocumentListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Document">Document</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WireTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransfer">WireTransfer</a>

Methods:

- <code title="post /wire_transfers">client.WireTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransferNewParams">WireTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /wire_transfers/{wire_transfer_id}">client.WireTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /wire_transfers">client.WireTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransferListParams">WireTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransfer">WireTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /wire_transfers/{wire_transfer_id}/approve">client.WireTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransferService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /wire_transfers/{wire_transfer_id}/cancel">client.WireTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransferService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/wire_transfers/{wire_transfer_id}/reverse">client.WireTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransferService.Reverse">Reverse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/wire_transfers/{wire_transfer_id}/submit">client.WireTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransferService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransfer">WireTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CheckTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransfer">CheckTransfer</a>

Methods:

- <code title="post /check_transfers">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransferNewParams">CheckTransferNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /check_transfers/{check_transfer_id}">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /check_transfers">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransferListParams">CheckTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransfer">CheckTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /check_transfers/{check_transfer_id}/approve">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransferService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /check_transfers/{check_transfer_id}/cancel">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransferService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /check_transfers/{check_transfer_id}/stop_payment">client.CheckTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransferService.StopPayment">StopPayment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransferStopPaymentParams">CheckTransferStopPaymentParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Entities

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Entity">Entity</a>

Methods:

- <code title="post /entities">client.Entities.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityNewParams">EntityNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entities/{entity_id}">client.Entities.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entities">client.Entities.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityListParams">EntityListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Entity">Entity</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /entities/{entity_id}/archive">client.Entities.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /entities/{entity_id}/address">client.Entities.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityService.UpdateAddress">UpdateAddress</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityUpdateAddressParams">EntityUpdateAddressParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## BeneficialOwners

Methods:

- <code title="post /entity_beneficial_owners">client.Entities.BeneficialOwners.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityBeneficialOwnerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityBeneficialOwnerNewParams">EntityBeneficialOwnerNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /entity_beneficial_owners/archive">client.Entities.BeneficialOwners.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityBeneficialOwnerService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityBeneficialOwnerArchiveParams">EntityBeneficialOwnerArchiveParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /entity_beneficial_owners/address">client.Entities.BeneficialOwners.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityBeneficialOwnerService.UpdateAddress">UpdateAddress</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntityBeneficialOwnerUpdateAddressParams">EntityBeneficialOwnerUpdateAddressParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## SupplementalDocuments

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SupplementalDocument">SupplementalDocument</a>

Methods:

- <code title="post /entities/{entity_id}/supplemental_documents">client.Entities.SupplementalDocuments.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntitySupplementalDocumentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntitySupplementalDocumentNewParams">EntitySupplementalDocumentNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entity_supplemental_documents">client.Entities.SupplementalDocuments.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntitySupplementalDocumentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EntitySupplementalDocumentListParams">EntitySupplementalDocumentListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SupplementalDocument">SupplementalDocument</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InboundACHTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransfer">InboundACHTransfer</a>

Methods:

- <code title="get /inbound_ach_transfers/{inbound_ach_transfer_id}">client.InboundACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransferService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundACHTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransfer">InboundACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /inbound_ach_transfers">client.InboundACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransferListParams">InboundACHTransferListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransfer">InboundACHTransfer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /inbound_ach_transfers/{inbound_ach_transfer_id}/decline">client.InboundACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransferService.Decline">Decline</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundACHTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransfer">InboundACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /inbound_ach_transfers/{inbound_ach_transfer_id}/notification_of_change">client.InboundACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransferService.NotificationOfChange">NotificationOfChange</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundACHTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransferNotificationOfChangeParams">InboundACHTransferNotificationOfChangeParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransfer">InboundACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /inbound_ach_transfers/{inbound_ach_transfer_id}/transfer_return">client.InboundACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransferService.TransferReturn">TransferReturn</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundACHTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransferTransferReturnParams">InboundACHTransferTransferReturnParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundACHTransfer">InboundACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InboundWireDrawdownRequests

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundWireDrawdownRequest">InboundWireDrawdownRequest</a>

Methods:

- <code title="get /inbound_wire_drawdown_requests/{inbound_wire_drawdown_request_id}">client.InboundWireDrawdownRequests.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundWireDrawdownRequestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundWireDrawdownRequestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundWireDrawdownRequest">InboundWireDrawdownRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /inbound_wire_drawdown_requests">client.InboundWireDrawdownRequests.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundWireDrawdownRequestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundWireDrawdownRequestListParams">InboundWireDrawdownRequestListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundWireDrawdownRequest">InboundWireDrawdownRequest</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WireDrawdownRequests

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireDrawdownRequest">WireDrawdownRequest</a>

Methods:

- <code title="post /wire_drawdown_requests">client.WireDrawdownRequests.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireDrawdownRequestService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireDrawdownRequestNewParams">WireDrawdownRequestNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireDrawdownRequest">WireDrawdownRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /wire_drawdown_requests/{wire_drawdown_request_id}">client.WireDrawdownRequests.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireDrawdownRequestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, wireDrawdownRequestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireDrawdownRequest">WireDrawdownRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /wire_drawdown_requests">client.WireDrawdownRequests.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireDrawdownRequestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireDrawdownRequestListParams">WireDrawdownRequestListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireDrawdownRequest">WireDrawdownRequest</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Events

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Event">Event</a>

Methods:

- <code title="get /events/{event_id}">client.Events.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Event">Event</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /events">client.Events.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventListParams">EventListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Event">Event</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# EventSubscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscription">EventSubscription</a>

Methods:

- <code title="post /event_subscriptions">client.EventSubscriptions.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscriptionNewParams">EventSubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscription">EventSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /event_subscriptions/{event_subscription_id}">client.EventSubscriptions.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventSubscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscription">EventSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /event_subscriptions/{event_subscription_id}">client.EventSubscriptions.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventSubscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscriptionUpdateParams">EventSubscriptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscription">EventSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /event_subscriptions">client.EventSubscriptions.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscriptionListParams">EventSubscriptionListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#EventSubscription">EventSubscription</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#File">File</a>

Methods:

- <code title="post /files">client.Files.<a href="https://pkg.go.dev/github.com/acme/acme-go#FileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#FileNewParams">FileNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files/{file_id}">client.Files.<a href="https://pkg.go.dev/github.com/acme/acme-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /files">client.Files.<a href="https://pkg.go.dev/github.com/acme/acme-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#FileListParams">FileListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#File">File</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Groups

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Group">Group</a>

Methods:

- <code title="get /groups/current">client.Groups.<a href="https://pkg.go.dev/github.com/acme/acme-go#GroupService.GetDetails">GetDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Group">Group</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OauthConnections

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#OauthConnection">OauthConnection</a>

Methods:

- <code title="get /oauth_connections/{oauth_connection_id}">client.OauthConnections.<a href="https://pkg.go.dev/github.com/acme/acme-go#OauthConnectionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, oauthConnectionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#OauthConnection">OauthConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /oauth_connections">client.OauthConnections.<a href="https://pkg.go.dev/github.com/acme/acme-go#OauthConnectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#OauthConnectionListParams">OauthConnectionListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#OauthConnection">OauthConnection</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CheckDeposits

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDeposit">CheckDeposit</a>

Methods:

- <code title="post /check_deposits">client.CheckDeposits.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDepositService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDepositNewParams">CheckDepositNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDeposit">CheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /check_deposits/{check_deposit_id}">client.CheckDeposits.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDepositService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkDepositID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDeposit">CheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /check_deposits">client.CheckDeposits.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDepositService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDepositListParams">CheckDepositListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDeposit">CheckDeposit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RoutingNumbers

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RoutingNumber">RoutingNumber</a>

Methods:

- <code title="get /routing_numbers">client.RoutingNumbers.<a href="https://pkg.go.dev/github.com/acme/acme-go#RoutingNumberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RoutingNumberListParams">RoutingNumberListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RoutingNumber">RoutingNumber</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccountStatements

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountStatement">AccountStatement</a>

Methods:

- <code title="get /account_statements/{account_statement_id}">client.AccountStatements.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountStatementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountStatementID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountStatement">AccountStatement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account_statements">client.AccountStatements.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountStatementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountStatementListParams">AccountStatementListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountStatement">AccountStatement</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Simulations

## AccountTransfers

Methods:

- <code title="post /simulations/account_transfers/{account_transfer_id}/complete">client.Simulations.AccountTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationAccountTransferService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountTransfer">AccountTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AccountStatements

Methods:

- <code title="post /simulations/account_statements">client.Simulations.AccountStatements.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationAccountStatementService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationAccountStatementNewParams">SimulationAccountStatementNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#AccountStatement">AccountStatement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ACHTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransferSimulation">ACHTransferSimulation</a>

Methods:

- <code title="post /simulations/inbound_ach_transfers">client.Simulations.ACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationACHTransferService.NewInbound">NewInbound</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationACHTransferNewInboundParams">SimulationACHTransferNewInboundParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransferSimulation">ACHTransferSimulation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/ach_transfers/{ach_transfer_id}/return">client.Simulations.ACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationACHTransferService.Return">Return</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationACHTransferReturnParams">SimulationACHTransferReturnParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/ach_transfers/{ach_transfer_id}/submit">client.Simulations.ACHTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationACHTransferService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, achTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#ACHTransfer">ACHTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardDisputes

Methods:

- <code title="post /simulations/card_disputes/{card_dispute_id}/action">client.Simulations.CardDisputes.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCardDisputeService.Action">Action</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardDisputeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCardDisputeActionParams">SimulationCardDisputeActionParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardDispute">CardDispute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardProfiles

Methods:

- <code title="post /simulations/card_profiles/{card_profile_id}/approve">client.Simulations.CardProfiles.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCardProfileService.Approve">Approve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardProfileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardProfile">CardProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardRefunds

Methods:

- <code title="post /simulations/card_refunds">client.Simulations.CardRefunds.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCardRefundService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCardRefundNewParams">SimulationCardRefundNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CheckTransfers

Methods:

- <code title="post /simulations/check_transfers/{check_transfer_id}/deposit">client.Simulations.CheckTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCheckTransferService.Deposit">Deposit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/check_transfers/{check_transfer_id}/mail">client.Simulations.CheckTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCheckTransferService.Mail">Mail</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkTransferID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckTransfer">CheckTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Documents

Methods:

- <code title="post /simulations/documents">client.Simulations.Documents.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationDocumentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationDocumentNewParams">SimulationDocumentNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Document">Document</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## DigitalWalletTokenRequests

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationDigitalWalletTokenRequestNewResponse">SimulationDigitalWalletTokenRequestNewResponse</a>

Methods:

- <code title="post /simulations/digital_wallet_token_requests">client.Simulations.DigitalWalletTokenRequests.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationDigitalWalletTokenRequestService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationDigitalWalletTokenRequestNewParams">SimulationDigitalWalletTokenRequestNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationDigitalWalletTokenRequestNewResponse">SimulationDigitalWalletTokenRequestNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CheckDeposits

Methods:

- <code title="post /simulations/check_deposits/{check_deposit_id}/reject">client.Simulations.CheckDeposits.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCheckDepositService.Reject">Reject</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkDepositID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDeposit">CheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/check_deposits/{check_deposit_id}/return">client.Simulations.CheckDeposits.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCheckDepositService.Return">Return</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkDepositID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDeposit">CheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/check_deposits/{check_deposit_id}/submit">client.Simulations.CheckDeposits.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCheckDepositService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, checkDepositID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CheckDeposit">CheckDeposit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Programs

Methods:

- <code title="post /simulations/programs">client.Simulations.Programs.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationProgramService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationProgramNewParams">SimulationProgramNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Program">Program</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InboundWireDrawdownRequests

Methods:

- <code title="post /simulations/inbound_wire_drawdown_requests">client.Simulations.InboundWireDrawdownRequests.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationInboundWireDrawdownRequestService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationInboundWireDrawdownRequestNewParams">SimulationInboundWireDrawdownRequestNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundWireDrawdownRequest">InboundWireDrawdownRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InboundFundsHolds

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationInboundFundsHoldReleaseResponse">SimulationInboundFundsHoldReleaseResponse</a>

Methods:

- <code title="post /simulations/inbound_funds_holds/{inbound_funds_hold_id}/release">client.Simulations.InboundFundsHolds.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationInboundFundsHoldService.Release">Release</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, inboundFundsHoldID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationInboundFundsHoldReleaseResponse">SimulationInboundFundsHoldReleaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InterestPayments

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InterestPaymentSimulationResult">InterestPaymentSimulationResult</a>

Methods:

- <code title="post /simulations/interest_payment">client.Simulations.InterestPayments.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationInterestPaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationInterestPaymentNewParams">SimulationInterestPaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InterestPaymentSimulationResult">InterestPaymentSimulationResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## WireTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransferSimulation">WireTransferSimulation</a>

Methods:

- <code title="post /simulations/inbound_wire_transfers">client.Simulations.WireTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationWireTransferService.NewInbound">NewInbound</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationWireTransferNewInboundParams">SimulationWireTransferNewInboundParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#WireTransferSimulation">WireTransferSimulation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Cards

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardAuthorizationSimulation">CardAuthorizationSimulation</a>

Methods:

- <code title="post /simulations/card_authorizations">client.Simulations.Cards.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCardService.Authorize">Authorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCardAuthorizeParams">SimulationCardAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardAuthorizationSimulation">CardAuthorizationSimulation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/card_settlements">client.Simulations.Cards.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCardService.Settlement">Settlement</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationCardSettlementParams">SimulationCardSettlementParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RealTimePaymentsTransfers

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundRealTimePaymentsTransferSimulationResult">InboundRealTimePaymentsTransferSimulationResult</a>

Methods:

- <code title="post /simulations/real_time_payments_transfers/{real_time_payments_transfer_id}/complete">client.Simulations.RealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationRealTimePaymentsTransferService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, realTimePaymentsTransferID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationRealTimePaymentsTransferCompleteParams">SimulationRealTimePaymentsTransferCompleteParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#RealTimePaymentsTransfer">RealTimePaymentsTransfer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /simulations/inbound_real_time_payments_transfers">client.Simulations.RealTimePaymentsTransfers.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationRealTimePaymentsTransferService.NewInbound">NewInbound</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationRealTimePaymentsTransferNewInboundParams">SimulationRealTimePaymentsTransferNewInboundParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#InboundRealTimePaymentsTransferSimulationResult">InboundRealTimePaymentsTransferSimulationResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PhysicalCards

Methods:

- <code title="post /simulations/physical_cards/{physical_card_id}/shipment_advance">client.Simulations.PhysicalCards.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationPhysicalCardService.ShipmentAdvance">ShipmentAdvance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, physicalCardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#SimulationPhysicalCardShipmentAdvanceParams">SimulationPhysicalCardShipmentAdvanceParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCard">PhysicalCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PhysicalCards

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCard">PhysicalCard</a>

Methods:

- <code title="post /physical_cards">client.PhysicalCards.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCardService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCardNewParams">PhysicalCardNewParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCard">PhysicalCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /physical_cards/{physical_card_id}">client.PhysicalCards.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, physicalCardID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCard">PhysicalCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /physical_cards/{physical_card_id}">client.PhysicalCards.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCardService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, physicalCardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCardUpdateParams">PhysicalCardUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCard">PhysicalCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /physical_cards">client.PhysicalCards.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCardListParams">PhysicalCardListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#PhysicalCard">PhysicalCard</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# CardPayments

Response Types:

- <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPayment">CardPayment</a>

Methods:

- <code title="get /card_payments/{card_payment_id}">client.CardPayments.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardPaymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPayment">CardPayment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /card_payments">client.CardPayments.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPaymentListParams">CardPaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#Page">Page</a>[<a href="https://pkg.go.dev/github.com/acme/acme-go">acme</a>.<a href="https://pkg.go.dev/github.com/acme/acme-go#CardPayment">CardPayment</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
