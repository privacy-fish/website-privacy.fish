---
title: "Launching privacy.fish with privacy above convenience"
description: "Why privacy.fish starts from a different premise than mainstream email providers, and why that changes everything from storage to recovery."
date: 2026-04-22
video: "/videos/data-minimization-by-default.mp4"
---

Most email services optimize for convenience first. privacy.fish does not.

The service starts from a simpler question: **what would email look like if the provider tried to retain as little useful data as possible?** That question affects every technical and operational decision on the platform.

## Why build it this way

A free and open society depends on private communication. If people cannot decide when and with whom to share their thoughts, every other right becomes harder to exercise in practice.

An email address is required to participate in most modern online services. That makes email infrastructure unusually powerful. When the default product model is "host everything forever, gather telemetry, and make recovery frictionless", users are pushed toward systems that ask for too much trust.

privacy.fish takes the opposite approach:

- Mail is stored encrypted using your SSH public keys.
- Retrieval and deletion happen over SSH / SFTP.
- Server-side retention is intentionally limited.
- Recovery depends on your backup SSH keys, not personal identity data.

## What "privacy first" changes in practice

This approach creates a product that is less polished than conventional webmail. There is no browser inbox, no "forgot password" flow tied to a phone number, and no promise of indefinite server-side storage.

Those are not missing features by accident. They are consequences of a stricter trust model.

When a provider avoids collecting unnecessary data, it loses some of the levers that mainstream services use to make recovery, support, and growth easier. privacy.fish accepts that tradeoff because privacy claims without product constraints are mostly marketing.

## Who this is for

privacy.fish is meant for people who prioritize privacy over convenience, are comfortable with SSH / SFTP workflows, and want a mailbox that minimizes retained data.

It is a poor fit for users who want classic webmail, permanent server-side storage, or a safety net built around personal identity data.

That narrower scope is intentional. The goal is not to be email for everyone. The goal is to be useful for people who want a different set of guarantees.
