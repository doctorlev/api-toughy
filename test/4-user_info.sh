#!/bin/bash

source ./0_settings.sh

data='{"Username":"lev","Password":"123"}'

USER_TOKEN=$(curl -X POST -d $data $HTTPAPIURL/auth | jq -r .token)

echo "USER_TOKEN: [$USER_TOKEN]"

curl -H 'Authorization: Bearer '$USER_TOKEN -X GET $HTTPAPIURL/userinfo
