
---
title: "Devel::Cover Grant Report (August)"
author: mdk
type: post
date: 2013-09-03 00:00:00 +0000 UTC
url: "/post/develcover_grant_report_august_1"
categories:

---

[Editor's note: Paul has posted 3 Grant reports together due to a very hectic workload and the need to do the work but hold on the reporting of the work. This is a large undertaking and I would like to thank him for his diligence, and tenacity, in filing the paperwork]

In accordance with the terms of my grant from TPF this is the monthly report for my work on improving Devel::Cover covering August 2013.

This month I released Devel::Cover versions 1.07 and 1.08.

Version 1.07 was primarily concerned with getting Devel::Cover to work with recent Perl releases: 5.18.1, 5.19.1, 5.19.2 and 5.19.3.  The main problem was with 5.18.1 which is a stable, maintenance release of Perl and which was causing Devel::Cover the fail its tests rather dramatically.  And the tests were showing up a real problem which stopped Devel::Cover working with Perl 5.8.1 at all.

The original bug report against Perl, filed by Jim Keenan, is RT #119351: Perl 5.18.1 breaks Devel::Cover.  And the analysis of the problem is at https://rt.perl.org/rt3//Public/Bug/Display.html?id=119351#txn-1245355

The minimal test case is:

$ perl5.18.0 -MB -E 'say B::main_cv->GV // "undef"'
B::SPECIAL=SCALAR(0x7b58e0)
$ perl5.18.1 -MB -E 'say B::main_cv->GV // "undef"'
undef

which shows the regression from 5.18.0 and, indeed, from every stable version since 5.6.1, which is the first version that Devel::Cover works with and, I think, the first version which supported this construct.  The reason that Devel::Cover has, and always had, a minimum requirement of Perl 5.6.1 was because that was the first version of Perl which supported sufficient introspection capabilities (B::) to make Devel::Cover possible.

The offending commit was e6c4c33, which was an important fix for a crash related to lexical subroutines, but which had this unfortunate and unintended consequence.

Further discussion in RT #119413: Reconsider e6c4c33 to ext/B/ in light of test failures, suggests that we will see a Perl 5.18.2 release fixing this regression sooner rather than later.

On a personal note, I'm annoyed that I didn't test any of the 5.18.1 release candidates against Devel::Cover.  Doing so could have prevented 5.18.1 being released with this regression.  I do usually test the release candidates against all the modules I use personally or in my work, but this release caught me during a rather busy time.

In any case, irrespective of what happens with upcoming Perl releases, Devel::Cover 1.07 and later work with every stable version from 5.6.1 to 5.18.1, and should work with 5.18.2 whether or not this regression is fixed.

The same regression, unsurprisingly, was in bleadperl, and was causing Devel::Cover to fail with recent development releases.  The fix for 5.18.1 also worked for the development releases, but Perl 5.19.2 and 5.19.3 had another problem.  The problem was that line numbers reported for certain constructs changed and whilst Devel::Cover still worked, the tests were failing because the line numbers against which coverage was reported didn't match.

This problem was reported in Perl RT #118931 and shows that the problem started with 2179133 and was fixed with bf1b738.  But since there are some cpantesters running 5.19.2 and 5.19.3, I caused the problematic tests to be skipped on those versions.  So as of the end of August, The current version of Devel::Cover should pass its tests against all stable versions of Perl since 5.6.1 and all 5.19.x development releases.

As usual, there were also some patches applied, pull requests merged, and other bugs fixed.

As I mentioned in previous reports, I'm still rather busy with other work, but I hope that I will be able to start devoting more time to Devel::Cover towards the end of September.

**<big>Closed Github tickets</big>**

*   66 test failed on Perl 5.18.1
*   67 typo fixes
*   68 Devel::Cover 1.06: many test failures with Perl 5.18.1

**<big>Merged pull requests<