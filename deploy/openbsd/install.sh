#!/bin/sh
set -eu

SERVICE_USER=_pfishsignup
SERVICE_GROUP=_pfishsignup
DATA_DIR=/var/db/pfish-signup
ENV_DIR=/etc/pfish-signup
ENV_FILE="$ENV_DIR/env"

if ! grep -q "^${SERVICE_GROUP}:" /etc/group; then
	groupadd "$SERVICE_GROUP"
fi

if ! id "$SERVICE_USER" >/dev/null 2>&1; then
	useradd -d "$DATA_DIR" -s /sbin/nologin -g "$SERVICE_GROUP" -L daemon "$SERVICE_USER"
fi

install -d -m 0700 -o "$SERVICE_USER" -g "$SERVICE_GROUP" "$DATA_DIR" "$DATA_DIR/pending"
install -d -m 0750 -o root -g "$SERVICE_GROUP" "$ENV_DIR"

if [ ! -f "$DATA_DIR/used.bin.overlay" ]; then
	install -m 0600 -o "$SERVICE_USER" -g "$SERVICE_GROUP" /dev/null "$DATA_DIR/used.bin.overlay"
fi

install -m 0755 pfish-signup /usr/local/bin/pfish-signup
install -m 0555 deploy/openbsd/rc.d/pfish_signup /etc/rc.d/pfish_signup

if [ ! -f "$ENV_FILE" ]; then
	install -m 0640 -o root -g "$SERVICE_GROUP" /dev/null "$ENV_FILE"
	echo "Created $ENV_FILE; fill it before starting the service."
fi

rcctl enable pfish_signup
