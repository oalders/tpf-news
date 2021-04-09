
---
title: "Parrot Grant Update for March - August, 2008"
author: Dave Rolsky
type: post
date: 2008-09-02 00:00:00 +0000 UTC
url: "/post/parrot_grant_update_for_march"
categories:
 - Grants

---

Apologies for the long delay in this report, I've had some ongoing health problems over the summer. Allison wrote this report to bring us up-to-date, and then we'll continue with the two-monthly schedule.

Parrot releases continue on their monthly schedule. The 0.6.0 release in March reflected major improvements to both the Parrot VM and the Perl 6 implementation on Parrot (Rakudo). The most notable change for Parrot was the implementation of the PMC design (PDD17), which was done by Allison and chromatic. The 0.6.1 release in April included significant enhancements to the Parrot Compiler Toolkit (PCT), as well as even more improvement in Rakudo. It also included a number of speedups in PCT. The 0.6.2 release in May launched the character sets design document (PDD28), updated the PHP implementation to use PCT, and added support for OpenGL. The 0.6.3 release in June marked substantial progress on the Ruby implementation (Cardinal), added a number of operators to Rakudo, added support for complex return types to PCT, and enhanced Parrot's operation on the g++ and gcc compilers.

The 0.6.4 release in July launched improvements to the Smalltalk implementation (ChitChat), migrated the PHP implementation (Pipp) over to PCT, added complex number operations and generic types to Rakudo, added support for smoke testing of Parrot using Smolder, enabled the latest version of Test::Harness, and made core improvements to bytecode disassembly, lexical subroutine handling, and register allocation. The 0.7.0 release in August included the implementation of the concurrency milestone (PDD25) and substantial improvements to the exception system, by Allison and chromatic. It also included the launch of the design document specifying Parrot's intermediate representation (PIR).

There have been several notable developments on the Parrot front. In June, the Parrot project incorporated the Parrot Foundation, a US non-profit foundation dedicated to supporting the development of Parrot and its language implementations. The management of the NLNet grant has been handed over to the new foundation.

The Perl Foundation recently announced a major new grant to support Perl 6 development. Ian Hague has donated $200,000, Roughly half of this money will go directly to supporting Perl 6 development (including Perl 6 on Parrot), with the rest going to help increase the organizational capacity of the Perl Foundation.

Jonathan Worthington has received funding from several community sources which will allow him to work full time several days a month. Jonathan has been putting this time to use both on the core Parrot VM subsystems and Rakudo. Some of the areas he's worked on include IO and Junctions in Rakudo, as well as the object subsystem in Parrot and Rakudo. He has been documenting his ongoing progress in his "use Perl journal":http://use.perl.org/~JonathanWorthington/journal/.

Patrick has also been doing funded work on Rakudo, including major work on the OO subsystem and compiler tools. He has also been helping put together a plan for future Rakudo work, as well as documenting how Parrot and Rakudo are designed, which is very important in enabling new contributors to the project. Most recently, he worked on implementing list and pair parsing for Rakudo, which are core pieces of the Perl 6 language. Patrick has also been documenting his progress in his "use Perl journal":http://use.perl.org/~pmichaud/journal/.

Parrot participated in Google's Summer of Code program. Andrew Whitworth implemented a new tri-color, incremental garbage collector. This work marks a major step towards making Parrot production-ready. Andrew documented his work in his "use Perl journal":http://use.perl.org/~Whiteknight/journal/. Andrew's work will be integrated into the core of Parrot as part of the upcoming garbage collection milestone in the NLNet grant. Kevin Tew, a long-time Parrot 