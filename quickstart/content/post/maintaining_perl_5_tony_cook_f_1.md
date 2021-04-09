
---
title: "Maintaining Perl 5 (Tony Cook): February 2019 Grant Report"
author: Matthias Bloch
type: post
date: 2019-03-11 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_f_1"
categories:

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.
<pre>
Approximately 37 tickets were reviewed, and 9 patches were
applied

[Hours]         [Activity]
  0.85          #108276 re-test, apply to blead
 12.83          #124203 reproduce, debugging (mostly seems to be lock() at
                start of DB::sub), try to bisect
                #124203 bisect some more, review results, notice change in
                bug and bisect for that change
                #124203 debugging, appears to be locking up in Cwd::CLONE,
                try to disable DB::sub calls while cloning, get a bit
                further, locks in threads::join() (holding the $DBGR lock)
                and main::sub1, trying to get the $DBGR lock
                #124203 re-work DB::sub to avoid holding the lock,
                testing, clean up and polish, work on a regression test
                #124203 debug test failures, fixes, testing, polish,
                comment with patch
                #124203 look for similar issues, work up a test case, fix
                and comment with patch
  6.50          #130585 debugging
                #130585 try to re-work parsing to avoid the bug
                #130585 testing, polish, comment with patch
                #130585 re-work patch a little, comment with updated patch
  1.95          #131683 review, testing, comment
                #131683 review original discussion testing, minor fixes,
                apply to blead
  0.82          #132964 review and comment
                #132964 review new patch, testing, apply to blead
  0.75          review security queue, comments, make #133334, #13335
                public
  0.30          #133462 work on fix, testing, apply to blead
  0.72          #133523 (sec) reply
  0.60          #133638 review and comment, mail perldoc.perl.org
                maintainer
  1.03          #133660 work up a test and comment with patch
                #133660 re-test and apply to blead
  0.10          #133670 review and close
  0.28          #133695 review and comment
  0.70          #133771 review, research and comment
  2.08          #133778 debugging, comment with patch
                #133778 retest and apply to blead
  0.62          #133781 review and comment
  2.14          #133795 research and comment
                #133795 review toke.c, perly.y to try to find a solution
  1.20          #133803 review, comment, prepare to try an test
  0.83          #133810 test setup, testing, propose fixes for
                backporting, close #133760
  1.39          #133816 review apparent bisect, testing
                #133816 reproduce bisect, ask khw about it
  1.78          alpine linux tests, open #133818 and #133819
                #133818 discussion with khw on providing access to alpine
                linux
  2.10          #133822 debugging
                #133822 debugging, comment with patch, also comment on
                similar #133823
  0.15          #133827 comment
  1.75          #133830 research, testing, comment, some setlocate
                discussion with khw
  0.87          #133838 review weekend commits, review patch, testing,
                apply to blead
  0.40          #133846 review, test, apply to blead, research ERRSV
                history
  4.02          #133850 debugging, work on fix
                #133850 testing, check a similar case, comment with patch,
                search for other bugs with the same cause, comment with
                some similar cases
  0.27          #133851 briefly comment, debugging
  2.32          #133853 research, adhoc testing, work up a test case,
                testing, apply to blead
  0.10          check for un-perldeltaed changes
  1.07      