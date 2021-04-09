
---
title: "Maintaining Perl 5 (Tony Cook): January 2019 Grant Report"
author: Matthias Bloch
type: post
date: 2019-02-15 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_j_1"
categories:

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.
<pre>


Approximately 54 tickets were reviewed, and 11 patches were
applied

[Hours]         [Activity]
  8.88          #108276 review
                #108276 check over committed changes, look to re-work, ask
                list about PERL_OP_PARENT
                #108276 cleanup PERL_OP_PARENT detritus
                #108276 review old patches, re-work, testing
                #108276 more testing, work on optimize_op()
                #108276 research, cleanup, more testing. Comment with
                patch
  0.58          #121033 research, comment
  1.98          #123724 rebase, retesting, bump versions, apply to blead,
                make public
  0.05          #126191 check for activity and close
  2.69          #127521 review code, POSIX, work on implementation
                #127521 local commit with notes
  1.72          #130367 rebase, write new tests, testing, apply to blead
                #130367 perldelta
  1.90          #130585 debugging
                #130585 debugging
  0.32          #131506 research and comment
                #131506 close
  0.07          #131931 link to stack-not-refcounted meta ticket
  0.08          #131955 check for activity and close
  0.52          #132158 rebase, retest and apply to blead
  1.35          #132338 research, comment
  0.12          #132583 same and add the form sub-parse fix to backports
                (the other is in 5.28 already)
  3.76          #132777 prep, review code, research, work on some docs
                #132777 testing, more docs, comment with patch
  0.27          #133030 re-check, comment
  0.27          #133153 comment
  0.62          #133504 research and comment
  0.55          #133522 review patch, testing, apply to blead
  1.88          #133524 work up a fix and apply to blead
  3.88          #133575 prep and start gcc 8.2.0 build
                #133575 work up a patch, testing, comment with patch
                #133575 consider comment changes, research, testing, apply
                to blead
  0.60          #133590 try to work with metaconfig, report upstream to
                the metaconfig wizards
  0.77          #133721 review, write test, testing, apply test and fix to
                blead
  0.13          #133740 review new patch and comment
  0.13          #133744 review and comment
  0.10          #133746 review and close
  0.10          #133750 review and reject
  1.00          #133751 comment with patch
                #133751 retest, apply to blead
  0.55          #133752 follow up
                #133752 got a response privately due to rt.perl.org
                issues, forwarded request to perlbug-admin and closed
                tickets
  0.25          #133753 (sec) fix subject, comment
                #133753 (sec) testing, comment
  1.82          #133754 test setup, testing, research
  2.08          #133755 research, produce a patch and comment
                #133755 fix a typo, test, apply to blead
                #133755 comment and close
  0.25          #133756 review and comment briefly
  0.61          #133760 research and comment
                #133760 research (read the Configure source) and comment
  0.72          #133765 review, start prep to test, review code, comment
                #133765 research, add commit to votes file
  0.53          #133771 review patch and comment
                #133771 review new patch and comment
  0.97          #133776 review and comment
                #133776 review and comment
  3.32          #133782 diagnose, work on a fix, testing
                #133782 retest patch against blead, apply to blead
                #133782 re-test, fix ticket number in patch, apply to
