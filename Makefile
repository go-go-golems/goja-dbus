.PHONY: all gifs docker-lint lint lintmax gosec govulncheck test build glazed-lint-build glazed-lint docsctl-install docsctl-export docsctl-validate logcopter-generate logcopter-check xgoja-doctor xgoja-build goreleaser-check goreleaser-snapshot goreleaser tag-major tag-minor tag-patch release bump-go-go-golems install

all: gifs

VERSION=v0.1.14
GORELEASER_ARGS ?= --skip=sign --snapshot --clean
GORELEASER_TARGET ?= --single-target
GLAZED_LINT_BIN ?= /tmp/glazed-lint
GLAZED_LINT_PKG ?= github.com/go-go-golems/glazed/cmd/tools/glazed-lint
GLAZED_VERSION ?= $(shell GOWORK=off go list -m -f '{{.Version}}' github.com/go-go-golems/glazed 2>/dev/null)
GLAZED_LINT_FLAGS ?=
DOCSCTL_BIN ?= $(shell command -v docsctl 2>/dev/null || echo /tmp/docsctl)
DOCSCTL_PKG ?= github.com/go-go-golems/glazed/cmd/docsctl
DOCSCTL_SQLITE ?= .docsctl/help.sqlite
DOCSCTL_PACKAGE ?= goja-dbus
DOCSCTL_VERSION ?= v0.0.0-local

TAPES=$(wildcard doc/vhs/*tape)
gifs: $(TAPES)
	for i in $(TAPES); do vhs < $$i; done

docker-lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run -v

lint: glazed-lint
	GOWORK=off golangci-lint run -v

lintmax: glazed-lint
	GOWORK=off golangci-lint run -v --max-same-issues=100

glazed-lint-build:
	@echo "Building glazed-lint from Glazed module..."
	@if [ -n "$(GLAZED_VERSION)" ] && [ "$(GLAZED_VERSION)" != "(devel)" ]; then \
		echo "Installing $(GLAZED_LINT_PKG)@$(GLAZED_VERSION)"; \
		GOBIN=$(dir $(GLAZED_LINT_BIN)) GOWORK=off go install $(GLAZED_LINT_PKG)@$(GLAZED_VERSION); \
	else \
		echo "Installing $(GLAZED_LINT_PKG) from current module graph"; \
		GOBIN=$(dir $(GLAZED_LINT_BIN)) GOWORK=off go install $(GLAZED_LINT_PKG); \
	fi

glazed-lint: glazed-lint-build
	GOWORK=off go vet -vettool=$(GLAZED_LINT_BIN) $(GLAZED_LINT_FLAGS) ./cmd/... ./pkg/...

gosec:
	GOWORK=off go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec -exclude-generated -exclude=G101,G304,G301,G306 -exclude-dir=.history ./...

govulncheck:
	GOWORK=off go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

test:
	GOWORK=off go test ./...

build:
	GOWORK=off go generate ./...
	GOWORK=off go build ./...

docsctl-install:
	@if [ ! -x "$(DOCSCTL_BIN)" ]; then \
		echo "Installing $(DOCSCTL_PKG)@$(GLAZED_VERSION)"; \
		GOBIN=$(dir $(DOCSCTL_BIN)) GOWORK=off go install $(DOCSCTL_PKG)@$(GLAZED_VERSION); \
	fi

docsctl-export:
	mkdir -p $(dir $(DOCSCTL_SQLITE))
	GOWORK=off go run ./cmd/goja-dbus help export --format sqlite --output-path $(DOCSCTL_SQLITE)

docsctl-validate: docsctl-install docsctl-export
	$(DOCSCTL_BIN) validate --file $(DOCSCTL_SQLITE) --package $(DOCSCTL_PACKAGE) --version $(DOCSCTL_VERSION)

xgoja-doctor:
	cd ../go-go-goja && GOWORK=off go run ./cmd/xgoja doctor -f ../goja-dbus/cmd/goja-dbus/xgoja.yaml

xgoja-build:
	cd ../go-go-goja && GOWORK=off go run ./cmd/xgoja build -f ../goja-dbus/cmd/goja-dbus/xgoja.yaml --output ../goja-dbus/dist/goja-dbus

logcopter-generate:
	GOWORK=off go generate ./...

logcopter-check:
	GOWORK=off go tool logcopter-gen -area-prefix go-go-golems.goja-dbus -strip-prefix github.com/go-go-golems/goja-dbus -check ./pkg/... ./cmd/...

goreleaser-check:
	GOWORK=off goreleaser check

goreleaser-snapshot:
	GOWORK=off goreleaser release --skip=sign --snapshot --clean --single-target

goreleaser:
	GOWORK=off goreleaser release $(GORELEASER_ARGS) $(GORELEASER_TARGET)

tag-major:
	git tag $(shell svu major)

tag-minor:
	git tag $(shell svu minor)

tag-patch:
	git tag $(shell svu patch)

release:
	git push origin --tags
	GOWORK=off GOPROXY=proxy.golang.org go list -m github.com/go-go-golems/goja-dbus@$(shell svu current)

bump-go-go-golems:
	@deps="$$(awk '/^require[[:space:]]+github\.com\/go-go-golems\// { print $$2 } /^[[:space:]]*github\.com\/go-go-golems\// { print $$1 }' go.mod | sort -u)"; \
	if [ -z "$$deps" ]; then \
		echo "No github.com/go-go-golems dependencies in go.mod"; \
	else \
		echo "Bumping go-go-golems dependencies:"; \
		echo "$$deps"; \
		for dep in $$deps; do GOWORK=off go get "$${dep}@latest"; done; \
	fi
	GOWORK=off go mod tidy

GOJA_DBUS_BINARY ?= $(shell command -v goja-dbus 2>/dev/null || echo $(HOME)/bin/goja-dbus)
install: xgoja-build
	install -D ./dist/goja-dbus $(GOJA_DBUS_BINARY)
