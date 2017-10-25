#!/bin/bash

source ./0_settings.sh

data='{"Username":"lev","Password":"123"}'

curl -X POST -d $data $HTTPAPIURL/users | jq .
