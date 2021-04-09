
---
title: "Grant Proposal: Curating Perl 6 Documentation"
author: Coke
type: post
date: 2019-09-17 00:00:00 +0000 UTC
url: "/post/grant_proposal_curating_perl_6"
categories:
 - Sign in

---

The Grants Committee has received the following grant proposal for the September/October 2019 round.
Before the Committee members vote, we would like to solicit feedback from the Perl
community on the proposal.

Due to my delay in posting this, I'm extending the comment period.
Please review the proposal below and please comment here by September 27th, 2019.
The Committee members will start the voting process following that.

- Name: Richard Hainsworth (on IRC #perl6 as finanalyst)
- Amount requested: $5000

# Synopsis
There is a need for full time work on Perl 6 Documentation. A new system
is being implemented with the inevitable problems. There are some 300
open issues on the github repo for perl6/doc.

When considering Perl 6 documentation, it is important to distinguish
between several issues, some of which are easy to resolve, others
involve matters of language design and arose emotion. This curation
proposal is to facilitate progress on each of these issues, but will
directly seek to close certain types of issues.

The [currently on-going documentation
redesign](https://github.com/perl6/doc/issues/2983) will split the
existing `perl6/doc` repo into at least four new repos. One repo will be
dedicated solely to the content of the documents, and the others to the
software for processing and testing the documents. This means that the
300 or so issues already registered will need to be reassigned to
appropriate repos.

Some of the 300 issues will disappear as a result of the split.

This proposal is in essence a continuation of the work done by JJ Merelo
under a previous Perl Foundation Grant, see his [final
report](https://jj.github.io/TPF-Grant/Final), together with his
continuing development and the GSoC projects.

Several contributors to Perl 6 regularly update the documentation. This
is a process that should be facilitated by curation. In the categories
of issues below, some of the issues naturally generate emotion. The task
of a curator is not to make decisions and make changes arbitrarily, but
to identify what sort of changes are being proposed, and hence
facilitate progress.

# Usefulness to Perl 6 community
Documentation is essential to any language. Perl 6 is fairly well
documented, but given the size of the language, there remain parts where
documentation does not exist or is not systematic. This is an on-going
process and works quite well.

However, documenting something does not mean the documentation is easily
available. There remain issues with the delivery to the user of the
documentation. In essence, the documentation site is in need of some TLC.

JJ Merelo's
[proposal](http://news.perlfoundation.org/2018/02/grant-proposal-curating-and-im.html)
already sets out much of the case for this proposal. This proposal was
made at the suggestion of JJ.

# Categories of Issues
- Documenting Perl 6 as used by a programmer, which can be divided into
     - questions that pertain to the Perl 6 language
     The documentation is constantly being updated and improved.
     The curation proposal is to facilitate answers being included into
the documentation. Sometimes, there is an issue where documented
behaviour, eg., sample code, does not work.

    - the release of 6.d required changes to the documentation, some of
these issues still need resolving.

     - questions that pertain to the Rakudo implementation
     There was a debate about whether or how to document classes and
routines that are implementation specific. The issue - I understand from
JJ - has been resolved and everything is being documented, with Rakudo
specific classes etc, being mentioned.

- The software that renders documentation content into ** html ** for
the online version
The software for doing this was created in an organic manner and has
become fragile. Considerable effort has already been spent on
refactoring the software. There is a milestone issue th