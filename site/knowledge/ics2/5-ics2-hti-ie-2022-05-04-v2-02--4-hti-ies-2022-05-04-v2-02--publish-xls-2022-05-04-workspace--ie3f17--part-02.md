---
title: "IE3F17 - part 2"
source_title: "IE3F17"
source_path: "5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F17.md"
section: "IE3F17"
chunk: 2
chunk_count: 2
generated: true
---

# IE3F17 - part 2

Source document: IE3F17
Source path: `5. ICS2-HTI-IE-(2022-05-04)-v2.02/4. HTI-IEs-(2022-05-04)-v2.02/publish_XLS_2022-05-04_workspace/IE3F17.md`
Chunk: 2 of 2

| IE3F17 | 2 | Attribute | 1..1 |         Identification number | an..17 | M | R3005 |  | Declarant, coded | R123 | CIS/Declarant/identification | identificationNumber |  |  |
| IE3F17 | 2 | Composition | 1..1 |         Address |  | M | R3002 |  | Address | 04A | CIS/Declarant/Address | address |  |  |
| IE3F17 | 3 | Attribute | 1..1 |             City | an..35 | M |  |  | City name | 241 | CIS/Declarant/Address/cityName | city |  |  |
| IE3F17 | 3 | Attribute | 1..1 |             Country | a2 | M |  |  | Country, coded | 242 | CIS/Declarant/Address/country | country | CL718 | CL718 |
| IE3F17 | 3 | Attribute | 0..1 |             Sub-division | an..35 | Ο |  |  | Country sub-entity name | 243 | CIS/Declarant/Address/countrySubDivisionName | subDivision |  |  |
| IE3F17 | 3 | Attribute | 0..1 |             Street | an..70 | O |  |  | Street and number/P.O. Box | 239 | CIS/Declarant/Address/line | street |  |  |
| IE3F17 | 3 | Attribute | 0..1 |             Postcode | an..17 | C |  | C3034 | Postcode identification | 245 | CIS/Declarant/Address/postcode | postCode |  |  |
| IE3F17 | 3 | Attribute | 0..1 |             Street additional line | an..70 | O |  |  |  |  | CIS/Declarant/Address/additionalLine | streetAdditionalLine |  |  |
| IE3F17 | 3 | Attribute | 0..1 |             Number | an..35 | O |  |  |  |  | CIS/Declarant/Address/streetNumber | number |  |  |
| IE3F17 | 3 | Attribute | 0..1 |             P.O. Box | an..70 | O |  |  |  |  | CIS/Declarant/Address/poBox | poBox |  |  |
| IE3F17 | 2 | Composition | 1..9 |         Communication |  | M |  |  | Communication | 25A | CIS/Declarant/Communication | communication |  |  |
| IE3F17 | 3 | Attribute | 1..1 |             Identifier | an..512 | M | R3006 |  | Communication number | 240 | CIS/Declarant/Communication/identification | identifier |  |  |
| IE3F17 | 3 | Attribute | 1..1 |             Type | an..3 | M |  |  | Communication number type | 253 | CIS/Declarant/Communication/type | type | CL707 | CL707 |
