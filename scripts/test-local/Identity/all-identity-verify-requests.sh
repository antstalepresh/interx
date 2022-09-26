#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="ALL-IDENTITY-VERIFY-REQUESTS-ID-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)
verifyIdentityRecord $VALIDATOR_ADDRESS validator 1 1000ukex

INTERX_GATEWAY="127.0.0.1:11000"
TOTAL_RECORDS_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/all_identity_verify_requests" | jq '.verifyRecords | length' || echo "")
TOTAL_RECORDS_CLI=$(sekaid query customgov all-identity-record-verify-requests --output=json | jq '.verify_records | length')

[ $TOTAL_RECORDS_CLI != $TOTAL_RECORDS_INTERX ] && echoErr "ERROR: Expected number of records to be '$TOTAL_RECORDS_CLI', but got '$TOTAL_RECORDS_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"