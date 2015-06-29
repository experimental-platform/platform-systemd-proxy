#!/bin/bash
set -e

SRC_PATH=$(pwd)

# we're calling docker from within a container with a path (/data/jenkins) mounted into this container (/var/jenkins)
# so the newly created container needs a different path (/data/jenkins).
docker run --rm -v /data${SRC_PATH#/var}:/usr/src/sproxy -w /usr/src/sproxy golang:1.4 /bin/bash -c 'go get -d && go build -v'
