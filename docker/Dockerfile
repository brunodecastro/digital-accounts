FROM golang:latest AS builder

LABEL maintainer="Bruno de Castro Oliveira <brunnodecastro@gmail.com>"

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

# app workdir
WORKDIR /go/src/app

COPY go.mod go.sum ./

# get go dependancies
RUN go mod download

# copy project to container
COPY . .

# Test before running
RUN go test -v ./...

# build the go app
RUN go build -o build/digital-accounts cmd/main.go

# Expose port to the outside world
EXPOSE 9090

# run the app executable
ENTRYPOINT ./build/digital-accounts