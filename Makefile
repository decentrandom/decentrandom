DEP := $(shell command -v dep 2> /dev/null)

get_tools:
ifndef DEP
	@echo "Installing dep"
	go get -u -v github.com/golang/dep/cmd/dep
else
	@echo "Dep is already installed..."
endif

get_vendor_deps:
	@echo "--> Generating vendor directory via dep ensure"
	@rm -rf .vendor-new
	@dep ensure -v -vendor-only

update_vendor_deps:
	@echo "--> Running dep ensure"
	@rm -rf .vendor-new
	@dep ensure -v -update

build: update_rand_lite_docs
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/randd.exe ./cmd/randd
	go build $(BUILD_FLAGS) -o build/randcli.exe ./cmd/randcli
	go build $(BUILD_FLAGS) -o build/randkeyutil.exe ./cmd/randkeyutil
else
	go build $(BUILD_FLAGS) -o build/randd ./cmd/randd
	go build $(BUILD_FLAGS) -o build/randcli ./cmd/randcli
	go build $(BUILD_FLAGS) -o build/randkeyutil ./cmd/randkeyutil
endif

update_rand_lite_docs:
	@statik -src=client/lcd/swagger-ui -dest=client/lcd -f

install:
	go install ./cmd/randd
	go install ./cmd/randcli

build-linux:
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

build-docker-randdnode:
	$(MAKE) -C networks/local