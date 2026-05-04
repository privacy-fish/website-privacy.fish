---
title: "Jurisdiction"
description: "Why Norway was selected for privacy.fish and how jurisdiction affects logging, disclosure, and legal exposure."
weight: 70
video: "/videos/norways-privacy-jurisdiction.mp4"
---


## Introduction

Privacy.Fish compared jurisdictions by one narrow question: what would the state force us to log as an email provider?

Every required log becomes data we must collect, protect, retain, secure, and potentially disclose. Norway was the best jurisdiction we found for the Privacy.Fish model because its mandatory retained data is narrow, explicit, and predictable: account, source IP:port, and timestamp for 12 months.

Norway is also clear about what this duty is not supposed to include. Under the Norwegian IP-retention duty, destination information must not be retained. This is the important difference: Norway requires a narrow access-identification log, not broad mail transaction logging.

Sources:
- [Norwegian Electronic Communications Act, unofficial English translation by Nkom](https://nkom.no/english/guidelines-for-providers/_/attachment/download/cccdd933-8254-4024-820e-168a96f071f0%3Ac6fa2c909f0187b6c31b6f8c0f40f328b7ce46a8/Electronic%20Communications%20Act%20%28Ekom%20Act%29%20%E2%80%93%20Unofficial%20Translation.pdf)
- [Norwegian Electronic Communications Regulation, unofficial English translation by Nkom](https://nkom.no/english/guidelines-for-providers/_/attachment/download/13dc9675-b0f4-42e0-840e-dba0c735db69%3A470133b1352c9038804c5c0b0d7ed967c13c9e0f/Regulation%20on%20Electronical%20Communications%20Network%20and%20Services%20%E2%80%93%20Unofficial%20version.pdf)
- [Nordic data-retention comparison: Norway](https://pub.norden.org/temanord2024-532/9-norway.html)

## EU

The EU does not currently have one single EU-wide data-retention law that tells every email provider exactly what to log.

The old EU Data Retention Directive required providers to retain traffic and location data for services including internet access, internet email, and internet telephony. In 2014, the Court of Justice of the European Union declared that directive invalid because it interfered too seriously with the rights to privacy and data protection.

Since then, the legal situation in the EU has been fragmented. Each member state has its own national retention rules, court decisions, national-security exceptions, and telecommunications laws. Some countries have no general retention rules currently in force, some have targeted retention, and some still impose broad or complex metadata-retention obligations.

For Privacy.Fish, this means the EU itself is not enough of an answer. Each country has to be checked separately.

Why Norway is better:
- Norway gives one narrow and explicit rule we can design around.
- EU countries vary heavily.
- Some EU countries have broader retention duties than Norway.
- Some EU countries have less blanket retention, but more legal uncertainty.
- The EU situation is unstable enough that “hosted in the EU” is not precise privacy information.

Sources:
- [CJEU press release declaring the Data Retention Directive invalid](https://curia.europa.eu/jcms/upload/docs/application/pdf/2014-04/cp140054en.pdf)
- [European Commission page on data retention](https://home-affairs.ec.europa.eu/policies/internal-security/lawful-access-data/data-retention_en)
- [EU Data Retention Directive FAQ](https://europa.eu/rapid/press-release_MEMO-14-269_en.htm)
- [Eurojust report on CJEU data-retention case law](https://www.eurojust.europa.eu/sites/default/files/assets/files/data-retention-report-cjeu-eurojust-13-11-2024.pdf)
- [Council document on future EU data-retention rules](https://data.consilium.europa.eu/doc/document/WK-11640-2025-INIT/en/pdf)
- [Cullen International update on national data-retention laws in Europe](https://www.cullen-international.com/news/2025/07/Update-on-national-data-retention-laws-in-Europe.html)

## Norway

Norway is the winner for Privacy.Fish.

Norwegian law requires providers to retain information needed to identify which subscriber used a public IP address at a given time. Where needed, this includes the port number. In the Privacy.Fish model, this becomes one narrow access log: account, source IP:port, and timestamp.

The retention period is 12 months.

The Norwegian rule also says that destination information shall not be retained. That matters because it means the law is about identifying account access from an IP address and time, not about keeping a general history of who you wrote to, which server your mail was delivered to, or what destinations you contacted.

Why Norway is better:
- Required retained data is narrow.
- The retention period is explicit.
- Destination logging is explicitly excluded under this duty.
- The rule is understandable enough to design around.
- It does not require routine sender, recipient, mailbox-event, message-ID, delivery-status, or destination-server logging for this model.

Sources:
- [Norwegian Electronic Communications Act, unofficial English translation by Nkom](https://nkom.no/english/guidelines-for-providers/_/attachment/download/cccdd933-8254-4024-820e-168a96f071f0%3Ac6fa2c909f0187b6c31b6f8c0f40f328b7ce46a8/Electronic%20Communications%20Act%20%28Ekom%20Act%29%20%E2%80%93%20Unofficial%20Translation.pdf)
- [Norwegian Electronic Communications Regulation, unofficial English translation by Nkom](https://nkom.no/english/guidelines-for-providers/_/attachment/download/13dc9675-b0f4-42e0-840e-dba0c735db69%3A470133b1352c9038804c5c0b0d7ed967c13c9e0f/Regulation%20on%20Electronical%20Communications%20Network%20and%20Services%20%E2%80%93%20Unofficial%20version.pdf)
- [Nordic data-retention comparison: Norway](https://pub.norden.org/temanord2024-532/9-norway.html)

## Switzerland

Switzerland is worse than Norway for the Privacy.Fish model.

Switzerland has a strong privacy reputation, but its surveillance and metadata-retention framework can bring much broader email metadata into scope than Norway. Swiss law has long required certain telecommunications metadata to be retained for six months. Swiss official material describes retained metadata in general as information about who communicated with whom, when, for how long, where, and by what technical means.

For email-related services, Swiss surveillance rules can include metadata such as sender and recipient addresses, protocol, mailbox events, delivery status, and sending or receiving mail-server information, depending on provider classification and the concrete obligation.

Switzerland is also more legally complex. Some pure email or over-the-top communication services may have more limited duties than traditional telecommunications providers. But that complexity is itself a downside compared with Norway’s narrow and explicit rule.

Why Norway is better:
- Norway requires account, source IP:port, and timestamp.
- Switzerland can require broader email metadata.
- Norway explicitly excludes destination information under the IP-retention duty.
- Switzerland’s rules create more classification and compliance uncertainty.
- Switzerland is harder to explain honestly as a simple data-minimizing jurisdiction for this service model.

Sources:
- [Swiss PTSS FAQ on metadata retention](https://www.li.admin.ch/en/documentation/faq)
- [Swiss BÜPF information page](https://www.li.admin.ch/de/themen/das-buepf)
- [Swiss VÜPF ordinance text on Fedlex](https://www.fedlex.admin.ch/eli/cc/2018/32/de)
- [Swiss ordinance text via LexFind](https://www.lexfind.ch/tolv/149472/de)
- [Laux Lawyers analysis on Swiss surveillance duties for messaging services](https://www.lauxlawyers.ch/ueberwachung-von-instant-messaging-bundesgericht-schafft-klarheit/)
- [Swico guide on authority requests and surveillance orders](https://www.swico.ch/media/filer_public/78/f1/78f151c6-ff8b-4ecf-b856-c1530e1a3fd5/leitfaden_behordenanfragen_3_anordnung_zur_fernmeldeuberwachung.pdf)
- [Internet Society Switzerland criticism of Swiss surveillance ordinance](https://www.isoc.ch/swiss-surveillance-ordinance-encryption-threat-vupf-oscpt/)
- [Digitale Gesellschaft: data retention in Switzerland](https://www.digitale-gesellschaft.ch/vds-suisse/index_en.html)

## Germany

Germany is worse than Norway for the Privacy.Fish model.

Germany’s main problem is not only data retention. The bigger problem is lawful-interception infrastructure.

German regulator BNetzA says operators of email systems are exempt from the TKÜV readiness obligation only as long as no more than 100,000 users are connected to the system. Above that threshold, email operators must maintain technical and organizational readiness for lawful interception.

Germany also has technical rules for handover points and full recording of lawful intercepts under TR TKÜV. BNetzA describes TR TKÜV as the technical directive that defines requirements for full records of telecommunications intercepts, provision of information, and configuration of handover points to authorized organizations.

Germany also has specific requirements for electronic authority interfaces. BNetzA says obligated providers with 100,000 or more contractual partners must maintain both an ETSI-based interface and an email-based transmission method for certain information procedures, while smaller obligated providers must keep the email-based procedure available.

Why Norway is better:
- Norway requires one narrow access log.
- Germany can require large email providers to become interception-ready.
- Germany has handover/interface requirements that Privacy.Fish does not want to design around.
- Germany’s retention rules have repeatedly been politically and legally contested.
- Norway is simpler: account, source IP:port, timestamp.

Sources:
- [BNetzA: email surveillance obligations and 100,000-user threshold](https://www.bundesnetzagentur.de/DE/Fachthemen/Telekommunikation/OeffentlicheSicherheit/Ueberwachung_Auskunftsert/EMailUeberwachung/node.html)
- [BNetzA: intercepts and provision of information](https://www.bundesnetzagentur.de/EN/Areas/Telecommunications/ServicerProviderObligation/PublicSafety/Intercepts/start.html)
- [BNetzA: TR TKÜV supplementary information](https://www.bundesnetzagentur.de/EN/Areas/Telecommunications/ServicerProviderObligation/PublicSafety/Intercepts/SupplementaryInformation/SupplementaryInformation_node.html)
- [BNetzA: electronic authority interface thresholds](https://www.bundesnetzagentur.de/DE/Fachthemen/Telekommunikation/OeffentlicheSicherheit/Ueberwachung_Auskunftsert/BDA/artikel.html)
- [German TR TKÜV draft text via EU TRIS](https://technical-regulation-information-system.ec.europa.eu/de/notification/26192/text/D/EN)
- [Statewatch note on German interception obligations for email providers](https://www.statewatch.org/statewatch-database/germany-telecommunication-providers-to-intercept-customers/)
- [Lexology note on German classification of internet-based email services](https://www.lexology.com/library/detail.aspx?g=62df0a70-3b37-4a27-a7e7-dcf41ba2e469)

## Netherlands

The Netherlands is not a simple “worse than Norway” case. It is more complicated.

The old Dutch blanket data-retention law was suspended by the District Court of The Hague in 2015 after the CJEU invalidated the EU Data Retention Directive. As a result, Dutch telecom and internet providers were no longer required under that law to retain general telephone and internet traffic data for fixed periods. The old Dutch regime had required telecom data retention for at least 6 months, and commentary describes the old requirement as including 12 months for telephony data and 6 months for internet access and email-related data.

That means the Netherlands may look attractive if the only question is blanket metadata retention. But the absence of the old blanket-retention law does not mean “no surveillance obligations.” Dutch public telecommunications providers can still be subject to lawful-interception and disclosure obligations. NBIP describes lawful interception and lawful disclosure as obligations for ISPs, VoIP providers, and other public telecommunications services, where authorized agencies can demand traffic interception or customer-identifying data.

This is the key difference from Norway for Privacy.Fish: the Netherlands may currently have less simple blanket retention, but the legal model is less clean to explain. Norway gives us a narrow, explicit, public rule: account, source IP:port, and timestamp for 12 months. The Netherlands requires a deeper classification question: whether the service is treated as a public electronic communications service, what interception or disclosure obligations apply, and what operational readiness is expected.

Why Norway is better:
- Norway gives a narrow, explicit retention duty we can design around.
- The Netherlands does not currently have the same old blanket retention law, but lawful-interception and disclosure obligations still exist.
- The Dutch answer depends more on provider classification and operational interpretation.
- Norway is easier to explain honestly to users: one required access log, not a broader or unclear mail-metadata regime.

Sources:
- [Dutch court invalidated the national data-retention law](https://peacepalacelibrary.nl/blog/2015/data-retention-saga-dutch-court-declared-national-data-retention-law-invalid)
- [Privacy First summary of the telecom retention case](https://privacyfirst.nl/en/court-cases/telecom-retention-obligation/)
- [Lexology summary of Dutch data-retention suspension](https://www.lexology.com/library/detail.aspx?g=0860c791-fc73-435a-a243-1701a03a7265)
- [NBIP lawful interception explanation](https://www.nbip.nl/en/lawful-interception/)
- [Dutch government business page for telecom-provider requirements](https://business.gov.nl/regulations/requirements-telecom-providers/)
- [PILP dossier on the Telecommunications Retention Act](https://pilp.nu/en/dossier/the-telecommunications-retention-act/)

## Sweden

Sweden is worse or less clean than Norway for the Privacy.Fish model.

Swedish retention rules are broader and more category-dependent. The Nordic comparison describes Swedish retention around data needed to trace and identify source and destination, date, time, duration, type of communication, user equipment, and location for covered services.

For internet access, Sweden retains data such as subscriber or user data, assigned IP address, login and logout times, and equipment data. Retention periods vary by category. Data related to internet access is stored for 10 months, while data related to telephony or message processing through a mobile termination point is generally stored for 6 months. Certain location data is stored for a shorter period.

This is less suitable for Privacy.Fish because the retained categories are broader and more dependent on service classification than Norway’s account, source IP:port, timestamp model.

Why Norway is better:
- Norway’s retained fields are narrower.
- Sweden’s rules include broader categories.
- Sweden’s retention periods vary by communication type.
- Norway explicitly excludes destination information under the IP-retention duty.
- Norway is easier to explain as a privacy-first jurisdiction.

Sources:
- [Nordic data-retention comparison: Sweden](https://pub.norden.org/temanord2024-532/10-sweden.html)
- [Nordic data-retention comparison: Nordic overview](https://pub.norden.org/temanord2024-532/5-nordic-overview.html)

## Denmark

Denmark is legally heavier and more complex than Norway.

Denmark historically imposed broad data retention. After legal changes and CJEU pressure, Denmark moved toward a more complicated model involving targeted retention, regional retention, and general or undifferentiated retention in certain national-security contexts.

The Nordic comparison describes targeted retention aimed at convicted persons, communication equipment, persons, geographical areas, and concrete assessments. Denmark also has rules around general retention for national security and internet access. Retention periods can be around one year depending on the activated rule.

This is a bad fit for Privacy.Fish because the answer is not a simple field list. The Danish regime depends on targeting, geography, crime categories, national-security assessments, and implementing decisions.

Why Norway is better:
- Norway has a narrow, general rule Privacy.Fish can design around.
- Denmark’s regime is more complex.
- Denmark can involve targeted, regional, and national-security retention models.
- Norway gives users a clearer answer about what we are forced to retain.

Sources:
- [Nordic data-retention comparison: Denmark](https://pub.norden.org/temanord2024-532/6-denmark.html)
- [IT-Pol: the new Danish data-retention law](https://itpol.dk/articles/new-Danish-data-retention-law-2022)
- [European Parliament document on Danish data retention](https://www.europarl.europa.eu/meetdocs/2014_2019/plmrep/COMMITTEES/LIBE/DV/2020/01-20/Danish_data_retention_EN.pdf)
- [Eurojust report on CJEU data-retention case law](https://www.eurojust.europa.eu/sites/default/files/assets/files/data-retention-report-cjeu-eurojust-13-11-2024.pdf)

## Finland

Finland may be lighter for some small providers, but it is less clean than Norway.

Finland’s retention rules under FECA apply to selected provider categories. The Ministry of the Interior decides which companies are subject to the retention obligation. That means Finland may not cover every small provider in practice, but the rule is less predictable and more administrative than Norway’s narrow statutory rule.

If a provider is covered, the retained categories can be broader than Norway’s access log. The Nordic comparison describes categories such as subscriber name and address, subscription identifiers, transaction-identifying data, message type, recipient, time and duration, and internet-access user or device identifiers.

Retention periods vary by service category: 12 months for mobile and text-message data, 6 months for internet telephony, and 9 months for internet access.

Why Norway is better:
- Finland may not cover every provider, but the coverage question is less predictable.
- If covered, Finland’s retained categories are broader.
- Finland uses different retention periods by service type.
- Norway gives one clearer rule for this model.

Sources:
- [Nordic data-retention comparison: Finland](https://pub.norden.org/temanord2024-532/7-finland.html)
- [Nordic data-retention comparison: Nordic overview](https://pub.norden.org/temanord2024-532/5-nordic-overview.html)

## Iceland

Iceland is worse than Norway for the Privacy.Fish model.

Iceland has a shorter retention period than Norway, but the scope appears broader. Icelandic telecommunications companies must retain a minimum record of users’ electronic-communications traffic for 6 months.

The Nordic comparison says providers must be able to identify which customer used a telephone number, IP address, or username, and provide information on connections, dates, who was connected, and amount of data transferred.

That is broader than Norway’s model. Norway requires account, source IP:port, and timestamp; Iceland appears to require broader connection records.

Why Norway is better:
- Iceland’s retention period is shorter, but the scope appears broader.
- Iceland can include information about connections and counterparties.
- Norway’s duty is narrower.
- Norway explicitly excludes destination information under the IP-retention duty.

Sources:
- [Nordic data-retention comparison: Iceland](https://pub.norden.org/temanord2024-532/8-iceland.html)
- [Iceland Electronic Communications Act No. 70/2022](https://fjarskiptastofa.is/library?itemid=175f315e-40a7-4b2e-83b0-4c377b96b098)
- [Freedom House note on Icelandic data retention](https://freedomhouse.org/country/iceland/freedom-net/2024)

## United Kingdom

The United Kingdom is worse than Norway for the Privacy.Fish model.

The UK can require data retention by notice. A data-retention notice under the Investigatory Powers Act 2016 may require retention of relevant communications data for up to 12 months.

The categories can be much broader than Norway’s narrow access log. UK material describes relevant communications data as data that may identify or assist in identifying senders, recipients, sources, destinations, IP addresses, ports, times, durations, communication type, method, pattern, system data, and location-type information.

The UK also has secrecy and notice-based powers. That is a bad fit for a transparent privacy-first email provider, because the obligation may not be a simple public rule users can understand.

Why Norway is better:
- Norway’s duty is public, narrow, and explicit.
- The UK can impose broader retention duties by notice.
- The UK system can involve secrecy.
- UK categories can include source and destination information.
- Norway is easier to describe honestly to users.

Sources:
- [UK Notices Regime Code of Practice](https://www.gov.uk/government/publications/notices-regime-code-of-practice/notices-regime-code-of-practice-accessible)
- [UK Investigatory Powers Act 2016, section 87](https://www.legislation.gov.uk/ukpga/2016/25/section/87)
- [UK Communications Data Code of Practice, 2025 PDF](https://assets.publishing.service.gov.uk/media/68419a3d58bd5a53d0211cf7/Communications_Data_Code_of_Practice_-_June_2025.pdf)
- [UK Communications Data Code of Practice, older PDF](https://assets.publishing.service.gov.uk/media/5bf43801ed915d183c9c53ba/Communications_Data_Code_of_Practice.pdf)

## United States

The United States is not a good fit, even though it does not have a simple Norway-style blanket email-provider retention rule.

The US Stored Communications Act allows government entities to compel disclosure of stored content and non-content records with legal process. It also allows preservation requests requiring a provider to preserve records and other evidence already in its possession for 90 days, renewable for another 90 days.

This means the US may look lighter if the only question is “is there a standing universal access-log rule?” But for Privacy.Fish, the broader compelled-disclosure, preservation, gag-order, and surveillance environment makes it less attractive.

The US model is less about one public mandatory log and more about a broad legal process framework that can compel disclosure or preservation of what the provider has.

Why Norway is better:
- Norway’s rule is narrow and explicit.
- The US has broad compelled-disclosure and preservation powers.
- The US legal environment is less predictable for a privacy-first email provider.
- Norway is a better fit for a transparent public promise about what is routinely retained.

Sources:
- [18 U.S.C. § 2703, Required disclosure and preservation under the Stored Communications Act](https://www.law.cornell.edu/uscode/text/18/2703)
- [18 U.S.C. § 2704, Backup preservation](https://www.law.cornell.edu/uscode/text/18/2704)
- [47 U.S.C. § 1002, CALEA assistance capability requirements](https://www.law.cornell.edu/uscode/text/47/1002)
- [ACLU article on email preservation requests](https://www.aclu.org/news/privacy-technology/government-cannot-force-e-mail-companies-copy-and-save-your)
- [EFF overview of mandatory data retention proposals](https://www.eff.org/issues/mandatory-data-retention)

## Canada

Canada is not as predictable as Norway.

Canada does not currently have the same simple, standing, Norway-style email-provider access-log rule in the sources reviewed here. However, Canadian lawful-access law has been moving toward stronger provider obligations.

Bill C-22 proposes powers allowing regulations for retention of metadata, including transmission data, for reasonable periods not exceeding one year. The Canadian government describes this as metadata retention and says it would not include content, web-browsing history, or social media activity.

For Privacy.Fish, this is less attractive than Norway because the Canadian regime is active, politically contested, and moving toward broader metadata powers.

Why Norway is better:
- Norway already has a narrow, known rule.
- Canada’s regime is politically active and evolving.
- Proposed Canadian metadata-retention powers could become broader than Norway’s access-log model.
- Privacy.Fish benefits from a jurisdiction where the required retained data is already clear.

Sources:
- [Canadian Bill C-22 text](https://www.parl.ca/DocumentViewer/en/45-1/bill/C-22/first-reading)
- [Public Safety Canada backgrounder on Bill C-22 metadata retention](https://www.canada.ca/en/public-safety-canada/news/2026/03/backgrounder--securing-access-to-information-in-bill-c-22.html)
- [Canadian government note explaining preservation orders are not general data retention](https://www.canada.ca/en/news/archive/2013/11/modernizing-criminal-code.html)
- [Michael Geist analysis of Bill C-22 metadata retention](https://www.michaelgeist.ca/2026/03/the-lawful-access-privacy-risks-unpacking-bill-c-22s-expansive-metadata-retention-requirements/)

## Australia

Australia is much worse than Norway.

Australia has one of the clearest bad-fit regimes for Privacy.Fish. Telecommunications providers must retain a prescribed metadata dataset for at least 2 years.

The retained dataset can include subscriber/account/service/device information, source identifiers such as email address, IP address, and port number, destination identifiers such as recipient email address, IP address, and phone identifiers, time and duration, communication type, service used, and location information for equipment.

Australian guidance also says that if a provider offers an email service, the provider has data-retention obligations unless the email service is only for a person’s immediate circle.

Why Norway is better:
- Norway requires one narrow access log for 12 months.
- Australia requires much broader metadata for at least 2 years.
- Australia includes destination identifiers.
- Norway explicitly excludes destination information under the Norwegian IP-retention duty.
- Australia is the opposite of what Privacy.Fish wants to build.

Sources:
- [Australian Home Affairs data-retention obligations](https://www.homeaffairs.gov.au/about-us/our-portfolios/national-security/lawful-access-telecommunications/data-retention-obligations)
- [Australian prescribed data set](https://www.homeaffairs.gov.au/nat-security/files/dataset.pdf)
- [Australian data-retention industry FAQ](https://www.homeaffairs.gov.au/nat-security/files/data-retention-industry-faqs.pdf)
- [Australian data-retention guidelines for service providers](https://www.homeaffairs.gov.au/nat-security/files/data-retention-guidelines-service-providers.pdf)
- [Australian Parliamentary report on data-retention period](https://www.aph.gov.au/-/media/02_Parliamentary_Business/24_Committees/244_Joint_Committees/PJCIS/DataRetention2014/Chapter4.pdf)

## New Zealand

New Zealand is not better than Norway for this model.

New Zealand does not appear to have a simple blanket email-provider retention duty like Australia’s in the sources reviewed here. However, New Zealand’s TICSA framework imposes interception-capability and network-security obligations on telecommunications network operators and service providers.

TICSA is about interception capability and network security. For Privacy.Fish, the issue is not only routine logging. It is whether the legal environment expects providers to design infrastructure around interception capability.

Why Norway is better:
- Norway requires a narrow retained access log.
- New Zealand has interception-capability obligations.
- Interception-readiness obligations are a worse fit for Privacy.Fish’s security philosophy.
- Norway is easier to explain as a data-minimizing jurisdiction.

Sources:
- [New Zealand NCSC: About TICSA](https://www.ncsc.govt.nz/what-we-do/regulations-and-standards/ticsa-network-security/about-ticsa/)
- [New Zealand Police TICSA overview](https://www.police.govt.nz/advice/businesses-and-organisations/ticsa)
- [New Zealand TICSA legislation](https://www.legislation.govt.nz/act/public/2013/0091/latest/DLM5177923.html)
- [New Zealand Interception Guide for TICSA](https://www.tcf.org.nz/industry-hub/industry-codes/interception-guide-for-ticsa)
- [New Zealand Privacy Principle 9: retention of personal information](https://www.privacy.org.nz/privacy-principles/9/)

## France

France is worse than Norway.

France has broader traffic and location data retention mechanisms, especially in national-security contexts. French electronic-communications law and legal analysis describe obligations for electronic communications operators to retain connection or traffic data, often for one year, in order to respond to law-enforcement or national-security needs.

French law also has extensive national-security and lawful-access machinery around electronic communications. Even if content is not the ordinary target of metadata retention, the metadata categories and legal context are broader than Norway’s narrow access-identification log.

Why Norway is better:
- Norway’s retained data is narrower.
- Norway’s rule is more explicit and easier to design around.
- France’s national-security retention model is broader.
- France is less suitable for a minimal email provider trying to retain only one narrow access log.

Sources:
- [Global Network Initiative: France country legal framework](https://clfr.globalnetworkinitiative.org/country/france/)
- [Lexology overview of telecom and internet access in France](https://www.lexology.com/library/detail.aspx?g=b74a9239-bb45-4354-a6a1-821884f1ca17)
- [ICLG France telecoms, media and internet overview](https://iclg.com/practice-areas/telecoms-media-and-internet-laws-and-regulations/france)
- [Legal 500 France TMT PDF](https://www.legal500.com/guides/wp-content/uploads/sites/1/2025/08/France-TMT.pdf)
- [ISDC surveillance of telecommunications providers report](https://www.isdc.ch/media/2668/25-059-e-final-report-20251002_sign%C3%A9.pdf)
- [CJEU/French data-retention case note via eucrim](https://eucrim.eu/news/cjeu-french-legislation-on-data-retention-for-the-purpose-of-combating-market-abuse-offences-unlawful/)

## Belgium

Belgium is worse or at least much more complex than Norway.

Belgium has a revised targeted and differentiated metadata-retention regime. BIPT describes Article 107/5 of the 2005 Electronic Communications Act as amended by the Act of 20 July 2022 on the collection and retention of identification data and metadata in the electronic communications sector and the provision of such data to authorities.

Belgian data-retention law has been repeatedly litigated and revised. That creates more legal complexity than Norway’s narrow rule.

The Belgian system is not a simple “one access log” model. It is a broader metadata-retention and evidence-provision framework with multiple moving parts and repeated court challenges.

Why Norway is better:
- Norway has one narrow access-log duty.
- Belgium has broader metadata-retention machinery.
- Belgium has repeated legal uncertainty around retention rules.
- Norway is easier to document clearly for users.

Sources:
- [BIPT: provision of electronic evidence to authorities](https://www.bipt.be/operators/telecommunications/security/provision-of-electronic-evidence)
- [EDRi/Liga letter on Belgian draft data-retention law](https://edri.org/wp-content/uploads/2022/06/EDRi-Liga-Letter-Draft-Law-Data-Retention-BE_EN.pdf)
- [EuroISPA note on Belgian data-retention uncertainty](https://www.euroispa.org/2024/10/data-retention-rules-in-belgium-uncertainty-remains-after-third-constitutional-court-ruling/)
- [NautaDutilh Belgium cybersecurity PDF mentioning 2022 data-retention act](https://www.nautadutilh.com/documents/32/2023_Lexology_GTDT_Cybersecurity_-_Belgium6212859.1_1.pdf)
- [Telenet Law Enforcement Disclosure Report 2024](https://www2.telenet.be/content/dam/www-telenet-corp/duurzaamheidsdocumenten-2025/Telenet%20group%20NV_Law%20Enforcement%20Disclosure%20Report_2024.pdf)

## Ireland

Ireland is legally complex and historically broader than Norway.

Ireland’s Communications (Retention of Data) Act 2011 included internet access, internet email, and internet telephony categories. The retained data included data necessary to trace and identify the source of a communication, such as user IDs, and other fields related to internet access, internet email, and internet telephony.

The Irish framework used a one-year retention period for internet-related data and a two-year period for telephone data. The body of email messages was not the target; the obligation concerned traffic-related information. But this is still broader than Norway’s narrow account, source IP:port, timestamp model.

Ireland’s framework has also been heavily affected by EU case law and later amendments. That makes it less clean for Privacy.Fish than Norway.

Why Norway is better:
- Norway’s required retained data is narrower and clearer.
- Ireland’s framework historically covered internet email metadata.
- Ireland’s framework is more complex and legally contested.
- Norway gives a simple public answer about what Privacy.Fish must retain.

Sources:
- [Irish Communications Retention of Data Act 2011](https://www.irishstatutebook.ie/eli/2011/act/3/enacted/en/html)
- [Irish Communications Retention of Data Act 2011, Schedule 2](https://www.irishstatutebook.ie/eli/2011/act/3/schedule/2/enacted/en/html)
- [Irish revised act text](https://revisedacts.lawreform.ie/eli/2011/act/3/revised/en/html)
- [Irish Revenue manual on the Communications Retention of Data Act](https://www.revenue.ie/en/tax-professionals/tdm-wm/investigations-prosecutions-enforcement/enforcement/communications-retention-of-data-act.pdf)
- [Irish Legal News analysis of new data-retention laws](https://www.irishlegal.com/articles/michael-madden-new-data-retention-laws-for-ireland)
- [DWT analysis of Irish internet traffic data retention](https://www.dwt.com/insights/2011/02/ireland-requires-storage-of-internet-traffic-data)
- [Independent Examiner review of Irish retention law](https://independentexaminer.ie/wp-content/uploads/2025/02/236667_a01fdf60-6d1c-43cb-be6b-9440b8d36d6f-3.pdf)

## Luxembourg

Luxembourg is unclear and not a stronger choice than Norway.

Luxembourg has been working on traffic and location data-retention legislation after EU case-law developments. Government sources describe a draft law regulating the retention and use of traffic and location data for national security, serious crime, and serious threats to public security.

That makes Luxembourg difficult to use as a clean privacy jurisdiction for the Privacy.Fish model. The relevant framework is not as straightforward as Norway’s narrow access-retention duty.

Why Norway is better:
- Norway has a known and narrow rule.
- Luxembourg’s retention framework is less clear for the Privacy.Fish model.
- Luxembourg’s draft-law environment is harder to explain publicly.
- Privacy.Fish needs a jurisdiction where the required retained data can be described precisely.

Sources:
- [Luxembourg government announcement on traffic and location data-retention draft law](https://me.gouvernement.lu/en/actualites.gouvernement2024%2Ben%2Bactualites%2Btoutes_actualites%2Bcommuniques%2B2023%2B01-janvier%2B25-tanson-loi-retention-donnees-caractere-personnel.html)
- [DLA Piper Luxembourg data-protection law overview mentioning draft bill 8148](https://www.dlapiperdataprotection.com/index.html?c=LU&t=law)
- [Luxembourg CNPD opinion on Bill 8148](https://cnpd.public.lu/en/actualites/national/2024/05/mai-2024-avis-retention-dp.html)
- [Lexology note on Luxembourg retention-law reform](https://www.lexology.com/pro/content/luxembourg-set-reform-retention-law)
- [Luxembourg Times note on electronic data-retention reform](https://www.luxtimes.lu/luxembourg/luxembourg-moves-to-limit-electronic-data-it-keeps/1358023.html)

## Austria

Austria is interesting, but not as straightforward as Norway.

Austria’s implementation of the old EU Data Retention Directive was struck down by the Austrian Constitutional Court in 2014 after the CJEU invalidated the directive. The court criticized the mass interference with fundamental rights and the conditions around storage, deletion, and access.

This makes Austria potentially attractive if only looking for absence of blanket data retention. However, Austria still has lawful-interception and data-provision frameworks under telecommunications and criminal-procedure law. Lawful interception is not the same thing as routine data retention, but it matters for a privacy-first email provider because it affects operational obligations.

For Privacy.Fish, Norway is still easier to defend because Norway’s mandatory retained data is explicit, narrow, and known. Austria may be interesting, but the exact current provider-specific email obligations require deeper local-law review.

Why Norway is better:
- Austria’s old blanket retention law was struck down, but that does not automatically mean “no relevant obligations.”
- Austria still has lawful-interception and data-provision frameworks.
- Norway gives a clear public answer: account, source IP:port, timestamp for 12 months.
- Norway is easier to document precisely.

Sources:
- [Austrian Constitutional Court press release on data retention](https://www.vfgh.gv.at/downloads/press_releasedataretention.pdf)
- [Austrian Constitutional Court PDF bulletin](https://www.vfgh.gv.at/downloads/Bulletin_2014-2_G_47-2012__G_59-2012__G_62_70_71-2012_27.06..pdf)
- [RTR English translation of Austrian Telecommunications Act 2021](https://www.rtr.at/rtr/service/rechtsvorschriften/gesetze/TKG_2021_en-gb.pdf)
- [Austrian TKG 2021 PDF via RIS](https://www.ris.bka.gv.at/Dokumente/Erv/ERV_2021_1_190/ERV_2021_1_190.pdf)
- [FRA Austria surveillance study mentioning annulled data retention](https://fra.europa.eu/sites/default/files/fra_uploads/austria-study-data-surveillance-at.pdf)
- [Deutsche Telekom Austria transparency report legal basis](https://www.telekom.com/en/company/data-privacy-and-security/news/austria-363540)
- [EFF-style data-retention cases briefing mentioning Austria](https://www.patrick-breyer.de/wp-content/uploads/2021/09/Briefing-Data-Retention-Cases.pdf)

## Comparison Table

| Jurisdiction | Required logging / retention fit | Main issue compared with Norway |
|---|---|---|
| Norway | Account, source IP:port, timestamp for 12 months | Best fit |
| EU | No single EU-wide rule after the Data Retention Directive was invalidated | Every country must be checked individually |
| Switzerland | Broader email metadata can be in scope | Sender, recipient, protocol, mailbox-event, and server metadata |
| Germany | Interception readiness for large email systems | TKÜV duties above 100,000 users |
| Netherlands | No old blanket-retention law, but interception and disclosure duties remain | Less clear than Norway |
| Sweden | Broader categories and variable periods | Source, destination, equipment, timing, and category-dependent retention |
| Denmark | Targeted, regional, and national-security retention model | More complex than Norway |
| Finland | Selected providers; broader fields if covered | Less predictable and broader when applicable |
| Iceland | 6 months, but broader connection records | Can include connection counterparties and transferred data |
| United Kingdom | Retention notices can require broad communications data | Secret notices and broader categories |
| United States | No simple blanket rule, but broad compelled disclosure and preservation | Less predictable legal exposure |
| Canada | No simple current Norway-style rule, but lawful-access expansion proposed | Metadata retention up to one year proposed |
| Australia | Broad prescribed metadata dataset for at least 2 years | Much broader and longer than Norway |
| New Zealand | No simple blanket rule found, but interception-capability duties | Interception design obligations |
| France | Broader national-security traffic/location retention | Broader and less suitable |
| Belgium | Targeted/differentiated metadata retention | Complex and repeatedly litigated |
| Ireland | Historically broader internet/email retention framework | More complex and legally contested |
| Luxembourg | Draft/unclear traffic-location retention framework | Less clear than Norway |
| Austria | Old blanket retention struck down, but interception frameworks remain | Interesting, but less clear than Norway |
