
---
title: "Grant Proposal: Improving the Robustness of Unicode Support in Rakudo on MoarVM"
author: Coke
type: post
date: 2017-04-04 00:00:00 +0000 UTC
url: "/post/grant_proposal"
categories:
 - Grants
 - Sign in

---


The Grants Committee has received the following grant proposal for the March round.
Before the Committee members vote, we would like to solicit feedback from the Perl community on the proposal.

Review the proposal below and please comment here by April 12th, 2017.
The Committee members will start the voting process following that
and the conclusion will be announced approximately in one week.

# Improving the Robustness of Unicode Support in Rakudo on MoarVM

- Name:

    Samantha McVey (samcv)

- Amount Requested:

    7,500 USD

## Synopsis

Implement Unicode Collation Algorithm, improve speed and spec conformance of the
text normalizer. Improve test coverage for Unicode specs and document our
compliance or lack of compliance with the Unicode spec.

## Benefits to the Perl Community

As Perl 6 starts to take off, it is increasingly important to provide robust
Unicode support. Perl 6 already provides some of the best Unicode support on
 many levels compared to other programming languages. The goal of this project
 is to make Perl 6's Unicode support production ready.

## Deliverables and Inch Stones

- General
    - Any deficits of our Unicode coverage will be documented in the course of work on this
      project. This is very important
      due to the vastness of the Unicode standard. Deficits will have tests
      written, unless such a thing would not be possible to test or input is needed
      from the rest of the Perl 6 team. In any of these cases, they will be
      documented in my reports for future and current developers of Perl 6 to reference.
    - Tests will be written to cover all of the relevant Unicode 9.0 test files, as
      well as making current ones more robust when checking the breaking of graphemes.
- Unicode Names
    - Hangul Syllables and other Unicode names shall be programmatically determined
      when generating the Unicode database.
- Unicode Collation Algorithm
    - Fully implement the Unicode Collation Algorithm at least for language
      nonspecific sorting.
    - Assess needs in supporting language and country specific collation, and write
     a report of these needs.
- Text Normalization
    - Improve the performance of the text normalizer and also allow the normalizer
      to save state across multiple characters to properly support Grapheme Breaking
      for all of Unicode 9.0 and beyond.
- Unicode Database Generation
    - The script used to generate the Unicode database shall be made deterministic,
      and produce the same output file on every run.
      At the current time about half of the file changes even if no changes are made to the
      script. This is an issue that will be solved.
    - In addition to the above, the original script assumed that property values for
    Unicode characters were unique. This causes issues when there is a conflict between
    these. The rewritten script shall resolve this problem.
    - Rewrite the Perl 5 script used to generate the Unicode database in Perl 6.
      This is also part of the previous item, since a rewrite is needed, it should be
      done in Perl 6 to help make it more maintainable.
    - Implement all relevant remaining Unicode properties from Unicode 9.0. This
      includes the properties needed to support the deliverables listed above.
    - Try to reduce the memory footprint of the Unicode database. Currently
      the unicode.o binary file created is 4.1MB. I hope to cut that in half.

## Project Details

I have already started working on the rewrite of the Unicode database generation
which is in [this public repository](https://github.com/samcv/UCD).

Some background into this. The original Unicode database generation script we
currently use was added in 2012 and much of the script has not changed much since.
Although this current script is somewhat adequate, as I have been working on
Unicode support these past months it became clear that the script was not easily
maintainable and had some issues which would have required an extensive
rewrite of much of the script. It became clear that a full rewrite was prudent.

I have done lots of fielding and preliminary work as well as the knowledge I have
gained by working with the current script we use and the MoarVM internals. 

### Plan

#### Month 1: Get database generation finalized

- Implement testing for resolving codepoints to bitfield rows, to ensure
  all codepoints resolve to the correct rows
- Integrate testing of database values into the current Unicode database rewrite
 (this will not be in roast, but will be done to ensure correctness during 
 development)
- Work with timotimo on C implementation on indexed decoding of base40 compressed
  Unicode Names (some work already done in [UCD repo](https://github.com/samcv/UCD)
- Implement programmatic generation of codepoint names for Hangul Syllables
- Implement remaining properties and data in the rewrite and achieve pairity with
  our current script

#### Month 2: Finish the Unicode Collation Algorithm and integrate with MoarVM, add test coverage to roast

- Assess needs for language/country specific sorting for Unicode Collation Algorithm (UCA)
- Implement weights of codepoint sequences (only single codepoint is implemented currently)
- Implement decomposition of codepoints if no Collation weights are found
- Integrate the rewrite with MoarVM codebase, as well as rewrite all sections
  of MoarVM
- Make needed changes in MoarVM so Property Values are not assumed unique.
- Integrate with MoarVM codebase
- Merge into MoarVM repository

## Project Schedule

I can begin work as soon as the grant is approved.

## Completeness Criteria

Tests will be commited to roast, and the rewritten database and Unicode Collation
Algorithm implementation will be merged into MoarVM.

## Bio

Although I am a fairly recent addition to the Perl 6 core developers, in a short
few months I have been very busy. I have two Perl 6 modules, IRC::TextColor
and URL::Find and I am the lead developer of the Perl 6 syntax highlighter
for Atom/Github as well as for docs.perl6.org. I converted the site from
using the old Pygments highlighter to the new highlighter.

My contributions to Perl 6 have been focused on Unicode support in Perl 6,
making changes throughout Rakudo, NQP and MoarVM to achieve this. All of the
work I have already done on improving Unicode support in Perl 6 shows I am
capable of completing this project and am the best person for this grant.
In addition, I have already started work on rewriting the Unicode Database
generation and shrinking the size of the data needed to be loaded on startup.

### Incomplete list of Unicode work I have done in the last few months

#### Tests:

- Fixed several errata in roast related to our Unicode support, which had often
  been present for a long time.
- Added a test based on GraphemeBreakTest.txt from Unicode and many others
  to Unicode 9.0
- Updated other tests for Unicode 9.0 and reworked others for compliance.

#### MoarVM:

- Implemented a simplistic implementation of the Unicode Collation Algorithm.
- Added `coll` operator, `.collate` method and the `$*COLLATION` variable
to Rakudo. For more information see the
[documentation](https://docs.perl6.org/language/experimental#Collation) I have added on it.
- Added support for named codepoint sequences, which includes the Named Sequences,
  Emoji Sequences and Emoji ZWJ Sequences. `"\c[woman gesturing OK]"`
- Implemented
[Unicode Name Aliases](https://github.com/MoarVM/MoarVM/commit/81618648) in
getting codepoints by name and [documentation](https://docs.perl6.org/language/unicode#Name_Aliases)
- Implemented the 'Extend' Grapheme\_Cluster\_Break property which was new in Unicode 9.0.
  We previously had no support for this property. This caused errors in grapheme
  breakup and incorrect character count and segmentation.
- Implemented many other Grapheme\_Cluster\_Break fixes and added support for
  most Emoji sequences. This fixes most Emoji made up of multiple codepoints,
  fixing character counts and text segmentation for these extended grapheme clusters.
- Improved the speed of radix 50% for non-ASCII decimal digits
(converts strings to their numeric representation/values)
- Improved the speed of text normalization, making slurping a Unicode heavy text
  file 14% faster
- Improved the speed m:i/ / regex matching between 1.8x and 3.3x (depending on not finding a match / finding a match at the beginning).
- Added a multitude of properties to our Unicode database.

#### Rakudo:

- Added support for a large number Unicode properties, handling Bool/Str/Int
  return types for `uniprop`
- Implemented `uniprops` method in Rakudo


