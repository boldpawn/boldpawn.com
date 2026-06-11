---
title: "Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30 - part 6"
source_title: "Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-04 Process arrival of means of transport (HTI)-(2021-04-30)-v1.30.md"
section: "Control notification branches"
chunk: 6
chunk_count: 7
generated: true
---

# Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30 - part 6

Source document: Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-04 Process arrival of means of transport (HTI)-(2021-04-30)-v1.30.md`
Chunk: 6 of 7

## Control notification branches

- **E-04-06 — Control notification to Person notifying the arrival received** → **T-04-07 — Notify control to Person notifying the arrival** *(sends IE3N08 E_CON_NOT)* → **E-04-07 — Control notified to Person notifying the arrival** *(end event)*
- **E-04-21 — Control notification for Notify Party received** → gateway **G-04-08 — Notify party to be notified about controls?**
  - **Yes** → **T-04-27 — Notify control to Notify party** *(sends IE3N08 E_CON_NOT)* → **E-04-22 — Control notified to Notify Party** *(end event)*
  - **No** → *(end)*
- **E-04-23 — Control notification for Person filing the ENS received** → **T-04-36 — Notify control to Person filing the ENS** *(sends IE3N08 E_CON_NOT)* → **E-04-24 — Control notified to Person filing the ENS** *(end event)* — *ICS2 R3*
