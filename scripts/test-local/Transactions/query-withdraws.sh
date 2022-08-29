#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="TX-WITHDRAW-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

ACCOUNT_ADDRESS=$(showAddress validator)
deleteAccount testuser1
addAccount testuser1
TESTUSER_ADDRESS=$(showAddress testuser1)

TXRESULT="0x"$(sekaid tx bank send validator $TESTUSER_ADDRESS 5ukex --keyring-backend=test --chain-id=$NETWORK_NAME --fees 100ukex --output=json --yes --home=$SEKAID_HOME | txAwait 180 | jsonQuickParse "txhash" 2> /dev/null || echo "error")

INTERX_GATEWAY="127.0.0.1:11000"
RESULT_FROM_INTERX=$(curl --fail $INTERX_GATEWAY/api/withdraws?account=$ACCOUNT_ADDRESS | jq '.transactions' | jq '.["'"$TXRESULT"'"].txs[0]' || echo "error")
RESULT_ADDRESS=$(echo $RESULT_FROM_INTERX | jq '.address')
RESULT_DENOM=$(echo $RESULT_FROM_INTERX | jq '.denom')
RESULT_AMOUNT=$(echo $RESULT_FROM_INTERX | jq '.amount')

[ "$RESULT_ADDRESS" != '"'$TESTUSER_ADDRESS'"' ] && echoErr "ERROR: Expected receiver address to be '$TESTUSER_ADDRESS', but got '$RESULT_ADDRESS'" && exit 1
[ "$RESULT_DENOM" != '"ukex"' ] && echoErr "ERROR: Expected denom to be ukex, but got '$RESULT_DENOM'" && exit 1
[ "$RESULT_AMOUNT" != "5" ] && echoErr "ERROR: Expected transfer amount to be 5, but got '$RESULT_AMOUNT'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"