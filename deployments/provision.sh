#!/bin/sh

mkdir -p node1
mkdir -p node2

geth account new --password ./password --datadir node1
geth account new --password ./password --datadir node2
