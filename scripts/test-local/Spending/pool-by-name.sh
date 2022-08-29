#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="SPENDING-POOL-BY-NAME-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

POOL_NAME=$(sekaid query spending pool-names --output=json | jq '.names[0]' | tr -d '"' || echo "")

INTERX_GATEWAY="127.0.0.1:11000"
POOL_BALANCE_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/spending-pools?name=$POOL_NAME" | jq '.pool.balance' | tr -d '"' || echo "0")
POOL_BALANCE_CLI=$(sekaid query spending pool-by-name $POOL_NAME --output=json | jq '.pool.balance' | tr -d '"' || echo "0")

[ $POOL_BALANCE_CLI != $POOL_BALANCE_INTERX ] && echoErr "ERROR: Expected number of pools to be '$POOL_BALANCE_CLI', but got '$POOL_BALANCE_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"