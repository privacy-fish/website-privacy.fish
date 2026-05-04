---
title: "OpenBSD setup"
description: "Basic OpenBSD notes for generating SSH keys and preparing for privacy.fish access."
weight: 7
---

OpenBSD is a natural fit for the privacy.fish workflow because SSH tooling is available by default.

## Generate a key

Create a dedicated SSH key for privacy.fish:

```sh
ssh-keygen -t ed25519 -f ~/.ssh/privacy-fish
```

Add the public key from `~/.ssh/privacy-fish.pub` to the signup form.

## Keep a backup key

Add a second key from another trusted device if you have one. Recovery only works through configured backup SSH keys.

## After provisioning

After your account is created, use the connection details from the getting started guide and your local OpenBSD SSH / SFTP tools to access the service.
