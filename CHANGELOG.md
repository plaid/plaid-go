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
