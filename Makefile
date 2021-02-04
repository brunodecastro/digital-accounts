
# A valid GOPATH is required to use the `go get` command.
# If $GOPATH is not specified, $HOME/go will be used by default
GOPATH := $(if $(GOPATH),$(GOPATH),~/go)


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
	go test -v -cover ./...

.PHONY: test-coverage-report
test-coverage-report:
	@echo "Running tests with coverage report"
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: fmt
fmt:
	@echo "Running go fmt"
	go fmt ./...

.PHONY: vet
vet:
	@echo "Running go vet"
	go vet ./...

.PHONY: review-code-and-test
review-code-and-test: fmt vet test

.PHONY: build-swagger
build-swagger:
	@echo "Building swagger doc files"
	swag init -g cmd\main.go

.PHONY: start
start:
	@docker-compose -f docker/docker-compose.yml up -d

.PHONY: start-build
start-build:
	@docker-compose -f docker/docker-compose.yml up -d --build

.PHONY: stop
stop:
	@docker-compose -f docker/docker-compose.yml down


install-swag:
ifeq (,$(wildcard test -f $(GOPATH)/bin/swag))
	@echo "  >  Installing swagger"
	@-bash -c "go get github.com/swaggo/swag/cmd/swag"
endif

swag: install-swag
	@bash -c "$(GOPATH)/bin/swag init --parseDependency -g ./cmd/main.go"



