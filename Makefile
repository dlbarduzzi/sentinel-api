.PHONY: run
run:
	@go run ./cmd/sentinel

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: lint
lint:
	@golangci-lint run -c ./.golangci.yml ./...

.PHONY: test
test:
	@go test -count=1 ./... -v --cover --coverprofile=coverage.out

.PHONY: test/coverage
test/coverage: test
	@go tool cover -html=coverage.out
