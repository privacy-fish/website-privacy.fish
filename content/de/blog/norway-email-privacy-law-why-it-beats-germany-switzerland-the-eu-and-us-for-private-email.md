---
title: "Norwegisches E-Mail-Datenschutzrecht: Warum es Deutschland, die Schweiz, die EU und die USA für private E-Mail schlägt"
description: "Norwegisches E-Mail-Datenschutzrecht gegenüber Deutschland, Schweiz, UK, Australien und USA: was Privacy.Fish nach Ekomloven protokollieren muss."
date: 2026-05-05
draft: false
video: "/videos/blog/norway-email-privacy-law-why-it-beats-germany-switzerland-the-eu-and-us-for-private-email/norway-email-privacy-law.mp4"
show_cta: false
---

## Dein E-Mail-Anbieter ist Teil deines Threat Models

Gesetze zur Vorratsdatenspeicherung sind ein Grund, warum E-Mail-Privatsphäre nicht nur eine Frage von Verschlüsselung oder Produktdesign ist. Sie ist auch eine Frage des Landes, in dem der Anbieter arbeitet, und was dieses Land den Anbieter speichern, offenlegen oder technisch ermöglichen lässt. Regierungen brauchen oft nicht dein Gerät, wenn das Gesetz deinen Anbieter und dessen Daten erreichen kann. Damit ist die Rechtsordnung deines E-Mail-Anbieters Teil deines Privatsphäre-Modells.

Unsere Antwort hat zwei Teile: Betrieb aus Norwegen und Vermeidung von Daten, die wir nicht brauchen.

Für die längere operative Dokumentation siehe unseren [Rechtsordnungsvergleich](/documentation/jurisdiction/ "Rechtsordnungsvergleich"), unsere [Datenschutzinformationen](/documentation/privacy-information/ "Datenschutzinformationen") und unsere [E-Mail-Infrastrukturarchitektur](/documentation/email-infrastructure-architecture/ "E-Mail-Infrastrukturarchitektur").

## Die unbequeme Realität der E-Mail-Vorratsdatenspeicherung im Westen

### Deutschland und die Schweiz: Die Datenschutzlieblinge mit Zähnen

Nach deutschem [TKG § 170](https://www.gesetze-im-internet.de/tkg_2021/__170.html "TKG § 170") muss jeder Betreiber einer Telekommunikationsanlage, mit der öffentlich zugängliche Telekommunikationsdienste erbracht werden, ab Betriebsbeginn und auf eigene Kosten technische Einrichtungen für gesetzlich angeordnete Telekommunikationsüberwachung sowie organisatorische Vorkehrungen für deren unverzügliche Umsetzung bereithalten. Außerdem muss der Betreiber die Bundesnetzagentur informieren und nachweisen, dass Einrichtungen und Verfahren der Verordnung und technischen Richtlinie entsprechen. [TKÜV § 3](https://www.gesetze-im-internet.de/tk_v_2005/__3.html "TKÜV § 3") wendet diese Regeln auf Betreiber von Anlagen für öffentlich zugängliche Telekommunikationsdienste an, nimmt aber bestimmte Anlagen aus, darunter Anlagen mit höchstens 10.000 Nutzern und Anlagen, die ausschließlich für nummernunabhängige interpersonelle Telekommunikationsdienste oder kennungsunabhängige WLAN-Internetzugangsdienste mit höchstens 100.000 Nutzern genutzt werden. [TKG § 174](https://www.gesetze-im-internet.de/tkg_2021/__174.html "TKG § 174") verlangt sichere elektronische Schnittstellen für Behördenanfragen: Anbieter mit 100.000 oder mehr Vertragspartnern müssen sowohl die Schnittstelle als auch das E-Mail-basierte Übermittlungsverfahren bereithalten, kleinere verpflichtete Anbieter müssen das E-Mail-basierte Verfahren bereithalten.

Die Schweizer [VÜPF](https://www.fedlex.admin.ch/filestore/fedlex.data.admin.ch/eli/cc/2018/32/20240326/de/pdf-a/fedlex-data-admin-ch-eli-cc-2018-32-20240326-de-pdf-a-1.pdf "VÜPF") definiert E-Mail-Überwachungsarten, die Metadaten und in manchen Fällen Inhalt plus Metadaten umfassen. Artikel 58 zur Echtzeitüberwachung von E-Mail-Metadaten umfasst Login-/Logout-Status, AAA- oder Teilnehmerinformationen, Client- und Server-IP-Adressen und Ports, Protokoll, Zeitstempel, Datenvolumen, Absender- und Empfängeradressen sowie IP-Adressen und Ports der sendenden und empfangenden Mailserver für Senden, Empfangen, Mailbox-Verarbeitung, Download und Upload. Artikel 59 umfasst Echtzeitüberwachung von E-Mail-Inhalten plus die zugehörigen Metadaten aus Artikel 58. Artikel 62 zur rückwirkenden E-Mail-Metadatenüberwachung umfasst vergangene Ereigniszeiten, Ereignisarten, Teilnehmerkennungen, Aliasadressen, Absender-/Empfängeradressen, Protokoll, Server-/Client-IP-Adressen und Ports, Zustellstatus, Mailbox-Login/-Logout, verfügbare Download-/Upload-/Lösch-/Verarbeitungs-/Add-Message-Ereignisse sowie Namen und IP-Adressen sendender und empfangender E-Mail-Server.

### Das westliche Muster: Speichern, sichern, abfangen

Das wiederkehrende Muster ist nicht ein einzelnes Datenschutzproblem, sondern drei verschiedene rechtliche Befugnisse: erzwungene Metadatenspeicherung, erzwungene Sicherung oder Offenlegung bereits vorhandener Daten und verpflichtende Abfang- oder Auslieferungsfähigkeit. Im Vereinigten Königreich kann eine [retention notice nach dem Investigatory Powers Act 2016](https://www.legislation.gov.uk/ukpga/2016/25/section/87 "retention notice nach dem Investigatory Powers Act 2016") einen Telekommunikationsbetreiber verpflichten, alle relevanten Kommunikationsdaten oder bestimmte Kategorien davon bis zu 12 Monate zu speichern, einschließlich Internet Connection Records; [section 95](https://www.legislation.gov.uk/ukpga/2016/25/section/95 "section 95") verbietet dem Betreiber außerdem, Existenz oder Inhalt der Anordnung ohne Genehmigung offenzulegen. In Australien verlangen [Telecommunications (Interception and Access) Act 1979 sections 187A, 187AA and 187C](https://www.legislation.gov.au/C2004A02124/2026-01-22/2026-01-22/text/original/pdf/2 "Telecommunications (Interception and Access) Act 1979 sections 187A, 187AA and 187C"), dass erfasste Diensteanbieter Teilnehmer-, Konto-, Geräte-, Quellen-, Ziel-, Datums-, Zeit-, Dauer-, Kommunikationsart- und bestimmte Standortdaten zwei Jahre speichern, während [section 191](https://www.legislation.gov.au/C2004A02124/2026-01-22/2026-01-22/text/original/pdf/2 "section 191") Netzbetreiber ohne spezielle Ausnahme verpflichtet, Abfang- und Auslieferungsfähigkeit vorzuhalten. In den Vereinigten Staaten erlaubt [18 U.S.C. § 2703](https://www.govinfo.gov/link/uscode/18/2703 "18 U.S.C. § 2703") der Regierung, gespeicherte Kommunikationsdaten und Records sowie deren Sicherung für 90 Tage plus weitere 90 Tage bei erneuter Anforderung zu verlangen; [18 U.S.C. § 2713](https://www.govinfo.gov/link/uscode/18/2713 "18 U.S.C. § 2713") wendet diese Sicherungs-, Backup- und Offenlegungspflichten auf Daten im Besitz, Gewahrsam oder unter Kontrolle des Anbieters an, selbst wenn sie außerhalb der USA gespeichert sind.

### Norwegen: IP, Zeit, Port, kein Ziel-Log

Nach norwegischem [Ekomloven § 3-13](https://lovdata.no/lov/2024-12-13-76/%C2%A73-13 "Ekomloven § 3-13") müssen Anbieter elektronischer Kommunikationsnetze für öffentliche elektronische Kommunikationsdienste und Anbieter solcher Dienste Informationen speichern, die erforderlich sind, um Teilnehmer anhand einer öffentlichen IP-Adresse und Kommunikationszeit zu identifizieren, oder anhand öffentlicher IP-Adresse, Zeit und Portnummer, wenn dieselbe öffentliche IP-Adresse gleichzeitig mehreren Teilnehmern zugewiesen ist. Derselbe Abschnitt sagt, dass Zielinformationen nicht gespeichert werden dürfen, und legt die Speicherfrist auf 12 Monate ab dem Tag fest, an dem die Kommunikation endet. [Ekomforskriften § 3-5](https://lovdata.no/forskrift/2024-12-20-3410/%C2%A73-5 "Ekomforskriften § 3-5") ergänzt, dass die gespeicherten Daten die verwendete Zeitzone enthalten, gegen Verlust geschützt und während Erhebung, Protokollierung, Speicherung oder Offenlegung nicht verändert werden dürfen.

Das bedeutet nicht „kein rechtmäßiger Zugriff“: [Ekomloven § 3-12](https://lovdata.no/lov/2024-12-13-76/%C2%A73-12 "Ekomloven § 3-12") verlangt, dass Anbieter Netze und Dienste so betreiben, dass gesetzlich vorgeschriebener Zugriff auf Informationen über Endnutzer und elektronische Kommunikation gesichert ist, und [§ 3-14](https://lovdata.no/lov/2024-12-13-76/%C2%A73-14 "§ 3-14") regelt die Offenlegung gespeicherter IP-Adressdaten an Polizei oder Staatsanwaltschaft für Ermittlungen zu schweren Straftaten und bestimmten Delikten.

Norwegen ist [kein EU-Mitgliedstaat](https://www.eu-norway.org/eu/norway-and-the-eu/ "kein EU-Mitgliedstaat"), auch wenn große Teile der Beziehung zur EU über den EWR laufen.

## Wie Privacy.Fish norwegisches Recht in echte Privatsphäre übersetzt

Die rechtliche Grundlage hilft nur, wenn der Dienst so gebaut ist, dass alles außerhalb dieser Grundlage gar nicht erst entsteht.

### Was wir tun, um andere vom Zugriff auf deine Daten abzuhalten

Norwegisches Recht hilft nur, wenn unsere eigene Infrastruktur keine zusätzlichen Daten erzeugt, die Angreifer, Insider, Zahlungsanbieter oder spätere Rechtsanfragen erreichen können.

#### Minimale Protokollierung

Privacy.Fish trennt verpflichtende Zugriffsdaten von operativem Security-Logging. Die Zugriffsdaten, die nach der [norwegischen IP-Speicherpflicht](https://lovdata.no/lov/2024-12-13-76/%C2%A73-13 "norwegischen IP-Speicherpflicht") erforderlich sind, werden für diesen gesetzlichen Zweck gespeichert; andere Server-Logzeilen werden in ein internes Analysesystem ohne Internetzugang, ohne Drittanbieter-Reputationsabfragen und ohne externe Analytics gestreamt. Wenn die Analyse nichts Verdächtiges findet, werden diese Logzeilen automatisch verworfen statt gespeichert. Wenn etwas falsch aussieht, behalten wir nur die Logzeilen, die nötig sind, um das Problem zu debuggen, den Vorfall zu bearbeiten oder Beweise zu sichern, wo rechtliche Schritte dies verlangen.

#### Minimale Datenerhebung

Wir bauen unseren Dienst mit [OpenSSH](https://www.openssh.com/ "OpenSSH"), um E-Mail herunterzuladen und sie danach sofort von dir löschen zu lassen. Das tun wir nicht nur, weil OpenSSH eine ungewöhnlich lange kryptografische und operative Bewährungsprobe hinter sich hat, sondern weil die beste Art, Daten zu schützen, darin besteht, sie als Anbieter nicht zu speichern. Mit mehreren Geräten kannst du E-Mail auf alle herunterladen oder Backups auf Wechseldatenträgern, externen Festplatten oder deinem eigenen Cloud-Speicher halten.

Privacy.Fish nutzt SFTP für heruntergeladene E-Mail, damit der Server als kurzlebiger Zustellspeicher dient, nicht als dein dauerhaftes Postfach. Sobald eine Nachricht angekommen ist, kannst du die verschlüsselte `.eml`-Datei auf ein Gerät, mehrere Geräte, lokale Festplatten oder dein eigenes Backup-System herunterladen und sie danach sofort von unseren Servern löschen.

Unser Ansatz ist weniger bequem als ein synchronisiertes Webmail-Konto, aber er ändert das Risikomodell: Das langfristige Archiv liegt bei dir, während der Anbieter weniger Mailbox-Daten schützen, offenlegen, verlieren oder gestohlen bekommen kann.

### Einmalzahlungen und Löschung der Zahlung-zu-Konto-Verknüpfung nach 14 Tagen

Wiederkehrende Abrechnung re-identifiziert das Konto, weil jede Verlängerung dieselbe operative Frage erneut beantworten muss: Für welches Konto ist diese Zahlung? Eine monatliche Kartenbelastung, jährliche PayPal-Verlängerung, wiederkehrende SEPA-Überweisung oder Abo-Verwaltungsseite braucht entweder kontobezogene Metadaten in der Zahlungsreferenz, einen Kundendatensatz im Abrechnungssystem des Anbieters oder beides. Selbst wenn der Anbieter alles andere minimiert, wird diese wiederkehrende Abrechnungsbeziehung zu einer langlebigen Identitätsbrücke zwischen Zahlungsdaten und Mailbox-Konto.

Privacy.Fish löst die Verbindung zwischen Zahlung und Kontoname mit einem temporären Zahlungscode. Während der Anmeldung erhält die Kontoanfrage einen zufälligen Zahlungscode; wenn die Zahlung mit diesem Code eingeht, können wir das gewünschte Postfach erstellen, ohne nach echtem Namen, Wiederherstellungs-E-Mail, Telefonnummer oder dauerhaftem Abrechnungskonto zu fragen. Der Code existiert, um eine eingehende Zahlung einer ausstehenden Kontoerstellung zuzuordnen, nicht um als dauerhafte Kundenkennung zu bleiben.

Für norwegische Verbraucherverträge im Fernabsatz gibt [Angrerettloven § 20](https://lovdata.no/lov/2014-06-20-27/%C2%A720 "Angrerettloven § 20") Verbrauchern ein Widerrufsrecht, [§ 21](https://lovdata.no/lov/2014-06-20-27/%C2%A721 "§ 21") setzt die gewöhnliche Widerrufsfrist für Dienstleistungsverträge auf 14 Tage ab Vertragsschluss, und [§ 26](https://lovdata.no/lov/2014-06-20-27/%C2%A726 "§ 26") regelt die Pflichten des Verbrauchers, wenn der Widerruf genutzt wird, nachdem die Leistung begonnen hat. Während dieser Zeit kann ein Kunde kündigen, indem er uns kontaktiert und den Zahlungscode angibt, weil dieser Code die Zahlung-Konto-Zuordnung auffindbar macht, die wir für die Erstattung brauchen. Wir behalten diese Zahlungscode-zu-Konto-Zuordnung 14 Tage nach Aktivierung des Kontos; das kann in manchen Fällen etwas länger als das strikte Minimum sein. Danach löschen wir sie.

Ab diesem Zeitpunkt haben wir noch eine Liste erhaltener Zahlungen und eine Liste aktiver Konten, aber unsere Systeme können nicht mehr beantworten, welche Zahlung welches Mailbox-Konto erstellt hat. Dieses Löschmodell ist mit Abos nicht möglich, weil der Anbieter genug Abrechnungszustand behalten muss, um das Abo später zu belasten, zu verlängern, fehlschlagen zu lassen, zu kündigen oder zu erstatten.

### Der echte Preis, das Richtige zu tun

Der geschäftliche Preis ist simpel: Abos würden Privacy.Fish wahrscheinlich profitabler machen. Wir haben trotzdem eine Einmalzahlung gewählt, weil wiederkehrende Abrechnung genau die langlebige Konto-Zahlungs-Verbindung schaffen würde, die wir vermeiden wollen, um deine Privatsphäre zu schützen.

Das heißt, der Dienst hängt stärker von Vertrauen, Weiterempfehlungen und Menschen ab, die verstehen, warum dieser Tradeoff wichtig ist. Wenn Privacy.Fish für dich nützlich ist, hilft eine Empfehlung an die richtigen Menschen mehr als jedes Abo-Upsell.

## Fazit: Rechtsordnung ist Teil von E-Mail-Privatsphäre

E-Mail-Privatsphäre ist teilweise ein technisches Problem, aber auch ein Rechtsordnungsproblem: Das Land hinter dem Anbieter entscheidet, was der Anbieter speichern, offenlegen oder für Zugriff bauen muss. Deutschland, die Schweiz, das Vereinigte Königreich, Australien und die Vereinigten Staaten zeigen, wie schnell „private E-Mail“ zu einer Frage von Metadatenspeicherung, Sicherungspflichten, Abfangbereitschaft oder wiederkehrender Abrechnungsidentität wird. Norwegen gibt Privacy.Fish eine engere rechtliche Grundlage, und wir nutzen das Produktdesign, um die verbleibende Datenoberfläche so klein wie möglich zu halten. Diese Entscheidung ist weniger bequem und geschäftlich weniger komfortabel als ein normales Mailbox-Abo, aber sie ist die einzige Richtung, die zu dem passt, was wir mit privater E-Mail meinen.

## FAQ: Norwegische E-Mail-Privatsphäre, Vorratsdatenspeicherung und Privacy.Fish

### Ist Norwegen in der EU?

Nein. Norwegen ist [kein EU-Mitgliedstaat](https://www.eu-norway.org/eu/norway-and-the-eu/ "kein EU-Mitgliedstaat"), nimmt aber über den EWR an großen Teilen des europäischen Rechts- und Wirtschaftsrahmens teil.

### Was muss Privacy.Fish nach norwegischem Recht standardmäßig protokollieren?

Für unser SSH-/SFTP-Kontozugangsmodell entspricht die routinemäßige Speicherpflicht nach [Ekomloven § 3-13](https://lovdata.no/lov/2024-12-13-76/%C2%A73-13 "Ekomloven § 3-13") Zeitstempel, Benutzername, Quell-IP-Adresse und Quell-Port für registrierten Kontozugriff, gespeichert für 12 Monate.

### Bedeutet norwegische Rechtsordnung keinen rechtmäßigen Zugriff?

Nein. [Ekomloven § 3-12](https://lovdata.no/lov/2024-12-13-76/%C2%A73-12 "Ekomloven § 3-12") verlangt, dass Anbieter Dienste für gesetzlich vorgeschriebenen Zugriff vorbereiten, und [§ 3-14](https://lovdata.no/lov/2024-12-13-76/%C2%A73-14 "§ 3-14") regelt die Offenlegung gespeicherter IP-Adressdaten an Polizei oder Staatsanwaltschaft in definierten Fällen.

### Warum nutzt Privacy.Fish Einmalzahlungen?

Einmalzahlungen erlauben uns, die Zahlungscode-zu-Konto-Zuordnung nach dem 14-Tage-Fenster zu löschen. Abos würden laufenden Abrechnungszustand erfordern, der eine Zahlungsbeziehung mit einem Konto verbindet. Die längere Erklärung steht in unseren [Datenschutzinformationen](/documentation/privacy-information/ "Datenschutzinformationen").
