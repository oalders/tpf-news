
---
title: "Maintaining Perl 5: Grant Report for March 2016"
author: Karen Pauley
type: post
date: 2016-05-26 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_grant_repor_28"
categories:
 - Grants

---

_Tony Cook writes:_

Approximately 48 tickets were reviewed, and 7 patches were applied

|Hours |        Activity |
|  1.57  |        investigate new ipsysv.t failures on darwin, create cpan |
|        |        #112827, fix in blead, #p5p unicode string behaviour |
|        |        discussion |
|  4.60  |        #122287 testing |
|        |        #122287 work on a patch to Configure/Makefile.SH |
|        |        #122287 more work, testing cross platform, comment with patch |
|        |        #122287 double check patch, expand comment on patch |
|  1.47  |        #124430 try to find why App::assh started working |
|        |        #124430 keep trying, start a bisect |
|  3.42  |        #125296 review, experiment and comment |
|        |        #125296 comment |
|  1.45  |        #125368 bisect and close |
|  3.30  |        #126162 apply patch manually, work on fixes |
|        |        #126162 more fixes, push to smoke-me and comment |



|  1.70  |        #126206 review ticket, try to work up a test case |
|        |        #126206 more try to work up a test case |
|  0.15  |        #126545 review discussion and tickets and comment |
|  0.32  |        #127080 review with aim to close, but unresolved issue, |
|        |        comment |
|  0.40  |        #127158 (sec) review discussion to see if anything needs |
|        |        to remain private (for now) |
|        |        #127158 make public |
|  7.15  |        #127231 ask List-MoreUtils maintainer for a fixed release |
|        |        #127231 partly track down to Params::Validate bug |
|        |        #127231 testing |
|        |        #127231 try to reproduce again |
|        |        #127231 open a PR against P::V and comment |
|        |        #127231 review autarch's suggested extra fix and comment on PR |
|  5.61  |        #127380 (sec) work on alternate patch |
|        |        #127380 (sec) more alternate patch work |
|        |        #127380 (sec) comment |
|        |        #127380 review discussion and comment |
|        |        #127380 comment, work on a patch |
|        |        #127380 reply to new comments, work on patch, comment with |
|        |        patch |
|        |        #127380 reply to new comments |
|  2.17  |        #127455 more debugging, work on hints to downgrade optimization |
|  4.03  |        #127494 testing, fix win32 jenkins failure |
|        |        #127494 testing, debugging, comment |
|  0.65  |        #127533 produce a simple patch and comment |
|        |        #127533 apply to blead |
| 1.50   |       #127543 review email from Alan Burlinson |
|        |        #127543 review, research and comment |
|        |        #127543 follow-up |
|  0.50  |        #127555 review, add to stack not refcounted ticket and |
|        |        briefly comment |
|  0.22  |        #127585 comment |
|  1.95  |        #127611 review and comment |
|        |        #127611 research and comment |
|        |        #127611 research and comment |
|  2.03  |        #127619 testing and comment |
|        |        #127619 review newest patches, testing, apply to blead |
|  1.21  |        #127635 review, produce an alternate patch |
|        |        #127635 re-test and apply alternate patch to blead |
|  0.43  |        #127636 review, test and apply to blead |
|  1.10  |        #127641 optimization, fix cmp_version.t failure |
|  1.32  |        #127657 build, try to figure out symbolizer |
|  3.48  |        #127663 review code, work on a patch |
|        |        #127663 more work on patch |
|  0.45  |        #127664 review patches, minor fixes, apply to blead |
|  0.53  |        #127687 research and comment |
|  6.75  |        #127708 research, discussion with khw on #p5p, comment |
|        |        #127708 testing, comments, extra patch, patch to add |
|        |        i_xlocale and d_duplocale probes |
|        |        #127708 