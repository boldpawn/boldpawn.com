---
title: "IE3N05"
source_title: "IE3N05"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N05.md"
section: "IE3N05"
chunk: 1
chunk_count: 1
generated: true
---

# IE3N05

Source document: IE3N05
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N05.md`
Chunk: 1 of 1

# IE3N05

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3N05 | 0 | Class |  | IE3N05 |  |  |  |  | CIS | N.A. | CIS | IE3N05 |  |  |
| IE3N05 | 1 | Attribute | 1..1 |     Document issue date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  | Document issuing date<br>Response issuing date | D011<br>D029 | CIS/issue | documentIssueDate |  |  |
| IE3N05 | 1 | Attribute | 1..1 |     MRN | an18 | M |  |  |  |  | CIS/masterReferenceNumber | MRN |  |  |
| IE3N05 | 1 | Composition | 1..1 |     Responsible Member State |  | M |  |  | ResponsibleMemberState | SC2 | CIS/ResponsibleMemberState | responsibleMemberState |  |  |
| IE3N05 | 2 | Attribute | 1..1 |         Country | a2 | M |  |  |  |  | CIS/ResponsibleMemberState/country | country | CL717 | CL717 |
| IE3N05 | 1 | Composition | 0..1 |     Transport document (Master level) |  | O |  |  | TransportContractDocument | 30B | CIS/TransportContractDocument | transportDocument |  |  |
| IE3N05 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocument/identification | documentNumber |  |  |
| IE3N05 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocument/type | type | CL754 | CL754 |
| IE3N05 | 1 | Composition | 1..1 |     Carrier |  | M |  |  | Carrier | 18A | CIS/Carrier | carrier |  |  |
| IE3N05 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Carrier identification | R012 | CIS/Carrier/identification | identificationNumber |  |  |
| IE3N05 | 1 | Composition | 1..999 |     Referral request details |  | M |  |  |  |  | CIS/ReferralRequest | referralRequestDetails |  |  |
| IE3N05 | 2 | Attribute | 1..1 |         Referral request reference | an..17 | M |  |  |  |  | CIS/ReferralRequest/Identification | referralRequestReference |  |  |
| IE3N05 | 2 | Attribute | 1..1 |         Request type | an..3 | M |  |  |  |  | CIS/ReferralRequest/requestType | requestType | CL735 | CL735 |
| IE3N05 | 2 | Composition | 0..9 |         Recommended HRCM Screening Method |  | O |  |  |  |  | CIS/ReferralRequest/recommendedHRCMScreeningMethod | recommendedHRCMScreeningMethod |  |  |
| IE3N05 | 3 | Attribute | 1..1 |             Method | an..3 | M |  |  |  |  | CIS/ReferralRequest/recommendedHRCMScreeningMethod/method | method | CL724 | CL724 |
| IE3N05 | 2 | Composition | 1..1 |         Transport document (House level) |  | M |  |  | TransportContractDocument | 30B | CIS/ReferralRequest/TransportContractDocument | transportDocumentHouse |  |  |
| IE3N05 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/ReferralRequest/TransportContractDocument/identification | documentNumber |  |  |
| IE3N05 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/ReferralRequest/TransportContractDocument/type | type | CL754 | CL754 |
