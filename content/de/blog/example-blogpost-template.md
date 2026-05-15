---
title: "Beispiel-Blogpost-Vorlage"
description: "Entwurfsbeispiel mit den Inhaltsmustern, die für privacy.fish-Blogposts verfügbar sind."
date: 2026-05-04
draft: true
video: "/videos/security-model.mp4"
show_cta: false
---

Nutze diesen Entwurf als Referenz für zukünftige Blogposts. Er enthält normale Absätze, Bilder, Videos, Codeblöcke, Zitate, Tabellen, Listen und eine kleine Callout-Box.

## Abschnittsüberschrift

Schreibe kurze Absätze mit jeweils einem klaren Punkt. Das Bloglayout hält die Inhaltsbreite lesbar, daher sollten lange Erklärungen in Abschnitte aufgeteilt werden.

![privacy.fish sicherer Mail-Workflow](/images/features/secure-mail-workflow.png)

## Inline-21:9-Video

Nutze dieses Muster, wenn ein Beitrag ein weiteres Video im Body braucht:

<div class="not-prose my-10 aspect-[21/9] overflow-hidden rounded-[2rem] border border-deep/8 bg-mist shadow-[0_30px_80px_-58px_rgba(6,47,73,0.32)]">
  <video class="h-full w-full object-cover object-center" muted playsinline preload="auto" data-limited-video data-play-count="1" poster="/videos/posters/data-minimization-by-default.jpg">
    <source src="/videos/data-minimization-by-default.mp4" type="video/mp4">
  </video>
</div>

## Codeblock

```sh
ssh-keygen -t ed25519 -f ~/.ssh/privacy-fish
cat ~/.ssh/privacy-fish.pub
```

## Zitat

> Jedes gespeicherte Geheimnis schafft ein weiteres operatives Risiko.

## Checkliste

- Erkläre den Tradeoff in klarer Sprache.
- Zeige exakt, was sich für den Benutzer ändert.
- Halte Aussagen spezifisch genug, um sie prüfen zu können.

## Vergleichstabelle

| Thema | Konventionelle E-Mail | privacy.fish |
| --- | --- | --- |
| Login | Passwortbasiertes Webmail | SSH-schlüsselbasierter Workflow |
| Speicherung | Langfristige Mailbox-Verwahrung | Kurze serverseitige Speicherung |
| Wiederherstellung | Identitäts- und Supportprozess | Backup-SSH-Schlüssel |

## Callout

<div class="not-prose my-10 rounded-[2rem] border border-ocean/16 bg-mist/55 p-6">
  <p class="text-sm font-bold tracking-[0.18em] text-amber uppercase">Redaktioneller Hinweis</p>
  <p class="mt-3 text-lg leading-8 text-deep/76">Nutze Callouts sparsam. Sie funktionieren am besten für Warnungen, Datenschutzhinweise oder Implementierungsdetails, die nicht im Haupttext untergehen sollen.</p>
</div>

## Schlussabschnitt

Ende damit, die konkrete Erkenntnis erneut zu formulieren. Vermeide generische Datenschutzsprache, außer der Beitrag hat bereits erklärt, was das System tatsächlich anders macht.
