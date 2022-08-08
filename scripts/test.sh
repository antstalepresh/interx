#!/usr/bin/env bash
set -e
set -x
. /etc/profile

echo "INFO: Cleaning up system resources"
go test -mod=readonly $(go list ./gateway/...) 
