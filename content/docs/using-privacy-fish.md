---
title: "Using Privacy.Fish"
weight: 20
---

@richtermac, I need to write each of the below sections for: CLI, macOS, Linux, Windows, Android, iPhone
So I need some kind of 6 boxes with the logo of the OS that the user can click on to get to
documentation/using-privacy-fish/command-line/account-creation
documentation/using-privacy-fish/android/account-creation

And so on. I'm not sure how else to sort this, but I'm happy for input. All of the headers below "# Using Privacy.Fish" are individual depending on which device you use. The only exceptions are "### Making a Private Payment" and "### How we Activate Your Account".

## Account Creation
### Installing the Client App
### Generating SSH Keys For Your Devices
### Using the Signup Form
### Making a Private Payment

### How we Activate Your Account

Within 24 hours of receiving your payment together with the associated payment code, Privacy.Fish administrators manually create your account.

The signup form does not ask for identifying contact information. Because of that, we have no way to notify you when the account is ready. The safest approach is to wait up to 24 hours after payment before trying to use the account.

Even though we offer instant payment methods like PayPal and credit cards, we deliberately do not connect the signup web server to the mail servers to automatically create new accounts. The web server running the signup form is not trusted to create accounts automatically on the email server infrastructure.

When you submit the signup form, your requested username, SSH public keys, and temporary payment code are stored in age-encrypted form on the web server. Within 24 hours, the data is downloaded, dencrypted by the administrator and then securely deleted from the webserver. Using the decrypted account data, our adminstrators manually check for new arrived payments, and create your account on the email servers.



## Connecting the Client App to Your Mail User Agent

## Sending Email
### Normal Delivery
### Unsafe TLS Warnings
### Using Aliases

## Receiving Email
### How Email is Downloaded and Decrypted in Your App

## Deleting Email From Our Servers

## Determining and Using Your Aliases

## Managing Devices and SSH Key Pairs
### Adding a Device
### Removing a Device
### Creating a Backup SSH Key
### Restoring an Account With a Backup SSH Key

