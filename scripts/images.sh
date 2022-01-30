#!/bin/bash

export GETH_CLIENT_IMAGE_NAME="learn-ethereum/gethclient:current"

COMMAND=$1
SUBCOMMAND=$2

function gethclient(){
    local cmd=$1
    case $cmd in
        "build")
            docker-compose -f ./build/builder.yaml build gethclient
            ;;
        "clean")
            docker rmi -fCLIENT}
            docker rmi -f $(docker images --filter "dangling=true" -q)
            ;;
        "run")
            docker run -it --rm -w /root ${GETH_CLIENT_IMAGE_NAME} /bin/bash
            ;;
        *)
            echo "Usage; $cmd <Command>"
            echo
            echo "Command:"
            echo " build      image"
            echo " clean      images"
            echo " run        images"
            exit 1
            ;;
    esac
}

case $COMMAND in

    "gethclient")
        gethclient $SUBCOMMAND
        ;;
    *)
        echo "Usage: $0 <Command>"
        echo
        echo "Command:"
        echo " gethclient  build or clean gethtool"
        exit 1
        ;;
esac