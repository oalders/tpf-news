
---
title: "Maintaining Perl 5 (Tony Cook): October 2020 Grant Report"
author: Matthias Bloch
type: post
date: 2020-11-05 00:00:00 +0000 UTC
url: "/post/tony-report-20-10"
categories:
 - Perl 5 Development

---

```

Approximately 34 tickets were reviewed, and 2 patches were
applied

[Hours]         [Activity]
  0.20          #16525 test and comment
  0.98          notifications catch up, including #17290
 40.03          #18005 polish
                #18005 add some more tests, work on commit messages
                #18005 minor re-work, some debugging, start a test run
                #18005 more polish, work on commit messages
                #18005 prep (including test-prep) for upstream patch, work
                on commit messages
                #18005 fix some un-updated config values and the problems
                encountered fixing that, clean up commits, (final?) test
                runs
                #18005 track down mingw failures (it’s using msvcrt.dll,
                sigh)
                #18005 re-work stat to avoid the broken msvcrt stat
                #18005 more work on stat(), discussion with xenu who was
                also working on it
                #18005 testing, debug and fix failing to close a handle,
                more stat tests for our implementation, start a MSVC test
                run, tests didn’t start
                #18005 re-work stat a bit, testing, start on utime() (to
                match stat)
                #18005 work utime, testing, minor fix, more testing
                #18005 research on leap seconds for time conversion,
                review changes and cleanup
                #18005 leap second discussion with xenu, re-work time_t &lt;&gt;
                FILETIME conversion to hopefully support leap seconds
                correctly (though we don’t enable leap seconds for
                perl.exe, XS or the registry might enable them), start
                make test runs
                #18005 commit time re-work, steal xenu’s verson check for
                unprivileged symlink creation, remove
                ${^WIN32_SLOPPY_STAT}, start test run
                #18005 treat junctions as symlinks for lstat and readlink,
                perlport updates, push to smoke-me
                #18005 re-work the regen_config_h target, after some mis-
                steps get the VC config working
                #18005 update other makefiles, testing, cleanup, push for
                CI
                #18005 review smoke results, make some fixes and start
                test run
                #18005 fix some more problems, more testing, push to check
                for more smoke
                #18005 ask George Greer to trigger a smoke for the pushed
                smoke-me
                #18005 try to diagnose win10 smoke failures
                #18005 more diag and fixes, push for another smoke run
                #18005 review latest smokes
                #18005 commit clean up, start a test run
  1.52          #18133 rebase, add tests, documentation updates, testing
                and push to update PR
  0.17          #18150 review and apply to blead
  0.12          #18169 briefly comment
  0.60          #18182 testing, briefly comment
                #18182 review
  2.07          #18198 review and comment
                #18198 re-check, research and comment
  0.22          #18205 briefly comment
  0.07          #18210 review
  0.30          #18212 review
  0.98          #18221 review, discussion with toddr
  0.03          #18231 review and apply to blead
  0.43          #18232 review, comment briefly
  1.46          #18256 research and comment
                review notifications, mostly #18256
  1.53          #18257 review code and move the dead code, test builds,
                make a bad build fail at compile time, push for CI
                #18257 review CI results, make PR 18266
  0.27          check notifications
  0.95          clean up tc function changes, update the associated dfly
                bug, push for smoke