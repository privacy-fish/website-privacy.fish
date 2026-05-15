---
title: "Architektur der E-Mail-Infrastruktur"
description: "Technischer Überblick über die privacy.fish E-Mail-Infrastruktur, Vertrauensgrenzen, OpenBSD-Stack und operatives Sicherheitsmodell."
weight: 40
video: "/videos/security-model.mp4"
---

# Konzepte der Infrastruktursicherheit

Das Sicherheitsmodell von Privacy.Fish kombiniert mehrere Ideen: so wenig wie möglich vertrauen, laufende Systeme klein halten und kompromittierte Server leicht ersetzbar machen.

## Dinge, denen wir letztlich vertrauen müssen

- stw.no WebUI für Serververwaltung und VPS
- cdn.openbsd.org für Installationsabbilder
- github.com
- Admin-Workstation
- SSH Private Key der Admin-Workstation
- Client-Geräte der Nutzer

## Hosting-Provider servetheworld.no

Alle Server befinden sich in Norwegen und werden bei servetheworld.no / stw.no betrieben. Die wichtigste Eigenschaft ist die Rechtsordnung: Die Server stehen unter norwegischem Recht.

## OpenBSD als Betriebssystem

OpenBSD ist die Basis für Server und Admin-Workstations. Es ist sicherheitsorientiert, standardmäßig restriktiv und bringt viele Sicherheitsmechanismen direkt im Basissystem mit.

## OpenSSH und OpenSMTPD

Der Stack nutzt OpenSSH für Nutzerinteraktion und Administration sowie OpenSMTPD für E-Mail-Verarbeitung. Beide gehören zum OpenBSD-Ökosystem und reduzieren den Bedarf an zusätzlicher Drittsoftware.

## Minimale eigene Software

Eigener Code wird auf das reduziert, was wirklich erforderlich ist. Weniger eigener Code bedeutet weniger Angriffsfläche und einfachere Überprüfbarkeit.

## Regelmäßige Neuaufbauten

Server werden als ersetzbare Infrastruktur behandelt. Regelmäßige Neuaufbauten sollen verhindern, dass gepatchte Exploits dauerhaften Zugriff hinterlassen.
