# Go embeds the package version in the generator.
GO_PACKAGE_VERSION=12.0.0

.PHONY: test
test:
	go test -v ./...
