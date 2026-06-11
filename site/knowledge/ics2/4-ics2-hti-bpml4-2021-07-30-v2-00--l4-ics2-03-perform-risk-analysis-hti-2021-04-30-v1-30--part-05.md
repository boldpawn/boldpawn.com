---
title: "Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30 - part 5"
source_title: "Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-03 Perform risk analysis (HTI)-(2021-04-30)-v1.30.md"
section: "Branch 1 — assessment complete to Person filing"
chunk: 5
chunk_count: 8
generated: true
---

# Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30 - part 5

Source document: Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-03 Perform risk analysis (HTI)-(2021-04-30)-v1.30.md`
Chunk: 5 of 8

### Branch 1 — assessment complete to Person filing

1. **E-03-01 — Assessment complete notification for Person Filing received** *(message start event)*
2. → gateway **G-03-01 — Person filing to be notified about assessment complete?**
   - **Yes** → **T-03-01 — Notify assessment complete to Person filing** *(sends IE3N03 E_ASM_CMP_NOT)* → *(merge)*
   - **No** → *(merge)*
3. → **E-03-02 — Assessment complete notified to Person Filing** *(end event)*
