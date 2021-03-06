#!/usr/bin/env bash

REDIS_PARAMS=$(echo $VCAP_SERVICES | jq '.[] | .[] | select(.tags[] | contains("redis"))')

export REDIS_HOST=$(echo $REDIS_PARAMS | jq '.credentials.host' | tr -d '"')
export REDIS_PORT=$(echo $REDIS_PARAMS | jq '.credentials.port')
export REDIS_PASSWORD=$(echo $REDIS_PARAMS | jq '.credentials.password' | tr -d '"')

envconfig-config

