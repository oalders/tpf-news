
---
title: "Perl 5.10 now available"
author: Andy Lester
type: post
date: 2007-12-19 00:00:00 +0000 UTC
url: "/post/perl_510_now_available"
categories:
 - Sign in

---

<p>
Today the Perl Foundation announces the release of Perl 5.10, the
first major upgrade to the wildly popular dynamic programming
language in over five years.  This latest version builds on the
successful 5.8.x series by adding powerful new language features
and improving the Perl interpreter itself.  The Perl development
team, called the Perl Porters, has taken features and inspiration
from the ambitious <a href="http://dev.perl.org/perl6/">Perl 6</a>
project, as well as from chiefly academic languages and blended
them with Perl's pragmatic view to practicality and usefulness.
</p>

<h2>Significant new language features</h2>

<p>
The most exciting change is the new <b>smart match operator</b>.
It implements a new kind of
comparison, the specifics of which are contextual based on the
inputs to the operator.  For example, to find if scalar <tt>$needle</tt> is in array <tt>@haystack</tt>,
simply use the new <tt>~~</tt> operator:
</p>
<pre>
  if ( $needle ~~ @haystack ) ...
</pre>
<p>
The result is that all comparisons now
just Do The Right Thing, a hallmark of Perl programming.
Building on the smart-match operator, Perl <b>finally gets a
<tt>switch</tt> statement</b>, and it goes far beyond the kind
of traditional <tt>switch</tt> statement found in languages like
C, C++ and Java.
</p>

<p>
Regular expressions are now far more powerful.  Programmers
can now use <b>named captures</b> in regular expressions, rather than counting parentheses for 
positional captures.  Perl 5.10 also supports recursive patterns,
making many useful constructs, especially in parsing, now possible.
Even with these new features, the regular expression engine has
been tweaked, tuned and sped up in many cases.
</p>

<p>
Other improvements include <b>state variables</b> that allow variables to
persist between calls to subroutines; user defined pragmata that
allow users to write modules to influence the way Perl behaves; a
<b>defined-or operator</b>; field hashes for inside-out objects and
<b>better error messages</b>.
</p>

<h2> Interpreter improvements </h2>

<p>
It's not just language changes.  The Perl interpreter itself is
<b>faster with a smaller memory footprint</b>, and has several UTF-8 and
threading improvements.  The Perl <b>installation is now
    relocatable</b>, a blessing for systems administrators and operating
system packagers.  The source code is more portable, and of course many
small bugs have been fixed along the way.  It all adds up to the best
Perl yet.
</p>

<p>
For a list of all changes in Perl 5.10, see Perl 5.10's <a href="http://search.cpan.org/dist/perl-5.10.0/pod/perl5100delta.pod">perldelta</a> document included
with the source distribution.  For a gentler introduction of just the high points, the slides for 
Ricardo Signes' <a href="http://www.slideshare.net/rjbs/perl-510-for-people-who-arent-totally-insane">Perl 5.10 For People Who Aren't Totally Insane</a> talk are well worth reading.
</p>

<p>
Don't think that the Perl Porters are resting on their laurels.
As Rafael Garcia-Suarez, the release manager for Perl 5.10, said:
"I would like to thank every one of the Perl Porters for their
efforts. I hope we'll all be proud of what Perl is becoming, and
ready to get back to the keyboard for 5.12."
</p>

<h2> Where to get Perl </h2>

<p>
Perl is a standard feature in almost every operating system today
except Windows.  Users who don't want to wait for their operating
system vendor to release a package can dig into Perl 5.10 by
downloading it from CPAN, the Comprehensive Perl Archive Network,
at <a href="http://search.cpan.org/dist/perl/">http://search.cpan.org/dist/perl/</a>,
or from the Perl home page at <a href="http://www.perl.org/">www.perl.org</a>.
</p>

<p>
Windows users can also take advantage of the power of Perl by
compiling a source distribution from CPAN, or downloading one of
two easily installed binary distributions.
<a href="http://strawberryp