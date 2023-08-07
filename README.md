# plaid-go [![CircleCI](https://circleci.com/gh/plaid/plaid-go.svg?style=svg)](https://circleci.com/gh/plaid/plaid-go) [![GoDoc](https://godoc.org/github.com/plaid/plaid-go?status.svg)](https://godoc.org/github.com/plaid/plaid-go/plaid)

The official Go client library for the [Plaid API](https://plaid.com/docs). The library is now generated from our [OpenAPI schema](https://github.com/plaid/plaid-openapi). For older manually-written versions of this client library, go [here](https://github.com/plaid/plaid-go/commit/0d3f02cddaa8fd637e84dccf2175a4a1a7dd0e07) for the latest non-generated version.

The latest version of the library supports only the latest version of the Plaid API (currently 2020-09-14). 

For more information about the Plaid API, including reference documentation, see the [Plaid API docs](https://www.plaid.com/docs/api).

## Table of Contents

- [plaid-go](#plaid-go)
  * [Install](#install)
  * [Getting Started](#getting-started)
    + [Calling Endpoints](#calling-endpoints)
    + [Errors](#errors)
  * [Authentication](#authentication)
  * [Examples](#examples)
  * [Contributing](#contributing)
  * [License](#license)


## Install

Library versions follow Semantic Versioning ("SemVer") and are formatted as `v1.0.0`. The plaid-go client library is typically updated on a monthly basis, and the canonical source for the latest version number is the [client library changelog](https://github.com/plaid/plaid-go/blob/master/CHANGELOG.md). As of `v1.0.0`, we've moved to support `GOMODULES`.

Edit your go.mod to include `github.com/plaid/plaid-go/v15 {VERSION}`

```console
$ go get github.com/plaid/plaid-go/v15@{VERSION}
```

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

Each endpoint will require an appropriate request model, and will return either the response model or an error.

### Dates

Dates and datetimes in requests, which are represented as strings in the API and in previous client library versions, are represented in this version of the Go client library as `time.Time` types. For more information, see the [Go documentation on the `time` package](https://pkg.go.dev/time).

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

## Authentication

First, you get your `client_id` and `secret` from your dashboard account. Authentication is handled by setting the `client_id` and `secret` on the configuration object.

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

    "github.com/plaid/plaid-go/v15/plaid"
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

## Contributing

Please see [Contributing](CONTRIBUTING.md) for guidelines and instructions for local development.

## License

[MIT](LICENSE)
