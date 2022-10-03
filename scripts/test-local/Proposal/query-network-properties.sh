#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="NETWORK-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

NETWORK_RESULT=$(showNetworkProperties | jq '.properties.inflation_rate' | tr -d '"' || exit 1)

INTERX_GATEWAY="127.0.0.1:11000"
NETWORK_RESULT_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/network_properties" | jq '.properties.inflation_rate' | tr -d '"' || exit 1)

[ $NETWORK_RESULT != $NETWORK_RESULT_INTERX ] && echoErr "ERROR: Expected network inflation rate to be '$NETWORK_RESULT', but got '$NETWORK_RESULT_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"