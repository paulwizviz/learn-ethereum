ARG OS_VER

FROM ubuntu:${OS_VER} AS builder

ARG SOLC_VER

RUN apt update && \
    apt install -y sudo wget gcc g++ cmake make libboost-all-dev git z3 cvc4 nodejs npm protobuf-compiler && \
    mkdir /opt/workspace; cd /opt/workspace && \
    git clone --recursive --branch ${SOLC_VER} https://github.com/ethereum/solidity.git && \
    cd solidity && \
    ./scripts/build.sh

FROM ubuntu:${OS_VER}

COPY --from=builder /opt/workspace/solidity/build/solc/solc /usr/local/bin/solc
