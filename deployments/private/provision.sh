#!/bin/bash

mkdir -p network/node1
mkdir -p network/node2

geth --datadir network/node1 account new
geth --datadir network/node2 account new

chown -R 1000:1000 network