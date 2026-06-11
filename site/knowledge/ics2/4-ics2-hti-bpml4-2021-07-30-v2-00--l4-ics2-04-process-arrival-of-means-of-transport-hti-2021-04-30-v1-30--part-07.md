---
title: "Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30 - part 7"
source_title: "Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-04 Process arrival of means of transport (HTI)-(2021-04-30)-v1.30.md"
section: "Incorrect House consignment state notification branches"
chunk: 7
chunk_count: 7
generated: true
---

# Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30 - part 7

Source document: Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-04 Process arrival of means of transport (HTI)-(2021-04-30)-v1.30.md`
Chunk: 7 of 7

## Incorrect House consignment state notification branches

- **E-04-04 — Incorrect House consignment state notification for Person notifying the arrival received** → **T-04-06 — Notify House consignments in incorrect state to Person notifying the arrival** *(sends IE3N07 E_HCS_INC_NOT)* → **E-04-05 — Incorrect House consignment state notification sent to Person notifying the arrival** *(end event)*
- **E-04-17 — Incorrect House consignment state notification for Person filling ENS received** → gateway **G-04-18 — Person filing the ENS to be notified of the house consignment(s) in incorrect state?** — *ICS2 R3*
  - **Yes** → **T-04-29 — Notify House consignment in incorrect state to Person filing the ENS** *(sends IE3N07 E_HCS_INC_NOT)* → **E-04-18 — Incorrect House consignment state notification sent to Person filling ENS** *(end event)*
  - **No** → *(end)*
- **E-04-19 — Incorrect House consignment state notification for Notify party received** → gateway **G-04-17 — Notify Party to be notified?** — *ICS2 R3*
  - **Yes** → **T-04-33 — Notify House consignment in incorrect state to Notify Party** *(sends IE3N07 E_HCS_INC_NOT)* → **E-04-20 — Incorrect House consignment state notification sent to Notify Party** *(end event)*
  - **No** → *(end)*
