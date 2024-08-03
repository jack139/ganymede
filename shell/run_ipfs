#!/bin/bash

echo "===============开启系统服务================="
export LIBP2P_FORCE_PNET=1

echo "===============开启IPFS服务================="
nohup ipfs daemon --enable-namesys-pubsub > /tmp/ipfs.log 2>&1 &
