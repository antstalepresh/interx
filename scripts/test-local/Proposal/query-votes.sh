#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="VOTES-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

PROPOSAL_RESULT=$(setPoorNetworkMessages validator asdf)

INTERX_GATEWAY="127.0.0.1:11000"
PROPOSAL_ID_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/proposals?all=true" | jq '.proposals | length' | tr -d '"' || exit 1)

voteYes $PROPOSAL_ID_INTERX validator

NUM_VOTERS_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/votes/$PROPOSAL_ID_INTERX" | jq '. | length' | tr -d '"' || exit 1)
NUM_VOTERS_CLI=$(sekaid query customgov votes $PROPOSAL_ID_INTERX --output=json | jq '.votes | length' | tr -d '"')

[[ $NUM_VOTERS_INTERX != $NUM_VOTERS_CLI ]] && echoErr "ERROR: Expected number of votes to be '$NUM_VOTERS_CLI', but got '$NUM_VOTERS_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"