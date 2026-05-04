---
title: "Risk Management and Compromise Response"
description: "How privacy.fish thinks about live compromise, impact boundaries, operational risk, and response after an incident."
weight: 45
video: "/videos/secure-email-workflow.mp4"
---

# Risk Management and Compromise Response

This page describes how Privacy.Fish thinks about live compromise, impact, and response. It is intentionally separate from the infrastructure design page, because this is risk management: what can happen, what can be detected, and what has to be done after a compromise is suspected or confirmed.

## How a Successful Attack Can Affect Privacy.Fish Infrastructure

The impact of a live compromise depends on which server is compromised. Privacy.Fish runs two physical servers with OpenBSD. On those physical servers we use `vmm`, OpenBSD's virtualization tool, to run separate OpenBSD VMs for different parts of the email infrastructure, including `in`, `spam-in`, `out`, `spam-out`, `fetch`, `keys`, and `monitoring`. The backup server runs separately as a VPS hosted with stw.no on Proxmox, which means the underlying virtualization layer is not under our control. Root access on one VM is serious, but it is not automatically the same as root access on every VM or on the physical VMM host. A compromise of only the incoming mail server (VM) could intercept or manipulate incoming mail before it is encrypted, but would not automatically compromise outgoing mail or stored encrypted mail. A compromise of only the outgoing mail server could intercept or manipulate outgoing mail, but would not automatically compromise incoming mail or stored encrypted mail. A compromise of only the key-management server could manipulate customer SSH keys. A compromise of the backup or migration path could carry bad state forward. A compromise of monitoring could hide alerts, or we could not have checks in place to detect a specific malicious behavior. A compromise of a physical VMM host would be broader than compromise of a single guest VM.

An attacker with `root` access on a specific server can do whatever that server role allows: change configuration, alter files, inspect logs, interfere with running services, and manipulate traffic that passes through that server. Stored encrypted mail is the special case where tampering should be visible to the user: if an existing `.age` mail file is modified, age authentication makes decryption fail. That protects against silent modification of already-encrypted stored mail, but nothing else.
