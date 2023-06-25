PROJECT_PKG = github.com/bo-at-pleno/go-thumbs
BUILD_DIR = build
VERSION ?=$(shell git describe --tags --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)
# remove debug info from the binary & make it smaller
LDFLAGS += -s -w
# inject build info
LDFLAGS += -X ${PROJECT_PKG}/internal/app/build.Version=${VERSION} -X ${PROJECT_PKG}/internal/app/build.CommitHash=${COMMIT_HASH} -X ${PROJECT_PKG}/internal/app/build.BuildDate=${BUILD_DATE}
GOPATH ?= $(shell go env GOPATH)

debug-makefile:
	@echo $(LDFLAGS)
	@echo $(shell go env GOPATH)

start-docker-compose-test:
	docker-compose -f docker-compose-test.yml up -d

stop-docker-compose-test:
	docker-compose -f docker-compose-test.yml down

test-all:
	$(MAKE) start-docker-compose-test
	go test -v ./...
	${MAKE} stop-docker-compose-test

.PHONY: build debug-makefile
build:
	go build ${GOARGS} -tags "${GOTAGS}" -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/app ./cmd/app

run:
	go run ./cmd/app

gen:
	go generate ./...

deps:
	wire ./...

swagger:
	swag init --parseDependency -g cmd/app/main.go --output=./api

proto:
	protoc --go_out=plugins=grpc:. internal/grpc/schema/*.proto

install-tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin v1.33.0
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/onsi/ginkgo/v2/ginkgo
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
