---
title: "IE3A26 - part 3"
source_title: "IE3A26"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A26.md"
section: "IE3A26"
chunk: 3
chunk_count: 3
generated: true
---

# IE3A26 - part 3

Source document: IE3A26
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A26.md`
Chunk: 3 of 3

| IE3A26 | 2 | Attribute | 1..1 |         Name | an..70 | M |  |  | Declarant name | R124 | CIS/Declarant/name | name |  |  |
| IE3A26 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005<br>R3024 |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3A26 | 2 | Composition | 1..1 |         Address |  | M | R3002 |  | Address | 04A | CIS/Declarant/Address | address |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             City | an..35 | M |  |  | City name | 241 | CIS/Declarant/Address/cityName | city |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             Country | a2 | M |  |  | Country, coded | 242 | CIS/Declarant/Address/country | country | CL718 | CL718 |
| IE3A26 | 3 | Attribute | 0..1 |             Sub-division | an..35 | Ο |  |  | Country sub-entity name | 243 | CIS/Declarant/Address/countrySubDivisionName | subDivision |  |  |
| IE3A26 | 3 | Attribute | 0..1 |             Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Declarant/Address/line | street |  |  |
| IE3A26 | 3 | Attribute | 0..1 |             Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Declarant/Address/postcode | postCode |  |  |
| IE3A26 | 3 | Attribute | 0..1 |             Street additional line | an..70 | O |  |  |  |  | CIS/Declarant/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A26 | 3 | Attribute | 0..1 |             Number | an..35 | O |  |  |  |  | CIS/Declarant/Address/streetNumber | number |  |  |
| IE3A26 | 3 | Attribute | 0..1 |             P.O. Box | an..70 | O |  |  |  |  | CIS/Declarant/Address/poBox | poBox |  |  |
| IE3A26 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/Declarant/Communication | communication |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Declarant/Communication/identification | identifier |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/Declarant/Communication/type | type | CL707 | CL707 |
| IE3A26 | 1 | Composition | 0..99 |     HRCM screening results |  | O |  |  |  |  | CIS/HRCMScreeningResponse | hrcmScreeningResults |  |  |
| IE3A26 | 2 | Attribute | 1..1 |         Result | an..3 | M |  |  |  |  | CIS/HRCMScreeningResponse/result | result | CL726 | CL726 |
| IE3A26 | 2 | Composition | 2..9 |         Screening method |  | M |  |  |  |  | CIS/HRCMScreeningResponse/ScreeningMethod | screeningMethod |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             Method | an..3 | M |  |  |  |  | CIS/HRCMScreeningResponse/ScreeningMethod/method | method | CL724 | CL724 |
| IE3A26 | 2 | Composition | 0..99 |         Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/HRCMScreeningResponse/AdditionalInformation | additionalInformation |  |  |
| IE3A26 | 3 | Attribute | 0..1 |             Code | an..17 | O |  |  | Additional statement code | 226 | CIS/HRCMScreeningResponse/AdditionalInformation/statement | code | CL752 | CL752 |
| IE3A26 | 3 | Attribute | 0..1 |             Text | an..512 | O |  |  | Additional statement text | 225 | CIS/HRCMScreeningResponse/AdditionalInformation/statementDescription | text |  |  |
| IE3A26 | 3 | Attribute | 0..1 |             Type | an..3 | O |  |  | Additional statement type | 369 | CIS/HRCMScreeningResponse/AdditionalInformation/statementType | type | CL703 | CL703 |
| IE3A26 | 2 | Composition | 0..1 |         Authorized person |  | O |  |  | AuthorisedPerson | 13A | CIS/HRCMScreeningResponse/AuthorisedPerson | authorizedPerson |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             Name | an..70 | M |  |  | Authorised person | R067 | CIS/HRCMScreeningResponse/AuthorisedPerson/name | name |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             Identification number | an..17 | M |  |  | Authorised person, coded | R068 | CIS/HRCMScreeningResponse/AuthorisedPerson/identification | identificationNumber |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             Type of person | n1 | M |  |  |  |  | CIS/HRCMScreeningResponse/AuthorisedPerson/identificationType | typeOfPerson | CL729 | CL729 |
| IE3A26 | 2 | Composition | 0..9 |         Binary file |  | O |  |  | BinaryFile | 72B | CIS/HRCMScreeningResponse/BinaryFile | binaryFile |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             Identification | an..256 | M |  |  | Binary File Identifier | M001 | CIS/HRCMScreeningResponse/BinaryFile/identification |  |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             Filename | an..256 | M |  |  | Binary File Name | M005 | CIS/HRCMScreeningResponse/BinaryFile/fileName | filename |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             MIME | an..70 | M |  |  | Binary File MIME, coded | M007 | CIS/HRCMScreeningResponse/BinaryFile/MIME | MIME |  |  |
| IE3A26 | 3 | Attribute | 0..1 |             Description | an..256 | O |  |  | Binary File Description | M012 | CIS/HRCMScreeningResponse/BinaryFile/description | description |  |  |
| IE3A26 | 2 | Composition | 0..1 |         Facility place |  | O |  |  | FacilityPlace | 58A | CIS/HRCMScreeningResponse/FacilityPlace | facilityPlace |  |  |
| IE3A26 | 3 | Composition | 1..1 |             Address |  | M |  |  | Address | 04A | CIS/HRCMScreeningResponse/FacilityPlace/Address | address |  |  |
| IE3A26 | 4 | Attribute | 1..1 |                 City | an..35 | M |  |  | City name | 241 | CIS/HRCMScreeningResponse/FacilityPlace/Address/cityName | city |  |  |
| IE3A26 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/HRCMScreeningResponse/FacilityPlace/Address/country | country | CL718 | CL718 |
| IE3A26 | 4 | Attribute | 0..1 |                 Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/HRCMScreeningResponse/FacilityPlace/Address/countrySubDivisionName | subDivision |  |  |
| IE3A26 | 4 | Attribute | 0..1 |                 Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/HRCMScreeningResponse/FacilityPlace/Address/line | street |  |  |
| IE3A26 | 4 | Attribute | 0..1 |                 Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/HRCMScreeningResponse/FacilityPlace/Address/postcode | postCode |  |  |
| IE3A26 | 4 | Attribute | 0..1 |                 Street additional line | an..70 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A26 | 4 | Attribute | 0..1 |                 Number | an..35 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/streetNumber | number |  |  |
| IE3A26 | 4 | Attribute | 0..1 |                 P.O. Box | an..70 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/poBox | poBox |  |  |
| IE3A26 | 2 | Composition | 0..1 |         Transport document (House level) |  | O |  |  | TransportContractDocument | 30B | CIS/HRCMScreeningResponse/TransportContractDocument | transportDocumentHouseLevel |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/HRCMScreeningResponse/TransportContractDocument/identification | documentNumber |  |  |
| IE3A26 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/HRCMScreeningResponse/TransportContractDocument/type | type | CL754 | CL754 |
