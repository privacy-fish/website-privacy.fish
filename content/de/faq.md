---
title: "FAQ"
type: "faq"
blocks:
  - block: faq
    id: faq
    heading: "FAQ"
    subheading: "Kurze Antworten, bevor du entscheidest, ob privacy.fish zu deinem E-Mail-Workflow passt."
    layout: "stacked"
    items:
      - question: "Ist privacy.fish für alle geeignet?"
        answer: "Nein. privacy.fish ist für Menschen gebaut, die Datenschutz und Sicherheit über Bequemlichkeit stellen, und passt am besten, wenn du mit einem technischeren Workflow zurechtkommst."
      - question: "Muss ich technisch versiert sein, um privacy.fish zu nutzen?"
        answer: "Du musst kein Experte sein, solltest aber mit Kommandozeile, SSH und SFTP grundsätzlich umgehen können."
      - question: "Wie greife ich auf meine E-Mails zu?"
        answer: "Du lädst age-verschlüsselte <code class=\"font-mono text-ocean\">.eml</code>-Dateien per SFTP herunter. Ein optionales clientseitiges Werkzeug kann heruntergeladene E-Mails über POP3 an einen Mailclient wie Thunderbird bereitstellen."
      - question: "Gibt es Webmail?"
        answer: "Nein. Der Dienst ist darauf ausgelegt, verschlüsselte E-Mails auf deine eigenen Geräte herunterzuladen, statt ein klassisches Webmail-Postfach auf dem Server zu betreiben."
      - question: "Wie viele Geräte kann ich nutzen?"
        answer: "Jedes Konto unterstützt bis zu 10 SSH Public Keys, also bis zu 10 autorisierte Geräte."
      - question: "Kann ich mehr als eine E-Mail-Adresse erstellen?"
        answer: "Ja. Jedes Konto enthält eine Hauptadresse und unbegrenzt viele Aliase, einschließlich unterstützter privacy.fish-Domains wie <code class=\"font-mono text-ocean\">@pfi.sh</code>."
      - question: "Wie lange bleiben E-Mails auf dem Server?"
        answer: "E-Mails werden von dir heruntergeladen oder nach 14 Tagen automatisch gelöscht. privacy.fish ist bewusst kein dauerhaftes serverseitiges Postfacharchiv."
      - question: "Ist privacy.fish ein Abonnement?"
        answer: "Nein. Der Zugang wird mit einer einmaligen Zahlung für die Lebensdauer des Dienstes freigeschaltet."
      - question: "Was passiert, wenn ich meine SSH-Schlüssel verliere?"
        answer: "Wiederherstellung funktioniert nur über dein Backup-SSH-Schlüsselpaar. Ohne gültigen Backup-Schlüssel gibt es keinen alternativen Wiederherstellungsweg."
      - question: "Scannt ihr E-Mails oder trackt Website-Besucher?"
        answer: "Nein. privacy.fish nutzt keine Drittanbieter-Analytics oder Tracking-Skripte auf der Website und scannt deine E-Mails nicht über minimale Anti-Spam-Schutzmaßnahmen hinaus."
      - question: "Warum ist iPhone-Unterstützung nicht verfügbar?"
        answer: "iPhone wird derzeit nicht unterstützt, weil die benötigten privacy.fish-Client-Werkzeuge dort nicht so laufen können, wie der Dienst es verlangt."
---
