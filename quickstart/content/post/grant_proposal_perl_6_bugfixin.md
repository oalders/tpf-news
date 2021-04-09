
---
title: "Grant Proposal: Perl 6: Bugfixing and Performance of Rationals / Fixing Constraints on Constants"
author: Coke
type: post
date: 2018-04-02 00:00:00 +0000 UTC
url: "/post/grant_proposal_perl_6_bugfixin"
categories:
 - Grants
 - Sign in

---

# Perl 6: Bugfixing and Performance of Rationals
# Fixing Constraints on Constants

- Name:

    Zoffix Znet

- Amount Requested:

    USD 1,999

## Synopsis

The proposal is to perform two pieces of work on the Rakudo Perl 6 compiler
along with a third bonus piece:

- 1. Implement support for type constraints on constants and polish
    some of the rough edges with `=` and `.=` constants initializer calls.
- 2. Fix several bugs and a race condition as well as fix problems in
    edge cases in instantiation and literals in core `Rational` types, also
    try to improve performance in this area.
- 3. _BONUS Work:_ Fix bugs in native `uint64` attributes behaving
    like signed `int64` types and try to use these attributes to boost
    performance of `Rat` type.



## Benefits to the Perl Community

- Performance is one of the bigger downsides of Perl 6. This proposal
    includes some performance improvements to one of the most used core types.
    Preliminary measurements show 1.5x improvement to instantiation of
    `Rat` type, and 9% improvement in a bench that measures all method calls.
    The crude measurements of the _BONUS Work_ item show an _additional_
    16% improvement in denominator value access, promising further improvements
    in numerous operators and methods of `Rat` type.
- Out-of-the-box support for `Rational` types is one of the flagship
    features of Perl 6, yet users trying out this feature are likely to
    come upon bugs and precision loss with literals.
    This proposal fixes those problems, by producing new non-infectious
    high-precision `MidRat` object when the user provides high precision
    in their literals. The allomorphiness of `MidRat`, for example, allows us
    to define `pi` as a high-precision `MidRat` that will function similar
    current `Num`-based `pi` in normal operations, yet offer extra precision
    in `FatRat` operations.
- The data race that currently exists in `Rational` types
    is rare and manifests itself not in crash, but a wrong result
    for a mathematical operation, which makes it hard to catch in real-life
    programs, even if they have a test suite. This proposal resolves that issue.
- Some mathematical operations on zero-denominator `Rational`s
    give entirely incorrect results and some object-identity bugs exist
    with all `Rational` types, which makes them problematic to use in Sets,
    Bags, and Mixes. This proposal resolves all of those issues.
- Type constraints on constants currently do get parsed, but they're
    not hooked into anything and don't perform any type checks. This makes
    unaware users think their code is doing work that in reality isn't done
    and can lead to bugs in their programs. This proposal resolves that issue.
- The way constants behave with different sigils, assigned values,
    and initializers is poorly defined and sparsely implemented. Users
    encountering these issues may get the impression the language is unfinished.
    This proposal resolves those issues and fully specs all of that behaviour.

## Deliverables

_Prior to the submission of this Grant Proposal, I drafted two work proposals,
presented to the Rakudo core developers, that spec the behaviour of type
constraints on
constants (["Constants Type Constraints" work proposal](https://github.com/rakudo/rakudo/blob/master/docs/archive/constants-type-constraints-proposal-2018-02-10.md)) as well as detail the work on fixing
various issues in
Rational types (["Polishing Rationals" work proposal](https://github.com/rakudo/rakudo/blob/master/docs/archive/2018-03-04--Polishing-Rationals.md)). After discussions with TimToady, the Rationals
proposal underwent two revisions with regard to how high-precision literals
are to behave. I'm still waiting for TimToady to comment on the final revision,
and depending on that feedback the object type for high-precision literals
that's produced might change slightly from what's described in this
Grant Proposal_.

The deliverables for this Grant Proposal are:

- Full implementation of the ["Constants Type Constraints" work proposal](https://github.com/rakudo/rakudo/blob/master/docs/archive/constants-type-constraints-proposal-2018-02-10.md), including Perl 6
    specification tests covering all of the mentioned features, and
    documentation fully documenting them.

    Some of the key items in this work proposal are:

    - 1. Implementation of `.=` initializer with `@`- and `%`-
        sigiled constants.
    - 2. Implementation of type check on the bound value when a type constraint
        is specified.
    - 3. Implementation of coercion of non-`Positional` values (for `@`
        sigils) and non-`Associative` values (for `%` sigils) to `List` and
        `Map` types respectively.
    - 4. Implementation of more helpful errors when some unsupported
        constructs are used.

- Full implementation of the ["Polishing Rationals" work proposal](https://github.com/rakudo/rakudo/blob/master/docs/archive/2018-03-04--Polishing-Rationals.md), including Perl 6
    specification tests covering all of the fixed bugs and implemented features,
    as well as full documentation for the user-facing changes.

    Some of the key items in this work proposal are:

    - 1. Implement a `MidRat` type that's an allomorph of `Rat` and
        `FatRat` types and will function as a `Rational` type capable of
        delivering high-precision data in operations with `FatRat`, yet
        degrading to a lower-precision `Rat` in other operations.
    - 2. Implement a `MidRatStr` type that's an allomorph of a `MidRat`
        and `Str` types.
    - 3. Make `Rational` number literals return a `MidRat` type instead
        of a broken `Rat`, when the requested denominator size is over 64 bits
    - 3. Make `&val`, quote words, and other allomorph-creating constructs
        create a `MidRatStr` instead of a broken `RatStr`,
        when the requested denominator size is over 64 bits
    - 4. Make `Rat.new` create a `MidRat` object, if after reduction the
        resultant `Rational` would have denominator over 64 bits.
    - 5. Fix a known data race by making `Rational` a fully-immutable type.
        This will involve the **removal** of optimization in `infix:<+>`
        and `infix:<->` operators.
    - 6. Perform optimization of various methods. Preliminary measurements show
        that despite the de-optimization to fix the data race (above) these
        optimizations will _improve_ the overall performance of the `Rat` type.
    - 7. Fix wrong results in many mathematical operations on
        zero-denominator `Rational`s
    - 8. Attempt performance enhancement of `Rational`
        types by extracting zero-denominator `Rational` handling logic into
        a separate `ZeroDenominatorRational` role. If no improvement is seen,
        this role will not be implemented.

- If successful, the _BONUS Work_ on `uint64` attributes
    will deliver fixed bugs in unsigned
    native type attributes, which currently erroneously behave as signed, as
    well as make `Rat` type use `uint64` type as denominator. Some nqp
    extops will be implemented in Rakudo
    (`nqp::p6gcd` and `nqp::p6div` ops used when creating `Rat`s,
    but possibly some functionality will be wrapped into a `nqp::p6rat` op).
    These ops will make use of the native-type denominators to compile to more
    efficient variants of `gcd` and `div` (`nqp::gcd_i` and `nqp::div_i` as opposed to slower
    big-int `_I` alternatives).

## Project Details

The full details for the two primary deliverables are documented in
the ["Constants Type Constraints" work proposal](https://github.com/rakudo/rakudo/blob/master/docs/archive/constants-type-constraints-proposal-2018-02-10.md) and the ["Polishing Rationals" work proposal](https://github.com/rakudo/rakudo/blob/master/docs/archive/2018-03-04--Polishing-Rationals.md).

The documents were originally prepared to detail the
planned work to other Perl 6 core developers and include considered-but-rejected
ideas as well as preliminary performance improvement
measurements that were made.

The original version of the "Polishing Rationals" work proposal received
feedback from TimToady with the proposal to overload a `RatStr` type
to be a non-infectious high-precision `Rational` type rather than upgrading high-precision literals to `FatRat` type. A trial implementation of that idea
identified a performance and semantical issue with that approach,
and in the latest revision of the "Polishing Rationals" work proposal
proposed new `MidRat`/`MidRatStr` types used instead. TimToady is yet to
review that revision, and depending on feedback he gives, this small detail
might change from what's currently described.

\----------------------------------

For the _BONUS Work_ item, I marked it as "BONUS Work" because it requires me
learning new things about MoarVM and the QAST-MAST compiler. As I currently
don't know much about them, I can't say with 100% certainty that I will
succeed in fixing the bugs mentioned in the _BONUS Work_. However, if I am
successful, the new knowledge will also allow me to fix many other bugs in my
role as a Rakudo core developer, even after this Grant has completed.

## Inch-stones

- 1.1 Write Perl 6 Specification tests specifying behaviour described
    in the ["Constants Type Constraints" work proposal](https://github.com/rakudo/rakudo/blob/master/docs/archive/constants-type-constraints-proposal-2018-02-10.md)
- 1.2 Implement type-constraint checking logic in
    `type_declarator:sym<constant>($/)` method in the Actions
- 1.3 Implement typecheck and coercers on constants with `@` and `%`
    sigils.
- 1.4 Implement additional typed exceptions described in the work proposal
    and include additional Grammar rules in the
    `type_declarator:sym<constant>($/)` token to give more useful error
    messages when unsupported constructs are used.
- 1.5 Fully document the newly-clarified behaviour on
    [https://docs.perl6.org/language/terms#Constants](https://docs.perl6.org/language/terms#Constants)
- 2.1 Write Perl 6 Specification tests specifying behaviour described
    in the ["Polishing Rationals" work proposal](https://github.com/rakudo/rakudo/blob/master/docs/archive/2018-03-04--Polishing-Rationals.md)
- 2.2 Remove `.REDUCE-ME` and `&DON'T_DIVIDE_NUMBERS` routines and their
    use, and always go through reduction for all operations instead.
- 2.3 Optimize `Rational.new` and `&DIVIDE_NUMBERS` routines
    using techniques that produced preliminary performance improvement
    measurements when the work proposal was drafted
- 2.4 Make `Rational.new` normalize zero-denominator `Rational`s to
    `<1/0>`, `<-1/0>`, `<0/0>`. Fix zero-denominator math
    issues, if any still remain after this change.
- 2.5 Evaluate whether any significant performance advantages are produced
    by demarcating zero-denominator rationals with
    `ZeroDenominatorRational` role and extracting zero-denominator-rational-specific
    logic into separate multi candidates for operators and routines. If there
    are such advantages, fully implement zero-denominator rationals in
    terms of that role.
- 2.6 Implement `MidRat` and `MidRatStr` types.
- 2.7 Make `Rat.new` return `MidRat` if after reduction the denominator
    is over 64 bits.
- 2.8 Ensure `Rational` literals produce `MidRat` types when the given
    precision produces denominators over 64-bit. This may fall-out automatically
    from `Rat.new` change in (2.7); if not, modify Grammar and Actions to
    support this behaviour.
- 2.9 Make `&val` produce `MidRatStr` for `Rational`s with denominators
    over 64-bit. This may fall-out automatically from `Rat.new` change in
    (2.8), but if not, modify the logic in `&val` routine.
- 2.10 Ensure quotewords and other allomorph-creating constructs (such
    as arguments to `sub MAIN`) create `MidRatStr` objects for `Rational`s
    with denominators over 64 bits. This may fall-out automatically from
    changes in (2.6) and (2.8).
- 2.11 Document the `MidRat` and `MidRatStr` types as well as the
    newly-clarified behaviour for `Rational` literals, zero-denominator
    `Rational`s, and `Rat.new`.
- 3.1 _BONUS Work_: Write tests in Perl 6 Specification covering unsigned
    attributes and variables, ensuring to test the behaviour that's expected
    from specifically unsigned types (and isn't present in signed types)
    exists (some of these tests should already exists as TODOs)
- 3.2 _BONUS Work_: study how native types are implemented and implement
    full support for unsigned types on attributes
- 3.3 _BONUS Work_: Change parametarization of the `Rat` type to be
    `[Int, uint64]` and change supporting routines like `&DIVIDE_NUMBERS` to
    handle native type arguments without having to box them.
- 3.4 _BONUS Work_: implement `nqp::p6gcd` and `nqp::p6div` ops
    that compile to `nqp::gcd_i` and `nqp::div_i` if their parameters are
    native attributes (and to `_I` variants otherwise). Alternatively or
    in addition to, implement more convenient ops for handling of `Rational`s;
    what that might be will be more obvious after all of the work on Rationals
    in item (2) is completed.
- 4.1 Go through the bug ticket queue and marked all tickets fixed by
    this Grant as resolved, writing explicit tests to cover the reported bugs,
    as needed.

## Project Schedule

I can begin work as soon as this Grant Proposal is approved.

I estimate it will take me 1 month to finish the work on the constants, 1 month
to finish the work on `Rational`s, and 2 months to finish the _BONUS Work_
on native attributes and corresponding optimizations on `Rat` with native
denominator.

## Completeness Criteria

- 1. Commits implementing the work in this proposal will be committed to [rakudo repository](https://github.com/rakudo/rakudo/) (the _BONUS Work_ portion
will also see commits made to [nqp repositor](https://github.com/perl6/nqp/) and [MoarVM repository](https://github.com/MoarVM/MoarVM)).
- 2. Commits specifying behaviour in the work proposals as well as testing
    fixed bugs will be commited to the [Perl 6
    Specification repository](https://github.com/perl6/roast).
- 3. Commits documenting all newly-clarified behaviour will be made
    to the [Perl 6 Documentation repository](https://github.com/perl6/doc).

**NOTE:** the extra NQP operators described in the _BONUS Work_ will be
implemented only for the MoarVM backend. Their implementation for JVM and
JavaScript backends is outside the scope of this Grant Proposal.

## Bio

I've been a member of the Perl 6 Core Development Team since summer of 2016.
I have previously successfully completed the [Perl 6 IO Grant](http://news.perlfoundation.org/2017/01/grant-proposal-standardization.html) with The Perl Foundation and to date have made **1,181** commits to [The Perl 6 Documentation](https://github.com/perl6/doc), **1,415** commits to
[The Perl 6 Specification](https://github.com/perl6/roast), and **1,641** commits to the [Rakudo repository](https://github.com/rakudo/rakudo/). I also wrote
numerous Perl 6 tutorials published
on the [Rakudo.Party](https://rakudo.party/) website.

My previous work in the area of constants and initializers
includes fixing bugs with `.=` initializer
on attributes, sigiless variables, and `$`- and sigiless constants for types
with `::` in their names as well as compound types such as `Array[Int]` (Rakudo commits `7793f420e`, `abea32429`, `562edfc50`, `8ba3c86e7`, and
`700a07747`)

My previous work with `Rational`s includes fixing numerous issues with
zero-denominator rationals, fixing several bugs due to `.REDUCE-ME`
optimization that I plan to remove/redesign as part of this grant, as well
as performing some optimizations to routines involving `Rational`s.
(Rakudo commits `8aeaf469`, `6c299bf9f`, `cb2476f9b`, `73182d4e9`, `748d1a57b`,
`c91bcc2af`, `042cb7413`, `893d09ffa`, `6dbe85eda`, `aac9efcbd`,
`79553d0fc`, `637241703`, `7434a8f73`, and `b5aa3c591`).

My previous work on QAST/MAST/MoarVM (which would be involved to complete _BONUS Work_) involves fixing spurious "useless use" warnings,
re-designing the `WhateverCode` currier to be more performant (both in
compile-time and in run-time), fixing numerous bugs with
`QAST::Block` migration, implementation of hypered versions of extended
method call variants (e.g. `$foo>>.?&elems`), fixing crashes with native
types in conditionals, as well as implementation of optimizations in
`dispatch:<var>` method calls, conditionals with natives, and
implementation of `nqp::chainstatic` op on MoarVM backend for staticalizing
of chained op calls. (Rakudo commits `fb3dfa567`, `c0c7756f4`, `1ee89b540`,
`e8c6c259c`, `752bb8b38`, `58de239cc`, `ef2dc1b8b`, `3c4041eab`, and
`97359ae42`; NQP commits `71658ad85`, `df45cb8ce`, and `414520586`). I
also fixed bugs in zeroing of VMArray elements, flushing of TTYs on WSL, EOF
detection, and `abs_n` and `div_i` ops.
(MoarVM commits `8d94732aa`, `4541cf6f6`, `cb4c1941a`, `09482f9bd`,
`43c926f9e`, `912f96783`, `3f3045d6e`, `cfb0bffc0`, `9d7bee40e`,
`fd7300d27`, and `1c1746e52`).


