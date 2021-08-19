# Contributing

Instructions for contributing to [plaid-go][1]. A go client library for the [Plaid API][2]. This library is fully generated from the [Plaid OpenAPI spec](3). The library is generated internally, so don't try to iterate from this repo.

## Running Tests

1. To build the docker image for the client tests, run `docker build -t plaid-go .`.
2. Go to the [Plaid Dashboard](https://dashboard.plaid.com/) and copy and paste your `client_id` and sandbox `secret` into the following command.
3. Run `docker run --rm -e CLIENT_ID=$CLIENT_ID -e SECRET=$SECRET plaid-go`.

If you wish to run a single test, do something like this:

```sh
CLIENT_ID="" SECRET="" go test -v ./... -run TESTNAME
```

[1]: https://github.com/plaid/plaid-go
[2]: https://plaid.com
[3]: https://github.com/plaid/plaid-openapi
