#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="PROPOSAL-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

PROPOSAL_RESULT=$(setPoorNetworkMessages validator asdf)

INTERX_GATEWAY="127.0.0.1:11000"
PROPOSAL_NUMS_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/proposals?all=true&reverse=true" | jq '.proposals | length' || exit 1)
PROPOSAL_NUMS_CLI=$(showProposals | jq '.proposals | length')

[[ $PROPOSAL_NUMS_INTERX != $PROPOSAL_NUMS_CLI ]] && echoErr "ERROR: Expected number of proposals to be '$PROPOSAL_NUMS_CLI', but got '$PROPOSAL_NUMS_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"