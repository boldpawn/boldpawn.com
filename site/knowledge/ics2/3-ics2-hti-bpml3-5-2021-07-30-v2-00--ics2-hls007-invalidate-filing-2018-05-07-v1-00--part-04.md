---
title: "Sub-process HLS007: Invalidate filing (2018-05-07) v1.0 - part 4"
source_title: "Sub-process HLS007: Invalidate filing (2018-05-07) v1.0"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS007 Invalidate filing-(2018-05-07)-v1.00.md"
section: "Process flow (Responsible Member State lane)"
chunk: 4
chunk_count: 4
generated: true
---

# Sub-process HLS007: Invalidate filing (2018-05-07) v1.0 - part 4

Source document: Sub-process HLS007: Invalidate filing (2018-05-07) v1.0
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS007 Invalidate filing-(2018-05-07)-v1.00.md`
Chunk: 4 of 4

## Process flow (Responsible Member State lane)

1. Start events:
   - **IFE01 — Invalidation request received** *(start event)* — triggered by the **Invalidation request** data object.
   - **IFE03 — 200 days timer expiry** *(timer start event)*.
2. From **IFE01** → **GET01 — Perform syntax & semantic validation** *(task)*
3. → gateway **Syntax & semantic validation successful?**
   - **No** → **GET03 — Notify error** *(task; produces Error notification)* → **IFE02 — Invalidation request rejected** *(end event)*
   - **Yes** → **IFT03 — Send invalidation request reception acknowledgement** *(task; produces Invalidation request reception ACK)*
4. → **GET04 — Perform ENS lifecycle validation** *(task)*
5. → gateway **ENS lifecycle validation successful?**
   - **No** → **GET03 — Notify error** → **IFE02 — Invalidation request rejected** *(end event)*
   - **Yes** → **IFT01 — Notify invalidation request acceptance** *(task; produces Invalidation request acceptance notification)*
6. The **IFT01** path and the **IFE03 — 200 days timer expiry** start event merge *(gateway)* → **IFT02 — Invalidate filing** *(task)* → **IFE04 — Filing invalidated** *(end event)*
