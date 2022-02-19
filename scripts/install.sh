#!/bin/bash
set -e

CURRENT_DIR=$(pwd)
UTILS_VER=$(utilsVersion 2> /dev/null || echo "")

# Installing utils is essential to simplify the setup steps
if [ -z "$UTILS_VER" ] ; then
    echo "INFO: KIRA utils were NOT installed on the system, setting up..."
    KIRA_UTILS_BRANCH="v0.0.1"
    cd /tmp && rm -fv ./i.sh
    wget https://raw.githubusercontent.com/KiraCore/tools/$KIRA_UTILS_BRANCH/bash-utils/install.sh -O ./i.sh
    chmod 555 -v ./i.sh && ./i.sh "$KIRA_UTILS_BRANCH" "/var/kiraglob"
fi

# navigate to current direcotry and load global environment variables
cd $CURRENT_DIR
loadGlobEnvs

if ($(isNullOrEmpty "$SEKAID_BRANCH")) ; then
    SEKAID_BRANCH="master"
    echoWarn "WARNING: SEKAI branch 'SEKAID_BRANCH' env variable was undefined, the '$SEKAID_BRANCH' branch will be used during installation process!" && sleep 1
    setGlobEnv SEKAID_BRANCH "$SEKAID_BRANCH"
fi

if ($(isNullOrEmpty "$GOBIN")) ; then
    GOBIN=${HOME}/go/bin
    echoWarn "WARNING: GOBIN env variable was undefined, the '$GOBIN' will be used during installation process!" && sleep 1
fi

go get "github.com/KiraCore/sekai@${SEKAID_BRANCH}"
go build -o "${GOBIN}/interxd"

