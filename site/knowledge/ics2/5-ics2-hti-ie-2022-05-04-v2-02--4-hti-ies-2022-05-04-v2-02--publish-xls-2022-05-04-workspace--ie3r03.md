---
title: "IE3R03"
source_title: "IE3R03"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3R03.md"
section: "IE3R03"
chunk: 1
chunk_count: 1
generated: true
---

# IE3R03

Source document: IE3R03
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3R03.md`
Chunk: 1 of 1

# IE3R03

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3R03 | 0 | Class |  | IE3R03 |  |  |  |  | CIS | N.A. | CIS | IE3R03 |  |  |
| IE3R03 | 1 | Attribute | 1..1 |     Document issue date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  | Document issuing date<br>Response issuing date | D011<br>D029 | CIS/issue | documentIssueDate |  |  |
| IE3R03 | 1 | Attribute | 1..1 |     MRN | an18 | M |  |  |  |  | CIS/masterReferenceNumber | MRN |  |  |
| IE3R03 | 1 | Composition | 1..1 |     Responsible Member State |  | M |  |  | ResponsibleMemberState | SC2 | CIS/ResponsibleMemberState | responsibleMemberState |  |  |
| IE3R03 | 2 | Attribute | 1..1 |         Country | a2 | M |  |  |  |  | CIS/ResponsibleMemberState/country | country | CL717 | CL717 |
| IE3R03 | 1 | Composition | 0..1 |     Representative |  | O |  |  | Agent | 05A | CIS/Agent | representative |  |  |
| IE3R03 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005 |  | Agent, coded | R004 | CIS/Agent/identification | identificationNumber |  |  |
| IE3R03 | 1 | Composition | 0..1 |     Transport document (Master level) |  | O |  |  | TransportContractDocument | 30B | CIS/TransportContractDocument | transportDocument |  |  |
| IE3R03 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocument/identification | documentNumber |  |  |
| IE3R03 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocument/type | type | CL754 | CL754 |
| IE3R03 | 1 | Composition | 1..1 |     Declarant |  | M |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3R03 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005 |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3R03 | 1 | Composition | 1..999 |     HRCM screening results |  | M |  |  |  |  | CIS/HRCMScreeningResponse | hrcmScreeningResults |  |  |
| IE3R03 | 2 | Attribute | 1..1 |         Referral request reference | an..17 | M |  |  |  |  | CIS/HRCMScreeningResponse/identifier | referralRequestReference |  |  |
| IE3R03 | 2 | Attribute | 1..1 |         Result | an..3 | M |  |  |  |  | CIS/HRCMScreeningResponse/result | result | CL726 | CL726 |
| IE3R03 | 2 | Composition | 2..9 |         Screening method |  | M |  |  |  |  | CIS/HRCMScreeningResponse/ScreeningMethod | screeningMethod |  |  |
| IE3R03 | 3 | Attribute | 1..1 |             Method | an..3 | M |  |  |  |  | CIS/HRCMScreeningResponse/ScreeningMethod/method | method | CL724 | CL724 |
| IE3R03 | 2 | Composition | 0..99 |         Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/HRCMScreeningResponse/AdditionalInformation | additionalInformation |  |  |
| IE3R03 | 3 | Attribute | 0..1 |             Code | an..17 | O |  |  | Additional statement code | 226 | CIS/HRCMScreeningResponse/AdditionalInformation/statement | code | CL752 | CL752 |
| IE3R03 | 3 | Attribute | 0..1 |             Text | an..512 | O |  |  | Additional statement text | 225 | CIS/HRCMScreeningResponse/AdditionalInformation/statementDescription | text |  |  |
| IE3R03 | 3 | Attribute | 0..1 |             Information type | an..3 | O |  |  | Additional statement type | 369 | CIS/HRCMScreeningResponse/AdditionalInformation/statementType | type | CL703 | CL703 |
| IE3R03 | 2 | Composition | 0..1 |         Authorised person |  | O |  |  | AuthorisedPerson | 13A | CIS/HRCMScreeningResponse/AuthorisedPerson | authorizedPerson |  |  |
| IE3R03 | 3 | Attribute | 1..1 |             Name | an..70 | M |  |  | Authorised person | R067 | CIS/HRCMScreeningResponse/AuthorisedPerson/name | name |  |  |
| IE3R03 | 3 | Attribute | 1..1 |             Identification number | an..17 | M |  |  | Authorised person, coded | R068 | CIS/HRCMScreeningResponse/AuthorisedPerson/identification | identificationNumber |  |  |
| IE3R03 | 3 | Attribute | 1..1 |             Type of person | n1 | M |  |  |  |  | CIS/HRCMScreeningResponse/AuthorisedPerson/identificationType | typeOfPerson | CL729 | CL729 |
| IE3R03 | 2 | Composition | 0..9 |         Binary attachment |  | O |  |  | BinaryFile | 72B | CIS/HRCMScreeningResponse/BinaryFile | binaryFile |  |  |
| IE3R03 | 3 | Attribute | 1..1 |             Identification | an..256 | M |  |  | Binary File Identifier | M001 | CIS/HRCMScreeningResponse/BinaryFile/identification |  |  |  |
| IE3R03 | 3 | Attribute | 1..1 |             Filename | an..256 | M |  |  | Binary File Name | M005 | CIS/HRCMScreeningResponse/BinaryFile/fileName | filename |  |  |
| IE3R03 | 3 | Attribute | 1..1 |             MIME | an..70 | M |  |  | Binary File MIME, coded | M007 | CIS/HRCMScreeningResponse/BinaryFile/MIME | MIME |  |  |
| IE3R03 | 3 | Attribute | 0..1 |             Description | an..256 | O |  |  | Binary File Description | M012 | CIS/HRCMScreeningResponse/BinaryFile/description | description |  |  |
| IE3R03 | 2 | Composition | 0..1 |         Screening facility |  | Ο |  |  | FacilityPlace | 58A | CIS/HRCMScreeningResponse/FacilityPlace | facilityPlace |  |  |
| IE3R03 | 3 | Composition | 1..1 |             Address |  | M | R3002 |  | Address | 04A | CIS/HRCMScreeningResponse/FacilityPlace/Address | address |  |  |
| IE3R03 | 4 | Attribute | 1..1 |                 City | an..35 | M |  |  | City name | 241 | CIS/HRCMScreeningResponse/FacilityPlace/Address/cityName | city |  |  |
| IE3R03 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/HRCMScreeningResponse/FacilityPlace/Address/country | country | CL718 | CL718 |
| IE3R03 | 4 | Attribute | 0..1 |                 Sub-division | an..35 | Ο |  |  | Country sub-entity name | 243 | CIS/HRCMScreeningResponse/FacilityPlace/Address/countrySubDivisionName | subDivision |  |  |
| IE3R03 | 4 | Attribute | 0..1 |                 Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/HRCMScreeningResponse/FacilityPlace/Address/line | street |  |  |
| IE3R03 | 4 | Attribute | 0..1 |                 Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/HRCMScreeningResponse/FacilityPlace/Address/postcode | postCode |  |  |
| IE3R03 | 4 | Attribute | 0..1 |                 Street additional line | an..70 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/additionalLine | streetAdditionalLine |  |  |
| IE3R03 | 4 | Attribute | 0..1 |                 Number | an..35 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/streetNumber | number |  |  |
| IE3R03 | 4 | Attribute | 0..1 |                 P.O. Box | an..70 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/poBox | poBox |  |  |
| IE3R03 | 2 | Composition | 1..1 |         Transport document (House level) |  | M |  |  | TransportContractDocument | 30B | CIS/HRCMScreeningResponse/TransportContractDocument | transportDocumentHouseLevel |  |  |
| IE3R03 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/HRCMScreeningResponse/TransportContractDocument/identification | documentNumber |  |  |
| IE3R03 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/HRCMScreeningResponse/TransportContractDocument/type | type | CL754 | CL754 |
