#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="PLAN-NEXT-UPGRADE-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

INTERX_GATEWAY="127.0.0.1:11000"
PLAN_DATA_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/upgrade/next_plan" | jq '.plan' || exit 1)
PLAN_NAME_INTERX=$(echo $PLAN_DATA_CLI | jq '. | name' || echo "null")
PLAN_PROPOSAL_INTERX=$(echo $PLAN_DATA_CLI | jq '. | proposalID' || echo "null")

PLAN_DATA_CLI=$(showNextPlan | jq '.plan')
PLAN_NAME_CLI=$(echo $PLAN_DATA_CLI | jq '. | name' || echo "null")
PLAN_PROPOSAL_CLI=$(echo $PLAN_DATA_CLI | jq '. | proposalID' || echo "null")

[[ $PLAN_NAME_CLI != $PLAN_NAME_INTERX ]] && echoErr "ERROR: Expected name of plan to be '$PLAN_NAME_CLI', but got '$PLAN_NAME_INTERX'" && exit 1
[[ $PLAN_PROPOSAL_CLI != $PLAN_PROPOSAL_INTERX ]] && echoErr "ERROR: Expected proposal ID of plan to be '$PLAN_PROPOSAL_CLI', but got '$PLAN_PROPOSAL_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"