#!/bin/bash
set -e
set -x
. /etc/profile

rm -rfv ./bin 
mkdir -p ./bin/arm64 ./bin/amd64 ./bin/deb

go mod tidy
go mod verify

# To see available archotectures, run: go tool dist list 
env GOOS=linux GOARCH=arm64 go build -o ./bin/arm64
env GOOS=linux GOARCH=amd64 go build -o ./bin/amd64

PKG_CONFIG_FILE=./nfpm.yaml && \
 PKG_CONFIG_FILE_amd64=./nfpm_amd64.yaml && \
 PKG_CONFIG_FILE_arm64=./nfpm_arm64.yaml

rm -rfv $PKG_CONFIG_FILE_amd64 && cp -v $PKG_CONFIG_FILE $PKG_CONFIG_FILE_amd64
rm -rfv $PKG_CONFIG_FILE_arm64 && cp -v $PKG_CONFIG_FILE $PKG_CONFIG_FILE_arm64

function pcgConfigure() {
    local ARCH="$1"
    local VERSION="$2"
    local FILE="$3"
    sed -i "s/\${ARCH}/$ARCH/" $FILE
    sed -i "s/\${VERSION}/$VERSION/" $FILE
}

VERSION="v0.1.18.15"

pcgConfigure "amd64" "$VERSION" $PKG_CONFIG_FILE_amd64
pcgConfigure "arm64" "$VERSION" $PKG_CONFIG_FILE_arm64

nfpm pkg --packager deb --target ./bin/deb/ -f $PKG_CONFIG_FILE_amd64
nfpm pkg --packager deb --target ./bin/deb/ -f $PKG_CONFIG_FILE_arm64

echoInfo "INFO: Sucessfully published INTERX deb packages into ./bin"
