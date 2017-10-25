#!/bin/bash

export GOPATH=$(cd ../../../ && echo $(pwd))

echo $(env | grep GOPATH)

# vendoring:
#go get github.com/go-redis/redis
#go get github.com/google/uuid

APPNAME="storageapi"

# compiling go file
go fmt ./... && go build -o bin/main.bin

#kill and build the image with go app
docker kill storagemanage
docker rm storagemanage

docker rmi -f $APPNAME
docker build -t $APPNAME .

# deploying container (with the name 'storagemanage')
docker run -d --name storagemanage -p 8081:8082 $APPNAME

# starting logs:
docker logs -f storagemanage
