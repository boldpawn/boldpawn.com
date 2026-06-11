---
title: "Sub-process HLS004: Process arrival of means of transport (2021-04-30) v1.20 - part 4"
source_title: "Sub-process HLS004: Process arrival of means of transport (2021-04-30) v1.20"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS004 Process arrival of means of transport-(2021-04-30)-v1.20.md"
section: "Process flow (Member State of actual first entry lane)"
chunk: 4
chunk_count: 4
generated: true
---

# Sub-process HLS004: Process arrival of means of transport (2021-04-30) v1.20 - part 4

Source document: Sub-process HLS004: Process arrival of means of transport (2021-04-30) v1.20
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS004 Process arrival of means of transport-(2021-04-30)-v1.20.md`
Chunk: 4 of 4

## Process flow (Member State of actual first entry lane)

1. **ANE01 — Arrival notification of means of transport received** *(start event)*
2. → **GET01 — Perform syntax & semantic validation** *(task)*
3. → gateway **Syntax & semantic validation successful?**
   - **No** → **GET03 — Notify error** *(task; produces Error notification)* → **ANE02 — Arrival of means of transport notification rejected** *(end event)*
   - **Yes** → **GET02 — Register and generate MRN** *(task; produces MRN)*
4. → **GET05 — Retrieve relevant ENS information and control recommendation** *(task)*
5. → gateway **Is retrieved ENS in correct state?**
   - **No** → **ANT01 — Notify ENS in incorrect state** *(task; produces ENS state error notification)* → **ANE03 — ENS state not updated** *(end event)*
   - **Yes** → **ANT02 — Stop 200 days timer** *(task)* → **GET06 — Set Control decision** *(task)* → gateway **Goods to be controlled?**
     - **Yes** → **GET07 — Notify control decision** *(task; produces Control notification)* → **ANE04 — Control decision notified** *(end event)*
     - **No** → **ANE05 — No control at first port or airport of arrival required** *(end event)*
