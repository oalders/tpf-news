
---
title: "Act Voyager ---- long story"
author: Elizabeth Mattijsen
type: post
date: 2015-01-02 00:00:00 +0000 UTC
url: "/post/act_voyager_----_long_story"
categories:
 - Grants
 - Hackathons

---

Happy New Year from outer space...

since the last report a lot has happened on our journey with Act-Voyager... So, let me try to write down the episode of this saga...

In the months leading to the Act-hackathon there where two things being worked on...

Firstly, there is now a 'Act-out-of-the-Box' that makes it super easy to start hacking on Act (current version). With three simple commands, any developer can be up and running and start writing their own bits.

<pre>
    curl -Os https://raw.githubusercontent.com/Act-Voyager/Act-out-of-the-box/master/Vagrantfile
    vagrant up
    vagrant ssh
</pre>
And then, point your favourite browser to: http://localhost:8080/voyager to see a very bare Act-instance.

What it will need to work, is of course VirtualBox and Vagrant. It will download a Vagrant base image from my company server, which is preconfigured with Apache, mod_perl, PostgreSQL and the dependencies of Act. Thanks to many contributions of others it is workable and is something that needs updates as the Act-Voyager moves on... one of them, is including DBIx::Class as a core dependency in the Act-out-of-the-Box base image.

And secondly, I started working (and finished) the DBIx::Class Schema for Act. During the London Perl Workshop, I spend quite some time with 'SysPete' who gave a lot of advice on what to do with the DBIx::Class Schema for Act. There is soo much possible and nice to do with DBIx::Class, but for Act-Voyager it was advised to just stick to what was needed... a nice sugar-coated schema (using DBIx::Class::Candy) and documented attributes to explain what the purpose of each is. DBIx::Class::Schema::Loader does do POD in a excessive technical manner. Ribasushi had to convince me that it was better to remove verbose defaults, which enhance readability for experienced users.

The schema has already been proven to be usable in some branches of the Act Voyager project...

Which brings me to the star-map of the Act-Voyager project... "a Journey through the Universe" where we explore new galaxies and make small adjustments to the plans as we go... with one goal, a Next Generation Act, extensible and fun to hack on.

The original plan needed some adjustments... and it will need more whilst traveling along the stars. I am very fortunate to have a Grant-Manager that understands that development of the Next Generation Act is a project that needs flexibility. She does agree on the new plan below, which already had some adjustments since the Hackathon in Lyon.

<ol> 
<li>Orientation and fundamentals ( 1.000 USD ) FINISHED, needs continuos updates</li>
<ul>
<li>Make Act 'hackable; with "Act-out-of-the-Box"</li>
<li>DBIx::Class schema</li>
</ul>
<li>Port Core Act to DBIx::Class schema ( 1.000 USD )</li>
<ul>
<li>Make Act::Object sub-classes into proxies for DBIx::Class::Result classes</li>
<li>They have the same attributes and methods as the original ones</li>
<li>Retaining the tests that currently exists</li>
</ul>
<li>Dancer2 implementation ( 1.500 USD )</li>
<ul>
<li>move mod_perl handlers to Dancer2 routes</li>
<li>have Dancer2 routes fill in the original template</li>
</ul>
<li>RESTapi ( 1.000 USD )</li>
<ul>
<li>define all useful resources</li>
<li>write the POST, GET, PUT and DELETE methods (and more)</li>
<li>test, test and test</li>
<li>write documentation</li>
</ul>
<li>Theme Based Templates ( 1.500 USD )</li>
<ul>
<li>Find a proper frontend framework (Foundation)</li>
<li>Define Layouts</li>
<li>Define Pages</li>
<li>Redo the entire ACT with new templates and themes</li>
<li>Decide which SASS variable will be editable</li>
<li>Run several default Themes and Colorschemes</li>
<li>Write a admin/theme-selector</li>
</ul>
</ol>

The next two milestones, porting Core Act and building a Dancer2 implementation, are the steps that will ensure that all the old conferences (and the current) can be viewed and visited by