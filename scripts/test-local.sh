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

echoInfo "INFO: Testing permissions by address query..."
./scripts/test-local/Permission/query-permissions-by-address.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing network properties query..."
./scripts/test-local/Proposal/query-network-properties.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing proposal by id query..."
./scripts/test-local/Proposal/query-proposal-by-id.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing voters query..."
./scripts/test-local/Proposal/query-voters.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing votes query..."
./scripts/test-local/Proposal/query-votes.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing all roles query..."
./scripts/test-local/Role/query-all-roles.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing roles by address query..."
./scripts/test-local/Role/query-roles-by-address.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing pool by name query..."
./scripts/test-local/Spending/pool-by-name.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing pool name by account query..."
./scripts/test-local/Spending/pool-name-by-account.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing spending pools query..."
./scripts/test-local/Spending/spending-pools.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing tokens aliases query..."
./scripts/test-local/Tokens/query-tokens-aliases.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing tokens rates query..."
./scripts/test-local/Tokens/query-tokens-rates.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing query block by height or hash query..."
./scripts/test-local/Transactions/query-block-by-height-or-hash.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing block transactions query..."
./scripts/test-local/Transactions/query-block-transactions.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing query blocks query..."
./scripts/test-local/Transactions/query-blocks.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing query transaction result..."
./scripts/test-local/Transactions/query-transaction-result.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing query transactions result..."
./scripts/test-local/Transactions/query-transactions.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing unconfirmed transactions query..."
./scripts/test-local/Transactions/query-unconfirmed-transactions.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing transaction hash query..."
./scripts/test-local/Transactions/transaction-hash.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing ubi records by name query..."
./scripts/test-local/Ubi/query-ubi-records-by-name.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing ubi records query..."
./scripts/test-local/Ubi/query-ubi-records.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing current upgrade plan query..."
./scripts/test-local/Upgrade/current-upgrade-plan.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing next upgrade plan query..."
./scripts/test-local/Upgrade/next-upgrade-plan.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing validator status query..."
./scripts/test-local/Validators/query-validator-status.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Testing validators query..."
./scripts/test-local/Validators/query-validators.sh || ( systemctl2 stop sekai && exit 1 )

# echoInfo "INFO: Testing evm account query..."
# ./scripts/test-local/Evm/query-accounts.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Stopping local network..."
./scripts/test-local/network-stop.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Success, all local tests passed, elapsed: $(prettyTime $(timerSpan))"