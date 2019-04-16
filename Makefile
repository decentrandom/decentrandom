#!/usr/bin/make -f

PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
GOBIN ?= $(GOPATH)/bin
GOSUM := $(shell which gosum)

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

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

# process linker flags

ldflags = -X github.com/decentrandom/decentrandom/version.Version=$(VERSION) \
	-X github.com/decentrandom/decentrandom/version.Commit=$(COMMIT) \
  -X "github.com/decentrandom/decentrandom/version.BuildTags=$(build_tags)"

ifneq ($(GOSUM),)
ldflags += -X github.com/decentrandom/decentrandom/version.GoSumHash=$(shell $(GOSUM) go.sum)
endif

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/decentrandom/decentrandom/types.DBBackend=cleveldb
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

########################################
### All

all: clean go-mod-cache install

########################################
### CI

ci: get_tools install

########################################

########################################
### Build/Install

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/randd.exe ./cmd/randd
	go build -mod=readonly $(BUILD_FLAGS) -o build/randcli.exe ./cmd/randcli
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/randd ./cmd/randd
	go build -mod=readonly $(BUILD_FLAGS) -o build/randcli ./cmd/randcli
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

update_rand_lite_docs:
	@statik -src=client/lcd/swagger-ui -dest=client/lcd -f

install: go.sum check-ledger update_rand_lite_docs
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/randd
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/randcli

dist:
	@bash publish/dist.sh
	@bash publish/publish.sh

########################################
### Tools & dependencies

get_tools:
	go get github.com/rakyll/statik
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

update_tools:
	@echo "--> Updating tools to correct version"
	$(MAKE) --always-make get_tools

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: get_tools
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

clean:
	rm -rf snapcraft-local.yaml build/

distclean: clean
	rm -rf vendor/

########################################
### Packaging

snapcraft-local.yaml: snapcraft-local.yaml.in
	sed "s/@VERSION@/${VERSION}/g" < $< > $@

# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: install install_debug dist clean \
lint benchmark \
build-linux \
format check-ledger \
go-mod-cache