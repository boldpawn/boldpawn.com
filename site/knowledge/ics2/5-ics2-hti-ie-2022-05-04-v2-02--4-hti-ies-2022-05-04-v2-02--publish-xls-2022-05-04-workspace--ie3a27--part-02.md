---
title: "IE3A27 - part 2"
source_title: "IE3A27"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A27.md"
section: "IE3A27"
chunk: 2
chunk_count: 4
generated: true
---

# IE3A27 - part 2

Source document: IE3A27
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A27.md`
Chunk: 2 of 4

| IE3A27 | 4 | Attribute | 1..1 |                 Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/Consignee/Communication/type | type | CL707 | CL707 |
| IE3A27 | 2 | Composition | 1..9999999 |         Goods item |  | M |  |  | ConsignmentItem | 29A | CIS/Consignment/ConsignmentItem | goodsItem |  |  |
| IE3A27 | 3 | Attribute | 1..1 |             Goods item number | n..5 | M | R3016<br>R3024 |  | Sequence number | 006 | CIS/Consignment/ConsignmentItem/sequence | goodsItemNumber |  |  |
| IE3A27 | 3 | Composition | 0..99 |             Supporting documents |  | O |  |  | AdditionalDocument | 02A | CIS/Consignment/ConsignmentItem/AdditionalDocument | supportingDocuments |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Reference number | an..70 | M |  |  | Additional document reference number | D005 | CIS/Consignment/ConsignmentItem/AdditionalDocument/identification | referenceNumber |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Type | an4 | M |  |  | Additional document type, coded | D006 | CIS/Consignment/ConsignmentItem/AdditionalDocument/type | type | CL013 | CL013 |
| IE3A27 | 3 | Composition | 0..99 |             Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/Consignment/ConsignmentItem/AdditionalInformation | additionalInformation |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 Code | an..17 | O |  |  | Additional statement code | 226 | CIS/Consignment/ConsignmentItem/AdditionalInformation/statement | code | CL701 | CL701 |
| IE3A27 | 4 | Attribute | 0..1 |                 Text | an..512 | O |  |  | Additional statement text | 225 | CIS/Consignment/ConsignmentItem/AdditionalInformation/statementDescription | text |  |  |
| IE3A27 | 3 | Composition | 0..99 |             Additional supply chain actor |  | O |  |  | AEOMutualRecognitionParty | 16C | CIS/Consignment/ConsignmentItem/AEOMutualRecognitionParty | additionalSupplyChainActor |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Identification number | an..17 | M |  |  | AEO mutual recognition party, coded | R143 | CIS/Consignment/ConsignmentItem/AEOMutualRecognitionParty/identification | identificationNumber |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Role | a..3 | M |  |  | Role code | R005 | CIS/Consignment/ConsignmentItem/AEOMutualRecognitionParty/role | role | CL704 | CL704 |
| IE3A27 | 3 | Composition | 1..1 |             Commodity |  | M |  |  | Commodity | 23A | CIS/Consignment/ConsignmentItem/Commodity | commodity |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Description of goods | an..512 | M |  |  | Description of goods | 137 | CIS/Consignment/ConsignmentItem/Commodity/description | descriptionOfGoods |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 CUS code | an9 | O | R3020 |  |  |  | CIS/Consignment/ConsignmentItem/Commodity/cusCode | cusCode | CL719 | CL719 |
| IE3A27 | 4 | Composition | 0..1 |                 Commodity code |  | C |  | C3048 |  |  | CIS/Consignment/ConsignmentItem/Commodity/Commodity code | commodityCode | CL706 | CL706 |
| IE3A27 | 5 | Attribute | 1..1 |                     Harmonized System sub-heading code | an6 | M |  |  |  |  | CIS/Consignment/ConsignmentItem/Commodity/Commodity code/Harmonized System sub-heading code | harmonizedSystemSubHeadingCode |  |  |
| IE3A27 | 5 | Attribute | 0..1 |                     Combined nomenclature code | an2 | O |  |  |  |  | CIS/Consignment/ConsignmentItem/Commodity/Commodity code/Combined nomenclature code | combinedNomenclatureCode |  |  |
| IE3A27 | 4 | Composition | 0..99 |                 Dangerous goods |  | O |  |  | DangerousGoods | 12C | CIS/Consignment/ConsignmentItem/Commodity/DangerousGoods | dangerousGoods |  |  |
| IE3A27 | 5 | Attribute | 1..1 |                     UN Number | an4 | M |  |  | United Nations Dangerous Goods Identifier (UNDG) | 506 | CIS/Consignment/ConsignmentItem/Commodity/DangerousGoods/UNDG | UNNumber | CL101 | CL101 |
| IE3A27 | 3 | Composition | 1..1 |             Weight |  | M |  |  | GoodsMeasure | 65A | CIS/Consignment/ConsignmentItem/GoodsMeasure | weight |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Gross mass | n..16,6 | M |  |  | Gross weight item level | 126 | CIS/Consignment/ConsignmentItem/GoodsMeasure/grossMass | grossMass |  |  |
| IE3A27 | 3 | Composition | 1..99 |             Packaging |  | M |  |  | Packaging | 93A | CIS/Consignment/ConsignmentItem/Packaging | packaging |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 Shipping marks | an..512 | C |  | C3035 | Shipping marks | 142 | CIS/Consignment/ConsignmentItem/Packaging/marksNumbers | shippingMarks |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 Number of packages | n..8 | C |  | C3035 | Number of packages | 144 | CIS/Consignment/ConsignmentItem/Packaging/quantity | numberOfPackages |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Type of packages | an2 | M |  |  | Type of packages identification, coded | 141 | CIS/Consignment/ConsignmentItem/Packaging/type | typeOfPackages | CL017 | CL017 |
| IE3A27 | 3 | Composition | 0..9999 |             Transport equipment |  | C |  | C3013 | TransportEquipment | 31B | CIS/Consignment/ConsignmentItem/TransportEquipment | transportEquipment |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Container identification number | an..17 | M |  |  | Equipment identification number | 159 | CIS/Consignment/ConsignmentItem/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3A27 | 2 | Composition | 1..1 |         Consignor |  | M |  |  | Consignor | 30A | CIS/Consignment/Consignor | consignor |  |  |
| IE3A27 | 3 | Attribute | 1..1 |             Name | an..70 | M |  |  | Consignor - name | R020 | CIS/Consignment/Consignor/name | name |  |  |
| IE3A27 | 3 | Attribute | 0..1 |             Identification number | an..17 | O |  |  | Consignor, coded | R021 | CIS/Consignment/Consignor/identification | identificationNumber |  |  |
| IE3A27 | 3 | Attribute | 1..1 |             Type of person | n1 | M |  |  | Consignor, type of code |  | CIS/Consignment/Consignor/identificationType | typeOfPerson | CL729 | CL729 |
| IE3A27 | 3 | Composition | 1..1 |             Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/Consignor/Address | address |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 City | an..35 | M |  |  | City name | 241 | CIS/Consignment/Consignor/Address/cityName | city |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/Consignor/Address/country | country | CL718 | CL718 |
| IE3A27 | 4 | Attribute | 0..1 |                 Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/Consignor/Address/countrySubDivisionName | subDivision |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/Consignor/Address/line | street |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/Consignor/Address/postcode | postCode |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/Consignor/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 Number | an..35 | O |  |  |  |  | CIS/Consignment/Consignor/Address/streetNumber | number |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/Consignor/Address/poBox | poBox |  |  |
| IE3A27 | 3 | Composition | 0..9 |             Communication |  | O |  |  | Communication | 25A | CIS/Consignment/Consignor/Communication | communication |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/Consignor/Communication/identification | identifier |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/Consignor/Communication/type | type | CL707 | CL707 |
| IE3A27 | 2 | Composition | 1..1 |         Transport charges |  | M |  |  | Freight | 62A | CIS/Consignment/Freight | transportCharges |  |  |
| IE3A27 | 3 | Attribute | 1..1 |             Method of payment | a1 | M |  |  | Transport charges method of payment, coded | 098 | CIS/Consignment/Freight/paymentMethod | methodOfPayment | CL116 | CL116 |
| IE3A27 | 2 | Composition | 1..1 |         Place of delivery |  | M |  |  | GoodsReceiptPlace | 66A | CIS/Consignment/GoodsReceiptPlace | placeOfDelivery |  |  |
| IE3A27 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3005 | Goods receipt place | L003 | CIS/Consignment/GoodsReceiptPlace/name | location |  |  |
| IE3A27 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Goods receipt place, coded | L004 | CIS/Consignment/GoodsReceiptPlace/identification | UNLOCODE | CL244 | CL244 |
| IE3A27 | 3 | Composition | 0..1 |             Address |  | C |  | C3005 | Address | 04A | CIS/Consignment/GoodsReceiptPlace/Address | address |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/GoodsReceiptPlace/Address/country | country | CL718 | CL718 |
| IE3A27 | 2 | Composition | 1..99999 |         Consignment (House level) |  | M |  |  | HouseConsignment | 21C | CIS/Consignment/HouseConsignment | consignmentHouseLevel |  |  |
| IE3A27 | 3 | Attribute | 1..1 |             Container indicator | n1 | M |  |  | Container transport code | 096 | CIS/Consignment/HouseConsignment/container | containerIndicator | CL708 | CL708 |
| IE3A27 | 3 | Attribute | 1..1 |             Total gross mass | n..16,6 | M |  |  | Total gross weight | 131 | CIS/Consignment/HouseConsignment/totalGrossMass | totalGrossMass |  |  |
| IE3A27 | 3 | Composition | 1..1 |             Place of acceptance |  | M |  |  | AcceptancePlace | 01A | CIS/Consignment/HouseConsignment/AcceptancePlace | placeOfAcceptance |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 Location | an..35 | C |  | C3004 | Place of acceptance | L011 | CIS/Consignment/HouseConsignment/AcceptancePlace/name | location |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 UNLOCODE | an..17 | O |  |  | Place of acceptance, coded | L099 | CIS/Consignment/HouseConsignment/AcceptancePlace/identification | UNLOCODE | CL244 | CL244 |
| IE3A27 | 4 | Composition | 0..1 |                 Address |  | C |  | C3004 | Address | 04A | CIS/Consignment/HouseConsignment/AcceptancePlace/Address | address |  |  |
| IE3A27 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/AcceptancePlace/Address/country | country | CL718 | CL718 |
| IE3A27 | 3 | Composition | 0..99 |             Supporting documents |  | O |  |  | AdditionalDocument | 02A | CIS/Consignment/HouseConsignment/AdditionalDocument | supportingDocuments |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Reference number | an..70 | M |  |  | Additional document reference number | D005 | CIS/Consignment/HouseConsignment/AdditionalDocument/identification | referenceNumber |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Type | an4 | M |  |  | Additional document type, coded | D006 | CIS/Consignment/HouseConsignment/AdditionalDocument/type | type | CL013 | CL013 |
| IE3A27 | 3 | Composition | 0..99 |             Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/Consignment/HouseConsignment/AdditionalInformation | additionalInformation |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 Code | an..17 | O |  |  | Additional statement code | 226 | CIS/Consignment/HouseConsignment/AdditionalInformation/statement | code | CL701 | CL701 |
| IE3A27 | 4 | Attribute | 0..1 |                 Text | an..512 | O |  |  | Additional statement text | 225 | CIS/Consignment/HouseConsignment/AdditionalInformation/statementDescription | text |  |  |
| IE3A27 | 3 | Composition | 0..99 |             Additional supply chain actor |  | O |  |  | AEOMutualRecognitionParty | 16C | CIS/Consignment/HouseConsignment/AEOMutualRecognitionParty | additionalSupplyChainActor |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Identification number | an..17 | M |  |  | AEO mutual recognition party, coded | R143 | CIS/Consignment/HouseConsignment/AEOMutualRecognitionParty/identification | identificationNumber |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Role | a..3 | M |  |  | Role code | R005 | CIS/Consignment/HouseConsignment/AEOMutualRecognitionParty/role | role | CL704 | CL704 |
| IE3A27 | 3 | Composition | 1..1 |             Consignee |  | M |  |  | Consignee | 27A | CIS/Consignment/HouseConsignment/Consignee | consignee |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Name | an..70 | M |  |  | Consignee name | R014 | CIS/Consignment/HouseConsignment/Consignee/name | name |  |  |
| IE3A27 | 4 | Attribute | 0..1 |                 Identification number | an..17 | O |  |  | Consignee, coded | R015 | CIS/Consignment/HouseConsignment/Consignee/identification | identificationNumber |  |  |
| IE3A27 | 4 | Attribute | 1..1 |                 Type of person | n1 | M |  |  |  |  | CIS/Consignment/HouseConsignment/Consignee/identificationType | typeOfPerson | CL729 | CL729 |
| IE3A27 | 4 | Composition | 1..1 |                 Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/HouseConsignment/Consignee/Address | address |  |  |
| IE3A27 | 5 | Attribute | 1..1 |                     City | an..35 | M |  |  | City name | 241 | CIS/Consignment/HouseConsignment/Consignee/Address/cityName | city |  |  |
