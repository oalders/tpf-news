
---
title: "Maintaining Perl 5 (Tony Cook): Grant Report for May/June 2017"
author: Makoto Nozaki
type: post
date: 2017-08-04 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_g"
categories:
 - Grants

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.


**May report**

<pre>
Approximately 12 tickets were reviewed.

[Hours]         [Activity]
  0.57          . in @INC follow-up for File::Fetch #11
  1.98          #128207 debugging
                #128207 more debugging, comment
  0.53          #130591 review discussion, provide updated patch
  0.17          #131000 review discussion
  1.31          #131062 work on a patch
                #131062 debugging, comment
  1.32          #131070 testing, research
  1.02          #131214 research, comment
                #131214 research freebsd updates, comment
 12.22          #131221 debugging
                #131221 debugging, research, work on a fix
                #131221 debugging, threads support
                #131221 more threads support, debugging
                #131221 debugging
                #131221 threads debugging, trying for a fix
                #131221 stumble into a fix, leak testing
  0.82          #131337 test setup (network, package repo issues)
                #131337 start gcc build from ports
  3.43          #131345 try to reproduce, comment
                #131345 testing, research, comment with other things to
                try
                #131345 research, comment
  0.63          #131348 try to reproduce, comment
  1.92          review perldelta, some fixes
======
 25.92 hours total
</pre>

**June report**

<pre>
Approximately 31 tickets were reviewed, and 5 patches were
applied

[Hours]         [Activity]
  0.15          -FS switch thread research and comment
  1.57          #123981 research
                #123981 research, testing, apply to blead
  9.69          #124368 review discussion, try to understand compile/rune-
                time regexp compilation
                #124368 more reading code, work on adding warning
                #124368 testing, more reading code, fixes, tests, try to
                work out moving /o handling
                #124368 find some more non-warnable cases for the warning,
                debugging, fixes
                #124368 more code research, try to work up a test case,
                check for usethreads differences
  3.44          #127663 rebase, testing
                #127663 test fix, research
  1.32          #128207 review patches, testing, comment
  0.83          #129183 review discussion, research, apply to blead
  5.61          #130981 review ticket, try to understand deparse
                #130981 more understanding, try to make private flag
                #130981 work it out, work on deparse, and tests
                #130981 comment with patch
  0.82          #131050 testing, debugging, comment
  5.13          #131221 testing, writing regression tests
                #131221 more tests, testing, comment with patch
                #131221 more testing, apply to blead
                #131221 fix unthreaded builds
  1.32          #131263 testing, debugging, comment with patch
                #131263 testing, apply to blead
  0.58          #131345 research, comment
  1.88          #131522 debugging
                #131522 debugging
  2.33          #131526 reproduce, try to debug
                #131526 debugging, try a fix, try another fix, testing
                #131526 apply fix to blead
  1.57          #131544 research, comment
  3.74          #131546 research, produce simle patch, comment
                #131546 review discussion, work on patches, comment
                #131546 try to get unknown layers to produce a nonsense
                errno, wording for scalar open fails, testing, comment
  2.60          #131551 testing and comment
                #131551 more testing, comment, more testing in response to
