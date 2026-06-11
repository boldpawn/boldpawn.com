---
title: "IE3F34 - part 1"
source_title: "IE3F34"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F34.md"
section: "IE3F34"
chunk: 1
chunk_count: 3
generated: true
---

# IE3F34 - part 1

Source document: IE3F34
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F34.md`
Chunk: 1 of 3

# IE3F34

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3F34 | 0 | Class |  | IE3F34 |  |  |  |  | CIS | N.A. | CIS |  |  |  |
| IE3F34 | 1 | Attribute | 1..1 |     LRN | an..22 | M | R3003 |  | Functional reference number | D026 | CIS/functionalReference | LRN |  |  |
| IE3F34 | 1 | Attribute | 1..1 |     Document issue date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  | Document issuing date<br>Response issuing date | D011<br>D029 | CIS/issue | documentIssueDate |  |  |
| IE3F34 | 1 | Attribute | 1..1 |     Specific circumstance indicator | an3 | M |  |  | Specific circumstances, coded | 504 | CIS/specificCircumstances | SpecificCircumstanceIndicator | CL742 | CL742 |
| IE3F34 | 1 | Attribute | 1..1 |     Re-entry indicator | n1 | M |  |  |  |  | CIS/reEntryIndicator | reEntryIndicator | CL027 | CL027 |
| IE3F34 | 1 | Composition | 1..1 |     Split Consignment |  | M |  |  |  |  | CIS/SplitConsignment | splitConsignment |  |  |
| IE3F34 | 2 | Attribute | 1..1 |         Split consignment indicator | n1 | M |  |  |  |  | CIS/SplitConsignment/splitIndicator | splitConsignmentIndicator | CL027 | CL027 |
| IE3F34 | 2 | Attribute | 0..1 |         Previous MRN | an18 | C |  | C3001 |  |  | CIS/SplitConsignment/previousMasterReferenceNumber | previousMRN |  |  |
| IE3F34 | 1 | Composition | 0..1 |     Representative |  | O |  |  | Agent | 05A | CIS/Agent | representative |  |  |
| IE3F34 | 2 | Attribute | 1..1 |         Name | an..70 | M |  |  | Agent - name | R003 | CIS/Agent/name | name |  |  |
| IE3F34 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3033<br>R3005 |  | Agent, coded | R004 | CIS/Agent/identification | identificationNumber |  |  |
| IE3F34 | 2 | Attribute | 1..1 |         Status | n1 | M |  |  | Agent status, coded | 102 | CIS/Agent/function | status | CL094 | CL094 |
| IE3F34 | 2 | Composition | 1..1 |         Address |  | M | R3002 |  | Address | 04A | CIS/Agent/Address | address |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             City | an..35 | M |  |  | City name | 241 | CIS/Agent/Address/cityName | city |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Country | a2 | M |  |  | Country, coded | 242 | CIS/Agent/Address/country | country | CL718 | CL718 |
| IE3F34 | 3 | Attribute | 0..1 |             Sub-division | an..35 | Ο |  |  | Country sub-entity name | 243 | CIS/Agent/Address/countrySubDivisionName | subDivision |  |  |
| IE3F34 | 3 | Attribute | 0..1 |             Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Agent/Address/line | street |  |  |
| IE3F34 | 3 | Attribute | 0..1 |             Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Agent/Address/postcode | postCode |  |  |
| IE3F34 | 3 | Attribute | 0..1 |             Street additional line | an..70 | O |  |  |  |  | CIS/Agent/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F34 | 3 | Attribute | 0..1 |             Number | an..35 | O |  |  |  |  | CIS/Agent/Address/streetNumber | number |  |  |
| IE3F34 | 3 | Attribute | 0..1 |             P.O. Box | an..70 | O |  |  |  |  | CIS/Agent/Address/poBox | poBox |  |  |
| IE3F34 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/Agent/Communication | communication |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Agent/Communication/identification | identifier |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/Agent/Communication/type | type | CL707 | CL707 |
| IE3F34 | 1 | Composition | 1..1 |     Active border transport means |  | M | R3035 |  | BorderTransportMeans | 15A | CIS/BorderTransportMeans | activeBorderTransportMeans |  |  |
| IE3F34 | 2 | Attribute | 1..1 |         Identification number | an..35 | M |  |  | Identification of means of transport crossing the border, coded | T006 | CIS/BorderTransportMeans/identification | identificationNumber |  |  |
| IE3F34 | 2 | Attribute | 1..1 |         Type of identification | n2 | M |  |  | Type of identification of means of transport crossing the border, coded | T018 | CIS/BorderTransportMeans/identificationType | typeOfIdentification | CL750 | CL750 |
| IE3F34 | 2 | Attribute | 1..1 |         Type of means of transport | an..4 | M |  |  | Type of means of transport crossing the border, coded | T010 | CIS/BorderTransportMeans/type | typeOfMeansOfTransport | CL751 | CL751 |
| IE3F34 | 2 | Attribute | 1..1 |         Nationality | a2 | M |  |  | Nationality of means of transport crossing the border, coded | T014 | CIS/BorderTransportMeans/registrationNationality | nationality | CL718 | CL718 |
| IE3F34 | 2 | Attribute | 1..1 |         Mode of transport | n1 | M |  |  | Mode of transport crossing the border, coded | T022 | CIS/BorderTransportMeans/mode | ModeOfTransport | CL018 | CL018 |
| IE3F34 | 2 | Attribute | 0..1 |         Actual date and time of departure | an20 (YYYY-MM-DDThh:mm:ssZ) | O | R3053 |  |  |  | CIS/BorderTransportMeans/actualDepartureDateTime | actualDateAndTimeOfDeparture |  |  |
| IE3F34 | 2 | Attribute | 0..1 |         Estimated date and time of departure | an20 (YYYY-MM-DDThh:mm:ssZ) | O | R3053 |  |  |  | CIS/BorderTransportMeans/estimatedDepartureDateTime | estimatedDateAndTimeOfDeparture |  |  |
| IE3F34 | 2 | Attribute | 0..1 |         Estimated date and time of arrival | an20 (YYYY-MM-DDThh:mm:ssZ) | O | R3053 |  |  |  | CIS/BorderTransportMeans/estimatedArrivalDateTime | estimatedDateAndTimeOfArrival |  |  |
| IE3F34 | 2 | Composition | 2..99 |         Countries of routing of means of transport |  | M | R3018<br>R3013<br>R3014 |  | Itinerary | 81A | CIS/BorderTransportMeans/Itinerary | countriesOfRoutingOfMeansOfTransport |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Sequence number | n..2 | M |  |  | Sequence number | 006 | CIS/BorderTransportMeans/Itinerary/sequence | sequenceNumber |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Country | a2 | M |  |  | Country(ies ) of routing, coded | 064 | CIS/BorderTransportMeans/Itinerary/routingCountry | country | CL718 | CL718 |
| IE3F34 | 1 | Composition | 1..1 |     Consignment (Master level) |  | M |  |  | Consignment | 28A | CIS/Consignment | consignmentMasterLevel |  |  |
| IE3F34 | 2 | Composition | 1..1 |         Place of acceptance |  | M |  |  | AcceptancePlace | 01A | CIS/Consignment/AcceptancePlace | placeOfAcceptance |  |  |
| IE3F34 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3003 | Place of acceptance | L011 | CIS/Consignment/AcceptancePlace/name | location |  |  |
| IE3F34 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Place of acceptance, coded | L099 | CIS/Consignment/AcceptancePlace/identification | UNLOCODE | CL244 | CL244 |
| IE3F34 | 3 | Composition | 0..1 |             Address |  | C |  | C3003 | Address | 04A | CIS/Consignment/AcceptancePlace/Address | address |  |  |
| IE3F34 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/AcceptancePlace/Address/country | country | CL718 | CL718 |
| IE3F34 | 2 | Composition | 0..99 |         Supporting documents |  | O |  |  | AdditionalDocument | 02A | CIS/Consignment/AdditionalDocument | supportingDocuments |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Additional document reference number | D005 | CIS/Consignment/AdditionalDocument/identification | referenceNumber |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Additional document type, coded | D006 | CIS/Consignment/AdditionalDocument/type | type | CL013 | CL013 |
| IE3F34 | 2 | Composition | 0..99 |         Additional supply chain actor |  | O |  |  | AEOMutualRecognitionParty | 16C | CIS/Consignment/AEOMutualRecognitionParty | additionalSupplyChainActor |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Identification number | an..17 | M |  |  | AEO mutual recognition party, coded | R143 | CIS/Consignment/AEOMutualRecognitionParty/identification | identificationNumber |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Role | a..3 | M |  |  | Role code | R005 | CIS/Consignment/AEOMutualRecognitionParty/role | role | CL704 | CL704 |
| IE3F34 | 2 | Composition | 1..1 |         Carrier |  | M |  |  | Carrier | 18A | CIS/Consignment/Carrier | carrier |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Name | an..70 | M |  |  | Carrier - name | R011 | CIS/Consignment/Carrier/name | name |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Identification number | an..17 | M | R3005 |  | Carrier identification | R012 | CIS/Consignment/Carrier/identification | identificationNumber |  |  |
| IE3F34 | 3 | Composition | 1..1 |             Address |  | M | R3002 |  | Address | 04A | CIS/Consignment/Carrier/Address | address |  |  |
| IE3F34 | 4 | Attribute | 1..1 |                 City | an..35 | M |  |  | City name | 241 | CIS/Consignment/Carrier/Address/cityName | city |  |  |
| IE3F34 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/Carrier/Address/country | country | CL718 | CL718 |
| IE3F34 | 4 | Attribute | 0..1 |                 Sub-division | an..35 | O |  |  | Country sub-entity name | 243 | CIS/Consignment/Carrier/Address/countrySubDivisionName | subDivision |  |  |
| IE3F34 | 4 | Attribute | 0..1 |                 Street | an..70 | Ο |  |  | Street and number/P.O. Box | 239 | CIS/Consignment/Carrier/Address/line | street |  |  |
| IE3F34 | 4 | Attribute | 0..1 |                 Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Consignment/Carrier/Address/postcode | postCode |  |  |
| IE3F34 | 4 | Attribute | 0..1 |                 Street additional line | an..70 | O |  |  |  |  | CIS/Consignment/Carrier/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F34 | 4 | Attribute | 0..1 |                 Number | an..35 | Ο |  |  |  |  | CIS/Consignment/Carrier/Address/streetNumber | number |  |  |
| IE3F34 | 4 | Attribute | 0..1 |                 P.O. Box | an..70 | Ο |  |  |  |  | CIS/Consignment/Carrier/Address/poBox | poBox |  |  |
| IE3F34 | 3 | Composition | 0..9 |             Communication |  | O |  |  | Communication | 25A | CIS/Consignment/Carrier/Communication | communication |  |  |
| IE3F34 | 4 | Attribute | 1..1 |                 Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Consignment/Carrier/Communication/identification | identifier |  |  |
| IE3F34 | 4 | Attribute | 1..1 |                 Type | an..3 | M |  |  | Communication number type | 253 | CIS/Consignment/Carrier/Communication/type | type | CL707 | CL707 |
| IE3F34 | 2 | Composition | 1..1 |         Place of delivery |  | M |  |  | GoodsReceiptPlace | 66A | CIS/Consignment/GoodsReceiptPlace | placeOfDelivery |  |  |
| IE3F34 | 3 | Attribute | 0..1 |             Location | an..35 | C |  | C3005 | Goods receipt place | L003 | CIS/Consignment/GoodsReceiptPlace/name | location |  |  |
| IE3F34 | 3 | Attribute | 0..1 |             UNLOCODE | an..17 | O |  |  | Goods receipt place, coded | L004 | CIS/Consignment/GoodsReceiptPlace/identification | UNLOCODE | CL244 | CL244 |
| IE3F34 | 3 | Composition | 0..1 |             Address |  | C |  | C3005 | Address | 04A | CIS/Consignment/GoodsReceiptPlace/Address | address |  |  |
| IE3F34 | 4 | Attribute | 1..1 |                 Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/GoodsReceiptPlace/Address/country | country | CL718 | CL718 |
| IE3F34 | 2 | Composition | 1..99999 |         Consignment (House level) |  | M |  |  | HouseConsignment | 21C | CIS/Consignment/HouseConsignment | consignmentHouseLevel |  |  |
| IE3F34 | 3 | Attribute | 1..1 |             Container indicator | n1 | M |  |  | Container transport code | 096 | CIS/Consignment/HouseConsignment/container | containerIndicator | CL708 | CL708 |
| IE3F34 | 3 | Attribute | 1..1 |             Total gross mass | n..16,6 | M |  |  | Total gross weight | 131 | CIS/Consignment/HouseConsignment/totalGrossMass | totalGrossMass |  |  |
| IE3F34 | 3 | Composition | 1..1 |             Place of acceptance |  | M |  |  | AcceptancePlace | 01A | CIS/Consignment/HouseConsignment/AcceptancePlace | placeOfAcceptance |  |  |
| IE3F34 | 4 | Attribute | 0..1 |                 Location | an..35 | C |  | C3004 | Place of acceptance | L011 | CIS/Consignment/HouseConsignment/AcceptancePlace/name | location |  |  |
| IE3F34 | 4 | Attribute | 0..1 |                 UNLOCODE | an..17 | O |  |  | Place of acceptance, coded | L099 | CIS/Consignment/HouseConsignment/AcceptancePlace/identification | UNLOCODE | CL244 | CL244 |
| IE3F34 | 4 | Composition | 0..1 |                 Address |  | C |  | C3004 | Address | 04A | CIS/Consignment/HouseConsignment/AcceptancePlace/Address | address |  |  |
| IE3F34 | 5 | Attribute | 1..1 |                     Country | a2 | M |  |  | Country, coded | 242 | CIS/Consignment/HouseConsignment/AcceptancePlace/Address/country | country | CL718 | CL718 |
| IE3F34 | 3 | Composition | 0..99 |             Supporting documents |  | O |  |  | AdditionalDocument | 02A | CIS/Consignment/HouseConsignment/AdditionalDocument | supportingDocuments |  |  |
| IE3F34 | 4 | Attribute | 1..1 |                 Reference number | an..70 | M |  |  | Additional document reference number | D005 | CIS/Consignment/HouseConsignment/AdditionalDocument/identification | referenceNumber |  |  |
| IE3F34 | 4 | Attribute | 1..1 |                 Type | an4 | M |  |  | Additional document type, coded | D006 | CIS/Consignment/HouseConsignment/AdditionalDocument/type | type | CL013 | CL013 |
