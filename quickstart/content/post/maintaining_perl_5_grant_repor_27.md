
---
title: "Maintaining Perl 5: Grant Report for February 2016"
author: Karen Pauley
type: post
date: 2016-04-08 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_grant_repor_27"
categories:

---

_Tony Cook writes:_

Approximately 57 tickets were reviewed or worked on, and 14 patches were applied.

[perl #126410] and [perl #124387] are related in that both problems are related to the same set of changes.

Between perl 5.8 and 5.16 inclusive, perl cached the CV for DESTROY in overload magic for the stash.  This was four times faster than the simple method lookup previously done each time an object needed to be destroyed.

This was much faster for destruction, but meant that overload magic was created for any stash that objects were destroyed for, taking up 576 bytes (on x64) for a stash with a DESTROY method, even if it did no overloading.



In 5.18 this mechanism was changed to store the DESTROY CV in the stash's stash slot, and stored (CV*)1 if there was no DESTROY method.

To ensure that a stale CV wasn't used if the DESTROY method was changed, it kept the current PL_sub_generation in the stash's MRO metadata, and checked it on each call to DESTROY.

This caused two problems:

[perl #126410] Code that checked the stash slot unexpectedly found (CV*)1 in the slot, which is non-NULL, and so attemptted to treat it as a hash.  Which crashed:

bc.  $ perl -Mversion -MDevel::Peek -e 'version->new("1.0"); Dump(\%version::)'
 SV = IV(0x1db4e68) at 0x1db4e78
   REFCNT = 1
   FLAGS = (TEMP,ROK)
   RV = 0x1db5388
   SV = PVHV(0x1dbb8f0) at 0x1db5388
     REFCNT = 3
     FLAGS = (OOK,OVERLOAD,SHAREKEYS)
     STASH = 0x18Segmentation fault

[perl #124387] The change inadvertently didn't handle AUTOLOAD for the DESTROY method.

To fix this, instead of storing the DESTROY CV in the stash slot I kept it in the stash's MRO metadata, alongside the sub generation.

A second patch added AUTOLOAD support, which I managed to do incorrectly, which caused CPAN failured, reported as [perl #127494].

The problem was my code called gv_fetchmethod_pvn_flags() with the GV_AUTOLOAD flag, and whether that was actually and AUTOLOAD or DESTROY method, cached the CV.

When looking for an AUTOLOAD method the various APIs set $AUTOLOAD to the expected method name[1] so the AUTOLOAD method can load, call or emulate the correct method.

By caching the AUTOLOAD method, that initialization wasn't done, so AUTOLOAD would see the name of the previous method that was autoloaded instead of DESTROY.

To fix this, I only cache the DESTROY CV if it's from DESTROY, not if it's AUTOLOAD.

It might have been a little faster to store whether the CV stored was DESTROY or AUTOLOAD, and perform the required setup when calling AUTOLOAD, but it didn't seem worthwhile - I expect the processing done by AUTOLOAD itself to significantly drown that out.

[1] it works a bit differently if AUTOLOAD is XS


| Hours |         Activity |
|  4.48  |        #118295 (sec) review upstream changes, do a simple fix, |
|        |        try a selective backport, just pull in upstream module |
|        |        #118295 testing, comment with patches |
|  1.17  |        #124387 check over smoke failures, debug failure and work |
|        |        out a fix, start a test run |
|        |        #124387 update the smoke-me branch |
| 10.92  |        #125303 testing |
|        |        #125303 debug win32 perl installation issues |
|        |        #125303 prep, reply release manager perldelta query |
|        |        #125303 testing, start 5.20.2 build |
|        |        #125303 deal with failure and fix git, write bisect script |
|        |        and test it, start bisect |
|        |        #125303 testing, try a new bisect |
|        |        #125303 more bisect work |
|        |        #125303 review bisect results, manual confirmation, comment |
|        |        #125303 testing, produce a patch, open PR and comment |
|  4.03  |        #125351 testing |
|        |        #125351 debugging |
|        |        #125351 debugging |
|  0.37  |        #125471 make a patch and comment |
|  0.47  |        #125540 re-test and apply to blead |
|  2.28  |        #125569 testing, comment |
|  0.57   |       #125833 update patch |
|  1.08   |       #126045 re-test, apply to blead |
|  1.32   |       #126410/#124387 testing, minor optimization, apply to blead |
|  0.62   |       #126472 review and summarize cpan fixes |
|  0.70   |       #126544 fix documentation, create pull request for |
|         |       IPC::Cmd |
|  0.07   |       #126815 test against 125540 fix and merge into 125540 |
|  0.45   |       #126871 try to reproduce |
|  0.47   |       #127000 review |
|  0.70   |       #127231 testing, comment |
|  0.30   |       #127234 review discussion, research |
|  0.13   |       #127260 comment |
|  0.45   |       #127323 research, comment |
|  0.70   |       #127334 produce a fix and apply to blead |
|  1.03   |       #127351 testing, minor fix, push to blead |
|  0.22   |       #127380 (sec) research and comment |
|  0.49   |       #127384 review, research and comment |
|         |       #127384 review discussion |
|  2.33   |       #127386 produce a patch and comment |
|         |       #127386 comment with new patch |
|         |       #127386 comment some more |
|         |       #127386 consider several variations, apply to blead |
|  4.55   |       #127392 testing |
|         |       #127392 debugging to find cause of extra memory use |
|         |       #127392 testing and comment |
|  0.30   |       #127426 jenkins errors, push fix |
|  9.89   |       #127455 regex_sets testing, irc discussion, create #127455 |
|         |       #127455 start building modern gcc on solaris |
|         |       #127455 more gcc (dep) building, find AS isn't supported :( |
|         |       #127455 start bisect for mktables crash |
|         |       #127455 debugging |
|         |       #127455 debugging |
|         |       #127455 debugging |
|  1.02   |       #127474 review, testing, apply to blead |
|  4.08   |       #127494 debug, find cause and work on a patch, testing |
|         |       #127494 re-check, testing, apply to blead and comment |
|  0.07   |       #127505 close |
|  1.60   |       #127508 testing (after dealing with MS's breakage of my |
|         |       git installation), apply to blead |
|  0.71   |      #127511 look over new solaris thread-dirh.t test failures, |
|         |     open ticket, email discussion with jhi |
|         |       #127511 review current smoker status, comment and close |
|  0.53   |       #127514 review, testing, apply to blead |
|  0.05   |       #127522 review discussion and close |
|  0.95   |       #127532 review, testing, apply to blead |
|  0.55   |       #127533 research and comment |
|  7.78   |       #127543 dtrace build failures |
|         |       #127543 continued, test possible causing commit |
|         |       #127543 testing, research |
|         |       #127543 testing, produce a patch and comment |
|         |       #127543 testing, long comment |
|         |       #127543 testing and reply |
|         |       #127543 test simple fix and create smoke-me |
|  0.08   |       #127547 comment and close |
|  2.83   |       #127556 work up a fix and comment |
|         |       #127556 testing, apply to blead |
|  0.85   |       #127584 review, testing and apply to blead |
|  0.85   |       #127588 testing, apply to blead |
|  0.60   |       #127599 review and note to khw |
|  0.78   |       #127619 testing and comment |
|  0.22   |       #127620 review, test and apply to blead |
|  0.07   |       #127627 review, point it out to jhi |
|  0.10   |       #127628 point OP at CPAN package |
|  0.28   |       armv7 smoke failure, try Configure, start full build |
|  0.23   |       BBC reviews (122136 amongst them) |
|  0.88   |       cygwin Win32API-File/t/file.t test failure |
|  0.73   |       List::MoreUtil regression test (since it's breaking too much CPAN) |
|  0.85   |       more cygwin locale testing, debugging, email |
|  0.50   |       push a coverity build |
|  0.25   |       reply khw re cygwin locale issue |
|  0.78   |       research and mail cygwin list about broken locales |
|  0.33   |       solaris threaded/dtrace/debugging build failure, try to |
|         |       find breaking commit by inspecting, start bisect |
|  1.08   |       test/debug khw's cygwin locale skip patch |
|  0.25   |       testing rjbs' Time::HiRes move, method return value |
|         |       refcounting discussion on #p5p |

**79.97 Hours Total**

