GO_SRC_PACKAGES =$(shell go list ./... | sed /vendor/d )
GOLINT_SRC = ./vendor/golang.org/x/lint/golint

# vanity
GREEN = \033[0;32m
MAGENTA = \033[0;35m
BLUE = \033[0;34m
RESET = \033[0;0m

# setup
.PHONY: setup
setup: install-dep vendor bin/golint

export DEP_RELEASE_TAG = v0.4.1
.PHONY: install-dep
install-dep:
	@echo "$(BLUE)[installing dep@$(DEP_RELEASE_TAG)]$(RESET)"
	@./scripts/install_dep.sh

vendor: Gopkg.toml Gopkg.lock
	@echo "$(GREEN)installing vendored dependencies...$(RESET)"
	@# use the vendor-only flag to prevent us from removing dependencies before
	@# they are added to the docker container
	@dep ensure -v --vendor-only

bin/golint: vendor
	@echo "$(MAGENTA)building $(@)...$(RESET)"
	@go build -o $(@) $(GOLINT_SRC)

.PHONY: test
test:
	@echo "$(MAGENTA)running go tests...$(RESET)"
	@go test -v $(GO_SRC_PACKAGES)

.PHONY: lint
lint: bin/golint
	@echo "$(MAGENTA)linting $(GO_SRC_PACKAGES)$(RESET)"
	# TODO: lint errors should fail this step (use -set_exit_status)
	@bin/golint $(GO_SRC_PACKAGES)

# releasing
.PHONY: release-%
release-%:
	@echo "$(BLUE)staging for a new release$(RESET)"
	go run ./internal/cmd/cmd.go release $(@:release-%=%)
	@echo "$(BLUE)the package has been prepared for release, please review the changes made to"
	@echo "the git repo by inspecting 'git log' and 'git tag --list | tail -n 5'"
	@echo "before pushing the changes with 'git push --all'$(RESET)"
