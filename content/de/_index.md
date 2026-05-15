---
title: "privacy.fish"
description: "Ein privacy-first E-Mail-Dienst, der gespeicherte Daten, Angriffsfläche und Vertrauen minimiert."
logo: "images/logo.svg"
blocks:
  - block: hero
    id: hero
    eyebrow: "In Norwegen gehostet"
    heading: "E-Mail: Kompromisslos Privat und Sicher"
    prefix_highlight: "E-Mail:"
    highlight: "Privat und Sicher"
    subheading: "Private E-Mail-Infrastruktur mit Datenminimierung und Open-Source-Basis."
    summary_heading: "Sicher entworfen, um Privatsphäre zu schützen"
    primary_cta_text: "Konto erstellen"
    primary_cta_url: "/signup/"
    secondary_cta_text: "Code"
    secondary_cta_url: "https://github.com/privacy-fish"
    callouts:
      - value: "Privacy by Design"
        label: "Es werden nur Daten gesammelt, die für den Betrieb der Mailserver erforderlich sind. Nicht mehr benötigte Daten werden gelöscht."
      - value: "Überprüfbare Sicherheit"
        label: "Unser Code ist Open Source und kann von der Community und externen Cybersecurity-Firmen geprüft werden."
    stats:
      - value: "10 Geräte"
        label: "Pro Konto"
      - value: "0 Kompromisse"
        label: "Bei dem, was zählt"
    price:
      value: "20,00 € Einmalzahlung"
      label: "Eine Registrierungsgebühr. Keine Abos und keine wiederkehrenden Kosten."
    compat:
      id: compat
      label: "Kompatibel mit"
      unsupported_prefix: "Kein"
      items:
        - name: "Command Line"
          supported: true
        - name: "Linux"
          supported: true
        - name: "macOS"
          supported: true
        - name: "Windows"
          supported: true
        - name: "Android"
          supported: true
        - name: "iPhone-Support"
          supported: false

  - block: intro
    id: intro
    eyebrow: "Norwegische Privatsphäre"
    heading: "<span class=\"text-sun\">Privater und sicherer</span> <span class=\"text-deep\">Zugang zum Internet</span> <span class=\"text-ocean\">ist ein Menschenrecht</span>"
    lead: "Eine E-Mail-Adresse ist Voraussetzung für die Teilnahme am modernen Internet. Sie ist das Erste, wonach die meisten Dienste fragen. Ihre Privatsphäre sollte geschützt werden."
    paragraphs:
      - "Deine E-Mail-Adresse ist der Schlüssel zu einem großen Teil deines Online-Lebens. Wenn dein E-Mail-Anbieter sie verfolgt, wird deine Identität leichter profilierbar, angreifbar und offengelegt."
      - "Privacy.Fish behandelt E-Mail als private Infrastruktur, nicht als Datenquelle."
      - "Wir sammeln nur, was für den Betrieb der Mailserver erforderlich ist, veröffentlichen unseren Code und richten jede technische Entscheidung auf Datenschutz und Sicherheit aus."
    primary_cta_text: "Unsere Unternehmensphilosophie lesen"
    primary_cta_url: "/documentation/company-philosophy/"

  - block: counter
    id: counter
    items:
      - value: "∞"
        label: "Privacy-Aliase"
      - value: "20 €"
        label: "Einmalzahlung"
      - value: "10"
        label: "Geräte"
      - value: "0"
        label: "Kompromisse bei Datenschutz und Sicherheit"

  - block: features
    id: features
    eyebrow: "Was wir anders machen"
    heading: "Was uns von anderen unterscheidet"
    subheading: "Jede Designentscheidung stellt deine Privatsphäre zuerst — von Verschlüsselung im Ruhezustand bis zum Umgang mit Zahlungen und Logs."
    groups:
      - title: "Datenminimierung standardmäßig"
        video: "/videos/data-minimization-by-default.mp4"
        image_position: "left"
        items:
          - "Wir wollen keine Daten, die wir nicht benötigen, und löschen Daten lieber, als sie dauerhaft schützen zu müssen."
          - "Keine Analytics, keine Tracker, keine Website-, Mail- oder SSH-Auth-Logs."
          - "Du lädst E-Mails selbst von unseren Servern herunter und löschst sie. Nicht gelöschte E-Mails werden nach 14 Tagen gelöscht."
      - title: "Private Zahlung"
        video: "/videos/private-payment.mp4"
        video_aspect: "16:9"
        image_position: "right"
        items:
          - "Nach der Anmeldung erhältst du einen temporären Zahlungscode, der 30 Tage gültig ist."
          - "Sobald die Zahlung eingegangen und das Konto erstellt ist, wird der Code gelöscht, sodass die Zahlung nicht mehr mit dem Konto verknüpft werden kann."
          - "Wir unterstützen Bargeld per Brief, Kryptowährungen, SEPA-Überweisung, PayPal und Kreditkarten."
      - title: "Norwegische Datenschutz-Rechtsordnung"
        video: "/videos/norways-privacy-jurisdiction.mp4"
        image_position: "left"
        items:
          - "Norwegisches Recht verlangt nur die Speicherung, wann du dich mit welcher IP:Port-Kombination angemeldet hast, für 12 Monate."
          - "Andere Rechtsordnungen bringen breitere Überwachungspflichten mit sich. Die Schweiz kann zum Beispiel Sender-, Empfänger-, Protokoll-, Mailbox-Event- und Server-Metadaten verlangen, während Deutschland große E-Mail-Anbieter zu Überwachungsinfrastruktur verpflichten kann."
          - "VPN-Zugang wird empfohlen und Tor-.onion-Adressen sind für alle Server verfügbar."
      - title: "Sicherheitsmodell"
        video: "/videos/security-model.mp4"
        image_position: "right"
        items:
          - "Gebaut mit besonders sicherer Open-Source-Software (OpenBSD, OpenSMTPD, OpenSSH) und minimalem eigenen Code. Alles ist auditierbarer Open-Source-Code."
          - "Um zu verhindern, dass gepatchte Exploits dauerhaften Zugriff hinterlassen, werden alle Server wöchentlich neu aufgebaut. Nur gespeicherte E-Mails und deine SSH Public Keys werden migriert."
          - "Admin-Workstations sind OpenBSD Raspberry Pis, werden monatlich ersetzt und dürfen per Firewall nur unsere Server und stw.no erreichen."
      - title: "Sicherer E-Mail-Workflow"
        video: "/videos/secure-email-workflow.mp4"
        image_position: "left"
        items:
          - "Sicherheit vor Bequemlichkeit: kein Webmail, kein Passwort-Login, kein IMAP und kein POP3 — nur SSH- und SFTP-Zugriff über unsere Client-App."
          - "Deine E-Mails werden mit age und deinen SSH Public Keys verschlüsselt gespeichert und nach 14 Tagen oder nach deiner Löschung sicher entfernt."
          - "Wenn der Mailserver eines Empfängers seine Identität nicht mit einem gültigen TLS-Zertifikat beweisen kann, fragt die App, ob du abbrechen oder trotzdem senden möchtest."

  - block: audience
    id: audience
    eyebrow: "Solltest du es nutzen?"
    heading: "Privatsphäre erfordert eine Entscheidung"
    for_title: "Gebaut für"
    for_items:
      - "Du bist datenschutzbewusst und willst einen E-Mail-Anbieter, der so wenig Daten wie möglich sammelt."
      - "Du möchtest nicht, dass deine E-Mails oder persönlichen Daten dauerhaft auf fremden Servern liegen."
      - "Du bevorzugst einen kleinen, fokussierten Dienst, der ausschließlich für bestmöglichen Datenschutz und Sicherheit gebaut ist."
    not_for_title: "Nicht gebaut für"
    not_for_items:
      - "Du brauchst einen Anbieter, der deine E-Mails dauerhaft speichert, statt sie auf deine Geräte herunterzuladen."
      - "Du brauchst Browser-Webmail, auch wenn dadurch mehr Metadaten oder lesbare E-Mails beim Anbieter landen."
      - "Du brauchst Extras wie Kalender, Kontakte, Cloud-Speicher oder Office-Tools, auch wenn sie zusätzliche angreifbare Software bedeuten."

  - block: how-it-works
    id: how-it-works
    eyebrow: "So funktioniert es"
    heading: "Von der Anmeldung bis zum Posteingang"
    subheading: "Nach Ablauf der gesetzlichen Verbraucher-Widerrufsfrist von 14 Tagen löschen wir dauerhaft alle Informationen darüber, welches Konto mit welcher Zahlung verbunden war."
    steps:
      - icon: "key"
        title: "App nutzen und Schlüsselpaare auf deinen Geräten erzeugen"
        text: "Installiere die Open-Source-App von Privacy.Fish und erzeuge SSH-Schlüsselpaare für jedes Gerät. Kopiere die Public Keys auf das Gerät, mit dem du das Anmeldeformular ausfüllst."
      - icon: "credit-card"
        title: "Privat anmelden und bezahlen"
        text: "Kopiere deine Public Keys in das Anmeldeformular und nutze unsere datenschutzfreundlichen Zahlungsmethoden. Wir erstellen dein Konto innerhalb von 24 Stunden nach Zahlung."
      - icon: "desktop"
        title: "App mit deinem E-Mail-Client verbinden"
        text: "Richte deine E-Mail-App auf die lokalen SMTP- und POP3-Ports dieser App aus. Sie übernimmt Senden, Abrufen und Entschlüsseln deiner E-Mails über unsere Server."

  - block: tradeoffs
    id: tradeoffs
    eyebrow: "Unser Vertrauensmodell"
    heading: "Du überprüfst alles"
    subheading: "„Every secret creates a potential failure point.“ — [Bruce Schneier](https://en.wikipedia.org/wiki/Bruce_Schneier 'Bruce Schneier')"
    items:
      - icon: "code"
        title: "Alles Open Source"
        text: "Alle Software für Aufbau und Administration von Privacy.Fish ist öffentlich auf [github](https://github.com/privacy-fish 'github'), damit Infrastruktur geprüft statt blind vertraut werden kann."
      - icon: "scale-balanced"
        title: "Norwegische Rechtsordnung"
        text: "Wir haben Norwegen wegen der für unseren Dienst besten datenschutzfreundlichen Gesetze gewählt, auch wenn dort die Steuern recht hoch sind. Privatsphäre kommt zuerst."
      - icon: "coins"
        title: "Gleiche Anreize"
        text: "Privacy.Fish wird durch eine einmalige Kontogebühr finanziert, nicht durch Werbung, Tracking, Abos oder Verkauf von Datenzugriff."

  - block: pricing
    id: pricing
    eyebrow: "Preise"
    heading: "Einmalzahlung"
    subheading: "Keine Abos und keine versteckten Gebühren. Die Zuordnung von Zahlung zu Konto wird nach Kontoerstellung gelöscht."
    price: "20"
    badge: "Einmalzahlung"
    note: "Für lebenslangen Unternehmenszugang"
    cta_text: "Konto erstellen"
    cta_url: "/signup/"
    provisioning_title: "Manuelle Bereitstellung"
    provisioning_copy: "Konten werden innerhalb von 24 Stunden nach Zahlung von einem Administrator erstellt."
    included_heading: "Enthalten"
    features:
      - "Hauptbenutzername, 10 zufällige Aliase und unbegrenzt rotierende Aliase"
      - "Alle unsere Domains funktionieren mit deinem Hauptbenutzernamen und allen Aliasen"
      - "Maximal 10 Geräte pro Konto"
    payment_label: "Zahlungsmethoden:"
    payment_methods:
      - "Bargeld per Brief"
      - "Kryptowährungen"
      - "SEPA-Überweisung"
      - "Kreditkarte"
      - "PayPal"
    payment_methods_joiner: " und "
    refund_label: "Erstattungsinformation:"
    refund_information: "Erstattungen sind nur möglich, bevor das Konto erstellt wurde, weil die Zahlungs-zu-Konto-Zuordnung direkt danach zerstört wird."
---
