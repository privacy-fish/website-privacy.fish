---
title: "References"
description: "Reference notes for account limits, supported domains, storage duration, compatibility, and operational assumptions."
weight: 80
video: "/videos/data-minimization-by-default.mp4"
---

## Limits
### Devices per Account

You can use 10 devices per account. You can do more if you use SSH wrong.

We allow up to 10 SSH public keys uploaded. These are used to give you access to the service, as well as encrypt your received email.

Each device should have its own private and public SSH key. You should not copy private keys from one device to another, but always generate a new SSH keypair for each device. The only exception is when you create a backup SSH keypair that you want to store offline, for example on a USB stick.

### Rotating Aliases
### Supported Domains

Your username and alias always works with all domains listed here:

- @privacy.fish
- @pfi.sh
- privacyfish.net
- datenschutzfisch.de
- detaeschutzfischli.ch

### Storage Duration and Size

When our servers receive an email for you, it is encrypted before the encrypted email is written to disk. You then have up to 14 days to download the email. You can then delete the email from our servers, or it will be automatically deleted after 14 days.

Your email is 100 MB in size and can not be extended. If your mailbox is full, incoming email will be rejected and the person that send you an email will receive an email with the error "SMTP 552 Mailbox full".

100 MB is plenty of space for two weeks. Here is [https://www.emailtooltester.com/en/blog/email-usage-statistics/](the source we base this calculation on):

```
80.6 Emails per day × 0.075 MB per email × 7 days ≈ 42.3 MB per week
```

We recommend to download your emails at least weekly and then delete them from our servers.

## Compatibility
### Supported Operating Systems

Currently only the command line client is supported.

We are working on graphical apps for all other operating systems.

### Supported Email Clients

The client app uses your downloaded and decrypted emails (.eml files) to serve them over POP3 to any email client app / mail user agent you choose. All email clients that support POP3 and SMTP should work.

### Known Limitations




## Troubleshooting

### Customer Support

If you are experiencing problems with using the service, please open an issue on [https://github.com/privacy-fish](github.com/privacy-fish) describing your problem and if it makes sense a screenshot.

For technical issues, we do not offer support by email or any other way than github.

### Account Recovery

As we delete the payment to user account association 14 days after receiving your payment, we can only identify which natural person created the account until then. After that, we have no way of knowing which human actually owns this account.

TODO describe the way key priviledges work







## Legal

### Legal Contact

Please send us an email to legal@privacy.fish

### Terms and Conditions

You can find our terms and conditions at https://privacy.fish/legal/terms-and-conditions

### Refunds

Refunds are not possible for cash-by-mail and cryptocurrency payments.

Refunds are only possible until 14 days after we have received your payment. If you have created your account within that time, please send an email to refunds@privacy.fish and include the payment code you used to send the payment, as well as the time and payment method you used.

We will then reply to your email within five working days.

### Abuse Handling

We act only on valid Norwegian legal orders, or when needed to protect Privacy.Fish infrastructure during an attack.

#### Unauthorized Payments

If you have detected an unauthorized payment to us, please contact legal@privacy.fish with a screenshot of the transaction.

#### Unauthorized Account Use

As we have no way of identifying which human owns which Privacy.Fish account, we can not aid you in recovering your account in any way.

### You Received Spam or Virus From a @privacy.fish Account

Please forward the Spam or Malicious Email to spam@privacy.fish

### Lawful Requests

Please send us an email to legal@privacy.fish
