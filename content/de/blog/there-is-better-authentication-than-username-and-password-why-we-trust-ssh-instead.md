---
title: "Es gibt bessere Authentifizierung als Benutzername und Passwort: Warum wir SSH vertrauen"
description: "Warum Privacy.Fish SSH-Schlüsselpaare statt Benutzername-und-Passwort-Login für E-Mail-Zugang nutzt."
date: 2026-05-06
draft: false
video: "/videos/blog/there-is-better-authentication-than-username-and-password-why-we-trust-ssh-instead/ssh-as-sign-in.mp4"
show_cta: false
---

## Das Problem mit klassischen Passwörtern

Benutzername und Passwort sind bequem, aber operativ riskant. Passwörter werden wiederverwendet, können gephisht werden, können schwach sein und müssen serverseitig in irgendeiner Form überprüfbar bleiben.

Für Privacy.Fish ist das zu viel Angriffsfläche.

## Wie SSH-Schlüssel funktionieren

Der Nutzer erzeugt ein Schlüsselpaar: einen privaten Schlüssel und einen passenden Public Key. Der Public Key wird bei der Anmeldung hinterlegt. Der private Schlüssel bleibt auf dem Gerät des Nutzers.

Beim Login beweist der Client kryptografisch, dass er den privaten Schlüssel besitzt. Der Server muss den privaten Schlüssel nie sehen und kein wiederverwendbares Postfachpasswort speichern.

## Warum das zu age-Verschlüsselung passt

Kompatible SSH Public Keys können auch als Empfänger für age-verschlüsselte Dateien genutzt werden. Damit passt dasselbe Schlüsselmodell sowohl für Authentifizierung als auch für verschlüsselte E-Mail-Speicherung.

## Warum wir den Usability-Preis akzeptieren

SSH-Schlüssel sind technischer als Passwörter. Nutzer müssen Schlüssel erzeugen, private Schlüssel schützen und Backups planen. Privacy.Fish akzeptiert diesen Aufwand, weil er uns erlaubt, weniger Daten zu sammeln, weniger Geheimnisse zu speichern und weniger Login-Oberflächen zu betreiben.

Das ist nicht der richtige Kompromiss für alle. Für Privacy.Fish ist es der richtige Kompromiss.
