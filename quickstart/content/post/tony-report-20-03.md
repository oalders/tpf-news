
---
title: "Maintaining Perl 5 (Tony Cook): March 2020 Grant Report"
author: Matthias Bloch
type: post
date: 2020-05-20 00:00:00 +0000 UTC
url: "/post/tony-report-20-03"
categories:
 - Perl 5 Development

---

Approximately 51 tickets were reviewed, and 4 patches were
applied
```
[Hours]         [Activity]
  0.58          sec #125 comment with intent to allocate CVE-ID
                sec #125 request CVE-ID
  1.77          sec #130 review discussion, comment
                sec #130 follow-up
                sec #130 request CVE ID
  0.23          #13510 look up ref and comment
  0.57          #15463/#14876 research and close
  0.25          #16300 check code and comment
  0.12          #16825 review and comment
  0.40          #17134 review discussion and comment
  0.35          #17240 look up some references and comment
  0.48          #17446 review and comment
  0.92          #17458 make PR with fix for documentation nit
  0.12          #17460 note that this is now documented
  0.12          #17484 check test results and merge
  0.12          #17556 review and apply to blead
  0.73          #17571 research
                #17571 respond
  2.18          #17583 re-check and testing, update PR
  3.65          #17597 testing, research, work through win32 configuration
                #17597 testing, research, work on a fix
                #17597 testing, push to github branch for CI
                #17597 look over smoke reports pointed out by khw, skip
                the test on hpux/solaris
  3.75          #17601 testing, update upstream ticket with more
                information
                #17601 testing and trace back the issue, comment on this
                and upstream
                #17601 explain win32 data imports
  1.05          #17609 testing, discussion with khw in IRC, briefly
                comment
  0.25          #17611 brief update
  2.08          #17618 comment
                #17618 debugging, briefly comment
  2.35          #17619 review, testing and comment
  0.48          #17622 review, irc discussion
  0.42          #17627 review and IRC discussion
                #17627 brief irc discussion
  0.75          #17628 review and briefly comment
                #17628 review and apply to blead
  1.00          #17640 review and think
  0.83          #17644 try to reproduce, ask for more information
  0.27          #17645 review and briefly comment on testing the issue
  1.40          #17650 review
                #17650 more review, research, comments
  1.70          #17655 debugging and reading code, comment
  0.17          #17656 review and brief IRC discussion
  0.97          #17660 review, testing, write a regression test, apply to
                blead
  4.73          #17661 debugging, learn something new about eval, fix
                attempts
                #17661 find the cause, work on a fix and a test, push for
                CI
                #17661 look for similar debugger bugs
                #17661 work up a test, work on a fix
  0.50          #17662 review discussion, review code and comment
  0.80          #17666 review, re-test and apply to blead
  1.62          #17667 try to find a way to test
  7.23          #17668 review, code reading, do a test build
                #17668 testing, review code some more
                #17668 more debugging, track down cause, work on a fix,
                testing
                #17668 make test still running
                #17668 review test results and make PR
  0.32          #17672 respond to comment
  0.84          #17675 check our coverity result for similar issue and
                comment
                #17675 review response, comment and close
  1.98          :utf8 debugging
  1.48          :utf8 debugging (broken at eof for partial)
  1.13          :utf8 review eof handling
  1.07          CID 270661 more code review, work up a fix and push for CI
  0.27          cid 270661: rebase and fix commit message, push for re-CI
  0.42          look over more github notifications
  0.33          respond to query from James Keenan