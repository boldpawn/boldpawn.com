---
title: "Sub-process L4-ICS2-02-01: Relate ENS filings (HTI) (2021-04-30) v1.30 - part 5"
source_title: "Sub-process L4-ICS2-02-01: Relate ENS filings (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-02-01 Relate ENS filings (HTI)-(2021-04-30)-v1.30.md"
section: "Branch 1 — notify Person filing"
chunk: 5
chunk_count: 6
generated: true
---

# Sub-process L4-ICS2-02-01: Relate ENS filings (HTI) (2021-04-30) v1.30 - part 5

Source document: Sub-process L4-ICS2-02-01: Relate ENS filings (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-02-01 Relate ENS filings (HTI)-(2021-04-30)-v1.30.md`
Chunk: 5 of 6

### Branch 1 — notify Person filing

1. **E-02-01-01 — ENS not complete notification for Person Filing received** *(message start event)*
2. → gateway **G-02-01-01 — Person filing to be notified?**
   - **Yes** → **T-02-01-01 — Notify ENS not complete to Person filing** *(sends IE3N02 E_ENS_NCP_NOT)* → *(merge gateway)*
   - **No** → *(merge gateway)*
3. → **E-02-01-02 — ENS not complete notified to Person filing** *(end event)*
