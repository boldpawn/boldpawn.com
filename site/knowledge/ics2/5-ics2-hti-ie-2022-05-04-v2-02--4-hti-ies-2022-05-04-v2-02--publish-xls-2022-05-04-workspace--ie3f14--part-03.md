---
title: "IE3F14 - part 3"
source_title: "IE3F14"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F14.md"
section: "IE3F14"
chunk: 3
chunk_count: 3
generated: true
---

# IE3F14 - part 3

Source document: IE3F14
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F14.md`
Chunk: 3 of 3

| IE3F14 | 4 | Attribute | 1..1 |                 Type of identification | n2 | M |  |  | Type of identification of passive transport means, coded | T027 | CIS/Consignment/HouseConsignment/PassiveTransportMeans/identificationType | typeOfIdentification | CL750 | CL750 |
| IE3F14 | 4 | Attribute | 1..1 |                 Type of means of transport | an..4 | M |  |  |  |  | CIS/Consignment/HouseConsignment/PassiveTransportMeans/type | typeOfMeansOfTransport | CL751 | CL751 |
| IE3F14 | 4 | Attribute | 0..1 |                 Nationality | a2 | O |  |  | Nationality of passive means of transport, coded | T028 | CIS/Consignment/HouseConsignment/PassiveTransportMeans/registrationNationality | nationality | CL718 | CL718 |
| IE3F14 | 3 | Composition | 0..99 |             Supplementary declarant |  | C |  | C3020 | SupplementaryDeclarant | 18C | CIS/Consignment/HouseConsignment/SupplementaryDeclarant | supplementaryDeclarant |  |  |
| IE3F14 | 4 | Attribute | 1..1 |                 Identification number | an..17 | M | R3005 |  | Supplementary declarant, coded | R146 | CIS/Consignment/HouseConsignment/SupplementaryDeclarant/identification | identificationNumber |  |  |
| IE3F14 | 4 | Attribute | 1..1 |                 Supplementary filing type | an..3 | M |  |  |  |  | CIS/Consignment/HouseConsignment/SupplementaryDeclarant/supplementaryFilingType | supplementaryFilingType | CL745 | CL745 |
| IE3F14 | 3 | Composition | 1..1 |             Transport document (House level) |  | M |  |  | TransportContractDocument | 30B | CIS/Consignment/HouseConsignment/TransportContractDocument | transportDocumentHouseLevel |  |  |
| IE3F14 | 4 | Attribute | 1..1 |                 Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/Consignment/HouseConsignment/TransportContractDocument/identification | documentNumber |  |  |
| IE3F14 | 4 | Attribute | 1..1 |                 Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/Consignment/HouseConsignment/TransportContractDocument/type | type | CL754 | CL754 |
| IE3F14 | 3 | Composition | 0..9999 |             Transport equipment |  | C |  | C3014 | TransportEquipment | 31B | CIS/Consignment/HouseConsignment/TransportEquipment | transportEquipment |  |  |
| IE3F14 | 4 | Attribute | 1..1 |                 Container size and type | an..10 | M |  |  | Equipment size and type identification | 152 | CIS/Consignment/HouseConsignment/TransportEquipment/characteristic | containerSizeAndType | CL710 | CL710 |
| IE3F14 | 4 | Attribute | 1..1 |                 Container packed status | an..3 | M |  |  | Transport equipment loaded status | 154 | CIS/Consignment/HouseConsignment/TransportEquipment/fullness | containerPackedStatus | CL709 | CL709 |
| IE3F14 | 4 | Attribute | 1..1 |                 Container supplier type | an..3 | M |  |  | Equipment supplier type, coded | 151 | CIS/Consignment/HouseConsignment/TransportEquipment/supplierPartyType | containerSupplierType | CL711 | CL711 |
| IE3F14 | 4 | Attribute | 1..1 |                 Container identification number | an..17 | M | R3021 |  | Equipment identification number | 159 | CIS/Consignment/HouseConsignment/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3F14 | 4 | Attribute | 0..1 |                 Number of seals | n..4 | O |  |  | Number of seals | 227 | CIS/Consignment/HouseConsignment/TransportEquipment/sealsAffixed | numberOfSeals |  |  |
| IE3F14 | 4 | Composition | 0..99 |                 Seal |  | C |  | C3017 | Seal | 44B | CIS/Consignment/HouseConsignment/TransportEquipment/Seal | seal |  |  |
| IE3F14 | 5 | Attribute | 1..1 |                     Identifier | an..20 | M |  |  | Seal number | 165 | CIS/Consignment/HouseConsignment/TransportEquipment/Seal/identification | identifier |  |  |
| IE3F14 | 3 | Composition | 0..1 |             Reference number/UCR |  | O |  |  | UCR | 35B | CIS/Consignment/HouseConsignment/UCR | referenceNumberUCR |  |  |
| IE3F14 | 4 | Attribute | 1..1 |                 Reference number/UCR | an..35 | M |  |  | Trader reference | 009 | CIS/Consignment/HouseConsignment/UCR/traderAssignedReference | referenceNumberUCR |  |  |
| IE3F14 | 1 | Composition | 1..1 |     Declarant |  | M |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3F14 | 2 | Attribute | 1..1 |         Name | an..70 | M |  |  | Declarant name | R124 | CIS/Declarant/name | name |  |  |
| IE3F14 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005 |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3F14 | 2 | Composition | 1..1 |         Address |  | M | R3002 |  | Address | 04A | CIS/Declarant/Address | address |  |  |
| IE3F14 | 3 | Attribute | 1..1 |             City | an..35 | M |  |  | City name | 241 | CIS/Declarant/Address/cityName | city |  |  |
| IE3F14 | 3 | Attribute | 1..1 |             Country | a2 | M |  |  | Country, coded | 242 | CIS/Declarant/Address/country | country | CL718 | CL718 |
| IE3F14 | 3 | Attribute | 0..1 |             Sub-division | an..35 | Ο |  |  | Country sub-entity name | 243 | CIS/Declarant/Address/countrySubDivisionName | subDivision |  |  |
| IE3F14 | 3 | Attribute | 0..1 |             Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Declarant/Address/line | street |  |  |
| IE3F14 | 3 | Attribute | 0..1 |             Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Declarant/Address/postcode | postCode |  |  |
| IE3F14 | 3 | Attribute | 0..1 |             Street additional line | an..70 | O |  |  |  |  | CIS/Declarant/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F14 | 3 | Attribute | 0..1 |             Number | an..35 | O |  |  |  |  | CIS/Declarant/Address/streetNumber | number |  |  |
| IE3F14 | 3 | Attribute | 0..1 |             P.O. Box | an..70 | O |  |  |  |  | CIS/Declarant/Address/poBox | poBox |  |  |
| IE3F14 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/Declarant/Communication | communication |  |  |
| IE3F14 | 3 | Attribute | 1..1 |             Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Declarant/Communication/identification | identifier |  |  |
| IE3F14 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/Declarant/Communication/type | type | CL707 | CL707 |
