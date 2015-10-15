#!/bin/bash
set -e

# get all dependencies
go get -d

# test the build
go build -v >/dev/stderr | true
if [[ ${PIPESTATUS[0]} -eq 0 ]]; then
    exit 0
fi

