
---
title: "Perl 6 Performance and Reliability Engineering: November 2017"
author: Coke
type: post
date: 2017-11-19 00:00:00 +0000 UTC
url: "/post/perl_6_performance_and_reliabi_4"
categories:
 - Grants

---

As part of the Perl 6 core development fund, Jonathan Worthington has completed another 200 hour block of hours, and his report of what
was completed follows the break.

Many thanks to the TPF sponsors of this and other grants. If you're interested in supporting work like this, please donate: http://donate.perlfoundation.org/

----


# Grant Completion Report: Perl 6 Performance and Reliability Engineering

At the end of July, I was granted a [200 hour extension](http://news.perlfoundation.org/2017/07/grant-extension-approved-perl-.html)
of my Perl 6 Performance and Reliability Engineering grant. I used this time
primarily to focus on MoarVM's dynamic optimizer, although did many other
fixes and improvements aside from that.

## Background on the dynamic optimizer improvements

Modern runtimes, especially for dynamic languages, rely on dynamic optimization
techniques. These analyze the behavior of the program at runtime, and use that
data to drive optimization. MoarVM's dynamic optimizer is typically referred to
as "the specializer", or "spesh" for short, and this nicely captures its core
strategy: taking code with lots of potential dynamism, and producing one or
more specialized versions of the code with as much of that dynamic behavior
stripped out as possible.

The specializer was planned as part of MoarVM from the start, although its
implementation came after the initial public release of the VM. Soon after
that, the focus switched to the Christmas release of Perl 6, where nailing down
semantics was a much bigger focus than optimization. Since then, the
specializer was improved in numerous ways, however various limitations in its
design started to show up repeatedly as we analyzed Perl 6 program performance.
Furthermore, stress testing showed up a range of bugs in the optimizer that had
potential to cause incorrect behavior in user code.

Therefore, for the sake of both performance and reliability, it was desirable
to invest some time improving the specializer.

## Specialization in the background

Modern hardware is parallel, and it is desirable to find ways to take advantage
of that. Originally, the specializer would do optimization work on the same
thread that user code was running on. This not only paused execution in order
to do optimization, but it also meant that multiple threads running the same
code (say, in a data parallel program) would all race to do and install the
same optimization.

I introduced a background thread for performing specializations. This not only
meant that optimization work and JIT compilation would not interupt execution,
but also resolved the issue of multiple threads scrambling to do the same
optimization. Since there was now only one thread producing optimizations, some
locking logic could also go away. A further upshot of this is that even a
single-threaded user program can now get some advantage from multi-core
hardware.

One downside of this is that the exact timing of specializations being installed
becomes unpredictable, and this can make debugging more difficult. Therefore, I
added a deterministic mode, enabled by environment variable, which makes a
thread pause while the optimization thread does its work. This, for a
single-threaded user program, offers deterministic behavior.

## Better data for better decisions

The specializer's decision making about what to optimize, and how to optimize
it, will only be as good as the data available to it. The data model
prior to my work under this grant was rather limiting. For example, it was not
possible to get a high level overview of what was monomorphic code (same types
always), polymorphic code (handful of different types) and megamorphic code
(loads of different types). There were also too few samples to know if a type
that was seen to differ once was really rare or not. When there are only ten or
so samples, and a type differs one time, then it could vary up to 10% of the
time; that will tend to be too costly to leave deoptimization to handle. By
contrast, if there are a hundred samples and it happens one time, then it is
much safer to leave that slow path to be handled by the interpreter, for the
sake of running the common case potentially a lot faster.

I implemented a lightweight interpreter logging mechanism. It writes compact
logs about encountered types, invocation targets, and more into a sequential
thread-local buffer. When the buffer is filled, it is sent to the specialization
thread. Over on that thread, the recorded events are played back in a stack
simulation, and a data model built up that aggregates the statistics. This is
then used by a planner to decide what optimizations to produce.

Along the way, I introduced a new kind of specialization, which specializes
only on the shape of the callsite (how many arguments and which named arguments)
rather than the incoming types. This means that megamorphic code (that is,
code called on many different types) can now receive some optimization, as well
as compilation into machine code. Before, a few specializations were produced,
and then everything else was left to run slowly under the interpreter.

## New optimizations

Besides allowing for better decision making, I introduced some new optimizations
as well as making existing ones more powerful.

* I enabled many more calls to be inlined (a powerful optimization where a
  call to a small routine is replaced with the code of the routine itself).
  This was achieved by using the new statistics to spot when the target of a
  call was reliably the same, and introducing a guard clause. Prior to this,
  inlining was only available to methods resolved through the cache or subs
  in the setting or the outermost scope. I also handled the case where the
  passed arguments were consistently typed, but it had not been possible for
  the optimizer to prove that, again using guard clauses.
* I implemented inlining of closures (that is, code that refers to variables
  in an outer scope).
* I made dead code removal happen far more eagerly, and improved the
  quality of type information available in code following the eliminated
  conditional. This is a significant improvement for parameters with default 
  values, as well as branches based on types or constants.
* I made frames that are reliably promoted from the call stack onto the heap be
  allocated right away on the heap, to save the promotion cost. (This promotion
  happens when a callframe needs to be referenced by a heap object.)
* I changed the way that control exception's flow is represented to be more
  accurate, enabling elimination of handlers that become unreachable once the
  code they cover also becomes unreachable. The change also resulted in more
  accurate type information propagation, which can aid other optimizations.
* I made the optimization that rewrites control exceptions into far cheaper
  `goto` instructions apply into inlines.

## Specializer fixes

The specializer usually only operates on "hot" code, so that the time that it
spends optimizing will have maximum benefit. However, it is possible to set an
environment variable that lowers these thresholds, making the specializer try
to optimize every bit of code that is encountered. This, combined with the
deterministic mode, provides a means to stress test the optimizer, by forcing
it to consider and optimize a great deal more code that it usually would do.
Running the NQP and Rakudo builds , together with the Perl 6 test suite, in
this way can flush out bugs that would not show up when only optimizing hot
code.

Prior to my work under this grant, failures would show up from this stress
testing as soon as the NQP build. After a good amount of bug hunting and fixing,
the NQP build and tests, together with the Rakudo build and basic tests, are
completely clean under this stresstest. The handful of remaining failures in
the stresstest are a result of inlining distorting backtraces (at the moment,
the inlined frames do not appear in the backtrace), thus causing some
error-reporting tests to fail.

The fixes included addressing a long-standing but rarely occurring crash
involving the intersection of specialization, multiple dispatch, and
destructuring in signatures; a number of different crashes boiled down to this.
Another important range of fixes involved eliminating poor assumptions around
`Scalar`s and aliasing, which aside from fixing bugs also stands us in a better
position to implement escape analysis (which requires alias analysis) in the
future.

## Notable results from the specializer work

The specialization improvements showed up in a number of benchmarks. Two that
are particularly worth calling out are:

* The daily `Text::CSV` module benchmark runs hit new lows thanks to the
  specializer improvements.
* The "read a million line UTF-8 file" benchmark that I've discussed before,
  where Rakudo on MoarVM used to be just a bit slower than Perl 5, is now won
  by Rakudo. This is a result of better code quality after specialization.

## Improved GC thread sync-up

I re-worked the way that garbage collection synchronizes running threads, to
eliminate busy-waiting. The idea of the original design was that running threads
would quickly react to the GC interupt flag being set on them. However, this
presumed that the threads were all really running, which is never a certainty
given CPU cores are competed over by many processes. Furthermore, under tools
like valgrind and callgrind, which serialize all threads onto a single CPU
core, the busy-waiting strategy produced hugely distorted results, and greatly
increased the time these useful, but already slow, tools would take. Now the
synchronization is done using condition variables, meaning that both the kernel
and tools like valgrind/callgrind have a much better idea of what is happening.
While callgrind showed a large (10%-15%) reduction in CPU cycles in some
multi-threaded programs, the improvements under normal running conditions were,
as expected, much smaller, though still worthwhile.

## Other work

Along with the improvements described above, I also:

* Added support to `Proc::Async` to plumb the output handle of one process into
  the standard input of another process, with the handles being chained together
  at the file descriptor level.
* Hunted down and fixed a SEGV when many processes were spawned and all gave an
  error.
* Fixed RT #131384 (panic due to bug in error reporting path of ASCII decoder)
* Fixed RT #131611 (sigilles variable in coercion could generate internal
  compiler error)
* Fixed RT #131740 (wrong `$*THREAD` after `await` due to lack of
  invalidation of the dynamic variable lookup cache)
* Fixed RT #131365 and RT #131383 (`getc` and `readchars` read too many chars
  at the end of the file)
* Fixed RT #131673 (`is rw` with anonymous parameter reported error incorrectly)
* Fixed MoarVM issue 611 (memory errors arising from certain usage patterns of
  the decode stream)
* Fixed MoarVM issue 562 (SEGV from a particular use of the `calframe(...)`
  function)
* Fixed native callbacks when the callback is made on a thread other than the
  one that passed the callback in the first place
* Avoided a linear lookup, knocking 5% off the code-gen phase of compiling
  `CORE.setting`
* Removed the long-unused `Lexotic` feature from MoarVM, which allowed some
  code cleanup (this used to be used to handle `return`, but it now uses the
  control exception system)

## Dissemination

I wrote a 4-part series on my blog about the MoarVM specializer. The posts walk
through the MoarVM specializer's design and functionality, and mention the many
improvements done as a result of this grant - explaining why the new way of
doing things represents an improvement from the previous way. The posts are:

* [Part 1: Gathering Data](https://6guts.wordpress.com/2017/08/06/moarvm-specializer-improvements-part-1-gathering-data/)
* [Part 2: Planning](https://6guts.wordpress.com/2017/09/17/moarvm-specializer-improvements-part-2-planning/)
* [Part 3: Optimizing Code](https://6guts.wordpress.com/2017/11/05/moarvm-specializer-improvements-part-3-optimizing-code/)
* [Part 4: Argument Guards](https://6guts.wordpress.com/2017/11/09/moarvm-specializer-improvements-part-4-argument-guards/)

I also travelled to the Swiss Perl Workshop and delivered a talk about the
MoarVM specializer, titled "How does deoptimization help us go faster?". The 
[slides](http://jnthn.net/papers/2017-spw-deopt.pdf) and [video](https://www.youtube.com/watch?v=3umNn1KnlCY)
were published online.

## Summary

This latest grant extension enabled me to spend a significant amount of time on
the MoarVM dynamic optimizer, both addressing bugs as well as overhauling the
way information about program execution is collected and used. The new data
allows for better decision making, and its availability allowed me to implement
some new optimizations. Furthermore, I moved optimization work to take place on
a background thread, so as not to interupt program execution. Aside from this
work, I fixed other bugs and made some performance improvements unrelated to
the dynamic optimizer. Finally, I gave a presentation about dynamic optimization
at the Swiss Perl Workshop, and wrote an extensive 4-part blog series explaining
the MoarVM optimization process.

