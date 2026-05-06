---
title: "Private Signup Form: The Most Privacy-Friendly One-Time Payment Flow You Can Build"
description: "How Privacy.Fish built a private signup form with self-hosted captcha, one-time payments, temporary payment codes, and minimal signup data."
date: 2026-05-06
draft: false
video: "/videos/blog/the-most-privacy-friendly-one-time-payment-form-you-can-build/anonymous-online-payment.mp4"
show_cta: false
---

Privacy.Fish has the checkout complexity of a lemonade stand: one product, one price, one payment, no subscription plan called Premium Plus Ultra. That makes our life easier than a normal webshop with carts, shipping addresses, discount codes, repeat customers, and refund workflows. Most checkout systems are built to remember the customer forever. We wanted the opposite: confirm that the right account was paid for, then delete the link between payment and account as soon as we can.

Our answer is a private signup form built around one-time payments, temporary payment codes, self-hosted captcha for bot protection, and deletion of the payment-to-account link after 14 days, when the Norwegian consumer withdrawal period has expired.

The implementation is public: the [code for this website](https://github.com/privacy-fish/website-privacy.fish) is open source.

## Why Private Signup Forms Are Harder Than They Look

A normal signup form is usually the beginning of a long identity chain: email verification, password resets, card tokens, invoices, support tickets, fraud checks, analytics, marketing tags, subscription renewals, chargeback handling, and customer dashboards. Each piece may be reasonable alone, but together they make the service remember far more than the first screen suggests.

For Privacy.Fish, the goal is different. We want to know that an account was paid for. We do not want to keep knowing who paid for it.

### The Problem With Normal Checkout Systems

Most checkout systems are built for conversion, accounting, fraud prevention, upsells, subscriptions, and customer support. Privacy is usually added later as a policy document. The technical result is predictable: the checkout provider can see the browser session, IP address, device data, payment method choice, amount, timing, product, and some form of customer identifier.

That may be acceptable for a shop shipping physical goods. It is a bad default for a privacy service.

### The Problem With Subscriptions

Subscriptions are worse for payment privacy because they create a permanent link between a customer, a payment relationship, and a service account. The provider must remember enough state to charge again, retry failed payments, cancel, renew, refund, handle disputes, and answer billing questions.

Privacy.Fish uses a one-time payment because it can be matched once, used to create the account, kept long enough for the refund window, and then disconnected from the account.

## Step 1: Use A Self-Hosted Captcha Method Without Google

A public signup form needs abuse control. If anyone can ask unlimited username-availability questions or reserve unlimited accounts, bots can turn the form into a username-enumeration tool or a pending-account queue. But dropping Google reCAPTCHA or another large third-party challenge provider into the form creates its own privacy problem.

Privacy.Fish uses [Cap](https://capjs.js.org/), an open source captcha system, and serves the widget from our own site instead of loading the challenge from Google or a third-party CDN. The [signup template](https://github.com/privacy-fish/website-privacy.fish/blob/main/layouts/_partials/blocks/signup.html#L1-L13) points the browser to the locally served Cap widget and WebAssembly asset, and the signup backend verifies the Cap token against our configured Cap endpoint before allowing username checks or account reservation.

### Why Captcha Attempts Are Limited

Solving the captcha does not give unlimited access to the username checker. A successful human check creates a short-lived credential with five attempts, and the server consumes that credential when the user checks a username or reserves an account. The relevant backend path is [`handleLookupSession`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/server/server.go#L128-L151).

The form also has rate limits by source IP and limits successful signups. That does not make enumeration impossible. It makes enumeration more expensive, slower, and noisier, without handing the signup session to a third-party captcha provider.

For logs, Privacy.Fish uses the model described in our [logfile privacy information](/documentation/privacy-information/#logfile-privacy-information/): internal security analysis, no third-party lookups, and deletion of ordinary lines when no alert is triggered.

## Step 2: Make The User List Harder To Steal From A Compromised Webserver

A private signup form still needs to answer a basic question: is this username available? The naive implementation is a database query against our email servers user table. That works, but it makes the signup server a valuable target.

### Fixed-Size Username Storage

The webserver hosting the Privacy.Fish website and signup form needs to answer "is this username already taken?", but we do not want it to hold a plaintext list of all real account names. So the deployed username database, `used.bin`, is not a normal customer table. It is a fixed-size binary database, currently one million entries by default. Each entry is 32 bytes: either the `HMAC-SHA256` output for a real username, or 32 bytes of random data.

The real username list is processed on the deployment workstation before it is uploaded to this website's webserver. For every real username, the deployment process computes a keyed hash. The pepper is the secret key for that hash; without it, a copied `used.bin` file is not enough to test guessed usernames offline. The salt is additional input mixed into the hash calculation; it separates this deployment's username database from other databases and from older deployments. Both values are rotated on every deployment, so an attacker who obtains an old and a new copy of `used.bin` cannot simply diff the two files to see which entries are new and infer which accounts were just created.

After the real usernames have been converted into 32-byte entries, the rest of the one-million-entry database is filled with random 32-byte padding entries. The whole database is then shuffled before upload. The fixed size avoids leaking the account count through the file length, and the shuffle avoids putting real entries in predictable positions. The signup backend also validates that `used.bin` has exactly the configured size before using it.

Between deployments, new accounts still have to count as "taken". We do not rebuild the full one-million-entry `used.bin` database every time an account is created, so newly provisioned usernames are written to a small side file, `used.bin.overlay`, in the same HMAC format as the main database. When someone checks a username, the signup server checks both places: the fixed-size `used.bin` database and the overlay file. On the next full deployment, the overlay entries are folded into a freshly rebuilt `used.bin`, the pepper and salt are rotated, and the temporary overlay can start fresh again.

This is not magic encryption of usernames. Usernames are guessable: if an attacker controls the webserver and also obtains the current pepper, salt, and username database, they can still test guessed names. The point is more specific: a stolen file should not immediately be a readable customer list, should not reveal the exact account count, should not make real entries obvious among padding entries, and should not expose stable hashes that can be compared against the previous deployment's database to identify newly created accounts.

### HMAC Checks Instead Of Plain Username Lookups

When a user checks a username, the server normalizes it, computes an `HMAC-SHA256` value with a deploy-rotated pepper and salt, and checks that value against the used and pending stores. The relevant code is [`UsernameHMAC`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/server/server.go#L370-L375) and the [`usernameTaken`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/server/server.go#L354-L367) path.

This is not a password hash. Usernames are low-entropy by nature; many can be guessed. The point is narrower: the signup server can do availability checks without carrying a simple plaintext list of every registered account.

### Why Enumeration Still Needs To Be Expensive

No public username checker can honestly promise that usernames cannot be probed. If an attacker can ask whether `alice` is taken, the answer itself leaks one bit of information. The practical defense is to limit the rate, limit attempts per human check, avoid exposing the whole database at rest, and make large-scale probing visible.

That is what the Privacy.Fish signup form is designed to do.

## Step 3: Ask For Login Keys Instead Of Passwords

Privacy.Fish asks for SSH public keys during signup because we need them to provide the service. While this has nothing to do with signup form security, it still might be of interest to you.

An SSH key pair is used for cryptographic login and for encrypting email with age: the public key can be stored on our servers, while the private key stays on the user's device. When the user connects later, the SSH protocol runs a public-key authentication exchange: the server checks that the client can produce a valid signature with the matching private key, without the private key or a reusable mailbox password being sent to us.

For the signup form, the privacy point is simple: we collect the requested username and the authentication method needed to operate the account. We do not ask for a contact email or phone number, because that is user data we do not want to have in the first place. The backend accepts only `ssh-ed25519` public keys, deduplicates them, and canonicalizes them to the key type and key material; optional OpenSSH comments such as `user@laptop` are stripped because they can accidentally reveal names, devices, or workplaces. The exact parser is [`internal/sshkey/parse.go`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/sshkey/parse.go#L18-L47).

## Step 4: Create A Temporary Payment Code

When the user presses the reserve button, the browser sends the requested username and SSH public keys to the Privacy.Fish signup backend. The backend validates the username, canonicalizes the public keys, generates a temporary payment reference, and stores a pending signup record with the username, canonicalized SSH public keys, payment reference, and creation time. The backend then sends the payment reference back to the browser, so the user can include it with whichever payment method they choose. The relevant code is [`handleSignup`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/server/server.go#L185-L257).

The payment code is intentionally not the username. It is meant to travel through the payment channel: a bank transfer reference, PayPal message, crypto confirmation, or note inside a cash letter can all contain the temporary payment code. What should not travel through those channels is the account identifier. The user should not put their requested email address, real name, existing email address, phone number, or other identifying text next to the payment code.

### Why The Account Is Reserved Before Payment Arrives

Cash by letter, SEPA transfers, and cryptocurrency payments are not always instant. All payment methods, including credit card and PayPal, still require manual account activation because the webserver hosting the Privacy.Fish website and signup form is deliberately decoupled from the mail infrastructure.

So the account is reserved first. The username is held for 30 days, and the user gets a payment code. When payment arrives with that code, Privacy.Fish can create the correct account.

Because the signup form does not collect contact information, Privacy.Fish cannot send "your account is ready" notifications. That is less convenient, but it avoids another permanent identity channel.

## Step 5: Offer Payment Methods Without One Central Checkout Provider

Most sites put every payment method behind one checkout provider. That is convenient, but it also creates a central observer that can see the payment code, browser session, IP address, device data, payment choice, amount, timing, and often the payment identity.

Privacy.Fish avoids that where possible. We prefer separate payment paths instead of one central processor for everything.

### Cash By Letter

Cash by letter is the strongest practical option we offer. The user puts EUR 20 in an envelope, includes only the temporary payment code, and sends it to Privacy.Fish. We explicitly tell users not to send coins, because coins tend to tear through paper when letters are processed.

Cash is not magic anonymity. Postal systems, cameras, fingerprints, handwriting, printer marks, postage payment, and local postal rules can all matter. But from our side, Privacy.Fish receives payment and a code, not a bank account, card number, PayPal identity, cryptocurrency address, or billing profile.

### Cryptocurrency

Cryptocurrency is supported for convenience, not because it is automatically private. Public blockchains are public, many users buy crypto through KYC exchanges, and wallet infrastructure can leak information.

Privacy.Fish uses cold receiving wallets so private keys do not live on the website or mail servers. For BTC, ETH, and USDT, common wallets and exchanges do not give us a reliable payment-code memo field, so users submit the transaction hash and payment code through a separate crypto confirmation form.

### SEPA, PayPal, Credit Card, And Other Convenience Methods

SEPA, PayPal, credit cards, and similar convenience methods are not private from the payment systems that process them. Banks, PayPal, card networks, card issuers, acquirers, payment processors, fraud systems, and compliance systems may see payment identity, account details, device data, timestamps, transaction references, and other payment metadata depending on the method.

The privacy improvement here is not that these payment methods become anonymous. They do not. The improvement is that Privacy.Fish does not route every method through one additional central checkout provider, and the payment message should still contain only the temporary payment code, not the requested account name.

## Step 6: Delete The Link Between Payment And Account

The strongest part of the design is not the payment code itself. A payment code is just a temporary identifier. The privacy feature is deleting the association after it has served its purpose.

During account creation, Privacy.Fish needs to know three things together: the requested username, SSH public keys, and payment code that proves payment. On the disk of the webserver hosting the Privacy.Fish website and signup form, pending signup data is stored as an age-encrypted payload for the administrator's public key; the pending record keeps metadata such as the username HMAC and encrypted ciphertext. The relevant code is [`internal/store/pending.go`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/store/pending.go#L59-L91).

That does not mean the live web process never sees the submitted form. A form handler necessarily receives the form while processing the request. The point is that pending signup data is not left on the disk of the webserver hosting the Privacy.Fish website and signup form in plaintext, and that webserver cannot create mail accounts automatically because, for security reasons, there intentionally is no network path from it to the mail infrastructure for that.

After payment and account creation, Privacy.Fish keeps the payment-code-to-account association only as long as needed for the consumer refund period. Then the association is deleted. From then on, Privacy.Fish has a list of received payments and a list of active accounts, but no system record that answers which payment created which account.

Subscriptions cannot work this way. A subscription requires ongoing billing state. The provider has to know which account belongs to which billing relationship so it can renew, fail, cancel, refund, or dispute the subscription later.

## Final Thoughts: What Most Privacy-Friendly Means

Privacy.Fish exists for one stubborn reason: to build email infrastructure that collects as little user data as we can get away with, even when that makes the product less convenient and the business model less profitable.

In signup terms, that means no real name, no contact email, no password, no subscription account, no central checkout provider unless a payment method requires it, no username inside the payment message, and no permanent payment-to-account link. Or shorter: collect as little data as possible.

A private signup form is not created by adding a privacy checkbox to a normal checkout. It has to be designed around data minimization from the start. For Privacy.Fish, that means self-hosted bot protection, limited username checks, SSH keys instead of passwords, temporary payment codes, separate payment paths, and deletion of the payment-to-account link.
