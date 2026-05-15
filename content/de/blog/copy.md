---
title: "Das datenschutzfreundlichste Einmalzahlungsformular, das man bauen kann"
description: "Wie Privacy.Fish ein datenschutzfreundliches Einmalzahlungsformular mit temporären Zahlungscodes, ohne Abos und mit minimalen Anmeldedaten gebaut hat."
date: 2026-05-06
draft: true
video: "/videos/blog/norway-email-privacy-law-why-it-beats-germany-switzerland-the-eu-and-us-for-private-email/norway-email-privacy-law.mp4"
show_cta: false
---

Für uns ist es leicht, das zu sagen: Wir haben nur ein Produkt, einen Preis dafür und verlangen nur eine Einmalzahlung. Zahlungsinformationen über unsere Kunden aus unseren Abläufen herauszuhalten, ist wahrscheinlich einfacher als bei einem großen Webshop mit wiederkehrenden Kunden.

Unser Ziel ist, alle Systeme und Geschäftsentscheidungen so privat wie möglich zu entwerfen. Für unser Modell war wichtig, nur zu wissen, dass ein Konto bezahlt wurde, nicht wer es bezahlt hat.

Wir haben wahrscheinlich das privateste und trotzdem praktische Anmeldeformular mit Einmalzahlungsfunktion gebaut, das man sich vorstellen kann. Zerlegen wir das Anmeldeformular Stück für Stück und erklären, worum es geht:

## 1 Menschlichkeitsprüfung zuerst

Um zu verhindern, dass Bots mit dem Anmeldeformular Chaos anrichten, verlangen wir ein Captcha, bevor es losgehen kann.

Das Open-Source-Captcha-Backend (Link ergänzen, welches wir verwenden) wird auf demselben Server wie die Website gehostet und protokolliert oder sendet keine Daten an Dritte.

## 2 Benutzernamen wählen

- Wir erzeugen eine DB mit allen echten Benutzernamen plus zufälligen Benutzernamen, bis wir auf 1 Million Einträge kommen, und hashen + salzen + peppern die Einträge.
- Wenn der Benutzer prüft, ob eine E-Mail-Adresse frei ist, fragt es den Server, der die Auswahl des Benutzers hasht und die DB prüft.
- Bots können nicht einfach enumerieren, weil Captcha.

Beim Deployment laden wir neue Kontoregistrierungen herunter, fügen sie zur DB hinzu, bauen die DB neu und rotieren Salt + Pepper bei jedem Deployment. So ändern sich die Hashes ständig und ein Angreifer kann DB-Änderungen nicht über Zeit verfolgen.

## 3 SSH Public Keys hinzufügen

Das hat nichts mit der Privatsphäre des Signups zu tun und ist nur eine technische Notwendigkeit unseres Dienstes.

## 4 Zahlungsreferenz

Beschreibe, wie der Zahlungslink funktioniert.

## 5 Zahlungsmethoden

Schreibe darüber, wie wir Bargeld per Brief anbieten, und scherze darüber, dass man dazuschreiben muss, keine Münzen zu senden.

Sage, dass wir Cold Wallets für Krypto nutzen, sodass wir nur die Blockchain auf neue Zahlungen überwachen können, und dass wir ein separates Kontaktformular nur für Krypto-Zahlungen haben, in dem Benutzer Transaktionshash + Zahlungscode senden, damit wir wissen, für welches Konto die Zahlung war. Das Kontaktformular liegt auf unserer Website, damit sie keine E-Mail / WhatsApp senden und damit keinen Drittanbieter einbeziehen müssen.

Sage, dass wir auch SEPA, PayPal und Kreditkarte anbieten, aber dass wir denken, dass 20 Euro per Brief einmalig eine sehr niedrige Hürde sind.

## 6 Einrichtungsanleitungen

Das hat nichts mit der Privatsphäre des Signups zu tun und ist nur eine technische Notwendigkeit unseres Dienstes.
