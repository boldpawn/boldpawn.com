---
title: "Sub-process L4-ICS2-07: Amend ENS filing (HTI) (2021-07-30) v2.00 - part 4"
source_title: "Sub-process L4-ICS2-07: Amend ENS filing (HTI) (2021-07-30) v2.00"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-07 Amend ENS filing (HTI)-(2021-07-30)-v2.00.md"
section: "Main process flow (Trader Interface lane)"
chunk: 4
chunk_count: 8
generated: true
---

# Sub-process L4-ICS2-07: Amend ENS filing (HTI) (2021-07-30) v2.00 - part 4

Source document: Sub-process L4-ICS2-07: Amend ENS filing (HTI) (2021-07-30) v2.00
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-07 Amend ENS filing (HTI)-(2021-07-30)-v2.00.md`
Chunk: 4 of 8

## Main process flow (Trader Interface lane)

1. **E-07-01 — Amendment of ENS filing received by HTI** *(message start; IE3Axx E_ENS_xxx_AMD)*
2. → **T-07-01 — Perform syntactical & semantical validation on amendment of ENS filing**
3. → gateway **G-07-01 — Syntactical & semantical validation successful?**
   - **No** → **T-07-04 — Notify error in amendment of ENS filing** *(sends IE3N99 E_ERR_NOT)* → **E-07-05 — Amendment of ENS filing rejected** *(end event)*
   - **Yes** → **T-07-02 — Submit amendment of ENS filing to CR** → **E-07-02 — Amendment of ENS filing submitted** *(end event)*
