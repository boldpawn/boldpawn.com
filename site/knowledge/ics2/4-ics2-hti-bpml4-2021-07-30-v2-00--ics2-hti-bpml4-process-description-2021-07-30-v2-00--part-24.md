---
title: "Introduction - part 24"
source_title: "Introduction"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 24
chunk_count: 82
generated: true
---

# Introduction - part 24

Source document: Introduction
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/ICS2-HTI-BPML4 Process Description-(2021-07-30)-v2.00.md`
Chunk: 24 of 82

#### Tasks

##### T-02-01-01 - Notify ENS not complete to Person filing

| Notify ENS not complete to Person filing | T-02-01-01 |
| --- | --- |
| System component: TI | System component: TI |
| Input: ENS not complete notification received from CR | Input: ENS not complete notification received from CR |
| Description: The Person filing is informed that the ENS is still not complete, via the message IE3N02 [Msg: E_ENS_NCP_NOT]. The notification will also include a list of the filers that this person has indicated as obliged to file and who have not filed yet. | Description: The Person filing is informed that the ENS is still not complete, via the message IE3N02 [Msg: E_ENS_NCP_NOT]. The notification will also include a list of the filers that this person has indicated as obliged to file and who have not filed yet. |
| Output (Final situation): The Person filing is notified about the ENS still not being complete. | Output (Final situation): The Person filing is notified about the ENS still not being complete. |


##### T-02-01-03 - Notify ENS filing pending to Person that has not yet filed

| Notify ENS filing pending to Person that has not yet filed | T-02-01-03 |
| --- | --- |
| System component: TI | System component: TI |
| Input: ENS not complete notification received from CR | Input: ENS not complete notification received from CR |
| Description: The Person that has not yet filed is informed that he is obliged to file an ENS filing, via the message IE3N11 [Msg: E_ENS_PND_NOT]. | Description: The Person that has not yet filed is informed that he is obliged to file an ENS filing, via the message IE3N11 [Msg: E_ENS_PND_NOT]. |
| Output (Final situation): The Person that has not yet filed is notified about the ENS filing pending. | Output (Final situation): The Person that has not yet filed is notified about the ENS filing pending. |
