
---
title: "January report of the Perl 6 Development Grant of Jonathan Worthington"
author: Matthias Bloch
type: post
date: 2019-02-11 00:00:00 +0000 UTC
url: "/post/january_report_of_the_perl_6_d"
categories:

---

Jonathan writes: 

January was a busy and productive month for my Perl 6 grant work.

Back in November, I was working on allowing us to lower lexical variables into locals. This is possible when they are not closed over, and allows for generation of more efficient code, which is in turn much easier for backends - MoarVM, the JVM, and JavaScript - to deal with in their own optimizers. It can also shorten the lifetimes of objects, leading to better memory performance. I completed and merged the almost-complete work to be able to lower normal variable declarations (such as `my $a = 42`). This will combine well with the ongoing work on escape analysis, since it will be able to eliminate many allocations of Scalar containers.

The next step was to take on the lowering of `$_`. In Perl 6.d, we made this a normal lexical rather than a dynamic lexical (available through `CALLER::`) in order to allow further analysis and optimization. Relatively few Perl 6 users were aware of its dynamic nature anyway, and fewer still making use of it. With the topic variable of course showing up in numerous common idioms, it would be a pity to then have people avoid them in the name of performance, due to the dynamic nature of `$_` impeding analysis. Thus the change in 6.d.

This month I worked on implementing the most immediate optimization that this enabled: being able to also lower `$_` into a local. In a loop like `for 1..10_000_000 { something() }`, previously we would have to allocate an `Int` object for every single iteration. Since `$_` was dynamic, we could not be sure that `something()` - or anything it in turned called - would not access it. Under the new semantics, we can lower it into a local, and then dead code analysis can see that the boxing is useless and discard it, saving a memory allocation every iteration. Even were it used, for example in an array index, we now have opened the door to being able to use escape anslysis in the future to also avoid the allocation.

This work had some amount of fallout, and turned up some rather dubious semantics around regexes matching against the topic variable when in boolean or sink context. This included some curious "at a distance" behavior, where things either worked by a happy accident, or could potentially match against completely unexpected data further down the callstack thanks to the thing they were expected to match against being undefined! I proposed a cleaner set of behaviors here and, with a lack of objections, implemented them.

Along with these optimizations, I also implemented more aggressive scope flattening in Rakudo's static optimizer. When we can show that a particular scope's existence would not be missed - such as the body of an `if` statement which doesn't have any variables in it that get closed over - we can flatten it into the surrouding block. This is a bit cheaper at runtime and leads to more chances to do lexical to local lowering, which - as already mentioned - is the gateway to further optimizations.

Back over in MoarVM, I continued my work on escape analysis and scalar replacement. I took a small detour to implement retention of deoptimization information in specialized bytecode, so we can recover it when inlining. This will allow us to more aggressively optimize code that we inline. I didn't yet enable the full set of optimizations we might do there, but I was able to enable dead instruction elimination, which can save some work (for example, if we see a return value is unused, we can strip out the work that produces it). This more detailed data was also required for the analysis algorithm that works out which object allocations that we eliminate as part of a scalar repalcement need to be restored when deoptimizing. I got most of that worked out, with a few issues left before it could be merged. (Spoiler: those were resolved in early February an