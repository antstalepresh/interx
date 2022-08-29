#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="ALL-IDENTITY-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

INTERX_GATEWAY="127.0.0.1:11000"
IDENTITY_TOTAL_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/all_identity_records" | jq '.records | length' || echo "0")
IDENTITY_TOTAL_CLI=$(sekaid query customgov all-identity-records --output=json --home=$SEKAID_HOME | jq '.records | length' || echo "0")

[ $IDENTITY_TOTAL_INTERX != $IDENTITY_TOTAL_CLI ] && echoErr "ERROR: Expected number of records to be '$IDENTITY_TOTAL_CLI', but got '$IDENTITY_TOTAL_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"