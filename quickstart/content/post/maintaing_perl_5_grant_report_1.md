
---
title: "Maintaining Perl 5: Grant Report for December 2014"
author: Karen Pauley
type: post
date: 2015-02-03 00:00:00 +0000 UTC
url: "/post/maintaing_perl_5_grant_report_1"
categories:
 - Grants
 - Sign in

---

_Tony Cook writes:_

Approximately 36 tickets were reviewed, and 5 patches were applied.

I spent some time this month working on getting minitest working on Win32.

In 2004 the Win32 makefiles were changed to use harness instead of TEST for minitest, presumably at some time after that Test::Harness was changed to load IO unconditionally, breaking minitest on Win32.

So I switched minitest back to using TEST and then worked through the failing tests to resolve their failures.  See [perl #123394] for details.


| Hours |         Activity |
|  1.75  |        #108276 debugging, try some extra assertions |
|  2.82  |        #122287 setup to build on omnios, reproduce failure, try |
|           |     patch, test patch on other dtrace platforms, fails on OS X |
|  0.50  |        #122635 apply to blead fixing merge, re-test, bump all modules |
| 0.32   |       #123063 re-test patch and apply to blead |
|  2.28  |        #123065 review, testing, start documentation, start 32-bit iv build |
|           |     #123065 copy/adapt to other big*, testing |
|           |     #123065 polish, comment with patch |
|  0.48  |        #123105 review |
|  5.03  |        #123245 review, try to understand code, start bisect |
|           |     #123245 debugging |
|           |     #123245 produce fix and comment |
|           |     #123245 have trouble reproducing, put test together |
|           |     #123245 review test results, push to blead, close |
|  0.63  |        #123266 review, test on non glibc and glibc, push to blead |
|  0.78  |        #123312 review -V output, research local::lib and comment |
|           |     #123312 review latest response and close |
|  0.40  |        #123333 reproduce, review, test and apply to blead |
|  0.25  |        #123345 review tie documentation and implementation |
|  0.25  |        #123353 review, test, apply to blead |
|  1.22  |        #123359 review patch, unrelated #p5p file locking discussion, comment |
|  2.82  |        #123364 reproduce, debugging |
|           |     #123364 code review, debugging, comment |
|  1.99  |        #123394 comment |
|           |     #123394 testing, adjusting patches based on comments |
|  0.97  |        #123398 merge duplicate ticket, test, think about a fix, comment |
|           |     #123398 comment |
|  0.70  |        #123402 review, testing, comment |
|  0.33  |        #123413 review discussion and comment |
|  1.08  |        #123415 review, research and comment |
|           |     #123415 review and comment |
|  1.15  |        #123418 review, testing, push to blead, perldelta |
|  7.62  |        #123443 review, produce a patch, see a possible issue in seek code too |
|           |     #123443 re-review size patch, comment with patches |
|           |     #123443 re-test and apply to blead |
|           |     #123443 try to reproduce jenkins win32 failure, reproduce, |
|           |     learn too much about PerlIO and fix |
|  2.10  |        #81074 review latest discussion, code, consider new solution |
|  0.50  |        #9319 review discussion |
|  0.15  |        fix g++ builds |
|  0.17  |        investigate jenkins failure, sprout fixes before I can push |
|  2.67  |        more minitest, consistent build options, test failure diagnosis |
|  0.53  |        research and comment on khw's uudecode thread |
|  1.05  |        research, reply bulk88's malloc email |
|  0.65  |        try to fix minitest on Win32 |
|  0.33  |        try to reproduce rjbs's OS X test failures |
|  2.33  |        try to reproduce rjbs's OS X test failures some more, |
|           |     track down deparse noise |
|  0.55  |        update Data::Dumper documentation as useqq no longer |
|           |     forces pure-perl |
|  2.57  |        win32 minitest - clean up commits, op/coreamp.t |
|  3.70  |        win32 minitest - finish coreamp.t, fork.t, filetest.t, |
|        