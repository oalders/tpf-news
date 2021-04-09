
---
title: "Maintaining Perl 5: Grant Report for May 2015"
author: Karen Pauley
type: post
date: 2015-07-12 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_grant_repor_18"
categories:
 - Grants
 - Sign in

---

_Tony Cook writes:_

Approximately 37 tickets were reviewed or worked on, and 2 patches were applied.

bulk88's patch in [perl #123867] adds support for building perl in parallel on Win32 using dmake's -P option.  I've spent some time trying to improve it, but ended up running into limitations in dmake, in particular it seems to call child makefiles with -S preventing any parallel builds in the child.

If a GNU make makefile is added [perl #123440] in 5.23 this may provide a better base for parallel builds, though the current version referenced only builds using GCC.

| Hours |       Activity |
|  0.90    |      #122368 review, research and comment |
|  1.45    |      #123011 testing update, and comment |
|  0.40    |      #123264 review discussion, check code, produce simple |
|             |   patch and comment |
|  4.70    |      #123440 review discussion, setup, testing |
|             |   #123440 testing, fixes |
|             |   #123440 testing, make into patches |
|             |   #123440 work on MSVC support |
|  0.33    |      #123575 review patch and comment |
|  0.58    |      #123616 review discussion |
| 10.81   |       #123867 review new patch, comment |
|             |   #123867 review new patch, testing, comment |
|             |   #123867 work on make_ext_mk.pl |
|             |   #123867 more on make_ext_mk.pl |
|             |   #123867 deps, dynamic exts |
|             |   #123867 testing, try to convince dmake to parallel build child makefile |
|             |   #123867 more trying to work with dmake |
|  4.24    |      #123998 research and comment |
|             |   #123998 backport possible fix for smoke, testing, comment |
|             |   #123998 backport another fix, testing, push for smoke, comment |
|             |   #123998 housekeeping, testing. Apply to blead |
|  1.02    |      #124181 review, testing |
|  3.41    |      #124187 review core PL_compcv handling |
|             |   #124187 debugging, testing, push to blead, comment |
|  7.96    |      #124203 debugging, attempt a fix |
|             |   #124203 debugging |
|             |   #124203 testing, debugging |
|             |   #124203 code review, testing |
|             |   #124203 testing , debugging |
|             |   #124203 code review |
|  5.46    |      #124256 review code, documentation, debugging |
|             |   #124256 more review, work on a patch, comment |
|             |   #124256 adjust for lookahead, testing |
|             |   #124256 debugging, produce a new patch and comment |
|  0.18    |      #124349 review discussion |
|  0.45    |      #124368 review discussion and briefly comment |
|  1.02    |      #124409 debug and post patch upstream, comment  |
|  0.82    |      #124443 debugging and comment |
|  0.35    |      #124446 comment on future policy thread |
|  3.10    |      #125096 review, start gcc-5.1.0 build |
|             |   #125096 testing |
|             |   #125096 testing, try to adapt Configure |
|             |   #125096 fixes, testing, comment with patch |
|  0.43    |      #125106 review, research |
|  0.23    |      #125112 review and comment | 
|  2.05    |      #125115 review code, testing, comment with patch |
|             |   #125115 create a test, testing, comment, add to post-5.22 branch |
|  0.97    |      #125121 review code, find it's fixed, propose for maint- |
|             |   5.20, comment |
|  0.97    |      #125147 review and comment |
|             |   #125147 review new patch, add to post 5.22 and comment |
|  0.65    |      #125150 review, test and add to post 5.22 branch, comment |
|  4.43    |      #125203 setup, testing, debugging |
|             |   #125203 look for VMS NaN fix in 5.21.5 commits, discussion |
|             |   #125203 research, review generated code |
|  1.28    |      #125217 review, research |
|             |   #125217 review, testing