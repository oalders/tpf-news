
---
title: "Devel::Cover Grant Report for August"
author: mdk
type: post
date: 2012-09-13 00:00:00 +0000 UTC
url: "/post/develcover_grant_report_for_au"
categories:

---

Paul Johnson writes:

In accordance with the terms of my grant from TPF this is the monthly
report for my work on improving Devel::Cover covering August 2012.

This month I released Devel::Cover 0.93.

The bulk of this report is taken from my weekly reports, so if you have
read them there is little new here.

One of the nice things about having this grant and being able to spend
more time working on Devel::Cover than I would otherwise have been able
to is that I'm getting feedback from my reports and checkins, and that I
have more time to be able to follow up on it.

Over the years I've received a lot of mail regarding Devel::Cover and
I've not always been able to properly reply to it all.  In some cases
I'm sure I've not replied at all.  If anyone has sent me something that
needs a reply, please feel free to remind me either by mail or,
preferably, on github if appropriate.

So I've started looking back at some of my outstanding mail (starting
with those where I have been prompted).  One of the first I looked at
was about a problem I have known about for many years.

When you want to know your coverage there are two main phases: the first
in which you exercise your code and the coverage data is collected, and
the second in which that data is displayed, hopefully in a format which
makes it easy to understand.  In each phase you can limit the amount of
coverage information both by limiting the files for which you collect
data and by limiting the criteria.

In general it is best to collect as little data as possible in the first
phase.  This will reduce the size of the coverage database but, more
importantly, it will make the running of the tests faster.  But there
are times when, having collected the data, you only want to display a
subset of it.  Thus you can also filter the data when you generate your
report.

When the report is generated, by default, a summary of the coverage is
printed.  This summary has always been a summary of all the data in the
database, rather than a summary of the data in the report.  I remember
looking into this many years ago and deciding that it was going to be
complicated to fix and this is probably why I've never seriously looked
at the problem up to now.  But now I have, and the fix wasn't as
complicated as I had imagined that it would be.  There's probably a
lesson of some sort there.

I also took a look at RT 49916 which relates to file and directory
permissions with respect to code that changes EUID.  I think this is
fixed but I'm waiting to hear back before closing the ticket.

I had a look into a cpantesters report on bleadperl which pointed back
to perl #113684, which is a report I've written about before.  Father
Chrysostomos has made some commits to address the problem and this has
lead to a minor change in B::Deparse which affects Devel::Cover.  This
is somewhat dictated by the limitations of B::Deparse.  The full details
are in the RT thread, but the long and short of it all is that these
changes are here to stay and so I will adjust Devel::Cover accordingly
as soon as 5.17.3 is released.

I've also been running cpancover jobs every so often.  This takes quite
a time to finish but I like to think of CPAN and cpancover as my
extended test suite.  When I see something that doesn't look right I can
then go in and investigate in a bit more detail.  I started running
cpancover with perl 5.16.1-RC1 to see how it pans out.  See the results at
"http://cpancover.com":http://cpancover.com

At the QA hackathon this year I put together a Vim plugin to display
coverage information within Vim.  I really like this plugin because I
can see my coverage information right where I am editing my code, but I
noticed that this plugin wasn't honouring the coverage criteria it was
given, so I fixed that up too.

Finally, I started looking into RT 69240.  Initially, this looked like
it should be a reasonably simple problem, but when I eventually
unearthed a CPAN module where I could reproduce the problem it slowly
became clear that I was on the trail of a longstanding problem that I
had known about for many years but had never been able to reliably
reproduce, and that I had never had a sufficient chunk of time to be
able to track down.

Now, thanks to this grant, I do.  This is exactly the sort of problem I
had hoped to be able to solve with this grant.  I have started off by
trying to catch up on the backlog of bugs that have been reported, but
then problems such as this are the priority before, with luck, moving on
to new features.

As is often the case, all the work was in finding the problem.  Once
found, the solutions was all but trivial.

The problem was in the heart of probably the most complicated part of
Devel::Cover.  That is the code which tries to manage the dynamic nature
of Perl.

Code coverage in a static language such as C is relatively
straightforward.  Source code is compiled to object code and that object
code doesn't change.  (Self-modifying code is an illusion.  It doesn't
exist.)  Everything you need to know about the code is known before a
line of it has been executed.

In Perl, and similar languages, the separation of compile time and run
time is not so clear cut.  Code being executed can get perl to compile
new code into already existing modules.  This poses challenges to a tool
such as Devel::Cover which tries to collect information about the code
being executed.

If you consider just subroutine coverage, in "normal", static code, I
can note the subroutines in a file and produce a mapping from the
subroutines to the position in the source code where they are defined,
producing an ordering on the subroutines.  When a subroutine is executed
I can note that the nth subroutine in the file is covered.

In dynamic code new subroutines can spring into existence whilst the
code is being executed.  This happens via a string eval in some guise.
When this happens I can tag the new sub onto the end of my list of subs,
and this works well.

But the problem becomes more difficult when in two different runs,
different subroutines are created.  A naïve solution here can lead to
subroutine $n meaning different subroutines in different runs.  So we
need to be clever and recognise when one of these new subroutines
matches an identical subroutine from a previous run, and when it has
never been seen before.

If all this functions correctly we should never get to a situation where
we have information that subroutine $n has been executed, but we only
know about n-1 subroutines.  If that situation does occur, we get the
"ignoring extra subroutine" message and coverage will be lost.

The bug existed in the code which managed how these lists of criteria
were maintained between runs.  There may still be bugs in this area, but
it was great to be able to knock this one over.  The actual commit was
"https://github.com/pjcj/Devel--Cover/commit/997426eecb16899d0be425853478e2b5bdf9a1ee":https://github.com/pjcj/Devel--Cover/commit/997426eecb16899d0be425853478e2b5bdf9a1ee
if you want to see the simple fix to the hard-to-find problem.

I should probably note that there are other solutions to this problem.
Early versions of Devel::Cover stored the location of each construct
together with the information about its coverage.  This does work well
but is very expensive on storage requirements and the CPU required
manage this extra data.  Coverage always has an overhead and the greater
the overhead the less people will be inclined to use it, so I try hard
to keep the overhead to a minimum.

Whilst tracking down and fixing this problem I also fixed more than ten
other bits and bobs that I noticed, as well as other peripheral matters.
The full details are in the commits.  A couple of those bits and bobs
are probably quite important, actually.

Oh, And then I did fix up the remainder of the problem in the original
bug report.

I also got a message from Nuno Carvalho who is packaging Devel::Cover
for Debian.  It seems that PodVersion isn't happy when the module
description is longer than one line.  Since that's not a very good idea
anyway I fixed up the affected modules so with luck the Debian packaging
can go ahead.

Finally, I fixed up Devel::Cover to work around mod_perl2 setting $^X
to httpd.  The mod_perl folk are also going to fix that, hopefully for
2.0.8.

As usual, I made other various fixes and updates.

The work I have completed in the time covered by this report is:

Closed RT tickets:

* 68517 summary, report total from cover tool includes ignored files

* 77818 tests fail due to spaces in @INC (Devel::Cover::Inc issue)

Fixed cpantesters reports:

* "http://www.cpantesters.org/cpan/report/a7ec86de-db7e-11e1-97c0-97aef2681f56":http://www.cpantesters.org/cpan/report/a7ec86de-db7e-11e1-97c0-97aef2681f56

You can see the commits at "https://github.com/pjcj/Devel--Cover/commits/master
":https://github.com/pjcj/Devel--Cover/commits/master

* Hours worked: 42:55
* Total hours works on grant: 131:45


#devel::cover #paul-johnson
