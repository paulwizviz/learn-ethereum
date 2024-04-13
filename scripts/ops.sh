#!/bin/bash

if [ "$(basename $(realpath .))" != "learn-ethereum" ]; then
    echo "You are outside the scope of the project"
    exit 0
else
    . ./scripts/images.sh
    . ./scripts/network.sh
    . ./scripts/solc.sh
fi

COMMAND=$1
SUBCOMMAND1=$2
SUBCOMMAND2=$3

function image() {
    local cmd=$1
    case $cmd in
        "build:node")
            build_node
            ;;
        "build:solc")
            build_solc
            ;;
        "clean")
            clean_image
            ;;
        *)
            echo "Usage: $0 image [comand]

command:
    build:node   build geth and related images
    build:solc   build image of solidity compiler from source
    clean        all images"
            ;;
    esac
}

function network(){
    local cmd=$1
    case $cmd in
        "private")
            private_network $SUBCOMMAND2
            ;;
        *)
            echo "Usage: $0 network [command]

command:
    private    operations related to private network"
            ;;
    esac
}

function solidity(){
    local cmd=$1
    case $cmd in
        "compile")
            compile_sol
            ;;
        "abi")
            gen_abi
            ;;
        *)
            echo "$0 solidity [command]

command:
    compile   solidity file in ./solidity
    abi       generate ABI from solidity file"
            ;;
    esac
}

case $COMMAND in
    "image")
        image $SUBCOMMAND1
        ;;
    "network")
        network $SUBCOMMAND1
        ;;
    "solidity")
        solidity $SUBCOMMAND1
        ;;
    *)
        echo "Usage: $0 <command>

command:
    image      operation to build or clean images
    network    operations related to network
    solidity   operations related to solidity file"
        ;;
esac