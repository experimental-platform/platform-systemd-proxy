#!/bin/bash
set -e

# get all dependencies
go get -d

# test the build
go build -v >/dev/stderr | true
if [[ ${PIPESTATUS[0]} -eq 0 ]]; then
    exit 0
fi

# yeah, no patch...
apt-get update
apt-get install -y patch
# an unsuccessfull build might be caused by the dbus assertion bug
# so we patch it w/ pull request #99 and then try again.
# pull request: https://github.com/coreos/go-systemd/issues/98
filename=$(go build -v 2>&1 | awk -F ':' '/dbus.go/ {print $1}')
echo "Filemname: ${filename}"
hash=$(md5sum ${filename}|awk '{print $1}')
echo "Hash: ${hash}"
if [[ ${hash} == 3b75aae9d8a95a39d84b4c2160d938a2 ]]; then
    cp 99.patch $(dirname ${filename})
    cd $(dirname ${filename})
    patch -p 1 dbus.go < 99.patch
    cd -
    go build -v
fi

