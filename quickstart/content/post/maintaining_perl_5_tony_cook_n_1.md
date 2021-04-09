
---
title: "Maintaining Perl 5 (Tony Cook): November 2018 Grant Report"
author: Matthias Bloch
type: post
date: 2018-12-16 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_n_1"
categories:

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.
<pre>
Approximately 34 tickets were reviewed, and 6 patches were
applied.

[Hours]         [Activity]
  7.90          #123638 review existing patch, work on fixes, testing
                #123638 review test results, debugging, research
                #123638 more research, work on a fix, debugging, testing
                #123638 debugging, partial fix
                #123638 debugging
  3.15          #123724 (sec) review, work out extra work needed, work on
                fix
                #123724 (sec) fix a related problem in blead, comment with
                patch
                #123724 (sec) research
  0.40          #125760 testing, comment
  0.63          #126706 retest, apply to blead
  1.68          #131065 (sec) review, research
                #131065 (sec) more research, look for related ticket
 12.74          #132147 (sec) apply security warning patch, work on trying
                to reproduce
                #132147 (sec) debugging
                #132147 (sec) debugging, work on a fix, polish, write
                tests, testing
                #132147 (sec) fuzz testing, fix an issue, more testing
                #132147 (sec) add more validation, improve and extend fuzz
                testing, more fuzz testing runs, testing
                #132147 (sec) polish, comment with patches, fuzz tools
                #132147 rebase, testing, skip on mismatched architecture
                for test data
                #132147 more testing, apply to blead
                #132147 ticket clean up
  1.00          #133440 research, re-work
  0.45          #133511 review, test, apply to blead
  0.17          #133640 briefly comment
  0.72          #133643 review, make public and link to meta ticket
  8.96          #133659 comment
                #133659 research, comment
                #133659 work on finishing in-place edit on non-error exit
                #133659 updates to tests, work on fix within die/eval vs
                normal scope exit
                #133659 try to break scope cleanup detection, rethink
                approach
                #133659 review scope handling code, try to work up a
                solution
                #133659 more work on possible solution, comments with work
                to date
                #133659 retest and apply to blead
  0.58          #133662 review patch, testing, apply to blead
  0.58          #133668 review patch, testing, apply to blead
  0.17          #133670 comment
  2.00          #133673 research, comment
                #133673 research
  2.70          #133686 work on tests for the numeric.c conversions,
                review other patches, research comments
  0.42          #133688 research, briefly comment
  0.83          Configure alignment probe testing
  2.10          diagnose cygwin build issues, format issue in utf8.c
  1.25          lgtm, lint research
  3.30          more security issue backport
                more security issue backport
  1.50          more security issue backport, debugging
  2.00          more security issue backport, finally find problem, work
                on another backport
  2.10          more security issue backport, give up on 133423, since I
                donâ€™t think itâ€™s applicable
  1.92          more security issue backport, test failures 
  1.41          more utf8-readline re-work
                more utf8-readline re-work
  0.92          perldelta updates
  1.17          security backports â€“ discussion with khw
  2.07          security backports, testing, email to downstream
  1.47          security issue backport checks for downstream
  1.25          security ticke