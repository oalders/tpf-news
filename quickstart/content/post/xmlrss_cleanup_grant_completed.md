
---
title: "XML::RSS Cleanup Grant Completed - Final Report"
author: Rosellyne Thompson
type: post
date: 2007-05-02 00:00:00 +0000 UTC
url: "/post/xmlrss_cleanup_grant_completed"
categories:
 - Grants

---

<p>I am pleased to announce that Shlomi Fish has completed his <a href="http://news.perl-foundation.org/2007/01/grants_awards.html">XML::RSS cleanup grant</a>. In his own words, Shlomi has summarised the work he's done and offers his thanks to those who helped him transform XML::RSS into a high quality tool for the community:</p>

<blockquote><p>This is a summary of a <a href="http://www.perlfoundation.org/">Perl Foundation</a> XML::RSS grant, I've been doing. Its scope was for two or three months from 1 January, 2007 onwards, with a lot of work done by me before the grant application was submitted.</p>

<p>I've kept a journal on my work by writing entries in <a href="http://use.perl.org/~Shlomif%20Fish/journal/">my use.perl.org 
weblog</a>. Here are the links to them.</p>
<p>Before receiving the grant:</p>
<ol>
<li>
<a href="http://use.perl.org/~Shlomi+Fish/journal/31348">XML::RSS</a>
</li>
<li>
<a href="http://use.perl.org/~Shlomi+Fish/journal/31411">XML::RSS Meta Bug</a>
</li>
<li>
<a href="http://use.perl.org/~Shlomi+Fish/journal/31607">XML::RSS Update</a>
</li>
</ol>
<p>After receiving the grant:</p>
<ol>
<li>
<a href="http://use.perl.org/~Shlomi+Fish/journal/32081">XML::RSS - Grant Application Was Accepted + Update</a>
</li>
<li>
<a href="http://use.perl.org/~Shlomi+Fish/journal/32121">Yet Another XML-RSS 
Update</a>
</li>
<li>
<a href="http://use.perl.org/~Shlomi+Fish/journal/32214">XML::RSS Update: Full 
Test Coverage</a>
</li>
<li><a href="http://use.perl.org/~Shlomi+Fish/journal/32305">XML::RSS Update: Refactoring</a></li>
</ol>
<p>Now for some general commentary. <a href="http://www.askbjoernhansen.com/">Ask Bjørn Hansen</a> has been doing
extraordinary work taking care of the perl.org infrastructure, and also became the maintainer of <a href="http://search.cpan.org/dist/XML-RSS/">the XML::RSS module</a>. He has been very helpful in commenting on my patches and applying them, guiding me through the process, and providing feedback. </p>

<p>Our process in working on XML::RSS was the following:</p>

<ol>
<li>
<p>At first we did some work on reducing the number of open bugs to a minimum.
This either involved correcting the code, and closing duplicate bugs or ones that were too obscure.</p>
</li>
<li>
<p>After that and during development, we worked on expanding and extending the 
XML::RSS test suite. <a href="http://search.cpan.org/dist/Devel-Cover/">Devel-Cover</a> by Paul Johnson was a huge help here, in locating things we did not cover yet.
</p>

<p>We emphasised making sure the tests were meaningful, and reflected the functionality of the module.</p>

<p>
During the testing stage, some hard-to-find bugs were fixed.
</p>
</li>
<li>
<p>
After we had a 100% test coverage, we started with mercilessly refactoring the code. Many methods were extracted and encapsulated, an XML element generation code was added and eventually, each RSS version backend became its own class. Finally, the parsing was made cleaner and more robust, while overcoming some of XML::Parser's inherent shortcomings.
</p>
<p>For now, it seems the code is squeaky clean.</p>
</li>
</ol>
<p>
All these modifications were incorporated into version 1.29_02 of XML::RSS, which is a development version. Expect them in the stable XML::RSS version 1.30 soon.
</p>

<p>I'd like to thank Ask for his co-operation and support; <a href="http://search.cpan.org/~pjcj/">Paul Johnson</a> for writing the Devel-Cover module; the TPF people (<a href="http://use.perl.org/~Ovid/">Curtis "Ovid" Poe</a> and others) for accepting my grant, and for guiding me through the grant process; and all the people who reported bugs or supplied feedback.</p>

<p>Peace, Love, Camels and web feed mangling!</p></blockquote>
<p>It's been a pleasure working with Shlomi and I'd like to take this opportunity to wish him every success in his future projects!</p>



#grants #perl #shlomi-fis