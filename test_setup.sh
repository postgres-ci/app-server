#!/bin/sh

echo "Get gvt (https://github.com/FiloSottile/gvt)"

go get -u github.com/FiloSottile/gvt

echo "Restore vendors"

gvt restore
mkdir -p /go/src/github.com/postgres-ci/app-server
cp -r . /go/src/github.com/postgres-ci/app-server/