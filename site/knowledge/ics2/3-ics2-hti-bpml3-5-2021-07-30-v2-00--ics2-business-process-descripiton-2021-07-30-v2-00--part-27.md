---
title: "Introduction - part 27"
source_title: "Introduction"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 27
chunk_count: 65
generated: true
---

# Introduction - part 27

Source document: Introduction
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md`
Chunk: 27 of 65

| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: ENS in the state 'ENS pending completeness' | Input: ENS in the state 'ENS pending completeness' |
| Description A partial master filing triggers the start of the timer for linking it with the related house level filings. The time allocated for the linking of partial filings should allow sufficient time for the completion of the risk analysis according to the legally set deadlines. Those deadlines depend on the given ENS scenario (i.e. transport mode and situations that determine time-limits – see Articles 105, 106, 107 UCC-DA). | Description A partial master filing triggers the start of the timer for linking it with the related house level filings. The time allocated for the linking of partial filings should allow sufficient time for the completion of the risk analysis according to the legally set deadlines. Those deadlines depend on the given ENS scenario (i.e. transport mode and situations that determine time-limits – see Articles 105, 106, 107 UCC-DA). |
| Output (Final situation): The timer for completing the ENS by linking house level filings with a partial master level filing is started. | Output (Final situation): The timer for completing the ENS by linking house level filings with a partial master level filing is started. |


##### PET11 - Check data consistency

| Check data consistency | PET11 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: Complete ENS | Input: Complete ENS |
| Description Data consistency of linked partial ENS filings will be checked. This is done in order to ensure the appropriate ENS data consistency is provided for the risk analysis. The consistency check will take into account inconsistencies such as major weight and piece count discrepancies, changed actors, discrepancies in the referenced transport documents at the level of master and underlying house filings as well as differences in air pre-loading and pre-arrival filings (algorithms and reference data for this task to be defined). | Description Data consistency of linked partial ENS filings will be checked. This is done in order to ensure the appropriate ENS data consistency is provided for the risk analysis. The consistency check will take into account inconsistencies such as major weight and piece count discrepancies, changed actors, discrepancies in the referenced transport documents at the level of master and underlying house filings as well as differences in air pre-loading and pre-arrival filings (algorithms and reference data for this task to be defined). |
| Output (Final situation): Data consistency check results, which will either be 'ok', or 'flagged', in the case the result is identified inconsistency. | Output (Final situation): Data consistency check results, which will either be 'ok', or 'flagged', in the case the result is identified inconsistency. |


##### PET12 – Relate ENS filings

| Relate ENS filings | PET12 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: ENS in the state 'ENS pending completeness' for which the linking expiration timer has expired | Input: ENS in the state 'ENS pending completeness' for which the linking expiration timer has expired |
| Description This automated task tries to relate ENS filings by taking into account other data elements with common attributes in terms of their content (e.g. container numbers, shipping marks etc.) in addition to the ULK used for task PET09 – Link related filings and check completeness. In case the house level partial ENS filings which are potentially to be related to a given partial master level filing are found, those filings will be proposed for the final manual determination and linking. | Description This automated task tries to relate ENS filings by taking into account other data elements with common attributes in terms of their content (e.g. container numbers, shipping marks etc.) in addition to the ULK used for task PET09 – Link related filings and check completeness. In case the house level partial ENS filings which are potentially to be related to a given partial master level filing are found, those filings will be proposed for the final manual determination and linking. |
| Output (Final situation): Related ENS filings found or not found. | Output (Final situation): Related ENS filings found or not found. |


##### PET13 - Notify ENS not complete

| Notify ENS not complete | PET13 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: ENS in the state 'ENS pending completeness' | Input: ENS in the state 'ENS pending completeness' |
| Description The person filing and the carrier when different from the person filing and has requested to be informed and is connected to the customs system will be informed that the ENS is not complete. The notification will include the relevant information i.e. which ENS filings from which person expected to file are still missing. | Description The person filing and the carrier when different from the person filing and has requested to be informed and is connected to the customs system will be informed that the ENS is not complete. The notification will include the relevant information i.e. which ENS filings from which person expected to file are still missing. |
| Output (Final situation): The person filing and where applicable the carrier are informed about the fact that no complete ENS was lodged so far. | Output (Final situation): The person filing and where applicable the carrier are informed about the fact that no complete ENS was lodged so far. |


##### PET17 – Start timer for RMS risk analysis completion

| Start timers for IMS risk analysis tasks completion | PET17 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: PLACI - 7+1 data elements or ENS | Input: PLACI - 7+1 data elements or ENS |
| Description: The appropriate timer value for the completion of e-risk analysis on 7+1 data elements or on an ENS is calculated and started. The timer values need to be agreed as per scenario. | Description: The appropriate timer value for the completion of e-risk analysis on 7+1 data elements or on an ENS is calculated and started. The timer values need to be agreed as per scenario. |
| Output (Final situation): The timer value for the RMS is known and countdown is started. | Output (Final situation): The timer value for the RMS is known and countdown is started. |
