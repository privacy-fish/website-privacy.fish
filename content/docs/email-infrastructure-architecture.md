---
title: "Email Infrastructure Architecture"
weight: 40
---

# Infrastructure Security Concepts

The Privacy.Fish infrastructure security model combines several ideas: trusting as little as possible, keeping running systems small, and making compromised servers easy to replace. This section explains the main building blocks of that approach: OpenBSD, OpenSSH, OpenSMTPD, minimal custom code, short-lived server installations, replaceable admin workstations, Norwegian hosting, and the parts of the infrastructure that cannot be fully controlled.

Privacy.Fish trades server-side convenience for a smaller trust model: fewer services, less retained data, fewer passwords, short-lived servers, and mail stored primarily on devices you control.

Security is never finished. There is always another control, another audit, or another layer that could be added. Privacy.Fish publishes its setup and design decisions on [github.com/privacy-fish](https://github.com/privacy-fish) so professional security reviewers and the wider technical community can inspect them, challenge them, and help improve the system over time.

## Things We Have to Trust Ultimately

In order to be able to do any administrative work, we have to trust the following things ulimately:

- stw.no WebUI for managing the server and the VPS (needs username + password + TOTP 2FA)
- cdn.openbsd.org for providing us with install.img (VPS) and install.iso files (hardware server)
- github.com
- the admin workstation (described below)
- the admin workstation SSH private key to administrate the servers
- the client to be able to protect his master key

## Hosting Provider servetheworld.no

All our servers are located in Norway and hosted with servetheworld.no / stw.no. To run all of our production infrastructure, we rent two dedicated blade servers as well as two Proxmox hosted VMs (VPS) from stw.no.

ServeTheWorld is a Norwegian hosting provider based in Oslo. For Privacy.Fish, the most important property is jurisdiction: the servers are hosted in Norway, under Norwegian law. Privacy.Fish is a Norwegian LLC.

STW states that it deploys its dedicated servers in Oslo data centers under Norwegian law. They also state that privacy is supported by Norwegian legislation and that customer data is not handed to another party without a Norwegian court order. ([Dedicated servers](https://stw.no/en/dedicated-server/), [ServeTheWorld overview](https://stw.no/en/))

From a physical and operational security perspective, STW describes its Oslo data center as a high-security facility with capacity for around 4,000 servers, protection against fire, theft, power failure and flooding, redundant network lines, redundant UPS, A/B power, firewalling and diesel generator backup. For colocation, STW also describes redundant cooling, redundant power paths, security around the clock and 24/7 emergency personnel. Their terms include a GDPR data processing agreement, which makes STW a data processor for hosted customer data, but Privacy.Fish still avoids treating the hosting provider as a privacy boundary: mail content is age-encrypted before storage, and the server design minimizes the data that STW or the infrastructure ever needs to handle in readable form. ([Colocation](https://stw.no/en/colocation/), [Data Processing Agreement](https://stw.no/wp-content/uploads/legal/nor-legal/Databehandleravtale_%28DBA%29.pdf))

## OpenBSD as Operating System

OpenBSD is widely considered to be one of the most secure operating systems. All our servers and admin workstations run OpenBSD.

The reason we use OpenBSD is not only its reputation, but its security philosophy. OpenBSD is designed to be secure by default: non-essential services are disabled, the base system is continuously audited, and many security mechanisms are built into the operating system itself instead of added later as optional hardening. Examples include privilege separation, privilege revocation, chroot jailing, W^X memory protection, randomized memory allocation, `pledge(2)`, and `unveil(2)`. For Privacy.Fish, this matters because we want the operating system to help reduce the damage a compromised process can do, even before our own application design is considered. ([OpenBSD Security](https://www.openbsd.org/security.html), [OpenBSD Innovations](https://www.openbsd.org/innovations.html))

## Mostly OpenBSD base packages

Our email hosting stack is mainly built around two programs that are shipped with every OpenBSD installation by default: OpenSSH and OpenSMTPD. They are part of the so-called OpenBSD base system, which means they are maintained, audited, documented, and released together with the operating system instead of being extra third-party services we have to add and track separately.

## Minimal Custom Code

The largest part of our codebase is [configuration management code](https://github.com/privacy-fish/playbook-infrastructure-privacy.fish) to install and configure software on the servers.

TODO links to the custom scripts / tools

## OpenSMTPD For Everything SMTP

OpenSMTPD is our open source mail server of choice. It is open source software developed as part of the OpenBSD project and is even included in OpenBSD’s base system. The mail server talks to gmail.com when you send or receive mail to/from example@gmail.com. It also fits our security approach because it is small, simple, and designed with security in mind.

## SSH For Client Interaction

SSH (Secure Shell) is the standard way to securely log in to most of today's Linux and OpenBSD servers over the internet. It uses an encrypted connection and usually cryptographic keys (like your SSH public keys) instead of reusable passwords.

We trust OpenSSH because OpenSSH is widely deployed, open source, and heavily scrutinized. Shodan currently lists OpenSSH as the top SSH product with more than 40 million indexed instances, and the OpenSSH project describes it as the premier tool for remote login with the SSH protocol. This popularity is not proof of safety by itself, but it does mean the software is continuously exercised, audited, attacked, and patched in real-world use. ([OpenSSH](https://www.openssh.com/), [Shodan SSH search](https://www.shodan.io/search?query=ssh)).

We need OpenSSH installed to administer our servers anyway, so we use it for more than one purpose. Using SSH for client interaction gives Privacy.Fish one key-based authentication system, fewer exposed login services, and no reusable mailbox password.


### Why SFTP instead of IMAP/POP3

Privacy.Fish prioritizes privacy and security above features. Our job is simple: send and deliver email, as private and secure as possible. We think this approach fits privacy-conscious email users who prefer fewer server-side features over giving us more mailbox data, even if that state is encrypted.

IMAP is useful because it lets clients work with server-side folders, message flags, search, selective downloads, synchronization, and long-lived mailboxes. IMAP servers such as [Dovecot](https://www.dovecot.org/), [Cyrus IMAP](https://www.cyrusimap.org/), and [Courier-IMAP](https://www.courier-mta.org/imap/) can provide synchronized folders, read/unread state, sent-mail copies, search, indexes, concurrent clients, and long-lived server mailboxes. That is useful, but it requires extra running code and extra mailbox state on the server, even if encrypted.

Privacy.Fish gives up those features deliberately. We avoid collecting and storing that mailbox state because we do not want that data, not even encrypted. Privacy.Fish is temporary encrypted delivery storage, not a long-term server-side mail archive. That also fits the 14-day retention model.

With Privacy.Fish, multiple devices can download and decrypt the same encrypted email before one of the devices deletes it from our servers. Your mail client can still sort emails into folders, search them, and store them indefinitely on your own devices. What Privacy.Fish does not support is server-side synchronization of folder moves, read/unread state, sent-mail copies, search indexes, or other mailbox state between devices. If you want that, it has to happen on your devices or in your own tooling, for example with a Thunderbird plugin.

In our opinion, the gain on security by avoiding the extra running lines of code from dovecot, courier or cyrus, as well as avoiding storing the associated mailbox states, even if encrypted, is worth the loss of convenience.

### Why SSH Instead of SMTP Auth

SMTP AUTH is the usual way for email clients to authenticate before sending mail. It commonly relies on account passwords or password-like tokens, and it adds another authentication path that has to be configured, monitored, secured, and exposed to the internet.

Privacy.Fish avoids that by using SSH for authenticated sending. The client or app submits a prepared email message through an SSH command. A minimal `.eml` file looks like this:

```text
From: j.doe@privacy.fish
To: friend@example.com
Subject: Hello

Hello from Privacy.Fish.
```

On the CLI, sending that file looks like this:

```bash
cat msg.eml | ssh j.doe@out.mail.privacy.fish
```

This lets us use the same SSH key-based account model for sending mail that we already use for downloading mail and managing devices. The server can restrict what the SSH key is allowed to do, avoid reusable mailbox passwords, and keep the authentication surface smaller than running both SSH and SMTP AUTH.

### Why SSH For User Authentication Management

We chose SFTP for managing OpenSSH public keys because we already had SSH installed ;)



## Role Isolation Using OpenBSD VMM

We currently run multiple OpenBSD VMs on one physical server (with the "failover" server being hosted in a different datacenter) and each VM has a dedicated role. For example, incoming mail, outgoing mail, stored encrypted mail, key management, and monitoring run in separate VMs. Communication between VMs happens over SSH and is restricted as much as possible.

This isolation limits the impact of some compromises: root access on one VM does not automatically mean root access on every other VM. However, OpenBSD VMM is still another layer of running code and therefore another security boundary we have to trust. The stronger alternative would be to run each role on separate physical hardware. That is a possible future improvement, but for now we focus resources on the controls that provide the most security value for the current size of the service. Separate physical servers would reduce the risk of an attacker moving from one VM to another through the virtualization layer.




## Limiting Attacker Persistence With Rebuilds

After a server compromise, the clean way forward is to set up a fresh server and migrate only the required data, which has to be verified not to have been manipulated.

Privacy.Fish does this every week because it is always reasonable to assume that a 0day vulnerability may have affected us, and now it was found and patched. Weekly rebuilds are meant to remove access from exploits that no longer work after the patch. They do not protect against an unpatched 0day that still works.

### How 0day Exploits Affect Privacy.Fish

0day exploits are software bugs that are not publicly known yet, or at least not yet patched by the affected software maintainers. OpenBSD has a strong security record, but it is not magic: serious vulnerabilities still get found and patched.

In January 2020, [Qualys found an OpenSMTPD bug](https://www.qualys.com/2020/01/28/cve-2020-7247/lpe-rce-opensmtpd.txt) where a specially crafted email address could make the mail server execute shell commands as root in some configurations. In February 2020, [Qualys found another OpenSMTPD bug](https://www.qualys.com/2020/02/24/cve-2020-8794/lpe-rce-opensmtpd-default-install.txt) where a malicious or compromised remote mail server could inject commands into OpenSMTPD's internal delivery state. Also in February 2020, an OpenBSD `vmm` issue was reported where a VM guest could write into host memory; OpenBSD published a [February 17, 2020 errata patch](https://www.openbsd.org/errata66.html).

Because of this, it is reasonable to treat any internet-facing server as potentially compromisable by a sufficiently capable attacker. The infrastructure is therefore designed around a simple question: if an attacker gets root access today, and the exploit they used is later found and patched, how do we make sure that access does not survive the next rebuild? This does not protect against an unpatched 0day that still works after the rebuild. In that case, the attacker could simply exploit the fresh server again. The rebuild model is mainly about preventing persistence after the original entry point has been closed.

0day exploits are a part of the modern IT world and there is no way around it. Writing all of the software we use ourselves, in a "very secure way", would not remove the bugs every human (and machine) introduces into code. Because of this, we chose software that is already very well audited.

### Reducing Attacker-Persistence By Reinstalling Everything From Scratch Once a Week

In a traditional long-lived server setup, the operating system may run for months or years. If an attacker uses a 0day exploit before it is publicly known and patched, installing the security update later may close the original bug, but it may not remove the attacker's existing foothold. A hidden backdoor, modified binary, changed configuration file, or added key can remain on the server after the patch is installed.

Privacy.Fish tries to reduce that risk by treating servers as replaceable infrastructure. Instead of only upgrading servers in place, we regularly rebuild them from scratch using configuration management. The current target lifetime for a physical server installation is one week, and running systems check and install OpenBSD security upgrades every 15 minutes during that time.

This changes what an attacker has to achieve. After the 0day exploit that the attacker used is found and patched and the servers are rebuilt, the attacker would need to persist in the very small amount of data that is carried from the old server to the new one, or exploit the method that transfers that data (rsync) - otherwise the next rebuild would lock him out for good.

### Attacker Persistence in Dynamic Data

Dynamic data is data our users generate and modify, so we have to migrate this data from the old to the new server (at least) weekly. That is:

- stored encrypted user emails (`.age` files)
- signed account-key state, including master keys, mail-operation keys, recovery keys, versions, previous-state hashes, and signatures

The actual `~/.ssh/authorized_keys` files are generated from signed account-key state during server setup and propagation.

Everything else about the server is controlled by code that is on [github.com/privacy-fish](https://github.com/privacy-fish). The configuration management code and all other automations are published on GitHub. The list of registered users is not published, because it is private operational data, but it is used by configuration management to recreate the required user accounts on a fresh OpenBSD server installation.

#### Mitigating The Manipulation Of Stored Encrypted Email Files

The current protection comes from the age encryption of the email files itself. Stored emails are encrypted for the users' SSH public keys, and age encryption includes authentication. If an existing encrypted mail file is modified, decryption fails instead of producing silently changed mail. The client application interprets age decryption errors clearly and warns the user when encrypted mail cannot be decrypted.

This does not detect every possible mailbox manipulation on the fetch server. If an attacker has root access on the fetch server, they could delete stored `.age` files, replace them, or add new `.age` files encrypted to the user's public keys. Modified files are detected by age decryption failure, but deletion, replacement, and newly added valid `.age` files are not detected by age alone.

Signing stored encrypted mail files is a TODO. It is tracked as infrastructure work in the [playbook-infrastructure-privacyfish issues](https://github.com/privacy-fish/playbook-infrastructure-privacyfish/issues).

Adding signatures would interfere with attacks in the following way:

| Attack | Age protects | Signing protects | How signing protects | Impact if undetected |
| --- | --- | --- | --- | --- |
| Modify a stored `.age` file | Yes | Yes | Age already detects this during decryption, because modified encrypted files fail age authentication. Signing would add an earlier check: any byte change would make signature verification fail before decryption. | The user could receive silently changed encrypted mail. With age alone this should already fail during decryption, but signatures would catch it earlier. |
| Replace a stored `.age` file with a different file | No | Yes | The replacement file would not match the original signature unless the attacker also had the signing key. The fetch server does not have access to the signing server; only the incoming mail path can request signatures for newly encrypted mail. | The client could accept a different encrypted message as if it were the original one. |
| Add a new `.age` file | No | Yes | The client would require every downloaded `.age` file to have a valid Privacy.Fish signature. Unsigned files, or files with invalid signatures, would be rejected. | The client could process a fake encrypted message, for example to confuse the user or attack the mail client. |
| Delete a stored `.age` file | No | No | Per-message signatures would not detect deletion if the attacker deleted both the encrypted file and its matching signature file. | The user may never see the deleted message. Detecting deletion would require a signed mailbox list or similar global mailbox state. |

...IF an attacker manages to gain access on the fetch server. Trusting OpenBSD and OpenSSH is a very fair bet.

We continuously improve our service and prefer to honestly tell our users what is implemented and what is not (yet).

#### Mitigating The Manipulation Of SSH Authorized Keys

Privacy.Fish uses a dedicated key-management server where the users can upload, delete and change their SSH public keys that are used to send and download email as well as for encrypting email files to store. Only the user's master SSH keys can log into the keys server and perform updates. This keeps normal mail keys separate from account-management keys.

Privacy.Fish treats key changes as signed user intent. Instead of trusting the current `authorized_keys` file on the key-management server as the source of truth, the user uploads signed account-key state: which master keys manage the account, and which mail-operation keys are allowed to access mail. New mail-operation keys must be signed by the current trusted master key. A new master key must be signed by the previous trusted master key, or by an already registered recovery key.

The signed account-key state is stored outside the key-management server and is used to generate the actual `authorized_keys` files for the mail servers. Fresh servers regenerate `authorized_keys` from the last verified signed state instead of copying the old server's file directly. This means that an attacker with root access on the key-management server can still break live service or submit invalid key data, but cannot create a valid new key state unless they also have the user's current trusted master private key.

When migrating from an old server to a fresh server, stored user emails are copied with `rsync`, with the new server pulling the files from the old server. SSH public keys are regenerated from the last verified signed account-key state.

### Conclusion

Because the server setup is automated, a new mail server only needs minutes to be installed with OpenBSD, configured, the old data migrated and the IPs switched, and it is ready to handle Privacy.Fish email.

The goal is not to claim that compromise is impossible. The goal is to make attacker persistence as unlikely as technically possible by keeping servers short-lived, rebuilding them quickly and reliably using automation, and carrying forward only the minimum data required to keep the service working.

That is the Privacy.Fish tradeoff in one sentence: smaller server state, fewer running services, less retained data, and more responsibility on the user side, especially around devices, keys, and local mail storage.





## Minimal And Replaceable Admin Workstation

Compromising an administrator can be easier than attacking a hardened server directly.

The admin workstation is arguably even more attackable than the servers, because it is not located inside a datacenter, but in the administrator's home. We exclusively work from home office.

The administrator's workstation is, like the servers, designed to be replaced completely regularly and consists of:
- Raspberry Pi
- MicroSD card
- External Screen
- Keyboard
- Mouse
- Network Cable (or Wi-Fi where unavoidable)
- 2FA TOTP device for stw.no

We use configuration management code available on [github.com/privacy-fish](https://github.com/privacy-fish) to set up the workstations so you can take a look for yourself.

The only "dynamic data" that has to be migrated is the administrator's SSH private key. The password to login to stw.no, as well as the password to use the SSH private key, can be memorized.

The administrator's workstation has pf firewall rules in place to only allow the admin to SSH to the servers and to open the WebUI of stw.no, our server hosting provider.

Additionally we use mullvad.net on all our workstations, just in case, with a VPN to Norway.

This is the most reasonably secure workstation setup we could think of.





## Things We Cannot Control

The rebuild model reduces the chance that an attacker can persist inside a mail server, but it does not eliminate every dependency around the server. Some layers cannot be fully verified or regenerated by Privacy.Fish alone: the hosting provider control panel, the provider account recovery process, server hardware and firmware below the operating system, source code and configuration management repositories on GitHub, backups, and the migration process between old and new servers.

The mitigations are layered. Provider access is protected with 2FA and restricted to the admin workstation. OpenBSD base-system updates are signed and applied through the normal OpenBSD update process. Server state is generated from configuration management instead of being hand-maintained. Backups are meant to restore only the data that cannot be regenerated. Migration should verify the dynamic data that is copied into a fresh server.

If our source repositories on GitHub, the firmware / BIOS of the servers, backups, our STW account or STW itself, or the migration process were compromised, a rebuild could still carry risk forward. The goal is to keep those dependencies visible, small, and replaceable where possible.










## Which Servers exist


### alpha.mail.privacy.fish (hardware)
#### in.mail.privacy.fish (VMM)
#### spam-in.mail.privacy.fish (VMM)

#### out.mail.privacy.fish (VMM)
#### spam-out.mail.privacy.fish (VMM)

#### fetch.mail.privacy.fish (VMM)
#### keys.mail.privacy.fish (VMM)

### omega.mail.privacy.fish (hardware)

Omega is the exact same hardware as the physical alpha blade server. Using configuration management and automation, we can set up the Privacy.Fish email infrastructure the same way as on alpha, which we do regularly.

After migrating the emails and SSH public keys of our users to the other server, we switch over the public IP addresses of our services to the freshly set up server.

Hence, we switch operations between alpha and omega weekly, while each time setting up the server from scratch.

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
