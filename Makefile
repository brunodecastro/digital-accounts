
.PHONY: build
build:
    go build -o build/digital-accounts cmd/main.go

.PHONY: run
run:
    @docker-compose up -d

.PHONY: stop
stop:
    @docker-compose down

.PHONY: logs
logs:
    @docker-compose logs -f

test:
    go test -cover ./...

code-review:
    fmt vet test

