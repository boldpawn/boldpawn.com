---
title: "Introduction - part 49"
source_title: "Introduction"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 49
chunk_count: 82
generated: true
---

# Introduction - part 49

Source document: Introduction
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md`
Chunk: 49 of 82

#### Tasks

##### T-07-01 – Perform syntactical & semantical validation on amendment of ENS filing

| Perform syntactical & semantical validation on amendment of ENS filing | T-07-01 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Amendment of ENS filing received | Input: Amendment of ENS filing received |
| Description: The received amendment is validated in an automated process. It comprises the validation against the message rules and conditions, as well as the relevant reference data which is valid at the date of the reception. The messages rules and conditions and relevant reference data are defined in [R02] ICS2 Information Exchange Messages. In case the ENS filing contains characters which are not Cyrillic, Latin or Greek, the ENS filing is rejected with “102” “Not allowed Unicode character” with pointer(s) to the elements resulting in the rejection. | Description: The received amendment is validated in an automated process. It comprises the validation against the message rules and conditions, as well as the relevant reference data which is valid at the date of the reception. The messages rules and conditions and relevant reference data are defined in [R02] ICS2 Information Exchange Messages. In case the ENS filing contains characters which are not Cyrillic, Latin or Greek, the ENS filing is rejected with “102” “Not allowed Unicode character” with pointer(s) to the elements resulting in the rejection. |
| Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated | Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated |


##### T-07-02 – Submit amendment of ENS filing to CR

| Submit amendment of ENS filing to the CR | T-07-02 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Positively validated amendment of ENS filing | Input: Positively validated amendment of ENS filing |
| Description: The amendment of ENS filing is submitted to the CR via the message IE4AXX [Msg: C_ENS_xxx_AMD]. | Description: The amendment of ENS filing is submitted to the CR via the message IE4AXX [Msg: C_ENS_xxx_AMD]. |
| Output (Final situation): The amendment of ENS filing was submitted to the CR. | Output (Final situation): The amendment of ENS filing was submitted to the CR. |


##### T-07-03 - Notify completion of amendment of ENS filing

| Notify completion of amendment of ENS filing | T-07-03 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Amendment of ENS filing completion notification received | Input: Amendment of ENS filing completion notification received |
| Description: The message IE3N10 [Msg: E_AMD_NOT] is generated and sent to the Person filing the amendment. It informs about successful amendment of the ENS filing. | Description: The message IE3N10 [Msg: E_AMD_NOT] is generated and sent to the Person filing the amendment. It informs about successful amendment of the ENS filing. |
| Output (Final situation): The Person filing is notified that the amendment of ENS filing was processed successfully. | Output (Final situation): The Person filing is notified that the amendment of ENS filing was processed successfully. |


##### T-07-04 – Notify error in amendment of ENS filing

| Notify error in amendment of ENS filing | T-07-04 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Notification of error in amendment of ENS filing received | Input: Notification of error in amendment of ENS filing received |
| Description: The information about a validation error is sent to the Person filing the amendment via the message IE3N99 [Msg: E_ERR_NOT]. | Description: The information about a validation error is sent to the Person filing the amendment via the message IE3N99 [Msg: E_ERR_NOT]. |
| Output (Final situation): The Person filing the amendment is informed about unsuccessful ENS lifecycle validation of the amendment of ENS filing. | Output (Final situation): The Person filing the amendment is informed about unsuccessful ENS lifecycle validation of the amendment of ENS filing. |


##### T-07-05 – Notify ENS lifecycle validation error

| Notify ENS lifecycle validation error | T-07-05 |
| --- | --- |
| System component: TI | System component: TI |
| Input: ENS lifecycle validation error received | Input: ENS lifecycle validation error received |
| Description: The information about unsuccessful ENS lifecycle validation is sent to the Person filing the amendment via the message IE3N01 [Msg: E_ELF_VLD_NOT]. | Description: The information about unsuccessful ENS lifecycle validation is sent to the Person filing the amendment via the message IE3N01 [Msg: E_ELF_VLD_NOT]. |
| Output (Final situation): The Person filing the amendment is informed about unsuccessful ENS lifecycle validation of the amendment of ENS filing. | Output (Final situation): The Person filing the amendment is informed about unsuccessful ENS lifecycle validation of the amendment of ENS filing. |


##### T-07-14 – Notify error

| Notify error | T-07-14 |
| --- | --- |
| System component: TI | System component: TI |
| Input: ENS filing amendment rejection received | Input: ENS filing amendment rejection received |
| Description: The Person filing is notified about the fact that the ENS filing amendment was rejected due to low quality data and about the errors that caused its rejection via the message IE3N99 [Msg: E_ERR_NOT]. | Description: The Person filing is notified about the fact that the ENS filing amendment was rejected due to low quality data and about the errors that caused its rejection via the message IE3N99 [Msg: E_ERR_NOT]. |
| Output (Final situation): The notification of error is generated and sent to the Person filing. | Output (Final situation): The notification of error is generated and sent to the Person filing. |
