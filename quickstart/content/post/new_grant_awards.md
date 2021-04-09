
---
title: "New Grant Awards"
author: Curtis "Ovid" Poe
type: post
date: 2007-04-01 00:00:00 +0000 UTC
url: "/post/new_grant_awards"
categories:
 - Grants

---

<p>The Perl Foundation is pleased to announce two new grant awards.  The first is adding new policies to Perl::Critic.  The second is improving the Smolder project.</p>

<p>Note that the second project involves the TAP::Parser module (which was known as TAPx::Parser at the time the grant application was submitted).  This was a project I started, currently maintained on the CPAN by Andy Armstrong (I'm still involved, as are others).  This is slated to be the replacement for Test::Harness.  Because we're seeing more grant applications involving this module, I have decided that I will abstain from all future votes for applications which specifically require adding TAP::Parser support (if it's peripheral to the project, that's OK).  Approving money for people integrating a project that I started doesn't pass my personal "smell test".  This does mean that it's possible we'll have worthwhile projects which come up one vote short, but in this case, Smolder was unanimously approved by everyone else.</p>

<h3>Policies for Perl::Critic</h3>

<p><strong>Name</strong></p>

<p>Chris Dolan</p>

<p><strong>Email</strong></p>

<p>cdolan@cpan.org</p>

<p><strong>Project Title</strong></p>

<p>Policies for Perl::Critic</p>

<p><strong>Synopsis</strong></p>

<p>I propose to implement a selection of new policies for Perl::Critic</p>

<p><strong>Benefits to the Perl Community</strong></p>

<p>More Perl::Critic policies will help developers to make their code
more consistent and maintainable by giving them more ways to comply
with a set of best practices.</p>

<p><strong>Deliverables</strong></p>

<p>I will implement the twenty PBP policy ideas in the TODO.pod of
Perl::Critic v1.01.</p>

<p><strong>Project Details</strong></p>

<p>Perl::Critic is helping Perl overcome it's reputation for hard-to-
maintain code by aiding developers enforce a consistent coding
style. As of v1.01, Perl::Critic has 93 pluggable policies, 71 of
which come from Damian Conway's Perl Best Practices book. Our TODO
list contains another twenty PBP policies thought to be
implementable. I propose to complete those twenty. Each of them
will be documented and tested to the level of quality expected for
the Perl::Critic project.</p>

<p><strong>The policy list:</strong></p>

<p>http://search.cpan.org/~thaljef/Perl-Critic-1.01/
TODO.pod#OTHER_PBP_POLICIES_THAT_SEEM_FEASIBLE_TO_IMPLEMENT

<p><strong>Project Schedule</strong></p>

<p>2 months, summer 2007

<p><strong>Bio</strong></p>

<p>I contribute to CPAN as CLOTHO and CDOLAN. I am one of the project
leaders for Perl::Critic. In 2006, I implemented 27 Perl::Critic and
Perl::Critic::More policies. I spearheaded the effort to get our
policy test coverage to almost 100%. I organized the Perl::Critic
volunteers for the 2006 Chicago Hackathon. The Perl::Critic
maintainer, Jeff Thalhammer, has read and approved this proposal.</p>

<p><strong>Amount Requested</strong></p>

<p>$2000 = 20 policies x 4 hours per policy x $25/hr

<hr/>

<h3>Improving Smolder</h3>

<p><strong>NAME</strong></p>

<p>Michael Peters</p>

<p><strong>EMAIL</strong></p>

<p>mpeters@plusthree.com

<p><strong>PROJECT TITLE</strong></p>

<p>Improving Smolder

<p><strong>SYNOPSIS</strong></p>

<p>Smolder (http://sourceforge.net/projects/smolder) is a web-based smoke
test aggregator. Meaning it allows developers and QA testers to upload
or monitor the test results from their projects. It works well, but has
several shortcomings that if addressed would allow for wider adoption
(and hence improvement) and to be more universally useable to the Perl
community as a companion to CPAN and CPAN Testers. This work would
address those issues.</p>

<p><strong>DELIVERABLES</strong></p>

<ol>
<li>Remove custom XML format in favor of using plain TAP and
TAPx::Parser.</li>

  <li>Extend Smolder to handle small CPAN style modules more easily and
automatically.</li>

  <li>Setup a Smolder server for the CGI::Application community to serve
as a testing ground and public display for their 110+ CPAN modules.</li>

  <li>Add per-project and per-developer RSS feeds as an alternative to
email notification.</li>
</ol>

<p><strong>PROJECT DETAILS</strong></p>

<ol>
  <li>Smolder currently runs tests expecting TAP as an output. However
since there were no standalone TAP parsers when Smolder was
originally written, "smolder_prove" (a derivative of "prove")
converted that TAP after the tests were run into XML (or YAML) so
that it could be sent to the Smolder server and parsed there. Now
with the ever increasing stability of TAPx::Parser this is no longer
an issue. This task would involve fixing Smolder to accept plain TAP
and use TAPx::Parser to parse it. We'll leave the XML/YAML support
as-is, but deprecate it in favor of using TAP.</li>

  <li>To make Smolder more useful to the Perl community it needs to not
only handle large, complex projects (which was the reason it was
built) but also handle smaller CPAN sized modules and applications.
To do this Smolder can not only act as a repository of test results,
but should also be able to run those tests itself in an automated
fashion. Support will be added to fetch (or update) a project from
an SCM (SVN initially, but will be flexible enough to add others
later). If changes have been made since the last time it checked,
the build/test cycle will be run and the results stored and the
normal notification cycle will occur.
<br/><br/>
To make this work with the current CPAN module build/test cycle
using ExtUtils::MakeMaker or Module::Build I will use, verify and
extend the work already started by Andy Armstrong, Ovid and the
Perl-QA community
(http://beta.nntp.perl.org/group/perl.qa/2007/01/msg7759.html) to
allow the "TAPx::" family of modules to be injected into that cycle
so that we can extract the full report.</li>

  <li>After Deliverable #2 is finished it will be possible for developers
(or communities) to setup a Smolder server to test their CPAN
modules automatically as soon as any changes are made. It will also
allow module users to submit better bug reports if the test suite is
failing on a particular OS or Perl version. Simply run the
build/test cycle (providing EU::MM or M::B with the correct
arguments so that TAPx::Parser is used) and upload the results to
the Smolder server. Much better than the typical
copy-paste-the-terminal-into-an-email method used today, and that
report can be referenced by URL in bug reports or other emails.
<br/><br/>
This Deliverable will focus on setting up such a Smolder server for
the CGI::Application community for as many modules as the community
and authors wish to add. I will find a donor for the hardware which
shouldn't be too difficult since Smolder doesn't use many resources.</li>

  <li>Currently Smolder sends email notifications if there are test
failures. Extend this by adding an RSS feed (using the newly
re-furbished XML::RSS). These feeds will be customizeable
per-developer as well as providing public feeds per-project.</li>
</ol>

<p><strong>PROJECT SCHEDULE</strong></p>

<ol>
  <li>Start off with the TAPx::Parser integration. - 10 hrs</li>

  <li>Extend Smolder to checkout project code via SVN - 5 hrs</li>

  <li>Extend Smolder to run checked out code and make sure was can get
full TAP reports from EU::MM and M::B - 10 hrs</li>

  <li>Add RSS feeds - 5 hrs</li>

  <li>Setup Smolder installation for CGI::Application community. - 5 hrs</li>
</ol>

<p><strong>BIO</strong></p>

<p>I am a Perl/mod_perl developer living in the Washington, DC suburbs and
working for Plus Three, LP (http://plusthree.com). Most of my work
involves building communication and fundraising tools used by political
and non-profit groups througout the US. One of my main responsibilities
at Plus Three is to build and maintain our testing tools and
infrastructure that we use on our core products.</p>

<p>I'm an active member of the CGI::Application, mod_perl, Krang
(http://krang.sf.net) and Perl-QA communities. I maintain several
modules on CPAN (http://search.cpan.org/~wonko/) and contribute to
several more. I am also the primary developer of Smolder and have even
spoken about it at the Pittsburgh Perl Workshop in 2006.</p>

<p><strong>AMOUNT REQUESTED</strong></p>

<p>$750</p>



#grant #grants #tap #ppi #perl::critic #smolder #michael-peters #chris-dolan
