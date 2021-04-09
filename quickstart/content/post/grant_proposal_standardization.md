
---
title: "Grant Proposal: Standardization, Test Coverage, and Documentation of Perl 6 I/O Routines"
author: Makoto Nozaki
type: post
date: 2017-01-17 00:00:00 +0000 UTC
url: "/post/grant_proposal_standardization"
categories:
 - Grants
 - Sign in

---

The Grants Committee has received one grant proposal **Standardization, Test Coverage, and Documentation of Perl 6 I/O Routines** for the January round.  Before the Committee members vote, we would like to solicit feedback from the Perl community on the proposal.

Review the proposal below and please comment here by January 23rd, 2017.  The Committee members will start the voting process following that and the conclusion will be announced approximately in one week.

# Standardization, Test Coverage, and Documentation of Perl 6 I/O Routines

- Name:

    Zoffix Znet

- Amount Requested:

    USD 999

## Synopsis

Many of Perl 6's I/O routines currently have inconsistencies in behaviour
between method and subroutine forms along with inconsistencies in calling
forms and failure modes compared to the rest of the language. Also some of
them, despite being speculated in the synopses and implemented in Rakudo, have
no test coverage in the Perl 6 Specification, and as such they remain
undocumented and not part of the Perl 6 Language. Some of the routines that do
have presence in the Specification, have sparse test coverage, leaving
some of their functionality entirely untested.

The work funded by this grant seeks to resolve those inconsistencies, as well
as provide full test coverage and documentation.

## Benefits to the Perl Community

Inconsistent interfaces are difficult to master and are bug-prone due to
programmers forgetting about varying details of individual routines. By making
I/O routines take arguments and indicate failures consistent to the way it's
done in the rest of the language, they will be much easier for programmers to
learn and use. Along with a more pleasant interface, the Perl Community will
also benefit by needing to provide less support for people trying to learn the
language.

The largest benefit will come from the full test coverage that will no doubt
reveal bugs that can be fixed before they're encountered by users in
production code. The coverage will also protect from future bugs being
introduced inadvertently. Also, the currently unspecced routines will be
officially part of the language, and so will be documented and available for
use.

### __As bonus benefits:__

The timing of this work seeks to align with the production of "Learning Perl 6"
book, in order for it to contain valid information on I/O routines and
describe only routines that are actually part of the language, which will
avoid confusion for its readers and folks helping those readers learn Perl 6.

Lastly, the bonus deliverables (described below) will benefit further
development of Rakudo/Perl 6 by elucidating unspecced, unused, or unwanted
routines. The data can also be used to produce a teaching/reference aid
(e.g. flash cards with routine names or IRC bot for routine reference).




## Deliverables

__Scope:__ by _I/O routines_ this grant proposal means subroutines and
methods of the [Rakudo implementation](https://github.com/rakudo/rakudo/)
found in `src/core/io_operators.pm`,
`src/core/IO.pm`, `src/core/IO/ArgFiles.pm`, `src/core/IO/Handle.pm`,
`src/core/IO/Notification.pm`, `src/core/IO/Path.pm`, `src/core/IO/Pipe.pm`,
`src/core/IO/Special.pm`, and `src/core/IO/Spec.pm` as well as its
subclasses found in `src/core/IO/Spec/` directory. Note that `IO::Socket` and
its subclasses are NOT in the scope of this grant proposal.

Deliverables:

- I/O Action Plan Report
- Commits in [rakudo](https://github.com/rakudo/rakudo/) implementing
the Action Plan.
- Commits in [Perl 6 Specification](https://github.com/perl6/roast)
providing full coverage for I/O routines.
- Commits in [Perl 6 Documentation](https://github.com/perl6/doc)
providing full coverage for undocumented I/O routines, as well as any
corrections for existing I/O documentation.
- Bonus deliverable: The "Map" of Rakudo Routines
- Bonus objective: whenever possible, I intend to fix any of the bugs
found by the new test coverage. Should a bug prove to be hard to fix,
the test exposing it will be fudged and a ticket for the bug will be filed.

## Project Details

- (Bonus deliverable): __The "Map" of Rakudo Routines__
This will be semi-automatically generated (by introspection) list of all
subroutines and publicly accessible object methods provided by Rakudo
implementation, together with information on what calling convention they
use and how they fail (e.g. returning some object, throwing, or returning
a `Failure` object). This data will guide the decision on how the calling convetion/failure mode for I/O routines should be standardized. I intend
to work on this deliverable while this grant proposal is deliberated. As
such, it's a bonus deliverable and will be completed regardless of whether
this grant is approved.
- __I/O Action Plan Report__
This will be a Markdown document detailing how the existing routines will
change, what the new routines will do, if any are to be added, and whether
any currently unspecced routines are to be removed. Parts of this document
will eventually be re-used as documention for currently undocumented
routines.
    - The scope of all of the changes will be limited by the _6.c_ Perl 6
    language specification, so none of the changes would have to be deferred
    until _6.d_ language release.
    - This document will be placed into
    [rakudo](https://github.com/rakudo/rakudo/) repository when ready and
    other core developers will be asked to reivew it and make comments on the
    proposed changes.
- __Commits__
The commits to rakudo/roast/doc repositories will implement changes in
the Action Plan, and test and document any I/O routines that are still
left untested/undocumented after those changes.

## Inch-stones

These are the inch-stones I intend to follow:

### Routine Map

- Use Perl 6's introspection facilities to locate core subroutines and
classes in `CORE::` lexical scope. Further introspect those classes to obtain
`Method` objects. Further introspect the candidates and signatures of all of
these routines to obtain the calling convention list (e.g. we pass arguments
as adverbs or as positionals, etc)
- Further introspect the arguments and programmatically invoke the
routines with some probably-acceptable values to obtain the output range, in
particular failure modes. Some heuristics or manual massaging will be needed
here.
- Part of this work already exists in [a repo](https://github.com/zoffixznet/perl6-map). If the grant is approved, this repository will be
transfered to [Perl 6 GitHub organization](https://github.com/perl6/)

### I/O Action Plan Report

- Using the routine map, figure out a better interface for I/O routines
and then see what can be changed without breaking the 6.c language spec. Even
without that map, I already can see some targets for change: the routines that
take `$test` named arg as a predefined string of test letters in particular
order, as well as the inconsistency that most methods `fail()` while most
subroutines `throw()`. Based on a brief [conversation with
Larry Wall](https://irclog.perlgeek.de/perl6-dev/2016-12-28#i_13814034), the
likely change would be to make all of them consistently `fail()`.
- Inspect the [code of NewIO branch](https://github.com/rakudo/rakudo/compare/6bbb56f4c598ba0fef49ba9b11671df675019366...newio)
for any ideas that can be salvaged for 6.c Language. The work in this branch
was created by lizmat++ in 2014-2015, but due to unfortunate circumstances
never got merged to master prior to 6.c release. My understanding is it offers
a lot of improvements to our current IO, but based on brief perusal of the
code, I suspect it would have a lot of conflict with 6.c Language
Specification tests. Under this grant, I would like to see if any of the
improvements that work offers can be applied to current 6.c language.
- Using the [MoarVM coverage report
tool](http://perl6.party/post/Perl-6-Core-Hacking-Can-Has-Moar-Cover) (and by
grepping the [roast repo](https://github.com/perl6/roast/)), find
routines that are entirely unspecced and decide whether they
should stay or be removed. For example, the currently unspecced `indir()`
routine is a likely candidate for staying, while I've seen some calls on IRC
to remove the unspecced `tmpdir()`.
- I forcee the largest part of this grant work to be the writing of
the tests. Here are some examples of the rough current coverage state of IO
routines:
[src/core/io\_operators.pm](http://perl6.wtf/src_core_io_operators.pm.coverage.html),
[IO::Argfiles](http://perl6.wtf/src_core_IO_ArgFiles.pm.coverage.html),
[IO::Special](http://perl6.wtf/src_core_IO_Special.pm.coverage.html),
[IO::Spec::Win32](http://perl6.wtf/src_core_IO_Spec_Win32.pm.coverage.html),
and [IO::Handle](http://perl6.wtf/src_core_IO_Handle.pm.coverage.html). The
routines whose names are on red lines indicate they're untested in roast
(this report was generated on Oct 7 and as I recall only IO::ArgFiles.lines and
IO::Handle.seek received any further tests since then).
- A lot of the decisions made by the Action Plan will be able to close
many of the tickets opened by _brian d foy_ while gathering info on our
I/O routines that is going to be included in _Learning Perl 6_ book and the
concerns raised in the tickets will be used as input for the Action Plan as
well. (Some of them are:
[RT#130460](https://rt.perl.org/Public/Bug/Display.html?id=130460),
[RT#130456](https://rt.perl.org/Public/Bug/Display.html?id=130456),
[RT#130454](https://rt.perl.org/Public/Bug/Display.html?id=130454),
[RT#130455](https://rt.perl.org/Public/Bug/Display.html?id=130455),
[RT#130489](https://rt.perl.org/Public/Bug/Display.html?id=130489),
[RT#130490](https://rt.perl.org/Public/Bug/Display.html?id=130490)).
- Write down the action plan in Markdown format, make it available
in Rakudo repository, and invite other core members in `#perl6-dev` IRC
channel to comment on it.
- After a 1 week review period, update the Action Plan to reflect
any received feedback, and proceed to implement it.

### Commits

The commits will implement the Action Plan. Roast commits will be based on
the changes to routines as well as the report generated by the coverage tool.
And the doc commits will be done by manually searching and reading exiting
documentation and amending it as needed.

(N.B.: I'm aware the coverage tool is currently busted by a commit that changed
filenames in `.file` method for core routines; however MasterDuke++ promised
to fix it, and if they won't be able to find time to do so, the fix should be
simple enough that I'd fix the tool myself).

## Project Schedule

I already began work on the routine map generator and will complete it by
the time the decision on this grant proposal is available. After that,
I expect to spend 2 weeks preparing the I/O Action Plan Report, 1 week for its
review by other core members, and 2 weeks for its full implementation
(including tests and docs). I also allow for extra 2 weeks for any unforseen
delays in any of the steps.

If my understanding of [the date when the decision on this grant would be
ready](http://news.perlfoundation.org/2017/01/call-for-grant-proposals-jan-2.html) is correct, I intend to finish the work before the end of March, 2017.

## Completeness Criteria

- [rakudo repository](https://github.com/rakudo/rakudo/tree/nom/docs)
will contain the IO Action Plan document and it will be fully implemented.
- All of the I/O routines will have tests in
[roast](https://github.com/perl6/roast/) and documented on
[docs.perl6.org](https://docs.perl6.org). If any of the currently implemented
but unspecced routines are decided against being included in Perl 6 Language,
their implementation will no longer be available in Rakudo.
- The test coverage tool will report all I/O routines as covered and
the information will be visible on [perl6.wtf](http://perl6.wtf) (_Perl 6's
Wonderful Test Files_) website. Note: due to current experimental status of the
coverage tool, its report may still show some lines or conditionals untested
despite them actually being tested; however, it _will_ show the lines where
routines' names are specified as covered.

## Sidenote Comments

The grant amount requested may be low compared to the described amount of work
because I'd like to still view part of my time on this work as donated to
Perl 6. Those who see me on IRC may notice I bounced around the ideas in this
proposal before.

The grant will let me finish this work much sooner and in a more complete
state than I would be able to otherwise.

## Bio

I'm a 30-year old Canadian who lived near Toronto for the past 14 years.
I also spent some years of my life living in USA (New Jersey) and, before that,
in Russia (Siberia), where I was born.

I started with Perl 5 about 12 years ago and have since released over
200 Perl 5 modules on CPAN and for the last 10 years held a single job,
large part of which is web development with Perl 5.

In the fall of 2015, I switched my focus to Perl 6 and to date released 34
Perl 6 modules and delivered a couple of Perl 6 presentations at the Toronto
Perl Mongers meetings.

Around July, 2016 I joined the Rakudo Perl 6 Core Development Team. I also have
been Rakudo's release manager every month since the 2016.06 release.

My notable deliverables to the Perl 6 Community involve the creation of the
web app driving [modules.perl6.org](https://modules.perl6.org); nearly total
automation of the Rakudo's release process, including the development of
[perl6.fail](http://perl6.fail) web app for RT interfacing and release status
tracking; and writing all of the tutorials on
[perl6.party](http://perl6.party) website.

To date, I have authored __461 commits__ to
[Rakudo](https://github.com/rakudo/rakudo/) and __1,823 commits__ to
repositories in [Perl 6 GitHub organization](https://github.com/perl6/)
(__523__ of which have been to the
[Perl 6 Specification](https://github.com/perl6/roast/) repository).


