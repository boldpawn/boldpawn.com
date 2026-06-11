---
title: "IE3N99"
source_title: "IE3N99"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N99.md"
section: "IE3N99"
chunk: 1
chunk_count: 1
generated: true
---

# IE3N99

Source document: IE3N99
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N99.md`
Chunk: 1 of 1

# IE3N99

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3N99 | 0 | Class |  | IE3N99 |  |  |  |  | CIS | N.A. | CIS | IE3N99 |  |  |
| IE3N99 | 1 | Attribute | 0..1 |     Functional reference | an..35 | O |  |  | Functional reference number | D026 | CIS/functionalReference | functionalReference |  |  |
| IE3N99 | 1 | Attribute | 1..1 |     Notification date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  |  |  | CIS/notificationDate | notificationDate |  |  |
| IE3N99 | 1 | Composition | 0..1 |     Addressed Member State |  | O |  |  | AddressedMemberState | SC2 | CIS/AddressedMemberState | addressedMemberState |  |  |
| IE3N99 | 2 | Attribute | 1..1 |         Country | a2 | M |  |  |  |  | CIS/AddressedMemberState/country | country | CL717 | CL717 |
| IE3N99 | 1 | Composition | 0..1 |     Representative |  | O |  |  | Agent | 05A | CIS/Agent | representative |  |  |
| IE3N99 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Agent, coded | R004 | CIS/Agent/identification | identificationNumber |  |  |
| IE3N99 | 1 | Composition | 0..1 |     Transport document (Master level) |  | O |  |  | TransportContractDocument | 30B | CIS/TransportContractDocument | transportDocument |  |  |
| IE3N99 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocument/identification | documentNumber |  |  |
| IE3N99 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocument/type | type | CL754 | CL754 |
| IE3N99 | 1 | Composition | 0..1 |     Declarant |  | O |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3N99 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3N99 | 1 | Composition | 0..1 |     Customs office of first entry |  | O |  |  | EntryOffice | 51A | CIS/EntryOffice | customsOfficeOfFirstEntry |  |  |
| IE3N99 | 2 | Attribute | 1..1 |         Reference number | an8 | M |  |  | Office of entry, coded | G004 | CIS/EntryOffice/identification | referenceNumber | CL141 | CL141 |
| IE3N99 | 1 | Composition | 1..99999 |     Error |  | M |  |  | Error | 53A | CIS/Error | error |  |  |
| IE3N99 | 2 | Attribute | 0..1 |         Technical error message | an..1024 | O |  |  |  |  | CIS/Error/technicalErrorMessage | technicalErrorMessage |  |  |
| IE3N99 | 2 | Attribute | 0..1 |         Description | an..512 | O |  |  | Error description | 488 | CIS/Error/description | description |  |  |
| IE3N99 | 2 | Attribute | 1..1 |         Validation code | an..8 | M |  |  | Error, coded | 377 | CIS/Error/validation | validationCode | CL723 | CL723 |
| IE3N99 | 2 | Composition | 0..1 |         Pointer |  | O |  |  | Pointer | 97A | CIS/Error/Pointer | pointer |  |  |
| IE3N99 | 3 | Attribute | 1..1 |             Message element path | an..512 | M |  |  |  |  | CIS/Error/Pointer/messageElementPath | messageElementPath |  |  |
