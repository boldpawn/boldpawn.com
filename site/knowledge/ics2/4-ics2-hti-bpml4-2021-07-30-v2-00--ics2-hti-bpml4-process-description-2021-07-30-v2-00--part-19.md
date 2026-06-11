---
title: "Introduction - part 19"
source_title: "Introduction"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 19
chunk_count: 82
generated: true
---

# Introduction - part 19

Source document: Introduction
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md`
Chunk: 19 of 82

#### Tasks

##### T-01-01 – Perform syntactical and semantical validation

| Perform syntactical and semantical validation | T-01-01 |
| --- | --- |
| System component: TI | System component: TI |
| Input: ENS filing | Input: ENS filing |
| Description: The received ENS filing is validated in an automated process. It comprises the validation against the message rules and conditions, as well as the relevant reference data which are valid at the date of the reception. The messages rules and conditions and relevant reference data are defined in [R02] ICS2 Information Exchange Messages. In case the ENS filing contains characters which are not Cyrillic, Latin or Greek, the ENS filing is rejected with “102” “Not allowed Unicode character” with pointer(s) to the elements resulting in the rejection. | Description: The received ENS filing is validated in an automated process. It comprises the validation against the message rules and conditions, as well as the relevant reference data which are valid at the date of the reception. The messages rules and conditions and relevant reference data are defined in [R02] ICS2 Information Exchange Messages. In case the ENS filing contains characters which are not Cyrillic, Latin or Greek, the ENS filing is rejected with “102” “Not allowed Unicode character” with pointer(s) to the elements resulting in the rejection. |
| Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated the ENS filing will obtain the state 'Rejected' | Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated the ENS filing will obtain the state 'Rejected' |


##### T-01-02 – Register ENS filing and assign MRN

| Register ENS filing and assign MRN | T-01-02 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Valid ENS filing | Input: Valid ENS filing |
| Description: Valid ENS filing obtains the state 'Registered'; an MRN is generated and assigned to the received ENS filing. | Description: Valid ENS filing obtains the state 'Registered'; an MRN is generated and assigned to the received ENS filing. |
| Output (Final situation): The ENS filing was registered and MRN has been generated and assigned to it. | Output (Final situation): The ENS filing was registered and MRN has been generated and assigned to it. |


##### T-01-03 – Submit registered ENS filing

| Submit registered ENS filing | T-01-03 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Registered ENS filing | Input: Registered ENS filing |
| Description: The registered ENS filing is submitted from the TI to the CR for further processing via the message IE4FXX [Msg: C_ENS_xxx_DEC] (The xxx is placeholder for the number of the filing, i.e. F10). | Description: The registered ENS filing is submitted from the TI to the CR for further processing via the message IE4FXX [Msg: C_ENS_xxx_DEC] (The xxx is placeholder for the number of the filing, i.e. F10). |
| Output (Final situation): The ENS filing is available at the CR. | Output (Final situation): The ENS filing is available at the CR. |


##### T-01-04 – Notify successful registration and MRN to Person filing

| Notify successful registration and MRN to Person filing | T-01-04 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Registered ENS filing | Input: Registered ENS filing |
| Description: The registration of the ENS filing and the MRN assigned to it are notified to the Person filing via the message IE3R01 [Msg: E_ENS_REG_RSP]. | Description: The registration of the ENS filing and the MRN assigned to it are notified to the Person filing via the message IE3R01 [Msg: E_ENS_REG_RSP]. |
| Output (Final situation): The Person filing is notified about the registration of the ENS filing and the MRN assigned to it. | Output (Final situation): The Person filing is notified about the registration of the ENS filing and the MRN assigned to it. |


##### T-01-05 – Notify successful registration and MRN to Carrier

| Notify successful registration and MRN to Carrier | T-01-05 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Registered ENS filing | Input: Registered ENS filing |
| Description: The registration of the ENS filing and the MRN assigned to it are notified to the Carrier via the message IE3R01 [Msg: E_ENS_REG_RSP]. | Description: The registration of the ENS filing and the MRN assigned to it are notified to the Carrier via the message IE3R01 [Msg: E_ENS_REG_RSP]. |
| Output (Final situation): The Carrier is notified about the registration of the ENS filing and the MRN assigned to it. | Output (Final situation): The Carrier is notified about the registration of the ENS filing and the MRN assigned to it. |


##### T-01-06 - Notify error

| Notify error | T-01-06 |
| --- | --- |
| System component: TI | System component: TI |
| Input: ENS filing with unsuccessful validation results | Input: ENS filing with unsuccessful validation results |
| Description: The Person filing is notified about the fact that the ENS filing was rejected and about the errors that caused its rejection via the message IE3N99 [Msg: E_ERR_NOT]. | Description: The Person filing is notified about the fact that the ENS filing was rejected and about the errors that caused its rejection via the message IE3N99 [Msg: E_ERR_NOT]. |
| Output (Final situation): The notification of error is generated and sent to the Person filing. | Output (Final situation): The notification of error is generated and sent to the Person filing. |


##### T-01-09 – Notify ENS lifecycle validation error

| Notify ENS lifecycle validation error | T-01-09 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Unsuccessful validation results | Input: Unsuccessful validation results |
| Description: An error was sent by the CR and it is communicated to the Person filing through the TI. The error notification will be sent via the message IE3N01 [Msg: E_ELF_VLD_NOT] which includes the indication(s)/explanation(s) of the error(s) which caused the failure of the validity check of the ENS filing. | Description: An error was sent by the CR and it is communicated to the Person filing through the TI. The error notification will be sent via the message IE3N01 [Msg: E_ELF_VLD_NOT] which includes the indication(s)/explanation(s) of the error(s) which caused the failure of the validity check of the ENS filing. |
| Output (Final situation): The Person filing is notified of the errors of the ENS filing. | Output (Final situation): The Person filing is notified of the errors of the ENS filing. |
