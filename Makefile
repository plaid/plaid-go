# Go embeds the package version in the generator.
GO_PACKAGE_VERSION=10.4.0

.PHONY: test
test:
	go test -v ./...
