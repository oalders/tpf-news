
---
title: "Perl 6 IO Grant: March 2017 Report"
author: Coke
type: post
date: 2017-04-04 00:00:00 +0000 UTC
url: "/post/perl_6_io_grant_march_2017_rep"
categories:
 - Grants

---

*Zoffix Znet provided this report on March 28, 2017*

# Perl 6 IO TPF Grant: Monthly Report (March, 2017)

This document is the March, 2017 progress report for [TPF Standardization,
Test Coverage, and Documentation of Perl 6 I/O Routines
grant](http://news.perlfoundation.org/2017/01/grant-proposal-standardization.html)

## Timing

My delivery of the Action Plan was one week later than I originally
expected to deliver it. The delay let me assess some of the big-picture
consistency issues, which led to proposal to remove 15 methods from IO::Handle
and to iron out naming and argument format for several other routines.

I still hope to complete all the code modifications prior to end of weekend of
April 15, so all of these can be included in the next Rakudo Star release. And
a week after, I plan to complete the grant.

Note: to minimize user impact, some of the changes may be included only in
6.d language, which will be available in 2017.04 release only if the user uses
`use v6.d.PREVIEW` pragma.

## IO Action Plan

I finished the IO Action Plan, [placed it into `/doc` of rakudo's repository](https://github.com/rakudo/rakudo/blob/nom/docs/2017-IO-Grant--Action-Plan.md), and made it available to other core devs for
review for a week (period ends on April 1st). The Action Plan is 16 pages long and
contains 26 sections detailing proposed changes.

Overall, I proposed many more smaller changes than I originally expected and
fewer larger, breaking changes than I originally expected. This has to do
with a much better understanding of how rakudo's IO routines are "meant to" be
used, so I think the improvement of the documentation under this grant will
be much greater than I originally anticipated.

A lot of this has to do with lack of explanatory documentation for how to
manipulate and traverse paths. This had the effect that users were using the
`$*SPEC` object (157 instances of its use in the ecosystem!) and its routines
for that goal, which is rather awkward.
This concept is prevalent enough that I even wrote [`SPEC::Func`](https://github.com/zoffixznet/perl6-SPEC-Func) module in the
past, due to user demand, and certain books whose draft copies I read used
the `$*SPEC` as well.

In reality, `$*SPEC` is an internal-ish thing and unless you're writing your
own IO abstractions, you never need to use it directly. The changes and
additions to the `IO::Path` methods done under this grant will make traversing
paths even more pleasant, and the new tutorial documentation I plan to write
under this grant will fully describe the Right Way??? to do it all.

In fact, removal of `$*SPEC` in future language versions is currently under
consideration...

## Removal of `$*SPEC`

lizmat++ pointed out that we can gain significant performance improvements by
removing `$*SPEC` infrastructure and moving it into module-space. For example,
a benchmark of slurping a 10-line file shows that removal of all the
path processing code makes benched program run more than 14x faster. When
benching `IO::Path` creation, dynamic var lookup alone takes up 14.73% of
the execution time.

The initial plan was to try and make IO routines handle all OSes in a unified
way (e.g. using `/` on Windows), however it was found this would create
several ambiguities and would be buggy, even if fast.

However, I think there are still a lot of improvements that can be gained
by making `$*SPEC` infrastructure internal. So we'd still have the
`IO::Spec`-type modules but they'll have a private API we can optimize freely,
and we'll get rid of the dynamic lookups, consolidate what code we can into
`IO::Path`, while keeping the functionality that differs between OSes in the
`::Spec` modules.

Since this all sounds like guestimation and there's a significant-ish use of
`$*SPEC` in the ecosystem, the plan now is to implement it all in a module
first and see whether it works well and offers any significant performance
improvements. If it does, I believe it should be possible to swap `IO::Path`
to use the fast version in `6.d` language, while still leaving `$*SPEC` dynvar
and its modules in core, as deprecated, for removal in `6.e`.

This won't be done under this grant, and while trying not to over-promise, I
hope to release this module some time in May-June. So keep an eye out for it; I
already picked out a name: [`FastIO`](https://modules.perl6.org/repo/FastIO)

## newio Branch

As per original goals of the grant, I reviewed the code in Rakudo's 2014???2015
`newio` branch, to look for any salvagable ideas. I did not have any masterplan
design documents to go with it and I tried a few commits but did not find one
that didn't have merge conflicts and compiled (it kept complaining about
ModuleLoader), so my understanding of it comes solely from reading the source
code, and may be off from what the original author intended it to be.

The major difference between `newio` and Rakudo's current IO structure is
type hierarchy and removal of `$*SPEC`. `newio` provides `IO::Pathy` and
`PIO` roles which are done by `IO::File`, `IO::Dir`, `IO::Local`, `IO::Dup`,
`IO::Pipe`, and `IO::Huh` classes that represent various IO objects. The current
Rakudo's system has fewer abstractions: `IO::Path` represents a path to an IO
object and `IO::Handle` provides read/write access to it, with `IO::Pipe`
handling pipes, and no special objects for directories (their contents are
obtained via `IO::Path.dir` method and their attributes are modified via
`IO::Path` methods).

Since 6.d language is *additive* to 6.c language, completely revamping the
type hierarchy may be challenging and messy. I'm also not entirely sold on what
appears to be one of the core design ideas in `newio`: most of the
abstractions are of IO objects as *they were at the object instantiation time*. An `IO::Pathy` object represents an IO item *that exists*, despite there being
no guarantees that it actually does. Thus, `IO::File`'s `.f` and `.e` methods
always return `True`, while its `.d` method always returns `False`. This
undoubtedly gives a performance enhancement, however, if
`$ rm foo` were executed after `IO::File` object's creation, the `.e` method
would no longer return correct data and if then `$ mkdir foo` were
executed, both `.f` and `.d` methods would be returning incorrect data.

Until recently, Rakudo cached the result of `.e` call and [that produced
unexpected by user behaviour](https://rt.perl.org/Ticket/Display.html?id=130889). I think the
issue will be greatly exacerbated if this sort of caching is extended to entire
objects and many of their methods.

However, I do think the removal of `$*SPEC` is a good idea. And as described in
previous section I will try to make a `FastIO` module, using ideas from `newio`
branch, for possible inclusion in future language versions.

## Experimental MoarVM Coverage Reporter

As was mentioned in my grant proposal, the coverage reporter was busted by
the upgrade of information returned by `.file` and `.line` methods on core
routines.
MasterDuke++ made several commits fixing numerous issues to the coverage
parser and last night I identified the final piece of the breakage. The
annotations and hit reports all use the new `SETTING::src/core/blah` file
format. The setting file has `SETTING::src/core/blah` markers inside of it.
The parser however, still thinks it's being fed the old `gen/moar/CORE.setting`
 filenames, so once I teach it to calculate proper offsets
into the setting file, we'll have coverage reports on [perl6.wtf](https://perl6.wtf) back up and running and I'll be able to use them
to judge IO routine test coverage required for this grant.

## Performance Improvements

Although not planned by the original grant, I was able to make the following
performance enhancements to IO routines. So hey! Bonus deliverables \o/:

- [rakudo/fa9aa47](https://github.com/rakudo/rakudo/commit/fa9aa47) Make `R::I::SET_LINE_ENDING_ON_HANDLE` **4.1x** Faster
- [rakudo/0111f10](https://github.com/rakudo/rakudo/commit/0111f10) Make IO::Spec::Unix.catdir **3.9x** Faster
- [rakudo/4fdebc9](https://github.com/rakudo/rakudo/commit/4fdebc9) Make IO::Spec::Unix.split **36x** Faster
    - Affects IO::Path's .parent, .parts, .volume, .dirname, and .basename
    - Measurement of first call to .basename shows it's now **6x-10x** faster
- [rakudo/dcf1bb2](https://github.com/rakudo/rakudo/commit/dcf1bb2) Make IO::Spec::Unix.rel2abs **35%** faster
- [rakudo/55abc6d](https://github.com/rakudo/rakudo/commit/55abc6d) Improve IO::Path.child perf on `*nix`:
    - make IO::Path.child **2.1x** faster on `*nix`
    - make IO::Spec::Unix.join **8.5x** faster
    - make IO::Spec::Unix.catpath **9x** faster
- [rakudo/4032953](https://github.com/rakudo/rakudo/commit/4032953) Make IO::Handle.open **75%** faster
- [rakudo/4eef6db](https://github.com/rakudo/rakudo/commit/4eef6dbf3789502205c254314f27ef33af549adc) Make IO::Spec::Unix.is-absolute about **4.4x** faster
- [rakudo/ae5e510](https://github.com/rakudo/rakudo/commit/ae5e510fd00c5544d99f56d6aeb17d222d56b220) Make IO::Path.new **7%** faster when creating from Str
- [rakudo/0c6281](https://github.com/rakudo/rakudo/commit/0c6281518e5c78113121968df0cf7404aa949dd3) Make IO::Pipe.lines use IO::Handle.lines for **3.2x** faster performance

## Performance Improvements Made By Other Core Developers

lizmat++ also made these improvements in IO area:

- [rakudo/b4d80c0](https://github.com/rakudo/rakudo/commit/b4d80c0) Make .IO.slurp about **2x** as fast
- [rakudo/9da50e3](https://github.com/rakudo/rakudo/commit/9da50e3) Introducing IO::Handle.iterator
- [rakudo/9019a5b](https://github.com/rakudo/rakudo/commit/9019a5b) Streamline IO::Handle.get/getc
- [rakudo/4bc826d](https://github.com/rakudo/rakudo/commit/4bc826d) Streamline IO::Handle.get

Along with the commits above, she also made IO::Handle.lines faster and
eliminated a quirk that required
custom `.lines` implementation in IO::Pipe (which is a subclass of IO::Handle).
Due to that, I was able to remove old IO::Pipe.lines implementation and make it
use new-and-improved IO::Handle.lines, which made the
method about 3.2x faster.

## Bugs

### Will (attempt to) fix as part of the grant

- `IO::Pipe` inherits `.t` method from from `IO::Handle` to check if the handle
is a TTY, however, attempt to call it causes a segfault. MasterDuke++ already
found the candidate for the offending code
([MoarVM/Issue#561](https://github.com/MoarVM/MoarVM/issues/561)) and this
should be resolved by the time this grant is completed.

### Don't think I will be able to fix these as part of the grant

- Found a strange error generated when `IO::Pipe`'s buffer is filled up.
This is too deep in the guts for me to know how to resolve yet, so I filed it as
    [RT#131026](https://rt.perl.org/Ticket/Display.html?id=131026)

### Already Fixed

- Made IO::Path.new 7% faster when creating from paths strings and fixed
    failure to detect [rakudo/ae5e51](https://github.com/rakudo/rakudo/commit/ae5e510fd00c5544d99f56d6aeb17d222d56b220)
- Made IO::Spec::Unix.is-absolute about 4.4x faster     
    [rakudo/4eef6d](https://github.com/rakudo/rakudo/commit/4eef6dbf37)
- Found that IO::Path had a vestigial .pipe method that delegated to a
    non-existant IO::Handle method. Removed in [rakudo/a01d67](https://github.com/rakudo/rakudo/commit/a01d6794d2d37b574011198cc4928f77f8c33361)
- Fixed IO::Pipe.lines not accepting a Whatever as limit, which is accepted by
    all other .lines. [rakudo/0c6281](https://github.com/rakudo/rakudo/commit/0c6281518e5c78113121968df0cf7404aa949dd3)
    Tests in [roast/465795](https://github.com/perl6/roast/commit/465795c458041e66e33e32e2de2b8cd358be5961) and [roast/add852](https://github.com/perl6/roast/commit/add852b082a2fca83dbefe03d890dd5939c5ff45)
- Fixed issues due to caching of `IO::Handle.e`. Reported as
    [RT#130889](https://rt.perl.org/Ticket/Display.html?id=130889). Fixed in
    [rakudo/76f718](https://github.com/rakudo/rakudo/commit/76f71878da61731f33b457e84c7b0e801c64af66).
    Tests in [roast/908348](https://github.com/perl6/roast/commit/908348eef18b1c33f1bd8d879b9bb16f002fb6f7)
- Rejected [rakudo PR#666](https://github.com/rakudo/rakudo/pull/666)
    and resolved [RT#126262](https://rt.perl.org/Ticket/Display.html?id=126262) by explaining why the methods return `Str` objects instead of `IO::Path` on
    ticket/PR and improving the documentation by
    fixing mistakes ([doc/ccae74](https://github.com/perl6/doc/commit/ccae74a1502285d8b82697b68a8e26a31ca762d7)) and expanding ([doc/3cf943](https://github.com/perl6/doc/commit/3cf943d86bef3744146e31e106815a00a2a81f4a)) on what the methods do exactly.
- IO::Path.Bridge was defunct, as it was trying to call .Bridge on Str, which
    does not exist. Resolved the issue by deleting this method in [rakudo/212cc8](https://github.com/rakudo/rakudo/commit/212cc8ae5d)
- Per demand, made `IO::Path.dir` a multi, so module-space can augment it with
    other candidates that add more functionality. [rakudo/fbe7ace](https://github.com/rakudo/rakudo/commit/fbe7ace6fc19d86ac1cb0519654e4239c1a17129)

