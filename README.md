# privacy.fish Website

Marketing and documentation site for `privacy.fish`, built with Hugo and Tailwind CSS v4.

## Quick Start

```bash
npm install
hugo server
```

The site runs at `http://localhost:1313/`.

For a production build:

```bash
hugo --cleanDestinationDir
```

## Stack

- Hugo
- Tailwind CSS v4 via Hugo's built-in `css.TailwindCSS`
- Markdown content for the homepage, blog, and documentation pages
- Partial-based block layout for the homepage sections
- Go backend for the privacy-preserving signup flow

## Current Structure

The homepage is driven by the `blocks` array in [`content/_index.md`](./content/_index.md).

Each homepage block maps to a partial in [`layouts/_partials/blocks`](./layouts/_partials/blocks):

- `hero`
- `intro`
- `counter`
- `features`
- `compat`
- `audience`
- `how-it-works`
- `tradeoffs`
- `threat`
- `pricing`
- `cta`

Additional site sections:

- [`content/blog`](./content/blog) for blog posts
- [`content/documentation`](./content/documentation) for documentation pages

The shared page templates live in:

- [`layouts/home.html`](./layouts/home.html)
- [`layouts/_default/list.html`](./layouts/_default/list.html)
- [`layouts/_default/single.html`](./layouts/_default/single.html)

## Main Files To Edit

- [`content/_index.md`](./content/_index.md): homepage copy, block order, CTAs, pricing, and section content
- [`content/blog`](./content/blog): blog landing copy and posts
- [`content/documentation`](./content/documentation): documentation landing copy and guides
- [`hugo.toml`](./hugo.toml): site metadata, menus, base URL, logo path, and signup URL
- [`assets/css/main.css`](./assets/css/main.css): Tailwind v4 theme tokens and global base styles
- [`assets/images`](./assets/images): logo and any Hugo-managed image assets
- [`layouts/_partials/header.html`](./layouts/_partials/header.html): header and nav
- [`layouts/_partials/footer.html`](./layouts/_partials/footer.html): footer
- [`layouts/_partials/blocks/*.html`](./layouts/_partials/blocks): homepage section layouts
- [`layouts/_partials/components/image-responsive.html`](./layouts/_partials/components/image-responsive.html): responsive Hugo image rendering
- [`static/js/site.js`](./static/js/site.js): mobile menu, smooth anchors, and nav behavior

## Tailwind CSS v4

Tailwind is configured in [`assets/css/main.css`](./assets/css/main.css) using `@theme`.

The file is intentionally limited to:

- design tokens
- base element styles
- global typography defaults

Prefer Tailwind utility classes in templates over adding custom component CSS.

## Build And Deploy

- Local development: `hugo server`
- Production build: `hugo build --gc --minify --cleanDestinationDir`
- Debian/systemd backend deploy files: [`deploy/install.sh`](./deploy/install.sh) and [`deploy/pfish-signup.service`](./deploy/pfish-signup.service)
- OpenBSD backend deploy files: [`deploy/openbsd`](./deploy/openbsd)

## Signup Backend

The signup backend is the Go service `pfish-signup`. It listens on `127.0.0.1:8080` by default and serves:

- `POST /api/lookup-session`
- `POST /api/check-username`
- `POST /api/signup`
- `GET /healthz`

The `/signup/` page is wired for the full flow:

1. The user enters a username and up to 10 `ssh-ed25519` public keys.
2. The user solves the locally hosted Cap captcha.
3. The page calls `/api/lookup-session` to mint a 5-try credential.
4. The page calls `/api/check-username` when the user presses "Check availability".
5. The page calls `/api/signup` when the user presses "Reserve account for 30 days".
6. The page displays the one-time payment reference and next steps.

To render the real Cap widget, set `cap_site_key` in `hugo.toml` or the signup page front matter and serve the widget script from this site:

```toml
cap_site_key = "your-cap-site-key"
cap_widget_src = "/js/cap-widget.js"
```

Do not load the widget from a third-party CDN in production. Download it during deployment and serve it from `static/js/cap-widget.js`.

Build and test:

```bash
GOCACHE="$PWD/.cache/go-build" GOMODCACHE="$PWD/.cache/go-mod" go test ./...
GOCACHE="$PWD/.cache/go-build" GOMODCACHE="$PWD/.cache/go-mod" go build ./...
```

Required environment for the service:

- `SITE_ORIGIN`: public origin, for example `https://privacy.fish`.
- `CAP_SITE_KEY`: Cap site key.
- `CAP_SECRET`: Cap secret key.
- `ADMIN_AGE_PUBLIC_KEY`: admin age public key; the private key stays on the admin workstation.
- `USERNAME_HMAC_PEPPER`: hex, at least 32 decoded bytes, rotated per deploy.
- `USERNAME_HMAC_SALT`: hex, at least 16 decoded bytes, rotated per deploy.

Optional environment:

- `LISTEN_ADDR`: default `127.0.0.1:8080`.
- `DATA_DIR`: default `/var/lib/pfish-signup`; use `/var/db/pfish-signup` on OpenBSD.
- `CAP_API_BASE`: default `http://127.0.0.1:3000`.
- `USED_STORE_SIZE`: default `1000000`; must match `used.bin`.
- `PENDING_TTL`: default `720h`.
- `CREDENTIAL_TRIES`: default `5`.
- `CREDENTIAL_TTL`: default `10m`.

Local Mac smoke test:

```bash
mkdir -p .local/pfish-data
printf 'postmaster\nroot\n' > .local/users.txt
export USERNAME_HMAC_PEPPER="$(openssl rand -hex 32)"
export USERNAME_HMAC_SALT="$(openssl rand -hex 16)"
go run ./deploy/used-bin-build.go \
  -users .local/users.txt \
  -pepper "$USERNAME_HMAC_PEPPER" \
  -salt "$USERNAME_HMAC_SALT" \
  -size 1024 \
  -out .local/pfish-data/used.bin

go run filippo.io/age/cmd/age-keygen -o .local/admin-age.txt
export ADMIN_AGE_PUBLIC_KEY="$(awk '/public key:/ {print $4}' .local/admin-age.txt)"
export DATA_DIR="$PWD/.local/pfish-data"
export USED_STORE_SIZE=1024
export SITE_ORIGIN="http://localhost:1313"
export CAP_SITE_KEY="dev-site"
export CAP_SECRET="dev-secret"
```

For manual endpoint testing, run Cap locally or this dev-only fake `siteverify` endpoint in another terminal:

```bash
node -e 'require("http").createServer((req,res)=>{res.writeHead(200,{"content-type":"application/json"});res.end("{\"success\":true}")}).listen(3000,"127.0.0.1")'
```

Then start the backend:

```bash
GOCACHE="$PWD/.cache/go-build" GOMODCACHE="$PWD/.cache/go-mod" go run ./cmd/pfish-signup
```

Browser testing needs `/api/*` and the Hugo site on the same origin, matching production. Use Caddy or another local reverse proxy that serves Hugo output and proxies `/api/*` to `127.0.0.1:8080`. Running `hugo server` alone will not proxy `/api/*` to the backend.

Deploy procedure for username privacy:

1. On the admin workstation, generate fresh `USERNAME_HMAC_PEPPER` and `USERNAME_HMAC_SALT`.
2. Build a new fixed-size `used.bin` from the canonical provisioned-user list with `go run ./deploy/used-bin-build.go`.
3. Upload `used.bin` to `$DATA_DIR/used.bin`.
4. Update `/etc/pfish-signup/env` with the new pepper, salt, Cap config, site origin, and admin age public key.
5. Restart `pfish-signup`.

Admin runbook:

1. `sudo -u pfish-signup pfish-signup admin list-pending`
2. `sudo -u pfish-signup pfish-signup admin export <id>` and decrypt the base64 age ciphertext on the admin workstation.
3. Provision the mailbox manually via Ansible.
4. `sudo -u pfish-signup pfish-signup admin consume <id> <username-hmac-hex>` to append the username HMAC to `used.bin.overlay` and delete the pending file.
5. `sudo -u pfish-signup pfish-signup admin expire-old` can be run anytime to remove expired pending signups.

OpenBSD uses the `_pfishsignup` service user and reads the same environment file through the `rc.d` service. Use the OpenBSD-specific runbook in [`deploy/openbsd/README.md`](./deploy/openbsd/README.md).

Threat model summary: username checks are gated by Cap-issued credentials, rate limits, and constant-time scans of `used.bin`, `used.bin.overlay`, and pending records. Pending signup payloads are age-encrypted to the admin public key, so the server can store them but cannot decrypt them. `used.bin` is rebuilt every deploy with a fresh pepper and salt, padded to a fixed entry count, and shuffled so real usernames are not distinguishable from padding entries. `used.bin.overlay` is a short-term bridge for newly provisioned users between rebuilds and should be folded into the canonical user list during the next rebuild.

## Notes

- Generated output is written to `public/`.
