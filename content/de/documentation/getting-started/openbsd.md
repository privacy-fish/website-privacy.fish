---
title: "OpenBSD-Einrichtung"
description: "Grundlegende OpenBSD-Hinweise zum Erzeugen von SSH-Schlüsseln und zur Vorbereitung des privacy.fish-Zugangs."
weight: 7
video: "/videos/security-model.mp4"
---

OpenBSD passt gut zum privacy.fish-Workflow, weil SSH-Werkzeuge standardmäßig verfügbar sind.

## Schlüssel erzeugen

Erzeuge einen dedizierten SSH-Schlüssel für privacy.fish:

```sh
ssh-keygen -t ed25519 -f ~/.ssh/privacy-fish
```

Füge den Public Key aus `~/.ssh/privacy-fish.pub` in das Signup-Formular ein.

## Backup-Schlüssel behalten

Füge einen zweiten Schlüssel von einem anderen vertrauenswürdigen Gerät hinzu, wenn du eines hast. Wiederherstellung funktioniert nur über konfigurierte Backup-SSH-Schlüssel.

## Nach der Bereitstellung

Nachdem dein Konto erstellt wurde, nutze die Verbindungsdaten aus der Getting-Started-Anleitung und deine lokalen OpenBSD SSH / SFTP-Werkzeuge, um auf den Dienst zuzugreifen.
