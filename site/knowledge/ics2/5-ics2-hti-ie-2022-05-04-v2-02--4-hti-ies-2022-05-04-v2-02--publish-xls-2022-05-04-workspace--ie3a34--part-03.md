---
title: "IE3A34 - part 3"
source_title: "IE3A34"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A34.md"
section: "IE3A34"
chunk: 3
chunk_count: 3
generated: true
---

# IE3A34 - part 3

Source document: IE3A34
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A34.md`
Chunk: 3 of 3

| IE3A34 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/GoodsReceiptPlace/Address/country | country | CL718 | CL718 |
| IE3A34 | 3 | Composition | 2..99 |             Countries of routing of consignment |  | M | R3022 |  | Itinerary | 81A | CIS/Consignment/HouseConsignment/Itinerary | countriesOfRoutingOfConsignment |  |  |
| IE3A34 | 4 | Attribute | 1..1 |                 Sequence number | n..2 | M |  |  | Sequence number | 006 | CIS/Consignment/HouseConsignment/Itinerary/sequence | sequenceNumber |  |  |
| IE3A34 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country(ies ) of routing, coded | 064 | CIS/Consignment/HouseConsignment/Itinerary/routingCountry | country | CL718 | CL718 |
| IE3A34 | 3 | Composition | 0..99 |             Passive border transport means |  | O |  |  | PassiveTransportMeans | 19C | CIS/Consignment/HouseConsignment/PassiveTransportMeans | passiveBorderTransportMeans |  |  |
| IE3A34 | 4 | Attribute | 1..1 |                 Identification number | an..35 | M |  |  | Identification of passive transport means, coded | T026 | CIS/Consignment/HouseConsignment/PassiveTransportMeans/identification | identificationNumber |  |  |
| IE3A34 | 4 | Attribute | 1..1 |                 Type of identification | n2 | M |  |  | Type of identification of passive transport means, coded | T027 | CIS/Consignment/HouseConsignment/PassiveTransportMeans/identificationType | typeOfIdentification | CL750 | CL750 |
| IE3A34 | 4 | Attribute | 1..1 |                 Type of means of transport | an..4 | M |  |  |  |  | CIS/Consignment/HouseConsignment/PassiveTransportMeans/type | typeOfMeansOfTransport | CL751 | CL751 |
| IE3A34 | 4 | Attribute | 1..1 |                 Nationality | a2 | M |  |  | Nationality of passive means of transport, coded | T028 | CIS/Consignment/HouseConsignment/PassiveTransportMeans/registrationNationality | nationality | CL718 | CL718 |
| IE3A34 | 3 | Composition | 1..1 |             Transport document (House level) |  | M |  |  | TransportContractDocument | 30B | CIS/Consignment/HouseConsignment/TransportContractDocument | transportDocumentHouseLevel |  |  |
| IE3A34 | 4 | Attribute | 1..1 |                 Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/Consignment/HouseConsignment/TransportContractDocument/identification | documentNumber |  |  |
| IE3A34 | 4 | Attribute | 1..1 |                 Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/Consignment/HouseConsignment/TransportContractDocument/type | type | CL754 | CL754 |
| IE3A34 | 3 | Composition | 0..9999 |             Transport equipment |  | C |  | C3051 | TransportEquipment | 31B | CIS/Consignment/HouseConsignment/TransportEquipment | transportEquipment |  |  |
| IE3A34 | 4 | Attribute | 1..1 |                 Container size and type | an..10 | M |  |  | Equipment size and type identification | 152 | CIS/Consignment/HouseConsignment/TransportEquipment/characteristic | containerSizeAndType | CL710 | CL710 |
| IE3A34 | 4 | Attribute | 1..1 |                 Container packed status | an..3 | M |  |  | Transport equipment loaded status | 154 | CIS/Consignment/HouseConsignment/TransportEquipment/fullness | containerPackedStatus | CL709 | CL709 |
| IE3A34 | 4 | Attribute | 1..1 |                 Container supplier type | an..3 | M |  |  | Equipment supplier type, coded | 151 | CIS/Consignment/HouseConsignment/TransportEquipment/supplierPartyType | containerSupplierType | CL711 | CL711 |
| IE3A34 | 4 | Attribute | 1..1 |                 Container identification number | an..17 | M | R3021 |  | Equipment identification number | 159 | CIS/Consignment/HouseConsignment/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3A34 | 3 | Composition | 1..1 |             Reference number/UCR |  | M |  |  | UCR | 35B | CIS/Consignment/HouseConsignment/UCR | referenceNumberUCR |  |  |
| IE3A34 | 4 | Attribute | 1..1 |                 Reference number/UCR | an..35 | M |  |  | Trader reference | 009 | CIS/Consignment/HouseConsignment/UCR/traderAssignedReference | referenceNumberUCR |  |  |
| IE3A34 | 2 | Composition | 1..1 |         Place of loading |  | M |  |  | LoadingLocation | 83A | CIS/Consignment/LoadingLocation | placeOfLoading |  |  |
| IE3A34 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3007 | Place of loading | L009 | CIS/Consignment/LoadingLocation/name | location |  |  |
| IE3A34 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Place of loading, coded | L010 | CIS/Consignment/LoadingLocation/identification | UNLOCODE | CL244 | CL244 |
| IE3A34 | 3 | Composition | 0..1 |             Address |  | C |  | C3007 | Address | 04A | CIS/Consignment/LoadingLocation/Address | address |  |  |
| IE3A34 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/LoadingLocation/Address/country | country | CL718 | CL718 |
| IE3A34 | 2 | Composition | 1..1 |         Place of unloading |  | M |  |  | UnloadingLocation | 38B | CIS/Consignment/UnloadingLocation | placeOfUnloading |  |  |
| IE3A34 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3008 | Place of discharge | L012 | CIS/Consignment/UnloadingLocation/name | location |  |  |
| IE3A34 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Place of discharge, coded | L013 | CIS/Consignment/UnloadingLocation/identification | UNLOCODE | CL244 | CL244 |
| IE3A34 | 3 | Composition | 0..1 |             Address |  | C |  | C3008 | Address | 04A | CIS/Consignment/UnloadingLocation/Address | address |  |  |
| IE3A34 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/UnloadingLocation/Address/country | country | CL718 | CL718 |
| IE3A34 | 1 | Composition | 1..1 |     Declarant |  | M |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3A34 | 2 | Attribute | 1..1 |         Name | an..70 | M |  |  | Declarant name | R124 | CIS/Declarant/name | name |  |  |
| IE3A34 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005 |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3A34 | 2 | Composition | 1..1 |         Address |  | M | R3002 |  | Address | 04A | CIS/Declarant/Address | address |  |  |
| IE3A34 | 3 | Attribute | 1..1 |             City | an..35 | M |  |  | City name | 241 | CIS/Declarant/Address/cityName | city |  |  |
| IE3A34 | 3 | Attribute | 1..1 |             Country | a2 | M |  |  | Country, coded | 242 | CIS/Declarant/Address/country | country | CL718 | CL718 |
| IE3A34 | 3 | Attribute | 0..1 |             Sub-division | an..35 | Ο |  |  | Country sub-entity name | 243 | CIS/Declarant/Address/countrySubDivisionName | subDivision |  |  |
| IE3A34 | 3 | Attribute | 0..1 |             Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Declarant/Address/line | street |  |  |
| IE3A34 | 3 | Attribute | 0..1 |             Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Declarant/Address/postcode | postCode |  |  |
| IE3A34 | 3 | Attribute | 0..1 |             Street additional line | an..70 | O |  |  |  |  | CIS/Declarant/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A34 | 3 | Attribute | 0..1 |             Number | an..35 | O |  |  |  |  | CIS/Declarant/Address/streetNumber | number |  |  |
| IE3A34 | 3 | Attribute | 0..1 |             P.O. Box | an..70 | O |  |  |  |  | CIS/Declarant/Address/poBox | poBox |  |  |
| IE3A34 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/Declarant/Communication | communication |  |  |
| IE3A34 | 3 | Attribute | 1..1 |             Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Declarant/Communication/identification | identifier |  |  |
| IE3A34 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/Declarant/Communication/type | type | CL707 | CL707 |
| IE3A34 | 1 | Composition | 1..1 |     Customs office of first entry |  | M |  |  | EntryOffice | 51A | CIS/EntryOffice | customsOfficeOfFirstEntry |  |  |
| IE3A34 | 2 | Attribute | 1..1 |         Reference number | an8 | M |  |  | Office of entry, coded | G004 | CIS/EntryOffice/identification | referenceNumber | CL141 | CL141 |
