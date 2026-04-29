# OpenBSD deployment notes

These files prepare the signup backend for an OpenBSD VPS.

The Go binary can be cross-compiled from macOS or Linux:

```sh
GOOS=openbsd GOARCH=amd64 go build -o pfish-signup ./cmd/pfish-signup
```

Copy `pfish-signup`, `public/`, and the `deploy/openbsd` directory to the server.

## Install backend

Run as root from the repository root on the OpenBSD server:

```sh
sh deploy/openbsd/install.sh
```

The script creates:

- `_pfishsignup` user and group.
- `/var/db/pfish-signup` and `/var/db/pfish-signup/pending`.
- `/etc/pfish-signup/env`.
- `/usr/local/bin/pfish-signup`.
- `/etc/rc.d/pfish_signup`.

Fill `/etc/pfish-signup/env` using `deploy/openbsd/env.example`.

Copy the generated `used.bin` to:

```text
/var/db/pfish-signup/used.bin
```

Set permissions:

```sh
chown _pfishsignup:_pfishsignup /var/db/pfish-signup/used.bin
chmod 0600 /var/db/pfish-signup/used.bin
```

Start the service:

```sh
rcctl start pfish_signup
rcctl check pfish_signup
```

Inspect logs:

```sh
tail -f /var/log/daemon
```

Health check:

```sh
ftp -o - http://127.0.0.1:8080/healthz
```

Expected output:

```text
ok
```

## Admin commands

Run admin commands with the same environment as the daemon:

```sh
doas -u _pfishsignup sh -c 'set -a; . /etc/pfish-signup/env; pfish-signup admin list-pending'
doas -u _pfishsignup sh -c 'set -a; . /etc/pfish-signup/env; pfish-signup admin export <id>'
doas -u _pfishsignup sh -c 'set -a; . /etc/pfish-signup/env; pfish-signup admin consume <id> <username-hmac-hex>'
doas -u _pfishsignup sh -c 'set -a; . /etc/pfish-signup/env; pfish-signup admin expire-old'
```

## Caddy on OpenBSD

If using Caddy on OpenBSD, install it with packages and use `deploy/openbsd/Caddyfile` as the starting point.

```sh
pkg_add caddy
install -d -m 0755 /etc/caddy
install -m 0644 deploy/openbsd/Caddyfile /etc/caddy/Caddyfile
rcctl enable caddy
rcctl start caddy
```

Deploy Hugo output to:

```text
/var/www/privacy.fish
```

Keep `pfish-signup` and Cap bound to loopback. Only ports 80 and 443 should be public.

## OpenBSD-specific differences from Debian

- Use `rcctl` and `/etc/rc.d/pfish_signup` instead of systemd.
- Use `_pfishsignup` instead of `pfish-signup` as the service account.
- Use `/var/db/pfish-signup` instead of `/var/lib/pfish-signup`.
- Use `/var/log/daemon` or syslog configuration instead of `journalctl`.
- Use `pkg_add caddy` and `rcctl` if you keep Caddy.
- If replacing Caddy with native OpenBSD `httpd`, add a separate reverse proxy layer such as `relayd` for `/api/*` and `/captcha/*`.
