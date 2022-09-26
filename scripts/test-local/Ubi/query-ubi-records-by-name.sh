#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="UBI-RECORD-BY-NAME-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

RECORD_NAME_CLI=$(sekaid query ubi ubi-records --output=json | jq '.records[0].name' | tr -d '"' || echo "null")

INTERX_GATEWAY="127.0.0.1:11000"
RECORD_AMOUNT_CLI=$(sekaid query ubi ubi-record-by-name $RECORD_NAME_CLI --output=json | jq '.record.amount' || echo "0")
RECORD_AMOUNT_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/ubi-records?name=$RECORD_NAME_CLI" | jq '.record.amount' || echo "0")

[ $RECORD_AMOUNT_CLI != $RECORD_AMOUNT_INTERX ] && echoErr "ERROR: Expected number of ubi records to be '$RECORD_AMOUNT_CLI', but got '$RECORD_AMOUNT_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"