---
title: "Privacy Information"
weight: 60
---

# Privacy Information

## Payment Privacy Information

Privacy.Fish uses a one-time payment instead of recurring subscription fees for one reason: protecting your privacy.

A one-time payment creates less ongoing payment metadata than monthly billing or prepaid subscription periods. It also makes private payment methods more realistic, because sending 20 € by letter to another country is a much smaller hurdle than doing that every month.

The hurdle to having a private and secure email account should be low.

We considered making the service donation-based, similar to Signal Messenger, but decided against it to reduce spam account creation and keep the service usable.

### How the Signup Payment Code Protects Your Privacy

Privacy.Fish uses temporary payment codes to prevent your payment information from permanently identifying the email account you paid for.

The goal is simple: payment confirms that an account should be created, but after creation, the deleted link can never again associate that payment with an account.

What we want to permanently know is limited to two separate facts: which email accounts have been paid for, and which payments we have received. We do not want to keep a permanent record of which payment created which account.

We of course have no influence over what data the post office stores for cash by letter, what our Norwegian bank stores for SEPA transfers, what our payment partners store for credit cards and PayPal, or what cryptocurrency blockchains store about the money you send us.

When you create an account by submitting a username and your SSH public keys, that username is reserved for 30 days. If payment is not completed within those 30 days, the username becomes available for registration again.

When you make the payment, you include the temporary payment code with the transaction. This lets us match the received payment to the reserved username and create the correct account.

We keep the payment-code-to-username association for 14 days so refund and withdrawal requests can still be handled during the Norwegian legal refund period for consumers.

After those 14 days, we permanently delete the payment-code-to-username association. From that point on, Privacy.Fish no longer has any record that links your payment data to your email account. Even if you could prove to us that you made a payment, we could not say which account it was for.

### Why We Only Require a One Time Payment

Privacy.Fish charges once instead of keeping a monthly or yearly subscription alive. This is worse for predictable recurring revenue, and we assume it will make us lose money compared with a normal subscription model.

We do it anyway because it lets us delete the payment-to-account association. With a subscription, Privacy.Fish would have to keep the data linking the account to a recurring payment method so we could charge it again, handle failed renewals, and manage cancellations. With a one-time payment, we only need the temporary payment code long enough to create the account and handle the legal refund period of 14 days. After that, the link between payment and username can be deleted.


### Payment Provider Privacy Information

The following chapters describe how much each payment option violates your privacy. Ordered from least to worst.

#### Cash By Letter

Cash by letter is our preferred payment method for users who want the strongest practical payment privacy.

##### Cash By Letter Recipient Post Address

Please put 20 euros into an envelope, add a letter that contains ONLY! the temporary payment code and NOT! your username, and send it to:

```text
Privacy.Fish
Fjordstreet 5
12345 Oslo
Norway
```

This gives Privacy.Fish only the information we need: a payment code and the payment itself.

##### Privacy Information for Cash By Letter

Cash by letter is still not magic privacy. Depending on how you prepare and send the letter, you may still expose information to other parties. The envelope, paper, and banknote can contain fingerprints, DNA traces, handwriting, printer marks, or other physical traces. If you send the letter from a post office, shop, kiosk, or other staffed location, there may be surveillance cameras. If you pay postage by credit card or another traceable payment method, that payment is logged. If you add a return address, the letter can be identified more easily.

Postal operators and shipping companies may also collect shipment data, but the exact data depends on the country, postal operator, shipping product, and how you pay.

We do not know whether every postal operator scans letter contents, but we assume it is technically possible to detect that an envelope contains a banknote or to extract the payment code without opening the letter. The privacy benefit is therefore not that cash by letter leaves no trace anywhere.

The benefit is that Privacy.Fish does not receive a bank account, card number, PayPal account, cryptocurrency address, or other digital payment identifier that most likely is directly linked to you.

Do not submit your username, email address, real name, return address, or other identifying information with the cash letter. Submit only the temporary payment code and the payment itself.



#### Cryptocurrencies

##### Cryptocurrency Payment Addresses

Please send 20 euro to:

| Cryptocurrency | Example address |
| --- | --- |
| Bitcoin (BTC) | `bc1qexampleaddressdonotuse0000000000000000000000` |
| Ethereum (ETH) | `0xExampleAddressDoNotUse0000000000000000000000` |
| Tether (USDT, ERC-20) | `0xExampleAddressDoNotUse0000000000000000000000` |

##### Communicating The Payment Code

BTC, ETH, and USDT do not provide a reliable payment-code memo field across common wallets and exchanges. After sending the transaction, submit the temporary payment code, cryptocurrency, and transaction hash through the [crypto payment confirmation form](https://github.com/signup/crypto-confirmation-form). Do NOT submit your username, email address, real name, or other identifying information! The confirmation form exists so you can send us the temporary payment code alongside the blockchain transaction, so Privacy.Fish can create the reserved account. This association is deleted after the Norwegian legal refund period for consumers of 14 days after the payment has been received.

##### Privacy Information for Cryptocurrency Payments

Cryptocurrency payments are supported for convenience, not because they are automatically private. For most users, cryptocurrency is bought, held, and sent through a regulated exchange such as Coinbase, Kraken, Binance, or another provider that knows the user's identity. If you pay Privacy.Fish from such an exchange, the exchange knows who you are, when you paid, how much you paid, and the Privacy.Fish address you paid to.

Public blockchains add another problem: transactions are visible forever. Depending on the cryptocurrency, the sender address, receiver address, amount, time, transaction fee, and later movement of funds can be analyzed by anyone. If you use wallet software that connects to third-party nodes or APIs, those providers may also learn your IP address and the addresses or transactions you query.

Privacy.Fish could use a fresh receiving address for every incoming payment. This would prevent an attacker from looking at a single public address and seeing every other payment to the same address. However, using multiple incoming addresses does not create the privacy benefit many people expect. If the user paid from a Know Your Customer (KYC) exchange, the exchange still knows the withdrawal. If Privacy.Fish later moves many received payments together, blockchain analysis can link those receiving addresses again. And because Privacy.Fish is a company, cryptocurrency eventually has to be exchanged, spent, or converted into ordinary money. That cash-out process can create another point where transactions become linkable.

Because of that, Privacy.Fish does not pretend that managing many cryptocurrency receiving addresses gives users strong payment privacy. We use one long-lived cold receiving address per supported cryptocurrency. The private keys are generated and stored offline, not on the website or mail servers. This keeps the operational setup simple and keeps private keys away from internet-facing systems, but it also means payments to that address are publicly visible as payments to Privacy.Fish.

If you want strong practical payment privacy, use cash by letter. If you use cryptocurrency, assume that the exchange, blockchain observers, wallet infrastructure, or later cash-out process may be able to link the payment to Privacy.Fish.

If you can obtain and send cryptocurrency legally anonymously, then sending cryptocurrency is better for your privacy than paying from an exchange account connected to your identity. It is also better to send from a wallet used only for this payment, on a device and network you trust, instead of from your everyday phone, browser wallet, or exchange account. This still does not make the payment guaranteed anonymous, but it reduces the number of places where the payment can be linked back to you.

Submit the temporary payment code, cryptocurrency used, and the transaction hash through the [crypto payment confirmation form](https://github.com/signup/crypto-confirmation-form). Do not attempt to send us your username, email address, real name, or other identifying information.


#### SEPA Bank Transfer

##### SEPA IBAN and BIC

Please send 20 Euros to:

BANK Name: Sparbank Oslo
IBAN: NE 1234 5678 9012 3456 78
BIC: 45678938
Subject: **YOUR-PAYMENT-CODE-HERE**

##### Privacy Information for SEPA Bank Transfers

Our bank can see and store information about incoming bank transfers. Based on SpareBank 1's privacy information, this can include financial information such as account numbers and transaction data, as well as payment information and other data required to provide banking services, prevent abuse, meet accounting duties, and comply with legal obligations. SpareBank 1 says personal data is generally stored as long as necessary for the purpose it was collected for, unless law or regulation requires longer storage. Their retention examples include documentation for anti-money-laundering and counter-terrorist-financing work for up to 10 years after a completed transaction or ended customer relationship, and information required under Norwegian bookkeeping rules for up to 10 years. ([SpareBank 1 privacy information](https://www.sparebank1.no/nb/bank/om-oss/personvern.html), [SpareBank 1 privacy policy PDF](https://www.sparebank1.no/content/dam/SB1/vedlegg/veiledninger/Privacy-policy.pdf))

For an incoming SEPA transfer, Privacy.Fish therefore assumes that our bank may store at least the sender name, sender account or IBAN, recipient account, amount, timestamp, payment reference / message field, and internal transaction metadata. We do not control the privacy policy, retention period, compliance systems, or fraud-monitoring systems of the bank you use to send the transfer.

Your own bank may store more information than we ever receive, and we have no way to know exactly what they store.

Do not submit your username, email address, real name, or other identifying information with the SEPA bank transfer. Submit only the temporary payment code in the transaction subject.

#### PayPal

##### PayPal Payment Address

Please send 20 Euros to:

```text
paypal@privacy.fish
```

##### Privacy Information for PayPal Payments

PayPal is convenient, but it is not a private payment method. PayPal can process account identifiers, contact information, financial information, transaction information, device information, IP addresses, cookies, fraud signals, and other data needed to provide its services, prevent abuse, resolve disputes, comply with law, and operate its payment network. PayPal may also disclose information to other account holders, merchants, service providers, financial institutions, authorities, and other parties when needed to complete transactions or comply with legal obligations. ([PayPal Privacy Statement](https://www.paypal.com/li/legalhub/privacy-full?locale.x=en_US), [PayPal User Agreement](https://www.paypal.com/us/legalhub/paypal/useragreement-full?locale.x=us_US))

For a PayPal payment to Privacy.Fish, assume PayPal may know who you are, which PayPal account or payment method you used, when you paid, how much you paid, the recipient account, the transaction ID, and the payment message containing the temporary payment code. Privacy.Fish may also see the sender information PayPal exposes to the recipient account. We cannot control what PayPal stores, how long PayPal keeps it, what fraud or compliance systems process it, or what your linked bank or card provider stores.

Do not submit your username, email address, real name, or other identifying information with the PayPal transaction. Submit only the temporary payment code in the PayPal transaction message.

#### Why We Do Not Use a Payment Processor

For most payment methods, Privacy.Fish avoids a central checkout or payment processor because it would collect even more data than the payment method already exposes by itself. A central payment processor can see the payment code, browser session, IP address, device data, payment method choice, transaction timing, amount, and often the payment identity behind the transaction.

Instead, we prefer separate payment paths where possible. Cash by letter does not have to touch a payment processor at all. SEPA exposes data to the banks involved, but not to an additional checkout provider. Cryptocurrency exposes blockchain and exchange data, but does not have to expose browser checkout metadata to a payment processor. PayPal exposes data to PayPal, but does not require another processor in front of it.

Credit cards are the exception. Card payments require a payment processor, so they are the least private payment option we offer.

#### Credit Card

For credit cards, sadly, we have to use a payment processor.

##### Credit Card Payment Processor

Privacy.Fish uses Dintero Checkout for credit card payments. Dintero is a Norwegian payment provider based in Oslo, and Dintero Checkout is their online checkout product for merchants. ([Dintero Checkout](https://dintero.com/products/dintero-checkout), [Dintero Terms for Checkout](https://dintero.com/legal/terms-for-checkout))

##### Privacy Information for Credit Card Payments

Credit card payments are the least private payment method offered by Privacy.Fish. When you pay by card, Dintero, the card network, the acquirer, your card issuer, and possibly other payment-service providers can process information about the transaction. Dintero's privacy policy says it may process contact information, account information, payment card data, account numbers, browser and device data, usage data, purchase history, transaction-related data, fraud-prevention data, and data required by financial, accounting, or other legal obligations. Dintero also says the merchant is typically a controller for data needed to manage the purchase, while a payment service provider is a controller for data needed to process the payment. ([Dintero Privacy Policy](https://dintero.com/no/legal/privacy-policy), [Dintero End-Customer Terms](https://dintero.com/no/legal/end-customer-terms))

Privacy.Fish does not want your card number. The point of using Dintero is that the card payment is handled by their payment infrastructure instead of Privacy.Fish storing or processing raw card details. But this also means a third-party checkout and card-payment chain exists, and that chain can see much more payment metadata than Privacy.Fish wants to know.

Do not submit your username, email address, real name, or other identifying information with the credit card payment. Submit only the temporary payment code in the payment form.






## Logfile Privacy Information

### Logfiles Mail Servers

Privacy.Fish is required by norwegian law to keep logfiles of registered accounts for SSH and SFTP logins to the servers, this means each time you send or fetch mail, or when you modify your SSH public keys. For each login, we store the timestamp, username, source IP address and source port. We are required to keep this data for 12 months for Norwegian legal compliance.

Besides this legal requirement, Privacy.Fish is striving to "have as little data as possible".

Logfiles from the servers are streamed to an internal analysis system, processed for security alerts, and discarded when no alert is triggered. The analysis system does not send log data to third parties or perform external lookups such as IP reputation checks. When suspicious activity is detected, Privacy.Fish stores the alert and the minimal log lines needed to understand it. Administrators may temporarily enable more detailed logging during incident response, but this is limited to the time required to deal with the incident and disabled again after debugging. After debugging, those temporary logs are deleted. If logs have to be retained for legal reasons or an incident report, they are minimized and anonymized as much as possible.

### Logfiles Web Server

## Website Privacy Information
