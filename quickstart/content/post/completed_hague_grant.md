
---
title: "Completed Hague Grant"
author: Karen Pauley
type: post
date: 2010-03-03 00:00:00 +0000 UTC
url: "/post/completed_hague_grant"
categories:
 - Grants

---

Jonathan Worthington has successfully completed his Hague Grant for "Rakudo Signature Improvements":http://news.perlfoundation.org/2009/09/hague_grant_application_rakudo.html.  This is his final report:

**Introduction**

A while back, I filed an "application":http://news.perlfoundation.org/2009/09/hague_grant_application_rakudo.html for a grant to work on a wide range of improvements to Rakudo's signature, capture and binding implementation, in order to improve performance, bring us in line with the specification and implement a range of features that we did not have support for.

I have now completed the grant tasks, and am filing this report to discuss the work that was done under it. I have missed the originally planned schedule for the grant to be completed. This was in a large part due to a lot of Rakudo development since October taking place in a branch, "ng":http://use.perl.org/~pmichaud/journal/39874, which set out to refactor and improve many aspects of Rakudo and get us in good shape for the Rakudo * release in a couple of months time. It made far more sense to implement some aspects of this grant in the branch rather than in master, and the branch also needed my attention in other key areas. The branch has recently become master, and since then I've been able to complete the final pieces of the grant and bring it to completion.

**Deliverables Status**

I proposed six deliverables, and have completed all of them.

D1 called for an extensive refactoring of signature internals and parameter binding. Rakudo now has a single custom signature binder, designed to handle its binding needs in accordance with the Perl 6 specification. This is a far better situation than before, where we used the Parrot binder and then layered extra binding semantics atop of it, meaning we had to make two passes through the parameters.

Having our own custom binder that takes into account the needs of Perl 6 has made it possible to do all of the other items in this grant in a clean and performant way. It has also enabled other bits of progress outside the scope of the grant; for example, I was able to get a first cut of coercion (the "as" trait) of parameters in signatures implemented in little over half an hour a few days ago.

We also have compile-time objects to represent signatures and parameters in the compiler. This has allowed much cleaner code in the actions that build the AST. In the future, this approach will also be highly useful for building extra static checking and an optimizer.

D2 required implementing support for binding named arguments to positional parameters. I considered this in the initial re-design of the binder, so it just worked "out of the box" once I switched Rakudo over to using it.

The target of D3 - unpacking of arguments - was made much easier by having a custom binder. The binder was designed to be re-entrant, and therefore unpacking arguments was implemented by specifying how various data structures could coerce themselves into a capture, and then making a recursive call with the capture that was produced.

For D4, I spent some time re-working capture support overall, and as well as being able to write a capture into a signature - and to unpack it with a sub-signature - we can also interpolate a capture into the arguments for a call. Before this didn't really work - it would lose the named parameters, for example.

Implementing D5 - support for multiple return values - first meant working with Larry to figure out the correct semantics. It turned out they were easier to get in place in the refactored "ng" branch - which has now become the master branch. I also implemented the support for binding the return values against a signature; this means all the power of unpacking arguments is available for returns too. It's all routed through the same underlying binder implementation.

Finally, for D6 I first put a proposal into the specification f