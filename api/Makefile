SHELL := /bin/bash
APP_NAME=api
TEST_PKGS := $(shell go list ./... | grep -v 'mock_*' | grep -v 'tester' | grep -v 'proto' | grep -v 'cmd'| grep -v 'api')
# build with version infos
VERSION_DIR = main
BUILD_DATE = $(shell date +%FT%T)
GIT_COMMIT = $(shell git log --pretty=format:'%h' -n 1)
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

GOLDFLAGS += -X "${VERSION_DIR}.date=${BUILD_DATE}"
GOLDFLAGS += -X "${VERSION_DIR}.commit=${GIT_COMMIT}"
GOLDFLAGS += -X "${VERSION_DIR}.branch=${GIT_BRANCH}"
GOLDFLAGS += -X "${VERSION_DIR}.version=${APP_VERSION}"

RED=\033[0;31m
GREEN=\033[0;32m
BLUE=\033[0;34m
NC=\033[0m

help: Makefile
	@printf "${BLUE}Choose a command run:${NC}\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/    /'

## make update: Update package
.PHONY: update
update:
	@go get -u

## make tidy: Run go mod tidy
.PHONY: tidy
tidy:
	@go mod tidy

## make build: Go build the project
.PHONY: build
build:
	@cd repo && packr2
	@go build -ldflags '${GOLDFLAGS}'  -o bin/${APP_NAME} ./cmd/${APP_NAME}
	@printf "${GREEN}Build ${APP_NAME} successfully!${NC}\n"

## make install: Go install the project
.PHONY: install
install:
	@rm -rf config/logs
	@cd repo && packr2
	@go install -ldflags '${GOLDFLAGS}'  ./cmd/${APP_NAME}
	@printf "${GREEN}Install ${APP_NAME} successfully!${NC}\n"

## make test: Run go unittest
.PHONY: test
test:
	go generate ./...
	go test -timeout 300s ${TEST_PKGS} -count=1

## make test-coverage: Test project with cover
.PHONY: test-coverage
test-coverage:
	go generate ./...
	@go test -timeout 300s -short -coverprofile cover.out -covermode=atomic ${TEST_PKGS}
	@cat cover.out | grep -v "pb.go" | grep -v "atomic" >> ../coverage.txt || echo "no test"

## make docker: Docker build the project
.PHONY: docker
docker:
	@docker build -t ${APP_NAME}:latest .
