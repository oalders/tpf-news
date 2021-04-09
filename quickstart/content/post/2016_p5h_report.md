
---
title: "Perl 5 Core Hackathon (p5hack) Report"
author: Todd Rinaldo
type: post
date: 2017-02-21 00:00:00 +0000 UTC
url: "/post/2016_p5h_report"
categories:
 - Conferences
 - Perl 5 Development

---

<img alt="perl5-hack-logo-04.jpg" src="http://news.perlfoundation.org/perl5-hack-logo-04.jpg" width="600" height="103" class="mt-image-none" style="" />

(Graphics credit to Leonardo Maia.)

Thanks to The Perl Foundation and our beloved sponsors, a large portion of the critical contributors to the Perl 5 Porters (p5p), the core group of Perl 5 developers, were able to get together on November 11th, 2016 at the Booking.com headquarters in Amsterdam, The Netherlands, for our very first Perl 5 Core hackathon, nicknamed **Perl 5 Hack**.

This is our report.

---

It was a four-day event, composed of discussions (centralized and specialized) on a myriad of important topics, knowledge sharing, and cooperative hacking. It provided us with the ability to cut down several months worth of discussions and loss of time due to timezone differences.

We worked on several topics, such as the `Configure` phase, core modules, security updates to Perl library loading order ("dot in `@INC`), improving compilation under C++11, Core API and internals, replacements for the buggy `utf8` IO layer, conditional lexical variable warnings, and deparsing some internal warnings information (`${^WARNING_BITS}`).

This alone would have been quite the fit, but we also achieved a few things along with those discussions. Some of our achievements include:

* Cleaning up dead hints code (pre-ELF)
* Improved our release guide
* Improved our Configure process
	* Removing hundreds of complicated C conditionals
	* Faster Configure script
* Cleaned up our install output
* Improved our API documentation
* Massive speed improvements to core modules
* Explored the new Test2 testing framework for Perl
* Reviewed and listed all current deprecations
* Tied all deprecations to an exact version, to improve transparency and communicate better our intentions regarding deprecated features and syntax
* Cleaned up core code:
	* Removed obscure and unused SVs
	* Removed dead code
	* Moved `Unicode::Normalize` to upstream blead
* Had done optimizations, such as when assigning to references.
* Reduced memory usage in core modules: `Data::Dumper` and `Time::Local`.
* Fuzzed core modules
* Have rewritten much of our readline implementation
* Introduced the indented `HEREDOC` feature, being released in the upcoming 5.26

We raised a lot of topics and we made numerous important and valuable changes to the core code and core utilities, but we have also held in-depth discussions on the following topics:

* Our COW implementation and where to take it in the future
* Efficiency of our hash bucket sizing
* Overloading indication
* Our policy on editing dual-life modules
* OP tree cloning
* Subroutine Signatures and their introspection
* The internals of the Perl to C compiler and what may be gleaned from it on improving perl
* The contribution hurdles when working with RT
* Vulnerability triage, classification, and handling - covering our security handling policies
* EBCDIC-portability
* Unicode (and specifically grapheme delimiters)
* Regular Expression internals
* Vtable-based hashes

Most of these have led to changes in policy, improvements or cleanups in core, and a "leg-up" on additional research into those topics.

All in all, despite it only being the first hackathon, it had already proved successful and beneficial. For this reason we intend to holding future hackathons, while both preparing better for them, as well as communicating better our work on them onward to the community.

Thank you to all attendees: Aaron Crane, Abigail, Brian Fraser, Dagfinn Mannsåker, Dave Mitchell, Dennis Kaarsemaker, Gonzalo Diethelm, H. Merijn Brand, John Lightsey, J. Nick Koston.  Karl Williamson, Leon Timmermans, Lukas Mai, Matthew Horsfall, Mattia Barbon, Nicholas Clark, Nicolas Rochelemagne, Sawyer X, Steffen Mueller, Stevan Little, Todd Rinaldo, Viekntiy Fusenov, and Yves Orton,

And than