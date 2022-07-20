# Go embeds the package version in the generator.
GO_PACKAGE_VERSION=8.1.0

.PHONY: test
test:
	go test -v ./...
