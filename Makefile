.SHELLFLAGS=-c -o pipefail
SHELL=/bin/bash

export CGO_LDFLAGS_ALLOW="-Wl,-z,now"
export CGO_LDFLAGS="-ldl -llz4 -lm"

export GO111MODULE = on
export GOFLAGS = -mod=vendor
export CGO_LDFLAGS_ALLOW="-Wl,-z,now"
export CGO_LDFLAGS="${BUILDER_LINKER_FLAGS}"

BRANCH := $(shell git branch --show-current)
CLEANBRANCH := $(shell git branch --show-current | sed 's/\//-/')
SHORT := $(shell git rev-parse --short HEAD)
DEVTAG := ${CLEANBRANCH}-${SHORT}

LINT_BIN		:= $(shell go env GOPATH)/bin/golangci-lint

ifeq (${BRANCH}, main)
	TAG := $(shell cat .version)
else
	TAG := "$(shell cat .version)-${DEVTAG}"
endif

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

test: fmt $(info Running tests...)
	# NOTE: make sure to export the values for the env variables in the terminal prompt before running the tests
	export BYBIT_API_KEY=${BYBIT_API_KEY}
	export BYBIT_API_SECRET=${BYBIT_API_SECRET}
	CGO_ENABLED=0 go test -coverprofile cover.out ./...
	CGO_ENABLED=0 go tool cover -html=cover.out -o cover.html

.PHONY: all

# NOTE: Install godoc to run this command: `go install golang.org/x/tools/cmd/godoc@latest`
docs:
	godoc -http=:6060

tag:
	@echo "[!!!] Releasing git tag ${TAG} [!!!]"
	git tag ${TAG}
	git push --tags
	@echo "[!!!] Done! [!!!]"