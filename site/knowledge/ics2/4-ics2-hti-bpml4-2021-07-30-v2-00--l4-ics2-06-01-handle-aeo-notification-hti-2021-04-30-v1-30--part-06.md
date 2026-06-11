---
title: "Sub-process L4-ICS2-06-01: Handle AEO notification (HTI) (2021-04-30) v1.30 - part 6"
source_title: "Sub-process L4-ICS2-06-01: Handle AEO notification (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-06-01 Handle AEO notification (HTI)-(2021-04-30)-v1.30.md"
section: "Branch 2 — notify Carrier (AEOS)"
chunk: 6
chunk_count: 6
generated: true
---

# Sub-process L4-ICS2-06-01: Handle AEO notification (HTI) (2021-04-30) v1.30 - part 6

Source document: Sub-process L4-ICS2-06-01: Handle AEO notification (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-06-01 Handle AEO notification (HTI)-(2021-04-30)-v1.30.md`
Chunk: 6 of 6

### Branch 2 — notify Carrier (AEOS)

1. **E-06-01-07 — Control notification for Carrier (AEOS) received from CR** *(message start event)*
2. → gateway **G-06-01-02 — Carrier (AEOS) to be notified?**
   - **Yes** → **T-06-01-02 — Notify Carrier (AEOS) about controls** *(sends IE3N09 E_AEO_CON_NOT)* → *(merge)*
   - **No** → *(merge)*
3. → **E-06-01-08 — Carrier (AEOS) notified** *(end event)*
