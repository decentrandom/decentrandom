include Makefile.ledger
all: install
install: go.sum
		GO111MODULE=on go install -tags "$(build_tags)" ./cmd/randd
		GO111MODULE=on go install -tags "$(build_tags)" ./cmd/randcli
go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/randd.exe ./cmd/gaia/cmd/randd
	go build -mod=readonly $(BUILD_FLAGS) -o build/randcli.exe ./cmd/gaia/cmd/randcli
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/randd ./cmd/gaia/cmd/randd
	go build -mod=readonly $(BUILD_FLAGS) -o build/randcli ./cmd/gaia/cmd/randcli
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build