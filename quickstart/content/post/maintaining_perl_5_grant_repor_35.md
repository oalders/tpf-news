
---
title: "Maintaining Perl 5: Grant Report for November 2016"
author: Makoto Nozaki
type: post
date: 2016-12-14 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_grant_repor_35"
categories:
 - Grants

---

*The Perl Foundation received the following report from Tony Cook.*

Approximately 31 tickets were reviewed or worked on, and 8 patches
were applied.

<pre>
[Hours]         [Activity]
  0.37          #116945 review and comment
                #116945 try to review branch, comment
  2.05          #122112 research, review code
  3.66          #123638 (sec) work on patch, testing
                #123638 (sec) perldiag wording, work on regression test
                #123638 (sec) final tests, comment with patch
  3.58          #126706 testing, research
                #126706 try an alternate approach (and fail)
  0.65          #128967 research, testing, revert a test changed only to
                satisfy the commit reverted to fix this ticket
  1.42          #129000 testing, apply to blead with minor commit message
                changes
</pre>


<pre>
  5.10          #129238 debugging
                #129238 finish new API, work on using it
                #129238 more work on using it
  0.83          #129272 (sec) testing, dup of #129000
  4.83          #129274 (sec) debugging
                #129274 (sec) debugging
                #129274 make public, comment with minor patch
  0.88          #129292 (sec) minimize
                #129292 (sec) more minimize, comment
  0.42          #129826 testing, apply to blead and close
  1.18          #129881 (sec) minimize, comment
  1.27          #129963 (sec) minimize, debugging, comment
  1.32          #129975 (sec) minimize, testing, comment
  5.46          #129990 ask for patch as attachment, briefly review
                without testing
                #129990 review, try a sanitize build, track down other
                sanitize failures
                #129990 testing
                #129990 review code and try to break it, benchmarking,
                comment
                #129990 testing, comment
  0.97          #129991 minimize, comment
  6.28          more debugging, open ticket #129997
                #129997 testing, work on some alternate approaches,
                comment with patch
                #129997 work on a regression test, apply to blead, merge
                into (sec) #129995 and close
                #129997 comments
  3.39          #130080 testing, comment
                #130080 testing, reproduce, email Steve Hay
                #130080 comment
  0.95          #130082 review and comment
                #130082 review new patch, testing, apply to blead
  0.80          #130100 (sec) research, comment
  2.00          #130101 review and comment
                #130101 comment
                #130101 review, testing, extra patches for bumping version
                and updating changes list, apply to blead
 24.08          #130108 test setup, debugging
                #130108 work on a fix, testing, comment with patch
                #130108 testing on 10.3, comment with patch for build
                issue
                #130108 debugging, comment
                #130108 testing on Solaris, testing on freebsd 11, trying
                to build older perl with dtrace on freebsd
                #130108 check reference supplied by swills in irc, irc
                discussion
                #130108 work on a Makefile.SH patch
                #130108 more dtrace build re-work
                #130108 more dtrace build re-work, cross-platform testing
                #130108 work out darwin failure, work on getting it to
                probe correctly in a FreeBSD jail
                #130108 move -xnolibs check earlier, testing
                #130108 cross-platform test setup
                #130108 test built binaries, polish, post patch to ticket
  0.90          #130128 review, testing, apply to blead
  0.73          #130133 review discussion, testing, apply to blead
  0.73          #130143 research, comment
  1.00          #130193 revie