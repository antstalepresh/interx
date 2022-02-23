# Local Development on Windows

## Re/Install WSL on Windows 10 (PowerShell)

```
wsl --terminate Ubuntu-20.04
wsl --unregister Ubuntu-20.04

wsl --install -d Ubuntu-20.04
wsl --setdefault Ubuntu-20.04
wsl --set-version Ubuntu-20.04 2

```

## Essential Dependencies

```
# Open Ubuntu 20.04 WSL 2.0 console

# Install Essential Dependecies
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - && \
 apt-get update -y && \
 apt-get install -y --allow-unauthenticated --allow-downgrades --allow-remove-essential --allow-change-held-packages \
    software-properties-common curl wget git nginx apt-transport-https file build-essential net-tools hashdeep \
    protobuf-compiler golang-goprotobuf-dev golang-grpc-gateway golang-github-grpc-ecosystem-grpc-gateway-dev lsb-release \
    clang cmake gcc g++ pkg-config libudev-dev libusb-1.0-0-dev iputils-ping nano jq python python3 python3-pip gnupg \
    bash libglu1-mesa lsof bc dnsutils psmisc netcat  make nodejs tar unzip xz-utils yarn zip p7zip-full ca-certificates \
	bridge-utils containerd docker.io 

# install systemd alternative
wget https://raw.githubusercontent.com/gdraheim/docker-systemctl-replacement/master/files/docker/systemctl.py -O /usr/local/bin/systemctl2 && \
 chmod +x /usr/local/bin/systemctl2 && \
 systemctl2 --version
 
# uninstall golang if needed
( go clean -modcache -cache -n || echo "Failed to cleanup go cache" ) && \
  ( rm -rfv "$GOROOT" || echo "Failed to cleanup go root" ) && \
  ( rm -rfv "$GOBIN" || echo "Failed to cleanup go bin" ) && \
  ( rm -rfv "$GOPATH" || echo "Failed to cleanup go path" ) && \
  ( rm -rfv "$GOCACHE" || echo "Failed to cleanup go cache" )

# install golang
GO_VERSION="1.17.2" && ARCH=$(([[ "$(uname -m)" == *"arm"* ]] || [[ "$(uname -m)" == *"aarch"* ]]) && echo "arm64" || echo "amd64") && \
GO_TAR=go${GO_VERSION}.linux-${ARCH}.tar.gz && rm -rfv /usr/local/go && cd /tmp && rm -fv ./$GO_TAR && \
 wget https://dl.google.com/go/${GO_TAR} && \
 tar -C /usr/local -xvf $GO_TAR && touch ~/.bash_aliases && \
 if ! grep -q GOPATH ~/.bash_aliases ; then
  echo "export GOROOT=/usr/local/go" >> ~/.bash_aliases
  echo "export GOBIN=/usr/local/go/bin" >> ~/.bash_aliases
  echo "export GOPATH=/home/go" >> ~/.bash_aliases
  echo "export GOCACHE=/home/go/cache" >> ~/.bash_aliases
  echo "export PATH=\$PATH:\$GOROOT:\$GOBIN:\$GOPATH" >> ~/.bash_aliases
  . ~/.bashrc
  go version
else
  go version
fi

# mount C drive or other disk where repo is stored
mkdir -p /mnt/c && \
 echo "mount -t drvfs C: /mnt/c" >> ~/.bash_aliases

# set env variable to your local repos
echo "SEKAI_REPO=\"/mnt/c/Users/asmodat/Desktop/KIRA/KIRA-CORE/GITHUB/sekai\"" >> ~/.bash_aliases && \
 echo "INTERX_REPO=\"/mnt/c/Users/asmodat/Desktop/KIRA/KIRA-CORE/GITHUB/interx\"" >> ~/.bash_aliases

# set home directory of your repos
echo "SEKAID_HOME=/root/.sekaid" >> ~/.bash_aliases

# Ensure you have Docker Desktop installed: https://code.visualstudio.com/blogs/2020/03/02/docker-in-wsl2 & reboot your entire host machine

. ~/.bashrc
cd $INTERX_REPO
```

## Installation

```
go get github.com/KiraCore/sekai@master

make install
```