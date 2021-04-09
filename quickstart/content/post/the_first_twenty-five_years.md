
---
title: "The First Twenty-Five Years"
author: mdk
type: post
date: 2012-12-18 00:00:00 +0000 UTC
url: "/post/the_first_twenty-five_years"
categories:
 - Perl 5 Development
 - Perl 6 Development
 - Perl Foundation
 - Sign in

---

<big><big>Introduction</big></big>

<big>Opening Remarks</big>
Being asked to write a piece celebrating the twenty-fifth anniversary, in essence a Silver Celebration, since the first release of the Perl Programming language was both a joy and a terror. Where would I start, what would I include, what approach should I take? It is a significant prospect as the sheer depth of history can only be matched by the breadth of influence that Perl has had over the previous quarter of a century. Then there was the inevitable consequences of having some information opinion based, common misconceptions or just not documented.[1]

My reaction was to think of all the elements that I knew made up the Perl Language and Community and to hold them aloft and groan as unlike Atlas I am unable to hold up the heavens.[2] There was an awfully large number of diverse elements that made up the phenomena known as Perl. So what then to do? 

This, therefore, is a personal journey with some elements gained from common research, I have decided to do a retrospective of the community and history, there are many things you will find familiar, hopefully some that you will find new and one or two that are established fact and good to be reminded off. 

I am indebted to all the others who did a large portion of the leg work from the general nature of being members of the organisations, development, community and projects; to the specific where I have leaned heavily on the knowledge of luminaries such as Dave Cross whose excellent presentation at the London Perl Workshop forms the Year by Year Repose. I am also deeply indebted to the information contained on the many Perl sites and to the elements on Wikipedia, please see the references for more information.
All mistakes, conjecture, opinions and assorted paraphernalia are my own, where I have used material from others that may be in question it is a failure in my judgement for its inclusion and not the fault of the original author.

I would ask you to note that each section that has footnotes has then printed at the base of that section before I move on and therefore the numbering correctly re-starts each time. The Sections are: Introduction; The Community; The Events; The Code; The Written Word; Towards A Future; A Year by Year Repose; Sources.

I would also like to thank Karen Pauley for being an excellent editor and accurately questioning all the silly things I have taken out, anything silly that remains is still my responsibility.

<big>In the Beginning...</big>
The beginnings are rather simple, and maybe a little mundane; Larry Wall (Tim Toady) released version 1.0 to the newsgroup comp.sources.misc on the 18th December 1987 while working as a programmer at Unisys. Perl was intended, we believe, to be a Unix scripting language to make report processing easier borrowing from sh, Awk and Sed.
Perl 2 was released in 1988 and added more features and a better regular expression engine, this was followed by Perl 3 in 1989, Perl 4 in 1991 to coincide with the Camel Book, until finally we have Perl 5 in 1994. Perl 5 was a major shift in the release of version numbers which I will touch upon below. Perl 6 started its life cycle in 2000 with a different principle to other versions of Perl. It was a complete re-write of the language and would start as a language specification before a release leading to the now apocryphal 'released in time for Christmas' line.

The development of Perl 5 continued both before and alongside that of Perl 6,[3] though it is interesting to note that there have been several major changes which could have been considered major version releases in previous versions of Perl. It is common these days for Perl 5 to be considered the language name with its point release being the version, hence the current version of Perl 5 is 16. Perl 5.6 was released in 2000; 5.8 was released and under constant development during 2002-08; version .10 released in 2007 for the 20th anniversary of Perl; version .12 was released in 2010, at this time Perl moved to a monthly release cycle with a yearly stable release, hence .14 was released in 2011 and .16 in 2012.

Perl 6 has influenced a number of changes, or backwards implemented as some would have it, in Perl 5 as well as having a number of implementations of its language specification such as Pugs (by Audrey Tang in Haskell), Niecza and Rakudo*. It has been a long standing source of humour that Perl 6 would be released in time for Christmas, but we didn't say -which- Christmas, the truth is that there are implementations of Perl 6 and the language itself is a specification that is under development. There are a number of projects, modules and libraries on CPAN that use Perl 6 and the production deployments continue to grow.

This is just the start of the journey as we now take a look at some of the many parts of the quarter century of Perl progress...

**Notes**
<small>[1] Then there is the inevitable variance in story by conjecture that allows us to ponder what is really truth and what is merely just recalled as being the truth.

[2] If I recall correctly he only held up the skies not the whole planet unlike that lying statue suggests.

[3] We shall have to note that there were the wilderness years when two communities arose that seemed not only divergent in focus, but also disagreeable in nature. These wilderness years coincided with the rise of languages such as PhP, Python and Ruby (though more specifically the popularity of the Rails framework).

These arguments were the inevitable consequence of the enthusiasm for each language by the active members of those communities. Today we can proudly say that we are one Perl community with two Perl languages, Perl 5 and its variants and Perl 6 with its. From the two, it is apparent that Perl 6 has pushed the envelope in regards to language design and specifications leading to implementations such as Pugs and Rakudo while also inspiring changes in Perl 5 such as Moose. </small>

<img alt="25 anniversary.png" src="http://news.perlfoundation.org/25%20anniversary.png" width="200" class="mt-image-center" style="text-align: center; display: block; margin: 0px auto 0px;" />

<big><big>The Community</big></big>

It is a widely held belief[1] that Perl has one of the best, if not the best, communities in the Open Source world. It is certainly true that it is a long established community and one that can be very rewarding to be a member.

From the very early days of the language there was a strong move towards 'grassroots' or natural community development. Before Perl had reached a point release between 1987-88 we had seen the first JAPH (Just Another Perl Hacker) snippet run wild across the newsgroups and that was likely the start of many programmers need to be identified as having Perl skills.

In those early days the majority of communication was conducted on newsgroups, in particular the now famous comp.lang.perl where Perl pioneers such as Randal Schwartz, Tom Christiansen, Len Wiseman, Steve Simmons and Marion Hakanson (see: https://groups.google.com/forum/?hl=en&fromgroups#!forum/comp.lang.perl) along with Larry Wall, met and discussed Perl issues. 

The first O'Reilly Perl conference was not until 1997 in San Jose, California, a decade since the language was first released and three years after Perl 5 was birthed to the world. At that conference many of the people who had been communicating simply with lines of text met face to face and realised that wasn't such a bad idea. Following this conference the very first Perl Monger group, the New York Perl Mongers was formed by brian d foy and other members of the New York Perlers which was quickly followed by a group in Boston.

<big>Perl Mongers</big>
The O'Reilly Perl Conference would eventually mutate into the OSCON conference which still runs in Portland each July. The Perl Mongers grew to one hundred groups in its first year as recorded at the second O'Reilly Perl Conference and now has more than one hundred and fifty groups.

<img alt="pm.gif" src="http://news.perlfoundation.org/pm.gif" width="150" class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
The original idea was that Perl Monger group names would be based on the name of a geographical location, hence we have London.pm, Boston.pm etc. A pun was used for the .pm as this is an extension for a Perl file and also an acronym of Perl Mongers. Today a number of other groups are also part of the Perl Monger family, there is Dahut.pm named for a mythical beast and refers to a non-geographical group of organisers/members; there is Les Mongueurs du Perl (The Mongers of Perl) which is taken by a collection of French groups who organise events such as the French Perl Workshop, Perl QA Hackathon and OSDC (Open Source Developers Conference); in the community we also have drinkers.pm a group dedicated to the consumption of sweet liquor and LGBT.pm (who might also be known as Equality.pm) which was started by @techpractical to reflect the broad nature of our diverse community.

Many of the Perl Mongers groups organise technical presentations and social meetings. These meetings act as a good way to pass information on latest developments in Perl or on the broader technical world and jobs markets, we also use the community resource jobs.perl.org for the same reason.

The Perl Community have always organised themselves into running a variety of local events, from regular Perl meetings, hackathons, training sessions, workshops and even the yearly YAPCs (Yet Another Perl Conference). Conferences are usually hosted by a local group who bids for the privilege, some groups have run more than one event and many have run more than one in a single year.

<big>White Camel</big>
Such is the importance of community that Perl has its own awards that are handed to community members who have made a significant contribution to the state of Perl and the community. This award was initiated by the Perl Mongers and O'Reilly & Associates at the Perl Conference in 1999 and has been awarded yearly since then, the scheme is managed by brian d foy.

<big>Perl Monks</big>
The Perl Mongers have always been more of a social endeavour focussing on the meeting in the real world of like minded individuals to share and express their thoughts and converse[2], but since late 1999 there has been a location for the technically curious.
<img alt="perlmonks.jpg" src="http://news.perlfoundation.org/perlmonks.jpg" width="200" class="mt-image-none" class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
If you have a question about Perl and wish to progress along a path of almost devout following then you can enter the Monastery of Perl and join the Perl Monks.[3] The Perl Monks website, FAQ and posts are intended to be a user-focussed resource of information, in many ways it is almost a CPAN for documentation. The site predates many of the more social systems that have become popular in recent years but contains a vast store of collected knowledge and experience to help new initiates upon their path to understanding.

<big>The Organisations</big>
There are a number of organisations in the Perl Community[4] which seems confusing to those who are uninitiated with the vast complexity of a global community with diverse aims and needs. There are legal matters of dealing with different countries, the focus of the organisation, such is the vast range of different elements to the Perl world one organisation would have quite a split in attention.

The largest of these organisations are The Perl Foundation, YAPC Europe Foundation, Japanese Perl Association and The Enlightened Perl Organisation. These groups can be seen as complementary, not contradictory, groups with overlapping areas of focus and differing duties.

**The Perl Foundation**
In 1999 Kevin Lenzo formed the organisation Yet Another Society (YAS) to help to organise the Perl conferences in North America. This organisation would eventually become known as the Perl Foundation. The original premise of the Perl Foundation was to help run the grassroots events in the Perl Community for North America, and to help the conference organisers with banking and organisation.

<img alt="tpf_logo_transparent.png" src="http://news.perlfoundation.org/tpf_logo_transparent.png" width="300" class="mt-image-center" style="text-align: center; display: block; margin: 10px auto 40px;" />
Over the last decade the focus of the Perl Foundation has shifted and grown so that along with helping to fund, organise and support North American Perl conferences it has evolved to handle grants and bounties for extending, amending and evolving both Perl 5 and Perl 6 throughout the world.

The Perl Foundation has also secured the rights for the Perl logo, artistic licence 2.0 and trademarks, the Onion and the name Perl when used in conjunction are a Trademark in many countries. It has helped the JPA overturn a trademark case in Japan which would have had potential issues for Perl.

Today, under the steady guidance of the current President, Karen Pauley, the Perl Foundation has become deeply involved with the worldwide Perl community. The Perl Foundation has extended its mission to improve and evolve Perl and to promote and support community. There are committees to handle marketing, grants, funding and the recent Perl Advocacy committee to work in the community.

The Perl Foundation is run entirely by volunteers who give up their time simply to help support the greater community, they are also proud to embrace diversity and the evolution of the community, the Perl Foundation survives solely from the generous sponsorship, donations and fund raising.

**YAPC Europe Foundation**
<img alt="yef.png" src="http://news.perlfoundation.org/yef.png" width="200" class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
YEF was formed in 2003 with the mission to help promote and support the growing YAPC::Europe. YEF was a desire to offer the same support and services to the European Conference scene that TPF was founded for in the United States.
YEF provide a framework to help organise many of the yearly conferences in Europe as well as supporting the other European events such as hackathons. YEF survives on donations and like TPF is staffed entirely by volunteers.

**Japanese Perl Association**
<img alt="jpa_logo-small.jpg" src="http://news.perlfoundation.org/jpa_logo-small.jpg" width="150" height="75" class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
The JPA was founded to help organise the yearly YAPC::Asia which has been held for many years in Tokyo due to the concentration of organisers and the need for a centrally accessible location. YAPC::Asia is currently the world's largest Perl conference with over 600 atendees.

In recent years the JPA has also fought, and won, an important battle for the control of the name Perl in Japan with the help of the Perl Foundation

**Enlightened Perl Organisation**
<img alt="logo-large.png" src="http://news.perlfoundation.org/logo-large.png" width="150" class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
The EPO was formed in March 2008 to celebrate and encourage the shift from what was perceived as a Renaissance in Perl to a more Modern Perl for the twenty-first century. Since the release of Perl 5.8 there was a shift in thinking for Perl 5 and a recognition that it was a language in its own right sharing a common root with Perl 6 but having a separate identity. The EPO was formed to further promote and support this evolving situation and community.

Unlike the three other organisations the EPO does not support any one specific conference or agenda but rather lends its support and business services to a range of them, from the Perl QA Hackathon, the London Perl Workshop, YAPC::Brazil and YAPC::Europe. 

The EPO also runs a number of initiatives and projects for the broader community such as the Send-a-Newbie initiative whose aim is to give financial aid in sending people new to a Perl conference to a major event; the support of CPAN Testers with the raising of funds and handling of costs; supporting grants for development and sponsoring diverse elements such as logo competitions. 

The EPO is a membership organisation and gains funds from membership, donation and sponsorship. It hopes to promote the tighter cohesion between business and community and to further enhance this union. It is staffed entirely by volunteers.

<big>The Logos</big>
A retrospective of the Perl World has to make some mention of the many logos of the Perl Community. Once again this is a rather eclectic list, there are many logos that were omitted such as logos for CPAN Testers, DBIx::Clagass, Catalyst, Dancer and Mojolicious in an effort to tell a slice of the story from the most visible elements.

**The Camel**
<img alt="perl_camel.jpg" src="http://news.perlfoundation.org/perl_camel.jpg" width="150" class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
The first logo that can be said to be used in conjunction with Perl is the Camel. This was the logo that O'Reilly decided to use for the Programming Perl book and became synonymous in the early days with the Perl Community. The Camel is owned by O'Reilly, but they allow the Perl Community to use it on its associated community sites and projects.

There have been a few variants on the Camel logo such as the one used to show the 'Programming Republic of Perl' a 90s construct, to its usage as a symbol of community excellence in the White Camel awards. 

Many community groups and individuals have their own Camel, Castaway and brian d foy both have a personal soft camel, as do the Italian Perl Mongers, London.pm (Amelia), EPO (Niles), Orlando Perl Mongers and others.

**The Onion**
<img alt="onion_logo.png" src="http://news.perlfoundation.org/onion_logo.png" width="120" class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
The logo of the Perl Foundation, and a trademark when used in conjunction with Perl, the Onion is now recognised as a semi-official symbol for the Perl community. The TPF did not seek to supplant the use of the Camel but merely give an official image that the community could call its own. Since the TPF is a community managed organisation the logo is therefore protected for community uses which is granted to a wide number of groups and initiatives.

**Camelia**
<img alt="800px-Camelia.svg.png" src="http://news.perlfoundation.org/800px-Camelia.svg.png" width="150" class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
In creating the logo for Perl 6 Larry Wall drew inspiration from the people he always intended to target. Larry would love to see a new generation of programmers in our world, he would love to see a new generation of inventors, his aim was to be friendly and encouraging and with her bright colours Camelia certainly addresses those concerns.

There is also for me an interesting parable in the choice of a butterfly for Perl 6 as the butterfly is an example of a morphic creature. The butterfly has passed through an intermediary, chrysalis, phase and changes from an earth bound crawling insect into a beautiful creature of flight. What better metaphor for the creation of a language from one base system into something else?

**The Raptor**
<img alt="raptor-and-number.png" src="http://news.perlfoundation.org/raptor-and-number.png" width="150" class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
I have written before, see note 5 below,  how we came to use the image of the Raptor, or more correctly the Velociraptor,[5] as a logo for Perl 5. The main use is to signify the two different communities and languages. As with Perl 6 where the use of Camelia is not enforced, the use of the Velociraptor is the choice of the individual. 

The current Raptor was created after a conversation between several people on the EPO irc channel, the two most relevant being myself and Sebastian Riedel, a popular image was created after that conversation by Sebastian and is available for use with a CC licence on Github.[6]

**CPAN and MetaCPAN**
<img alt="cpan.png" src="http://news.perlfoundation.org/cpan.png" width="150"  class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
For many years there has been a CPAN logo that reminded the visitor of libraries of knowledge contained in a single location. In the early days of MetaCPAN they used a variant of the same logo but in 2011 a competition, sponsored by the EPO, was held to find a new logo for MetaCPAN with the eventual winner being used for the whole of the project from this date.

Many of the logos used by projects, communities and libraries in the Perl community are created by the members of that group, they are identifying marks and in the case of some are used to great effect. A particular favourite and relative newcomer is the logo used by Cluj.pm for their meetings and events.

**Notes**
<small>[1] Certainly it is widely held in the Perl community and when I have personally spoken to people who have attended a variety of language conferences they applaud the general spirit of the Perl community.

[2] Sometimes over voluminous amounts of fermented vegetable products.

[3] There is a high level of insightful jest, punning and complementary mockery in the religious overtones used by the denizens of the Perl Monks, none of it is intended to be offensive.

[4] Leading me to make witticisms such as why have only one organisation when a dozen can do the same job.

[5] http://mdk.per.ly/2011/03/02/evolution-of-the-velociraptor/

[6] https://github.com/kraih/perl-raptor</small>

<img alt="25 anniversary.png" src="http://news.perlfoundation.org/25%20anniversary.png" width="200" class="mt-image-center" style="text-align: center; display: block; margin: 0px auto 0px;" />

<big><big>The Events</big></big>

<big>Yet Another Perl Conference (YAPC)</big>
In 1987, ten years since the release of Perl 1.0, a group of aspiring Perl programmers gathered together and held the very first event of the Perl world, the first Perl Conference (TPC). Within a single year of that event they had changed the whole nature of the Perl world in both the number of groups, contributors and social structures that were formed. 

These early conferences were held in conjunction with O'Reilly media who would, in 1999, take this event and transform it into the Open Source Conference (OSCON).
In the same year, 1999, the first ever Yet Another Perl Conference (YAPC) was held in the United States. The community had grown to hold its own large events. In 2000 Europe joined the fray and held its first YAPC. Since then we have held yearly ::EU and ::NA events; we were joined by an Asia conference which is currently the world's largest Perl event; a conference in South America which evolved into the ::Brasil YAPC and the yearly YAPC::Russia.

For 2013-14 there are plans to hold a conference in Australia,[1] ::Australia/::Oz(?), and even plans to attempt a YAPC/Workshop in India and perhaps even Africa.

<big>Workshops</big>
A Perl Workshop often indicates a smaller event to the YAPC and is held on only one or two days, it may also be country or language specific. The first ever was held in Germany in 1999. Since then a number of countries, including the United Kingdom who hold the largest Perl workshop with the LPW attracting close to 250 guests each year for one day, hold Perl workshops.

There have been a number of Perl workshops, some return annually to the same location, others move around a region and are hosted by different user groups. To date we have had workshops such as the:[2]

* German Perl Workshop
* Nordic Perl Workshop
* London Perl Workshop
* French Perl Workshop
* Frozen Perl
* Orlando Perl Workshop
* Israel Perl Workshop
* Korean Perl Workshop
* Baltimore Perl Workshop
* Italian Perl Workshop
* Kiev Perl Workshop
* Russian Perl Workshop
* Portuguese Perl Workshop
* Pitsburgh Perl Workshop

The coming year should reflect the growing number of Perl Workshops with the possibility of events being held in new regions.

<big>Other Events</big>
Perl has been a proud member and participant in a number of other events that are both specific to our community and a part of the larger community of Open Source. With local and international events such as: the UKUUG/F/LOSS Spring and Autumn Conferences in the United Kingdom, the Dynamic Languages Conference in the United Kingdom, the Open Source Developers Conference in France, OSCON in America and FOSDEM in Belgium.

Perl has also, since 2008, held the Quality Assurance Hackathon in a different country each year to gather many of the luminaries of the Perl QA world who gather and create new standards and work on projects relating to QA in Perl.

In 2012 we also saw the first ever Quack and Hack Hackathon sponsored by DuckDuckGo, the Moving to Moose Hackathon, and the Perl Reunification Summit held by Djikmat.

Perl has also participated in both the Google Summer of Code and Google Code-In that are held in the spring and autumn respectively.

**Notes**
<small>[1] There was a Perl conference as a part of the OSDC in Australia, but to my knowledge no stand alone event similar to a YAPC.

[2] A comprehensive list was hard to compile due to there being more than one system of hosting events and some being singular occurrences. I apologise for any omissions and ask that they be noted in the comments for future inclusion and reference.</small>

<img alt="25 anniversary.png" src="http://news.perlfoundation.org/25%20anniversary.png" width="200" class="mt-image-center" style="text-align: center; display: block; margin: 0px auto 0px;" />

<big><big>The Code</big></big>

With a period of twenty-five years, thousands of contributors and an eco-system that encourages participation and sharing it is inevitable that there would be a number of projects, libraries, modules and scripts that have collected and sometimes evolved.[1] Before the advent of Perl 5 the resources for collecting these scripts were few and far between, and one or two have fallen into legend and are now taken out by wizened Perlers around a flickering light where the Tales of Terror are shared and Matt's Script Archive[2] comes into its own magnificent glory. In these Enlightened days it is easy to mock those early pioneers and to smile fondly at some of the erroneous efforts, but they were the only resource of their time and they were formative in the evolution.

<big>CPAN</big>
With the release of Perl 5 in 1994 a new boundary was made between the code of the past, Perl 4, and that of the future. This was complemented by the creation of Perl's largest resource the mighty Comprehensive Perl Archive Network. CPAN is the largest archive of its type and has been emulated in other scripting languages to varying degrees of success. It is a resource where code authors can request an account, Pause Account, and then upload or contribute to code that is in the archive. CPAN predates GitHub by many years and was the model for how one joined a project by being awarded a commit-bit, or co-maint of a module or script.

The evolution of CPAN has brought us to the modern day where we have MetaCPAN, a web service and API for accessing the archive that is CPAN and for working within it. We also have CPANPLUS and CPANM which are abridged version of the complete archive. Many companies complete their own DarkPAN, a personal version for use in-house and there is also the vast BackPAN which contains scripts and libraries that may not be compatible.

The great advantage of the CPAN was that we were suddenly able to run distributed projects across multiple contributors, timezones and countries. No longer would we need a single archive of scripts maintained by the author of the original project, we could share, contribute, manage, and even surrender our work to other people. This allowed many projects to grow, in fact it facilitated a thriving community of contributors. CPAN is also a tool for installing modules and libraries and managing your own Perl with, if required, custom dependencies and versions.

CPAN has spawned an entire section of itself called the Acme namespace, an area where the code that is left there may be a prototype, a test, an early attempt to solve an issue, a throwaway hack that others might use or just a bit of fun. To view the whole of CPAN visually you can use the Map of CPAN[3] which allows you to pan and zoom around the archive as if it were a geographical location, it has also spawned a video showing how the map has changed and evolved over time.[4]

<big>Perl 5 Porters</big>
With the release of Perl 5 we also saw the creation of a new sub-group of programmers in the Perl community, the P5P or Perl 5 Porters. This group was originally responsible for porting from Perl 4 to Perl 5 and since then they have managed the releases, code quality, evolution and management of the Perl 5 Core. From amongst their ranks they elect a PumpKing, a single person responsible for directing the group and development, in essence a chair of the board.

The current Perl 5 PumpKing is Ricardo Signes who has guided Perl 5 through a significant period where we come to terms with the evolving nature of a written release schedule and forward plan for the evolution of the core. The previous PumpKing was Jesse Vincent who managed the core through its period of change from general releases when the code was 'ready' to a schedule of releases, development, deprecation; these two have changed the manner in which both the development of the Perl Core and the manner of releases are managed, streamlining the process, making it efficient and providing us with a recognised plan for continuing development and forward progress while understanding the need for careful management of backwards compatibility.

<big>Strawberry Perl</big>
The first release of Perl for the Mac appeared in 1992, and with the advent of OSX it has been a part of every Mac release,[5] but we had to wait until 2007 for a CPAN available version of Perl for Windows. Strawberry Perl was a response to the need for a OS Licenced version of Windows Perl.[6]

The project has brought the ethos of CPAN, Perl and F/LOSS to a traditionally closed architecture and continues to be evolved with each new version of Perl. The maintainers do their best to make the latest version of Perl available on Windows at almost the same time that it is released to Unix based systems.

It should be noted that these days Strawberry Perl usually comes bundled with Padre an IDE that is developed in Perl and suitable for using in a multiple of environments. Padre has a strong following amongst many developers who like to use a powerful IDE of this type, and is a good default for those developing in a Windows environment.

<big>Perl 6</big>
In the year 2000 there was a movement to create a new version of Perl. Instead of it being an evolution of existing code the plan was to formulate a list of what was required in the language and what was not required, from that a language specification would be created and this would be the template for implementations of that specification. This was the Perl 6 project.

Since then we have seen twelve years of continuous research, development, implementation and release of code and practices from the Perl 6 team. The work undergone on this project so far has inspired a number of releases, including the currently popular Rakudo* project.

There is no full specification-complete version of Perl 6 released to the world, and it has been a source of some jest that we have waited so long. This is however underselling the important work already carried out and the fact that many of the practices, methods and code has been back-ported to the Perl 5 core. There are many modules and libraries written in Perl 6 and some authors have ported their Perl 5 work to the language.

<big>Moose</big>
One of the largest projects, and changes to Perl 5 inspired by the work on Perl 6 has been the creation of the Moose postmodern object system. Moose borrows heavily from Perl 6 along with CLOS (Lisp), Smalltalk, Java, Beta, OCAML, Ruby etc. And uses them in a thoroughly Perl 5 manner.

Moose extends the Perl 5 object system in order to make object orientated programming easier, allowing programmers to think more about the task they need to perform and less about the mechanics of how to achieve that task. Moose has now been incorporated into a number of projects and its use in Perl 5 re-defines to many the maturity of Perl 5 Object Orientated programming.

<big>The Well-Known Projects</big>
In the early days of web frameworks, that used MVC (Model-View-Controller), Perl had a number of projects. Some popular libraries were Mason (an MVC framework) and Class DBI for work with the databases, there was also Template Toolkit for designing re-usable view elements. These were often used in conjunction with custom, client specific, cgi frameworks to create a vast number of business models and systems.

These systems, and they are still used to this day, led to a range of younger projects as a new generation of programmers appeared on the scene. The first widely popular framework was Catalyst, now considered the largest of the Perl frameworks, its strength lies in the vast number of modules and elements already written for it allowing almost every eventuality to be considered. Catalyst has been followed by the popular frameworks of Dancer and Mojolicious.

Dancer seems ideal for completing and controlling multiple sites and systems and has gained a following with people wishing to move away from cgi while leveraging the best elements of CPAN.

The original author of Catalyst created a new framework, Mojolicious, which is designed to be dependency free. Mojolicious is a out-of-the-box solution to web development, rather than using CPAN it comes almost complete with many of the essential elements in the core of the system. Mojolicious seems designed to allow the rollout of a project and a complete solution for systems. 

Dancer and Catalyst attempt to supply a higher level of granular management, situation agnostic, configurable environment by leveraging a vast choice of dependencies customised to target specific solutions.

<big>But What About...</big>
So we arrive at the contention, at this point Perl people will be tearing their hair out and screaming abuse in my general direction,[7] as I come to conclude this section without mentioning their favourite Perl tool. Which is the issue, Perl has so many good tools and it has a culture that not only celebrates sharing and contributing but in learning. 

One of the major habits in the Perl community is learning from other languages, we are inspired by and cautious of the way other languages work and implement. This is most evident when Perl gurus like Miyagawa look at WSGI and create PSGI or Matt S. Trout who looked at the way other languages did remote server management over ssh and decided that Perl could have Tak or his playing with the parser that created Devel::Declare. Perl is not re-inventing the wheel, more it looks at what works and what doesn't work and decides to implement it in a Perlish manner and hopefully a more code complete way.

There were a number of projects I could have spoken about here, I could have given some time to how Plack has revolutionised how our code is deployed to servers, or spoken about dependency management using Carton. I could have discussed www::mechanize and its importance in Perl being used as language-of-choice for web page automation, I never mentioned Web::Simple, Rose::DB, Web::Graphite, POE, Distzilla, the .Net, www or xml namespaces, or a hundred other influential modules that have shaped our world.

The strength of Perl is the huge ecosystem that has evolved with it, we have seen efforts to control and understand that ecology with the creation of CPAN, MetaCPAN etc. The challenge for the coming years is the aggregation and understanding of this vast resource, MetaCPAN makes a start on how we can access the resource in a modern fashion, but we are only just beginning to see how we can present this information.

**Notes**
<small>[1] But which of these would I have to focus on in a twenty-five year retrospective? Whichever I include may seem contentious and those that I exclude just as much an issue. Tracing the formative projects, code and changes would be a grand task and I would be the first to encourage it to be done, but not within the confines of this retrospective.[8]

[2] For accuracy I should note that I played a little with history for the essence of this paragraph, the Year by Year Repose will confirm that Matt's Script Archive appeared in the same year as CPAN.

[3] Map of CPAN: http://mapofcpan.org/#/

[4] Video of the Map of CPAN: http://vimeo.com/51893508

[5] Though there are some issues to do with which version is used.

[6] To my knowledge before Strawberry Perl gave us a functioning and complete Windows installable version of Perl one had to use Cygwin or Active State to have Perl on Windows.

[7] perhaps they might even burn an effigy of me.

[8] It's not as if people wont be coming after me for missing out Perl Golf or Perl Poetry.</small>

<img alt="25 anniversary.png" src="http://news.perlfoundation.org/25%20anniversary.png" width="200" class="mt-image-center" style="text-align: center; display: block; margin: 0px auto 0px;" />

<big><big>Code Completeness</big></big>

I have mentioned that the Perl community and ecosystem varies from other languages in some ways, or more likely evolved earlier due to its maturity. One of the major differences is in the testing culture, this is not to suggest that other languages do not have such a culture, but that they do not have a culture like the one in Perl.

The testing culture of Perl has led to a vast range of modules and systems that help programmers to develop competent code that is feature rich but also complete in its quality. To begin we created the Test Anything Protocol (TAP) which by its name suggests that testing is prevalent, many projects are not prototyped in code but defined in tests, we also use methods such as test-driven deployment.

<big>CPAN Testers</big>
<img alt="Screen Shot 2012-08-27 at 12.15.37.png" src="http://news.perlfoundation.org/Screen%20Shot%202012-08-27%20at%2012.15.37.png" width="120" class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
In 1998 Graham Barr and Chris Nandor conceived the system that is CPAN Testers. This important group takes every module uploaded to CPAN and runs it against every operating system and version of Perl that it will be deployed against, this generates a report of any issues which is automatically reported to the author of the module. This system is repeated to ensure that any changes or evolutions in code are accounted for. This is a high degree of quality control that produces hundreds of thousands of reports each month, there are currently over 26 million reports.

CPAN Testers is maintained like most projects in the Perl world by volunteers and donations from the companies who use its vital statistics.

CPAN testers is complemented by Perl Brew which is a project that allows you to run multiple versions of Perl on the same install of an operating system allowing a variable code environment for developers.

<big>Code Veracity</big>
There is a suspicion that has been inherited from the early days of Perl that the community is populated by random hackers who spew out illegible scripts to solve single issues. This is heightened by people referring to Perl as duct tape or the glue of the internet and has led through some poor examples and low quality programming to a reputation that the language does not deserve.

Much of this suspicion comes from the powerful TIMTOWTDI (There Is More That One Way To Do It, Tim Toady) which allows a wider variety of answers to the same question. In the community we also use variants of BCINABT (Because Consistency Is Not A Bad Thing, bicarbonate) to suggest that although one can do almost anything doesn't mean one should and that there are generally agreed ways of doing something.

To this end we have elements such as 'use strict' and 'use warnings' which can be placed at the beginning of a Perl file to prevent random hackery. We also have many modules that help the neophyte, to the grand master, program.  Modules such as the mighty Devel::Cover, and TAP::Harness, which allow you to verify your code, along with Perl Critic, Perl Tidy which will aid you in building compliant code. Much of this work has been inspired by, and conforms to the sublime Perl Best Practices by Damian Conway and to the general need to keep consistency of quality while allowing freedom of implementation.

Veracity is also a by-word for the standards ensured by the Perl 5 Porters. For a number of years David Mitchell has been working on fixing existing bugs in the Perl 5 Core, after a very successful year of fundraising by the Perl Foundation in 2011 on the Perl 5 Core Maintenance Fund he was joined by Nicholas Clark. Between them these two, along with the many volunteers who work without specific grants, have been improving the standard and quality of the code in the Perl 5 Core. 

<big>Code Quality</big>
Verifying code quality; a mature attitude to testing while planning, developing and deploying systems; a suite of tools to help with ensuring best practices and internal project standards; these elements have gifted the Perl community with a respect for code quality. This led to the start of a yearly conference that focuses on Quality Assurance. The QA Hackathon (Perl) was started in Oslo in 2008 and has been held in a different country for each year since then. The focus of the event is to gather some of the finest developers in the Perl community to one location to discuss, develop and plan improvements to code quality over the coming year. In 2013 it will once again return to the United Kingdom.

Perl has always relied on code quality and veracity, it is why we have a fondness for ensuring good backwards compatibility. Because of the many years of development and deployment of sometimes critical systems there is a culture of ensuring that the latest code is as compliant with existing code as is possible. This is not always possible, but the work that is done is by turn surprising and amazing.

<img alt="25 anniversary.png" src="http://news.perlfoundation.org/25%20anniversary.png" width="200" class="mt-image-center" style="text-align: center; display: block; margin: 0px auto 0px;" />

<big><big>The Written Word</big></big>

The very first documentation for Perl was back in 1987/88 with the creation of the original man perl for use on Unix systems. Like any software project the quality of this document has improved and varied over the many years of Perl evolution. From an early age Perl attracted books teaching you how to get the best from the language. The very first book, Programming Perl, the now legendary Camel book was printed in 1991 (further editions in 1996, 2000 and 2012), it was followed by books such as Learning Perl in 1993 (with editions in 1997, 2001, 2005, 2008 and 2011), Beginning Perl (first version in 2000, 2004 and 2010), the less well known but respected Perl Cookbook in 1998 (second edition in 2003), the seminal Higher Order Perl in 2005 and in that year we also were graced with Damian Conway's Perl Best Practices, the Modern Perl in 2011 (with a second edition in 2012) and this year a new Beginning Perl (2012) by Curtis 'ovid' Poe.

There are many other books that have become well known in the Perl world. One of particular note for me and those developing web frameworks, and in particular using Catalyst as a solution, was the Definitive Guide to Catalyst, Kieren Diemet & Matt S. Trout (with others). this book not only demonstrated how to use Catalyst but gave a good grounding in how to develop, deploy and manage applications for modern usage.
However the wealth of Perl material is not just contained within published books, the projects themselves often have their own manuals, FAQs, Wikis, Advent Calendars and guides. Famously Perl runs several advent calendars each year and has the websites such as Learn Perl,[1] the Perl FAQs,[2] and Perl Doc.[3]

Over the many years of evolution we have also seen publications such as the Perl Journal arise, this was joined by the German publication $foo Magazin, these are complemented by initiatives such as Perl News, the idiosyncratic Perl Weekly, blogs.perl.org, and the Ironman Competition and aggregated blog site.

With many Perl programmers answering questions on sites such as Slashdot, on Facebook groups such as Perl Programers -  MetaCPAN - CPAN; on LinkedIn in Perl Mongers - Perl and Perl Foundation, twitter, irc on Perl's own IRC server irc.perl.org - as well as freenode #perl, and in the long standing gateway to Perl knowledge the Perl Monks the supplementary knowledge to slot between references and articles is massive. To complete this there are many user groups such as the Perl Mongers who maintain mailing lists and meetings to further share the written information.

<small>[1] http://learn.perl.org/

[2] http://learn.perl.org/faq/

[3] http://perldoc.perl.org/</small>

<img alt="25 anniversary.png" src="http://news.perlfoundation.org/25%20anniversary.png" width="200" class="mt-image-center" style="text-align: center; display: block; margin: 0px auto 0px;" />

<big>Towards A Future</big>

Had anyone asked Larry Wall on December 18th 1987, 'hey, Larry, what do you see in the next twenty-five years for Perl?' I doubt his reply would have been accurate, my thoughts would be along the lines of 'hell, I wasn't planning on using it much after I leave this job, why would I be using it in a new century?'. Now we could ask Larry what he actually would have said, but like the witty backronyms for Perl itself we probably couldn't trust him to be serious with us.

So what does the future hold for Perl? Well I don't have a crystal ball but I cannot see the language fading from usage in the next quarter century, the truth of the matter is that even though there are languages that can do some of the things that Perl does, some of them do some things better, others do things Perl wasn't designed for, there is no language that has been designed to do the things that Perl is very good at doing. No language in the current scripting languages seems to have the flexibility, maturity and extensibility of Perl. The main power of Perl has always been its ability to quickly adapt, and be adapted, to suit purposes. It has been referred to as glue, as tape, as a Swiss Army Knife with a chainsaw attachment, as a command line processing language. The truth is that it is all of these and more.

The power of Perl is that it does not limit your expectations, it doesn't control your involvement, to some this is a scary prospect, to those who understand that there are protocols and methods for controlling and throttling that power Perl becomes a set of tools and the toolbox that holds them.

Perl will continue to evolve as a number of implementations of the same language, I would imagine we have many years of Perl 5 evolution still to come and at the same time we will see an implementation of the Perl 6 specification. But I would also hazard that the Perl 6 developers will add and evolve their specification and move onwards as this is not a community that is used to standing still.

The greatest challenges we will face for Perl is a shifting end-user base that will become more reliant on devices that are feature focussed but crippled in application choice, the rise in mobile devices will continue and Perl will need to evolve to work with that. A better challenge for us to face would be the integration with electronically aware, and connected devices and systems, the apocryphal internet of things, in this Perl could be a powerful tool. I also believe that the more we see a divergence of language uses in the other scripting languages the more they will face issues in their core designs, issues that Perl avoids due to its malleable nature, what some believe is the crippling factor for Perl is likely to be its saving grace as it has the power and flexibility to cope with the shifting goalposts of an increasingly technologically reliant world.

**The Changing Community**
<img alt="Screen Shot 2012-12-20 at 11.16.18.png" src="http://news.perlfoundation.org/Screen%20Shot%202012-12-20%20at%2011.16.18.png" width="200" class="mt-image-right" style="float: right; margin: 0 0 20px 20px;" />
For the community itself we face a number of challenges, statistically we seem to be generally an ageing body who are predominantly white, male, and often bearded.[1] Without having a full statistical package of a number of varying communities over a period of years I cannot say if we are a part of or separate to a trend that may be prevalent in the industry itself. Certainly in my experience there have been slight, and I do mean slight, changes in the participation at the conferences I have managed. 
There is a balance between attracting a wider mix of people into our community, making it an attractive, safe and profitable place for them to grow and forcing a change that may not be supportable. These are questions that we will seek to address in the future as we attempt to evolve our community along with our language, to be more welcoming, accepting and tolerant of people, it should be natural to embrace diversity.
To this end the Perl Foundation have taken the bold step of creating a Advocacy Committee to do work in the Perl Community led by the widely respected Yaakov Sloman. In the coming years he will be working hard to grow and evolve our community to match the maturity and strength of our language.

**A Chance to Meet and Greet Each Other**
An element that has grown each year that I have been a member of the Perl community is the number of organised events by region. These are the hackathons, workshops and technical meetings. In the past year we have seen a rise of people wishing to hold conferences in countries where there has never been one. YAPC Europe will next year be held in Kiev, and as previously mentioned a number of new workshops and conferences will be held in Asia, and Australia.

Also growing is the number of commits to CPAN and the number of programmers with a Pause Identity, as the statistics on CPAN have attested, and after a few years of slow growth Perl is once again rising on many of the indexes that track jobs, development and usage.

**To the next 25 years...**
So as we celebrate the start of the 26th Year of Perl with its 25th Anniversary I hope you will be cheered by the wealth of the Perl World. The future for Perl seems to be full of challenges, surprises, evolution and, hopefully even more growth and diversity. See you at the 50th Birthday.[1]

**Notes**
<small>[1] Only one thing is certain, in twenty-five years time they are going to have to pay me to write another of these.</small>

<img alt="25 anniversary.png" src="http://news.perlfoundation.org/25%20anniversary.png" width="200" class="mt-image-center" style="text-align: center; display: block; margin: 0px auto 0px;" />

<big><big>A Year by Year Repose</big></big>

It is with grateful thanks to Dave Cross that I print here a list of the changes and developments he has recorded for the last twenty-five years of Perl. Like my own list it is not comprehensive, but it does contain many elements that were significant in the history of our language.

<big>1987/8</big>
Perl 1.0 (and 2.0) Released!
**1987/8 Releases**
Dec-18 1.000
Jan-30 1.001..10
Feb-02 1.011..14
Jun-05 2.000
Jun-21 2.001
**1987/8 Community**
First JAPH
**1987/8 Technical**
Perl
**1987/8 Books**
man perl

<big>1989</big>
Perl 3.0 Released
**1989 Releases**
Oct-18 3.000
Nov-11 3.002..4
Nov-18 3.005
Dec-22 3.006..8
**1989 Community**
comp.lang.perl
**1989 Technical**
Binary data
**1989 Books**

<big>1990</big>
Black Perl
**1990 Releases**
Mar 3.009..18
Aug 3.019..28
Oct 3.029..37
Nov 3.040..41
**1990 Community**
Perl poetry
**1990 Technical**
1990 Books

<big>1991</big>
Pink Camel published
**1991 Releases**
Jan 3.042..044
Mar 4.000
Apr 4.001..3
Jun 4.004..10
Nov 4.011..19
**1991 Community**
Artistic Licence
**1991 Technical**
Perl 4
**1991 Books**
Programming Perl

<big>1992</big>
MacPerl
**1992 Releases**
Jun 4.020..35
**1992 Community**
**1992 Technical**
MacPerl
**1992 Books**

<big>1993</big>
Pink Llama published
**1993 Releases**
Feb-05 4.036
**1993 Community**
Perl Packrats
**1993 Technical**
cgi-lib.pl
**1993 Books**
Learning Perl

<big>1994</big>
Perl 5 Released
**1994 Releases**
Oct-17 5.000
**1994 Community**
Perl 5 Porters
**1994 Technical**
Perl 5
DBI
CGI
**1994 Books**

<big>1995</big>
CPAN launched
**1995 Releases**
Mar-13 5.001
**1995 Community**
perl.com
perl.org
comp.lang.perl.misc
**1995 Technical**
CPAN
Schwartzian Transform
Matt's Script Archive
LWP
DBD::mysql
RSA encryption/decryption
**1995 Books**

<big>1996</big>
The Perl Journal
**1996 Releases**
Feb-29 5.002
Mar-25 5.002_01
Jun-25 5.003
Jul-31 5.003_01
Aug 5.003_02..03
Sep 5.003_04..05
Oct 5.003_06..07
Nov 5.003_08..10
Dec 5.003_11..17
**1996 Community**
The Perl Journal
**1996 Technical**
mod_perl
CPAN shell
Test::Harness
Zeroth Annual Obfuscated Perl Contest
**1996 Books**
Perl 5 Desktop Reference
CGI Programming on the World Wide Web
Programming Perl (2ed)

<big>1997</big>
The Perl Conference
**1997 Releases**
Jan 5.003_19..23
Feb 5.003_24..28
Mar 5.003_90..95
Apr 5.003_96..98
May 5.003_99..99a
May-15 5.004
Jun-13 5.004_01
Aug-07 5.004_02
Sep-05 5.004_03
Oct-15 5.004_04
**1997 Community**
The Perl Conference
New York Perl Mongers
**1997 Technical**
purl
**1997 Books**
Learning Perl (2ed)
Advanced Perl Programming
Learning Perl on Win32 Systems
CGI Cookbook
Effective Perl Programming
Perl Resource Kit

<big>1998</big>
London Perl Mongers
**1998 Releases**
Jul-22 5.005
Jul-27 5.005_01
Aug-08 5.005_02
**1998 Community**
The Perl Institute
London Perl Mongers
**1998 Technical**
CPAN Testers
OnePerl
HTML::Mason
POE
XML::Parser
**1998 Books**
MacPerl: Power and Ease
The Perl Cookbook (1ed)
Perl 5 Pocket Reference (2ed)
Perl in a Nutshell
Perl for Dummies (2ed)

<big>1999</big>
First YAPC
**1999 Releases**
Mar-28 5.005_03
Apr-29 5.004_05
**1999 Community**
TPC becomes OSCON
First YAPC::NA
German Perl Workshop
Perlmonth
White Camel Awards
Yet Another Society
Rhizomatic (irc.perl.org)
**1999 Technical**
Template Toolkit
Class::DBI
Perl Power Tools
Perl Golf
**1999 Books**
Learning Perl/Tk
Programming Web Graphics with Perl & Gnu Software
Writing Apache Modules with Perl and C
Object Oriented Perl
Mastering Algorithms with Perl
Perl CD Bookshelf
Elements of Programming with Perl
Teach Yourself Perl in 24 Hours
Perl and CGI for the World Wide Wed (1ed)

<big>2000</big>
Coffee Cups
**2000 Releases**
Mar-22 5.6.0
**2000 Community**
First YAPC::EU
Perl Monks
First Perl Whirl
Perl Advent Calendar
**2000 Technical**
Perl 6 (announced)
Testing Revolution
ActivePerl
Symbol::Approx::Sub
**2000 Books**
Programming the Perl DBI
Perl 5 Pocket Reference (3ed)
Perl for System Administration (1ed)
CGI Programming with Perl (2ed)
CGI Programming 101 (1ed)
Beginning Perl (1ed)
Programming Perl (3ed)
Debugging Perl
Perl for Dummies (3ed)
The mod_perl Pocket Reference

<big>2001</big>
Damian Conway works for us
**2001 Releases**
Apr-08 5.6.1
**2001 Community**
use.perl.org launched
Perl Monks and Perl Mongers join YAS
YAS forms The Perl Foundation
YAPC Auctions
Parrot (April Fool)
**2001 Technical**
Acme::*
Parrot
Test::Builder
Movable Type
**2001 Books**
Network Programming with Perl
Data Munging with Perl
Perl: The Complete Reference
Perl CD Bookshelf Version 2.0
Professional Perl Programming
Instant Perl Modules
Writing CGI Applications with Perl
Perl Debugged
Professional Perl Development
Instant CGI/Perl
Perl and CGI for the World Wide Wed (2ed)
Learning Perl (3ed)
Perl Developer's Dictionary
Perl for the Web
MySQL and Perl for the Web
Perl for Web Site Management
Beginning Perl for Bioinformatics
Core Perl

<big>2002</big>
nms
**2002 Releases**
Jul-18 5.8.0
Dec-18 1.0.15
**2002 Community**
nms
The Perl Review
**2002 Technical**
CPANPLUS
WWW::Mechanize
**2002 Books**
mod_perl Developer's Cookbook
Modern Perl Programming
Web Development with Apache and Perl
Perl and XML
Perl in a Nutshell (2ed)
Graphics Programming with Perl
Teach Yourself Perl in 21 Days (2ed)
Perl Pocket Reference (4ed)
Perl & LWP
Perl Black Book
Perl for C Programmers
Extending and Embedding Perl
Writing Perl Modules for CPAN
Perl for Oracle DBAs
Teach Yourself CGI in 24 Hours (2ed)
Embedding Perl in HTML with Mason
XML and Perl
The Perl CD Bookshelf Version 3.0
Best of the Perl Journal: Computer Science and Perl Programming
Programming Web Services with Perl
Perl Graphics Programming

<big>2003</big>
Perl On New Internal Engine
**2003 Releases**
Sep-25 5.8.1
Nov-04 5.8.2
Nov-15 5.6.2
Dec-18 1.0.16
**2003 Community**
PONIE
A Conference Toolkit
**2003 Technical**
Email::Simple
DateTime
**2003 Books**
Best of the Perl Journal: Web Graphics & Perl/Tk
Perl for Dummies (4ed)
Practical mod_perl
Mastering Perl for Bioinformatics
Learning Perl Objects, References and Modules
Perl 6 Essentials
Perl Cookbook (2ed)
Perl in Easy Steps
Perl Template Toolkit

<big>2004</big>
Catalyst
**2004 Releases**
Jan-14 5.8.3
Feb-23 5.005_04
Apr-21 5.8.4
Jul-19 5.8.5
Nov-27 5.8.6
**2004 Community**
Catalyst community
**2004 Technical**
Maypole
Catalyst
**2004 Books**
Perl Debugger Pocket Reference
CGI Programming 101 (2ed)
Perl Medic
The Perl CD Bookshelf Version 4.0
Perl 6 and Parrot Essentials (2ed)
Best of the Perl Journal: Games, Diversions & Perl Culture
Beginning Perl (2ed)
Randal Schwartz's Perls of Wisdom

<big>2005</big>
DBIx::Class
**2005 Releases**
May-30 5.8.7
**2005 Community**
Google Summer of Code
Planet Perl
Perl Mongers Census
**2005 Technical**
DBIx::Class
Pugs
**2005 Books**
Pro Perl
Higher Order Perl
Teach Yourself Perl in 24 Hours (3ed)
Advanced Perl Programming (2ed)
Perl Best Practices
Perl Testing: A Developer's Notebook
Learning Perl (4ed)

<big>2006</big>
Moose
**2006 Releases**
Jan-31 5.8.8
**2006 Community
2006 Technical**
Moose
Plagger
**2006 Books**
Beginning Perl Web Development
Perl Hacks
Minimal Perl

<big>2007</big>
Strawberry Perl
**2007 Releases**
Dec-18 5.10.0
**2007 Community**
Perl Buzz
$foo Magazin
**2007 Technical**
Strawberry Perl
**2007 Books**
Mastering Perl
mod_perl 2 Users' Guide
Perl by Example (4ed)
Catalyst

<big>2008</big>
Marketing
**2008 Releases**
Dec-14 5.8.9
**2008 Community**
QA Hackathon
Copenhagen Marketing BOF
Enlightened Perl Organisation
**2008 Technical**
Test Anything Protocol
Dist::Zilla
Mojolicious
**2008 Books**
Learning Perl (5ed)
Professional Perl and MySQL

<big>2009</big>
PSGI
**2009 Releases**
Aug-22 5.10.1
**2009 Community**
Send-a-Newbie
Iron Man
blogs.perl.org
New perl.org
**2009 Technical**
PSGI/Plack
Dancer
**2009 Books**
Automating System Administration with Perl
Definitive Guide to Catalyst

<big>2010</big>
Rakudo *
**2010 Releases**
Apr-12 5.12.0
May-16 5.12.1
Sep-06 5.12.2
**2010 Community**
Annual release cycle
**2010 Technical**
Rakudo *
App::cpanminus
metacpan
**2010 Books**
Catalyst 5.8: The Perl MVC Framework
Effective Perl Programming (2ed)
Beginning Perl (3ed)

<big>2011</big>
Money!
**2011 Releases**
Jan-21 5.12.3
May-14 5.14.0
Jun-16 5.14.1
Jun-20 5.12.4
Sep-26 5.14.2
**2011 Community**
Perl Weekly
Perl News
New perl.com
TPF Perl Core Maintenance Fund
**2011 Technical**
Carton
**2011 Books**
Learning Perl (6ed)
Perl 5 Pocket Reference (5ed)
Modern Perl (2ed)

<big>2012</big>
New Camel
**2012 Releases**
May-20 5.16.0
Aug-08 5.16.1
Oct-12 5.14.3
Nov-01 5.16.2
Nov-10 5.12.5
**2012 Community**
Community Summit
Perl Advocacy Committee
**2012 Technical**
???? (we'll tell you at the end of the year ;) ).
**2012 Books**
Intermediate Perl (2ed)
Programming Perl (4ed)
Beginning Perl

<img alt="25 anniversary.png" src="http://news.perlfoundation.org/25%20anniversary.png" width="200" class="mt-image-center" style="text-align: center; display: block; margin: 0px auto 0px;" />

<big><big>Sources</big></big>

The following is the short list of sources, and resources, I used to fact check and create the article, like the piece itself it is in no way comprehensive and I will accept any corrections, revisions or additions and take all responsibility for inaccuracies or omissions.

Dave Cross, London Perl Workshop presentation, ©2012

The Perl Foundation, http://www.perlfoundation.org
YAPC Europe Foundation, http://www.yapceurope.org/ , ©2003-12
Enlightened Perl Organisation: http://www.enlightenedperl.org, ©2008-2012 
Japanese Perl Association: http://japan.perlassociation.org/

Strawberry Perl: http://strawberryperl.com
Moose: http://moose.perl.com
Rakudo: http://rakudo.org/
Comp.lang.perl: https://groups.google.com/forum/?hl=en&fromgroups#!forum/comp.lang.perl
MetaCPAN: https://metacpan.org
Tap::Harness: https://metacpan.org/module/TAP::Harness
www::mechanize: https://metacpan.org/module/WWW::Mechanize
Devel::Cover: https://metacpan.org/module/Devel::Cover
Devel::Declare: https://metacpan.org/module/Devel::Declare#AUTHORS
Map of CPAN: http://mapofcpan.org/#/
Video of the Map of CPAN: http://vimeo.com/51893508
CPAN Testers: http://cpantesters.org/
$Foo Magazin: http://www.perl-magazin.de/
Ironman: http://ironman.enlightenedperl.org
Send-A-Newbie: http://sendanewbie.enlightenedperl.org
**Perl: **
Perl: http://www.perl.org/
Perl Monks: http://www.perlmonks.org/
Perl Mongers: http://www.pm.org/
Learn Perl: http://learn.perl.org/
Perl FAQ: http://learn.perl.org/faq/
PerlDoc: http://perldoc.perl.org/

**Wikipedia:**
Perl: http://en.wikipedia.org/wiki/Perl, ©2012 
Programming Republic of Perl Logo: http://en.wikipedia.org/wiki/File:Programming-republic-of-perl.png, ©2012 
Perl Camel: http://en.wikipedia.org/wiki/File:Perl-camel-small.png, ©2012 
Onion Logo: http://en.wikipedia.org/wiki/File:Onion_64x64.png, ©2012 
Perl: http://en.wikipedia.org/wiki/Perl, ©2012 
JAPH: http://en.wikipedia.org/wiki/Just_another_Perl_hacker, ©2012
Perl Mongers, http://en.wikipedia.org/wiki/Perl_Mongers, ©2012

Other:
Evolution of the Velociraptor: http://mdk.per.ly/2011/03/02/evolution-of-the-velociraptor/, ©Mark Keating 2011
Raptor on Github: https://github.com/kraih/perl-raptor, © Sebastian Riedel 2012
Article by kamaal: http://news.ycombinator.com/item?id=4869647


#perl #perl-5 #perl-6
