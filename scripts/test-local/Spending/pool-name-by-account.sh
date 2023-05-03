#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="SPENDING-POOL-NAME-BY-ACCOUNT-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

INTERX_GATEWAY="127.0.0.1:11000"
TOTAL_POOLS_CLI=$(sekaid query spending pools-by-account $VALIDATOR_ADDRESS --output=json | jq '.pools | length' || echo "0")
TOTAL_POOLS_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/spending-pools?account=$VALIDATOR_ADDRESS" | jq '.pools | length' || echo "0")

[[ $TOTAL_POOLS_CLI != $TOTAL_POOLS_INTERX ]] && echoErr "ERROR: Expected number of pools to be '$POOL_BALATOTAL_POOLS_CLINCE_CLI', but got '$TOTAL_POOLS_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"