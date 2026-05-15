---
title: "E-Mail-Infrastrukturarchitektur"
description: "Technischer Überblick über die privacy.fish E-Mail-Infrastruktur, Vertrauensgrenzen, OpenBSD-Stack und operatives Sicherheitsmodell."
weight: 40
video: "/videos/security-model.mp4"
---

# Konzepte der Infrastruktursicherheit

Das Infrastruktursicherheitsmodell von Privacy.Fish kombiniert mehrere Ideen: so wenig wie möglich vertrauen, laufende Systeme klein halten und kompromittierte Server leicht ersetzbar machen. Dieser Abschnitt erklärt die wichtigsten Bausteine dieses Ansatzes: OpenBSD, OpenSSH, OpenSMTPD, minimaler eigener Code, kurzlebige Serverinstallationen, ersetzbare Admin-Workstations, norwegisches Hosting und die Teile der Infrastruktur, die nicht vollständig kontrolliert werden können.

Privacy.Fish tauscht serverseitige Bequemlichkeit gegen ein kleineres Vertrauensmodell: weniger Dienste, weniger gespeicherte Daten, weniger Passwörter, kurzlebige Server und E-Mail, die primär auf Geräten gespeichert wird, die du kontrollierst.

Sicherheit ist nie fertig. Es gibt immer eine weitere Kontrolle, ein weiteres Audit oder eine weitere Schicht, die ergänzt werden könnte. Privacy.Fish veröffentlicht Setup und Designentscheidungen auf [github.com/privacy-fish](https://github.com/privacy-fish "github.com/privacy-fish"), damit professionelle Security-Reviewer und die breitere technische Community sie prüfen, herausfordern und das System mit der Zeit verbessern können.

## Dinge, denen wir letztlich vertrauen müssen

Um administrative Arbeit überhaupt durchführen zu können, müssen wir den folgenden Dingen letztlich vertrauen:

- stw.no WebUI zur Verwaltung der Server und VPS (benötigt Benutzername + Passwort + TOTP 2FA)
- cdn.openbsd.org für die Bereitstellung von install.img- (VPS) und install.iso-Dateien (Hardware-Server)
- github.com
- der unten beschriebenen Admin-Workstation
- dem SSH Private Key der Admin-Workstation zur Administration der Server
- dem Client, damit er seinen Master Key schützen kann

## Hosting-Provider servetheworld.no

Alle unsere Server befinden sich in Norwegen und werden bei servetheworld.no / stw.no gehostet. Für unsere Produktionsinfrastruktur mieten wir zwei dedizierte Blade-Server sowie zwei Proxmox-gehostete VMs (VPS) von stw.no.

ServeTheWorld ist ein norwegischer Hosting-Anbieter mit Sitz in Oslo. Für Privacy.Fish ist die wichtigste Eigenschaft die Rechtsordnung: Die Server werden in Norwegen unter norwegischem Recht gehostet. Privacy.Fish ist eine norwegische LLC.

STW gibt an, dedizierte Server in Oslo-Rechenzentren unter norwegischem Recht bereitzustellen. STW gibt außerdem an, dass Privatsphäre durch norwegische Gesetzgebung unterstützt wird und Kundendaten nicht ohne norwegische gerichtliche Anordnung an Dritte herausgegeben werden. ([Dedicated servers](https://stw.no/en/dedicated-server/ "Dedicated servers"), [ServeTheWorld overview](https://stw.no/en/ "ServeTheWorld overview"))

Aus physischer und operativer Sicherheitsperspektive beschreibt STW sein Rechenzentrum in Oslo als Hochsicherheitsanlage mit Kapazität für etwa 4.000 Server, Schutz gegen Feuer, Diebstahl, Stromausfall und Überflutung, redundante Netzleitungen, redundante UPS, A/B-Stromversorgung, Firewalling und Dieselgenerator-Backup. Für Colocation beschreibt STW außerdem redundante Kühlung, redundante Strompfade, Rund-um-die-Uhr-Sicherheit und 24/7-Notfallpersonal. Die Bedingungen enthalten eine DSGVO-Auftragsverarbeitungsvereinbarung, wodurch STW für gehostete Kundendaten ein Auftragsverarbeiter ist. Privacy.Fish behandelt den Hosting-Anbieter trotzdem nicht als Datenschutzgrenze: Mail-Inhalte werden vor der Speicherung age-verschlüsselt, und das Serverdesign minimiert die Daten, die STW oder die Infrastruktur überhaupt in lesbarer Form verarbeiten müssen. ([Colocation](https://stw.no/en/colocation/ "Colocation"), [Data Processing Agreement](https://stw.no/wp-content/uploads/legal/nor-legal/Databehandleravtale_%28DBA%29.pdf "Data Processing Agreement"))

## OpenBSD als Betriebssystem

OpenBSD gilt weithin als eines der sichersten Betriebssysteme. Alle unsere Server und Admin-Workstations laufen mit OpenBSD.

Wir nutzen OpenBSD nicht nur wegen seines Rufs, sondern wegen seiner Sicherheitsphilosophie. OpenBSD ist darauf ausgelegt, standardmäßig sicher zu sein: nicht notwendige Dienste sind deaktiviert, das Basissystem wird kontinuierlich auditiert, und viele Sicherheitsmechanismen sind direkt in das Betriebssystem eingebaut, statt später als optionale Härtung ergänzt zu werden. Beispiele sind Privilege Separation, Privilege Revocation, chroot Jailing, W^X Memory Protection, randomisierte Speicherallokation, `pledge(2)` und `unveil(2)`. Für Privacy.Fish ist das wichtig, weil das Betriebssystem helfen soll, den Schaden eines kompromittierten Prozesses zu reduzieren, noch bevor unser eigenes Anwendungsdesign berücksichtigt wird. ([OpenBSD Security](https://www.openbsd.org/security.html "OpenBSD Security"), [OpenBSD Innovations](https://www.openbsd.org/innovations.html "OpenBSD Innovations"))

## Hauptsächlich OpenBSD-Basispakete

Unser E-Mail-Hosting-Stack ist hauptsächlich um zwei Programme gebaut, die mit jeder OpenBSD-Installation standardmäßig ausgeliefert werden: OpenSSH und OpenSMTPD. Sie sind Teil des sogenannten OpenBSD-Basissystems, was bedeutet, dass sie zusammen mit dem Betriebssystem gepflegt, auditiert, dokumentiert und veröffentlicht werden, statt zusätzliche Drittanbieter-Dienste zu sein, die wir ergänzen und separat verfolgen müssen.

## Minimaler eigener Code

Der größte Teil unserer Codebasis ist [Configuration-Management-Code](https://github.com/privacy-fish/playbook-infrastructure-privacy.fish "Configuration-Management-Code"), um Software auf den Servern zu installieren und zu konfigurieren.

TODO Links zu den eigenen Skripten / Tools

## OpenSMTPD für alles rund um SMTP

OpenSMTPD ist unser Open-Source-Mailserver der Wahl. Es ist Open-Source-Software, die als Teil des OpenBSD-Projekts entwickelt wird, und sogar im OpenBSD-Basissystem enthalten ist. Der Mailserver spricht mit gmail.com, wenn du E-Mail an/von example@gmail.com sendest oder empfängst. Er passt auch zu unserem Sicherheitsansatz, weil er klein, einfach und mit Sicherheit im Blick entworfen ist.

## SSH für Client-Interaktion

SSH (Secure Shell) ist der Standardweg, um sich sicher über das Internet auf den meisten heutigen Linux- und OpenBSD-Servern einzuloggen. Es nutzt eine verschlüsselte Verbindung und normalerweise kryptografische Schlüssel (wie deine SSH Public Keys) statt wiederverwendbarer Passwörter.

Wir vertrauen OpenSSH, weil OpenSSH weit verbreitet, Open Source und stark geprüft ist. Shodan listet OpenSSH derzeit als Top-SSH-Produkt mit mehr als 40 Millionen indexierten Instanzen, und das OpenSSH-Projekt beschreibt es als führendes Tool für Remote Login mit dem SSH-Protokoll. Diese Popularität ist für sich genommen kein Sicherheitsbeweis, bedeutet aber, dass die Software in der Praxis kontinuierlich genutzt, auditiert, angegriffen und gepatcht wird. ([OpenSSH](https://www.openssh.com/ "OpenSSH"), [Shodan SSH search](https://www.shodan.io/search?query=ssh "Shodan SSH search")).

Wir brauchen OpenSSH ohnehin zur Administration unserer Server, also nutzen wir es für mehr als einen Zweck. SSH für Client-Interaktion gibt Privacy.Fish ein schlüsselbasiertes Authentifizierungssystem, weniger exponierte Login-Dienste und kein wiederverwendbares Mailbox-Passwort.

### Warum SFTP statt IMAP/POP3

Privacy.Fish priorisiert Privatsphäre und Sicherheit vor Funktionen. Unsere Aufgabe ist einfach: E-Mail so privat und sicher wie möglich senden und zustellen. Wir denken, dass dieser Ansatz zu datenschutzbewussten E-Mail-Benutzern passt, die weniger serverseitige Funktionen bevorzugen, statt uns mehr Mailbox-Daten zu geben, selbst wenn dieser Zustand verschlüsselt ist.

IMAP ist nützlich, weil Clients mit serverseitigen Ordnern, Nachrichtenflags, Suche, selektiven Downloads, Synchronisierung und langlebigen Mailboxen arbeiten können. IMAP-Server wie [Dovecot](https://www.dovecot.org/ "Dovecot"), [Cyrus IMAP](https://www.cyrusimap.org/ "Cyrus IMAP") und [Courier-IMAP](https://www.courier-mta.org/imap/ "Courier-IMAP") können synchronisierte Ordner, Gelesen/Ungelesen-Zustand, Gesendet-Kopien, Suche, Indizes, parallele Clients und langlebige serverseitige Mailboxen bereitstellen. Das ist nützlich, erfordert aber zusätzlichen laufenden Code und zusätzlichen Mailbox-Zustand auf dem Server, selbst wenn verschlüsselt.

Privacy.Fish verzichtet bewusst auf diese Funktionen. Wir vermeiden das Sammeln und Speichern dieses Mailbox-Zustands, weil wir diese Daten nicht wollen, nicht einmal verschlüsselt. Privacy.Fish ist temporärer verschlüsselter Zustellspeicher, kein langfristiges serverseitiges Mailarchiv. Das passt auch zum 14-Tage-Speichermodell.

Mit Privacy.Fish können mehrere Geräte dieselbe verschlüsselte E-Mail herunterladen und entschlüsseln, bevor eines der Geräte sie von unseren Servern löscht. Dein Mail-Client kann E-Mails weiterhin in Ordner sortieren, durchsuchen und unbegrenzt auf deinen eigenen Geräten speichern. Was Privacy.Fish nicht unterstützt, ist serverseitige Synchronisierung von Ordnerbewegungen, Gelesen/Ungelesen-Zustand, Gesendet-Kopien, Suchindizes oder anderem Mailbox-Zustand zwischen Geräten. Wenn du das willst, muss es auf deinen Geräten oder in deinem eigenen Tooling passieren, zum Beispiel mit einem Thunderbird-Plugin.

Unserer Meinung nach ist der Sicherheitsgewinn durch das Vermeiden zusätzlicher laufender Codezeilen von Dovecot, Courier oder Cyrus sowie das Vermeiden der zugehörigen Mailbox-Zustände, selbst wenn verschlüsselt, den Komfortverlust wert.

### Warum SSH statt SMTP AUTH

SMTP AUTH ist der übliche Weg, wie E-Mail-Clients sich vor dem Senden authentifizieren. Es basiert häufig auf Kontopasswörtern oder passwortähnlichen Tokens und ergänzt einen weiteren Authentifizierungspfad, der konfiguriert, überwacht, gesichert und im Internet exponiert werden muss.

Privacy.Fish vermeidet das, indem es SSH für authentifiziertes Senden nutzt. Der Client oder die App übermittelt eine vorbereitete E-Mail-Nachricht über einen SSH-Befehl. Eine minimale `.eml`-Datei sieht so aus:

```text
From: j.doe@privacy.fish
To: friend@example.com
Subject: Hello

Hello from Privacy.Fish.
```

Auf der CLI sieht das Senden dieser Datei so aus:

```bash
cat msg.eml | ssh j.doe@out.mail.privacy.fish
```

Das erlaubt uns, dasselbe SSH-schlüsselbasierte Kontomodell zum Senden zu verwenden, das wir bereits zum Herunterladen von Mail und Verwalten von Geräten nutzen. Der Server kann einschränken, was der SSH-Schlüssel tun darf, wiederverwendbare Mailbox-Passwörter vermeiden und die Authentifizierungsfläche kleiner halten, als wenn sowohl SSH als auch SMTP AUTH betrieben würden.

### Warum SSH für Benutzer-Authentifizierungsverwaltung

Wir haben SFTP für die Verwaltung von OpenSSH Public Keys gewählt, weil SSH ohnehin installiert war ;)

## Rollenisolierung mit OpenBSD VMM

Wir betreiben derzeit mehrere OpenBSD-VMs auf einem physischen Server (wobei der „Failover“-Server in einem anderen Rechenzentrum gehostet wird), und jede VM hat eine dedizierte Rolle. Zum Beispiel laufen eingehende Mail, ausgehende Mail, gespeicherte verschlüsselte Mail, Key Management und Monitoring in getrennten VMs. Kommunikation zwischen VMs erfolgt über SSH und wird so weit wie möglich eingeschränkt.

Diese Isolation begrenzt die Wirkung mancher Kompromittierungen: Root-Zugriff auf einer VM bedeutet nicht automatisch Root-Zugriff auf jeder anderen VM. OpenBSD VMM ist jedoch trotzdem eine weitere Schicht laufenden Codes und damit eine weitere Sicherheitsgrenze, der wir vertrauen müssen. Die stärkere Alternative wäre, jede Rolle auf separater physischer Hardware zu betreiben. Das ist eine mögliche zukünftige Verbesserung, aber derzeit fokussieren wir Ressourcen auf die Kontrollen, die für die aktuelle Dienstgröße den größten Sicherheitswert liefern. Separate physische Server würden das Risiko verringern, dass ein Angreifer über die Virtualisierungsschicht von einer VM zur anderen wechselt.

## Angreifer-Persistenz durch Neuaufbauten begrenzen

Nach einer Serverkompromittierung ist der saubere Weg, einen frischen Server aufzusetzen und nur die erforderlichen Daten zu migrieren, die überprüft werden müssen, damit sie nicht manipuliert wurden.

Privacy.Fish tut das jede Woche, weil es immer vernünftig ist anzunehmen, dass eine 0day-Schwachstelle uns betroffen haben könnte und nun gefunden und gepatcht wurde. Wöchentliche Neuaufbauten sollen Zugriff von Exploits entfernen, die nach dem Patch nicht mehr funktionieren. Sie schützen nicht gegen einen ungepatchten 0day, der weiterhin funktioniert.

### Wie 0day-Exploits Privacy.Fish beeinflussen

0day-Exploits sind Softwarefehler, die noch nicht öffentlich bekannt oder zumindest noch nicht von den betroffenen Software-Maintainern gepatcht sind. OpenBSD hat eine starke Sicherheitsbilanz, ist aber keine Magie: Ernsthafte Schwachstellen werden weiterhin gefunden und gepatcht.

Im Januar 2020 fand [Qualys einen OpenSMTPD-Bug](https://www.qualys.com/2020/01/28/cve-2020-7247/lpe-rce-opensmtpd.txt "Qualys fand einen OpenSMTPD-Bug"), bei dem eine speziell konstruierte E-Mail-Adresse den Mailserver in manchen Konfigurationen dazu bringen konnte, Shell-Befehle als root auszuführen. Im Februar 2020 fand [Qualys einen weiteren OpenSMTPD-Bug](https://www.qualys.com/2020/02/24/cve-2020-8794/lpe-rce-opensmtpd-default-install.txt "Qualys fand einen weiteren OpenSMTPD-Bug"), bei dem ein bösartiger oder kompromittierter Remote-Mailserver Befehle in den internen Delivery-State von OpenSMTPD injizieren konnte. Ebenfalls im Februar 2020 wurde ein OpenBSD-`vmm`-Problem gemeldet, bei dem ein VM-Gast in Host-Speicher schreiben konnte; OpenBSD veröffentlichte einen [Errata-Patch vom 17. Februar 2020](https://www.openbsd.org/errata66.html "Errata-Patch vom 17. Februar 2020").

Deshalb ist es vernünftig, jeden internetexponierten Server als potenziell kompromittierbar durch einen ausreichend fähigen Angreifer zu behandeln. Die Infrastruktur ist daher um eine einfache Frage gebaut: Wenn ein Angreifer heute root-Zugriff bekommt und der genutzte Exploit später gefunden und gepatcht wird, wie stellen wir sicher, dass dieser Zugriff den nächsten Neuaufbau nicht überlebt? Das schützt nicht gegen einen ungepatchten 0day, der nach dem Neuaufbau weiterhin funktioniert. In diesem Fall könnte der Angreifer den frischen Server einfach erneut ausnutzen. Das Neuaufbau-Modell geht hauptsächlich darum, Persistenz zu verhindern, nachdem der ursprüngliche Einstiegspunkt geschlossen wurde.

0day-Exploits sind Teil der modernen IT-Welt, und es gibt keinen Weg daran vorbei. Alle Software, die wir nutzen, selbst „sehr sicher“ zu schreiben, würde die Bugs nicht entfernen, die jeder Mensch (und jede Maschine) in Code einführt. Deshalb wählen wir Software, die bereits sehr gut auditiert ist.

### Angreifer-Persistenz reduzieren, indem alles einmal pro Woche von Grund auf neu installiert wird

In einem traditionellen langlebigen Server-Setup läuft das Betriebssystem vielleicht Monate oder Jahre. Wenn ein Angreifer einen 0day-Exploit nutzt, bevor er öffentlich bekannt und gepatcht ist, kann das spätere Installieren des Sicherheitsupdates den ursprünglichen Fehler schließen, aber den bestehenden Fuß in der Tür des Angreifers nicht entfernen. Eine versteckte Backdoor, veränderte Binary, geänderte Konfigurationsdatei oder hinzugefügter Schlüssel kann nach dem Patch auf dem Server bleiben.

Privacy.Fish versucht, dieses Risiko zu reduzieren, indem Server als ersetzbare Infrastruktur behandelt werden. Statt Server nur in-place zu aktualisieren, bauen wir sie regelmäßig mit Configuration Management von Grund auf neu. Die aktuelle Ziel-Lebensdauer einer physischen Serverinstallation ist eine Woche, und laufende Systeme prüfen und installieren OpenBSD-Sicherheitsupdates während dieser Zeit alle 15 Minuten.

Das ändert, was ein Angreifer erreichen muss. Nachdem der 0day-Exploit, den der Angreifer verwendet hat, gefunden und gepatcht wurde und die Server neu aufgebaut wurden, müsste der Angreifer in der sehr kleinen Datenmenge persistieren, die vom alten Server auf den neuen übertragen wird, oder die Methode ausnutzen, die diese Daten überträgt (`rsync`). Andernfalls würde ihn der nächste Neuaufbau endgültig aussperren.

### Angreifer-Persistenz in dynamischen Daten

Dynamische Daten sind Daten, die unsere Benutzer erzeugen und verändern. Wir müssen diese Daten daher mindestens wöchentlich vom alten auf den neuen Server migrieren. Das sind:

- gespeicherte verschlüsselte Benutzer-E-Mails (`.age`-Dateien)
- signierter Account-Key-Zustand, einschließlich Master Keys, Mail-Operation-Keys, Recovery Keys, Versionen, vorheriger State-Hashes und Signaturen

Die eigentlichen `~/.ssh/authorized_keys`-Dateien werden während Server-Setup und Propagation aus signiertem Account-Key-Zustand erzeugt.

Alles andere am Server wird durch Code kontrolliert, der auf [github.com/privacy-fish](https://github.com/privacy-fish "github.com/privacy-fish") liegt. Der Configuration-Management-Code und alle anderen Automatisierungen sind auf GitHub veröffentlicht. Die Liste registrierter Benutzer ist nicht veröffentlicht, weil sie private operative Daten ist, wird aber vom Configuration Management genutzt, um die erforderlichen Benutzerkonten auf einer frischen OpenBSD-Serverinstallation neu zu erstellen.

#### Manipulation gespeicherter verschlüsselter E-Mail-Dateien abmildern

Der aktuelle Schutz kommt von der age-Verschlüsselung der E-Mail-Dateien selbst. Gespeicherte E-Mails werden für die SSH Public Keys der Benutzer verschlüsselt, und age-Verschlüsselung enthält Authentifizierung. Wenn eine bestehende verschlüsselte Maildatei verändert wird, schlägt die Entschlüsselung fehl, statt still geänderte Mail zu erzeugen. Die Client-Anwendung interpretiert age-Entschlüsselungsfehler klar und warnt den Benutzer, wenn verschlüsselte Mail nicht entschlüsselt werden kann.

Das erkennt nicht jede mögliche Mailbox-Manipulation auf dem Fetch-Server. Wenn ein Angreifer Root-Zugriff auf dem Fetch-Server hat, könnte er gespeicherte `.age`-Dateien löschen, ersetzen oder neue `.age`-Dateien hinzufügen, die für die Public Keys des Benutzers verschlüsselt sind. Veränderte Dateien werden durch age-Entschlüsselungsfehler erkannt, aber Löschung, Ersetzung und neu hinzugefügte gültige `.age`-Dateien werden durch age allein nicht erkannt.

Das Signieren gespeicherter verschlüsselter Maildateien ist ein TODO. Es wird als Infrastrukturarbeit in den [playbook-infrastructure-privacyfish issues](https://github.com/privacy-fish/playbook-infrastructure-privacyfish/issues "playbook-infrastructure-privacyfish issues") verfolgt.

Signaturen würden Angriffe folgendermaßen stören:

| Angriff | Age schützt | Signatur schützt | Wie Signatur schützt | Auswirkung, wenn unerkannt |
| --- | --- | --- | --- | --- |
| Eine gespeicherte `.age`-Datei verändern | Ja | Ja | Age erkennt das bereits bei der Entschlüsselung, weil veränderte verschlüsselte Dateien an der age-Authentifizierung scheitern. Signieren würde eine frühere Prüfung ergänzen: Jede Byte-Änderung würde die Signaturprüfung vor der Entschlüsselung scheitern lassen. | Der Benutzer könnte still geänderte verschlüsselte Mail erhalten. Mit age allein sollte das bereits bei der Entschlüsselung fehlschlagen, aber Signaturen würden es früher erkennen. |
| Eine gespeicherte `.age`-Datei durch eine andere Datei ersetzen | Nein | Ja | Die Ersatzdatei würde nicht zur ursprünglichen Signatur passen, es sei denn, der Angreifer hätte auch den Signaturschlüssel. Der Fetch-Server hat keinen Zugriff auf den Signing-Server; nur der eingehende Mailpfad kann Signaturen für neu verschlüsselte Mail anfordern. | Der Client könnte eine andere verschlüsselte Nachricht akzeptieren, als wäre sie die ursprüngliche. |
| Eine neue `.age`-Datei hinzufügen | Nein | Ja | Der Client würde verlangen, dass jede heruntergeladene `.age`-Datei eine gültige Privacy.Fish-Signatur hat. Unsigned files oder Dateien mit ungültiger Signatur würden abgelehnt. | Der Client könnte eine gefälschte verschlüsselte Nachricht verarbeiten, zum Beispiel um den Benutzer zu verwirren oder den Mailclient anzugreifen. |
| Eine gespeicherte `.age`-Datei löschen | Nein | Nein | Pro-Message-Signaturen würden Löschung nicht erkennen, wenn der Angreifer sowohl die verschlüsselte Datei als auch die passende Signaturdatei löscht. | Der Benutzer sieht die gelöschte Nachricht möglicherweise nie. Löschungserkennung würde eine signierte Mailbox-Liste oder ähnlichen globalen Mailbox-Zustand erfordern. |

...WENN ein Angreifer Zugriff auf den Fetch-Server erlangt. OpenBSD und OpenSSH zu vertrauen ist eine sehr faire Wette.

Wir verbessern unseren Dienst kontinuierlich und ziehen es vor, unseren Benutzern ehrlich zu sagen, was implementiert ist und was (noch) nicht.

#### Manipulation von SSH Authorized Keys abmildern

Privacy.Fish nutzt einen dedizierten Key-Management-Server, auf dem Benutzer ihre SSH Public Keys hochladen, löschen und ändern können, die zum Senden und Herunterladen von E-Mail sowie zur Verschlüsselung von E-Mail-Dateien verwendet werden. Nur die Master-SSH-Schlüssel des Benutzers können sich auf dem Keys-Server anmelden und Updates durchführen. Das hält normale Mail-Schlüssel getrennt von Account-Management-Schlüsseln.

Privacy.Fish behandelt Key-Änderungen als signierte Benutzerabsicht. Statt der aktuellen `authorized_keys`-Datei auf dem Key-Management-Server als Quelle der Wahrheit zu vertrauen, lädt der Benutzer signierten Account-Key-Zustand hoch: welche Master Keys das Konto verwalten und welche Mail-Operation-Keys auf Mail zugreifen dürfen. Neue Mail-Operation-Keys müssen vom aktuell vertrauenswürdigen Master Key signiert sein. Ein neuer Master Key muss vom vorherigen vertrauenswürdigen Master Key oder von einem bereits registrierten Recovery Key signiert sein.

Der signierte Account-Key-Zustand wird außerhalb des Key-Management-Servers gespeichert und verwendet, um die tatsächlichen `authorized_keys`-Dateien für die Mailserver zu erzeugen. Frische Server erzeugen `authorized_keys` aus dem letzten verifizierten signierten Zustand neu, statt die Datei des alten Servers direkt zu kopieren. Das bedeutet, dass ein Angreifer mit Root-Zugriff auf dem Key-Management-Server den Live-Dienst weiterhin stören oder ungültige Key-Daten übermitteln kann, aber keinen gültigen neuen Key-Zustand erstellen kann, sofern er nicht auch den aktuellen vertrauenswürdigen Master Private Key des Benutzers hat.

Bei der Migration von einem alten Server auf einen frischen Server werden gespeicherte Benutzer-E-Mails mit `rsync` kopiert, wobei der neue Server die Dateien vom alten Server zieht. SSH Public Keys werden aus dem letzten verifizierten signierten Account-Key-Zustand neu erzeugt.

### Fazit

Weil das Server-Setup automatisiert ist, braucht ein neuer Mailserver nur Minuten, um mit OpenBSD installiert, konfiguriert, mit den alten Daten migriert und mit umgeschalteten IPs bereit für Privacy.Fish-E-Mail zu sein.

Das Ziel ist nicht zu behaupten, dass Kompromittierung unmöglich ist. Das Ziel ist, Angreifer-Persistenz technisch so unwahrscheinlich wie möglich zu machen, indem Server kurzlebig gehalten, schnell und zuverlässig per Automation neu aufgebaut und nur die minimalen Daten weitergetragen werden, die nötig sind, damit der Dienst funktioniert.

Das ist der Privacy.Fish-Tradeoff in einem Satz: weniger Serverzustand, weniger laufende Dienste, weniger gespeicherte Daten und mehr Verantwortung auf der Benutzerseite, besonders bei Geräten, Schlüsseln und lokaler Mail-Speicherung.

## Minimale und ersetzbare Admin-Workstation

Einen Administrator zu kompromittieren kann einfacher sein, als einen gehärteten Server direkt anzugreifen.

Die Admin-Workstation ist möglicherweise sogar angreifbarer als die Server, weil sie nicht in einem Rechenzentrum steht, sondern im Zuhause des Administrators. Wir arbeiten ausschließlich aus dem Home Office.

Die Workstation des Administrators ist, wie die Server, dafür entworfen, regelmäßig vollständig ersetzt zu werden, und besteht aus:

- Raspberry Pi
- MicroSD-Karte
- Externem Bildschirm
- Tastatur
- Maus
- Netzwerkkabel (oder WLAN, wo unvermeidbar)
- 2FA-TOTP-Gerät für stw.no

Wir nutzen Configuration-Management-Code auf [github.com/privacy-fish](https://github.com/privacy-fish "github.com/privacy-fish"), um die Workstations einzurichten, damit du es selbst ansehen kannst.

Die einzigen „dynamischen Daten“, die migriert werden müssen, sind der SSH Private Key des Administrators. Das Passwort für den Login bei stw.no sowie das Passwort zur Nutzung des SSH Private Keys können auswendig gelernt werden.

Auf der Administrator-Workstation sind pf-Firewall-Regeln aktiv, die nur SSH zu den Servern und das Öffnen der WebUI von stw.no, unserem Server-Hosting-Anbieter, erlauben.

Zusätzlich nutzen wir mullvad.net auf allen Workstations, nur für den Fall, mit einem VPN nach Norwegen.

Das ist das vernünftig sicherste Workstation-Setup, das uns eingefallen ist.

## Dinge, die wir nicht kontrollieren können

Das Neuaufbau-Modell reduziert die Wahrscheinlichkeit, dass ein Angreifer in einem Mailserver persistieren kann, beseitigt aber nicht jede Abhängigkeit rund um den Server. Einige Schichten können von Privacy.Fish allein nicht vollständig verifiziert oder regeneriert werden: das Hosting-Provider-Control-Panel, der Account-Recovery-Prozess des Providers, Server-Hardware und Firmware unterhalb des Betriebssystems, Quellcode- und Configuration-Management-Repositories auf GitHub, Backups und der Migrationsprozess zwischen alten und neuen Servern.

Die Abmilderungen sind geschichtet. Provider-Zugriff ist mit 2FA geschützt und auf die Admin-Workstation beschränkt. OpenBSD-Basissystem-Updates sind signiert und werden über den normalen OpenBSD-Updateprozess angewendet. Serverzustand wird aus Configuration Management erzeugt statt manuell gepflegt. Backups sollen nur die Daten wiederherstellen, die nicht regeneriert werden können. Migration soll die dynamischen Daten prüfen, die in einen frischen Server kopiert werden.

Wenn unsere Quell-Repositories auf GitHub, Firmware / BIOS der Server, Backups, unser STW-Konto oder STW selbst oder der Migrationsprozess kompromittiert wären, könnte ein Neuaufbau trotzdem Risiko weitertragen. Das Ziel ist, diese Abhängigkeiten sichtbar, klein und wo möglich ersetzbar zu halten.

## Welche Server existieren

### alpha.mail.privacy.fish (Hardware)
#### in.mail.privacy.fish (VMM)
#### spam-in.mail.privacy.fish (VMM)

#### out.mail.privacy.fish (VMM)
#### spam-out.mail.privacy.fish (VMM)

#### fetch.mail.privacy.fish (VMM)
#### keys.mail.privacy.fish (VMM)

### omega.mail.privacy.fish (Hardware)

Omega ist exakt dieselbe Hardware wie der physische Alpha-Blade-Server. Mit Configuration Management und Automation können wir die Privacy.Fish-E-Mail-Infrastruktur genauso wie auf Alpha aufsetzen, was wir regelmäßig tun.

Nach der Migration der E-Mails und SSH Public Keys unserer Benutzer auf den anderen Server schalten wir die öffentlichen IP-Adressen unserer Dienste auf den frisch eingerichteten Server um.

Daher wechseln wir den Betrieb wöchentlich zwischen Alpha und Omega, wobei der Server jedes Mal von Grund auf neu aufgesetzt wird.

### backup.mail.privacy.fish (VPS)
### www.privacy.fish (VPS)

## Liste relevanter Git-Repositories
### Ansible Playbook Mail
### Ansible Playbook Web
### website-privacy.fish

## Server für die E-Mail-Infrastruktur neu erstellen

## Server für die Web-Infrastruktur neu erstellen

## Wöchentlicher Server-Lifecycle

## Monitoring
### Setup
### Alerting
### Incident Response
### Incident Reports

## Backups
### Was gesichert wird

/var/log/legal-access.log - SSH / SFTP-Verbindungen: IP + Port + Zeitstempel
Für 12 Monate

/home/<user>/Maildir
Für 7 Tage

/home/user/.ssh/authorized_keys
Für 7 Tage

### Backup-Löschung nach Löschung

TODO E-Mails in Backups löschen, wenn sie per SFTP gelöscht werden

### Restore-Prozess

Die Backups sind nur für einen E-Mail-Server-Systemausfall gedacht. Wir bieten unseren Benutzern keine Wiederherstellung von Backups an.
