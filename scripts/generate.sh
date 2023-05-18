#!/bin/bash

set -ex

PKG="git.vseinstrumenti.net/fd/proto"

PROTOS=$(find . -name "*.proto")

PROTOC_BIN=protoc
PROTOC_GEN_GO_BIN=$(which protoc-gen-go)
PROTOC_GEN_GO_GRPC_BIN=$(which protoc-gen-go-grpc)

for proto in $PROTOS
do
  $PROTOC_BIN \
    --go_opt=module=$PKG \
    -I/usr/include \
    --proto_path=. \
    --proto_path=./shared \
    --proto_path=$(dirname $proto) \
    --plugin=$PROTOC_GEN_GO_BIN \
    --go_out=. \
    $proto;

  $PROTOC_BIN \
    --go-grpc_opt=module=$PKG \
    -I/usr/include \
    --proto_path=. \
    --proto_path=./shared \
    --proto_path=$(dirname $proto) \
    --plugin=PROTOC_GEN_GO_GRPC_BIN \
    --go-grpc_out=. \
    $proto;
done

echo "Done"
