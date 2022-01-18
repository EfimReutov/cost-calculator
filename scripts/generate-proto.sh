#!/bin/bash

rm -r gen/proto
mkdir -p gen/proto

protoc -I proto --go_out=. --go-grpc_out=. ./proto/**/*.proto