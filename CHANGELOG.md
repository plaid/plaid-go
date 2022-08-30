This library is generated from an OpenAPI schema (OAS). See full changelog [here](https://github.com/plaid/plaid-openapi/blob/master/CHANGELOG.md) for schema changes.

# 8.2.1
- bump golang.org/x/oauth version
- bump github.com/stretchr/testify version
- bump (indirect dependency) golang.org/x/net version
- bump (indirect dependency) golang.org/x/crypto => golang.org/x/crypto version

# 8.2.0
- Updating to OAS 2020-09-14_1.161.5

## OpenAPI Schema Changes
### 2020-09-14_1.161.5

- Reverts the changes made in 2020-09-14_1.157.0

### 2020-09-14_1.161.4

- Update description for webhook `USER_PERMISSION_REVOKED`

### 2020-09-14_1.161.3
 
- Bug fix: Put quotes around the '529' account type to prevent generated client libraries from treating it as an integer.
- Add `RECURRING_TRANSACTIONS_UPDATE` to description of sandbox webhook testing endpoint

### 2020-09-14_1.161.2

- Change the Item schema to refer to `PlaidError` rather than `Error` to avoid namespace conflicts in client libraries.

### 2020-09-14_1.161.1

- Add required fields back to AssetReport schema

### 2020-09-14_1.161.0
- Remove `access_token` request parameter from `/beta/partner/v1/customers/create`

### 2020-09-14_1.160.0

- Add `user` fields to the `/link/token/create` `user` object.

### 2020-09-14_1.159.0

- Add `RECURRING_TRANSACTIONS_UPDATE` webhook specification
- Fix typo in `/sandbox/item/fire_webhook` description

### 2020-09-14_1.158.1

- Add x-plaid-validation for Credit Relay Tokens

### 2020-09-14_1.158.0

- Add `address` and `id_number` fields to the `/link/token/create` `user` object.

### 2020-09-14_1.157.0

- Update and add schemas referred by `AssetReportGetResponse`

### 2020-09-14_1.156.0

- Add `payment_profile_id` to `/transfer/authorization/create` and `/transfer/create`
- Make `access_token` and `account_id` optional in `/transfer/authorization/create` and `/transfer/create`

### 2020-09-14_1.155.0

- Add `payment_profile_id` as a field under `transfer` for `/link/token/create`

### 2020-09-14_1.154.0

- Create `/credit/bank_income/pdf/get` to allow customers to retrieve the bank income product as a PDF

### 2020-09-14_1.153.1
-- Add `personal_finance_category_icon_url` to `/beta/transactions/v1/enhance` response

### 2020-09-14_1.153.0
-- Add new endpoint `/beta/partner/v1/customers/create`

### 2020-09-14_1.152.0
- Add `form1099s` as part of the `credit/payroll_income/get` response.
- Add `Credit1099` object and corresponding subtypes.

### 2020-09-14_1.151.1
- Update required fields list for guaranteed ACH customers

### 2020-09-14_1.151.0
- Mark `verification` field under `paystubs` object in `/credit/payroll_income/get` as deprecated

### 2020-09-14_1.150.0
- Make `user_present` under `/transfer/authorization/create` nullable

### 2020-09-14_1.149.1
- Update fields description for guaranteed ACH customers

### 2020-09-14_1.149.0

- Add `gave_consent` as a field under `identity_verification` for `/link/token/create`
- Remove `identity_verification.consent`, which is deprecated, from documentation for `/link/token/create`

### 2020-09-14_1.148.0

- Add `user_present` a an optional field under `/transfer/authorization/create`, update fields description for guaranteed ACH customers

### 2020-09-14_1.147.1

- Remove deprecated `reversed` status from Transfer schema
- Remove deprecated `reverse_swept` status from Transfer Event schema

### 2020-09-14_1.147.0

- Add new endpoints `/credit/relay/get`, `/credit/relay/refresh` and `/credit/relay/remove`

# 8.1.0
- Updating to OAS 2020-09-14_1.146.0

## OpenAPI Schema Changes
### 2020-09-14_1.146.0

- make item_logins field not required

### 2020-09-14_1.145.0

- Make changes to IdentityMatchRequest to support options data

### 2020-09-14_1.144.0

- Remove unsupported error_code (`PRODUCTS_NOT_SUPPORTED`) for the `force_error` config in Sandbox

### 2020-09-14_1.143.0

- Make changes to support crypto in Investments product

### 2020-09-14_1.142.0

- Add new endpoints `/payment_profile/get` and `/payment_profile/remove`

### 2020-09-14_1.141.1

- Bug fix: add `identity_verification` to `Products` array

### 2020-09-14_1.141.0

- Add a new endpoint `/payment_profile/create`

### 2020-09-14_1.140.1

- Add `stated_income_sources` as a field under `income_verification` for `/link/token/create`

### 2020-09-14_1.140.0

- Add `bin` as a field under `institution_metadata` for `/link/token/create`

### 2020-09-14_1.139.0

- Add `international` IBAN numbers to `WalletNumbers` field in `/wallet/get` and `/wallet/list` responses.
- Wrap the `iban` field in the `WalletTransactionCounterpartyNumbers` under new field `international`.

### 2020-09-14_1.138.0

- Add new `identity/match` endpoint, and `IdentityMatchRequest` and `IdentityMatchResponse`

### 2020-09-14_1.137.1

- Update comment for the `recipient.name` field with recommendation to avoid long strings and special characters.

### 2020-09-14_1.137.0

- Reverts the changes made in 2020-09-14_1.135.0. ProductAccess and AccountProductAccess once again have all fields explicitly defined.

### 2020-09-14_1.136.0

- Add `products` in `asset_report/create` endpoint.

### 2020-09-14_1.135.0

- Convert ProductAccess and AccountProductAccess to optional.Map

### 2020-09-14_1.134.0

- Rename `merchant_website` and `merchant_logo_url` in `/beta/transactions/v1/enhance` to `website` and `logo_url`

### 2020-09-14_1.133.0

- Update `/credit/payroll_income/get` response to have a `pull_id` instead of an `income_report_token` and add `pull_id` to `/credit/employment/get` response

### 2020-09-14_1.132.1

- Made several fields nullable for `/signal/evaluate`

### 2020-09-14_1.132.0

- Add `/credit/audit_copy_token/remove` to invalidate Audit Copy tokens

# 8.0.0

- Note: we're updating the version from v3.5.0 to v8.0.0 because of this issue [here](https://github.com/plaid/plaid-go/issues/235). This should help make this package's versioning a lot more friendly for dependabot and other tools that use tags for version comparison. Going forward, we should be following the typical versioning system and not skip any other major versions.
- Updating to OAS 2020-09-14_1.131.3

## OpenAPI Schema Changes
### 2020-09-14_1.131.3
- Update external URL for `/transfer/migrate_account` endpoint

### 2020-09-14_1.131.2
- Add `TRANSFER_LIMIT_REACHED` to enum `TransferAuthorizationDecisionRationaleCode`

### 2020-09-14_1.131.1
- Update `/payment_initiation/consent/*` external docs URLs

### 2020-09-14_1.131.0
- Add `iban` to `counterparty`'s `numbers` field in `/wallet/transaction/execute`

### 2020-09-14_1.130.0
- Update description of `/user/create` to make it more clear what would happen when a client calls the user creation endpoint with the same client_user_id multiple times.

### 2020-09-14_1.129.0
- Add `signal_description` for each risk signal in `/beta/credit/payroll_income/risk_signals/get`
- Change `DocumentRiskSignalInstitutionMetadata` to be nullable

### 2020-09-14_1.128.6
- Fix category rules formatting

### 2020-09-14_1.128.5
- Change description in `UserCustomPassword` to reflect that only top level fields are optional not all fields

### 2020-09-14_1.128.4
- Move Category Rules description to `include_personal_finance_category` flag

# 3.5.0
- Updating to OAS 2020-09-14_1.128.4

## OpenAPI Schema Changes
### 2020-09-14_1.128.4
- Add `identity_verification` as a new optional parameter for `/link/token/create`

### 2020-09-14_1.128.3
- Add category rules beta to `personal_finance_category` field description
in `/transactions/get`

### 2020-09-14_1.128.2
- Update description of `/investments/transactions/get` endpoint
- Update description of `cost_basis` field

### 2020-09-14_1.128.1
- Add `merchant_website` and `merchant_logo_url` to `/beta/transactions/v1/enhance` response

### 2020-09-14_1.128.0
- Add `auth_type_select_enabled`, `automated_microdeposits_enabled`, `instant_match_enabled`, and `same_day_microdeposits_enabled` to `/link/token/create`
- Marked `flow_type` parameter as deprecated

### 2020-09-14_1.127.0
- Add `consent_id` filter to `/payment_initiation/payment/list`
- Add `consent_id` to `/link/token/create`

### 2020-09-14_1.126.0

- Added `/credit/relay/create` endpoint

### 2020-09-14_1.125.0

- Add `accounts` object to `/credit/payroll_income/get` response

### 2020-09-14_1.124.0
- Added `include_fast_report` in `asset_report/create` endpoint. Added `fast_report` in `asset_report/get` endpoint

### 2020-09-14_1.123.1
- Removed `reversed` and `reverse_swept` from possible event_type values in TransferEventListRequest docs

### 2020-09-14_1.123.0

- Change `tokens` to `report_tokens` in `/credit/audit_copy_token/create` request

### 2020-09-14_1.122.0

- Add `request_id` to all Identity Verification and Monitor responses

### 2020-09-14_1.121.2
- Added sample user_id to INCOME_VERIFICATION webhook

### 2020-09-14_1.121.1
- Pre-release API refinements to Monitor and Flow endpoints

# 3.4.0
- Updating to OAS 2020-09-14_1.121.0

## OpenAPI Schema Changes
### 2020-09-14_1.121.0
- Add TimestampNullable type

### 2020-09-14_1.120.0
- Removed `Uploaded`, `Created` and `APPROVAL_STATUS_APPROVED` enum strings from `PayrollItemStatus` field.

### 2020-09-14_1.119.0
- Add `returned` to TransferStatus enum
- Add `return_swept` to TransferSweepStatus enum
- Add `returned` and `return_swept` to TransferEventType enum

### 2020-09-14_1.118.0
- Added `employee_type` and `last_paystub_date` to `/credit/employment/get` response

### 2020-09-14_1.117.1
- Make `/payment_initiation/consent/create` API more strict

### 2020-09-14_1.117.0
- Add `/credit/audit_copy_token/create` endpoint

### 2020-09-14_1.116.0
- Add `/wallet/list` endpoint

### 2020-09-14_1.115.2
- Update description fields to fix formatting errors
- Reflect that `error.suggested_action` is `nullable`

### 2020-09-14_1.115.1
- Update OpenAPI spec

### 2020-09-14_1.115.0
- Added `income_report_token` to `/credit/payroll_income/get` response

### 2020-09-14_1.114.2
- Add `/wallet/create` endpoint

### 2020-09-14_1.114.1
- Add beta `additional_consented_products` field to `/link/token/create`

### 2020-09-14_1.113.1
- Updated `/transactions/recurring/get` description

### 2020-09-14_1.113.0
- Add webhooks for new Monitor and Identity Verification products

### 2020-09-14_1.112.0
- Add endpoints for new Monitor and Identity Verification products

### 2020-09-14_1.111.15
- Remove `emi_recipient_id` from Payment Initiation Recipient

### 2020-09-14_1.111.14
- Add optional `iban` and `bacs` fields to `options` in the `/payment_initiation/consent/create` request

### 2020-09-14_1.111.13
- Updated `/transactions/sync` description

### 2020-09-14_1.111.12
- Add more accurate enum documentations to `/transactions/recurring/get` API doc

### 2020-09-14_1.111.11
- Additional documentation for `/transactions/sync`

### 2020-09-14_1.111.9
- Remove deprecated field `createdAt` from `/application/get` response

### 2020-09-14_1.111.8
- Add field validation to `BankTransferDirection`

### 2020-09-14_1.111.7
- Remove deprecated field `createdAt` from `/application/get` response

### 2020-09-14_1.111.6
- Add external doc link to `transactions/recurring/get`

### 2020-09-14_1.111.5
- Updating the API doc for Recurring Transactions

### 2020-09-14_1.111.4
- Add `DisplayName` in `/application/get` response

### 2020-09-14_1.111.3
- Updated sample responses for all Transfer endpoints

### 2020-09-14_1.111.2
- Changing `beta/transactions/rules/` routes to `beta/transactions/rules/v1`

### 2020-09-14_1.111.1
- Fixing `InsitututionMetadata` typo to `InstitutionMetadata` in private `/beta/credit/payroll_income/risk_signals/get`endpoint response

### 2020-09-14_1.111.0
- Added `require_guarantee`, `guarantee_decision`, and `guarantee_decision_rationale` to `/transfer/intent` in order to support Guarantee when using Transfer UI.

### 2020-09-14_1.110.2
- Add additional supported `type` enums in `WalletTransaction`.

### 2020-09-14_1.110.1
- Add Additional History billing information for /asset_report/create.

### 2020-09-14_1.110.0
- Add `user_id` to income verification webhook payload

### 2020-09-14_1.109.1
- Make `consent_id` field nullable in `PaymentInitiationPayment`.

### 2020-09-14_1.109.0
- Replace `initiated_refunds` with `refund_ids` in the `/payment_initiation/payment/get` and `/payment_initiation/payment/list` responses

### 2020-09-14_1.108.0
- Added `/beta/credit/payroll_income/risk_signals/get` endpoint (currently private)

### 2020-09-14_1.107.4
- Remove unsupported ACH classes from `bank_transfer/` and `transfer/` endpoints.

### 2020-09-14_1.107.3
- Add `enable_multiple_items` parameter for bank income.

### 2020-09-14_1.107.2
- Fix typo in `institution_name` parameter for credit endpoints.

### 2020-09-14_1.107.0
- Added `reference` and `idempotency_key` fields to the `payment_initiation/payment/reverse` request.

### 2020-09-14_1.106.0
- Added `is_update_mode` to `income_verification` in the `/link/token/create` body

### 2020-09-14_1.105.2
- Consolidate item schemas

### 2020-09-14_1.105.1
- Removed `client_id` and `secret` as required fields from `/transfer/intent/{get,create}` to match actual API behavior.

### 2020-09-14_1.105.0
- Add `/credit/payroll_income/refresh` endpoint

### 2020-09-14_1.104.0
- Added `/signal/prepare`

### 2020-09-14_1.102.0
- Added `/watchlist_screening/individual/list` (currently private)
- Added `/watchlist_screening/individual/create` (currently private)
- Added `/watchlist_screening/individual/get` (currently private)
- Added `/watchlist_screening/individual/update` (currently private)
- Added `/watchlist_screening/individual/history/list` (currently private)
- Added `/watchlist_screening/individual/review/list` (currently private)
- Added `/watchlist_screening/individual/review/create` (currently private)
- Added `/watchlist_screening/individual/hit/list` (currently private)
- Added `/watchlist_screening/entity/list` (currently private)
- Added `/watchlist_screening/entity/create` (currently private)
- Added `/watchlist_screening/entity/get` (currently private)
- Added `/watchlist_screening/entity/update` (currently private)
- Added `/watchlist_screening/entity/history/list` (currently private)
- Added `/watchlist_screening/entity/hit/list` (currently private)
- Added `/watchlist_screening/entity/review/list` (currently private)
- Added `/watchlist_screening/entity/review/create` (currently private)
- Added `/watchlist_screening/individual/program/list` (currently private)
- Added `/watchlist_screening/individual/program/get` (currently private)
- Added `/watchlist_screening/entity/program/list` (currently private)
- Added `/watchlist_screening/entity/program/get` (currently private)
- Added `/dashboard_user/list` (currently private)
- Added `/dashboard_user/get` (currently private)
- Added `/identity_verification/list` (currently private)
- Added `/identity_verification/get` (currently private)
- Added `/identity_verification/retry` (currently private)
- Modified `/identity_verification/create` (currently private)

### 2020-09-14_1.101.0
- Add endpoint for `/credit/bank_income/refresh`

### 2020-09-14_1.100.0
- Add `include_original_description`, `include_personal_finance_category` options to `/transactions/sync` request.

### 2020-09-14_1.99.0
- API changes for /credit/employment/get

### 2020-09-14_1.98.1
- Add `gusto` as processor partner

### 2020-09-14_1.98.0
- Add `user_token` as a request parameter for `/sandbox/public_token/create`

### 2020-09-14_1.97.1
- Remove `auth`, `transactions_updates`, `investments_updates`, and `identity` as required fields from Item status to match actual API behavior.

# 3.3.0
- Updating to OAS 2020-09-14_1.97.0

## OpenAPI Schema Changes
### 2020-09-14_1.97.0
- [Beta] Rename some `Credit` refs that were preventing client library generation from completing successfully

### 2020-09-14_1.96.0
- remove unused `payroll_income_id` from `/credit/payroll_income/get` field
- add status object to items in `/credit/payroll_income/get` response body

### 2020-09-14_1.95.1
- Add `TransferEventsUpdateWebhook` schema

### 2020-09-14_1.95.0
- Add `institution_data` parameter to `/link/token/create`

### 2020-09-14_1.94.2
- Tidy up YAML

### 2020-09-14_1.94.1
- Add `highnote` processor to `/processor/token/create`

# 3.2.0
- Updating to OAS 2020-09-14_1.94.0

## OpenAPI Schema Changes
### 2020-09-14_1.94.0
- Add `use_case`, `company_legal_name`, `city`, `region`, `country_code`, `postal_code` as a required response field of `Application`

### 2020-09-14_1.93.2
- Remove `income_verification_id` from income webhook example
- Fix incorrect URL for `/user/create` endpoint

### 2020-09-14_1.93.1
- Remove deprecated `income_verification_id` from income webhooks
- Standardize income webhook casing

### 2020-09-14_1.93.0
- Add several new fields to `/signal/evaluate` response
- 
### 2020-09-14_1.92.4
- Add `/sandbox/transfer/fire_webhook` endpoint

### 2020-09-14_1.91.4
- Mark certain Income endpoints as deprecated in favor of the new `/credit/*` endpoints.

### 2020-09-14_1.91.3
- Add `checkout` processor to `/processor/token/create`

### 2020-09-14_1.91.2
- Add `webhook_type` parameter to `/sandbox/item/fire_webhook`
- Support for investments transactions, investments holdings and liabilities `DEFAULT_UPDATE` webhooks

### 2020-09-14_1.90.2
- Add new warning type to `/credit/bank_income/get` response

### 2020-09-14_1.90.1
- Add `marqeta` and `solid` as Auth processor partners
- Fix schema of `cause` parameter for Asset Reports
- Fix some invalid examples
 
### 2020-09-14_1.90.0
- Add `/credit/employment/get` endpoint
- Add optional `access_tokens` array to `/credit/payroll_income/precheck` request

### 2020-09-14_1.89.3
- Update description of `/sandbox/item/fire_webhook`

### 2020-09-14_1.89.2
- Update description of `accounts/get`

### 2020-09-14_1.89.1
- Added `AUTH_DATA_UPDATE` webhook code as valid input to `/sandbox/item/fire_webhook`
- Update description for `/sandbox/item/fire_webhook`

### 2020-09-14_1.89.0
- Add `/transfer/migrate_account` endpoint

### 2020-09-14_1.88.2
- Fix operationId for `/credit/payroll_income/precheck`

### 2020-09-14_1.88.1
- Remove deprecated fields from `/item/application/list`

### 2020-09-14_1.88.0
- Add `wire_routing_number` parameter to `/bank_transfer/migrate_account`

### 2020-09-14_1.87.1
- Specify minimum length of 1 for `description` on `TransferIntentCreateRequest`

### 2020-09-14_1.87.0
- Add `consent_id` support in the Institutions Search request

### 2020-09-14_1.86.1
- Add `apex_clearing` as a processor partner

### 2020-09-14_1.86.0
- Introduce Credit Payroll Income APIs
- Introduce Credit Precheck API

### 2020-09-14_1.85.1
- Add `/identity_verification/create` endpoint, kept private for now

### 2020-09-14_1.85.0
- Add `status` field to `ConnectedApplication`

# 3.1.0
- Updating to OAS 2020-09-14_1.84.5

## OpenAPI Schema Changes
### 2020-09-14_1.84.5
- Added missing `asset_report_id` field to `/asset_report/relay/refresh`

### 2020-09-14_1.84.4
- Change summary description and url for `/credit/bank_income/get`

### 2020-09-14_1.84.3
- Slight wording change for `/credit/bank_income/get` response fields

### 2020-09-14_1.84.3
- Move `user_token` to top level of `link/token/create` request 

### 2020-09-14_1.84.2
- Correct typo in enum value for Investment subtypes (`person` -> `pension`)

### 2020-09-14_1.84.1
- Fix schema to properly handle personal finance categories in `/transactions/get`

### 2020-09-14_1.84.0
- Add `user_token` parameter to `link/token/create`

### 2020-09-14_1.83.1
- Add new fields to `/credit/bank_income/get` response

### 2020-09-14_1.83.0
- Remove `permitted` decision for `/transfer/authorization/create`

### 2020-09-14_1.82.0
- Add beta field `consented_products` to `/item/get/` endpoint response

### 2020-09-14_1.82.0
- Revamp LinkTokenCreate.IncomeVerificationOptions for GA

### 2020-09-14_1.81.0
- Add `/transaction/rules/create`, `/transaction/rules/list` and `/transaction/rules/remove` endpoints

### 2020-09-14_1.80.0
- Added `/user/create` endpoint

# 3.0.0
- Updating to OAS 2020-09-14_1.79.0

## Breaking changes
- Many enum fields have been de-anonymized and renamed, new names can be found in the `2020-09-14_1.64.15` change message
- Non integer numbers are now `float64` fields

## OpenAPI Schema Changes
### 2020-09-14_1.77.4
- Remove the word "Asset" before "Relay" in every asset report relay related responses and request objects

### 2020-09-14_1.77.3
- Add "AssetReport" at the beginning of relay related responses and request objects to match the same pattern as other assets related objects

### 2020-09-14_1.77.2
- Add `ProductAccess` fields for upcoming partner

# 2.2.0
- Updating to OAS 2020-09-14_1.78.0
- Go library 2.1.0 was erroneously released as a `minor` version; 2.1.0 changes will be re-released as a `major` version shortly.

## OpenAPI Schema Changes
### 2020-09-14_1.78.0
- Revert to `2020-09-14_1.64.14`

# 2.1.0
- Updating to OAS 2020-09-14_1.77.1

## OpenAPI Schema Changes
### 2020-09-14_1.77.1
- Fix extraneous field in enum that caused issue in code generation
- Added `asset_report_id` to the example for `/asset_report/relay/refresh`

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
