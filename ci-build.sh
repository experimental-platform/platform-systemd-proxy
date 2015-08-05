#!/bin/bash
set -e

SRC_PATH=$(pwd)

# we're calling docker from within a container with a path (/data/jenkins) mounted into this container (/var/jenkins)
# so the newly created container needs a different path (/data/jenkins).
docker run --rm --volume=/data${SRC_PATH#/var}:/usr/src/sproxy --workdir=/usr/src/sproxy golang:1.4 /usr/src/sproxy/patch_dbus.sh
