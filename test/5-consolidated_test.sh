#!/bin/bash

source ./0_settings.sh

data='{"Username":"lev","Password":"123"}'

echo "1 - Creating User (setting in Redis)"
curl -X POST -d $data $HTTPAPIURL/users | jq .

echo "2 - Checking User (exists in Redis) "
echo "Expecting: success"
data='{"KeyName":"lev"}'
curl -X GET -d $data $HTTPAPIURL/users 2>/dev/null | jq .

echo "Expecting: failure"
data='{"KeyName":"unknownuser"}'
curl -X GET -d $data $HTTPAPIURL/users 2>/dev/null | jq .


echo "3 - Getting Token"
data='{"Username":"lev","Password":"123"}'
curl -X POST -d $data $HTTPAPIURL/auth | jq .


echo "4 - User Info - autorized access"
USER_TOKEN=$(curl -X POST -d $data $HTTPAPIURL/auth | jq -r .token)
echo "USER_TOKEN: [$USER_TOKEN]"
curl -H 'Authorization: Bearer '$USER_TOKEN -X GET $HTTPAPIURL/userinfo
