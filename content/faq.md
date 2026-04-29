---
title: "FAQ"
type: "faq"
blocks:
  - block: faq
    id: faq
    heading: "FAQ"
    subheading: "Short answers before you decide if privacy.fish fits how you want to handle email."
    layout: "stacked"
    items:
      - question: "Is privacy.fish for everyone?"
        answer: "No. privacy.fish is designed for people who prioritize privacy and security above convenience, and it fits best if you are comfortable with a more technical workflow."
      - question: "Do I need to be technical to use privacy.fish?"
        answer: "You do not need to be an expert, but privacy.fish is built for people comfortable with command-line, SSH, and SFTP workflows."
      - question: "How do I access my email?"
        answer: "You download age-encrypted <code class=\"font-mono text-ocean\">.eml</code> files over SFTP. An optional client-side tool can serve downloaded mail to an email client such as Thunderbird via POP3."
      - question: "Does privacy.fish include webmail?"
        answer: "No. The service is designed around downloading encrypted mail to your own devices instead of keeping a traditional webmail mailbox on the server."
      - question: "How many devices can I use?"
        answer: "Each account supports up to 10 SSH public keys, which means you can authorize up to 10 devices."
      - question: "Can I create more than one email address?"
        answer: "Yes. Each account includes one main address and an unlimited number of aliases, including supported privacy.fish domains such as <code class=\"font-mono text-ocean\">@pfi.sh</code>."
      - question: "How long is email kept on the server?"
        answer: "Emails are downloaded by you or automatically deleted after 14 days. privacy.fish is intentionally not permanent server-side storage."
      - question: "Is privacy.fish a subscription?"
        answer: "No. Access is unlocked with a one-time payment for the lifetime of the service rather than an ongoing subscription."
      - question: "What happens if I lose my SSH keys?"
        answer: "Recovery depends on your backup SSH key pair. Without a valid backup key, there is no alternative recovery path."
      - question: "Do you scan emails or track website visitors?"
        answer: "No. privacy.fish does not use third-party analytics or tracking scripts on the website, and it does not scan your emails beyond minimal anti-spam protections."
      - question: "Why is iPhone support listed as unavailable?"
        answer: "iPhone is not supported because the required privacy.fish client tooling cannot run there in the way the service requires."
---
