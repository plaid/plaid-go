# Contributing

Instructions for contributing to [plaid-go][1]. A go client library for the [Plaid API][2]. This library is fully generated from the [Plaid OpenAPI spec][3]. 

This library cannot directly accept PRs from the public as it is generated from internal Plaid sources on the internal Plaid GitHub instance and any changes made directly to this repo will be overwritten. If you submit a PR and it is accepted, a member of Plaid engineering will copy and paste your change into the upstream, internal version of this repo rather than merging your PR. Plaid employees should make any changes on the internal Plaid GitHub and not on the public repo. 

## Running Tests

The tests run against a local [WireMock](https://wiremock.org/) server that returns canned Plaid API responses, so they don't need Plaid credentials. You need Go 1.21 or later and Docker.

1. Start WireMock and point the tests at it:

   ```sh
   docker compose -f wiremock/docker-compose.test.yml up -d --wait
   export WIREMOCK_URL=http://localhost:$(docker compose -f wiremock/docker-compose.test.yml port wiremock 8080 | cut -d: -f2)
   ```

2. Run the tests with `go test ./...`.
3. Stop WireMock with `docker compose -f wiremock/docker-compose.test.yml down`.

If you wish to run a single test, start WireMock as above, then run:

```sh
go test -v ./... -run TESTNAME
```

Each test checks how many requests WireMock received for it, so a test fails if it runs a second time against the same WireMock server. Before you rerun, clear WireMock's request log with `curl -X DELETE $WIREMOCK_URL/__admin/requests`, or restart WireMock.

[1]: https://github.com/plaid/plaid-go
[2]: https://plaid.com
[3]: https://github.com/plaid/plaid-openapi
