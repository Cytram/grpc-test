#!/bin/bash

protoc -I proto/ --go_out=plugins=grpc:proto proto/process/processor.proto
python -m grpc_tools.protoc -I.  --python_out=. --grpc_python_out=. proto/process/processor.proto
