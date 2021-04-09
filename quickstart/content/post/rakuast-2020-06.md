
---
title: "June 2020 Grant Report for RakuAST by Jonathan Worthington"
author: Moritz Lenz
type: post
date: 2020-06-19 00:00:00 +0000 UTC
url: "/post/rakuast-2020-06"
categories:
 - Grants
 - Raku Development

---

Jonathan writes the following as his first report for his RakuAST grant. Enjoy!

---

It's been about a month since the approval of my [RakuAST grant](https://news.perlfoundation.org/post/gp_rakuast). This report covers the grant work performed since then.

Under the RakuAST architecture, language elements are modeled by AST nodes, which can be thought of as document object model elements, where the document in question is a Raku program. Today, the overall organization of the compiler part of Rakudo involves:

* A parser, implemented as a Raku grammar
* The actions, which build up QAST - a lower-level AST, thus defining the  execution aspects of the language. This is quite a large gap to bridge.
* The world, which deals the declarational aspects of the language. There's  a lot of those, so it's also a rather large piece of code.
* An optimizer, which works on the QAST tree

A large part of the Raku codebase consists of implementations of meta-objects and the standard library, which will remain almost entirely the same.

Under the new RakuAST architecture, each type of AST nodes captures all of the non-syntactic knowledge about that kind of program element. This includes:

* How it evaluates (defined by producing the appropriate QAST, typically  incorporating that produced by child nodes)
* How any related meta-objects are produced
* How it is resolved, if it involves some kind of symbol lookup
* How it is checked for semantic problems
* How it is optimized

Before the grant commenced, I had been exploring these ideas, and hadimplemented a small number of AST nodes. Since the start of the grant, I have expanded that to cover:

* Sub and method calls
* Arguments, including named arguments and flattening
* The fat arrow construction (which is currently the only way to pass a named argument; colonpairs will come later)
* The `while`, `until`, `repeat while`, `repeat until`, and `loop` loop
  constructs
* The `for` iteration construct
* The `if`/`with`/`elsif`/`orwith`/`else` conditional construct
* The `without` conditional construct
* Array and hash indexers (`@a[1]`, `%h{'x'}`)
* Parenthesized expressions and the `[...]` array composer
* A range of infix operators that need special compilation, rather than  simply being a function call

I also worked on how to model sink context marking (also known as void context in various other languages). The current code doing this work in Rakudo is difficult to follow and not terribly efficient. That's in no small part because it is done as a pass after the QAST construction, and has to try and recover a lot of information and tweak the QAST in order to do the sinking required. Since RakuAST models the language at a higher level and defers producing QAST until much later, a far cleaner solution to the sink analysis problem is possible. As with the rest of the RakuAST design, nodes themselves know if they are interested in being sunk and how to propagate sink context. I'm optimistic that the model I've created will be flexible enough to handle all the sinking-related requirements - the proof
will be in passing the test suite, of course.

All of this is important, but what of the big goal I set out for June in the grant application? As a reminder, I wrote this:

> In June I would start work on a means to get the compiler frontend to use RakuAST, most probably enabled by an environment variable. This need not deal with such concerns as BEGIN, module usage, or precompilation, but it should be possible to run various simple scripts in this mode successfully.

At present, this is all achieved, except that "simple" should probably be read as "very simple". That's fine, because I've still got the rest of June to make it cover a wider range of language constructs.

After spending some time looking at all of the ways the current grammar is tied up with the actions and world, I decided the best way forward would be to build that up "from scratch" too. The current grammar, actions,
world, and optimizer live under `src/Perl6/`; we'd never got to doing the rename there yet. And we won't: I'm building the RakuAST-based grammar, actions, and AST nodes under `src/Raku`. If one sets `RAKUDO_RAKUAST=1` in the environment, the new compiler implementation is used. Eventually, we'll
switch over to the RakuAST-based implementation, and drop the current one.

Building up the "new" parser is mostly a case of copying from the current one, chopping out various pieces that were tied into the previous compilation
model and, where needed, substituting them with the new approach. In some cases, there is no substitution; things - at least at the parser level - get simpler. Since RakuAST can be constructed synthetically, a much
stricter syntactic/semantic division is required.

The new actions are vastly more straightforward, because they are building up an AST that is far closer to the language being parsed. It wouldn't surprise me if the actions end up between a quarter and a third of their current number of lines of code. Of course, the logic that used to be in the actions mostly has to go somewhere else - that somewhere being the nodes.
I'm finding that even then it *still* ends up being a net simplification, because by the time we produce QAST nodes we're in a position to produce the correct thing immediately, rather than producing something that has to be tweaked later.

The monolithic `World` object dealing in declarations has no direct successor in the RakuAST-based architecture. It's spread over:

* The AST nodes themselves, which deal with the production of meta-objects   and some aspects of symbol resolution
* The resolver, which keeps track of things like "which scope are we in  now", so we know where to go looking for symbols when we want to do  resolutions. There are two resolver implementations: a "batch" one for  when we `EVAL` an AST, and a "live" one that is used when we are parsing  Raku source code and building up the AST as we go. The needs are a bit  different: in an AST, there isn't really a BEGIN time, nor are there any  syntactic ambiguities that need symbol lookups to resolve.
* A constant table, for building and interning literals

It's probable that some further pieces will be discovered when I reach dealing with the `BEGIN`-time things.

So far, the RakuAST-based compiler passes a number of the sanity tests in the Rakudo test suite. I plan to work though those and get them all to
compile successfully - however, doing so will probably entail a slight course correction. I didn't really plan to work on the quoting and regex aspects of the language until July. However, even a number of the basic
sanity tests rely on interpolation and quote words - so I'll probably end up dealing with a bunch of the quoting-related bits during June. I'd also not really expected to start dealing with meta-operators so soon, but things
like `~=` and `+=` also show up really quite early in the sanity tests.

Finally, it's interesting to note that RakuAST-based compilation comes out faster - at least for the tiny test scripts I'm able to try out now. I know it's some mix of:

* The initialization of the compilation phase being O(1) in the number of   top-level symbols in CORE.setting (the built-ins), rather than O(n) in   the current compiler. This win is constant in the size of the code being  compiled; it's thus best considered as a startup time improvement (though  you'll win it for every `EVAL` or module dependency that has to be  re-compiled too.) This improvement will be retained, so we can cross our  fingers that it's a dominant factor in the speedup I'm observing.
* The grammar being far smaller, and the parser engine having less  initialization to do. This won't remain the case, so let's hope it's a  small part of the speed-up observed (but if it isn't, it still is a useful  data point.)
* Potentially, design improvements meaning the compilation is more  efficient. I don't think the tiny test scripts I'm running so far are  creating enough work for this to be a significant factor, however.

It will be interesting to see how the performance situation develops as I go on; it's always nice to see smaller times, but it's too early to read too
much into them yet.

All in all, this feels like a good start on the grant. There's much to come, but it's exciting to already have some basic things running on the new RakuAST-based compiler architecture.

