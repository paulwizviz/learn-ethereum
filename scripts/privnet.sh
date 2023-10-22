#!/bin/bash

export ETH_ADMIN_IMAGE_NAME="ethereum/client-go:alltools-stable"
export ETH_NODE_IMAGE_NAME="ethereum/client-go:stable"
export NETWORK_NAME="learn-ethereum_private"

COMMAND=$1

case $COMMAND in
    "admin")
        docker-compose -f ./deployments/private/docker-compose.yaml run -w /root admin /bin/sh
        ;;
    "node1")
        docker-compose -f ./deployments/private/docker-compose.yaml run -w /root node1
        ;;
    "clean")
        docker rm -f $(docker ps -a)
        docker rmi -f ${ETH_ADMIN_IMAGE_NAME}
        docker rmi -f ${ETH_NODE_IMAGE_NAME}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        ;;
    *)
        echo "Usage; $0 <Command>"
        echo
        echo "Command:"
        echo " admin   open a shell in admin container"
        echo " build   image"
        echo " clean   images and containers"
        exit 1
        ;;
esac