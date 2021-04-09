
---
title: "Grant Proposal: Implement Perl Binding for libuv"
author: Coke
type: post
date: 2019-11-11 00:00:00 +0000 UTC
url: "/post/gp-perl5bindinguv"
categories:
 - Grants
 - Sign in

---

# UV

- Name:

    Paul Evans

- Amount Requested:

    USD 3,698 (*Converted at time of submission from GBP*)

## Synopsis

Implement a Perl binding for [libuv](https://libuv.org), wrapping as many of
the features and types as is practical and useful for Perl 5.

## Benefits to the Perl Community

The `libuv` library provides a multi-platform event system, and is the basis
for the `nodejs` JavaScript and `moarvm` Perl 6 runtimes and the `neovim`
text editor, to give a few examples. It is widely used as a support library by
a wide variety of software and has language bindings for C#, Python 3, C++ and
Lua, as well as being accessible directly via NodeJS. Adding Perl 5 to this
list will allow programmers familiar with it from other langauges to make use
of it from Perl 5, and will bring to Perl 5 several useful abilities the
library provides that most other event systems do not.

## Deliverables

The primary deliverable would be a newer version of [UV](https://metacpan.org/pod/UV) whose internal
implementation has been reworked to have a better architecture to more easily
maintain going forward and that eliminates the current leaks and hangs on
polling.

A number of known bugs already exist and a few are collected on GitHub (at
[https://github.com/genio/p5-UV/issues](https://github.com/genio/p5-UV/issues)). Other issues are collected on the
[Mojo::Reactor::UV](https://metacpan.org/pod/Mojo::Reactor::UV) module, which could be considered to be a real-world
example usage: [https://github.com/Grinnz/Mojo-Reactor-UV/issues](https://github.com/Grinnz/Mojo-Reactor-UV/issues).

In addition to that first release, subsequent releases would be made to add
the missing features of the current implementation, namely requests:
[http://docs.libuv.org/en/v1.x/request.html](http://docs.libuv.org/en/v1.x/request.html). The currently missing handle
objects would also be added.

Upon fully implementing the features of the library, a secondary deliverable
will be a set of blog posts and presentations to educate potential users,
and to demonstrate its use. Some blog posts may draw contrasts and
similarities to other languages, helping to further emphasize the impact of
this cross-platform library in all of these languages. In particular,
parallels to Perl 6, Python 3, and NodeJS may be useful as those languages
are ones that potential users may also be familiar with.

## Project Details

While a binding already exists in the namespace [UV](https://metacpan.org/pod/UV), it suffers many
shortcomings and bugs. The current maintainer is willing to work with me on
rebuilding it in a better form, now that the existing structure is more
understood. We should be able to preserve much of the existing documentation
and unit tests, which will be helpful in getting a good start.

## Inch-stones

The heart of [UV](https://metacpan.org/pod/UV), [UV::Loop](https://metacpan.org/pod/UV::Loop) has been implemented, but is in need of a
rewrite to fix it and clear up the bugs.

Once this is done, the various handle types need to be wrapped; around 15
different types in total. Writing a good implementation for any one of them is
likely to take an amount of time and research, but from that point the
remaining ones should follow along somewhat easier, as they all follow a
similar pattern.

## Project Schedule

There are two main phases to this project.

The first phase will involve rebuilding the XS implementation of the current
portion of `libuv` which has already been attempted, sufficient to cover the
existing unit tests for the loop itself and the currently-implemented handle
types (`check`, `handle`, `idle`, `poll`, `prepare` and `timer`). As
mentioned, it is likely that the existing documentation and unit tests of
these can be retained and provided with a better code implementation. The
tests will help illustrate that the implementation has been completed.

This first ph