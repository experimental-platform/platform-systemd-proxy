#!/bin/bash
# THIS ONLY WORK IN OUR CI!
docker run --rm -v /data/jenkins/jobs/${JOB_NAME}/workspace:/usr/src/sproxy -w /usr/src/sproxy golang:1.4 /bin/bash -c 'go get -d && go build -v'
