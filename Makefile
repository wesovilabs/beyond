GOLANG_VERSION=1.13
COMMIT = $(shell git log --pretty=format:'%H' -n 1)
VERSION    = $(shell git describe --tags --always)
BUILD_DATE = $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS   = -ldflags "\
 -X github.com/wesovilabs/goa/goa.Commit=$(COMMIT) \
 -X github.com/wesovilabs/goa/goa.Version=$(VERSION) \
 -X github.com/wesovilabs/goa/goa.BuildDate=$(BUILD_DATE) \
 -X github.com/wesovilabs/goa/goa.Compiler=$(GOLANG_VERSION)"

# Go
GO  = GOFLAGS=-mod=vendor go
GOBUILD  = CGO_ENABLED=0 $(GO) build $(LDFLAGS)

all: fmt lint test

clean: ; @ ## Remove temporal files
	rm -f coverage.txt;
	rm -r vendor

deps: ; @ ## Download dependencies
	${GO} mod vendor
	${GO} mod download

test: ; @ ## Run tests
	${GO} test  -v ./...

test-coverage: ; @ ## Run tests with coverage
	${GO} test -race -coverprofile=coverage.txt -covermode=atomic ./...

fmt: ; @ ## Format code
	${GO} fmt ./...

lint: fmt ; @ ## Format code and run linter
	${GO} run github.com/golangci/golangci-lint/cmd/golangci-lint run --verbose

benchmark: ; @ ## Run benchmark tests
	${GO} test -bench Benchmark.+ -failfast -run -Benchmark.+ -v ./benchmark/...


.PHONY: build
build: ; @ ## build exeutable for your current osm
	$(GOBUILD) -o build/goa


.PHONY: build-all
build-all: ; @ ## Build binary files
	GOARCH=amd64 GOOS=linux  $(GOBUILD) -o build/goa.linux main.go
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o build/goa.darwin main.go
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o build/goa.exe main.go

.PHONY: init
init: ; @ ## Setup the git hooks
	chmod +x scripts/.githooks/*
	git config core.hooksPath scripts/.githooks

docker-%: ; @ ## Run commands inside a docker container
	docker run --rm --workdir /app -v $(CURDIR):/app golang:$(GOLANG_VERSION) \
	make $*

run:
	go run main.go \
		--project github.com/wesovilabs/goa/testdata/basic \
		--goPath /Users/ivan/Workspace/Wesovilabs/goa/testdata/basic \
		--output /Users/ivan/Workspace/Wesovilabs/goa/testdata/generated \
		--verbose

help:
	@grep -E '^[a-zA-Z_-]+[%]?:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
