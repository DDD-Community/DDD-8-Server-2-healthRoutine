GO_BIN_PATH := $(shell go env GOPATH)/bin

SQLC_CMD := $(GO_BIN_PATH)/sqlc

install:
	go mod download all
	go get -d github.com/kyleconroy/sqlc/cmd/sqlc@v1.16.0
	go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.16.0

install-oas-mac:
	brew install openapi-generator

versions:
	$(SQLC_CMD) version

sqlc:
	$(SQLC_CMD) generate

oas:
	cd $(PROJECT_DIR)/docs && openapi-generator generate -i docs.yaml -g openapi-yaml -o gen

.PHONY: install sqlc install-oas-mac oas