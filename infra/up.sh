#!/bin/bash

#start redis

docker run -d --name redis1 -p 6379:6379 redis
