
---
title: "Maintaining Perl 5 (Tony Cook): November 2019 Grant Report"
author: Matthias Bloch
type: post
date: 2019-12-18 00:00:00 +0000 UTC
url: "/post/tony-report-19-11"
categories:
 - Grants
 - Perl 5 Development

---

Approximately 58 tickets were reviewed, and 3 patches were
applied
```
[Hours]         [Activity]
  0.27          sec #127 research and comment
  0.12          #14409 review and close
  0.77          #17071 re-check old suggested commit and make PR #17313
  1.95          #17145 testing and close
  0.81          #17154 research and comment
                #17154 research and comment
  0.28          #17187 review and comment
  0.10          #17222 review and comment
  0.62          #17228 debugging, comment
  0.91          #17229 look at trying to avoid allocating a bit to the
                unicode_strings feature
                #17229 comment
  1.27          #17230 debugging, comment
  0.93          #17234 review, work on a fix and make PR 17248
  0.82          #17238 review, testing, apply to blead
  1.18          #17241 debugging, comment with recursive backtrace, more
                testing
                #17241 recheck, mea culpa
  0.65          #17242 review, local test edits, comment
  1.74          #17245 setup to test
                #17245 review build failure, research, comment, local fix,
                restart test (including deps, yay)
  0.13          #17248 recheck and apply to blead, comment on #17234 to
                request OP tests it
  0.30          #17249 review and apply to blead
  0.53          #17254 work up a fix and make PR 17257
  0.15          #17255 comment
  0.80          #17268 reproduce, basic debugging and comment
  0.53          #17269 testing and comment
  0.73          #17270 review code, debugging
  0.92          #17279 review, work on a fix and make PR #17289
  0.10          #17281 re-review and approve
  0.10          #17287 testing, comment
  1.27          #17288 work on a fix, testing
 10.00          #17293 review, research and comment, work on bringing
                Tie::H::NC into core
                #17293 more moving code, testing
                #17293 fix some issues, testing
                #17293 more testing, polish, create PR #17305
                #17293 update re/pat.t to actually run under minitest
                #17293 rebase to test CI updates
  0.15          #17300 review and merge
  0.82          #17306 re-test and make PR #17311
  0.38          #17322 review test results, work on a fix, update PR
  0.68          #17331 comment
  0.35          #19 review, research and comment (expect to reject this)
  0.72          :utf8 - failquiet handling and testing
  1.10          :utf8 - finish up failquiet, time for something more
                complex, replacement
  1.50          :utf8 - more replacement, find a bug and fix it
  1.75          :utf8 finish discard re-work, work on replacement
                processing
  2.07          :utf8 more replacement
                :utf8 more replacement
  1.68          :utf8 re-work discard handling, fail to rebase, start re-
                working from original commit
  2.72          :utf8 re-work, testing
  2.35          continue review security issue list
  2.02          cygwin github action: fix handling when called for a PR,
                more testing
  1.88          cygwin github action: mostly get it working, some cleanup,
                test the cleanup
  2.25          cygwin tests to diagnose github action failures, try to re-
                work cygwin github action
  1.20          more cygwin github action testing, debugging
  3.00          more review coverity scan results, work on error handling
                for POSIX::SigSet
  1.23          perldelta updates
  0.57          reproduce and fix a deparse issue introduced with the
                faster feature handling change
  0.88          review and apply some PRs, look over 17267
  1.55          review coverity scan results
  0.08          review test results, squash cleanup, make PR 17323
  0.67          review tickets
  0.40 