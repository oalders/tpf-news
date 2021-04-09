
---
title: "2009Q1 Grant Proposal: Web.pm - a lightweight web framework for Perl 6"
author: Alberto SimÃµes
type: post
date: 2009-02-06 00:00:00 +0000 UTC
url: "/post/2009q1_grant_proposal_webpm_-"
categories:
 - Grants
 - Sign in

---

<dl>
<dt><strong>Names:</strong>

<dd>
<p>Ilya Belikin, Carl Mäsak, Stephen Weeks</p>
</dd>
</li>
<dt><strong>Email:</strong>

<dd>
<p>Ilya Belikin ([hidden email]), Carl Mäsak ([hidden email]),
Stephen Weeks ([hidden email])</p>
</dd>
</li>
<dt><strong>Amount Requested:</strong>

<dd>
<p>$3000.</p>
</dd>
<dl>
<dt><strong>Ilya: 1320$, for 12h/week.</strong>

<dt><strong>Carl: 1100$, for 10h/week.</strong>

<dt><strong>Stephen: 550$, for 5h/week.</strong>

</dl>
</dl>
<p>
</p>
<h2>Synopsis</h2>
<p>Two years ago, Juerd proposed the Web.pm module for Perl 6. Meant to
functionally replace Perl 5's CGI.pm, it is a revamping of the
fundamentals in order to better serve today's web development needs.</p>
<p>Apart from implementing a working web development toolkit, we intend
to integrate templating as a desirable default, make sessions and
sticky form values (as described in Juerd's email) work, and organize
the framework according to the Model-View-Controller (MVC) pattern and
the REST architecture.</p>
<p>The structure of the Web framework will be modular, and a web
application writer can choose the parts that make sense for the
particular application developed. As the user's application matures,
more high-level modules can be swapped in to combat the increasing
complexity.</p>
<p>We intend to develop Web.pm in the open, seeking frequent input from
people on #perl6 and #parrot, and from the mailing lists
perl6-compiler, perl6-users and perl6-language.</p>


<h2>Benefits to the Perl Community</h2>
<p>Having a simple, well thought-out web development toolkit can be an
enormous benefit to Perl 6 and its growing developer community. There
is good opportunity at this time to coordinate an effort to create
such a toolkit. A web toolkit will increase Rakudo visibility, and
bring more people to try Perl 6 programming.</p>
<p>Applications like November show that it is possible to create a Perl 6
web application today, but that much convenience functionality has to
be reinvented in the web application itself. The fact that a bridge
between Rakudo and Perl 5/CPAN does not exist at this point is an
obstacle, but also an opportunity to sit down and think about the
parts that we want to improve or remake rather than just inherit.</p>
<p>As the past six months of November clearly show, development of
real-world applications is one of the most effective ways to generate
feedback to the Rakudo development team, and to help focus on bug
fixes and features needed for everyday tasks.</p>
<p>
</p>
<h2>Deliverables</h2>
<dl>
<dt><strong>HTTP::Request, HTTP::Response and Web.pm modules that run on the
latest release of Rakudo.</strong>

<dt><strong>A URI module based on <a href="http://www.ietf.org/rfc/rfc3986.txt" class="rfc">RFC 3986</a>, using Perl 6 grammars.</strong>

<dt><strong>Web::Tags module for (X)HTML tags generation.</strong>

<dt><strong>Dispatcher and Routines (with REST support) modules for better
controllers code organization.</strong>

<dt><strong>A simple Perl6-ish templater as a replacement for HTML::Template.</strong>

<dt><strong>A slightly more advanced XML-aware templating system, similar to
Python's Genshi.</strong>

<dt><strong>Three web applications that make use of the Web module. (On the
assumption that having three clients to an API gives enough clues
about the different needs clients might have.) The first web
application will most likely be November, the wiki, and the second
Maya, a blogging engine. The third will be a proof-of-concept pastebin
for Perl 6 code with color coding.</strong>

<dt><strong>A tutorial clearly showing the strengths of the Web module
framework, why/when it should be used, and how to get started using
it.</strong>

</dl>
<p>The modules URI, HTML::Template, and Dispatcher already exist in a
working alpha state within the November repository, and are being u