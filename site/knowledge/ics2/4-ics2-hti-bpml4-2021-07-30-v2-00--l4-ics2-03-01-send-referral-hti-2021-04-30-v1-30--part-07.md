---
title: "Sub-process L4-ICS2-03-01: Send referral (HTI) (2021-04-30) v1.30 - part 7"
source_title: "Sub-process L4-ICS2-03-01: Send referral (HTI) (2021-04-30) v1.30"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-03-01 Send referral (HTI)-(2021-04-30)-v1.30.md"
section: "Branch C — notify Carrier about additional information request"
chunk: 7
chunk_count: 10
generated: true
---

# Sub-process L4-ICS2-03-01: Send referral (HTI) (2021-04-30) v1.30 - part 7

Source document: Sub-process L4-ICS2-03-01: Send referral (HTI) (2021-04-30) v1.30
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-03-01 Send referral (HTI)-(2021-04-30)-v1.30.md`
Chunk: 7 of 10

### Branch C — notify Carrier about additional information request

1. **E-03-01-32 — Request for additional information notification received** *(message start)*
2. → gateway **G-03-01-01 — Carrier to be notified about additional information request?**
   - **Yes** → **T-03-01-02 — Notify additional information request to Carrier** *(sends IE3N04 E_REF_RFI_NOT)* → *(merge)*
   - **No** → *(merge)*
3. → **E-03-01-33 — Request for additional information notified to Carrier** *(end event)*
