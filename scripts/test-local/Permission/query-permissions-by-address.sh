#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="PERMISSION-BY-ADDRESS-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

INTERX_GATEWAY="127.0.0.1:11000"
TOTAL_BLACKLIST_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/permissions_by_address/$VALIDATOR_ADDRESS" | jq '.blacklist | length' || echo "")
TOTAL_BLACKLIST_CLI=$(showPermissions validator | jq '.blacklist | length' || echo "")

[[ $TOTAL_BLACKLIST_CLI != $TOTAL_BLACKLIST_INTERX ]] && echoErr "ERROR: Expected number of blacklist to be '$TOTAL_BLACKLIST_CLI', but got '$TOTAL_BLACKLIST_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"