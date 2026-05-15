---
title: "Datenschutzinformationen"
description: "Datenschutzinformationen zu Zahlungscodes, Einmalzahlungen, Anbieterexposition, Serverlogs und Website-Daten."
weight: 60
video: "/videos/private-payment.mp4"
aliases:
  - /documentation/payment-and-account-creation/
---

# Datenschutzinformationen

## Datenschutzinformationen zur Zahlung

Privacy.Fish nutzt eine Einmalzahlung statt wiederkehrender Abo-Gebühren aus einem Grund: zum Schutz deiner Privatsphäre.

Eine Einmalzahlung erzeugt weniger laufende Zahlungsmetadaten als monatliche Abrechnung oder vorausbezahlte Abo-Zeiträume. Sie macht private Zahlungsmethoden auch realistischer, weil es eine viel kleinere Hürde ist, einmal 20 € per Brief in ein anderes Land zu senden, als das jeden Monat zu tun.

Die Hürde zu einem privaten und sicheren E-Mail-Konto sollte niedrig sein.

Wir haben überlegt, den Dienst spendenbasiert wie Signal Messenger zu machen, uns aber dagegen entschieden, um Spam-Kontoerstellung zu reduzieren und den Dienst nutzbar zu halten.

### Wie der Signup-Zahlungscode deine Privatsphäre schützt

Privacy.Fish nutzt temporäre Zahlungscodes, um zu verhindern, dass deine Zahlungsinformationen dauerhaft das E-Mail-Konto identifizieren, für das du bezahlt hast.

Das Ziel ist einfach: Die Zahlung bestätigt, dass ein Konto erstellt werden soll, aber nach der Erstellung kann die gelöschte Verbindung diese Zahlung nie wieder einem Konto zuordnen.

Was wir dauerhaft wissen wollen, ist auf zwei getrennte Fakten beschränkt: welche E-Mail-Konten bezahlt wurden und welche Zahlungen wir erhalten haben. Wir wollen keinen dauerhaften Datensatz behalten, der zeigt, welche Zahlung welches Konto erstellt hat.

Wir haben natürlich keinen Einfluss darauf, welche Daten die Post bei Bargeld per Brief speichert, was unsere norwegische Bank bei SEPA-Überweisungen speichert, was unsere Zahlungspartner bei Kreditkarten und PayPal speichern oder was Kryptowährungs-Blockchains über das Geld speichern, das du uns sendest.

Wenn du ein Konto erstellst, indem du einen Benutzernamen und deine SSH Public Keys übermittelst, wird dieser Benutzername für 30 Tage reserviert. Wenn die Zahlung nicht innerhalb dieser 30 Tage abgeschlossen wird, wird der Benutzername wieder für die Registrierung verfügbar.

Wenn du die Zahlung ausführst, gibst du den temporären Zahlungscode in der Transaktion an. Damit können wir die eingegangene Zahlung dem reservierten Benutzernamen zuordnen und das richtige Konto erstellen.

Wir behalten die Zahlungscode-zu-Benutzername-Zuordnung 14 Tage lang, damit Erstattungs- und Widerrufsanfragen während der norwegischen gesetzlichen Verbraucher-Erstattungsfrist noch bearbeitet werden können.

Nach diesen 14 Tagen löschen wir die Zahlungscode-zu-Benutzername-Zuordnung dauerhaft. Ab diesem Zeitpunkt hat Privacy.Fish keinen Datensatz mehr, der deine Zahlungsdaten mit deinem E-Mail-Konto verbindet. Selbst wenn du uns beweisen könntest, dass du eine Zahlung geleistet hast, könnten wir nicht sagen, für welches Konto sie war.

### Warum wir nur eine Einmalzahlung verlangen

Privacy.Fish verlangt einmalig Geld, statt ein monatliches oder jährliches Abo am Leben zu halten. Das ist schlechter für planbare wiederkehrende Einnahmen, und wir gehen davon aus, dass wir dadurch im Vergleich zu einem normalen Abo-Modell Geld verlieren.

Wir tun es trotzdem, weil wir dadurch die Zahlung-zu-Konto-Verbindung löschen können. Mit einem Abo müsste Privacy.Fish die Daten behalten, die das Konto mit einer wiederkehrenden Zahlungsmethode verbinden, damit wir erneut abbuchen, fehlgeschlagene Verlängerungen behandeln und Kündigungen verwalten können. Mit einer Einmalzahlung brauchen wir den temporären Zahlungscode nur lange genug, um das Konto zu erstellen und die gesetzliche Erstattungsfrist von 14 Tagen zu handhaben. Danach kann die Verbindung zwischen Zahlung und Benutzername gelöscht werden.

### Datenschutzinformationen zu Zahlungsanbietern

Die folgenden Kapitel beschreiben, wie stark jede Zahlungsoption deine Privatsphäre verletzt. Sortiert von am wenigsten bis am stärksten.

#### Bargeld per Brief

Bargeld per Brief ist unsere bevorzugte Zahlungsmethode für Benutzer, die die stärkste praktische Zahlungsprivatsphäre wollen.

#### Was du senden sollst

Sende keine Münzen, sondern nur Banknoten, weil Münzen das Papier bei der Briefverarbeitung durch Postunternehmen leicht zerreißen.

##### Empfängeradresse für Bargeld per Brief

Bitte lege 20 Euro in einen Umschlag, füge einen Brief hinzu, der NUR! den temporären Zahlungscode enthält und NICHT! deinen Benutzernamen, und sende ihn an:

```text
Privacy.Fish
Fjordstreet 5
12345 Oslo
Norway
```

Das gibt Privacy.Fish nur die Informationen, die wir brauchen: einen Zahlungscode und die Zahlung selbst.

##### Kosten für einen Brief an Privacy.Fish in Oslo, Norwegen

Wir haben recherchiert, dass das ungefähre Porto für einen kleinen gewöhnlichen Brief nach Norwegen mit einer Banknote und einem A4-Standardbrief für den Zahlungscode aus den meisten Ländern der Welt zwischen einem und fünf Euro liegt, wenn du den Brief nicht verfolgen willst und ihn als normalen Brief sendest.

Privacy.Fish bietet keine Erstattungen für Bargeld-per-Brief-Zahlungen an, unabhängig davon, ob die Zahlungen vollständig oder unvollständig waren.

##### Datenschutzinformationen für Bargeld per Brief

Bargeld per Brief ist trotzdem keine magische Privatsphäre. Je nachdem, wie du den Brief vorbereitest und sendest, kannst du weiterhin Informationen gegenüber anderen Parteien offenlegen. Umschlag, Papier und Banknote können Fingerabdrücke, DNA-Spuren, Handschrift, Druckermarken oder andere physische Spuren enthalten. Wenn du den Brief in einem Postamt, Geschäft, Kiosk oder an einem anderen bedienten Ort abgibst, kann es Überwachungskameras geben. Wenn du Porto per Kreditkarte oder mit einer anderen nachvollziehbaren Zahlungsmethode bezahlst, wird diese Zahlung protokolliert. Wenn du eine Absenderadresse hinzufügst, kann der Brief leichter identifiziert werden.

Postbetreiber und Versandunternehmen können ebenfalls Versanddaten sammeln, aber die genauen Daten hängen vom Land, Postbetreiber, Versandprodukt und der Zahlungsweise ab.

Wir wissen nicht, ob jeder Postbetreiber Briefinhalte scannt, aber wir nehmen an, dass es technisch möglich ist zu erkennen, dass ein Umschlag eine Banknote enthält, oder den Zahlungscode ohne Öffnen des Briefs auszulesen. Der Datenschutzvorteil ist deshalb nicht, dass Bargeld per Brief nirgendwo Spuren hinterlässt.

Der Vorteil ist, dass Privacy.Fish kein Bankkonto, keine Kartennummer, kein PayPal-Konto, keine Kryptowährungsadresse und keine andere digitale Zahlungskennung erhält, die höchstwahrscheinlich direkt mit dir verbunden ist.

Übermittle im Bargeldbrief nicht deinen Benutzernamen, deine E-Mail-Adresse, deinen echten Namen, deine Rücksendeadresse oder andere identifizierende Informationen. Übermittle nur den temporären Zahlungscode und die Zahlung selbst.

#### Kryptowährungen

##### Kryptowährungs-Zahlungsadressen

Bitte sende 20 Euro an:

| Kryptowährung | Beispieladresse |
| --- | --- |
| Bitcoin (BTC) | `bc1qexampleaddressdonotuse0000000000000000000000` |
| Ethereum (ETH) | `0xExampleAddressDoNotUse0000000000000000000000` |
| Tether (USDT, ERC-20) | `0xExampleAddressDoNotUse0000000000000000000000` |

##### Den Zahlungscode übermitteln

BTC, ETH und USDT bieten über gängige Wallets und Börsen hinweg kein verlässliches Zahlungs-Code-Memo-Feld. Nachdem du die Transaktion gesendet hast, übermittle den temporären Zahlungscode, die Kryptowährung und den Transaktionshash über das [Krypto-Zahlungsbestätigungsformular](https://github.com/signup/crypto-confirmation-form "Krypto-Zahlungsbestätigungsformular"). Übermittle NICHT deinen Benutzernamen, deine E-Mail-Adresse, deinen echten Namen oder andere identifizierende Informationen! Das Bestätigungsformular existiert, damit du uns den temporären Zahlungscode zusammen mit der Blockchain-Transaktion senden kannst, sodass Privacy.Fish das reservierte Konto erstellen kann. Diese Zuordnung wird nach der norwegischen gesetzlichen Verbraucher-Erstattungsfrist von 14 Tagen nach Zahlungseingang gelöscht.

##### Datenschutzinformationen für Kryptowährungszahlungen

Kryptowährungszahlungen werden aus Bequemlichkeit unterstützt, nicht weil sie automatisch privat sind. Für die meisten Benutzer werden Kryptowährungen über eine regulierte Börse wie Coinbase, Kraken, Binance oder einen anderen Anbieter gekauft, gehalten und gesendet, der die Identität des Benutzers kennt. Wenn du Privacy.Fish von einer solchen Börse bezahlst, weiß die Börse, wer du bist, wann du bezahlt hast, wie viel du bezahlt hast und an welche Privacy.Fish-Adresse du bezahlt hast.

Öffentliche Blockchains erzeugen ein weiteres Problem: Transaktionen sind dauerhaft sichtbar. Je nach Kryptowährung können Senderadresse, Empfängeradresse, Betrag, Zeit, Transaktionsgebühr und spätere Bewegungen der Gelder von jedem analysiert werden. Wenn du Wallet-Software verwendest, die sich mit Drittanbieter-Nodes oder APIs verbindet, können diese Anbieter auch deine IP-Adresse und die Adressen oder Transaktionen erfahren, die du abfragst.

Privacy.Fish könnte für jede eingehende Zahlung eine frische Empfangsadresse verwenden. Das würde verhindern, dass ein Angreifer eine einzelne öffentliche Adresse anschaut und jede andere Zahlung an dieselbe Adresse sieht. Mehrere eingehende Adressen schaffen aber nicht den Datenschutzvorteil, den viele erwarten. Wenn der Benutzer von einer Know-Your-Customer-(KYC)-Börse bezahlt, kennt die Börse die Auszahlung trotzdem. Wenn Privacy.Fish später viele empfangene Zahlungen zusammen bewegt, kann Blockchain-Analyse diese Empfangsadressen wieder verbinden. Und weil Privacy.Fish ein Unternehmen ist, muss Kryptowährung irgendwann getauscht, ausgegeben oder in normales Geld umgewandelt werden. Dieser Auszahlungsprozess kann einen weiteren Punkt schaffen, an dem Transaktionen verknüpfbar werden.

Deshalb tut Privacy.Fish nicht so, als würde die Verwaltung vieler Kryptowährungs-Empfangsadressen Benutzern starke Zahlungsprivatsphäre geben. Wir nutzen eine langlebige kalte Empfangsadresse pro unterstützter Kryptowährung. Die Private Keys werden offline erzeugt und gespeichert, nicht auf Website- oder Mailservern. Das hält den operativen Aufbau einfach und hält Private Keys von internetexponierten Systemen fern, bedeutet aber auch, dass Zahlungen an diese Adresse öffentlich als Zahlungen an Privacy.Fish sichtbar sind.

Wenn du starke praktische Zahlungsprivatsphäre willst, nutze Bargeld per Brief. Wenn du Kryptowährung nutzt, nimm an, dass Börse, Blockchain-Beobachter, Wallet-Infrastruktur oder späterer Cash-out-Prozess die Zahlung mit Privacy.Fish verbinden können.

Wenn du Kryptowährung legal anonym erhalten und senden kannst, ist das besser für deine Privatsphäre als eine Zahlung von einem Börsenkonto, das mit deiner Identität verbunden ist. Es ist außerdem besser, von einer Wallet zu senden, die nur für diese Zahlung verwendet wird, auf einem Gerät und Netzwerk, dem du vertraust, statt von deinem Alltags-Handy, Browser-Wallet oder Börsenkonto. Das macht die Zahlung trotzdem nicht garantiert anonym, reduziert aber die Anzahl der Stellen, an denen sie zu dir zurückverknüpft werden kann.

Übermittle den temporären Zahlungscode, die verwendete Kryptowährung und den Transaktionshash über das [Krypto-Zahlungsbestätigungsformular](https://github.com/signup/crypto-confirmation-form "Krypto-Zahlungsbestätigungsformular"). Versuche nicht, uns deinen Benutzernamen, deine E-Mail-Adresse, deinen echten Namen oder andere identifizierende Informationen zu senden.

#### SEPA-Überweisung

##### SEPA IBAN und BIC

Bitte sende 20 Euro an:

BANK Name: Sparbank Oslo
IBAN: NE 1234 5678 9012 3456 78
BIC: 45678938
Betreff: **DEIN-ZAHLUNGSCODE-HIER**

##### Datenschutzinformationen für SEPA-Überweisungen

Unsere Bank kann Informationen über eingehende Banküberweisungen sehen und speichern. Nach den Datenschutzinformationen von SpareBank 1 kann dies Finanzinformationen wie Kontonummern und Transaktionsdaten sowie Zahlungsinformationen und andere Daten umfassen, die benötigt werden, um Bankdienstleistungen bereitzustellen, Missbrauch zu verhindern, Buchhaltungspflichten zu erfüllen und rechtlichen Verpflichtungen nachzukommen. SpareBank 1 sagt, personenbezogene Daten würden grundsätzlich so lange gespeichert, wie es für den Zweck erforderlich ist, für den sie erhoben wurden, sofern Gesetz oder Regulierung keine längere Speicherung verlangen. Ihre Speicherbeispiele umfassen Dokumentation für Geldwäsche- und Terrorismusfinanzierungsbekämpfung bis zu 10 Jahre nach abgeschlossener Transaktion oder beendeter Kundenbeziehung sowie Informationen, die nach norwegischen Buchführungsregeln bis zu 10 Jahre erforderlich sind. ([SpareBank 1 Datenschutzinformationen](https://www.sparebank1.no/nb/bank/om-oss/personvern.html "SpareBank 1 Datenschutzinformationen"), [SpareBank 1 Datenschutzrichtlinie PDF](https://www.sparebank1.no/content/dam/SB1/vedlegg/veiledninger/Privacy-policy.pdf "SpareBank 1 Datenschutzrichtlinie PDF"))

Für eine eingehende SEPA-Überweisung geht Privacy.Fish deshalb davon aus, dass unsere Bank mindestens Sendername, Senderkonto oder IBAN, Empfängerkonto, Betrag, Zeitstempel, Zahlungsreferenz / Nachrichtenfeld und interne Transaktionsmetadaten speichern kann. Wir kontrollieren nicht die Datenschutzrichtlinie, Speicherfrist, Compliance-Systeme oder Betrugsmonitoring-Systeme der Bank, die du für die Überweisung verwendest.

Deine eigene Bank kann mehr Informationen speichern, als wir jemals erhalten, und wir haben keine Möglichkeit zu wissen, was genau sie speichert.

Übermittle mit der SEPA-Überweisung nicht deinen Benutzernamen, deine E-Mail-Adresse, deinen echten Namen oder andere identifizierende Informationen. Übermittle nur den temporären Zahlungscode im Verwendungszweck.

#### PayPal

##### PayPal-Zahlungsadresse

Bitte sende 20 Euro an:

```text
paypal@privacy.fish
```

##### Datenschutzinformationen für PayPal-Zahlungen

PayPal ist bequem, aber keine private Zahlungsmethode. PayPal kann Konto-IDs, Kontaktinformationen, Finanzinformationen, Transaktionsinformationen, Gerätedaten, IP-Adressen, Cookies, Betrugssignale und andere Daten verarbeiten, die benötigt werden, um Dienste bereitzustellen, Missbrauch zu verhindern, Streitfälle zu lösen, Gesetze einzuhalten und das Zahlungsnetzwerk zu betreiben. PayPal kann Informationen auch an andere Kontoinhaber, Händler, Dienstleister, Finanzinstitute, Behörden und andere Parteien offenlegen, wenn dies zur Durchführung von Transaktionen oder zur Erfüllung rechtlicher Pflichten erforderlich ist. ([PayPal Privacy Statement](https://www.paypal.com/li/legalhub/privacy-full?locale.x=en_US "PayPal Privacy Statement"), [PayPal User Agreement](https://www.paypal.com/us/legalhub/paypal/useragreement-full?locale.x=us_US "PayPal User Agreement"))

Für eine PayPal-Zahlung an Privacy.Fish solltest du annehmen, dass PayPal wissen kann, wer du bist, welches PayPal-Konto oder welche Zahlungsmethode du verwendet hast, wann du bezahlt hast, wie viel du bezahlt hast, welches Empfängerkonto du bezahlt hast, welche Transaktions-ID entstand und welche Zahlungsnachricht den temporären Zahlungscode enthält. Privacy.Fish kann außerdem Senderinformationen sehen, die PayPal dem Empfängerkonto offenlegt. Wir können nicht kontrollieren, was PayPal speichert, wie lange PayPal es behält, welche Betrugs- oder Compliance-Systeme es verarbeiten oder was deine verknüpfte Bank oder dein Kartenanbieter speichert.

Übermittle mit der PayPal-Transaktion nicht deinen Benutzernamen, deine E-Mail-Adresse, deinen echten Namen oder andere identifizierende Informationen. Übermittle nur den temporären Zahlungscode in der PayPal-Transaktionsnachricht.

#### Warum wir keinen Zahlungsprozessor verwenden

Für die meisten Zahlungsmethoden vermeidet Privacy.Fish einen zentralen Checkout oder Zahlungsprozessor, weil dieser noch mehr Daten sammeln würde, als die Zahlungsmethode selbst bereits offenlegt. Ein zentraler Zahlungsprozessor kann Zahlungscode, Browsersitzung, IP-Adresse, Gerätedaten, Wahl der Zahlungsmethode, Transaktionszeitpunkt, Betrag und oft die Zahlungsidentität hinter der Transaktion sehen.

Stattdessen bevorzugen wir getrennte Zahlungswege, wo möglich. Bargeld per Brief muss überhaupt keinen Zahlungsprozessor berühren. SEPA legt Daten gegenüber den beteiligten Banken offen, aber nicht gegenüber einem zusätzlichen Checkout-Anbieter. Kryptowährung legt Blockchain- und Börsendaten offen, muss aber keine Browser-Checkout-Metadaten gegenüber einem Zahlungsprozessor offenlegen. PayPal legt Daten gegenüber PayPal offen, verlangt aber keinen weiteren Prozessor davor.

Kreditkarten sind die Ausnahme. Kartenzahlungen benötigen einen Zahlungsprozessor und sind deshalb die am wenigsten private Zahlungsoption, die wir anbieten.

#### Kreditkarte

Für Kreditkarten müssen wir leider einen Zahlungsprozessor verwenden.

##### Kreditkarten-Zahlungsprozessor

Privacy.Fish nutzt Dintero Checkout für Kreditkartenzahlungen. Dintero ist ein norwegischer Zahlungsanbieter mit Sitz in Oslo, und Dintero Checkout ist ihr Online-Checkout-Produkt für Händler. ([Dintero Checkout](https://dintero.com/products/dintero-checkout "Dintero Checkout"), [Dintero Terms for Checkout](https://dintero.com/legal/terms-for-checkout "Dintero Terms for Checkout"))

##### Datenschutzinformationen für Kreditkartenzahlungen

Kreditkartenzahlungen sind die am wenigsten private Zahlungsmethode, die Privacy.Fish anbietet. Wenn du per Karte bezahlst, können Dintero, das Kartennetzwerk, der Acquirer, dein Kartenherausgeber und möglicherweise andere Zahlungsdienstleister Informationen über die Transaktion verarbeiten. Dinteros Datenschutzrichtlinie sagt, dass Dintero Kontaktinformationen, Kontoinformationen, Zahlungskartendaten, Kontonummern, Browser- und Gerätedaten, Nutzungsdaten, Kaufhistorie, transaktionsbezogene Daten, Betrugspräventionsdaten und Daten verarbeiten kann, die durch Finanz-, Buchhaltungs- oder andere gesetzliche Verpflichtungen erforderlich sind. Dintero sagt außerdem, dass der Händler typischerweise Controller für Daten ist, die zur Verwaltung des Kaufs benötigt werden, während ein Zahlungsdienstleister Controller für Daten ist, die zur Verarbeitung der Zahlung benötigt werden. ([Dintero Privacy Policy](https://dintero.com/no/legal/privacy-policy "Dintero Privacy Policy"), [Dintero End-Customer Terms](https://dintero.com/no/legal/end-customer-terms "Dintero End-Customer Terms"))

Privacy.Fish will deine Kartennummer nicht. Der Sinn von Dintero ist, dass die Kartenzahlung über deren Zahlungsinfrastruktur abgewickelt wird, statt dass Privacy.Fish Rohkartendaten speichert oder verarbeitet. Das bedeutet aber auch, dass eine Drittanbieter-Checkout- und Kartenzahlungskette existiert, und diese Kette kann deutlich mehr Zahlungsmetadaten sehen, als Privacy.Fish wissen will.

Übermittle bei der Kreditkartenzahlung nicht deinen Benutzernamen, deine E-Mail-Adresse, deinen echten Namen oder andere identifizierende Informationen. Übermittle nur den temporären Zahlungscode im Zahlungsformular.

## Logfile-Datenschutzinformationen

### Logfiles der Mailserver

Privacy.Fish ist nach norwegischem Recht verpflichtet, Logfiles registrierter Konten für SSH- und SFTP-Logins zu den Servern zu speichern. Das betrifft jeden Zeitpunkt, an dem du E-Mail sendest oder abrufst oder deine SSH Public Keys änderst. Für jeden Login speichern wir Zeitstempel, Benutzername, Quell-IP-Adresse und Quell-Port. Wir müssen diese Daten für die norwegische Rechtskonformität 12 Monate speichern.

Neben dieser gesetzlichen Pflicht strebt Privacy.Fish an, „so wenig Daten wie möglich“ zu haben.

Logfiles von den Servern werden an ein internes Analysesystem gestreamt, für Sicherheitsalarme verarbeitet und verworfen, wenn kein Alarm ausgelöst wird. Das Analysesystem sendet keine Logdaten an Dritte und führt keine externen Abfragen wie IP-Reputationschecks aus. Wenn verdächtige Aktivität erkannt wird, speichert Privacy.Fish den Alarm und die minimalen Logzeilen, die nötig sind, um ihn zu verstehen. Administratoren können während Incident Response vorübergehend detaillierteres Logging aktivieren, aber nur für die Zeit, die zur Bearbeitung des Vorfalls nötig ist, und deaktivieren es danach wieder. Nach dem Debugging werden diese temporären Logs gelöscht. Wenn Logs aus rechtlichen Gründen oder für einen Incident Report behalten werden müssen, werden sie so weit wie möglich minimiert und anonymisiert.

### Logfiles des Webservers

## Website-Datenschutzinformationen
