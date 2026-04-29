---
title: "Payment privacy and account creation"
description: "How payment data is handled, which payment methods are accepted, and why accounts are created manually."
weight: 5
---

privacy.fish tries to minimize the amount of personal linkage between payment activity and long-term account records.

## Payment model

The service uses a one-time payment for lifetime access to one account.

Accepted payment methods include:

- credit card
- PayPal
- cash
- voucher
- cryptocurrencies
- SEPA bank transfer

Anonymous payment options are available through Bitcoin or cash by mail.

## Retention of payment data

Payment data is retained for only **three months**. After that, it should no longer be linkable to your email account.

This is another example of the broader minimization strategy: keep only what is required to operate the service, then stop keeping it.

## Why manual account creation matters

Newly paid accounts are manually created by administrators. The signup and payment side is physically decoupled from the backend mail infrastructure.

That choice slows down activation, but it reduces automation and trust coupling between the public-facing purchase flow and the mail system itself.

## Refunds

Refunds are possible within one week of payment, except for cash and cryptocurrencies.
