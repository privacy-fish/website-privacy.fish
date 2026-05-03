---
title: "Administrator Workstation Design"
weight: 50
---

The core concept behind the hardware of a Privacy.Fish administrators workstation is to harden it, but still consider it exploitable and hence completely replace it frequently.

The admin should only need three things to setup a new workstation:
- stw.no username and password
- stw.no 2FA TOTP device
- SSH private key file for the Privacy.Fish email servers
- SSH private key password

## Hardware Acquisition

The safest route to a secure workstation is to frequently replace it completely. For this, a Raspberry Pi with an SD card with Raspberry Pi OS pre-installed or similar will do niceley, as they are cheap and readily available. Peripherals like keyboard, mouse, screen can also not permanently be trusted. Additionally, one USB stick is required to install OpenBSD onto the SD card directly from the Pi before using it.

## Setting Up a New Workstation

Using the Raspberry Pi OS from the SD card, OpenBSD can be flashed onto the USB stick, which can then be used to install OpenBSD onto the micro SD card.

## Generating a new SSH Private Key

Using the following command:

```bash
ssh-keygen -t ed25519 -a 100 -o
```

A secure password has to be given to the SSH private key.

## Migrating the SSH Private Key to a New Workstation

An ed25519 private key is reasonably short and can be typed of from another screen, or printed on paper. There is no need to use a USB stick or similar to migrate only the SSH private key.
