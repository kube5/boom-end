RED=\033[0;31m
GREEN=\033[0;32m
BLUE=\033[0;34m
NC=\033[0m

SUBDIR=api user game

help: Makefile
	@printf "${BLUE}Choose a command run:${NC}\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/    /'

## make prepare: Preparation before development
.PHONY: prepare
prepare:
	@cd scripts && bash prepare.sh

## make proto: Generate proto
.PHONY: proto
proto:
	cd common && make proto

## make build: Go build all services
.PHONY: build
build:
	@for subdir in $(SUBDIR); \
	do \
		echo "building" $$subdir; \
		cd $$subdir && make build && cd ..; \
    done

## make test: Run all services unittest
.PHONY: test
test:
	@for subdir in $(SUBDIR); \
	do \
		echo "testing" $$subdir; \
		cd $$subdir && make test && cd ..; \
    done

## make test-coverage: Test all services with cover
.PHONY: test-coverage
test-coverage:
	echo "mode: atomic" >> coverage.txt
	@for subdir in $(SUBDIR); \
	do \
		echo "test-coverage" $$subdir; \
		cd $$subdir && make test-coverage && cd ..; \
    done

## make install: Go install all services
.PHONY: install
install:
	@for subdir in $(SUBDIR); \
	do \
		echo "installing" $$subdir; \
		cd $$subdir && make install && cd ..; \
    done

## make docker: Docker build all services
docker:
	@for subdir in $(SUBDIR); \
	do \
		echo "docker" $$subdir; \
		cd $$subdir && make docker && cd ..; \
    done

