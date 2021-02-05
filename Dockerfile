#####################################
# Step that builds the app executable
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

# app workdir
WORKDIR /go/app

# copy go.mo and go.sum to download dependencies
COPY go.mod go.sum ./

# get go dependencies
RUN go mod download

# copy project to container
COPY . .

# build the go app
RUN go build -o build/digital-accounts cmd/main.go


###################################
# Step that runs the app executable
FROM alpine:latest

LABEL maintainer="Bruno de Castro Oliveira <brunnodecastro@gmail.com>"

# app workdir
WORKDIR /app

# copy the build app to final image
COPY --from=builder /go/app/build/digital-accounts .

# copy migration files
COPY app/persistence/database/postgres/migrations ./migrations

# copy doc files
COPY docs/ ./docs

# Expose port to the outside world
EXPOSE $PORT

# run the app executable
ENTRYPOINT ./digital-accounts

