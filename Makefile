# Go embeds the package version in the generator.
GO_PACKAGE_VERSION=21.0.0

.PHONY: test
test:
	go test -v ./...
