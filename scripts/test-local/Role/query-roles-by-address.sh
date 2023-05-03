#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="ROLES-BY-ADDRESS-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

INTERX_GATEWAY="127.0.0.1:11000"
TOTAL_ROLES_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/gov/roles_by_address/$VALIDATOR_ADDRESS" | jq '.roleIds | length' || echo "")
TOTAL_ROLES_CLI=$(showRoles validator | jq '.roleIds | length' || echo "")

[[ $TOTAL_ROLES_CLI != $TOTAL_ROLES_INTERX ]] && echoErr "ERROR: Expected number of roles to be '$TOTAL_ROLES_CLI', but got '$TOTAL_ROLES_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"