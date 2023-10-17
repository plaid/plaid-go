# Contributing

Instructions for contributing to [plaid-go][1]. A go client library for the [Plaid API][2]. This library is fully generated from the [Plaid OpenAPI spec](3). 

This library cannot directly accept PRs from the public as it is generated from internal Plaid sources on the internal Plaid GitHub instance and any changes made directly to this repo will be overwritten. If you submit a PR and it is accepted, a member of Plaid engineering will copy and paste your change into the upstream, internal version of this repo rather than merging your PR. Plaid employees should make any changes on the internal Plaid GitHub and not on the public repo. 

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
