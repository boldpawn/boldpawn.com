---
title: "Sub-process L4-ICS2-00: Master process (HTI) (2018-10-15) v1.10 - part 4"
source_title: "Sub-process L4-ICS2-00: Master process (HTI) (2018-10-15) v1.10"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-00 Master process (HTI)-(2018-10-15)-v1.10.md"
section: "Responsible Member State lane"
chunk: 4
chunk_count: 5
generated: true
---

# Sub-process L4-ICS2-00: Master process (HTI) (2018-10-15) v1.10 - part 4

Source document: Sub-process L4-ICS2-00: Master process (HTI) (2018-10-15) v1.10
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-00 Master process (HTI)-(2018-10-15)-v1.10.md`
Chunk: 4 of 5

### Responsible Member State lane

- **E-00-01 — ENS filing received** *(start event)*
  → **L4-ICS2-01 — Register filing** *(sub-process)*
  → **L4-ICS2-02 — Prepare ENS for risk analysis** *(sub-process)*
  → **L4-ICS2-03 — Perform risk analysis** *(sub-process)*
  → **E-00-04 — Entry process completed** *(end event)*
- **E-00-02 — Amendment of ENS filing received** *(start event)*
  → **L4-ICS2-07 — Amend filing** *(sub-process)*
  → joins **L4-ICS2-02 — Prepare ENS for risk analysis**
- **E-00-03 — Timer for completion of linking expired** *(timer start event)*
  → **L4-ICS2-02-01 — Relate ENS with House consignments** *(sub-process)*
  → **L4-ICS2-03 — Perform risk analysis**
- **E-00-05 — Invalidation request received** *(start event)* and **E-00-06 — 200 day timer expired** *(timer start event)*
  → *(gateway)* → **L4-ICS2-08 — Invalidate filing** *(sub-process)*
  → **E-00-07 — ENS filing invalidated** *(end event)*
