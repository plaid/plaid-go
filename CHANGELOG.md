This library is generated from an OpenAPI schema (OAS). See full changelog [here](https://github.com/plaid/plaid-openapi/blob/master/CHANGELOG.md) for schema changes.

## 1.2.0
Updated to OAS version 2020-09-14_1.31.1.

## 1.1.0
Updated to OAS version 2020-09-14_1.26.1.

## 1.0.0
- Moved officially to `GOMODULES`, as well as GA'd the generated version of `plaid-go`! This is ported over from our beta branch / release.
- Pinned to OpenAPI version `2020-09-14_1.21.0`. Make sure to check the OpenAPI changelog.

## 0.1.0-beta-2
- Introduce the `ToPlaidError` helper function to convert generic errors to the plaid.Error struct
- Type fixes from `OpenAPI version 2020-09-14_1.20.6`, see full changelog [here](https://github.com/plaid/plaid-openapi/blob/master/CHANGELOG.md#2020-09-14_11912)

## 0.1.0-beta-1
We first are reversioning this package, as we are now using GOMODULES.

This version represents a transition in how we maintain our external client libraries. We are now using an [API spec](https://github.com/plaid/plaid-openapi) written in `OpenAPI 3.0.0` and running our definition file through [OpenAPITool's `go` generator](https://github.com/OpenAPITools/openapi-generator).

**Go Migration Guide:**

### Client initialization
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

### Endpoints
All endpoint rquests now take a request model and the functions have been renamed to move the
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

### Errors
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

## 7.3.0
- Add support for `options` to `/payment_initiation/payment/create`

## 7.2.0
- Add support for `last_updated_datetime` to `/accounts/balance/get`
- Add Standing Orders support to Payment Initiation

## 7.1.0
- Add institution status types for health incidents and investment updates

## 7.0.0

- Add back (deprecated) `/payment_initiation/payment/token/create` endpoint
- Add back (deprecated) `/item/public_token/create` endpoint

BREAKING CHANGES:

- Add Standing Order support to the `/payment_initiation/payment/create` endpoint

## 6.0.0

BREAKING CHANGES:

- The library has been pinned to the '2020-09-14' API release. Visit the [docs](https://plaid.com/docs/versioning/) to see what changed.
- the `/item/public_token/create` endpoint has been disabled in favor of the /link/token/create
    endpoint
- The `/item/add_token/create endpoint` has been disabled in favor of the /link/token/create
- The `/payment_initiation/payment/token/create` endpoint has been disabled in favor of the /link/token/create
    endpoint
- The `/item/remove` endpoint will no longer return a `removed` boolean.
- The `/institutions/get`, `/institutions/get_by_id`, and `/institutions/search` now require
    `country_codes` to be passed in.

## 5.1.0

- Add support for Link Token get endpoint ([#142](https://github.com/plaid/plaid-go/pull/142))
  - `/link/token/get`

## 5.0.0

BREAKING CHANGES:

- Add BACS as a parameter to `/recipient/create` ([#137](https://github.com/plaid/plaid-go/pull/137))

## 4.0.0

- Add `MerchantName` to `Transaction` struct

BREAKING CHANGES:

- Removes the public key as input to `ClientOptions`. The public key is no longer needed by the API.
- Add support for the `/link/token/create` endpoint

## 3.1.0

- `AuthorizedDate` and `PaymentChannel` added to the `Transaction` struct
- `Item` added to the `GetAuthResponse` struct

## 3.0.0

- Adds support for `/sandbox/item/set_verification_status`
- `PaymentRecipientAddress` can now be null
- Removed support for deprecated `/item/access_token/update_version` endpoint

BREAKING CHANGES:

- Removed `client.UpdateAccessTokenVersion`
- `PaymentRecipientAddress` can now be null
