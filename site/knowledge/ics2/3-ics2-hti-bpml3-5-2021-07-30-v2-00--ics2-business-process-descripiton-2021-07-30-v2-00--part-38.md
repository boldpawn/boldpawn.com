---
title: "Introduction - part 38"
source_title: "Introduction"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 38
chunk_count: 65
generated: true
---

# Introduction - part 38

Source document: Introduction
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md`
Chunk: 38 of 65

#### Tasks

##### GET01 - Perform syntax and semantic validation

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
| Description: The Arrival notification is not valid and rejected. The person notifying the arrival is notified of the rejection and of the errors that caused the rejection. | Description: The Arrival notification is not valid and rejected. The person notifying the arrival is notified of the rejection and of the errors that caused the rejection. |
| Output (Final situation): The notification of error is generated and sent to the person notifying the arrival and the process ends. | Output (Final situation): The notification of error is generated and sent to the person notifying the arrival and the process ends. |


##### GET02 – Register and generate MRN

| Register and generate MRN | GET02 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry | Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry |
| Input: Arrival notification | Input: Arrival notification |
| Description: Valid Arrival notification is registered, an MRN is generated and allocated to the received Arrival notification. The person notifying the arrival and the notify party, if any, is notified of the registration and the MRN which was allocated to the Arrival notification. | Description: Valid Arrival notification is registered, an MRN is generated and allocated to the received Arrival notification. The person notifying the arrival and the notify party, if any, is notified of the registration and the MRN which was allocated to the Arrival notification. |
| Output (Final situation): The Arrival notification was registered and MRN has been generated. The person notifying the arrival and the notify party, if any, has been notified. The Arrival notification obtains the state 'registered'. | Output (Final situation): The Arrival notification was registered and MRN has been generated. The person notifying the arrival and the notify party, if any, has been notified. The Arrival notification obtains the state 'registered'. |


##### GET05 – Retrieve relevant ENS information and control recommendation

| Retrieve relevant ENS information and control recommendation | GET05 |
| --- | --- |
| Organisation: National Customs Administration of the Member State of actual first entry and/or presentation | Organisation: National Customs Administration of the Member State of actual first entry and/or presentation |
| Input: Registered Arrival notification or Presentation notification | Input: Registered Arrival notification or Presentation notification |
| Description: ENS data related to the given journey of the means of transport for which the arrival has been notified is to be retrieved in an automated process. To match ENS data with the means of transport the following data elements are used depending on the mode of transport: Maritime: The IMO-number and the estimated date and time of arrival at the first place of arrival in the Customs territory of the Union. Air: The flight number and the estimated date and time of arrival at the first place of arrival in the Customs territory of the Union or alternatively by a list of MAWB numbers covered by related ENS data Independently from the mode of transport a list of MRNs allocated to the related ENS. The arrival can only be notified once when the means of transport has arrived at the first port/airport in the EU. The status of the related ENS needs to be in a state which is prior to the arrival of the means of transport (prior to an arrival notification received for the means of transport). The related ENS will obtain the state 'arrived' after this task is executed. If the ENS is not in a correct state it will not be further processed. In case of the split shipments the ENS will obtain the state 'partially arrived'. In case of a Presentation notification the related ENS will be matched by the MRN of the ENS (or the MRN of a partial ENS filing) referred to in the presentation notification. For air and maritime transport the ENS has to be in state 'arrived' or a later state before actions upon presentation can be processed. | Description: ENS data related to the given journey of the means of transport for which the arrival has been notified is to be retrieved in an automated process. To match ENS data with the means of transport the following data elements are used depending on the mode of transport: Maritime: The IMO-number and the estimated date and time of arrival at the first place of arrival in the Customs territory of the Union. Air: The flight number and the estimated date and time of arrival at the first place of arrival in the Customs territory of the Union or alternatively by a list of MAWB numbers covered by related ENS data Independently from the mode of transport a list of MRNs allocated to the related ENS. The arrival can only be notified once when the means of transport has arrived at the first port/airport in the EU. The status of the related ENS needs to be in a state which is prior to the arrival of the means of transport (prior to an arrival notification received for the means of transport). The related ENS will obtain the state 'arrived' after this task is executed. If the ENS is not in a correct state it will not be further processed. In case of the split shipments the ENS will obtain the state 'partially arrived'. In case of a Presentation notification the related ENS will be matched by the MRN of the ENS (or the MRN of a partial ENS filing) referred to in the presentation notification. For air and maritime transport the ENS has to be in state 'arrived' or a later state before actions upon presentation can be processed. |
| Output (Final situation): In case of an arrival notification the ENS obtains the state 'arrived' or 'partially arrived'. In case of Presentation notification the ENS obtains the state 'presented'. | Output (Final situation): In case of an arrival notification the ENS obtains the state 'arrived' or 'partially arrived'. In case of Presentation notification the ENS obtains the state 'presented'. |


##### ANT01 – Notify ENS data is in incorrect state

| Notify ENS data is in incorrect state | ANT01 |
| --- | --- |
| Organisation: National Customs Administration of the Member State of actual first entry | Organisation: National Customs Administration of the Member State of actual first entry |
| Input: Retrieved ENS which is in incorrect state | Input: Retrieved ENS which is in incorrect state |
| Description: A notification to the person notifying the arrival and the notify party, if any, and the person filing (in case of maritime) is generated informing that ENS is in incorrect state and cannot obtain the state 'arrived'. | Description: A notification to the person notifying the arrival and the notify party, if any, and the person filing (in case of maritime) is generated informing that ENS is in incorrect state and cannot obtain the state 'arrived'. |
| Output (Final situation): Person notifying the arrival and the notify party, if any, and person filing (in case of maritime) are notified that ENS in incorrect state and did not obtain the state 'arrived'. | Output (Final situation): Person notifying the arrival and the notify party, if any, and person filing (in case of maritime) are notified that ENS in incorrect state and did not obtain the state 'arrived'. |


##### ANT02 – Stop 200 days timer

| Stop 200 days timer | ANT02 |
| --- | --- |
| Organisation: National Customs Administration of the Member State of actual first entry | Organisation: National Customs Administration of the Member State of actual first entry |
| Input: Retrieved ENS which is in correct state | Input: Retrieved ENS which is in correct state |
| Description: The 200 days timer is stopped. | Description: The 200 days timer is stopped. |
| Output (Final situation): The 200 days timer is stopped. | Output (Final situation): The 200 days timer is stopped. |


##### GET06 – Set control decision

| Issue control decision | GET06 |
| --- | --- |
| Organisation: National Customs Administration of Member State of first entry or presentation or control (at final destination) | Organisation: National Customs Administration of Member State of first entry or presentation or control (at final destination) |
| Input: Received control recommendation from RMS | Input: Received control recommendation from RMS |
| Description The Member State of first entry or presentation or control (at final destination) decides on whether a control is performed or not and if a control is performed on the type of control taking into account the recommendation received from RMS. | Description The Member State of first entry or presentation or control (at final destination) decides on whether a control is performed or not and if a control is performed on the type of control taking into account the recommendation received from RMS. |
| Output (Final situation): Control decision is issued. | Output (Final situation): Control decision is issued. |


##### GET07 – Notify control decision

| Notify control decision | GET07 |
| --- | --- |
| Organisation: National Customs Administration of the Member State of actual first entry or presentation or control (at final destination) | Organisation: National Customs Administration of the Member State of actual first entry or presentation or control (at final destination) |
| Input: Control recommendation from RMS or own e-Screening results | Input: Control recommendation from RMS or own e-Screening results |
| Description: Control decisions for the identified items are notified to the person notifying the arrival, the notify party, if any, and the person filing (in case of maritime). In case of arrival this will force the unloading of the selected items to customs in order to make them available to customs for physical controls even when the unloading of those items at this place was originally not foreseen by that person. | Description: Control decisions for the identified items are notified to the person notifying the arrival, the notify party, if any, and the person filing (in case of maritime). In case of arrival this will force the unloading of the selected items to customs in order to make them available to customs for physical controls even when the unloading of those items at this place was originally not foreseen by that person. |
| Output (Final situation): The control decision for items which require a control upon entry is notified to the person notifying the arrival, the notify party, if any, and the person filing (in case of maritime). | Output (Final situation): The control decision for items which require a control upon entry is notified to the person notifying the arrival, the notify party, if any, and the person filing (in case of maritime). |
