#!/bin/bash
set -e

SRC_PATH=$(pwd)
PROJECT_NAME="github.com/experimental-platform/platform-systemd-proxy"

docker run -v "${SRC_PATH}:/go/src/$PROJECT_NAME" -w "/go/src/$PROJECT_NAME" golang:1.4 ./patch_dbus.sh
