# plaid-go [![CircleCI](https://circleci.com/gh/plaid/plaid-go.svg?style=svg)](https://circleci.com/gh/plaid/plaid-go) [![GoDoc](https://godoc.org/github.com/plaid/plaid-go?status.svg)](https://godoc.org/github.com/plaid/plaid-go/plaid)

A Go client library for the [Plaid API](https://plaid.com/docs).

## Table of Contents

- [plaid-go](#plaid-go)
    * [Install](#install)
    * [Versioning](#versioning)
    * [Documentation](#documentation)
    * [Getting Started](#getting-started)
    * [Developing](#developing)
    * [Examples](#examples)
        * [Payment Initiation](#payment-initiation)
    * [Support](#support)
    * [License](#license)

## Install

```console
$ go get github.com/plaid/plaid-go
```

## Versioning

Each major version of `plaid-go` targets a specific version of the Plaid API:

| API version | plaid-go release |
| ----------- | ------------------ |
| [`2020-09-14`][api-version-2020-09-14] (**latest**) | `7.x.x` |
| [`2019-05-29`][api-version-2019-05-29] (**latest**) | `2.x.x` |
| [`2018-05-22`][api-version-2018-05-22] | `1.x.x` |
| `2017-03-08` | not supported |

For information about what has changed between versions and how to update your integration, head to the [version changelog][version-changelog].

## Documentation

The module supports all Plaid API endpoints.

GoDoc: [![GoDoc](https://godoc.org/github.com/plaid/plaid-go?status.svg)](https://godoc.org/github.com/plaid/plaid-go/plaid)

## Getting Started

### Calling Endpoints

To call an endpoint you must create a `Client` object.

```go
import (
    "net/http"
    "os"

    "github.com/plaid/plaid-go/plaid"
)

clientOptions := plaid.ClientOptions{
    os.Getenv("PLAID_CLIENT_ID"),
    os.Getenv("PLAID_SECRET"),
    plaid.Sandbox, // Available environments are Sandbox, Development, and Production
    &http.Client{}, // This parameter is optional
}
client, err := plaid.NewClient(clientOptions)
```

Each endpoint returns an object which contains the parsed JSON from the HTTP response.

### Errors

All non-200 responses will return a plaid.Error instance.

For more information on Plaid response codes, head to the [docs](https://plaid.com/docs/errors/).

## Developing

1. Download this repo into your Go source directory
2. Run `make setup` pull down all dependencies etc

### Tests

To run the tests you'll need to sign up for a Sandbox account on https://plaid.com as they perform real API requests.

Once you have these you can run `make test`, passing your Sandbox credentials as environment variables:

```shell
PLAID_CLIENT_ID=aabbcc PLAID_PUBLIC_KEY=ddeeff PLAID_SECRET=ffeedd make test
```

## Examples

### Payment Initiation
For more information about this product, head to the [Payment Initiation docs](https://plaid.com/docs/payment-initiation/).

#### Create payment recipient using BACS with no IBAN or address
```go
paymentRecipientCreateResp, err := client.CreatePaymentRecipient("John Doe",
plaid.OptionalRecipientCreateParams{
    BACS: &plaid.PaymentRecipientBacs{
        Account:  "26207729",
        SortCode: "560029",
    },
})
recipientID := paymentRecipientCreateResp.RecipientID
```

#### Create payment
```go
paymentCreateResp, err := client.CreatePayment(recipientID, "TestPayment", plaid.PaymentAmount{
    Currency: "GBP",
    Value:    100.0,
}, nil)
paymentID := paymentCreateResp.PaymentID
paymentStatus := paymentCreateResp.Status
```

#### Create Link Token (for Payment Initiation only)
```go
linkTokenResp, err := client.CreateLinkToken(plaid.LinkTokenConfigs{
    User: &plaid.LinkTokenUser{
        ClientUserID: "123-test-user-id",
    },
    ClientName:        "Plaid Test App",
    Products:          []string{"payment_initiation"},
    CountryCodes:      []string{"GB"},
    Language:          "en",
    Webhook:           "https://webhook-uri.com",
    PaymentInitiation: &plaid.PaymentInitiation{PaymentID: paymentID},
})
linkToken := linkTokenResp.LinkToken
```

## Support

Open an [issue](https://github.com/plaid/plaid-go/issues/new)!

## License

[MIT](https://github.com/plaid/plaid-go/blob/master/LICENSE)

[version-changelog]: https://plaid.com/docs/api/versioning/
[api-version-2018-05-22]: https://plaid.com/docs/api/versioning/#2018-05-22
[api-version-2019-05-29]: https://plaid.com/docs/api/versioning/#2019-05-29
[api-version-2020-09-14]: https://plaid.com/docs/api/versioning/#2020-09-14
