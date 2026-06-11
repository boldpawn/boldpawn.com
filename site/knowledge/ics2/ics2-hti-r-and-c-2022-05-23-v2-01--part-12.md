---
title: "Introduction - part 12"
source_title: "Introduction"
source_path: "ICS2-HTI-R&C-(2022-05-23)-v2.01.md"
section: "Conditions"
chunk: 12
chunk_count: 12
generated: true
---

# Introduction - part 12

Source document: Introduction
Source path: `ICS2-HTI-R&C-(2022-05-23)-v2.01.md`
Chunk: 12 of 12

# Conditions

| C3001 | IF 'SPLIT CONSIGNMENT/Split consignment indicator' is set to '1', THEN this data attribute is M, ELSE it is not used. |
| --- | --- |


| C3002 | IF 'SPLIT CONSIGNMENT/Split consignment indicator' is set to '1', THEN this data attribute is O, ELSE it is not used. |
| --- | --- |


| C3003 | IF 'CONSIGNMENT (MASTER LEVEL)/PLACE OF ACCEPTANCE/UNLOCODE' is not provided, THEN this attribute/class is M, ELSE it is not used. |
| --- | --- |


| C3004 | IF 'CONSIGNMENT (HOUSE LEVEL)/PLACE OF ACCEPTANCE/UNLOCODE' is not provided, THEN this attribute/class is M, ELSE it is not used. |
| --- | --- |


| C3005 | IF 'CONSIGNMENT (MASTER LEVEL)/PLACE OF DELIVERY/UNLOCODE' is not provided, THEN this attribute/class is M, ELSE it is not used. |
| --- | --- |


| C3006 | IF 'CONSIGNMENT (HOUSE LEVEL)/PLACE OF DELIVERY/UNLOCODE' is not provided, THEN this attribute/class is M, ELSE it is not used. |
| --- | --- |


| C3007 | IF 'CONSIGNMENT (MASTER LEVEL)/PLACE OF LOADING/UNLOCODE' is not provided, THEN this attribute/class is M, ELSE it is not used. |
| --- | --- |


| C3008 | IF 'CONSIGNMENT (MASTER LEVEL)/PLACE OF UNLOADING/UNLOCODE' is not provided, THEN this attribute/class is M, ELSE it is not used. |
| --- | --- |


| C3012 | IF 'ACTIVE BORDER TRANSPORT MEANS/Mode of transport' is set to '1' or '8', THEN this data attribute is M, ELSE it is not used. |
| --- | --- |


| C3013 | IF 'CONSIGNMENT (MASTER LEVEL)/Container indicator' is set to '1', THEN this class is M at 'CONSIGNMENT (MASTER LEVEL)' or 'CONSIGNMENT (MASTER LEVEL)/GOODS ITEM', ELSE it is not used. |
| --- | --- |


| C3014 | IF 'CONSIGNMENT (HOUSE LEVEL)/Container indicator' is set to '1', THEN this class is M at 'CONSIGNMENT (HOUSE LEVEL)' or 'CONSIGNMENT (HOUSE LEVEL)/GOODS ITEM', ELSE it is not used. |
| --- | --- |


| C3015 | IF 'CONSIGNMENT (MASTER LEVEL)/TRANSPORT EQUIPMENT/Number of seals' is used, THEN this classis M, ELSE it is not used. |
| --- | --- |


| C3016 | IF 'CONSIGNMENT (MASTER LEVEL)/GOODS ITEM/TRANSPORT EQUIPMENT/Number of seals' is used, THEN this classis M, ELSE it is not used. |
| --- | --- |


| C3017 | IF 'CONSIGNMENT (HOUSE LEVEL)/TRANSPORT EQUIPMENT/Number of seals' is used, THEN this classis M, ELSE it is not used. |
| --- | --- |
| C3018 | IF 'CONSIGNMENT (HOUSE LEVEL)/GOODS ITEM/TRANSPORT EQUIPMENT/Number of seals' is used, THEN this classis M, ELSE it is not used. |


| C3019 | IF 'ACTIVE BORDER TRANSPORT MEANS/Mode of transport' is set to '4', THEN this data attribute is M, ELSE it is not used. |
| --- | --- |


| C3020 | IF 'CONSIGNMENT (HOUSE LEVEL)/PLACE OF DELIVERY/UNLOCODE' 2 first characters or 'CONSIGNMENT (HOUSE LEVEL)/PLACE OF DELIVERY/ADDRESS/Country' refers to EU MS, THEN this class is M, ELSE it is O. |
| --- | --- |


| C3021 | IF 'CONSIGNMENT (MASTER LEVEL)/ADDITIONAL INFORMATION/Code' is '10600', THEN this class is O, ELSE it is M. |
| --- | --- |


| C3022 | IF 'CONSIGNMENT (HOUSE LEVEL)/ADDITIONAL INFORMATION/Code' is '10600', THEN this class is O, ELSE it is M. |
| --- | --- |


| C3023 | IF 'CONSIGNMENT (MASTER LEVEL)/ADDITIONAL INFORMATION/Code' is '10600', THEN this class is M, ELSE it is O. |
| --- | --- |


| C3024 | IF 'CONSIGNMENT (HOUSE LEVEL)/ADDITIONAL INFORMATION/Code' is '10600', THEN this class is M, ELSE it is O. |
| --- | --- |


| C3025 | IF 'CONSIGNMENT (HOUSE LEVEL)/CONSIGNOR/Type of person' and 'CONSIGNMENT (HOUSE LEVEL)/CONSIGNEE/Type of person' is '1 (Natural person)', THEN attribute is O, ELSE it is M. |
| --- | --- |


| C3034 | IF 'ADRESS/Country' in CL733 has a country code with postal code indicator 'C', THEN this attribute in the same class is O, ELSE it is M. |
| --- | --- |


| C3035 | IF 'CONSGINMENT (MASTER LEVEL)/GOODS ITEM/PACKAGING/Type of packages' is 'VQ', 'VG', 'VL', 'VY', 'VR', 'VS', 'VO', 'NE', 'NF', 'NG', THEN this attribute is not used, ELSE it is M. |
| --- | --- |


| C3036 | IF 'CONSGINMENT (HOUSE LEVEL)/GOODS ITEM/PACKAGING/Type of packages' is 'VQ', 'VG', 'VL', 'VY', 'VR', 'VS', 'VO', 'NE', 'NF', 'NG', THEN this attribute is not used, ELSE it is M. |
| --- | --- |


| C3037 | Conditions IF 'CONSIGNMENT (HOUSE LEVEL)/ADDITIONAL INFORMATION/Code' contains value ‘10900’ and 'CONSIGNMENT (HOUSE LEVEL)/ADDITIONAL FISCAL REFERENCES' class is not provided THEN this class/attribute is M, ELSE it is O. |
| --- | --- |


| C3038 | Conditions IF 'CONSIGNMENT (HOUSE LEVEL)/ADDITIONAL INFORMATION/Code' contains value ‘10900’ THEN this class/attribute is M, ELSE it is O. |
| --- | --- |
|  |  |
| C3046 | IF ‘TRANSPORT DOCUMENT (HOUSE LEVEL)/Type’ is ‘703 - House air waybill’ or ‘740 - Airway bill’ (general)’, THEN this class is M, ELSE this class is O. |
|  |  |
| C3048 | IF 'CONSIGNMENT (MASTER LEVEL)/CONSIGNOR/Type of person' and 'CONSIGNMENT (MASTER LEVEL)/CONSIGNEE/Type of person' is '1 (Natural person)', THEN attribute is O, ELSE it is M. |
|  |  |
| C3049 | If the Arrival notification is provided via a National Arrival System (not via the STI), then the ‘PERSON NOTIFYING THE ARRIVAL’ is O, else it is M. |
|  |  |
| C3050 | IF 'CONSIGNMENT (MASTER LEVEL)/Container indicator' is set to '1', THEN this class is M, ELSE it is not used. |
|  |  |
| C3051 | IF 'CONSIGNMENT (HOUSE LEVEL)/Container indicator' is set to '1', THEN this class is M, ELSE it is not used. |


*End of document*
