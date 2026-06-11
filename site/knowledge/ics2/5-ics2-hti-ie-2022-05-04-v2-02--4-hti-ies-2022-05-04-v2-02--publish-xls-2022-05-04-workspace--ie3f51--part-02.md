---
title: "IE3F51 - part 2"
source_title: "IE3F51"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F51.md"
section: "IE3F51"
chunk: 2
chunk_count: 4
generated: true
---

# IE3F51 - part 2

Source document: IE3F51
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F51.md`
Chunk: 2 of 4

| IE3F51 | 4 | Attribute | 0..1 |                 Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/Consignor/Address/countrySubDivisionName | subDivision |  |  |
| IE3F51 | 4 | Attribute | 0..1 |                 Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/Consignor/Address/line | street |  |  |
| IE3F51 | 4 | Attribute | 0..1 |                 Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/Consignor/Address/postcode | postCode |  |  |
| IE3F51 | 4 | Attribute | 0..1 |                 Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/Consignor/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F51 | 4 | Attribute | 0..1 |                 Number | an..35 | O |  |  |  |  | CIS/Consignment/Consignor/Address/streetNumber | number |  |  |
| IE3F51 | 4 | Attribute | 0..1 |                 P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/Consignor/Address/poBox | poBox |  |  |
| IE3F51 | 3 | Composition | 0..9 |             Communication |  | O |  |  | Communication | 25A | CIS/Consignment/Consignor/Communication | communication |  |  |
| IE3F51 | 4 | Attribute | 1..1 |                 Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/Consignor/Communication/identification | identifier |  |  |
| IE3F51 | 4 | Attribute | 1..1 |                 Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/Consignor/Communication/type | type | CL707 | CL707 |
| IE3F51 | 2 | Composition | 1..1 |         Transport charges |  | M |  |  | Freight | 62A | CIS/Consignment/Freight | transportCharges |  |  |
| IE3F51 | 3 | Attribute | 1..1 |             Method of payment | a1 | M |  |  | Transport charges method of payment, coded | 098 | CIS/Consignment/Freight/paymentMethod | methodOfPayment | CL116 | CL116 |
| IE3F51 | 2 | Composition | 1..1 |         Place of delivery |  | M |  |  | GoodsReceiptPlace | 66A | CIS/Consignment/GoodsReceiptPlace | placeOfDelivery |  |  |
| IE3F51 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3005 | Goods receipt place | L003 | CIS/Consignment/GoodsReceiptPlace/name | location |  |  |
| IE3F51 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Goods receipt place, coded | L004 | CIS/Consignment/GoodsReceiptPlace/identification | UNLOCODE | CL244 | CL244 |
| IE3F51 | 3 | Composition | 0..1 |             Address |  | C |  | C3005 | Address | 04A | CIS/Consignment/GoodsReceiptPlace/Address | address |  |  |
| IE3F51 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/GoodsReceiptPlace/Address/country | country | CL718 | CL718 |
| IE3F51 | 2 | Composition | 1..99999 |         Consignment (House level) |  | M |  |  | HouseConsignment | 21C | CIS/Consignment/HouseConsignment | consignmentHouseLevel |  |  |
| IE3F51 | 3 | Attribute | 1..1 |             Container indicator | n1 | M |  |  | Container transport code | 096 | CIS/Consignment/HouseConsignment/container | containerIndicator | CL708 | CL708 |
| IE3F51 | 3 | Attribute | 1..1 |             Total gross mass | n..16,6 | M |  |  | Total gross weight | 131 | CIS/Consignment/HouseConsignment/totalGrossMass | totalGrossMass |  |  |
| IE3F51 | 3 | Composition | 1..1 |             Place of acceptance |  | M |  |  | AcceptancePlace | 01A | CIS/Consignment/HouseConsignment/AcceptancePlace | placeOfAcceptance |  |  |
| IE3F51 | 4 | Attribute | 0..1 |                 Location | an..35 | C |  | C3004 | Place of acceptance | L011 | CIS/Consignment/HouseConsignment/AcceptancePlace/name | location |  |  |
| IE3F51 | 4 | Attribute | 0..1 |                 UNLOCODE | an..17 | O |  |  | Place of acceptance, coded | L099 | CIS/Consignment/HouseConsignment/AcceptancePlace/identification | UNLOCODE | CL244 | CL244 |
| IE3F51 | 4 | Composition | 0..1 |                 Address |  | C |  | C3004 | Address | 04A | CIS/Consignment/HouseConsignment/AcceptancePlace/Address | address |  |  |
| IE3F51 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/AcceptancePlace/Address/country | country | CL718 | CL718 |
| IE3F51 | 3 | Composition | 0..99 |             Supporting documents |  | O |  |  | AdditionalDocument | 02A | CIS/Consignment/HouseConsignment/AdditionalDocument | supportingDocuments |  |  |
| IE3F51 | 4 | Attribute | 1..1 |                 Reference number | an..70 | M |  |  | Additional document reference number | D005 | CIS/Consignment/HouseConsignment/AdditionalDocument/identification | referenceNumber |  |  |
| IE3F51 | 4 | Attribute | 1..1 |                 Type | an4 | M |  |  | Additional document type, coded | D006 | CIS/Consignment/HouseConsignment/AdditionalDocument/type | type | CL013 | CL013 |
| IE3F51 | 3 | Composition | 0..99 |             Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/Consignment/HouseConsignment/AdditionalInformation | additionalInformation |  |  |
| IE3F51 | 4 | Attribute | 0..1 |                 Code | an..17 | O |  |  | Additional statement code | 226 | CIS/Consignment/HouseConsignment/AdditionalInformation/statement | code | CL701 | CL701 |
| IE3F51 | 4 | Attribute | 0..1 |                 Text | an..512 | O |  |  | Additional statement text | 225 | CIS/Consignment/HouseConsignment/AdditionalInformation/statementDescription | text |  |  |
| IE3F51 | 3 | Composition | 0..9 |             Additional supply chain actor |  | O |  |  | AEOMutualRecognitionParty | 16C | CIS/Consignment/HouseConsignment/AEOMutualRecognitionParty | additionalSupplyChainActor |  |  |
| IE3F51 | 4 | Attribute | 1..1 |                 Identification number | an..17 | M |  |  | AEO mutual recognition party, coded | R143 | CIS/Consignment/HouseConsignment/AEOMutualRecognitionParty/identification | identificationNumber |  |  |
| IE3F51 | 4 | Attribute | 1..1 |                 Role | a..3 | M |  |  | Role code | R005 | CIS/Consignment/HouseConsignment/AEOMutualRecognitionParty/role | role | CL704 | CL704 |
| IE3F51 | 3 | Composition | 1..9999 |             Goods item |  | M |  |  | ConsignmentItem | 29A | CIS/Consignment/HouseConsignment/ConsignmentItem | goodsItem |  |  |
| IE3F51 | 4 | Attribute | 1..1 |                 Goods item number | n..5 | M | R3017 |  | Sequence number | 006 | CIS/Consignment/HouseConsignment/ConsignmentItem/sequence | goodsItemNumber |  |  |
| IE3F51 | 4 | Composition | 0..99 |                 Supporting documents |  | O |  |  | AdditionalDocument | 02A | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalDocument | supportingDocuments |  |  |
| IE3F51 | 5 | Attribute | 1..1 |                     Reference number | an..70 | M |  |  | Additional document reference number | D005 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalDocument/identification | referenceNumber |  |  |
| IE3F51 | 5 | Attribute | 1..1 |                     Type | an4 | M |  |  | Additional document type, coded | D006 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalDocument/type | type | CL013 | CL013 |
| IE3F51 | 4 | Composition | 0..99 |                 Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalInformation | additionalInformation |  |  |
| IE3F51 | 5 | Attribute | 0..1 |                     Code | an..17 | O |  |  | Additional statement code | 226 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalInformation/statement | code | CL701 | CL701 |
| IE3F51 | 5 | Attribute | 0..1 |                     Text | an..512 | O |  |  | Additional statement text | 225 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalInformation/statementDescription | text |  |  |
| IE3F51 | 4 | Composition | 0..9 |                 Additional supply chain actor |  | O |  |  | AEOMutualRecognitionParty | 16C | CIS/Consignment/HouseConsignment/ConsignmentItem/AEOMutualRecognitionParty | additionalSupplyChainActor |  |  |
| IE3F51 | 5 | Attribute | 1..1 |                     Identification number | an..17 | M |  |  | AEO mutual recognition party, coded | R143 | CIS/Consignment/HouseConsignment/ConsignmentItem/AEOMutualRecognitionParty/identification | identificationNumber |  |  |
| IE3F51 | 5 | Attribute | 1..1 |                     Role | a..3 | M |  |  | Role code | R005 | CIS/Consignment/HouseConsignment/ConsignmentItem/AEOMutualRecognitionParty/role | role | CL704 | CL704 |
| IE3F51 | 4 | Composition | 1..1 |                 Commodity |  | M |  |  | Commodity | 23A | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity | commodity |  |  |
| IE3F51 | 5 | Attribute | 1..1 |                     Description of goods | an..512 | M |  |  | Description of goods | 137 | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/description | descriptionOfGoods |  |  |
| IE3F51 | 5 | Attribute | 0..1 |                     CUS code | an9 | O | R3020 |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/cusCode | cusCode | CL719 | CL719 |
| IE3F51 | 5 | Composition | 0..1 |                     Commodity code |  | C |  | C3025 |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/Commodity code | commodityCode | CL706 | CL706 |
| IE3F51 | 6 | Attribute | 1..1 |                         Harmonized System sub-heading code | an6 | M |  |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/Commodity code/Harmonized System sub-heading code | harmonizedSystemSubHeadingCode |  |  |
| IE3F51 | 6 | Attribute | 0..1 |                         Combined nomenclature code | an2 | O |  |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/Commodity code/Combined nomenclature code | combinedNomenclatureCode |  |  |
| IE3F51 | 5 | Composition | 0..99 |                     Dangerous goods |  | O |  |  | DangerousGoods | 12C | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/DangerousGoods | dangerousGoods |  |  |
| IE3F51 | 6 | Attribute | 1..1 |                         UN Number | an4 | M |  |  | United Nations Dangerous Goods Identifier (UNDG) | 506 | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/DangerousGoods/UNDG | UNNumber | CL101 | CL101 |
| IE3F51 | 4 | Composition | 1..1 |                 Weight |  | M |  |  | GoodsMeasure | 65A | CIS/Consignment/HouseConsignment/ConsignmentItem/GoodsMeasure | weight |  |  |
| IE3F51 | 5 | Attribute | 1..1 |                     Gross mass | n..16,6 | M |  |  | Gross weight item level | 126 | CIS/Consignment/HouseConsignment/ConsignmentItem/GoodsMeasure/grossMass | grossMass |  |  |
| IE3F51 | 4 | Composition | 1..99 |                 Packaging |  | M |  |  | Packaging | 93A | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging | packaging |  |  |
| IE3F51 | 5 | Attribute | 0..1 |                     Shipping marks | an..512 | C |  | C3036 | Shipping marks | 142 | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging/marksNumbers | shippingMarks |  |  |
| IE3F51 | 5 | Attribute | 0..1 |                     Number of packages | n..8 | C |  | C3036 | Number of packages | 144 | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging/quantity | numberOfPackages |  |  |
| IE3F51 | 5 | Attribute | 1..1 |                     Type of packages | an2 | M |  |  | Type of packages identification, coded | 141 | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging/type | typeOfPackages | CL017 | CL017 |
| IE3F51 | 4 | Composition | 0..99 |                 Passive border transport means |  | O | R3008 |  | PassiveTransportMeans | 19C | CIS/Consignment/HouseConsignment/ConsignmentItem/PassiveTransportMeans | passiveBorderTransportMeans |  |  |
| IE3F51 | 5 | Attribute | 1..1 |                     Identification number | an..35 | M |  |  | Identification of passive transport means, coded | T026 | CIS/Consignment/HouseConsignment/ConsignmentItem/PassiveTransportMeans/identification | identificationNumber |  |  |
| IE3F51 | 5 | Attribute | 1..1 |                     Type of identification | n2 | M |  |  | Type of identification of passive transport means, coded | T027 | CIS/Consignment/HouseConsignment/ConsignmentItem/PassiveTransportMeans/identificationType | typeOfIdentification | CL750 | CL750 |
| IE3F51 | 5 | Attribute | 1..1 |                     Type of means of transport | an..4 | M |  |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/PassiveTransportMeans/type | typeOfMeansOfTransport | CL751 | CL751 |
| IE3F51 | 4 | Composition | 0..9999 |                 Transport equipment |  | C |  | C3014 | TransportEquipment | 31B | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment | transportEquipment |  |  |
| IE3F51 | 5 | Attribute | 1..1 |                     Container size and type | an..10 | M |  |  | Equipment size and type identification | 152 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/characteristic | containerSizeAndType | CL710 | CL710 |
| IE3F51 | 5 | Attribute | 1..1 |                     Container packed status | an..3 | M |  |  | Transport equipment loaded status | 154 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/fullness | containerPackedStatus | CL709 | CL709 |
| IE3F51 | 5 | Attribute | 1..1 |                     Container supplier type | an..3 | M |  |  | Equipment supplier type, coded | 151 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/supplierPartyType | containerSupplierType | CL711 | CL711 |
| IE3F51 | 5 | Attribute | 1..1 |                     Container identification number | an..17 | M | R3021 |  | Equipment identification number | 159 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/identification | containerIdentificationNumber |  |  |
