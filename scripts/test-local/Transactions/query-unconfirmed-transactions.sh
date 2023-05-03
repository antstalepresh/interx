#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="TX-UNCONFIRMED-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)
addAccount testuser6
TESTUSER_ADDRESS=$(showAddress testuser6)

RESULT=$(sekaid tx bank send validator $TESTUSER_ADDRESS 5ukex --keyring-backend=test --chain-id=$NETWORK_NAME --fees 100ukex --output=json --yes --home=$SEKAID_HOME 2> /dev/null || exit 1)

INTERX_GATEWAY="127.0.0.1:11000"
RESULT_FROM_INTERX=$(curl --fail $INTERX_GATEWAY/api/unconfirmed_txs | jq '.txs[0]' || exit 1)

RESULT_FROM=$(echo $RESULT_FROM_INTERX | jq '.msgs[0].data.from_address' | tr -d '"')
RESULT_TO=$(echo $RESULT_FROM_INTERX | jq '.msgs[0].data.to_address' | tr -d '"')
RESULT_DENOM=$(echo $RESULT_FROM_INTERX | jq '.msgs[0].data.amount[0].denom' | tr -d '"')
RESULT_AMOUNT=$(echo $RESULT_FROM_INTERX | jq '.msgs[0].data.amount[0].amount' | tr -d '"')
[[ $VALIDATOR_ADDRESS !=  $RESULT_FROM ]] && echoErr "ERROR: Expected tx sender address to be '$VALIDATOR_ADDRESS', but got '$RESULT_FROM'" && exit 1
[[ $TESTUSER_ADDRESS !=  $RESULT_TO ]] && echoErr "ERROR: Expected tx receiver address to be '$TESTUSER_ADDRESS', but got '$RESULT_TO'" && exit 1
[[ "ukex" !=  $RESULT_DENOM ]] && echoErr "ERROR: Expected tx denom to be ukex, but got '$RESULT_DENOM'" && exit 1
[[ "5" !=  $RESULT_AMOUNT ]] && echoErr "ERROR: Expected tx amount to be 5, but got '$RESULT_AMOUNT'" && exit 1


echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"