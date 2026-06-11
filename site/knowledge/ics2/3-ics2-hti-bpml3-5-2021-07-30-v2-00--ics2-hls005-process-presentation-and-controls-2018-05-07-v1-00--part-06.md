---
title: "Sub-process HLS005: Process presentation and controls (2018-05-07) v1.0 - part 6"
source_title: "Sub-process HLS005: Process presentation and controls (2018-05-07) v1.0"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS005 Process presentation and controls-(2018-05-07)-v1.00.md"
section: "Member State of Control (at final destination) lane"
chunk: 6
chunk_count: 6
generated: true
---

# Sub-process HLS005: Process presentation and controls (2018-05-07) v1.0 - part 6

Source document: Sub-process HLS005: Process presentation and controls (2018-05-07) v1.0
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS005 Process presentation and controls-(2018-05-07)-v1.00.md`
Chunk: 6 of 6

### Member State of Control (at final destination) lane

1. **PCE04 — Control recommendation received** *(start event)* — triggered by PCT03 from the presentation/control lane.
2. → **GET06 — Set Control decision** *(task)*
3. → **GET07 — Notify control decision** *(task; produces Control notification)*
4. → **PCT01 — Perform controls** *(task)*
5. → **PCT02 — Document control results** *(task)*
6. → gateway **Goods to be released for further customs treatment?**
   - **Yes** → **PCE02 — Goods released for further customs treatment** *(end event)*
   - **No** → **PCE03 — Goods not released for further customs treatment** *(end event)*
