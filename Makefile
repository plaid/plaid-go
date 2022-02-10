# Go embeds the package version in the generator.
GO_PACKAGE_VERSION=2.0.1

.PHONY: test
test:
	go test -v ./...
