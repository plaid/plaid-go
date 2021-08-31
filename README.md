# plaid-go [![CircleCI](https://circleci.com/gh/plaid/plaid-go.svg?style=svg)](https://circleci.com/gh/plaid/plaid-go) [![GoDoc](https://godoc.org/github.com/plaid/plaid-go?status.svg)](https://godoc.org/github.com/plaid/plaid-go/plaid)

The official Go client library for the [Plaid API](https://plaid.com/docs). The library is now generated from our [OpenAPI schema](https://github.com/plaid/plaid-openapi) and is only compatible with the 2020-09-14 Plaid API version. The current library version contains breaking changes from the 7.x.x library version. For the previous version of this client library, also targeting Plaid API version 2020-09-14, [see version 7.3.0](https://github.com/plaid/plaid-go/releases/tag/7.3.0-archived).

## Table of Contents

- [plaid-go](#plaid-go)
  * [Install](#install)
  * [Getting Started](#getting-started)
    + [Calling Endpoints](#calling-endpoints)
    + [Errors](#errors)
  * [Authentication](#authentication)
  * [Contributing](#contributing)
  * [License](#license)

## Install

Versions look something like `v1.0.0`. As of `v1.0.0`, we've moved to support `GOMODULES`.

Edit your go.mod to include `github.com/plaid/plaid-go {VERSION}`

```console
$ go get github.com/plaid/plaid-go@{VERSION}
```

## Documentation

The module supports all Plaid API endpoints.  For complete information about
the API, head to the [docs](https://plaid.com/docs/api).

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

### Errors

In the case one of the endpoints you call returns an error, you can get the Plaid error object with the following:

```go
response, httpResponse, err := client.PlaidApi.Endpoint(...)
plaidErr, err := plaid.ToPlaidError(err)
fmt.Println(plaidErr.ErrorMessage)
```

## Authentication

First, you get your `client_id` and `secret` from your dashboard account. Authentication is handled by setting the `client_id` and `secret` on the configuration object.

## Contributing

Please see [Contributing](CONTRIBUTING.md) for guidelines and instructions for local development.

## License

[MIT](LICENSE)

