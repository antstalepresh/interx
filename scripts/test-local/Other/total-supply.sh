#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="TOTAL-SUPPLY-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

INTERX_GATEWAY="127.0.0.1:11000"
AMOUNT_SUPPLY_CLI=$(sekaid query bank total --output=json | jq '.supply[5].amount' || echo "0")
AMOUNT_SUPPLY_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/supply" | jq '.supply[5].amount' || echo "0")

[[ $AMOUNT_SUPPLY_CLI != $AMOUNT_SUPPLY_INTERX ]] && echoErr "ERROR: Expected amount to be '$AMOUNT_SUPPLY_CLI', but got '$AMOUNT_SUPPLY_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"