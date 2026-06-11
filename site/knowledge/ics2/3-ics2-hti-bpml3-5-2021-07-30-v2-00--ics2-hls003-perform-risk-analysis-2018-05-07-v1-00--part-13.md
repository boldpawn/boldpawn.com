---
title: "Sub-process HLS003: Perform risk analysis (2018-05-07) v1.0 - part 13"
source_title: "Sub-process HLS003: Perform risk analysis (2018-05-07) v1.0"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS003 Perform risk analysis-(2018-05-07)-v1.00.md"
section: "Air pre-loading risk analysis flow"
chunk: 13
chunk_count: 15
generated: true
---

# Sub-process HLS003: Perform risk analysis (2018-05-07) v1.0 - part 13

Source document: Sub-process HLS003: Perform risk analysis (2018-05-07) v1.0
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS003 Perform risk analysis-(2018-05-07)-v1.00.md`
Chunk: 13 of 15

### Air pre-loading risk analysis flow

1. **PEE01 — 7+1 data ready for air pre-loading risk analysis** *(start)* → risk assessment.
2. gateway **Risk identified?**
   - leads to information-gathering loop: **RAT06 — Request additional information from EO** / **RAT10 — Request additional information from IMS** → **RAE01 / RAE03** responses received → **RAT07 — Assess additional information from EO** / **RAT11 — Assess additional information from IMS**.
   - gateway **Additional information needed … / Further additional information needed …?** loops back for more information when needed.
3. gateway **Issue DNL?**
   - **Yes** → **RAT12 — Set RA result** → **RAT14 — Issue DNL** → **RAE05 — DNL issued**; **RAT13 — Send assessment complete notification** (produces Assessment complete notification).
4. End: **RAE04 — Air pre-loading risk analysis complete** *(end event)*.
