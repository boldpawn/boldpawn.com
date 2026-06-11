---
title: "IE3N06"
source_title: "IE3N06"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N06.md"
section: "IE3N06"
chunk: 1
chunk_count: 1
generated: true
---

# IE3N06

Source document: IE3N06
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N06.md`
Chunk: 1 of 1

# IE3N06

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3N06 | 0 | Class |  | IE3N06 |  |  |  |  | CIS | N.A. | CIS | IE3N06 |  |  |
| IE3N06 | 1 | Attribute | 1..1 |     LRN | an..22 | M | R3003 |  | Functional reference number | D026 | CIS/functionalReference | LRN |  |  |
| IE3N06 | 1 | Attribute | 1..1 |     Document issue date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  | Document issuing date<br>Response issuing date | D011<br>D029 | CIS/issue | documentIssueDate |  |  |
| IE3N06 | 1 | Composition | 0..9 |     Notify Party |  | O |  |  | Agent | 05A | CIS/Agent | notifyParty |  |  |
| IE3N06 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005 |  | Agent, coded | R004 | CIS/Agent/identification | identificationNumber |  |  |
| IE3N06 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/Agent/Communication | communication |  |  |
| IE3N06 | 3 | Attribute | 1..1 |             Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Agent/Communication/identification | identifier |  |  |
| IE3N06 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/Agent/Communication/type | type | CL707 | CL707 |
| IE3N06 | 1 | Composition | 1..1 |     Active border transport means |  | M |  |  | BorderTransportMeans | 15A | CIS/BorderTransportMeans | activeBorderTransportMeans |  |  |
| IE3N06 | 2 | Attribute | 0..1 |         Identification number | an..35 | C |  | C3012 | Identification of means of transport crossing the border, coded | T006 | CIS/BorderTransportMeans/identification | identificationNumber |  |  |
| IE3N06 | 2 | Attribute | 0..1 |         Type of identification | n2 | C | R3011<br>R3012 | C3012 | Type of identification of means of transport crossing the border, coded | T018 | CIS/BorderTransportMeans/identificationType | typeOfIdentification | CL750 | CL750 |
| IE3N06 | 2 | Attribute | 1..1 |         Mode of transport | n1 | M |  |  | Mode of transport crossing the border, coded | T022 | CIS/BorderTransportMeans/mode | ModeOfTransport | CL018 | CL018 |
| IE3N06 | 2 | Attribute | 1..1 |         Actual date and time of arrival | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  |  |  | CIS/BorderTransportMeans/actualArrivalDateTime | actualDateAndTimeOfArrival |  |  |
| IE3N06 | 2 | Attribute | 0..1 |         Estimated date and time of arrival | an20 (YYYY-MM-DDThh:mm:ssZ) | O | R3027 |  |  |  | CIS/BorderTransportMeans/estimatedArrivalDateTime | estimatedDateAndTimeOfArrival |  |  |
| IE3N06 | 2 | Attribute | 0..1 |         Conveyance reference number | an..17 | C |  | C3019 | Conveyance reference number | 149 | CIS/BorderTransportMeans/journey | conveyanceReferenceNumber |  |  |
| IE3N06 | 1 | Composition | 0..9999 |     Related transport document |  | O | R3048 |  | TransportContractDocument | 30B | CIS/TransportContractDocument | relatedTransportDocument |  |  |
| IE3N06 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocument/identification | documentNumber |  |  |
| IE3N06 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocument/type | type | CL754 | CL754 |
| IE3N06 | 1 | Composition | 0..9999 |     Related MRN |  | C | R3027<br>R3048 |  | PreviousDocument | 99A | CIS/PreviousDocument | relatedMRN |  |  |
| IE3N06 | 2 | Attribute | 1..1 |         MRN | an18 | M |  |  | Previous document number | D018 | CIS/PreviousDocument/identification | identification |  |  |
| IE3N06 | 1 | Composition | 0..1 |     Person notifying the arrival |  | C |  | C3049 | Submitter | 17B | CIS/PersonNotifyingArrival | personNotifyingTheArrival |  |  |
| IE3N06 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005 |  | Submitter, coded | R059 | CIS/PersonNotifyingArrival/identification | identificationNumber |  |  |
| IE3N06 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/PersonNotifyingArrival/Communication | communication |  |  |
| IE3N06 | 3 | Attribute | 1..1 |             Identifier | an..512 | M |  |  | Communication number | 240 | CIS/PersonNotifyingArrival/Communication/identification | identifier |  |  |
| IE3N06 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/PersonNotifyingArrival/Communication/type | type | CL707 | CL707 |
| IE3N06 | 1 | Composition | 1..1 |     Actual customs office of first entry |  | M |  |  | EntryOffice | 51A | CIS/EntryOffice | actualCustomsOfficeOfFirstEntry |  |  |
| IE3N06 | 2 | Attribute | 1..1 |         Reference number | an8 | M |  |  | Office of entry, coded | G004 | CIS/EntryOffice/identification | referenceNumber | CL141 | CL141 |
