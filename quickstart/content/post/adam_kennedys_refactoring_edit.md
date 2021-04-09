
---
title: "Adam Kennedy's Refactoring Editor Grant"
author: Curtis "Ovid" Poe
type: post
date: 2007-03-25 00:00:00 +0000 UTC
url: "/post/adam_kennedys_refactoring_edit"
categories:
 - Grants

---

<p>First off, let's get the bad news out of the way.  Here's Adam's final report on the PPI refactoring editor:</p>

<hr/>
<p>It is with some sadness I think I will have to finally draw a line  under
this effort and consider it over. Without the appearance of a  pure-Perl
cross-platform (most likely Wx) programmer's text editor  on the CPAN, I don't
see how I can complete the proof of concept.</p>

<p>The entire grant was based on the availability of such an editor,  and the
administration overhead of getting it to the CPAN has meant  that hasn't
happened.</p>

<p>In summary, I believe most or all the secondary goals have been  completed
(caching support, various bits and pieces of PPI  functionality, the rewriting
of File::HomeDir, the creation of  File::ShareDir, Module::Install::Share,
File::UserConfig and a few  other modules, and helping out with various
PPI/Perl::Critic things).</p>

<p>The primary goal has failed (proof of concept refactoring Perl editor).</p>

<p>An unexpected additional goal has been completed, with the creation  of
Vanilla/Strawberry Perl Win32 distributions, bug-fixing a vast  number of CPAN
and core modules to make them work on Win32, and the  creation of
<a href="http://win32.perl.org/">win32.perl.org</a>.</p>

<p>Adam K</p>
<hr/>

<p>He was to be paid $5,000US for this grant, with $2,500 up front and the rest due on completion.  Given that he didn't complete his grant, some might be surprised that we've decided to make the final payment any.  What follows is why.</p>

<p>When I took over the role of grant committee secretary, one of the first things I did was a review of our payment policies.  We needed to balance the interests of the Perl community with the needs of the applicants.  One of the ideas which was floated was a complicated system of milestones that have to be achieved for each payment of a grant to be approved.  While this is the route that we have for the Parrot grant process, normal grants are such small amounts and the burden and costs of international transfers for folks outside of the US are such that a "multiple milestone" process quickly becomes very unwieldy.</p>

<p>As a result, we pay half up front and half on completion.  But what is "completion"?  In one grant process I've reviewed (not a TPF Grant Committee process, I might add), I watched as much concern was expressed over vague grants getting awarded for poor code because the grants were vague enough that "success" was easy to claim.  Other grants, with clear milestones, excellent code and well-defined goals were at risk of being denied because they were defined so rigidly that success is hard to achieve.</p>

<p>We don't want this to happen at TPF.  If people produce good work and achieve a reasonable substitute for their original goal, we consider -- on a case-by-case basis -- paying all or some of their remaining grant award.  If they <em>don't</em> produce work, we won't pay them.  It's that simple.  Of course, "a reasonable substitute" can be hard to define.  In Adam's case, the amount of work he has put in to make Win32 Perl solid and much easier to use in many cases.  Further, the tremendous number of bug fixes he's supplied to make core modules work on Win32, clearly the world's most widely used operating system, is a huge win for the Perl community.</p>

<p>Given that the Grant Committee tries to focus on grants with the widest benefit to the Perl community -- a goal we freely admit we don't always reach -- we quite likely would have approved a grant for the work Adam has accomplished and, as a result, Adam will be paid the final payment.</p>

<p>I strongly encourage all Win32 Perl users to check out <a href="http://vanillaperl.com/">the Vanilla Perl home page</a>.  It's mission is simple:</p>

<blockquote>The Vanilla Perl Project is an experiment to provide binary Perl distributions for the Microsoft Windows