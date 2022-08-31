# Go embeds the package version in the generator.
GO_PACKAGE_VERSION=8.2.1

.PHONY: test
test:
	go test -v ./...
