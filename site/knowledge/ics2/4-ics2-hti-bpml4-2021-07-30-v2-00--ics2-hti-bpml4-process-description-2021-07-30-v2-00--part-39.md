---
title: "Introduction - part 39"
source_title: "Introduction"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 39
chunk_count: 82
generated: true
---

# Introduction - part 39

Source document: Introduction
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md`
Chunk: 39 of 82

#### Tasks

##### T-04-01 - Perform syntactical and semantical validation on received arrival notification

| Perform syntactical and semantical validation on received arrival notification | T-04-01 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Arrival notification | Input: Arrival notification |
| Description: The arrival notification is validated in an automated process. It comprises the validation against the message rules and conditions, as well as the relevant reference data which are valid at the date of the reception. The messages rules and conditions and relevant reference data are defined in [R02] ICS2 Information Exchange Messages. | Description: The arrival notification is validated in an automated process. It comprises the validation against the message rules and conditions, as well as the relevant reference data which are valid at the date of the reception. The messages rules and conditions and relevant reference data are defined in [R02] ICS2 Information Exchange Messages. |
| Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated the input arrival notification will obtain the state 'Rejected' | Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated the input arrival notification will obtain the state 'Rejected' |


##### T-04-02 – Notify error

| Notify error | T-04-02 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Unsuccessful validation results | Input: Unsuccessful validation results |
| Description: The file is not valid and is rejected. The person filing is notified of the rejection and of the errors that caused the rejection via the message IE3N99 [Msg: E_ERR_NOT]. | Description: The file is not valid and is rejected. The person filing is notified of the rejection and of the errors that caused the rejection via the message IE3N99 [Msg: E_ERR_NOT]. |
| Output (Final situation): The notification of error is generated and sent to the person filing. | Output (Final situation): The notification of error is generated and sent to the person filing. |


##### T-04-03 – Register arrival notification and assign MRN

| Register arrival notification and assign MRN | T-04-03 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Valid arrival notification | Input: Valid arrival notification |
| Description: The valid arrival notification is registered; an MRN is generated and allocated to the arrival notification. The arrival notification obtains the state 'Registered'. | Description: The valid arrival notification is registered; an MRN is generated and allocated to the arrival notification. The arrival notification obtains the state 'Registered'. |
| Output (Final situation): The arrival notification was registered and MRN has been generated. | Output (Final situation): The arrival notification was registered and MRN has been generated. |


##### T-04-04 – Submit arrival notification

| Submit arrival notification | T-04-04 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Registered arrival notification from Task: T-04-03 – Register arrival notification and assign MRN | Input: Registered arrival notification from Task: T-04-03 – Register arrival notification and assign MRN |
| Description: The registered arrival notification is submitted to the CR via the message IE4N07 [Msg: C_ARV_NOT]. | Description: The registered arrival notification is submitted to the CR via the message IE4N07 [Msg: C_ARV_NOT]. |
| Output (Final situation): The arrival notification is available in the CR. | Output (Final situation): The arrival notification is available in the CR. |


##### T-04-05 – Notify successful registration and MRN to Person filing

| Notify successful registration and MRN to Person filing | T-04-05 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Registered arrival notification from Task: T-04-03 – Register arrival notification and assign MRN | Input: Registered arrival notification from Task: T-04-03 – Register arrival notification and assign MRN |
| Description: The registration and assigned MRN of the arrival notification is notified to the Person filing via the message IE3R04 [Msg: E_ARV_REG_RSP]. | Description: The registration and assigned MRN of the arrival notification is notified to the Person filing via the message IE3R04 [Msg: E_ARV_REG_RSP]. |
| Output (Final situation): The Person filing is notified of the successful registration of the arrival notification and the MRN assigned to it. | Output (Final situation): The Person filing is notified of the successful registration of the arrival notification and the MRN assigned to it. |


##### T-04-06 – Notify House consignment in incorrect state to Person filing

| Notify House consignment in incorrect state to Person filing | T-04-06 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Incorrect state notification received from CR | Input: Incorrect state notification received from CR |
| Description: The message IE3N07 [Msg: E_HCS_INC_NOT] is sent to the person filing informing him that ENS is in incorrect state and its House consignments cannot obtain the state 'Arrived'. | Description: The message IE3N07 [Msg: E_HCS_INC_NOT] is sent to the person filing informing him that ENS is in incorrect state and its House consignments cannot obtain the state 'Arrived'. |
| Output (Final situation): Person filing is notified that ENS in incorrect state and its House consignments did not obtain the state 'Arrived'. | Output (Final situation): Person filing is notified that ENS in incorrect state and its House consignments did not obtain the state 'Arrived'. |


##### T-04-07 – Notify control to Person filing

| Notify control to Person filing | T-04-07 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Control notification received from CR | Input: Control notification received from CR |
| Description: Control notification for the identified items is sent to the Person filing via the message IE3N08 [Msg: E_CON_NOT]. The Person filing will be forced to unload the selected goods in order to make them available to customs for controls, regardless of whether the presentation of those goods was originally foreseen at this place or not. | Description: Control notification for the identified items is sent to the Person filing via the message IE3N08 [Msg: E_CON_NOT]. The Person filing will be forced to unload the selected goods in order to make them available to customs for controls, regardless of whether the presentation of those goods was originally foreseen at this place or not. |
| Output (Final situation): The control decision for items which require a control upon entry is notified to Person filing. | Output (Final situation): The control decision for items which require a control upon entry is notified to Person filing. |


##### T-04-26 – Notify successful registration and MRN to notify party

| Notify successful registration and MRN to notify party | T-04-26 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Outcome of Task: T-04-05 – Notify successful registration and MRN to Person filing | Input: Outcome of Task: T-04-05 – Notify successful registration and MRN to Person filing |
| Description: The registration and assigned MRN of the arrival notification is notified to notify party via the message IE3R04 [Msg: E_ARV_REG_RSP]. | Description: The registration and assigned MRN of the arrival notification is notified to notify party via the message IE3R04 [Msg: E_ARV_REG_RSP]. |
| Output (Final situation): The notify party is notified of the successful registration of the arrival notification and the MRN assigned to it. | Output (Final situation): The notify party is notified of the successful registration of the arrival notification and the MRN assigned to it. |


##### T-04-27 – Notify control to notify party

| Notify control to notify party | T-04-27 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Control notification received from CR | Input: Control notification received from CR |
| Description: Control notification for the identified items is sent to notify party via the message IE3N08 [Msg: E_CON_NOT]. The notify party will be forced to organise to unload and present the selected goods in order to make them available to customs for controls, regardless of whether the presentation of those goods was originally foreseen at this place or not. | Description: Control notification for the identified items is sent to notify party via the message IE3N08 [Msg: E_CON_NOT]. The notify party will be forced to organise to unload and present the selected goods in order to make them available to customs for controls, regardless of whether the presentation of those goods was originally foreseen at this place or not. |
| Output (Final situation): The control decision for items which require a control upon entry is notified to notify party. | Output (Final situation): The control decision for items which require a control upon entry is notified to notify party. |


##### T-04-29 – Notify House consignments in incorrect state to Person filing the ENS

| Notify House consignments in incorrect state to Person filing the ENS | T-04-29 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Incorrect state notification received from CR | Input: Incorrect state notification received from CR |
| Description: The message IE3N07 [Msg: E_HCS_INC_NOT] is sent to the Person filing the ENS notifying that ENS contains House consignments in incorrect state. | Description: The message IE3N07 [Msg: E_HCS_INC_NOT] is sent to the Person filing the ENS notifying that ENS contains House consignments in incorrect state. |
| Output (Final situation): Person filing the ENS is notified that the ENS contains House consignments in incorrect state. | Output (Final situation): Person filing the ENS is notified that the ENS contains House consignments in incorrect state. |


##### T-04-33 – Notify House consignments in incorrect state to Notify Party

| Notify House consignments in incorrect state to Notify Party | T-04-33 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Incorrect state notification received from CR | Input: Incorrect state notification received from CR |
| Description: The message IE3N07 [Msg: E_HCS_INC_NOT] is sent to the Notify Party notifying that ENS contains House consignments in incorrect state. | Description: The message IE3N07 [Msg: E_HCS_INC_NOT] is sent to the Notify Party notifying that ENS contains House consignments in incorrect state. |
| Output (Final situation): Notify Party is notified that the ENS contains House consignments in incorrect state. | Output (Final situation): Notify Party is notified that the ENS contains House consignments in incorrect state. |


##### T-04-36 – Notify control to Person filing the ENS

| Notify control to Person notifying the arrival | T-04-36 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Control notification received from CR | Input: Control notification received from CR |
| Description: Control notification for the identified items is sent to the Person filing the ENS via the message IE3N08 [Msg: E_CON_NOT]. | Description: Control notification for the identified items is sent to the Person filing the ENS via the message IE3N08 [Msg: E_CON_NOT]. |
| Output (Final situation): The control decision for items which require control upon entry is notified to Person filing the ENS. | Output (Final situation): The control decision for items which require control upon entry is notified to Person filing the ENS. |
