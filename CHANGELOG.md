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
