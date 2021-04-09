
---
title: "RakuAST Grant Report from Jonathan Worthington, 2020-09"
author: Moritz Lenz
type: post
date: 2020-10-06 00:00:00 +0000 UTC
url: "/post/jnthn-grant-report-rakuast-2020-09"
categories:
 - Grants

---

Jonathan Worthington has shared his latest progress report from the RakuAST grant, covering the work done in 2020-09.

In addition, he announced that [he is taking a break from Raku and Rakudo development](https://6guts.wordpress.com/2020/10/05/taking-a-break-from-raku-core-development/). His work his highly appreciated, the resting period well-desevered, and we expect his work to continue in a month or a few months.

Here's his progress report, in his own words:

---

In [last month's report](https://news.perlfoundation.org/post/grant-report-jonathan-worthington-rakuast-2020-08), I mentioned
that the RakuAST-based compiler frontend had reached the point of being able to attempt the specification test suite, thanks
to it gaining support for basic `use` statements. For now, it depends on using the standard compiler frontend to precompile the test modules. Back then, it was capable of passing 39 spectest files in full.

Fast forward a month, and that number stands at 209 passing spectest files. On the one hand, that's quite a step forward.
On the other hand, it means there's still quite a bit to go. Part of the challenge here is that spectest files may aim to
test a particular feature, but since a lot of interesting behaviors involve the interactions of features, they often end
up touching quite a lot of the language. Thus implementing support for a particular language construct's compilation via
the new Raku AST does not automatically mean we're now going to be able to pass the spectest file(s) that are most
obviously associated with it. It may, however, mean we unblock some other spectest files that aren't directly about the
construct, but were blocking on it! For example, implementing support for the `.WHAT` type introspection operation resulted in a bunch of new passing spectest files. It's also not just about feature interactions. Tests are code, and people like to factor
code nicely, and that can also entail use of further language constructs beyond those under test.

Here's a summary of the advances made over the last month; generally, when a language construct is mentioned, then RakuAST nodes were developed and also the work was done to construct these in the new RakuAST-based compiler frontend.

* Get multi-part package lookups of types working
* Add the Pod grammar to the new compiler frontend (which won quite a bunch of spectests that contained Pod comments)
* Get single-part package name installation working, so that simple package declarations result in installation of a type name that can be resolved; the key realization here is that installation in a package is most naturally handled as a `BEGIN`-time side-effect
* Get method definitions working, including production of the implicit invocant parameter
* Implement the `self` term and the `.foo` term for calls on the topic
* Implement basic attribute declaration and attribute access
* Implement `.=` (call method and assign) and `.` (call method) infixes; these are interesting in that they parse a postfix after them rather than a term, and this also needed modeling in the AST
* Implement the terms `now`, `time`, `rand`, empty set, `*`, and `**` terms (however not yet the whatever-currying semantics)
* Implement lookups of named constant terms (`True` and `False` being the most obvious examples)
* Implement all forms of colonpair (so far, only as terms rather than as part of a name)
* Wire up compilation of dynamic variables and match variables to the already existing AST nodes
* Implement the literal hash index (`%foo<x>`) and call (`$foo()`) postcircumfixes
* Make `EVAL` work when running the RakuAST-based frontend, so we can use RakuAST to compile a program that calls `EVAL` that uses the RakuAST-based compiler
* Make the new compiler frontend parse and pass along types on lexical variables (the AST already supported this)
* Implement term (`\foo`), 