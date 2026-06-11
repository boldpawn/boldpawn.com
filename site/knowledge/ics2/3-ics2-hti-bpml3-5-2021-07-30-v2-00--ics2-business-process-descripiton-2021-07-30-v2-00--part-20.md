---
title: "Introduction - part 20"
source_title: "Introduction"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 20
chunk_count: 65
generated: true
---

# Introduction - part 20

Source document: Introduction
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md`
Chunk: 20 of 65

#### Tasks

##### GET01 - Perform syntax and semantic validation

| Perform syntax and semantic validation | GET01 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry | Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry |
| Input: ENS filing, Arrival notification, Amendment of ENS filing, Invalidation request | Input: ENS filing, Arrival notification, Amendment of ENS filing, Invalidation request |
| Description: The input file is validated in an automated process. The validation comprises validation of message syntax, checks of the compliance with business rules and conditions as well as reference data. If the input file is not valid it is rejected; If the input file is valid it is further processed. | Description: The input file is validated in an automated process. The validation comprises validation of message syntax, checks of the compliance with business rules and conditions as well as reference data. If the input file is not valid it is rejected; If the input file is valid it is further processed. |
| Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated the input file will obtain the state 'rejected' | Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated the input file will obtain the state 'rejected' |


##### GET03 - Notify error

| Notify error | GET03 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: Unsuccessful validation results - containing the error description | Input: Unsuccessful validation results - containing the error description |
| Description: The file is not valid and rejected. The person filing is notified of the rejection and of the errors that caused the rejection. | Description: The file is not valid and rejected. The person filing is notified of the rejection and of the errors that caused the rejection. |
| Output (Final situation): The notification of error is generated and sent to the person filing and the process ends. | Output (Final situation): The notification of error is generated and sent to the person filing and the process ends. |


##### GET02 - Register and generate MRN

| Register and generate MRN | GET02 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry | Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry |
| Input: Valid ENS filing or Arrival notification | Input: Valid ENS filing or Arrival notification |
| Description: Valid file is registered; an MRN is generated and allocated to the received filing. In the case of received filing F25 (reference to the MAWB containing the list of underlying HAWB(s)), no additional MRN will be generated and the already generated MRN(s) for the previously received and related filing(s) F23 will be used. In the case of ENS filing containing minimum data-set for air cargo pre-loading processes, besides the above conditions that need to be met for notifying the carrier, the air carrier also needs to be identifiable from the data contained in the ENS filing: MAWB reference number. | Description: Valid file is registered; an MRN is generated and allocated to the received filing. In the case of received filing F25 (reference to the MAWB containing the list of underlying HAWB(s)), no additional MRN will be generated and the already generated MRN(s) for the previously received and related filing(s) F23 will be used. In the case of ENS filing containing minimum data-set for air cargo pre-loading processes, besides the above conditions that need to be met for notifying the carrier, the air carrier also needs to be identifiable from the data contained in the ENS filing: MAWB reference number. |
| Output (Final situation): The file was registered and MRN has been generated. The person filing and the carrier (where applicable) has been notified. The file obtains the state 'registered'. | Output (Final situation): The file was registered and MRN has been generated. The person filing and the carrier (where applicable) has been notified. The file obtains the state 'registered'. |


##### RFT01 - Start 200 days timer

| Start 200 days timer | RFT01 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: Registered ENS filing | Input: Registered ENS filing |
| Description: The 200 days timer is started for the registered ENS filing. When the goods declared in the ENS filing do not arrive within 200 days the filing will be invalidated (legally not deemed being lodged). | Description: The 200 days timer is started for the registered ENS filing. When the goods declared in the ENS filing do not arrive within 200 days the filing will be invalidated (legally not deemed being lodged). |
| Output (Final situation): 200 days countdown started. | Output (Final situation): 200 days countdown started. |
