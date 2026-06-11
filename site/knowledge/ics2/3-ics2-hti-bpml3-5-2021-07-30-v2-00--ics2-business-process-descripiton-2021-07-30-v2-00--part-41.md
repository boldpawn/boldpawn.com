---
title: "Introduction - part 41"
source_title: "Introduction"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md"
section: "Business Process Model"
chunk: 41
chunk_count: 65
generated: true
---

# Introduction - part 41

Source document: Introduction
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md`
Chunk: 41 of 65

#### Business Process Model


*Figure 7: HLS005 Process presentation and controls sub-process*

##### GET05 – Retrieve relevant ENS information and control recommendation

| Retrieve relevant ENS information and control recommendation | GET05 |
| --- | --- |
| Organisation: National Customs Administration of the Member State of actual first entry and/or presentation | Organisation: National Customs Administration of the Member State of actual first entry and/or presentation |
| Input: Registered Arrival notification or Presentation notification | Input: Registered Arrival notification or Presentation notification |
| Description: ENS data related to the given journey of the means of transport for which the arrival has been notified is to be retrieved in an automated process. To match ENS data with the means of transport the following data elements are used depending on the mode of transport: Maritime: The IMO-number and the estimated date and time of arrival at the first place of arrival in the Customs territory of the Union. Air: The flight number and the estimated date and time of arrival at the first place of arrival in the Customs territory of the Union or alternatively by a list of MAWB numbers covered by related ENS data Independently from the mode of transport a list of MRNs allocated to the related ENS. The arrival can only be notified once when the means of transport has arrived at the first port/airport in the EU. The status of the related ENS needs to be in a state which is prior to the arrival of the means of transport (prior to an arrival notification received for the means of transport). The related ENS will obtain the state 'arrived' after this task is executed. If the ENS is not in a correct state it will not be further processed. In case of the split shipments the ENS will obtain the state 'partially arrived'. In case of a Presentation notification the related ENS will be matched by the MRN of the ENS (or the MRN of a partial ENS filing) referred to in the presentation notification. For air and maritime transport the ENS has to be in state 'arrived' or a later state before actions upon presentation can be processed. | Description: ENS data related to the given journey of the means of transport for which the arrival has been notified is to be retrieved in an automated process. To match ENS data with the means of transport the following data elements are used depending on the mode of transport: Maritime: The IMO-number and the estimated date and time of arrival at the first place of arrival in the Customs territory of the Union. Air: The flight number and the estimated date and time of arrival at the first place of arrival in the Customs territory of the Union or alternatively by a list of MAWB numbers covered by related ENS data Independently from the mode of transport a list of MRNs allocated to the related ENS. The arrival can only be notified once when the means of transport has arrived at the first port/airport in the EU. The status of the related ENS needs to be in a state which is prior to the arrival of the means of transport (prior to an arrival notification received for the means of transport). The related ENS will obtain the state 'arrived' after this task is executed. If the ENS is not in a correct state it will not be further processed. In case of the split shipments the ENS will obtain the state 'partially arrived'. In case of a Presentation notification the related ENS will be matched by the MRN of the ENS (or the MRN of a partial ENS filing) referred to in the presentation notification. For air and maritime transport the ENS has to be in state 'arrived' or a later state before actions upon presentation can be processed. |
| Output (Final situation): In case of an arrival notification the ENS obtains the state 'arrived' or 'partially arrived'. In case of Presentation notification the ENS obtains the state 'presented'. | Output (Final situation): In case of an arrival notification the ENS obtains the state 'arrived' or 'partially arrived'. In case of Presentation notification the ENS obtains the state 'presented'. |


##### PCT04 – Stop 200 days timer

| Stop 200 days timer | PCT04 |
| --- | --- |
| Organisation: National Customs Administration of the Member State of presentation | Organisation: National Customs Administration of the Member State of presentation |
| Input: Retrieved ENS which is in correct state | Input: Retrieved ENS which is in correct state |
| Description: The 200 days timer is stopped. | Description: The 200 days timer is stopped. |
| Output (Final situation): The 200 days timer is stopped. | Output (Final situation): The 200 days timer is stopped. |


##### GET06 – Set control decision

| Issue control decision | GET06 |
| --- | --- |
| Organisation: National Customs Administration of Member State of first entry or presentation or control (at final destination) | Organisation: National Customs Administration of Member State of first entry or presentation or control (at final destination) |
| Input: Received control recommendation from RMS | Input: Received control recommendation from RMS |
| Description The Member State of first entry or presentation or control (at final destination) decides on whether a control is performed or not and if a control is performed on the type of control taking into account the recommendation received from RMS. | Description The Member State of first entry or presentation or control (at final destination) decides on whether a control is performed or not and if a control is performed on the type of control taking into account the recommendation received from RMS. |
| Output (Final situation): Control is issued. | Output (Final situation): Control is issued. |


##### GET07 – Notify control decision

| Notify control decision | GET07 |
| --- | --- |
| Organisation: National Customs Administration of the Member State of actual first entry or presentation or control (at final destination) | Organisation: National Customs Administration of the Member State of actual first entry or presentation or control (at final destination) |
| Input: Control recommendation from RMS or own e-Screening results | Input: Control recommendation from RMS or own e-Screening results |
| Description: Control decisions for the identified items are notified to the person filing the arrival or presentation notification. In case of arrival this will force the presentation of the selected items to customs in order to make them available for physical controls even when the presentation of those items at this place was originally not foreseen by that person. | Description: Control decisions for the identified items are notified to the person filing the arrival or presentation notification. In case of arrival this will force the presentation of the selected items to customs in order to make them available for physical controls even when the presentation of those items at this place was originally not foreseen by that person. |
| Output (Final situation): The control decision for items which require a control upon entry is notified to the person filing. | Output (Final situation): The control decision for items which require a control upon entry is notified to the person filing. |


##### PCT01 – Perform controls

| Perform controls | PCT01 |
| --- | --- |
| Organisation: National Customs Administration of the Member State of presentation or control | Organisation: National Customs Administration of the Member State of presentation or control |
| Input: Issued control decision | Input: Issued control decision |
| Description: Controls are performed as decided in task GET06. | Description: Controls are performed as decided in task GET06. |
| Output (Final situation): The controls were performed as decided. | Output (Final situation): The controls were performed as decided. |


##### PCT02 – Document control results

| Document control results | PCT02 |
| --- | --- |
| Organisation: National Customs Administration of the Member State of presentation or control | Organisation: National Customs Administration of the Member State of presentation or control |
| Input: Control results | Input: Control results |
| Description: The results of the performed controls are documented. | Description: The results of the performed controls are documented. |
| Output (Final situation): The control results are documented. | Output (Final situation): The control results are documented. |


##### PCT03 – Communicate control recommendation to final destination

| Communicate control recommendation to final destination | PCT03 |
| --- | --- |
| Organisation: National Customs Administration of the Member State of control | Organisation: National Customs Administration of the Member State of control |
| Input: Issued control decision | Input: Issued control decision |
| Description: The control recommendation is communicated to the Member State of control (at final destination) when different from the Member State where the goods were presented first after their entry into the customs territory of the European Union. This provides the opportunity to perform the controls at the most appropriate place and not necessarily at the border customs office. | Description: The control recommendation is communicated to the Member State of control (at final destination) when different from the Member State where the goods were presented first after their entry into the customs territory of the European Union. This provides the opportunity to perform the controls at the most appropriate place and not necessarily at the border customs office. |
| Output (Final situation): The issued control recommendation was communicated to the Member State of control (at final destination). | Output (Final situation): The issued control recommendation was communicated to the Member State of control (at final destination). |
