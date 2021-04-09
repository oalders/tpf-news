
---
title: "Maintaining Perl 5 (Tony Cook): November 2020 Grant Report"
author: Matthias Bloch
type: post
date: 2020-12-12 00:00:00 +0000 UTC
url: "/post/tony-report-20-11"
categories:
 - Perl 5 Development

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund).
We thank the TPF sponsors to make this grant possible.
<pre>
[Hours]         [Activity]
  2.08          more win32 test noise, fix some, PR 377 for EU::MM,
                File::Path bug report cpan#133671
  1.32          #16825 research
  4.43          #17296 find another issue, work on a test and fix
                #17296 more fixes, make PR #18327
  1.18          go through PRs, checking #17724 (ask sawyerx for a
                comment), #17848 (comment), #17909 (some research), #17999
                review and some testing
  1.70          #17926 research
                #17926 rebase, re-test and apply to blead, perldelta
                updates
  1.54          review notifications mostly #17999
                #17999 comment
                #17999 comment
 10.58          #18005 more commit clean up, start a test run
                #18005 review results, add Test::Harness fix, testing,
                start a developer mode (symlinks enabled) test run
                #18005 diagnose build failure (possibly from rebase), fix
                and testing
                #18005 rebase on recent fixes and testing
                #18005 more testing, push for smoke-me
                #18005 more re-work commits, minor fixes, testing, push
                for smoke-me
                #18005 review smoke results
                #18005 more review smoke results, make PR 18306
                #18005 re-work based on xenu’s comments
                #18005 rebase win32-symlink code, make tests pass pre-
                vista, including fixing a test bug
  0.55          #18094 review and apply to blead
  0.32          #18133 check and apply to blead
  5.61          #18232 review some more, work up a erroneous case and
                comment
                #18232 review, work up a failing test case
                #18232 debugging, testing, comment
                #18232 testing, partly squash and apply to blead
  3.56          review notifications,mostly #18256
                #18256 code review, testing
  0.80          #18262 review and briefly comment
  0.47          #18272 rebase, squash and apply to blead
  0.65          #18279 research and briefly comment
  0.52          #18283 review and apply to blead
  0.67          #18285 review, research and comment
                #18285 comment
  1.63          #18293 debug, work on a fix, add a test, make PR 18307
  0.15          #18297 review and apply to blead
  0.10          #18298 review and comment
  1.35          #18308 review and comment
                #18308 fix an unrelated failure
  1.03          #18322 research and comment
                #18322 comment
  1.22          #18325 review, discussion with khw
  1.08          #18337 revise IO documentation PR based on comments
  0.45          #18339 review, test and apply to blead
  0.12          #18341 check fix was merged and close
  0.13          #18354 review, apply to blead
  0.17          #18359 review and briefly comment
  3.64          #18364 reproduce with gcc, try to figure out why MSVC
                isn’t failing the same way
                #18364 figure it out, discussion with khw
                #18364 testing
  1.93          look over OS X failures, reproduce and diagnose, work on
                an EU::MM PR to fix it #379
                my EU::MM #379: haarg made a better PR, close mine
  0.95          :utf8  rebase
  1.55          :utf8 more rebase
  2.20          check other sysv ipc ops for UTF-8 issues, find one in
                semop (test and fix), another in msgsnd/rcv (work on
                tests)
  2.23          document IO
  1.45          dragonfly changes: review smokes, reproduce freebsd issue
                (fixed in blead), r