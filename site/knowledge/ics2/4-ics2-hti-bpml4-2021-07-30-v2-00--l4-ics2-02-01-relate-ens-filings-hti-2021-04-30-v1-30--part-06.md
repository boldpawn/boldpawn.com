---
title: "Sub-process L4-ICS2-02-01: Relate ENS filings (HTI) (2021-04-30) v1.30 - part 6"
source_title: "Sub-process L4-ICS2-02-01: Relate ENS filings (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-02-01 Relate ENS filings (HTI)-(2021-04-30)-v1.30.md"
section: "Branch 2 — notify Person that has not yet filed"
chunk: 6
chunk_count: 6
generated: true
---

# Sub-process L4-ICS2-02-01: Relate ENS filings (HTI) (2021-04-30) v1.30 - part 6

Source document: Sub-process L4-ICS2-02-01: Relate ENS filings (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-02-01 Relate ENS filings (HTI)-(2021-04-30)-v1.30.md`
Chunk: 6 of 6

### Branch 2 — notify Person that has not yet filed

1. **E-02-01-06 — ENS pending notification for Person that has not yet filed** *(message start event)*
2. → gateway **G-02-01-03 — Person that has not yet filed to be notified?**
   - **Yes** → **T-02-01-03 — Notify ENS filing pending to Person that has not yet filed** *(sends IE3N11 E_ENS_PND_NOT)* → *(merge gateway)*
   - **No** → *(merge gateway)*
3. → **E-02-01-07 — ENS not complete notified to Person that has not yet filed** *(end event)*
