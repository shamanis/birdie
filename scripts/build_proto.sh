#!/bin/bash

protoc -I . api/birdie.proto --go_out=. --go-grpc_out=.
git add internal/app/birdie/birdie.pb.go
git add internal/app/birdie/birdie_grpc.pb.go
