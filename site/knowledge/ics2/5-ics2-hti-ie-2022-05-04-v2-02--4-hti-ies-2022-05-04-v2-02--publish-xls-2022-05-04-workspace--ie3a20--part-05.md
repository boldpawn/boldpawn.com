---
title: "IE3A20 - part 5"
source_title: "IE3A20"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A20.md"
section: "IE3A20"
chunk: 5
chunk_count: 5
generated: true
---

# IE3A20 - part 5

Source document: IE3A20
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A20.md`
Chunk: 5 of 5

| IE3A20 | 3 | Attribute | 1..1 |             Identification | an..256 | M |  |  | Binary File Identifier | M001 | CIS/HRCMScreeningResponse/BinaryFile/identification |  |  |  |
| IE3A20 | 3 | Attribute | 1..1 |             Filename | an..256 | M |  |  | Binary File Name | M005 | CIS/HRCMScreeningResponse/BinaryFile/fileName | filename |  |  |
| IE3A20 | 3 | Attribute | 1..1 |             MIME | an..70 | M |  |  | Binary File MIME, coded | M007 | CIS/HRCMScreeningResponse/BinaryFile/MIME | MIME |  |  |
| IE3A20 | 3 | Attribute | 0..1 |             Description | an..256 | O |  |  | Binary File Description | M012 | CIS/HRCMScreeningResponse/BinaryFile/description | description |  |  |
| IE3A20 | 2 | Composition | 0..1 |         Facility place |  | O |  |  | FacilityPlace | 58A | CIS/HRCMScreeningResponse/FacilityPlace | facilityPlace |  |  |
| IE3A20 | 3 | Composition | 1..1 |             Address |  | M |  |  | Address | 04A | CIS/HRCMScreeningResponse/FacilityPlace/Address | address |  |  |
| IE3A20 | 4 | Attribute | 1..1 |                 City | an..35 | M |  |  | City name | 241 | CIS/HRCMScreeningResponse/FacilityPlace/Address/cityName | city |  |  |
| IE3A20 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/HRCMScreeningResponse/FacilityPlace/Address/country | country | CL718 | CL718 |
| IE3A20 | 4 | Attribute | 0..1 |                 Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/HRCMScreeningResponse/FacilityPlace/Address/countrySubDivisionName | subDivision |  |  |
| IE3A20 | 4 | Attribute | 0..1 |                 Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/HRCMScreeningResponse/FacilityPlace/Address/line | street |  |  |
| IE3A20 | 4 | Attribute | 0..1 |                 Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/HRCMScreeningResponse/FacilityPlace/Address/postcode | postCode |  |  |
| IE3A20 | 4 | Attribute | 0..1 |                 Street additional line | an..70 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A20 | 4 | Attribute | 0..1 |                 Number | an..35 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/streetNumber | number |  |  |
| IE3A20 | 4 | Attribute | 0..1 |                 P.O. Box | an..70 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/poBox | poBox |  |  |
| IE3A20 | 2 | Composition | 0..1 |         Transport document (House level) |  | O |  |  | TransportContractDocument | 30B | CIS/HRCMScreeningResponse/TransportContractDocument | transportDocumentHouseLevel |  |  |
| IE3A20 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/HRCMScreeningResponse/TransportContractDocument/identification | documentNumber |  |  |
| IE3A20 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/HRCMScreeningResponse/TransportContractDocument/type | type | CL754 | CL754 |
