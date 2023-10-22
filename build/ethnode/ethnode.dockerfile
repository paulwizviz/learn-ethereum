# Builder
FROM ubuntu:22.04 AS builder

ENV GOROOT=/usr/local/go
ENV GOPATH=/opt/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN apt-get update && \
    apt-get -y install wget gcc g++ make git nodejs npm protobuf-compiler && \
    cd /tmp; wget https://dl.google.com/go/go1.21.0.linux-amd64.tar.gz && \
    tar -xvf go1.21.0.linux-amd64.tar.gz; mv go /usr/local/ && \
    mkdir /opt/workspace; cd /opt/workspace && \
    git clone  --depth 1 --branch v1.10.15 https:///github.com/ethereum/go-ethereum; cd /opt/workspace/go-ethereum && \
    make all

# Node
FROM ubuntu:22.04

COPY --from=builder /opt/workspace/go-ethereum/build/bin/abidump /usr/local/bin/abidump
COPY --from=builder /opt/workspace/go-ethereum/build/bin/abigen /usr/local/bin/abigen
COPY --from=builder /opt/workspace/go-ethereum/build/bin/bootnode /usr/local/bin/bootnode
COPY --from=builder /opt/workspace/go-ethereum/build/bin/checkpoint-admin /usr/local/bin/checkpoint-admin
COPY --from=builder /opt/workspace/go-ethereum/build/bin/clef /usr/local/bin/clef
COPY --from=builder /opt/workspace/go-ethereum/build/bin/devp2p /usr/local/bin/devp2p
COPY --from=builder /opt/workspace/go-ethereum/build/bin/ethkey /usr/local/bin/ethkey
COPY --from=builder /opt/workspace/go-ethereum/build/bin/evm /usr/local/bin/evm
COPY --from=builder /opt/workspace/go-ethereum/build/bin/faucet /usr/local/bin/geth
COPY --from=builder /opt/workspace/go-ethereum/build/bin/geth /usr/local/bin/abidump
COPY --from=builder /opt/workspace/go-ethereum/build/bin/p2psim /usr/local/bin/p2psim
COPY --from=builder /opt/workspace/go-ethereum/build/bin/puppeth /usr/local/bin/puppeth
COPY --from=builder /opt/workspace/go-ethereum/build/bin/rlpdump /usr/local/bin/rlpdump