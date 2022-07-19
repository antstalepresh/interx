#!/usr/bin/env bash
set -e
set -x
. /etc/profile

echo "INFO: Ensuring essential dependencies are installed & up to date"
SYSCTRL_DESTINATION=/usr/local/bin/systemctl2
if [ ! -f $SYSCTRL_DESTINATION ] ; then
    safeWget /usr/local/bin/systemctl2 \
     https://raw.githubusercontent.com/gdraheim/docker-systemctl-replacement/9cbe1a00eb4bdac6ff05b96ca34ec9ed3d8fc06c/files/docker/systemctl.py \
     "e02e90c6de6cd68062dadcc6a20078c34b19582be0baf93ffa7d41f5ef0a1fdd" && \
    chmod +x $SYSCTRL_DESTINATION && \
    systemctl2 --version
fi

UTILS_VER=$(bashUtilsVersion 2> /dev/null || echo "")
UTILS_OLD_VER="false" && [[ $(versionToNumber "$UTILS_VER" || echo "0") -ge $(versionToNumber "v0.1.5" || echo "1") ]] || UTILS_OLD_VER="true" 

# Installing utils is essential to simplify the setup steps
if [ "$UTILS_OLD_VER" == "true" ] ; then
    echo "INFO: KIRA utils were NOT installed on the system, setting up..." && sleep 2
    TOOLS_VERSION="v0.1.5" && mkdir -p /usr/keys && FILE_NAME="bash-utils.sh" && \
     if [ -z "$KIRA_COSIGN_PUB" ] ; then KIRA_COSIGN_PUB=/usr/keys/kira-cosign.pub ; fi && \
     echo -e "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE/IrzBQYeMwvKa44/DF/HB7XDpnE+\nf+mU9F/Qbfq25bBWV2+NlYMJv3KvKHNtu3Jknt6yizZjUV4b8WGfKBzFYw==\n-----END PUBLIC KEY-----" > $KIRA_COSIGN_PUB && \
     wget "https://github.com/KiraCore/tools/releases/download/$TOOLS_VERSION/${FILE_NAME}" -O ./$FILE_NAME && \
     wget "https://github.com/KiraCore/tools/releases/download/$TOOLS_VERSION/${FILE_NAME}.sig" -O ./${FILE_NAME}.sig && \
     cosign verify-blob --key="$KIRA_COSIGN_PUB" --signature=./${FILE_NAME}.sig ./$FILE_NAME && \
     chmod -v 555 ./$FILE_NAME && ./$FILE_NAME bashUtilsSetup "/var/kiraglob" && . /etc/profile && \
     echoInfo "Installed bash-utils $(bashUtilsVersion)"

    globSet KIRA_COSIGN_PUB "$KIRA_COSIGN_PUB"
else
    echoInfo "INFO: KIRA utils are up to date, latest version $UTILS_VER" && sleep 2
fi

TEST_NAME="NETWORK-START" && timerStart $TEST_NAME

ARCHITECURE=$(getArch)
PLATFORM="$(getPlatform)"
INTERX_VERSION=$(./scripts/version.sh InterxVersion)
SEKAID_VERSION=$(./scripts/version.sh SekaiVersion)
DEFAULT_GRPC_PORT=9090
DEFAULT_RPC_PORT=26657
DEFAULT_INTERX_PORT=11000
PING_TARGET="127.0.0.1"
CFG_grpc="dns:///$PING_TARGET:$DEFAULT_GRPC_PORT"
CFG_rpc="http://$PING_TARGET:$DEFAULT_RPC_PORT"

set +x
echoInfo "          INFO: $TEST_NAME - Integration Test - START"
echoInfo "INTERX Version: $INTERX_VERSION"
echoInfo "SEKAID Version: $SEKAID_VERSION"
echoInfo "  Architecture: $ARCHITECURE"
echoInfo "      Platform: $PLATFORM"
set -x

BIN_DEST="/usr/local/bin/sekaid" && \
  safeWget ./sekaid.deb "https://github.com/KiraCore/sekai/releases/download/$SEKAID_VERSION/sekai-linux-${ARCHITECURE}.deb" \
  "$KIRA_COSIGN_PUB" && dpkg-deb -x ./sekaid.deb ./ && cp -fv "./bin/sekaid" $BIN_DEST && chmod -v 755 $BIN_DEST

BIN_DEST="/usr/local/bin/sekai-utils.sh" && \
  safeWget ./bin/sekai-utils.sh "https://github.com/KiraCore/sekai/releases/download/$SEKAID_VERSION/sekai-utils.sh" \
  "$KIRA_COSIGN_PUB" && chmod -v 755 ./bin/sekai-utils.sh && ./bin/sekai-utils.sh sekaiUtilsSetup && chmod -v 755 $BIN_DEST && . /etc/profile

FILE=/usr/local/bin/sekai-env.sh && \
safeWget $FILE "https://github.com/KiraCore/sekai/releases/download/$SEKAID_VERSION/sekai-env.sh" \
  "$KIRA_COSIGN_PUB" && chmod -v 755 $FILE && echo "source $FILE" >> /etc/profile && . /etc/profile

[ "$SEKAID_VERSION" != "$(sekaid version)" ] && echoErr "ERROR: Expected installed sekaid version to be '$SEKAID_VERSION', but got '$(sekaid version)'" && exit 1
[ "$INTERX_VERSION" != "$(interxd version)" ] && echoErr "ERROR: Expected installed interxd version to be '$INTERX_VERSION', but got '$(interxd version)'" && exit 1

echoInfo "INFO: Environment cleanup...."
NETWORK_NAME="localnet-1"
setGlobEnv SEKAID_HOME ~/.sekaid-$NETWORK_NAME
setGlobEnv INTERXD_HOME ~/.interxd-$NETWORK_NAME
setGlobEnv NETWORK_NAME $NETWORK_NAME
loadGlobEnvs

rm -rfv "$SEKAID_HOME" "$INTERXD_HOME"
mkdir -p "$SEKAID_HOME" "$INTERXD_HOME/cache"

echoInfo "INFO: Starting new network..."
sekaid init --overwrite --chain-id=$NETWORK_NAME "KIRA TEST LOCAL VALIDATOR NODE" --home=$SEKAID_HOME
addAccount validator
echo $(addAccount interx | jq .mnemonic | xargs) > $INTERXD_HOME/interx.mnemonic
echo $(addAccount faucet | jq .mnemonic | xargs) > $INTERXD_HOME/faucet.mnemonic
sekaid add-genesis-account $(showAddress validator) 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=test --home=$SEKAID_HOME
sekaid add-genesis-account $(showAddress faucet) 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=test --home=$SEKAID_HOME
sekaid gentx-claim validator --keyring-backend=test --moniker="GENESIS VALIDATOR" --home=$SEKAID_HOME

cat > /etc/systemd/system/sekai.service << EOL
[Unit]
Description=Local KIRA Test Network
After=network.target
[Service]
MemorySwapMax=0
Type=simple
User=root
WorkingDirectory=/root
ExecStart=/usr/local/bin/sekaid start --home=$SEKAID_HOME --trace
Restart=always
RestartSec=5
LimitNOFILE=4096
[Install]
WantedBy=default.target
EOL

systemctl2 enable sekai 
systemctl2 start sekai

echoInfo "INFO: Waiting for network to start..." && sleep 3

systemctl2 status sekai

echoInfo "INFO: Checking network status..."
NETWORK_STATUS_CHAIN_ID=$(showStatus | jq .NodeInfo.network | xargs)

if [ "$NETWORK_NAME" != "$NETWORK_STATUS_CHAIN_ID" ] ; then
    echoErr "ERROR: Incorrect chain ID from the status query, expected '$NETWORK_NAME', but got $NETWORK_STATUS_CHAIN_ID"
fi

echoInfo "INFO: Waiting for next block to be produced..."
timeout 60 sekai-utils awaitBlocks 2
BLOCK_HEIGHT=$(showBlockHeight)
timeout 60 sekai-utils awaitBlocks 2
NEXT_BLOCK_HEIGHT=$(showBlockHeight)

if [ $BLOCK_HEIGHT -ge $NEXT_BLOCK_HEIGHT ] ; then
    echoErr "ERROR: Failed to produce next block height, stuck at $BLOCK_HEIGHT"
fi

echoInfo "INFO: Printing sekai status..."
showStatus | jq
globSet validator_node_id "$(showStatus | jsonParse 'NodeInfo.id')"

echoInfo "INFO: Initalizing interxd..."

interxd init --cache_dir="$INTERXD_HOME/cache" --home="$INTERXD_HOME" --grpc="$CFG_grpc" --rpc="$CFG_rpc" --port="$INTERNAL_API_PORT" \
    --signing_mnemonic="$INTERXD_HOME/interx.mnemonic" \
    --faucet_mnemonic="$INTERXD_HOME/faucet.mnemonic" \
    --port="$DEFAULT_INTERX_PORT" \
    --node_type="validator" \
    --seed_node_id="" \
    --sentry_node_id="" \
    --validator_node_id="$(globGet validator_node_id)" \
    --addrbook="$(globFile KIRA_ADDRBOOK)" \
    --faucet_time_limit=30 \
    --faucet_amounts="100000ukex,20000000test,300000000000000000samolean,1lol" \
    --faucet_minimum_amounts="1000ukex,50000test,250000000000000samolean,1lol" \
    --fee_amounts="ukex 1000ukex,test 500ukex,samolean 250ukex,lol 100ukex"

cat > /etc/systemd/system/interx.service << EOL
[Unit]
Description=Local KIRA Test Network
After=network.target
[Service]
MemorySwapMax=0
Type=simple
User=root
WorkingDirectory=/root
ExecStart=$GOBIN/interxd start --home="$INTERXD_HOME"
Restart=always
RestartSec=5
LimitNOFILE=4096
[Install]
WantedBy=default.target
EOL

systemctl2 enable interx 
systemctl2 start interx

echoInfo "INFO: Waiting for interx to start..." && sleep 3

systemctl2 status interx

INTERX_GATEWAY="127.0.0.1:$DEFAULT_INTERX_PORT"

echoInfo "INFO: Waiting for next block to be produced..."
BLOCK_HEIGHT=$(curl --fail $INTERX_GATEWAY/api/status | jsonParse "interx_info.latest_block_height" || echo "0")
timeout 60 sekai-utils awaitBlocks 2
NEXT_BLOCK_HEIGHT=$(curl --fail $INTERX_GATEWAY/api/status | jsonParse "interx_info.latest_block_height" || echo "0")

if [ $BLOCK_HEIGHT -ge $NEXT_BLOCK_HEIGHT ] ; then
    echoErr "ERROR: INTERX failed to catch up with the latest sekai block height, stuck at $BLOCK_HEIGHT"
fi

echoInfo "INFO: Printing interx status..."
curl --fail $INTERX_GATEWAY/api/status | jq

set +x
echoInfo "INFO: SEKAID $(sekaid version) is running"
echoInfo "INFO: INTERXD $(interxd version) is running"
echoInfo "INFO: NETWORK-START - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"