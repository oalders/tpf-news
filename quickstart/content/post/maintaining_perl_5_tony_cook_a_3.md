
---
title: "Maintaining Perl 5 (Tony Cook): August 2019 Grant Report"
author: Matthias Bloch
type: post
date: 2019-09-16 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_a_3"
categories:

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.
<pre>
Approximately 54 tickets were reviewed, and 9 patches were
applied

[Hours]         [Activity]
  0.13          #118551 track down fix and close
  1.10          #124256 rebase, testing, apply to blead
  0.12          #125096 re-check and close
  1.73          #126991 work on a fix
  0.72          #131136 track down fix and close, and for other Storable
                tickets too…
  0.50          #131990 track down the fix and close
  0.13          #131999 track down fix and close
  1.26          #132777 review and some re-work, testing, apply to blead
                #132777 research and comment
  0.68          #133695 re-test and apply to blead
  0.33          #133878, #134269 make public
  1.78          #133981 minor fix, testing, try to diagnose an unrelated
                porting test failure, comment
  0.78          #134138 re-work patch and comment
  1.39          #134171 testing
                #134171 research, comment
                #134171 research and comment
  0.25          #134219 review and request more information
                #134219 review and merge into #134266
  3.39          #134230 debugging, research, work on a fix
                #134230 work up a regression test, comment with patch
                #134230 testing, apply to blead
  0.38          #134238 review and close
  1.70          #134241 testing
                #134241 more testing
  2.02          #134265 re-work supplied patches and comment
                #134265 re-test modified patch and apply to blead
  3.04          #134266 work up a basic fix… which breaks tests
                #134266 work out problem, testing, comment with patch
                #134266 re-test and apply to blead
  0.37          #134269 (sec) comment
  0.17          #134270 merge into 134126
  1.62          #134277 debugging
  1.02          #134288 review, testing
                #134288 finish up, apply to blead
  1.15          #134290 review and comment
                #134290 research, comment and close.
  0.17          #134295 re-revert the exit(0) fix, since I didn’t make any
                progress in fixing the issues
  1.12          #134312 research and comment
                #134312 research and comment some more
  0.80          #134313 research and comment
  0.50          #134320 review for proposal to backport, along with
                several other commits
  4.72          #134325 debugging
                #134325 debugging, discussion with khw, comment
                #134325 debugging
                #134325 research, testing and comment
  0.13          #134326 also fixed by khw’s 134329 fix, comment
  0.10          #134327 also fixed by khw’s 134329 fix, comment
  0.40          #134328 testing, try khw’s 134329 fix, comment
  1.42          #134329 testing, start bisect
  0.43          #134337 comment
  1.88          #134349 work on understanding code, work up a test, fix
  0.85          #134360 review, research, apply to blead, PR to metaconfig
  0.40          #134361 comment
                #134361 research, follow-up
  1.98          #134362 (sec) comment
                #134362 (sec) follow-up
                #134362 (sec) work on a fix
                #134362 (sec) testing, comment with patch
  4.91          #134363 review and request more information
                #134363 research, fail to reproduce
                #134363 more research
                #134363 more research into gcc implementation
                #134363 more research, comment
                #134363 comment
  0.37          #134365 review discussion
                #134365 comment with simple patch
  1.37          #134369 (sec) research, testing, comment