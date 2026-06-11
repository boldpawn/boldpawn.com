---
title: "IE3A15 - part 2"
source_title: "IE3A15"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A15.md"
section: "IE3A15"
chunk: 2
chunk_count: 3
generated: true
---

# IE3A15 - part 2

Source document: IE3A15
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A15.md`
Chunk: 2 of 3

| IE3A15 | 4 | Composition | 0..99 |                 Additional supply chain actor |  | O |  |  | AEOMutualRecognitionParty | 16C | CIS/Consignment/HouseConsignment/ConsignmentItem/AEOMutualRecognitionParty | additionalSupplyChainActor |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Identification number | an..17 | M |  |  | AEO mutual recognition party, coded | R143 | CIS/Consignment/HouseConsignment/ConsignmentItem/AEOMutualRecognitionParty/identification | identificationNumber |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Role | a..3 | M |  |  | Role code | R005 | CIS/Consignment/HouseConsignment/ConsignmentItem/AEOMutualRecognitionParty/role | role | CL704 | CL704 |
| IE3A15 | 4 | Composition | 1..1 |                 Commodity |  | M |  |  | Commodity | 23A | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity | commodity |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Description of goods | an..512 | M |  |  | Description of goods | 137 | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/description | descriptionOfGoods |  |  |
| IE3A15 | 5 | Attribute | 0..1 |                     CUS code | an9 | O | R3020 |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/cusCode | cusCode | CL719 | CL719 |
| IE3A15 | 5 | Composition | 0..1 |                     Commodity code |  | C |  | C3025 |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/Commodity code | commodityCode | CL706 | CL706 |
| IE3A15 | 6 | Attribute | 1..1 |                         Harmonized System sub-heading code | an6 | M |  |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/Commodity code/Harmonized System sub-heading code | harmonizedSystemSubHeadingCode |  |  |
| IE3A15 | 6 | Attribute | 0..1 |                         Combined nomenclature code | an2 | O |  |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/Commodity code/Combined nomenclature code | combinedNomenclatureCode |  |  |
| IE3A15 | 5 | Composition | 0..99 |                     Dangerous goods |  | O |  |  | DangerousGoods | 12C | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/DangerousGoods | dangerousGoods |  |  |
| IE3A15 | 6 | Attribute | 1..1 |                         UN Number | an4 | M |  |  | United Nations Dangerous Goods Identifier (UNDG) | 506 | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/DangerousGoods/UNDG | UNNumber | CL101 | CL101 |
| IE3A15 | 4 | Composition | 1..1 |                 Weight |  | M |  |  | GoodsMeasure | 65A | CIS/Consignment/HouseConsignment/ConsignmentItem/GoodsMeasure | weight |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Gross mass | n..16,6 | M |  |  | Gross weight item level | 126 | CIS/Consignment/HouseConsignment/ConsignmentItem/GoodsMeasure/grossMass | grossMass |  |  |
| IE3A15 | 4 | Composition | 1..99 |                 Packaging |  | M |  |  | Packaging | 93A | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging | packaging |  |  |
| IE3A15 | 5 | Attribute | 0..1 |                     Shipping marks | an..512 | C |  | C3036 | Shipping marks | 142 | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging/marksNumbers | shippingMarks |  |  |
| IE3A15 | 5 | Attribute | 0..1 |                     Number of packages | n..8 | C |  | C3036 | Number of packages | 144 | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging/quantity | numberOfPackages |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Type of packages | an2 | M |  |  | Type of packages identification, coded | 141 | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging/type | typeOfPackages | CL017 | CL017 |
| IE3A15 | 4 | Composition | 0..9999 |                 Transport equipment |  | C |  | C3014 | TransportEquipment | 31B | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment | transportEquipment |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Container size and type | an..10 | M |  |  | Equipment size and type identification | 152 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/characteristic | containerSizeAndType | CL710 | CL710 |
| IE3A15 | 5 | Attribute | 1..1 |                     Container packed status | an..3 | M |  |  | Transport equipment loaded status | 154 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/fullness | containerPackedStatus | CL709 | CL709 |
| IE3A15 | 5 | Attribute | 1..1 |                     Container supplier type | an..3 | M |  |  | Equipment supplier type, coded | 151 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/supplierPartyType | containerSupplierType | CL711 | CL711 |
| IE3A15 | 5 | Attribute | 1..1 |                     Container identification number | an..17 | M | R3021 |  | Equipment identification number | 159 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3A15 | 5 | Attribute | 0..1 |                     Number of seals | n..4 | O |  |  | Number of seals | 227 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/sealsAffixed | numberOfSeals |  |  |
| IE3A15 | 5 | Composition | 0..99 |                     Seal |  | C |  | C3018 | Seal | 44B | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/Seal | seal |  |  |
| IE3A15 | 6 | Attribute | 1..1 |                         Identifier | an..20 | M |  |  | Seal number | 165 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/Seal/identification | identifier |  |  |
| IE3A15 | 3 | Composition | 1..1 |             Consignor |  | M |  |  | Consignor | 30A | CIS/Consignment/HouseConsignment/Consignor | consignor |  |  |
| IE3A15 | 4 | Attribute | 1..1 |                 Name | an..70 | M |  |  | Consignor - name | R020 | CIS/Consignment/HouseConsignment/Consignor/name | name |  |  |
| IE3A15 | 4 | Attribute | 0..1 |                 Identification number | an..17 | O |  |  | Consignor, coded | R021 | CIS/Consignment/HouseConsignment/Consignor/identification | identificationNumber |  |  |
| IE3A15 | 4 | Attribute | 1..1 |                 Type of person | n1 | M |  |  | Consignor, type of code |  | CIS/Consignment/HouseConsignment/Consignor/identificationType | typeOfPerson | CL729 | CL729 |
| IE3A15 | 4 | Composition | 1..1 |                 Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/HouseConsignment/Consignor/Address | address |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     City | an..35 | M |  |  | City name | 241 | CIS/Consignment/HouseConsignment/Consignor/Address/cityName | city |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/Consignor/Address/country | country | CL718 | CL718 |
| IE3A15 | 5 | Attribute | 0..1 |                     Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/HouseConsignment/Consignor/Address/countrySubDivisionName | subDivision |  |  |
| IE3A15 | 5 | Attribute | 0..1 |                     Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/HouseConsignment/Consignor/Address/line | street |  |  |
| IE3A15 | 5 | Attribute | 0..1 |                     Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/HouseConsignment/Consignor/Address/postcode | postCode |  |  |
| IE3A15 | 5 | Attribute | 0..1 |                     Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignor/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A15 | 5 | Attribute | 0..1 |                     Number | an..35 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignor/Address/streetNumber | number |  |  |
| IE3A15 | 5 | Attribute | 0..1 |                     P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignor/Address/poBox | poBox |  |  |
| IE3A15 | 4 | Composition | 0..9 |                 Communication |  | O |  |  | Communication | 25A | CIS/Consignment/HouseConsignment/Consignor/Communication | communication |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/HouseConsignment/Consignor/Communication/identification | identifier |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/HouseConsignment/Consignor/Communication/type | type | CL707 | CL707 |
| IE3A15 | 3 | Composition | 1..1 |             Transport charges |  | M |  |  | Freight | 62A | CIS/Consignment/HouseConsignment/Freight | transportCharges |  |  |
| IE3A15 | 4 | Attribute | 1..1 |                 Method of payment | a1 | M |  |  | Transport charges method of payment, coded | 098 | CIS/Consignment/HouseConsignment/Freight/paymentMethod | methodOfPayment | CL116 | CL116 |
| IE3A15 | 3 | Composition | 1..1 |             Place of delivery |  | M |  |  | GoodsReceiptPlace | 66A | CIS/Consignment/HouseConsignment/GoodsReceiptPlace | placeOfDelivery |  |  |
| IE3A15 | 4 | Attribute | 0..1 |                 Location | an..35 | C |  | C3006 | Goods receipt place | L003 | CIS/Consignment/HouseConsignment/GoodsReceiptPlace/name | location |  |  |
| IE3A15 | 4 | Attribute | 0..1 |                 UNLOCODE | an..17 | O |  |  | Goods receipt place, coded | L004 | CIS/Consignment/HouseConsignment/GoodsReceiptPlace/identification | UNLOCODE | CL244 | CL244 |
| IE3A15 | 4 | Composition | 0..1 |                 Address |  | C |  | C3006 | Address | 04A | CIS/Consignment/HouseConsignment/GoodsReceiptPlace/Address | address |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/GoodsReceiptPlace/Address/country | country | CL718 | CL718 |
| IE3A15 | 3 | Composition | 0..1 |             Goods shipment |  | C |  | C3020 | GoodsShipment | 67A | CIS/Consignment/HouseConsignment/GoodsShipment | goodsShipment |  |  |
| IE3A15 | 4 | Composition | 1..1 |                 Buyer |  | M |  |  | Buyer | 16A | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer | buyer |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Name | an..70 | M |  |  | Buyer - name | R009 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/name | name |  |  |
| IE3A15 | 5 | Attribute | 0..1 |                     Identification number | an..17 | O |  |  | Buyer, coded | R010 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/identification | identificationNumber |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Type of person | n1 | M |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/identificationType | typeOfPerson | CL729 | CL729 |
| IE3A15 | 5 | Composition | 1..1 |                     Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address | address |  |  |
| IE3A15 | 6 | Attribute | 1..1 |                         City | an..35 | M |  |  | City name | 241 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/cityName | city |  |  |
| IE3A15 | 6 | Attribute | 1..1 |                         Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/country | country | CL718 | CL718 |
| IE3A15 | 6 | Attribute | 0..1 |                         Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/countrySubDivisionName | subDivision |  |  |
| IE3A15 | 6 | Attribute | 0..1 |                         Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/line | street |  |  |
| IE3A15 | 6 | Attribute | 0..1 |                         Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/postcode | postCode |  |  |
| IE3A15 | 6 | Attribute | 0..1 |                         Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A15 | 6 | Attribute | 0..1 |                         Number | an..35 | O |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/streetNumber | number |  |  |
| IE3A15 | 6 | Attribute | 0..1 |                         P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Address/poBox | poBox |  |  |
| IE3A15 | 5 | Composition | 0..9 |                     Communication |  | O |  |  | Communication | 25A | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Communication | communication |  |  |
| IE3A15 | 6 | Attribute | 1..1 |                         Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Communication/identification | identifier |  |  |
| IE3A15 | 6 | Attribute | 1..1 |                         Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/HouseConsignment/GoodsShipment/Buyer/Communication/type | type | CL707 | CL707 |
| IE3A15 | 4 | Composition | 1..1 |                 Seller |  | M |  |  | Seller | 09B | CIS/Consignment/HouseConsignment/GoodsShipment/Seller | seller |  |  |
| IE3A15 | 5 | Attribute | 1..1 |                     Name | an..70 | M |  |  | Seller - name | R050 | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/name | name |  |  |
| IE3A15 | 5 | Attribute | 0..1 |                     Identification number | an..17 | O |  |  | Seller, coded | R051 | CIS/Consignment/HouseConsignment/GoodsShipment/Seller/identification | identificationNumber |  |  |
