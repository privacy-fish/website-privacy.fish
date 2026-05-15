---
title: "There Is Better Authentication Than Username And Password: Why We Trust SSH Instead"
description: "Why Privacy.Fish uses SSH keypairs instead of username-and-password email login."
date: 2026-05-06
draft: false
video: "/videos/blog/there-is-better-authentication-than-username-and-password-why-we-trust-ssh-instead/ssh-as-sign-in.mp4"
show_cta: false
---


## The Problems with Classic Username and Password Authentication

Username + password authentication is the default choice for most web services because everyone understands it, every framework supports it, and it fits normal login, reset, and support workflows.

The cost is a predictable operational security problem: users reuse passwords, phishing pages can collect them, weak passwords can be guessed, and username/password pairs leaked from one service can be tried automatically against other services (credential stuffing). The service also has to maintain password-verification data that becomes valuable breach material.

It also provides poor device separation. If the same password is entered into every client and stored wherever those clients decide to store it, compromise of one device, app profile, or copied credential can become compromise of the whole account.


## How SSH Keypair Authentication Works

SSH keypair authentication works differently from username + password login. The user creates a keypair: a private key and its corresponding public key. The public key is provided at signup and copied to our email servers. The private key stays with the user. When the user connects, the server does not ask for the private key and the client does not upload it. Instead, OpenSSH performs a cryptographic challenge-response exchange: the client signs authentication data with the private key, and the server verifies that signature with the public key it already has.

We can do even more: it just so happens that compatible SSH public keys can also be used as recipients for age-encrypted files, which lets us use the same keypair model for stored email encryption.

### Why SSH Public Keys Can Be Public

An SSH public key can be posted publicly on the internet because it only lets others check signatures made by the corresponding private key; it cannot create those signatures itself. Hence, an attacker compromising our users' SSH public keys does not gain the ability to log into their accounts.

Even a well-configured password-hash database can be attacked by testing password guesses offline; an SSH public key does not give the attacker a comparable path to derive the matching private key. The sensitive part is the private key, and that should remain on the user's device, hardware token, encrypted backup, or offline storage.

### The User Experience of SSH Keypair Authentication

For Privacy.Fish users, this means account access starts with managing keypairs instead of memorizing a mailbox password.

Each of our users can have two Account Control Public Keys and up to eight Email Public Keys.

The eight Email Public Keys can be used to send, download, and delete emails from our servers.

The two Account Control Public Keys can add or remove Account Control Public Keys as well as Email Public Keys, as well as everything Email Public Keys can do. We recommend storing them securely and offline.

Users who prefer to keep it simple and only have one device can sign up and use the service with only one keypair, used as the Account Control Public Key.

That is unfamiliar for many email users. People understand passwords intuitively because every website uses them. Keypairs are more technical. We chose SSH keypair authentication because Privacy.Fish puts privacy and security above convenience.




## Why We Trust OpenSSH as the Authentication System for Our Email Service

This is not only laziness, although yes, OpenSSH being installed by default on our OpenBSD-based servers certainly helps. Privacy.Fish uses OpenSSH for user authentication because it offers us a mature, heavily tested authentication stack.

### OpenSSH Is Widely Deployed and Heavily Audited

[OpenSSH](https://www.openssh.com/) is one of the most widely deployed remote-access stacks in the world. It has been attacked, audited, patched, hardened, and operationally abused for decades. That does not make it perfect. It does mean that trusting OpenSSH is a very different risk decision from trusting a small custom login protocol or a webmail-specific authentication layer.

There is another important distinction: TLS around IMAP, POP3, or SMTP AUTH protects the transport, but the application still often ends in a password-style login. SSH combines encrypted transport and public-key authentication in one mature protocol. We do not think "SSH good, TLS bad" is the right mental model. The real issue is that password-based email protocols usually require more exposed authentication surfaces than we want to run.

### One Authentication Stack Is Better than Two

If Privacy.Fish ran IMAP for reading, plus POP3 for downloading and deleting, plus SMTP AUTH for sending, plus webmail for browser access, every one of those surfaces would need to accept user authentication. Each surface would need rate limiting, monitoring, hardening, logs, patching, abuse handling, and configuration review. Even if every component were well-chosen open source software, the total attack surface would be larger.

So we keep the user authentication model centered on SSH. That is not normal email UX, but it is a much smaller and cleaner design.

### The Same Public-Key Model Fits Age Encryption for Storing Emails

The same public-key model also fits how Privacy.Fish stores email in encrypted form. For each incoming message, Privacy.Fish writes an age-encrypted file whose recipients are the user's compatible SSH public keys. One of the matching private keys are needed to decrypt that `.age` file, so the server can store encrypted delivery storage.

[Age](https://github.com/FiloSottile/age) also gives authenticated encryption. If an encrypted email file is modified, decryption fails instead of silently producing altered plaintext. That does not solve every possible server-side manipulation problem, but it gives us the most important baseline property: stored email should not be useful to read without the user's private key.


## Why We Accept the Usability Cost

SSH keypair authentication asks more from users, but it lets Privacy.Fish ask for and store less.

### SSH Keys Are More Work than Passwords

SSH keys are harder than passwords. Users need to generate keys, keep private keys safe, copy public keys during signup, understand which key belongs to which device, and keep backups.

A password login is easier for most people, especially because most people are not focused on their privacy when choosing an email provider. Privacy.Fish is built for users who decided that protecting their privacy is worth an extra step because they want authentication that is considered more secure than username + password login.



## Is SSH Authentication Worth Considering for Your Service?

We use SSH keys because they fit the Privacy.Fish philosophy: collect less, store less, expose less, and let the user control the important secrets. Password-based email is convenient, but that convenience tends to move risk and data into the provider's hands. SSH keys move more responsibility to the user, but they let us avoid reusable mailbox passwords, skip entire IMAP/POP3/SMTP/Webmail AUTH password surfaces, and build email access around OpenSSH and age instead of a traditional mailbox-login stack.

That is not the right tradeoff for everyone. It is the right tradeoff for Privacy.Fish.
