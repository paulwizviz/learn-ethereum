#!/bin/bash

export GETH_CLIENT_IMAGE_NAME="learn-ethereum/gethclient:current"
export GETH_CLIENT_CONTAINER_NAME="gethcontainer"

COMMAND=$1

case $COMMAND in
    "build")
        docker-compose -f ./build/builder.yaml build
        ;;
    "clean")
        docker rm -f ${GETH_CLIENT_CONTAINER_NAME} 
        docker rmi -f ${GETH_CLIENT_IMAGE_NAME}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    "shell")
        docker run -it --name ${GETH_CLIENT_CONTAINER_NAME} --rm -w /root ${GETH_CLIENT_IMAGE_NAME} /bin/bash
        ;;
    *)
        echo "Usage; $0 <Command>"
        echo
        echo "Command:"
        echo " build   image"
        echo " clean   images and containers"
        echo " shell   open a shell in a running container"
        exit 1
        ;;
esac