
---
title: "TPF awards Hague Grant to Jerry Gay, core 'Rakudo' Perl 6 team member, regarding S19 command line"
author: Richard Dice
type: post
date: 2008-11-09 00:00:00 +0000 UTC
url: "/post/tpf_awards_hague_grant_to_jerr"
categories:
 - Grants
 - Parrot development
 - Perl 6 Development
 - Perl Foundation
 - Sign in

---

The Perl Foundation is pleased to announce the second "Hague Grant":http://www.perlfoundation.org/ian_hague_perl_6_development_grants. It is being awarded to Jerry Gay, core Parrot hacker and 'Rakudo' Perl 6 implementation hacker.  The details of Jerry's grant proposal are below.  The work will be to define the S19 synopsis pertaining to command-line interaction with Perl 6, and to provide a Rakudo implementation of the synopsis.

Jesse Vincent, the project manager of the Perl 6 effort, has agreed to volunteer as the grant manager for Jerry's grant. Jesse will provide updates to TPF on Jerry's status and will judge acceptances of the various milestones and deliverables.  Larry Wall has also agreed to act as the acceptor for the synopsis-definition deliverable of the grant.

We look forward to Jerry's success on this grant project and we are proud to be able to support him in this work.

The details of the grant follow.


<hr>
<p>

Name: Jerry Gay

Project Title: Synopsis 19 - Command Line Syntax, Design and Implementation

Synopsis:
Perl 6 won't feel like Perl until it's callable from the command-line
with a well-defined, practical syntax. This proposal aims to deliver a
design and implementation of the Perl 6 command-line syntax. This project
implements a core component of Perl 6--one which will define how users
interact with Perl 6 every day.


Benefits to Perl 6 Development:
The Perl 6 specification requires the development and testing of
Synopsis 19 - Command Line Syntax (S19). The release of any Perl 6
implementation, including Rakudo Perl, will require an implementation
of this Synopsis. This project will provide the required Synopsis, and
an implementation for Rakudo.


Deliverables:
D1. Documentation of Perl 6 command line syntax working draft in
    Synopsis form (Synopsis 19). This preliminary document will be
    reviewed by the design team and by the community via at least one
    round of feedback. The Synopsis may be revised multiple times before
    final approval; these revisions are outside the scope of this item.

D2. Implementation of a set of tests covering main design points of
    Synopsis 19. These tests are meant to aid in the development of
    the design laid out in the Synopsis--specifically to uncover edge
    cases and ambiguities. The tests are not meant to provide full
    coverage of the Synopsis. The tests will reside in the Pugs
    repository, with the official Perl 6 tests in the t/spec/ directory.

D3. Implementation of a command line parsing library for Parrot.
    This library will be based on Synopsis 19, and much as the Parrot
    Grammar Engine (PGE) is based on Synopsis 5, this library will be
    factored in such a way as to make it useful for Parrot, Rakudo,
    and other high-level languages targeting Parrot. To meet this need,
    the library will become part of the Parrot Compiler Toolkit (PCT).

D4. Implementation of a subset of the Perl 6 command line syntax in
    Rakudo, following the design of Synopsis 19, and based on the Parrot
    command line parsing library (D3). This library will be customized to
    support Perl 6 syntax, and I suspect it will be written in Perl 6 or
    NQP. Implementation of command line options will be limited to those
    which may be created using features available in Rakudo at the time of
    implementation. Tests addressing these features (D2) will be added to
    the suite of regression tests run by Rakudo.

D5. Implementation of the Perl 6 MAIN() subroutine semantics, as given
    in Synopses 6 and 19, and other relevant passages.

D6. Timely reports on implementation progress (probably weekly).


Project Details:
I will pull ideas from various sources during the course of this project,
including:
* Perl 5 command l