ARG GO_VER
ARG OS_VER

# Go builder
FROM golang:${GO_VER} as builder

WORKDIR /opt

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go mod download && \
    go build -o ./build/bin/mongocmd ./cmd/simple

# App container
FROM ${OS_VER}

COPY --from=builder /opt/build/bin/mongocmd /usr/local/bin/mongocmd 