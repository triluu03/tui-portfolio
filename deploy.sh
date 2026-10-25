#!/usr/bin/env bash
set -euo pipefail

SERVER=root@37.27.92.253

log() {
  printf '[%s] %s\n' "$(date '+%Y-%m-%d %H:%M:%S')" "$*"
}

# Keep the terminal open when the script finishes (success or failure) so the
# user can read the output. Automatically skipped when not run interactively.
if [[ -t 0 ]]; then
  trap 'read -r -p "Script finished. Press Enter to close..."' EXIT
fi

log "Starting deploy to ${SERVER}"

log "Building linux/arm64 binary..."
GOOS=linux GOARCH=arm64 go build -o ./builds/portfolio .
log "Build complete"

log "Uploading binary..."
scp ./builds/portfolio "${SERVER}:/home/portfolio/portfolio_new"
ssh "${SERVER}" -i ~/.ssh/id_hetzner '
  chown portfolio:portfolio /home/portfolio/portfolio_new &&
  chmod +x /home/portfolio/portfolio_new &&
  mv /home/portfolio/portfolio_new /home/portfolio/portfolio &&
  sudo systemctl restart portfolio &&
  systemctl is-active portfolio
'

log "Deploy finished successfully"
