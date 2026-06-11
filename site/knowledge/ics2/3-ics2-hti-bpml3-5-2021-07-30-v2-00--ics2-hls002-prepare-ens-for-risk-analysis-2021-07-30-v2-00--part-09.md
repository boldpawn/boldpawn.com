---
title: "Sub-process HLS002: Prepare ENS for risk analysis (2021-07-30) v2.00 - part 9"
source_title: "Sub-process HLS002: Prepare ENS for risk analysis (2021-07-30) v2.00"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS002 Prepare ENS for risk analysis-(2021-07-30)-v2.00.md"
section: "Main process flow"
chunk: 9
chunk_count: 9
generated: true
---

# Sub-process HLS002: Prepare ENS for risk analysis (2021-07-30) v2.00 - part 9

Source document: Sub-process HLS002: Prepare ENS for risk analysis (2021-07-30) v2.00
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-HLS002 Prepare ENS for risk analysis-(2021-07-30)-v2.00.md`
Chunk: 9 of 9

## Main process flow

1. **RFE02 — Filing registered** *(start)* → **GET08 — Cleanse ENS filing** → **GET01 — Perform syntax & semantic validation** → gateway **Syntax & semantic validation successful?**
   - **No** → **GET03 — Notify error** → **GEE01 — Filing rejected** *(end; produces Error notification)*
   - **Yes** → **GET09 — Check stops words/phrases** → gateway **Sufficient meaningful data?**
     - **No** → **GET03 — Notify error** → **GEE01 — Filing rejected** *(end)*
     - **Yes** → **GET04 — Perform ENS lifecycle validation** → gateway **ENS lifecycle validation successful?**
       - **No** → **GET03 — Notify error** → **GEE01 — Filing rejected** *(end)*
       - **Yes** → **PET01 — Send filing acceptance notification** → **PET04 — Manage timer for completion of risk analysis**

2. Linking / completeness branch — entered via **PET04** and via **AFE03 — Amended ENS filing** *(start)* and **PEE05 — Timer for completion of linking expired** *(timer start)*:
   - gateway **Master filing?** routes to **PET09 — Link related filings and check completeness**, **PET10 — Start linking expiration timer**, **PET11 — Check data consistency**, **PET12 — Relate ENS filings**.
   - gateway **Found potentially related ENS filings?** determines whether ENS filings are related (**PET12 — Relate ENS filings**) or the filing is held as **PEE03 — House filing pending linking**.
   - gateway **Full filing? / Complete ENS?**:
     - if not complete → **PET13 — Notify ENS not complete** → **PEE06 — ENS not complete** *(end; produces ENS not complete notification)*, or held as **PEE02 — ENS pending completeness**.
     - if complete → continue to risk-analysis preparation.

3. Air pre-loading risk-analysis branch:
   - gateway **Applicable to air pre-loading risk analysis?**
     - **Yes** → **PET05 — Determine applicability for air pre-loading risk analysis** → **PET06 — Extract relevant data for air pre-loading risk analysis** → **PEE01 — 7+1 data ready for air pre-loading risk analysis**, then **PET07 — Determine IMS** → **PET08 — Start timers for IMS risk analysis tasks completion** / **PET17 — Start timer for RMS risk analysis completion**.
     - gateway **Is an Air pre-loading risk analysis on-going?** → **RAE04 — Air pre-loading risk analysis complete** when applicable.
     - **No** → **PET07 — Determine IMS** → **PET08 — Start timers for IMS risk analysis tasks completion** / **PET17 — Start timer for RMS risk analysis completion**.

4. End state: **PEE04 — ENS ready for full risk analysis** *(end event)*.
