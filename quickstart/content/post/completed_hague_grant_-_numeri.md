
---
title: "Completed Hague Grant - Numeric and Real Support"
author: Karen Pauley
type: post
date: 2010-08-04 00:00:00 +0000 UTC
url: "/post/completed_hague_grant_-_numeri"
categories:
 - Grants

---

I am pleased to announce that Solomon Foster has successfully completed his Hague Grant for "Numeric and Real Support":http://news.perlfoundation.org/2010/03/hague-grant-application-numeri.html.  I would like to thank Solomon and his grant manager, Jonathan Worthington, for all their work on this project.  

_Solomon Foster writes:_

**Introduction**

While I worked on this grant, Rakudo has matured into a capable compiler ready to go out into the world as Rakudo Star. The work done under this "grant":http://news.perlfoundation.org/2010/03/hague-grant-application-numeri.html ensures it will have a solid implementation of the S32 Numeric specification. And, as is usually the case when part of the Perl 6 specification is implemented for the first time, my experiences doing so have led to improvements and clarifications.

**Deliverables Status**

All deliverables have been achieved, and some additional work completed as well.

D1. The Numeric role is now in place as the root of all number types in Rakudo. Furthermore, the "Numeric":http://github.com/rakudo/rakudo/blob/master/src/core/Numeric.pm spec has been modified to include casting methods, the trigonometric base conversion methods, and a tentative approach to allow Numeric types that do not know of each other to be sorted together.

D2. The "Real":http://github.com/rakudo/rakudo/blob/master/src/core/Real.pm role is also in place. All of the built-in numeric types other than Complex now do Real. In addition to cleaning up the Real spec a bit, I introduced the Bridge casting method to make it very easy to write a new Real type which is fully interoperable with any other properly-coded Real type. 

D3. Int, Rat, and Num now do the Real role in Rakudo. So far this seems to work very well, particularly because Int and Rat have many of their methods provided by Real.

D4. Complex is now a Numeric type composed of two Reals. Again, this is a good fit, a testament to the solid fundamental design of this area of the spec.

D5. The trig functionality is now implemented in Num and Complex, with Cool forwarding trig methods to the Numeric cast of their arguments, and Real forwarding them to the Bridge cast. In practice, this means that any Real type other than Num automatically gets its trig functionality by delegation to Num, though the Bridge approach allows implementations to select a different type to handle this.

D6. I added the "real-bridge.t":http://svn.pugscode.org/pugs/t/spec/S32-num/real-bridge.t test file to the spectests to test that an arbitrary new Real type will easily fit properly into the existing scheme. I didn't add a Numeric-specific test file, as fewer tests were required there and they seemed to blend naturally into existing test files.

D7. With approximately 15,000 tests deleted or added, by far the biggest test changes were to the trig tests. I overhauled them to cover more usage cases while using far fewer actual test cases. In the process I found several trig bugs and fixed them.

In addition to this, I fixed the long-broken handling of long numeric constants with decimal points in Rakudo, as it turned out to be a major impediment to revising the trig tests. Though the resulting patch needs to be refactored, it represents a major improvement in an area users are likely to encounter.

**Dissemination**

I have written "fourteen blog posts":http://justrakudoit.wordpress.com/ on the Numeric work, as well as remaining active on the #perl6 channel. I was also part of the Perl 6 contingent at "YAPC::NA":http://yapc2010.com/yn2010/.

**Conclusions**

This grant has advanced the state of the art in Rakudo, both in terms of the numerics classes themselves and in terms of the bugs turned up and fixed in the general implementation of roles. It has improved the specification and the test suite. In addition, both the Numeric and Real roles are implemented in almost p