
---
title: "TPF awards first Hague Grant to Patrick Michaud, head of the 'Rakudo' Perl 6 implementation effort"
author: Richard Dice
type: post
date: 2008-11-09 00:00:00 +0000 UTC
url: "/post/tpf_awards_first_hague_grant_t"
categories:
 - Awards
 - Grants
 - Parrot development
 - Perl 6 Development
 - Perl Foundation
 - Sign in

---

The Perl Foundation is pleased to announce the first "Hague Grant":http://www.perlfoundation.org/ian_hague_perl_6_development_grants.  It is being awarded to Patrick Michaud, the head of the 'Rakudo' Perl 6 implementation effort on top of the Parrot VM.  Conceptually it is an extension of Patrick's earlier Mozilla Foundation / Perl Foundation Perl 6 development grant that he worked on between late 2007 and mid 2008.

The details of Patrick's grant proposal are below.  Jesse Vincent, the project manager of the Perl 6 effort, has agreed to volunteer as the grant manager for Patrick's grant.  Jesse will provide updates to TPF on Patrick's status and will judge acceptances of the various milestones and deliverables.

We look forward to Patrick's continued success on the Rakudo Perl 6 implementation and we are proud to be able to support him in this work over the next 4 months.

The details of the grant follow.



<hr>

Name:           Patrick R. Michaud

Project Title:  Rakudo Perl and PCT improvements

Synopsis:
This project will aim to coordinate and implement many key features
needed for Rakudo Perl, the implementation of Perl 6 on Parrot.  In
particular, this project aims to continue and extend the work performed
under my existing grants from the Mozilla Foundation and TPF.
As with the the Mozilla/TPF grants, the overarching goal of this
project will be to (1) focus on areas that enable others to more
easily and directly participate in Perl 6 development, and (2) to
work on features critical to Perl 6 implementation for which I
have special skills.  These critical features include compiler
architecture and design, protoregex and longest token matching
in Perl 6 regular expressions, Rakudo alignment with STD.pm,
builtin and external library development, and continued
work on the official test suite.


Benefits to Perl 6 Development:
All of the items listed in the synopsis above are on the critical
path for a Perl 6 release implementation.  In particular, it is
vital that the core parsing and builtin library features are
correct with respect to the Synopses and STD.pm specifications;
achieving these requires some intensive development and coordination
effort that will be provided by this project.


Deliverables:
D1. Implementation of protoregexes and longest token matching in
    Parrot's Parser Grammar Engine (PGE).  This is needed for completion
    of regex features described in Synopsis 5, as well as the ability
    to use STD.pm in deliverable #2 below.

D2. Alignment of Rakudo's parser and grammar with STD.pm.  STD.pm
    is the reference model for parsing Perl 6 programs, Rakudo Perl
    must be consistent with STD.pm to be an accurate implementation of
    Perl 6.

D3. Continue coordination and development of Perl 6 builtins for Rakudo,
    and investigate development of a common p6-based "Prelude" that can
    be used as a reference (or even directly) by other Rakudo and
    other Perl 6 implementations.  These builtins come from the
    Synopses and form the basis of any working Perl 6 environment.

D4. Develop and improve Rakudo's ability to make use of external
    libraries, including those written in Perl 6 and other Parrot
    target languages.

D5. Continue developing, reviewing, and refactoring the official
    Perl 6 test suite and coordinating test suite features among 
    various implementations.

D6. Create additional tests and test suites for the PGE and
    library components developed under the project.

D7. Publish weekly reports on Perl 6 implementation progress.  This will
    certainly include information about Rakudo Perl, language spec changes,
    and test suite improvements, but can also include information from
    IRC discussions and other Perl 6 implementation efforts.  This is
    needed to continue visibility into Perl 6 progress and to recruit
    more contributors to Perl 6 development.


Project Details:
The work performed under the Mozilla Foundation/TPF grant has been
highly effective at building development momentum for Rakudo Perl
(Perl 6 on Parrot).  The activities listed under this project seek
to continue and dramatically extend the work performed under that
grant.  In the past few months many more of the precise details needed
for a Perl 6 release have come into focus -- these include:
  * a correct, efficient implementation of parsing tools utilizing
    the concepts embodied in Synopsis 5 and STD.pm,
  * a cleanly organized set of builtin functions and classes that
    function properly and that others can understand and extend, and
  * continued oversight, coordination, and reporting of compiler,
    library, and test suite implementation efforts.

In particular, this project and its deliverables follow the
same general principles of the Mozilla/TPF grants; in particular:
(1) implement key core components (e.g., longest-token-matching)
that are necessary to a Perl 6 compiler and require intensive
or specialized design or development effort, (2) clear obstacles
and coordinate development efforts to allow others to be active
participants, and (3) promote awareness of progress on Perl 6 to
bring more contributors into the process.


Project Schedule:
The project will take approximately four calendar months.  It will 
begin as soon as possible after the conclusion of the Mozilla/TPF 
grant in July 2008.  The test suite development and reporting tasks 
(D5-D7) are ongoing throughout the project, other work is divided 
into four specific "monthly" milestones:

M1. PGE internal refactors and initial protoregex implementation (D1),
    selected protoregex constructs added to Rakudo's grammar (D2),
    interface design for pre-compilation and external libraries (D3, D4).

M2. Completed protoregex implementation (D1),
    initial implementation of longest token matching in PGE (D1),
    completed Rakudo grammar migration to protoregexes (D2),
    beginning of P6-based "Prelude" implementation for Rakudo (D3),
    initial implementations of external HLL library support (D4).

M3. Substantially complete longest token matching in PGE (D1),
    Rakudo using protoregexes and longest token matching in grammar (D2),
    full protoregex and LTM support in Rakudo (D3),
    integrate P6-based Prelude into Rakudo implementation (D3),
    continued improvement of HLL library support (D4).

M4. Completed PGE protoregex and longest token matching implementation (D1),
    update Rakudo bottom-up parser to align with STD.pm (D2),
    Rakudo and/or compiler tools able to parse STD.pm directly (D2),
    substantial implementation of a P6-based Prelude for Rakudo (D3),
    substantially complete external HLL libraries support (D4).
    updated documentation and tests for all work performed, final report.

Throughout the project we should also see substantial improvements
in overall Rakudo Perl compiler performance as a direct result of 
the work performed in this project.


Report Schedule:
In addition to the weekly "Perl 6 Status" reports given above (D7),
at the end of each project milestone (approximately monthly) there 
will be a progress report detailing progress on the corresponding
specific deliverables listed above.  These milestone reports will 
also serve as checkpoints to authorize interim project completion
payments.


Public Repository and License:
In general, all code and documents will be developed and 
stored either in the Parrot or Pugs subversion repositories as
appropriate.  Work on the official Perl 6 test suite will
be committed to the Pugs repository; improvements to PGE
and Rakudo will be committed to the Parrot repository with
any copyrights to be held by The Perl Foundation.  All work
will be licensed under the Artistic License 2.0 .


Bio:
Since 2004 I have been the lead architect, developer, and project
manager for what is now Rakudo Perl (Perl 6 on Parrot), the Parser
Grammar Engine (PGE), and the Parrot Compiler Toolkit (PCT).  My
previous work on these and other components provides me with unique
experience and insight into the critical needs for designing and
the core Perl 6 features described by this project.


Amount Requested:
4 months at US$5,000/month  == US$20,000

Country of Origin:
United States

Okay to publish proposal? :
Absolutely.


