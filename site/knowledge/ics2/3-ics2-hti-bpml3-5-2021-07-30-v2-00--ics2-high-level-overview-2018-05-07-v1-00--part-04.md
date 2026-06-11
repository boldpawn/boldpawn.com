---
title: "ICS2 High Level End-to-End Process (2018-05-07) v1.0 - part 4"
source_title: "ICS2 High Level End-to-End Process (2018-05-07) v1.0"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-High level overview-(2018-05-07)-v1.00.md"
section: "Responsible Member State lane"
chunk: 4
chunk_count: 6
generated: true
---

# ICS2 High Level End-to-End Process (2018-05-07) v1.0 - part 4

Source document: ICS2 High Level End-to-End Process (2018-05-07) v1.0
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-High level overview-(2018-05-07)-v1.00.md`
Chunk: 4 of 6

### Responsible Member State lane

- **HLE01 — ENS filing received** *(start event)*
  → **HLS001 — Register filing** *(sub-process)*
  → **HLS002 — Prepare ENS for risk analysis** *(sub-process)*
  → **HLS003 — Perform risk analysis** *(sub-process)*
  → **HLE06 — Entry process completed** *(end event)*
- **HLE02 — Amendment of ENS filing received** *(start event)*
  → **HLS006 — Amend filing** *(sub-process)*
  → joins **HLS002 — Prepare ENS for risk analysis**
- **HLE03 — Invalidation request received** *(start event)* and **HLE04 — 200 days timer expired** *(timer start event)*
  → *(gateway)* → **HLS007 — Invalidate filing** *(sub-process)*
  → **HLE05 — ENS filing invalidated** *(end event)*
