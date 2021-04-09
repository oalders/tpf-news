
---
title: "Final Grant Report - MoarVM JIT Compiler Expression Backend"
author: Mark A Jensen
type: post
date: 2019-09-14 00:00:00 +0000 UTC
url: "/post/final_grant_report_-_moarvm_ji"
categories:
 - Grants
 - Sign in

---

While a number of intended deliverables for [Bart's
grant](http://news.perlfoundation.org/2018/10/grant-proposal-moarvm-jit-comp.html)
remain unmet, he's made significant contributions to the expression
backend and has identified unexpected roadblocks to the remaining
tasks (as outlined in previous reports) that should yield to
additional preparatory work.

In light of this, the Grants Committee will be considering Bart's
report below and voting on a payment for the currently accomplished work of 50% of the original amount
requested. The Grants Committee
will consider a revised proposal for the remainder of the work via the
[usual proposal
process](http://news.perlfoundation.org/2019/09/call-for-grant-proposals-septe-3.html#more).

Comments are welcome, as always.

MAJ


*Report for grant project _MoarVM JIT Compiler Expression Backend
Maturation_*

*Bart Wiegmans, September 12, 2019*

This is an intermediate report after the first major deliverable of
the grant project has been merged and released to the perl community.

This project aims to improve and mature the new 'expression' JIT
backend for MoarVM by the following deliverables:

-   Floating point support in the expression JIT backend.

-   Code quality improvements (optimizations) based on expression-tree
rewriting optimizer, specifically:

    -   Completing the work on the jit-expr-optimizer branch,
        including several common optimizations.

    -   Implement optimizations in the register allocator.

-   Replaced specialized code in the legacy JIT with portable
expression tree structures, specifically:

    -   Invocation of subroutines

    -   Optimization guards

    -   NativeCall invocations

-   Improved handling of irregular instructions by the register
allocator.

The first of those deliverables, floating point support, has been
merged into the MoarVM main tree. I learned that this was a
considerable larger challenge than initially anticipated. Aside from
the necessary DynASM changes, changes were needed in:

-   The expression template precompiler ([specifically a type
system](http://brrt-to-the-future.blogspot.com/2019/01/a-short-posts-about-types-and.html))

-   The register allocator (this was expected), including an improved
register representation

-   Compilation of boolean operators (floating point comparisons set a
different set of CPU flags than do integer comparisons, to handle NaN
correctly)

Along with floating point support, improvements to register allocation
algorithm have also been merged and released. I've researched in using
a somewhat different algorithm ([reverse linear scan
allocation](http://brrt-to-the-future.blogspot.com/2019/03/reverse-linear-scan-allocation-is.html))
but have deprioritized it. This is an interesting algorithm that I
expect to do better than the current 'regular' linear scan algorithm
in many cases, but unlike the current algorithm it isn't fully
general. Handling irregular register requirements has been put on hold
until I make a final decision.

Merging floating point supports has unblocked progress on subroutine
and nativecall invocations. The limiting factor currently is that I
don't have a good testcase for the nativecall case.

The implementation of (de)optimization guards in the expression JIT is
currently in progress and nearing completion.

The expression optimization work has stumbled on a roadblock described
in one of my earlier
[reports](http://news.perlfoundation.org/2019/03/grant-report---moarvm-jit-comp-2.html) - the current IR does not reify the order of operations, which is
instead handled in a later phase (during instruction selection). This
means that it is nearly impossible to safely implement optimizations
since an optimization may accidentally move the write of a value after
a read of it. The good news is that removing that roadblock should
also enable the expre