#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="KIRA-STATUS-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

INTERX_GATEWAY="127.0.0.1:11000"
STATUS_NODE_ID_CLI=$(sekaid status | jq '.NodeInfo.id' || echo "0")
STATUS_NODE_ID_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/status" | jq '.node_info.id' || echo "0")

[[ $STATUS_NODE_ID_CLI != $STATUS_NODE_ID_INTERX ]] && echoErr "ERROR: Expected node id of the status to be '$STATUS_NODE_ID_CLI', but got '$STATUS_NODE_ID_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"