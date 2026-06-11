---
title: "Sub-process L4-ICS2-03-01: Send referral (HTI) (2021-04-30) v1.30 - part 9"
source_title: "Sub-process L4-ICS2-03-01: Send referral (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-03-01 Send referral (HTI)-(2021-04-30)-v1.30.md"
section: "Branch E — notify Carrier about HRCM screening request"
chunk: 9
chunk_count: 10
generated: true
---

# Sub-process L4-ICS2-03-01: Send referral (HTI) (2021-04-30) v1.30 - part 9

Source document: Sub-process L4-ICS2-03-01: Send referral (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-03-01 Send referral (HTI)-(2021-04-30)-v1.30.md`
Chunk: 9 of 10

### Branch E — notify Carrier about HRCM screening request

1. **E-03-01-35 — HRCM screening request notification received** *(message start)*
2. → gateway **G-03-01-02 — Carrier to be notified about HRCM screening request?**
   - **Yes** → **T-03-01-05 — Notify HRCM screening request to Carrier** *(sends IE3N05 E_REF_RFS_NOT)* → *(merge)*
   - **No** → *(merge)*
3. → **E-03-01-36 — HRCM screening request notified** *(end event)*
