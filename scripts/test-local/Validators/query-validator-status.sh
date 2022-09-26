#!/usr/bin/env bash
# To run test locally: make network-start && ./scripts/test-local/account-balances.sh
set -e
set -x
. /etc/profile

TEST_NAME="VALIDATOR-STATUS-QUERY" && timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

VALIDATOR_ADDRESS=$(showAddress validator)

RESULT=$(sekaid query customstaking validators --moniker="" --output=json | jq '.validators' 2> /dev/null || exit 1)
RESULT_=$(echo $RESULT | jq '.[]')

TOTAL_VAL=$(echo $RESULT | jq '. | length' | tr -d '"')
ACTIVE_VAL=0
PAUSED_VAL=0
INACTIVE_VAL=0
JAILED_VAL=0
WAITING_VAL=0

for i in "${RESULT_[@]}"
do
  VAL_STATUS=$(echo $i | jq '.status' | tr -d '"')
  [ $VAL_STATUS == "ACTIVE" ] && ((++ACTIVE_VAL))
  [ $VAL_STATUS == "PAUSED" ] && ((++PAUSED_VAL))
  [ $VAL_STATUS == "INACTIVE" ] && ((++INACTIVE_VAL))
  [ $VAL_STATUS == "JAILED" ] && ((++JAILED_VAL))
  [ $VAL_STATUS == "WAITING" ] && ((++WAITING_VAL))
done

INTERX_GATEWAY="127.0.0.1:11000"
RESULT_FROM_INTERX=$(curl --fail $INTERX_GATEWAY/api/valopers?status_only=true || exit 1)
TOTAL_VAL_INTERX=$(echo $RESULT_FROM_INTERX | jq '.total_validators')
ACTIVE_VAL_INTERX=$(echo $RESULT_FROM_INTERX | jq '.active_validators')
PAUSED_VAL_INTERX=$(echo $RESULT_FROM_INTERX | jq '.paused_validators')
INACTIVE_VAL_INTERX=$(echo $RESULT_FROM_INTERX | jq '.inactive_validators')
JAILED_VAL_INTERX=$(echo $RESULT_FROM_INTERX | jq '.jailed_validators')
WAITING_VAL_INTERX=$(echo $RESULT_FROM_INTERX | jq '.waiting_validators')

[ $TOTAL_VAL != $TOTAL_VAL_INTERX ] && echoErr "ERROR: Expected total validator amount to be '$TOTAL_VAL', but got '$TOTAL_VAL_INTERX'" && exit 1
[ $ACTIVE_VAL != $ACTIVE_VAL_INTERX ] && echoErr "ERROR: Expected active validator amount to be '$ACTIVE_VAL', but got '$ACTIVE_VAL_INTERX'" && exit 1
[ $PAUSED_VAL != $PAUSED_VAL_INTERX ] && echoErr "ERROR: Expected paused validator amount to be '$PAUSED_VAL', but got '$PAUSED_VAL_INTERX'" && exit 1
[ $INACTIVE_VAL != $INACTIVE_VAL_INTERX ] && echoErr "ERROR: Expected inactive validator amount to be '$INACTIVE_VAL', but got '$INACTIVE_VAL_INTERX'" && exit 1
[ $JAILED_VAL != $JAILED_VAL_INTERX ] && echoErr "ERROR: Expected jailed validator amount to be '$JAILED_VAL', but got '$JAILED_VAL_INTERX'" && exit 1
[ $WAITING_VAL != $WAITING_VAL_INTERX ] && echoErr "ERROR: Expected waiting validator amount to be '$WAITING_VAL', but got '$WAITING_VAL_INTERX'" && exit 1

echoInfo "INFO: $TEST_NAME - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"