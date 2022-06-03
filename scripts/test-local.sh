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

echoInfo "INFO: Testing valopers query..."
./scripts/test-local/valopers-query.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Stopping local network..."
./scripts/test-local/network-stop.sh || ( systemctl2 stop sekai && exit 1 )

echoInfo "INFO: Success, all local tests passed, elapsed: $(prettyTime $(timerSpan))"