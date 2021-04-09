
---
title: "Maintaining Perl 5 (Tony Cook): January 2020 Grant Report"
author: Matthias Bloch
type: post
date: 2020-03-30 00:00:00 +0000 UTC
url: "/post/tony-report-20-01"
categories:
 - Perl 5 Development

---

Approximately 54 tickets were reviewed, and 8 patches were
applied
```
[Hours]         [Activity]
  0.67          PR #12 review, comment
                #12 follow-up
  0.23          #12147 research and close
  3.28          #17157 reproduce, research
                #17157 test afresh’s patch, add a patch, testing cross
                platform, make PR
  1.31          #17186 add docs, update $VERSION, update PR
                #17186 comment
  0.20          #17335 briefly review and comment
  1.73          #17338 work up a test case, a fix, testing, make PR#17484
  0.33          #17373 review, apply to blead
  0.18          #17388 comment
  0.69          #17393 review and comment briefly
                #17393 review and apply to blead
  1.37          #17397 debugging
  0.68          #17401 review and comment
  0.17          #17404 review and briefly comment
  1.85          #17407 research and comment
                #17407 briefly comment
                #17407 research and comment
  1.61          #17409 review and comment
                #17409 review and comment
  1.57          #17419 try to diagnose build problem, work up a fix and
                apply to blead (PR needs to rebase)
                #17419 review the content changes and apply to blead
  2.18          #17421 review, research and comments, prepare for some
                testing
                #17421 testing and comment
  0.52          #17422 simple fixes and make PR #17483
  2.46          #17423 review and comment, irc discussion on #17421, re-
                work the test code, testing
                #17423 review test results, push and make PR #17425
  0.63          #17424 review installer code, testing, make PR #17432
  1.33          #17425 fix a problem with patch, update subject for commit
                #17425 check possible temp files left behind, could not
                reproduce
                #17425 make a suggested change
                #17425 recheck and apply to blead
  2.21          #17428 research, debugging, testing, comment
                #17428 make test case, fix, create PR #17479
  0.42          #17432 recheck and apply to blead
  0.72          #17458 research, comment
  0.89          #17477 review, testing and comment
                #17477 review and briefly comment
  1.02          #17478 darwin also a problem, update PR
                #17478 cygwin too
                #17478 clean up commits and force push for CI
                #17478 review test results and apply to blead, perldelta
  0.58          #17479 apply to blead and perldelta
  2.48          #17483 some fixes
                #17483 re-test, update ChangeLog, testing
                #17483 check results, apply to blead and perldelta
  0.50          #17500 review, test something, briefly comment
                #17500 brief comment
  0.73          #17501 review, diagnose and comment, report new bug
                against BBC reported module
  0.22          #17502 review, #p5p pack discussion
  0.33          #17503 review
  0.83          :utf8 debugging
  1.80          :utf8 debugging warn reporting
  1.05          :utf8 debugging, find an issue, research
  1.87          :utf8 fix issue above, testing
  2.02          :utf8 get back up to speed, fixes to WIP replacement code
  2.60          :utf8 main structure done, fill out utility functions
  1.25          :utf8 message handling, basic testing
  1.48          :utf8 more fixes
  1.40          :utf8 more non-fast_gets
  1.18          :utf8 reconsider code structure
  1.42          :utf8 review how fast_gets() is set, more non-fast_gets
                code
  1.02          :utf8 start non-fast_gets path
  1.77          :utf8 testing, debugging
  2.23          :utf8 update allow- tests for replace* options, debugging
  0.47          Felipe follow-up: comment
  0.17          Felipe follow-up: comment 