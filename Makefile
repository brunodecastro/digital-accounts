.PHONY: build
build:
	@echo "Building Digital Accounts Api"
	go build -o build/digital-accounts cmd/main.go

.PHONY: run
run:
	@echo "Running Digital Accounts Api"
	go run cmd/main.go

.PHONY: test
test:
	@echo "Running tests"
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage"
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: start
start:
	@docker-compose -f docker/docker-compose.yml up -d

.PHONY: start-build
start-build:
	@docker-compose -f docker/docker-compose.yml up -d --build

.PHONY: stop
stop:
	@docker-compose -f docker/docker-compose.yml down





