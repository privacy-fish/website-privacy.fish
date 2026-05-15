---
title: "Risikomanagement und Reaktion auf Kompromittierung"
description: "Wie privacy.fish über Live-Kompromittierung, Wirkungsgrenzen, operatives Risiko und Reaktion nach einem Vorfall denkt."
weight: 45
video: "/videos/secure-email-workflow.mp4"
---

# Risikomanagement und Reaktion auf Kompromittierung

Diese Seite beschreibt, wie Privacy.Fish über Kompromittierungen, Auswirkungen und Reaktion denkt.

## Wie ein erfolgreicher Angriff wirken kann

Die Auswirkungen hängen davon ab, welcher Server kompromittiert ist. Privacy.Fish trennt Rollen wie eingehende Mail, ausgehende Mail, gespeicherte verschlüsselte Mail, Schlüsselverwaltung und Monitoring.

Root-Zugriff auf eine VM ist ernst, bedeutet aber nicht automatisch Root-Zugriff auf alle anderen Rollen. Ein kompromittierter Eingangsserver könnte eingehende Mail vor der Verschlüsselung beeinflussen. Ein kompromittierter Ausgangsserver könnte ausgehende Mail beeinflussen. Ein kompromittierter Key-Management-Server könnte SSH Public Keys manipulieren.

## Reaktion

Nach einem bestätigten oder vermuteten Vorfall werden betroffene Systeme ersetzt, benötigte Daten geprüft migriert und die Ursache geschlossen. Das Modell geht davon aus, dass langlebige Server ein Persistenzrisiko darstellen.
