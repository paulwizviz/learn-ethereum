ARG GO_VER
ARG OS_VER

FROM golang:${GO_VER}${OS_VER} AS builder

ARG GETH_VER

RUN apk update && \
    apk add git make && \
    cd /opt/workspace && \
    git clone  --depth 1 --branch ${GETH_VER} https:///github.com/ethereum/go-ethereum; cd /opt/workspace/go-ethereum && \
    go mod tidy; make all && \
    cd /opt/workspace/genesis; go mod tidy; go build -o ./build/genesis ./cmd/genesis

FROM alpine:${OS_VER}

COPY --from=builder /opt/workspace/go-ethereum/build/bin/* /usr/local/bin