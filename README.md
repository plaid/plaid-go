[![CircleCI](https://circleci.com/gh/plaid/plaid-go.svg?style=svg)](https://circleci.com/gh/plaid/plaid-go)

# plaid-go

A Go client library for the [Plaid API](https://plaid.com/docs).

## Table of Contents

- [plaid-go](#plaid-go)
    * [Install](#install)
    * [Versioning](#versioning)
    * [Documentation](#documentation)
    * [Getting Started](#getting-started)
    * [Contributing](#contributing)
    * [License](#license)

## Install

```console
$ go get github.com/plaid/plaid-go
```

## Versioning

Each major version of `plaid-go` targets a specific version of the Plaid API:

| API version | plaid-go release |
| ----------- | ------------------ |
| [`2018-05-22`][api-version-2018-05-22] (**latest**) | `6.x.x` |
| `2017-03-01` | `5.x.x` |

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
    os.Getenv("PLAID_PUBLIC_KEY"),
    plaid.Sandbox, // Available environments are Sandbox, Development, and Production
    &http.Client{}, // This parameter is optional
    plaid.LatestAPIVersion, // e.g. "2018-05-22"
}
client, err := plaid.NewClient(clientOptions)
```

Each endpoint returns an object which contains the parsed JSON from the HTTP response.

### Errors

All non-200 responses will return a plaid.Error instance.

For more information on Plaid response codes, head to the [docs](https://plaid.com/docs/api#errors).

## Support

Open an [issue](https://github.com/plaid/plaid-python/issues/new)!

## License

[MIT](https://github.com/plaid/plaid-go/blob/master/LICENSE)

