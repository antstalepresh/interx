#!/bin/bash
set -e
echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list
sudo apt update
sudo apt install nfpm -y

echo "Installing build-essential, golang"
sudo apt install -y build-essential
wget https://go.dev/dl/go1.17.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
rm go1.17.7.linux-amd64.tar.gz
../make publish
