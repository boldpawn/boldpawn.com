---
title: "Introduction - part 49"
source_title: "Introduction"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 49
chunk_count: 65
generated: true
---

# Introduction - part 49

Source document: Introduction
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md`
Chunk: 49 of 65

#### Tasks

##### GET01 – Perform syntax and semantic validation

| Perform syntax and semantic validation | GET01 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry | Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry |
| Input: ENS filing, Arrival notification, Amendment of ENS filing, Invalidation request | Input: ENS filing, Arrival notification, Amendment of ENS filing, Invalidation request |
| Description: The input file is validated in an automated process. The validation comprises validation of message syntax, checks of the compliance with business rules and conditions as well as reference data. If the input file is not valid it is rejected; If the input file is valid it is further processed. | Description: The input file is validated in an automated process. The validation comprises validation of message syntax, checks of the compliance with business rules and conditions as well as reference data. If the input file is not valid it is rejected; If the input file is valid it is further processed. |
| Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated the input file will obtain the state 'rejected' | Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated the input file will obtain the state 'rejected' |


##### GET03 – Notify error

| Notify error | GET03 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry | Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry |
| Input: Unsuccessful validation results – containing the error description | Input: Unsuccessful validation results – containing the error description |
| Description: The file is not valid and rejected. The person filing is notified of the rejection and of the errors that caused the rejection. | Description: The file is not valid and rejected. The person filing is notified of the rejection and of the errors that caused the rejection. |
| Output (Final situation): The notification of error is generated and sent to the person filing and the process ends. | Output (Final situation): The notification of error is generated and sent to the person filing and the process ends. |


##### IFT03 – Send invalidation request reception acknowledgement

| Send invalidation request reception acknowledgement | AFT01 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: Invalidation request | Input: Invalidation request |
| Description: A notification is generated and sent to the person filing. This informs about successful syntax and semantic validation. | Description: A notification is generated and sent to the person filing. This informs about successful syntax and semantic validation. |
| Output (Final situation): The person filing is notified that its invalidation request passed syntax and semantic validation successfully. | Output (Final situation): The person filing is notified that its invalidation request passed syntax and semantic validation successfully. |


##### GET04 – Perform ENS-lifecycle validation

| Perform ENS lifecycle validation | GET04 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: ENS filing, Amendment of ENS filing, or Invalidation request | Input: ENS filing, Amendment of ENS filing, or Invalidation request |
| Description The ENS lifecycle validation checks Whether the ENS filing is in the correct state regarding the ENS lifecycle’s sequence of the process (E.g. whether an ENS filing refers to a journey of a means of transport which arrived already in the past, or the goods covered by the ENS filing were already presented to the customs and effectively the ENS cannot be filed for the same goods already presented, or ENS cannot be amended or invalidated) and; Whether relevant key data; in particular data elements which are used as a unique linking key are unique. In the case of indicated split shipments (i.e. ENS filing containing defined split shipment code indicator) where unique linking key data can be re-used for the remaining goods after the shipment was split, the ENS filing will be processed. For an ENS amendment; That the state of the original ENS filing is not 'presented' or at a state where, for example amendment of the ENS filing is not possible anymore (i.e. customs decided to perform controls, or trader notified by the customs that goods are to be controlled). That no particulars were amended which are not allowed to be amended. The person filing the amendment needs to be the declarant or the representative of the ENS filing referred to. For an invalidation request; The related ENS filing is required to be in a state prior to the state 'presented'. The person filing the invalidation request needs to be the declarant or the representative of the ENS filing referred to. If the ENS lifecycle validation is successful, the input file obtains the state 'accepted' and the sub-process continues. If the ENS lifecycle validation fails the input file obtains the state 'rejected' and the sub-process ends. | Description The ENS lifecycle validation checks Whether the ENS filing is in the correct state regarding the ENS lifecycle’s sequence of the process (E.g. whether an ENS filing refers to a journey of a means of transport which arrived already in the past, or the goods covered by the ENS filing were already presented to the customs and effectively the ENS cannot be filed for the same goods already presented, or ENS cannot be amended or invalidated) and; Whether relevant key data; in particular data elements which are used as a unique linking key are unique. In the case of indicated split shipments (i.e. ENS filing containing defined split shipment code indicator) where unique linking key data can be re-used for the remaining goods after the shipment was split, the ENS filing will be processed. For an ENS amendment; That the state of the original ENS filing is not 'presented' or at a state where, for example amendment of the ENS filing is not possible anymore (i.e. customs decided to perform controls, or trader notified by the customs that goods are to be controlled). That no particulars were amended which are not allowed to be amended. The person filing the amendment needs to be the declarant or the representative of the ENS filing referred to. For an invalidation request; The related ENS filing is required to be in a state prior to the state 'presented'. The person filing the invalidation request needs to be the declarant or the representative of the ENS filing referred to. If the ENS lifecycle validation is successful, the input file obtains the state 'accepted' and the sub-process continues. If the ENS lifecycle validation fails the input file obtains the state 'rejected' and the sub-process ends. |
| Output (Final situation): The state of the input file will either be changed to 'accepted' or 'rejected'. | Output (Final situation): The state of the input file will either be changed to 'accepted' or 'rejected'. |


##### IFT01 – Notify invalidation request acceptance

| Notify invalidation request acceptance | IFT01 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: Invalidation request | Input: Invalidation request |
| Description: Notification of acceptance of the invalidation request is generated and sent to the person filing and the process continues. | Description: Notification of acceptance of the invalidation request is generated and sent to the person filing and the process continues. |
| Output (Final situation): The person filing is notified of the acceptance of the invalidation request. | Output (Final situation): The person filing is notified of the acceptance of the invalidation request. |


##### IFT02 – Invalidate filing

| Invalidate filing | IFT02 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: Invalidation request or 200 days timer expiry. | Input: Invalidation request or 200 days timer expiry. |
| Description: The ENS filing for which the 200 days timer expired or the ENS filing referred to in the invalidation request will be invalidated by assigning the status 'invalidated'. | Description: The ENS filing for which the 200 days timer expired or the ENS filing referred to in the invalidation request will be invalidated by assigning the status 'invalidated'. |
| Output (Final situation): The ENS filing is no longer processed and will be moved to an archive. | Output (Final situation): The ENS filing is no longer processed and will be moved to an archive. |
