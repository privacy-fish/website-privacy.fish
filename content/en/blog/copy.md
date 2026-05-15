---
title: "The Most Privacy-Friendly One-Time Payment Form You Can Build"
description: "How Privacy.Fish built a privacy-friendly one-time payment form with temporary payment codes, no subscriptions, and minimal signup data."
date: 2026-05-06
draft: true
video: "/videos/blog/norway-email-privacy-law-why-it-beats-germany-switzerland-the-eu-and-us-for-private-email/norway-email-privacy-law.mp4"
show_cta: false
---


Its easy for us to say - we have only one product, one price for it and require only a one time payment. stripping our payments of privacy information about our customers might be simpler than stripping a large webshop with recurring customers.

it is our goal to design all systems and business decisions to be as private as possible. for our model it was important to know only that an account was payed for, not who payed for it.

we have build the possibly most private while still being practical signup form with one-time payment function you can think of. lets break up the signup form piece by piece and explain what its about:


## 1 Human check first

to prevent bots reaking havoc with the signup form, we require a captcha before we can start.

the open source captcha backend (add a link which one we use) is hosted on the same server as the website and does not log or send data to third parties.

## 2 Choose your username

- we generate a db with all real usernames + random usernames until we get to 1 million entries, and hash + salt + pepper the entries.
- when the user checks a email address is free, it asks the server, which hashes the users choice and checks the db
- bots cant enumarate cause captcha

when deploying, we download new account registrations, add them to the db, then rebuild db and rotate salt + pepper each time we deploy the website. this way the hashes change all the time and attacker cant track db changes over time.


## 3 Add your SSH public keys

has nothing to do with signup for privacy and is just technicality of our service.


## 4 Payment reference

describe how the payment link works


## 5 Payment methods

write about how we offer cash by letter, and joke that you have to write dont send coins.

say we use cold wallets for crypto, so we can just monitor the blockchain for new payments and we have a seperate contact form just for crypto payments where users send the transaction hash + payment code, so we know which acccount the payment was for (contact form on our website, so they dont have to send an email / whatsapp -> involve a third party)

say that we offer sepa, paypal and cc too, but say that we think that 20 euro by letter is a very low barrier too to send one time.



## 6 Setup guides

has nothing to do with signup for privacy and is just technicality of our service.
