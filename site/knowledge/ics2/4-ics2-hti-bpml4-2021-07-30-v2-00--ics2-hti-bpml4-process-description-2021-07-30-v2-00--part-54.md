---
title: "Introduction - part 54"
source_title: "Introduction"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 54
chunk_count: 82
generated: true
---

# Introduction - part 54

Source document: Introduction
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md`
Chunk: 54 of 82

#### Tasks

##### T-08-01 – Perform syntactical and semantical validation on ENS filing invalidation request

| Perform syntactical & semantical validation on ENS filing invalidation request | T-08-01 |
| --- | --- |
| System component: TI | System component: TI |
| Input: ENS filing invalidation request received | Input: ENS filing invalidation request received |
| Description: The received ENS filing invalidation request is validated in an automated process. It comprises the validation against the message rules and conditions, as well as the relevant reference data which is valid at the date of the reception. The messages rules and conditions and relevant reference data are defined in [R02] ICS2 Information Exchange Messages. | Description: The received ENS filing invalidation request is validated in an automated process. It comprises the validation against the message rules and conditions, as well as the relevant reference data which is valid at the date of the reception. The messages rules and conditions and relevant reference data are defined in [R02] ICS2 Information Exchange Messages. |
| Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated. | Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated. |


##### T-08-02 – Notify error in invalidation of ENS filing

| Notify error in invalidation of ENS filing | T-08-02 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Unsuccessful ENS filing invalidation request validation results | Input: Unsuccessful ENS filing invalidation request validation results |
| Description: The information about unsuccessful semantical and syntactical validation is sent to the Person filing the ENS filing invalidation request via the message IE3N99 [Msg: E_ERR_NOT]. | Description: The information about unsuccessful semantical and syntactical validation is sent to the Person filing the ENS filing invalidation request via the message IE3N99 [Msg: E_ERR_NOT]. |
| Output (Final situation): The Person filing the invalidation request is informed about unsuccessful semantical and syntactical validation. | Output (Final situation): The Person filing the invalidation request is informed about unsuccessful semantical and syntactical validation. |


##### T-08-03 – Submit ENS filing invalidation request to CR

| Submit ENS filing invalidation request to CR | T-08-03 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Positively validated ENS filing invalidation request | Input: Positively validated ENS filing invalidation request |
| Description: The ENS filing invalidation request is submitted to the CR via the message IE4Q06 [Msg: C_INV_REQ]. | Description: The ENS filing invalidation request is submitted to the CR via the message IE4Q06 [Msg: C_INV_REQ]. |
| Output (Final situation): The ENS filing invalidation request is available at the CR. | Output (Final situation): The ENS filing invalidation request is available at the CR. |


##### T-08-04 – Notify ENS lifecycle validation error

| Notify ENS lifecycle validation error | T-08-04 |
| --- | --- |
| System component: TI | System component: TI |
| Input: ENS filing invalidation request error received from CR. | Input: ENS filing invalidation request error received from CR. |
| Description: The information about unsuccessful semantical and syntactical validation or ENS lifecycle validation is sent to the Person filing the ENS filing invalidation request via the message IE3N01 [Msg: E_ELF_VLD_NOT]. | Description: The information about unsuccessful semantical and syntactical validation or ENS lifecycle validation is sent to the Person filing the ENS filing invalidation request via the message IE3N01 [Msg: E_ELF_VLD_NOT]. |
| Output (Final situation): The Person filing the invalidation request is informed about unsuccessful semantical and syntactical validation or ENS lifecycle validation. | Output (Final situation): The Person filing the invalidation request is informed about unsuccessful semantical and syntactical validation or ENS lifecycle validation. |


##### T-08-05 - Notify invalidation of ENS filing completion

| Notify invalidation of ENS filing completion | T-08-05 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Invalidation of ENS filing completion notification received | Input: Invalidation of ENS filing completion notification received |
| Description: The message IE3R07 [Msg: E_INV_ACC_RSP] is generated and sent to the Person filing the invalidation request. It informs about successful invalidation of the ENS filing. | Description: The message IE3R07 [Msg: E_INV_ACC_RSP] is generated and sent to the Person filing the invalidation request. It informs about successful invalidation of the ENS filing. |
| Output (Final situation): The Person filing is notified that the invalidation of ENS filing was processed successfully. | Output (Final situation): The Person filing is notified that the invalidation of ENS filing was processed successfully. |
