---
title: "Sub-process HLS001: Register filing (2018-10-15) v1.10 - part 4"
source_title: "Sub-process HLS001: Register filing (2018-10-15) v1.10"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS001 Register filing-(2018-10-15)-v1.10.md"
section: "Process flow (Responsible Member State lane)"
chunk: 4
chunk_count: 4
generated: true
---

# Sub-process HLS001: Register filing (2018-10-15) v1.10 - part 4

Source document: Sub-process HLS001: Register filing (2018-10-15) v1.10
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS001 Register filing-(2018-10-15)-v1.10.md`
Chunk: 4 of 4

## Process flow (Responsible Member State lane)

1. **RFE01 — ENS filing received** *(start event)* — triggered by the incoming **ENS filing** data object.
2. → **GET01 — Perform syntax & semantic validation** *(task)*
3. → **Syntax & semantic validation successful?** *(gateway)*
   - **No** → **GET03 — Notify error** *(task)* → **GEE01 — Filing rejected** *(end event)*
     - GET03 produces the **Error notification** data object.
   - **Yes** → **GET02 — Register and generate MRN** *(task)* → **RFT01 — Start 200days timer** *(task)* → **RFE02 — Filing registered** *(end event)*
