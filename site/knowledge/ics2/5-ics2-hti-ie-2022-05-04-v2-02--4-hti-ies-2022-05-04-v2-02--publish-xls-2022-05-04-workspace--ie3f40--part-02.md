---
title: "IE3F40 - part 2"
source_title: "IE3F40"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F40.md"
section: "IE3F40"
chunk: 2
chunk_count: 2
generated: true
---

# IE3F40 - part 2

Source document: IE3F40
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F40.md`
Chunk: 2 of 2

| IE3F40 | 4 | Attribute | 0..1 |                 Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/Consignor/Address/countrySubDivisionName | subDivision |  |  |
| IE3F40 | 4 | Attribute | 0..1 |                 Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/Consignor/Address/line | street |  |  |
| IE3F40 | 4 | Attribute | 0..1 |                 Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/Consignor/Address/postcode | postCode |  |  |
| IE3F40 | 4 | Attribute | 0..1 |                 Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/Consignor/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F40 | 4 | Attribute | 0..1 |                 Number | an..35 | O |  |  |  |  | CIS/Consignment/Consignor/Address/streetNumber | number |  |  |
| IE3F40 | 4 | Attribute | 0..1 |                 P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/Consignor/Address/poBox | poBox |  |  |
| IE3F40 | 3 | Composition | 0..9 |             Communication |  | O |  |  | Communication | 25A | CIS/Consignment/Consignor/Communication | communication |  |  |
| IE3F40 | 4 | Attribute | 1..1 |                 Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/Consignor/Communication/identification | identifier |  |  |
| IE3F40 | 4 | Attribute | 1..1 |                 Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/Consignor/Communication/type | type | CL707 | CL707 |
| IE3F40 | 2 | Composition | 1..1 |         Place of loading |  | M |  |  | LoadingLocation | 83A | CIS/Consignment/LoadingLocation | placeOfLoading |  |  |
| IE3F40 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3007 | Place of loading | L009 | CIS/Consignment/LoadingLocation/name | location |  |  |
| IE3F40 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Place of loading, coded | L010 | CIS/Consignment/LoadingLocation/identification | UNLOCODE | CL244 | CL244 |
| IE3F40 | 3 | Composition | 0..1 |             Address |  | C |  | C3007 | Address | 04A | CIS/Consignment/LoadingLocation/Address | address |  |  |
| IE3F40 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/LoadingLocation/Address/country | country | CL718 | CL718 |
| IE3F40 | 2 | Composition | 0..99 |         Passive border transport means |  | O |  |  | PassiveTransportMeans | 19C | CIS/Consignment/PassiveTransportMeans | passiveBorderTransportMeans |  |  |
| IE3F40 | 3 | Attribute | 1..1 |             Identification number | an..35 | M |  |  | Identification of passive transport means, coded | T026 | CIS/Consignment/PassiveTransportMeans/identification | identificationNumber |  |  |
| IE3F40 | 3 | Attribute | 1..1 |             Type of identification | n2 | M |  |  | Type of identification of passive transport means, coded | T027 | CIS/Consignment/PassiveTransportMeans/identificationType | typeOfIdentification | CL750 | CL750 |
| IE3F40 | 3 | Attribute | 1..1 |             Type of means of transport | an..4 | M |  |  |  |  | CIS/Consignment/PassiveTransportMeans/type | typeOfMeansOfTransport | CL751 | CL751 |
| IE3F40 | 3 | Attribute | 1..1 |             Nationality | a2 | M |  |  | Nationality of passive means of transport, coded | T028 | CIS/Consignment/PassiveTransportMeans/registrationNationality | nationality | CL718 | CL718 |
| IE3F40 | 2 | Composition | 1..1 |         Transport document (Master level) |  | M |  |  | TransportContractDocument | 30B | CIS/Consignment/TransportContractDocument | transportDocumentMasterLevel |  |  |
| IE3F40 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/Consignment/TransportContractDocument/identification | documentNumber |  |  |
| IE3F40 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/Consignment/TransportContractDocument/type | type | CL754 | CL754 |
| IE3F40 | 2 | Composition | 0..9999 |         Transport equipment |  | C |  | C3050 | TransportEquipment | 31B | CIS/Consignment/TransportEquipment | transportEquipment |  |  |
| IE3F40 | 3 | Attribute | 1..1 |             Container identification number | an..17 | M |  |  | Equipment identification number | 159 | CIS/Consignment/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3F40 | 3 | Attribute | 0..1 |             Number of seals | n..4 | O |  |  | Number of seals | 227 | CIS/Consignment/TransportEquipment/sealsAffixed | numberOfSeals |  |  |
| IE3F40 | 3 | Composition | 0..9 |             Seal |  | C |  | C3015 | Seal | 44B | CIS/Consignment/TransportEquipment/Seal | seal |  |  |
| IE3F40 | 4 | Attribute | 1..1 |                 Identifier | an..20 | M |  |  | Seal number | 165 | CIS/Consignment/TransportEquipment/Seal/identification | identifier |  |  |
| IE3F40 | 2 | Composition | 0..1 |         Reference number/UCR |  | O |  |  | UCR | 35B | CIS/Consignment/UCR | referenceNumberUCR |  |  |
| IE3F40 | 3 | Attribute | 1..1 |             Reference number/UCR | an..35 | M |  |  | Trader reference | 009 | CIS/Consignment/UCR/traderAssignedReference | referenceNumberUCR |  |  |
| IE3F40 | 2 | Composition | 1..1 |         Place of unloading |  | M |  |  | UnloadingLocation | 38B | CIS/Consignment/UnloadingLocation | placeOfUnloading |  |  |
| IE3F40 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3008 | Place of discharge | L012 | CIS/Consignment/UnloadingLocation/name | location |  |  |
| IE3F40 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Place of discharge, coded | L013 | CIS/Consignment/UnloadingLocation/identification | UNLOCODE | CL244 | CL244 |
| IE3F40 | 3 | Composition | 0..1 |             Address |  | C |  | C3008 | Address | 04A | CIS/Consignment/UnloadingLocation/Address | address |  |  |
| IE3F40 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/UnloadingLocation/Address/country | country | CL718 | CL718 |
| IE3F40 | 1 | Composition | 1..1 |     Declarant |  | M |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3F40 | 2 | Attribute | 1..1 |         Name | an..70 | M |  |  | Declarant name | R124 | CIS/Declarant/name | name |  |  |
| IE3F40 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005 |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3F40 | 2 | Composition | 1..1 |         Address |  | M | R3002 |  | Address | 04A | CIS/Declarant/Address | address |  |  |
| IE3F40 | 3 | Attribute | 1..1 |             City | an..35 | M |  |  | City name | 241 | CIS/Declarant/Address/cityName | city |  |  |
| IE3F40 | 3 | Attribute | 1..1 |             Country | a2 | M |  |  | Country, coded | 242 | CIS/Declarant/Address/country | country | CL718 | CL718 |
| IE3F40 | 3 | Attribute | 0..1 |             Sub-division | an..35 | Ο |  |  | Country sub-entity name | 243 | CIS/Declarant/Address/countrySubDivisionName | subDivision |  |  |
| IE3F40 | 3 | Attribute | 0..1 |             Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Declarant/Address/line | street |  |  |
| IE3F40 | 3 | Attribute | 0..1 |             Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Declarant/Address/postcode | postCode |  |  |
| IE3F40 | 3 | Attribute | 0..1 |             Street additional line | an..70 | O |  |  |  |  | CIS/Declarant/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F40 | 3 | Attribute | 0..1 |             Number | an..35 | O |  |  |  |  | CIS/Declarant/Address/streetNumber | number |  |  |
| IE3F40 | 3 | Attribute | 0..1 |             P.O. Box | an..70 | O |  |  |  |  | CIS/Declarant/Address/poBox | poBox |  |  |
| IE3F40 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/Declarant/Communication | communication |  |  |
| IE3F40 | 3 | Attribute | 1..1 |             Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Declarant/Communication/identification | identifier |  |  |
| IE3F40 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/Declarant/Communication/type | type | CL707 | CL707 |
| IE3F40 | 1 | Composition | 1..1 |     Customs office of first entry |  | M |  |  | EntryOffice | 51A | CIS/EntryOffice | customsOfficeOfFirstEntry |  |  |
| IE3F40 | 2 | Attribute | 1..1 |         Reference number | an8 | M |  |  | Office of entry, coded | G004 | CIS/EntryOffice/identification | referenceNumber | CL141 | CL141 |
