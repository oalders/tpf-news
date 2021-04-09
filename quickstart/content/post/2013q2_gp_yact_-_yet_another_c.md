
---
title: "2013Q2 GP: YACT - Yet Another Conference Tool"
author: Alberto SimÃµes
type: post
date: 2013-05-04 00:00:00 +0000 UTC
url: "/post/2013q2_gp_yact_-_yet_another_c"
categories:
 - Sign in

---

<dl>

<dt id="Name">Name:</dt>
<dd>

<p>Torsten Raudssus (Getty)</p>

</dd>
<dt id="Amount-Requested">Amount Requested:</dt>
<dd>

<p>$3000</p>

</dd>
</dl>

<h2 id="Synopsis">Synopsis</h2>

<p>The current Act software and their instances are an often discussed topic in the world of Perl. The migrating of those instances, and the move forward to more modern Perl solutions in the system are often discussed. Last year we were able to address and start a concept at the Quack and Hack Europe 2012, we called it YACT - Yet Another Conference Toolkit, which targets to be a &quot;in place&quot; replacement for the existing Act software. It accesses the same database, and (should) be usable with the same templates. It also implements the idea of using git (and so GitHub) to maintain the Act setup for a conference.</p>

<p>YACT will not replace the Act name, Act will stay as the name for the complete project/platform/service, YACT is just a specific implementation state.</p>

<p>Additionally YACT is supposed to be also a management for Perl Monger groups and so should animate people to startup a Perl Monger group &quot;instantly&quot; and get all the nice shiny fuzz to manage it. This also means that we need a more flexible theme based style for this. A better support for localization is also a very important point here.</p>

<p>The current state is a very good database setup already, but the web interface setup is not really started and no administrative tools are added yet. There is still lots of work todo. (Current state: <a>&quot;/github.com/perl-act/yact&quot; in  https:</a>)</p>


<h2 id="Benefits-to-the-Perl-Community">Benefits to the Perl Community</h2>

<p>The most important and most used Act instance is managed since many years, nearly alone by Maddingue (S&eacute;bastien Aperghis-Tramoni) and the French Perl Mongers, but is still suffering from the problem of integrating help easily. Alone the missing scalability disallows other system infrastructure to support the Act in a productive manner (which many many companies would I bet). We would open a new door to relieve those hard working Perl people from their stress level.</p>

<p>The new features will allow to make a more decent and more perfect conference experience. Especially with the experience of the French Perl Mongers in making events, we planned on several very important features which would give help for all event organizers in the future, like checklists or features for managing common processes inside the act itself (like setting up extra pages for extra situation like a special dinner with extra reserveration informations).</p>

<p>Also a very important point is making the experience for the user more interesting, like the idea of integrating Booking.com a bit more to make it very easy to manage the accommodation. I hope there that we will find some love from Booking.com to get all the informations we need for this :).</p>

<h2 id="Deliverables">Deliverables</h2>

<pre><code> * DBIx::Class schema of the Act DB

 * Dancer implementation of the web interface to the Act DB

 * Simple CLI application for administration of the Act instance

 * Testsuite for YACT::DB, YACT::Web and YACT::App

 * Making git the conference holding instead of SVN

 * Integrating template themes

 * Better integration of localization

 * &quot;Making a conference&quot; checklist feature

 * Support for custom hostnames and conference independent landing pages

 * Making a scalable DB setup (read only at least)</code></pre>

<h2 id="Project-Details">Project Details</h2>

<p>The implementation itself is pretty straight forward. The existing urls and endpoints will be translated over to Dancer step by step. On those transfer the required functions are added to the result classes on the DBIx::Class schema. The same goes for the development of the application. This way we have a clean and straight forward model, we also can bind to XMPP, IRC, and all other protocols to allow more interesting accessing features (IRC bot in your conference channel?).</p>

<p>The administrative tools will be made with Moo and as primitive as possible. In general we avoid to use bigger systems (so far we have no Moose in play, so why bring it in). The template engine will be staying to Template::Toolkit, but we might implement a switch for this to other templates (very optional).</p>

<p>For scalability I would consider some database experts to find a very sane simple way to get read only sharing of the database to outposts. With some simple tricks we will then drive write access to the specific central instance. A serious spreaded backup could be added here, too.</p>

<p>The templates will be heading towards bootstrap, as this is one of the most accepted concepts on this and offers by itself already themes (which will not replace that we offer totally different &quot;layout&quot; themes for this).</p>

<p>For the localization I will be using my own Locale::Simple, which is not much more than a wrapper around gettext (Oh please god, someone make a grant for fixing the xs implementation). This allows us even to make out of code from the DuckDuckGo Community Platform, where this is open source, to make an own translation platform (another grant?).</p>

<p>The optional migration process should be &quot;perfect&quot; in the way to try to link the accounts and so allow migration of all the conference informations of a person.</p>

<h2 id="Inch-stones">Inch-stones</h2>

<p>The project is really a bit bigger, then one grant might cover. Those are the Inch-Stones I see, but depending on how much work those steps cost, additional grants for those sub tasks could be made (and also given out to other people)</p>

<pre><code> * Finish up the existing YACT source code to deliver all the base features
   required for most acts (as reference I would take a German Perl
   Workshop).

 * Integrating modern template themes. For this I already have
aknowledgement
   of several organizations (shadow.cat) and private persons (Marek Suppa,
   Valentin Schmid) who want to supply results here.

 * Making administration tools for setting up acts and managing them most
easy

 * Providing a cross-conference, permanent user profile people can point to
   from their other profiles (MetaCPAN, LinkedIn, etc). Probably with
nickname
   in the URL as well, instead of meaningless number, something like:
   http://act.yapc.eu/profile/getty

 * Adding scalability features to allow it to run &quot;readable&quot; from several
   points in the world, a central write point would be still fine enough for
   the current situation (I hope YAPC::Asia agrees).</code></pre>

<p>Definitly outside this grant are those points, even tho I think there could happening by the community automatically, if the points before are reached.</p>

<pre><code> * Migrating over all features not so common in use (API, Payment, ...)

 * Migrating all non-French Perl Mongers Acts into the French Perl Mongers
Act

 * Adding new modern features (Twitter, Facebook, Flickr, Hotels, ...)</code></pre>

<h2 id="Project-Schedule">Project Schedule</h2>

<p>I want to start as soon as possible with the project, as my personal schedule allows me to take the time. I target with help of the community to finish the first Inch-stone in 3 weeks, seeing the already existing state. The additional assets like templates and the involvement into the existing administration of the main act instances might bend the process to a longer time, indepedent of the code evolvement.</p>

<p>So far the community and the French Perl Mongers are very supportive, which will help accomplish the target soon. S&eacute;bastien Aperghis-Tramoni already has a clear process plan for the migration process.</p>

<h2 id="Completeness-Criteria">Completeness Criteria</h2>

<p>There is a conference and a Perl Monger group running YACT which access the existing Act DB parallel to the other Act instance on the French Perl Monger hosting.</p>

<h2 id="Bio">Bio</h2>

<p>I&#39;m doing Perl since now around 4 years and worked over 15 years as system developer and technical leader for several agencies in PHP before that. With Perl I started to get a complete new experience of how you can integrate yourself into the community. I always loved this aspect of Perl most, and with my seat in the Marketing Committee I try to bring this experience also to the outer world. The Act system is the very important key here, because the Perl community delivers the best community driven conferences and we should concentrate on those key features of our community. I currently work mostly as contractor for DuckDuckGo which allowed me to visit many conferences in the last year, and additionally allowed me to organize 2 events on my own, the 2 Quack and Hacks, one in US, one in France together with the French Perl Mongers.</p>


