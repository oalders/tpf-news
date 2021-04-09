
---
title: "Hague Perl 6 Grant for Jonathan Worthington completed and accepted"
author: Richard Dice
type: post
date: 2009-04-28 00:00:00 +0000 UTC
url: "/post/hague_perl_6_grant_for_jonatha"
categories:
 - Grants
 - Perl 6 Development

---

In late November 2007 Jonathan Worthington submitted a "Hague Perl 6 Grant proposal":http://www.perlfoundation.org/ian_hague_perl_6_development_grants entitled "'Rakudo Dispatch and Role Enhancements'":http://news.perlfoundation.org/2008/11/hague_grant_request_rakudo_dis.html.  I am pleased to announce that Jonathan submitted his final Hague grant report yesterday (included below the fold), and it was approved and accepted by chromatic, the grant manager for the grant.  This successfully completes this Hague Grant for Perl 6 development.  We thank Jonathan for his great work on this grant and his efforts in advancing the goal of a completed Perl 6 implementation.


<big>**Introduction**</big>
In the months during which I have worked on my Hague Grant, Rakudo Perl 
6 has taken some huge steps forward - not just as a result of my work 
and this grant, but as the result of the efforts of a growing Rakudo 
Perl 6 user and developer community. When I submitted my grant proposal 
I noted that we passed over 4,000 tests. Today we pass over 10,000, are 
more stable and have a much wider feature set. Thus my work on the 
deliverables of this grant have been part of a wider scene of extensive 
Rakudo progress.

<big>**Deliverables Status**</big>
All deliverables have been achieved. In places, my work has gone beyond 
what the grant required. Reactions to my work have been positive, both 
from users and other developers.

**D1.** We now register types in the symbol table at compile time and the 
hack that saw us through the early days is removed. Furthermore, I have 
used this to build some other features, including compile time detection 
of type re-declaration.

**D2.** Junction auto-threading is now fully implemented for both single and 
multiple dispatch. The Perl 6 multi-dispatcher knows how to generate and 
cache junction dispatchers for future performance too. As various 
built-ins have moved over to the Perl 6 prelude, they too have gained 
the auto-threading automatically. I also reviewed and added a range of 
tests[1] to exercise the auto-threading, including ensuring that 
interaction with multis and auto-threading as well as nested junctions 
work as expected.

**D3.** I implemented submethods and reviewed and enabled the tests for them.

**D4.** The handles trait verb has been extensively refactored and now 
handles many more cases. I reviewed, enabled and added to the spectest 
suite tests[2] relating to this, to improve test coverage.

**D5.** This has by far been the most significant and complicated 
deliverable of the grant, and what has been achieved in this area 
exceeds what was originally required. First and foremost, we now have an 
implementation of parametric roles. They can be declared, composed and 
mixed in, the selection falling naturally out of the usual 
multi-dispatch semantics as is specified.

In addition to this, I have converted the Positional, Associative and 
Callable roles to be parametric. This, along with a little bit of extra 
work in the compiler, has now given us typed arrays, hashes and routines 
- features from S09 and S06. This has served as a great concrete test of 
parametric roles. It has also won us sigil-based multi-dispatch and 
multi-dispatch based upon typed data structures too.

Test coverage for parametric roles was weak (understandable, because 
nobody had implemented them before), and I have expanded this greatly 
during my work[3],[4],[5]. Some tricky areas - including parametric roles 
recursively parametrized and parametric role subtyping - a feature not 
in the specification at the time I started the grant - are now well 
exercised.

I rounded this work off by updating S14[6] (the specification for roles, 
which has been broken out from S12 since I started this grant). It is 
now more detailed and complete.

<big>**Dissemination**</