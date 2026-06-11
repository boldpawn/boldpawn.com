---
title: "Sub-process HLS006: Amend ENS filing (2021-07-30) v2.00 - part 4"
source_title: "Sub-process HLS006: Amend ENS filing (2021-07-30) v2.00"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS006 Amend ENS filing-(2021-07-30)-v2.00.md"
section: "Process flow (Responsible Member State lane)"
chunk: 4
chunk_count: 4
generated: true
---

# Sub-process HLS006: Amend ENS filing (2021-07-30) v2.00 - part 4

Source document: Sub-process HLS006: Amend ENS filing (2021-07-30) v2.00
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS006 Amend ENS filing-(2021-07-30)-v2.00.md`
Chunk: 4 of 4

## Process flow (Responsible Member State lane)

1. **AFE01 — Amendment of ENS filing received** *(start event)* — triggered by the **Amendment of ENS filing** data object.
2. → **GET08 — Cleanse ENS filing** *(task)*
3. → **GET01 — Perform syntax & semantic validation** *(task)*
4. → gateway **Syntax & semantic validation successful?**
   - **No** → **GET03 — Notify error** *(task; produces Error notification)* → **AFE02 — Amendment rejected** *(end event)*
   - **Yes** → **GET09 — Check stops words/phrases** *(task)* → gateway **Sufficient meaningful data?**
     - **No** → **GET03 — Notify error** → **AFE02 — Amendment rejected** *(end event)*
     - **Yes** → **AFT01 — Send amendment reception acknowledgement** *(task; produces Amendment reception ACK)*
5. → **AFT02 — Retrieve original ENS filing** *(task)*
6. → gateway **Original ENS filing exists?**
   - **No** → **GET03 — Notify error** → **AFE02 — Amendment rejected** *(end event)*
   - **Yes** → **GET04 — Perform ENS lifecycle validation** *(task)*
7. → gateway **ENS lifecycle validation successful?**
   - **No** → **GET03 — Notify error** → **AFE02 — Amendment rejected** *(end event)*
   - **Yes** → **AFT03 — Amend ENS filing** *(task; produces Amendment completion notification)* → **AFE03 — Amended ENS filing** *(end event)*
