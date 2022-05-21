.SHELLFLAGS=-c -o pipefail
SHELL=/bin/bash

export CGO_LDFLAGS_ALLOW="-Wl,-z,now"
export CGO_LDFLAGS="-ldl -llz4 -lm"

export GO111MODULE = on
export GOFLAGS = -mod=vendor
export CGO_LDFLAGS_ALLOW="-Wl,-z,now"
export CGO_LDFLAGS="${BUILDER_LINKER_FLAGS}"

LINT_BIN		:= $(shell go env GOPATH)/bin/golangci-lint

deps:
	go mod download
	go mod tidy
	go mod verify
	go mod vendor

fmt: ; $(info Running goimports to format the codebase...) @
	@goimports -w -e $$(find . -type f -name '*.go' -not -path "*/vendor/*")

$(LINT_BIN):  ; $(info Getting lint...) @
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.41.1

lint: $(LINT_BIN) ; $(info Running lint...) @
	CGO_LDFLAGS_ALLOW="-Wl,-z,now" CGO_LDFLAGS="${BUILDER_LINKER_FLAGS}" $(LINT_BIN) --timeout 10m run
	CGO_LDFLAGS_ALLOW="-Wl,-z,now" CGO_LDFLAGS="${BUILDER_LINKER_FLAGS}" go vet ./...

test: $(info Running tests...)
	CGO_ENABLED=0 go test -coverprofile cover.out ./...
	CGO_ENABLED=0 go tool cover -html=cover.out -o cover.html

.PHONY: all