# POSTGRES_DATA_DIR=postgresdata
# POSTGRES_DOCKER_IMAGE=circleci/postgres:9.6-alpine
# POSTGRES_PORT=5432
# POSTGRES_DB_NAME=civil_crawler
# POSTGRES_USER=docker
# POSTGRES_PSWD=docker

GOCMD=go
GOGEN=$(GOCMD) generate
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCOVER=$(GOCMD) tool cover
ABIGEN=abigen

LIB_GEN_MAIN=cmd/libgen/main.go

ABI_DIR=abi

## List of expected dirs for generated code
GENERATED_DIR=pkg/generated
GENERATED_CONTRACT_DIR=$(GENERATED_DIR)/contract

GOMETALINTER_INSTALLER=scripts/gometalinter_install.sh
GOMETALINTER_VERSION_TAG=v2.0.11


GO:=$(shell command -v go 2> /dev/null)
# DOCKER:=$(shell command -v docker 2> /dev/null)
APT:=$(shell command -v apt-get 2> /dev/null)

## Reliant on go and $GOPATH being set
.PHONY: check-go-env
check-go-env:
ifndef GO
	$(error go command is not installed or in PATH)
endif
ifndef GOPATH
	$(error GOPATH is not set)
endif

## NOTE: If installing on a Mac, use Docker for Mac, not Docker toolkit
## https://www.docker.com/docker-mac
# .PHONY: check-docker-env
# check-docker-env:
# ifndef DOCKER
# 	$(error docker command is not installed or in PATH)
# endif

.PHONY: install-dep
install-dep: check-go-env ## Installs dep
	@mkdir -p $(GOPATH)/bin
	@curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

.PHONY: install-linter
install-linter: check-go-env ## Installs linter
	sh $(GOMETALINTER_INSTALLER) -b $(GOPATH)/bin $(GOMETALINTER_VERSION_TAG)

ifdef APT
	@sudo apt-get install golang-race-detector-runtime || true
endif

.PHONY: install-cover
install-cover: check-go-env ## Installs code coverage tool
	@$(GOGET) -u golang.org/x/tools/cmd/cover

.PHONY: install-abigen
install-abigen: check-go-env ## Installs the Ethereum abigen tool
	@$(GOGET) -u github.com/ethereum/go-ethereum/cmd/abigen

.PHONY: setup
setup: check-go-env install-dep install-linter install-cover install-abigen ## Sets up the tooling.

# .PHONY: postgres-setup-launch
# postgres-setup-launch:
# ifeq ("$(wildcard $(POSTGRES_DATA_DIR))", "")
# 	mkdir -p $(POSTGRES_DATA_DIR)
# 	docker run \
# 		-v $$PWD/$(POSTGRES_DATA_DIR):/tmp/$(POSTGRES_DATA_DIR) -i -t $(POSTGRES_DOCKER_IMAGE) \
# 		/bin/bash -c "cp -rp /var/lib/postgresql /tmp/$(POSTGRES_DATA_DIR)"
# endif
# 	docker run -e "POSTGRES_USER="$(POSTGRES_USER) -e "POSTGRES_PASSWORD"=$(POSTGRES_PSWD) -e "POSTGRES_DB"=$(POSTGRES_DB_NAME) \
# 	    -v $$PWD/$(POSTGRES_DATA_DIR)/postgresql:/var/lib/postgresql -d -p $(POSTGRES_PORT):$(POSTGRES_PORT) \
# 		$(POSTGRES_DOCKER_IMAGE);

# .PHONY: postgres-check-available
# postgres-check-available:
# 	@for i in `seq 1 10`; \
# 	do \
# 		nc -z localhost 5432 2> /dev/null && exit 0; \
# 		sleep 3; \
# 	done; \
# 	exit 1;

# .PHONY: postgres-start
# postgres-start: check-docker-env postgres-setup-launch postgres-check-available ## Starts up a development PostgreSQL server
# 	@echo "Postgresql launched and available"

# .PHONY: postgres-stop
# postgres-stop: check-docker-env ## Stops the development PostgreSQL server
# 	@docker stop `docker ps -q`
# 	@echo 'Postgres stopped'

.PHONY: generate-civil-contracts
generate-civil-contracts: ## Builds the contract wrapper code from the ABIs in /abi for Civil.
ifneq ("$(wildcard $(ABI_DIR)/*.abi)", "")
	@mkdir -p $(GENERATED_CONTRACT_DIR)

	@# Produce the contract bin/abi and binding files
	@$(ABIGEN) -abi ./$(ABI_DIR)/CivilTCR.abi -bin ./$(ABI_DIR)/CivilTCR.bin -type CivilTCRContract -out ./$(GENERATED_CONTRACT_DIR)/CivilTCRContract.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/Newsroom.abi -bin ./$(ABI_DIR)/Newsroom.bin -type NewsroomContract -out ./$(GENERATED_CONTRACT_DIR)/NewsroomContract.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/CivilPLCRVoting.abi -bin ./$(ABI_DIR)/CivilPLCRVoting.bin -type CivilPLCRVotingContract -out ./$(GENERATED_CONTRACT_DIR)/CivilPLCRVotingContract.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/Parameterizer.abi -bin ./$(ABI_DIR)/Parameterizer.bin -type ParameterizerContract -out ./$(GENERATED_CONTRACT_DIR)/ParameterizerContract.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/Government.abi -bin ./$(ABI_DIR)/Government.bin -type GovernmentContract -out ./$(GENERATED_CONTRACT_DIR)/GovernmentContract.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/EIP20.abi -bin ./$(ABI_DIR)/EIP20.bin -type EIP20Contract -out ./$(GENERATED_CONTRACT_DIR)/EIP20.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/DummyTokenTelemetry.abi -bin ./$(ABI_DIR)/DummyTokenTelemetry.bin -type DummyTokenTelemetryContract -out ./$(GENERATED_CONTRACT_DIR)/DummyTokenTelemetry.go -pkg contract

	@# Produce the bin/abi files
	@# NOTE(PN): The ABIs for these need to have the Data types replaced with "string" before this will successfully work.
	@# This is due to abigen no being able to handle user defined structs, but not needed
	@# for our purposes
	@cp ./$(ABI_DIR)/AttributeStore.abi ./$(ABI_DIR)/AttributeStore.abi.bak
	@sed -i "" 's/AttributeStore\.Data\ storage/string/g' ./$(ABI_DIR)/AttributeStore.abi
	@$(GORUN) $(LIB_GEN_MAIN) -abi ./$(ABI_DIR)/AttributeStore.abi -bin ./$(ABI_DIR)/AttributeStore.bin -type AttributeStoreContract -out ./$(GENERATED_CONTRACT_DIR)/AttributeStoreContract.go -pkg contract
	@mv ./$(ABI_DIR)/AttributeStore.abi.bak ./$(ABI_DIR)/AttributeStore.abi

	@cp ./$(ABI_DIR)/DLL.abi ./$(ABI_DIR)/DLL.abi.bak
	@sed -i "" 's/DLL\.Data\ storage/string/g' ./$(ABI_DIR)/DLL.abi
	@$(GORUN) $(LIB_GEN_MAIN) -abi ./$(ABI_DIR)/DLL.abi -bin ./$(ABI_DIR)/DLL.bin -type DLLContract -out ./$(GENERATED_CONTRACT_DIR)/DLLContract.go -pkg contract
	@mv ./$(ABI_DIR)/DLL.abi.bak ./$(ABI_DIR)/DLL.abi

else
	$(error No abi files found; copy them to /abi after generation)
endif

## gometalinter config in .gometalinter.json
.PHONY: lint
lint: ## Runs linting.
	@gometalinter ./...

.PHONY: build
build: ## Builds the repo, mainly to ensure all the files will build properly
	$(GOBUILD) ./...

.PHONY: test
test: ## Runs unit tests and tests code coverage
	@echo 'mode: atomic' > coverage.txt && $(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

.PHONY: test-integration
test-integration: ## Runs tagged integration tests
	@echo 'mode: atomic' > coverage.txt && PUBSUB_EMULATOR_HOST=localhost:8042 $(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s -tags=integration ./...

.PHONY: cover
cover: test ## Runs unit tests, code coverage, and runs HTML coverage tool.
	@$(GOCOVER) -html=coverage.txt

.PHONY: clean
clean: ## go clean and clean up of artifacts
	@$(GOCLEAN) ./... || true
	@rm coverage.txt || true
	@rm build || true

## Some magic from http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

