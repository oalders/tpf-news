
---
title: "Perl 6 Performance and Reliability Engineering: Grant Report"
author: Makoto Nozaki
type: post
date: 2017-04-22 00:00:00 +0000 UTC
url: "/post/perl_6_performance_and_reliabi_3"
categories:
 - Grants
 - Sign in

---

*This is a grant report by Jonathan Worthington on his grant under Perl 6 Core Development Fund. We thank the TPF sponsors to make this grant possible.*

I have completed the second 200 hours of my [Perl 6 performance and reliability
engineering](http://news.perlfoundation.org/2016/02/grant-proposal-perl-6-performa.html)
grant, funded by the [Perl 6 core development fund](http://www.perlfoundation.org/perl_6_core_development_fund).
This report summarizes the work done during those 200 hours. In accordance with
community feedback, the vast majority of effort has been put into reliability
rather than performance.

## Concurrency robustness

The main area of focus in this grant period has been making Perl 6's concurrency
support more robust. While work remains to be done, the improvement over the
last several months has been noticeable. It is also an important area for me to
focus on, given the small number of people in the community with the skills,
time, and patience (or, perhaps, stubbornness) to track down and resolve these
problems. Here is a summary of the issues resolved.



* Fixed a bug affecting use of `callwith` in multiple threads
* Fixed RT #128809 (closure-related bug involving `s///` construct, which showed
  up in concurrent scenarios)
* Fixed RT #129213 (synchronous socket `accept` could block GC in other threads,
  thus blocking program progess)
* Determined RT #128694 fixed, added test (`zip-latest` on two intervals would hang)
* Eliminated use of in-place rope flattening, which violated the immutability
  of strings and could thus cause various crashes (especially in hash access);
  this resolved many failures, amongst them the one reported in RT #129781, and
  also made hash lookups using rope strings keys more efficient as a bonus
* Fixed warnings due to over-sharing of `$/` between threads when using grammars
  in parallel (mostly fixed RT #128833)
* Fixed a `Lock.protect` bug when we unwound the stack due to control exceptions
  (discovered independently of, but also resolved, RT #126774)
* Fixed RT #129949 (GC crash resulting from missing rooting of sent value in
  concurrent blocking queue)
* Fixed RT #129834 (sporadic behavior when concurrently creating `Proc::Async`
  objects and obtaining handles)
* Audited and fixed vulnerable cases of the `once` construct
* Fixed RT #129994 (long-lived native call on one thread could block GC in other
  threads)
* Fixed RT #125782 (uninformative error reporting when a `Promise` is broken)
* Examined RT #127960; concluded it is fixed, but added it as a stress test
  since it's a bit distinct from the other test for the same underlying bug
* Fixed a bug where method caches could be revealed to other threads before
  they were fully deserialized, causing falsely missed lookups
* Fixed a data race inside of the NativeCall setup code
* Fixed RT #130064 (trying to rethrow an exception that was never thrown before
  leaked an internal error; this showed up in Promise.break("foo"))
* Fixed scoping/cloning problem with `LAST`/`NEXT`/`QUIT` phasers in `supply`,
  `react`, `whenever`, and `for` constructs
* Fixed a bug with `QUIT` phasers mishandling exceptions thrown synchronously
  with the `.tap`
* Switched to using `Supplier::Preserving` on the taps of stdout/stderr in
  `Proc::Async`, to avoid various innocent-looking usage patterns losing output
* Fixed RT #128991 (messages could seep out of a `supply` block even after it was
  considered done)
* Fixed a GC corruption bug involving `Proc::Async` that caused occasional crashes
* Tracked down and fixed two data races in the `supply`/`whenever` implementation
  and in `Supply.interval`
* Fixed RT #130168 (`Supply.interval(...)` with very small interval would only
  ever emit 1 value)
* Fixed interaction of native callbacks, GC blocking, and embedding, which
  afflicted `Inline::Perl6`
* Fixed use-after-free that could occur as part of inlining fixups when in a
  multi-threaded program
* Fixed precompilation of the `OO::Monitors` module
* Fixed RT #130266 (premature frees of handles in various async socket error
  handling cases)
* Fixed SEGVs when GC stressing was applied to S15-nfg/many-threads.t and 
  S17-supply/syntax.t
* Fixed incorrect reporting of some errors on threads, which could show up as
  if they were compile-time errors
* Fixed thread safety issue in the `>>.foo` implementation
* Fixed a miscompilation of `||=`, `&&=`, and `//=`, making them a good bit more
  efficient along the way
* Add various bits of missing concurrency control in precompilation management,
  thus fixing parallel use of precompilation (this will help towards faster
  `p6doc` builds)

## String decoding improvements in asynchronous I/O

Previously, decoding of bytes to strings for both `IO::Socket::Async` and
`Proc::Async` was done at the VM level. This created a number of fragilities
with regard to decoding errors. Due to time constraints, different encodings
besides UTF-8 had not been implemented for these classes either, leaving users
of them to do decoding manually if they needed anything else.

To rectify these issues, I first made the VM-backed decoders directly available
to userland. These will, in the future, be exposed as a Perl 6-level API, and
we'll support user-space encodings. For now, it meant I could move the code that
orchestrates the decoding of strings in async I/O into Perl 6 space, fixing the robustness issues. This
also means that string decoding for different spawned processes and socket
connections can be done in the thread pool, rather than using the
event-processing thread. Along the way, I added support for different encodings.

Finally, there were some issues around the way async sockets and processes
worked with regard to NFG. I resolved these issues and made sure there was
test coverage of the various edge cases.

## Non-blocking await and react support

I did the initial round of work to provide support for non-blocking `await` and
`react`. At present, these constructs will block a real OS thread, even if used
on a thread in the thread pool. The changes, available via. `use v6.d.PREVIEW`,
mean that thread-pool threads will be returned to the thread pool to do other
work, and the code following the `await` and `react` will be scheduled once the
result is available or processing is complete. This is implemented using
continuations (much like `gather`/`take`, except in this case the continuation
may be resumed on a different OS thread). The result is that Perl 6 programs
will be able to have hundreds or thousands of outstanding `react`s and `await`s,
with just a handful of real OS-threads required to process them.

This is just the initial implementation; further work will be required to make
this feature ready to be the default in Perl 6.d.

## Memory leak fixes and memory use improvements

The highlight of the memory management improvements was a simplification to
the lifetime management of register working sets in MoarVM. This resulted from
the elimination of a couple of speculative features that were not yet being
utilized by Rakudo, and in one case never would have been anyway. Coupled with
a range of cleanups and some code streamlining, the result was a 10% reduction
in peak memory use for CORE.setting compilation, and 20% off the compilation
runtime. I also:

* Fixed a bug that caused bogus multi-dispatch cache misses for calls with many
  named arguments, leading to the cache growing indefinitely with duplicate
  entries
* Fixed a regex interpolation memory leak; it boiled down to unclaimed entries
  left behind in the serialization context weakhash
* Fixed leaks of asynchronous task handles
* Fixed a leak in decode stream cleanup
* Improved memory allocation measurement in I/O, meaning that full GC collection
  decisions are made more accurately in I/O-heavy programs
* Fixed a memory leak involving `Proc::Async`
* Fixed a memory leak when a synchronous socket failed to connect
* Tracked down and resolved the last remaining leaks that showed up in
  `perl6-valgrind-m -e ''`, meaning it is now clean. (Previously, some cleanup
  was missed at VM shutdown)

## Unicode-related work

I did a number of Unicode improvements, as well as discussing with and reviewing
Pull Requests from a new contributor who is now doing a bunch of Unicode
work for Perl 6. My own contributions code wise were:

* Initial support for Unicode 9 (updating the character database, basic NFG
  tweaks)
* A rewrite of the UTF8-C8 encoding to eliminate various bugs (some falling
  under RT #128184), including a buffer overrun and not properly round-tripping
  valid but non-NFC input

## Other assorted bugs

I also took care of a range of other bugs, which don't fit into any of the
previously mentioned areas of work.

* Fixed RT #128703 (1 R, 2 R, 3 lost values)
* Fixed RT #129088 (lack of backtaces for `sprintf` and friends)
* Fixed RT #129249 (mis-compilation of `/$<cat>=@(...)/`)
* Fixed RT #129306 (erorr reporting bug involving sub-signatures)
* Partially fixed and tested RT #129278 (native attributive parameter binding broken)
  and noted on the RT the less common cases that remain to be fixed
* Fixed RT #129430 (sigilless parameters were declared too late to use in where
  clauses)
* Fixed RT #129827 (sub { 42.return }() ended up being code-gen'd without a return
  handler)
* Fixed RT #129772 (poor error reporting when you tried to invoke a native
  parameter; it blew up code-gen and gave no location info)
* Tracked down and fixed a pre-comp management bug on Windows due to a file not being
  closed and then trying to rename over it
* Fixed RT #129968 (error-reporting crash in the case of redeclaration errors in
  nested packages)
* Fixed a bug with augmenting nested classes
* Fixed RT #129921 (internal warning when producing exception for `my $!a`)
* Hunted down a bug in the JIT-compilation of the `nqp::exception()` op and fixed it
* Fixed RT #130107 (accidentally treated `nread == 0` as an error in a couple of
  places in MoarVM)
* Fixed RT #130081 (did not backtrack into a regex `TOP` in a grammar to try and
  make it match until the end of the string)
* Fixed RT #130294 (SEGV that occasionally occurred during some cases of deep recursion)
* Fixed RT #128516 (SEGV when composing meta-object held in an attribute)
* Fixed RT #130465 (ignoremark not applied with backslashed literals)
* Fixed RT #130208 (putting multi-line Pod documentation on a role would pun it)
* Fixed RT #130615 (code-gen of `$a++` in sink context for native `$a` was a lot
  slower than in non-sink context)
* Fixed RT #130637 (two grammar constructs produced malformed NFAs, which gave
  wrong results or could even SEGV in MoarVM; MoarVM was made to validate NFAs
  more strongly, which shook out the second issue besides the reported one)
* Investigated RT #129291 (SEGV involving processes with output of one fed into
  input of the other); applied likely fix
* Investigated and fixed an issue with `$/` setting, arising from changes to 
  the implementation of `match`
* Fixed a bug that occasionally caused spectest crashes; was an interaction
  between dynamic lexical caching, inlining, deoptimization, and garbage
  collection
* Fixed a rare crash related to assignment when the type constraint was a
  refinement type
* Fixed MoarVM #120 and #426, in which a failed debug annotation lookup led to
  boxing a NULL string
* Fixed a couple of places where dynamic optimization could accidentally
  trigger garbage collection; the optimizer assumes this can never happen
* Fixed RT #123989 and RT #125135 (callsame et al could sometimes have the
  dispatcher stolen by the wrong target invokee)

## Other tasks

On top of this, some time was spent reviewing pull requests to Rakudo, NQP,
and MoarVM, providing feedback, and merging them when appropriate. I also
commented on a range of RT tickets besides those I fixed myself. Various other
small cleanups and additions resulted from this work, ranging from typo fixes
in comments up to improvements to GC debugging macros added while finding bugs.

