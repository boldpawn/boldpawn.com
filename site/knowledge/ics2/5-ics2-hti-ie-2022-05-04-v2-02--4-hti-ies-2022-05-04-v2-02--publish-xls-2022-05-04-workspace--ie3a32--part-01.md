---
title: "IE3A32 - part 1"
source_title: "IE3A32"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A32.md"
section: "IE3A32"
chunk: 1
chunk_count: 2
generated: true
---

# IE3A32 - part 1

Source document: IE3A32
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3A32.md`
Chunk: 1 of 2

# IE3A32

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3A32 | 0 | Class |  | IE3A32 |  |  |  |  | CIS | N.A. | CIS | IE3A32 |  |  |
| IE3A32 | 1 | Attribute | 1..1 |     Document issue date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  | Document issuing date<br>Response issuing date | D011<br>D029 | CIS/issue | documentIssueDate |  |  |
| IE3A32 | 1 | Attribute | 0..1 |     Referral request reference | an..17 | O |  |  |  |  | CIS/requestReference | referralRequestReference |  |  |
| IE3A32 | 1 | Attribute | 1..1 |     MRN | an18 | M | R3004 |  |  |  | CIS/masterReferenceNumber | MRN |  |  |
| IE3A32 | 1 | Attribute | 1..1 |     Specific circumstance indicator | an3 | M | R3024 |  | Specific circumstances, coded | 504 | CIS/specificCircumstances | SpecificCircumstanceIndicator | CL742 | CL742 |
| IE3A32 | 1 | Composition | 1..1 |     Addressed Member State |  | M |  |  | AddressedMemberState | SC2 | CIS/AddressedMemberState | addressedMemberState |  |  |
| IE3A32 | 2 | Attribute | 1..1 |         Country | a2 | M |  |  |  |  | CIS/AddressedMemberState/country | country | CL717 | CL717 |
| IE3A32 | 1 | Composition | 0..1 |     Representative |  | O | R3024 |  | Agent | 05A | CIS/Agent | representative |  |  |
| IE3A32 | 2 | Attribute | 1..1 |         Name | an..70 | M |  |  | Agent - name | R003 | CIS/Agent/name | name |  |  |
| IE3A32 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3033<br>R3005<br>R3024 |  | Agent, coded | R004 | CIS/Agent/identification | identificationNumber |  |  |
| IE3A32 | 2 | Attribute | 1..1 |         Status | n1 | M |  |  | Agent status, coded | 102 | CIS/Agent/function | status | CL094 | CL094 |
| IE3A32 | 2 | Composition | 1..1 |         Address |  | M | R3002 |  | Address | 04A | CIS/Agent/Address | address |  |  |
| IE3A32 | 3 | Attribute | 1..1 |             City | an..35 | M |  |  | City name | 241 | CIS/Agent/Address/cityName | city |  |  |
| IE3A32 | 3 | Attribute | 1..1 |             Country | a2 | M |  |  | Country, coded | 242 | CIS/Agent/Address/country | country | CL718 | CL718 |
| IE3A32 | 3 | Attribute | 0..1 |             Sub-division | an..35 | Ο |  |  | Country sub-entity name | 243 | CIS/Agent/Address/countrySubDivisionName | subDivision |  |  |
| IE3A32 | 3 | Attribute | 0..1 |             Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Agent/Address/line | street |  |  |
| IE3A32 | 3 | Attribute | 0..1 |             Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Agent/Address/postcode | postCode |  |  |
| IE3A32 | 3 | Attribute | 0..1 |             Street additional line | an..70 | O |  |  |  |  | CIS/Agent/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A32 | 3 | Attribute | 0..1 |             Number | an..35 | O |  |  |  |  | CIS/Agent/Address/streetNumber | number |  |  |
| IE3A32 | 3 | Attribute | 0..1 |             P.O. Box | an..70 | O |  |  |  |  | CIS/Agent/Address/poBox | poBox |  |  |
| IE3A32 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/Agent/Communication | communication |  |  |
| IE3A32 | 3 | Attribute | 1..1 |             Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Agent/Communication/identification | identifier |  |  |
| IE3A32 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/Agent/Communication/type | type | CL707 | CL707 |
| IE3A32 | 1 | Composition | 1..1 |     Consignment (Master level) |  | M |  |  | Consignment | 28A | CIS/Consignment | consignmentMasterLevel |  |  |
| IE3A32 | 2 | Composition | 1..99999 |         Consignment (House level) |  | M |  |  | HouseConsignment | 21C | CIS/Consignment/HouseConsignment | consignmentHouseLevel |  |  |
| IE3A32 | 3 | Attribute | 0..1 |             Total amount invoiced |  | C |  | C3038 | Total invoice amount | 109 | CIS/Consignment/HouseConsignment/invoice | totalAmountInvoiced |  |  |
| IE3A32 | 4 | Attribute | 1..1 |                 Value | n..16,2 | M |  |  |  |  | CIS/Consignment/HouseConsignment/invoice/Amount. Content | value |  |  |
| IE3A32 | 4 | Attribute | 1..1 |                 Currency | a3 | M |  |  |  |  | CIS/Consignment/HouseConsignment/invoice/Amount Currency. Identifier | currency | CL352 | CL352 |
| IE3A32 | 3 | Attribute | 1..1 |             Total gross mass | n..16,6 | M |  |  | Total gross weight | 131 | CIS/Consignment/HouseConsignment/totalGrossMass | totalGrossMass |  |  |
| IE3A32 | 3 | Composition | 0..1 |             Additions and deductions |  | C |  | C3037 | ChargeDeduction | 58B | CIS/Consignment/HouseConsignment/ChargeDeduction | additionsAndDeductions |  |  |
| IE3A32 | 4 | Attribute | 1..1 |                 Amount |  | M |  |  | Other charges/deductions | 181 | CIS/Consignment/HouseConsignment/ChargeDeduction/otherChargeDeduction | amount |  |  |
| IE3A32 | 5 | Attribute | 1..1 |                     Value | n..16,2 | M |  |  |  |  | CIS/Consignment/HouseConsignment/ChargeDeduction/otherChargeDeduction/Amount. Content | value |  |  |
| IE3A32 | 5 | Attribute | 1..1 |                     Currency | a3 | M |  |  |  |  | CIS/Consignment/HouseConsignment/ChargeDeduction/otherChargeDeduction/Amount Currency. Identifier | currency | CL352 | CL352 |
| IE3A32 | 3 | Composition | 0..1 |             Additional fiscal references |  | O |  |  | DomesticDutyTaxParty | 55B | CIS/Consignment/HouseConsignment/DomesticDutyTaxParty | additionalFiscalReferences |  |  |
| IE3A32 | 4 | Attribute | 1..1 |                 VAT identification number | an..17 | M |  |  | DomesticDutyTaxParty, coded | R119 | CIS/Consignment/HouseConsignment/DomesticDutyTaxParty/identification | VATIdentificationNumber |  |  |
| IE3A32 | 4 | Attribute | 1..1 |                 Role | an..3 | M |  |  | Role code | R005 | CIS/Consignment/HouseConsignment/DomesticDutyTaxParty/role | role | CL756 | CL756 |
| IE3A32 | 3 | Composition | 0..99 |             Supporting documents |  | O |  |  | AdditionalDocument | 02A | CIS/Consignment/HouseConsignment/AdditionalDocument | supportingDocuments |  |  |
| IE3A32 | 4 | Attribute | 1..1 |                 Reference number | an..70 | M |  |  | Additional document reference number | D005 | CIS/Consignment/HouseConsignment/AdditionalDocument/identification | referenceNumber |  |  |
| IE3A32 | 4 | Attribute | 1..1 |                 Type | an4 | M |  |  | Additional document type, coded | D006 | CIS/Consignment/HouseConsignment/AdditionalDocument/type | type | CL013 | CL013 |
| IE3A32 | 3 | Composition | 0..99 |             Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/Consignment/HouseConsignment/AdditionalInformation | additionalInformation |  |  |
| IE3A32 | 4 | Attribute | 0..1 |                 Code | an..17 | O |  |  | Additional statement code | 226 | CIS/Consignment/HouseConsignment/AdditionalInformation/statement | code | CL701 | CL701 |
| IE3A32 | 4 | Attribute | 0..1 |                 Text | an..512 | O |  |  | Additional statement text | 225 | CIS/Consignment/HouseConsignment/AdditionalInformation/statementDescription | text |  |  |
| IE3A32 | 3 | Composition | 0..9 |             Additional supply chain actor |  | O |  |  | AEOMutualRecognitionParty | 16C | CIS/Consignment/HouseConsignment/AEOMutualRecognitionParty | additionalSupplyChainActor |  |  |
| IE3A32 | 4 | Attribute | 1..1 |                 Identification number | an..17 | M |  |  | AEO mutual recognition party, coded | R143 | CIS/Consignment/HouseConsignment/AEOMutualRecognitionParty/identification | identificationNumber |  |  |
| IE3A32 | 4 | Attribute | 1..1 |                 Role | a..3 | M |  |  | Role code | R005 | CIS/Consignment/HouseConsignment/AEOMutualRecognitionParty/role | role | CL704 | CL704 |
| IE3A32 | 3 | Composition | 1..1 |             Consignee |  | M |  |  | Consignee | 27A | CIS/Consignment/HouseConsignment/Consignee | consignee |  |  |
| IE3A32 | 4 | Attribute | 1..1 |                 Name | an..70 | M |  |  | Consignee name | R014 | CIS/Consignment/HouseConsignment/Consignee/name | name |  |  |
| IE3A32 | 4 | Attribute | 0..1 |                 Identification number | an..17 | O |  |  | Consignee, coded | R015 | CIS/Consignment/HouseConsignment/Consignee/identification | identificationNumber |  |  |
| IE3A32 | 4 | Attribute | 1..1 |                 Type of person | n1 | M |  |  |  |  | CIS/Consignment/HouseConsignment/Consignee/identificationType | typeOfPerson | CL729 | CL729 |
| IE3A32 | 4 | Composition | 1..1 |                 Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/HouseConsignment/Consignee/Address | address |  |  |
| IE3A32 | 5 | Attribute | 1..1 |                     City | an..35 | M |  |  | City name | 241 | CIS/Consignment/HouseConsignment/Consignee/Address/cityName | city |  |  |
| IE3A32 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/Consignee/Address/country | country | CL718 | CL718 |
| IE3A32 | 5 | Attribute | 0..1 |                     Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/HouseConsignment/Consignee/Address/countrySubDivisionName | subDivision |  |  |
| IE3A32 | 5 | Attribute | 0..1 |                     Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/HouseConsignment/Consignee/Address/line | street |  |  |
| IE3A32 | 5 | Attribute | 0..1 |                     Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/HouseConsignment/Consignee/Address/postcode | postCode |  |  |
| IE3A32 | 5 | Attribute | 0..1 |                     Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignee/Address/additionalLine | streetAdditionalLine |  |  |
| IE3A32 | 5 | Attribute | 0..1 |                     Number | an..35 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignee/Address/streetNumber | number |  |  |
| IE3A32 | 5 | Attribute | 0..1 |                     P.O. Box | an..70 | O |  |  |  |  | CIS/Consignment/HouseConsignment/Consignee/Address/poBox | poBox |  |  |
| IE3A32 | 4 | Composition | 0..9 |                 Communication |  | O |  |  | Communication | 25A | CIS/Consignment/HouseConsignment/Consignee/Communication | communication |  |  |
| IE3A32 | 5 | Attribute | 1..1 |                     Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/HouseConsignment/Consignee/Communication/identification | identifier |  |  |
| IE3A32 | 5 | Attribute | 1..1 |                     Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/HouseConsignment/Consignee/Communication/type | type | CL707 | CL707 |
| IE3A32 | 3 | Composition | 1..9999 |             Goods item |  | M |  |  | ConsignmentItem | 29A | CIS/Consignment/HouseConsignment/ConsignmentItem | goodsItem |  |  |
| IE3A32 | 4 | Attribute | 0..1 |                 Item amount invoiced |  | O |  |  | Item amount | 112 | CIS/Consignment/HouseConsignment/ConsignmentItem/itemCharge | itemAmountInvoiced |  |  |
| IE3A32 | 5 | Attribute | 1..1 |                     Value | n..16,2 | M |  |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/itemCharge/Amount. Content | value |  |  |
| IE3A32 | 5 | Attribute | 1..1 |                     Currency | a3 | M |  |  |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/itemCharge/Amount Currency. Identifier | currency | CL352 | CL352 |
| IE3A32 | 4 | Attribute | 1..1 |                 Goods item number | n..5 | M | R3017<br>R3024 |  | Sequence number | 006 | CIS/Consignment/HouseConsignment/ConsignmentItem/sequence | goodsItemNumber |  |  |
| IE3A32 | 4 | Attribute | 0..1 |                 Type of goods | an..3 | C |  | C3037 |  |  | CIS/Consignment/HouseConsignment/ConsignmentItem/postalType | typeOfGoods | CL749 | CL749 |
| IE3A32 | 4 | Composition | 0..99 |                 Supporting documents |  | O |  |  | AdditionalDocument | 02A | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalDocument | supportingDocuments |  |  |
| IE3A32 | 5 | Attribute | 1..1 |                     Reference number | an..70 | M |  |  | Additional document reference number | D005 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalDocument/identification | referenceNumber |  |  |
| IE3A32 | 5 | Attribute | 1..1 |                     Type | an4 | M |  |  | Additional document type, coded | D006 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalDocument/type | type | CL013 | CL013 |
| IE3A32 | 4 | Composition | 0..99 |                 Additional information |  | O |  |  | AdditionalInformation | 03A | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalInformation | additionalInformation |  |  |
| IE3A32 | 5 | Attribute | 0..1 |                     Code | an..17 | O |  |  | Additional statement code | 226 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalInformation/statement | code | CL701 | CL701 |
| IE3A32 | 5 | Attribute | 0..1 |                     Text | an..512 | O |  |  | Additional statement text | 225 | CIS/Consignment/HouseConsignment/ConsignmentItem/AdditionalInformation/statementDescription | text |  |  |
| IE3A32 | 4 | Composition | 0..9 |                 Additional supply chain actor |  | O |  |  | AEOMutualRecognitionParty | 16C | CIS/Consignment/HouseConsignment/ConsignmentItem/AEOMutualRecognitionParty | additionalSupplyChainActor |  |  |
| IE3A32 | 5 | Attribute | 1..1 |                     Identification number | an..17 | M |  |  | AEO mutual recognition party, coded | R143 | CIS/Consignment/HouseConsignment/ConsignmentItem/AEOMutualRecognitionParty/identification | identificationNumber |  |  |
