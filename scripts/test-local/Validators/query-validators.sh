#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="VALIDATOR-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

RESULT_NUM=$(sekaid query customstaking validators --moniker="" --output=json | jq '.validators | length' 2> /dev/null || exit 1)
SIGNING_INFO_SEKAI=$(sekaid query customslashing signing-infos --output=json | jq '.info[0]' 2> /dev/null || exit 1)
START_HEIGHT_SEKAI=$(echo $SIGNING_INFO_SEKAI | jq '.start_height' | tr -d '"')
MISCHANCE_CONFIDENCE_SEKAI=$(echo $SIGNING_INFO_SEKAI | jq '.mischance_confidence' | tr -d '"')
LAST_PRESENT_BLOCK_SEKAI=$(echo $SIGNING_INFO_SEKAI | jq '.last_present_block' | tr -d '"')
MISSED_BLOCKS_COUNTER_SEKAI=$(echo $SIGNING_INFO_SEKAI | jq '.missed_blocks_counter' | tr -d '"')
PRODUCED_BLOCKS_COUNTER_SEKAI=$(echo $SIGNING_INFO_SEKAI | jq '.produced_blocks_counter' | tr -d '"')

INTERX_GATEWAY="127.0.0.1:11000"
RESULT_FROM_INTERX=$(curl --fail $INTERX_GATEWAY/api/valopers | jq '.validators' || exit 1)
RESULT_FROM_INTERX_NUM=${#RESULT_FROM_INTERX[@]}
SIGNING_INFO_INTERX=$(echo $RESULT_FROM_INTERX | jq '.[0]' 2> /dev/null || exit 1)
START_HEIGHT_INTERX=$(echo $SIGNING_INFO_INTERX | jq '.start_height' | tr -d '"')
MISCHANCE_CONFIDENCE_INTERX=$(echo $SIGNING_INFO_INTERX | jq '.mischance_confidence' | tr -d '"')
LAST_PRESENT_BLOCK_INTERX=$(echo $SIGNING_INFO_INTERX | jq '.last_present_block' | tr -d '"')
MISSED_BLOCKS_COUNTER_INTERX=$(echo $SIGNING_INFO_INTERX | jq '.missed_blocks_counter' | tr -d '"')
PRODUCED_BLOCKS_COUNTER_INTERX=$(echo $SIGNING_INFO_INTERX | jq '.produced_blocks_counter' | tr -d '"')

[ $RESULT_NUM != $RESULT_FROM_INTERX_NUM ] && echoErr "ERROR: Expected validator amount to be '$RESULT_NUM', but got '$RESULT_FROM_INTERX_NUM'" && exit 1
if [ $RESULT_NUM -ge 1 ] ; then
    [ $START_HEIGHT_SEKAI != $START_HEIGHT_INTERX ] && echoErr "ERROR: Expected start height to be '$START_HEIGHT_SEKAI', but got '$START_HEIGHT_INTERX'" && exit 1
    [ $MISCHANCE_CONFIDENCE_SEKAI != $MISCHANCE_CONFIDENCE_INTERX ] && echoErr "ERROR: Expected mischance confidence to be '$MISCHANCE_CONFIDENCE_SEKAI', but got '$MISCHANCE_CONFIDENCE_INTERX'" && exit 1
    [ $LAST_PRESENT_BLOCK_SEKAI -ge 1 ] && [ $LAST_PRESENT_BLOCK_INTERX == 0 ] && echoErr "ERROR: Expected last present block to be '$LAST_PRESENT_BLOCK_SEKAI', but got '$LAST_PRESENT_BLOCK_INTERX'" && exit 1
    [ $MISSED_BLOCKS_COUNTER_SEKAI != $MISSED_BLOCKS_COUNTER_INTERX ] && echoErr "ERROR: Expected missed blocks counter to be '$MISSED_BLOCKS_COUNTER_SEKAI', but got '$MISSED_BLOCKS_COUNTER_INTERX'" && exit 1
    [ $PRODUCED_BLOCKS_COUNTER_SEKAI -ge 1] && [ $PRODUCED_BLOCKS_COUNTER_INTERX == 0 ] && echoErr "ERROR: Expected produced blocks counter to be '$PRODUCED_BLOCKS_COUNTER_SEKAI', but got '$PRODUCED_BLOCKS_COUNTER_INTERX'" && exit 1
fi

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"