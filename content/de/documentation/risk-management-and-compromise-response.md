---
title: "Risikomanagement und Reaktion auf Kompromittierung"
description: "Wie privacy.fish über Live-Kompromittierung, Wirkungsgrenzen, operatives Risiko und Reaktion nach einem Vorfall denkt."
weight: 45
video: "/videos/secure-email-workflow.mp4"
---

# Risikomanagement und Reaktion auf Kompromittierung

Diese Seite beschreibt, wie Privacy.Fish über Live-Kompromittierung, Auswirkungen und Reaktion denkt. Sie ist absichtlich von der Infrastruktur-Design-Seite getrennt, weil es hier um Risikomanagement geht: was passieren kann, was erkannt werden kann und was nach einer vermuteten oder bestätigten Kompromittierung getan werden muss.

## Wie ein erfolgreicher Angriff die Privacy.Fish-Infrastruktur beeinflussen kann

Die Wirkung einer Live-Kompromittierung hängt davon ab, welcher Server kompromittiert ist. Privacy.Fish betreibt zwei physische Server mit OpenBSD. Auf diesen physischen Servern nutzen wir `vmm`, OpenBSDs Virtualisierungstool, um getrennte OpenBSD-VMs für verschiedene Teile der E-Mail-Infrastruktur zu betreiben, darunter `in`, `spam-in`, `out`, `spam-out`, `fetch`, `keys` und `monitoring`. Der Backup-Server läuft separat als VPS bei stw.no auf Proxmox, was bedeutet, dass die darunterliegende Virtualisierungsschicht nicht unter unserer Kontrolle steht. Root-Zugriff auf eine VM ist ernst, aber nicht automatisch dasselbe wie Root-Zugriff auf jede VM oder auf den physischen VMM-Host. Eine Kompromittierung nur des eingehenden Mailservers (VM) könnte eingehende Mail abfangen oder manipulieren, bevor sie verschlüsselt wird, würde aber nicht automatisch ausgehende Mail oder gespeicherte verschlüsselte Mail kompromittieren. Eine Kompromittierung nur des ausgehenden Mailservers könnte ausgehende Mail abfangen oder manipulieren, aber nicht automatisch eingehende Mail oder gespeicherte verschlüsselte Mail. Eine Kompromittierung nur des Key-Management-Servers könnte Kunden-SSH-Schlüssel manipulieren. Eine Kompromittierung des Backup- oder Migrationspfads könnte schlechten Zustand weitertragen. Eine Kompromittierung des Monitorings könnte Alarme verstecken, oder wir könnten keine Prüfungen haben, um ein bestimmtes bösartiges Verhalten zu erkennen. Eine Kompromittierung eines physischen VMM-Hosts wäre breiter als die Kompromittierung einer einzelnen Gast-VM.

Ein Angreifer mit `root`-Zugriff auf einen bestimmten Server kann alles tun, was diese Serverrolle erlaubt: Konfiguration ändern, Dateien verändern, Logs einsehen, laufende Dienste stören und Verkehr manipulieren, der durch diesen Server läuft. Gespeicherte verschlüsselte Mail ist der Sonderfall, bei dem Manipulation für den Benutzer sichtbar sein sollte: Wenn eine bestehende `.age`-Maildatei verändert wird, lässt die age-Authentifizierung die Entschlüsselung fehlschlagen. Das schützt gegen stille Veränderung bereits verschlüsselter gespeicherter Mail, aber gegen nichts anderes.
