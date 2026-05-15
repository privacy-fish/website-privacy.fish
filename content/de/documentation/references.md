---
title: "Referenzen"
description: "Referenzhinweise zu Kontolimits, unterstützten Domains, Speicherdauer, Kompatibilität und operativen Annahmen."
weight: 80
video: "/videos/data-minimization-by-default.mp4"
---

## Limits
### Geräte pro Konto

Du kannst 10 Geräte pro Konto verwenden. Du kannst mehr verwenden, wenn du SSH falsch benutzt.

Wir erlauben bis zu 10 hochgeladene SSH Public Keys. Diese werden genutzt, um dir Zugang zum Dienst zu geben und deine empfangenen E-Mails zu verschlüsseln.

Jedes Gerät sollte seinen eigenen privaten und öffentlichen SSH-Schlüssel haben. Du solltest Private Keys nicht von einem Gerät auf ein anderes kopieren, sondern für jedes Gerät immer ein neues SSH-Schlüsselpaar erzeugen. Die einzige Ausnahme ist ein Backup-SSH-Schlüsselpaar, das du offline speichern willst, zum Beispiel auf einem USB-Stick.

### Rotierende Aliase
### Unterstützte Domains

Dein Benutzername und Alias funktionieren immer mit allen hier aufgeführten Domains:

- @privacy.fish
- @pfi.sh
- privacyfish.net
- datenschutzfisch.de
- detaeschutzfischli.ch

### Speicherdauer und Größe

Wenn unsere Server eine E-Mail für dich empfangen, wird sie verschlüsselt, bevor die verschlüsselte E-Mail auf die Festplatte geschrieben wird. Danach hast du bis zu 14 Tage Zeit, die E-Mail herunterzuladen. Du kannst sie anschließend von unseren Servern löschen, oder sie wird nach 14 Tagen automatisch gelöscht.

Deine Mailbox ist 100 MB groß und kann nicht erweitert werden. Wenn deine Mailbox voll ist, werden eingehende E-Mails abgewiesen und die Person, die dir eine E-Mail gesendet hat, erhält eine E-Mail mit dem Fehler „SMTP 552 Mailbox full“.

100 MB ist für zwei Wochen viel Platz. Hier ist [die Quelle, auf der diese Berechnung basiert](https://www.emailtooltester.com/en/blog/email-usage-statistics/ "die Quelle, auf der diese Berechnung basiert"):

```
80.6 E-Mails pro Tag × 0.075 MB pro E-Mail × 7 Tage ≈ 42.3 MB pro Woche
```

Wir empfehlen, deine E-Mails mindestens wöchentlich herunterzuladen und sie danach von unseren Servern zu löschen.

## Kompatibilität
### Unterstützte Betriebssysteme

Derzeit wird nur der Kommandozeilen-Client unterstützt.

Wir arbeiten an grafischen Apps für alle anderen Betriebssysteme.

### Unterstützte E-Mail-Clients

Die Client-App nutzt deine heruntergeladenen und entschlüsselten E-Mails (`.eml`-Dateien), um sie über POP3 jeder E-Mail-Client-App / jedem Mail User Agent deiner Wahl bereitzustellen. Alle E-Mail-Clients, die POP3 und SMTP unterstützen, sollten funktionieren.

### Bekannte Einschränkungen

## Fehlerbehebung

### Kundensupport

Wenn du Probleme bei der Nutzung des Dienstes hast, öffne bitte ein Issue auf [github.com/privacy-fish](https://github.com/privacy-fish "github.com/privacy-fish") und beschreibe dein Problem. Wenn es sinnvoll ist, füge einen Screenshot hinzu.

Für technische Probleme bieten wir keinen Support per E-Mail oder auf anderem Weg als GitHub an.

### Kontowiederherstellung

Da wir die Zahlung-zu-Benutzerkonto-Zuordnung 14 Tage nach Eingang deiner Zahlung löschen, können wir nur bis dahin identifizieren, welche natürliche Person das Konto erstellt hat. Danach haben wir keine Möglichkeit mehr zu wissen, welcher Mensch dieses Konto tatsächlich besitzt.

TODO beschreiben, wie Key-Privilegien funktionieren

## Rechtliches

### Rechtlicher Kontakt

Bitte sende uns eine E-Mail an legal@privacy.fish

### Allgemeine Geschäftsbedingungen

Unsere Allgemeinen Geschäftsbedingungen findest du unter https://privacy.fish/legal/terms-and-conditions

### Erstattungen

Erstattungen sind für Bargeld-per-Brief- und Kryptowährungszahlungen nicht möglich.

Erstattungen sind nur bis 14 Tage nach Eingang deiner Zahlung möglich. Wenn du dein Konto innerhalb dieser Zeit erstellt hast, sende bitte eine E-Mail an refunds@privacy.fish und gib den Zahlungscode an, den du für die Zahlung verwendet hast, sowie Zeitpunkt und Zahlungsmethode.

Wir antworten dann innerhalb von fünf Arbeitstagen auf deine E-Mail.

### Missbrauchsbearbeitung

Wir handeln nur bei gültigen norwegischen Rechtsanordnungen oder wenn es nötig ist, die Privacy.Fish-Infrastruktur während eines Angriffs zu schützen.

#### Nicht autorisierte Zahlungen

Wenn du eine nicht autorisierte Zahlung an uns festgestellt hast, kontaktiere bitte legal@privacy.fish mit einem Screenshot der Transaktion.

#### Nicht autorisierte Kontonutzung

Da wir keine Möglichkeit haben zu identifizieren, welcher Mensch welches Privacy.Fish-Konto besitzt, können wir dir bei der Wiederherstellung deines Kontos in keiner Weise helfen.

### Du hast Spam oder Viren von einem @privacy.fish-Konto erhalten

Bitte leite die Spam- oder bösartige E-Mail an spam@privacy.fish weiter.

### Rechtmäßige Anfragen

Bitte sende uns eine E-Mail an legal@privacy.fish
