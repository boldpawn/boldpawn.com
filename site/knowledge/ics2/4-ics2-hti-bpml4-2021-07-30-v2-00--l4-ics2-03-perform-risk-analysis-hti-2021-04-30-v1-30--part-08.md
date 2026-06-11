---
title: "Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30 - part 8"
source_title: "Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-03 Perform risk analysis (HTI)-(2021-04-30)-v1.30.md"
section: "Branch 4 — DNL to Carrier"
chunk: 8
chunk_count: 8
generated: true
---

# Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30 - part 8

Source document: Sub-process L4-ICS2-03: Perform risk analysis (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-03 Perform risk analysis (HTI)-(2021-04-30)-v1.30.md`
Chunk: 8 of 8

### Branch 4 — DNL to Carrier

1. **E-03-31 — DNL notification for Carrier received** *(message start event)*
2. → gateway **G-03-03 — Carrier to be notified about DNL?**
   - **Yes** → **T-03-04 — Notify DNL to Carrier** *(sends IE3Q01 E_DNL_REQ)* → *(merge)*
   - **No** → *(merge)*
3. → **E-03-32 — DNL notified to Carrier** *(end event)*
