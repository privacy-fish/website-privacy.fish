---
title: "How the data lifecycle works"
description: "How mail is stored, downloaded, deleted, and backed up on privacy.fish."
weight: 2
---

The platform is designed to minimize retained data. That principle matters more than convenience features.

## Storage at rest

Mail is stored encrypted using your SSH public keys with **age**.

The provider is not offering a typical permanently hosted mailbox with a polished server-side interface. The goal is to reduce the amount of useful plaintext data sitting on the service.

## Retrieval and deletion

Mail is handled through **SSH / SFTP** workflows.

In practice, that means:

- you download age-encrypted `.eml` files,
- optional client-side tooling can make downloaded mail available to desktop mail clients via POP3,
- you can manually delete mail from the server over SFTP.

## Automatic deletion

If you do not manually delete mail, it is automatically deleted after **14 days**.

That retention window exists to reduce the amount of mail the service holds at any given time. privacy.fish is not meant to be the long-term archive for your inbox.

## Backups

Operational backups are made hourly and kept for seven days.

That does **not** mean deleted messages are a user-facing restore feature. You should assume that if something matters, you need to download and preserve it yourself.

## Practical rule

Treat the server as a short-term encrypted holding area and your own devices as the long-term location for important mail.
