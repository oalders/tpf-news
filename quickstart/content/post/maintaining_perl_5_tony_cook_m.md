
---
title: "Maintaining Perl 5 (Tony Cook): March 2018 Grant Report"
author: Matthias Bloch
type: post
date: 2018-05-01 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_m"
categories:

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.
<pre>
[Hours]         [Activity]
  3.25          #127743 fix MSVC debugging builds
                #127743 work on reducing Win32/gcc Storable warnings
                #127743 testing warning fixes
  7.31          #130683 review the change and try to figure out what’s
                going wrong
                #130683 debugging
                #130683 debugging
                #130683 work up a possible fix, testing
                #130683 review code, try some alternatives
                #130683 look for related issues outside of %ENV
  4.10          #131844 (sec) comment
                #131844 (sec) debugging, consider approaches, make changes
                to fix and testing, comment with updated patches
  4.08          #131984 work on scope code to support fix
                #131984 finish up code, testing
                #131984 fix some issues, try to work up a test
  0.37          #132055 (sec) comment
  0.87          #132227 (sec) backport discussion with khw
                #132227 (sec) review khw’s fix, discussion
  0.58          #132764 review ticket and code
  2.54          #132800 profile mktables, look for optimization
                possibilities
                #132800 work on a mktables optimization
                #132800 validate results, fixes, profile again, comment
                with patch
  0.88          #132866 research, comment
  0.63          #132870 test that existing tests cover the issue, apply to
                blead
  0.65          #132878 review, testing, apply to blead
  0.98          #132893 more testing, apply to blead
  0.77          #132924 testing, research, apply to blead
  3.19          #132925 review patches, review mktables, mk_invlists.pl,
                work on alternate fixes, comment with patches
                #132925 testing, minor fix, more testing
                #132925 investigate unrelated test failure, rebase with
                conflicts, test_porting, apply to blead
  1.23          #132926 review, ask khw about it, abstract security
                discussion in #p5p
  0.69          #132929 review patch, research
                #132929 more review, testing, apply to blead
  1.30          #132943 review, testing, #p5p discussion of glibc
                duplocale issue probing, apply to blead
  0.28          #132949 testing, comment
  1.00          #132954 review, research and comment
  0.47          #132992 review patch and GNUmakefile
  1.35          #132999 testing, work up a fix and briefly comment with
                patch
  8.15          #133009 work on a fix, testing
                #133009 fix dmake makefile, testing, comment with patch
                #133009 try to break my patch, make stacksize more
                paranoid (while working on perlsecpolicy)
                #133009 more testing, comment with updated patch
                #133009 re-work probing some more, testing
                #133009 cross platform tests
  1.27          #133033 testing, comment
                #133033 research and comment
  0.90          clean up some security tickets
  1.13          khw’s perl.dll discussion on #p5p
  2.60          more perlsecpolicy
  1.57          more security backporting, ASAN discussion in #p5p
  0.17          more security support
  1.42          more try to backport, ask khw for help
  0.78          rebase maint votes (with lots of conflicts)
  0.50          rebase security branches
  0.48          research on security issues, comment
  1.30          review maint-votes and vote
  1.60          security issues: respond to downstream query, reproduce
                reported issue
  1.43          security list catch up, start per