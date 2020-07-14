#!/usr/bin/make -f

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building relayer..."
	@go build -mod=readonly -o build/ relayer/cmd/relayer
else
	@echo "building relayer..."
	@go build -mod=readonly -o build/ relayer/cmd/relayer
endif

install: go.sum
	@echo "installing relayer..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/ relayer/cmd/relayer
