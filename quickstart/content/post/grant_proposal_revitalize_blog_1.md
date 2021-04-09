
---
title: "Grant Proposal: Revitalize blogs.perl.org"
author: Coke
type: post
date: 2017-06-06 00:00:00 +0000 UTC
url: "/post/grant_proposal_revitalize_blog_1"
categories:
 - Grants
 - Sign in

---


The Grants Committee has received the following grant proposal for the May/June round.
Before the Committee members vote, we would like to solicit feedback from the Perl community on the proposal.

Review the proposal below and please comment here by June 15th, 2017.
The Committee members will start the voting process following that
and the conclusion will be announced approximately in one week.

# Revitalize blogs.perl.org

- Name:

    Andr� Walker

- Amount Requested:

    USD 3,000

## Synopsis

[blogs.perl.org](http://blogs.perl.org) is in need of replacement. I
am applying for finishing the work to replace it.

## Benefits to the Perl Community

For a long time now, BPO (blogs.perl.org) has suffered from
[numerous issues](https://github.com/blogs-perl-org/blogs.perl.org/issues). It
is running on a custom MovableType which we cannot fix, has many
bugs in its behavior, and frankly requires being replaced.

There was a grant request to this effect in 2015, but development
stalled, and I would like to complete the work.  What I
propose to deliver will have many of the same benefits to the Perl
Community:

- A simple, stable blogging platform for the Perl Community
- Migrated content from the current system
- Open sourced, hackable, publicly available on Github

## Deliverables

- Host blogs.perl.org source on a public GitHub repository
- Add continuous integration and testing
- Migrate MovableType.org database to blogs.perl.org schema (posts, comments, and users)
- Basic set of features:
    - Registration of new users
    - Login / Reset password (something missing in BPO now)
    - Content editing: Create / Update / Delete
    - Support for tags and categories
    - Support for comments (create, delete)
- Host a preview version for current users to test
- Release and publish final blogs.perl.org
    - Research and resolve hosting platform
- After release, respond to issues as they come up

## Project Details

There are currently at least three forks of the PearlBee project. The original
project, a fork for the 2015 grant (intended for post-grant merge), and another
created by myself with additional people from Booking.com.

I plan to review the differences between the BPO-specific version and the
upstream and merge any useful commits. I intend to remove the snowflake
nature of BPO and provide a general-purpose blogging platform software
in Perl which could be used outside BPO.

This blogging platform (written to host BPO) will be written on top of Dancer2,
DBIx::Class, and DBI, with a PostgreSQL database imported from the
existing BPO, and session management provided by Memcached. The nature
of the work will allow making adjustments to this setup, as I will
implement generic interfaces instead of depending on specific
technologies (PostgreSQL, Memcached, etc.).

Unit and integration tests will be performed through Travis CI, and source code
made available on GitHub.

I'll begin by reviewing the code of both forks and preparing a suitable
base. The fork on which I worked provides a good start, as it cleaned
up the code considerably and changed the API endpoints to a more
consistent interface.

On top of this, I will work on cherry-picking any valuable commits from Evozon's
official and BPO forks, as noted above.

Next, I'll make sure all the basic features work, preferably covered by tests.
Abilities like registering, adding a post, editing, commenting, etc.

Then I would focus on migrating the existing database into the new schema. As
the previous grant request stated, it's possible to reuse the password hashing
scheme so people can log in without having to reset their passwords. If I
find this to be an issue, we will work on a plan for migrating users' accounts
without breakage. Existing user accounts will be covered in the grant
work regardless of which approach I take.

I will maintain support for the following co