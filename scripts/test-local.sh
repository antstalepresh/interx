#!/usr/bin/env bash
set -e
set -x
. /etc/profile

echo "INFO: Started local tests in '$PWD'..."
timerStart

echoInfo "INFO: Stopping local network..."
./scripts/test-local/network-stop.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Launching local network..."
./scripts/test-local/network-start.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing account all balances query..."
./scripts/test-local/Account/query-all-balances.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing data reference query..."
./scripts/test-local/Data/query-data-reference.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing all identity records query..."
./scripts/test-local/Identity/all-identity-records.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing all identity verify requests query..."
./scripts/test-local/Identity/all-identity-verify-requests.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing identity record by id query..."
./scripts/test-local/Identity/identity-record-by-id.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing identity record verify request by approver query..."
./scripts/test-local/Identity/identity-record-verify-request-by-approver.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing identity record verify request by requester query..."
./scripts/test-local/Identity/identity-record-verify-request-by-requester.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing identity records by address query..."
./scripts/test-local/Identity/identity-records-by-address.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing kira status query..."
./scripts/test-local/Other/kira-status.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing total supply query..."
./scripts/test-local/Other/total-supply.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing valopers query..."
./scripts/test-local/valopers-query.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Stopping local network..."
./scripts/test-local/network-stop.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Success, all local tests passed, elapsed: $(prettyTime $(timerSpan))"