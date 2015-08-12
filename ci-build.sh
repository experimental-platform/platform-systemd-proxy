#!/bin/bash
set -e

SRC_PATH=$(pwd)

docker run --volume=${SRC_PATH}:/usr/src/sproxy --workdir=/usr/src/sproxy golang:1.4 /usr/src/sproxy/patch_dbus.sh
