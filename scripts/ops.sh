#!/bin/bash

if [ "$(basename $(realpath .))" != "learn-ethereum" ]; then
    echo "You are outside the scope of the project"
    exit 0
else
    . ./scripts/images.sh
    . ./scripts/network.sh
fi

COMMAND=$1
SUBCOMMAND1=$2
SUBCOMMAND2=$3

function image() {
    local cmd=$1
    case $cmd in
        "build")
            build_image
            ;;
        "clean")
            clean_image
            ;;
        *)
            echo "Usage: $0 image [comand]

command:
    build   images
    clean   all images"
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

case $COMMAND in
    "image")
        image $SUBCOMMAND1
        ;;
    "network")
        network $SUBCOMMAND1
        ;;
    "geth")
        docker run -it --rm ${ETH_NODE_IMAGE_NAME} /bin/sh
        ;;
    *)
        echo "Usage: $0 <command>

command:
    image    operation to build or clean images
    network  operations related to network"
        ;;
esac