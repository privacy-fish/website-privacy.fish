---
title: "Design der Administrator-Workstation"
description: "Wie die Admin-Workstation als ersetzbare Infrastruktur behandelt und um minimales vertrauenswürdiges Material herum neu aufgebaut wird."
weight: 50
video: "/videos/security-model.mp4"
---

Das Grundkonzept der Privacy.Fish-Admin-Workstation ist, sie zu härten, aber dennoch als potenziell kompromittierbar zu betrachten und regelmäßig vollständig zu ersetzen.

Eine neue Workstation sollte nur wenige Dinge benötigen:

- stw.no Benutzername und Passwort
- stw.no 2FA-TOTP-Gerät
- SSH Private Key für die Privacy.Fish-Mailserver
- Passwort für den SSH Private Key

## Hardwarebeschaffung

Eine günstige, leicht ersetzbare Plattform wie ein Raspberry Pi ist für dieses Modell geeignet. Auch Peripheriegeräte werden nicht dauerhaft als vertrauenswürdig betrachtet.

## Neue Workstation einrichten

OpenBSD wird frisch installiert. Nur das minimal notwendige vertrauenswürdige Material wird übernommen.

## SSH Private Key migrieren

Ein ed25519 Private Key ist kurz genug, um kontrolliert von einem anderen Bildschirm abzutippen oder auf Papier zu übertragen. Für nur diesen Schlüssel ist kein USB-Stick erforderlich.
