#!/usr/bin/env bash

set -e

clear

echo "MTProtoUI Installer"
echo

echo "1) English"
echo "2) Русский"

read -p "Select language: " LANG

echo
echo "Panel port setup"

read -p "Enter panel port [default: 8080]: " PANEL_PORT

if [ -z "$PANEL_PORT" ]; then
    PANEL_PORT=8080
fi

echo

if ! command -v docker &> /dev/null
then
    echo "Installing Docker..."
    curl -fsSL https://get.docker.com | sh
else
    echo "Docker already installed"
fi

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

sed -i "s/8080:8080/${PANEL_PORT}:8080/g" docker-compose.yml

docker compose up -d --build

echo
echo "Installation completed"
echo

SERVER_IP=$(curl -s ifconfig.me)

echo "Panel:"
echo "http://$SERVER_IP:$PANEL_PORT"
