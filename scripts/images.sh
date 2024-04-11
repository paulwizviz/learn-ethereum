#!/bin/bash

export ETH_NODE_IMAGE_NAME="learn-ethereum/ethnode:current"
export SOLC_IMAGE_NAME="learn-ethereum/solc:current"

function build_image() {
    docker-compose -f ./build/builder.yaml build
}

function clean_image(){
    docker rmi -f ${ETH_NODE_IMAGE_NAME}
    docker rmi -f ${SOLC_IMAGE_NAME}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}