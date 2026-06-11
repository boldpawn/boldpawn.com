---
title: "Sub-process HLS005: Process presentation and controls (2018-05-07) v1.0 - part 5"
source_title: "Sub-process HLS005: Process presentation and controls (2018-05-07) v1.0"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS005 Process presentation and controls-(2018-05-07)-v1.00.md"
section: "Member State of Presentation/control (at final destination) lane"
chunk: 5
chunk_count: 6
generated: true
---

# Sub-process HLS005: Process presentation and controls (2018-05-07) v1.0 - part 5

Source document: Sub-process HLS005: Process presentation and controls (2018-05-07) v1.0
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS005 Process presentation and controls-(2018-05-07)-v1.00.md`
Chunk: 5 of 6

### Member State of Presentation/control (at final destination) lane

1. Start events: **ANE04 — Control decision notified** and **PCE01 — Goods presented**.
2. → **GET05 — Retrieve relevant ENS information and control recommendation** *(task)*
3. → **PCT04 — Stop 200 days timer** *(task)*
4. → **GET06 — Set Control decision** *(task)*
5. → gateway **Goods to be controlled?**
   - **No** → goods released → **PCE02 — Goods released for further customs treatment** *(end event)*
   - **Yes** → gateway **Goods to be controlled at the final destination?**
     - **Yes** → **PCT03 — Communicate control recommendation to final destination** *(task)* → **PCE05 — Goods to be controlled at final destination** *(end event)* — sends the control recommendation to the Member State of Control lane.
     - **No** → **GET07 — Notify control decision** *(task; produces Control notification)* → **PCT01 — Perform controls** *(task)* → **PCT02 — Document control results** *(task)* → gateway **Goods to be released for further customs treatment?**
       - **Yes** → **PCE02 — Goods released for further customs treatment** *(end event)*
       - **No** → **PCE03 — Goods not released for further customs treatment** *(end event)*
