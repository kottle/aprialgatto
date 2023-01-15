#!/bin/bash

GOPTH=/Users/eliofrancesconi/.go
BINPATH=$GOPTH/bin

SRC=.
SRC_DIR=.
DST_DIR=../
protoc -I=$SRC --plugin=$BINPATH/protoc-gen-go  --go-grpc_out=$DST_DIR $SRC_DIR/detection.proto
protoc -I=$SRC --plugin=$BINPATH/protoc-gen-go  --go_out=$DST_DIR  $SRC_DIR/detection.proto
