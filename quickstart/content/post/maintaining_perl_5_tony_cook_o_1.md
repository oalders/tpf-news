
---
title: "Maintaining Perl 5 (Tony Cook): October 2018 Grant Report"
author: Matthias Bloch
type: post
date: 2018-11-09 00:00:00 +0000 UTC
url: "/post/maintaining_perl_5_tony_cook_o_1"
categories:

---

This is a monthly report by Tony Cook on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.
<pre>
Approximately 49 tickets were reviewed, and 6 patches were
applied

[Hours]         [Activity]
  2.18          #125760 re-test branch and apply to blead
                #125760 perldelta
 11.90          #126706 get tests working, work on installer with
                install_name_tool
                #126706 re-work to use @rpath
                #126706 polish, testing
                #126706 fixes, testing, comment with patch
                #126706 look at fixing embed tests, re-work to avoid rpath
                #126706 more re-work, simplify, testing, polish, comment
                with patch
  1.27          #131649 (sec) find fix in blead/5.28, backport to 5.26,
                comment with patch
  1.73          #132147 (sec) review, look for supposed other project
                fixes
  2.44          #132782 review patches, work on tests, find an existing
                vec() bug, rebuild for debugging
                #132782 debugging, comment
  1.10          #133396 review ticket and code
  1.58          #133423 (sec) check blead, try to forward-port patch, test
                failures
                #133423 (sec) testing, comment
  0.35          #133439 retesting, apply to blead
  3.64          #133440 review discussion, research, work on improving the
                error messages
                #133440 more work on improving errors
  0.73          #133442 review, testing
                #133442 push to blead
  0.13          #133494 re-check, apply to blead
  0.75          #133511 review, research and comment
  1.40          #133519 try to reproduce, fail on win7, setup win10
                #133519 try to reproduce on win10
  1.58          #133523 (sec) review regexp code, comment
                #133523 (sec) review regexp code, comment
  2.54          #133535 debugging
                #133535 debugging, comment
  0.85          #133550 research, testing, apply to blead
  0.32          #133567 test, apply to blead
  0.95          #133582 comment with patch
  0.35          #133585 review and comment
  0.12          #133597 (sec) reject
  1.82          #133603 try to reproduce with gcc 8.1.0
                #133603 manage to reproduce, try to debug
  4.54          #133604 push skip, debugging, track down cause, testing
                #133604 work on regression tests, testing, apply to blead
  0.42          #133610 comment
  1.35          #133620 (sec) reproduce, work on bisect
                #133620 (sec) more work on bisect, comment
  0.27          #133630 comment
  0.17          ask khw about two of the issues, some tracking admin
  2.07          bang head against encoding issues trying to apply patches
  0.20          comment on private File::Slurp thread
  0.55          comment svtype thread
  0.50          debugging binmode
  1.37          diagnose :utf8 recv fatal tests on Win32 (binmode doesnâ€™t
                appear to work), encounter some build issues along the way
  2.05          feature sysio_bytes
  2.35          feature sysio_bytes, add some :utf8 fatal tests for recv,
                send, testing
  1.38          feature sysio_bytes, debugging
  1.95          feature sysio_bytes, more testing
  1.00          feature sysio_bytes, rebase, debugging, testing
  1.92          feature sysio_bytes: docs, testing
  1.33          find cause of encoding issues, finish patching
  1.07          more security fix checks
  1.78          polish, perldeta, open RFC 133610
  2.08          reply security email, work on checking security fixes for
                conflicts
  1.25          request CVE IDs
  2.03          respond to security email from sawyerx, try to work out
               