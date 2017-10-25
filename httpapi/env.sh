#!/bin/bash

export GOPATH=$(cd ../../../ && echo $(pwd))

echo $(env | grep GOPATH)
