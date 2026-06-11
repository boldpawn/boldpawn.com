---
title: "IE3F11 - part 3"
source_title: "IE3F11"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F11.md"
section: "IE3F11"
chunk: 3
chunk_count: 5
generated: true
---

# IE3F11 - part 3

Source document: IE3F11
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F11.md`
Chunk: 3 of 5

| IE3F11 | 4 | Attribute | 1..1 |                 Identification number | an..17 | M |  |  | AEO mutual recognition party, coded | R143 | CIS/Consignment/HouseConsignment/AEOMutualRecognitionParty/identification | identificationNumber |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Role | a..3 | M |  |  | Role code | R005 | CIS/Consignment/HouseConsignment/AEOMutualRecognitionParty/role | role | CL704 | CL704 |
| IE3F11 | 3 | Composition | 0..1 |             Consignee |  | C |  | C3022 | Consignee | 27A | CIS/Consignment/HouseConsignment/Consignee | consignee |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Name | an..70 | M |  |  | Consignee name | R014 | CIS/Consignment/HouseConsignment/Consignee/name | name |  |  |
| IE3F11 | 4 | Attribute | 0..1 |                 Identification number | an..17 | O |  |  | Consignee, coded | R015 | CIS/Consignment/HouseConsignment/Consignee/identification | identificationNumber |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Type of person | n1 | M |  |  |  |  | CIS/Consignment/HouseConsignment/Consignee/identificationType | typeOfPerson | CL729 | CL729 |
| IE3F11 | 4 | Composition | 1..1 |                 Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/HouseConsignment/Consignee/Address | address |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     City | an..35 | M |  |  | City name | 241 | CIS/Consignment/HouseConsignment/Consignee/Address/cityName | city |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/Consignee/Address/country | country | CL718 | CL718 |
| IE3F11 | 5 | Attribute | 0..1 |                     Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/HouseConsignment/Consignee/Address/countrySubDivisionName | subDivision |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/HouseConsignment/Consignee/Address/line | street |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/HouseConsignment/Consignee/Address/postcode | postCode |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignee/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Number | an..35 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignee/Address/streetNumber | number |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignee/Address/poBox | poBox |  |  |
| IE3F11 | 4 | Composition | 0..9 |                 Communication |  | O |  |  | Communication | 25A | CIS/Consignment/HouseConsignment/Consignee/Communication | communication |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/HouseConsignment/Consignee/Communication/identification | identifier |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/HouseConsignment/Consignee/Communication/type | type | CL707 | CL707 |
| IE3F11 | 3 | Composition | 1..9999 |             Goods item |  | M |  |  | ConsignmentItem | 29A | CIS/Consignment/HouseConsignment/ConsignmentItem | goodsItem |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Goods item number | n..5 | M | R3017 |  | Sequence number | 006 | CIS/Consignment/HouseConsignment/ConsignmentItem/sequence | goodsItemNumber |  |  |
| IE3F11 | 4 | Composition | 0..99 |                 Supporting documents |  | O |  |  | AdditionalDocument | 02A | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalDocument | supportingDocuments |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Reference number | an..70 | M |  |  | Additional document reference number | D005 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalDocument/identification | referenceNumber |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Type | an4 | M |  |  | Additional document type, coded | D006 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalDocument/type | type | CL013 | CL013 |
| IE3F11 | 4 | Composition | 0..99 |                 Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalInformation | additionalInformation |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Code | an..17 | O |  |  | Additional statement code | 226 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalInformation/statement | code | CL701 | CL701 |
| IE3F11 | 5 | Attribute | 0..1 |                     Text | an..512 | O |  |  | Additional statement text | 225 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalInformation/statementDescription | text |  |  |
| IE3F11 | 4 | Composition | 0..99 |                 Additional supply chain actor |  | O |  |  | AEOMutualRecognitionParty | 16C | CIS/Consignment/HouseConsignment/ConsignmentItem/AEOMutualRecognitionParty | additionalSupplyChainActor |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Identification number | an..17 | M |  |  | AEO mutual recognition party, coded | R143 | CIS/Consignment/HouseConsignment/ConsignmentItem/AEOMutualRecognitionParty/identification | identificationNumber |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Role | a..3 | M |  |  | Role code | R005 | CIS/Consignment/HouseConsignment/ConsignmentItem/AEOMutualRecognitionParty/role | role | CL704 | CL704 |
| IE3F11 | 4 | Composition | 1..1 |                 Commodity |  | M |  |  | Commodity | 23A | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity | commodity |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Description of goods | an..512 | M |  |  | Description of goods | 137 | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/description | descriptionOfGoods |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     CUS code | an9 | O | R3020 |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/cusCode | cusCode | CL719 | CL719 |
| IE3F11 | 5 | Composition | 0..1 |                     Commodity code |  | C |  | C3025 |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/Commodity code | commodityCode | CL706 | CL706 |
| IE3F11 | 6 | Attribute | 1..1 |                         Harmonized System sub-heading code | an6 | M |  |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/Commodity code/Harmonized System sub-heading code | harmonizedSystemSubHeadingCode |  |  |
| IE3F11 | 6 | Attribute | 0..1 |                         Combined nomenclature code | an2 | O |  |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/Commodity code/Combined nomenclature code | combinedNomenclatureCode |  |  |
| IE3F11 | 5 | Composition | 0..99 |                     Dangerous goods |  | O |  |  | DangerousGoods | 12C | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/DangerousGoods | dangerousGoods |  |  |
| IE3F11 | 6 | Attribute | 1..1 |                         UN Number | an4 | M |  |  | United Nations Dangerous Goods Identifier (UNDG) | 506 | CIS/Consignment/HouseConsignment/ConsignmentItem/Commodity/DangerousGoods/UNDG | UNNumber | CL101 | CL101 |
| IE3F11 | 4 | Composition | 1..1 |                 Weight |  | M |  |  | GoodsMeasure | 65A | CIS/Consignment/HouseConsignment/ConsignmentItem/GoodsMeasure | weight |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Gross mass | n..16,6 | M |  |  | Gross weight item level | 126 | CIS/Consignment/HouseConsignment/ConsignmentItem/GoodsMeasure/grossMass | grossMass |  |  |
| IE3F11 | 4 | Composition | 1..99 |                 Packaging |  | M |  |  | Packaging | 93A | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging | packaging |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Shipping marks | an..512 | C |  | C3036 | Shipping marks | 142 | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging/marksNumbers | shippingMarks |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Number of packages | n..8 | C |  | C3036 | Number of packages | 144 | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging/quantity | numberOfPackages |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Type of packages | an2 | M |  |  | Type of packages identification, coded | 141 | CIS/Consignment/HouseConsignment/ConsignmentItem/Packaging/type | typeOfPackages | CL017 | CL017 |
| IE3F11 | 4 | Composition | 0..9999 |                 Transport equipment |  | C |  | C3014 | TransportEquipment | 31B | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment | transportEquipment |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Container size and type | an..10 | M |  |  | Equipment size and type identification | 152 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/characteristic | containerSizeAndType | CL710 | CL710 |
| IE3F11 | 5 | Attribute | 1..1 |                     Container packed status | an..3 | M |  |  | Transport equipment loaded status | 154 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/fullness | containerPackedStatus | CL709 | CL709 |
| IE3F11 | 5 | Attribute | 1..1 |                     Container supplier type | an..3 | M |  |  | Equipment supplier type, coded | 151 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/supplierPartyType | containerSupplierType | CL711 | CL711 |
| IE3F11 | 5 | Attribute | 1..1 |                     Container identification number | an..17 | M | R3021 |  | Equipment identification number | 159 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Number of seals | n..4 | O |  |  | Number of seals | 227 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/sealsAffixed | numberOfSeals |  |  |
| IE3F11 | 5 | Composition | 0..99 |                     Seal |  | C |  | C3018 | Seal | 44B | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/Seal | seal |  |  |
| IE3F11 | 6 | Attribute | 1..1 |                         Identifier | an..20 | M |  |  | Seal number | 165 | CIS/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/Seal/identification | identifier |  |  |
| IE3F11 | 3 | Composition | 1..1 |             Consignor |  | M |  |  | Consignor | 30A | CIS/Consignment/HouseConsignment/Consignor | consignor |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Name | an..70 | M |  |  | Consignor - name | R020 | CIS/Consignment/HouseConsignment/Consignor/name | name |  |  |
| IE3F11 | 4 | Attribute | 0..1 |                 Identification number | an..17 | O |  |  | Consignor, coded | R021 | CIS/Consignment/HouseConsignment/Consignor/identification | identificationNumber |  |  |
| IE3F11 | 4 | Attribute | 1..1 |                 Type of person | n1 | M |  |  | Consignor, type of code |  | CIS/Consignment/HouseConsignment/Consignor/identificationType | typeOfPerson | CL729 | CL729 |
| IE3F11 | 4 | Composition | 1..1 |                 Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/HouseConsignment/Consignor/Address | address |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     City | an..35 | M |  |  | City name | 241 | CIS/Consignment/HouseConsignment/Consignor/Address/cityName | city |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/Consignor/Address/country | country | CL718 | CL718 |
| IE3F11 | 5 | Attribute | 0..1 |                     Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/HouseConsignment/Consignor/Address/countrySubDivisionName | subDivision |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/HouseConsignment/Consignor/Address/line | street |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/HouseConsignment/Consignor/Address/postcode | postCode |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignor/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     Number | an..35 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignor/Address/streetNumber | number |  |  |
| IE3F11 | 5 | Attribute | 0..1 |                     P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignor/Address/poBox | poBox |  |  |
| IE3F11 | 4 | Composition | 0..9 |                 Communication |  | O |  |  | Communication | 25A | CIS/Consignment/HouseConsignment/Consignor/Communication | communication |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/HouseConsignment/Consignor/Communication/identification | identifier |  |  |
| IE3F11 | 5 | Attribute | 1..1 |                     Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/HouseConsignment/Consignor/Communication/type | type | CL707 | CL707 |
| IE3F11 | 3 | Composition | 1..1 |             Transport charges |  | M |  |  | Freight | 62A | CIS/Consignment/HouseConsignment/Freight | transportCharges |  |  |
