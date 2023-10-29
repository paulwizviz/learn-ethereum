#!/bin/bash

mkdir -p node1
mkdir -p node2

geth --datadir node1 account new
geth --datadir node2 account new

chown -R 1000:1000 *