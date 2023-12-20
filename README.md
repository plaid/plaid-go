# plaid-go [![GoDoc](https://godoc.org/github.com/plaid/plaid-go?status.svg)](https://godoc.org/github.com/plaid/plaid-go/plaid)

The official Go client library for the [Plaid API](https://plaid.com/docs). The library is now generated from our [OpenAPI schema](https://github.com/plaid/plaid-openapi). If you are currently on a manually-written version of this library from August 2021 or earlier, see the [migration guide](#migration-guide).

The latest version of the library supports only the latest version of the Plaid API (currently 2020-09-14). 

For more information about the Plaid API, including reference documentation, see the [Plaid API docs](https://www.plaid.com/docs/api).

## Table of Contents

* [Installation](#installation)
* [Getting Started](#getting-started)
  + [Calling Endpoints](#calling-endpoints)
  + [Dates](#errors)
  + [Errors](#errors)
* [Examples](#examples)
* [Migration Guide](#migration-guide)
* [Contributing](#contributing)
* [License](#license)


## Installation

Library versions follow Semantic Versioning ("SemVer") and are formatted as `v1.0.0`. The plaid-go client library is typically updated on a monthly basis, and the canonical source for the latest version number is the [client library changelog](https://github.com/plaid/plaid-go/blob/master/CHANGELOG.md). New versions are published as [GitHub tags](https://github.com/plaid/plaid-go/tags), not as Releases. New versions are also published on [pkg.go.dev](https://pkg.go.dev/github.com/plaid/plaid-go). Make sure to look at the highest tagged version.

As of `v1.0.0`, we've moved to support `GOMODULES`.

Edit your go.mod to include `github.com/plaid/plaid-go/v20 {VERSION}`

```console
$ go get github.com/plaid/plaid-go/v20@{VERSION}
```

All users are strongly recommended to use a recent version of the library, as older versions do not contain support for new endpoints and fields. For more details, see the [Migration Guide](#migration-guide).

## Getting Started

### Calling Endpoints

To call an endpoint, you must create a Plaid API client. Here's how to configure it:

```go
configuration := plaid.NewConfiguration()
configuration.AddDefaultHeader("PLAID-CLIENT-ID", {VALUE})
configuration.AddDefaultHeader("PLAID-SECRET", {VALUE})
configuration.UseEnvironment(plaid.Production)
client := plaid.NewAPIClient(configuration)
```

Get your `client_id` and `secret` from your [dashboard account](https://dashboard.plaid.com).

Each endpoint will require an appropriate request model, and will return either the response model or an error.

### Dates

Dates and datetimes in requests, which are represented as strings in the APIs, are represented in this version of the Go client library as `time.Time` types. For more information, see the [Go documentation on the `time` package](https://pkg.go.dev/time).

Time zone information is required for request fields that accept datetimes. Failing to include time zone information (or specifying a string, instead of an instance of the `time` package) will result in an error.

If the API reference documentation for a request field specifies `format: date`, the following is acceptable:

```go
import (
    "time"
)

const (
    // Choose any arbitrary date in the desired format
    YYYYMMDD = "2006-01-02"
)

myDate := time.Date(2019, time.December, 6, 22, 35, 49, 0, time.UTC).Format(YYYYMMDD)
// Returns 2019-12-06
```

If the API reference documentation for a request field specifies `format: date-time`, the following is acceptable:


```go
import (
  "time"
)

myDate := time.Date(2019, time.December, 6, 22, 35, 49, 0, time.UTC)
// Returns 2019-12-06 22:35:49 +0000 UTC
```


### Errors

In the case one of the endpoints you call returns an error, you can get the Plaid error object with the following:

```go
response, httpResponse, err := client.PlaidApi.Endpoint(...)
plaidErr, err := plaid.ToPlaidError(err)
fmt.Println(plaidErr.ErrorMessage)
```

If you need to set up specific plaid errors to test your code, there is a helper function you can use to create a GenericOpenAPIError with embedded PlaidError:
```go
	plaidError := plaid.NewPlaidError(plaid.PLAIDERRORTYPE_ITEM_ERROR, "PRODUCT_NOT_READY", "", plaid.NullableString{})
	genericOpenAPIError := plaid.MakeGenericOpenAPIError([]byte{}, "400 Bad Request", *plaidError)
````

## Examples

The sections below show examples for some common API calls in Go. For more examples, see the [API reference docs](https://plaid.com/docs/api/).

#### Get Link Token to initialize Link

```go
user := plaid.LinkTokenCreateRequestUser{
    ClientUserId: "USER_ID_FROM_YOUR_DB",
}
request := plaid.NewLinkTokenCreateRequest(
  "Plaid Test",
  "en",
  []plaid.CountryCode{plaid.COUNTRYCODE_US},
  user,
)
request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})
request.SetLinkCustomizationName("default")
request.SetWebhook("https://webhook-uri.com")
request.SetRedirectUri("https://domainname.com/oauth-page.html")
request.SetAccountFilters(plaid.LinkTokenAccountFilters{
  Depository: &plaid.DepositoryFilter{
    AccountSubtypes: []plaid.AccountSubtype{plaid.ACCOUNTSUBTYPE_CHECKING, plaid.ACCOUNTSUBTYPE_SAVINGS},
  },
})
resp, _, err := testClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

linkToken := resp.GetLinkToken()
```

#### Create an Item using Link

Exchange a `public_token` from Plaid Link for a Plaid access token:
```go
exchangePublicTokenReq := plaid.NewItemPublicTokenExchangeRequest(sandboxPublicTokenResp.GetPublicToken())
exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
  *exchangePublicTokenReq,
).Execute()
accessToken := exchangePublicTokenResp.GetAccessToken()
```

#### Retrieve transactions (preferred method)

```go
request := plaid.NewTransactionsSyncRequest(
  accessToken
)

transactionsResp, _, err := testClient.PlaidApi.TransactionsSync(ctx).TransactionsSyncRequest(*request).Execute()
```

#### Retrieve transactions (older method)

```go
const iso8601TimeFormat = "2006-01-02"
startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
endDate := time.Now().Format(iso8601TimeFormat)

request := plaid.NewTransactionsGetRequest(
  accessToken,
  startDate,
  endDate,
)

options := plaid.TransactionsGetRequestOptions{
  Count:  plaid.PtrInt32(100),
  Offset: plaid.PtrInt32(0),
}

request.SetOptions(options)

transactionsResp, _, err := testClient.PlaidApi.TransactionsGet(ctx).TransactionsGetRequest(*request).Execute()
```

#### Retrieve real-time balance data

```go
balancesGetReq := plaid.NewAccountsBalanceGetRequest(accessToken)

balancesGetResp, _, err = testClient.PlaidApi.AccountsBalanceGet(ctx).AccountsBalanceGetRequest(
  *balancesGetReq,
).Execute()
```

#### Create a Dwolla bank account token

Exchange a Plaid Link `public_token` for an API `access_token`. Then exchange that `access_token` and the Plaid Link `account_id` (received along with the `public_token`) for a Dwolla processor token:

```go
import (
    "context"
    "os"

    "github.com/plaid/plaid-go/v20/plaid"
)

configuration := plaid.NewConfiguration()
configuration.AddDefaultHeader("PLAID-CLIENT-ID", os.Getenv("CLIENT_ID"))
configuration.AddDefaultHeader("PLAID-SECRET", os.Getenv("SECRET"))
configuration.UseEnvironment(plaid.Sandbox)

client := plaid.NewAPIClient(configuration)
ctx := context.Background()

// If not testing in Sandbox, remove these four lines and instead use a publicToken obtained from Link
sandboxInstitution := "ins_109508"
testProducts := []string{"auth"}
sandboxPublicTokenResp, _, err := client.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(
  *plaid.NewSandboxPublicTokenCreateRequest(
    sandboxInstitution,
    testProducts,
  ),
).Execute()
publicToken := sandboxPublicTokenResp.GetPublicToken()

// Exchange the publicToken for an accessToken
exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
  *plaid.NewItemPublicTokenExchangeRequest(publicToken),
).Execute()
accessToken := exchangePublicTokenResp.GetAccessToken()

// Get Accounts
accountsGetResp, _, err := client.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
  *plaid.NewAccountsGetRequest(accessToken),
).Execute()
accountID := accountsGetResp.GetAccounts()[0].GetAccountId()

// Create processor token
processorTokenCreateResp, _, err := client.PlaidApi.ProcessorTokenCreate(ctx).ProcessorTokenCreateRequest(
  *plaid.NewProcessorTokenCreateRequest(accessToken, accountID, "dwolla"),
).Execute()
```

## Migration guide

### Post-August 2021 to latest

Migrating from a version released on or after August 2021 to a recent version should involve very minor integration changes. Many customers will not need to make changes to their integrations at all. To see a list of all potentially-breaking changes since your current version, see the [client library changelog](https://github.com/plaid/plaid-go/blob/master/CHANGELOG.md) and search for "Breaking changes in this version". Breaking changes are annotated at the top of each major version header.

### Pre-August 2021 to latest

The client library was rewritten in August 2021 to be generated from OpenAPI. To update, you will need to update your integration as follows:

#### Client initialization
From:
```go
import (
  "github.com/plaid/plaid-go/plaid"
)

client, err := plaid.NewClient(plaid.ClientOptions{
  os.Getenv("PLAID_CLIENT_ID"),
  os.Getenv("PLAID_SECRET"),
  plaid.Sandbox,
})

```

To:
```go
import (
  "github.com/plaid/plaid-go/plaid"
)

configuration := plaid.NewConfiguration()
configuration.AddDefaultHeader("PLAID-CLIENT-ID", os.Getenv("CLIENT_ID"))
configuration.AddDefaultHeader("PLAID-SECRET", os.Getenv("SECRET"))
configuration.UseEnvironment(plaid.Sandbox)

client := plaid.NewAPIClient(configuration)
```

#### Endpoints
All endpoint requests now take a request model and the functions have been renamed to move the
verb to the end (e.g. `GetBalances` is now `BalancesGet`). We now accept a `context` object in order to give the caller more control over the [http request](https://pkg.go.dev/net/http#Request.WithContext).

From:
```go
response, err := client.CreateSandboxPublicToken(SandboxInstitution, TestProducts)
```

To:
```go
response, httpResponse, err := client.PlaidApi.SandboxPublicTokenCreate(context.Background()).SandboxPublicTokenCreateRequest(
    *plaid.NewSandboxPublicTokenCreateRequest(
      SandboxInstitution,
      TestProducts,
    ),
   ).Execute()
```

#### Errors
From:
```go
response, err := client.CreateSandboxPublicToken(SandboxInstitution, TestProducts)

plaidErr := err.(plaid.Error)
```

To:
```go
response, httpResponse, err := client.PlaidApi.SandboxPublicTokenCreate(context.Background()).SandboxPublicTokenCreateRequest(
    *plaid.NewSandboxPublicTokenCreateRequest(
      SandboxInstitution,
      TestProducts,
    ),
   ).Execute()

plaidErr, _ := plaid.ToPlaidError(err)
```


#### Balances example (Old way)
```
// get balances for all accounts
balanceResp, err := client.GetBalances(accessToken)

// get balances for selected accounts
options := GetBalancesOptions{
  AccountIDs: []string{balanceResp.Accounts[0].AccountID},
}
balanceResp, err = client.GetBalancesWithOptions(accessToken, options)
```

#### Balances example (New way)
```
balancesGetReq := plaid.NewAccountsBalanceGetRequest(accessToken)

balancesGetResp, _, err = testClient.PlaidApi.AccountsBalanceGet(ctx).AccountsBalanceGetRequest(
  *balancesGetReq,
).Execute()
```

#### Using optional parameters example (old way)

```
const iso8601TimeFormat = "2006-01-02"
startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
endDate := time.Now().Format(iso8601TimeFormat)
transactionsResp, err := client.GetTransactions(accessToken, startDate, endDate)

// Or, using optional parameters:
startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
endDate := time.Now().Format(iso8601TimeFormat)
options := GetTransactionsOptions{
  StartDate:  startDate,
  EndDate:    endDate,
  AccountIDs: []string{},
  Count:      2,
  Offset:     1,
}
transactionsResp, err := client.GetTransactionsWithOptions(accessToken, options)
```

#### Using optional parameters example (new way)

```
const iso8601TimeFormat = "2006-01-02"
startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
endDate := time.Now().Format(iso8601TimeFormat)
options := plaid.TransactionsGetRequestOptions{
    IncludePersonalFinanceCategory := true,
}
request.SetOptions(options)

request := plaid.NewTransactionsGetRequest(
  accessToken,
  startDate,
  endDate,
)

options := plaid.TransactionsGetRequestOptions{
  Count:  plaid.PtrInt32(100),
  Offset: plaid.PtrInt32(0),
  IncludePersonalFinanceCategory: true,
}

request.SetOptions(options)

transactionsResp, _, err := testClient.PlaidApi.TransactionsGet(ctx).TransactionsGetRequest(*request).Execute()
```

## Contributing

Please see [Contributing](CONTRIBUTING.md) for guidelines and instructions for local development.

## License

[MIT](LICENSE)
