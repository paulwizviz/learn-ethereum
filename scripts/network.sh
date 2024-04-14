#!/bin/bash

export NETWORK_NAME=learn-ethereum_gethnet

export SHARED_VOLUME="gethnet_shared"

function geth_network(){
    local cmd=$1
    case $cmd in
        "admin")
            docker-compose -f ./deployments/gethnet.yaml run admin /bin/sh
            ;;
        "start")
            docker-compose -f ./deployments/gethnet.yaml up
            ;;
        "stop")
            docker-compose -f ./deployments/gethnet.yaml down
            ;;
        *)
            echo "Usage: $0 network gethnet [command]
    
    command:
        admin   open up an admin console
        start   gethnet network
        stop    gethnet network"
            ;;
    esac
}

function remove_volume(){
    docker volume rm ${SHARED_VOLUME}
}