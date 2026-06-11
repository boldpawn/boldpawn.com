---
title: "Introduction - part 11"
source_title: "Introduction"
source_path: "ICS2-HTI-R&C-(2022-05-23)-v2.01.md"
section: "Rules"
chunk: 11
chunk_count: 12
generated: true
---

# Introduction - part 11

Source document: Introduction
Source path: `ICS2-HTI-R&C-(2022-05-23)-v2.01.md`
Chunk: 11 of 12

# Rules

| R3002 | In the 'ADDRESS' class either 'ADDRESS/Street' and 'ADDRESS/Number' or 'ADDRESS/P.O.Box' has to be used. |
| --- | --- |


| R3003 | This attribute has to be unique as per Declarant (DECLARANT/Identification number). |
| --- | --- |


| R3004 | The MRN shall be generated using the following structure: Last two digits of year of the registration (n2) Identifier of the country where the message is addressed to (a2) Unique identifier per year and country (an11) Trader interface identifier (a1) Procedure identifier (a1) Check digit (an1) |
| --- | --- |


| R3005 | EORI number shall be declared. |
| --- | --- |


| R3006 | When 'COMMUNICATION/Type' is 'TE' the 'COMMUNICATION/Identifier' has to have an international phone number format as defined in the ITU-T recommendation E.123 (02/2001). The format must be compliant with the pattern "^\\+(?:[0-9] ?){6,14}[0-9]$”. |
| --- | --- |


| R3007 | Either 'ACTIVE BORDER TRANSPORT MEANS/Conveyance reference number' or 'ACTIVE BORDER TRANSPORT MEANS/Identification number' and 'ACTIVE BORDER TRANSPORT MEANS/Type of identification' must be used. |
| --- | --- |


| R3008 | Either 'CONSIGNMENT (HOUSE LEVEL)/PASSIVE BORDER TRANSPORT MEANS' or 'CONSIGNMENT (HOUSE LEVEL)/GOODS ITEM/PASSIVE BORDER TRANSPORT MEANS' must be used. |
| --- | --- |


| R3009 | When 'ACTIVE BORDER TRANSPORT MEANS/Type of identification' is '10' this attribute must contain IMO ship identification number and its value must be made of the three letters "IMO" followed by a six-digit sequential unique number and a check digit. The check digit is calculated by multiplying each of the first six digits by a factor of 2 to 7 corresponding to their position from right to left. The rightmost digit of this sum is the check digit. |
| --- | --- |


| R3010 | When 'ACTIVE BORDER TRANSPORT MEANS/Type of identification' is '80' this attribute must contain unique European Vessel Identification Number (ENI) and its value must have an eight digit format. |
| --- | --- |


| R3011 | When 'ACTIVE BORDER TRANSPORT MEANS/Mode of transport' is '1' the value of this attribute must be '10'. |
| --- | --- |


| R3012 | When 'ACTIVE BORDER TRANSPORT MEANS/Mode of transport' is '8' the value of this attribute must be '80'. |
| --- | --- |


| R3013 | 'ACTIVE BORDER TRANSPORT MEANS/Countries of routing of means of transport' first country in the sequence must be the same as 'CONSIGNMENT (MASTER LEVEL)/PLACE OF LOADING/UNLOCODE' or 'CONSIGNMENT (MASTER LEVEL)/PLACE OF LOADING/ADDRESS/Country'. |
| --- | --- |


| R3014 | 'ACTIVE BORDER TRANSPORT MEANS/Countries of routing of means of transport' last country in the sequence must be the same as 'CONSIGNMENT (MASTER LEVEL)/PLACE OF UNLOADING/UNLOCODE' or 'CONSIGNMENT (MASTER LEVEL)/PLACE OF UNLOADING/ADDRESS/Country'. |
| --- | --- |


| R3016 | Each ‘Goods Item Number’ must be unique throughout CONSIGNMENT (MASTER LEVEL). The items shall be numbered in a sequential fashion, starting from '1' for the first item and incrementing the numbering by '1' for each following item. |
| --- | --- |


| R3017 | Each ‘Goods Item Number’ must be unique throughout CONSIGNMENT (HOUSE LEVEL). The items shall be numbered in a sequential fashion, starting from '1' for the first item and incrementing the numbering by '1' for each following item. |
| --- | --- |


| R3018 | In the sequence of 'ACTIVE BORDER TRANSPORT MEANS/Countries of routing of means of transport' the first country that is an EU Member State must be the country of the 'CUSTOMS OFFICE OF FIRST ENTRY'. |
| --- | --- |


| R3019 | When used this attribute must contain at least 6 and maximum 8 digits. |
| --- | --- |


| R3020 | The format of this attribute is 9999999-9. |
| --- | --- |
| R3021 | The format of this attribute is ZZZZ9999999 (ISO 6346). |


| R3022 | 'COUNTRIES OF ROUTING OF CONSIGNMENT' must contain countries of 'PLACE OF ACCEPTANCE' and 'PLACE OF DELIVERY'. |
| --- | --- |


| R3024 | This class/attribute cannot be amended. |
| --- | --- |


| R3025 | Either 'DO NOT LOAD DETAILS/Receptacle identification numbers' or 'DO NOT LOAD DETAILS/TRANSPORT DOCUMENT (HOUSE LEVEL)' or 'DO NOT LOAD DETAILS/TRANSPORT EQUIPMENT' must be used. |
| --- | --- |


| R3026 | When ADDITIONAL INFORMATION RESPONSE/BINARY ATTACHMENT/hashCode is used ADDITIONAL INFORMATION RESPONSE/BINARY ATTACHMENT/hashCodeAlgorithmnID must be used. |
| --- | --- |


| R3027 | At least one of the following shall be provided: 'ACTIVE BORDER TRANSPORT MEANS/Estimated date and time of arrival' or 'RELATED MRN'. |
| --- | --- |


| R3028 | Either 'MRN' or 'TRANSPORT DOCUMENT' must be used. |
| --- | --- |


| R3029 | Either 'EXAMINATION PLACE/Place of examination' or 'EXAMINATION PLACE/Reference number' must be used. |
| --- | --- |


| R3030 | At least one of the classes 'RECEPTACLE', 'GOODS ITEM', 'CONSIGNMENT (HOUSE LEVEL)', 'PACKAGING', 'TRANSPORT DOCUMENT', 'TRANSPORT EQUIPMENT' must be used. |
| --- | --- |


| R3031 | This class should be provided, if available, either at CONSIGNMENT level or CONSIGNMENT (HOUSE LEVEL)/GOODS ITEM level. Where the same information is applicable to all Goods items in the same Consignment, it should be provided at CONSIGNMENT level only. |
| --- | --- |


| R3032 | Either 'CARRIER/Identification number' or 'CARRIER/Name' and 'CARRIER/ADDRESS' must be used. |
| --- | --- |


| R3033 | 'REPRESENTATIVE/Identification number' cannot be the same as 'DECLARANT/Identification number'. |
| --- | --- |


| R3034 | Either the ‘CUSTOMS OFFICE OF FIRST ENTRY’ or ‘ADDRESSED MEMBER STATE’ is to be provided. |
| --- | --- |


| R3035 | Either the ‘Estimated date and time of departure’ or ‘Actual date and time of departure’ is to be provided. |
| --- | --- |
| R3046 | This message can contain 'ENS QUERY PARAMETERS' class only with one of the two in the attribute 'ENS QUERY PARAMETERS/Parameter' per combination: Combination 1: ENSS or ENSH Combination 2: ENSC or ENSF or ENSP |
|  |  |
| R3048 | If the ‘RELATED TRANSPORT DOCUMENT’ is used, then the ‘RELATED MRN’ cannot be used. |
|  |  |
| R3051 | Only value “2 – Sub-house level filing” can be used. |
| R3052 | Each ‘Transport document (House level)’ must be unique throughout the ENS filing. |
| R3053 | The 'Estimated date and time of arrival' cannot be earlier than the 'Actual date and time of departure' or 'Estimated date and time of departure'. |
|  |  |
