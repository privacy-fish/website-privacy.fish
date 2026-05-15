---
title: "Benutzerhandbuch"
description: "Benutzerorientierte Hinweise zu Kontoerstellung, SSH-Schlüsseln, Zahlung, Bereitstellung, E-Mail-Zugriff, Senden, Empfangen und Löschung."
weight: 20
video: "/videos/secure-email-workflow.mp4"
# @richtermac, die folgenden Abschnitte müssen jeweils für CLI, macOS, Linux, Windows, Android und iPhone geschrieben werden
# Dafür werden wahrscheinlich 6 Boxen mit Betriebssystem-Logo benötigt, die auf die jeweilige Anleitung verlinken
# documentation/user-manual/command-line/account-creation
# documentation/user-manual/android/account-creation
#
# Und so weiter. Alle Überschriften unter "# Benutzerhandbuch" sind geräteabhängig. Ausnahmen sind "### Eine private Zahlung ausführen" und "### Wie wir dein Konto aktivieren".
---

## Kontoerstellung
### Client-App installieren
### SSH-Schlüssel für deine Geräte erzeugen
### Das Anmeldeformular verwenden
### Eine private Zahlung ausführen

### Wie wir dein Konto aktivieren

Innerhalb von 24 Stunden nach Eingang deiner Zahlung zusammen mit dem zugehörigen Zahlungscode erstellen Privacy.Fish-Administratoren dein Konto manuell.

Das Anmeldeformular fragt nicht nach identifizierenden Kontaktinformationen. Deshalb haben wir keine Möglichkeit, dich zu benachrichtigen, wenn das Konto bereit ist. Der sicherste Ansatz ist, nach der Zahlung bis zu 24 Stunden zu warten, bevor du versuchst, das Konto zu nutzen.

Auch wenn wir sofortige Zahlungsmethoden wie PayPal und Kreditkarten anbieten, verbinden wir den Signup-Webserver absichtlich nicht mit den Mailservern, um neue Konten automatisch zu erstellen. Dem Webserver, auf dem das Anmeldeformular läuft, wird nicht vertraut, automatisch Konten auf der E-Mail-Server-Infrastruktur zu erstellen.

Wenn du das Anmeldeformular absendest, werden dein gewünschter Benutzername, deine SSH Public Keys und der temporäre Zahlungscode in age-verschlüsselter Form auf dem Webserver gespeichert. Innerhalb von 24 Stunden werden die Daten heruntergeladen, vom Administrator entschlüsselt und danach sicher vom Webserver gelöscht. Mit den entschlüsselten Kontodaten prüfen unsere Administratoren manuell neu eingegangene Zahlungen und erstellen dein Konto auf den E-Mail-Servern.

## Client-App mit deinem Mail User Agent verbinden

## E-Mail senden
### Normale Zustellung
### Warnungen bei unsicherem TLS
### Aliase verwenden

## E-Mail empfangen
### Wie E-Mail in deiner App heruntergeladen und entschlüsselt wird

## E-Mail von unseren Servern löschen

## Aliase bestimmen und verwenden

## Geräte und SSH-Schlüsselpaare verwalten
### Ein Gerät hinzufügen
### Ein Gerät entfernen
### Einen Backup-SSH-Schlüssel erstellen
### Ein Konto mit einem Backup-SSH-Schlüssel wiederherstellen
