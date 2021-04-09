
---
title: "Maintaining Perl 5 (Tony Cook): Grant Report for July 2017"
author: Makoto Nozaki
type: post
date: 2017-09-06 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_g_1"
categories:
 - Grants
 - Perl 5 Development

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.


<pre>Approximately 38 tickets were reviewed, and 3 patches were
applied

[Hours]         [Activity]
  0.12          #p5p POSIX 2008 locale api discussion with khw
  4.00          #124368 work up a decent test, find a difference between
                threaded/non-threaded, work up a patch, testing
                #124368 more testing, re-work patch
                #124368 cleanup, testing, comment with patch
  8.81          #127663 work on non-*at() function alternative checks
                #127663 more non-*at() function
                #127663 more non-*at() function code, need to write path
                manip code
                #127663 no at-function chdir handling, testing
                #127663 debugging
                #127663 debugging
                #127663 handling change of dir
  0.17          #130651 (sec) review
  0.63          #130841 (sec) review discussion and comment
  0.97          #130917 (sec) review discussion, testing, comment
  4.68          #131065 (sec) review discussion, research, comment
                #131065 (sec) research
                #131065 (sec) work on a patch
  0.77          #131546 testing, updates, apply to blead
  0.32          #131559 review cpan ticket discussion
  3.78          #131570 work on a simpler repro case
                #131570 try some more, testing, apply to blead
  1.43          #131588 make public, repackage OP’s patch with a test
  1.15          #131597 testing, apply to blead
  0.13          #131598 (sec) review discussion
  4.03          #131618 (sec) work on a patch
                #131618 (sec) work on a patch
                #131618 (sec) testing, debugging
  1.30          #131648 (sec) debugging
  6.17          #131665 (sec) write a test, testing, research, more
                testing
                #131665 (sec) comment with patch
                #131665 (sec) add a text version, research, review
                supplied patch, comment
                #131665 (sec) review patch, research, testing
                #131665 (sec) comment with alt patch
  0.83          #131685 read discussion and comment
  2.68          #131690 produce a patch and comment
                #131690 research, reject ticket
  0.18          #131697 comment
  2.72          #131746 try to reproduce
                #131746 fail to reproduce, research and comment
  4.97          #131777 debugging, work on patch
                #131777 more work on patch, debugging error recovery
                #131777 more work on error recovery, update tests,
                perldiag, comment with patch
                #131777 fix a typo
  0.12          #131787, #131788 reject some security tickets (not perl5
                core)
  1.98          #131793 (sec) debugging
                #131793 make public, work on a fix, comment with patch
  1.47          cygwin64 test failures: debugging
  0.08          cygwin64 test failures: try a partial fix
  0.45          cygwin64: write up a full commit, testing
  0.20          locale api research after discussion with khw
  1.87          more review perlunicode, perlunifaq, comment
  1.30          more utf8 stuff, debugging cygwin64 issues too
  2.15          more utf8 stuff, editing for patch, comment with patch
  0.15          review list discussion (hash ties)
  0.70          review security queue
  1.33          utf8 docs: more fixups, apply to blead, comment
  0.30          utf8.pm function naming: read discussion, review
                perlunicode
  0.90          utf8.pm: more research, briefly comment
  0.40          utf8: some fixes
  2.18          utf8::* discussion: work on a better patch, review
    