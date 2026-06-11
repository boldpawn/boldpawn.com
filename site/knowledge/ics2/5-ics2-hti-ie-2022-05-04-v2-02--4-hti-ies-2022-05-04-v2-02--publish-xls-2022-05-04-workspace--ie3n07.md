---
title: "IE3N07"
source_title: "IE3N07"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N07.md"
section: "IE3N07"
chunk: 1
chunk_count: 1
generated: true
---

# IE3N07

Source document: IE3N07
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N07.md`
Chunk: 1 of 1

# IE3N07

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3N07 | 0 | Class |  | IE3N07 |  |  |  |  | CIS | N.A. | CIS | IE3N07 |  |  |
| IE3N07 | 1 | Attribute | 1..1 |     MRN | an18 | M |  |  |  |  | CIS/masterReferenceNumber | MRN |  |  |
| IE3N07 | 1 | Attribute | 1..1 |     Notification date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  |  |  | CIS/notificationDate | notificationDate |  |  |
| IE3N07 | 1 | Composition | 0..1 |     Representative |  | O |  |  | Agent | 05A | CIS/Agent | representative |  |  |
| IE3N07 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Agent, coded | R004 | CIS/Agent/identification | identificationNumber |  |  |
| IE3N07 | 1 | Composition | 1..1 |     Transport document (Master level) |  | M |  |  | TransportContractDocument | 30B | CIS/TransportContractDocument | transportDocument |  |  |
| IE3N07 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocument/identification | documentNumber |  |  |
| IE3N07 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocument/type | type | CL754 | CL754 |
| IE3N07 | 1 | Composition | 0..99999 |     Transport document (House level) |  | O |  |  | TransportContractDocument | 30B | CIS/TransportContractDocumentHouse | transportDocumentHouse |  |  |
| IE3N07 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocumentHouse/identification | documentNumber |  |  |
| IE3N07 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocumentHouse/type | type | CL754 | CL754 |
| IE3N07 | 1 | Composition | 0..1 |     Declarant |  | O |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3N07 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3N07 | 1 | Composition | 0..1 |     Person notifying the arrival |  | O |  |  | Submitter | 17B | CIS/PersonNotifyingArrival | personNotifyingTheArrival |  |  |
| IE3N07 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Submitter, coded | R059 | CIS/PersonNotifyingArrival/identification | identificationNumber |  |  |
| IE3N07 | 1 | Composition | 1..1 |     Customs office of first entry |  | M |  |  | EntryOffice | 51A | CIS/EntryOffice | customsOfficeOfFirstEntry |  |  |
| IE3N07 | 2 | Attribute | 1..1 |         Reference number | an8 | M |  |  | Office of entry, coded | G004 | CIS/EntryOffice/identification | referenceNumber | CL141 | CL141 |
| IE3N07 | 1 | Composition | 1..99999 |     Error |  | M |  |  | Error | 53A | CIS/Error | error |  |  |
| IE3N07 | 2 | Attribute | 1..1 |         Description | an..512 | M |  |  | Error description | 488 | CIS/Error/description | description |  |  |
| IE3N07 | 2 | Attribute | 1..1 |         Validation code | an..8 | M |  |  | Error, coded | 377 | CIS/Error/validation | validationCode | CL723 | CL723 |
| IE3N07 | 2 | Composition | 1..1 |         Pointer |  | M |  |  | Pointer | 97A | CIS/Error/Pointer | pointer |  |  |
| IE3N07 | 3 | Attribute | 1..1 |             Message element path | an..512 | M |  |  |  |  | CIS/Error/Pointer/messageElementPath | messageElementPath |  |  |
| IE3N07 | 1 | Composition | 0..1 |     Notify party |  | O |  |  | RepresentativePerson | 06B | CIS/RepresentativePerson | notifyParty |  |  |
| IE3N07 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  |  |  | CIS/RepresentativePerson/identification | identificationNumber |  |  |
