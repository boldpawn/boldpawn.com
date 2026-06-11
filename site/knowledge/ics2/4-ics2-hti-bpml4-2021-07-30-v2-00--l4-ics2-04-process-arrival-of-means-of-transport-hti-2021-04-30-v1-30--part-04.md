---
title: "Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30 - part 4"
source_title: "Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-04 Process arrival of means of transport (HTI)-(2021-04-30)-v1.30.md"
section: "Main registration flow (Trader Interface lane)"
chunk: 4
chunk_count: 7
generated: true
---

# Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30 - part 4

Source document: Sub-process L4-ICS2-04: Process arrival of means of transport (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-04 Process arrival of means of transport (HTI)-(2021-04-30)-v1.30.md`
Chunk: 4 of 7

## Main registration flow (Trader Interface lane)

1. **E-04-01 — Arrival notification of the means of transport received** *(message start; IE3N06 E_ARV_NOT)*
2. → **T-04-01 — Perform syntactical and semantical validation on received arrival notification**
3. → gateway **G-04-01 — Syntactical & semantical validation successful?**
   - **No** → **T-04-02 — Notify error** *(sends IE3N99 E_ERR_NOT)* → **E-04-02 — Arrival of means of transport notification rejected** *(end event)*
   - **Yes** → **T-04-03 — Register arrival notification and assign MRN** → **T-04-04 — Submit arrival notification** → **E-04-03 — Arrival notification submitted** *(end event)*
