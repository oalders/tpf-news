
---
title: "Grant Proposal: Future::AsyncAwait"
author: Coke
type: post
date: 2018-10-04 00:00:00 +0000 UTC
url: "/post/grant_proposal_futureasyncawai"
categories:
 - Grants

---

The Grants Committee has received the following grant proposal for the Sep/Oct round.
Before the Committee members vote, we would like to solicit feedback from the Perl
community on the proposal.

**Update: You have until October 17th!**
Review the proposal below and please comment here by October 10th, 2018. 
The Committee members will start the voting process following that.

# Future::AsyncAwait

- Name:

    Paul Evans

- Amount Requested:

    GBP 4,800 (GBP 200/day for 24 days)

    At time of writing, equivalent to USD 6,256 or EUR 5,388.

## Synopsis

Fix and polish the implementation of the [Future::AsyncAwait](https://metacpan.org/pod/Future::AsyncAwait) CPAN module by
resolving the bugs that currently prevent it from being useful.

## Benefits to the Perl Community

This CPAN module implements the `async/await` syntax, which has emerged as a
common standard among various other languages. At the time of writing this
proposal, C#, Python 3, ECMAScript 6 (JavaScript), Dart, and Rust all implement
this syntax. It centres around the use of "futures" or "promises" as containers
of values to be provided by some possibly-asynchronous or background work, and
greatly improves on the readability of programs written using them. Logical
flow reads very similarly to standard synchronous call-style.

The examples in the CPAN module should help to illustrate the readability
advantages, as do other examples written in the other languages mentioned
above. There can be little doubt that `async/await` is gaining traction as the
standard syntax for writing control-flow around future-based concurrency across
many diverse programming languages.

By using this particular syntax, Perl 5 gains the same semantics with the same
syntax as used by these other languages. This helps readers to use the same
concepts as they may already be familiar with from using those languages.

## Deliverables

The primary deliverable would be a newer version of [Future::AsyncAwait](https://metacpan.org/pod/Future::AsyncAwait) whose
internal implementation is robust and fully-functional in a variety of test
cases and conditions.

A number of known bugs already exist and are collected on RT (at
[https://rt.cpan.org/Dist/Display.html?Queue=Future-AsyncAwait](https://rt.cpan.org/Dist/Display.html?Queue=Future-AsyncAwait)). Not all of
the tickets in this queue are necessary for completion of this project, as
some are "wish-list" items or design discussions about possible further
extensions or additions.

In addition to this CPAN release, a secondary deliverable will be a set of
blog posts and presentations to educate potential users - bringing the new
syntax to their attention, and demonstrating its use. Some blog posts may draw
contrasts and similarities to other languages, helping to further emphasize the
close relationship between the syntax across these languages. In particular,
parallels to Python 3 and ECMAScript 6 may be useful, as both of those
languages are ones that potential users may also be familiar with.

## Project Details

In its current state, the CPAN module implements a perl parser plugin that
parses the newly-defined keywords, and provides implementation semantics that
work in many simple test cases. The distribution contains a selection of unit
test files that demonstrate these cases.

However, the implementation of the suspend/resume functionality behind the
keywords is as yet insufficient for many larger, real-world applications to use
in existing code. This is due to the incomplete understanding of some of the
internal details of the `perl` interpreter, which causes some situations to
end up in a mismatched state, making it crash.

Central to this project's success is performing further research and
understanding on these internals. Once the operation of the interpreter is
better understood, the shortcomings of the current implementation of
[Future::AsyncAwait](https://metacpan.org/pod/Future::AsyncAwait) can be identified and improved upon. This may yield
further details that need to be understood which have yet to come to light,
but with any luck the process will eventually terminate yielding a new syntax
extension that is robust and useful.

## Inch-stones

Initial progress will be made from a better understanding of certain details of
the `perl` interpreter, primarily on the subject of the interaction between
`JMPENV_PUSH` and `PL_top_env`. A discussion of the topic can be seen here:
[https://www.nntp.perl.org/group/perl.perl5.porters/2018/08/msg251864.html](https://www.nntp.perl.org/group/perl.perl5.porters/2018/08/msg251864.html).

It is plausible there may be other subjects that arise once this one is fixed,
so the project may continue to switch between research and implementation
phases, as bugfixes uncover new areas that don't yet work.

There are particular bugs on the RT queue that need fixing; these are:

- [https://rt.cpan.org/Ticket/Display.html?id=126037](https://rt.cpan.org/Ticket/Display.html?id=126037)
- [https://rt.cpan.org/Ticket/Display.html?id=126036](https://rt.cpan.org/Ticket/Display.html?id=126036)
- [https://rt.cpan.org/Ticket/Display.html?id=125613](https://rt.cpan.org/Ticket/Display.html?id=125613)
- [https://rt.cpan.org/Ticket/Display.html?id=123062](https://rt.cpan.org/Ticket/Display.html?id=123062)

Additionally to these, to support the secondary goals of developer education,
a number of blog posts can be written:

- A general introduction to the `async/await` syntax as provided by
[Future::AsyncAwait](https://metacpan.org/pod/Future::AsyncAwait). Suitable for all Perl audiences.
- A comparison between [Future::AsyncAwait](https://metacpan.org/pod/Future::AsyncAwait) and similar syntax provided by
Python 3. This will serve as a useful comparison between the two languages, and
may help to draw some existing Python developers. Where possible it should
include some "Perlish rewrites" of some standard Python documentation, to
drive home that similarity.
- A comparison between [Future::AsyncAwait](https://metacpan.org/pod/Future::AsyncAwait) and ECMAScript 6, in similar style
to above, for similar reasons.

Thirdly, a presentation for a Perl conference can be prepared to demonstrate
the new language feature - perhaps as a follow-up to the talk I did for TPCiA
in 2017 ([https://www.youtube.com/watch?v=Xf7rStpNaT0](https://www.youtube.com/watch?v=Xf7rStpNaT0)), except this time to
report on some real-world success stories of the feature being used in
production cases.

Finally, the details of the perl interpreter discovered along the way can be
better documented, so that even if ultimately it proves impossible to fix the
implementation of [Future::AsyncAwait](https://metacpan.org/pod/Future::AsyncAwait), at least others can benefit from the
new understanding discovered in that research, which may be helpful to other
similar projects.

## Project Schedule

There are two main phases to this project: research into the existing operation
of the `perl` interpreter, and fixing the existing implementation of
suspend/resume semantics as a result of the discoveries made.

The first phase - research - should be achievable with around 10 days of effort
through a combination of code review and instrumented builds, resulting in a
textual description of the relevant "moving parts" of the interpreter. This
includes time to write down documentation of the discoveries.

Based on this understanding, the second phase - fixing the implementation - can
take place. This is likely to take a further 10 days including building more
unit test files to cover the new test cases.

In addition, the three blog posts and conference presentation will take around
4 days to prepare.

This brings the total to 24 days.

I'm currently a self-employed contractor with approximately three weeks each
month taken by existing clients. Working at a rate of 5 days per month, I
believe the project will come to a conclusion after around 5 months.

I intend to work at a rate of 5 days per month so regular checkpoints can be
established.

## Completeness Criteria

At time of writing there are no known CPAN modules able to use
[Future::AsyncAwait](https://metacpan.org/pod/Future::AsyncAwait) because it is too unstable at present for even the
smallest of unit-tests in other code to successfully pass. A useful moment to
consider as a completion target may be when the module is sufficiently robust
that it can be used as a dependency by at least a few other CPAN modules that
currently use [Future](https://metacpan.org/pod/Future).

For example, any of my [Device::Chip](https://metacpan.org/pod/Device::Chip) driver modules, that are currently
heavily [Future](https://metacpan.org/pod/Future)-based at present, would greatly benefit using this new
syntax.

CPAN shows a great number of other modules using [Future](https://metacpan.org/pod/Future) by many authors,
and any of these could also be used as test-cases for the robustness of
[Future::AsyncAwait](https://metacpan.org/pod/Future::AsyncAwait):

[https://metacpan.org/requires/distribution/Future?sort=\[\[2,1\]\]](https://metacpan.org/requires/distribution/Future?sort=[[2,1]])

## Bio

I am Paul Evans, PEVANS on CPAN ([https://metacpan.org/author/PEVANS](https://metacpan.org/author/PEVANS)).

I have been a CPAN maintainer for over 10 years, and currently have 155
distributions under my name. Among this set of modules are a number of
dual-life core modules - [List::Util](https://metacpan.org/pod/List::Util) and [IO::Socket::IP](https://metacpan.org/pod/IO::Socket::IP) being two that
may be among the most heavily-depended upon on CPAN. I have a number of
XS-based modules, including some such as [Syntax::Keyword::Try](https://metacpan.org/pod/Syntax::Keyword::Try) that provide
keyword plugins to extend the Perl syntax. I am familiar with many parts of the
core perl interpreter, and am well-known to many of the perl5-porters group.

I maintain a blog on a variety of programming topics, often posting on
Perl-related matters. [http://leonerds-code.blogspot.com/search/label/perl](http://leonerds-code.blogspot.com/search/label/perl).

I have spoken at most London Perl Workshops in the past few years, and attend
(and sometimes talk at) the European occurance of what was formerly called
YAPC::EU, most recently called TPCiG. I maintain a YouTube playlist of
recordings of talks I have given.
[https://www.youtube.com/playlist?list=PL9-uV\_AVx5FOzWJIvpuebmyiIuNd4L7GJ](https://www.youtube.com/playlist?list=PL9-uV_AVx5FOzWJIvpuebmyiIuNd4L7GJ)


