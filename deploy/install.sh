#!/bin/sh
set -eu

SERVICE_USER=pfish-signup
DATA_DIR=/var/lib/pfish-signup
ENV_DIR=/etc/pfish-signup

if ! id "$SERVICE_USER" >/dev/null 2>&1; then
  useradd --system --home "$DATA_DIR" --shell /usr/sbin/nologin "$SERVICE_USER"
fi

install -d -m 0700 -o "$SERVICE_USER" -g "$SERVICE_USER" "$DATA_DIR" "$DATA_DIR/pending"
install -d -m 0700 -o "$SERVICE_USER" -g "$SERVICE_USER" "$ENV_DIR"
if [ ! -f "$DATA_DIR/used.bin.overlay" ]; then
  install -m 0600 -o "$SERVICE_USER" -g "$SERVICE_USER" /dev/null "$DATA_DIR/used.bin.overlay"
fi
install -m 0755 pfish-signup /usr/local/bin/pfish-signup
install -m 0644 deploy/pfish-signup.service /etc/systemd/system/pfish-signup.service

if [ ! -f "$ENV_DIR/env" ]; then
  install -m 0600 -o "$SERVICE_USER" -g "$SERVICE_USER" /dev/null "$ENV_DIR/env"
  echo "Created $ENV_DIR/env; fill it before starting the service."
fi

systemctl daemon-reload
systemctl enable pfish-signup
