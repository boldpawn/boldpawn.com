---
title: "IE3A43 - part 2"
source_title: "IE3A43"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A43.md"
section: "IE3A43"
chunk: 2
chunk_count: 2
generated: true
---

# IE3A43 - part 2

Source document: IE3A43
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A43.md`
Chunk: 2 of 2

| IE3A43 | 4 | Composition | 1..1 |                 Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/HouseConsignment/Consignor/Address | address |  |  |
| IE3A43 | 5 | Attribute | 1..1 |                     City | an..35 | M |  |  | City name | 241 | CIS/Consignment/HouseConsignment/Consignor/Address/cityName | city |  |  |
| IE3A43 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/Consignor/Address/country | country | CL718 | CL718 |
| IE3A43 | 5 | Attribute | 0..1 |                     Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/HouseConsignment/Consignor/Address/countrySubDivisionName | subDivision |  |  |
| IE3A43 | 5 | Attribute | 0..1 |                     Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/HouseConsignment/Consignor/Address/line | street |  |  |
| IE3A43 | 5 | Attribute | 0..1 |                     Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/HouseConsignment/Consignor/Address/postcode | postCode |  |  |
| IE3A43 | 5 | Attribute | 0..1 |                     Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignor/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A43 | 5 | Attribute | 0..1 |                     Number | an..35 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignor/Address/streetNumber | number |  |  |
| IE3A43 | 5 | Attribute | 0..1 |                     P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignor/Address/poBox | poBox |  |  |
| IE3A43 | 4 | Composition | 0..9 |                 Communication |  | O |  |  | Communication | 25A | CIS/Consignment/HouseConsignment/Consignor/Communication | communication |  |  |
| IE3A43 | 5 | Attribute | 1..1 |                     Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/HouseConsignment/Consignor/Communication/identification | identifier |  |  |
| IE3A43 | 5 | Attribute | 1..1 |                     Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/HouseConsignment/Consignor/Communication/type | type | CL707 | CL707 |
| IE3A43 | 3 | Composition | 0..1 |             Postal charges |  | C |  | C3038 | CustomsValuation | 41A | CIS/Consignment/HouseConsignment/CustomsValuation | postalCharges |  |  |
| IE3A43 | 4 | Attribute | 1..1 |                 Postal charges |  | M |  |  | Freight costs | 117 | CIS/Consignment/HouseConsignment/CustomsValuation/freightCharge | postalCharges |  |  |
| IE3A43 | 5 | Attribute | 1..1 |                     Value | n..16,2 | M |  |  |  |  | CIS/Consignment/HouseConsignment/CustomsValuation/freightCharge/Amount. Content | value |  |  |
| IE3A43 | 5 | Attribute | 1..1 |                     Currency | a3 | M |  |  |  |  | CIS/Consignment/HouseConsignment/CustomsValuation/freightCharge/Amount Currency. Identifier | currency | CL352 | CL352 |
| IE3A43 | 3 | Composition | 1..1 |             Transport document (House level) |  | M | R3052 |  | TransportContractDocument | 30B | CIS/Consignment/HouseConsignment/TransportContractDocument | transportDocumentHouseLevel |  |  |
| IE3A43 | 4 | Attribute | 1..1 |                 Reference number | an..70 | M | R3024 |  | Transport document number | D023 | CIS/Consignment/HouseConsignment/TransportContractDocument/identification | documentNumber |  |  |
| IE3A43 | 4 | Attribute | 1..1 |                 Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/Consignment/HouseConsignment/TransportContractDocument/type | type | CL754 | CL754 |
| IE3A43 | 3 | Composition | 0..1 |             Reference number/UCR |  | O |  |  | UCR | 35B | CIS/Consignment/HouseConsignment/UCR | referenceNumberUCR |  |  |
| IE3A43 | 4 | Attribute | 1..1 |                 Reference number/UCR | an..35 | M |  |  | Trader reference | 009 | CIS/Consignment/HouseConsignment/UCR/traderAssignedReference | referenceNumberUCR |  |  |
| IE3A43 | 1 | Composition | 1..1 |     Declarant |  | M |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3A43 | 2 | Attribute | 1..1 |         Name | an..70 | M |  |  | Declarant name | R124 | CIS/Declarant/name | name |  |  |
| IE3A43 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005<br>R3024 |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3A43 | 2 | Composition | 1..1 |         Address |  | M | R3002 |  | Address | 04A | CIS/Declarant/Address | address |  |  |
| IE3A43 | 3 | Attribute | 1..1 |             City | an..35 | M |  |  | City name | 241 | CIS/Declarant/Address/cityName | city |  |  |
| IE3A43 | 3 | Attribute | 1..1 |             Country | a2 | M |  |  | Country, coded | 242 | CIS/Declarant/Address/country | country | CL718 | CL718 |
| IE3A43 | 3 | Attribute | 0..1 |             Sub-division | an..35 | Ο |  |  | Country sub-entity name | 243 | CIS/Declarant/Address/countrySubDivisionName | subDivision |  |  |
| IE3A43 | 3 | Attribute | 0..1 |             Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Declarant/Address/line | street |  |  |
| IE3A43 | 3 | Attribute | 0..1 |             Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Declarant/Address/postcode | postCode |  |  |
| IE3A43 | 3 | Attribute | 0..1 |             Street additional line | an..70 | O |  |  |  |  | CIS/Declarant/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A43 | 3 | Attribute | 0..1 |             Number | an..35 | O |  |  |  |  | CIS/Declarant/Address/streetNumber | number |  |  |
| IE3A43 | 3 | Attribute | 0..1 |             P.O. Box | an..70 | O |  |  |  |  | CIS/Declarant/Address/poBox | poBox |  |  |
| IE3A43 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/Declarant/Communication | communication |  |  |
| IE3A43 | 3 | Attribute | 1..1 |             Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Declarant/Communication/identification | identifier |  |  |
| IE3A43 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/Declarant/Communication/type | type | CL707 | CL707 |
