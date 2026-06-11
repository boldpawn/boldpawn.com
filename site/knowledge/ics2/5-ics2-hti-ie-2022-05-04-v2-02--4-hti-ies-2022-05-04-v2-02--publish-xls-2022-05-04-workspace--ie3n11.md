---
title: "IE3N11"
source_title: "IE3N11"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N11.md"
section: "IE3N11"
chunk: 1
chunk_count: 1
generated: true
---

# IE3N11

Source document: IE3N11
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N11.md`
Chunk: 1 of 1

# IE3N11

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3N11 | 0 | Class |  | IE3N11 |  |  |  |  | CIS | N.A. | CIS | IE3N11 |  |  |
| IE3N11 | 1 | Attribute | 1..1 |     Notification date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  |  |  | CIS/notificationDate | notificationDate |  |  |
| IE3N11 | 1 | Composition | 1..1 |     Transport document (Master level) |  | M |  |  | TransportContractDocument | 30B | CIS/TransportContractDocument | transportDocument |  |  |
| IE3N11 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocument/identification | documentNumber |  |  |
| IE3N11 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocument/type | type | CL754 | CL754 |
| IE3N11 | 1 | Composition | 1..1 |     Customs office of first entry |  | M |  |  | EntryOffice | 51A | CIS/EntryOffice | customsOfficeOfFirstEntry |  |  |
| IE3N11 | 2 | Attribute | 1..1 |         Reference number | an8 | M |  |  | Office of entry, coded | G004 | CIS/EntryOffice/identification | referenceNumber | CL141 | CL141 |
| IE3N11 | 1 | Composition | 1..1 |     Supplementary declarant |  | M |  |  | SupplementaryDeclarant | 18C | CIS/SupplementaryDeclarant | supplementaryDeclarant |  |  |
| IE3N11 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Supplementary declarant, coded | R146 | CIS/SupplementaryDeclarant/identification | identificationNumber |  |  |
