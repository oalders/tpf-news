
---
title: "April 1 - Act"
author: Elizabeth Mattijsen
type: post
date: 2015-04-01 00:00:00 +0000 UTC
url: "/post/april_1_-_act"
categories:
 - Grants
 - Sign in

---

Since the beginning of the year, some smaller steps had been made and things are looking good!

First of all the the database schema had been changed to provide some additional things on request of some organisers. Now I realise that I also need to make new versions of the Act-out-of-the-Box vagrant-image once in a while. The DBIx::Class has been updated accordingly and it all works like a charm

Secondly, I started playing around with the DBIx::Class implementation and natural stumbled upon the first hurdle... Conferences are not in the database, those are neatly stored in a bunch of INI files. Adding another table to the database which will be populated from the INI-files is giving Act-Voyager at least a top-level entry point for the relation databases, finally there are relations between the rows in the separate tables.

Those INI files provide much more information and only some is really about the conference itself, like dates and location. Other things that are in it are the number of rooms what names they have, the price tiers and tickets or other products. Information that will have to end up in other tables soon. But, this all sounds like the start of the admin-tool, which I had given a lower priority.

While playing with the DBIx::Class schema and being very pleased that I could build simple scripts showing me a list of talks or users related to some conference, I started to write the first bit of the REST api, as a starting point... and here a whole new story unfolds itself:

Do You Speak-a My Language?

For entire population groups on planet earth it looks like they do not have to concern about language difference, for them there exists only 1 language, 1 language to rule all... And for those people it is fine to organise a conference or workshop in their local language. And it's perfectly okay to let Act perform in a way as if your local language would be the only one to consider. However, as organiser of the Dutch Perl Workshop, we care about being open to people that do not come from the Netherlands and do not speak Dutch and therefore run our web-site bilingual and provide most of the texts in English as well. I even suggested to one of our people to set it up in German.

Act caters for all that, no problem - only a technical design problem here.

Retrieving data in multiple languages from a REST api is not uncommon, and the IETF even came up with a solution called Content-Negotiation. Basically, 1 single canonical URI for 1 resource, using the"Accept-Language" header-field to let client and server decide what content will be delivered. There is no need to put a language tag inside the URL to do so. But it makes you wonder if there is the same way to create update or delete resources in a specific language... so, without causing url-polution.

Nope, no answer is to be found!

Until lately, when I figured out how to do HTTP-Authoring in a multilingual REST api, simply by using the 'Content-Language' header-field in the request. It is possible within the restraints of the RFCs to update a resource with a representation in another language, or only delete the language variant for a resource instead of entirely whipping it of the database.

In Cluj (Romania) and lately at AmsterdamX I gave a demonstration, a proof of concept. The implementation is quite easy and simple. Moreover it does also separate the REST api objects from the internal data structure. So many API designers couple their tables directly to REST resources, which is a big mistake on itself. Such solutions make it impossible to change underlying structures without corrupting the REST api. Having said that... for Act it is essential to have them separate because once the new version is up and running on top of the current database there will come a migration to a new database.

On GitHub you can find the AnimalSanctuary that does show some implentations on the above concept. Sepa