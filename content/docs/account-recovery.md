---
title: "Account recovery"
description: "Recovery depends on backup SSH key pairs. There is no identity-based fallback."
weight: 3
---

privacy.fish does not use phone numbers, identity questions, or personal profile data for recovery.

## Recovery model

Account recovery relies on **backup SSH key pairs** that you control.

This is consistent with the service goal of collecting as little identifying information as possible. It also means recovery is less forgiving than mainstream products.

## What this means for you

You are responsible for:

- generating recovery-capable key material,
- storing backups safely,
- making sure you can still access them when a device is lost.

## What happens if you lose your keys

There is no alternate recovery channel.

That is not a warning banner added for dramatic effect. It is a real constraint of the product. If you lose both access and your backup SSH keys, there is no identity-based process that can safely recreate trust without undermining the entire privacy model.

## Recommended practice

- Keep at least one offline backup of your recovery keys.
- Test that the backup is usable before you need it.
- Do not store your only recovery path on the same device you use every day.

If your threat model includes device loss, key hygiene is not optional.
