---
title: "Privates Anmeldeformular: Der datenschutzfreundlichste Einmalzahlungs-Flow"
description: "Wie Privacy.Fish ein privates Anmeldeformular mit selbst gehostetem Captcha, Einmalzahlung, temporären Zahlungscodes und minimalen Anmeldedaten baut."
date: 2026-05-06
draft: false
video: "/videos/blog/the-most-privacy-friendly-one-time-payment-form-you-can-build/anonymous-online-payment.mp4"
show_cta: false
---

Privacy.Fish hat absichtlich ein sehr einfaches Geschäftsmodell: ein Produkt, ein Preis, eine Zahlung, kein Abo. Das macht es möglich, Zahlungen so zu gestalten, dass wir nach der Kontoerstellung nicht dauerhaft wissen müssen, welche Zahlung zu welchem Konto gehörte.

## Warum private Signup-Formulare schwieriger sind, als sie aussehen

Ein normales Signup-Formular beginnt oft eine lange Identitätskette: E-Mail-Verifikation, Passwort-Reset, Karten-Token, Rechnungen, Support-Tickets, Betrugsprüfung, Analytics, Marketing-Tags und Abo-Verwaltung.

Privacy.Fish will nur wissen, dass ein Konto bezahlt wurde. Wir wollen nicht dauerhaft wissen, wer dafür bezahlt hat.

## Captcha vor Benutzernamenprüfung

Die Benutzernamenprüfung ist die dynamische Schnittstelle der Website. Deshalb verlangt Privacy.Fish zuerst ein selbst gehostetes Captcha. Danach sind nur wenige Prüfversuche erlaubt, damit massenhafte Enumeration teuer und langsam wird.

## SSH Public Keys statt Passwörter

Die Anmeldung fragt nach SSH Public Keys, weil diese für den Dienst erforderlich sind. Wir fragen nicht nach Telefonnummern, Kontakt-E-Mail-Adressen oder Passwörtern, die wir später schützen müssten.

## Temporärer Zahlungscode

Nach der Reservierung erhält der Nutzer einen zufälligen Zahlungscode. Dieser Code wird bei der Zahlung angegeben, nicht der gewünschte Benutzername. Sobald Zahlung und Konto zusammengeführt wurden und die rechtliche Frist abgelaufen ist, wird die Zuordnung gelöscht.

## Fazit

Ein privates Signup-Formular entsteht nicht durch eine Datenschutz-Checkbox. Es muss um Datenminimierung herum gebaut werden: selbst gehostete Bot-Abwehr, begrenzte Benutzernamenprüfungen, SSH-Schlüssel, temporäre Zahlungscodes und Löschung der Zahlungs-zu-Konto-Verknüpfung.
