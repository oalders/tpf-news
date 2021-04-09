
---
title: "Paul Johnson, Devel::Cover Grant Report for May 2012"
author: mdk
type: post
date: 2012-06-15 00:00:00 +0000 UTC
url: "/post/paul_johnson_develcover_grant"
categories:

---

In accordance with the terms of my grant from TPF this is the monthly
report for my work on improving Devel::Cover covering May 2012.

This report is essentially the same as my report for the first week, but
that will not be the usual case.  If you read that report you can safely
skip this one.

You may recall that recently I submitted a grant proposal to The Perl
Foundation to work on Devel::Cover.  I'd like to thank everyone who
thought that this might be a good idea, and I'm pleased to say that the
proposal was accepted and I have started work.  Modelled on the
successful grants of Dave and Nick, one of the conditions of the grant
is that I should produce weekly and monthly reports of my progress, and
these reports will be sent to "perl-qa@perl.org":mailto:perl-qa@perl.org.  If you are not
interested in these reports, please set up a filter, or just ignore
them.

I suppose there are two deliverables this month - Devel::Cover 0.87 and
"http://cpancover.com":http://cpancover.com

My plan is to work primarily on existing, relatively simple bugs and
problems before getting down to the complicated bugs and the more
interesting work of adding functionality, and most of my work so far has
indeed fallen into that category.  There have been a couple of
deviations from that - one planned and one not.  Much of the rest of my
work was related to Moose and Mouse.  I've had a report about Moo which
I need to look at, but I've heard nothing about Mo so I'm not sure if
there are any problems there.

Very shortly after I started work perl-5.16.0 was released.  So the second bug
that I looked at was RT 75314 which concerned a regression on condition cover
in 5.16.  Unfortunately this problem went to the core of one of the most
tricky parts of Devel::Cover and it took me quite some time to get to the
bottom of it.  The good news though, is that without this grant I would never
have been able to solve this problem and so condition and branch coverage in
5.16 would have been broken.

The bulk of the fix is found in f26ba32 and it relates to an
optimisation from David Mitchell:
"http://perl5.git.perl.org/perl.git/blobdiff/0e1b3a4b35c4f6798b244c5b82edcf759e9e6806..db4d68cf2dda3f17:/op.c":http://perl5.git.perl.org/perl.git/blobdiff/0e1b3a4b35c4f6798b244c5b82edcf759e9e6806..db4d68cf2dda3f17:/op.c

Nick's bisecting code in the perl core was very helpful in fingering
this commit but, really, I should have remembered it.  When Dave
announced this optimisation, I suspected that it would affect
Devel::Cover, but I had not fully followed up on that.  The problem for
Devel::Cover is that I was depending on that op ordering to properly
calculate condition coverage.

Devel::Cover has two ways of collecting condition coverage, depending on
whether or not the condition is short circuited.  If it is short
circuited, then when we get to the condition all we have to do is look
at the top of the stack to see whether the LHS is true or false.  If
not, we have to remember that we are interested in the value of the RHS
when it has been evaluated.

Devel::Cover does this by hijacking the op_next pointer to look at the
value of the RHS and collect its coverage before moving on to the
expected op.  The optimisation meant that we never evaluated the RHS and
so Devel::Cover didn't know its value.  So Devel::Cover fell back to its
default behaviour in such cases which is to assume that the RHS is
false, as it is in cases when we never return a value from the RHS.
This is somewhat simplified, but gives you the idea of what is going on.
(Incidently, this is also the primary reason that Devel::Cover doesn't
run under threads.)

The solution to this involves grovelling around the optree a little more
to find the conditional ops and, whilst it solves most of the problems,
there is still a little work to be done in this area.  (Where "a little"
means that the problem is simple, but the so