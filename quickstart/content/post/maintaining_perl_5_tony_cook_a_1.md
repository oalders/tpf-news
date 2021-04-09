
---
title: "Maintaining Perl 5 (Tony Cook): August 2018 Grant Report"
author: Matthias Bloch
type: post
date: 2018-09-16 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_a_1"
categories:

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.
<pre>
Approximately 44 tickets were reviewed, and 4 patches were applied.

[Hours]         [Activity]
  1.23          #p5p discussion with khw on shared hash problem
  4.73          #132158 testing, debugging, work on a fix, more testing
                #132158 consider other fixes, testing, comment with
                patches
  1.60          #132655 (sec) work on test code, produce a fix and a
                patch, comment with patch
  0.72          #132683 re-check patch, test and apply to blead
  0.35          #133002 review, make public, merge into 132609
  1.71          #133204 (sec) research, comment
                #133204 (sec) comment
  0.30          #133250 (sec) review, comment
  0.37          #133292 review and briefly comment
  1.87          #133314 work on test code
                #133314 more testing, apply to blead
  8.73          #133326 debugging, work on fixes
                #133326 more work on fixes, testing
                #133326 debugging, clean up, testing
                #133326 polish, commit message, comment with patch
                #133326 re-test, apply to blead
  0.40          #133331 (sec) review, ask for more information
                #133331 (sec) briefly comment
  2.45          #133334 (sec) debugging
                #133334 (sec) debugging
                #133334 (sec) comment
  0.05          #133335 (sec) brief testing and comment
  0.95          #133345 (sec) testing, comment
  0.95          #133376 review patch, test build, testing, apply to blead
  0.07          #133406 (sec) comment and close
  0.55          #133417 review patches, original discussion, comment about
                the commit messages
                #133417 review, research and comment
  2.29          #133422 review ticket, look over code
                #133422 work on a fix
                #133422 more work on fix, testing
  0.38          #133423 (sec) private IRC discussion with khw
  0.30          #133431 review and briefly comment
  0.55          apm821xx cross build thread, research, comment
  0.47          Briefly review khw’s khw-core branch
  0.95          hacking on feature.h thread, review code, testing, comment
  2.07          reply sawyer email on unicode filenames on Win32
  1.88          review maint-votes, add some commits, cherry pick/test two
                commits
  0.37          review security issues, make 133345 public
  0.77          review security queue
  1.53          security queue review/clean up
  0.33          security ticket round up
  1.17          utf8_readline: consider optimizations
  1.38          utf8_readline: croak tests, add another test and find a
                bug
  1.45          utf8_readline: debug/fix surrogate bug, more tests
  2.98          utf8_readline: debugging, fix short handling bug, more
                testing, find surrogate handling issue, debugging, croak
                on surrogate bug
  1.87          utf8_readline: debugging, possible fix, testing
  1.72          utf8_readline: debugging, testing, debug a test failure
                and consider possible fixes
  1.87          utf8_readline: more optimizing non-mutating stream
  2.45          utf8_readline: more optimizing, building/testing + #p5p
                discussion of tainted \p{} property names with khw
  2.08          utf8_readline: more tests
  1.27          utf8_readline: polish, testing
  0.58          utf8_readline: rebase, testing, review
  1.83          utf8_readline: simplify some code based (oldish)
                discussions with khw, testing
  1.90          utf8_readline: work on fix for optimized code consuming
                input incorrectly
  1.53          utf8_r