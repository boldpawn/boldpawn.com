---
title: "Introduction - part 15"
source_title: "Introduction"
source_path: "3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md"
section: "Business context"
chunk: 15
chunk_count: 65
generated: true
---

# Introduction - part 15

Source document: Introduction
Source path: `3. ICS2-HTI-BPML3.5-(2021-07-30)-v2.00/ICS2-Business Process Descripiton-(2021-07-30)-v2.00.md`
Chunk: 15 of 65

## Business context


*Figure 1: ICS2 Business context*

The purpose of the above figure is to present ICS2 business context and depict different business interfaces between ICS2 components and other customs processes (see Figure 1).

Entry of the goods into the EU can be described as a 5 step process, consisting of the lodgement of the Entry summary declaration (ENS), notification of the arrival of the means of transport, presentation of goods, temporary storage and placing the goods under a customs procedure.

ICS2 business process scope covers 3 steps out of 5: lodgement of the ENS, notification of the arrival of the means of transport and presentation of goods. In addition, this process interacts with other customs processes in the global customs context.

First of all, the ENS filing can be lodged not only directly but also be an input from other customs processes. Art. 130 of UCC stipulates that ENS particulars can be combined with a declaration for temporary storage or a customs declaration for release for free circulation, transit, warehousing, free zones, temporary admission, end use and inward processing. These particulars then are to be extracted and submitted as input to ICS2 process in a form of a full ENS filing.

Secondly, ICS2 process interacts with the reference data management process to obtain the necessary reference data, e.g. EORI number, AEO authorisation, TARIC, Customs Office list, other code lists.

ICS2 process provides inputs to other processes as well. Art. 145 (5) of UCC allows a reuse of the ENS data for temporary storage declaration. Art.144 of UCC DA permits using the ENS data for customs declaration for release for free circulation of the postal consignments. In these cases the complete ENS data will be re-used by the respective declarant for the relevant customs clearance processes.

Within the overall process of ICS2 there are interactions between the different components of the process: ENS data is provided to the risk management process as an input for the risk analysis, which returns the risk analysis results to be fed back into the ENS process. Next step is the control recommendation which is fed into the presentation of goods process from which control results are returned back to the ENS process and, in addition, provided to the risk management process.
