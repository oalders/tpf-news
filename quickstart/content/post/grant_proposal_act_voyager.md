
---
title: "Grant Proposal: Act Voyager"
author: Coke
type: post
date: 2017-09-20 00:00:00 +0000 UTC
url: "/post/grant_proposal_act_voyager"
categories:
 - Grants
 - Sign in

---

The Grants Committee has received the following grant proposal for the September/October round.
Before the Committee members vote, we would like to solicit feedback from the Perl community on the proposal.

Review the proposal below and please comment here by September 26th, 2017.
The Committee members will start the voting process following that and the conclusion
will be announced the first week of October.

# Act Voyager

- Name:

    Theo van Hoesel

- Amount Requested:

    USD 7,500

## Synopsis

A Modern rebuild of the Conference Tool Kit.



Act, as we know it, has been build more than a decade ago for Apache 1 and
mod-perl 1.3 architecture. And it has served the Perl Community very well over
all those years and hundreds of events.

But over the last year we have suffered from major outages. And since the
operating sytem it self being used to run on, has passed it LTS, more recent
compilers being packed, it has become extremely difficult (if not impossible) to
rebuild Act from scratch.

Also, Act as it is now, is no longer maintained in a pro-active manner, due to
it's archaic design - and people having other priorities.

Lastly, due to the flexibility off how it has been designed it comes with such a
complexity, that it is hard and difficult for current organisers to deliver new
websites for each and every conference or workshop they want to promote. Of
course, this also relates to the fact that most of the Perl developers are not
the best frontend designers.

Reasons enough to stop using Act, as a toolkit, it has become rusty and unfit
for the job. It has come to the point that Act should be to REST.

## Benefits to the Perl Community

The future is quite clear. Act requires a giant leap forward to keep up with the
demands of today users.

Visitors that come to the website should experience that Perl, although a
language with a longstanding history, is still keeping up and pushing the bar to
a higher level for competing alternatives. The user experience should be up to
modern expectations. Therefor Act Voyager will include a revamped frontend. And
hopefully will make it more likely to attract new people to the Community,
rather then putting people off on their first impression (as we know from Perl
Monks).

Organisers should be enabled to quickly and easaly set up a new event. That
should be as easy as setting a date, selecting a frontend theme and additional
styling. The admin part of Act would be decoupled from the public website. This
makes it much easier to implement, without messing with the public website
interface. This admin part would also need new tools that connect with Social
Media and allow simple things as sending emails. Previous organisers will
recognise all the hassle to get it up and running and most would wish it was
more easy to communicate with attendees and prospective Community members.

## Deliverables

### Act-out-of-the-Box

A 'must have' for working on anything related to Act, more specific, anything
related to the Legacy code base.

### Open API spec (aka Swagger)

This YAML specification of an API will describe the interface.

Rather than just writing code and deliver something that would look okay, it's
important to have a good definition of what the API looks like - much like Test
Driven Development.

### API server

Based on the above Specification, there will an API server that talks to the
Legacy Act databases.

### Act CLI

As a side product, mostly for testing purposes, there will be a simple command
line interface. Extremely useful for admins and people that dislike modern web.

### Theme based web-server

A Dancer2 application that will provide the HTML, JavaScript, and CSS needed to
access a Perl Event web page.

### Improve Modules

- Dancer2::Plugin::HTTP::Auth::Extensible
- Dancer2::Plugin::HTTP::Caching
- Dancer2::Plugin::HTTP::ConditionalRequest
- Dancer2::Plugin::HTTP::ContentNegotiation
- LPW::UserAgent::Caching
- HTTP::Caching

## Project Details

### Act-out-of-the-Box

Since it's a nightmare to build Act from scratch - DO NOT TRY THIS AT HOME - it
comes with Apache 1.3.x and mod-perl 1.31. It also includes a recent checkout
of the Git repository and it's dependancies.

This will benefit organisers to work on their website and test locally before
committing and merging with the 'Act-Conferences' repository.

It will also help developers that desire to fix bugs in the legacy codebase of
Act... or add new features.

This part has almost finished. I have gone through dozens and dozens off
iterations to get it working. Luckily I had a back up from a previous attempt
from some years ago. It only needs some nice script to make it easy for an
organiser to add a conference, linked to GitHub.

### Open API specification

There are many tools available around the Oen API specification that will help
backend developers to test their implementations. There are also tools available
to run a API simulation. And many more...

The Specification will be a living document and contributions are welcome. But
whatever the specification will be, that will be what client and server will
to agree on. This will open up the way developers can work on separate products.
With a proper Specification, it is possible to built different clients (like iOS
applications) and servers.

The specification can be just one big file and will soon become messy. This year
I started out writing the specification, split over several files. This will
make it much more easy to maintain. The current setup passes the OAS validation.

More endpoints and objects will need to be added. But I also hope that people
interested in the project will comment on the repo and make pull-requests.

### API server

Most likely, this will be build on top of Dancer2, using a few plugins that deal
with the correct handling of HTTP request headers in requests and responses.

### Web Server

This will be a minimal viable product.

Themes used for this server should obtain their data through the REST api. And
since it has a strict separation between data and presentation, actually anyone
could build their own server that talks to the API ... when following the
Specification.

The server will also include a 'page' that enable a user to manage their own
settings and make payments.

### Improve Modules

There are some modules I already published on CPAN, but need a little more after
care. Some lack proper tests, some actually fail after changes in Dancer2.

I've written those modules, because I felt they were needed to make this entire
project succeed. But the Perl Dancer developers have a good level of minimum
requirements before being acknowledged and accepted as real Dancer2 plugins.

LWP::UserAgent::Caching, and it's dependancy HTTP::Caching, require a bit more
development. So far, it's the only RFC compliant library.

The module will be needed in the CLI client. The REST api will only serve
'simple' (JSON) objects. And because off that, they are cacheable. But there was
no UserAgent that was capable of handling HTTP requests or responses as it was
defined by the RFC's.

## Inch-stones

### Act-out-of-the-Box

- create script to add a conference and checkout wit GitHub
- save as a Vagrant Box
- write documentation for organisers and developers

### Open API Specification

- add specification for `user`
- add specification for `talk`
- add specification for `confernce`
- add specification for `attendee`
- add specification for `...`

### Improving Modules

- Dancer2::Plugin::HTTP::AuthExtensible

    probably quite some rewrite, but there is quite some willingness to support that
    work from the Dancer Core developers

- Dancer2::Plugin::HTTP::Caching
    - add tests and anything else that would enhance the kwalitee
- Dancer2::Plugin::HTTP::ConditionalRequest
    - add tests and anything else that would enhance the kwalitee
- Dancer2::Plugin::HTTP::ContentNegotiation
    - add tests and anything else that would enhance the kwalitee
- HTTP::Caching
    - Responses should have `Age` header
    - successful `Delete` request should delete the cached data

### Moo Classes for business and transport objects and data access objects.

This requires a little chat with some respectable Perl developers. But basically
I want a true separation of the business object and how it has been retrieved or
stored or updated.

This makes it possible to use the same object on the CLI and on the server, with
the only difference that the former talks to the REST api and the latter to the
database.

### API Server

- create simple route
- add authentication and authorization
- add more routes as described in the Open API specification

### CLI

- figure out how to easy build a CLI from a Swagger file

    Tina Muller has done some tremendous work on that

- get it working for a `user` object
- add more command options

    as the Specification is becoming more complete, and the API get more endpoints,
    the CLI will grow with it.

### Theme Based WEB Server

This requires some more investigation, but I have very strong connection with
some very reputable front end developer that knows how to build a front end.

But it basically comes down - when using Angular2 with Bootstrap 4 to just go
along with the growth of the API.

## Project Schedule

The project is quite big, but within the first few weeks, it will at least have
made good progress on the first two deliverables.

The hard part will be to get an initial 'prototype' of the REST api working. I
expect that to have finished within two months after having the Vagrant Box up
and running.

The CLI will follow probably another month later.

From there on, it's a matter of adding more endpoints and command options as the
Swagger files gets updated.

A first release of the web-server will probably follow four to six months after
the CLI.

Since the size of the project, it would be about a year before the MVP will be
ready.

## Completeness Criteria

Well, a minimal viable product, a working theme-based website. It should at
least cover 80% of the common user stories for a general conference attendee.
This includes a personal web-page to edit and manage settings.

For the other admin tasks, I would stick to a proper CLI for now.

## Bio

I have initially started work on this project some years ago and am very eager
to continue work on this.

I've done a lot of research on REST api's and given some presentations on that.

In the last three years I've also organised three Perl Workshops in the
Netherlands - and was the main organiser of The Perl Conference in Amsterdam.

This has given me a lot of insight in how Act has been designed and what the
difficulties are with the current legacy version.

A year ago I did give a presentation during the Perl Dancer conference. This
was a summary of my work done so far and how all the puzzle pieces would fit
together to build Act-Voyager.

Compared to three years ago, I've grown much more in experience as a Perl
Developer. This new application is much better than the previous and based on
personal experience working on on the project intermittently.

While organising TPCiA, I was facing quite some difficulties that caused delay
in further development of Act Voyager. Now that has finished, time has come
available.


