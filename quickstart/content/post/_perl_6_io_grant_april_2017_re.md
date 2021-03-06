
---
title: "Perl 6 IO Grant: April 2017 Report"
author: Coke
type: post
date: 2017-04-26 00:00:00 +0000 UTC
url: "/post/_perl_6_io_grant_april_2017_re"
categories:
 - Grants

---


*Zoffix Znet provided this report on April 19, 2017*

# Perl 6 IO TPF Grant: Monthly Report (April, 2017)

This document is the April, 2017 progress report for [TPF Standardization,
Test Coverage, and Documentation of Perl 6 I/O Routines
grant](http://news.perlfoundation.org/2017/01/grant-proposal-standardization.html)

## Timing

As proposed to and approved by the Grant Manager, I've extended the due date
for this grant by 1 extra month, in exchange for doing some extra optimization
work on IO routines at no extra cost. The new completion date is May 22nd;
right after the next Rakudo compiler release.

## Communications

I've created and published three notices as part of this grant, informing the
users on what is changing and how to best upgrade their code, where needed:

- [Upgrade Information for Changes Due to IO Grant Work](http://rakudo.org/2017/04/02/upgrade-information-for-changes-due-to-io-grant-work/)
- [PART 2: Upgrade Information for Changes Due to IO Grant Work](http://rakudo.org/2017/04/03/part-2-upgrade-information-for-changes-due-to-io-grant-work/)
- [PART 3: Information on Changes Due to IO Grant Work](http://rakudo.org/2017/04/17/final-notes-on-changes-due-to-io-grant-work/)


## IO Action Plan Progress

Most of the [IO Action Plan](https://github.com/rakudo/rakudo/blob/nom/docs/2017-IO-Grant--Action-Plan.md) has been implemented and got shipped in Rakudo's 2017.04.2 release. The remaining items are:

- Implement better way to signal closed handle status (was omited from release due to original suggestion to do this not being ideal; likely better to do this on the VM end)
- Implement IO::CatHandle as a generalized IO::ArgFiles (was omited from release because it was decided to make it mostly-usable wherever IO::Handle can be used, and IO::ArgFiles is far from that, having only a handful of methods implemented)
- Optimization of the way we perform `stat` calls for multiple file tests (entirely internal that requires no changes to users' code)


## Documentation and Coverage Progress

In my spreadsheet with all the IO routines and their candidates, the totals
show that 40% have been documented and tested. Some of the remaining 60% may
already have tests/docs added when implementing IO Action Plan or ealier and
merely need checking and verification.

## Optimizations

Some of the optimizations I promised to deliver in exchange for grant deadline
extension were already done on IO::Spec::Unix and IO::Path
routines and have made it into the 2017.04.2 release. Most of the optimizations
that will be done in the upcoming month will be done in IO::Spec::Win32 and
will largely affect Windows users.

#### IO Optimizations in 2017.04.2 Done by Other Core Members:

- Elizabeth Mattijsen made .slurp 2x faster [rakudo/b4d80c0](https://github.com/rakudo/rakudo/commit/b4d80c0)
- Samantha McVey made nqp::index???which is used in path operations???2x faster [rakudo/f1fc879](https://github.com/rakudo/rakudo/commit/f1fc879)
- IO::Pipe.lines was made 3.2x faster by relying on work done by Elizabeth Mattijsen [rakudo/0c62815](https://github.com/rakudo/rakudo/commit/0c62815)

## Tickets Resolved

The following tickets have been resolved as part of the grant:

- [RT#130460 Can we relax indir's test on the target directory?](https://rt.perl.org/Ticket/Display.html?id=130460): Resolved by making indir's default test to be only `:d`
- [RT#130900: nul in pathname](https://rt.perl.org/Ticket/Display.html?id=130900): Resolved by checking the path for nul and throwing when found
- [RT#127407: [RFC] (1) add method IO::Path.stemname; (2) expand method IO::Path.parts](https://rt.perl.org/Ticket/Update.html?id=127407): RFC rejected, but the described functionality is now available via the .extension method that was made much more powerful as part of IO Action Plan

Possibly more tickets were addressed by the IO Action Plan implementation, but
they still need further review.

## Bugs Fixed

- Fixed a bug in `IO::Path.resolve` with combiners tucked on the path
    separator. Fix in
[rakudo/9d8e391f3b](https://github.com/rakudo/rakudo/commit/9d8e391f3b452efeb06c860e37293c696b244786);
tests in
[roast/92217f75ce](https://github.com/perl6/roast/commit/92217f75ce9392466eb14fc562911fc21a83a02e). The bug was identified by Timo Paulssen while testing secure implementation of IO::Path.child

#### IO Bug Fixes in 2017.04.2 Done by Other Core Members:

- Timo Paulssen fixed a bug with IO::Path types not being accepted by
`is native` NativeCall trait [rakudo/99840804](https://github.com/rakudo/rakudo/commit/9984080413c9c236f9e8a35fded928a65cb9810a)
- Elizabeth Mattijsen fixed an issue in assignment to dynamics. This made it possible to `temp` `$*TMPDIR` variable [rakudo/1b9d53](https://github.com/rakudo/rakudo/commit/1b9d53ce5ed1e9e1c546f39fbec4fe342667ed1c)
- Jonathan Worthington fixed a crash when slurping large files in binary mode with &slurp or IO::Path.slurp [rakudo/d0924f1a2](https://github.com/rakudo/rakudo/commit/d0924f1a287bd2966c1dec156d135f2447ca21da)
- Jonathan Worthington fixed a bug with binary slurp reading zero bytes when another thread is causing a lot of GC [rakudo/756877e](https://github.com/rakudo/rakudo/commit/756877e)

## Commits

So far, I've commited 192 IO grant commits to rakudo/roast/doc repos.

### Rakudo

69 IO grant commits:

- [`c6fd736`](https://github.com/rakudo/rakudo/commit/c6fd736) `Make IO::Spec::Win32.is-absolute about 63x faster`
- [`7112a08`](https://github.com/rakudo/rakudo/commit/7112a08) `Add :D on invocant for file tests`
- [`8bacad8`](https://github.com/rakudo/rakudo/commit/8bacad8) `Implement IO::Path.sibling`
- [`a98b285`](https://github.com/rakudo/rakudo/commit/a98b285) `Remove IO::Path.child-secure`
- [`0b5a41b`](https://github.com/rakudo/rakudo/commit/0b5a41b) `Rename IO::Path.concat-with to .add`
- [`9d8e391`](https://github.com/rakudo/rakudo/commit/9d8e391) `Fix IO::Path.resolve with combiners; timotimo++`
- [`1887114`](https://github.com/rakudo/rakudo/commit/1887114) `Implement IO::Path.child-secure`
- [`b8458d3`](https://github.com/rakudo/rakudo/commit/b8458d3) `Reword method child for cleaner code`
- [`51e4629`](https://github.com/rakudo/rakudo/commit/51e4629) `Amend rules for last part in IO::Path.resolve`
- [`9a2446c`](https://github.com/rakudo/rakudo/commit/9a2446c) `Move Bool return value to signature`
- [`214198b`](https://github.com/rakudo/rakudo/commit/214198b) `Implement proper args for IO::Handle.lock`
- [`c95c4a7`](https://github.com/rakudo/rakudo/commit/c95c4a7) `Make IO::Path/IO::Special do IO role`
- [`fd503f8`](https://github.com/rakudo/rakudo/commit/fd503f8) `grant] Remove role IO and its .umask method"`
- [`0e36bb2`](https://github.com/rakudo/rakudo/commit/0e36bb2) `Make IO::Spec::Win32!canon-cat 2.3x faster`
- [`40217ed`](https://github.com/rakudo/rakudo/commit/40217ed) `Swap .child to .concat-with in all the guts`
- [`490ffd1`](https://github.com/rakudo/rakudo/commit/490ffd1) `Do not use self.Str in IO::Path errors`
- [`1f689a9`](https://github.com/rakudo/rakudo/commit/1f689a9) `Fix up IO::Handle.Str`
- [`c01ebea`](https://github.com/rakudo/rakudo/commit/c01ebea) `Make IO::Path.mkdir return invocant on success`
- [`d46e8df`](https://github.com/rakudo/rakudo/commit/d46e8df) `Add IO::Pipe .path and .IO methods`
- [`6ee71c2`](https://github.com/rakudo/rakudo/commit/6ee71c2) `Coerce mode in IO::Path.mkdir to Int`
- [`0d9ecae`](https://github.com/rakudo/rakudo/commit/0d9ecae) `Remove multi-dir &mkdir`
- [`ff97083`](https://github.com/rakudo/rakudo/commit/ff97083) `Straighten up rename, move, and copy`
- [`7f73f92`](https://github.com/rakudo/rakudo/commit/7f73f92) `Make IO::Path.new-from-absolute-path private`
- [`da1dea2`](https://github.com/rakudo/rakudo/commit/da1dea2) `Fix &symlink and &link`
- [`8c09c84`](https://github.com/rakudo/rakudo/commit/8c09c84) `Fix symlink and link routines`
- [`90da80f`](https://github.com/rakudo/rakudo/commit/90da80f) `Rework read methods in IO::Path/IO::Handle`
- [`c13480c`](https://github.com/rakudo/rakudo/commit/c13480c) `IO::Path.slurp: make 12%-35% faster; propagate Failures`
- [`f1b4af7`](https://github.com/rakudo/rakudo/commit/f1b4af7) `Implement IO::Handle.slurp`
- [`184d499`](https://github.com/rakudo/rakudo/commit/184d499) `Make IO::Handle.Supply respect handle's mode`
- [`b6838ee`](https://github.com/rakudo/rakudo/commit/b6838ee) `Remove .f check in .z`
- [`6a8d63d`](https://github.com/rakudo/rakudo/commit/6a8d63d) `Implement :completely param in IO::Path.resolve`
- [`e681498`](https://github.com/rakudo/rakudo/commit/e681498) `Make IO::Path throw when path contains NUL byte`
- [`b4358af`](https://github.com/rakudo/rakudo/commit/b4358af) `Delete code for IO::Spec::Win32.catfile`
- [`0a442ce`](https://github.com/rakudo/rakudo/commit/0a442ce) `Remove type constraint in IO::Spec::Cygwin.canonpath`
- [`0c8bef5`](https://github.com/rakudo/rakudo/commit/0c8bef5) `Implement :parent in IO::Spec::Cygwin.canonpath`
- [`a0b82ed`](https://github.com/rakudo/rakudo/commit/a0b82ed) `Make IO::Path::* actually instantiate a subclass`
- [`67f06b2`](https://github.com/rakudo/rakudo/commit/67f06b2) `Run S32-io/io-special.t test file`
- [`954e69e`](https://github.com/rakudo/rakudo/commit/954e69e) `Fix return value of IO::Special methods`
- [`a432b3d`](https://github.com/rakudo/rakudo/commit/a432b3d) `Remove IO::Path.abspath (part 2)`
- [`94a6909`](https://github.com/rakudo/rakudo/commit/94a6909) `Clean up IO::Spec::Unix.abs2rel a bit`
- [`966a7e3`](https://github.com/rakudo/rakudo/commit/966a7e3) `Implement IO::Path.concat-with`
- [`50aea2b`](https://github.com/rakudo/rakudo/commit/50aea2b) `Restore IO::Handle.IO`
- [`15a25da`](https://github.com/rakudo/rakudo/commit/15a25da) `Fix ambiguity in empty extension vs no extension`
- [`b1e7a01`](https://github.com/rakudo/rakudo/commit/b1e7a01) `Implement IO::Path.extension 2.0`
- [`b62d1a7`](https://github.com/rakudo/rakudo/commit/b62d1a7) `Give $*TMPDIR a container`
- [`099512b`](https://github.com/rakudo/rakudo/commit/099512b) `Clean up and improve all spurt routines`
- [`cb27bce`](https://github.com/rakudo/rakudo/commit/cb27bce) `Clean up &open and IO::Path.open`
- [`4c31903`](https://github.com/rakudo/rakudo/commit/4c31903) `Add S32-io/chdir-process.t to list of test files to run`
- [`5464b82`](https://github.com/rakudo/rakudo/commit/5464b82) `Improve &*chdir`
- [`2483d68`](https://github.com/rakudo/rakudo/commit/2483d68) `Fix regression in &chdir's failure mode`
- [`ca1acb7`](https://github.com/rakudo/rakudo/commit/ca1acb7) `Fix race in &indir(IO::Path ???)`
- [`a0ef2ed`](https://github.com/rakudo/rakudo/commit/a0ef2ed) `Improve &chdir, &indir, and IO::Path.chdir`
- [`aa62cd5`](https://github.com/rakudo/rakudo/commit/aa62cd5) `Remove &tmpdir and &homedir`
- [`a5800a1`](https://github.com/rakudo/rakudo/commit/a5800a1) `Implement IO::Handle.spurt`
- [`36ad92a`](https://github.com/rakudo/rakudo/commit/36ad92a) `Remove 15 methods from IO::Handle`
- [`87987c2`](https://github.com/rakudo/rakudo/commit/87987c2) `Remove role IO and its .umask method`
- [`9d8d7b2`](https://github.com/rakudo/rakudo/commit/9d8d7b2) `Log all changes to plan made during review period`
- [`0c7e4a0`](https://github.com/rakudo/rakudo/commit/0c7e4a0) `Do not capture args in .IO method`
- [`c360ac2`](https://github.com/rakudo/rakudo/commit/c360ac2) `Fix smartmatch of Cool ~~ IO::Path`
- [`fa9aa47`](https://github.com/rakudo/rakudo/commit/fa9aa47) `Make R::I::SET_LINE_ENDING_ON_HANDLE 4.1x Faster`
- [`0111f10`](https://github.com/rakudo/rakudo/commit/0111f10) `Make IO::Spec::Unix.catdir 3.9x Faster`
- [`4fdebc9`](https://github.com/rakudo/rakudo/commit/4fdebc9) `Make IO::Spec::Unix.split 36x Faster`
- [`dcf1bb2`](https://github.com/rakudo/rakudo/commit/dcf1bb2) `Make IO::Spec::Unix.rel2abs 35% faster`
- [`55abc6d`](https://github.com/rakudo/rakudo/commit/55abc6d) `Improve IO::Path.child perf on *nix`
- [`4032953`](https://github.com/rakudo/rakudo/commit/4032953) `Make IO::Handle.open 75% faster`
- [`a01d679`](https://github.com/rakudo/rakudo/commit/a01d679) `Remove IO::Path.pipe`
- [`212cc8a`](https://github.com/rakudo/rakudo/commit/212cc8a) `Remove IO::Path.Bridge`
- [`76f7187`](https://github.com/rakudo/rakudo/commit/76f7187) `Do not cache IO::Path.e results`
- [`dd4dfb1`](https://github.com/rakudo/rakudo/commit/dd4dfb1) `Fix crash in IO::Special .WHICH/.Str`

### Perl 6 Specification

47 IO grant commits:

- [`3b36d4d`](https://github.com/perl6/roast/commit/3b36d4d) `Test IO::Path.sibling`
- [`7a063b5`](https://github.com/perl6/roast/commit/7a063b5) `Fudge .child-secure tests`
- [`39677c4`](https://github.com/perl6/roast/commit/39677c4) `IO::Path.concat-with got renamed to .add`
- [`92217f7`](https://github.com/perl6/roast/commit/92217f7) `Test IO::Path.child-secure with combiners`
- [`f3c5dae`](https://github.com/perl6/roast/commit/f3c5dae) `Test IO::Path.child-secure`
- [`a716962`](https://github.com/perl6/roast/commit/a716962) `Amend rules for last part in IO::Path.resolve`
- [`4194755`](https://github.com/perl6/roast/commit/4194755) `Test IO::Handle.lock/.unlock`
- [`64ff572`](https://github.com/perl6/roast/commit/64ff572) `Cover IO::Path/IO::Pipe's .Str/.path/.IO`
- [`637500d`](https://github.com/perl6/roast/commit/637500d) `Spec IO::Pipe.path/.IO returns IO::Path type object`
- [`8fa49e1`](https://github.com/perl6/roast/commit/8fa49e1) `Test link routines`
- [`416b746`](https://github.com/perl6/roast/commit/416b746) `Test symlink routines`
- [`d4353b6`](https://github.com/perl6/roast/commit/d4353b6) `Rewrite .l on broken symlinks test`
- [`7e4a2ae`](https://github.com/perl6/roast/commit/7e4a2ae) `Swap .slurp-rest to .slurp`
- [`a4c53b0`](https://github.com/perl6/roast/commit/a4c53b0) `Use bin IO::Handle to test its .Supply`
- [`feecaf0`](https://github.com/perl6/roast/commit/feecaf0) `Expand file tests`
- [`a809f0f`](https://github.com/perl6/roast/commit/a809f0f) `Expand IO::Path.resolve tests`
- [`ee7f05b`](https://github.com/perl6/roast/commit/ee7f05b) `Move is-path sub to top so it can be reused`
- [`b16fbd3`](https://github.com/perl6/roast/commit/b16fbd3) `Add tests to check nul byte is rejected`
- [`8f73ad8`](https://github.com/perl6/roast/commit/8f73ad8) `Change \0 roundtrip test to \t roundtrip test`
- [`7c7fbb4`](https://github.com/perl6/roast/commit/7c7fbb4) `Cover :parent arg in IO::Spec::Cygwin.canonpath`
- [`896033a`](https://github.com/perl6/roast/commit/896033a) `Cover IO::Spec::QNX.canonpath`
- [`c3c51ed`](https://github.com/perl6/roast/commit/c3c51ed) `Cover IO::Spec::Win32.basename`
- [`d8707e7`](https://github.com/perl6/roast/commit/d8707e7) `Cover IO::Spec::Unix.basename`
- [`bd8d167`](https://github.com/perl6/roast/commit/bd8d167) `Test IO::Path::* instantiate a subclass`
- [`43ec543`](https://github.com/perl6/roast/commit/43ec543) `Cover methods of IO::Special`
- [`e5dc376`](https://github.com/perl6/roast/commit/e5dc376) `Expand IO::Path.accessed tests`
- [`0e47f25`](https://github.com/perl6/roast/commit/0e47f25) `Test IO::Path.concat-with`
- [`305f206`](https://github.com/perl6/roast/commit/305f206) `Test empty-string extensions in IO::Path.extension`
- [`2f09f18`](https://github.com/perl6/roast/commit/2f09f18) `Fix incorrect test`
- [`b23e53e`](https://github.com/perl6/roast/commit/b23e53e) `Test IO::Path.extension`
- [`1d4e881`](https://github.com/perl6/roast/commit/1d4e881) `Test $*TMPDIR can be temped`
- [`79ff022`](https://github.com/perl6/roast/commit/79ff022) `Expand &spurt and IO::Path.spurt tests`
- [`ba3e7be`](https://github.com/perl6/roast/commit/ba3e7be) `Merge S32-io/path.t and S32-io/io-path.t`
- [`3c4e81b`](https://github.com/perl6/roast/commit/3c4e81b) `Test IO::Path.Str works as advertised`
- [`86c5f9c`](https://github.com/perl6/roast/commit/86c5f9c) `Delete qp{} tests`
- [`430ab89`](https://github.com/perl6/roast/commit/430ab89) `Test &*chdir`
- [`86f79ce`](https://github.com/perl6/roast/commit/86f79ce) `Expand &chdir tests`
- [`73a5448`](https://github.com/perl6/roast/commit/73a5448) `Remove two fudged &chdir tests`
- [`04333b3`](https://github.com/perl6/roast/commit/04333b3) `Test &indir fails with non-existent paths by default`
- [`bd46836`](https://github.com/perl6/roast/commit/bd46836) `Amend &indir race tests`
- [`f48198f`](https://github.com/perl6/roast/commit/f48198f) `Test &indir`
- [`5a7a365`](https://github.com/perl6/roast/commit/5a7a365) `Expand IO::Spec::*.tmpdir tests`
- [`14b6844`](https://github.com/perl6/roast/commit/14b6844) `Use Numeric instead of IO role in dispatch test`
- [`8d6ca7a`](https://github.com/perl6/roast/commit/8d6ca7a) `Cover IO::Path.ACCEPTS`
- [`091931a`](https://github.com/perl6/roast/commit/091931a) `Expand &open tests`
- [`465795c`](https://github.com/perl6/roast/commit/465795c) `Test IO::Path.lines(*) does not crash`
- [`63370fe`](https://github.com/perl6/roast/commit/63370fe) `Test IO::Special .WHICH/.Str do not crash`

### Documentation

76 IO grant commits:

- [`61cb776`](https://github.com/perl6/doc/commit/61cb776) `Document IO::Path.sibling`
- [`b9c9117`](https://github.com/perl6/doc/commit/b9c9117) `Toss IO::Path.child-secure`
- [`6ca67e4`](https://github.com/perl6/doc/commit/6ca67e4) `Start sketching out Definitive IO Guide???`
- [`81a5806`](https://github.com/perl6/doc/commit/81a5806) `Amend IO::Path.resolve: :completely`
- [`c5524ef`](https://github.com/perl6/doc/commit/c5524ef) `Rename IO::Path.concat-with to .add`
- [`3145979`](https://github.com/perl6/doc/commit/3145979) `Document IO::Path.child-secure`
- [`160c6a2`](https://github.com/perl6/doc/commit/160c6a2) `Document IO::Handle.lock/.unlock`
- [`53f2b99`](https://github.com/perl6/doc/commit/53f2b99) `Document role IO's new purpose`
- [`2aaf12a`](https://github.com/perl6/doc/commit/2aaf12a) `Document IO::Handle.Str`
- [`bd4fa68`](https://github.com/perl6/doc/commit/bd4fa68) `Document IO::Handle/IO::Pipe.IO`
- [`fd8a5ed`](https://github.com/perl6/doc/commit/fd8a5ed) `Document IO::Pipe.path`
- [`8d95371`](https://github.com/perl6/doc/commit/8d95371) `Expand IO::Handle/IO::Pipe.path docs`
- [`60b9227`](https://github.com/perl6/doc/commit/60b9227) `Change return value for mkdir`
- [`47b0526`](https://github.com/perl6/doc/commit/47b0526) `Explicitly spell out caveats of IO::Path.Str`
- [`923ea05`](https://github.com/perl6/doc/commit/923ea05) `Straighten up mkdir docs`
- [`aeeec94`](https://github.com/perl6/doc/commit/aeeec94) `Straighten up copy, move, rename`
- [`fff866f`](https://github.com/perl6/doc/commit/fff866f) `Fix docs for symlink/link routines`
- [`f83f78c`](https://github.com/perl6/doc/commit/f83f78c) `Use idiomatic Perl 6 in example`
- [`e60da5c`](https://github.com/perl6/doc/commit/e60da5c) `List utf-* alias examples too since they're common`
- [`0f49bb5`](https://github.com/perl6/doc/commit/0f49bb5) `List Rakudo-supported encodings in open()`
- [`017acd4`](https://github.com/perl6/doc/commit/017acd4) `Improve docs for IO::Path.slurp`
- [`56b50fe`](https://github.com/perl6/doc/commit/56b50fe) `Document IO::Handle.slurp`
- [`2aa3c9f`](https://github.com/perl6/doc/commit/2aa3c9f) `Document new behaviour of IO::Handle.Supply`
- [`a30fae6`](https://github.com/perl6/doc/commit/a30fae6) `Avoid potential confusion with use of word "object"`
- [`372545c`](https://github.com/perl6/doc/commit/372545c) `Straighten up file test docs`
- [`1527d32`](https://github.com/perl6/doc/commit/1527d32) `Document :completely arg to IO::Path.resolve`
- [`4090446`](https://github.com/perl6/doc/commit/4090446) `Improve chmod docs`
- [`d436f3c`](https://github.com/perl6/doc/commit/d436f3c) `Document IO::Spec::* don't do any validation`
- [`69b2082`](https://github.com/perl6/doc/commit/69b2082) `Document IO::Path.chdir`
- [`b9de84f`](https://github.com/perl6/doc/commit/b9de84f) `Remove DateTime tutorial from IO::Path docs`
- [`45e84ad`](https://github.com/perl6/doc/commit/45e84ad) `Move IO::Path.path to attributes`
- [`0ca2295`](https://github.com/perl6/doc/commit/0ca2295) `Reword/expand IO::Path intro prose`
- [`dbdc995`](https://github.com/perl6/doc/commit/dbdc995) `Document IO::Spec::*.catpath`
- [`50e5565`](https://github.com/perl6/doc/commit/50e5565) `Document IO::Spec::*.catdir and .catfile`
- [`28b6283`](https://github.com/perl6/doc/commit/28b6283) `Document IO::Spec::*.canonpath`
- [`a1cb80b`](https://github.com/perl6/doc/commit/a1cb80b) `Document IO::Spec::Win32.basename`
- [`5c1d3b6`](https://github.com/perl6/doc/commit/5c1d3b6) `Document IO::Spec::Unix.basename`
- [`9102b51`](https://github.com/perl6/doc/commit/9102b51) `Fix up IO::Path.basename`
- [`e9b6809`](https://github.com/perl6/doc/commit/e9b6809) `Document IO::Path::* subclasses`
- [`a43ecb9`](https://github.com/perl6/doc/commit/a43ecb9) `Document IO::Path's $.SPEC and $.CWD`
- [`7afd9c4`](https://github.com/perl6/doc/commit/7afd9c4) `Remove unrelated related classes`
- [`6bd0f98`](https://github.com/perl6/doc/commit/6bd0f98) `Dissuade readers from using IO::Spec*`
- [`184342c`](https://github.com/perl6/doc/commit/184342c) `Document IO::Special.what`
- [`56256d0`](https://github.com/perl6/doc/commit/56256d0) `Minor formatting improvements in IO::Special`
- [`1cd7de0`](https://github.com/perl6/doc/commit/1cd7de0) `Fix up type graph`
- [`b3a9324`](https://github.com/perl6/doc/commit/b3a9324) `Expand/fix up IO::Path.accessed`
- [`1973010`](https://github.com/perl6/doc/commit/1973010) `Document IO::Path.ACCEPTS`
- [`cc62dd2`](https://github.com/perl6/doc/commit/cc62dd2) `Kill IO::Path.abspath`
- [`1f75ddc`](https://github.com/perl6/doc/commit/1f75ddc) `Document IO::Spec*.abs2rel`
- [`24a6ea9`](https://github.com/perl6/doc/commit/24a6ea9) `Toss all of the TODO methods in IO::Spec*`
- [`65cc372`](https://github.com/perl6/doc/commit/65cc372) `Document IO::Path.concat-with`
- [`b9e692e`](https://github.com/perl6/doc/commit/b9e692e) `Document new IO::Path.extension`
- [`d5abceb`](https://github.com/perl6/doc/commit/d5abceb) `Write docs for all spurt routines`
- [`b8fba97`](https://github.com/perl6/doc/commit/b8fba97) `Point out my $*CWD = chdir ??? is an error`
- [`b78d4fd`](https://github.com/perl6/doc/commit/b78d4fd) `Include type names in links to methods`
- [`bdd18f1`](https://github.com/perl6/doc/commit/bdd18f1) `Fix desc of IO::Path.Str`
- [`a53015a`](https://github.com/perl6/doc/commit/a53015a) `Clarify value of IO::Path.path`
- [`5aa614f`](https://github.com/perl6/doc/commit/5aa614f) `Improve suggestion for Perl 5's opendir`
- [`bf377c7`](https://github.com/perl6/doc/commit/bf377c7) `Document &indir`
- [`e5225be`](https://github.com/perl6/doc/commit/e5225be) `Fix URL to &*chdir`
- [`e1a299c`](https://github.com/perl6/doc/commit/e1a299c) `Reword "defined as" for &*chdir`
- [`3fdc6dc`](https://github.com/perl6/doc/commit/3fdc6dc) `Document &*chdir`
- [`1d0e433`](https://github.com/perl6/doc/commit/1d0e433) `Document &chdir`
- [`d050d4b`](https://github.com/perl6/doc/commit/d050d4b) `Remove IO::Path.chdir prose`
- [`839a6b3`](https://github.com/perl6/doc/commit/839a6b3) `Expand docs for $*HOME and $*TMPDIR`
- [`db36655`](https://github.com/perl6/doc/commit/db36655) `Remove tip to use $*SPEC to detect OS`
- [`0511e07`](https://github.com/perl6/doc/commit/0511e07) `Document IO::Spec::*.tmpdir`
- [`cc6539b`](https://github.com/perl6/doc/commit/cc6539b) `Remove 8 methods from IO::Handle`
- [`335a98d`](https://github.com/perl6/doc/commit/335a98d) `Remove mention of role IO`
- [`cc496eb`](https://github.com/perl6/doc/commit/cc496eb) `Remove mention of IO.umask`
- [`3cf943d`](https://github.com/perl6/doc/commit/3cf943d) `Expand IO::Path.relative`
- [`ccae74a`](https://github.com/perl6/doc/commit/ccae74a) `Fix incorrect information for IO::Path.absolute`
- [`d02ae7d`](https://github.com/perl6/doc/commit/d02ae7d) `Remove IO::Handle.rw and .rwx`
- [`69d32da`](https://github.com/perl6/doc/commit/69d32da) `Remove IO::Handle.z`
- [`110efb4`](https://github.com/perl6/doc/commit/110efb4) `No need for .ends-with`
- [`fd7a41b`](https://github.com/perl6/doc/commit/fd7a41b) `Improve code example`


