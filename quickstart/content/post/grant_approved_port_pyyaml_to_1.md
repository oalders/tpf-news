
---
title: "Grant approved:  Port PyYAML to Perl"
author: Curtis "Ovid" Poe
type: post
date: 2006-08-22 00:00:00 +0000 UTC
url: "/post/grant_approved_port_pyyaml_to_1"
categories:
 - Grants

---

Here's Brian Ingerson's "Port PyYAML to Perl" grant application.

Name: Ingy dÃƒÂ¶t Net
Email: ingy@ttul.org
Project Title: Port PyYAML to Perl
Amount Requested: $3000

Synopsis: >
PyYAML(1) is a pure Python implementation of the YAML 1.1
specification(2). It is considered the highest quality YAML
implementation to date. It has an excellent API (including streaming
SAX-like support), a comprehensive test suite, full unicode support,
very good tagging/type support, follows the spec precisely and is
written in a style that makes it easy to port. It has already been
ported to pure Ruby(3), and is considered by the YAML community to be
the standard reference implementation. It would greatly benefit both
the YAML and Perl communities to port this code to a Perl module.

Benefits to the Perl Community: |
* A solid, dependable pure Perl implementation of YAML
* Ability to share YAML reliably with Ruby and Python
* A streaming interface as well as the Dump/Load API
* Full support for YAML implicit types
* More control for how YAML objects get serialized
* Better error messages

YAML is currently used for many important Perl projects such as SVK(4),
Plagger(5), Catalyst(6), Jifty(7) and Kwiki(8) to name a few. It is used
to store CPAN metadata(9). Currently most of these use cases are small
and simple config files.

YAML was designed to scale to much more elaborate use cases. For
instance, the Perl 6(10) community, where it is used as a kind of Perl
6 bytecode(11) to transport a Pugs AST from Haskell to Perl 6, to Perl
5 and back.

For YAML to be used as a solid platform for flawless data transport, it
needs rock solid implementations. This project will deliver that.

Deliverables: |
* Port all the classes of PyYAML to Perl
* Port all of the PyYAML tests and make them pass
* Write full documentation for all the classes and the API
* Release this module as YAML.pm, replacing the legacy code

Relevant History: >
The YAML serialization project was started in May 2001 with the intent
of bringing a simple yet complete, plain text, human friendly data
serialization language to all programming languages. It was designed
with an information module that very closely follows the Perl data
model of potentially tagged (blessed) scalars, arrays and hashes with
potentially cyclic references. This makes it an excellent choice for
languages like Perl, Python, Ruby and PHP.

The language took several years to design, has a very complete
specification document and has had two major versions, 1.0 and 1.1.
The YAML language specification 1.1 has been fairly stable for over a
year and a half now.

In late 2001, Brian Ingerson wrote the original YAML.pm(12) Perl
implementation. This is still the primary implementation used by the
Perl world today. Although it is generally useful it has many
problems. Its parser is very incomplete. It parses a subset of YAML
1.0. It generally can parse what it emits, but there are bugs in that
regard too. Its type resolution is severely lacking. The code is
rather poorly written, and lacks significantly in its test suite.

In early 2002, Brian talked coworker Steve Howell into writing an
implementation for Python. Dubbed PyYaml, it was indeed written, but
was generally considered even worse than the YAML.pm. For years,
Python has had very lacking YAML support.

Later in 2002, Why The Lucky Stiff began writing a Ruby
implementation. The Ruby community rallied behind him, and YAML
started showing up all over the place in Ruby projects. In 2003, Why
decided to write a C implementation called Syck(13), aka libsyck. This
library was written following the YAML 1.0 spec. It was amazingly
fast, and was soon ported to Python, PHP and OCaml. Syck became the
YAML codebase that is distributed the Ruby itself.

In 2004, the YAML spec went through some significant change, and was
thus dubbed YAML 1.1. While the basic syntax remained the same, the
tagging system and directive declaration was revamped.

In early 2006, Audrey Tang ported Syck to Perl as YAML::Syck(14). The
code is very fast and offers improvement over YAML.pm, but it too has
many problems. It is still YAML 1.0, with its own set of bugs. It
doesn't produce good error messages. It ocassionally segfaults. This
is almost entirely the fault of libsyck, not YAML::Syck. It also lacks
most of the advanced features of YAML.pm, which allow more control of
how objects get serialized.

In Spring of 2006, Kirill Simonov, began a project code named PyYAML
3000. It later simply took on the name PyYAML, replacing the old code
base. This new Python code followed the spec in every way, and even
exposed a few bugs in the spec. It was intended from the start to be a
reference implementation, and is currently considered the best
implementation available.

Also in Spring of 2006, Ola Bini ported PyYAML to Ruby, giving Ruby
the same benefits that Python was now enjoying. This left Perl as the
language with the worst YAML support. A sad thing.

In June 2006, Kirill Simonov was awarded a Google Summer of Code grant(15)
to write a C implementation of YAML tools and bind them to Python.
This code will eventually become libyaml, and will be bound to many
languages, including Perl.

NOTE: My proposal does not include porting libyaml (which doesn't even
exist yet) to Perl. Although when it is ready, I will likely
distribute a binding of it with YAML.pm.

Project Details: >
Most users of YAML and YAML.pm are only aware of the simple Dump/Load
interface, but in reality there is a stack of processing that happens.
PyYAML gives programming access to each of those layers of the YAML
stack(16).

The Dump stack looks like this:
* Dumper
* Representer
* Serializer
* Emitter

The Load stack looks like this:
* Loader
* Constructor
* Composer
* Parser
* Scanner

PyYAML has classes for all these processes. These classes will be
ported as accurately as possible to Perl while taking defacto standard
Perl coding styles into account.

The resulting module distribution will have no dependency
modules/distributions if possible. It will likely use Test::Base for
testing, but will be packaged using the Module::Install method that
does not require Test::Base to be installed.

The current YAML.pm has a YAML::Node generic node class that is used
as a meta object under certain circumstances, to acheive certain goals
such as customized serialization. The new project will attempt to make
this, and other details like this, backwards compatible with the
current YAML.pm.

This work is intended to be done in the open and will have a
publically readable Subversion repository(17), a Trac, a mailing
list(18) and an irc channel(19).

Project Schedule: >
This project should take me a month or two. I can begin working on it
immediately.

Bio: >
Ingy dÃƒÂ¶t Net is a prolific Perl module author and one of the three
people who designed the YAML serialization language. He wrote the
original YAML.pm, which has served the community fairly well, but is
severely lacking in many ways. Ingy is intimately aware of both the
serialization needs of Perl and of the multilayered architecture of a
properly written YAML implementation.

Note: I realize this proposal seemingly contradicts what I once said in
my use.perl.org journal(20), but... what the heck!

References:
- 1) http://pyyaml.org/
- 2) http://yaml.org/spec/cvs/current.html
- 3) http://rbyaml.rubyforge.org/
- 4) http://svk.elixus.org/
- 5) http://plagger.org/trac
- 6) http://www.catalystframework.org/
- 7) http://www.jifty.org/view/HomePage
- 8) http://search.cpan.org/dist/Kwiki/
- 9) http://search.cpan.org/src/INGY/YAML-0.62/META.yml
- 10) http://dev.perl.org/perl6/
- 11) http://perlsoc2006.blogspot.com/2006/06/progress-marches-on.html
- 12) http://search.cpan.org/dist/YAML/
- 13) http://whytheluckystiff.net/syck/
- 14) http://search.cpan.org/dist/YAML-Syck/
- 15) http://code.google.com/soc/psf/appinfo.html?csaid=75A3CBE3EC4B3DB2
- 16) http://yaml.org/spec/cvs/current.html#id859109
- 17) http://svn.kwiki.org/
- 18) https://lists.sourceforge.net/lists/listinfo/yaml-core
- 19) irc://irc.freenode.net/#yaml
- 20) http://use.perl.org/~ingy/journal/

