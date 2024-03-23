ARG OS_VER

FROM ubuntu:${OS_VER}

ARG GO_VER
ARG GETH_VER

ENV GOROOT=/usr/local/go
ENV GOPATH=/opt/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN apt-get update && \
    apt-get -y install wget gcc g++ make git nodejs npm protobuf-compiler && \
    cd /tmp; wget https://go.dev/dl/go${GO_VER}.linux-amd64.tar.gz && \
    tar -xvf go${GO_VER}.linux-amd64.tar.gz; mv go /usr/local/ && \
    mkdir /opt/workspace; cd /opt/workspace && \
    git clone  --depth 1 --branch ${GETH_VER} https:///github.com/ethereum/go-ethereum; cd /opt/workspace/go-ethereum && \
    make all; ln -s /opt/workspace/go-ethereum/build/bin/* /usr/local/bin/
