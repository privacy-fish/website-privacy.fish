---
title: "Privates Anmeldeformular: Der datenschutzfreundlichste Einmalzahlungs-Flow, den man bauen kann"
description: "Wie Privacy.Fish ein privates Anmeldeformular mit selbst gehostetem Captcha, Einmalzahlungen, temporären Zahlungscodes und minimalen Anmeldedaten gebaut hat."
date: 2026-05-06
draft: false
video: "/videos/blog/the-most-privacy-friendly-one-time-payment-form-you-can-build/anonymous-online-payment.mp4"
show_cta: false
---

Privacy.Fish hat die Checkout-Komplexität eines Limonadenstands: ein Produkt, ein Preis, eine Zahlung, kein Abo namens Premium Plus Ultra. Das macht unser Leben einfacher als bei einem normalen Webshop mit Warenkörben, Lieferadressen, Rabattcodes, wiederkehrenden Kunden und Erstattungsabläufen. Die meisten Checkout-Systeme sind dafür gebaut, sich Kunden dauerhaft zu merken. Wir wollten das Gegenteil: bestätigen, dass das richtige Konto bezahlt wurde, und danach die Verbindung zwischen Zahlung und Konto löschen, sobald wir können.

Unsere Antwort ist ein privates Anmeldeformular rund um Einmalzahlungen, temporäre Zahlungscodes, ein selbst gehostetes Captcha gegen Bots und die Löschung der Zahlung-zu-Konto-Verknüpfung nach 14 Tagen, sobald die norwegische Verbraucher-Widerrufsfrist abgelaufen ist.

Die Implementierung ist öffentlich: Der [Code für diese Website](https://github.com/privacy-fish/website-privacy.fish "Code für diese Website") ist Open Source.

## Warum private Anmeldeformulare schwieriger sind, als sie aussehen

Ein normales Anmeldeformular ist meistens der Anfang einer langen Identitätskette: E-Mail-Verifikation, Passwort-Resets, Kartentokens, Rechnungen, Support-Tickets, Betrugsprüfungen, Analytics, Marketing-Tags, Abo-Verlängerungen, Chargeback-Bearbeitung und Kunden-Dashboards. Jedes Teil kann einzeln sinnvoll sein, aber zusammen sorgen sie dafür, dass der Dienst viel mehr behält, als der erste Bildschirm vermuten lässt.

Für Privacy.Fish ist das Ziel anders. Wir wollen wissen, dass ein Konto bezahlt wurde. Wir wollen nicht dauerhaft wissen, wer es bezahlt hat.

### Das Problem mit normalen Checkout-Systemen

Die meisten Checkout-Systeme sind für Conversion, Buchhaltung, Betrugsprävention, Upsells, Abos und Kundensupport gebaut. Datenschutz wird meistens später als Policy-Dokument ergänzt. Das technische Ergebnis ist vorhersehbar: Der Checkout-Anbieter kann Browsersitzung, IP-Adresse, Gerätedaten, gewählte Zahlungsmethode, Betrag, Zeitpunkt, Produkt und irgendeine Form von Kundenkennung sehen.

Für einen Shop, der physische Waren versendet, mag das akzeptabel sein. Für einen Datenschutzdienst ist es ein schlechter Standard.

### Das Problem mit Abos

Abos sind für Zahlungsprivatsphäre schlechter, weil sie eine permanente Verbindung zwischen Kunde, Zahlungsbeziehung und Dienstkonto schaffen. Der Anbieter muss genug Zustand behalten, um erneut abzubuchen, fehlgeschlagene Zahlungen zu wiederholen, zu kündigen, zu verlängern, zu erstatten, Streitfälle zu bearbeiten und Abrechnungsfragen zu beantworten.

Privacy.Fish nutzt eine Einmalzahlung, weil sie einmal zugeordnet, zur Kontoerstellung genutzt, lange genug für das Erstattungsfenster behalten und danach vom Konto getrennt werden kann.

## Schritt 1: Ein selbst gehostetes Captcha ohne Google nutzen

Ein öffentliches Anmeldeformular braucht Missbrauchsschutz. Wenn jeder unbegrenzt Benutzernamen-Verfügbarkeitsfragen stellen oder unbegrenzt Konten reservieren kann, können Bots das Formular in ein Tool zur Benutzernamen-Enumeration oder in eine Warteschlange für Fake-Konten verwandeln. Aber Google reCAPTCHA oder einen anderen großen Drittanbieter-Challenge-Dienst in das Formular zu setzen, erzeugt ein eigenes Datenschutzproblem.

Privacy.Fish nutzt [Cap](https://capjs.js.org/ "Cap"), ein Open-Source-Captcha-System, und liefert das Widget von unserer eigenen Website aus, statt die Challenge von Google oder einem Drittanbieter-CDN zu laden. Das [Signup-Template](https://github.com/privacy-fish/website-privacy.fish/blob/main/layouts/_partials/blocks/signup.html#L1-L13 "Signup-Template") zeigt den Browser auf das lokal bereitgestellte Cap-Widget und WebAssembly-Asset, und das Signup-Backend prüft den Cap-Token gegen unseren konfigurierten Cap-Endpunkt, bevor Benutzernamenprüfungen oder Konto-Reservierungen erlaubt werden.

### Warum Captcha-Versuche begrenzt sind

Das gelöste Captcha gibt keinen unbegrenzten Zugriff auf den Benutzernamen-Checker. Eine erfolgreiche Menschlichkeitsprüfung erstellt eine kurzlebige Berechtigung mit fünf Versuchen, und der Server verbraucht diese Berechtigung, wenn der Benutzer einen Namen prüft oder ein Konto reserviert. Der relevante Backend-Pfad ist [`handleLookupSession`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/server/server.go#L128-L151 "`handleLookupSession`").

Das Formular hat außerdem Rate Limits nach Quell-IP und Limits für erfolgreiche Anmeldungen. Das macht Enumeration nicht unmöglich. Es macht Enumeration teurer, langsamer und auffälliger, ohne die Anmeldesitzung an einen Drittanbieter-Captcha-Dienst zu übergeben.

Für Logs nutzt Privacy.Fish das Modell aus unseren [Logfile-Datenschutzinformationen](/documentation/privacy-information/#logfile-privacy-information/ "Logfile-Datenschutzinformationen"): interne Sicherheitsanalyse, keine Drittanbieter-Abfragen und Löschung normaler Zeilen, wenn kein Alarm ausgelöst wird.

## Schritt 2: Die Benutzerliste schwerer von einem kompromittierten Webserver stehlbar machen

Ein privates Anmeldeformular muss trotzdem eine einfache Frage beantworten: Ist dieser Benutzername verfügbar? Die naive Implementierung ist eine Datenbankabfrage gegen die User-Tabelle unserer E-Mail-Server. Das funktioniert, macht den Signup-Server aber zu einem wertvollen Ziel.

### Feste Benutzernamen-Speichergröße

Der Webserver, auf dem die Privacy.Fish-Website und das Anmeldeformular laufen, muss beantworten können: „Ist dieser Benutzername bereits vergeben?“ Wir wollen aber nicht, dass er eine Klartextliste aller echten Kontonamen enthält. Deshalb ist die bereitgestellte Benutzerdatenbank `used.bin` keine normale Kundentabelle. Sie ist eine binäre Datenbank mit fester Größe, derzeit standardmäßig eine Million Einträge. Jeder Eintrag ist 32 Byte groß: entweder das `HMAC-SHA256`-Ergebnis für einen echten Benutzernamen oder 32 zufällige Bytes.

Die echte Benutzerliste wird auf der Deployment-Workstation verarbeitet, bevor sie auf den Webserver dieser Website hochgeladen wird. Für jeden echten Benutzernamen berechnet der Deployment-Prozess einen keyed hash. Der Pepper ist der geheime Schlüssel für diesen Hash; ohne ihn reicht eine kopierte `used.bin`-Datei nicht aus, um geratene Benutzernamen offline zu testen. Das Salt ist zusätzliche Eingabe in die Hash-Berechnung; es trennt diese Deployment-Datenbank von anderen Datenbanken und älteren Deployments. Beide Werte werden bei jedem Deployment rotiert, sodass ein Angreifer, der eine alte und eine neue Kopie von `used.bin` erhält, die Dateien nicht einfach diffen kann, um neue Einträge und damit neu erstellte Konten zu erkennen.

Nachdem die echten Benutzernamen in 32-Byte-Einträge umgewandelt wurden, wird der Rest der Ein-Million-Einträge-Datenbank mit zufälligen 32-Byte-Padding-Einträgen gefüllt. Danach wird die gesamte Datenbank vor dem Upload gemischt. Die feste Größe verhindert, dass die Kontozahl über die Dateigröße sichtbar wird, und das Mischen verhindert, dass echte Einträge an vorhersehbaren Positionen liegen. Das Signup-Backend prüft außerdem, dass `used.bin` exakt die konfigurierte Größe hat, bevor es sie nutzt.

Zwischen Deployments müssen neue Konten trotzdem als „vergeben“ zählen. Wir bauen nicht jedes Mal die vollständige Ein-Million-Einträge-Datenbank neu, wenn ein Konto erstellt wird. Deshalb werden neu bereitgestellte Benutzernamen in eine kleine Seitendatei `used.bin.overlay` geschrieben, im selben HMAC-Format wie die Hauptdatenbank. Wenn jemand einen Benutzernamen prüft, prüft der Signup-Server beide Stellen: die feste `used.bin`-Datenbank und die Overlay-Datei. Beim nächsten vollständigen Deployment werden die Overlay-Einträge in eine frisch neu gebaute `used.bin` integriert, Pepper und Salt rotiert, und das temporäre Overlay kann wieder leer starten.

Das ist keine magische Verschlüsselung von Benutzernamen. Benutzernamen sind ratbar: Wenn ein Angreifer den Webserver kontrolliert und auch aktuellen Pepper, Salt und Benutzerdatenbank erhält, kann er weiterhin geratenen Namen testen. Der Punkt ist enger: Eine gestohlene Datei soll nicht sofort eine lesbare Kundenliste sein, nicht die genaue Kontozahl verraten, echte Einträge nicht offensichtlich unter Padding-Einträgen zeigen und keine stabilen Hashes offenlegen, die mit einer früheren Deployment-Datenbank verglichen werden können, um neu erstellte Konten zu identifizieren.

### HMAC-Prüfungen statt Klartext-Benutzernamen-Lookups

Wenn ein Benutzer einen Namen prüft, normalisiert der Server ihn, berechnet einen `HMAC-SHA256`-Wert mit deploy-rotiertem Pepper und Salt und prüft diesen Wert gegen die verwendeten und ausstehenden Stores. Der relevante Code ist [`UsernameHMAC`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/server/server.go#L370-L375 "`UsernameHMAC`") und der [`usernameTaken`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/server/server.go#L354-L367 "`usernameTaken`")-Pfad.

Das ist kein Passwort-Hash. Benutzernamen haben von Natur aus wenig Entropie; viele lassen sich raten. Der Punkt ist enger: Der Signup-Server kann Verfügbarkeitsprüfungen durchführen, ohne eine einfache Klartextliste aller registrierten Konten zu tragen.

### Warum Enumeration trotzdem teuer sein muss

Kein öffentlicher Benutzernamen-Checker kann ehrlich versprechen, dass Namen nicht abgefragt werden können. Wenn ein Angreifer fragen kann, ob `alice` vergeben ist, leakt die Antwort selbst ein Bit Information. Die praktische Verteidigung ist, die Rate zu begrenzen, Versuche pro Menschlichkeitsprüfung zu begrenzen, nicht die ganze Datenbank im Ruhezustand offenzulegen und großflächiges Probieren sichtbar zu machen.

Genau dafür ist das Privacy.Fish-Anmeldeformular gebaut.

## Schritt 3: Login-Schlüssel statt Passwörter verlangen

Privacy.Fish fragt während der Anmeldung nach SSH Public Keys, weil wir sie für den Dienst brauchen. Das hat zwar nichts mit der Sicherheit des Anmeldeformulars selbst zu tun, ist aber trotzdem relevant.

Ein SSH-Schlüsselpaar wird für kryptografischen Login und für die Verschlüsselung von E-Mail mit age genutzt: Der Public Key kann auf unseren Servern liegen, während der Private Key auf dem Gerät des Benutzers bleibt. Wenn der Benutzer später verbindet, führt das SSH-Protokoll einen Public-Key-Authentifizierungsablauf aus: Der Server prüft, dass der Client eine gültige Signatur mit dem passenden Private Key erzeugen kann, ohne dass der Private Key oder ein wiederverwendbares Mailbox-Passwort an uns gesendet wird.

Für das Anmeldeformular ist der Datenschutzpunkt simpel: Wir sammeln den gewünschten Benutzernamen und die Authentifizierungsmethode, die wir für den Betrieb des Kontos brauchen. Wir fragen nicht nach Kontakt-E-Mail oder Telefonnummer, weil das Nutzerdaten sind, die wir gar nicht erst haben wollen. Das Backend akzeptiert nur `ssh-ed25519` Public Keys, dedupliziert sie und kanonisiert sie auf Key-Typ und Key-Material; optionale OpenSSH-Kommentare wie `user@laptop` werden entfernt, weil sie unbeabsichtigt Namen, Geräte oder Arbeitsplätze verraten können. Der genaue Parser ist [`internal/sshkey/parse.go`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/sshkey/parse.go#L18-L47 "`internal/sshkey/parse.go`").

## Schritt 4: Einen temporären Zahlungscode erstellen

Wenn der Benutzer den Reservieren-Button drückt, sendet der Browser den gewünschten Benutzernamen und die SSH Public Keys an das Privacy.Fish-Signup-Backend. Das Backend validiert den Benutzernamen, kanonisiert die Public Keys, erzeugt eine temporäre Zahlungsreferenz und speichert einen ausstehenden Signup-Datensatz mit Benutzername, kanonisierten SSH Public Keys, Zahlungsreferenz und Erstellungszeit. Danach sendet das Backend die Zahlungsreferenz zurück an den Browser, damit der Benutzer sie mit der gewählten Zahlungsmethode angeben kann. Der relevante Code ist [`handleSignup`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/server/server.go#L185-L257 "`handleSignup`").

Der Zahlungscode ist absichtlich nicht der Benutzername. Er soll durch den Zahlungskanal reisen: Eine Banküberweisungsreferenz, PayPal-Nachricht, Krypto-Bestätigung oder Notiz in einem Bargeldbrief kann den temporären Zahlungscode enthalten. Was nicht durch diese Kanäle reisen soll, ist die Konto-ID. Der Benutzer sollte nicht seine gewünschte E-Mail-Adresse, seinen echten Namen, seine bestehende E-Mail-Adresse, Telefonnummer oder andere identifizierende Texte neben den Zahlungscode setzen.

### Warum das Konto reserviert wird, bevor die Zahlung eingeht

Bargeld per Brief, SEPA-Überweisungen und Kryptowährungszahlungen sind nicht immer sofort. Alle Zahlungsmethoden, auch Kreditkarte und PayPal, erfordern trotzdem manuelle Kontoaktivierung, weil der Webserver mit der Privacy.Fish-Website und dem Anmeldeformular absichtlich von der Mail-Infrastruktur getrennt ist.

Also wird das Konto zuerst reserviert. Der Benutzername wird 30 Tage gehalten, und der Benutzer erhält einen Zahlungscode. Wenn die Zahlung mit diesem Code eingeht, kann Privacy.Fish das richtige Konto erstellen.

Weil das Anmeldeformular keine Kontaktinformationen sammelt, kann Privacy.Fish keine „Dein Konto ist fertig“-Benachrichtigung senden. Das ist weniger bequem, vermeidet aber einen weiteren dauerhaften Identitätskanal.

## Schritt 5: Zahlungsmethoden ohne einen zentralen Checkout-Anbieter anbieten

Die meisten Websites packen alle Zahlungsmethoden hinter einen Checkout-Anbieter. Das ist bequem, schafft aber auch einen zentralen Beobachter, der Zahlungscode, Browsersitzung, IP-Adresse, Gerätedaten, Zahlungswahl, Betrag, Zeitpunkt und oft die Zahlungsidentität sehen kann.

Privacy.Fish vermeidet das, wo möglich. Wir bevorzugen getrennte Zahlungswege statt eines zentralen Prozessors für alles.

### Bargeld per Brief

Bargeld per Brief ist die stärkste praktische Option, die wir anbieten. Der Benutzer legt 20 EUR in einen Umschlag, fügt nur den temporären Zahlungscode hinzu und sendet ihn an Privacy.Fish. Wir sagen ausdrücklich, dass Benutzer keine Münzen senden sollen, weil Münzen Papier bei der Briefverarbeitung leicht zerreißen.

Bargeld ist keine magische Anonymität. Postsysteme, Kameras, Fingerabdrücke, Handschrift, Druckermarken, Porto-Zahlung und lokale Postregeln können alle eine Rolle spielen. Aber aus unserer Sicht erhält Privacy.Fish Zahlung und Code, nicht Bankkonto, Kartennummer, PayPal-Identität, Kryptowährungsadresse oder Abrechnungsprofil.

### Kryptowährung

Kryptowährung wird aus Bequemlichkeit unterstützt, nicht weil sie automatisch privat ist. Öffentliche Blockchains sind öffentlich, viele Benutzer kaufen Krypto über KYC-Börsen, und Wallet-Infrastruktur kann Informationen leaken.

Privacy.Fish nutzt kalte Empfangs-Wallets, damit Private Keys nicht auf Website- oder Mailservern liegen. Für BTC, ETH und USDT geben gängige Wallets und Börsen uns kein verlässliches Zahlungscode-Memo-Feld, deshalb übermitteln Benutzer Transaktionshash und Zahlungscode über ein separates Krypto-Bestätigungsformular.

### SEPA, PayPal, Kreditkarte und andere bequeme Methoden

SEPA, PayPal, Kreditkarten und ähnliche Komfortmethoden sind gegenüber den Zahlungs­systemen, die sie verarbeiten, nicht privat. Banken, PayPal, Kartennetzwerke, Kartenherausgeber, Acquirer, Zahlungsprozessoren, Betrugssysteme und Compliance-Systeme können je nach Methode Zahlungsidentität, Kontodetails, Gerätedaten, Zeitstempel, Transaktionsreferenzen und andere Zahlungsmetadaten sehen.

Die Datenschutzverbesserung ist hier nicht, dass diese Zahlungsmethoden anonym werden. Das werden sie nicht. Die Verbesserung ist, dass Privacy.Fish nicht jede Methode durch einen zusätzlichen zentralen Checkout-Anbieter leitet, und dass die Zahlungsnachricht weiterhin nur den temporären Zahlungscode enthalten soll, nicht den gewünschten Kontonamen.

## Schritt 6: Die Verbindung zwischen Zahlung und Konto löschen

Der stärkste Teil des Designs ist nicht der Zahlungscode selbst. Ein Zahlungscode ist nur eine temporäre Kennung. Das Datenschutzfeature ist die Löschung der Zuordnung, nachdem sie ihren Zweck erfüllt hat.

Während der Kontoerstellung muss Privacy.Fish drei Dinge zusammen kennen: den gewünschten Benutzernamen, SSH Public Keys und den Zahlungscode, der die Zahlung belegt. Auf der Festplatte des Webservers, auf dem die Privacy.Fish-Website und das Anmeldeformular laufen, werden ausstehende Signup-Daten als age-verschlüsseltes Payload für den Public Key des Administrators gespeichert; der Pending-Datensatz hält Metadaten wie den Username-HMAC und den verschlüsselten Ciphertext. Der relevante Code ist [`internal/store/pending.go`](https://github.com/privacy-fish/website-privacy.fish/blob/main/internal/store/pending.go#L59-L91 "`internal/store/pending.go`").

Das bedeutet nicht, dass der laufende Webprozess das übermittelte Formular nie sieht. Ein Formular-Handler erhält das Formular notwendigerweise während der Verarbeitung der Anfrage. Der Punkt ist, dass ausstehende Signup-Daten nicht im Klartext auf der Festplatte des Webservers zurückbleiben und dass dieser Webserver keine Mailkonten automatisch erstellen kann, weil es aus Sicherheitsgründen absichtlich keinen Netzwerkpfad von ihm zur Mail-Infrastruktur gibt.

Nach Zahlung und Kontoerstellung behält Privacy.Fish die Zahlungscode-zu-Konto-Zuordnung nur so lange, wie sie für die Verbraucher-Erstattungsfrist benötigt wird. Danach wird die Zuordnung gelöscht. Ab dann hat Privacy.Fish eine Liste erhaltener Zahlungen und eine Liste aktiver Konten, aber keinen Systemdatensatz, der beantwortet, welche Zahlung welches Konto erstellt hat.

Abos können so nicht funktionieren. Ein Abo erfordert laufenden Abrechnungszustand. Der Anbieter muss wissen, welches Konto zu welcher Abrechnungsbeziehung gehört, damit er das Abo verlängern, fehlschlagen lassen, kündigen, erstatten oder bestreiten kann.

## Schlussgedanken: Was „möglichst datenschutzfreundlich“ bedeutet

Privacy.Fish existiert aus einem sturen Grund: Wir wollen E-Mail-Infrastruktur bauen, die so wenig Nutzerdaten sammelt, wie wir irgendwie können, auch wenn das Produkt dadurch weniger bequem und das Geschäftsmodell weniger profitabel wird.

Für die Anmeldung bedeutet das: kein echter Name, keine Kontakt-E-Mail, kein Passwort, kein Abo-Konto, kein zentraler Checkout-Anbieter außer dort, wo eine Zahlungsmethode ihn erfordert, kein Benutzername in der Zahlungsnachricht und keine dauerhafte Zahlung-zu-Konto-Verbindung. Oder kürzer: so wenig Daten wie möglich sammeln.

Ein privates Anmeldeformular entsteht nicht dadurch, dass man einem normalen Checkout eine Datenschutz-Checkbox hinzufügt. Es muss von Anfang an um Datenminimierung herum entworfen werden. Für Privacy.Fish bedeutet das selbst gehosteten Bot-Schutz, begrenzte Benutzernamenprüfungen, SSH-Schlüssel statt Passwörter, temporäre Zahlungscodes, getrennte Zahlungswege und Löschung der Zahlung-zu-Konto-Verbindung.
