#!/bin/bash

export ETH_NODE_IMAGE_NAME="learn-ethereum/node:current"
export ETH_ADMIN_NODE="admin"
export NETWORK_NAME="learn-ethereum_private-network"

COMMAND=$1

case $COMMAND in
    "admin")
        docker run -it --rm --name $ETH_ADMIN_NODE -w /root \
            -v $PWD/deployments/private/network:/root/network \
            -v $PWD/deployments/private/provision.sh:/root/network/provision.sh \
            $ETH_NODE_IMAGE_NAME /bin/bash
        ;;
    "boot")
        docker-compose -f ./deployments/private/docker-compose.yaml run -w /root bootnode /bin/bash
        ;;
    "node1")
        docker-compose -f ./deployments/private/docker-compose.yaml run -w /root node1 /bin/bash
        ;;
    "node2")
        docker-compose -f ./deployments/private/docker-compose.yaml run -w /root node2 /bin/bash
        ;;
    "build")
        docker-compose -f ./build/ethnode/builder.yaml build
        ;;
    "clean")
        docker rm -f ${ETH_ADMIN_NODE}
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
        ;;
esac