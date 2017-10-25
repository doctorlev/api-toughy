#!/bin/bash

echo "stopping redis container"

docker kill redis1

docker rm redis1
