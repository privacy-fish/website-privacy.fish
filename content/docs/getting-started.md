---
title: "Getting started"
description: "What you need before signing up, what happens after payment, and the basic lifecycle of a privacy.fish account."
weight: 1
---

privacy.fish is built for users who are comfortable with SSH / SFTP and who want stronger privacy guarantees than mainstream webmail providers usually offer.

## Before you sign up

Make sure the service matches your workflow.

You should be comfortable with:

- using SSH keys,
- downloading files via SFTP,
- keeping your own local backups,
- living without classic browser-based webmail.

You should not sign up expecting:

- indefinite server-side storage,
- account recovery through personal identity checks,
- a polished convenience layer that hides the underlying workflow.

## What you buy

A one-time payment of **€ 20** buys one account for the lifetime of the service.

An account includes:

- one main address,
- unlimited addresses across supported domains,
- support for up to 10 SSH public keys,
- access on supported desktop platforms and Android workflows.

## What happens after payment

Accounts are not provisioned automatically on the backend. The signup and payment side is intentionally decoupled from the mail infrastructure.

After payment:

1. an administrator reviews the purchase,
2. the account is manually created,
3. provisioning is completed within 24 hours.

This is slower than consumer SaaS onboarding, but it reduces how tightly the public signup flow is coupled to the mail backend.

## First things to do

When your account is ready:

1. confirm that your SSH public keys are correctly associated,
2. save your backup key material somewhere safe,
3. test SFTP access,
4. make sure you understand the deletion policy before treating the service like storage.

If you lose your recovery keys and your configured devices, there is no alternate identity-based recovery path.
