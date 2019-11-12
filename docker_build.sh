#!/bin/bash

docker build -f Dockerfile.image-processor -t cytram/image-processor . && docker push cytram/image-processor
docker build -f Dockerfile.api-service -t cytram/api-service . && docker push cytram/api-service
