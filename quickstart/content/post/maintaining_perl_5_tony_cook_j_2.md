
---
title: "Maintaining Perl 5 (Tony Cook): June 2019 Grant Report"
author: Matthias Bloch
type: post
date: 2019-07-24 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_j_2"
categories:

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.
<pre>
Approximately 40 tickets were reviewed, and 18 patches were
applied

[Hours]         [Activity]
 11.77          #122112 review, testing, comment
                #122112 research, set up VM
                #122112 try to diagnose issue on android, remove some
                extraneous debug output
                #122112 debugging, testing older versions
                #122112 more testing, debugging â€“ apparently the signal is
                blocked
                #122112 work up a fix for io/pipe.t, try to work out
                io/eintr.t
                #122112 re-test and apply to blead
  0.38          #126991 review history, comment
  0.15          #130560 review and comment
  1.38          #130585 re-test and apply to blead
  0.43          #130632 try to reproduce and fail, find the probably fix
                and comment
  1.12          #133850 re-test and apply to blead
  0.73          #133913 re-test and apply to blead
  3.62          #133936 re-work patch, comment with updated patch
                #133936 retest and apply to blead
                #133936 work on a fix, testing, comment with patch
                #133936 re-test, apply to blead
  2.12          #133989 diagnose, work on a fix, comment with patch
                #133989 re-test and apply to blead
  1.00          #134008 re-testing, minor re-work, apply to blead
  0.47          #134035 re-test and apply to blead
  0.42          #134045 re-test and apply to blead
  0.90          #134057 try to regain access to HP-UX, ask Tux
                #134057 apply to blead after Tuxâ€™s response
  0.22          #134065 look into this again, debugging
  0.43          #134072 re-test, minor comment change, apply to blead
  4.50          #134125 work on a patch and comment with it
                #134125 re-test and apply to blead
                #134125 work on a patch to make it fatal instead
                #134125 re-test and apply to blead
  0.20          #134137 review and comment
  1.68          #134138 research, review and comment
  0.32          #134164 review and comment
  1.87          #134169 review related tickets, research, testing comment
                #134169 review new patch, testing, apply to blead
                #134169 ticket clean up, perldelta
  0.48          #134172 review, check with valgrind, check for
                valgrind/asan on freebsd
  1.08          #134175 review, work on tests, testing
                #134175 re-work tests, apply to blead
  4.84          #134177 comment
                #134177 work on a fix
                #134177 work on adding G_RETHROW to eval_sv
                #134177 more work, testing, comment with patch
                #134177 minor re-work, consider adding eval bytes/unicode
                support, but doesnâ€™t seem to be worth it
                #134177 briefly comment
  4.61          #134179 debugging (not as simple as it looked)
                #134179 work out the issue, work on fix, testing, comment
                with patch
                #134179 fix minor patch issue, check possibly related
                ticket, apply to blead
  1.25          #134187 research, comment
                #134187 research, comment with brief patch
  3.51          #134189 setup test env, testing, work on fix
                #134189 finish fix, comment with patch
                #134189 re-work, comment with new patch
                #134189 re-test and apply to blead
  3.05          #134193 work up a fix and comment with patch
                #134193 research, work on a minor fix, comment with patch
                #134193 re-test and apply to blead
  1.23          #134194 debugg