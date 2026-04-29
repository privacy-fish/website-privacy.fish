---
title: "Why mail expires after 14 days"
description: "The 14-day retention window is a security and trust decision, not a storage optimization trick."
date: 2026-04-19
---

privacy.fish does not try to be a permanent archive. Mail on the server is temporary by design.

## The point of limited retention

Every extra day of retained email increases the value of the server as a target. It also increases the amount of data that could be exposed through operator error, compromise, or legally binding disclosure.

That is why privacy.fish uses a limited retention model:

- Mail is downloaded by you.
- You can manually delete it through SFTP.
- If you do nothing, it is automatically deleted after 14 days.

This does not eliminate risk, but it constrains it. The platform is designed to keep less, for less time.

## Why not offer "restore my deleted mail?"

Because the service is not built around permanent server-side custody.

Backups exist for operational resilience, not as a user-facing archive. They are made hourly and kept for seven days, but deleted mail is not something you can ask the service to recover for you. That would push the system back toward the trust-heavy model privacy.fish is trying to avoid.

If you want long-term retention, the responsibility belongs on your own devices and backups.

## What users should assume

Treat privacy.fish as a delivery and short-term holding system, not as the final resting place for your inbox.

If a message matters, download it promptly and back it up locally. The service is strongest when the server is a temporary staging point rather than a permanent repository.
