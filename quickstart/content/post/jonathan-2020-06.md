
---
title: "June report of the Raku Development Grant of Jonathan Worthington"
author: Matthias Bloch
type: post
date: 2020-07-10 00:00:00 +0000 UTC
url: "/post/jonathan-2020-06"
categories:
 - Raku Development

---

Jonathan writes:

My grant time in June was spent on the new unified dispatch mechanism; my [May report](https://news.perlfoundation.org/post/jonathan-2020-05) provides some background on what that is, and this report will make more sense having read the May report first. In May, I reached the point of having a significant part of the new dispatch mechanism implemented, and had exercised it with some unit tests. However, these were all rather artificial. During June, I started applying it to rather more realistic problems.

Today, Rakudo (the Raku compiler and standard library) uses a mechanism called "specializer plugins", which allow it to "explain" a range of language constructs to MoarVM (the virtual machine), in order that it can optimize them decently well. In fact, the overall approach I explored in specializer plugins has been hugely informative for the design and implementation of the generalized dispatch mechanism that I am currently working on. I decided that replacing the specializer plugins with the new generalized dispatch mechanism was a good place to start. Of note, I replaced (in a branch):

* The return value type check plugin. The new generalized dispatch mechanism will allow for a shorter instruction sequence for these checks in the bytecode. However, there's another advantage. Specializer plugins work by looking at a number of arguments and mapping them - based on their matching of some guards - to something to invoke. In the case where we replace the type check with an exact type guard, the thing to invoke is the identity function. This invocation has overhead. We then rely on inlining to inline the identity function to eliminate that cost. Thus on hot paths we eventually have quite nicely optimized code. However, it takes time for the optimizer to kick in, and inlining costs something too. With the new dispatcher, we can do identity results without having to invoke anything, and the optimizer will be able to deal with them much more cheaply than using the inlining machinery too.
* The return value container handling plugin. It also often wants to just return identity, so gets the benefits described above. And in the case it wants to decontainerize a `Scalar`, we don't need an invocation either, since dispatch programs can read attributes. (It's not done yet, but if it's a `Proxy`, we shall be able to wire the dispatch directly to a call to its `FETCH` function.)
* The private method call plugin. Again, these come out with a shorter bytecode sequence.
* The qualified method call plugin. Guess what? Yup, a shorter bytecode sequence for these too. Also, error handling is done in the dispatcher, rather than needing a further delegation.
* The maybe method call plugin (for `.?foo`). This produces `Nil` if there is no such method, which is done rather cheaply done now (again, thanks to being able to produce value results without an invoke). Shorter bytecode too.
* The assign plugin. In fact, this is only partially complete, and still missing its optimized versions. Shorter bytecode again.

Return value type checks can involve coercion types, which means that we don't just need to check the type, but also to coerce it into another one. Since this involves a method call or potentially even a multi method call, this required me to start looking at implementing those too - ahead of switching all normal calls over to them. Standard method resolution is quite straightforward. Multiple dispatch, on the other hand, is rather more interesting.

Rather than wedge the current multi-dispatcher into the new dispatcher with minimal changes, I've decided to revisit how it works to take advantage of the new opportunities. This is still a work in progress. The main change is that I'm factoring it as a two-step process: 

1. Perform a pre-filtering step on the candidates based on properties that we can guard