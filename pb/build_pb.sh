#!/usr/bin/env bash

protoc  \
    --proto_path=.  \
    --go_out=plugins=grpc:$GOPATH/src/github.com/zhiruchen/archi/pb \
    archi.proto