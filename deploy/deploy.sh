#!/bin/bash
set -euo pipefail

IMAGE="${1:?Usage: deploy.sh IMAGE_URI}"
DATA_DIR="/mnt/disks/studyhub-data"
ENV_FILE="/opt/studyhub/studyhub.env"
CONTAINER="studyhub"

sudo mkdir -p "$DATA_DIR"
sudo chown -R 10001:10001 "$DATA_DIR"

sudo docker pull "$IMAGE"
sudo docker rm -f "$CONTAINER" 2>/dev/null || true

sudo docker run -d \
  --name "$CONTAINER" \
  --restart unless-stopped \
  --env-file "$ENV_FILE" \
  -p 8080:8080 \
  -v "$DATA_DIR:/data" \
  "$IMAGE"

sudo docker image prune -f
sudo docker ps --filter "name=$CONTAINER"
