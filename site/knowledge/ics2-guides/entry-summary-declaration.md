---
title: "Guide: Entry Summary Declaration (ENS)"
description: "Curated guide for answering detailed questions about Entry Summary Declaration, ENS filings, lifecycle, messages, and related ICS2 HTI source documents."
tags:
  - ics2
  - ens
  - entry-summary-declaration
  - hti
  - customs
---

# Guide: Entry Summary Declaration (ENS)

Use this guide as the starting point for questions about the Entry Summary Declaration in ICS2. It is a curated map over the generated source chunks in `knowledge/ics2/`. Start here, then read the cited source chunks for the exact table rows, message structures, rules, and process steps.

## Canonical Terms

Entry Summary Declaration is abbreviated as ENS.

Important synonyms and related terms:

- Entry Summary Declaration
- ENS
- ENS filing
- full ENS filing
- partial ENS filing
- complete ENS
- ENS lifecycle
- master level ENS filing
- house level ENS filing
- PLACI
- MRN
- person filing
- carrier
- Trader Interface
- Common Repository
- risk analysis
- referral
- amendment
- invalidation
- consultation

## Core Meaning

In the ICS2 definitions, Entry Summary Declaration (ENS) is the act by which a person informs customs authorities, in the prescribed form, manner, and time limit, that goods are to be brought into the customs territory of the Union.

An ENS filing is the dataset submitted for that declaration. It can be a full filing or a partial filing, depending on the transport mode, business model, and legal data requirements. A complete ENS can be represented either by one full ENS filing or by multiple partial ENS filings that together provide the required particulars.

Primary source:

- `knowledge/ics2/ics2-hti-definitions-2020-03-16-v1-11--part-11.md`

## Mental Model

Think about ENS in three layers:

1. Legal/business concept: the Entry Summary Declaration obligation and its required data.
2. Lifecycle/process concept: the filing is lodged, registered, linked or completed, assessed for risk, possibly amended or invalidated, and later connected to arrival, presentation, controls, and consultation.
3. Message/data concept: the HTI exchanges concrete XML-like message structures such as IE3Fxx declarations, IE3Axx amendments, IE3R01 registration responses, and IE3Nxx notifications.

When answering elaborate questions, combine all three layers. A good answer should explain the concept, identify the process step, and then point to the relevant messages or data elements.

## HTI Scope

The HTI specifications cover lodgement and registration of ENS filings, full or partial, and arrival notifications through the Trader Interface. They also cover exchanges between economic operators and customs administrations, referrals, ENS consultation, and invalidation and amendment of ENS filings.

The HTI package does not cover every possible ENS-related customs declaration scenario. The main document explicitly excludes customs declarations containing ENS particulars under Art. 130 UCC, postal-item ENS filings for non-air modes, and express-consignment ENS filings for non-air modes.

Primary source:

- `knowledge/ics2/ics2-hti-main-2022-05-23-v2-02--part-08.md`

## Lifecycle Overview

The ENS lifecycle is the end-to-end status flow of an ENS from filing through final states such as presentation, with risk analysis, controls, and control-result documentation where relevant.

Typical lifecycle path:

1. ENS filing is received through the Trader Interface.
2. Syntactical and semantical validation is performed.
3. If validation fails, the filing is rejected and an error notification is sent.
4. If validation succeeds, the filing is registered and an MRN is assigned.
5. The registered filing is submitted onward and the person filing, and sometimes the carrier, receive registration or acceptance information.
6. The ENS is prepared for risk analysis, including completeness and linking checks when partial filings must be related.
7. Customs risk analysis may produce assessment completion, referral requests, DNL, HRCM screening requests, control notifications, or related notifications.
8. Later flows can amend, invalidate, consult, arrive, present, or control the ENS depending on state and business scenario.

Useful source chunks:

- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-00-master-process-hti-2018-10-15-v1-10--part-04.md`
- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-01-register-filing-hti-2021-07-30-v2-00--part-03.md`
- `knowledge/ics2/3-ics2-hti-bpml3-5-2021-07-30-v2-00--ics2-hls002-prepare-ens-for-risk-analysis-2021-07-30-v2-00--part-05.md`
- `knowledge/ics2/3-ics2-hti-bpml3-5-2021-07-30-v2-00--ics2-hls002-prepare-ens-for-risk-analysis-2021-07-30-v2-00--part-06.md`

## Registration

The L4 register-filing process begins when an ENS filing declaration is received as an IE3Fxx message. The Trader Interface performs syntactical and semantical validation. If successful, it registers the ENS filing, assigns an MRN, submits the registered filing, and notifies the person filing. The carrier can also be notified when the conditions for carrier notification are met.

Main messages:

- IE3Fxx / E_ENS_xxx_DEC: incoming ENS filing declaration
- IE3R01 / E_ENS_REG_RSP: ENS registration response
- IE3N99 / E_ERR_NOT: error notification
- IE3N01 / E_ELF_VLD_NOT: ENS lifecycle validation error notification

Primary source:

- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-01-register-filing-hti-2021-07-30-v2-00--part-03.md`

## Completeness and Relating Partial Filings

An ENS can be complete because one full ENS filing contains all required particulars, or because multiple partial filings together contain the required particulars. This is why the sources distinguish full ENS filing, house level ENS filing, master level ENS filing, complete ENS, and Unique Linking Key.

When partial filings are involved, related filings and house consignments matter. The system can notify that an ENS is not complete or that an ENS is pending because another person has not yet filed. These flows are especially important when answering questions about completeness, pending ENS filings, linking, or responsibilities across parties.

Useful source chunks:

- `knowledge/ics2/ics2-hti-definitions-2020-03-16-v1-11--part-11.md`
- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-02-01-relate-ens-filings-hti-2021-04-30-v1-30--part-02.md`
- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-02-01-relate-ens-filings-hti-2021-04-30-v1-30--part-03.md`

Important messages:

- IE3N02 / E_ENS_NCP_NOT: ENS not complete notification
- IE3N11 / E_ENS_PND_NOT: ENS pending notification

## Risk Analysis and Referrals

After registration and preparation, the ENS is used for customs risk analysis. Risk analysis can lead to assessment completion, additional information requests, high-risk cargo and mail screening requests, do-not-load outcomes, control notifications, and other referral or notification flows.

Useful source chunks:

- `knowledge/ics2/3-ics2-hti-bpml3-5-2021-07-30-v2-00--ics2-hls003-perform-risk-analysis-2018-05-07-v1-00--part-01.md`
- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-03-perform-risk-analysis-hti-2021-04-30-v1-30--part-01.md`
- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-03-01-send-referral-hti-2021-04-30-v1-30--part-01.md`

Relevant messages from the IE overview:

- IE3Q01 / E_DNL_REQ: Do Not Load request
- IE3Q02 / E_REF_RFI_REQ: additional information request
- IE3Q03 / E_REF_RFS_REQ: HRCM screening request
- IE3R02 / E_REF_RFI_RSP: additional information response
- IE3R03 / E_REF_RFS_RSP: HRCM screening response
- IE3N03 / E_ASM_CMP_NOT: assessment complete notification
- IE3N04 / E_REF_RFI_NOT: additional information request notification
- IE3N05 / E_REF_RFS_NOT: HRCM screening request notification

## Arrival, Presentation, and Controls

Arrival and presentation processing connect the ENS to the actual arrival of the means of transport and goods. The system retrieves relevant ENS information for the journey, checks state, and can move the ENS to arrived, partially arrived, or presented states depending on the scenario.

Useful source chunks:

- `knowledge/ics2/3-ics2-hti-bpml3-5-2021-07-30-v2-00--ics2-business-process-descripiton-2021-07-30-v2-00--part-38.md`
- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-04-process-arrival-of-means-of-transport-hti-2021-04-30-v1-30--part-01.md`
- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--ics2-hti-bpml4-process-description-2021-07-30-v2-00--part-40.md`

## Amendment

ENS filing amendments are submitted through IE3Axx messages. The amendment process validates the amendment, submits valid amendments, rejects invalid ones, and handles completion or lifecycle validation notifications.

Important points:

- IE3Axx / E_ENS_xxx_AMD is the incoming amendment message family.
- IE3N10 / E_AMD_NOT can notify amendment completion.
- IE3N01 / E_ELF_VLD_NOT can notify lifecycle validation errors.
- IE3N99 / E_ERR_NOT can notify errors.
- The information-exchange main document contains principles for how amendments replace or affect previous filings.

Useful source chunks:

- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-07-amend-ens-filing-hti-2021-07-30-v2-00--part-03.md`
- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-07-amend-ens-filing-hti-2021-07-30-v2-00--part-04.md`
- `knowledge/ics2/5-ics2-hti-ie-2022-05-04-v2-02--4-ics2-hti-ie-2022-05-04-v2-02--part-18.md`

## Invalidation

Invalidation is a separate lifecycle operation from amendment. It has its own request and response messages and must respect lifecycle validation. Use invalidation sources when questions ask whether an ENS filing can be withdrawn, cancelled, invalidated, or rejected after registration.

Useful source chunks:

- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-08-invalidate-filing-hti-2021-04-30-v1-30--part-03.md`
- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-08-invalidate-filing-hti-2021-04-30-v1-30--part-06.md`

Relevant messages:

- IE3Q04 / E_INV_REQ: invalidation request
- IE3R07 / E_INV_ACC_RSP: invalidation acceptance response
- IE3N01 / E_ELF_VLD_NOT: lifecycle validation error notification

## Consultation

ENS consultation is a supporting process for economic operators to request information about specific ENS filings or related entities. Consultation can be based on MRN, master level transport document, or house level transport document. The results can include relevant notifications and ENS-related entities, subject to access rights.

Useful source chunks:

- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--l4-ics2-13-consult-ens-hti-2018-05-07-v1-00--part-02.md`
- `knowledge/ics2/4-ics2-hti-bpml4-2021-07-30-v2-00--ics2-hti-bpml4-process-description-2021-07-30-v2-00--part-70.md`

Relevant messages:

- IE3Q05 / E_ENS_CNS: ENS consultation request
- IE3R08 / E_ENS_CNS_RES: ENS consultation results

## Message Families

The information exchange overview is the best source for mapping message IDs to process usage.

Key families:

- IE3Fxx / E_ENS_xxx_DEC: ENS filing declarations
- IE3Axx / E_ENS_xxx_AMD: ENS filing amendments
- IE3Qxx: requests, including DNL, referrals, invalidation, and consultation
- IE3Rxx: responses, including registration response, referral responses, invalidation acceptance, and consultation results
- IE3Nxx: notifications, including lifecycle validation, not complete, assessment complete, referral notifications, and pending notifications

Primary source:

- `knowledge/ics2/5-ics2-hti-ie-2022-05-04-v2-02--4-ics2-hti-ie-2022-05-04-v2-02--part-16.md`

## How to Answer Common Questions

For "What is an Entry Summary Declaration?":

1. Use the definition in `ics2-hti-definitions...part-11.md`.
2. Distinguish ENS from ENS filing.
3. Explain full, partial, and complete ENS.
4. Mention that the HTI specifications cover lodgement, registration, amendment, invalidation, consultation, and related customs exchanges.

For "How is an ENS submitted and registered?":

1. Read the L4 register-filing process.
2. Explain IE3Fxx receipt, syntactical and semantical validation, MRN assignment, submission, registration response, and rejection paths.
3. Include carrier notification only where applicable.

For "How does ENS risk analysis work?":

1. Start with preparation for risk analysis and completeness.
2. Then read the risk-analysis and referral process chunks.
3. Explain assessment completion, additional information requests, HRCM screening, DNL, and control notifications as possible outputs.

For "What is the difference between full and partial ENS filing?":

1. Use definitions for complete ENS, full ENS filing, house level ENS filing, master level ENS filing, and ULK.
2. Explain that multiple partial filings can compose a complete ENS.
3. Use relate-ENS-filing chunks for not-complete and pending notifications.

For "Which messages are involved?":

1. Use the IE overview chunk first.
2. Then read the individual IE3Fxx, IE3Axx, IE3Qxx, IE3Rxx, or IE3Nxx generated chunks for detailed data elements.

## Search Recipes

Use these query phrases with `search_content`:

- `Entry Summary Declaration ENS definition`
- `Complete ENS full partial filing`
- `ENS filing received register MRN IE3Fxx IE3R01`
- `Prepare ENS for risk analysis completeness linking`
- `Relate ENS filings not complete pending IE3N02 IE3N11`
- `ENS amendment IE3Axx lifecycle validation`
- `ENS invalidation IE3Q04 IE3R07`
- `ENS consultation IE3Q05 IE3R08 MRN transport document`
- `ENS arrival state arrived partially arrived presented`
- `ENS filing data elements IE3F`
- `ENS amendment data elements IE3A`

## Answer Quality Checklist

For elaborate answers, include:

- The exact term and abbreviation.
- Whether the question is about the legal concept, a filing dataset, a process, a lifecycle state, or a message.
- The actor or component involved: Person filing, carrier, Trader Interface, Common Repository, National Entry System, customs authority, or economic operator.
- The relevant process: registration, preparation, risk analysis, referral, arrival, amendment, invalidation, or consultation.
- The relevant message IDs and names.
- Whether the answer concerns full ENS, partial ENS, complete ENS, house level, master level, or PLACI.
- Source chunk paths that support the answer.

