---
title: "Sub-process L4-ICS2-13: Consult ENS (HTI) (2018-05-07) v1.0 - part 4"
source_title: "Sub-process L4-ICS2-13: Consult ENS (HTI) (2018-05-07) v1.0"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-13 Consult ENS (HTI)-(2018-05-07)-v1.00.md"
section: "Process flow (Harmonized Trader Interface lane)"
chunk: 4
chunk_count: 4
generated: true
---

# Sub-process L4-ICS2-13: Consult ENS (HTI) (2018-05-07) v1.0 - part 4

Source document: Sub-process L4-ICS2-13: Consult ENS (HTI) (2018-05-07) v1.0
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-13 Consult ENS (HTI)-(2018-05-07)-v1.00.md`
Chunk: 4 of 4

## Process flow (Harmonized Trader Interface lane)

1. **E-13-01 — ENS consultation request received** *(message start; IE3Q05 E_ENS_CNS)*
2. → **T-13-01 — Submit ENS consultation request** *(task)*
3. → **E-13-02 — ENS consultation results received** *(intermediate message event)*
4. → **T-13-02 — Send ENS consultation results** *(task; sends IE3R08 E_ENS_CNS_RES)*
5. → **E-13-03 — ENS consultation request processed** *(end event)*
