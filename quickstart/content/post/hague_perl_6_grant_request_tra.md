
---
title: "Hague Perl 6 Grant request, \"Traits, Introspection and More Dispatch Work For Rakudo\", open for comment"
author: Richard Dice
type: post
date: 2009-04-28 00:00:00 +0000 UTC
url: "/post/hague_perl_6_grant_request_tra"
categories:
 - Grants
 - Perl 6 Development
 - Sign in

---

Coming off the completion of his first Hague grant, Jonathan Worthington has submitted a request for a new Ian Hague Perl 6 development grant for his proposal "Traits, Introspection and More Dispatch Work For Rakudo". A part of the Hague grant process is that submitted grant requests may, as opted by the submitter, be provided for public and community comment.

Jonathan's grant request is included here, below. Any interested Perl community members may provide their comments regarding this grant request here.


<big>**Name**</big>
Jonathan Worthington

<big>**Project Title**</big>
Traits, Introspection and More Dispatch Work For Rakudo

<big>**Synopsis**</big>
This project aims to add additional functionality specified in S12 and 
S14 to Rakudo Perl 6.

<big>**Benefits to Perl 6 Development**</big>
Today, Rakudo is one of the most advanced, and certainly the most actively developed, Perl 6 implementation. I am seeking this grant to fund my time to continue implementing features in Rakudo that are on the path to it being a complete implementation of the Perl 6.0.0 specification. In fact, the work proposed under this grant would bring us very close to a complete implementation of S14, and fill out a substantial amount of the bits of S12 that Rakudo currently does not support.

<big>**Deliverables**</big>

**D1.** Implementation of .WALK, .can returning an iterator, dispatches of the form $foo.@candidates, and all of callsame, callwith, nextsame, nextwith and lastcall, as specified in S12. This work will also give coverage of the section of S12 under the heading "Interface Consistency". Furthermore, it will layer Rakudo's dispatch more neatly  onto Parrot's own dispatch mechanisms, improving Rakudo performance.

**D2.** Refactoring of .wrap and .unwrap, as specified in S06 to work in terms of the same candidate lists as these S12 features.

**D3.** Implementation of the parallel dispatch section in S12.

**D4.** Filling out of the default Perl 6 metaclass to support the introspection capabilities specified in S12.

**D5.** Implementation of defining methods on the metaclass, as described in the Class methods section of S12.

**D6.** Implementation of traits, as specified in S14.

<big>**Project Details**</big>
In my last grant, I did some significant improvements to dispatch in Rakudo. However, there's still some work to go. D1, D2, D3 and D5 aim to fill out what is missing, and will incorporate a refactor that will bring improvements to method dispatch performance in Rakudo.

Rakudo currently allows very limited introspection. D4 will fill out the meta-model and provide much fuller introspection capabilities. This work will also bring us closer to implementing the HOW API suggested by the smop project, which is likely to be incorporated into S12. This also means the foundations for people to define their own metaclasses will have been laid by the end of this grant. Completing this section of the 
grant will require getting input and consensus on some currently under-specified bits of the introspection interface and ensuring those clarifications are made explicit in S12.

Currently we cheat a bit to support the Perl 6 built-in traits in Rakudo. In D6, I aim to implement the traits section of S14, refactoring existing trait handling to match up with this model while providing support for user-defined traits.

<big>**Project Schedule**</big>
Work will begin as soon as the grant is approved. I expect to accomplish the majority of the deliverables during May, and if all goes smoothly aim to wrap up the grant for the end of June or early July. Here is a rough timetable.

Middle of May: Substantial progress on D1, D4 started

End of May: D1 and D2 mostly complete, notable progress on D3, D4 and D5

Middle of June: D1 through D5 mostly complete