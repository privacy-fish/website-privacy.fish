---
title: "How the Privacy.Fish Infrastructure Works"
weight: 40
---

## Core Design Concept

The infrastructure core concept is to assign as little trust as possible, and through a system of automation and configuration-management make it very easy to setup "everything from scratch" while only migrating minimal data from the previous systems to minimize attacker persistence. On top of this, only minimal custom code is used and a very high security operating system is chosen (OpenBSD), along only very few packages on the servers.

One of the core trusts lies in SSH to connect to the servers to send email, and SFTP to download email and delete it from the server, as well as manage SSH public keys. We deem it to be more secure to manage a large number of SSH accounts rather than SMTP and IMAP user + password combinations.

For SMTP communication with other servers, opensmtpd is chosen, which does a quick spam check on incoming email and then age-encrypts it with the users SSH public keys before storing it, until the user fetches the encrypted emails via SFTP. The user can then choose to delete the emails from the server or wait for the server to do it automatically after 14 days. Each user gets a 100 MB mailbox, which should do even for a busy two weeks.

## stw.no Hosting Provider Overview

All our servers are located in Norway and hosted with stw.no. We rent two dedicated blade servers as well as two Proxmox hosted VMs (VPS) from stw.no.

## Which Servers exist
### alpha.mail.privacy.fish (hardware)
### omega.mail.privacy.fish (hardware)
### backup.mail.privacy.fish (VPS)
### www.privacy.fish (VPS)

## List of Relevant Git Repositories
### Ansible Playbook Mail
### Ansible Playbook Web
### website-privacy.fish

## Recreating the Servers For The Email Infrastructure

## Recreating the Servers For The Web-Infrastructure

## Weekly Server Lifecycle

## Monitoring
### Setup
### Alerting
### Incident Response
### Incident Reports

## Backups
### What Is Backed Up

/var/log/legal-access.log - SSH / SFTP connections: IP + Port + timestamp
For 12 months

/home/<user>/Maildir
For 7 days

/home/user/.ssh/authorized_keys
For 7 days

### Backup Clearing After Deletion

TODO delete mails in backups when they are deleted via SFTP

### Restore Process

The backups are only for a email server system failure. We do not offer restoring of backups to our users.

