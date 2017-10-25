#!/bin/bash

source ./0_settings.sh

echo "Expecting: success"
data='{"KeyName":"lev"}'
curl -X GET -d $data $HTTPAPIURL/users 2>/dev/null | jq .

echo "Expecting: failure"
data='{"KeyName":"unknownuser"}'
curl -X GET -d $data $HTTPAPIURL/users 2>/dev/null | jq .
