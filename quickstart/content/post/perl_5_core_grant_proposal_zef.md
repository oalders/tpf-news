
---
title: "Perl 5 Core Grant Proposal: Zefram maintaining the Perl 5 core"
author: Makoto Nozaki
type: post
date: 2017-06-04 00:00:00 +0000 UTC
url: "/post/perl_5_core_grant_proposal_zef"
categories:
 - Grants
 - Perl 5 Development
 - Sign in

---

TPF Board has received a new grant application under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund).

Before we vote on this proposal, we would like to get feedback from the Perl community. Please leave feedback in the comments or, if you prefer, email your comments to makoto at perlfoundation.org.

**IRC nickname:** Zefram 

**project title:** Zefram maintaining the Perl 5 core 

###synopsis

I'd like a grant to work on the Perl 5 core, concentrating on 
tricky and obscure issues of the core internals. The grant would extend 
to the full range of a core committer's activities, along the lines of 
my past work on the core. The grant would allow me to put much more 
time into the Perl 5 core than I have been able to recently. 

###benefits to Perl 5 core maintenance

this grant would overcome the present 
limitation on the amount of time I can devote to Perl, namely my need for 
an income. The core would benefit from the increased application of my 
rare knowledge and skills, including my deep understanding of much of 
the core internals. It would especially benefit through the fixing of 
troublesome bugs, and more generally through the addressing of issues 
that require in-depth familiarity with the internals. 

###deliverable elements

Like the similar grants for Dave Mitchell, Nicholas Clark, and Tony 
Cook, I do not propose specific technical deliverables for this project. 
The details of what work I tackle would be led by the exigencies of bug 
reports and the direction of discussions among core developers (on the 
public mailing list, the IRC channel, and at times on the closed security 
report mailing list). The main deliverable, then, is time spent on core 
maintenance work, at a rate of USD 50 per hour. 

Secondary deliverables are the required reports, weekly to p5p and 
monthly to TPF. These will list the issues worked on. The grant manager 
representing p5p can see how this relates to my code commits and mailing 
list activity, providing an opportunity to raise any concerns. 

###project details

This project covers the full range of core maintenance activity that I 
have performed in the past. Within this mix, I would favour activity 
of which the fewest people are capable: work that takes advantage of my 
particular knowledge of core internals and my particular abilities. 
In general, the highest priority work would be the diagnosis and 
resolution of bugs that appear to involve obscure internals. (Note that 
diagnosis can radically change the appearance of a bug in this respect.) 

Activities covered by this grant would include, but are not limited to: 

* diagnosing reported bugs 
* fixing bugs, whether reported or not 
* discussing core issues on the p5p mailing list 
* writing and improving documentation 
* code refactoring 
* release engineering 
* implementation of new core features, where approved through discussion 

Specifically excluded from this project scope is work on unapproved new 
features, because of the risk of (perception of) abuse. I do not seek 
a license to arbitrarily perform major speculative work on TPF's dime. 
But where a new feature has been discussed prior to implementation, 
and there is the appropriate consensus that it should be implemented, 
then the endeavour is no longer an individual's speculation, and that 
implementation work can fall within the scope of this project. Looking 
back at features I have added in the past, I developed call-checker 
and parser plugins without discussion, so they would not have been in 
scope for this type of grant. Subroutine signatures, on the other hand, 
I developed in response to a fairly detailed consensus reached on the 
IRC channel, and with the pumpking's specific blessing, so that would 
have been in scope. 

###project schedule

This project c