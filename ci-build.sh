#!/bin/bash
# THIS ONLY WORK IN OUR CI!
docker run --rm -v /data/jenkins/jobs/docker-systemd-proxy/workspace:/usr/src/sproxy -w /usr/src/sproxy golang:1.4 go build -v
