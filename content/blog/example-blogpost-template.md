---
title: "Example blog post template"
description: "Draft example showing the content patterns available for privacy.fish blog posts."
date: 2026-05-04
draft: false
video: "/videos/security-model.mp4"
show_cta: false
---

Use this draft as a reference for future blog posts. It includes normal paragraphs, images, videos, code blocks, quotes, tables, lists, and a small callout block.

## Section heading

Write short paragraphs with one clear point each. The blog layout keeps the content width readable, so long explanations should be split into sections.

![privacy.fish secure mail workflow](/images/features/secure-mail-workflow.png)

## Inline 21:9 video

Use this pattern when a post needs another video inside the body:

<div class="not-prose my-10 aspect-[21/9] overflow-hidden rounded-[2rem] border border-deep/8 bg-mist shadow-[0_30px_80px_-58px_rgba(6,47,73,0.32)]">
  <video class="h-full w-full object-cover object-center" muted playsinline preload="auto" data-limited-video data-play-count="1" poster="/videos/posters/data-minimization-by-default.jpg">
    <source src="/videos/data-minimization-by-default.mp4" type="video/mp4">
  </video>
</div>

## Code block

```sh
ssh-keygen -t ed25519 -f ~/.ssh/privacy-fish
cat ~/.ssh/privacy-fish.pub
```

## Quote

> Every retained secret creates another operational risk.

## Checklist

- Explain the tradeoff in plain language.
- Show exactly what changes for the user.
- Keep claims specific enough to verify.

## Comparison table

| Topic | Conventional email | privacy.fish |
| --- | --- | --- |
| Login | Password-based webmail | SSH-key based workflow |
| Storage | Long-term mailbox custody | Short server-side retention |
| Recovery | Identity and support process | Backup SSH keys |

## Callout

<div class="not-prose my-10 rounded-[2rem] border border-ocean/16 bg-mist/55 p-6">
  <p class="text-sm font-bold tracking-[0.18em] text-amber uppercase">Editorial note</p>
  <p class="mt-3 text-lg leading-8 text-deep/76">Use callouts sparingly. They work best for warnings, privacy notes, or implementation details that should not be buried in the main text.</p>
</div>

## Final section

End by restating the concrete takeaway. Avoid generic privacy language unless the post has already explained what the system actually does differently.
