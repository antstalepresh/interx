#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="IDENTITY-RECORD-BY-ID-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

INTERX_GATEWAY="127.0.0.1:11000"
IDENTITY_ADDRESS_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/identity_record/1" | jq '.record.address' || echo "")
IDENTITY_ADDRESS_CLI=$(sekaid query customgov identity-record 1 --output=json --home=$SEKAID_HOME | jq '.record.address' || echo "")

[ $IDENTITY_ADDRESS_INTERX != $IDENTITY_ADDRESS_CLI ] && echoErr "ERROR: Expected address of identity to be '$IDENTITY_ADDRESS_CLI', but got '$IDENTITY_ADDRESS_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"