---
title: "IE3N03"
source_title: "IE3N03"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N03.md"
section: "IE3N03"
chunk: 1
chunk_count: 1
generated: true
---

# IE3N03

Source document: IE3N03
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3N03.md`
Chunk: 1 of 1

# IE3N03

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3N03 | 0 | Class |  | IE3N03 |  |  |  |  | CIS | N.A. | CIS | IE3N03 |  |  |
| IE3N03 | 1 | Attribute | 1..1 |     MRN | an18 | M |  |  |  |  | CIS/masterReferenceNumber | MRN |  |  |
| IE3N03 | 1 | Attribute | 1..1 |     Completion date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  |  |  | CIS/completionDate | completionDate |  |  |
| IE3N03 | 1 | Composition | 0..1 |     Assessment complete |  | O |  |  |  |  | CIS/AssessmentComplete | assessmentComplete |  |  |
| IE3N03 | 2 | Composition | 0..999 |         Receptacle |  | O |  |  |  |  | CIS/AssessmentComplete/PostalReceptacle | receptacle |  |  |
| IE3N03 | 3 | Attribute | 1..1 |             Receptacle identification number | an..35 | M |  |  |  |  | CIS/AssessmentComplete/PostalReceptacle/receptacleIdentificationNumber | receptacleIdentificationNumber |  |  |
| IE3N03 | 2 | Composition | 0..99999 |         Transport document (House level) |  | O |  |  | TransportContractDocument | 30B | CIS/AssessmentComplete/TransportContractDocument | transportDocument |  |  |
| IE3N03 | 3 | Attribute | 1..1 |             Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/AssessmentComplete/TransportContractDocument/identification | documentNumber |  |  |
| IE3N03 | 3 | Attribute | 1..1 |             Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/AssessmentComplete/TransportContractDocument/type | type | CL754 | CL754 |
| IE3N03 | 2 | Composition | 0..9999 |         Transport equipment |  | O |  |  | TransportEquipment | 31B | CIS/AssessmentComplete/TransportEquipment | transportEquipment |  |  |
| IE3N03 | 3 | Attribute | 1..1 |             Container identification number | an..17 | M |  |  | Equipment identification number | 159 | CIS/AssessmentComplete/TransportEquipment/identification | containerIdentificationNumber |  |  |
| IE3N03 | 1 | Composition | 1..1 |     Responsible Member State |  | M |  |  | ResponsibleMemberState | SC2 | CIS/ResponsibleMemberState | responsibleMemberState |  |  |
| IE3N03 | 2 | Attribute | 1..1 |         Country | a2 | M |  |  |  |  | CIS/ResponsibleMemberState/country | country | CL717 | CL717 |
| IE3N03 | 1 | Composition | 0..1 |     Representative |  | O |  |  | Agent | 05A | CIS/Agent | representative |  |  |
| IE3N03 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Agent, coded | R004 | CIS/Agent/identification | identificationNumber |  |  |
| IE3N03 | 1 | Composition | 0..1 |     Transport document (Master level) |  | O |  |  | TransportContractDocument | 30B | CIS/TransportContractDocument | transportDocument |  |  |
| IE3N03 | 2 | Attribute | 1..1 |         Reference number | an..70 | M |  |  | Transport document number | D023 | CIS/TransportContractDocument/identification | documentNumber |  |  |
| IE3N03 | 2 | Attribute | 1..1 |         Type | an4 | M |  |  | Transport document type, coded | D024 | CIS/TransportContractDocument/type | type | CL754 | CL754 |
| IE3N03 | 1 | Composition | 0..1 |     Carrier |  | O |  |  | Carrier | 18A | CIS/Carrier | carrier |  |  |
| IE3N03 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Carrier identification | R012 | CIS/Carrier/identification | identificationNumber |  |  |
| IE3N03 | 1 | Composition | 0..1 |     Declarant |  | O |  |  | Declarant | 57B | CIS/Declarant | declarant |  |  |
| IE3N03 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
