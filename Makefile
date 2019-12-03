# POSTGRES_DATA_DIR=postgresdata
# POSTGRES_DOCKER_IMAGE=circleci/postgres:9.6-alpine
# POSTGRES_PORT=5432
# POSTGRES_DB_NAME=civil_crawler
# POSTGRES_USER=docker
# POSTGRES_PSWD=docker

GOVERSION=go1.12.7

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

PUBSUB_SIM_DOCKER_IMAGE=kinok/google-pubsub-emulator:latest
REDIS_DOCKER_IMAGE=redis:4.0.14

ABI_DIR=abi

## List of expected dirs for generated code
GENERATED_DIR=pkg/generated
GENERATED_CONTRACT_DIR=$(GENERATED_DIR)/contract

# curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin vX.Y.Z
GOLANGCILINT_URL=https://install.goreleaser.com/github.com/golangci/golangci-lint.sh
GOLANGCILINT_VERSION_TAG=v1.16.0


GO:=$(shell command -v go 2> /dev/null)
DOCKER:=$(shell command -v docker 2> /dev/null)
APT:=$(shell command -v apt-get 2> /dev/null)
GOVERCURRENT=$(shell go version |awk {'print $$3'})


UNAME=$(shell uname)

SED=sed -i ""
ifeq ($(UNAME),Linux)
	SED=sed -i
endif

## Reliant on go and $GOPATH being set
.PHONY: check-go-env
check-go-env:
ifndef GO
	$(error go command is not installed or in PATH)
endif
ifndef GOPATH
	$(error GOPATH is not set)
endif
ifneq ($(GOVERCURRENT), $(GOVERSION))
	$(error Incorrect go version, needs $(GOVERSION))
endif

## NOTE: If installing on a Mac, use Docker for Mac, not Docker toolkit
## https://www.docker.com/docker-mac
.PHONY: check-docker-env
check-docker-env:
ifndef DOCKER
	$(error docker command is not installed or in PATH)
endif

.PHONY: install-gobin
install-gobin: check-go-env ## Installs gobin tool
	@GO111MODULE=off go get -u github.com/myitcv/gobin

# Use commit until they cut a release with our fix
.PHONY: install-conform
install-conform: install-gobin ## Installs conform
	@gobin github.com/autonomy/conform@7bed9129bc73b5a6fd2b6d8c12f7c024ea4b7107

.PHONY: install-linter
install-linter: check-go-env ## Installs linter
	@curl -sfL $(GOLANGCILINT_URL) | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCILINT_VERSION_TAG)

.PHONY: install-cover
install-cover: install-gobin ## Installs code coverage tool
	@gobin -u golang.org/x/tools/cmd/cover

# Update to matching version
.PHONY: install-abigen
install-abigen: install-gobin ## Installs the Ethereum abigen tool
	@gobin github.com/ethereum/go-ethereum/cmd/abigen@v0.0.0-20190528221609-008d250e3c57

.PHONY: setup-githooks
setup-githooks: ## Setups any git hooks in githooks
	@ln -f -s ../../githooks/commit-msg .git/hooks

.PHONY: setup
setup: check-go-env install-conform install-linter install-cover install-abigen setup-githooks ## Sets up the tooling.

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

.PHONY: pubsub-setup-launch
pubsub-setup-launch:
	@docker run -it -d -p 8042:8042 $(PUBSUB_SIM_DOCKER_IMAGE)

.PHONY: pubsub-start
pubsub-start: check-docker-env pubsub-setup-launch ## Starts up the pubsub simulator
	@echo 'Google pubsub simulator up'

.PHONY: pubsub-stop
pubsub-stop: check-docker-env ## Stops the pubsub simulator
	@docker stop `docker ps -q --filter "ancestor=$(PUBSUB_SIM_DOCKER_IMAGE)"`
	@echo 'Google pubsub simulator down'

.PHONY: redis-setup-launch
redis-setup-launch:
	@docker run -it -d -p 6379:6379 $(REDIS_DOCKER_IMAGE)

.PHONY: redis-start
redis-start: check-docker-env redis-setup-launch ## Starts up local test redis
	@echo 'Redis up'

.PHONY: redis-stop
redis-stop: check-docker-env ## Stops the local test redis
	@docker stop `docker ps -q --filter "ancestor=$(REDIS_DOCKER_IMAGE)"`
	@echo 'Redis down'

## Used to test lock redis pools
.PHONY: redis-setup-launch-b
redis-setup-launch-b:
	@docker run -it -d -p 6378:6379 $(REDIS_DOCKER_IMAGE)

.PHONY: redis-start-b
redis-start-b: check-docker-env redis-setup-launch-b ## Starts up local test redis
	@echo 'Redis up'

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
	@$(ABIGEN) -abi ./$(ABI_DIR)/CVLToken.abi -bin ./$(ABI_DIR)/CVLToken.bin -type CVLTokenContract -out ./$(GENERATED_CONTRACT_DIR)/CVLToken.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/DummyTokenTelemetry.abi -bin ./$(ABI_DIR)/DummyTokenTelemetry.bin -type DummyTokenTelemetryContract -out ./$(GENERATED_CONTRACT_DIR)/DummyTokenTelemetry.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/CivilTokenController.abi -bin ./$(ABI_DIR)/CivilTokenController.bin -type CivilTokenControllerContract -out ./$(GENERATED_CONTRACT_DIR)/CivilTokenController.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/NoOpTokenController.abi -bin ./$(ABI_DIR)/NoOpTokenController.bin -type NoOpTokenControllerContract -out ./$(GENERATED_CONTRACT_DIR)/NoOpTokenController.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/NewsroomFactory.abi -bin ./$(ABI_DIR)/NewsroomFactory.bin -type NewsroomFactory -out ./$(GENERATED_CONTRACT_DIR)/NewsroomFactory.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/CreateNewsroomInGroup.abi -bin ./$(ABI_DIR)/CreateNewsroomInGroup.bin -type CreateNewsroomInGroupContract -out ./$(GENERATED_CONTRACT_DIR)/CreateNewsroomInGroup.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/MultiSigWalletFactory.abi -bin ./$(ABI_DIR)/MultiSigWalletFactory.bin -type MultiSigWalletFactoryContract -out ./$(GENERATED_CONTRACT_DIR)/MultiSigWalletFactory.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/MultiSigWallet.abi -bin ./$(ABI_DIR)/MultiSigWallet.bin -type MultiSigWalletContract -out ./$(GENERATED_CONTRACT_DIR)/MultiSigWallet.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/ECRecovery.abi -bin ./$(ABI_DIR)/ECRecovery.bin -type ECRecoveryContract -out ./$(GENERATED_CONTRACT_DIR)/ECRecovery.go -pkg contract
	@$(ABIGEN) -abi ./$(ABI_DIR)/RootCommits.abi -bin ./$(ABI_DIR)/RootCommits.bin -type RootCommitsContract -out ./$(GENERATED_CONTRACT_DIR)/RootCommits.go -pkg contract

	@# Produce the bin/abi files
	@# NOTE(PN): The ABIs for these need to have the Data types replaced with "string" before this will successfully work.
	@# This is due to abigen not being able to handle user defined structs, but not needed
	@# for our purposes
	@cp ./$(ABI_DIR)/AttributeStore.abi ./$(ABI_DIR)/AttributeStore.abi.bak
	@$(SED) 's/AttributeStore\.Data\ storage/string/g' ./$(ABI_DIR)/AttributeStore.abi
	@$(GORUN) $(LIB_GEN_MAIN) -abi ./$(ABI_DIR)/AttributeStore.abi -bin ./$(ABI_DIR)/AttributeStore.bin -type AttributeStoreContract -out ./$(GENERATED_CONTRACT_DIR)/AttributeStoreContract.go -pkg contract
	@mv ./$(ABI_DIR)/AttributeStore.abi.bak ./$(ABI_DIR)/AttributeStore.abi

	@cp ./$(ABI_DIR)/DLL.abi ./$(ABI_DIR)/DLL.abi.bak
	@$(SED) 's/DLL\.Data\ storage/string/g' ./$(ABI_DIR)/DLL.abi
	@$(GORUN) $(LIB_GEN_MAIN) -abi ./$(ABI_DIR)/DLL.abi -bin ./$(ABI_DIR)/DLL.bin -type DLLContract -out ./$(GENERATED_CONTRACT_DIR)/DLLContract.go -pkg contract
	@mv ./$(ABI_DIR)/DLL.abi.bak ./$(ABI_DIR)/DLL.abi

	@cp ./$(ABI_DIR)/MessagesAndCodes.abi ./$(ABI_DIR)/MessagesAndCodes.abi.bak
	@$(SED) 's/MessagesAndCodes\.Data\ storage/string/g' ./$(ABI_DIR)/MessagesAndCodes.abi
	@$(GORUN) $(LIB_GEN_MAIN) -abi ./$(ABI_DIR)/MessagesAndCodes.abi -bin ./$(ABI_DIR)/MessagesAndCodes.bin -type MessagesAndCodesContract -out ./$(GENERATED_CONTRACT_DIR)/MessagesAndCodesContract.go -pkg contract
	@mv ./$(ABI_DIR)/MessagesAndCodes.abi.bak ./$(ABI_DIR)/MessagesAndCodes.abi

else
	$(error No abi files found; copy them to /abi after generation)
endif

## golangci-lint config in .golangci.yml
.PHONY: lint
lint: check-go-env ## Runs linting.
	@golangci-lint run ./...

.PHONY: conform
conform: check-go-env ## Runs conform (commit message linting)
	@conform enforce

.PHONY: build
build: check-go-env ## Builds the repo, mainly to ensure all the files will build properly
	$(GOBUILD) ./...

.PHONY: test
test: check-go-env ## Runs unit tests and tests code coverage
	@echo 'mode: atomic' > coverage.txt && $(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=60s ./...

.PHONY: test-integration
test-integration: check-go-env ## Runs tagged integration tests
	@echo 'mode: atomic' > coverage.txt && PUBSUB_EMULATOR_HOST=localhost:8042 $(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=60s -tags=integration ./...

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

