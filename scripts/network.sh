#!/bin/bash

export NETWORK_NAME=learn-ethereum_network

function private_network(){
    local cmd=$1
    case $cmd in
        "start")
            docker-compose -f ./deployments/privnet.yaml up
            ;;
        "stop")
            docker-compose -f ./deployments/privnet.yaml down
            ;;
        *)
            echo "Usage: $0 network private [command]
    
    command:
        start     private network
        stop      private network"
            ;;
    esac
}