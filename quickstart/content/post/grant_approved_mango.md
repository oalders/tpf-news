
---
title: "Grant approved:  Mango"
author: Curtis "Ovid" Poe
type: post
date: 2006-08-22 00:00:00 +0000 UTC
url: "/post/grant_approved_mango"
categories:
 - Grants
 - Sign in

---

Here's Christopher Laco's "Mango" grant application.

NAME:
Christopher H Laco

EMAIL:
claco@chrislaco.com

PROJECT:
Mango

AMOUNT:
$1500

DURATION:
2-3 Months, starting immediately.

SYNOPSIS:
Mango is a Web 2.0 Ecommerce Solution built using the Handel and Catalyst
frameworks.

PREFACE:
Since graduating college, I've spent the last 10 years working on various
forms of ecommerce solutions. Most of that time has been in connecting
various website checkout processes to backend business systems using
custom created socket clients. I've also been involved in the other website
programming aspects as well as internal applications used for administration
of the ecommerce sites themselves.

During that time, I also ran a small web hosting service out of my residence
using *cough*Windows, ASP, FreeBSD, Perl, AxKit, and now Catalyst. The need
to convert an existing ASP based ecommerce site to *nix and the apparent
lack of basic cart/order/checkout modules on CPAN gave birth to Handel.
Handel is a set of modules to do cart|order based operations, as well as a
plugin based checkout order processing pipeline that can be configured
differently per app, instance, or even process call.

Handel has recently undergone a major refactor to make use of DBIx::Class
and the flexibility it provides. It now allows people to make use of
existing schemas as well as the ability to customize the use of Handel
towards each websites needs. Along with the refactor, I've also stayed tried
and true to the TDD model of development, and the large test suite in Handel
has helped ensure a certain level of stability during the refactor.

While Handel itself includes some basic Catalyst helpers to generate some
basic code, there has been interest from others and myself in providing a
larger ecommerce "out of the box" solution.

DESCRIPTION:
Admittedly, the Mango project is in its planning stages. A rough list of
the features can be found at:

http://handelframework.com/wiki/Mango::ProposedFeatures

Mango itself will be built in 2 distinct layers: the web layer, and the
application core.

The core will be a set of reusable modules that contain no web specific
logic. As such, they should be reusable in any application that wants to
administer Mango based solutions, be it a GUI app, a cron job, SOAP|XMLRPC
listeners, or the Catalyst website itself.

The web layer will of course be Catalyst, and the necessary MVC components
needed to allow the core components functionality to be utilized from the
web.

Because of Handel, most of the shopping cart, wishlist, order, and checkout
code is already in place. With the latest Handel version 1.0 pending, those
processes can easily use a schema that is customized and expanded to better
meet the needs or an entire commerce package such as this.

BENEFITS:
The main benefit of Mango, aside from the obvious application itself, is one
of promotion. In the realm of buzzwords like Web 2.0, MVC, Ruby on Rails,
and Django, I feel Mango would help publicize Perl and Catalyst in a way
that most people can relate to.

Additionally, Mango, coupled with Catalyst, Handel, and DBIx::Class
will provide a tremendous kickstart to anyone wanting to use Perl to power
their website to sell production the internet, or add ecommerce
functionality to existing Catalyst powered websites.

DELIVERABLES:
A Catalyst application that provides:

* Provider|Adapter Based API
* Use AJAX w/ Full Downgrade
* Shopping Cart|Wishlists
* Save|Restore
* RSS|Atom|Email A Friend
* Wishlists
* Product Pages
* Sales
* PayPal|Google Checkout Integration
* Shipping Estimates|Tracking
* Product Catalog Maintenance
* Complex Sales
* Buy X, Get Y Free
* Buy 1-5 at $X, 6-10 at $Y, ...
* Everything $X Off
* Everything %Y Off
* Administrative Configuration
* Users
* Products
* Orders
* Pages|Templates
* Roles|Permissions

ADDITIONAL RESOURCES:

http://handelframework.com/
http://svn.handelframew