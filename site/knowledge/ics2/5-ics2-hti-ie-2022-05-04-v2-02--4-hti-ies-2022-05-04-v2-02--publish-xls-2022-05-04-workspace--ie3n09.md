---
title: "IE3N09"
source_title: "IE3N09"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N09.md"
section: "IE3N09"
chunk: 1
chunk_count: 1
generated: true
---

# IE3N09

Source document: IE3N09
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N09.md`
Chunk: 1 of 1

# IE3N09

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3N09 | 0 | Class |  | IE3N09 |  |  |  |  | CIS | N.A. | CIS | IE3N09 |  |  |
| IE3N09 | 1 | Attribute | 1..1 |     MRN | an18 | M |  |  |  |  | CIS/masterReferenceNumber | MRN |  |  |
| IE3N09 | 1 | Attribute | 1..1 |     Notification date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  |  |  | CIS/notificationDate | notificationDate |  |  |
| IE3N09 | 1 | Attribute | 0..1 |     Scheduled control date | an20 (YYYY-MM-DDThh:mm:ssZ) | O |  |  |  |  | CIS/scheduledControlDate | scheduledControlDate |  |  |
| IE3N09 | 1 | Composition | 1..1 |     Customs office of control |  | M |  |  | ControlOffice | SC2 | CIS/ControlOffice | customsOfficeOfControl |  |  |
| IE3N09 | 2 | Attribute | 1..1 |         Reference number | an8 | M |  |  | Government agency name, coded |  | CIS/ControlOffice/identification | referenceNumber | CL141 | CL141 |
| IE3N09 | 1 | Composition | 0..1 |     Representative |  | O |  |  | Agent | 05A | CIS/Agent | representative |  |  |
| IE3N09 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Agent, coded | R004 | CIS/Agent/identification | identificationNumber |  |  |
| IE3N09 | 1 | Composition | 0..1 |     Transport document (Master level) |  | O |  |  | TransportContractDocument | 30B | CIS/TransportContractDocument | transportDocument |  |  |
| IE3N09 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocument/identification | documentNumber |  |  |
| IE3N09 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocument/type | type | CL754 | CL754 |
| IE3N09 | 1 | Composition | 0..1 |     Carrier |  | O |  |  | Carrier | 18A | CIS/Carrier | carrier |  |  |
| IE3N09 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Carrier identification | R012 | CIS/Carrier/identification | identificationNumber |  |  |
| IE3N09 | 1 | Composition | 1..99999 |     Control |  | M |  |  | Control | 38A | CIS/Control | control |  |  |
| IE3N09 | 2 | Composition | 1..1 |         Examination place |  | M | R3029 |  | ExaminationPlace | 54A | CIS/Control/ExaminationPlace | examinationPlace |  |  |
| IE3N09 | 3 | Attribute | 0..1 |             Place of examination | an..256 | O |  |  | Place of physical examination | L069 | CIS/Control/ExaminationPlace/name | placeOfExamination |  |  |
| IE3N09 | 3 | Attribute | 0..1 |             Reference number | an..17 | O |  |  | Place of physical examination, coded | L029 | CIS/Control/ExaminationPlace/identification | referenceNumber |  |  |
| IE3N09 | 2 | Composition | 1..1 |         Control subject |  | M |  |  |  |  | CIS/Control/ControlSubject | controlSubject |  |  |
| IE3N09 | 3 | Composition | 1..1 |             Consignment (Master level) |  | M | R3030 |  | Consignment | 28A | CIS/Control/ControlSubject/Consignment | consignmentMasterLevel |  |  |
| IE3N09 | 4 | Composition | 0..999 |                 Receptacle |  | O |  |  |  |  | CIS/Control/ControlSubject/Consignment/PostalReceptacle | receptacle |  |  |
| IE3N09 | 5 | Attribute | 1..1 |                     Receptacle identification number | an..35 | M |  |  |  |  | CIS/Control/ControlSubject/Consignment/PostalReceptacle/receptacleIdentificationNumber | receptacleIdentificationNumber |  |  |
| IE3N09 | 4 | Composition | 0..9999 |                 Goods item |  | O |  |  | ConsignmentItem | 29A | CIS/Control/ControlSubject/Consignment/ConsignmentItem | goodsItem |  |  |
| IE3N09 | 5 | Attribute | 1..1 |                     Goods item number | n..5 | M |  |  | Sequence number | 006 | CIS/Control/ControlSubject/Consignment/ConsignmentItem/sequence | goodsItemNumber |  |  |
| IE3N09 | 5 | Composition | 0..99 |                     Packaging |  | O |  |  | Packaging | 93A | CIS/Control/ControlSubject/Consignment/ConsignmentItem/Packaging | packaging |  |  |
| IE3N09 | 6 | Attribute | 1..1 |                         Shipping marks | an..512 | M |  |  | Shipping marks | 142 | CIS/Control/ControlSubject/Consignment/ConsignmentItem/Packaging/marksNumbers | shippingMarks |  |  |
| IE3N09 | 6 | Attribute | 1..1 |                         Type of packages | an2 | M |  |  | Type of packages identification, coded | 141 | CIS/Control/ControlSubject/Consignment/ConsignmentItem/Packaging/type | typeOfPackages | CL017 | CL017 |
| IE3N09 | 5 | Composition | 0..9999 |                     Transport equipment |  | O |  |  | TransportEquipment | 31B | CIS/Control/ControlSubject/Consignment/ConsignmentItem/TransportEquipment | transportEquipment |  |  |
| IE3N09 | 6 | Attribute | 1..1 |                         Container identification number | an..17 | M |  |  | Equipment identification number | 159 | CIS/Control/ControlSubject/Consignment/ConsignmentItem/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3N09 | 4 | Composition | 0..99999 |                 Consignment (House level) |  | O |  |  | HouseConsignment | 21C | CIS/Control/ControlSubject/Consignment/HouseConsignment | consignmentHouseLevel |  |  |
| IE3N09 | 5 | Composition | 0..9999 |                     Goods item |  | O |  |  | ConsignmentItem | 29A | CIS/Control/ControlSubject/Consignment/HouseConsignment/ConsignmentItem | goodsItem |  |  |
| IE3N09 | 6 | Attribute | 1..1 |                         Goods item number | n..5 | M |  |  | Sequence number | 006 | CIS/Control/ControlSubject/Consignment/HouseConsignment/ConsignmentItem/sequence | goodsItemNumber |  |  |
| IE3N09 | 6 | Composition | 0..99 |                         Packaging |  | O |  |  | Packaging | 93A | CIS/Control/ControlSubject/Consignment/HouseConsignment/ConsignmentItem/Packaging | packaging |  |  |
| IE3N09 | 7 | Attribute | 1..1 |                             Shipping marks | an..512 | M |  |  | Shipping marks | 142 | CIS/Control/ControlSubject/Consignment/HouseConsignment/ConsignmentItem/Packaging/marksNumbers | shippingMarks |  |  |
| IE3N09 | 7 | Attribute | 1..1 |                             Type of packages | an2 | M |  |  | Type of packages identification, coded | 141 | CIS/Control/ControlSubject/Consignment/HouseConsignment/ConsignmentItem/Packaging/type | typeOfPackages | CL017 | CL017 |
| IE3N09 | 6 | Composition | 0..99 |                         Passive border transport means |  | O |  |  | PassiveTransportMeans | 19C | CIS/Control/ControlSubject/Consignment/HouseConsignment/ConsignmentItem/PassiveTransportMeans | passiveBorderTransportMeans |  |  |
| IE3N09 | 7 | Attribute | 1..1 |                             Identification number | an..35 | M |  |  | Identification of passive transport means, coded | T026 | CIS/Control/ControlSubject/Consignment/HouseConsignment/ConsignmentItem/PassiveTransportMeans/identification | identificationNumber |  |  |
| IE3N09 | 7 | Attribute | 1..1 |                             Type of identification | n2 | M |  |  | Type of identification of passive transport means, coded | T027 | CIS/Control/ControlSubject/Consignment/HouseConsignment/ConsignmentItem/PassiveTransportMeans/identificationType | typeOfIdentification | CL750 | CL750 |
| IE3N09 | 7 | Attribute | 0..1 |                             Nationality | a2 | O |  |  | Nationality of passive means of transport, coded | T028 | CIS/Control/ControlSubject/Consignment/HouseConsignment/ConsignmentItem/PassiveTransportMeans/registrationNationality | nationality | CL718 | CL718 |
| IE3N09 | 6 | Composition | 0..9999 |                         Transport equipment |  | O |  |  | TransportEquipment | 31B | CIS/Control/ControlSubject/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment | transportEquipment |  |  |
| IE3N09 | 7 | Attribute | 1..1 |                             Container identification number | an..17 | M |  |  | Equipment identification number | 159 | CIS/Control/ControlSubject/Consignment/HouseConsignment/ConsignmentItem/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3N09 | 5 | Composition | 0..99 |                     Passive border transport means |  | O |  |  | PassiveTransportMeans | 19C | CIS/Control/ControlSubject/Consignment/HouseConsignment/PassiveTransportMeans | passiveBorderTransportMeans |  |  |
| IE3N09 | 6 | Attribute | 1..1 |                         Identification number | an..35 | M |  |  | Identification of passive transport means, coded | T026 | CIS/Control/ControlSubject/Consignment/HouseConsignment/PassiveTransportMeans/identification | identificationNumber |  |  |
| IE3N09 | 6 | Attribute | 1..1 |                         Type of identification | n2 | M |  |  | Type of identification of passive transport means, coded | T027 | CIS/Control/ControlSubject/Consignment/HouseConsignment/PassiveTransportMeans/identificationType | typeOfIdentification | CL750 | CL750 |
| IE3N09 | 6 | Attribute | 0..1 |                         Nationality | a2 | O |  |  | Nationality of passive means of transport, coded | T028 | CIS/Control/ControlSubject/Consignment/HouseConsignment/PassiveTransportMeans/registrationNationality | nationality | CL718 | CL718 |
| IE3N09 | 5 | Composition | 0..1 |                     Transport document (House level) |  | O |  |  | TransportContractDocument | 30B | CIS/Control/ControlSubject/Consignment/HouseConsignment/TransportContractDocument | transportDocumentHouseLevel |  |  |
| IE3N09 | 6 | Attribute | 1..1 |                         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/Control/ControlSubject/Consignment/HouseConsignment/TransportContractDocument/identification | documentNumber |  |  |
| IE3N09 | 6 | Attribute | 1..1 |                         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/Control/ControlSubject/Consignment/HouseConsignment/TransportContractDocument/type | type | CL754 | CL754 |
| IE3N09 | 5 | Composition | 0..9999 |                     Transport equipment |  | O |  |  | TransportEquipment | 31B | CIS/Control/ControlSubject/Consignment/HouseConsignment/TransportEquipment | transportEquipment |  |  |
| IE3N09 | 6 | Attribute | 1..1 |                         Container identification number | an..17 | M |  |  | Equipment identification number | 159 | CIS/Control/ControlSubject/Consignment/HouseConsignment/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3N09 | 4 | Composition | 0..99 |                 Packaging |  | O |  |  | Packaging | 93A | CIS/Control/ControlSubject/Consignment/Packaging | packaging |  |  |
| IE3N09 | 5 | Attribute | 1..1 |                     Shipping marks | an..512 | M |  |  | Shipping marks | 142 | CIS/Control/ControlSubject/Consignment/Packaging/marksNumbers | shippingMarks |  |  |
| IE3N09 | 5 | Attribute | 1..1 |                     Type of packages | an2 | M |  |  | Type of packages identification, coded | 141 | CIS/Control/ControlSubject/Consignment/Packaging/type | typeOfPackages | CL017 | CL017 |
| IE3N09 | 4 | Composition | 0..9999 |                 Transport equipment |  | O |  |  | TransportEquipment | 31B | CIS/Control/ControlSubject/Consignment/TransportEquipment | transportEquipment |  |  |
| IE3N09 | 5 | Attribute | 1..1 |                     Container identification number | an..17 | M |  |  | Equipment identification number | 159 | CIS/Control/ControlSubject/Consignment/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3N09 | 1 | Composition | 0..1 |     Declarant |  | O |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3N09 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
