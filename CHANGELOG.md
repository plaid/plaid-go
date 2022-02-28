This library is generated from an OpenAPI schema (OAS). See full changelog [here](https://github.com/plaid/plaid-openapi/blob/master/CHANGELOG.md) for schema changes.

# 2.1.0
- Updating to OAS 2020-09-14_1.77.0

## Breaking changes
- Many fields that were previously `float32` values are now `float64` values.

## OpenAPI Schema Changes
### 2020-09-14_1.77.0
- Explicitly set `format: double` for non-integer numbers so generated fields prefer float64

### 2020-09-14_1.76.0
- Add three new endpoints for Assets: `/asset_report/relay/create`, `/asset_report/relay/get`, and `/asset_report/relay/rmeove`

### 2020-09-14_1.75.0
- Added `/asset_report/relay/refresh` endpoint

### 2020-09-14_1.74.0
- Add `recurring_transactions` to list of products

### 2020-09-14_1.73.0
- Add new endpoint for `/credit/bank_income/get`

### 2020-09-14_1.72.0
- Updated documentation URLs for all product endpoints. They can now be found
at `/docs/api/products/<product-name>/#endpoint` instead of `/docs/api/products/#endpoint`

### 2020-09-14_1.71.0
- internal changes

### 2020-09-14_1.70.0
- Remove deprecated `income_verification_id` from `/sandbox/income/fire_webhook`

### 2020-09-14_1.69.1
- Reorder processors enum

### 2020-09-14_1.69.0
- Added `/beta/transactions/v1/enhance` endpoint

### 2020-09-14_1.68.1
- Added `status` object to sample responses for `/institutions/get` and `institutions/search` endpoints

### 2020-09-14_1.68.0
- Mark `include_personal_finance_category_beta` property as deprecated.
- Add new argument `include_personal_finance_category` to TransactionsGetRequestOptions.
- Update docs for `/transactions/get` request and response, referencing personal_finance_category taxonomy csv file.

### 2020-09-14_1.67.1
- internal changes

### 2020-09-14_1.67.0
- Removed unused `/income/verification/summary/get` endpoint

### 2020-09-14_1.66.0
- Added Payment Consent endpoints

### 2020-09-14_1.65.0
- Removed unused `/income/verification/paystub/get` endpoint

### 2020-09-14_1.64.15
- De-anonymized enums:
  - `PaymentInitiationPaymentReverseResponse.properties.status` => `PaymentInitiationRefundStatus`
  - `PaymentInitiationPaymentCreateResponse.properties.status` => `PaymentInitiationPaymentCreateStatus`
  - `PaymentInitiationRefund.properties.status` => `PaymentInitiationRefundStatus`
  - `PaymentAmount.properties.currency` => `PaymentAmountCurrency`
  - `InvestmentTransaction.properties.type` => `InvestmentTransactionType`
  - `InvestmentTransaction.properties.subtype` => `InvestmentTransactionSubtype`
  - `TransferAuthorizationDecisionRationale.properties.code` => `TransferAuthorizationDecisionRationaleCode`
  - `TransferAuthorizationGuaranteeDecisionRationale.properties.code` => `TransferAuthorizationGuaranteeDecisionRationaleCode`
  - `TransferAuthorization.properties.decision` => `TransferAuthorizationDecision`
  - `TransferEventListRequest.properties.transfer_type` => `TransferEventListTransferType`
  - `BankTransferEventListRequest.properties.bank_transfer_type` => `BankTransferEventListBankTransferType`
  - `BankTransferEventListRequest.properties.direction` => `BankTransferEventListDirection`
  - `TransferIntentCreate.properties.status` => `TransferIntentStatus`
  - `TransferIntentGet.properties.status` => `TransferIntentStatus`
  - `TransferIntentGet.properties.authorization_decision` => `TransferIntentAuthorizationDecision`
- `IncomeVerificationPrecheckMilitaryInfo.properties.branch` is now a string field (previously enum)

### 2020-09-14_1.64.15
- Made `last_statement_balance` and `minimum_payment_amount` `nullable` for credit card liabilities schema to reflect existing API behavior.

# 2.0.1
- Updating to OAS 2020-09-14_1.64.14
- Fix issue with nullable strings incorrectly reporting if they are set.

## OpenAPI Schema Changes
### 2020-09-14_1.64.14
- Made `last_payment_amount` and `last_statement_issue_date` `nullable` for credit card liabilities schema to reflect existing API behavior.
- Fix transfers examples to reflect more consistent usage of `region` field.

### 2020-09-14_1.64.13
- Deprecate `idempotency_key` parameter in transfer/create

### 2020-09-14_1.64.12
- Removed the unused `required_product_access` and `optional_product_access` parameters from `RequestedScopes`

### 2020-09-14_1.64.11
- Fix some examples that were not consistent with their schemas
- Add `adjustments` as an investments transaction type to make OpenAPI file consistent with values returned by the API
- Clarify description field for `marital_status` to reflect possible values

### 2020-09-14_1.64.10
- Updated the external docs URL for Bank Transfers sandbox endpoints

# 2.0.0
- Updating to OAS 2020-09-14_1.64.9

## Breaking changes
- `InstitutionsGetRequestOptions.Oauth` went from a `PtrBool` -> `NullableBool` as it's now nullable.
- `InstitutionsSearchRequestOptions` fields changed to nullable (so `NullableBool`):
   - `.Oauth`
   - `.IncludeAuthMetadata`
   - `.IncludePaymentInitiationMetadata`
- To migrate from `PtrBool` -> `NullableBool`, set the value to a `*plaid.NewNullableBool(plaid.PtrBool(bool))`. For an example, check out how `Oauth` is set in `/tests/institutions_test.go`.
- Split `AccountSubtype` enums and objects into per-endpoint ones to more accurately represent allowed values.
  - Fixed a bug with `LinkTokenCreateRequestAccountSubtypes` having unconditional maps as parameters instead of models
  - `DepositoryFilter`'s account subtypes uses `DepositoryAccountSubtypes`
  - `CreditFilter`'s account subtypes uses `CreditAccountSubtypes`
  - `LoanFilter`'s account subtypes uses `LoanAccountSubtypes`
  - `InvestmentFilter`'s account subtypes uses `InvestmentAccountSubtypes`
## OpenAPI Schema Changes
### 2020-09-14_1.64.9
- De-anonymized the object filters under `LinkTokenCreateRequestAccountSubtypes`, as anonymous objects aren't compatible with the generated CLibs.

### 2020-09-14_1.64.8
- Updated the description of the historical_balances array

### 2020-09-14_1.64.7
- Add new possible enums for income verification earnings breakdown canonical description

### 2020-09-14_1.64.6
- Hid a few product enum values that are deprecated or no longer valid for certain request fields. This affects the documentation only.

### 2020-09-14_1.64.5
- Make guarantee fields required in Transfer endpoints

### 2020-09-14_1.64.4
- Updated description for `failure_reason` field in Transfer endpoints

### 2020-09-14_1.64.3
- Make `repayment_id` required in `/transfer/repayment/return/list` endpoint

### 2020-09-14_1.64.2
- Update description for legal name field in `BankTransferUser` 

### 2020-09-14_1.64.1
- Update descriptions for `/transfer/repayment/list` and `/transfer/repayment/return/list` endpoints

### 2020-09-14_1.64.0
- Remove `scheme_automatic_downgrade` from `/payment_initiation/payment/create`

### 2020-09-14_1.63.1
- Update description for `/sandbox/transfer/sweep/simulate` endpoint

### 2020-09-14_1.63.0
- Refactor account subtype enums for greater specificity. This has no changes to the API but is a major semver change for Python, Node, Go, and Java client library interfaces to the AccountSubtype object within account filtering contexts in `/link/token/create`. The `AccountSubtype` namespace in this context is now prefixed with the AccountType. (Example for Node: Old: `AccountSubtype.checking` New: `DepositoryAccountSubtype.checking`)

### 2020-09-14_1.62.7
- Update description for `datetime` and `authorized_datetime` fields in Transactions endpoints

### 2020-09-14_1.62.6
- Make `sweep_id` / `sweep_amount` fields on Transfer Event nullable

### 2020-09-14_1.62.6
- Set `institution_status` to be nullable in `InstitutionsGetResponse`

### 2020-09-14_1.62.5
- Update external docs URLs for Transfer and Bank Transfer endpoints
- Update description for `ach_return_code` field in Transfer endpoints

### 2020-09-14_1.62.4
- Add `join_date` to `/application/get` and `/item/application/list`
- Remove `created_at` from `/application/get`

### 2020-09-14_1.62.3
- Updated various description fields for Income

### 2020-09-14_1.62.2
- Add `employment` as an available product in Product array.

### 2020-09-14_1.62.1
- Add `minItems` and `minLength` validation to various fields in `/institution/*` request schemas

### 2020-09-14_1.62.0
- Add guarantee_decision and guarantee_decision rationale fields to the transfer API
- Add repayment-related resources to the transfer API

### 2020-09-14_1.61.7
- Remove `receiver_pending` and `receiver_posted` from bank transfer event types.
- Remove `BankTransferReceiverDetails` from bank transfer event types.

### 2020-09-14_1.61.6
- Update description formatting for `sweep` and `amount` fields for sweep endpoints

### 2020-09-14_1.61.5
- Added `NEW_ACCOUNTS_AVAILABLE` webhook code as valid input to `/sandbox/item/fire_webhook`
- Update description for `/sandbox/item/fire_webhook`

### 2020-09-14_1.61.4
- Set the `minimum` for the `count` and `offset` fields in `InstitutionsGetRequest`
- Set `products`, `routing_numbers`, and `oauth` fields to be nullable in `InstitutionsGetRequestOptions`
- Set `products` to be nullable in `InstitutionsSearchRequest`
- Set `oauth`, `include_auth_metadata`, and `include_payment_initiation_metadata` fields to be nullable in `InstitutionsSearchRequestOptions`
- Set `payment_id` field to be nullable in `InstitutionsSearchPaymentInitiationOptions`

### 2020-09-14_1.61.3
- Adds `DOCUMENT_TYPE_NONE` enum value for document metadata

### 2020-09-14_1.61.2
- Relax length restrictions on the `currency` field in the `Pay` schema

### 2020-09-14_1.61.1
- Use new payment statuses in `PaymentStatusUpdateWebhook`

# 1.10.0
- Updating to OAS 2020-09-14_1.61.0

# 1.9.0
- Updating to OAS 2020-09-14_1.58.1

# 1.8.0
- Updating to OAS 2020-09-14_1.54.0

# 1.7.0
- Updating to OAS 2020-09-14_1.44.0

# 1.6.0
- Updating to OAS 2020-09-14_1.40.3

# 1.5.0
- Updating to OAS 2020-09-14_1.36.1

# 1.4.0
- Updating to OAS 2020-09-14_1.36.1

# 1.3.0
- Updating to OAS 2020-09-14_1.33.0.

# 1.2.0
- Updated to OAS 2020-09-14_1.31.1.

# 1.1.0
- Updated to OAS 2020-09-14_1.26.1.

# 1.0.0
- Moved officially to `GOMODULES`, as well as GA'd the generated version of `plaid-go`! This is ported over from our beta branch / release.
- Pinned to OpenAPI version `2020-09-14_1.21.0`. Make sure to check the OpenAPI changelog.

# 0.1.0-beta-2
- Introduce the `ToPlaidError` helper function to convert generic errors to the plaid.Error struct
- Type fixes from `OpenAPI version 2020-09-14_1.20.6`, see full changelog [here](https://github.com/plaid/plaid-openapi/blob/master/CHANGELOG.md#2020-09-14_11912)

# 0.1.0-beta-1
We first are reversioning this package, as we are now using GOMODULES.

This version represents a transition in how we maintain our external client libraries. We are now using an [API spec](https://github.com/plaid/plaid-openapi) written in `OpenAPI 3.0.0` and running our definition file through [OpenAPITool's `go` generator](https://github.com/OpenAPITools/openapi-generator).

**Go Migration Guide:**

## Client initialization
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

## Endpoints
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

## Errors
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

# 7.3.0
- Add support for `options` to `/payment_initiation/payment/create`

# 7.2.0
- Add support for `last_updated_datetime` to `/accounts/balance/get`
- Add Standing Orders support to Payment Initiation

# 7.1.0
- Add institution status types for health incidents and investment updates

# 7.0.0

- Add back (deprecated) `/payment_initiation/payment/token/create` endpoint
- Add back (deprecated) `/item/public_token/create` endpoint

BREAKING CHANGES:

- Add Standing Order support to the `/payment_initiation/payment/create` endpoint

# 6.0.0

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

# 5.1.0

- Add support for Link Token get endpoint ([#142](https://github.com/plaid/plaid-go/pull/142))
  - `/link/token/get`

# 5.0.0

BREAKING CHANGES:

- Add BACS as a parameter to `/recipient/create` ([#137](https://github.com/plaid/plaid-go/pull/137))

# 4.0.0

- Add `MerchantName` to `Transaction` struct

BREAKING CHANGES:

- Removes the public key as input to `ClientOptions`. The public key is no longer needed by the API.
- Add support for the `/link/token/create` endpoint

# 3.1.0

- `AuthorizedDate` and `PaymentChannel` added to the `Transaction` struct
- `Item` added to the `GetAuthResponse` struct

# 3.0.0

- Adds support for `/sandbox/item/set_verification_status`
- `PaymentRecipientAddress` can now be null
- Removed support for deprecated `/item/access_token/update_version` endpoint

BREAKING CHANGES:

- Removed `client.UpdateAccessTokenVersion`
- `PaymentRecipientAddress` can now be null
