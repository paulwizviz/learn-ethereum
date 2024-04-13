#!/bin/bash

function compile_sol(){
    local sol="HelloWorld.sol"
    docker run -v $(PWD)/solidity/$sol:/opt/solidity/$sol \
            -v $(PWD)/tmp/abi/:/opt/abi \
            ethereum/solc:stable --abi --bin /opt/solidity/$sol -o /opt/abi
}

function gen_abi(){
    local abi="HelloWorld.abi"
    local pkg="hello"
    local type="HelloWorld"
    docker run -v ${PWD}/tmp/abi/$abi:/opt/abi/$abi \
               -v ${PWD}/internal/$pkg:/opt/internal/$pkg \
               ${GETH_NODE_IMAGE_NAME} abigen --abi /opt/abi/$abi --pkg $pkg --type $type --out /opt/internal/$pkg/"${pkg}.go"

}

