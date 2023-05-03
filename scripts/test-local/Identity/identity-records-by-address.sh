#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="IDENTITY-RECORD-BY-ADDRESS-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

INTERX_GATEWAY="127.0.0.1:11000"
IDENTITY_KEY_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/identity_records/$VALIDATOR_ADDRESS" | jq '.records[0].key' || echo "")
IDENTITY_KEY_CLI=$(sekaid query customgov identity-records-by-addr $VALIDATOR_ADDRESS --output=json | jq '.records[0].key' || echo "")

[[ $IDENTITY_KEY_INTERX != $IDENTITY_KEY_CLI ]] && echoErr "ERROR: Expected key of identity record to be '$IDENTITY_KEY_CLI', but got '$IDENTITY_KEY_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"