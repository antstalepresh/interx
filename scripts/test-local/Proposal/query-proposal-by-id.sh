#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="PROPOSAL-QUERY-BY-ID" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

PROPOSAL_RESULT=$(setPoorNetworkMessages validator asdf)

INTERX_GATEWAY="127.0.0.1:11000"
PROPOSAL_ID_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/proposals?all=true&reverse=true" | jq '.proposals[0].proposal_id' | tr -d '"' || exit 1)
RESULT_INTERX=$(curl --fail $INTERX_GATEWAY/api/kira/gov/proposals/$PROPOSAL_ID_INTERX | jq '.proposal.content.messages[0]' || exit 1)
RESULT_CLI=$(showProposal $PROPOSAL_ID_INTERX | jq '.content.messages[0]')

[ $RESULT_INTERX != $RESULT_CLI ] && echoErr "ERROR: Expected proposal description to be '$RESULT_CLI', but got '$RESULT_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"