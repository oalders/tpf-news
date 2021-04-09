
---
title: "Maintaining Perl 5: Grant Report for January 2016"
author: Karen Pauley
type: post
date: 2016-02-22 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_grant_repor_26"
categories:
 - Grants

---

_Tony Cook writes:_

Approximately 43 tickets were reviewed, and 11 patches were applied

| Hours |         Activity |
|  2.38 |         #123737 testing, apply to blead |
|       |         #123737 testing, produce a patch and comment |
|       |         #123737 re-test, look for similar no_op() issues, push to blead |
|  2.15 |         #123788 re-check patch, testing |
|       |         #123788 more testing, apply to blead |
|  4.37 |         #124387 work on autoload on top of above |
|       |         #124387 testing, code archaeology, produce patches and comment |
|       |         #124387 review test results, comment with updated patch |
|  2.31 |         #125540 debugging |
|       |         #125540 produce a patch, testing and comment |


|  0.85 |         #126042 retest and apply to blead |
|       |         #126042 research |
|  1.60 |         #126045 review, IRC discussion with khw, comment with a simple patch |
|       |         #126045 adjust patch, irc discussion, comment with new patch |
|  1.05 |         #126240 check if I missed something and fix it, testing |
|       |         #126240 note the extra commit |
|  3.89 |         #126410 debugging |
|       |         #126410 work on a patch, switch to related ticket |
|  0.42 |         #126431 research and comment |
|  0.07 |         #126544 comment |
|  5.69 |         #126632 testing, review patch, comment |
|       |         #126632 testing, apply to blead |
|       |         #126632 check jenkins issue, fail to reproduce, ask Seveas about platform |
|       |         #126632 more jenkins diagnosis |
|       |         #126632 irc discussion, testing |
|       |         #126632 testing, ask Seveas to fix jenkins |
|       |         #126632 remove the extra diagnostics |
|  2.02 |         #126633 review davem's suggestion, testing, comment |
|       |         #126633 more testing, apply to blead |
|  0.90 |         #126877 review, testing, conflict fixes |
|       |         #126877 rebase, testing, apply to blead |
|  0.98 |         #126922 (sec) re-check, testing |
|       |         #126922 apply to blead, add to maint votes |
|  1.37 |        #126953 review patch, research handling of unknown keys by |
|       |         other implementations, review original discussion, comment |
|  1.01 |         #126991 research, produce a patch and comment |
|       |         #126991 testing, apply to blead |
|  1.13 |         #127054 debugging |
|       |         #127054 more testing and comment |
|  2.85 |         #127061 debugging, comment |
|       |         #127061 produce patch, testing and comment |
|  0.93 |         #127063 produce a patch and comment |
|  1.17 |         #127122 produce a patch and comment |
|       |         #127122 testing, apply to blead, create ticket 127333 for similar until() issue |
|  0.92 |         #127131 review, testing, push to blead, review kid51's author fixes and comment in #p5p |
|  2.00 |         #127149 review, testing, work on a patch |
|       |         #127149 comment |
| 13.12 |         #127158 (sec) review, testing |
|       |         #127158 (sec) testing, research |
|       |         #127158 (sec) ticket clean up, review discussion, work on a patch, testing, comment with patch, set all ccs of |original message as ticket ccs |
|       |         #127158 (sec) review discussion |
|       |         #127158 (sec) comment |
|       |         #127158 (sec) research, work on alternate patch |
|       |         #127158 (sec) more work on alternate patch, private irc discussion with rjbs |
|       |         #127158 (sec) produce new patch and comment |
|       |         #127158 (sec) review discussion |
|  1.10 |         #127231 testing |
|       |         #127231 more testing |
|  0.57 |         #127260 review and comment |
|  1.65 |         #127262 make a different reproducer, comment |
|       |         #127262 diagnose, try t