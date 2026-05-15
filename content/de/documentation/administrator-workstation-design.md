---
title: "Design der Administrator-Workstation"
description: "Wie die Admin-Workstation als ersetzbare Infrastruktur behandelt und um minimales vertrauenswürdiges Material herum neu aufgebaut wird."
weight: 50
video: "/videos/security-model.mp4"
---

Das Kernkonzept hinter der Hardware einer Privacy.Fish-Administrator-Workstation ist, sie zu härten, sie aber trotzdem als ausnutzbar zu betrachten und deshalb regelmäßig vollständig zu ersetzen.

Der Administrator sollte nur vier Dinge benötigen, um eine neue Workstation einzurichten:

- stw.no Benutzername und Passwort
- stw.no 2FA-TOTP-Gerät
- SSH Private-Key-Datei für die Privacy.Fish-E-Mail-Server
- Passwort für den SSH Private Key

## Hardwarebeschaffung

Der sicherste Weg zu einer sicheren Workstation ist, sie häufig vollständig zu ersetzen. Dafür reicht ein Raspberry Pi mit einer SD-Karte mit vorinstalliertem Raspberry Pi OS oder Ähnlichem gut aus, weil diese Geräte günstig und leicht verfügbar sind. Peripherie wie Tastatur, Maus und Bildschirm kann ebenfalls nicht dauerhaft vertraut werden. Zusätzlich wird ein USB-Stick benötigt, um OpenBSD direkt vom Pi auf die SD-Karte zu installieren, bevor sie genutzt wird.

## Neue Workstation einrichten

Mit dem Raspberry Pi OS von der SD-Karte kann OpenBSD auf den USB-Stick geschrieben werden. Dieser USB-Stick kann danach verwendet werden, um OpenBSD auf die Micro-SD-Karte zu installieren.

## Neuen SSH Private Key erzeugen

Mit folgendem Befehl:

```bash
ssh-keygen -t ed25519 -a 100 -o
```

Für den SSH Private Key muss ein sicheres Passwort vergeben werden.

## SSH Private Key auf eine neue Workstation migrieren

Ein ed25519 Private Key ist relativ kurz und kann von einem anderen Bildschirm abgetippt oder auf Papier gedruckt werden. Es ist nicht nötig, einen USB-Stick oder ähnliches zu verwenden, nur um den SSH Private Key zu migrieren.
