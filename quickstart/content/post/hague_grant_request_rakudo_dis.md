
---
title: "Hague Grant request, \"Rakudo Dispatch and Role Enhancements\", open for community comments"
author: Richard Dice
type: post
date: 2008-11-20 00:00:00 +0000 UTC
url: "/post/hague_grant_request_rakudo_dis"
categories:
 - Grants
 - Parrot development
 - Perl 6 Development
 - Perl Foundation
 - Sign in

---

"Jonathan Worthington":http://www.jnthn.net/ has submitted a request for an "Ian Hague Perl 6 development grant":http://www.perlfoundation.org/ian_hague_perl_6_development_grants for his proposal "Rakudo Dispatch and Role Enhancement".  A part of the Hague grant process is that submitted grant requests may, as opted by the submitter, be provided for public and community comment.

Jonathan's grant request is included here, below.  Any interested Perl community members may provide their comments regarding this grant request here.


Name:           Jonathan Worthington

Project Title: Rakudo Dispatch and Role Enhancements

Synopsis:
Less than a year ago, Rakudo implemented just a small fraction of the Perl 6 language specification, passed very few of the spectests and was far from ready to develop any real software with. By contrast, today Rakudo supports a wide range of Perl 6 features, passes over 4,000 spectests, can run as an Apache module and is being used to implement a November, a wiki.

Even though it has been a year of great progress, there still remains a lot of work to do to make Rakudo into a complete, production-ready implementation of Perl 6. This project aims to add some of the missing major OO and type-related features that are specified in the Perl 6 synopses, mostly in S12, and which are required to bring us closer to a complete Perl 6 implementation.

Benefits to Perl 6 Development:
All items listed in the deliverables of this proposal are required in any full implementation of Perl 6.

Deliverables:
D1. Register symbols in the namespace at compile time, allowing elimination of current "types must start with upper-case letters" hack and detection of re-defined routines, as specified in S06 and S12
D2. Implementation of junction auto-threading, as specified in S09

D3. Implementation of submethods, as specified in S12

D4. Finish implementation of delegation ("handles" trait verb), as specified in S12

D5. Implementation of parametric roles, as specified in S12

Project Details:
D1 involves removing a hack that has seen us through the early stages of Rakudo, making sure classes really are composed and installed in the namespace at compile time.

D2 through D4 all involve work on the dispatcher. Therefore, there will likely be an initial refactor that paves the way for them, followed by getting each of them in place after it. Here are some notes on each of them.

D2: Currently, we have some junction support, but instead of being hard-coded as it is now, it should really all fall naturally out of the dispatch mechanism. Furthermore, since the dispatch mechanism doesn't currently account for junctions, dispatches with junctional parameters don't auto-thread when they should.

D3: We don't currently implement submethods at all; they are needed to move us closer to a complete S12 implementation.

D4: We currently implement some use cases of the "handles" trait verb on attributes. However, they should also be available on methods. Furthermore, only strings and pairs are currently available; the specification allows for a role or class to be mentioned as well as regex or substitution to be used for * as a wild card.

Rakudo currently has a basic implementation of roles. However, roles can also take parameters, which may form part of their long name (essentially, multiple roles share the same short name but with different long names - just like multis). D5 involves implementing this.

Project Schedule:
Assuming this grant is approved before/in early December:

December 2008: D1 mostly complete, dispatcher refactored in preparation for D2 through D4, D3 complete

January 2009: D1 through D4 complete, progress on D5 to the point where can declare and do a parametric role (differentiation of roles by th