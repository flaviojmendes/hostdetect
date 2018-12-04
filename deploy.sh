#!/bin/sh
curl -H 'Content-Type: application/json' --data '{"build": true}' -X POST https://registry.hub.docker.com/u/brfutebol/hostdetect/trigger/$DOCKERHUB_TOKEN/