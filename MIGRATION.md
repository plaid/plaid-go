# Go SDK Migration Guide

Learn how to migrate to version 49 of the Plaid Go SDK, or continue using the previous SDK.

## Overview

Version 49 of `plaid-go` introduces a new SDK. Plaid can release updates to the new SDK more frequently, so new API features will be available in the SDK sooner after they launch.

The new SDK also requires less code for common tasks. You can make a call in a single statement, iterate over paginated results without writing your own loop, handle errors by type without converting them first, and enable retries for rate-limited requests and server errors.

### Summary of updates

- Endpoints are grouped by product, such as `c.Accounts.Get` and `c.Link.CreateToken`, instead of being methods on a single `PlaidApi` client.
- Each call is a single method that takes a request struct and returns the response and an error, instead of a chain of builder methods ending in `Execute()`.
- Request and response fields are plain struct fields, so you can build requests with struct literals and read responses without getter methods.
- Errors are typed by HTTP status and include the parsed Plaid error, so you can match them with `errors.As`.
- Many paginated endpoints provide an iterator, so you do not need to write your own pagination loop.
- Optional retries for requests that fail with a 408, 429, or 5xx status can be enabled with `option.WithMaxAttempts`. Retries are disabled by default.
- Dates are `time.Time` values instead of strings.

## Migration walkthrough

Version 49 also includes the previous SDK as a separate `legacy` package in the same module. This allows you to upgrade to version 49 without changing any calls, then migrate to the new SDK one endpoint at a time.

The overall process for migrating is:

1. [Install version 49](#install-version-49).
2. [Switch your imports to the legacy package](#switch-your-imports-to-the-legacy-package).
3. [Create a client with the new SDK](#create-a-client-with-the-new-sdk).
4. [Migrate calls one endpoint at a time](#migrate-calls-one-endpoint-at-a-time).
5. [Update pagination logic](#update-pagination-logic).
6. [Update error handling](#update-error-handling).
7. [Test your integration](#test-your-integration).
8. [Remove the legacy package](#remove-the-legacy-package).

## Install version 49

```bash
go get github.com/plaid/plaid-go/v49
```

> [!NOTE]
> Version 49 requires Go 1.21 or later. Version 48 supported Go 1.13 or later.

## Switch your imports to the legacy package

Version 49 contains two SDKs, each with its own import path:

| SDK          | Import path                            |
| ------------ | -------------------------------------- |
| New SDK      | `github.com/plaid/plaid-go/v49`        |
| Previous SDK | `github.com/plaid/plaid-go/v49/legacy` |

To use the previous SDK with version 49, update your import path:

```diff
- import "github.com/plaid/plaid-go/v48/plaid"
+ import "github.com/plaid/plaid-go/v49/legacy"
```

To update all version 48 imports in your project on macOS or Linux, run the following command from the root of your Go module:

```bash
gofmt -w -r '"github.com/plaid/plaid-go/v48/plaid" -> "github.com/plaid/plaid-go/v49/legacy"' . && go mod tidy
```

This command also reformats your Go files, and it will update files in `vendor/` if that directory is present (run `go mod vendor` afterward). After updating, run your project's tests.

The `legacy` package is still named `plaid`, so existing `plaid.*` references do not need to change.

## Create a client with the new SDK

Create a client with the new SDK alongside your existing client. Because both packages are named `plaid`, files that use both SDKs must import one of them under an alias. In the following example, the previous SDK is imported as `legacy`, so existing references to it in that file must use `legacy.*` instead of `plaid.*`:

```go
import (
    plaid "github.com/plaid/plaid-go/v49"
    "github.com/plaid/plaid-go/v49/client"
    "github.com/plaid/plaid-go/v49/option"

    legacy "github.com/plaid/plaid-go/v49/legacy"
)

// new SDK
c := client.NewClient(
    option.WithClientID(plaid.String(clientID)),
    option.WithSecret(secret),
    option.WithBaseURL(plaid.Environments.Sandbox),
)

// previous SDK
cfg := legacy.NewConfiguration()
cfg.AddDefaultHeader("PLAID-CLIENT-ID", clientID)
cfg.AddDefaultHeader("PLAID-SECRET", secret)
cfg.UseEnvironment(legacy.Sandbox)
legacyClient := legacy.NewAPIClient(cfg)
```

If your credentials are set in the `PLAID_CLIENT_ID` and `PLAID_SECRET` environment variables, you can omit them:

**New SDK, using environment variables**

```go
c := client.NewClient(option.WithBaseURL(plaid.Environments.Sandbox))
```

The new SDK reads `PLAID_CLIENT_ID`, `PLAID_SECRET`, and `PLAID_VERSION` from the environment when they are not passed explicitly.

Note that `option.WithClientID` takes a `*string`, while `option.WithSecret` takes a `string`. Use the `plaid.String()` helper to convert a string to a pointer.

> [!WARNING]
> Always set a non-empty base URL. As in version 48, a client initialized without an environment sends requests to Production. However, in the new SDK, `option.WithBaseURL("")` also sends requests to Production, whereas in version 48, `UseEnvironment("")` caused all requests to fail. If you read the base URL from configuration, return an error for unrecognized values rather than passing an empty string.

If your integration uses a dedicated Plaid endpoint or a proxy you control and trust to receive your Plaid credentials, configure it with `option.WithCustomEndpoint(baseURL, httpClient)`. This explicitly authorizes that destination. Use the standard Production or Sandbox configuration otherwise.

### Retries

Like the previous SDK, the new SDK does not retry failed requests by default. To enable retries for HTTP 408, 429, and 5xx responses, use `option.WithMaxAttempts(n)`. This limit includes the initial request: `WithMaxAttempts(2)` allows one retry.

## Migrate calls one endpoint at a time

Replace each call to the previous SDK with the equivalent call to the new SDK. The following table summarizes the differences between the two SDKs. The compiler will flag most of these changes; for changes that it will not flag, see [Test your integration](#test-your-integration).

| Change                  | Previous SDK                                                         | New SDK                                                         |
| ----------------------- | -------------------------------------------------------------------- | --------------------------------------------------------------- |
| Making a call           | `client.PlaidApi.AccountsGet(ctx).AccountsGetRequest(req).Execute()` | `c.Accounts.Get(ctx, &req)`                                     |
| Client grouping         | A single `PlaidApi` client                                           | Grouped by product: `c.Accounts`, `c.Auth`, `c.Item`, `c.Link`… |
| Return values           | `(resp, *http.Response, error)`                                      | `(resp, error)`                                                 |
| Building requests       | `plaid.NewAccountsGetRequest(token)`                                 | `&plaid.AccountsGetRequest{AccessToken: token}`                 |
| Setting optional fields | `req.SetCursor(cursor)`                                              | `Cursor: &cursor`                                               |
| Reading fields          | `resp.GetAccounts()`                                                 | `resp.Accounts`                                                 |
| Initialisms             | `GetItemId()`, `GetAccountId()`                                      | `.ItemID`, `.AccountID`                                         |
| Checking presence       | `resp.GetBalanceOk()`, `resp.HasBalance()`                           | `if resp.Balance != nil`                                        |
| Enums                   | `PRODUCTS_AUTH`, `COUNTRYCODE_US`                                    | `ProductsAuth`, `CountryCodeUs`                                 |
| Nullable fields         | `plaid.NullableString`                                               | `*string`, with a `SetX` method                                 |
| Errors                  | A single `GenericOpenAPIError` type                                  | Typed errors, which can be matched with `errors.As`             |
| Dates                   | `string`                                                             | `time.Time`                                                     |
| Downloads               | `*os.File`                                                           | `io.Reader`                                                     |
| Status and headers      | Second return value, `*http.Response`                                | `c.Accounts.WithRawResponse.Get(...)`                           |

The following example calls `/accounts/balance/get` in each SDK:

**Previous SDK**

```go
bal, _, err := client.PlaidApi.AccountsBalanceGet(ctx).AccountsBalanceGetRequest(
    *plaid.NewAccountsBalanceGetRequest(accessToken),
).Execute()
if err != nil {
    return err
}
for _, a := range bal.GetAccounts() {
    balances := a.GetBalances()
    fmt.Println(a.GetAccountId(), balances.GetCurrent())
}
```

**New SDK**

```go
bal, err := c.Accounts.GetBalance(ctx, &plaid.AccountsBalanceGetRequest{
    AccessToken: accessToken,
})
if err != nil {
    return err
}
for _, a := range bal.Accounts {
    if a.Balances.Current != nil {
        fmt.Println(a.AccountID, *a.Balances.Current)
    }
}
```

The following example calls `/link/token/create`. Required fields that the previous SDK took as constructor arguments, such as `client_name`, `language`, and `country_codes`, are now named struct fields, and nested objects such as `user` are set as pointers to structs:

**Previous SDK**

```go
user := plaid.NewLinkTokenCreateRequestUser("user-123")
req := plaid.NewLinkTokenCreateRequest("My App", "en", []plaid.CountryCode{plaid.COUNTRYCODE_US})
req.SetUser(*user)
req.SetProducts([]plaid.Products{plaid.PRODUCTS_TRANSACTIONS})

resp, _, err := client.PlaidApi.LinkTokenCreate(ctx).
    LinkTokenCreateRequest(*req).Execute()
if err != nil {
    return err
}
linkToken := resp.GetLinkToken()
```

**New SDK**

```go
resp, err := c.Link.CreateToken(ctx, &plaid.LinkTokenCreateRequest{
    ClientName:   "My App",
    Language:     "en",
    CountryCodes: []plaid.CountryCode{plaid.CountryCodeUs},
    User: &plaid.LinkTokenCreateRequestUser{
        ClientUserID: "user-123",
    },
    Products: []plaid.Products{plaid.ProductsTransactions},
})
if err != nil {
    return err
}
linkToken := resp.LinkToken
```

### Method and type names

Initialisms in type names are no longer capitalized, so `ACHClass` becomes `AchClass`. If you cannot find the new name for a method or type, search the `reference.md` file at the root of the SDK repository for the endpoint path.

### Nullable fields and empty arrays

In the new SDK, nullable fields are pointers, so check for `nil` before dereferencing them. Getter methods such as `a.GetMask()` are still available and safe to call, but for nullable fields they now return a pointer. As a result, code such as `fmt.Println(a.GetMask())` will still compile, but will print a memory address instead of the mask.

The previous SDK used `Nullable[T]` types to distinguish between an omitted field and a field explicitly set to `null`. The new SDK uses pointers and `SetX` methods instead. To omit a field, leave it `nil`. To send a value, set the field. To send an explicit `null`, call `SetX(nil)`:

```go
cursor := "abc"
req := &plaid.TransactionsSyncRequest{AccessToken: "access-sandbox-71e02f71-..."}
// cursor omitted entirely:            {"access_token":"access-sandbox-71e02f71-..."}

req.SetCursor(nil)
// cursor sent as an explicit null:    {"access_token":"access-sandbox-71e02f71-...","cursor":null}

req.SetCursor(&cursor)
// cursor sent with a value:           {"access_token":"access-sandbox-71e02f71-...","cursor":"abc"}
```

An empty slice set in a struct literal, such as `AccountIDs: []string{}`, is omitted from the request, whereas the previous SDK sent `"account_ids": []`. If you need to send an empty array rather than omit the field, call the setter, such as `SetAccountIDs([]string{})`.

### Default values

Many constructors in the previous SDK filled in default values for optional fields and sent them with the request. For example, `NewTransactionsSyncRequest()` sent `count: 100`, and `NewTransactionsGetRequestOptions()` sent `count: 100`, `offset: 0`, and `days_requested: 90`. These values are the defaults listed for each field in the [API reference](https://plaid.com/docs/api/). The new SDK has no constructors, so it sends only the fields you set, and Plaid applies its server-side defaults to the fields you omit.

Server-side defaults are not guaranteed to match the values that the previous SDK sent, so omitting a field can change your integration's behavior. If your integration depends on one of these values, set it explicitly.

### Downloads

Endpoints that return a file, such as `/asset_report/pdf/get`, now return an `io.Reader` instead of an `*os.File`. To save the PDF, copy the reader to a file:

```go
pdf, err := c.AssetReport.GetPdf(ctx, &plaid.AssetReportPdfGetRequest{
    AssetReportToken: assetReportToken,
})
if err != nil {
    return err
}
f, err := os.Create("asset_report.pdf")
if err != nil {
    return err
}
defer f.Close()
if _, err := io.Copy(f, pdf); err != nil {
    return err
}
```

### Status codes and headers

Calls no longer return the `*http.Response`. To read the HTTP status code or headers, make the call through `WithRawResponse`, which returns the status code, headers, and response body:

```go
resp, err := c.Accounts.WithRawResponse.Get(ctx, &plaid.AccountsGetRequest{AccessToken: accessToken})
if err != nil {
    return err
}
fmt.Println(resp.StatusCode, resp.Header.Get("Content-Type"))
accounts := resp.Body.Accounts
```

## Update pagination logic

In the new SDK, many paginated list endpoints return a page with an iterator, so you no longer need to write your own pagination loop. This applies to both cursor-based endpoints and offset-based endpoints.

**Previous SDK**

```go
req := plaid.NewWatchlistScreeningIndividualListRequest(programID)
for {
    resp, _, err := client.PlaidApi.WatchlistScreeningIndividualList(ctx).
        WatchlistScreeningIndividualListRequest(*req).Execute()
    if err != nil {
        return err
    }
    for _, s := range resp.GetWatchlistScreenings() {
        process(s)
    }
    next := resp.GetNextCursor()
    if next == "" {
        break
    }
    req.SetCursor(next)
}
```

**New SDK**

```go
page, err := c.WatchlistScreeningIndividual.List(ctx, &plaid.WatchlistScreeningIndividualListRequest{
    WatchlistProgramID: programID,
})
if err != nil {
    return err
}

iter := page.Iterator()
for iter.Next(ctx) {
    process(iter.Current())
}
if err := iter.Err(); err != nil {
    return err
}
```

`page.Results` contains the items in the current page, and `page.Response` contains the full response, including fields outside the list, such as `TotalTransactions` for `/transactions/get`. To paginate manually, call `page.GetNextPage(ctx)`, which returns `core.ErrNoPages` after the last page. The iterator stops when the server returns an empty page.

The following endpoints do not provide an iterator, so continue to use your existing pagination loop for them:

- `/transactions/sync` and `/processor/transactions/sync`, which return three arrays (`added`, `modified`, and `removed`) rather than a single list
- `/transfer/event/sync` and `/bank_transfer/event/sync`, which paginate by event ID (`after_id`)
- `/cashflow_report/get` and `/cashflow_report/transactions/get`

To determine whether an endpoint provides an iterator, check its return type. Endpoints that return `*core.Page[...]` provide an iterator.

### Sync transactions

To paginate `/transactions/sync`, keep your existing cursor loop. Only the call style has changed:

**Previous SDK**

```go
cursor := ""
for {
    req := plaid.NewTransactionsSyncRequest(accessToken)
    req.SetCursor(cursor)
    tx, _, err := client.PlaidApi.TransactionsSync(ctx).
        TransactionsSyncRequest(*req).Execute()
    if err != nil {
        return err
    }
    process(tx.GetAdded(), tx.GetModified(), tx.GetRemoved())
    if !tx.GetHasMore() {
        break
    }
    cursor = tx.GetNextCursor()
}
```

**New SDK**

```go
cursor := ""
for {
    tx, err := c.Transactions.Sync(ctx, &plaid.TransactionsSyncRequest{
        AccessToken: accessToken,
        Cursor:      &cursor,
    })
    if err != nil {
        return err
    }
    process(tx.Added, tx.Modified, tx.Removed)
    if !tx.HasMore {
        break
    }
    cursor = tx.NextCursor
}
```

## Update error handling

The previous SDK returns a single error type, which you convert with `plaid.ToPlaidError`. The new SDK returns typed errors that you can match with `errors.As`. Each error's `Body` field contains the Plaid error as a `*plaid.PlaidError`, so you can compare its fields directly. For example, `ITEM_LOGIN_REQUIRED` has a 400 status, so the new SDK returns it as a `BadRequestError`. To detect it and send the user through [update mode](https://plaid.com/docs/link/update-mode/):

**Previous SDK**

```go
if err != nil {
    plaidErr, convErr := plaid.ToPlaidError(err)
    if convErr == nil && plaidErr.ErrorCode == "ITEM_LOGIN_REQUIRED" {
        // send the user through update mode
    }
    return err
}
```

**New SDK**

```go
if err != nil {
    var badReq *plaid.BadRequestError
    var unauth *plaid.UnauthorizedError
    var rateLimited *plaid.TooManyRequestsError

    switch {
    case errors.As(err, &badReq):
        if badReq.Body != nil && badReq.Body.ErrorCode == "ITEM_LOGIN_REQUIRED" {
            // send the user through update mode
        }
    case errors.As(err, &unauth):
        // re-authenticate
    case errors.As(err, &rateLimited):
        // back off and retry
    }
    return err
}
```

The new SDK provides the following error types:

- `BadRequestError` (400)
- `UnauthorizedError` (401)
- `TooManyRequestsError` (429)
- `InternalServerError` (500)
- `ClientRequestError` (any other 4xx status)
- `ServerError` (any other 5xx status)

Each error type also includes the HTTP `StatusCode` and `Header`. Note that `PlaidError.Status` is now a `*int`, so check it for `nil` before reading it.

To handle all Plaid errors the same way regardless of status code, as with `plaid.ToPlaidError` in the previous SDK, match on the `GetBody()` method that all six error types share. For example, to print the type and code of any Plaid error:

**Print any Plaid error**

```go
var plaidErr interface{ GetBody() *plaid.PlaidError }
if errors.As(err, &plaidErr) && plaidErr.GetBody() != nil {
    fmt.Println(plaidErr.GetBody().ErrorType, plaidErr.GetBody().ErrorCode)
}
```

If the response body is empty or is not a Plaid error (for example, if it is an HTML page returned by a proxy), the SDK returns a `*core.APIError` that contains the status code but no parsed Plaid error.

## Test your integration

After migrating each endpoint, run your project's tests. Some changes do not cause compile errors, so confirm the following:

- Every call to `option.WithBaseURL` receives a valid environment URL, including when the URL is read from configuration. An [empty base URL](#create-a-client-with-the-new-sdk) sends requests to Production.
- Code reading [nullable fields](#nullable-fields-and-empty-arrays), including through getter methods such as `GetMask()`, checks for `nil` and dereferences the pointer.
- If your integration sends an explicit `null` or an empty array to mean something different from omitting the field, the request body still includes it. The new SDK sends these values only when you set them with a [setter method](#nullable-fields-and-empty-arrays). This is uncommon; fields that you do not use can be left unset.
- If your integration depends on a [default value](#default-values) that a previous SDK constructor filled in, such as `count` or `days_requested`, the request now sets that value explicitly.
- If you enable SDK [retries](#retries), confirm the total number of attempts is what you intend, especially if your code also retries failed requests.

## Remove the legacy package

After you have migrated every call, remove the `legacy` import and the previous SDK client.

## Getting help

For help with the Go SDK, open an issue in the [Plaid Go SDK repository](https://github.com/plaid/plaid-go/issues).
