
---
title: "July 2020 grant report from Jonathan Worthington on RakuAST"
author: Moritz Lenz
type: post
date: 2020-07-23 00:00:00 +0000 UTC
url: "/post/2020-07-grant-report-jnthn-rakuast"
categories:

---

Jonathan has been working hard on his [RakuAST grant](https://news.perlfoundation.org/post/gp_rakuast), and made some great progress. He'll be pausing development for roughly two weeks for some well-earned vacation.

Anyway, here is the progress report, in his own words:

---

Raku is not really one language, but a braid of languages: the main language, the regex language, and the quoting language. My work as of my [previous report](https://news.perlfoundation.org/post/rakuast-2020-06) had focused heavily on the main language, dipped into the quoting language only a little bit, and completely ignored the regex language. During the last month, I've made progress on all three.

The quoting language is the simplest of the three, although still quite rich with its various forms of interpolation. I created new RakuAST nodes as needed and wired them up to the new RakuAST-based compiler frontend, so as to enable handling of:

* Variable interpolation
* Closure interpolation
* Most of the quote "processors": words (split on whitespace, used for quote words), val (produce `IntStr`, `RatStr`, and friends where applicable, as used in the `<foo 42 bar>` construct), and exec (used for `qx`); these can also be combined in various ways.

I also implemented some amount of the constant folding possible here.

The major focus of the month's work was on the regex language. While I'm not at full coverage of it yet, I've got AST nodes that model a significant number of regex constructs:

* Alternation (LTM and sequential)
* Conjunction
* Literals, including quoted literals
* Groups (`[...]`)
* Capturing groups (`(...)`)
* Anchors (`^`, `$`, `^^`, `$$`, `<<`, `>>`)
* Match start and end markers (`<(`, `)>`)
* Some built-in character classes (`\w`, `\d`, `\n`, `\s`, and their negations)
* The pass and fail assertions (`<?>` and `<!>`)
* Named rule assertions, both capturing and non-capturing and with arguments (`<alpha>`, `<before foo>`, `<my-rule('with args')>`, `<.ident>`)
* Aliasing assertions (`<name=.ident>`)
* Lookahead assertions (`<?foo>`, `<!foo>`)
* The `?`, `+`, and `*` quantifiers along with the backtrack modifiers
* Quantifier separators (`%` and `%%`)

There are test cases covering the construction of these through the AST, as well as them being wired up to the RakuAST-based compiler frontend, meaning that the regex sanity test now passes (and plenty more regex tests would too, if only we could run them, but we can't until the new compiler frontend can load modules).

Work on RakuAST nodes to model the main language also moved along quite nicely. The main language presents the most significant modeling challenges. The regex and quoting languages are declarative in nature, and the information needed to compile an element is either held within the node, can be found by looking at immediate children, or can be passed down from a parent at the point we are compiling the nodes into a lower-level AST form. With the main language, things are not so simple.

Much hinges on how much of a burden we wish to place upon those constructing the AST. As an example, consider these three constructs that involve blocks:

    for 1..10 { .say }
    if $password eq 'correcthorsebatterystaple' { say "correct" }
    with get-result() { .say }

In all cases, the block is a `RakuAST::Block`. However, the first and the third cases have the block taking an implicit parameter, `$_`, while the block to the `if` does not. This does not solely affect the compilation of the block, but also the production of a `Signature` object.

One solution would be to require the `RakuAST::Block` be constructed with some kind of flag indicating whether or not it received a topic. This means the onus is on those constructing the AST to create the node correctly (or tweak the node appropriately after creating it). We might argue that this isn't so problematic: the compiler actions to produce it are written once, and there will be quasiquoting (effectively, producing an AST by writing the code that maps to it), so making the AST nodes directly by object creation is not the most common case. However, this overlooks some things:

* ASTs may also be transformed. For example, we might do a code transformation from an `if` to a `with` or vice versa. That would also be burdened with having to go and tweak the block object too.
* It allows us to represent things in the AST that don't map back to something you could do in code. I'm not outright trying to forbid every case of this, but I don't consider it desirable either.
* A general design principle I'm trying to follow is to have the new compiler frontend construct the right things in the first place, rather than tweak things it has already built. This makes the code more functional and easier to reason about, and hopefully a bit more efficient too. The design where the block carries the flag would imply we have to modify the block node.

None of these things alone is a huge deal, but together they point away from such a design. Instead, I picked a resolution-time solution. Resolve time could maybe be thought of as "graph time", because it's the time that we take what is outerly a tree and have various nodes connect themselves together in non-tree-like ways. For example, variable usages are linked to their declaration. Anyway, that's also a natural time for nodes that want to apply *implicit semantics* to a block child to do so, and thus it happens during that phase. 

I also introduced an "attach" mechanism that operates during resolution time. While the implicit block semantics mechanism is aimed at a fairly particular problem (and I may even spot a generalization in the future), the attach mechanism is aimed at things that want to attach themselves to some enclosing scope. For example, I've used to so far to have:

* `method`s attach to the enclosing package, so the meta-object can be formed properly
* `END` blocks attach to the enclosing compilation unit
* `CATCH` and `CONTROL` exceptions attach to the enclosing block
* `when` and `default` attach to (or at least make their existence known to) the enclosing block, in order that it can produce the appropriate `succeed` control exception handler

All of which are main language constructs that I implemented RakuAST for this month (and wired up in the compiler too).

Finally, here's a list of other main language constructs I also got RakuAST nodes for in place during the last month:

* Infix assignment meta-ops (`+=`)
* Bare blocks, which are automatically invoked when encountered
* Named and anonymous subroutines, which entailed introducing support for implicit declarations (to handle the per-block `$!` and `$/`)
* `our`-scoped variables, which entailed doing a decent amount of support for package scoping
* Numerous statement prefixes: `do`, `gather`, `quietly`, `lazy`, `eager`, `race`, `hyper` and some of `try` (it's missing its fatalize semantics still)

Working on RakuAST is causing me to revisit pretty much the entire Raku language and consider the semantics of each language element (noting that a huge amount of what we informally consider as "the Raku language" is actually its standard library, which I'm not really touching on at all). With all the knowledge of a language construct moving into a class modeling that construct, quite a few questions arise about where things belong. While we might well have an actions *class* in the current compiler frontend, doing a real object model for the language and its semantics shows up just how procedural the current compiler frontend code is.

That I'm finding the move to a genuinely OO design a challenge is, honestly, quite reassuring to me. Back when I was teaching software development, I often said that good OO design (rather than a procedural design spread over various classes) is difficult, and usually requires digging deeply into the domain. Ultimately, it's a modeling exercise, and the work on RakuAST sees me crunching years of learnings about the Raku language and compilers into a new way to model Raku programs. I hope the result will lead to as much `-Ofun` as I'm having on the journey.



