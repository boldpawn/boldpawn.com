---
title: "Introduction - part 24"
source_title: "Introduction"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md"
section: "Tasks"
chunk: 24
chunk_count: 65
generated: true
---

# Introduction - part 24

Source document: Introduction
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md`
Chunk: 24 of 65

#### Tasks

##### GET08 – Cleanse ENS filling

| Cleanse ENS filling | GET08 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: ENS filing, Amendment of ENS filing | Input: ENS filing, Amendment of ENS filing |
| Description The cleansing of the ENS data removes symbols and strings from the ENS filing according the common predefined rules agreed with the Member states. The cleansed version of the ENS filing is stored (and the non-cleansed for the purpose of enrichment) and the sub-process continues. | Description The cleansing of the ENS data removes symbols and strings from the ENS filing according the common predefined rules agreed with the Member states. The cleansed version of the ENS filing is stored (and the non-cleansed for the purpose of enrichment) and the sub-process continues. |
| Output (Final situation): The cleansed version of the ENS filing is stored. | Output (Final situation): The cleansed version of the ENS filing is stored. |


##### GET01 – Perform syntax and semantic validation

| Perform syntax and semantic validation | GET01 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry | Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry |
| Input: ENS filing, Arrival notification, Amendment of ENS filing, Invalidation request | Input: ENS filing, Arrival notification, Amendment of ENS filing, Invalidation request |
| Description: The input file is validated in an automated process. The validation comprises validation of message syntax, checks of the compliance with business rules and conditions as well as reference data. If the input file is not valid it is rejected. If the input file is valid it is further processed. | Description: The input file is validated in an automated process. The validation comprises validation of message syntax, checks of the compliance with business rules and conditions as well as reference data. If the input file is not valid it is rejected. If the input file is valid it is further processed. |
| Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated the input file will obtain the state 'rejected' | Output (Final situation): The validation will have the following possible outcomes: Successful; Unsuccessful, in which case the error description will be generated the input file will obtain the state 'rejected' |


##### GET09 – Check stop words/phrases

| Check stop words/phrases | GET09 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: ENS filing, Amendment of ENS filing | Input: ENS filing, Amendment of ENS filing |
| Description The cleansed version of the ENS filing is checked against the pre-defined stop words/phrases per data element. In case the data element contains only stop words/phrases it is considered empty and, if the data element is mandatory, the ENS filing is rejected due to poor data quality. The stop words/phrases are not removed from the data elements of the ENS filing. If the validation is successful, the ENS filing the state 'accepted' and the sub-process continues. If the validation fails the ENS filing obtains the state 'rejected' and the sub-process ends. | Description The cleansed version of the ENS filing is checked against the pre-defined stop words/phrases per data element. In case the data element contains only stop words/phrases it is considered empty and, if the data element is mandatory, the ENS filing is rejected due to poor data quality. The stop words/phrases are not removed from the data elements of the ENS filing. If the validation is successful, the ENS filing the state 'accepted' and the sub-process continues. If the validation fails the ENS filing obtains the state 'rejected' and the sub-process ends. |
| Output (Final situation): The state of the ENS filing will be either changed to 'accepted' or to 'rejected' due to insufficient data quality. | Output (Final situation): The state of the ENS filing will be either changed to 'accepted' or to 'rejected' due to insufficient data quality. |


##### GET04 - Perform ENS-lifecycle validation

| Perform ENS lifecycle validation | GET04 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: ENS filing, Amendment of ENS filing, or Invalidation request | Input: ENS filing, Amendment of ENS filing, or Invalidation request |
| Description The ENS lifecycle validation checks Whether the ENS filing is in the correct state regarding the ENS lifecycle’s sequence of the process (E.g. whether an ENS filing refers to a journey of a means of transport which arrived already in the past, or the goods covered by the ENS filing were already presented to the customs and effectively the ENS cannot be filed for the same goods already presented, or ENS cannot be amended or invalidated) and; Whether relevant key data; in particular data elements which are used as a unique linking key are unique. In the case of indicated split shipments (i.e. ENS filing containing defined split shipment code indicator) where unique linking key data can be re-used for the remaining goods after the shipment was split, the ENS filing will be processed. For an ENS amendment; That the state of the original ENS filing is not 'presented' or at a state where, for example amendment of the ENS filing is not possible anymore (i.e. customs decided to perform controls, or trader notified by the customs that goods are to be controlled). That no particulars were amended which are not allowed to be amended. The person filing the amendment needs to be the declarant or the representative of the ENS filing referred to. For an invalidation request; The related ENS filing is required to be in a state prior to the state 'presented'. The person filing the invalidation request needs to be the declarant or the representative of the ENS filing referred to. If the ENS lifecycle validation is successful, the input file obtains the state 'accepted' and the sub-process continues. If the ENS lifecycle validation fails the input file obtains the state 'rejected' and the sub-process ends. | Description The ENS lifecycle validation checks Whether the ENS filing is in the correct state regarding the ENS lifecycle’s sequence of the process (E.g. whether an ENS filing refers to a journey of a means of transport which arrived already in the past, or the goods covered by the ENS filing were already presented to the customs and effectively the ENS cannot be filed for the same goods already presented, or ENS cannot be amended or invalidated) and; Whether relevant key data; in particular data elements which are used as a unique linking key are unique. In the case of indicated split shipments (i.e. ENS filing containing defined split shipment code indicator) where unique linking key data can be re-used for the remaining goods after the shipment was split, the ENS filing will be processed. For an ENS amendment; That the state of the original ENS filing is not 'presented' or at a state where, for example amendment of the ENS filing is not possible anymore (i.e. customs decided to perform controls, or trader notified by the customs that goods are to be controlled). That no particulars were amended which are not allowed to be amended. The person filing the amendment needs to be the declarant or the representative of the ENS filing referred to. For an invalidation request; The related ENS filing is required to be in a state prior to the state 'presented'. The person filing the invalidation request needs to be the declarant or the representative of the ENS filing referred to. If the ENS lifecycle validation is successful, the input file obtains the state 'accepted' and the sub-process continues. If the ENS lifecycle validation fails the input file obtains the state 'rejected' and the sub-process ends. |
| Output (Final situation): The state of the input file will either be changed to 'accepted' or 'rejected'. | Output (Final situation): The state of the input file will either be changed to 'accepted' or 'rejected'. |


##### GET03 - Notify error

| Notify error | GET03 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry | Organisation: National Customs Administration of Responsible Member State or National Customs Administration of Member State of actual first entry |
| Input: Unsuccessful validation results - containing the error description | Input: Unsuccessful validation results - containing the error description |
| Description: The file is not valid and rejected. The person filing is notified of the rejection and of the errors that caused the rejection. | Description: The file is not valid and rejected. The person filing is notified of the rejection and of the errors that caused the rejection. |
| Output (Final situation): The notification of error is generated and sent to the person filing and the process ends. | Output (Final situation): The notification of error is generated and sent to the person filing and the process ends. |


##### PET01 - Send filing acceptance notification

| Send filing acceptance notification | PET01 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: Successful validation results | Input: Successful validation results |
| Description The registration of the filing and assigned MRN are notified to the person filing and to the Carrier when it has requested to be notified. | Description The registration of the filing and assigned MRN are notified to the person filing and to the Carrier when it has requested to be notified. |
| Output (Final situation) The person filing and where applicable the carrier were notified. | Output (Final situation) The person filing and where applicable the carrier were notified. |


##### PET04 – Manage timer for completion of risk analysis

| Manage timer for completion of risk analysis | PET04 |
| --- | --- |
| Organisation: National Customs Administration of Responsible Member State | Organisation: National Customs Administration of Responsible Member State |
| Input: Accepted ENS filing or Amended ENS filing | Input: Accepted ENS filing or Amended ENS filing |
