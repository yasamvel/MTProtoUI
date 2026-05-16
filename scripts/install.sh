#!/usr/bin/env bash

clear

echo "MTProto UI Next Installer"
echo

read -p "Use domain or IP? [domain/ip]: " MODE

if [ "$MODE" = "domain" ]; then
    read -p "Enter domain: " DOMAIN
    echo "Domain selected: $DOMAIN"
else
    echo "IP mode selected"
fi

echo
echo "Select proxy type:"
echo "1) classic"
echo "2) dd"
echo "3) ee"

read -p "Choice: " TYPE

echo
echo "Installing Docker..."

echo
echo "Starting stack..."

docker compose up -d

echo
echo "Installation completed"
