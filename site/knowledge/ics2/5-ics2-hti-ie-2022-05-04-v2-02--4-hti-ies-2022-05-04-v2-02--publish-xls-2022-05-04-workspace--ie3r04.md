---
title: "IE3R04"
source_title: "IE3R04"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3R04.md"
section: "IE3R04"
chunk: 1
chunk_count: 1
generated: true
---

# IE3R04

Source document: IE3R04
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3R04.md`
Chunk: 1 of 1

# IE3R04

| Message | Level | Type | Occurs | ICS2 Name | ICS2 Format | Status | Rules | Conditions | WCO Name | WCO Id | XPath | ICS2 XML Name | ICS2 CL (MS) | ICS2 CL (Trade) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| IE3R04 | 0 | Class |  | IE3R04 |  |  |  |  | CIS | N.A. | CIS | IE3R04 |  |  |
| IE3R04 | 1 | Attribute | 1..1 |     LRN | an..22 | M |  |  | Functional reference number | D026 | CIS/functionalReference | LRN |  |  |
| IE3R04 | 1 | Attribute | 1..1 |     MRN | an18 | M |  |  |  |  | CIS/masterReferenceNumber | MRN |  |  |
| IE3R04 | 1 | Attribute | 1..1 |     Registration date | an20 (YYYY-MM-DDThh:mm:ssZ) | M |  |  |  |  | CIS/registrationDate | registrationDate |  |  |
| IE3R04 | 1 | Composition | 0..1 |     Notify party |  | O |  |  | Agent | 05A | CIS/Agent | notifyParty |  |  |
| IE3R04 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Agent, coded | R004 | CIS/Agent/identification | identificationNumber |  |  |
| IE3R04 | 1 | Composition | 0..1 |     Person notifying the arrival |  | O |  |  | Submitter | 17B | CIS/PersonNotifyingArrival | personNotifyingTheArrival |  |  |
| IE3R04 | 2 | Attribute | 1..1 |         Identification number | an..17 | M |  |  | Submitter, coded | R059 | CIS/PersonNotifyingArrival/identification | identificationNumber |  |  |
| IE3R04 | 1 | Composition | 1..1 |     Customs office of first entry |  | M |  |  | EntryOffice | 51A | CIS/EntryOffice | customsOfficeOfFirstEntry |  |  |
| IE3R04 | 2 | Attribute | 1..1 |         Reference number | an8 | M |  |  | Office of entry, coded | G004 | CIS/EntryOffice/identification | referenceNumber | CL141 | CL141 |
