BINDIR:=bin
ROOT_PACKAGE:=$(shell go list .)
COMMAND_PACKAGES:=$(shell go list ./cmd | grep -v '^\([^/]*/\)\{4\}')
BINARIES:=$(COMMAND_PACKAGES:$(ROOT_PACKAGE)/%=$(BINDIR)/%)
GO_FILES:=$(shell find . -type f -name '*.go' -print)

# symbol table and dwarf
GO_LDFLAGS_SYMBOL:=
ifdef RELEASE
	GO_LDFLAGS_SYMBOL:=-w -s
endif
# static ldflag
GO_LDFLAGS_STATIC:=
ifdef RELEASE
	GO_LDFLAGS_STATIC:=-extldflags '-static'
endif
# build ldflags
GO_LDFLAGS:=$(GO_LDFLAGS_VERSION) $(GO_LDFLAGS_SYMBOL) $(GO_LDFLAGS_STATIC)
# build tags
GO_BUILD_TAGS:=debug
ifdef RELEASE
	GO_BUILD_TAGS:=release
endif
# race detector
GO_BUILD_RACE:=-race
ifdef RELEASE
	GO_BUILD_RACE:=
endif
# static build flag
GO_BUILD_STATIC:=
ifdef RELEASE
	GO_BUILD_STATIC:=-a -installsuffix netgo
	GO_BUILD_TAGS:=$(GO_BUILD_TAGS),netgo
endif
GO_BUILD:=-tags=$(GO_BUILD_TAGS) $(GO_BUILD_RACE) $(GO_BUILD_STATIC) -ldflags "$(GO_LDFLAGS)"

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: lint
lint: ## Run linters
	@echo "pinact"
	@pinact run --check || true
	@echo "actionlint"
	@actionlint || true
	@echo "ghalint run"
	@ghalint run || true
	@echo "ghalint act"
	@ghalint act || true
	@echo "zizmor"
	@zizmor -q . | grep -v 'No findings to report. Good job!' || true

.PHONY: gen
gen: ## Generate code
	@go generate ./...

.PHONY: build
build: $(BINARIES) ## Build all binaries. If RELEASE is set, it will build release binaries.
	@echo "Binaries built in $(BINDIR)/"

$(BINARIES): $(GO_FILES) .git/HEAD
	@CGO_ENABLED=0 go build -o $@ $(GO_BUILD) $(@:$(BINDIR)/%=$(ROOT_PACKAGE)/%)

.PHONY: test
test: ## Run tests
	@gotest -race -shuffle on -timeout 3s -count 2 ./...

.PHONY: run
run: ## Run the main application
	@air -c .air.toml -build.bin $(BINARIES)

.PHONY: clean
clean: ## Clean up build artifacts
	@$(RM) $(BINARIES)
