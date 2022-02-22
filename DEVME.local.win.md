# Local Development on Windows

## Re/Install WSL on Windows 10 (PowerShell)

```
# Uninstall Ubuntu
wsl --terminate Ubuntu-20.04 && \
 wsl --unregister Ubuntu-20.04

# Install Ubuntu
wsl --install -d Ubuntu-20.04 && \
 wsl --setdefault Ubuntu-20.04 && \
 wsl --set-version Ubuntu-20.04 2
```

## Essential Dependencies

```
# Open Ubuntu 20.04 WSL 2.0 console

# Install Essential Dependecies
apt-get install -y curl && curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - && apt-get update -y && \
 apt-get install -y --allow-unauthenticated --allow-downgrades --allow-remove-essential --allow-change-held-packages \
 software-properties-common wget git nginx apt-transport-https file build-essential net-tools hashdeep \
 protobuf-compiler golang-goprotobuf-dev golang-grpc-gateway golang-github-grpc-ecosystem-grpc-gateway-dev lsb-release \
 clang cmake gcc g++ pkg-config libudev-dev libusb-1.0-0-dev iputils-ping nano jq python python3 python3-pip gnupg \
 bash libglu1-mesa lsof bc dnsutils psmisc netcat  make nodejs tar unzip xz-utils yarn zip p7zip-full ca-certificates \
 containerd docker.io dos2unix

# install systemd alternative
wget https://raw.githubusercontent.com/gdraheim/docker-systemctl-replacement/master/files/docker/systemctl.py -O /usr/local/bin/systemctl2 && \
 chmod +x /usr/local/bin/systemctl2 && \
 systemctl2 --version

# install kira bash helper utils
BRANCH="v0.0.1" && cd /tmp && rm -fv ./i.sh && \
wget https://raw.githubusercontent.com/KiraCore/tools/$BRANCH/bash-utils/install.sh -O ./i.sh && \
 chmod 555 -v ./i.sh && ./i.sh "$BRANCH" "/var/kiraglob" && . /etc/profile && rm -fv ./i.sh
 
# uninstall golang if needed
( go clean -modcache -cache -n || echo "Failed to cleanup go cache" ) && \
  ( rm -rfv "$GOROOT" || echo "Failed to cleanup go root" ) && \
  ( rm -rfv "$GOBIN" || echo "Failed to cleanup go bin" ) && \
  ( rm -rfv "$GOPATH" || echo "Failed to cleanup go path" ) && \
  ( rm -rfv "$GOCACHE" || echo "Failed to cleanup go cache" )

# mount C drive or other disk where repo is stored
setGlobLine "mount -t drvfs C:" "mount -t drvfs C: /mnt/c"

# set env variable to your local repos (will vary depending on the user)
setGlobEnv SEKAI_REPO "/mnt/c/Users/asmodat/Desktop/KIRA/KIRA-CORE/GITHUB/sekai" && \
 setGlobEnv INTERX_REPO "/mnt/c/Users/asmodat/Desktop/KIRA/KIRA-CORE/GITHUB/interx" && \
 loadGlobEnvs

# set home directory of your repos
setGlobEnv SEKAID_HOME "/root/.sekaid" && \
 setGlobEnv INTERXD_HOME "/root/.interxd" && \
 loadGlobEnvs

# Ensure you have Docker Desktop installed: https://code.visualstudio.com/blogs/2020/03/02/docker-in-wsl2 & reboot your entire host machine
```

# Clean Clone
```
cd $HOME && rm -fvr ./interx && INTERX_BRANCH="master" && \
 git clone https://github.com/KiraCore/interx.git -b $INTERX_BRANCH && \
 cd ./interx
```

## Installation

```
cd $INTERX_REPO

go get github.com/KiraCore/sekai@master

make install
```

## Startup

```

    CFG_grpc="dns:///127.0.0.1:9090" && \
     CFG_rpc="http://127.0.0.1:26657" && \
     rm -rfv $INTERXD_HOME && mkdir -p $INTERXD_HOME/cache && \
     rm -rfv /interx/cache &&
     interxd init --cache_dir="$INTERXD_HOME/cache" --config="$INTERXD_HOME/config.json" --grpc="$CFG_grpc" --rpc="$CFG_rpc" --port="$INTERNAL_API_PORT" \
      --signing_mnemonic="$COMMON_DIR/signing.mnemonic" \
      --seed_node_id="$seed_node_id" \
      --sentry_node_id="$sentry_node_id" \
      --validator_node_id="$validator_node_id" \
      --addrbook="$KIRA_ADDRBOOK_FILE" \
      --faucet_time_limit=30 \
      --faucet_amounts="100000ukex,20000000test,300000000000000000samolean,1lol" \
      --faucet_minimum_amounts="1000ukex,50000test,250000000000000samolean,1lol" \
      --fee_amounts="ukex 1000ukex,test 500ukex,samolean 250ukex, lol 100ukex" \
      --version="$KIRA_SETUP_VER"

    seed_node_id=$(globGet seed_node_id)
    sentry_node_id=$(globGet sentry_node_id)
    validator_node_id=$(globGet validator_node_id)

    interxd init --cache_dir="$CACHE_DIR" --config="$CONFIG_PATH" --grpc="$CFG_grpc" --rpc="$CFG_rpc" --port="$INTERNAL_API_PORT" \
      --signing_mnemonic="$COMMON_DIR/signing.mnemonic" \
      --seed_node_id="$seed_node_id" \
      --sentry_node_id="$sentry_node_id" \
      --validator_node_id="$validator_node_id" \
      --addrbook="$KIRA_ADDRBOOK_FILE" \
      --faucet_time_limit=30 \
      --faucet_amounts="100000ukex,20000000test,300000000000000000samolean,1lol" \
      --faucet_minimum_amounts="1000ukex,50000test,250000000000000samolean,1lol" \
      --fee_amounts="ukex 1000ukex,test 500ukex,samolean 250ukex, lol 100ukex" \
      --version="$KIRA_SETUP_VER"

```