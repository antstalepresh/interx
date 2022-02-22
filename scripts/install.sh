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
    KIRA_UTILS_BRANCH="v0.0.1" && cd /tmp && rm -fv ./i.sh && \
    wget https://raw.githubusercontent.com/KiraCore/tools/$KIRA_UTILS_BRANCH/bash-utils/install.sh -O ./i.sh && \
    chmod 555 -v ./i.sh && ./i.sh "$KIRA_UTILS_BRANCH" "/var/kiraglob"
fi

# install golang if needed
if  ($(isNullOrEmpty "$GO_VER")) ; then
    GO_VERSION="1.17.7" && ARCH=$(([[ "$(uname -m)" == *"arm"* ]] || [[ "$(uname -m)" == *"aarch"* ]]) && echo "arm64" || echo "amd64") && \
     GO_TAR=go${GO_VERSION}.linux-${ARCH}.tar.gz && rm -rfv /usr/local/go && cd /tmp && rm -fv ./$GO_TAR && \
     wget https://dl.google.com/go/${GO_TAR} && chmod 555 -v $GO_TAR && \
     tar -C /usr/local -xvf $GO_TAR && \
     setGlobEnv GOROOT "/usr/local/go" && setGlobPath "\$GOROOT" && \
     setGlobEnv GOBIN "/usr/local/go/bin" && setGlobPath "\$GOBIN" && \
     setGlobEnv GOPATH "/home/go" && setGlobPath "\$GOPATH" && \
     setGlobEnv GOCACHE "/home/go/cache" && \
     loadGlobEnvs && \
     mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
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
go get "github.com/KiraCore/sekai@${SEKAI_BRANCH}"
# TODO: GO mod tidy requires protogen, resolve the pending issues
# go mod tidy
go build -o "${GOBIN}/interxd"
go mod verify
echoInfo "INFO: Sucessfully intalled INTERX $(interxd version)"
