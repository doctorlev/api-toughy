#!/bin/bash

source ./0_settings.sh

echo "Getting Token"

data='{"Username":"lev","Password":"123"}'

curl -X POST -d $data $HTTPAPIURL/auth | jq .
