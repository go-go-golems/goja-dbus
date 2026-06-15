.PHONY: gifs logcopter-generate logcopter-check xgoja-doctor xgoja-build

all: gifs

VERSION=v0.1.14
GORELEASER_ARGS ?= --skip=sign --snapshot --clean
GORELEASER_TARGET ?= --single-target

TAPES=$(wildcard doc/vhs/*tape)
gifs: $(TAPES)
	for i in $(TAPES); do vhs < $$i; done

docker-lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run -v

lint:
	GOWORK=off golangci-lint run -v

lintmax:
	GOWORK=off golangci-lint run -v --max-same-issues=100

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

xgoja-doctor:
	cd ../go-go-goja && GOWORK=off go run ./cmd/xgoja doctor -f ../goja-dbus/xgoja.yaml

xgoja-build:
	cd ../go-go-goja && GOWORK=off go run ./cmd/xgoja build -f ../goja-dbus/xgoja.yaml --output ../goja-dbus/dist/goja-dbus-xgoja

logcopter-generate:
	GOWORK=off go generate ./...

logcopter-check:
	GOWORK=off go tool logcopter-gen -area-prefix go-go-golems.goja-dbus -strip-prefix github.com/go-go-golems/goja-dbus -check ./pkg/...

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

GOJA_DBUS_DEMO_BINARY=$(shell which goja-dbus-demo)
install:
	GOWORK=off go build -o ./dist/goja-dbus-demo ./cmd/goja-dbus-demo && \
		cp ./dist/goja-dbus-demo $(GOJA_DBUS_DEMO_BINARY)
