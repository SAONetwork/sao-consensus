#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
TM_PKG_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')
COSMOS_PKG_VERSION := $(shell go list -m github.com/cosmos/cosmos-sdk | sed 's:.* ::')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= false
PROJECT_NAME = sao
DOCKER:=docker
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
HTTPS_GIT := https://github.com/SaoNetwork/sao.git

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc
endif

ifeq (secp,$(findstring secp,$(COSMOS_BUILD_OPTIONS)))
  build_tags += libsecp256k1_sdk
endif

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=sao \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=sao \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
		  -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_PKG_VERSION)

# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (badgerdb,$(findstring badgerdb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
  BUILD_TAGS += badgerdb
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(COSMOS_BUILD_OPTIONS)))
  CGO_ENABLED=1
  BUILD_TAGS += rocksdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(COSMOS_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb
endif

ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

all: install

build: go.sum
ifeq ($(OS), Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/$(shell go env GOOS)/saod.exe ./cmd/saod
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/$(shell go env GOOS)/saod ./cmd/saod
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/saod

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download
PHONY: go-mod-cache

go.sum: go.mod
	@echo "--> Ensuring dependencies have not been modified"
	@go mod verify

clean:
	rm -rf build/

########################################
### Linting

# Check url links in the repo are not broken.
# This tool checks local markdown links as well.
# Set to exclude riot links as they trigger false positives
lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify
.PHONY: lint

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/tendermint
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/cosmos/cosmos-sdk
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/SaoNetwork/sao
.PHONY: format


###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=v0.3
protoImageName=tendermintdev/sdk-proto-gen:$(protoVer)
containerProtoGen=$(PROJECT_NAME)-proto-gen-$(protoVer)
containerProtoGenAny=$(PROJECT_NAME)-proto-gen-any-$(protoVer)
containerProtoGenSwagger=$(PROJECT_NAME)-proto-gen-swagger-$(protoVer)
containerProtoFmt=$(PROJECT_NAME)-proto-fmt-$(protoVer)

proto-all: proto-gen proto-format proto-lint proto-swagger-gen

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protocgen.sh; fi

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGenSwagger}$$"; then docker start -a $(containerProtoGenSwagger); else docker run --name $(containerProtoGenSwagger) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protoc-swagger-gen.sh; fi

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name *.proto -exec clang-format -style=file -i {} \; ; fi

proto-lint:
	@$(DOCKER_BUF) lint --error-format=json

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=master

GOOGLE_PROTO_URL = https://raw.githubusercontent.com/googleapis/googleapis/master/google/api
PROTOBUF_GOOGLE_URL = https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf
COSMOS_PROTO_URL = https://raw.githubusercontent.com/cosmos/cosmos-proto/master

GOOGLE_PROTO_TYPES = third_party/proto/google/api
PROTOBUF_GOOGLE_TYPES = third_party/proto/google/protobuf
COSMOS_PROTO_TYPES = third_party/proto/cosmos_proto

GOGO_PATH := $(shell go list -m -f '{{.Dir}}' github.com/gogo/protobuf)
TENDERMINT_PATH := $(shell go list -m -f '{{.Dir}}' github.com/tendermint/tendermint)
COSMOS_PROTO_PATH := $(shell go list -m -f '{{.Dir}}' github.com/cosmos/cosmos-proto)
COSMOS_SDK_PATH := $(shell go list -m -f '{{.Dir}}' github.com/cosmos/cosmos-sdk)
IBC_GO_PATH := $(shell go list -m -f '{{.Dir}}' github.com/cosmos/ibc-go/v3)
ETHERMINT_PATH := $(shell go list -m -f '{{.Dir}}' github.com/tharsis/ethermint)

proto-update-deps:
	mkdir -p $(GOOGLE_PROTO_TYPES)
	curl -sSL $(GOOGLE_PROTO_URL)/annotations.proto > $(GOOGLE_PROTO_TYPES)/annotations.proto
	curl -sSL $(GOOGLE_PROTO_URL)/http.proto > $(GOOGLE_PROTO_TYPES)/http.proto
	curl -sSL $(GOOGLE_PROTO_URL)/httpbody.proto > $(GOOGLE_PROTO_TYPES)/httpbody.proto

	mkdir -p $(PROTOBUF_GOOGLE_TYPES)
	curl -sSL $(PROTOBUF_GOOGLE_URL)/any.proto > $(PROTOBUF_GOOGLE_TYPES)/any.proto

	mkdir -p client/docs
	cp $(COSMOS_SDK_PATH)/client/docs/swagger-ui/swagger.yaml client/docs/cosmos-swagger.yml
	cp $(IBC_GO_PATH)/docs/client/swagger-ui/swagger.yaml client/docs/ibc-go-swagger.yml

	mkdir -p $(COSMOS_PROTO_TYPES)
	cp $(COSMOS_PROTO_PATH)/cosmos.proto $(COSMOS_PROTO_TYPES)/cosmos.proto

	rsync -r --chmod 644 --include "*.proto" --include='*/' --exclude='*' $(GOGO_PATH)/gogoproto third_party/proto
	rsync -r --chmod 644 --include "*.proto" --include='*/' --exclude='*' $(TENDERMINT_PATH)/proto third_party
	rsync -r --chmod 644 --include "*.proto" --include='*/' --exclude='*' $(COSMOS_SDK_PATH)/proto third_party
	rsync -r --chmod 644 --include "*.proto" --include='*/' --exclude='*' $(IBC_GO_PATH)/proto third_party
	rsync -r --chmod 644 --include "*.proto" --include='*/' --exclude='*' $(ETHERMINT_PATH)/proto third_party
	cp -f $(IBC_GO_PATH)/third_party/proto/proofs.proto third_party/proto/proofs.proto

.PHONY: proto-all proto-gen proto-gen-any proto-swagger-gen proto-format proto-lint proto-check-breaking proto-update-deps

########################################
### Testing

# TODO tidy up cli tests to use same -Enable flag as simulations, or the other way round
# TODO -mod=readonly ?
# build dependency needed for cli tests
test-all: build
	# basic app tests
	@go test ./app -v
	# basic simulation (seed "4" happens to not unbond all validators before reaching 100 blocks)
	#@go test ./app -run TestFullAppSimulation        -Enabled -Commit -NumBlocks=100 -BlockSize=200 -Seed 4 -v -timeout 24h
	# other sim tests
	#@go test ./app -run TestAppImportExport          -Enabled -Commit -NumBlocks=100 -BlockSize=200 -Seed 4 -v -timeout 24h
	#@go test ./app -run TestAppSimulationAfterImport -Enabled -Commit -NumBlocks=100 -BlockSize=200 -Seed 4 -v -timeout 24h
	# AppStateDeterminism does not use Seed flag
	#@go test ./app -run TestAppStateDeterminism      -Enabled -Commit -NumBlocks=100 -BlockSize=200 -Seed 4 -v -timeout 24h

# run module tests and short simulations
test-basic: test
	@go test ./app -run TestFullAppSimulation        -Enabled -Commit -NumBlocks=5 -BlockSize=200 -Seed 4 -v -timeout 2m
	# other sim tests
	@go test ./app -run TestAppImportExport          -Enabled -Commit -NumBlocks=5 -BlockSize=200 -Seed 4 -v -timeout 2m
	@go test ./app -run TestAppSimulationAfterImport -Enabled -Commit -NumBlocks=5 -BlockSize=200 -Seed 4 -v -timeout 2m
	@# AppStateDeterminism does not use Seed flag
	@go test ./app -run TestAppStateDeterminism      -Enabled -Commit -NumBlocks=5 -BlockSize=200 -Seed 4 -v -timeout 2m

test:
	@go test $$(go list ./... | grep -v 'contrib')

# Run tests for migration cli command
test-migrate:
	@go test -v -count=1 ./migrate/...


.PHONY: all build-linux install clean build test test-cli test-all test-rest test-basic start-remote-sims
