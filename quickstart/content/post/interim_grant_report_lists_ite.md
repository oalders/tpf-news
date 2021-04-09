
---
title: "Interim Grant Report: Lists, Iterators, and Parcels"
author: Karen Pauley
type: post
date: 2012-07-31 00:00:00 +0000 UTC
url: "/post/interim_grant_report_lists_ite"
categories:
 - Grants

---

_Patrick Michaud writes:_

This is an interim report for my Hague Grant on "Lists, Iterators, and Parcels":http://news.perlfoundation.org/2010/07/hague-grant-application-lists.html.  This report has been far too long in coming; at the time of the original grant proposal (July 2010) it was expected that the design and implementation of Perl 6 lists and iterators was near completion, and that most of what remained to be done would be updating the documentation, tests, and Synopses.

Since that time we've essentially learned that the design as of July 2010 still had some holes in it, especially in the area of performance, and thus I have been working to solidify the underpinnings of both.  The lists and iteration implementation in Rakudo 2012.07 is some 40 to 50 times faster than the one that existed at the time of the original grant proposal, and is now written almost entirely in "Perl 6 instead of PIR":http://pmthium.com/2011/06/rakudo-nom-reaches-milestone-no-more-pir-files/.

As of this report, we now have an updated draft of "Synopsis 7":https://github.com/perl6/specs/blob/master/S07-lists.pod that reflects the best understanding of Perl 6 lists and iterators as of July 2012. In addition, I have given several presentations at "YAPC::EU 2011":http://yapceurope.lv/ye2011/talk/3520 and YAPC::NA 2012 ("1":http://act.yapcna.org/2012/talk/123 "2":http://act.yapcna.org/2012/talk/124) on Perl 6 lists and iteration. The draft is still a draft, even as recently as June 2012 there have been "some questions":https://github.com/perl6/specs/issues/18 regarding the funamental semantics of Perl 6 lists that may affect the ultimate design.  I am currently working with people on #perl6 and the Perl 6 design team to resolve the remaining issues and have them fully documented.

Here is the status of the deliverables identified in the original report:

D1: (Completion of the new list implementation in Rakudo Perl 6.)  The list implementation referred to in July 2010 was essentially completed in the following month, but it became quickly apparent that it had performance problems.  In Spring 2011 the list code was completely reimplemented in Perl 6 as part of the "nom" branch of development and was about 2x faster than previously.  Since then several improvements and refactors have been made to basic list and map operations such that list handling is now 40x to 50x faster than it was in 2010.

More changes are left to be made; there are still some spec-related questions regarding the handling of Parcel flattening and combining of lists, which are waiting on input from the Perl 6 design team. I expect these to be resolved within the next six weeks.

D2: (Review and updated tests for iterables and lists in official test suite.)  All of the existing tests have been updated for compliance with the current Lists and Iterables implementation and description.  A new section in the test suite is being created to specifically test the features defined by the new Synopsis 7 draft.  A key focus of this additional section will be to explicitly document and enforce the list-related design decisions that have been made by the design team over the past several years of development.

D3:  (A new version of Synopsis 7.)  A draft of a new Synopsis 7 has been checked into the perl6/specs repository. [3]  Over the next couple of weeks it will be enhanced with detailed descriptions of List-related methods and functions, as well as incorporate any new decisions made by the design team.

D4:  (Update Synopsis 8.)  After working on drafts of Synopsis 7, it seems that it will make more sense to document the Parcel and Capture classes directly in Synopsis 7, rather than in a separate synopsis.  The handling of slices has changed somewhat since July 2010, and now belongs more naturally in Synopsis 9 (which already covers a fair bit on slicing).  The end result of th