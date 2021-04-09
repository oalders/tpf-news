
---
title: "Maintaining Perl 5: Grant Report for January 2017"
author: Makoto Nozaki
type: post
date: 2017-02-22 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_grant_repor_37"
categories:
 - Grants

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.

<pre>
Approximately 38 tickets were reviewed, and 8 patches were
applied

[Hours]         [Activity]
 13.82          #122490 (sec) more merge conflicts
                #122490 (sec) more merge conflicts, track down warning
                sources
                #122490 (sec) track down warning sources, start merging
                test changes
                #122490 (sec) more test merging, testing, debugging
                #122490 (sec) debugging
                #122490 (sec) debugging
  0.97          #126228 build testing, apply to blead
 16.08          #127663 testing, apply hash seed env suppression patch,
                back to in-place changes
                #127663 work on chdir test, testing, debugging, make
                mg_freeext() API and fix docs
                #127663 cleanup, threads handling, threads test
                #127663 more threads testing, try to make it fail with
                fork
                #127663 more try to make it fail with fork and succeed,
                work on fix, code polish
                #127663 hoist some work back up, testing
                #127663 uncombine thread/fork child handling which I
                combined by accident, work on more tests and find a couple
                of cleanup issues
                #127663 more tests
                #127663 post patch to ticket
  0.22          #128528 (sec) review and comment
  0.88          #128998 track down when it was fixed, ticket management
  0.30          #129012 make public, comment and close
  1.88          #129024 review, make public, check fix backports to 5.24,
                non-trivial backport to 5.22, comment
  1.30          #129125 check, testing, apply to blead
  1.65          #129149 apply patch, test #130557 case, testing, make
                public apply to blead, comment on #130557
  0.08          #129187 check and merge into #129149
  0.95          #129190 rebase with some conflicts, testing, make public,
                apply to blead
  0.17          #129199 make public, comment and close
  2.62          #129274 (sec) try to find an alternative attack
                #129274 more trying to break it, write regression test,
                testing, make public, apply to blead
  2.12          #129292 review code, debugging, make public and comment
  1.77          #129298 review patches, research, consider whether changes
                improve perl
                #129298 more consideration, ask khw
  4.32          #129340 (sec) review code, think about solutions
                #129340 (sec) work on a solution, testing
                #129340 (sec) write a regression test, testing
                #129340 (sec) suggested changes, testing
                #129340 (sec) research, comment with updated patch
  0.50          #129342 (sec) test provided patch, create a test case and
                comment
  0.45          #129377 (sec) review patch, look for similar issues,
                comment
  1.32          #129840 (sec) review, testing
                #129840 get it to test, merge into 129377
  0.40          #129848 review and make public
  1.53          #129861 (sec) debugging
  0.42          #129887 (sec) review and comment
  0.82          #129963 research, make public and link to stack-not-
                refcounted meta ticket
  0.92          #129975 debugging, make public and link to stack-not-
                refcounted meta ticket
  0.28          #130100 make public and point at list discussion on
                removal
  0.73          #130256 debugging, make public and link to stack-not-
                refcounted meta ticket
  1.67          #130262 apply patch with nois