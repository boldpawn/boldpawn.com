---
title: "IE3Q01"
source_title: "IE3Q01"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3Q01.md"
section: "IE3Q01"
chunk: 1
chunk_count: 1
generated: true
---

# IE3Q01

Source document: IE3Q01
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3Q01.md`
Chunk: 1 of 1

# IE3Q01

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3Q01 | 0 | Class |  | IE3Q01 |  |  |  |  | CIS | N.A. | CIS | IE3Q01 |  |  |
| IE3Q01 | 1 | Attribute | 1..1 |     Document issue date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  | Document issuing date<br>Response issuing date | D011<br>D029 | CIS/issue | documentIssueDate |  |  |
| IE3Q01 | 1 | Attribute | 1..1 |     MRN | an18 | M |  |  |  |  | CIS/masterReferenceNumber | MRN |  |  |
| IE3Q01 | 1 | Composition | 1..1 |     Responsible Member State |  | M |  |  | ResponsibleMemberState | SC2 | CIS/ResponsibleMemberState | responsibleMemberState |  |  |
| IE3Q01 | 2 | Attribute | 1..1 |         Country | a2 | M |  |  |  |  | CIS/ResponsibleMemberState/country | country | CL717 | CL717 |
| IE3Q01 | 1 | Composition | 0..1 |     Representative |  | O |  |  | Agent | 05A | CIS/Agent | representative |  |  |
| IE3Q01 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Agent, coded | R004 | CIS/Agent/identification | identificationNumber |  |  |
| IE3Q01 | 1 | Composition | 0..1 |     Transport document (Master level) |  | O |  |  | TransportContractDocument | 30B | CIS/TransportContractDocument | transportDocument |  |  |
| IE3Q01 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocument/identification | documentNumber |  |  |
| IE3Q01 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocument/type | type | CL754 | CL754 |
| IE3Q01 | 1 | Composition | 0..1 |     Transport document (House level) |  | O |  |  | TransportContractDocument | 30B | CIS/TransportContractDocumentHouse | transportDocumentHouse |  |  |
| IE3Q01 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocumentHouse/identification | documentNumber |  |  |
| IE3Q01 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocumentHouse/type | type | CL754 | CL754 |
| IE3Q01 | 1 | Composition | 0..1 |     Carrier |  | O |  |  | Carrier | 18A | CIS/Carrier | carrier |  |  |
| IE3Q01 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Carrier identification | R012 | CIS/Carrier/identification | identificationNumber |  |  |
| IE3Q01 | 1 | Composition | 0..1 |     Declarant |  | O |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3Q01 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3Q01 | 1 | Composition | 1..1 |     Do not load details |  | M | R3025 |  |  |  | CIS/DoNotLoad | doNotLoadDetails |  |  |
| IE3Q01 | 2 | Composition | 0..999 |         Receptacle |  | O |  |  |  |  | CIS/DoNotLoad/PostalReceptacle | receptacle |  |  |
| IE3Q01 | 3 | Attribute | 1..1 |             Receptacle identifier | an..35 | M |  |  |  |  | CIS/DoNotLoad/PostalReceptacle/receptacleIdentificationNumber | receptacleIdentificationNumber |  |  |
| IE3Q01 | 2 | Composition | 0..1 |         Transport document (House level) |  | O |  |  | TransportContractDocument | 30B | CIS/DoNotLoad/TransportContractDocument | transportDocument |  |  |
| IE3Q01 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/DoNotLoad/TransportContractDocument/identification | documentNumber |  |  |
| IE3Q01 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/DoNotLoad/TransportContractDocument/type | type | CL754 | CL754 |
| IE3Q01 | 2 | Composition | 0..9999 |         Transport equipment |  | O |  |  | TransportEquipment | 31B | CIS/DoNotLoad/TransportEquipment | transportEquipment |  |  |
| IE3Q01 | 3 | Attribute | 1..1 |             Container identification number | an..17 | M |  |  | Equipment identification number | 159 | CIS/DoNotLoad/TransportEquipment/identification | containerIdentificationNumber |  |  |
