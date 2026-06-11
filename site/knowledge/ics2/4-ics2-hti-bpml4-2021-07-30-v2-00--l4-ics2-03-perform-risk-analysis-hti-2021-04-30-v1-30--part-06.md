---
title: "Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30 - part 6"
source_title: "Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-03 Perform risk analysis (HTI)-(2021-04-30)-v1.30.md"
section: "Branch 2 — assessment complete to Carrier"
chunk: 6
chunk_count: 8
generated: true
---

# Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30 - part 6

Source document: Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-03 Perform risk analysis (HTI)-(2021-04-30)-v1.30.md`
Chunk: 6 of 8

### Branch 2 — assessment complete to Carrier

1. **E-03-29 — Assessment complete notification for Carrier received** *(message start event)*
2. → gateway **G-03-02 — Carrier to be notified about assessment complete?**
   - **Yes** → **T-03-02 — Notify assessment complete to Carrier** *(sends IE3N03 E_ASM_CMP_NOT)* → *(merge)*
   - **No** → *(merge)*
3. → **E-03-30 — Assessment complete notified to Carrier** *(end event)*
