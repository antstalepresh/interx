#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="ACCOUNT-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

ACCOUNT_ADDRESS=$(showAddress validator)

RESULT_FROM_SEKAI=$(sekaid query account "$ACCOUNT_ADDRESS" --output=json)
RESULT_SUM_FROM_SEKAI="$(echo $RESULT_FROM_SEKAI | jq '.account_number')$(echo $RESULT_FROM_SEKAI | jq '."@type"')$(echo $RESULT_FROM_SEKAI | jq '.address')"

INTERX_GATEWAY="127.0.0.1:11000"

RESULT_FROM_INTERX=$(curl --fail $INTERX_GATEWAY/api/cosmos/auth/accounts/$ACCOUNT_ADDRESS | jq '.account' || echo "error")
RESULT_SUM_FROM_INTERX="$(echo $RESULT_FROM_INTERX | jq '.accountNumber')$(echo $RESULT_FROM_INTERX | jq '."@type"')$(echo $RESULT_FROM_INTERX | jq '.address')"

[ "$RESULT_SUM_FROM_SEKAI" != "$RESULT_SUM_FROM_INTERX" ] && echoErr "ERROR: Expected validator account info to be '$RESULT_SUM_FROM_SEKAI', but got '$RESULT_SUM_FROM_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"