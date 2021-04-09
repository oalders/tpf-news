
---
title: "Maintaining Perl 5 (Tony Cook): July 2019 Grant Report"
author: Matthias Bloch
type: post
date: 2019-08-11 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_j_3"
categories:

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.
<pre>

Approximately 30 tickets were reviewed, and 8 patches were
applied

[Hours]         [Activity]
  3.43          #120841 document ERRSV, CLEAR_ERRSV(), and the errno
                handling functions for internal use, comment with patch
                #120841 retest. Try to figure out some strange output in
                perlintern.pod, find the problem, improve apidoc.pl, apply
                to blead
  0.60          #125096 review old ticket, find it’s probably fixed and
                comment
  0.98          #126474 research and comment
  0.47          #130917 re-rebase the patch
  3.30          #133803 try to setup to test, but my ppc vm host doesn’t
                do 32-bit, research
                #133803 research
                #133803 long comment
  2.75          #133981 testing, research, comment
                #133981 more testing, work on a patch and comment with
                patch
  0.13          #134127 review and briefly comment
  2.35          #134169 try to make this more testable
                #134169 more work, testing, comment with patch
  4.27          #134172 benchmarking, work on fixing the issue
                #134172 work on fix, testing
                #134172 polish, check on FreeBSD
                #134172 review and comment
  0.95          #134177 re-test and apply to blead
  0.42          #134180 review and comment
  0.28          #134187 re-work and comment with new patch
                #134187 apply to blead
  0.08          #134210 review downstream release, close
  0.85          #134218 test alternative fix, apply to blead
 10.13          #134221 more work on this, get something working (needs
                polish)
                #134221 write a test, research portability to VMS, polish
                (still need win32 and to ask Craig about VMS)
                #134221 win32 implementation, testing, comment with patch
                #134221 reproduce issue on FreeBSD, work on a fix,
                testing, retest on Linux, comment with new patch
                #134221 research, email to Craig Berry re VMS
                #134221 test MSVC, fix reported issue
  1.58          #134238 comment
                #134238 review docs and comment
                #134238 comment
  3.67          #134241 try a quick test and request more information,
                work on preventing a type mismatch between compilation
                units that LTO complains about
                #134241 diagnose and get a successful build, but several
                tests fail
                #134241 fewer tests fail in a -O0 build, track down
                failure causes, comment
  0.58          #134259 review Configure and comment
  0.38          #134263 review patch, test and apply to blead
  0.28          #134270 research and comment
  0.08          #134272 review and briefly comment
  2.00          #134291 follow up on comment in irc from toddr
                #134291 re-test and apply to blead
 16.24          #134295 rebuild and reproduce
                #134295 side-track into solving other win32 build/test
                issues
                #134295 more side-track, ext/XS-APItest/t/locale.t fails
                #134295 more debugging locale.t, open ticket 134302
                #134295 testing, open another win32 build failure ticket,
                back to the original issue
                #134295 testing, try to set up to test #132863
                #134295 reproduce, try the suggested fix, try to test
                #134311, but it fails elsewhere
                #134295 debugging lockup in IO::Socket::SSL, rebuild with
                MSVC
                #134