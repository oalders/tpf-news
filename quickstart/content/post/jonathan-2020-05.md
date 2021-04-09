
---
title: "May report of the Raku Development Grant of Jonathan Worthington"
author: Matthias Bloch
type: post
date: 2020-06-16 00:00:00 +0000 UTC
url: "/post/jonathan-2020-05"
categories:
 - Raku Development
 - Sign in

---

Jonathan writes:

In my [April report](https://news.perlfoundation.org/post/jonathan-2020-04)
I mentioned that I had performed an amount of design work for a new
generalized dispatch mechanism in MoarVM. In May, I forged ahead with the
implementation work on this - although in reality, that work isn't just
churning out code, but also making lots of smaller design decisions along
the way too.

I started work on MoarVM in 2012. I've learned a thing or two about building
runtimes since then, and have applied many of those learnings to MoarVM.
However, thus far, that has tended to come in the form of additions to what
was already there - initially in aid of supporting all that the Raku language
wanted to offer, and later in terms of running it more efficiently. The
generalized dispatch work is a bit different: it's about asking how we can
do moar with less.

The entry point to it all is a new `dispatch` instruction. It identifies a
dispatcher - something that knows how to take a set of arguments and turn
them into an outcome - and provides it with some arguments. Those arguments
may be aimed at the dispatcher (such as a method name), a final piece of code
we choose to invoke (method arguments), or perhaps both (arguments in a
multiple dispatch). This instruction will come to replace all existing
invocation related instructions, including specializer plugin resolution,
method caching, and multi dispatch caching.

The dispatch instruction also marks the first time in MoarVM's history that
the calling conventions will change. The current invocation sequence for a
method call looks like this:

```
findmethod r3, r1, 'foo'
prepargs <callsite>
arg_o 0, r1
arg_o 1, r2
invoke_o r4, r3
```

That is, look up a method `foo` on the object in register `r1`, and store
it in the register `r3`. Then, take the code object in register `r3`, invoke
it with object arguments in registers `r1` and `r2`, and store the return
value into register r4. The callsite object provides the static properties
of the arguments (for exmaple, which ones are named and which ones are
flattening).

The new `dispatch` instruction would look more like this:

```
const_s r0, 'foo'
dispatch_o r4, 'raku-method', r0, r1, r2
```

First, we put the method name into register `r0`. Then, we do a dispatch using
the `raku-method` dispatcher, providing it with the arguments `r0`,
`r1` and `r2`, and put the result of the method call into `r4`. It turns a
32-byte sequence into a 24-byte sequence. This leads to rather more compact
bytecode. It also means we have just two instructions for the interpreter to
dispatch, not 5. The whole concept of an "args buffer" also goes away; the
register indices become a "map" of how to look up the arguments directly
from registers. This change of calling conventions runs pretty deep - but
thankfully, it was possible to introduce it alongside the current approach,
using it only when a `dispatch` instruction initiated things. Of course, the
end goal is that the old calling conventions will fade from use, and all code
associated with them will go away. Without going into more details,
I'll note that the argument processing and memory management with the new
approach is also overall simpler.

But what does `dispatch` do? The first time we reach a `dispatch`, we take
the string name of the dispatcher and resolve it to something we can invoke
to decide what will happen. That something receives the arguments as a
capture object. It can read the arguments, and use them to decide what to
do. For example, the `raku-method` dispatcher can take the invocant
argument, the method name argument, and call `$obj.^find_method('foo')` in
order to ask the MOP for the method we should call.

A dispatcher can then *delegate* to another dispatcher. In fact, all
user-defined dispatchers delegate. Only a handful of built-in ones, provided
by MoarVM, are "terminals": these do things like evaluating to a constant,
evaluating to one of the input arguments, invoking bytecode, or invoking
some VM primitive function. When one dispatcher delegates to another, it
provides it with an argument capture. This is derived from its input argument
capture. There are a set of operations for doing this; one can:

* Drop some arguments
* Insert an argument at a difference place, or duplicate it
* Insert an object or string constant argument
* Read an attribute from an argument, and insert the result of that read

These operations are part of building up what I call a "dispatch program":
a program consisting of a small set of simple instructions that map the
arguments into a final outcome (the outcome often being to run some bytecode,
which is what you'd expect of a dispatch). As well as capture transformations,
we can also add *guards* into the dispatch program: things that check the
types of arguments. If the guards are not met, the dispatch program fails.

To give an impression of how delegation will be used, let's consider a method
dispatch in Raku. It will:

* Start in a `raku-method` dispatch, which does a lookup of the method, and
  inserts the result at the start, and delegates to a `raku-method-resolved`
  dispatcher (which will also be reached by other forms, such as qualified
  method dispatch).
* In `raku-method-resolved`, check if this is a multiple dispatch routine
  or not, and delegate as appropriate. Before it does so, it removes the
  method name argument (it will retain this as "dispatch state" for if we
  need to resume the dispatch for a callsame; this piece is yet to implement,
  however).
* Assuming it was a single dispatch, end up in `raku-invoke`. This will see
  if it's a `Code` object of some kind, and if so extract the `$!do` attribute,
  which is a VM-level code handle. The Raku-level code object argument will
  be dropped, and and VM-level code handle inserted in its place.
* Finally reach the `boot-code` dispatcher, which is the terminal,
  and results in invoking bytecode.

After we've recorded a dispatch program, we transition the `dispatch`
instruction from an "uninitialized" state into a "monomorphic" state -
that is to say, we've set up one dispatch program, and so far as we know,
every dispatch we encounter will look just like the one we already saw.
And if we only ever call a method on the same type, for example, then that
is how things will stay. If, however, the dispatch program fails - perhaps
we are dispatching a method on a difference type - then we'll do a second
recording. At this point, the `dispatch` instruction is transitioned to a
"polymorphic" state. These states are stored in an array indexed based on
bytecode position of the dispatch instruction, a means of achieving a kind
of (polymorphic) inline cache, but without losing our ability to `mmap` the
bytecode for sharing across multiple processes.

The whole delegation from one dispatcher to another is erased in the dispatch
program that we end up with. Furthermore, the way the transformations and
guards are recorded, and then compiled, automatically eliminates duplicate
guards, duplicate attribute lookup paths, duplicate guards on those
attribute lookups, and so forth. All of the capture transformations are also
erased; an access into an N-times modified capture that maps to an original
argument will simply read that argument. In fact, in the common case where
the dispatcher gets its arguments of interest at the start (for example, the
method name), and the arguments for the final destination of the dispatch
are all at the end, then the entire load of shuffling around between the
dispatchers is optimized into a single addition to the arg indices pointer!

By the end of May, I'd largely got what I described so far implemented,
although was not at the point of having started using the new dispatcher in
Rakudo to replace the existing specializer plugins. Thus, I'll write about
that step in the next report.

Doing the dispatcher work did entail a rather significant, but worthwhile, detour.
We need to store the various bits of state associated with dispatchers
somewhere. That "somewhere" must be somewhere that we can "slice off" when
we take a continuation. However, so far, we'd only used our callstack to store
call frames, and even then we wiped it clean at the point we did a heap
promotion of a call frame. Thus, that all had to be refactored. The changes
are rather nice.

The new callstack can contain a range of kinds of record. Unlike before, even
if a frame is heap promoted (because a GC-managed object needs to point to it),
it still retains a presence on the callstack. The callstack can also store
dispatch records and run records. But there's more!

Previously, taking a continuation would force a heap promotion of all the
frames on the callstack, and we'd have to traverse each frame. Under the new
design, a continuation reset (the base of a delimited continuation) always
starts a new callstack segment. This means that we can cheaply find the base
of the continuation, and "slice" off the segment(s) containing its frames
just by NULLing out a the pointer at the appropriate point in the linked list of
segments. We also don't need to do heap promotion of all the frames; we just mark the
callstack entries just like we would if they were referenced by virtue of
being in dynamic scope rather than in a continuation. Invoking the
continuation is then sticking those previously sliced off segments on the
end of the segment linked list. (This is a slight lie: we do end up having
to walk the frames on one side or the other to clear up dynamic variable
cache entries. This is a solvable problem, and then we can really say that
continuations are O(stack segments) rather than O(call frames).)

As if having to revist continuations wasn't enough fun, I also had to revisit
deoptimization, and of note, uninlining. Inlining is an optimization where
we insert the body of a callee into the body of the caller to avoid the cost
of the call and, more importantly, to allow us to do inter-procedural
optimizations. However, since we do speculative optimizations, we have to be
able to undo that, thus uninlining: recreating the callstack as it would have
been if we had not done the inlining optimizations, at the point we discover
that one of the speculations we made is not valid.

Previously, uninlining would force all call frames to be heap promoted, which
was kinda nice in that it meant we were now dealing with a linked list. In a
case where a mixin causes us to have to deoptimize the whole callstack, it
was thus easy to insert any frames producing by uninlining in the middle. But
what when we don't have that option any more? The naive approach involves a
whole load of hairy rewriting of the call stack, loads of `memcpy`, and
no doubt trashing much of our CPU cache contents along the way. Thankfully,
having spent plenty of the time in the Perl community, I know that laziness
is a virtue. So what's the answer? Lazy deoptimization! We just mark call
frames that we need to deoptimize and potentially do uninling of, and then,
when we reach them while unwinding the callstack, go and do the required bit
of uninlining before we continue execution. This means we're only ever fiddling
around at the stack top.

Something really surprising - at least to me - happened when I did this. I've
always dreaded working on deopt, or having bugs there, and especially so if
there was uninlining involved. It was a case of "if you had to be clever to
write the code, you better be having a darn good day if you're going to debug
it". Only ever touching the stack top, rather than dabbling with a linked
list-ish thing, and recreating the frames outermost to innermost (it was the
other way round before) comes out *vastly* simpler. As in, I can probably
understand what the code is doing even on one of my average days now. This
is welcome. Oh, and we no longer have to heap promote frames because we're doing a
deoptimization, which makes speculative optimization cheaper in the caess reality
doesn't match up with our speculations.

All in all, I feel this is a solid start to the dispatcher work. Of course,
there's quite a bit more between this and something we're shipping in
production releases; I'll be working to close that gap in the month ahead.


```
72:17 Implementation work on new dispatch mechanism, as described above

Total reported time in this report: 72:17
Remaining approved hours on current grant period: 54:04
```
