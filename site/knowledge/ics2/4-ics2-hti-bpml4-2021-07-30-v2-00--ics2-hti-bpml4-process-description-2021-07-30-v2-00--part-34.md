---
title: "Introduction - part 34"
source_title: "Introduction"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 34
chunk_count: 82
generated: true
---

# Introduction - part 34

Source document: Introduction
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md`
Chunk: 34 of 82

#### Tasks

##### T-03-01-01 – Request additional information from Person filing

| Request additional information from Person filing | T-03-01-01 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Request for additional information from Person filing received | Input: Request for additional information from Person filing received |
| Description: The request for additional information is sent to the Person filing via the message IE3Q02 [Msg: E_REF_RFI_REQ]. | Description: The request for additional information is sent to the Person filing via the message IE3Q02 [Msg: E_REF_RFI_REQ]. |
| Output (Final situation): The request for additional information is sent to the Person filing. | Output (Final situation): The request for additional information is sent to the Person filing. |


##### T-03-01-02 – Notify additional information request to Carrier

| Notify additional information request to Carrier | T-03-01-02 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Request for additional information from Person filing received | Input: Request for additional information from Person filing received |
| Description: The Carrier is notified via the message IE3N04 [Msg: E_REF_RFI_NOT] that the Person filing was requested to provide additional information. | Description: The Carrier is notified via the message IE3N04 [Msg: E_REF_RFI_NOT] that the Person filing was requested to provide additional information. |
| Output (Final situation): The Carrier is notified about the additional information request that was sent to the Person filing. | Output (Final situation): The Carrier is notified about the additional information request that was sent to the Person filing. |


##### T-03-01-03 – Submit additional information response from Person filing

| Submit additional information response received from Person filing | T-03-01-03 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Additional information from Person filing received | Input: Additional information from Person filing received |
| Description: The response to the additional information request received from Person filing is sent to the CR via the message IE4R02 [Msg: C_REF_RFI_RSP]. In case no corresponding RfI request can be found, then an error notification is sent by the CR to the STI and then forwarded to the person filing. | Description: The response to the additional information request received from Person filing is sent to the CR via the message IE4R02 [Msg: C_REF_RFI_RSP]. In case no corresponding RfI request can be found, then an error notification is sent by the CR to the STI and then forwarded to the person filing. |
| Output (Final situation): The additional information from Person filing is available in the CR. | Output (Final situation): The additional information from Person filing is available in the CR. |


##### T-03-01-04 – Request HRCM screening from Person filing

| Request HRCM screening from Person filing | T-03-01-04 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Request for HRCM screening received | Input: Request for HRCM screening received |
| Description: The request for HRCM screening is sent to the Person filing via the message IE3Q03 [Msg: E_REF_RFS_REQ]. | Description: The request for HRCM screening is sent to the Person filing via the message IE3Q03 [Msg: E_REF_RFS_REQ]. |
| Output (Final situation): The request for HRCM screening is sent to the Person filing. | Output (Final situation): The request for HRCM screening is sent to the Person filing. |


##### T-03-01-05 – Notify HRCM screening request to Carrier

| Notify HRCM screening request to Carrier | T-03-01-05 |
| --- | --- |
| System component: TI | System component: TI |
| Input: Request for HRCM screening received | Input: Request for HRCM screening received |
| Description: The Carrier is notified via the message IE3N05 [Msg: E_REF_RFS_NOT] that the Person filing was requested to provide HRCM screening outcome. | Description: The Carrier is notified via the message IE3N05 [Msg: E_REF_RFS_NOT] that the Person filing was requested to provide HRCM screening outcome. |
| Output (Final situation): The Carrier is notified about the HRCM screening request that was sent to the Person filing. | Output (Final situation): The Carrier is notified about the HRCM screening request that was sent to the Person filing. |


##### T-03-01-06 – Submit HRCM screening outcome

| Submit HRCM screening outcome | T-03-01-06 |
| --- | --- |
| System component: TI | System component: TI |
| Input: HRCM screening outcome received | Input: HRCM screening outcome received |
| Description: The outcome of the HRCM screening is sent to the CR via the message IE4R03 [Msg: C_REF_RFS_RSP]. In case no corresponding RfS request can be found, then an error notification is sent by the CR to the STI and then forwarded to the person filing. | Description: The outcome of the HRCM screening is sent to the CR via the message IE4R03 [Msg: C_REF_RFS_RSP]. In case no corresponding RfS request can be found, then an error notification is sent by the CR to the STI and then forwarded to the person filing. |
| Output (Final situation): The HRCM screening outcome is available in the CR. | Output (Final situation): The HRCM screening outcome is available in the CR. |
