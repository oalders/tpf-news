
---
title: "Maintaining Perl 5 (Tony Cook): Grant Report for September 2017"
author: Makoto Nozaki
type: post
date: 2017-10-25 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_g_3"
categories:
 - Grants
 - Perl 5 Development

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.


<pre>Approximately 47 tickets were reviewed, and 7 patches were applied.

[Hours]         [Activity]
  7.13          #122112 work on a fix (and reentr discussion with khw)
                #122112 more work on a fix
                #122112 debugging
                #122112 debugging
                #122112 debugging, look for a different approach
  2.31          #124256 re-work patch
                #124256 more re-work patch, testing
  3.94          #125351 try to understand the parser
                #125351 try restoring stack, debug crashes
                #125351 experiments, testing, debugging
  1.82          #126706 re-work patch a bit, testing, debugging
                #126706 debugging, fix taint issue
  7.01          #127663 re-check, minor commit message changes, comment
                with new patch
                #127663 testing
                #127663 more testing, apply to blead
                #127663 reproduce, work on a fix, testing
                #127663 more testing
                #127663 comment on freebsd renameat() ticket
  1.98          #129990 review, testing, coverage checks, benchmarking,
                apply to blead
  2.36          #131582 (sec) work on CVE details, comment
                #131582, #131598 (sec) work with downstream
                #131582, #131598 (sec) work with downstream
  0.57          #131598 (sec) work on CVE details, comment
  3.32          #131648 (sec) debugging
                #131648 (sec) more debugging, briefly comment
  0.77          #131665 (sec) work on CVE details, comment
                #131665 (sec) comment
  1.58          #131683 comment
                #131683 comment
                #131683 review patch, discussion in #p5p
  0.70          #131685 research
  1.15          #131746 retesting, apply to blead
  1.45          #131777 re-base, re-check, re-test and apply to blead
  2.21          #131894 research how to fix non-hex digit handling for
                hex^Wbinaryfp numbers
                #131894 work on a fix, testing
                #131894 add more tests, testing, comment with patch
  1.18          #131954 (sec) research and comment
  1.51          #131982 research, comment
                #131982 work on a patch and comment with patch
  0.95          #131984 comment
                #131984 re-review, testing, apply to blead
  0.67          #132001 debugging, reject with comment
  3.24          #132008 testing, minor fixes, apply to blead
                #132008 look into regression, start setup of test VM
                #132008 discussion with Tux, finish VM setup, also discuss
                spam when smokes run in PERL_UNICODE= LC_ALL=some-utf8-
                locale mode, work on a fix and apply to blead
  0.17          #132017 (sec) comment
  1.74          #132055 (sec) reproduce, debugging
                #132055 (sec) debugging, ask khw to look at it
  0.20          #132063 (sec) debugging, ask khw again
  0.75          #132077 review patches and comment
  3.02          #132087 testing, research, comment with patch, comment
                with fixed patch
                #132087, #127663 re-test various portability fixes and
                apply to blead
  0.28          #132094 research and comment
  1.53          #132105 research and comment
                #132105 look at making SIG_SIZE a variable, check CPAN for
                usage and comment
  0.05          #132134 review and comment
  0.33          #132137 research and comment
  1.03          #132138 work to understand failure, fix and apply to blead
  0.20          #132139 track down when minitest was introduced and
         