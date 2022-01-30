#!/bin/bash

git clone https://github.com/ethereum/go-ethereum.git ~/go-ethereum

pushd ~/go-ethereum; make all; popd
