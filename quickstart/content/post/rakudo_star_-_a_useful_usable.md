
---
title: "Rakudo Star - a useful, usable, \"early adopter\" distribution of Perl 6"
author: Karen Pauley
type: post
date: 2010-07-29 00:00:00 +0000 UTC
url: "/post/rakudo_star_-_a_useful_usable"
categories:
 - Perl 6 Development
 - Sign in

---

On behalf of the Rakudo and Perl 6 development teams, I'm happy to announce the July 2010 release of "Rakudo Star", a useful and usable distribution of Perl 6.  The tarball for the July 2010 release is available from <http://github.com/rakudo/star/downloads>.

Rakudo Star is aimed at "early adopters" of Perl 6.  We know that it still has some bugs, it is far slower than it ought to be, and there are some advanced pieces of the Perl 6 language specification that aren't implemented yet.  But Rakudo Perl 6 in its current form is also proving to be viable (and fun) for developing applications and exploring a great new language.  These "Star" releases are intended to make Perl 6 more widely available to programmers, grow the Perl 6 codebase, and gain additional end-user feedback about the Perl 6 language and Rakudo's implementation of it.

In the Perl 6 world, we make a distinction between the language ("Perl 6") and specific implementations of the language such as "Rakudo Perl".  "Rakudo Star" is a distribution that includes release #31 of the [Rakudo Perl 6 "compiler](http://github.com/rakudo/rakudo), version 2.6.0 of the [Parrot Virtual Machine](http://parrot.org/), and various modules, documentation, and other resources collected from the Perl 6 community.  We plan to make Rakudo Star releases on a monthly schedule, with occasional special releases in response to important bugfixes or changes.

Some of the many cool Perl 6 features that are available in this release of Rakudo Star:

  * Perl 6 grammars and regexes
  * formal parameter lists and signatures
  * metaoperators
  * gradual typing
  * a powerful object model, including roles and classes
  * lazy list evaluation
  * multiple dispatch
  * smart matching
  * junctions and autothreading
  * operator overloading (limited forms for now)
  * introspection
  * currying
  * a rich library of builtin operators, functions, and types
  * an interactive read-evaluation-print loop
  * Unicode at the codepoint level
  * resumable exceptions

There are some key features of Perl 6 that Rakudo Star does not yet handle appropriately, although they will appear in upcoming releases.  Thus, we do not consider Rakudo Star to be a "Perl 6.0.0" or "1.0" release.

In many places we've tried to make Rakudo smart enough to inform the programmer that a given feature isn't implemented, but there are many that we've missed.  Bug reports about missing and broken features are welcomed.

See <http://perl6.org/> for links to much more information about  Perl 6, including documentation, example code, tutorials, reference materials, specification documents, and other supporting resources. Rakudo Star also contains a draft of a Perl 6 book -- see  <docs/UsingPerl6-draft.pdf> in the release tarball.

The development team thanks all of the contributors and sponsors for making Rakudo Star possible.  If you would like to contribute, see <http://rakudo.org/how-to-help>, ask on the perl6-compiler@perl.org mailing list, or join us on IRC #perl6 on freenode.

Rakudo Star releases are created on a monthly cycle or as needed in response to important bug fixes or improvements.  The next planned release of Rakudo Star will be on August 24, 2010.

Editor's notes

For questions, contact Perl Foundation Public Relations at pr@perlfoundation.org.

[Perl](http://www.perl.org)

Perl is a dynamic programming language created by Larry Wall and first released in 1987. Perl borrows features from a variety of other languages including C, shell scripting (sh), AWK, sed and Lisp. It is distributed with practically every version of Unix available and runs on a huge number of platforms, as diverse as Windows, Mac OS X, Solaris, z/OS, os400, QNX and Symbian.

[The Perl Foundation](http://www.perlfoundation.org/)

The Perl Foundation is dedicated to the advancement of the Perl progr