#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="STAKING-POOL-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VAL_ADDRESS=$(showAddress validator)
VALOPER_ADDRESS=$(showValidator validator | jq '.val_key' | tr -d '"')

INTERX_GATEWAY="127.0.0.1:11000"
STAKING_POOL_INTERX=$(curl --fail "$INTERX_GATEWAY/api/kira/staking-pool?validatorAddress=$VAL_ADDRESS" || exit 1)

TOKENS_NUMS_INTERX=$(echo $STAKING_POOL_INTERX | jq '.tokens | length')
DELEGATORS_NUMS_INTERX=$(echo $STAKING_POOL_INTERX | jq '.total_delegators')

TOKENS_NUMS_CLI=$(showTokenRates | jq '.data | length')
STAKING_POOL_DELEGATORS_NUM_CLI=$(sekaid query multistaking staking-pool-delegators $VALOPER_ADDRESS --output=json  | jq '.delegators | length')

[[ "$TOKENS_NUMS_INTERX" != "$TOKENS_NUMS_CLI" ]] && echoErr "ERROR: Expected number of tokens to be '$TOKENS_NUMS_CLI', but got '$TOKENS_NUMS_INTERX'" && exit 1
[[ "$DELEGATORS_NUMS_INTERX" != "$STAKING_POOL_DELEGATORS_NUM_CLI" ]] && echoErr "ERROR: Expected number of tokens to be '$TOKENS_NUMS_CLI', but got '$TOKENS_NUMS_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"