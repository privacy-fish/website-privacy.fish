---
title: "Rechtsordnung"
description: "Warum Norwegen für privacy.fish ausgewählt wurde und wie Rechtsordnung Protokollierung, Offenlegung und rechtliche Exposition beeinflusst."
weight: 70
video: "/videos/norways-privacy-jurisdiction.mp4"
---

## Einleitung

Privacy.Fish hat Rechtsordnungen anhand einer engen Frage verglichen: Was würde der Staat uns als E-Mail-Anbieter zu protokollieren zwingen?

Jedes verpflichtende Log wird zu Daten, die wir sammeln, schützen, speichern, sichern und potenziell offenlegen müssen. Norwegen war die beste Rechtsordnung, die wir für das Privacy.Fish-Modell gefunden haben, weil die verpflichtend gespeicherten Daten eng, explizit und vorhersehbar sind: Konto, Quell-IP:Port und Zeitstempel für 12 Monate.

Norwegen ist auch klar darüber, was diese Pflicht nicht umfassen soll. Unter der norwegischen IP-Speicherpflicht dürfen Zielinformationen nicht gespeichert werden. Das ist der wichtige Unterschied: Norwegen verlangt ein enges Zugriffsidentifikations-Log, keine breite Mail-Transaktionsprotokollierung.

Quellen:
- [Norwegian Electronic Communications Act, unofficial English translation by Nkom](https://nkom.no/english/guidelines-for-providers/_/attachment/download/cccdd933-8254-4024-820e-168a96f071f0%3Ac6fa2c909f0187b6c31b6f8c0f40f328b7ce46a8/Electronic%20Communications%20Act%20%28Ekom%20Act%29%20%E2%80%93%20Unofficial%20Translation.pdf "Norwegian Electronic Communications Act, unofficial English translation by Nkom")
- [Norwegian Electronic Communications Regulation, unofficial English translation by Nkom](https://nkom.no/english/guidelines-for-providers/_/attachment/download/13dc9675-b0f4-42e0-840e-dba0c735db69%3A470133b1352c9038804c5c0b0d7ed967c13c9e0f/Regulation%20on%20Electronical%20Communications%20Network%20and%20Services%20%E2%80%93%20Unofficial%20version.pdf "Norwegian Electronic Communications Regulation, unofficial English translation by Nkom")
- [Nordic data-retention comparison: Norway](https://pub.norden.org/temanord2024-532/9-norway.html "Nordic data-retention comparison: Norway")

## EU

Die EU hat derzeit kein einzelnes EU-weites Vorratsdatenspeicherungsgesetz, das jedem E-Mail-Anbieter genau sagt, was zu protokollieren ist.

Die alte EU-Vorratsdatenspeicherungsrichtlinie verlangte von Anbietern, Verkehrs- und Standortdaten für Dienste einschließlich Internetzugang, Internet-E-Mail und Internettelefonie zu speichern. 2014 erklärte der Gerichtshof der Europäischen Union diese Richtlinie für ungültig, weil sie zu stark in die Rechte auf Privatsphäre und Datenschutz eingriff.

Seitdem ist die Rechtslage in der EU fragmentiert. Jeder Mitgliedstaat hat eigene nationale Speicherregeln, Gerichtsentscheidungen, nationale Sicherheitsausnahmen und Telekommunikationsgesetze. Manche Länder haben derzeit keine allgemeinen Speicherregeln in Kraft, manche haben gezielte Speicherung, und manche legen weiterhin breite oder komplexe Metadaten-Speicherpflichten auf.

Für Privacy.Fish bedeutet das: Die EU selbst ist keine ausreichende Antwort. Jedes Land muss separat geprüft werden.

Warum Norwegen besser ist:
- Norwegen gibt eine enge und explizite Regel, um die herum wir designen können.
- EU-Länder unterscheiden sich stark.
- Manche EU-Länder haben breitere Speicherpflichten als Norwegen.
- Manche EU-Länder haben weniger pauschale Speicherung, aber mehr rechtliche Unsicherheit.
- Die EU-Situation ist instabil genug, dass „in der EU gehostet“ keine präzise Datenschutzinformation ist.

Quellen:
- [CJEU press release declaring the Data Retention Directive invalid](https://curia.europa.eu/jcms/upload/docs/application/pdf/2014-04/cp140054en.pdf "CJEU press release declaring the Data Retention Directive invalid")
- [European Commission page on data retention](https://home-affairs.ec.europa.eu/policies/internal-security/lawful-access-data/data-retention_en "European Commission page on data retention")
- [EU Data Retention Directive FAQ](https://europa.eu/rapid/press-release_MEMO-14-269_en.htm "EU Data Retention Directive FAQ")
- [Eurojust report on CJEU data-retention case law](https://www.eurojust.europa.eu/sites/default/files/assets/files/data-retention-report-cjeu-eurojust-13-11-2024.pdf "Eurojust report on CJEU data-retention case law")
- [Council document on future EU data-retention rules](https://data.consilium.europa.eu/doc/document/WK-11640-2025-INIT/en/pdf "Council document on future EU data-retention rules")
- [Cullen International update on national data-retention laws in Europe](https://www.cullen-international.com/news/2025/07/Update-on-national-data-retention-laws-in-Europe.html "Cullen International update on national data-retention laws in Europe")

## Norwegen

Norwegen ist der Gewinner für Privacy.Fish.

Norwegisches Recht verlangt von Anbietern, Informationen zu speichern, die benötigt werden, um zu identifizieren, welcher Teilnehmer zu einem bestimmten Zeitpunkt eine öffentliche IP-Adresse verwendet hat. Wo nötig, umfasst das die Portnummer. Im Privacy.Fish-Modell wird daraus ein enges Zugriffslog: Konto, Quell-IP:Port und Zeitstempel.

Die Speicherfrist beträgt 12 Monate.

Die norwegische Regel sagt außerdem, dass Zielinformationen nicht gespeichert werden dürfen. Das ist wichtig, weil es bedeutet, dass das Gesetz um die Identifizierung von Kontozugriff anhand einer IP-Adresse und Zeit geht, nicht um eine allgemeine Historie, wem du geschrieben hast, an welchen Server deine Mail zugestellt wurde oder welche Ziele du kontaktiert hast.

Warum Norwegen besser ist:
- Die verpflichtend gespeicherten Daten sind eng.
- Die Speicherfrist ist explizit.
- Zielprotokollierung ist unter dieser Pflicht ausdrücklich ausgeschlossen.
- Die Regel ist verständlich genug, um um sie herum zu designen.
- Sie verlangt für dieses Modell keine routinemäßige Protokollierung von Absender, Empfänger, Mailbox-Events, Message-ID, Zustellstatus oder Zielservern.

Quellen:
- [Norwegian Electronic Communications Act, unofficial English translation by Nkom](https://nkom.no/english/guidelines-for-providers/_/attachment/download/cccdd933-8254-4024-820e-168a96f071f0%3Ac6fa2c909f0187b6c31b6f8c0f40f328b7ce46a8/Electronic%20Communications%20Act%20%28Ekom%20Act%29%20%E2%80%93%20Unofficial%20Translation.pdf "Norwegian Electronic Communications Act, unofficial English translation by Nkom")
- [Norwegian Electronic Communications Regulation, unofficial English translation by Nkom](https://nkom.no/english/guidelines-for-providers/_/attachment/download/13dc9675-b0f4-42e0-840e-dba0c735db69%3A470133b1352c9038804c5c0b0d7ed967c13c9e0f/Regulation%20on%20Electronical%20Communications%20Network%20and%20Services%20%E2%80%93%20Unofficial%20version.pdf "Norwegian Electronic Communications Regulation, unofficial English translation by Nkom")
- [Nordic data-retention comparison: Norway](https://pub.norden.org/temanord2024-532/9-norway.html "Nordic data-retention comparison: Norway")

## Schweiz

Die Schweiz ist für das Privacy.Fish-Modell schlechter als Norwegen.

Die Schweiz hat einen starken Datenschutzruf, aber ihr Überwachungs- und Metadaten-Speicherrahmen kann deutlich breitere E-Mail-Metadaten erfassen als Norwegen. Schweizer Recht verlangt seit langem, dass bestimmte Telekommunikationsmetadaten sechs Monate gespeichert werden. Offizielle Schweizer Materialien beschreiben gespeicherte Metadaten allgemein als Informationen darüber, wer mit wem kommuniziert hat, wann, wie lange, wo und mit welchen technischen Mitteln.

Für E-Mail-bezogene Dienste können Schweizer Überwachungsregeln je nach Anbieterklassifikation und konkreter Verpflichtung Metadaten wie Absender- und Empfängeradressen, Protokoll, Mailbox-Events, Zustellstatus und Informationen zu sendenden oder empfangenden Mailservern umfassen.

Die Schweiz ist außerdem rechtlich komplexer. Manche reinen E-Mail- oder Over-the-Top-Kommunikationsdienste können geringere Pflichten haben als klassische Telekommunikationsanbieter. Diese Komplexität ist aber selbst ein Nachteil gegenüber Norwegens enger und expliziter Regel.

Warum Norwegen besser ist:
- Norwegen verlangt Konto, Quell-IP:Port und Zeitstempel.
- Die Schweiz kann breitere E-Mail-Metadaten verlangen.
- Norwegen schließt Zielinformationen unter der IP-Speicherpflicht ausdrücklich aus.
- Die Schweizer Regeln schaffen mehr Klassifikations- und Compliance-Unsicherheit.
- Die Schweiz ist schwerer ehrlich als einfache datenminimierende Rechtsordnung für dieses Dienstmodell zu erklären.

Quellen:
- [Swiss PTSS FAQ on metadata retention](https://www.li.admin.ch/en/documentation/faq "Swiss PTSS FAQ on metadata retention")
- [Swiss BÜPF information page](https://www.li.admin.ch/de/themen/das-buepf "Swiss BÜPF information page")
- [Swiss VÜPF ordinance text on Fedlex](https://www.fedlex.admin.ch/eli/cc/2018/32/de "Swiss VÜPF ordinance text on Fedlex")
- [Swiss ordinance text via LexFind](https://www.lexfind.ch/tolv/149472/de "Swiss ordinance text via LexFind")
- [Laux Lawyers analysis on Swiss surveillance duties for messaging services](https://www.lauxlawyers.ch/ueberwachung-von-instant-messaging-bundesgericht-schafft-klarheit/ "Laux Lawyers analysis on Swiss surveillance duties for messaging services")
- [Swico guide on authority requests and surveillance orders](https://www.swico.ch/media/filer_public/78/f1/78f151c6-ff8b-4ecf-b856-c1530e1a3fd5/leitfaden_behordenanfragen_3_anordnung_zur_fernmeldeuberwachung.pdf "Swico guide on authority requests and surveillance orders")
- [Internet Society Switzerland criticism of Swiss surveillance ordinance](https://www.isoc.ch/swiss-surveillance-ordinance-encryption-threat-vupf-oscpt/ "Internet Society Switzerland criticism of Swiss surveillance ordinance")
- [Digitale Gesellschaft: data retention in Switzerland](https://www.digitale-gesellschaft.ch/vds-suisse/index_en.html "Digitale Gesellschaft: data retention in Switzerland")

## Deutschland

Deutschland ist für das Privacy.Fish-Modell schlechter als Norwegen.

Deutschlands Hauptproblem ist nicht nur Vorratsdatenspeicherung. Das größere Problem ist Infrastruktur für rechtmäßige Überwachung.

Die deutsche Regulierungsbehörde BNetzA sagt, Betreiber von E-Mail-Systemen seien nur so lange von der TKÜV-Bereitschaftspflicht ausgenommen, wie nicht mehr als 100.000 Nutzer an das System angeschlossen sind. Oberhalb dieser Schwelle müssen E-Mail-Betreiber technische und organisatorische Bereitschaft zur rechtmäßigen Überwachung vorhalten.

Deutschland hat außerdem technische Regeln für Übergabepunkte und vollständige Aufzeichnung rechtmäßiger Überwachungen nach TR TKÜV. Die BNetzA beschreibt TR TKÜV als technische Richtlinie, die Anforderungen an vollständige Aufzeichnungen von Telekommunikationsüberwachungen, Auskunftserteilung und Konfiguration von Übergabepunkten zu berechtigten Stellen definiert.

Deutschland hat außerdem spezifische Anforderungen an elektronische Behördenschnittstellen. Die BNetzA sagt, verpflichtete Anbieter mit 100.000 oder mehr Vertragspartnern müssen sowohl eine ETSI-basierte Schnittstelle als auch ein E-Mail-basiertes Übermittlungsverfahren für bestimmte Auskunftsverfahren vorhalten, während kleinere verpflichtete Anbieter das E-Mail-basierte Verfahren bereithalten müssen.

Warum Norwegen besser ist:
- Norwegen verlangt ein enges Zugriffslog.
- Deutschland kann große E-Mail-Anbieter zu Überwachungsbereitschaft verpflichten.
- Deutschland hat Übergabe-/Schnittstellenanforderungen, um die Privacy.Fish nicht herum designen will.
- Deutschlands Speicherregeln waren wiederholt politisch und rechtlich umstritten.
- Norwegen ist einfacher: Konto, Quell-IP:Port, Zeitstempel.

Quellen:
- [BNetzA: email surveillance obligations and 100,000-user threshold](https://www.bundesnetzagentur.de/DE/Fachthemen/Telekommunikation/OeffentlicheSicherheit/Ueberwachung_Auskunftsert/EMailUeberwachung/node.html "BNetzA: email surveillance obligations and 100,000-user threshold")
- [BNetzA: intercepts and provision of information](https://www.bundesnetzagentur.de/EN/Areas/Telecommunications/ServicerProviderObligation/PublicSafety/Intercepts/start.html "BNetzA: intercepts and provision of information")
- [BNetzA: TR TKÜV supplementary information](https://www.bundesnetzagentur.de/EN/Areas/Telecommunications/ServicerProviderObligation/PublicSafety/Intercepts/SupplementaryInformation/SupplementaryInformation_node.html "BNetzA: TR TKÜV supplementary information")
- [BNetzA: electronic authority interface thresholds](https://www.bundesnetzagentur.de/DE/Fachthemen/Telekommunikation/OeffentlicheSicherheit/Ueberwachung_Auskunftsert/BDA/artikel.html "BNetzA: electronic authority interface thresholds")
- [German TR TKÜV draft text via EU TRIS](https://technical-regulation-information-system.ec.europa.eu/de/notification/26192/text/D/EN "German TR TKÜV draft text via EU TRIS")
- [Statewatch note on German interception obligations for email providers](https://www.statewatch.org/statewatch-database/germany-telecommunication-providers-to-intercept-customers/ "Statewatch note on German interception obligations for email providers")
- [Lexology note on German classification of internet-based email services](https://www.lexology.com/library/detail.aspx?g=62df0a70-3b37-4a27-a7e7-dcf41ba2e469 "Lexology note on German classification of internet-based email services")

## Niederlande

Die Niederlande sind kein einfacher „schlechter als Norwegen“-Fall. Sie sind komplizierter.

Das alte niederländische pauschale Vorratsdatenspeicherungsgesetz wurde 2015 vom Bezirksgericht Den Haag ausgesetzt, nachdem der EuGH die EU-Vorratsdatenspeicherungsrichtlinie für ungültig erklärt hatte. Dadurch mussten niederländische Telekommunikations- und Internetanbieter nach diesem Gesetz keine allgemeinen Telefon- und Internetverkehrsdaten mehr für feste Zeiträume speichern. Das alte niederländische Regime hatte Telekommunikationsdatenspeicherung für mindestens 6 Monate verlangt, und Kommentare beschreiben die alte Pflicht als 12 Monate für Telefoniedaten und 6 Monate für Internetzugang und E-Mail-bezogene Daten.

Das kann die Niederlande attraktiv aussehen lassen, wenn die einzige Frage pauschale Metadatenspeicherung ist. Aber das Fehlen des alten pauschalen Speichergesetzes bedeutet nicht „keine Überwachungspflichten“. Niederländische öffentliche Telekommunikationsanbieter können weiterhin rechtmäßigen Überwachungs- und Offenlegungspflichten unterliegen. NBIP beschreibt Lawful Interception und Lawful Disclosure als Pflichten für ISPs, VoIP-Anbieter und andere öffentliche Telekommunikationsdienste, bei denen berechtigte Behörden Verkehrsmitschnitt oder kundenidentifizierende Daten verlangen können.

Das ist der entscheidende Unterschied zu Norwegen für Privacy.Fish: Die Niederlande können derzeit weniger einfache pauschale Speicherung haben, aber das rechtliche Modell ist weniger sauber zu erklären. Norwegen gibt uns eine enge, explizite, öffentliche Regel: Konto, Quell-IP:Port und Zeitstempel für 12 Monate. Die Niederlande erfordern eine tiefere Klassifikationsfrage: ob der Dienst als öffentlicher elektronischer Kommunikationsdienst behandelt wird, welche Interception- oder Disclosure-Pflichten gelten und welche operative Bereitschaft erwartet wird.

Warum Norwegen besser ist:
- Norwegen gibt eine enge, explizite Speicherpflicht, um die herum wir designen können.
- Die Niederlande haben derzeit nicht mehr dasselbe alte pauschale Speichergesetz, aber rechtmäßige Überwachungs- und Offenlegungspflichten existieren weiterhin.
- Die niederländische Antwort hängt stärker von Anbieterklassifikation und operativer Interpretation ab.
- Norwegen ist ehrlicher einfach zu erklären: ein erforderliches Zugriffslog, kein breiteres oder unklareres Mail-Metadatenregime.

Quellen:
- [Dutch court invalidated the national data-retention law](https://peacepalacelibrary.nl/blog/2015/data-retention-saga-dutch-court-declared-national-data-retention-law-invalid "Dutch court invalidated the national data-retention law")
- [Privacy First summary of the telecom retention case](https://privacyfirst.nl/en/court-cases/telecom-retention-obligation/ "Privacy First summary of the telecom retention case")
- [Lexology summary of Dutch data-retention suspension](https://www.lexology.com/library/detail.aspx?g=0860c791-fc73-435a-a243-1701a03a7265 "Lexology summary of Dutch data-retention suspension")
- [NBIP lawful interception explanation](https://www.nbip.nl/en/lawful-interception/ "NBIP lawful interception explanation")
- [Dutch government business page for telecom-provider requirements](https://business.gov.nl/regulations/requirements-telecom-providers/ "Dutch government business page for telecom-provider requirements")
- [PILP dossier on the Telecommunications Retention Act](https://pilp.nu/en/dossier/the-telecommunications-retention-act/ "PILP dossier on the Telecommunications Retention Act")

## Schweden

Schweden ist für das Privacy.Fish-Modell schlechter oder weniger sauber als Norwegen.

Schwedische Speicherregeln sind breiter und stärker kategorieabhängig. Der nordische Vergleich beschreibt schwedische Speicherung rund um Daten, die benötigt werden, um Quelle und Ziel, Datum, Zeit, Dauer, Kommunikationsart, Benutzergerät und Standort für erfasste Dienste zu verfolgen und zu identifizieren.

Für Internetzugang speichert Schweden Daten wie Teilnehmer- oder Benutzerdaten, zugewiesene IP-Adresse, Login- und Logout-Zeiten und Gerätedaten. Speicherfristen unterscheiden sich je nach Kategorie. Daten zum Internetzugang werden 10 Monate gespeichert, während Daten zu Telefonie oder Nachrichtenverarbeitung über einen mobilen Endpunkt generell 6 Monate gespeichert werden. Bestimmte Standortdaten werden kürzer gespeichert.

Das ist weniger geeignet für Privacy.Fish, weil die gespeicherten Kategorien breiter und stärker von Dienstklassifikation abhängig sind als Norwegens Konto-, Quell-IP:Port-, Zeitstempel-Modell.

Warum Norwegen besser ist:
- Norwegens gespeicherte Felder sind enger.
- Schwedens Regeln umfassen breitere Kategorien.
- Schwedens Speicherfristen variieren nach Kommunikationsart.
- Norwegen schließt Zielinformationen unter der IP-Speicherpflicht ausdrücklich aus.
- Norwegen ist einfacher als privacy-first Rechtsordnung zu erklären.

Quellen:
- [Nordic data-retention comparison: Sweden](https://pub.norden.org/temanord2024-532/10-sweden.html "Nordic data-retention comparison: Sweden")
- [Nordic data-retention comparison: Nordic overview](https://pub.norden.org/temanord2024-532/5-nordic-overview.html "Nordic data-retention comparison: Nordic overview")

## Dänemark

Dänemark ist rechtlich schwerer und komplexer als Norwegen.

Dänemark hatte historisch breite Vorratsdatenspeicherung. Nach Rechtsänderungen und EuGH-Druck bewegte sich Dänemark zu einem komplizierteren Modell mit gezielter Speicherung, regionaler Speicherung und allgemeiner oder undifferenzierter Speicherung in bestimmten nationalen Sicherheitskontexten.

Der nordische Vergleich beschreibt gezielte Speicherung, die auf verurteilte Personen, Kommunikationsausrüstung, Personen, geografische Gebiete und konkrete Bewertungen zielt. Dänemark hat außerdem Regeln zu allgemeiner Speicherung für nationale Sicherheit und Internetzugang. Speicherfristen können je nach aktivierter Regel etwa ein Jahr betragen.

Das passt schlecht zu Privacy.Fish, weil die Antwort keine einfache Feldliste ist. Das dänische Regime hängt von Targeting, Geografie, Straftatkategorien, nationalen Sicherheitsbewertungen und Umsetzungsentscheidungen ab.

Warum Norwegen besser ist:
- Norwegen hat eine enge allgemeine Regel, um die herum Privacy.Fish designen kann.
- Dänemarks Regime ist komplexer.
- Dänemark kann gezielte, regionale und nationale Sicherheits-Speichermodelle umfassen.
- Norwegen gibt Benutzern eine klarere Antwort darauf, was wir speichern müssen.

Quellen:
- [Nordic data-retention comparison: Denmark](https://pub.norden.org/temanord2024-532/6-denmark.html "Nordic data-retention comparison: Denmark")
- [IT-Pol: the new Danish data-retention law](https://itpol.dk/articles/new-Danish-data-retention-law-2022 "IT-Pol: the new Danish data-retention law")
- [European Parliament document on Danish data retention](https://www.europarl.europa.eu/meetdocs/2014_2019/plmrep/COMMITTEES/LIBE/DV/2020/01-20/Danish_data_retention_EN.pdf "European Parliament document on Danish data retention")
- [Eurojust report on CJEU data-retention case law](https://www.eurojust.europa.eu/sites/default/files/assets/files/data-retention-report-cjeu-eurojust-13-11-2024.pdf "Eurojust report on CJEU data-retention case law")

## Finnland

Finnland kann für manche kleinen Anbieter leichter sein, ist aber weniger sauber als Norwegen.

Finnlands Speicherregeln nach FECA gelten für ausgewählte Anbietergruppen. Das Innenministerium entscheidet, welche Unternehmen der Speicherpflicht unterliegen. Das bedeutet, dass Finnland möglicherweise nicht jeden kleinen Anbieter praktisch erfasst, aber die Regel ist weniger vorhersehbar und administrativer als Norwegens enge gesetzliche Regel.

Wenn ein Anbieter erfasst ist, können die gespeicherten Kategorien breiter sein als Norwegens Zugriffslog. Der nordische Vergleich beschreibt Kategorien wie Name und Adresse des Teilnehmers, Anschlusskennungen, transaktionsidentifizierende Daten, Nachrichtentyp, Empfänger, Zeit und Dauer sowie Benutzer- oder Gerätekennungen für Internetzugang.

Speicherfristen variieren nach Dienstkategorie: 12 Monate für Mobilfunk- und SMS-Daten, 6 Monate für Internettelefonie und 9 Monate für Internetzugang.

Warum Norwegen besser ist:
- Finnland erfasst möglicherweise nicht jeden Anbieter, aber die Abdeckungsfrage ist weniger vorhersehbar.
- Falls erfasst, sind Finnlands gespeicherte Kategorien breiter.
- Finnland nutzt unterschiedliche Speicherfristen nach Diensttyp.
- Norwegen gibt eine klarere Regel für dieses Modell.

Quellen:
- [Nordic data-retention comparison: Finland](https://pub.norden.org/temanord2024-532/7-finland.html "Nordic data-retention comparison: Finland")
- [Nordic data-retention comparison: Nordic overview](https://pub.norden.org/temanord2024-532/5-nordic-overview.html "Nordic data-retention comparison: Nordic overview")

## Island

Island ist für das Privacy.Fish-Modell schlechter als Norwegen.

Island hat eine kürzere Speicherfrist als Norwegen, aber der Umfang erscheint breiter. Isländische Telekommunikationsunternehmen müssen für 6 Monate eine Mindestaufzeichnung des elektronischen Kommunikationsverkehrs von Benutzern speichern.

Der nordische Vergleich sagt, Anbieter müssen identifizieren können, welcher Kunde eine Telefonnummer, IP-Adresse oder einen Benutzernamen verwendet hat, und Informationen zu Verbindungen, Daten, wer verbunden war und übertragener Datenmenge bereitstellen.

Das ist breiter als Norwegens Modell. Norwegen verlangt Konto, Quell-IP:Port und Zeitstempel; Island scheint breitere Verbindungsaufzeichnungen zu verlangen.

Warum Norwegen besser ist:
- Islands Speicherfrist ist kürzer, aber der Umfang erscheint breiter.
- Island kann Informationen über Verbindungen und Gegenparteien umfassen.
- Norwegens Pflicht ist enger.
- Norwegen schließt Zielinformationen unter der IP-Speicherpflicht ausdrücklich aus.

Quellen:
- [Nordic data-retention comparison: Iceland](https://pub.norden.org/temanord2024-532/8-iceland.html "Nordic data-retention comparison: Iceland")
- [Iceland Electronic Communications Act No. 70/2022](https://fjarskiptastofa.is/library?itemid=175f315e-40a7-4b2e-83b0-4c377b96b098 "Iceland Electronic Communications Act No. 70/2022")
- [Freedom House note on Icelandic data retention](https://freedomhouse.org/country/iceland/freedom-net/2024 "Freedom House note on Icelandic data retention")

## Vereinigtes Königreich

Das Vereinigte Königreich ist für das Privacy.Fish-Modell schlechter als Norwegen.

Das UK kann Datenspeicherung per Notice verlangen. Eine Data Retention Notice nach dem Investigatory Powers Act 2016 kann die Speicherung relevanter Kommunikationsdaten für bis zu 12 Monate verlangen.

Die Kategorien können viel breiter sein als Norwegens enges Zugriffslog. UK-Material beschreibt relevante Kommunikationsdaten als Daten, die Sender, Empfänger, Quellen, Ziele, IP-Adressen, Ports, Zeiten, Dauern, Kommunikationsart, Methode, Muster, Systemdaten und standortartige Informationen identifizieren oder bei deren Identifizierung helfen können.

Das UK hat außerdem Geheimhaltungs- und Notice-basierte Befugnisse. Das passt schlecht zu einem transparenten privacy-first E-Mail-Anbieter, weil die Verpflichtung möglicherweise keine einfache öffentliche Regel ist, die Benutzer verstehen können.

Warum Norwegen besser ist:
- Norwegens Pflicht ist öffentlich, eng und explizit.
- Das UK kann breitere Speicherpflichten per Notice auferlegen.
- Das UK-System kann Geheimhaltung umfassen.
- UK-Kategorien können Quell- und Zielinformationen einschließen.
- Norwegen ist einfacher ehrlich zu erklären.

Quellen:
- [UK Notices Regime Code of Practice](https://www.gov.uk/government/publications/notices-regime-code-of-practice/notices-regime-code-of-practice-accessible "UK Notices Regime Code of Practice")
- [UK Investigatory Powers Act 2016, section 87](https://www.legislation.gov.uk/ukpga/2016/25/section/87 "UK Investigatory Powers Act 2016, section 87")
- [UK Communications Data Code of Practice, 2025 PDF](https://assets.publishing.service.gov.uk/media/68419a3d58bd5a53d0211cf7/Communications_Data_Code_of_Practice_-_June_2025.pdf "UK Communications Data Code of Practice, 2025 PDF")
- [UK Communications Data Code of Practice, older PDF](https://assets.publishing.service.gov.uk/media/5bf43801ed915d183c9c53ba/Communications_Data_Code_of_Practice.pdf "UK Communications Data Code of Practice, older PDF")

## Vereinigte Staaten

Die Vereinigten Staaten passen nicht gut, auch wenn sie keine einfache norwegenartige pauschale Speicherregel für E-Mail-Anbieter haben.

Der US Stored Communications Act erlaubt staatlichen Stellen, mit rechtlichem Verfahren die Offenlegung gespeicherter Inhalte und nicht-inhaltlicher Records zu erzwingen. Er erlaubt außerdem Preservation Requests, die einen Anbieter verpflichten, Records und andere Beweismittel, die bereits in seinem Besitz sind, für 90 Tage zu sichern, verlängerbar um weitere 90 Tage.

Das bedeutet, die USA können leichter aussehen, wenn die einzige Frage lautet: „Gibt es eine stehende universelle Zugriffslog-Regel?“ Für Privacy.Fish macht das breitere Umfeld aus erzwungener Offenlegung, Sicherung, Gag Orders und Überwachung die USA aber weniger attraktiv.

Das US-Modell dreht sich weniger um ein öffentliches Pflichtlog und mehr um einen breiten Rechtsprozessrahmen, der Offenlegung oder Sicherung dessen erzwingen kann, was der Anbieter hat.

Warum Norwegen besser ist:
- Norwegens Regel ist eng und explizit.
- Die USA haben breite Befugnisse zu erzwungener Offenlegung und Sicherung.
- Das US-Rechtsumfeld ist für einen privacy-first E-Mail-Anbieter weniger vorhersehbar.
- Norwegen passt besser zu einem transparenten öffentlichen Versprechen darüber, was routinemäßig gespeichert wird.

Quellen:
- [18 U.S.C. § 2703, Required disclosure and preservation under the Stored Communications Act](https://www.law.cornell.edu/uscode/text/18/2703 "18 U.S.C. § 2703, Required disclosure and preservation under the Stored Communications Act")
- [18 U.S.C. § 2704, Backup preservation](https://www.law.cornell.edu/uscode/text/18/2704 "18 U.S.C. § 2704, Backup preservation")
- [47 U.S.C. § 1002, CALEA assistance capability requirements](https://www.law.cornell.edu/uscode/text/47/1002 "47 U.S.C. § 1002, CALEA assistance capability requirements")
- [ACLU article on email preservation requests](https://www.aclu.org/news/privacy-technology/government-cannot-force-e-mail-companies-copy-and-save-your "ACLU article on email preservation requests")
- [EFF overview of mandatory data retention proposals](https://www.eff.org/issues/mandatory-data-retention "EFF overview of mandatory data retention proposals")

## Kanada

Kanada ist nicht so vorhersehbar wie Norwegen.

Kanada hat nach den hier geprüften Quellen derzeit nicht dieselbe einfache, stehende, norwegenartige Zugriffslog-Regel für E-Mail-Anbieter. Allerdings bewegt sich kanadisches Lawful-Access-Recht in Richtung stärkerer Anbieterpflichten.

Bill C-22 schlägt Befugnisse vor, die Regulierung für die Speicherung von Metadaten einschließlich Übertragungsdaten für angemessene Zeiträume von höchstens einem Jahr erlauben. Die kanadische Regierung beschreibt dies als Metadatenspeicherung und sagt, sie umfasse keine Inhalte, Browserhistorie oder Social-Media-Aktivität.

Für Privacy.Fish ist das weniger attraktiv als Norwegen, weil das kanadische Regime aktiv, politisch umstritten und auf breitere Metadatenbefugnisse zubewegt ist.

Warum Norwegen besser ist:
- Norwegen hat bereits eine enge, bekannte Regel.
- Kanadas Regime ist politisch aktiv und entwickelt sich.
- Vorgeschlagene kanadische Metadaten-Speicherbefugnisse könnten breiter werden als Norwegens Zugriffslog-Modell.
- Privacy.Fish profitiert von einer Rechtsordnung, in der die verpflichtend gespeicherten Daten bereits klar sind.

Quellen:
- [Canadian Bill C-22 text](https://www.parl.ca/DocumentViewer/en/45-1/bill/C-22/first-reading "Canadian Bill C-22 text")
- [Public Safety Canada backgrounder on Bill C-22 metadata retention](https://www.canada.ca/en/public-safety-canada/news/2026/03/backgrounder--securing-access-to-information-in-bill-c-22.html "Public Safety Canada backgrounder on Bill C-22 metadata retention")
- [Canadian government note explaining preservation orders are not general data retention](https://www.canada.ca/en/news/archive/2013/11/modernizing-criminal-code.html "Canadian government note explaining preservation orders are not general data retention")
- [Michael Geist analysis of Bill C-22 metadata retention](https://www.michaelgeist.ca/2026/03/the-lawful-access-privacy-risks-unpacking-bill-c-22s-expansive-metadata-retention-requirements/ "Michael Geist analysis of Bill C-22 metadata retention")

## Australien

Australien ist viel schlechter als Norwegen.

Australien hat eines der klarsten Bad-Fit-Regime für Privacy.Fish. Telekommunikationsanbieter müssen einen vorgeschriebenen Metadatensatz mindestens 2 Jahre speichern.

Der gespeicherte Datensatz kann Teilnehmer-/Konto-/Dienst-/Geräteinformationen, Quellenkennungen wie E-Mail-Adresse, IP-Adresse und Portnummer, Zielkennungen wie Empfänger-E-Mail-Adresse, IP-Adresse und Telefonkennungen, Zeit und Dauer, Kommunikationsart, verwendeten Dienst und Standortinformationen für Geräte umfassen.

Australische Leitlinien sagen außerdem, dass ein Anbieter, der einen E-Mail-Dienst anbietet, Data-Retention-Pflichten hat, sofern der E-Mail-Dienst nicht nur für den unmittelbaren Kreis einer Person ist.

Warum Norwegen besser ist:
- Norwegen verlangt ein enges Zugriffslog für 12 Monate.
- Australien verlangt deutlich breitere Metadaten für mindestens 2 Jahre.
- Australien umfasst Zielkennungen.
- Norwegen schließt Zielinformationen unter der norwegischen IP-Speicherpflicht ausdrücklich aus.
- Australien ist das Gegenteil dessen, was Privacy.Fish bauen will.

Quellen:
- [Australian Home Affairs data-retention obligations](https://www.homeaffairs.gov.au/about-us/our-portfolios/national-security/lawful-access-telecommunications/data-retention-obligations "Australian Home Affairs data-retention obligations")
- [Australian prescribed data set](https://www.homeaffairs.gov.au/nat-security/files/dataset.pdf "Australian prescribed data set")
- [Australian data-retention industry FAQ](https://www.homeaffairs.gov.au/nat-security/files/data-retention-industry-faqs.pdf "Australian data-retention industry FAQ")
- [Australian data-retention guidelines for service providers](https://www.homeaffairs.gov.au/nat-security/files/data-retention-guidelines-service-providers.pdf "Australian data-retention guidelines for service providers")
- [Australian Parliamentary report on data-retention period](https://www.aph.gov.au/-/media/02_Parliamentary_Business/24_Committees/244_Joint_Committees/PJCIS/DataRetention2014/Chapter4.pdf "Australian Parliamentary report on data-retention period")

## Neuseeland

Neuseeland ist für dieses Modell nicht besser als Norwegen.

Neuseeland scheint in den hier geprüften Quellen keine einfache pauschale Speicherpflicht für E-Mail-Anbieter wie Australien zu haben. Allerdings legt Neuseelands TICSA-Rahmen Telekommunikationsnetzbetreibern und Diensteanbietern Pflichten zur Abfangfähigkeit und Netzwerksicherheit auf.

TICSA dreht sich um Abfangfähigkeit und Netzwerksicherheit. Für Privacy.Fish ist das Problem nicht nur routinemäßiges Logging. Es geht auch darum, ob das rechtliche Umfeld erwartet, dass Anbieter Infrastruktur um Abfangfähigkeit herum designen.

Warum Norwegen besser ist:
- Norwegen verlangt ein enges gespeichertes Zugriffslog.
- Neuseeland hat Pflichten zur Abfangfähigkeit.
- Abfangbereitschaftspflichten passen schlechter zur Sicherheitsphilosophie von Privacy.Fish.
- Norwegen ist einfacher als datenminimierende Rechtsordnung zu erklären.

Quellen:
- [New Zealand NCSC: About TICSA](https://www.ncsc.govt.nz/what-we-do/regulations-and-standards/ticsa-network-security/about-ticsa/ "New Zealand NCSC: About TICSA")
- [New Zealand Police TICSA overview](https://www.police.govt.nz/advice/businesses-and-organisations/ticsa "New Zealand Police TICSA overview")
- [New Zealand TICSA legislation](https://www.legislation.govt.nz/act/public/2013/0091/latest/DLM5177923.html "New Zealand TICSA legislation")
- [New Zealand Interception Guide for TICSA](https://www.tcf.org.nz/industry-hub/industry-codes/interception-guide-for-ticsa "New Zealand Interception Guide for TICSA")
- [New Zealand Privacy Principle 9: retention of personal information](https://www.privacy.org.nz/privacy-principles/9/ "New Zealand Privacy Principle 9: retention of personal information")

## Frankreich

Frankreich ist schlechter als Norwegen.

Frankreich hat breitere Verkehrs- und Standortdaten-Speichermechanismen, besonders in nationalen Sicherheitskontexten. Französisches elektronisches Kommunikationsrecht und juristische Analyse beschreiben Pflichten für Betreiber elektronischer Kommunikationsdienste, Verbindungs- oder Verkehrsdaten, oft für ein Jahr, zu speichern, um auf Strafverfolgungs- oder nationale Sicherheitsbedürfnisse zu reagieren.

Französisches Recht hat außerdem umfangreiche nationale Sicherheits- und Lawful-Access-Maschinerie rund um elektronische Kommunikation. Auch wenn Inhalte nicht das gewöhnliche Ziel von Metadatenspeicherung sind, sind die Metadatenkategorien und der rechtliche Kontext breiter als Norwegens enges Zugriffsidentifikationslog.

Warum Norwegen besser ist:
- Norwegens gespeicherte Daten sind enger.
- Norwegens Regel ist expliziter und einfacher zu designen.
- Frankreichs nationales Sicherheits-Speichermodell ist breiter.
- Frankreich ist weniger geeignet für einen minimalen E-Mail-Anbieter, der nur ein enges Zugriffslog speichern will.

Quellen:
- [Global Network Initiative: France country legal framework](https://clfr.globalnetworkinitiative.org/country/france/ "Global Network Initiative: France country legal framework")
- [Lexology overview of telecom and internet access in France](https://www.lexology.com/library/detail.aspx?g=b74a9239-bb45-4354-a6a1-821884f1ca17 "Lexology overview of telecom and internet access in France")
- [ICLG France telecoms, media and internet overview](https://iclg.com/practice-areas/telecoms-media-and-internet-laws-and-regulations/france "ICLG France telecoms, media and internet overview")
- [Legal 500 France TMT PDF](https://www.legal500.com/guides/wp-content/uploads/sites/1/2025/08/France-TMT.pdf "Legal 500 France TMT PDF")
- [ISDC surveillance of telecommunications providers report](https://www.isdc.ch/media/2668/25-059-e-final-report-20251002_sign%C3%A9.pdf "ISDC surveillance of telecommunications providers report")
- [CJEU/French data-retention case note via eucrim](https://eucrim.eu/news/cjeu-french-legislation-on-data-retention-for-the-purpose-of-combating-market-abuse-offences-unlawful/ "CJEU/French data-retention case note via eucrim")

## Belgien

Belgien ist schlechter oder zumindest viel komplexer als Norwegen.

Belgien hat ein überarbeitetes gezieltes und differenziertes Metadaten-Speicherregime. BIPT beschreibt Artikel 107/5 des Electronic Communications Act von 2005 in der Fassung des Gesetzes vom 20. Juli 2022 über die Erhebung und Speicherung von Identifikationsdaten und Metadaten im elektronischen Kommunikationssektor und die Bereitstellung solcher Daten an Behörden.

Belgisches Vorratsdatenspeicherungsrecht wurde wiederholt angefochten und überarbeitet. Das schafft mehr rechtliche Komplexität als Norwegens enge Regel.

Das belgische System ist kein einfaches „ein Zugriffslog“-Modell. Es ist ein breiterer Metadaten-Speicher- und Beweisbereitstellungsrahmen mit mehreren beweglichen Teilen und wiederholten Gerichtsverfahren.

Warum Norwegen besser ist:
- Norwegen hat eine enge Zugriffslog-Pflicht.
- Belgien hat breitere Metadaten-Speichermaschinerie.
- Belgien hat wiederholte rechtliche Unsicherheit rund um Speicherregeln.
- Norwegen ist einfacher klar für Benutzer zu dokumentieren.

Quellen:
- [BIPT: provision of electronic evidence to authorities](https://www.bipt.be/operators/telecommunications/security/provision-of-electronic-evidence "BIPT: provision of electronic evidence to authorities")
- [EDRi/Liga letter on Belgian draft data-retention law](https://edri.org/wp-content/uploads/2022/06/EDRi-Liga-Letter-Draft-Law-Data-Retention-BE_EN.pdf "EDRi/Liga letter on Belgian draft data-retention law")
- [EuroISPA note on Belgian data-retention uncertainty](https://www.euroispa.org/2024/10/data-retention-rules-in-belgium-uncertainty-remains-after-third-constitutional-court-ruling/ "EuroISPA note on Belgian data-retention uncertainty")
- [NautaDutilh Belgium cybersecurity PDF mentioning 2022 data-retention act](https://www.nautadutilh.com/documents/32/2023_Lexology_GTDT_Cybersecurity_-_Belgium6212859.1_1.pdf "NautaDutilh Belgium cybersecurity PDF mentioning 2022 data-retention act")
- [Telenet Law Enforcement Disclosure Report 2024](https://www2.telenet.be/content/dam/www-telenet-corp/duurzaamheidsdocumenten-2025/Telenet%20group%20NV_Law%20Enforcement%20Disclosure%20Report_2024.pdf "Telenet Law Enforcement Disclosure Report 2024")

## Irland

Irland ist rechtlich komplex und historisch breiter als Norwegen.

Irlands Communications (Retention of Data) Act 2011 enthielt Kategorien für Internetzugang, Internet-E-Mail und Internettelefonie. Die gespeicherten Daten umfassten Daten, die erforderlich sind, um die Quelle einer Kommunikation zu verfolgen und zu identifizieren, etwa User IDs, sowie andere Felder zu Internetzugang, Internet-E-Mail und Internettelefonie.

Der irische Rahmen nutzte eine einjährige Speicherfrist für internetbezogene Daten und eine zweijährige Frist für Telefondaten. Der Inhalt von E-Mail-Nachrichten war nicht das Ziel; die Pflicht betraf verkehrsbezogene Informationen. Das ist aber weiterhin breiter als Norwegens enges Konto-, Quell-IP:Port-, Zeitstempel-Modell.

Irlands Rahmen wurde außerdem stark durch EU-Rechtsprechung und spätere Änderungen beeinflusst. Das macht ihn für Privacy.Fish weniger sauber als Norwegen.

Warum Norwegen besser ist:
- Norwegens verpflichtend gespeicherte Daten sind enger und klarer.
- Irlands Rahmen erfasste historisch Internet-E-Mail-Metadaten.
- Irlands Rahmen ist komplexer und rechtlich umstritten.
- Norwegen gibt eine einfache öffentliche Antwort darauf, was Privacy.Fish speichern muss.

Quellen:
- [Irish Communications Retention of Data Act 2011](https://www.irishstatutebook.ie/eli/2011/act/3/enacted/en/html "Irish Communications Retention of Data Act 2011")
- [Irish Communications Retention of Data Act 2011, Schedule 2](https://www.irishstatutebook.ie/eli/2011/act/3/schedule/2/enacted/en/html "Irish Communications Retention of Data Act 2011, Schedule 2")
- [Irish revised act text](https://revisedacts.lawreform.ie/eli/2011/act/3/revised/en/html "Irish revised act text")
- [Irish Revenue manual on the Communications Retention of Data Act](https://www.revenue.ie/en/tax-professionals/tdm-wm/investigations-prosecutions-enforcement/enforcement/communications-retention-of-data-act.pdf "Irish Revenue manual on the Communications Retention of Data Act")
- [Irish Legal News analysis of new data-retention laws](https://www.irishlegal.com/articles/michael-madden-new-data-retention-laws-for-ireland "Irish Legal News analysis of new data-retention laws")
- [DWT analysis of Irish internet traffic data retention](https://www.dwt.com/insights/2011/02/ireland-requires-storage-of-internet-traffic-data "DWT analysis of Irish internet traffic data retention")
- [Independent Examiner review of Irish retention law](https://independentexaminer.ie/wp-content/uploads/2025/02/236667_a01fdf60-6d1c-43cb-be6b-9440b8d36d6f-3.pdf "Independent Examiner review of Irish retention law")

## Luxemburg

Luxemburg ist unklar und keine stärkere Wahl als Norwegen.

Luxemburg arbeitet nach EU-Rechtsprechungsentwicklungen an Gesetzgebung zur Speicherung von Verkehrs- und Standortdaten. Regierungsquellen beschreiben einen Gesetzentwurf, der Speicherung und Nutzung von Verkehrs- und Standortdaten für nationale Sicherheit, schwere Kriminalität und ernsthafte Bedrohungen der öffentlichen Sicherheit regelt.

Das macht Luxemburg schwer als saubere Datenschutzrechtsordnung für das Privacy.Fish-Modell nutzbar. Der relevante Rahmen ist nicht so klar wie Norwegens enge Zugriffsspeicherpflicht.

Warum Norwegen besser ist:
- Norwegen hat eine bekannte und enge Regel.
- Luxemburgs Speicherrahmen ist für das Privacy.Fish-Modell weniger klar.
- Luxemburgs Gesetzentwurfsumfeld ist öffentlich schwerer zu erklären.
- Privacy.Fish braucht eine Rechtsordnung, in der die verpflichtend gespeicherten Daten präzise beschrieben werden können.

Quellen:
- [Luxembourg government announcement on traffic and location data-retention draft law](https://me.gouvernement.lu/en/actualites.gouvernement2024%2Ben%2Bactualites%2Btoutes_actualites%2Bcommuniques%2B2023%2B01-janvier%2B25-tanson-loi-retention-donnees-caractere-personnel.html "Luxembourg government announcement on traffic and location data-retention draft law")
- [DLA Piper Luxembourg data-protection law overview mentioning draft bill 8148](https://www.dlapiperdataprotection.com/index.html?c=LU&t=law "DLA Piper Luxembourg data-protection law overview mentioning draft bill 8148")
- [Luxembourg CNPD opinion on Bill 8148](https://cnpd.public.lu/en/actualites/national/2024/05/mai-2024-avis-retention-dp.html "Luxembourg CNPD opinion on Bill 8148")
- [Lexology note on Luxembourg retention-law reform](https://www.lexology.com/pro/content/luxembourg-set-reform-retention-law "Lexology note on Luxembourg retention-law reform")
- [Luxembourg Times note on electronic data-retention reform](https://www.luxtimes.lu/luxembourg/luxembourg-moves-to-limit-electronic-data-it-keeps/1358023.html "Luxembourg Times note on electronic data-retention reform")

## Österreich

Österreich ist interessant, aber nicht so geradlinig wie Norwegen.

Österreichs Umsetzung der alten EU-Vorratsdatenspeicherungsrichtlinie wurde 2014 vom österreichischen Verfassungsgerichtshof aufgehoben, nachdem der EuGH die Richtlinie für ungültig erklärt hatte. Das Gericht kritisierte den massiven Eingriff in Grundrechte und die Bedingungen rund um Speicherung, Löschung und Zugriff.

Das macht Österreich potenziell attraktiv, wenn man nur auf das Fehlen pauschaler Vorratsdatenspeicherung schaut. Österreich hat aber weiterhin rechtmäßige Überwachungs- und Datenbereitstellungsrahmen im Telekommunikations- und Strafprozessrecht. Rechtmäßige Überwachung ist nicht dasselbe wie routinemäßige Datenspeicherung, aber sie ist für einen privacy-first E-Mail-Anbieter relevant, weil sie operative Pflichten beeinflusst.

Für Privacy.Fish ist Norwegen weiterhin einfacher zu verteidigen, weil Norwegens verpflichtend gespeicherte Daten explizit, eng und bekannt sind. Österreich kann interessant sein, aber die genauen aktuellen anbieter- und e-mailspezifischen Pflichten erfordern tiefere lokale Rechtsprüfung.

Warum Norwegen besser ist:
- Österreichs altes pauschales Speichergesetz wurde aufgehoben, aber das bedeutet nicht automatisch „keine relevanten Pflichten“.
- Österreich hat weiterhin rechtmäßige Überwachungs- und Datenbereitstellungsrahmen.
- Norwegen gibt eine klare öffentliche Antwort: Konto, Quell-IP:Port, Zeitstempel für 12 Monate.
- Norwegen ist einfacher präzise zu dokumentieren.

Quellen:
- [Austrian Constitutional Court press release on data retention](https://www.vfgh.gv.at/downloads/press_releasedataretention.pdf "Austrian Constitutional Court press release on data retention")
- [Austrian Constitutional Court PDF bulletin](https://www.vfgh.gv.at/downloads/Bulletin_2014-2_G_47-2012__G_59-2012__G_62_70_71-2012_27.06..pdf "Austrian Constitutional Court PDF bulletin")
- [RTR English translation of Austrian Telecommunications Act 2021](https://www.rtr.at/rtr/service/rechtsvorschriften/gesetze/TKG_2021_en-gb.pdf "RTR English translation of Austrian Telecommunications Act 2021")
- [Austrian TKG 2021 PDF via RIS](https://www.ris.bka.gv.at/Dokumente/Erv/ERV_2021_1_190/ERV_2021_1_190.pdf "Austrian TKG 2021 PDF via RIS")
- [FRA Austria surveillance study mentioning annulled data retention](https://fra.europa.eu/sites/default/files/fra_uploads/austria-study-data-surveillance-at.pdf "FRA Austria surveillance study mentioning annulled data retention")
- [Deutsche Telekom Austria transparency report legal basis](https://www.telekom.com/en/company/data-privacy-and-security/news/austria-363540 "Deutsche Telekom Austria transparency report legal basis")
- [EFF-style data-retention cases briefing mentioning Austria](https://www.patrick-breyer.de/wp-content/uploads/2021/09/Briefing-Data-Retention-Cases.pdf "EFF-style data-retention cases briefing mentioning Austria")

## Vergleichstabelle

| Rechtsordnung | Passung der verpflichtenden Protokollierung / Speicherung | Hauptproblem im Vergleich zu Norwegen |
|---|---|---|
| Norwegen | Konto, Quell-IP:Port, Zeitstempel für 12 Monate | Beste Passung |
| EU | Keine einzelne EU-weite Regel nach Ungültigerklärung der Vorratsdatenspeicherungsrichtlinie | Jedes Land muss separat geprüft werden |
| Schweiz | Breitere E-Mail-Metadaten können erfasst sein | Absender, Empfänger, Protokoll, Mailbox-Event und Server-Metadaten |
| Deutschland | Überwachungsbereitschaft für große E-Mail-Systeme | TKÜV-Pflichten über 100.000 Nutzern |
| Niederlande | Kein altes pauschales Speichergesetz, aber Interception- und Disclosure-Pflichten bleiben | Weniger klar als Norwegen |
| Schweden | Breitere Kategorien und variable Fristen | Quelle, Ziel, Gerät, Zeit und kategorieabhängige Speicherung |
| Dänemark | Gezieltes, regionales und nationales Sicherheits-Speichermodell | Komplexer als Norwegen |
| Finnland | Ausgewählte Anbieter; breitere Felder, wenn erfasst | Weniger vorhersehbar und breiter, wenn anwendbar |
| Island | 6 Monate, aber breitere Verbindungsaufzeichnungen | Kann Verbindungsgegenparteien und übertragene Daten umfassen |
| Vereinigtes Königreich | Retention Notices können breite Kommunikationsdaten verlangen | Geheime Notices und breitere Kategorien |
| Vereinigte Staaten | Keine einfache pauschale Regel, aber breite erzwungene Offenlegung und Sicherung | Weniger vorhersehbare rechtliche Exposition |
| Kanada | Keine einfache aktuelle norwegenartige Regel, aber Lawful-Access-Ausweitung vorgeschlagen | Metadatenspeicherung bis zu einem Jahr vorgeschlagen |
| Australien | Breiter vorgeschriebener Metadatensatz für mindestens 2 Jahre | Viel breiter und länger als Norwegen |
| Neuseeland | Keine einfache pauschale Regel gefunden, aber Abfangfähigkeits-Pflichten | Pflichten zum Abfangdesign |
| Frankreich | Breitere nationale Sicherheits-Verkehrs-/Standortspeicherung | Breiter und weniger geeignet |
| Belgien | Gezielte/differenzierte Metadatenspeicherung | Komplex und wiederholt gerichtlich angefochten |
| Irland | Historisch breiterer Internet-/E-Mail-Speicherrahmen | Komplexer und rechtlich umstritten |
| Luxemburg | Entwurfs-/unklarer Verkehrs-Standort-Speicherrahmen | Weniger klar als Norwegen |
| Österreich | Alte pauschale Speicherung aufgehoben, aber Interception-Rahmen bleiben | Interessant, aber weniger klar als Norwegen |
