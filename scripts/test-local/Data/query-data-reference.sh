#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="DATA-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

INTERX_GATEWAY="127.0.0.1:11000"
DATA_NUMS_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/data_keys?limit=2&offset=0&count_total=true" | jq '.keys | length' || exit 1)
DATA_NUMS_CLI=$(showDataRegistryKeys | jq '. | length')

[ $DATA_NUMS_INTERX != $DATA_NUMS_CLI ] && echoErr "ERROR: Expected number of keys to be '$DATA_NUMS_CLI', but got '$DATA_NUMS_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"