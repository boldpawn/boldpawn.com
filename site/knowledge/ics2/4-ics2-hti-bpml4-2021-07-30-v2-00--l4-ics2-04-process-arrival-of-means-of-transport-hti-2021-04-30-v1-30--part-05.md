---
title: "Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30 - part 5"
source_title: "Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-04 Process arrival of means of transport (HTI)-(2021-04-30)-v1.30.md"
section: "Registration / MRN notification branches"
chunk: 5
chunk_count: 7
generated: true
---

# Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30 - part 5

Source document: Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-04 Process arrival of means of transport (HTI)-(2021-04-30)-v1.30.md`
Chunk: 5 of 7

## Registration / MRN notification branches

- **T-04-05 — Notify successful registration and MRN to Person notifying the arrival** *(sends IE3R04 E_ARV_REG_RSP)*
- gateway **G-04-07 — Notify party to be notified about registration of the Arrival notification and MRN?**
  - **Yes** → **T-04-26 — Notify successful registration and MRN to Notify party** *(sends IE3R04 E_ARV_REG_RSP)*
  - **No** → *(end)*
