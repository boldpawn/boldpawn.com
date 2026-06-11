---
title: "IE3A28 - part 3"
source_title: "IE3A28"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A28.md"
section: "IE3A28"
chunk: 3
chunk_count: 3
generated: true
---

# IE3A28 - part 3

Source document: IE3A28
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A28.md`
Chunk: 3 of 3

| IE3A28 | 5 | Attribute | 1..1 |                     Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/HouseConsignment/NotifyParty/Communication/identification | identifier |  |  |
| IE3A28 | 5 | Attribute | 1..1 |                     Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/HouseConsignment/NotifyParty/Communication/type | type | CL707 | CL707 |
| IE3A28 | 3 | Composition | 1..1 |             Transport document (House level) |  | M |  |  | TransportContractDocument | 30B | CIS/Consignment/HouseConsignment/TransportContractDocument | transportDocumentHouseLevel |  |  |
| IE3A28 | 4 | Attribute | 1..1 |                 Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/Consignment/HouseConsignment/TransportContractDocument/identification | documentNumber |  |  |
| IE3A28 | 4 | Attribute | 1..1 |                 Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/Consignment/HouseConsignment/TransportContractDocument/type | type | CL754 | CL754 |
| IE3A28 | 3 | Composition | 0..9999 |             Transport equipment |  | C |  | C3014 | TransportEquipment | 31B | CIS/Consignment/HouseConsignment/TransportEquipment | transportEquipment |  |  |
| IE3A28 | 4 | Attribute | 1..1 |                 Container identification number | an..17 | M |  |  | Equipment identification number | 159 | CIS/Consignment/HouseConsignment/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3A28 | 3 | Composition | 0..1 |             Reference number/UCR |  | O |  |  | UCR | 35B | CIS/Consignment/HouseConsignment/UCR | referenceNumberUCR |  |  |
| IE3A28 | 4 | Attribute | 1..1 |                 Reference number/UCR | an..35 | M |  |  | Trader reference | 009 | CIS/Consignment/HouseConsignment/UCR/traderAssignedReference | referenceNumberUCR |  |  |
| IE3A28 | 2 | Composition | 1..1 |         Place of loading |  | M |  |  | LoadingLocation | 83A | CIS/Consignment/LoadingLocation | placeOfLoading |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3007 | Place of loading | L009 | CIS/Consignment/LoadingLocation/name | location |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Place of loading, coded | L010 | CIS/Consignment/LoadingLocation/identification | UNLOCODE | CL244 | CL244 |
| IE3A28 | 3 | Composition | 0..1 |             Address |  | C |  | C3007 | Address | 04A | CIS/Consignment/LoadingLocation/Address | address |  |  |
| IE3A28 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/LoadingLocation/Address/country | country | CL718 | CL718 |
| IE3A28 | 2 | Composition | 1..1 |         Transport document (Master level) |  | M |  |  | TransportContractDocument | 30B | CIS/Consignment/TransportContractDocument | transportDocumentMasterLevel |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             Reference number | an..70 | M | R3024 |  | Transport document number | D023 | CIS/Consignment/TransportContractDocument/identification | documentNumber |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/Consignment/TransportContractDocument/type | type | CL754 | CL754 |
| IE3A28 | 2 | Composition | 1..1 |         Place of unloading |  | M |  |  | UnloadingLocation | 38B | CIS/Consignment/UnloadingLocation | placeOfUnloading |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3008 | Place of discharge | L012 | CIS/Consignment/UnloadingLocation/name | location |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Place of discharge, coded | L013 | CIS/Consignment/UnloadingLocation/identification | UNLOCODE | CL244 | CL244 |
| IE3A28 | 3 | Composition | 0..1 |             Address |  | C |  | C3008 | Address | 04A | CIS/Consignment/UnloadingLocation/Address | address |  |  |
| IE3A28 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/UnloadingLocation/Address/country | country | CL718 | CL718 |
| IE3A28 | 1 | Composition | 1..1 |     Declarant |  | M |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3A28 | 2 | Attribute | 1..1 |         Name | an..70 | M |  |  | Declarant name | R124 | CIS/Declarant/name | name |  |  |
| IE3A28 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005<br>R3024 |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3A28 | 2 | Composition | 1..1 |         Address |  | M | R3002 |  | Address | 04A | CIS/Declarant/Address | address |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             City | an..35 | M |  |  | City name | 241 | CIS/Declarant/Address/cityName | city |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             Country | a2 | M |  |  | Country, coded | 242 | CIS/Declarant/Address/country | country | CL718 | CL718 |
| IE3A28 | 3 | Attribute | 0..1 |             Sub-division | an..35 | Ο |  |  | Country sub-entity name | 243 | CIS/Declarant/Address/countrySubDivisionName | subDivision |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Declarant/Address/line | street |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Declarant/Address/postcode | postCode |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             Street additional line | an..70 | O |  |  |  |  | CIS/Declarant/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             Number | an..35 | O |  |  |  |  | CIS/Declarant/Address/streetNumber | number |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             P.O. Box | an..70 | O |  |  |  |  | CIS/Declarant/Address/poBox | poBox |  |  |
| IE3A28 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/Declarant/Communication | communication |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Declarant/Communication/identification | identifier |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/Declarant/Communication/type | type | CL707 | CL707 |
| IE3A28 | 1 | Composition | 1..1 |     Customs office of first entry |  | M |  |  | EntryOffice | 51A | CIS/EntryOffice | customsOfficeOfFirstEntry |  |  |
| IE3A28 | 2 | Attribute | 1..1 |         Reference number | an8 | M | R3024 |  | Office of entry, coded | G004 | CIS/EntryOffice/identification | referenceNumber | CL141 | CL141 |
| IE3A28 | 1 | Composition | 0..99 |     HRCM screening results |  | O |  |  |  |  | CIS/HRCMScreeningResponse | hrcmScreeningResults |  |  |
| IE3A28 | 2 | Attribute | 1..1 |         Result | an..3 | M |  |  |  |  | CIS/HRCMScreeningResponse/result | result | CL726 | CL726 |
| IE3A28 | 2 | Composition | 2..9 |         Screening method |  | M |  |  |  |  | CIS/HRCMScreeningResponse/ScreeningMethod | screeningMethod |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             Method | an..3 | M |  |  |  |  | CIS/HRCMScreeningResponse/ScreeningMethod/method | method | CL724 | CL724 |
| IE3A28 | 2 | Composition | 0..99 |         Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/HRCMScreeningResponse/AdditionalInformation | additionalInformation |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             Code | an..17 | O |  |  | Additional statement code | 226 | CIS/HRCMScreeningResponse/AdditionalInformation/statement | code | CL752 | CL752 |
| IE3A28 | 3 | Attribute | 0..1 |             Text | an..512 | O |  |  | Additional statement text | 225 | CIS/HRCMScreeningResponse/AdditionalInformation/statementDescription | text |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             Type | an..3 | O |  |  | Additional statement type | 369 | CIS/HRCMScreeningResponse/AdditionalInformation/statementType | type | CL703 | CL703 |
| IE3A28 | 2 | Composition | 0..1 |         Authorized person |  | O |  |  | AuthorisedPerson | 13A | CIS/HRCMScreeningResponse/AuthorisedPerson | authorizedPerson |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             Name | an..70 | M |  |  | Authorised person | R067 | CIS/HRCMScreeningResponse/AuthorisedPerson/name | name |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             Identification number | an..17 | M |  |  | Authorised person, coded | R068 | CIS/HRCMScreeningResponse/AuthorisedPerson/identification | identificationNumber |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             Type of person | n1 | M |  |  |  |  | CIS/HRCMScreeningResponse/AuthorisedPerson/identificationType | typeOfPerson | CL729 | CL729 |
| IE3A28 | 2 | Composition | 0..9 |         Binary file |  | O |  |  | BinaryFile | 72B | CIS/HRCMScreeningResponse/BinaryFile | binaryFile |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             Identification | an..256 | M |  |  | Binary File Identifier | M001 | CIS/HRCMScreeningResponse/BinaryFile/identification |  |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             Filename | an..256 | M |  |  | Binary File Name | M005 | CIS/HRCMScreeningResponse/BinaryFile/fileName | filename |  |  |
| IE3A28 | 3 | Attribute | 1..1 |             MIME | an..70 | M |  |  | Binary File MIME, coded | M007 | CIS/HRCMScreeningResponse/BinaryFile/MIME | MIME |  |  |
| IE3A28 | 3 | Attribute | 0..1 |             Description | an..256 | O |  |  | Binary File Description | M012 | CIS/HRCMScreeningResponse/BinaryFile/description | description |  |  |
| IE3A28 | 2 | Composition | 0..1 |         Facility place |  | O |  |  | FacilityPlace | 58A | CIS/HRCMScreeningResponse/FacilityPlace | facilityPlace |  |  |
| IE3A28 | 3 | Composition | 1..1 |             Address |  | M |  |  | Address | 04A | CIS/HRCMScreeningResponse/FacilityPlace/Address | address |  |  |
| IE3A28 | 4 | Attribute | 1..1 |                 City | an..35 | M |  |  | City name | 241 | CIS/HRCMScreeningResponse/FacilityPlace/Address/cityName | city |  |  |
| IE3A28 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/HRCMScreeningResponse/FacilityPlace/Address/country | country | CL718 | CL718 |
| IE3A28 | 4 | Attribute | 0..1 |                 Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/HRCMScreeningResponse/FacilityPlace/Address/countrySubDivisionName | subDivision |  |  |
| IE3A28 | 4 | Attribute | 0..1 |                 Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/HRCMScreeningResponse/FacilityPlace/Address/line | street |  |  |
| IE3A28 | 4 | Attribute | 0..1 |                 Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/HRCMScreeningResponse/FacilityPlace/Address/postcode | postCode |  |  |
| IE3A28 | 4 | Attribute | 0..1 |                 Street additional line | an..70 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A28 | 4 | Attribute | 0..1 |                 Number | an..35 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/streetNumber | number |  |  |
| IE3A28 | 4 | Attribute | 0..1 |                 P.O. Box | an..70 | O |  |  |  |  | CIS/HRCMScreeningResponse/FacilityPlace/Address/poBox | poBox |  |  |
