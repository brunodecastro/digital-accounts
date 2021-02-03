FROM golang:latest AS builder

LABEL maintainer="Bruno de Castro Oliveira <brunnodecastro@gmail.com>"

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

# app workdir
WORKDIR /go/src/app

# copy go.mo and go.sum to download dependencies
COPY go.mod go.sum ./

# get go dependencies
RUN go mod download

# copy project to container
COPY . .

# Test before running
# RUN go test -v ./...
RUN make test

# build the go app
# RUN go build -o build/digital-accounts cmd/main.go
RUN make build

# Expose port to the outside world
EXPOSE 9090

# run the app executable
ENTRYPOINT ./build/digital-accounts