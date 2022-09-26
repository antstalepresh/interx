#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="VALIDATOR-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

RESULT_NUM=$(sekaid query customstaking validators --moniker="" --output=json | jq '.validators | length' 2> /dev/null || exit 1)

INTERX_GATEWAY="127.0.0.1:11000"
RESULT_FROM_INTERX_NUM=$(curl --fail $INTERX_GATEWAY/api/valopers | jq '.validators | length' || exit 1)

[ $RESULT_NUM != $RESULT_FROM_INTERX_NUM ] && echoErr "ERROR: Expected validator amount to be '$RESULT_NUM', but got '$RESULT_FROM_INTERX_NUM'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"