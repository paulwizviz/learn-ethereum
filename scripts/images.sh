#!/bin/bash

export GETH_NODE_IMAGE_NAME="learn-ethereum/gethnode:current"
export SOLC_IMAGE_NAME="learn-ethereum/solc:current"
export GETH_TOOLS_IMAGE_NAME="go-eth-app/gethtool:current"

function build_geth() {
    docker-compose -f ./build/node/builder.yaml build
}

function build_solc(){
    docker-compose -f ./build/solc/builder.yaml build
}

function clean_image(){
    docker rmi -f ${GETH_NODE_IMAGE_NAME}
    docker rmi -f ${SOLC_IMAGE_NAME}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}