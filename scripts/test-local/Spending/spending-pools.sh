#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="SPENDING-POOLS-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

INTERX_GATEWAY="127.0.0.1:11000"
TOTAL_POOLS_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/spending-pools" | jq '.names | length' || echo "0")
TOTAL_POOLS_CLI=$(sekaid query spending pool-names --output=json | jq '.names | length' || echo "0")

[ $TOTAL_POOLS_CLI != $TOTAL_POOLS_INTERX ] && echoErr "ERROR: Expected number of pools to be '$TOTAL_POOLS_CLI', but got '$TOTAL_POOLS_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"