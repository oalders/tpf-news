
---
title: "Maintaining Perl 5 (Tony Cook): May 2019 Grant Report"
author: Matthias Bloch
type: post
date: 2019-06-21 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_m_3"
categories:

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.
<pre>
Approximately 30 tickets were reviewed, and 5 patches were
applied

[Hours]         [Activity]
  0.20          #121783 comment
  7.98          #122112 re-work to save the pid instead of the fd
                #122112 more re-work, testing
                #122112 debugging, re-work differently
                #122112 work on a simpler fix
                #122112 more work on simpler fix, testing
                #122112 more testing, commit message, comment with patch
                #122112 re-test and apply to blead
  0.55          #130585 rebase, testing, push to smoke branch
  0.98          #132782 review, testing, comment (via #p5p, since RT is
                misbehaving)
  0.89          #132964 comment
                #132964 oops, comment
                #132964 re-check, testing, apply to blead
  2.73          #133850 rebase, test old patch, work on fixing the other
                cases
                #133850 more work on fixing other cases, comment with
                patch
  2.30          #133958 discussion with khw, research, work on a fix,
                testing, write up a long commit message, comment with
                patch
                #133958 review smoke reports, comment
                #133958 re-test and apply patch
  7.02          #134031 review discussion, testing, reproduce
                #134031 more testing, research
                #134031 discussion in #p5p, testing, some debugging
                #134031 more testing, discussion, research expected C math
                precision, freebsd bug reports
  3.20          #134041 (sec) work on a fix
                #134041 (sec) work on a fix
                #134041 (sec) more work on a fix
  0.47          #134046 review responses, rebase, retest, apply to blead
  1.73          #134061 research and comment
                #134061 testing, fix a porting test failure, apply to
                blead
  0.78          #134065 more debugging
 13.90          #134072 debugging, comment
                #134072 debugging, research
                #134072 debugging
                #134072 debugging, work on a fix
                #134072 get test case working, try to make a test
                #134072 work out a test after some flailing around,
                testing
                #134072 track down a related issue (CVREF in stash instead
                of glob outside of main::, find it, comment with patch,
                add reference to experiment ticket
  0.45          #134080 research, comment
  3.99          #134111 debugging, research
                #134111 more research, comment
  0.73          #134114 review patch, research
  1.10          #134126 research, long comment
  0.85          bug report with patch for re debug and uniprops
  0.93          regexp runtime limits:
  2.02          regexp runtime limits: and CPU limits
  1.98          regexp runtime limits: donâ€™t tie it to re.pm <sigh>,
                tests, documentation (needs some polish)
  0.37          regexp runtime limits: exploring a maze of twisty macros
  0.65          regexp runtime limits: look for a better way of handling
                the debug flag
  0.98          regexp runtime limits: more efficiency, bang head against
                inconsistent naming
  1.55          regexp runtime limits: research
  2.30          regexp runtime limits: solve the macros, work on
                implementing limits
  2.07          regexp runtime limits: try to work up a memory limit test
                case, side-trip into security issue discussion with khw,
                find a test case, work on efficiency
  2.22          regexp runtime limits: work on memory li