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
