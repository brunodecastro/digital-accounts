#####################################
# Step that builds the app executable
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

# app workdir
WORKDIR /go/app

# copy go.mo and go.sum to download dependencies
COPY go.mod go.sum ./

# get go dependencies
RUN go mod download

# copy project to container
COPY . .

# Test before running
# RUN go test -v ./...
#RUN make test

# build the go app
RUN go build -o /build/digital-accounts github.com/brunodecastro/digital-accounts/cmd
#RUN make build


###################################
# Step that runs the app executable
FROM alpine:latest

LABEL maintainer="Bruno de Castro Oliveira <brunnodecastro@gmail.com>"

# app workdir
WORKDIR /app

# copy the build app to final image
COPY --from=builder /go/app/build/digital-accounts /digital-accounts

# Expose port to the outside world
EXPOSE $APP_SERVER_PORT

# run the app executable
ENTRYPOINT ["./digital-accounts"]

