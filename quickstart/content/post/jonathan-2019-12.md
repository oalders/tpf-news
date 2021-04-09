
---
title: "April-December report of the Perl 6 Development Grant of Jonathan Worthington"
author: Matthias Bloch
type: post
date: 2019-12-21 00:00:00 +0000 UTC
url: "/post/jonathan-2019-12"
categories:
 - Grants
 - Perl 6 Development

---

Jonathan writes:

This is a status update on my Perl 6 Performance and Reliability grant. First of all, I'd like to explain why there's been so little progress in the latter part of the year. In August, the process that led to the Raku rename started. This was designated as a language issue, and as the responsible person for the language label in the problem-solving process, it fell to me to help shepherd the issue through to a resolution. Most of the time and focus in those months went there rather than on grant work. Almost immediately after that, things got decidedly busy in my non-work life (good things, just time-consuming). After a good bit of rest over the holiday period, I'll be finally ready to dig back into work on Rakudo and MoarVM.

Meanwhile, here's a summary of the work that I did get done since I last filed a report. The following
work was completed and has already appeared in a release:

```
4:52	Work on finding and fixing memory leaks involving a Supply and a Channel that sends to itself
1:13	Hunt down and fix a memory leak affecting longer-running async servers
0:42	Allow the maximum inline size to be picked per language; re-tune it for NQP, achieving 5% off
	spectest times, less compilation memory use, etc.
0:33	Fix an object creation performance bug on objects with no attributes
0:40	Fix a couple of bugs related to BUILDPLAN, one of them where multiple inheritance and attribute
	initialization were broken
3:54	Various optimizations to speed up match object construction
0:43	Look into an async process spawning issue that had an odd failure mode; add improved error reporting
0:46	Fix failure to propagate an exception when a race nested inside another race died (similar for hyper)
3:49	Fix problems with profiling and debugging multi-threaded programs (assorted symptoms, one root cause)
2:39	Fix high resource usage by spesh stack simulation in some cases where specialized code called
	unspecialized code
55:00	Fix a crash in uninlining due to the callsite not being reinstated during deopt
39:00	Fix Proc not being falsey and not blowing up in sink context when the process that was run exited
	with a segfault
1:20	Fix a bug relating to compile-time use of EVAL
1:32	Fix a bimodal performance bug and a performance regression arising from specialization
	logging's interaction with the creation of new compilation units
1:12	Fix optimization of routines where a Scalar containing an unboxable object was passed to
	a native parameter
1:16	Better optimization of Nil assignments
2:20	Better optimize assignment into an untyped scalar
1:41	Make the bindlex instruction more inline-friendly, allowing for further inlining optimizations
6:26	Code generation improvements to produce less return handlers, resulting in more compact
	compilation output and bringing more code within the inline budget
0:49 	Simplify code generation of for 1..100 { } style loops
1:50	Work on speeding up invocation; > 2% off a benchmark's CPU cycles with a couple of changes
1:07	Assorted release-blocker bug fixes
```

I also continued work on escape analysis. Bringing this to a merge-ready state will be my main focus
when I get back to working on Rakudo and MoarVM in the new year. The main areas of work on this have
been:

```
12:31	Work on making the escape analysis algorithm into a partial escape analysis one. This also
	makes it capable of allocation sinking - that is, moving allocations into the codepaths
 	where they happen, or delaying them until they have to happen to allow further scalar
	replacement until that point
20:13	Work on escape analysis handling of big integer operations; this will be especially useful
	in that every Int is conceptually a big integer, and through scalar replacement of them we
	have a chance at making integer math far more efficient and free of needing to make any
	function cal