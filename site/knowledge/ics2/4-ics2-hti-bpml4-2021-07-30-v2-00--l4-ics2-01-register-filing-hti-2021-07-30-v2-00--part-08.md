---
title: "Sub-process L4-ICS2-01: Register filing (HTI) (2021-07-30) v2.00 - part 8"
source_title: "Sub-process L4-ICS2-01: Register filing (HTI) (2021-07-30) v2.00"
source_path: "4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-01 Register filing (HTI)-(2021-07-30)-v2.00.md"
section: "Main process flow (Trader Interface lane)"
chunk: 8
chunk_count: 8
generated: true
---

# Sub-process L4-ICS2-01: Register filing (HTI) (2021-07-30) v2.00 - part 8

Source document: Sub-process L4-ICS2-01: Register filing (HTI) (2021-07-30) v2.00
Source path: `4. ICS2-HTI-BPML4-(2021-07-30)-v2.00/L4-ICS2-01 Register filing (HTI)-(2021-07-30)-v2.00.md`
Chunk: 8 of 8

## Main process flow (Trader Interface lane)

1. **E-01-01 — ENS filing received** *(start)* — triggered by **IE3Fxx E_ENS_xxx_DEC**.
2. → **T-01-01 — Perform syntactical & semantical validation on received ENS filing**
3. → gateway **G-01-01 — Syntactical & semantical validation successful?**
   - **No** → **T-01-06 — Notify error** *(sends IE3N99 E_ERR_NOT)* → **E-01-03 — ENS filing rejected** *(end)*
   - **Yes** → **T-01-02 — Register ENS filing and assign MRN** → **T-01-03 — Submit registered ENS filing**
4. After submission, the registered filing awaits a response:
   - **E-01-06 — ENS filing acceptance notification for Person Filling received** → **T-01-04 — Notify successful registration and MRN to Person filing** *(sends IE3R01 E_ENS_REG_RSP)* → **E-01-02 — ENS filing registration notified to Person Filling** *(end)*
   - **E-01-09 — ENS filing acceptance notification for Carrier received** → gateway **G-01-02 — Carrier to be notified about successful registration?**
     - **Yes** → **T-01-05 — Notify successful registration and MRN to Carrier** *(sends IE3R01 E_ENS_REG_RSP)* → **E-01-10 — ENS filing registration notified to Carrier** *(end)*
     - **No** → *(end)*
   - **E-01-07 — ENS lifecycle validation error notification received** *(IE3N01 E_ELF_VLD_NOT)* → **T-01-09 — Notify ENS lifecycle validation error** → **E-01-08 — ENS lifecycle validation error notified** *(end)*
   - **E-01-11 — ENS filing rejection due to low quality data received** → **T-01-06 — Notify error** → **E-01-03 — ENS filing rejected** *(end)*
