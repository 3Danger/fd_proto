generate:
	bash ./scripts/generate.sh
	go generate ./...

install-mockery:
	go install github.com/vektra/mockery/v2@v2.32.4

