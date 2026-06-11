---
title: "IE3Q05"
source_title: "IE3Q05"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3Q05.md"
section: "IE3Q05"
chunk: 1
chunk_count: 1
generated: true
---

# IE3Q05

Source document: IE3Q05
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3Q05.md`
Chunk: 1 of 1

# IE3Q05

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3Q05 | 0 | Class |  | IE3Q05 |  |  |  |  | CIS | N.A. | CIS | IE3Q05 |  |  |
| IE3Q05 | 1 | Attribute | 1..1 |     Request notifications | n1 | M |  |  |  |  | CIS/dataRequest | dataRequest | CL027 | CL027 |
| IE3Q05 | 1 | Attribute | 1..1 |     Functional reference | an..35 | M | R3003 |  | Functional reference number | D026 | CIS/functionalReference | functionalReference |  |  |
| IE3Q05 | 1 | Attribute | 0..1 |     MRN | an18 | O | R3028 |  |  |  | CIS/masterReferenceNumber | MRN |  |  |
| IE3Q05 | 1 | Composition | 1..1 |     Requester |  | M |  |  | Submitter | 17B | CIS/Requester | requester |  |  |
| IE3Q05 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005 |  | Submitter, coded | R059 | CIS/Requester/identification | identificationNumber |  |  |
| IE3Q05 | 1 | Composition | 0..1 |     Transport document (Master level) |  | O | R3028 |  | TransportContractDocument | 30B | CIS/TransportContractDocument | transportDocument |  |  |
| IE3Q05 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocument/identification | documentNumber |  |  |
| IE3Q05 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocument/type | type | CL754 | CL754 |
| IE3Q05 | 1 | Composition | 0..1 |     Transport document (House level) |  | O | R3028 |  | TransportContractDocument | 30B | CIS/TransportContractDocumentHouse | transportDocumentHouse |  |  |
| IE3Q05 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocumentHouse/identification | documentNumber |  |  |
| IE3Q05 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocumentHouse/type | type | CL754 | CL754 |
