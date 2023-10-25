# Builder
FROM ubuntu:22.04 AS builder

ENV GOROOT=/usr/local/go
ENV GOPATH=/opt/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

COPY ./cmd /opt/workspace/genesis/cmd
COPY ./internal /opt/workspace/genesis/internal
COPY ./go.mod   /opt/workspace/genesis/go.mod
COPY ./go.sum   /opt/workspace/genesis/go.sum

RUN apt-get update && \
    apt-get -y install wget gcc g++ make git nodejs npm protobuf-compiler && \
    cd /tmp; wget https://dl.google.com/go/go1.21.3.linux-amd64.tar.gz && \
    tar -xvf go1.21.3.linux-amd64.tar.gz; mv go /usr/local/ && \
    mkdir /opt/workspace; cd /opt/workspace && \
    git clone  --depth 1 --branch master https:///github.com/ethereum/go-ethereum; cd /opt/workspace/go-ethereum && \
    go mod tidy; make all && \
    cd /opt/workspace/genesis; go mod tidy; go build -o ./build/genesis ./cmd/genesis

FROM ubuntu:22.04

COPY --from=builder /opt/workspace/go-ethereum/build/bin/* /usr/local/bin
COPY --from=builder /opt/workspace/genesis/build/genesis /usr/local/bin/genesis