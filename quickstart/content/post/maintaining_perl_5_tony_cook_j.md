
---
title: "Maintaining Perl 5 (Tony Cook): January 2018 Report"
author: Makoto Nozaki
type: post
date: 2018-02-12 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_j"
categories:
 - Grants
 - Perl 5 Development
 - Sign in

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.

<pre>Approximately 29 tickets were reviewed, and 6 patches were
applied

[Hours]         [Activity]
 34.99          #127743 try to work up makefile rules for new stack limit
                handling
                #127743 more makefile rules, work on fetching config from
                var instead of a constant
                #127743 more stack size calculation re-work
                #127743 more makefile rules, experiment with stack size
                probing
                #127743 testing, debugging, more stack size probing,
                documentation
                #127743 debugging, non-core testing
                #127743 non-core debugging, fix calculation for stack
                frames on 5.8.9, debug 5.10.0
                #127743 for 5.10.0 debugging
                #127743 track down to dependency issue, work on fix,
                testing
                #127743 testing, minor fixes
                #127743 more non-core testing, clean up, start on Win32
                build updates
                #127743 check non-Storable perl builds, back to Win32
                builds, locale discussion with khw
                #127743 more Win32 build work, nmake, gmake, start dmake
                #127743 fix deps, testing
                #127743 finish win32 testing, commit clean up, more regexp
                tests, start smoke
                #127743 review smoke results, diagnose cygwin failures,
                debugging
                #127743 continue debugging, find part of problem (no depth
                checks on retrieving array/refs)
                #127743 adjust max depth for cynwin64, prep to test on
                fedora
                #127743 try to repro fedora failures, canâ€™t, try to repro
                freebsd issues
                #127743 avoid testing huge levels of recursion
                #127743 review smoke results, be more conservative on all
                platforms, testing, push new smoke candidate
</pre>

<pre>
  0.46          #131844 (sec) refresh memory and comment
                #131844 (sec) review and comment
  1.45          #131878 re-review discussion, discussion on #p5p, minor
                change to patch, testing, apply to blead, make ticket
                public and close
  1.32          #131954 testing, apply to blead, make ticket public
  0.35          #132055 (sec) review and comment
  1.74          #132063 (sec) review, check possible related issues
                #132063 (sec) more check possible issues, comment
  2.33          #132147 (sec) review, research, work on patch
                #132147 (sec) testing, clean up, comment with patch
  0.30          #132189 (sec) reproduce and forward privately upstream
  1.03          #132227 (sec) debugging
  2.26          #132533 review patches, struggle applying them, comment
                #132533 review new patches, testing, fixes, apply to blead
  0.52          #132602 security assessment, make public, comments
  1.40          #132618 review, debugging, make public and link to meta
                ticket, comment
  0.37          #132625 (sec) testing, comment
  8.53          #132640 (sec) delve into sub-parsing, try not to delve too
                deeply
                #132640 try to understand the parser, research
                #132640 (sec) work on a fix, debugging
                #132640 (sec) get a working patch, try to improve it, make
                ticket public, merge into 125351, comment with patch
  1.74          #132648 testing, research
                #132648 mail cygwin mailing list
                #132648 testing
  0.18          #1326