#!/usr/bin/env bash

set -e

echo "Removing MTProtoUI..."

cd /opt/MTProtoUI || exit

docker compose down || true

cd /opt || exit

rm -rf /opt/MTProtoUI

docker rm -f mtproto-ui 2>/dev/null || true

docker image prune -af || true

echo
echo "MTProtoUI removed successfully"
