---
title: "Sub-process HLS003: Perform risk analysis (2018-05-07) v1.0 - part 14"
source_title: "Sub-process HLS003: Perform risk analysis (2018-05-07) v1.0"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS003 Perform risk analysis-(2018-05-07)-v1.00.md"
section: "Full risk analysis flow"
chunk: 14
chunk_count: 15
generated: true
---

# Sub-process HLS003: Perform risk analysis (2018-05-07) v1.0 - part 14

Source document: Sub-process HLS003: Perform risk analysis (2018-05-07) v1.0
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS003 Perform risk analysis-(2018-05-07)-v1.00.md`
Chunk: 14 of 15

### Full risk analysis flow

1. **PEE04 — ENS ready for full risk analysis** *(start)*.
2. Screening / information-gathering: **RAT08 — Request HRCM screening** → **RAE02 — HRCM screening results received** → **RAT09 — Assess HRCM screening results**; in parallel **RAT06/RAT10** request additional information from EO/IMS → **RAE01/RAE03** received → **RAT07/RAT11** assess.
   - gateways **Referrals needed … / Additional referrals needed from EO or risk mitigating information from IMS?** control the loops.
3. gateway **Risk identified?**
   - **Yes** → **RAT12 — Set RA result** → **RAT14 — Issue DNL** → **RAE05 — DNL issued**.
   - **No** → **RAT13 — Send assessment complete notification**.
4. **RAT18 — Set control recommendation and place of control** → gateway **Goods to be controlled?**
   - **Yes** → **RAT19 — Inform MS of Control** → (MS of control) **RAE15 — Control recommendation received** → **RAT26 — Set control decision** → **RAE16 — Control decision set**.
5. gateway **AEO to be informed?**
   - **Yes** → **RAT27 — Inform AEO about Controls** (produces Advanced control notification AEO).
6. End: **RAE09 — Risk analysis complete** *(end event)*.
