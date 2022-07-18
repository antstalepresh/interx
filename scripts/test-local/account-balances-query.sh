a#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="ACCOUNT-BALANCES-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)
VALIDATOR_BALANCE_FROM_SEKAI=$(showBalance validator ukex)

INTERX_GATEWAY="127.0.0.1:11000"

echoInfo "INFO: Waiting for next block to be produced..."
VALIDATOR_BALANCE_FROM_INTERX=$(curl --fail $INTERX_GATEWAY/api/cosmos/bank/balances/$VALIDATOR_ADDRESS | jq '.balances[3].amount' | cut -d '"' -f 2 || echo "0")

[ "$VALIDATOR_BALANCE_FROM_SEKAI" != "$VALIDATOR_BALANCE_FROM_INTERX" ] && echoErr "ERROR: Expected validator account balance to be '$VALIDATOR_BALANCE_FROM_SEKAI', but got '$VALIDATOR_BALANCE_FROM_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"