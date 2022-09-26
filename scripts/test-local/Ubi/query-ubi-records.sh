#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="UBI-RECORDS-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

INTERX_GATEWAY="127.0.0.1:11000"
TOTAL_RECORDS_CLI=$(sekaid query ubi ubi-records --output=json | jq '.records | length' || echo "0")
TOTAL_RECORDS_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/ubi-records" | jq '.records | length' || echo "0")

[ $TOTAL_RECORDS_CLI != $TOTAL_RECORDS_INTERX ] && echoErr "ERROR: Expected number of ubi records to be '$TOTAL_RECORDS_CLI', but got '$TOTAL_RECORDS_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"