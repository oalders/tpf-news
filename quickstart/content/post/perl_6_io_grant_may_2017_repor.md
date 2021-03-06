
---
title: "Perl 6 IO Grant: May 2017 Report (Complete)"
author: Coke
type: post
date: 2017-06-06 00:00:00 +0000 UTC
url: "/post/perl_6_io_grant_may_2017_repor"
categories:
 - Grants

---

*Zoffix Znet provided this report on May 29, 2017*

*The Grants Committee will vote on its completion and payment.
If you have feedback or questions, please comment here*

# COMPLETION Report / Perl 6 IO TPF Grant

This document is the May, 2017 progress report for [TPF Standardization,
Test Coverage, and Documentation of Perl 6 I/O Routines
grant](http://news.perlfoundation.org/2017/01/grant-proposal-standardization.html). I believe I reasonably satisfied the goals of the grant and consider it
completed. This is the final report and may reference some of the
work/commits previously mentioned in monthly reports.

# Thank You!

I'd like to thank all the donors that
[support The Perl Foundation](https://donate.perlfoundation.org/)
who made this grant possible. It was a wonderful learning experience for me,
and it brings me joy to look back and see Perl 6 improved due to this grant.

Thank You!

# Completeness Criteria

Here are the original completeness criteria (in bold) that are listed on the
original grant proposal and my comments on their status:

- **rakudo repository will contain the IO Action Plan document and it will
  be fully implemented.**
  The promised [document
  exists](https://github.com/rakudo/rakudo/blob/nom/docs/2017-IO-Grant--Action-Plan.md). It's fully implemented except for three items that I listed on the
  IO Action Plan, but which are currently a bit beyond my skill level to
  implement. I hope to do them eventually, but outside the scope of this grant.
  They are:
  - [IO::Handle's Closed status](https://github.com/rakudo/rakudo/blob/nom/docs/2017-IO-Grant--Action-Plan.md#iohandles-closed-status-issue-for-discussion). My original proposal
  would cause some perfomance issues, so it was decided to improve MoarVM
  errors instead.
  - [Optimize multiple stat calls](https://github.com/rakudo/rakudo/blob/nom/docs/2017-IO-Grant--Action-Plan.md#iopath-routines-that-involve-a-stat-call-issue-for-discussion). This
  involves creating a new nqp op, with code for it implemented in MoarVM and
  JVM backends.
  - [Use typed exceptions instead of X::AdHoc](https://github.com/rakudo/rakudo/blob/nom/docs/2017-IO-Grant--Action-Plan.md#use-typed-exceptions-instead-of-xadhoc-issue-for-discussion). I
  made typed exceptions be thrown whereever I could. The rest require VM-level
  exceptions and is on the same level as the handle closed status issue
  (first item above).
- **All of the I/O routines will have tests in roast and documented on
  docs.perl6.org. If any of the currently implemented but unspecced routines
  are decided against being included in Perl 6 Language, their implementation
  will no longer be available in Rakudo.**
  To the best of my knowledge, this is completed in full.
- **The test coverage tool will report all I/O routines as covered and the
  information will be visible on [perl6.wtf](https://perl6.wtf) (Perl 6's
  Wonderful Test Files)
  website. Note: due to current experimental status of the coverage tool, its
  report may still show some lines or conditionals untested despite them
  actually being tested; however, it will show the lines where routines' names
  are specified as covered.**
  To the best of my knowledge, all IO routines currently have tests covering
  them. Due to its experimental status, the coverage tool shows some
  attributes as uncovered. I did manually verify all the attributes/routines
  whose names the tool shows as uncovered contain tests for them.
  One exception is `IO::Notification` type (and `IO::Path.watch` method).
  While it has full coverage for
  OSX operating system, it lacks it for other OSes. I tried writing some tests for it, but
  it looks like the behaviour of the nqp op handling these is broken
  on Linux and the class needs more work.

# Extra Deliverables

I produced these extra deliverables while working on the grant:

- **[The Definitive I/O Guide](https://docs.perl6.org/language/io-guide).**
    Providing tutorial-like documentation for Perl 6's I/O, including
    documenting some of the bad practices I noticed in the ecosystem
    (and even a Perl 6 book!) and the correct way to perform those tasks.
    (N.B. as I write this report, the guide could still use a few extra
    sections to be considered "The Definitive"; I'll write them in upcoming
    weeks)
- **Performance improvements.** I made 23 performance-enhancing commits, with
    many commits making things more than 200% faster, with highest improvement
    making a routine 6300% faster.
- **[`Trait::IO` module](https://github.com/zoffixznet/perl6-Trait-IO).**
    Provides `does auto-close` pseudo-trait to simplify closing of IO handles.
- **[`IO::Path::ChildSecure` module](https://github.com/zoffixznet/perl6-IO-Path-ChildSecure).** Due to
    large ecosystem usage, `IO::Path.child` was left as is until 6.d language,
    at which point it will be made secure (as outlined in the IO Plan). This
    module provides the secure version in the mean time.
- **[`IO::Dir` module](https://github.com/zoffixznet/perl6-IO-Dir).** Provides
    `IO::Path.dir`-like functionality, with ability to close open directory
    without needing to fully exhaust the returned `Seq`.
- **[`Die` module](https://github.com/zoffixznet/perl6-Die).** Implements
    Perl-5-like behaviour for `&die` routine.
- **The "Map of Perl 6 Routines"** (or rather the "table") is available on
    [map.perl6.party](https://map.perl6.party/) with its code in
    [perl6/routine-map repo](https://github.com/perl6/routine-map). In near
    future, I plan to use it to identify incorrect or incomplete entries in
    our [documentation](https://docs.perl6.org)

In addition, I plan to complete these modules some time in the future; the
ideas for them were birthed while working on the grant:
- **`NL` module.** Targeted for use in one liners, the module will provide
  `$*NL` dynvar that behaves like Perl 5's `$.` variable (providing current
    `$*ARGFILES`'s file's line number). Its implementation became possible
    thanks to newly-implemented
    [`IO::CatHandle` type](https://docs.perl6.org/type/IO::CatHandle)
- **`FastIO` module.** A re-imagination of core IO, the biggest part of which
    will be the removal of (user-exposed) use of `IO::Spec::*` types
    and `$*SPEC`
    variable, which???it is believed???will provide improved performance over
    core IO. The module is a prototype for some of the proposals that were
    made during the IO grant and if it offers significant improvements over
    core IO, its ideas will be used by core IO in future language versions.

# Work Performed in May

For the work done in May,
many of my commits went into going through the IO routine list, and adding
missing tests and documentation, along with fixing bugs (and reporting new
ones I found).

The major work was implementation of the
[IO::CatHandle class](https://docs.perl6.org/type/IO::CatHandle) that fixed
all of the bugs and NYIs with the
[`$*ARGFILES`](/language/variables#index-entry-%24%2AARGFILES). This work
saw the addition of [372 lines of code](https://github.com/rakudo/rakudo/blob/608e886a9b01c976ee6002c9a0341c289b045f98/src/core/IO/CatHandle.pm),
[800 lines of tests](https://github.com/perl6/roast/blob/d36b8371304a5e62aa5669c1f2cd85c73a8de343/S32-io/io-cathandle.t) and [793 lines of documentation](https://github.com/perl6/doc/blob/856e8465096f51c699c6c68ff8dba4ea641f6b03/doc/Type/IO/CatHandle.pod6).

### Work by Other Core Members

jnthn++ completed the [handle encoding
refactor](https://github.com/rakudo/rakudo/compare/f2fca0c8c2...41bb13722c)
that will eventually let us get rid of using libuv for syncronous IO and, more
importantly, allow us to support user-defined encoders/decoders.

Along with fixing a bunch of bugs, this work altered the performance landscape
for IO operations (i.e. some operations may now be a bit faster, others
a bit slower), though overall the performance appeared to stay the same.

### Tickets Fixed

- [RT#130456 `$*HOME blows up if HOME isn't set`](https://rt.perl.org/Ticket/Display.html?id=130456)
- [RT#126935 `bad .perl for paths with pipe characters`](https://rt.perl.org/Ticket/Display.html?id=126935) (was already fixed; just added an extra test)
- [RT#131185 IO::Path.perl doesn't roundtrip](https://rt.perl.org/Ticket/Display.html?id=131185)
- [RT#130454 tmpdir tries to change the current working directory](https://rt.perl.org/Ticket/Display.html?id=130454) (routine was removed as part of IO work)
- [RT#130455 Should I be able to change the temporary directory?](https://rt.perl.org/Ticket/Display.html?id=130455)
- [RT#130715 IO::Handle::close shouldn't decide what's a failure](https://rt.perl.org/Ticket/Display.html?id=130715) (not a bug, but user's confusion; explained and rejected)
- [RT#131242 IO::Path.move/.copy hangs/trashes file when target/source are same](https://rt.perl.org/Ticket/Display.html?id=131242)

# Grant Commits

During this grant, I've made 417 commits, that are: 134 Rakudo commits + 23
performance-enchancing Rakudo commits + 114 Perl 6 Specification commits + 146
documentation commits,

## Performance Rakudo Commits

I've made 23 performance enchancing commits to Rakudo's repository:

- [4032953](https://github.com/rakudo/rakudo/commit/4032953) `Make IO::Handle.open 75% faster`
- [dcf1bb2](https://github.com/rakudo/rakudo/commit/dcf1bb2) `Make IO::Spec::Unix.rel2abs 35% faster`
- [c13480c](https://github.com/rakudo/rakudo/commit/c13480c) `IO::Path.slurp: make 12%-35% faster; propagate Failures`
- [0e36bb2](https://github.com/rakudo/rakudo/commit/0e36bb2) `Make IO::Spec::Win32!canon-cat 2.3x faster`
- [c6fd736](https://github.com/rakudo/rakudo/commit/c6fd736) `Make IO::Spec::Win32.is-absolute about 63x faster`
- [894ba82](https://github.com/rakudo/rakudo/commit/894ba82) `Make IO::Spec::Win32.split about 82% faster`
- [277b6e5](https://github.com/rakudo/rakudo/commit/277b6e5) `Make IO::Spec::Unix.rel2abs 2.9x faster`
- [74680d4](https://github.com/rakudo/rakudo/commit/74680d4) `Make IO::Path.is-absolute about 80% faster`
- [ff23416](https://github.com/rakudo/rakudo/commit/ff23416) `Make IO::Path.is-relative about 2.1x faster`
- [d272667](https://github.com/rakudo/rakudo/commit/d272667) `Make IO::Spec::Unix.join about 40% faster`
- [50429b1](https://github.com/rakudo/rakudo/commit/50429b1) `Make IO::Handle.put($x) about 5%-35% faster`
- [204ea59](https://github.com/rakudo/rakudo/commit/204ea59) `Make &say(**@args) 70%??? faster`
- [6d7fc8e](https://github.com/rakudo/rakudo/commit/6d7fc8e) `Make &put(**@args) up to 70% faster`
- [76af536](https://github.com/rakudo/rakudo/commit/76af536) `Make 1-arg IO::Handle.say up to 2x faster`
- [aa72bde](https://github.com/rakudo/rakudo/commit/aa72bde) `Remove dir's :absolute and :Str; make up to 23% faster`
- [48cf0e6](https://github.com/rakudo/rakudo/commit/48cf0e6) `Make IO::Spec::Cygwin.is-absolute 21x faster`
- [c96727a](https://github.com/rakudo/rakudo/commit/c96727a) `Fix combiners on SPEC::Win32.rel2abs; make 6% faster`
- [0547979](https://github.com/rakudo/rakudo/commit/0547979) `Make IO::Spec::Unix.path consistent and 4.6x faster`
- [8992af1](https://github.com/rakudo/rakudo/commit/8992af1) `Fix IO::Spec::Win32.path and make 26x faster`
- [7d6fa73](https://github.com/rakudo/rakudo/commit/7d6fa73) `Make IO::Spec::Win32.catpath 47x faster`
- [494659a](https://github.com/rakudo/rakudo/commit/494659a) `Make IO::Spec::Win32.join 26x faster`
- [6ca702f](https://github.com/rakudo/rakudo/commit/6ca702f) `Make IO::Spec::Unix.splitdir 7.7x faster`
- [2816ef7](https://github.com/rakudo/rakudo/commit/2816ef7) `Make IO::Spec::Win32.splitdir 25x faster`

## Non-Performance Rakudo Commits

Other than perf commits, I've also made 134 commits to the Rakudo's repository:

- [dd4dfb1](https://github.com/rakudo/rakudo/commit/dd4dfb1) `Fix crash in IO::Special .WHICH/.Str`
- [76f7187](https://github.com/rakudo/rakudo/commit/76f7187) `Do not cache IO::Path.e results`
- [212cc8a](https://github.com/rakudo/rakudo/commit/212cc8a) `Remove IO::Path.Bridge`
- [a01d679](https://github.com/rakudo/rakudo/commit/a01d679) `Remove IO::Path.pipe`
- [55abc6d](https://github.com/rakudo/rakudo/commit/55abc6d) `Improve IO::Path.child perf on *nix`
- [4fdebc9](https://github.com/rakudo/rakudo/commit/4fdebc9) `Make IO::Spec::Unix.split 36x Faster`
- [0111f10](https://github.com/rakudo/rakudo/commit/0111f10) `Make IO::Spec::Unix.catdir 3.9x Faster`
- [fa9aa47](https://github.com/rakudo/rakudo/commit/fa9aa47) `Make R::I::SET_LINE_ENDING_ON_HANDLE 4.1x Faster`
- [c360ac2](https://github.com/rakudo/rakudo/commit/c360ac2) `Fix smartmatch of Cool ~~ IO::Path`
- [0c7e4a0](https://github.com/rakudo/rakudo/commit/0c7e4a0) `Do not capture args in .IO method`
- [9d8d7b2](https://github.com/rakudo/rakudo/commit/9d8d7b2) `Log all changes to plan made during review period`
- [87987c2](https://github.com/rakudo/rakudo/commit/87987c2) `Remove `role IO` and its .umask method`
- [36ad92a](https://github.com/rakudo/rakudo/commit/36ad92a) `Remove 15 methods from IO::Handle`
- [a5800a1](https://github.com/rakudo/rakudo/commit/a5800a1) `Implement IO::Handle.spurt`
- [aa62cd5](https://github.com/rakudo/rakudo/commit/aa62cd5) `Remove &tmpdir and &homedir`
- [a0ef2ed](https://github.com/rakudo/rakudo/commit/a0ef2ed) `Improve &chdir, &indir, and IO::Path.chdir`
- [ca1acb7](https://github.com/rakudo/rakudo/commit/ca1acb7) `Fix race in &indir(IO::Path ???)`
- [2483d68](https://github.com/rakudo/rakudo/commit/2483d68) `Fix regression in &chdir's failure mode`
- [5464b82](https://github.com/rakudo/rakudo/commit/5464b82) `Improve &*chdir`
- [4c31903](https://github.com/rakudo/rakudo/commit/4c31903) `Add S32-io/chdir-process.t to list of test files to run`
- [cb27bce](https://github.com/rakudo/rakudo/commit/cb27bce) `Clean up &open and IO::Path.open`
- [099512b](https://github.com/rakudo/rakudo/commit/099512b) `Clean up and improve all spurt routines`
- [b62d1a7](https://github.com/rakudo/rakudo/commit/b62d1a7) `Give $*TMPDIR a container`
- [b1e7a01](https://github.com/rakudo/rakudo/commit/b1e7a01) `Implement IO::Path.extension 2.0`
- [15a25da](https://github.com/rakudo/rakudo/commit/15a25da) `Fix ambiguity in empty extension vs no extension`
- [50aea2b](https://github.com/rakudo/rakudo/commit/50aea2b) `Restore IO::Handle.IO`
- [966a7e3](https://github.com/rakudo/rakudo/commit/966a7e3) `Implement IO::Path.concat-with`
- [94a6909](https://github.com/rakudo/rakudo/commit/94a6909) `Clean up IO::Spec::Unix.abs2rel a bit`
- [a432b3d](https://github.com/rakudo/rakudo/commit/a432b3d) `Remove IO::Path.abspath (part 2)`
- [954e69e](https://github.com/rakudo/rakudo/commit/954e69e) `Fix return value of IO::Special methods`
- [67f06b2](https://github.com/rakudo/rakudo/commit/67f06b2) `Run S32-io/io-special.t test file`
- [a0b82ed](https://github.com/rakudo/rakudo/commit/a0b82ed) `Make IO::Path::* actually instantiate a subclass`
- [0c8bef5](https://github.com/rakudo/rakudo/commit/0c8bef5) `Implement :parent in IO::Spec::Cygwin.canonpath`
- [0a442ce](https://github.com/rakudo/rakudo/commit/0a442ce) `Remove type constraint in IO::Spec::Cygwin.canonpath`
- [b4358af](https://github.com/rakudo/rakudo/commit/b4358af) `Delete code for IO::Spec::Win32.catfile`
- [e681498](https://github.com/rakudo/rakudo/commit/e681498) `Make IO::Path throw when path contains NUL byte`
- [6a8d63d](https://github.com/rakudo/rakudo/commit/6a8d63d) `Implement :completely param in IO::Path.resolve`
- [b6838ee](https://github.com/rakudo/rakudo/commit/b6838ee) `Remove .f check in .z`
- [184d499](https://github.com/rakudo/rakudo/commit/184d499) `Make IO::Handle.Supply respect handle's mode`
- [f1b4af7](https://github.com/rakudo/rakudo/commit/f1b4af7) `Implement IO::Handle.slurp`
- [90da80f](https://github.com/rakudo/rakudo/commit/90da80f) `Rework read methods in IO::Path/IO::Handle`
- [8c09c84](https://github.com/rakudo/rakudo/commit/8c09c84) `Fix symlink and link routines`
- [da1dea2](https://github.com/rakudo/rakudo/commit/da1dea2) `Fix &symlink and &link`
- [7f73f92](https://github.com/rakudo/rakudo/commit/7f73f92) `Make IO::Path.new-from-absolute-path private`
- [ff97083](https://github.com/rakudo/rakudo/commit/ff97083) `Straighten up rename, move, and copy`
- [0d9ecae](https://github.com/rakudo/rakudo/commit/0d9ecae) `Remove multi-dir &mkdir`
- [6ee71c2](https://github.com/rakudo/rakudo/commit/6ee71c2) `Coerce mode in IO::Path.mkdir to Int`
- [d46e8df](https://github.com/rakudo/rakudo/commit/d46e8df) `Add IO::Pipe .path and .IO methods`
- [c01ebea](https://github.com/rakudo/rakudo/commit/c01ebea) `Make IO::Path.mkdir return invocant on success`
- [1f689a9](https://github.com/rakudo/rakudo/commit/1f689a9) `Fix up IO::Handle.Str`
- [490ffd1](https://github.com/rakudo/rakudo/commit/490ffd1) `Do not use self.Str in IO::Path errors`
- [40217ed](https://github.com/rakudo/rakudo/commit/40217ed) `Swap .child to .concat-with in all the guts`
- [fd503f8](https://github.com/rakudo/rakudo/commit/fd503f8) `Revert "Remove `role IO` and its .umask method"`
- [c95c4a7](https://github.com/rakudo/rakudo/commit/c95c4a7) `Make IO::Path/IO::Special do IO role`
- [214198b](https://github.com/rakudo/rakudo/commit/214198b) `Implement proper args for IO::Handle.lock`
- [9a2446c](https://github.com/rakudo/rakudo/commit/9a2446c) `Move Bool return value to signature`
- [51e4629](https://github.com/rakudo/rakudo/commit/51e4629) `Amend rules for last part in IO::Path.resolve`
- [b8458d3](https://github.com/rakudo/rakudo/commit/b8458d3) `Reword `method child` for cleaner code`
- [1887114](https://github.com/rakudo/rakudo/commit/1887114) `Implement IO::Path.child-secure`
- [9d8e391](https://github.com/rakudo/rakudo/commit/9d8e391) `Fix IO::Path.resolve with combiners; timotimo++`
- [0b5a41b](https://github.com/rakudo/rakudo/commit/0b5a41b) `Rename IO::Path.concat-with to .add`
- [a98b285](https://github.com/rakudo/rakudo/commit/a98b285) `Remove IO::Path.child-secure`
- [8bacad8](https://github.com/rakudo/rakudo/commit/8bacad8) `Implement IO::Path.sibling`
- [7112a08](https://github.com/rakudo/rakudo/commit/7112a08) `Add :D on invocant for file tests`
- [b2a64a1](https://github.com/rakudo/rakudo/commit/b2a64a1) `Fix $*CWD inside IO::Path.dir's :test Callable`
- [6fa4bbc](https://github.com/rakudo/rakudo/commit/6fa4bbc) `Straighten out &slurp/&spurt/&get/&getc/&close`
- [34b58d1](https://github.com/rakudo/rakudo/commit/34b58d1) `Straighten out &lines/&words`
- [d0cd137](https://github.com/rakudo/rakudo/commit/d0cd137) `Make dir take any IO(), not just Cool`
- [7412184](https://github.com/rakudo/rakudo/commit/7412184) `Make $*HOME default to Nil, not Any`
- [475d9bc](https://github.com/rakudo/rakudo/commit/475d9bc) `Fix display of backslashes in IO::Path.gist`
- [6ef2abd](https://github.com/rakudo/rakudo/commit/6ef2abd) `Revert "Fix display of backslashes in IO::Path.gist"`
- [134efd8](https://github.com/rakudo/rakudo/commit/134efd8) `Fix .perl for IO::Path and subclasses`
- [69320e7](https://github.com/rakudo/rakudo/commit/69320e7) `Fix .IO on :U of IO::Path subclasses`
- [eb8d006](https://github.com/rakudo/rakudo/commit/eb8d006) `Make IO::Handle.iterator a private lines iterator`
- [08a8075](https://github.com/rakudo/rakudo/commit/08a8075) `Fix IO::Path.copy/move when source/target are same`
- [973338a](https://github.com/rakudo/rakudo/commit/973338a) `Fix IO::Handle.comb/.split; make them .slurp`
- [b43ed18](https://github.com/rakudo/rakudo/commit/b43ed18) `Make IO::Handle.flush fail with typed exceptions`
- [276d4a7](https://github.com/rakudo/rakudo/commit/276d4a7) `Remove .tell info in IO::Handle.gist`
- [f4309de](https://github.com/rakudo/rakudo/commit/f4309de) `Fix IO::Spec::Unix.is-absolute for combiners on /`
- [06d8800](https://github.com/rakudo/rakudo/commit/06d8800) `Fix crash when setting .nl-in ...`
- [7e9496d](https://github.com/rakudo/rakudo/commit/7e9496d) `Make IO::Handle.encoding settable via .new`
- [95e49dc](https://github.com/rakudo/rakudo/commit/95e49dc) `Make IO::Handle.open respect attribute values`
- [6ed14ef](https://github.com/rakudo/rakudo/commit/6ed14ef) `Remove `:directory` from IO::Spec::*.split`
- [9021a48](https://github.com/rakudo/rakudo/commit/9021a48) `Make IO::Path.parts a Map instead of Hash`
- [a282b8c](https://github.com/rakudo/rakudo/commit/a282b8c) `Fix IO::Handle.perl.EVAL roundtrippage`
- [a412788](https://github.com/rakudo/rakudo/commit/a412788) `Make IO::Path.resolve set CWD to $!SPEC.dir-sep`
- [84502dc](https://github.com/rakudo/rakudo/commit/84502dc) `Implement $limit arg for IO::Handle.words`
- [613bdcf](https://github.com/rakudo/rakudo/commit/613bdcf) `Make IO::Handle.print/.put sig consistent`
- [0646d3f](https://github.com/rakudo/rakudo/commit/0646d3f) `Allow no-arg &prompt`
- [4a8aa27](https://github.com/rakudo/rakudo/commit/4a8aa27) `Implement IO::CatHandle.close`
- [4ad8b17](https://github.com/rakudo/rakudo/commit/4ad8b17) `Implement IO::CatHandle.get`
- [3b668b6](https://github.com/rakudo/rakudo/commit/3b668b6) `Implement IO::CatHandle.getc`
- [25b664a](https://github.com/rakudo/rakudo/commit/25b664a) `Implement IO::CatHandle.words`
- [7ebc386](https://github.com/rakudo/rakudo/commit/7ebc386) `Implement IO::CatHandle.slurp`
- [52b34b7](https://github.com/rakudo/rakudo/commit/52b34b7) `Implement IO::CatHandle.comb/.split`
- [beaa925](https://github.com/rakudo/rakudo/commit/beaa925) `Implement IO::CatHandle.read`
- [ccc90fd](https://github.com/rakudo/rakudo/commit/ccc90fd) `Implement IO::CatHandle.readchars`
- [40f4dc9](https://github.com/rakudo/rakudo/commit/40f4dc9) `Implement IO::CatHandle.Supply`
- [0c9aea7](https://github.com/rakudo/rakudo/commit/0c9aea7) `Implement IO::CatHandle.encoding`
- [ee1e185](https://github.com/rakudo/rakudo/commit/ee1e185) `Implement IO::CatHandle.eof`
- [80686a7](https://github.com/rakudo/rakudo/commit/80686a7) `Implement IO::CatHandle.t/.path/.IO/.native-descriptor`
- [993de50](https://github.com/rakudo/rakudo/commit/993de50) `Implement IO::CatHandle.gist/.Str/.opened/.open`
- [677c4ea](https://github.com/rakudo/rakudo/commit/677c4ea) `Implement IO::CatHandle.lock/.unlock/.seek/.tell`
- [e657ed1](https://github.com/rakudo/rakudo/commit/e657ed1) `Implement IO::CatHandle.chomp/.nl-in`
- [a452e42](https://github.com/rakudo/rakudo/commit/a452e42) `Implement IO::CatHandle.on-switch`
- [f539a62](https://github.com/rakudo/rakudo/commit/f539a62) `Swap IO::ArgFiles to IO::CatHandle impl`
- [fa7aa1c](https://github.com/rakudo/rakudo/commit/fa7aa1c) `Implement IO::CatHandle.perl method`
- [21fd2c4](https://github.com/rakudo/rakudo/commit/21fd2c4) `Remove IO::Path.watch`
- [65941b2](https://github.com/rakudo/rakudo/commit/65941b2) `Revert "Remove IO::Path.watch"`
- [a47a78f](https://github.com/rakudo/rakudo/commit/a47a78f) `Remove useless :SPEC/:CWD on some IO subs`
- [d13d9c2](https://github.com/rakudo/rakudo/commit/d13d9c2) `Throw out IO::Path.int`

## Perl 6 Specification Commits

I've made 114 commits to the Perl 6 Specification (roast) repository:

- [63370fe](https://github.com/rakudo/rakudo/commit/63370fe) `Test IO::Special .WHICH/.Str do not crash`
- [465795c](https://github.com/rakudo/rakudo/commit/465795c) `Test IO::Path.lines(*) does not crash`
- [091931a](https://github.com/rakudo/rakudo/commit/091931a) `Expand &open tests`
- [8d6ca7a](https://github.com/rakudo/rakudo/commit/8d6ca7a) `Cover IO::Path.ACCEPTS`
- [14b6844](https://github.com/rakudo/rakudo/commit/14b6844) `Use Numeric instead of IO role in dispatch test`
- [5a7a365](https://github.com/rakudo/rakudo/commit/5a7a365) `Expand IO::Spec::*.tmpdir tests`
- [f48198f](https://github.com/rakudo/rakudo/commit/f48198f) `Test &indir`
- [bd46836](https://github.com/rakudo/rakudo/commit/bd46836) `Amend &indir race tests`
- [04333b3](https://github.com/rakudo/rakudo/commit/04333b3) `Test &indir fails with non-existent paths by default`
- [73a5448](https://github.com/rakudo/rakudo/commit/73a5448) `Remove two fudged &chdir tests`
- [86f79ce](https://github.com/rakudo/rakudo/commit/86f79ce) `Expand &chdir tests`
- [430ab89](https://github.com/rakudo/rakudo/commit/430ab89) `Test &*chdir`
- [86c5f9c](https://github.com/rakudo/rakudo/commit/86c5f9c) `Delete qp{} tests`
- [3c4e81b](https://github.com/rakudo/rakudo/commit/3c4e81b) `Test IO::Path.Str works as advertised`
- [ba3e7be](https://github.com/rakudo/rakudo/commit/ba3e7be) `Merge S32-io/path.t and S32-io/io-path.t`
- [79ff022](https://github.com/rakudo/rakudo/commit/79ff022) `Expand &spurt and IO::Path.spurt tests`
- [1d4e881](https://github.com/rakudo/rakudo/commit/1d4e881) `Test $*TMPDIR can be `temp`ed`
- [b23e53e](https://github.com/rakudo/rakudo/commit/b23e53e) `Test IO::Path.extension`
- [2f09f18](https://github.com/rakudo/rakudo/commit/2f09f18) `Fix incorrect test`
- [305f206](https://github.com/rakudo/rakudo/commit/305f206) `Test empty-string extensions in IO::Path.extension`
- [0e47f25](https://github.com/rakudo/rakudo/commit/0e47f25) `Test IO::Path.concat-with`
- [e5dc376](https://github.com/rakudo/rakudo/commit/e5dc376) `Expand IO::Path.accessed tests`
- [43ec543](https://github.com/rakudo/rakudo/commit/43ec543) `Cover methods of IO::Special`
- [bd8d167](https://github.com/rakudo/rakudo/commit/bd8d167) `Test IO::Path::* instantiate a subclass`
- [d8707e7](https://github.com/rakudo/rakudo/commit/d8707e7) `Cover IO::Spec::Unix.basename`
- [c3c51ed](https://github.com/rakudo/rakudo/commit/c3c51ed) `Cover IO::Spec::Win32.basename`
- [896033a](https://github.com/rakudo/rakudo/commit/896033a) `Cover IO::Spec::QNX.canonpath`
- [7c7fbb4](https://github.com/rakudo/rakudo/commit/7c7fbb4) `Cover :parent arg in IO::Spec::Cygwin.canonpath`
- [8f73ad8](https://github.com/rakudo/rakudo/commit/8f73ad8) `Change \0 roundtrip test to \t roundtrip test`
- [b16fbd3](https://github.com/rakudo/rakudo/commit/b16fbd3) `Add tests to check nul byte is rejected`
- [ee7f05b](https://github.com/rakudo/rakudo/commit/ee7f05b) `Move is-path sub to top so it can be reused`
- [a809f0f](https://github.com/rakudo/rakudo/commit/a809f0f) `Expand IO::Path.resolve tests`
- [feecaf0](https://github.com/rakudo/rakudo/commit/feecaf0) `Expand file tests`
- [a4c53b0](https://github.com/rakudo/rakudo/commit/a4c53b0) `Use bin IO::Handle to test its .Supply`
- [7e4a2ae](https://github.com/rakudo/rakudo/commit/7e4a2ae) `Swap .slurp-rest to .slurp`
- [d4353b6](https://github.com/rakudo/rakudo/commit/d4353b6) `Rewrite .l on broken symlinks test`
- [416b746](https://github.com/rakudo/rakudo/commit/416b746) `Test symlink routines`
- [8fa49e1](https://github.com/rakudo/rakudo/commit/8fa49e1) `Test `link` routines`
- [637500d](https://github.com/rakudo/rakudo/commit/637500d) `Spec IO::Pipe.path/.IO returns IO::Path type object`
- [64ff572](https://github.com/rakudo/rakudo/commit/64ff572) `Cover IO::Path/IO::Pipe's .Str/.path/.IO`
- [4194755](https://github.com/rakudo/rakudo/commit/4194755) `Test IO::Handle.lock/.unlock`
- [a716962](https://github.com/rakudo/rakudo/commit/a716962) `Amend rules for last part in IO::Path.resolve`
- [f3c5dae](https://github.com/rakudo/rakudo/commit/f3c5dae) `Test IO::Path.child-secure`
- [92217f7](https://github.com/rakudo/rakudo/commit/92217f7) `Test IO::Path.child-secure with combiners`
- [39677c4](https://github.com/rakudo/rakudo/commit/39677c4) `IO::Path.concat-with got renamed to .add`
- [7a063b5](https://github.com/rakudo/rakudo/commit/7a063b5) `Fudge .child-secure tests`
- [3b36d4d](https://github.com/rakudo/rakudo/commit/3b36d4d) `Test IO::Path.sibling`
- [41b7f9f](https://github.com/rakudo/rakudo/commit/41b7f9f) `Test $*CWD in IO::Path.dir(:test) Callable`
- [18d9c04](https://github.com/rakudo/rakudo/commit/18d9c04) `Cover IO::Handle.spurt`
- [8f78ca6](https://github.com/rakudo/rakudo/commit/8f78ca6) `Test &words with IO::ArgFiles`
- [ea137f6](https://github.com/rakudo/rakudo/commit/ea137f6) `Cover IO::Handle.tell`
- [71a6423](https://github.com/rakudo/rakudo/commit/71a6423) `Add $*HOME tests`
- [95d68a2](https://github.com/rakudo/rakudo/commit/95d68a2) `Test IO::Path.gist does escapes of backslashes`
- [de89d25](https://github.com/rakudo/rakudo/commit/de89d25) `Revert "Test IO::Path.gist does escapes of backslashes"`
- [9e8b154](https://github.com/rakudo/rakudo/commit/9e8b154) `Test IO::Handle.close can be...`
- [853f76f](https://github.com/rakudo/rakudo/commit/853f76f) `Test IO::Pipe.close returns pipe's Proc`
- [d543e75](https://github.com/rakudo/rakudo/commit/d543e75) `Test IO::Handle.DESTROY closes the handle`
- [1ed18b4](https://github.com/rakudo/rakudo/commit/1ed18b4) `Add test for .perl.EVAL roundtrip with combiners`
- [704210c](https://github.com/rakudo/rakudo/commit/704210c) `Test we can roundtrip IO::Path.perl`
- [2689eb1](https://github.com/rakudo/rakudo/commit/2689eb1) `Test .IO on :U of IO::Path subclasses`
- [40353f1](https://github.com/rakudo/rakudo/commit/40353f1) `Test for IO::Handle:D { ... } loops over handle`
- [4fdb850](https://github.com/rakudo/rakudo/commit/4fdb850) `Test IO::Path.copy/move when source/target are same`
- [98917dc](https://github.com/rakudo/rakudo/commit/98917dc) `Test IO::Path.dir's absoluteness behaviour`
- [71eebc7](https://github.com/rakudo/rakudo/commit/71eebc7) `Test IO::Spec::Unix.extension`
- [4495615](https://github.com/rakudo/rakudo/commit/4495615) `Test IO::Handle.flush`
- [60f5a6d](https://github.com/rakudo/rakudo/commit/60f5a6d) `Test IO::Handle.t when handle is a TTY`
- [31e3993](https://github.com/rakudo/rakudo/commit/31e3993) `Test IO::Path*.gist`
- [c481433](https://github.com/rakudo/rakudo/commit/c481433) `Test .is-absolute method for / with combiners`
- [8ee0a0a](https://github.com/rakudo/rakudo/commit/8ee0a0a) `Test IO::Spec::Win32.rel2abs with combiners`
- [a41027f](https://github.com/rakudo/rakudo/commit/a41027f) `Test IO::Handle.nl-in can be set`
- [e82b798](https://github.com/rakudo/rakudo/commit/e82b798) `Test IO::Handle.open respects attributes`
- [2c29150](https://github.com/rakudo/rakudo/commit/2c29150) `Test IO::Handle.nl-in attribute`
- [03ce93b](https://github.com/rakudo/rakudo/commit/03ce93b) `Test IO::Handle.encoding can be set`
- [8ae81c0](https://github.com/rakudo/rakudo/commit/8ae81c0) `Test no-arg candidate of &note`
- [fb61306](https://github.com/rakudo/rakudo/commit/fb61306) `Test IO::Path.parts attribute`
- [7266522](https://github.com/rakudo/rakudo/commit/7266522) `Test return type of IO::Spec::Unix.path`
- [6ac3b4a](https://github.com/rakudo/rakudo/commit/6ac3b4a) `Test IO::Spec::Win32.path`
- [dbbea15](https://github.com/rakudo/rakudo/commit/dbbea15) `Test IO::Handle.perl.EVAL roundtrips`
- [5eb513c](https://github.com/rakudo/rakudo/commit/5eb513c) `Test IO::Path.resolve sets CWD to $!SPEC.dir-sep`
- [b0c4a7a](https://github.com/rakudo/rakudo/commit/b0c4a7a) `Test &words, IO::Handle.words, and IO::Path.words`
- [f3d1f67](https://github.com/rakudo/rakudo/commit/f3d1f67) `Test $limit arg with &lines/IO::*.lines`
- [4f5589b](https://github.com/rakudo/rakudo/commit/4f5589b) `Add test for handle leak in IO::Path.lines`
- [4d0f97a](https://github.com/rakudo/rakudo/commit/4d0f97a) `Add &put/IO::Handle.put tests`
- [125fe18](https://github.com/rakudo/rakudo/commit/125fe18) `Add &prompt tests`
- [939ca8d](https://github.com/rakudo/rakudo/commit/939ca8d) `Test IO::CatHandle.close`
- [9833012](https://github.com/rakudo/rakudo/commit/9833012) `Test IO::CatHandle.get`
- [2f65a72](https://github.com/rakudo/rakudo/commit/2f65a72) `Test IO::CatHandle.getc`
- [a4a7eaa](https://github.com/rakudo/rakudo/commit/a4a7eaa) `Test IO::CatHandle.words`
- [1131c09](https://github.com/rakudo/rakudo/commit/1131c09) `Add &put/IO::Handle.put tests`
- [80de9b6](https://github.com/rakudo/rakudo/commit/80de9b6) `Add &prompt tests`
- [bacfd9f](https://github.com/rakudo/rakudo/commit/bacfd9f) `Test IO::CatHandle.slurp`
- [e78e3c0](https://github.com/rakudo/rakudo/commit/e78e3c0) `Test IO::CatHandle.comb/.split`
- [f1c1125](https://github.com/rakudo/rakudo/commit/f1c1125) `Test IO::CatHandle.read`
- [e9e78e1](https://github.com/rakudo/rakudo/commit/e9e78e1) `Test IO::CatHandle.readchars`
- [0479087](https://github.com/rakudo/rakudo/commit/0479087) `Test IO::CatHandle.Supply`
- [71953e3](https://github.com/rakudo/rakudo/commit/71953e3) `Test IO::CatHandle.encoding`
- [db4847e](https://github.com/rakudo/rakudo/commit/db4847e) `Test IO::CatHandle.eof`
- [175ba45](https://github.com/rakudo/rakudo/commit/175ba45) `Test IO::CatHandle.t/.path/.IO/.native-descriptor`
- [c6cc66a](https://github.com/rakudo/rakudo/commit/c6cc66a) `Test IO::CatHandle.gist/.Str/.opened/.open`
- [dcdac1a](https://github.com/rakudo/rakudo/commit/dcdac1a) `Test IO::CatHandle.lock/.unlock/.seek/.tell`
- [f48c26e](https://github.com/rakudo/rakudo/commit/f48c26e) `Test IO::CatHandle.chomp/.nl-in`
- [8afd758](https://github.com/rakudo/rakudo/commit/8afd758) `Test IO::CatHandle.DESTROY`
- [c7eff2b](https://github.com/rakudo/rakudo/commit/c7eff2b) `Test IO::CatHandle.on-switch`
- [e87e20d](https://github.com/rakudo/rakudo/commit/e87e20d) `Test IO::CatHandle.next-handle`
- [28717f0](https://github.com/rakudo/rakudo/commit/28717f0) `Test IO::CatHandle.perl method`
- [432bf94](https://github.com/rakudo/rakudo/commit/432bf94) `Test IO::Path.watch`
- [ce1b637](https://github.com/rakudo/rakudo/commit/ce1b637) `Test IO::Handle.say`
- [0bb6298](https://github.com/rakudo/rakudo/commit/0bb6298) `Test IO::Handle.print-nl`
- [47c88ab](https://github.com/rakudo/rakudo/commit/47c88ab) `Test IO::Pipe.proc attribute`
- [945621d](https://github.com/rakudo/rakudo/commit/945621d) `Test IO::Path.SPEC attribute`
- [5fb4b63](https://github.com/rakudo/rakudo/commit/5fb4b63) `Test IO::Path.CWD/.path attributes`
- [d0e5701](https://github.com/rakudo/rakudo/commit/d0e5701) `Test IO::Path.Numeric and other .numeric methods`
- [94d7133](https://github.com/rakudo/rakudo/commit/94d7133) `Test 0-arg &say/&put/&print`
- [38c61cd](https://github.com/rakudo/rakudo/commit/38c61cd) `Test &slurp() and &slurp(IO::Handle)`

## Perl 6 Documentation Commits

I've made 146 commits to the Perl 6 Documentation repository:

- [fd7a41b](https://github.com/rakudo/rakudo/commit/fd7a41b) `Improve code example`
- [110efb4](https://github.com/rakudo/rakudo/commit/110efb4) `No need for `.ends-with``
- [69d32da](https://github.com/rakudo/rakudo/commit/69d32da) `Remove IO::Handle.z`
- [d02ae7d](https://github.com/rakudo/rakudo/commit/d02ae7d) `Remove IO::Handle.rw and .rwx`
- [ccae74a](https://github.com/rakudo/rakudo/commit/ccae74a) `Fix incorrect information for IO::Path.absolute`
- [3cf943d](https://github.com/rakudo/rakudo/commit/3cf943d) `Expand IO::Path.relative`
- [cc496eb](https://github.com/rakudo/rakudo/commit/cc496eb) `Remove mention of IO.umask`
- [335a98d](https://github.com/rakudo/rakudo/commit/335a98d) `Remove mention of `role IO``
- [cc6539b](https://github.com/rakudo/rakudo/commit/cc6539b) `Remove 8 methods from IO::Handle`
- [0511e07](https://github.com/rakudo/rakudo/commit/0511e07) `Document IO::Spec::*.tmpdir`
- [db36655](https://github.com/rakudo/rakudo/commit/db36655) `Remove tip to use $*SPEC to detect OS`
- [839a6b3](https://github.com/rakudo/rakudo/commit/839a6b3) `Expand docs for $*HOME and $*TMPDIR`
- [d050d4b](https://github.com/rakudo/rakudo/commit/d050d4b) `Remove IO::Path.chdir prose`
- [1d0e433](https://github.com/rakudo/rakudo/commit/1d0e433) `Document &chdir`
- [3fdc6dc](https://github.com/rakudo/rakudo/commit/3fdc6dc) `Document &*chdir`
- [e1a299c](https://github.com/rakudo/rakudo/commit/e1a299c) `Reword "defined as" for &*chdir`
- [e5225be](https://github.com/rakudo/rakudo/commit/e5225be) `Fix URL to &*chdir`
- [bf377c7](https://github.com/rakudo/rakudo/commit/bf377c7) `Document &indir`
- [5aa614f](https://github.com/rakudo/rakudo/commit/5aa614f) `Improve suggestion for Perl 5's opendir`
- [a53015a](https://github.com/rakudo/rakudo/commit/a53015a) `Clarify value of IO::Path.path`
- [bdd18f1](https://github.com/rakudo/rakudo/commit/bdd18f1) `Fix desc of IO::Path.Str`
- [b78d4fd](https://github.com/rakudo/rakudo/commit/b78d4fd) `Include type names in links to methods`
- [b8fba97](https://github.com/rakudo/rakudo/commit/b8fba97) `Point out my $*CWD = chdir ??? is an error`
- [d5abceb](https://github.com/rakudo/rakudo/commit/d5abceb) `Write docs for all spurt routines`
- [b9e692e](https://github.com/rakudo/rakudo/commit/b9e692e) `Document new IO::Path.extension`
- [65cc372](https://github.com/rakudo/rakudo/commit/65cc372) `Document IO::Path.concat-with`
- [24a6ea9](https://github.com/rakudo/rakudo/commit/24a6ea9) `Toss all of the TODO methods in IO::Spec*`
- [1f75ddc](https://github.com/rakudo/rakudo/commit/1f75ddc) `Document IO::Spec*.abs2rel`
- [cc62dd2](https://github.com/rakudo/rakudo/commit/cc62dd2) `Kill IO::Path.abspath`
- [1973010](https://github.com/rakudo/rakudo/commit/1973010) `Document IO::Path.ACCEPTS`
- [b3a9324](https://github.com/rakudo/rakudo/commit/b3a9324) `Expand/fix up IO::Path.accessed`
- [1cd7de0](https://github.com/rakudo/rakudo/commit/1cd7de0) `Fix up type graph`
- [56256d0](https://github.com/rakudo/rakudo/commit/56256d0) `Minor formatting improvements in IO::Special`
- [184342c](https://github.com/rakudo/rakudo/commit/184342c) `Document IO::Special.what`
- [6bd0f98](https://github.com/rakudo/rakudo/commit/6bd0f98) `Dissuade readers from using IO::Spec*`
- [7afd9c4](https://github.com/rakudo/rakudo/commit/7afd9c4) `Remove unrelated related classes`
- [a43ecb9](https://github.com/rakudo/rakudo/commit/a43ecb9) `Document IO::Path's $.SPEC and $.CWD`
- [e9b6809](https://github.com/rakudo/rakudo/commit/e9b6809) `Document IO::Path::* subclasses`
- [9102b51](https://github.com/rakudo/rakudo/commit/9102b51) `Fix up IO::Path.basename`
- [5c1d3b6](https://github.com/rakudo/rakudo/commit/5c1d3b6) `Document IO::Spec::Unix.basename`
- [a1cb80b](https://github.com/rakudo/rakudo/commit/a1cb80b) `Document IO::Spec::Win32.basename`
- [28b6283](https://github.com/rakudo/rakudo/commit/28b6283) `Document IO::Spec::*.canonpath`
- [50e5565](https://github.com/rakudo/rakudo/commit/50e5565) `Document IO::Spec::*.catdir and .catfile`
- [dbdc995](https://github.com/rakudo/rakudo/commit/dbdc995) `Document IO::Spec::*.catpath`
- [0ca2295](https://github.com/rakudo/rakudo/commit/0ca2295) `Reword/expand IO::Path intro prose`
- [45e84ad](https://github.com/rakudo/rakudo/commit/45e84ad) `Move IO::Path.path to attributes`
- [b9de84f](https://github.com/rakudo/rakudo/commit/b9de84f) `Remove DateTime tutorial from IO::Path docs`
- [69b2082](https://github.com/rakudo/rakudo/commit/69b2082) `Document IO::Path.chdir`
- [d436f3c](https://github.com/rakudo/rakudo/commit/d436f3c) `Document IO::Spec::* don't do any validation`
- [4090446](https://github.com/rakudo/rakudo/commit/4090446) `Improve chmod docs`
- [1527d32](https://github.com/rakudo/rakudo/commit/1527d32) `Document :completely arg to IO::Path.resolve`
- [372545c](https://github.com/rakudo/rakudo/commit/372545c) `Straighten up file test docs`
- [a30fae6](https://github.com/rakudo/rakudo/commit/a30fae6) `Avoid potential confusion with use of word "object"`
- [2aa3c9f](https://github.com/rakudo/rakudo/commit/2aa3c9f) `Document new behaviour of IO::Handle.Supply`
- [56b50fe](https://github.com/rakudo/rakudo/commit/56b50fe) `Document IO::Handle.slurp`
- [017acd4](https://github.com/rakudo/rakudo/commit/017acd4) `Improve docs for IO::Path.slurp`
- [0f49bb5](https://github.com/rakudo/rakudo/commit/0f49bb5) `List Rakudo-supported encodings in open()`
- [e60da5c](https://github.com/rakudo/rakudo/commit/e60da5c) `List utf-* alias examples too since they're common`
- [f83f78c](https://github.com/rakudo/rakudo/commit/f83f78c) `Use idiomatic Perl 6 in example`
- [fff866f](https://github.com/rakudo/rakudo/commit/fff866f) `Fix docs for symlink/link routines`
- [aeeec94](https://github.com/rakudo/rakudo/commit/aeeec94) `Straighten up copy, move, rename`
- [923ea05](https://github.com/rakudo/rakudo/commit/923ea05) `Straighten up mkdir docs`
- [47b0526](https://github.com/rakudo/rakudo/commit/47b0526) `Explicitly spell out caveats of IO::Path.Str`
- [60b9227](https://github.com/rakudo/rakudo/commit/60b9227) `Change return value for `mkdir``
- [8d95371](https://github.com/rakudo/rakudo/commit/8d95371) `Expand IO::Handle/IO::Pipe.path docs`
- [fd8a5ed](https://github.com/rakudo/rakudo/commit/fd8a5ed) `Document IO::Pipe.path`
- [bd4fa68](https://github.com/rakudo/rakudo/commit/bd4fa68) `Document IO::Handle/IO::Pipe.IO`
- [2aaf12a](https://github.com/rakudo/rakudo/commit/2aaf12a) `Document IO::Handle.Str`
- [53f2b99](https://github.com/rakudo/rakudo/commit/53f2b99) `Document `role IO`'s new purpose`
- [160c6a2](https://github.com/rakudo/rakudo/commit/160c6a2) `Document IO::Handle.lock/.unlock`
- [3145979](https://github.com/rakudo/rakudo/commit/3145979) `Document IO::Path.child-secure`
- [c5524ef](https://github.com/rakudo/rakudo/commit/c5524ef) `Rename IO::Path.concat-with to .add`
- [81a5806](https://github.com/rakudo/rakudo/commit/81a5806) `Amend IO::Path.resolve: :completely`
- [6ca67e4](https://github.com/rakudo/rakudo/commit/6ca67e4) `Start sketching out Definitive IO Guide???`
- [b9c9117](https://github.com/rakudo/rakudo/commit/b9c9117) `Toss IO::Path.child-secure`
- [61cb776](https://github.com/rakudo/rakudo/commit/61cb776) `Document IO::Path.sibling`
- [0fc39a6](https://github.com/rakudo/rakudo/commit/0fc39a6) `Fix typegraph`
- [9a63dc4](https://github.com/rakudo/rakudo/commit/9a63dc4) `Document IO::Path.cleanup`
- [2387ce3](https://github.com/rakudo/rakudo/commit/2387ce3) `Re-write IO::Handle.close docs`
- [0def0d1](https://github.com/rakudo/rakudo/commit/0def0d1) `Amend IO::Handle.close docs`
- [c7e32e2](https://github.com/rakudo/rakudo/commit/c7e32e2) `Document IO::Spec::Unix.curupdir`
- [fe489dc](https://github.com/rakudo/rakudo/commit/fe489dc) `Document IO::Spec::Unix.curdir`
- [83d5de0](https://github.com/rakudo/rakudo/commit/83d5de0) `Document IO::Spec::Unix.updir`
- [4804128](https://github.com/rakudo/rakudo/commit/4804128) `Document IO::Handle.DESTROY`
- [c991862](https://github.com/rakudo/rakudo/commit/c991862) `Add warning to dir about...`
- [eca21ff](https://github.com/rakudo/rakudo/commit/eca21ff) `Document copy/move behaviour for same target/source`
- [6c2b8b2](https://github.com/rakudo/rakudo/commit/6c2b8b2) `Document IO::Path/IO::Handle.comb`
- [fb29e04](https://github.com/rakudo/rakudo/commit/fb29e04) `Include exception used in IO::Path.resolve`
- [69d473f](https://github.com/rakudo/rakudo/commit/69d473f) `Document IO::Spec::*.devnull`
- [994d671](https://github.com/rakudo/rakudo/commit/994d671) `List IO::Dir as one of the means...`
- [4432ef3](https://github.com/rakudo/rakudo/commit/4432ef3) `Finish up IO::Path.dir docs`
- [64355c8](https://github.com/rakudo/rakudo/commit/64355c8) `Document IO::Spec::*.dir-sep`
- [914c100](https://github.com/rakudo/rakudo/commit/914c100) `Finish up IO::Path.dirname`
- [8d5e31c](https://github.com/rakudo/rakudo/commit/8d5e31c) `Document IO::Handle.encoding`
- [d5c36aa](https://github.com/rakudo/rakudo/commit/d5c36aa) `Finish off IO::Handle.eof`
- [e9de97e](https://github.com/rakudo/rakudo/commit/e9de97e) `Document IO::Spec::*.extension`
- [bf7ec00](https://github.com/rakudo/rakudo/commit/bf7ec00) `Document IO::Handle.flush`
- [25bce38](https://github.com/rakudo/rakudo/commit/25bce38) `Document IO::Path.succ`
- [8233960](https://github.com/rakudo/rakudo/commit/8233960) `Improve IO::Handle.t docs`
- [b4006a2](https://github.com/rakudo/rakudo/commit/b4006a2) `Be explicit what IO::Handle.opened returns`
- [c4f27a7](https://github.com/rakudo/rakudo/commit/c4f27a7) `Document IO::Path.pred`
- [860333f](https://github.com/rakudo/rakudo/commit/860333f) `Remove entirely-invented "File test operators"`
- [ab0bd7a](https://github.com/rakudo/rakudo/commit/ab0bd7a) `Document IO::Path.Numeric/.Int`
- [4f81f08](https://github.com/rakudo/rakudo/commit/4f81f08) `Improve IO::Handle.get docs`
- [c45d389](https://github.com/rakudo/rakudo/commit/c45d389) `Finish off IO::Handle.getc/&getc docs`
- [a4012e0](https://github.com/rakudo/rakudo/commit/a4012e0) `Document IO::Handle.gist`
- [d15b0c7](https://github.com/rakudo/rakudo/commit/d15b0c7) `Document IO::Path.gist`
- [1cf6932](https://github.com/rakudo/rakudo/commit/1cf6932) `Document IO::Spec::*.is-absolute`
- [4e88b84](https://github.com/rakudo/rakudo/commit/4e88b84) `Finish up IO::Path.is-absolute`
- [497e7f7](https://github.com/rakudo/rakudo/commit/497e7f7) `Finish off IO::Path.is-relative`
- [f7e75c1](https://github.com/rakudo/rakudo/commit/f7e75c1) `Document IO::Handle.nl-in`
- [e309ddd](https://github.com/rakudo/rakudo/commit/e309ddd) `Finish up &note`
- [81900cb](https://github.com/rakudo/rakudo/commit/81900cb) `Finish off IO::Path.parent`
- [59cbc38](https://github.com/rakudo/rakudo/commit/59cbc38) `Finish off IO::Path.parts`
- [b99a666](https://github.com/rakudo/rakudo/commit/b99a666) `Finish off IO::Path.path/.IO`
- [b070999](https://github.com/rakudo/rakudo/commit/b070999) `Document IO::Spec::*.path`
- [bace8ff](https://github.com/rakudo/rakudo/commit/bace8ff) `Document IO::Path*.perl`
- [dfdd845](https://github.com/rakudo/rakudo/commit/dfdd845) `Add "The Basics" section to TDIOG`
- [cdc701e](https://github.com/rakudo/rakudo/commit/cdc701e) `Add "What's an IO::Path Anyway?" section to TDIOG`
- [0d6d058](https://github.com/rakudo/rakudo/commit/0d6d058) `Add "Writing into files" Section to TDIOG`
- [a6365f3](https://github.com/rakudo/rakudo/commit/a6365f3) `Document IO::Handle.words/&words`
- [2e25c82](https://github.com/rakudo/rakudo/commit/2e25c82) `Document IO::Spec::*.join`
- [49e58bd](https://github.com/rakudo/rakudo/commit/49e58bd) `Document IO::Handle.lines`
- [1744820](https://github.com/rakudo/rakudo/commit/1744820) `Document IO::Path.lines`
- [f3f70a0](https://github.com/rakudo/rakudo/commit/f3f70a0) `Document IO::Path.words`
- [509f0e8](https://github.com/rakudo/rakudo/commit/509f0e8) `Fix incorrect suggested routine`
- [a6f1cbf](https://github.com/rakudo/rakudo/commit/a6f1cbf) `Fix up IO::Handle.print`
- [8f53830](https://github.com/rakudo/rakudo/commit/8f53830) `Fix up IO::Handle.print-nl`
- [dc50211](https://github.com/rakudo/rakudo/commit/dc50211) `Fix &prompt`
- [98965b3](https://github.com/rakudo/rakudo/commit/98965b3) `Fix up IO::Handle.split`
- [bd702e2](https://github.com/rakudo/rakudo/commit/bd702e2) `Fix up IO::Handle.comb`
- [6dd92b8](https://github.com/rakudo/rakudo/commit/6dd92b8) `Document IO::CatHandle`
- [edeb069](https://github.com/rakudo/rakudo/commit/edeb069) `Document IO::Path.split`
- [2d96596](https://github.com/rakudo/rakudo/commit/2d96596) `Document IO::Spec::*.split`
- [129c097](https://github.com/rakudo/rakudo/commit/129c097) `Document IO::Spec::*.splitdir`
- [b946960](https://github.com/rakudo/rakudo/commit/b946960) `Document IO::Spec::*.splitpath`
- [dcd7490](https://github.com/rakudo/rakudo/commit/dcd7490) `Fix rmdir docs`
- [2a7bd17](https://github.com/rakudo/rakudo/commit/2a7bd17) `Document IO::Spec::*.rel2abs`
- [f45241f](https://github.com/rakudo/rakudo/commit/f45241f) `Document IO::Spec::*.rootdir`
- [70a80ec](https://github.com/rakudo/rakudo/commit/70a80ec) `Document IO::Handle.put`
- [6f58ed0](https://github.com/rakudo/rakudo/commit/6f58ed0) `Polish IO::Handle.say`
- [3790a0f](https://github.com/rakudo/rakudo/commit/3790a0f) `Polish &put/&print/&say`
- [ebb6f53](https://github.com/rakudo/rakudo/commit/ebb6f53) `Document IO::Handle.nl-out attribute`
- [53c9c91](https://github.com/rakudo/rakudo/commit/53c9c91) `Document IO::Handle.chomp attribute`
- [ca2a3a0](https://github.com/rakudo/rakudo/commit/ca2a3a0) `Improve &open/IO::Handle.open docs`
- [856e846](https://github.com/rakudo/rakudo/commit/856e846) `Add Reading From Files section to TDIOG`


