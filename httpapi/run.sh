#!/bin/bash

export GOPATH=$(cd ../../../ && echo $(pwd))

echo $(env | grep GOPATH)

# vendoring:
#go get github.com/go-redis/redis
#go get github.com/google/uuid

APPNAME="resapi"

# compiling go file
go fmt ./... && go build -o bin/main.bin

#kill and build the image with go app
docker kill httpmanage
docker rm httpmanage

docker rmi -f $APPNAME
docker build -t $APPNAME .

# deploying container (with the name 'httpmanage')
docker run -d --name httpmanage -p 8080:8082 $APPNAME

# starting logs:
docker logs -f httpmanage
