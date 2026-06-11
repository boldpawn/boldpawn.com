---
title: "Sub-process L4-ICS2-08: Invalidate filing (HTI) (2021-04-30) v1.30 - part 4"
source_title: "Sub-process L4-ICS2-08: Invalidate filing (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-08 Invalidate filing (HTI)-(2021-04-30)-v1.30.md"
section: "Main process flow (Trader Interface lane)"
chunk: 4
chunk_count: 7
generated: true
---

# Sub-process L4-ICS2-08: Invalidate filing (HTI) (2021-04-30) v1.30 - part 4

Source document: Sub-process L4-ICS2-08: Invalidate filing (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-08 Invalidate filing (HTI)-(2021-04-30)-v1.30.md`
Chunk: 4 of 7

## Main process flow (Trader Interface lane)

1. **E-08-01 — ENS filing invalidation request received** *(message start; IE3Q04 E_INV_REQ)*
2. → **T-08-01 — Perform syntactical & semantical validation on ENS filing invalidation request**
3. → gateway **G-08-01 — Syntactical & semantical validation successful?**
   - **No** → **T-08-02 — Notify error in invalidation of ENS filing** *(sends IE3N99 E_ERR_NOT)* → **E-08-07 — ENS filing invalidation request rejected** *(end event)*
   - **Yes** → **T-08-03 — Submit ENS filing invalidation request to CR** → **E-08-02 — ENS filing invalidation request submitted** *(end event)*
