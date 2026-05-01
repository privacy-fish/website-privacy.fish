---
title: "Sign up"
description: "Pick a username and add your SSH public keys."
layout: "signup"
sitemap:
  changefreq: monthly
  priority: 0.8
blocks:
  - block: signup
    id: signup
    eyebrow: "Sign up"
    heading: "Reserve your privacy.fish address"
    subheading: |-
      Pick your username and add your SSH public keys.  
      We will verify them and walk you through the rest.
    human_step_label: "Human check first"
    human_heading: "Prove you are human before checking usernames."
    human_copy: "Solve the captcha once, then you get five username or signup attempts."
    captcha_label: "Captcha"
    captcha_disabled_copy: "Captcha is required before availability checks. The local captcha widget will appear here once it is enabled."
    username_step_label: "Choose your username"
    username_label: "Desired email address"
    username_help: "Lowercase letters, digits, dots and hyphens. 1 to 31 characters. Must start and end with a letter or digit."
    username_placeholder: "john.doe"
    username_domain: "@privacy.fish"
    check_text: "Check availability"
    username_alias_help: 'You will also receive matching addresses on every other supported domain (for example <code class="font-mono text-ocean">john.doe@pfi.sh</code>).'
    keys_step_label: "Add your SSH public keys"
    keys_label: "Up to 10 ed25519 keys, one per line"
    keys_help: 'Only <code class="font-mono text-ocean">ssh-ed25519</code> keys are accepted. Each key authorizes one device. Mail is age-encrypted to these keys, so only you can read it.'
    keys_placeholder: |-
      ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI... laptop
      ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI... phone
    keys_generate_help: 'Generate with <code class="font-mono text-ocean">ssh-keygen -t ed25519</code> on every device.'
    keys_comment_help: 'The optional comment after each key (the third part, often <code class="font-mono text-ocean">user@host</code>) is removed before submission to protect your privacy.'
    keys_counter_initial: "0 / 10"
    submit_text: "Reserve account for 30 days"
    help_url: "/docs/getting-started/"
    help_text: "Read the Getting Started guide"
    success_eyebrow: "Reserved"
    success_heading: "What happens now"
    success_copy: "Your account is reserved for 30 days. Include the payment reference below whenever or however you pay. We manually create accounts within 24 hours after receiving payment."
    payment_reference_label: "Payment reference"
    success_cards:
      - title: "Pay € 20"
        text: "Use this reference with bank transfer, crypto, cash by mail, or voucher payment."
      - title: "Manual provisioning"
        text: "We create your account after payment is received and verified."
      - title: "Association deleted"
        text: "The payment reference to account link is deleted after provisioning."
    new_signup_text: "Reserve another account"
---
