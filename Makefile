# Go embeds the package version in the generator.
GO_PACKAGE_VERSION=0.1.0.beta.2

.PHONY: test
test:
	go test -v ./...
