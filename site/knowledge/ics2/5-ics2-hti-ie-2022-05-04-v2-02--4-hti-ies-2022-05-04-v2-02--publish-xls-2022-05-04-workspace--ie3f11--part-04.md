---
title: "IE3F11 - part 4"
source_title: "IE3F11"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F11.md"
section: "IE3F11"
chunk: 4
chunk_count: 5
generated: true
---

# IE3F11 - part 4

Source document: IE3F11
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F11.md`
Chunk: 4 of 5

| IE3F11 | 4 | Attribute | 1..1 |                 Method of payment | a1 | M |  |  | Transport charges method of payment, coded | 098 | CIS/Consignment/HouseConsignment/Freight/paymentMethod | methodOfPayment | CL116 | CL116 |
| IE3F11 | 3 | Composition | 1..1 |             Place of delivery |  | M |  |  | GoodsReceiptPlace | 66A | CIS/Consignment/HouseConsignment/GoodsReceiptPlace | placeOfDelivery |  |  |
| IE3F11 | 4 | Attribute | 0..1 |                 Location | an..35 | C |  | C3006 | Goods receipt place | L003 | CIS/Consignment/HouseConsignment/GoodsReceiptPlace/name | location |  |  |
| IE3F11 | 4 | Attribute | 0..1 |                 UNLOCODE | an..17 | O |  |  | Goods receipt place, coded | L004 | CIS/Consignment/HouseConsignment/GoodsReceiptPlace/identification | UNLOCODE | CL244 | CL244 |
| IE3F11 | 4 | Composition | 0..1 |                 Address |  | C |  | C3006 | Address | 04A | CIS/Consignment/HouseConsignment/GoodsReceiptPlace/Address | address |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/GoodsReceiptPlace/Address/country | country | CL718 | CL718 |
| IE3F11 | 3 | Composition | 0..1 |             Goods shipment |  | C |  | C3020 | GoodsShipment | 67A | CIS/Consignment/HouseConsignment/GoodsShipment | goodsShipment |  |  |
| IE3F11 | 4 | Composition | 1..1 |                 Buyer |  | M |  |  | Buyer | 16A | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer | buyer |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Name | an..70 | M |  |  | Buyer - name | R009 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/name | name |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Identification number | an..17 | O |  |  | Buyer, coded | R010 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/identification | identificationNumber |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Type of person | n1 | M |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/identificationType | typeOfPerson | CL729 | CL729 |
| IE3F11 | 5 | Composition | 1..1 |                     Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address | address |  |  |
| IE3F11 | 6 | Attribute | 1..1 |                         City | an..35 | M |  |  | City name | 241 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/cityName | city |  |  |
| IE3F11 | 6 | Attribute | 1..1 |                         Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/country | country | CL718 | CL718 |
| IE3F11 | 6 | Attribute | 0..1 |                         Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/countrySubDivisionName | subDivision |  |  |
| IE3F11 | 6 | Attribute | 0..1 |                         Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/line | street |  |  |
| IE3F11 | 6 | Attribute | 0..1 |                         Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/postcode | postCode |  |  |
| IE3F11 | 6 | Attribute | 0..1 |                         Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F11 | 6 | Attribute | 0..1 |                         Number | an..35 | O |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/streetNumber | number |  |  |
| IE3F11 | 6 | Attribute | 0..1 |                         P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/poBox | poBox |  |  |
| IE3F11 | 5 | Composition | 0..9 |                     Communication |  | O |  |  | Communication | 25A | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Communication | communication |  |  |
| IE3F11 | 6 | Attribute | 1..1 |                         Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Communication/identification | identifier |  |  |
| IE3F11 | 6 | Attribute | 1..1 |                         Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Communication/type | type | CL707 | CL707 |
| IE3F11 | 4 | Composition | 1..1 |                 Seller |  | M |  |  | Seller | 09B | CIS/Consignment/HouseConsignment/GoodsShipment/Seller | seller |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Name | an..70 | M |  |  | Seller - name | R050 | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/name | name |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Identification number | an..17 | O |  |  | Seller, coded | R051 | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/identification | identificationNumber |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Type of person | n1 | M |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/identificationType | typeOfPerson | CL729 | CL729 |
| IE3F11 | 5 | Composition | 1..1 |                     Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Address | address |  |  |
| IE3F11 | 6 | Attribute | 1..1 |                         City | an..35 | M |  |  | City name | 241 | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Address/cityName | city |  |  |
| IE3F11 | 6 | Attribute | 1..1 |                         Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Address/country | country | CL718 | CL718 |
| IE3F11 | 6 | Attribute | 0..1 |                         Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Address/countrySubDivisionName | subDivision |  |  |
| IE3F11 | 6 | Attribute | 0..1 |                         Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Address/line | street |  |  |
| IE3F11 | 6 | Attribute | 0..1 |                         Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Address/postcode | postCode |  |  |
| IE3F11 | 6 | Attribute | 0..1 |                         Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F11 | 6 | Attribute | 0..1 |                         Number | an..35 | O |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Address/streetNumber | number |  |  |
| IE3F11 | 6 | Attribute | 0..1 |                         P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Address/poBox | poBox |  |  |
| IE3F11 | 5 | Composition | 0..9 |                     Communication |  | O |  |  | Communication | 25A | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Communication | communication |  |  |
| IE3F11 | 6 | Attribute | 1..1 |                         Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Communication/identification | identifier |  |  |
| IE3F11 | 6 | Attribute | 1..1 |                         Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/Communication/type | type | CL707 | CL707 |
| IE3F11 | 3 | Composition | 2..99 |             Countries of routing of consignment |  | M | R3022 |  | Itinerary | 81A | CIS/Consignment/HouseConsignment/Itinerary | countriesOfRoutingOfConsignment |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Sequence number | n..2 | M |  |  | Sequence number | 006 | CIS/Consignment/HouseConsignment/Itinerary/sequence | sequenceNumber |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country(ies ) of routing, coded | 064 | CIS/Consignment/HouseConsignment/Itinerary/routingCountry | country | CL718 | CL718 |
| IE3F11 | 3 | Composition | 0..1 |             Notify party |  | C |  | C3024 | NotifyParty | 89A | CIS/Consignment/HouseConsignment/NotifyParty | notifyParty |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Name | an..70 | M |  |  | Notify party | R045 | CIS/Consignment/HouseConsignment/NotifyParty/name | name |  |  |
| IE3F11 | 4 | Attribute | 0..1 |                 Identification number | an..17 | O | R3005 |  | Notify party, coded | R046 | CIS/Consignment/HouseConsignment/NotifyParty/identification | identificationNumber |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Type of person | n1 | M |  |  |  |  | CIS/Consignment/HouseConsignment/NotifyParty/identificationType | typeOfPerson | CL729 | CL729 |
| IE3F11 | 4 | Composition | 1..1 |                 Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/HouseConsignment/NotifyParty/Address | address |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     City | an..35 | M |  |  | City name | 241 | CIS/Consignment/HouseConsignment/NotifyParty/Address/cityName | city |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/NotifyParty/Address/country | country | CL718 | CL718 |
| IE3F11 | 5 | Attribute | 0..1 |                     Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/HouseConsignment/NotifyParty/Address/countrySubDivisionName | subDivision |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/HouseConsignment/NotifyParty/Address/line | street |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/HouseConsignment/NotifyParty/Address/postcode | postCode |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/NotifyParty/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Number | an..35 | O |  |  |  |  | CIS/Consignment/HouseConsignment/NotifyParty/Address/streetNumber | number |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/NotifyParty/Address/poBox | poBox |  |  |
| IE3F11 | 4 | Composition | 1..1 |                 Communication |  | M |  |  | Communication | 25A | CIS/Consignment/HouseConsignment/NotifyParty/Communication | communication |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/HouseConsignment/NotifyParty/Communication/identification | identifier |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/HouseConsignment/NotifyParty/Communication/type | type | CL707 | CL707 |
| IE3F11 | 3 | Composition | 0..99 |             Passive border transport means |  | O |  |  | PassiveTransportMeans | 19C | CIS/Consignment/HouseConsignment/PassiveTransportMeans | passiveBorderTransportMeans |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Identification number | an..35 | M |  |  | Identification of passive transport means, coded | T026 | CIS/Consignment/HouseConsignment/PassiveTransportMeans/identification | identificationNumber |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Type of identification | n2 | M |  |  | Type of identification of passive transport means, coded | T027 | CIS/Consignment/HouseConsignment/PassiveTransportMeans/identificationType | typeOfIdentification | CL750 | CL750 |
| IE3F11 | 4 | Attribute | 1..1 |                 Type of means of transport | an..4 | M |  |  |  |  | CIS/Consignment/HouseConsignment/PassiveTransportMeans/type | typeOfMeansOfTransport | CL751 | CL751 |
| IE3F11 | 4 | Attribute | 0..1 |                 Nationality | a2 | O |  |  | Nationality of passive means of transport, coded | T028 | CIS/Consignment/HouseConsignment/PassiveTransportMeans/registrationNationality | nationality | CL718 | CL718 |
| IE3F11 | 3 | Composition | 1..1 |             Transport document (House level) |  | M |  |  | TransportContractDocument | 30B | CIS/Consignment/HouseConsignment/TransportContractDocument | transportDocumentHouseLevel |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/Consignment/HouseConsignment/TransportContractDocument/identification | documentNumber |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/Consignment/HouseConsignment/TransportContractDocument/type | type | CL754 | CL754 |
| IE3F11 | 3 | Composition | 0..9999 |             Transport equipment |  | C |  | C3014 | TransportEquipment | 31B | CIS/Consignment/HouseConsignment/TransportEquipment | transportEquipment |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Container size and type | an..10 | M |  |  | Equipment size and type identification | 152 | CIS/Consignment/HouseConsignment/TransportEquipment/characteristic | containerSizeAndType | CL710 | CL710 |
| IE3F11 | 4 | Attribute | 1..1 |                 Container packed status | an..3 | M |  |  | Transport equipment loaded status | 154 | CIS/Consignment/HouseConsignment/TransportEquipment/fullness | containerPackedStatus | CL709 | CL709 |
