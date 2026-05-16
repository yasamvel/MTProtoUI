#!/usr/bin/env bash

set -e

clear

echo "MTProtoUI Installer"
echo

echo "1) English"
echo "2) Русский"

read -p "Select language: " LANG

echo
echo "Installing Docker..."

curl -fsSL https://get.docker.com | sh

echo
echo "Installing Git..."

apt-get update -y
apt-get install -y git curl

echo
echo "Cloning repository..."

cd /opt || exit

rm -rf MTProtoUI

git clone https://github.com/yasamvel/MTProtoUI.git

cd MTProtoUI || exit

docker compose up -d --build

echo
echo "Installation completed"
echo

SERVER_IP=$(curl -s ifconfig.me)

echo "Panel:"
echo "http://$SERVER_IP:8080"
