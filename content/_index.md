---
title: "privacy.fish"
description: "A privacy-first email service built to minimize retained data, attack surface, and trust."
logo: "images/logo.svg"
blocks:
  - block: hero
    id: hero
    eyebrow: "Hosted in Norway"
    heading: "Uncompromisingly private and secure"
    highlight: "private and secure"
    subheading: "Private and secure access to the internet is a human right."
    summary_heading: "Designed solely for your Privacy and Security"
    primary_cta_text: "Create Your Account"
    primary_cta_url: "https://privacy.fish/signup"
    secondary_cta_text: "How it works"
    secondary_cta_url: "#how-it-works"
    callouts:
      - value: "No compromise privacy"
        label: "We prioritize customer privacy above profit, usability, and everything else."
      - value: "No compromise security"
        label: "We prioritize IT security above usability, features, and everything else."
    stats:
      - value: "10 devices"
        label: "Per account"
      - value: "0 compromise"
        label: "On what matters to you"
    price:
      value: "€ 20"
      label: "Lifetime account"

  - block: compat
    id: compat
    label: "Compatible with"
    items:
      - name: "Command Line"
        supported: true
      - name: "Linux"
        supported: true
      - name: "macOS"
        supported: true
      - name: "Windows"
        supported: true
      - name: "Android"
        supported: true
      - name: "iPhone support"
        supported: false

  - block: intro
    id: intro
    eyebrow: "Why we exist"
    heading: "Email, the way it should be"
    lead: "privacy.fish exists for people who would rather adapt their workflow than hand over more data than necessary."
    paragraphs:
      - "A free and open society is built upon privacy. If people don't have the right to decide for themselves exactly when and with whom they want to share their thoughts and ideas, then all other rights are at risk."
      - "Every human being has the right to access the internet. An email address is required to sign up for almost all modern online services."
      - "<strong>privacy.fish</strong> email is focused on privacy and security above absolutely everything else. Zero compromise."
    primary_cta_text: "Create Your Account"
    primary_cta_url: "https://privacy.fish/signup"

  - block: counter
    id: counter
    items:
      - value: "€ 20"
        label: "per Account"
      - value: "10"
        label: "devices per Account"
      - value: "∞"
        label: "email addresses per account"
      - value: "0"
        label: "compromise on privacy and security"

  - block: features
    id: features
    eyebrow: "What we do differently"
    heading: "What separates us from all others"
    subheading: "Every design decision puts your privacy first — from encryption at rest to how we handle payments and logs."
    groups:
      - title: "Data minimization by default"
        image: "images/features/we-keep-only-what-is-technically-required.png"
        image_position: "left"
        items:
          - "We prefer deleting data over trying to protect it."
          - "We really do not want any of your personal data."
          - "No analytics, trackers, website-, mail- or ssh auth logs."
          - "You download and delete mail from our servers yourself."
          - "Mail you do not delete is deleted after 14 days."

      - title: "Private account and payment flow"
        image: "images/features/private-account-payment-flow.png"
        image_position: "right"
        items:
          - "The payment-to-account reference is deleted after provisioning."
          - "Cryptocurrencies and cash-by-mail payments are supported."
          - "Recovery only works using backup SSH keys."

      - title: "Private access under Norwegian jurisdiction"
        image: "images/features/private-access-norwegian-jurisdiction.png"
        image_position: "left"
        items:
          - "Tor .onion addresses are available for all servers."
          - "Norway is the world's best jurisdiction for a private email provider."

      - title: "Secure mail workflow"
        image: "images/features/secure-mail-workflow.png"
        image_position: "right"
        items:
          - "No webmail, no password login, no user comfort in favor of less security."
          - "Mail is age-encrypted using your SSH public keys."
          - "When you delete mail on our servers, it is overwritten with random data."
          - "Sending and receiving mail happens through SSH / SFTP."

      - title: "Drastically reduced attack surface"
        image: "images/features/drastically-reduced-attack-surface.png"
        image_position: "left"
        items:
          - "Built on OpenBSD, OpenSSH, and OpenSMTPD, with minimal custom code."
          - "All infrastructure is FOSS code, including the admin workstation setup."
          - "Signup is separated from the mail servers; paid accounts are manually provisioned."
          - "Spam and virus filtering run locally on our servers, without network access."
          - "Sending to a mail server with an invalid TLS certificate requires you to confirm or abort."
          - "Servers and admin workstations are rebuilt weekly; only emails, SSH keys, and Tor onion keys migrate."

  - block: audience
    id: audience
    eyebrow: "The fit"
    heading: "Is privacy.fish for you?"
    for_title: "Who is it for?"
    for_items:
      - "People who prioritize privacy over convenience"
      - "Technical users comfortable using SSH / SFTP"
      - "Journalists, activists, researchers, and privacy-conscious individuals"
      - "Users who want minimal retained data"
    not_for_title: "Who it is not for"
    not_for_items:
      - "Users who want classic webmail"
      - "Users who need permanent server-side storage"
      - "Users who want easy account recovery"
      - "Users who want polished convenience features over security guarantees"

  - block: how-it-works
    id: how-it-works
    eyebrow: "How it works"
    heading: "Private email, simply explained"
    subheading: "privacy.fish is fully command-line compatible. Client-side tools to serve downloaded mail via POP3 exist as FOSS code."
    steps:
      - icon: "key"
        title: "Up to 10 SSH keys per account"
        text: "Authorize one SSH public key per device — phone, laptop, workstation. Mail is age-encrypted to those keys, so only your devices can read it."
      - icon: "at"
        title: "One main address, unlimited aliases"
        text: "You get one main and an unlimited number of @privacy.fish email addresses — for example <code class=\"font-mono text-ocean\">john.doe@privacy.fish</code>."
      - icon: "globe"
        title: "Multiple top-level domains"
        text: "We support a variety of additional top-level domains, which you get automatically — like <code class=\"font-mono text-ocean\">john.doe@pfi.sh</code>."
      - icon: "cloud-arrow-down"
        title: "SFTP download of encrypted .eml files"
        text: "Download age-encrypted .eml files via SFTP. An optional client-side tool serves them to your email client (like Thunderbird) via POP3."
      - icon: "trash-can"
        title: "Manual or automatic deletion"
        text: "You can manually delete emails from the server via SFTP, or they are automatically deleted after 14 days."
      - icon: "desktop"
        title: "No webmail — you own your mail"
        text: "There is no traditional webmail client. You download your emails to one or more devices."

  - block: tradeoffs
    id: tradeoffs
    eyebrow: "Honest about the tradeoffs"
    heading: "What are the tradeoffs?"
    subheading: "We prioritize privacy and security above convenience. privacy.fish is a technical product. While it does not require deep expertise, it helps to understand how it works."
    items:
      - icon: "user-gear"
        title: "User experience"
        text: "Usability is more complicated than traditional email providers."
      - icon: "mobile-screen"
        title: "Platform compatibility"
        text: "iPhone is not supported, since there is no way to run privacy.fish client tooling there."
      - icon: "coins"
        title: "Profit"
        text: "A one-time payment keeps incentives aligned with your privacy, not perpetual monetization."

  - block: threat
    id: threat-model
    eyebrow: "Threat model"
    heading: "Security philosophy"
    subheading: "We are explicit about what we protect against, what we deliberately avoid, and what we cannot prevent."
    items:
      - label: "What we defend against"
        title: "No Privacy violations driven by profit"
        text: "And avoidable vulnerabilities in the email hosting stack."
      - label: "What we deliberately avoid"
        title: "No Collecting data we don't need"
        text: "Any personal data that is not required to operate the service is never collected."
      - label: "What our design minimizes"
        title: "Minimized Retained data, trust and attack surface"
        text: "We minimize data kept on our servers, the trust you have to place in us, and the attack surface of our infrastructure."
      - label: "What we cannot protect against"
        title: "Only Lawful Norwegian court orders"
        text: "And other legally binding requests from Norwegian authorities."

  - block: cta
    id: last-word
    eyebrow: "Last word"
    heading: "Private and secure access to the internet is a human right!"
    copy: "Join privacy.fish and take back control of your email."
    primary_cta_text: "Code"
    primary_cta_url: "https://github.com/fishprivacy"
    secondary_cta_text: "Read the docs"
    secondary_cta_url: "/docs/"

  - block: pricing
    id: pricing
    eyebrow: "Pricing"
    heading: "One payment. For life."
    subheading: "No subscriptions, no hidden fees, no data harvesting. You pay once, then the account is yours."
    price: "20"
    note: "One-time — for the lifetime of the service"
    cta_text: "Create Your Account"
    cta_url: "https://privacy.fish/signup"
    features:
      - "One account with unlimited email addresses"
      - "Includes all supported top-level domains"
      - "Account manually created by our admins within 24 hours of payment"
      - "Refunds possible within one week (except cash and cryptocurrencies)"
    payment_label: "Payment methods:"
    payment_methods:
      - "Credit card"
      - "PayPal"
      - "Cash"
      - "Voucher"
      - "Cryptocurrencies"
      - "SEPA bank transfer"

---
