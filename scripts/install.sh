#!/bin/bash
set -e
set -x
. /etc/profile

CURRENT_DIR=$(pwd)
UTILS_VER=$(utilsVersion 2> /dev/null || echo "")
GO_VER=$(go version 2> /dev/null || echo "")

# Installing utils is essential to simplify the setup steps
if [ -z "$UTILS_VER" ] ; then
    echo "INFO: KIRA utils were NOT installed on the system, setting up..."
    KIRA_UTILS_BRANCH="v0.0.1"
    cd /tmp && rm -fv ./i.sh
    wget https://raw.githubusercontent.com/KiraCore/tools/$KIRA_UTILS_BRANCH/bash-utils/install.sh -O ./i.sh
    chmod 555 -v ./i.sh && ./i.sh "$KIRA_UTILS_BRANCH" "/var/kiraglob"
fi

# install golang if needed
if  ($(isNullOrEmpty "$GO_VER")) ; then
    GO_VERSION="1.17.7" && ARCH=$(([[ "$(uname -m)" == *"arm"* ]] || [[ "$(uname -m)" == *"aarch"* ]]) && echo "arm64" || echo "amd64")
    GO_TAR=go${GO_VERSION}.linux-${ARCH}.tar.gz && rm -rfv /usr/local/go && cd /tmp && rm -fv ./$GO_TAR
    wget https://dl.google.com/go/${GO_TAR}
    tar -C /usr/local -xvf $GO_TAR
    setGlobEnv GOROOT "/usr/local/go" && setGlobPath "\$GOROOT"
    setGlobEnv GOBIN "/usr/local/go/bin" && setGlobPath "\$GOBIN"
    setGlobEnv GOPATH "/home/go" && setGlobPath "\$GOPATH"
    setGlobEnv GOCACHE "/home/go/cache"
    loadGlobEnvs 
    echoInfo "INFO: Sucessfully intalled $(go version)"
fi

# navigate to current direcotry and load global environment variables
cd $CURRENT_DIR
loadGlobEnvs

if ($(isNullOrEmpty "$SEKAI_BRANCH")) ; then
    SEKAI_BRANCH="master"
    echoWarn "WARNING: SEKAI branch 'SEKAI_BRANCH' env variable was undefined, the '$SEKAI_BRANCH' branch will be used during installation process!" && sleep 1
    setGlobEnv SEKAI_BRANCH "$SEKAI_BRANCH"
fi

if ($(isNullOrEmpty "$GOBIN")) ; then
    GOBIN=${HOME}/go/bin
    echoWarn "WARNING: GOBIN env variable was undefined, the '$GOBIN' will be used during installation process!" && sleep 1
fi

go clean -modcache
BUF_VER=$(buf --version 2> /dev/null || echo "")
PROTOC_VER=$(protoc --version 2> /dev/null || echo "")

if ($(isNullOrEmpty "$BUF_VER")) ; then
    GO111MODULE=on go install github.com/bufbuild/buf/cmd/buf@v1.0.0
    echoInfo "INFO: Sucessfully intalled buf $(buf --version)"
fi

if ($(isNullOrEmpty "$PROTOC_VER")) ; then
    GO111MODULE=on go install github.com/gogo/protobuf/protoc-gen-gofast@v1.3.2
    echoInfo "INFO: Sucessfully intalled buf $(protoc --version)"
fi  

go get "github.com/KiraCore/sekai@${SEKAI_BRANCH}"
go get "golang.org/x/net@v0.0.0-20210903162142-ad29c8ab022f"

echoInfo "Generating protobuf files..."
rm -rfv ./proto-gen
mkdir -p ./proto-gen
proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

for dir in $proto_dirs; do
	protoc \
		-I "./proto" \
		-I third_party/grpc-gateway/ \
		-I third_party/googleapis/ \
		-I third_party/proto/ \
		--go_out=plugins=grpc,paths=source_relative:./proto-gen \
		--grpc-gateway_out=paths=source_relative:./proto-gen \
		$(find "${dir}" -maxdepth 1 -name '*.proto')
done

echoInfo "Proto files were generated for:"
echoInfo echo ${proto_dirs[*]}
sleep 1

# TODO: GO mod tidy requires protogen, resolve the pending issues
go mod tidy

go build -o "${GOBIN}/interxd"
go mod verify
echoInfo "INFO: Sucessfully intalled INTERX $(interxd version)"
