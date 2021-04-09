
---
title: "Grant Proposal: Raku Ecosystem"
author: Jason A. Crome
type: post
date: 2020-07-30 00:00:00 +0000 UTC
url: "/post/grant_proprosal_raku_ecosystem"
categories:
 - Grants
 - Sign in

---

## Raku Ecosystem

- Name: tony o'dell
- Amount Requested: $12,000

## Synopsis

Redesign the raku/zef ecosystem to be robust and to make easier the distribution submission for the raku ecosystem.

## Benefits to the Raku Community

Currently the process for maintaining the ecosystem in raku is either uploading to cpan, which comes with its own set of limitations as cpan was not designed to handle the way raku uses distributions (the same distribution name can be used multiple times and pared down by the consumer by using :auth and/or :ver).  The other way this is handled right now is through a github repo containing a file that the distribution authors must update in order keep the ecosystem fresh, which comes with its own set of challenges and barrier to entry for the user.

This project would create an ecosystem that is both friendly to the end user, provides secure access and storage to distribution consumers, and promote the development of distribution from raku distribution authors.

## Deliverables

- Fault tolerant ecosystem
- Provide hooks for author tooling

## Project Details

The primary deliverable of this project would be a fault tolerant ecosystem that is both consumable via zef (so, a zef plugin) and a website similar to metacpan for browsing and finding packages.  Guidelines for distribution authors and tooling for distribution authors to test the quality of their upload.

The secondary deliverable would be to create an expandable API for further development (testing, quality checks, health checks, etc).  It is limited in scope to the design accomodation and not the implementation of what the hooks will or could be used for.

The operating costs of this project will be paid for by donations from the community or out of my pocket.  Donations exceeding operating costs will be used to further develop tooling and expand on the secondary deliverable.

The domain for this project has been suggested to use zef.pm but this is not set in stone and this proposal is open for using a different domain.

## Project Schedule

The primary deliverable of this project is reckoned to take about two months complete.  This includes the administration and automation of server maintenance, storage space, and writing the necessary code to store/save/retrieve these distributions both from a zef plugin (or another CLI) and from the WWW.

The secondary deliverable is likely to take closer to a month to design and build.

## Bio

I am Tony O'Dell.  I have written a good number of raku modules and have been writing in perl for about 15 years.  I have sys admin experience and am initial commiter and co-author of zef.  Prior to writing software as a full time job my primary area of expertise was in data warehouse management and design, and statistics.

# Addendum

The ecosystem in the context of this grant is meant to mean a repository where packages can be downloaded/consumed by raku package managers and processes.  The structure of the repository will be similar to CPAN, maintaining the ability to be mirrored, with  modifications to handle the added complexity of `:ver`, `:auth`, `:api` in raku.

The API in context of this grant is meant to enable some process that allows authors to write to the repository and upload their dists (more analogous to PAUSE).  The functions of the API will be around user creation and uploading distributions to the ecosystem.

This project is in no way meant to replace modules.rakudo.org or metacpan.  It will make modules.rakudo.org's job much easier in displaying/searching for modules.

## Milestones

1. Design and architecture of dist storage on servers, permissions, indexing, and mirroring
2. Build out for distribution with a `sandbox` environment and tests
3. Writing tooling around the API to manage users and allowing users to manage their dists available in the repository.  This milestone includes the zef