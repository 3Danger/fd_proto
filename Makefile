OS = $(shell uname | tr A-Z a-z)
export PATH := $(abspath bin/):${PATH}

GOLANGCI_VERSION ?= 1.52.2

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint
bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	curl -L https://github.com/golangci/golangci-lint/releases/download/v${GOLANGCI_VERSION}/golangci-lint-${GOLANGCI_VERSION}-${OS}-amd64.tar.gz | tar -zOxf - golangci-lint-${GOLANGCI_VERSION}-${OS}-amd64/golangci-lint > ./bin/golangci-lint-${GOLANGCI_VERSION} && chmod +x ./bin/golangci-lint-${GOLANGCI_VERSION}

tidy:
	go mod tidy

.PHONY: lint
lint: bin/golangci-lint ## Запуск линтеров
	bin/golangci-lint run

generate:
	bash ./scripts/generate.sh
	go generate ./...

install-mockery:
	go install github.com/vektra/mockery/v2@v2.32.4

