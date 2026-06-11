---
title: "IE3F20 - part 4"
source_title: "IE3F20"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F20.md"
section: "IE3F20"
chunk: 4
chunk_count: 5
generated: true
---

# IE3F20 - part 4

Source document: IE3F20
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F20.md`
Chunk: 4 of 5

| IE3F20 | 5 | Attribute | 0..1 |                     Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/HouseConsignment/NotifyParty/Address/countrySubDivisionName | subDivision |  |  |
| IE3F20 | 5 | Attribute | 0..1 |                     Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/HouseConsignment/NotifyParty/Address/line | street |  |  |
| IE3F20 | 5 | Attribute | 0..1 |                     Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/HouseConsignment/NotifyParty/Address/postcode | postCode |  |  |
| IE3F20 | 5 | Attribute | 0..1 |                     Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/NotifyParty/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F20 | 5 | Attribute | 0..1 |                     Number | an..35 | O |  |  |  |  | CIS/Consignment/HouseConsignment/NotifyParty/Address/streetNumber | number |  |  |
| IE3F20 | 5 | Attribute | 0..1 |                     P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/NotifyParty/Address/poBox | poBox |  |  |
| IE3F20 | 4 | Composition | 0..9 |                 Communication |  | O |  |  | Communication | 25A | CIS/Consignment/HouseConsignment/NotifyParty/Communication | communication |  |  |
| IE3F20 | 5 | Attribute | 1..1 |                     Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/HouseConsignment/NotifyParty/Communication/identification | identifier |  |  |
| IE3F20 | 5 | Attribute | 1..1 |                     Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/HouseConsignment/NotifyParty/Communication/type | type | CL707 | CL707 |
| IE3F20 | 3 | Composition | 1..1 |             Transport document (House level) |  | M |  |  | TransportContractDocument | 30B | CIS/Consignment/HouseConsignment/TransportContractDocument | transportDocumentHouseLevel |  |  |
| IE3F20 | 4 | Attribute | 1..1 |                 Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/Consignment/HouseConsignment/TransportContractDocument/identification | documentNumber |  |  |
| IE3F20 | 4 | Attribute | 1..1 |                 Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/Consignment/HouseConsignment/TransportContractDocument/type | type | CL754 | CL754 |
| IE3F20 | 3 | Composition | 0..9999 |             Transport equipment |  | C |  | C3014 | TransportEquipment | 31B | CIS/Consignment/HouseConsignment/TransportEquipment | transportEquipment |  |  |
| IE3F20 | 4 | Attribute | 1..1 |                 Container identification number | an..17 | M |  |  | Equipment identification number | 159 | CIS/Consignment/HouseConsignment/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3F20 | 3 | Composition | 0..1 |             Reference number/UCR |  | O |  |  | UCR | 35B | CIS/Consignment/HouseConsignment/UCR | referenceNumberUCR |  |  |
| IE3F20 | 4 | Attribute | 1..1 |                 Reference number/UCR | an..35 | M |  |  | Trader reference | 009 | CIS/Consignment/HouseConsignment/UCR/traderAssignedReference | referenceNumberUCR |  |  |
| IE3F20 | 2 | Composition | 1..1 |         Place of loading |  | M |  |  | LoadingLocation | 83A | CIS/Consignment/LoadingLocation | placeOfLoading |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3007 | Place of loading | L009 | CIS/Consignment/LoadingLocation/name | location |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Place of loading, coded | L010 | CIS/Consignment/LoadingLocation/identification | UNLOCODE | CL244 | CL244 |
| IE3F20 | 3 | Composition | 0..1 |             Address |  | C |  | C3007 | Address | 04A | CIS/Consignment/LoadingLocation/Address | address |  |  |
| IE3F20 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/LoadingLocation/Address/country | country | CL718 | CL718 |
| IE3F20 | 2 | Composition | 0..1 |         Notify party |  | O |  |  | NotifyParty | 89A | CIS/Consignment/NotifyParty | notifyParty |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Name | an..70 | M |  |  | Notify party | R045 | CIS/Consignment/NotifyParty/name | name |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             Identification number | an..17 | O | R3005 |  | Notify party, coded | R046 | CIS/Consignment/NotifyParty/identification | identificationNumber |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Type of person | n1 | M |  |  |  |  | CIS/Consignment/NotifyParty/identificationType | typeOfPerson | CL729 | CL729 |
| IE3F20 | 3 | Composition | 1..1 |             Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/NotifyParty/Address | address |  |  |
| IE3F20 | 4 | Attribute | 1..1 |                 City | an..35 | M |  |  | City name | 241 | CIS/Consignment/NotifyParty/Address/cityName | city |  |  |
| IE3F20 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/NotifyParty/Address/country | country | CL718 | CL718 |
| IE3F20 | 4 | Attribute | 0..1 |                 Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/NotifyParty/Address/countrySubDivisionName | subDivision |  |  |
| IE3F20 | 4 | Attribute | 0..1 |                 Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/NotifyParty/Address/line | street |  |  |
| IE3F20 | 4 | Attribute | 0..1 |                 Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/NotifyParty/Address/postcode | postCode |  |  |
| IE3F20 | 4 | Attribute | 0..1 |                 Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/NotifyParty/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F20 | 4 | Attribute | 0..1 |                 Number | an..35 | O |  |  |  |  | CIS/Consignment/NotifyParty/Address/streetNumber | number |  |  |
| IE3F20 | 4 | Attribute | 0..1 |                 P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/NotifyParty/Address/poBox | poBox |  |  |
| IE3F20 | 3 | Composition | 0..9 |             Communication |  | O |  |  | Communication | 25A | CIS/Consignment/NotifyParty/Communication | communication |  |  |
| IE3F20 | 4 | Attribute | 1..1 |                 Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/NotifyParty/Communication/identification | identifier |  |  |
| IE3F20 | 4 | Attribute | 1..1 |                 Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/NotifyParty/Communication/type | type | CL707 | CL707 |
| IE3F20 | 2 | Composition | 1..1 |         Transport document (Master level) |  | M |  |  | TransportContractDocument | 30B | CIS/Consignment/TransportContractDocument | transportDocumentMasterLevel |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/Consignment/TransportContractDocument/identification | documentNumber |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/Consignment/TransportContractDocument/type | type | CL754 | CL754 |
| IE3F20 | 2 | Composition | 0..9999 |         Transport equipment |  | C |  | C3013 | TransportEquipment | 31B | CIS/Consignment/TransportEquipment | transportEquipment |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Container identification number | an..17 | M |  |  | Equipment identification number | 159 | CIS/Consignment/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3F20 | 2 | Composition | 0..1 |         Reference number/UCR |  | O |  |  | UCR | 35B | CIS/Consignment/UCR | referenceNumberUCR |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Reference number/UCR | an..35 | M |  |  | Trader reference | 009 | CIS/Consignment/UCR/traderAssignedReference | referenceNumberUCR |  |  |
| IE3F20 | 2 | Composition | 1..1 |         Place of unloading |  | M |  |  | UnloadingLocation | 38B | CIS/Consignment/UnloadingLocation | placeOfUnloading |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3008 | Place of discharge | L012 | CIS/Consignment/UnloadingLocation/name | location |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Place of discharge, coded | L013 | CIS/Consignment/UnloadingLocation/identification | UNLOCODE | CL244 | CL244 |
| IE3F20 | 3 | Composition | 0..1 |             Address |  | C |  | C3008 | Address | 04A | CIS/Consignment/UnloadingLocation/Address | address |  |  |
| IE3F20 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/UnloadingLocation/Address/country | country | CL718 | CL718 |
| IE3F20 | 1 | Composition | 1..1 |     Declarant |  | M |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3F20 | 2 | Attribute | 1..1 |         Name | an..70 | M |  |  | Declarant name | R124 | CIS/Declarant/name | name |  |  |
| IE3F20 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005 |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3F20 | 2 | Composition | 1..1 |         Address |  | M | R3002 |  | Address | 04A | CIS/Declarant/Address | address |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             City | an..35 | M |  |  | City name | 241 | CIS/Declarant/Address/cityName | city |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Country | a2 | M |  |  | Country, coded | 242 | CIS/Declarant/Address/country | country | CL718 | CL718 |
| IE3F20 | 3 | Attribute | 0..1 |             Sub-division | an..35 | Ο |  |  | Country sub-entity name | 243 | CIS/Declarant/Address/countrySubDivisionName | subDivision |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Declarant/Address/line | street |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Declarant/Address/postcode | postCode |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             Street additional line | an..70 | O |  |  |  |  | CIS/Declarant/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             Number | an..35 | O |  |  |  |  | CIS/Declarant/Address/streetNumber | number |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             P.O. Box | an..70 | O |  |  |  |  | CIS/Declarant/Address/poBox | poBox |  |  |
| IE3F20 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/Declarant/Communication | communication |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Declarant/Communication/identification | identifier |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/Declarant/Communication/type | type | CL707 | CL707 |
| IE3F20 | 1 | Composition | 1..1 |     Customs office of first entry |  | M |  |  | EntryOffice | 51A | CIS/EntryOffice | customsOfficeOfFirstEntry |  |  |
| IE3F20 | 2 | Attribute | 1..1 |         Reference number | an8 | M |  |  | Office of entry, coded | G004 | CIS/EntryOffice/identification | referenceNumber | CL141 | CL141 |
| IE3F20 | 1 | Composition | 0..99 |     HRCM screening results |  | O |  |  |  |  | CIS/HRCMScreeningResponse | hrcmScreeningResults |  |  |
| IE3F20 | 2 | Attribute | 1..1 |         Result | an..3 | M |  |  |  |  | CIS/HRCMScreeningResponse/result | result | CL726 | CL726 |
| IE3F20 | 2 | Composition | 2..9 |         Screening method |  | M |  |  |  |  | CIS/HRCMScreeningResponse/ScreeningMethod | screeningMethod |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Method | an..3 | M |  |  |  |  | CIS/HRCMScreeningResponse/ScreeningMethod/method | method | CL724 | CL724 |
| IE3F20 | 2 | Composition | 0..99 |         Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/HRCMScreeningResponse/AdditionalInformation | additionalInformation |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             Code | an..17 | O |  |  | Additional statement code | 226 | CIS/HRCMScreeningResponse/AdditionalInformation/statement | code | CL752 | CL752 |
| IE3F20 | 3 | Attribute | 0..1 |             Text | an..512 | O |  |  | Additional statement text | 225 | CIS/HRCMScreeningResponse/AdditionalInformation/statementDescription | text |  |  |
| IE3F20 | 3 | Attribute | 0..1 |             Type | an..3 | O |  |  | Additional statement type | 369 | CIS/HRCMScreeningResponse/AdditionalInformation/statementType | type | CL703 | CL703 |
| IE3F20 | 2 | Composition | 0..1 |         Authorized person |  | O |  |  | AuthorisedPerson | 13A | CIS/HRCMScreeningResponse/AuthorisedPerson | authorizedPerson |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Name | an..70 | M |  |  | Authorised person | R067 | CIS/HRCMScreeningResponse/AuthorisedPerson/name | name |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Identification number | an..17 | M |  |  | Authorised person, coded | R068 | CIS/HRCMScreeningResponse/AuthorisedPerson/identification | identificationNumber |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Type of person | n1 | M |  |  |  |  | CIS/HRCMScreeningResponse/AuthorisedPerson/identificationType | typeOfPerson | CL729 | CL729 |
| IE3F20 | 2 | Composition | 0..9 |         Binary file |  | O |  |  | BinaryFile | 72B | CIS/HRCMScreeningResponse/BinaryFile | binaryFile |  |  |
| IE3F20 | 3 | Attribute | 1..1 |             Identification | an..256 | M |  |  | Binary File Identifier | M001 | CIS/HRCMScreeningResponse/BinaryFile/identification |  |  |  |
