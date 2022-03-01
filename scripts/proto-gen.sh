#!/usr/bin/env bash
set -e
set -x
. /etc/profile

CURRENT_DIR=$(pwd)
UTILS_VER=$(utilsVersion 2> /dev/null || echo "")
GO_VER=$(go version 2> /dev/null || echo "")

UTILS_OLD_VER="false" && [[ $(versionToNumber "$UTILS_VER" || echo "0") -ge $(versionToNumber "v0.0.10" || echo "1") ]] || UTILS_OLD_VER="true" 

# Installing utils is essential to simplify the setup steps
if [ "$UTILS_OLD_VER" == "true" ] ; then
    echo "INFO: KIRA utils were NOT installed on the system, setting up..." && sleep 2
    KIRA_UTILS_BRANCH="v0.0.2" && cd /tmp && rm -fv ./i.sh && \
    wget https://raw.githubusercontent.com/KiraCore/tools/$KIRA_UTILS_BRANCH/bash-utils/install.sh -O ./i.sh && \
    chmod 777 ./i.sh && ./i.sh "$KIRA_UTILS_BRANCH" "/var/kiraglob" && loadGlobEnvs
else
    echoInfo "INFO: KIRA utils are up to date, latest version $UTILS_VER" && sleep 2
fi

# install golang if needed
if  ($(isNullOrEmpty "$GO_VER")) || ($(isNullOrEmpty "$GOBIN")) ; then
    GO_VERSION="1.17.7" && ARCH=$(([[ "$(uname -m)" == *"arm"* ]] || [[ "$(uname -m)" == *"aarch"* ]]) && echo "arm64" || echo "amd64") && \
     OS_VERSION=$(uname) && GO_TAR=go${GO_VERSION}.${OS_VERSION,,}-${ARCH}.tar.gz && rm -rfv /usr/local/go && cd /tmp && rm -fv ./$GO_TAR && \
     wget https://dl.google.com/go/${GO_TAR} && \
     tar -C /usr/local -xvf $GO_TAR && rm -fv ./$GO_TAR && \
     setGlobEnv GOROOT "/usr/local/go" && setGlobPath "\$GOROOT" && \
     setGlobEnv GOBIN "/usr/local/go/bin" && setGlobPath "\$GOBIN" && \
     setGlobEnv GOPATH "/usr/home/go" && setGlobPath "\$GOPATH" && \
     setGlobEnv GOCACHE "/usr/home/go/cache" && \
     loadGlobEnvs && \
     mkdir -p "$GOPATH/src" "$GOPATH/bin" "$GOCACHE" && \
     chmod -R 777 "$GOPATH" && chmod -R 777 "$GOROOT" && chmod -R 777 "$GOCACHE"

    echoInfo "INFO: Sucessfully intalled $(go version)"
fi

# navigate to current direcotry and load global environment variables
cd $CURRENT_DIR
loadGlobEnvs

go clean -modcache
EXPECTED_INTERX_PROTO_DEP_VER="v0.0.2"
BUF_VER=$(buf --version 2> /dev/null || echo "")

if ($(isNullOrEmpty "$BUF_VER")) || [ "$INTERX_PROTO_DEP_VER" != "$EXPECTED_INTERX_PROTO_DEP_VER" ] ; then
    GO111MODULE=on 
    go install github.com/bufbuild/buf/cmd/buf@v1.0.0-rc10
    echoInfo "INFO: Sucessfully intalled buf $(buf --version)"

    setGlobEnv GOLANG_PROTOBUF_VERSION "1.27.1" && \
     setGlobEnv GOGO_PROTOBUF_VERSION "1.3.2" && \
     setGlobEnv GRPC_GATEWAY_VERSION "1.14.7" && \
     loadGlobEnvs

    go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@latest && \
     go install google.golang.org/protobuf/cmd/protoc-gen-go@v${GOLANG_PROTOBUF_VERSION} && \
     go install github.com/gogo/protobuf/protoc-gen-gogo@v${GOGO_PROTOBUF_VERSION} && \
     go install github.com/gogo/protobuf/protoc-gen-gogofast@v${GOGO_PROTOBUF_VERSION} && \
     go install github.com/gogo/protobuf/protoc-gen-gogofaster@v${GOGO_PROTOBUF_VERSION} && \
     go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v${GRPC_GATEWAY_VERSION} && \
     go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v${GRPC_GATEWAY_VERSION} && \
     go install github.com/gogo/protobuf/protoc-gen-gogotypes

    # Following command executes with error requiring us to silence it, however the executable is placed in $GOBIN
    # https://github.com/regen-network/cosmos-proto
    # Original setup originates from Docker Image tendermintdev/sdk-proto-gen:v0.2
    # reference: 
    go install github.com/regen-network/cosmos-proto/protoc-gen-gocosmos@v0.3.1 2> /dev/null || : 
    go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
    setGlobEnv INTERX_PROTO_DEP_VER "$EXPECTED_INTERX_PROTO_DEP_VER"
fi

COSMOS_BRANCH="v0.45.1"
SEKAI_BRANCH="master"

go get github.com/KiraCore/sekai@$SEKAI_BRANCH
go get github.com/cosmos/cosmos-sdk@$COSMOS_BRANCH

echoInfo "Cleaning up proto gen files..."
rm -rfv ./proto-gen
mkdir -p ./proto-gen ./proto
kira_dir=$(go list -f '{{ .Dir }}' -m github.com/KiraCore/sekai@$SEKAI_BRANCH)
cosmos_sdk_dir=$(go list -f '{{ .Dir }}' -m github.com/cosmos/cosmos-sdk@$COSMOS_BRANCH)

rm -rfv ./proto/cosmos ./proto/kira ./third_party/proto
mkdir -p ./third_party/proto
cp -rfv $cosmos_sdk_dir/proto/cosmos ./proto
cp -rfv $cosmos_sdk_dir/third_party/proto/* ./third_party/proto
cp -rfv $kira_dir/proto/kira ./proto

### This part is required by gocosmos_out
rm -rfv ./codec && mkdir -p codec/types
buf protoc -I "third_party/proto" --gogotypes_out=./codec/types third_party/proto/google/protobuf/any.proto
mv codec/types/google/protobuf/any.pb.go codec/types
rm -rfv codec/types/third_party
rm -rfv ./third_party/proto/gogoproto
rm -rfv ./third_party/proto/google
###

sed '/proto\.RegisterType/d' codec/types/any.pb.go > tmp && mv tmp codec/types/any.pb.go

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

echoInfo "Generating protobuf files..."

for dir in $proto_dirs; do
    proto_fils=$(find "${dir}" -maxdepth 1 -name '*.proto') 
    for fil in $proto_fils; do
        buf protoc \
          -I "./proto" \
          -I third_party/grpc-gateway/ \
		  -I third_party/googleapis/ \
		  -I third_party/proto/ \
          --go_out=paths=source_relative:./proto-gen \
          --go-grpc_out=paths=source_relative:./proto-gen \
          --grpc-gateway_out=logtostderr=true,paths=source_relative:./proto-gen \
          $fil || ( echoErr "ERROR: Failed proto build for: ${fil}" && sleep 2 && exit 1 )
    done
done

#TODO: Currently it is not possible for go to dicover the gocosmos_out plugin (might require to resolve some issues with path)
#--gocosmos_out=plugins=interfacetype+grpc,\
#Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:./proto-gen \

echoInfo "INFO: Success, all proto files were compiled!"
