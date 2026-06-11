---
title: "IE3R01"
source_title: "IE3R01"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3R01.md"
section: "IE3R01"
chunk: 1
chunk_count: 1
generated: true
---

# IE3R01

Source document: IE3R01
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3R01.md`
Chunk: 1 of 1

# IE3R01

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3R01 | 0 | Class |  | IE3R01 |  |  |  |  | CIS | N.A. | CIS | IE3R01 |  |  |
| IE3R01 | 1 | Attribute | 1..1 |     LRN | an..22 | M |  |  | Functional reference number | D026 | CIS/functionalReference | LRN |  |  |
| IE3R01 | 1 | Attribute | 1..1 |     MRN | an18 | M |  |  |  |  | CIS/masterReferenceNumber | MRN |  |  |
| IE3R01 | 1 | Attribute | 1..1 |     Registration date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  |  |  | CIS/registrationDate | registrationDate |  |  |
| IE3R01 | 1 | Attribute | 1..1 |     Specific circumstance indicator | an3 | M |  |  | Specific circumstances, coded | 504 | CIS/specificCircumstances | SpecificCircumstanceIndicator | CL742 | CL742 |
| IE3R01 | 1 | Composition | 0..1 |     Addressed Member State |  | O |  |  | AddressedMemberState | SC2 | CIS/AddressedMemberState | addressedMemberState |  |  |
| IE3R01 | 2 | Attribute | 1..1 |         Country | a2 | M |  |  |  |  | CIS/AddressedMemberState/country | country | CL717 | CL717 |
| IE3R01 | 1 | Composition | 0..1 |     Representative |  | O |  |  | Agent | 05A | CIS/Agent | representative |  |  |
| IE3R01 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Agent, coded | R004 | CIS/Agent/identification | identificationNumber |  |  |
| IE3R01 | 1 | Composition | 0..1 |     Transport document (Master level) |  | O |  |  | TransportContractDocument | 30B | CIS/TransportContractDocument | transportDocument |  |  |
| IE3R01 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocument/identification | documentNumber |  |  |
| IE3R01 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocument/type | type | CL754 | CL754 |
| IE3R01 | 1 | Composition | 0..1 |     Carrier |  | O |  |  | Carrier | 18A | CIS/Carrier | carrier |  |  |
| IE3R01 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Carrier identification | R012 | CIS/Carrier/identification | identificationNumber |  |  |
| IE3R01 | 1 | Composition | 0..1 |     Declarant |  | Ο |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3R01 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3R01 | 1 | Composition | 0..1 |     Customs office of first entry |  | O |  |  | EntryOffice | 51A | CIS/EntryOffice | customsOfficeOfFirstEntry |  |  |
| IE3R01 | 2 | Attribute | 1..1 |         Reference number | an8 | M |  |  | Office of entry, coded | G004 | CIS/EntryOffice/identification | referenceNumber | CL141 | CL141 |
