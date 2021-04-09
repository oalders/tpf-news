
---
title: "Hague Grant: Traits, Introspection and More Dispatch Work For Rakudo - Final Report"
author: Karen Pauley
type: post
date: 2009-09-05 00:00:00 +0000 UTC
url: "/post/hague_grant_traits_introspecti"
categories:
 - Grants

---

_Jonathan Worthington writes:_

**Introduction**

bq.   If it weren't for the fact that you said you were going to do it,
  I would've doubted it could be done at all in current Parrot.
    -- pmichaud

In many senses, this "Hague Grant":http://news.perlfoundation.org/2009/04/hague_perl_6_grant_request_tra.html was rather more challenging than my "last one":http://news.perlfoundation.org/2008/11/hague_grant_request_rakudo_dis.html. In the course of it, I've been dealing with major changes to some of the things most fundamental to any Perl 6 implementation, including method dispatch, trait application and metaclasses. Happily, I have been able to meet the challenges and complete all of the deliverables laid out in the grant application, and had no shortage of fun in doing so.

**Deliverables Status**

I proposed six deliverables. Here I will detail what was achieved in each of them.

D1. This deliverable was by far the most challenging of the grant. At the heart of it was completely replacing the existing method dispatcher, which allowed us to implement many features but was never going to be sufficiently performant, mapped onto Parrot's primitives poorly and did not provide a good way to support deferral.

The new dispatcher gave support for deferral from the start. Furthermore, it computes the list of candidates we could defer to lazily, so as to limit the cost to those methods that choose to defer. Since we had never supported the $foo.@candidates syntax, I was able to partly introduce it there, but there was no getting away from eventually needing a large patch that updated all other method calls forms to work with the new dispatcher. I was happy that the landing of this patch caused very little disruption to Rakudo users, despite it marking a very large change to the way something very fundamental to the language worked. As an additional bonus, replacing the dispatcher forced various corner cases to be explored, and shook out a lot of places where we had been cheating.

All of the promised trimmings on the side of the dispatcher replacement were also completed: .WALK, .can, lastcall, the hides trait modifier and the hidden trait. Additionally, dispatches for the form .+, .? and .* that take an indirect method or a list of methods are also now supported. All of these, and deferral, have had their test coverage reviewed and "improved":http://svn.pugscode.org/pugs/t/spec/S12-methods/.

D2. I refactored .wrap and .unwrap to work in terms of candidate lists. This also fixed some tests that the previous implementation had been failing.

D3. I implemented parallel dispatch and reviewed and extended "testing":http://svn.pugscode.org/pugs/t/spec/S12-methods/parallel-dispatch.t for this feature.

D4. I started off my work on improving Rakudo's introspection capabilities by reviewing, getting input on and then improving the specification itself to be more clear. I then worked on implementing these parts of the specification. Along the way I did a refactor to give Rakudo it's own metaclasses for roles and classes, which means we don't poke so much stuff into P6object in the Parrot namespace. The spectest suite was extremely lacking in test coverage for this area; I wrote over 150 new "tests":http://svn.pugscode.org/pugs/t/spec/S12-introspection/.

D5. I implemented support for defining methods on the metaclass, as specified in S12, and reviewed and enabled the tests for this.

D6. The traits part of the grant was something of an adventure. When I started asking for clarifications on various aspects of it during planning the implementation, it soon became clear that the design team wanted to re-visit some aspects of it. Shortly afterwards, some fairly sizable changes were made. Thus what I have implemented in Rakudo is a greatly updated specification. That needed a little extra refactoring, and s