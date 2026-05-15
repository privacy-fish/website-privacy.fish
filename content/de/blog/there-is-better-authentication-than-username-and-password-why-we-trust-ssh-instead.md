---
title: "Es gibt bessere Authentifizierung als Benutzername und Passwort: Warum wir stattdessen SSH vertrauen"
description: "Warum Privacy.Fish SSH-Schlüsselpaare statt Benutzername-und-Passwort-Login für E-Mail verwendet."
date: 2026-05-06
draft: false
video: "/videos/blog/there-is-better-authentication-than-username-and-password-why-we-trust-ssh-instead/ssh-as-sign-in.mp4"
show_cta: false
---

## Die Probleme klassischer Benutzername-und-Passwort-Authentifizierung

Benutzername + Passwort ist die Standardwahl für die meisten Webdienste, weil jeder sie versteht, jedes Framework sie unterstützt und sie zu normalen Login-, Reset- und Support-Abläufen passt.

Der Preis ist ein vorhersehbares operatives Sicherheitsproblem: Benutzer verwenden Passwörter mehrfach, Phishing-Seiten können sie einsammeln, schwache Passwörter können geraten werden, und Benutzername/Passwort-Paare aus einem Leak können automatisch gegen andere Dienste ausprobiert werden (Credential Stuffing). Der Dienst muss außerdem Passwort-Verifikationsdaten pflegen, die bei einem Datenleck wertvolles Material werden.

Außerdem bietet dieses Modell schlechte Gerätetrennung. Wenn dasselbe Passwort in jeden Client eingegeben und dort gespeichert wird, wo diese Clients es speichern wollen, kann die Kompromittierung eines Geräts, App-Profils oder kopierten Credentials zur Kompromittierung des gesamten Kontos werden.

## Wie SSH-Schlüsselpaar-Authentifizierung funktioniert

SSH-Schlüsselpaar-Authentifizierung funktioniert anders als Benutzername + Passwort. Der Benutzer erzeugt ein Schlüsselpaar: einen Private Key und den dazugehörigen Public Key. Der Public Key wird bei der Anmeldung angegeben und auf unsere E-Mail-Server kopiert. Der Private Key bleibt beim Benutzer. Wenn der Benutzer verbindet, fragt der Server nicht nach dem Private Key und der Client lädt ihn nicht hoch. Stattdessen führt OpenSSH einen kryptografischen Challenge-Response-Austausch aus: Der Client signiert Authentifizierungsdaten mit dem Private Key, und der Server prüft diese Signatur mit dem Public Key, den er bereits hat.

Wir können noch mehr tun: Es trifft sich gut, dass kompatible SSH Public Keys auch als Empfänger für age-verschlüsselte Dateien verwendet werden können. So können wir dasselbe Schlüsselpaar-Modell auch für gespeicherte E-Mail-Verschlüsselung nutzen.

### Warum SSH Public Keys öffentlich sein können

Ein SSH Public Key kann öffentlich ins Internet gestellt werden, weil er anderen nur erlaubt, Signaturen des zugehörigen Private Keys zu prüfen; er kann diese Signaturen nicht selbst erzeugen. Ein Angreifer, der die SSH Public Keys unserer Benutzer kompromittiert, erhält dadurch also nicht die Fähigkeit, sich in deren Konten einzuloggen.

Selbst eine gut konfigurierte Passwort-Hash-Datenbank kann durch Offline-Testen von Passwortvermutungen angegriffen werden; ein SSH Public Key gibt dem Angreifer keinen vergleichbaren Weg, den passenden Private Key abzuleiten. Der sensitive Teil ist der Private Key, und der sollte auf dem Gerät des Benutzers, einem Hardware-Token, einem verschlüsselten Backup oder Offline-Speicher bleiben.

### Die Nutzererfahrung von SSH-Schlüsselpaar-Authentifizierung

Für Privacy.Fish-Benutzer bedeutet das, dass Kontozugriff mit der Verwaltung von Schlüsselpaaren beginnt statt mit dem Merken eines Mailbox-Passworts.

Jeder unserer Benutzer kann zwei Account Control Public Keys und bis zu acht Email Public Keys haben.

Die acht Email Public Keys können verwendet werden, um E-Mails von unseren Servern zu senden, herunterzuladen und zu löschen.

Die zwei Account Control Public Keys können Account Control Public Keys und Email Public Keys hinzufügen oder entfernen und außerdem alles tun, was Email Public Keys tun können. Wir empfehlen, sie sicher und offline aufzubewahren.

Benutzer, die es einfach halten und nur ein Gerät haben wollen, können sich mit nur einem Schlüsselpaar anmelden und den Dienst nutzen, wobei dieses als Account Control Public Key verwendet wird.

Das ist für viele E-Mail-Benutzer ungewohnt. Menschen verstehen Passwörter intuitiv, weil jede Website sie verwendet. Schlüsselpaare sind technischer. Wir haben SSH-Schlüsselpaar-Authentifizierung gewählt, weil Privacy.Fish Datenschutz und Sicherheit über Bequemlichkeit stellt.

## Warum wir OpenSSH als Authentifizierungssystem für unseren E-Mail-Dienst vertrauen

Das ist nicht nur Faulheit, auch wenn OpenSSH standardmäßig auf unseren OpenBSD-basierten Servern installiert ist. Privacy.Fish nutzt OpenSSH für Benutzer­authentifizierung, weil es uns einen reifen, stark getesteten Authentifizierungsstack bietet.

### OpenSSH ist weit verbreitet und stark auditiert

[OpenSSH](https://www.openssh.com/ "OpenSSH") ist einer der meistverbreiteten Remote-Access-Stacks der Welt. Es wurde über Jahrzehnte angegriffen, auditiert, gepatcht, gehärtet und operativ missbraucht. Das macht es nicht perfekt. Es bedeutet aber, dass Vertrauen in OpenSSH eine völlig andere Risikoentscheidung ist als Vertrauen in ein kleines eigenes Login-Protokoll oder eine Webmail-spezifische Authentifizierungsschicht.

Es gibt noch einen wichtigen Unterschied: TLS um IMAP, POP3 oder SMTP AUTH schützt den Transport, aber die Anwendung endet oft weiterhin in einem passwortartigen Login. SSH kombiniert verschlüsselten Transport und Public-Key-Authentifizierung in einem reifen Protokoll. Wir halten „SSH gut, TLS schlecht“ nicht für das richtige Denkmodell. Das eigentliche Problem ist, dass passwortbasierte E-Mail-Protokolle meistens mehr exponierte Authentifizierungsflächen erfordern, als wir betreiben wollen.

### Ein Authentifizierungsstack ist besser als zwei

Wenn Privacy.Fish IMAP zum Lesen, plus POP3 zum Herunterladen und Löschen, plus SMTP AUTH zum Senden, plus Webmail für Browserzugriff betreiben würde, müsste jede dieser Flächen Benutzer­authentifizierung akzeptieren. Jede Fläche bräuchte Rate Limiting, Monitoring, Härtung, Logs, Patching, Missbrauchsbehandlung und Konfigurationsprüfung. Selbst wenn jede Komponente gut gewählte Open-Source-Software wäre, wäre die gesamte Angriffsfläche größer.

Also halten wir das Benutzer­authentifizierungsmodell auf SSH zentriert. Das ist keine normale E-Mail-UX, aber ein deutlich kleineres und saubereres Design.

### Dasselbe Public-Key-Modell passt zu age-Verschlüsselung für gespeicherte E-Mails

Dasselbe Public-Key-Modell passt auch dazu, wie Privacy.Fish E-Mail verschlüsselt speichert. Für jede eingehende Nachricht schreibt Privacy.Fish eine age-verschlüsselte Datei, deren Empfänger die kompatiblen SSH Public Keys des Benutzers sind. Einer der passenden Private Keys wird benötigt, um diese `.age`-Datei zu entschlüsseln, sodass der Server verschlüsselten Zustellspeicher halten kann.

[Age](https://github.com/FiloSottile/age "Age") bietet außerdem authentifizierte Verschlüsselung. Wenn eine verschlüsselte E-Mail-Datei verändert wird, schlägt die Entschlüsselung fehl, statt still geänderten Klartext zu erzeugen. Das löst nicht jedes mögliche serverseitige Manipulationsproblem, gibt uns aber die wichtigste Grundeigenschaft: Gespeicherte E-Mail sollte ohne den Private Key des Benutzers nicht lesbar sein.

## Warum wir den Usability-Preis akzeptieren

SSH-Schlüsselpaar-Authentifizierung verlangt Benutzern mehr ab, aber sie erlaubt Privacy.Fish, weniger zu verlangen und weniger zu speichern.

### SSH-Schlüssel sind mehr Arbeit als Passwörter

SSH-Schlüssel sind schwieriger als Passwörter. Benutzer müssen Schlüssel erzeugen, Private Keys sicher aufbewahren, Public Keys während der Anmeldung kopieren, verstehen, welcher Schlüssel zu welchem Gerät gehört, und Backups behalten.

Ein Passwort-Login ist für die meisten Menschen einfacher, besonders weil die meisten Menschen bei der Wahl eines E-Mail-Anbieters nicht auf maximale Privatsphäre fokussiert sind. Privacy.Fish ist für Benutzer gebaut, die entschieden haben, dass der Schutz ihrer Privatsphäre einen zusätzlichen Schritt wert ist, weil sie Authentifizierung wollen, die als sicherer gilt als Benutzername + Passwort.

## Ist SSH-Authentifizierung für deinen Dienst eine Überlegung wert?

Wir nutzen SSH-Schlüssel, weil sie zur Privacy.Fish-Philosophie passen: weniger sammeln, weniger speichern, weniger exponieren und den Benutzer die wichtigen Geheimnisse kontrollieren lassen. Passwortbasierte E-Mail ist bequem, aber diese Bequemlichkeit verschiebt Risiko und Daten tendenziell in die Hände des Anbieters. SSH-Schlüssel verschieben mehr Verantwortung zum Benutzer, lassen uns aber wiederverwendbare Mailbox-Passwörter vermeiden, ganze IMAP-/POP3-/SMTP-/Webmail-AUTH-Passwortflächen überspringen und E-Mail-Zugriff auf OpenSSH und age statt auf einen traditionellen Mailbox-Login-Stack aufbauen.

Das ist nicht für jeden der richtige Tradeoff. Für Privacy.Fish ist es der richtige Tradeoff.
