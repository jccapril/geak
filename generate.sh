#!/usr/bin/env sh

export PATH=$PATH:./Tools/

rm -rf model/*

Tools/protoc \
--proto_path=Dependency/jeak_proto/ \
--plugin=Tools/protoc-gen-go-grpc \
--go-grpc_out=model \
--plugin==Tools/protoc-gen-go \
--go_out=model \
-I Dependency/jeak_proto \
Dependency/jeak_proto/*.proto

