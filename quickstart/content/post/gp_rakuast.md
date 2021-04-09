
---
title: "Grant Proposal: RakuAST"
author: Coke
type: post
date: 2020-04-18 00:00:00 +0000 UTC
url: "/post/gp_rakuast"
categories:
 - Grants
 - Sign in

---

# Project title

RakuAST

# Project synopsis

I recently presented a proposal for RakuAST ([slides](http://jnthn.net/papers/2020-gpw-realizing-raku-macros.pdf),
[video](http://jnthn.net/papers/2020-gpw-realizing-raku-macros.mp4)), an
abstract syntax tree for the Raku language that will become part of the
language specification. I have also released my work on it so far, which is in
a [branch](https://github.com/rakudo/rakudo/tree/rakuast) in the Rakudo
repository.

So far, I've worked on it during spare moments. The purpose of this grant is
to accelerate progress by enabling me to spend more of my working time on it.
While I already am funded through an ongoing Performance and Reliability
grant - which I shall continue to work on - the RakuAST work does not naturally
fit within its scope, thus this grant application.

# Your name

Jonathan Worthington

# A discussion on the benefits of the project to the Perl Community

An AST can be thought of as a document object model for a programming language.
The goal of RakuAST is to provide an AST that is part of the Raku language
specification, and thus can be relied upon by the language user. Such an AST
is a prerequisite for a useful implementation of macros that actually solve
practical problems, but also offers further powerful opportunities for the
module developer. For example:

* Modules that use Raku as a translation target (for example, `ECMA262Regex`,
  a dependency of `JSON::Schema`) can produce a tree representation to `EVAL`
  rather than a string. This is more efficient, more secure, and more robust.
  (In the standard library, this could also be used to realize a more
  efficient `sprintf` implementation.)
* A web framework such as `Cro` could obtain program elements involved in
  validation, and translate a typical subset of them into JavaScript (or
  patterns for the HTML5 `pattern` attribute) to provide client side
  validation automatically.

RakuAST will also become the initial internal representation of Raku programs
used by Rakudo itself. That in turn gives an opportunity to improve the
compiler. The frontend compiler architecture of Rakudo has changed little in
the last 10 years. Naturally, those working on it have learned a few things in
that time, and implementing RakuAST provides a chance to fold those learnings
into the compiler. Better static optimization, use of parallel processing in
the compiler, and improvements to memory and time efficiency are all quite
reasonable expectations.

However, before many of those benefits can be realized, the work of designing
and implementing RakuAST, such that the object model covers the entire semantic
and declarational space of the language, must take place. This grant focuses
on that work.

# A list of deliverables

1. Class and role implementations defining an document object model for the
   Raku language, including both the regex and string sub-languages. These
   should all be constructable and introspectable from within the Raku
   language.

2. Generation of QAST, the backend-independent intermediate representation,
   from RakuAST nodes, such that one can execute an AST. This functionality is
   exposed as an overload of `EVAL`, since it's ultimately the same operation;
   the RakuAST form of `EVAL` can skip the parsing stage entirely, since it
   already has the AST.

3. Tests that cover `EVAL` of RakuAST. Every node should be test covered.

4. Integration of RakuAST into the compilation process. This functionality
   can be behind an environment variable rather being the default mode of
   compilation; this is to avoid a situation where there is pressure to
   enable it by default in order to consider the grant complete, when it
   would serve Raku users better to make the switch when it is really
   ready for that. However, it should be good enough to pass at least
   90% of the test suite on the MoarVM backend.

# Project details and a propose