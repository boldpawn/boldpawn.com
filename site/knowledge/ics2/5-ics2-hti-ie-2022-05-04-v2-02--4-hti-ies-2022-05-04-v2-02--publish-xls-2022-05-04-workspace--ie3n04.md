---
title: "IE3N04"
source_title: "IE3N04"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N04.md"
section: "IE3N04"
chunk: 1
chunk_count: 1
generated: true
---

# IE3N04

Source document: IE3N04
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N04.md`
Chunk: 1 of 1

# IE3N04

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3N04 | 0 | Class |  | IE3N04 |  |  |  |  | CIS | N.A. | CIS | IE3N04 |  |  |
| IE3N04 | 1 | Attribute | 1..1 |     Document issue date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  | Document issuing date<br>Response issuing date | D011<br>D029 | CIS/issue | documentIssueDate |  |  |
| IE3N04 | 1 | Attribute | 1..1 |     MRN | an18 | M |  |  |  |  | CIS/masterReferenceNumber | MRN |  |  |
| IE3N04 | 1 | Composition | 1..1 |     Responsible Member State |  | M |  |  | ResponsibleMemberState | SC2 | CIS/ResponsibleMemberState | responsibleMemberState |  |  |
| IE3N04 | 2 | Attribute | 1..1 |         Country | a2 | M |  |  |  |  | CIS/ResponsibleMemberState/country | country | CL717 | CL717 |
| IE3N04 | 1 | Composition | 1..1 |     Carrier |  | M |  |  | Carrier | 18A | CIS/Carrier | carrier |  |  |
| IE3N04 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Carrier identification | R012 | CIS/Carrier/identification | identificationNumber |  |  |
| IE3N04 | 1 | Composition | 1..999 |     Referral request details |  | M |  |  |  |  | CIS/ReferralRequest | referralRequestDetails |  |  |
| IE3N04 | 2 | Attribute | 1..1 |         Referral request reference | an..17 | M |  |  |  |  | CIS/ReferralRequest/Identification | referralRequestReference |  |  |
| IE3N04 | 2 | Attribute | 1..1 |         Request type | an..3 | M |  |  |  |  | CIS/ReferralRequest/requestType | requestType | CL735 | CL735 |
| IE3N04 | 2 | Composition | 0..1 |         Transport document (House level) |  | O |  |  | TransportContractDocument | 30B | CIS/ReferralRequest/TransportContractDocument | transportDocumentHouse |  |  |
| IE3N04 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/ReferralRequest/TransportContractDocument/identification | documentNumber |  |  |
| IE3N04 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/ReferralRequest/TransportContractDocument/type | type | CL754 | CL754 |
| IE3N04 | 2 | Composition | 0..1 |         Transport document (Master level) |  | O |  |  | TransportContractDocument | 30B | CIS/ReferralRequest/TransportContractDocument | transportDocumentMaster |  |  |
| IE3N04 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/ReferralRequest/TransportContractDocument/identification | documentNumber |  |  |
| IE3N04 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/ReferralRequest/TransportContractDocument/type | type | CL754 | CL754 |
| IE3N04 | 2 | Composition | 0..99 |         Supporting documents |  | O |  |  | AdditionalDocument | 02A | CIS/ReferralRequest/AdditionalDocument | supportingDocuments |  |  |
| IE3N04 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Additional document reference number | D005 | CIS/ReferralRequest/AdditionalDocument/identification | referenceNumber |  |  |
| IE3N04 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Additional document type, coded | D006 | CIS/ReferralRequest/AdditionalDocument/type | type | CL013 | CL013 |
| IE3N04 | 2 | Composition | 0..99 |         Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/ReferralRequest/AdditionalInformation | additionalInformation |  |  |
| IE3N04 | 3 | Attribute | 0..1 |             Code | an..17 | O |  |  | Additional statement code | 226 | CIS/ReferralRequest/AdditionalInformation/statement | code | CL752 | CL752 |
| IE3N04 | 3 | Attribute | 0..1 |             Text | an..512 | O |  |  | Additional statement text | 225 | CIS/ReferralRequest/AdditionalInformation/statementDescription | text |  |  |
| IE3N04 | 3 | Attribute | 0..1 |             Information type | an..3 | O |  |  | Additional statement type | 369 | CIS/ReferralRequest/AdditionalInformation/statementType | type | CL703 | CL703 |
| IE3N04 | 2 | Composition | 0..1 |         Pointer |  | O |  |  | Pointer | 97A | CIS/ReferralRequest/Pointer | pointer |  |  |
| IE3N04 | 3 | Attribute | 1..1 |             Message element path | an..512 | M |  |  |  |  | CIS/ReferralRequest/Pointer/messageElementPath | messageElementPath |  |  |
