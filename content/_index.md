---
title: "privacy.fish"
description: "A privacy-first email service built to minimize retained data, attack surface, and trust."
logo: "images/logo.svg"
blocks:
  - block: hero
    id: hero
    eyebrow: "Hosted in Norway"
    heading: "Email: Uncompromisingly Private and Secure"
    prefix_highlight: "Email:"
    highlight: "Private and Secure"
    subheading: "Private email hosting built around data minimization and open-source infrastructure."
    summary_heading: "Securely Designed to Protect Privacy"
    primary_cta_text: "Create Your Account"
    primary_cta_url: "https://privacy.fish/signup"
    secondary_cta_text: "Code"
    secondary_cta_url: "https://github.com/privacy-fish"
    callouts:
      - value: "Privacy by Design"
        label: "Only data required to run the email servers is collected, and any data no longer needed is deleted."
      - value: "Verifiable Security"
        label: "All our code is open source and can be audited by the community and external cybersecurity firms."
    stats:
      - value: "10 Devices"
        label: "Per account"
      - value: "0 Compromise"
        label: "On what matters"
    price:
      value: "20,00 € One time payment"
      label: "One registration fee. No subscriptions and no recurring costs."
    compat:
      id: compat
      label: "Compatible with"
      unsupported_prefix: "No"
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
    eyebrow: "Norwegian Privacy"
    heading: "<span class=\"text-sun\">Private and Secure</span> <span class=\"text-deep\">Access to the Internet</span> <span class=\"text-ocean\">Is a Human Right</span>"
    lead: "An email address is a requirement for participating in the modern internet. It is the first thing most modern services ask for. Its privacy should be protected."
    paragraphs:
      - "Your email address is the key to most of your online life. When your email provider tracks it, your identity becomes easier to profile, target, and expose."
      - "Privacy.Fish treats email as private infrastructure, not as a data source."
      - "We collect only what is required to run the mail servers, publish our code, and design every technical decision around privacy and security."
    primary_cta_text: "Read our company philosophy page"
    primary_cta_url: "/documentation/company-philosophy/"

  - block: counter
    id: counter
    items:
      - value: "∞"
        label: "Privacy Aliases"
      - value: "20 €"
        label: "One Time Payment"
      - value: "10"
        label: "Devices"
      - value: "0"
        label: "Compromise on Privacy and Security"

  - block: features
    id: features
    eyebrow: "What we do differently"
    heading: "What separates us from all others"
    subheading: "Every design decision puts your privacy first — from encryption at rest to how we handle payments and logs."
    groups:

      - title: "Data Minimization by Default"
        video: "/videos/data-minimization-by-default.mp4"
        image_position: "left"
        items:
          - "We do not want data we don't need and prefer deleting data over trying to protect it."
          - "No analytics, trackers, website-, mail- or ssh auth logs."
          - "You download and delete mail from our servers yourself. Mail you do not delete is deleted after 14 days."

      - title: "Private Payment"
        video: "/videos/private-payment.mp4"
        video_aspect: "16:9"
        image_position: "right"
        items:
          - "After signup, you receive a temporary payment code valid for 30 days."
          - "Once payment is received and the account is created, the code is deleted so the payment cannot be linked to the account anymore."
          - "We support cash by letter, cryptocurrencies, SEPA bank transfer, PayPal and credit cards."

      - title: "Norway’s Privacy Jurisdiction"
        video: "/videos/norways-privacy-jurisdiction.mp4"
        image_position: "left"
        items:
          - "Norwegian law only requires us to save when you logged in and from with IP:port, for 12 months."
          - "All other jurisdictions come with broader surveillance requirements. For example Switzerland can require sender, recipient, protocol, mailbox-event, and server metadata, while Germany requires large email providers to maintain lawful-interception infrastructure."
          - "VPN access is encouraged and tor .onion addresses are available for all servers."

      - title: "Security Model"
        video: "/videos/security-model.mp4"
        image_position: "right"
        items:
          - "Built using the most secure open source software (OpenBSD, OpenSMTPD, OpenSSH) and only minimal custom code. Everything is auditable open source code."
          - "To prevent patched exploits from leaving persistent access, all servers are rebuilt weekly, while only stored emails and your SSH public keys are migrated."
          - "Admin workstations are OpenBSD Raspberry Pis, replaced monthly with only the SSH private key migrated, and firewalled to only reach our servers and stw.no."

      - title: "Secure Email Workflow"
        video: "/videos/secure-email-workflow.mp4"
        image_position: "left"
        items:
          - "Security before convenience - There is no webmail, password login, IMAP, or POP3 - only SSH and SFTP access through our client app."
          - "Your emails are stored age-encrypted using your SSH public keys and are securely deleted after 14 days or when you delete them yourself."
          - "If a recipient’s mail server cannot prove its identity with a valid TLS certificate, the app asks you to cancel or send anyway."

  - block: audience
    id: audience
    eyebrow: "Should you use it?"
    heading: "Privacy Requires Making a Choice"
    for_title: "Built for"
    for_items:
      - "You are privacy-conscious and want an email provider that collects as little data as possible."
      - "You do not want your email or personal data permanently stored on someone else’s servers."
      - "You prefer a small, focused service built exclusively for the best possible privacy and security."
    not_for_title: "Not built for"
    not_for_items:
      - "You need your email provider to permanently store mail instead of downloading it to your own devices."
      - "You need browser email access, even if that means exposing even more metadata or readable mail to the provider."
      - "You need bundled extras like calendar, contacts, cloud storage, or office tools, even if they add attackable code."

  - block: how-it-works
    id: how-it-works
    eyebrow: "How it works"
    heading: "From Signup to Inbox"
    subheading: "After the legal consumer refund period of 14 days has expired, we will permanently delete all information on which account is associated with which received payment."
    steps:
      - icon: "key"
        title: "Use the App to Generate Keypairs on Your Devices"
        text: "Install the open source Privacy.Fish app and generate SSH keypairs for each device. Copy the public keys to the device you complete the signup form with."

      - icon: "credit-card"
        title: "Signup and Pay Privately"
        text: "Copy your list of public keys into the signup form and use our privacy-respecting payment methods. We create your account within 24 hours after payment."

      - icon: "desktop"
        title: "Connect the App to Your Email Client"
        text: "Point your email app to this app’s local SMTP and POP3 ports, and it will handle sending, fetching, and decrypting your mail through our servers."


  - block: tradeoffs # TODO rename me to trust model
    id: tradeoffs
    eyebrow: "Our Trust Model"
    heading: "You Verify Everything"
    subheading: "“Every secret creates a potential failure point.” — [Bruce Schneier](https://en.wikipedia.org/wiki/Bruce_Schneier)"
    items:
      - icon: "code"
        title: "Everything Open Source"
        text: "All software used to build and administrate Privacy.Fish is public on [github](https://github.com/privacy-fish) and our infrastructure can be inspected instead of blindly trusted."
      - icon: "scale-balanced"
        title: "Norwegian Jurisdiction"
        text: "We chose Norway for the best privacy-respecting laws for our service, even though they have rather high taxes. Bcause privacy comes first."
      - icon: "coins"
        title: "Aligned Incentives"
        text: "Privacy.Fish is funded by a one-time account fee, not ads, tracking, subscriptions, or selling access to your data."

  - block: pricing
    id: pricing
    eyebrow: "Pricing"
    heading: "One Time Payment"
    subheading: "No subscriptions or hidden fees. Payment to account association is deleted after account creation."
    price: "20"
    badge: "One-time payment"
    note: "For company lifetime access"
    cta_text: "Create Your Account"
    cta_url: "https://privacy.fish/signup"
    provisioning_title: "Manual provisioning"
    provisioning_copy: "Accounts are created by an administrator within 24 hours of payment."
    included_heading: "What’s included"
    features:
      - "Main username, 10 random and unlimited rotating aliases"
      - "All our domains work with your main username and all aliases"
      - "Maximum of 10 devices per account"
    payment_label: "Payment methods:"
    payment_methods:
      - "Cash by letter"
      - "Cryptocurrencies"
      - "SEPA bank transfer"
      - "Credit card"
      - "PayPal"
    refund_label: "Refund information:"
    refund_information: "Refunds are only possible before the account is created, because the payment-to-account association is destroyed right after."

---
