
---
title: "Maintaining Perl 5: Grant Report for April 2016"
author: Karen Pauley
type: post
date: 2016-06-18 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_grant_repor_29"
categories:

---

Tony has supplied the second report for his "Maintaining Perl 5":http://news.perlfoundation.org/2013/05/grant-application-maintaining.html grant, which was recently successfully "extended":http://news.perlfoundation.org/2016/02/maintaining-perl-5---grant-ext-1.html.

_Tony Cook writes:_

Approximately 40 tickets were reviewed, and 3 patches were applied.

| Hours |         Activity |
|  1.66 |         #121734 report cygwin issue upstream |
 |      |         #121734 test upstream fix and comment |
|       |         #121734 comment on POSIX::strxfrm() bug note |
|  0.27 |         #122551 review discussion, TR::Perl bugs |
|  2.23 |         #126188 debugging, try potential fixes, comment |
|  0.92 |         #126203 reproduce, comment and supply a trivial patch |
|  2.64 |         #127380 (sec) work on patches, comments |
|       |         #127380 find and post missing patch |
| 11.21 |         #127663 research |


|       |         #127663 research |
|       |         #127663 research, work on fix |
|       |         #127663 more work on fix |
|       |         #127663 more work on fix, mail craigb asking about VMS behaviour |
|       |         #127663 more work, research, work on path handling |
|  0.93 |         #127708 irc discussion, testing |
|  9.30 |         #127743 review, testing |
|       |         #127743 debugging, simple fix, start on complex fix |
|       |         #127743 code, tests, debugging |
|       |         #127743 testing, fixes, comment with patches |
|       |         #127743 look over change in cperl |
|  3.06 |         #127759 debugging |
|       |         #127759 debugging |
|  1.03 |         #127760 testing, update Maintainers.pl, customized,dat and |
|       |         apply to blead |
|  0.67 |         #127780 testing, produce and patch and comment |
|       |         #127780 fix patch and comment |
|  6.35 |         #127799 testing, debugging, brief comment, further debugging |
|       |         #127799 try a 32-bit build, code review, comment |
|       |         #127799 test Dave's fix |
|  1.13 |         #127805 research, fix a doc typo |
|  1.79 |         #127810 comment |
|       |         #127810 review patch, comment on security ticket, more review |
|       |         #127810 review latest discussion |
|  0.87 |         #127829 testing, comment |
|  0.87 |         #127834 (sec) ticket merging, review discussion, comment |
|  0.82 |         #127875 research, comment |
|  0.57 |         #127880 review, testing |
|  0.32 |         #127882 adjust commit message, re-test and apply to blead |
|  0.70 |         #127885 review patches and comment |
|  3.10 |         RC-1 testing, open #127894 and work on a fix, open #127895 |
|       |         #127894 review alternate approach, testing, apply to blead |
|  2.77 |         #127895 debugging, try freebsd 9.3, comment |
|  0.47 |         #127922 research and comment |
|  2.18 |         #127923 review, testing, porting fixes, push to smoke-me, |
|       |         comment |
|  0.17 |         #127936 testing, fix typo |
|  0.78 |         #127953 testing, comment |
|  0.52 |         #127970 (sec) comment |
|  0.58 |         #127981 test to see if recursion overflowing thread stack |
|  0.12 |         #127991 review and comment |
|  1.38 |         investigate darwin g++ failures, work on patch, testing |
|  2.57 |         investigate new cygwin failure, push a fix |
|  0.08 |         ipc-cmd start another run |
|  1.35 |         ipc-cmd win32 testing, test test script, start first run |
|  0.20 |         ipc-cmd: review first result, start second |
|  2.05 |         more win32 IPC-Cmd testing |
|  1.43 |         post-5.24 patches |
|  0.05 |         review test results and push to blead |
|  0.45 |         rjbs' github notices: test, provide a patch |
|  1.12 |         win32 jenkins errors, looks like stuck process, notify |
|       |         Seveas and list |
| 