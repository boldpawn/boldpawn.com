---
title: "Introduction - part 70"
source_title: "Introduction"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 70
chunk_count: 82
generated: true
---

# Introduction - part 70

Source document: Introduction
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md`
Chunk: 70 of 82

#### Tasks

##### T-13-01 – Submit ENS consultation request

| Submit ENS consultation request | T-13-01 |
| --- | --- |
| System component: TI | System component: TI |
| Input: ENS consultation request received | Input: ENS consultation request received |
| Description: The TI will submit the ENS consultation request to the CR via the message IE4Q09 [Msg: C_ENS_CNS]. | Description: The TI will submit the ENS consultation request to the CR via the message IE4Q09 [Msg: C_ENS_CNS]. |
| Output (Final situation): The ENS consultation request was submitted to the CR | Output (Final situation): The ENS consultation request was submitted to the CR |


##### T-13-02 – Send ENS consultation results

| Send ENS consultation results | T-13-02 |
| --- | --- |
| System component: TI | System component: TI |
| Input: ENS consultation results received from CR | Input: ENS consultation results received from CR |
| Description: The TI will send the ENS consultation results to the party that requested the consultation, via the message IE3R08 [Msg: E_ENS_CNS_RES]. If the Economic operator requesting the ENS consultation has indicated as “True” the “Request notifications” data, then the following list of notifications related to the requested MRN/Master level transport document/House level transport document needs to be provided (in case available). Based on the ENS Consultation request information (MRN, Master level transport document or House level transport document) the CR identifies the related entities and forwards the related identifiers to the STI (IE4R06). The STI retrieves the following stored notifications and forwards them to the EO – ENS lifecycle validation error notification, ENS not complete notification, Assessment complete notification, Additional information request notification, HRCM screening request notification, Control notification, AEO control notification (if authorized), Amendment notification, DNL, Additional information request, HRCM screening request and Invalidation acceptance response. When an MRN is provided in the ENS Consultation the CR identifies the ENS Filing to which the MRN was generated. The CR forwards to the STI the related identification information for the STI to retrieve the relevant notifications to the MRN (and the related ENS entities) for which the EO sending the ENS Consultation is the recipient. When a Master level or House level Transport document reference number is provided in the ENS Consultation the CR identifies the ENS Filing via which the Transport document was registered. The CR forwards to the STI the related identification information for the STI to retrieve the relevant notifications to the Master level or House level transport document (and the related ENS entities) for which the EO sending the ENS Consultation is the recipient. The retrieved notifications (and the related ENS entities) per House consignment, Master consignment or ENS Filing are forwarded to the requesting EO in accordance with its data access rights. When the consultation is for a House consignment, the notifications (and the related ENS entities) only related to the House consignment should be retrieved and forwarded (not the notifications in general related to the ENS to which this House consignments belongs). When the consultation is for a Master consignment, the notifications (and the related ENS entities) only related to the Master ENS Filing are retrieved and forwarded (not the notifications related to the linked House consignments). | Description: The TI will send the ENS consultation results to the party that requested the consultation, via the message IE3R08 [Msg: E_ENS_CNS_RES]. If the Economic operator requesting the ENS consultation has indicated as “True” the “Request notifications” data, then the following list of notifications related to the requested MRN/Master level transport document/House level transport document needs to be provided (in case available). Based on the ENS Consultation request information (MRN, Master level transport document or House level transport document) the CR identifies the related entities and forwards the related identifiers to the STI (IE4R06). The STI retrieves the following stored notifications and forwards them to the EO – ENS lifecycle validation error notification, ENS not complete notification, Assessment complete notification, Additional information request notification, HRCM screening request notification, Control notification, AEO control notification (if authorized), Amendment notification, DNL, Additional information request, HRCM screening request and Invalidation acceptance response. When an MRN is provided in the ENS Consultation the CR identifies the ENS Filing to which the MRN was generated. The CR forwards to the STI the related identification information for the STI to retrieve the relevant notifications to the MRN (and the related ENS entities) for which the EO sending the ENS Consultation is the recipient. When a Master level or House level Transport document reference number is provided in the ENS Consultation the CR identifies the ENS Filing via which the Transport document was registered. The CR forwards to the STI the related identification information for the STI to retrieve the relevant notifications to the Master level or House level transport document (and the related ENS entities) for which the EO sending the ENS Consultation is the recipient. The retrieved notifications (and the related ENS entities) per House consignment, Master consignment or ENS Filing are forwarded to the requesting EO in accordance with its data access rights. When the consultation is for a House consignment, the notifications (and the related ENS entities) only related to the House consignment should be retrieved and forwarded (not the notifications in general related to the ENS to which this House consignments belongs). When the consultation is for a Master consignment, the notifications (and the related ENS entities) only related to the Master ENS Filing are retrieved and forwarded (not the notifications related to the linked House consignments). |
| Output (Final situation): The ENS consultation results were sent to the party that requested the consultation. | Output (Final situation): The ENS consultation results were sent to the party that requested the consultation. |


##### T-13-03 – Identify relevant ENS information to be shared with the requesting party

| Identify relevant ENS information to be shared with the requesting party | T-13-03 |
| --- | --- |
| System component: CR | System component: CR |
| Input: Received ENS consultation request | Input: Received ENS consultation request |
| Description: The CR identifies the relevant ENS information that matches the inserted criteria and at the same time complies with the access rights that the party requesting the consultation has. | Description: The CR identifies the relevant ENS information that matches the inserted criteria and at the same time complies with the access rights that the party requesting the consultation has. |
| Output (Final situation): The relevant ENS information was identified. | Output (Final situation): The relevant ENS information was identified. |


##### T-13-04 – Send ENS consultation results to TI

| Send ENS consultation results to TI | T-13-04 |
| --- | --- |
| System component: CR | System component: CR |
| Input: Identified relevant ENS information | Input: Identified relevant ENS information |
| Description: The CR sends the identified relevant ENS information to the TI, via the message IE4R06 [Msg: C_ENS_CNS_RES]. | Description: The CR sends the identified relevant ENS information to the TI, via the message IE4R06 [Msg: C_ENS_CNS_RES]. |
| Output (Final situation): The ENS consultation results were sent to the TI. | Output (Final situation): The ENS consultation results were sent to the TI. |
