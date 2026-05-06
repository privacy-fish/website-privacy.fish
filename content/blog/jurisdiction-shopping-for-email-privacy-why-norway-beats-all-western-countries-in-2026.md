---
title: "Jurisdiction Shopping for Email Privacy: Why Norway Beats All Western Countries in 2026"
description: "Norway email privacy law vs Germany, Switzerland, UK, Australia and US data retention: what Privacy.Fish must log under Ekomloven."
date: 2026-05-05
draft: false
video: "/videos/security-model.mp4"
show_cta: false
---


## Your Email Provider Is Part Of Your Threat Model

Email data retention laws are one reason email privacy is not only a question of encryption or product design; it is also a question of which country the email provider chose to operate in, and what that country forces the provider to retain, disclose, or technically enable access to. Governments often do not need your device if the law can reach your provider and the data the provider has. That means your email provider’s jurisdiction is part of your privacy model.

Our answer has two parts: operate from Norway, and avoid storing data we do not need.

For the longer operational documentation, see our [jurisdiction comparison](/documentation/jurisdiction/), [privacy information](/documentation/privacy-information/), and [email infrastructure architecture](/documentation/email-infrastructure-architecture/).

## The Uncomfortable Reality Of Email Data Retention In The West

### Germany And Switzerland: The Privacy Darlings With Teeth

Under German [TKG § 170](https://www.gesetze-im-internet.de/tkg_2021/__170.html), anyone operating a telecommunications system used to provide publicly available telecommunications services must, from the start of operation and at their own cost, keep technical facilities for legally ordered telecommunications surveillance and organizational arrangements for immediate implementation; the operator must also notify the Bundesnetzagentur and prove that the facilities and procedures comply with the ordinance and technical directive. [TKÜV § 3](https://www.gesetze-im-internet.de/tk_v_2005/__3.html) applies these rules to operators of systems used for publicly available telecommunications services, but exempts certain systems, including systems with no more than 10,000 users and systems used exclusively for number-independent interpersonal telecommunications services or non-identifier WLAN internet access services with no more than 100,000 users. [TKG § 174](https://www.gesetze-im-internet.de/tkg_2021/__174.html) requires secure electronic interfaces for authority requests: providers with 100,000 or more contractual partners must keep both the interface and the e-mail-based transmission process available, while smaller obligated providers must keep the e-mail-based process available.

The Swiss [VÜPF](https://www.fedlex.admin.ch/filestore/fedlex.data.admin.ch/eli/cc/2018/32/20240326/de/pdf-a/fedlex-data-admin-ch-eli-cc-2018-32-20240326-de-pdf-a-1.pdf) defines email surveillance types that cover metadata and, in some cases, content plus metadata. Article 58 real-time email metadata surveillance includes login/logout status, AAA or subscriber information, client and server IP addresses and ports, protocol, timestamp, data volume, sender and recipient addresses, and sending/receiving mail-server IP addresses and ports for send, receive, mailbox processing, download, and upload events. Article 59 covers real-time email content plus the related Article 58 metadata. Article 62 retroactive email metadata surveillance includes past event date/time/type, participant identifiers, alias addresses, sender/recipient addresses, protocol, server/client IP addresses and ports, delivery status, mailbox login/logout, available download/upload/delete/process/add-message events, and names/IP addresses of sending and receiving email servers.

### The Western Playbook: Retain, Preserve, Intercept

The recurring pattern is not one single privacy problem; it is three different legal powers: compelled metadata retention, compelled preservation or disclosure of data the provider already has, and mandatory interception or delivery capability. In the United Kingdom, a [retention notice under the Investigatory Powers Act 2016](https://www.legislation.gov.uk/ukpga/2016/25/section/87) may require a telecommunications operator to retain all relevant communications data, or any description of it, for up to 12 months, including internet connection records; [section 95](https://www.legislation.gov.uk/ukpga/2016/25/section/95) also prohibits the operator from disclosing the existence or contents of the notice without permission. In Australia, [Telecommunications (Interception and Access) Act 1979 sections 187A, 187AA and 187C](https://www.legislation.gov.au/C2004A02124/2026-01-22/2026-01-22/text/original/pdf/2) require covered service providers to retain subscriber, account, device, source, destination, date, time, duration, communication-type and certain location data for two years, while [section 191](https://www.legislation.gov.au/C2004A02124/2026-01-22/2026-01-22/text/original/pdf/2) requires carriers without a specific determination to maintain interception and delivery capability; in the United States, [18 U.S.C. § 2703](https://www.govinfo.gov/link/uscode/18/2703) lets government require stored communications records and preservation for 90 days plus another 90 days on renewed request, and [18 U.S.C. § 2713](https://www.govinfo.gov/link/uscode/18/2713) applies those preservation, backup and disclosure duties to data in the provider's possession, custody or control even when stored outside the United States.

### Norway: IP, Time, Port, No Destination Log

Under Norwegian [Ekomloven § 3-13](https://lovdata.no/lov/2024-12-13-76/%C2%A73-13), providers of electronic communications networks used for public electronic communications services, and providers of such services, must retain information necessary to identify subscribers from a public IP address and time of communication, or from public IP address, time and port number where the same public IP address is assigned to multiple subscribers at the same time; the same section says destination information must not be retained and sets the retention period at 12 months from the day the communication ends. [Ekomforskriften § 3-5](https://lovdata.no/forskrift/2024-12-20-3410/%C2%A73-5) adds that the retained data must include the time zone used for timestamps, be secured against loss, and not be altered during collection, logging, retention or disclosure.

This does not mean "no lawful access": [Ekomloven § 3-12](https://lovdata.no/lov/2024-12-13-76/%C2%A73-12) requires providers to operate networks and services so legally mandated access to information about end-users and electronic communications is secured, and [§ 3-14](https://lovdata.no/lov/2024-12-13-76/%C2%A73-14) governs disclosure of retained IP-address data to police or prosecution authorities for serious-crime investigations and specified offences.

Norway is [not an EU Member State](https://www.eu-norway.org/eu/norway-and-the-eu/), although much of its relationship with the EU runs through the EEA.




## How Privacy.Fish Turns Norwegian Law Into Real Privacy

The legal baseline matters only if the service is built to keep everything outside that baseline from existing in the first place.

### What We Do To Keep Others From Accessing Your Data

Norwegian law helps only if our own infrastructure does not create extra data for attackers, insiders, payment providers, or later legal requests to reach.

#### Minimal Logging

Privacy.Fish separates mandatory access records from operational security logging. The access records required under the [Norwegian IP-retention duty](https://lovdata.no/lov/2024-12-13-76/%C2%A73-13) are kept for that legal purpose; other server log lines are streamed into an internal analysis system with no internet access, no third-party reputation lookups, and no external analytics. If the analysis finds nothing suspicious, those log lines are automatically discarded instead of stored. If something does look wrong, we keep only the log lines needed to debug the problem, handle the incident, or preserve evidence where legal action requires it.

#### Minimal Data Collection

We build our service using [OpenSSH](https://www.openssh.com/) to download email and then let you immediately delete it not only because OpenSSH has an unusually battle-tested cryptographic and operational track record, but because the only way to protect data is to not store it as a provider. Using multiple devices, you can download the email to all of them or keep backups on removable storage, external drives, or your own cloud storage.

Privacy.Fish uses SFTP for downloaded email so the server can act as short-lived delivery storage, not as your permanent mailbox. Once a message has arrived, you can download the encrypted `.eml` file to one device, several devices, local disks, or your own backup system, and then delete it from our servers immediately.

Our approach is less convenient than a synchronized webmail account, but it changes the risk model: the long-term archive lives with you, while the provider has less mailbox data to protect, disclose, lose, or have stolen.


### One-Time Payments And 14-Day Payment-To-Account Link Deletion

Recurring billing keeps re-identifying the account because every renewal has to answer the same operational question again: which account is this payment for? A monthly card charge, yearly PayPal renewal, recurring SEPA transfer, or subscription management page needs either account-identifying metadata in the payment reference, a customer record inside the provider's billing system, or both. Even if the provider minimizes everything else, that recurring billing relationship becomes a long-lived identity bridge between payment data and mailbox account.

Privacy.Fish solves the payment-to-account-name association with a temporary payment code. During signup, the account request gets a random payment code; when the payment arrives with that code, we can create the requested mailbox without asking for a real name, recovery email, phone number, or permanent billing account. The code exists to match one incoming payment to one pending account creation, not to remain as a durable customer identifier.

For Norwegian consumer distance contracts, [Angrerettloven § 20](https://lovdata.no/lov/2014-06-20-27/%C2%A720) gives the consumer a withdrawal right, [§ 21](https://lovdata.no/lov/2014-06-20-27/%C2%A721) sets the ordinary withdrawal period for service contracts at 14 days from the day the agreement was entered into, and [§ 26](https://lovdata.no/lov/2014-06-20-27/%C2%A726) governs the consumer's obligations when withdrawal is used after service delivery has begun. During that period, a customer can cancel by contacting us and providing the payment code, because the code is what lets us find the payment-account match needed to process the refund. We keep that payment-code-to-account association for 14 days after activating the account, which may be slightly longer than the strict minimum in some cases; after that, we delete it.

From then on, we still have a list of received payments and a list of active accounts, but our systems can no longer answer which payment created which mailbox account. That deletion model is not possible with subscriptions, because the provider must keep enough billing state to charge, renew, fail, cancel, or refund the subscription later.




### The Real Cost Of Doing The Right Thing

The business cost is simple: subscriptions would probably make Privacy.Fish more profitable. We chose a one-time payment anyway because recurring billing would create exactly the kind of long-lived account-payment link we are trying to avoid in order to protect your privacy.

That means the service depends more on trust, word of mouth, and people who understand why this tradeoff matters. If Privacy.Fish is useful to you, recommending it to the right people helps more than any subscription upsell ever could.


## The Bottom Line: Jurisdiction Is Part Of Email Privacy

Email privacy is partly a technical problem, but it is also a jurisdiction problem: the country behind the provider decides what the provider must retain, disclose, or build for access. Germany, Switzerland, the United Kingdom, Australia, and the United States show how quickly "private email" can become a question of metadata retention, preservation duties, interception readiness, or recurring billing identity. Norway gives Privacy.Fish a narrower legal baseline, and we use the product design to keep the remaining data surface as small as we can. That choice is less convenient and less commercially comfortable than a normal mailbox subscription, but it is the only direction that matches what we mean by private email.


## FAQ: Norway Email Privacy, Data Retention, And Privacy.Fish

### Is Norway In The EU?

No. Norway is [not an EU Member State](https://www.eu-norway.org/eu/norway-and-the-eu/), although it participates in large parts of the European legal and economic framework through the EEA.

### What Does Norwegian Law Require Privacy.Fish To Log By Default?

For our SSH/SFTP account-access model, the routine retention duty under [Ekomloven § 3-13](https://lovdata.no/lov/2024-12-13-76/%C2%A73-13) maps to timestamp, username, source IP address and source port for registered account access, retained for 12 months.

### Does Norwegian Jurisdiction Mean No Lawful Access?

No. [Ekomloven § 3-12](https://lovdata.no/lov/2024-12-13-76/%C2%A73-12) requires providers to arrange services for legally mandated access, and [§ 3-14](https://lovdata.no/lov/2024-12-13-76/%C2%A73-14) governs disclosure of retained IP-address data to police or prosecution authorities in defined cases.

### Why Does Privacy.Fish Use One-Time Payments?

One-time payments let us delete the payment-code-to-account association after the 14-day window; subscriptions would require ongoing billing state that links a payment relationship to an account. The longer explanation is in our [privacy information](/documentation/privacy-information/).
