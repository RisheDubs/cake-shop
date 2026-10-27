# Run from the api/ folder. On Windows use Git Bash or WSL, or run the go commands directly.

.PHONY: run test cover vet fmt lint build tidy

run:            ## Start the API locally on :8080
	go run ./cmd/server

test:           ## Run all tests
	go test ./...

cover:          ## Tests with coverage report (opens coverage.html)
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html

vet:            ## Static checks
	go vet ./...

fmt:            ## Format code
	gofmt -w .

lint:           ## Needs golangci-lint installed
	golangci-lint run

tidy:
	go mod tidy

build:          ## Linux ARM64 binary for Lambda (Graviton is cheaper)
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/bootstrap ./cmd/server

# Called by `sam build` (BuildMethod: makefile). Must output a file named "bootstrap".
build-ApiFunction:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -o $(ARTIFACTS_DIR)/bootstrap ./cmd/server
