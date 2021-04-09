
---
title: "Perl 6 Performance and Reliability Engineering grant report (Jonathan Worthington)"
author: Matthias Bloch
type: post
date: 2018-07-27 00:00:00 +0000 UTC
url: "/post/perl_6_performance_and_reliabi_5"
categories:

---

This report describes what I have accomplished in the [latest 200 hours](http://news.perlfoundation.org/2018/01/grant-extension-request-perl-6-1.html)
of my Perl 6 Performance and Reliability Engineering grant.

## A new MoarVM specializer plugin mechanism

The most significant new optimizations developed in this grant period center
around a new mechanism for helping the MoarVM specializer, which optimizes and
JIT-compiles hot code, to understand a wider range of Perl 6 constructs. I used
this mechanism to speed up:

* Private method calls taking place within roles (6.5x faster)
* Qualified method calls like `$obj.Foo::meth` (12x faster)
* Duck dispatches like `$obj.?foo` (2.5x faster if there is usually only one
  type really seen at the callsite in a particular program, and 1.5x faster
  otherwise)

This mechanism was also used in a re-work of `Scalar` containers, discussed a
little later in this report.

## MoarVM specializer usage analysis improvements

Various optimizations depend on having good information on what values are
used and where they are used in the program. Previously, we maintained a
simple usage count, and incremented that number to account for places where
the value would be needed if the code was deoptimized (that is, returning to
the interpreter if a speculative optimization turned out to be wrong).

First, I introduced use chains, which allow easy finding of instructions that
use the result of an operation. Along with this, I implemented a debug checker
that makes sure the usage information is maintained correctly as optimizations
are performed. This uncovered a number of usage bugs, all of which I fixed.

I used these chains to improve the elimination of useless `set` instructions
(which assign the contents of one register in the VM into another). Various
optimizations reduce other instructions to those, so it is an important cleanup step
that saves wasteful work in the optimized code. The previous implementation
of this optimization followed a number of adhoc rules, didn't rewrite the
graph in a way that upheld the SSA invariants, and still left behind a number
of `set` instructions that looked very possible to eliminate. The new algorithm
using the use chains deletes more useless `set` instructions and performs just
two different transformations, which between them cover what the adhoc rules
did beforehand and more.

I also used the chains to implement `box`/`unbox` pair analysis, discovering places
where we box a value into an object and then unbox it later. If that case we can
simply use the original unboxed value and potentially eliminate the `box` operation
and the memory allocation that goes with it. This optimization works, however it is not
yet very effective since it doesn't allow for elimination of a `box` that took
place in an inlined operation and unboxed outside of it.

The immediate reason that elimination of a `box` instruction in inlined code
could not take place is because we don't yet have a way to delete guard
instructions that we can prove are no longer needed, and the inlining code has
a guard on the return value type. However, even if we could do that, we didn't
have enough information to understand that deoptimization could no longer happen
at that point, and so we'd still retain the `box` that we'd really like to get
rid of.

In the course of this grant period, I took on the second part of the problem,
which is the more difficult of the two. This involved designing a new algorithm
to build more detailed information about what computations have to be retained
for the sake of surviving a deoptimization. Then, if we prove that such a
deoptimization can never happen, we may delete the instructions we retained
just in case.

Taken together, this work has given the optimizer far more detailed usage
information to work with, has already aided some optimizations, and will play
an important part in future optimizations too. Along the way, the optimizer
debug output was augmented to include this new information, which helps those
working on the optimizer to better understand why instructions are being
retained.

## More inlining

Inlining is an important optimization that sees the code of small callees
being incorporated into the caller, so they can be analyzed and optimized as
a single unit. It also eliminates the overhead of making the call. For Perl 6,
inlining is really important, because even basic operators like `+` are a
multiple-dispatch subroutine call.

Previously, we would only inline calls where we had already optimized the
callee. This usually worked out well, but in some cases this strategy did
not work out and the inlining opportunity would be missed. With the latest
changes, it will now be able to do inlining in such a case anyway.

Further, I enabled speculative inlining of void context calls. Previously
that had not mattered much, as we didn't generate very many of them.
However, I implemented a code-generation improvement in NQP that led to
many more void invoke instructions being generated. This got rid of a
wasteful throwaway boxing operation when the last instruction of a loop
block produced a natively typed value.

## Scalar improvements

`Scalar` containers are one of the most used, and most allocated, types in
most Perl 6 programs. For this reason, much of the code to deal with the
most common operations on `Scalar` containers was originally written in C. This made
sense when MoarVM lacked sophisticated optimization infrastructure. Today,
however, that C code is an opaque blob that the optimizer cannot see into.
We'd like to implement optimizations that will in some cases be able to
completely eliminate the allocation of short-lived `Scalar` containers, but
the optimizer not being able to see inside of them blocks this.

To enable all of that to happen, it was necessary to revisit `Scalar`. A
further problem was that auto-vivification was also done in a way that was hard
to optimize. I worked out a way to solve that problem and at the same time
make `Scalar` one attribute smaller, meaning we save 8 bytes (assuming a 64-bit
environment) off every `Scalar` container.

I used the new specializer plugin mechanism as part of my changes to `Scalar`
containers. This means that there is now the possibility to eliminate the
assignment type check and the `Nil` assignment special case when the optimizer
already knows the type that is being assigned.

This work fixed a number of shortcomings in the original container handling
code. One of those fixes (to `Proxy` handling in return) has been temporarily
removed again, since a number of ecosystem modules were relying on the bug. A
further more subtle precompilation related issue - again, an improvement on
the surface, but that has an interesting interaction with an ecosystem module -
is still under investigation.

These issues aside, I completed the `Scalar` overhaul, putting us in a good
place for the next round of interesting optimizations.

## Working on inlining, lookup, and context introspection problems

Inlining is essential for Perl 6 performance, however some features are quite
difficult to provide in combination with it. As inlining has become more and
more capable, these problems have been exposed. I spent some time on addressing
them, and implemented a solution for most cases.

The problem shows up in operations that want to walk caller and outer chains.
Inlining eliminates caller structure. Therefore, this has to be "reassembled"
sufficiently to answer the lookup or introspection in question. That's tricky,
but was being done in, for example, dynamic variable lookup. Other kinds of
lookup could not handle it, however.

I introduced a new abstraction - the frame walker - that exposes an iterator
over frames as they would exist had inlining not taken place. This allowed
for a great simplification to the dynamic variable lookup code, followed by
fixing a number of other forms of lookup.

This made things a great deal better. Remaining is to deal with some of the
most late-bound uses pseudo-packages like `CALLERS`, where they are stored
away rather than used for an immediate lookup. This is, thankfully, rarely
done, but still needs a fix.

## Regex compilation and evaluation improvements

I made a number of improvements to regex compilation and evaluation. The most
significant improvement helped with deciding which protoregex or alternation
branch to pick in the case that there was a huge number of different literal
prefixes. This was the case for parsing the `infix` category in the Perl 6
grammar, where there are a huge number of choices; the issue also shows up
for terms. The search for literal prefixes now has logarithmic rather than
linear complexity. In the compilation of the Rakudo CORE.setting, this led
to a 13% reduction in the time spent performing Longest Token Matching, and
took around 1% off the overall compilation time (which involves much more
than just parsing, though parsing is a significant part of it). User grammars
which also involve alternations or protoregexes with a large set of literal
prefixes will also benefit, since they use the same engine as we use to parse
Perl 6 source code.

I also worked to reduce some of the boilerplate code that is emitted when
compiling regexes, by analyzing whether it was needed. For example, the
capture stack handling code need not be generated if we can see that nothing
is captured.

Finally, I made calling of action methods a little cheaper by providing a
way to avoid a `can`/`findmethod` sequence of instructions; this required
two lookups in the method cache, whereas now only one is needed.

## Assorted additional performance improvements

* Fixed a performance bug in the thread pool scheduler where some sequences of
  operations would lead to an ever-growing call stack, thanks to mis-handling
  of continuations
* Reduced the number of allocations in some `Rat` operations
* Reduced the number of allocations in adverbed hash indexing operations, as
  well as other code using coercion types
* Specialization of `isnull` and `tryfindmeth`

## Assorted other fixes

* Fixed a memory leak when a hot loop repeatedly forced GC
* Fixed a memory leak involving `Promise.in(...)`
* Added missing functionality to get the process ID of a spawned asynchronous
  process
* Fixed a bug involving incorrect caching of type variable lookups
* Marked a routine using `samewith` as not being elligible for inlining, fixing a
  problem that could occur if that happened
* Fixed missing GC marking of contexts held by a serialization context, which
  showed up as an occasional segfault in some longer-running Perl 6 programs
* Fixed a bug involving lazy deserialization and vivification of lexicals in
  combination with inlining
* Marked `getaddrinfo` as blocking, so it taking a long time would not block
  GC and thus progress on other threads
* Fixed an `IO::Socket::Async` resource leak, and also a regression introduced
  by work to provide the host/port of a connection
* Fixed a bug where the results of `callframe` would become unpredictable if
  an onlystar `proto` was inlined
* Fixed a Rakudo code-gen bug for C-style `loop`s, whereby if it were the last
  statement in that `sub`, it would try to return a null, not a `Nil` like
  other loop constructs
* Corrected an off-by-one uninlining bug that could occur when an instruction
  at the very end of an inline was deleted
* Fixed return handlers of calls made in void context not getting the return
  value

## Enabling others

I also took the time to write detailed replies on a number of Rakudo tickets
to explain various issues, especially performance ones, such that others may
be able to look in to them.

## Dissemination

I wrote 5 blog posts:

* [Faster dispatches with MoarVM specializer plugins](https://6guts.wordpress.com/2018/06/09/faster-dispatches-with-moarvm-specializer-plugins/)
* [Better usage information in the MoarVM specializer](https://6guts.wordpress.com/2018/07/20/better-usage-information-in-the-moarvm-specializer/)
* [More precise deoptimization usage tracking](https://6guts.wordpress.com/2018/07/21/more-precise-deoptimization-usage-tracking/)
* [Dynamic lookups and context introspection with inlining](https://6guts.wordpress.com/2018/07/23/dynamic-lookups-and-context-introspection-with-inlining/)
* [Redesigning Rakudo’s Scalar](https://6guts.wordpress.com/2018/07/27/redesigning-rakudos-scalar/)

I also gave a [talk](http://act.yapc.eu/gpw2018/talk/7327) at the German
Perl Workshop about the MoarVM optimization infrastructure.

