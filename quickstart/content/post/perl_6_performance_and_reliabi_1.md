
---
title: "Perl 6 Performance and Reliability Grant Progress Report"
author: Karen Pauley
type: post
date: 2016-08-30 00:00:00 +0000 UTC
url: "/post/perl_6_performance_and_reliabi_1"
categories:
 - Grants
 - Sign in

---

_Jonathan Worthington writes:_

I have completed the initial 200 hours awarded under my "Perl 6 performance and reliability engineering grant":http://news.perlfoundation.org/2016/02/grant-proposal-perl-6-performa.html. This report summarizes what has been achieved in this time. I have also written a number of more detailed "blog posts":https://6guts.wordpress.com/ about my work.

**Tooling**

I implemented heap snapshots in MoarVM. This is a mechanism for taking recordings of what is in the heap after each garbage collection run. It can be used to understand the memory use of programs, but also to track down memory leaks. The snapshots are produced by passing the --profile=heap option when invoking Rakudo Perl 6. They can then be analyzed using a "tool":https://github.com/jnthn/p6-app-moarvm-heapanalyzer/, which I implemented in Perl 6. It makes good use of both native arrays and parallel processing, and so also serves as a good example of a Perl 6 program processing a non-trivial volume of data. The data appears to be something of a goldmine, and following up on and addressing everything raised by it will probably keep us busy for a good while. It has already been used to track down memory leaks and fix them.






**Performance**

My performance work largely focused on lower-level improvements, rather than optimizing the Perl 6 built-ins (which are already receiving plenty of attention). Each improvement is annotated with the components affected. I:

* Had the compiler code-generate accessor methods where possible, rather than them being closures added by a trait. This in turn made them inlinable. This made attribute accessors many times faster. (Rakudo)
* Re-designed and re-implemented the MoarVM multiple dispatch cache, so it can handle named parameters. I made it more memory compact and faster to search along the way. With changes to Rakudo to take advantage of it, multiple dispatches involving named arguments got much faster. This most notably impacted constructs like @a[$i]:exists, which got around 20 times faster. (MoarVM, Rakudo)
* Made a number of improvements to how return is implemented. This made it a real control exception, as per the language design documents. More notably, however, the changes made return a couple of times faster when used, cut the cost of nearly all routine invocations whether they used return or not, made it possible for more routines to be inlined, and reduced memory consumption. Further optimizations in this area will be possible thanks to these changes. (MoarVM, NQP, Rakudo)
* Optimized throwing of next/last/redo control exceptions, making them much cheaper in the common case. (Rakudo)
* Implemented lazily decoding the string heap, improving startup time. Gave 1.26MB less base Rakudo memory use, and shaved 2.7 million CPU cycles off startup. (MoarVM)
* Fixed a string decoding performance bug that made reading very long lines extremely slow. (MoarVM)
* Knocked 80% off get_boxed_ref, which is a hot path in Int math. (MoarVM)
* Eliminated generating various unrequired decont operations at code-gen time. (NQP)
* Significantly overhauled MoarVM's call frame handling, eliminating reference counts, simplifying memory management, fixing excessive GC time in programs that store a huge number of closures, and preparing the way for a number of future improvements. (MoarVM)
* Avoided various bits of NULLing on frame entry and initialization, especially in specialized (optimized) frames. (MoarVM)
* Fix bugs with a submitted patch that made serialization and compilation vastly faster for large compilation units (such as the Rakudo CORE.setting), so that it could be merged in. (MoarVM)
* Various additional optimizations to invocation, that added up to shaving a couple of perfect off an invocation-heavy benchmark. (MoarVM)

**Memory leak fixes and other memory use improvements**

Aided by the heap analyzer described earlier in this report, along with tools from the Valgrind suite, I tracked down and fixed a number of memory leaks, and also reduced memory use.

* Fixed a memory leak that could affect multi dispatch + constraints + flattening, and likely other situations. (MoarVM)
* Elimianted near-unused static frame array in MVMCompUnit, saving some hundreds of kilobytes off the Rakudo base memory. (MoarVM)
* Fixed a memory leak affecting EVAL and everything using it (some cases of regex interpolation were also impacted, for example). There were actually multiple problems, identified through using the heap analyzer. (NQP, MoarVM)
* GC performance analysis and tuning when we do full collections, to improve memory behavior for various kinds of program. (MoarVM)
* Eliminate caching of call contexts, reducing the size of all call frames and simplifying/cheapning context serialization. (MoarVM)
* Eliminated retention of barely-used bytecode maps produced during validation and only initialized frame instrumentation state if needed, adding up to 3.5MB of savings on Rakudo's base memory. (MoarVM)

**Concurrency bug fixes**

A number of concurrency bugs were tracked down and eliminated, improving the reliability of programs using Perl 6's parallel and concurrent programming features.

* Fixed a number of data races around thread spawning and the first GC run of the thread. (MoarVM)
* Fixed mis-handling of spurious condition variable wake-ups in Promise.result. (Rakudo)
* Fixed hang reported in RT #128628 by adding missing GC block/unblock around semaphore wait. (MoarVM)
* Fixed deadlock that could occasionally occur in the concurrent blocking queue used for task scheduling. (MoarVM)
* Added missing GC rooting around concurrency control constructs when marking themselves blocked/unblocked. (MoarVM)
* Fix a race condition in the Channel.Supply coercer. (Rakudo)
* Fix a circular waiting bug that led to occasional deadlocks in some uses of the supply and react syntax. (Rakudo)
* Made .close of a listening socket tap await the actual shutdown of the socket, fixing a race that caused instability in the async socket tets. (MoarVM, Rakudo)
* Give start blocks a fresh $/ and $!. (Rakudo)
* Track down the problem with S17-lowlevel/lock.t sometimes failing/crashing; correct a bug in the test, resolving the problem. (Spectest)
* Start re-working VMArray so mis-use of it across threads cannot cause crashes. (MoarVM)

**Other assorted fixes**

I fixed a selection of other problems, mostly coming from the RT bug queue. I've grouped them by the component that was primarily fixed.

Rakudo

* RT #127548 (crash involving uint64 attribute code-gen)
* RT #127660 (didn't pay attention to submethod Bool)
* RT #127629 (issues with conveying exceptions in Supply <-> Channel coercions)
* Fixed Mu.Str to use objectid, not memory address, eliminating some test instabilities
* RT #127540 (anon subs triggering a bogus redeclaration error)
* RT #128270 (mis-compilation of charset with ignoremark led to crashes, e.g. if used in combination with :g)
* RT #128581 (poor error reporting for my Array[Numerix] $x)
* RT #127749 (Seqs should not be stuck into the constants table)
* RT #127785 (parse error if trying to use where in a unit sub MAIN(...) signature)
* RT #127473 (compiler explodes on (;), (;;), [0;] and similar)
* RT #127394 (cannot write -> SomeSubtype:D $x { } as it produced a compiler error; now it works)
* RT #128552 (missing $?MODULE and ::?MODULE symbols)
* Fixed a hang in spectests on Windows, which ended up involving file lock mis-management in precompilation handling

MoarVM

* RT #127530 (SEGV on certain concatenating certain characters)
* RT #127272 (a JIT compilation bug in the string ge/le operators)
* RT #123602 and RT #127782 (repeat + concat + substr interaction bug)
* RT #126756 (SEGV on single utf8-c8 synthetic)
* RT #127748 (SEGV due to a GC invariant violation, which led to memory corruption)
* Fix various missing GC rootings and write barriers uncovered in stress testing

**Other assorted tasks**

Some time was spent on the following tasks:

* Reviewed various pull requests to Rakudo/NQP/MoarVM, providing feedback and/or merging them as appropriate.
* Got OSX Travis support for MoarVM setup, after a regression on OSX that got missed.
* The odd bit of bug queue wrangling (merging duplicates, closing already fixed issues, rejecting things that are not bugs, etc.)


