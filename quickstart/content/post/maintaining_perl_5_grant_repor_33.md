
---
title: "Maintaining Perl 5: Grant Report for August 2016"
author: Karen Pauley
type: post
date: 2016-10-11 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_grant_repor_33"
categories:
 - Grants

---

_Tony Cook writes:_

Approximately 42 tickets were reviewed, and 6 patches were applied.

| Hours |        Activity |
|  0.93  |        %zu on HP-UX issue |
|  3.22  |        #123981 updates, testing, comment |
|  5.19  |        #126482 debugging |
|        |        #126482 bisect, review changes |
|        |        #126482 more debugging, comment |
| 12.51  |        #127663 more tests, code review |
|        |        #127663 re-work, testing |


|        |        #127663 delete on abort, testing |
|        |        #127663 remove some later work, testing |
|        |        #127663 polish, testing |
|        |        #127663 review, push to smoke-me |
|        |        #127663 code review, testing, setup for 128954 testing |
|  0.60  |        #127810 review discussion, brief comment |
|  1.35  |        #127834 (sec) review follow-ups to upstream patches |
|  1.95  |        #128095 upstream ticket #116819 with patch |
|        |        #128095 produce alternative patch |
|  2.87  |        #128263 debugging, testing, work on fix |
|        |       #128263 more work on fix, write tests, comment with patch |
|  0.83  |        #128574 testing, apply to blead |
|  2.49  |        #128598 review discussion, produce patch, testing and comment |
|        |        #128598 review new discussion, comment |
|  1.30  |        #128685 try aristotle's code cross platform |
|        |        #128685 more testing, apply to blead |
|  0.58  |        #128740 testing, research and comment |
|  0.05  |        #128767 reject |
|  0.73  |        #128775 testing, review changes, review cpan changes, comment |
|  0.12  |        #128829 (sec) , #128845 review discussion |
|  1.30  |        #128843 debugging, review S_hextract, research |
|        |        #128843 more research, comment |
|  0.30  |        #128856 review and comment |
|  0.55  |        #128865 research and comment |
|  0.08  |        #128877 (sec) reject (CPAN upstream) |
|  2.58  |        #128930 review patch |
|        |        #128930 testing, find it was applied and close |
|        |        #128930 testing, apply to blead |
|  1.55  |        #128951 debugging |
|  2.07  |        #128953 review, attempt to apply, comment |
|        |       #128953 manually apply rejected chunks, apply to blead |
|  1.60  |        #128954 review, prep for testing (or try) |
|        |        #128954 porting tests, apply to blead |
|  0.23  |        #128972 review, research, comment |
|  2.24  |        #128988 (sec) diagnose, work on a fix, testing |
|        |        #128988 (sec) comment with patch |
|  0.22  |        #128996 (sec) simplify test code, comment |
|  1.35  |        #128997 (sec) simplify, comments |
|  3.50  |        #128998 (sec) debugging, work up a fix and comment |
|        |        #128998 (sec) try to work up a test some more, comment |
|  0.30  |        #129000 (sec) minify test case and comment |
|  2.60  |        #129012 (sec) check older versions, debugging |
|        |        #129012 (sec) debugging, comment |
|  1.52  |        #129021 (sec) diagnose, make a fix, try to work up a test |
|        |        #129021 (sec) work up a test, comment with patch |
|  1.10  |        #129024 (sec) reproduce with valgrind, reduce test case |
|        |        and comment |
|  0.63  |        #129029 build library, reproduce, reduce -D flags, test |
|       |         with valgrind, comment |
|  0.43 |         #129068 debugging |
|  0.90 |         #129069 comment |
|       |         #129069 review new patch, testing, apply to blead |
|  3.90 |         #129085 (sec) reproduce, debugging |
|       |         #129085 (sec) work on fix, testing, comment with patch |
|       |         #129085 (sec) comment |
|  0.88 |         #129086 (sec) debugging, work up a fix |
|  0.18 |         #129117 (sec) comment and reject |
|  0.83 |         #129125 (sec) minify, diagnose and comment |
|  0.97 |