FROM ubuntu:18.04

ENV GOROOT=/usr/local/go
ENV GOPATH=/opt/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN apt-get update && \
    apt-get -y install wget gcc g++ make git nodejs npm protobuf-compiler && \
    cd /tmp; wget https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz && \
    tar -xvf go1.15.6.linux-amd64.tar.gz; mv go /usr/local/ && \
    mkdir /opt/workspace; cd /opt/workspace && \
    git clone  --depth 1 --branch v1.10.15 https:///github.com/ethereum/go-ethereum; cd /opt/workspace/go-ethereum && \
    make all; ln -s /opt/workspace/go-ethereum/build/bin/* /usr/local/bin/
