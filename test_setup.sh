#!/bin/sh

echo "Restore vendors"

gvt restore

mkdir -p /go/src/github.com/postgres-ci/app-server/

cp -r . /go/src/github.com/postgres-ci/app-server/