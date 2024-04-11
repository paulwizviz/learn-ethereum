#!/bin/bash

if [ "$(basename $(realpath .))" != "learn-ethereum" ]; then
    echo "You are outside the scope of the project"
    exit 0
else
    . ./scripts/images.sh
fi

COMMAND=$1

case $1 in
    "compile")
        sol="HelloWorld.sol"
        docker run -v $(PWD)/solidity/$sol:/opt/solidity/$sol \
                -v $(PWD)/tmp/abi/:/opt/abi \
                ethereum/solc:stable --abi --bin /opt/solidity/$sol -o /opt/abi
        ;;
    "abigen")
        abi="HelloWorld.abi"
        pkg="hello"
        type="HelloWorld"
        docker run -v ${PWD}/tmp/abi/$abi:/opt/abi/$abi \
        -v ${PWD}/internal/$pkg:/opt/internal/$pkg \
        ${ETH_NODE_IMAGE_NAME} abigen --abi /opt/abi/$abi --pkg $pkg --type $type --out /opt/internal/$pkg/helloworld.go 
        ;;
    *)
    echo "$0 [command]
    
command:
    compile    solidity file" 
        ;;
esac
