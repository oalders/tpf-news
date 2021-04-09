
---
title: "Devel::Cover Grant Report for June"
author: mdk
type: post
date: 2012-07-20 00:00:00 +0000 UTC
url: "/post/develcover_grant_report_for_ju"
categories:

---

<big>by Paul Johnson</big>

In accordance with the terms of my grant from TPF this is the monthly report for my work on improving Devel::Cover covering June 2012.

This month I released Devel::Cover 0.88 and 0.89.

I looked into "https://rt.perl.org/rt3//Public/Bug/Display.html?id=113464":https://rt.perl.org/rt3//Public/Bug/Display.html?id=113464 which is a perl bug, but is related to warnings when running Devel::Cover.  I supplied a core patch to quieten the warnings which has been applied, but I also changed Devel::Cover to turn off warnings when calling code in B::Deparse since the consensus seems to be that we probably have better things to do with our time than ensure that B::Deparse is warnings free.

I spent some time on cpancover.  I received some feedback from my initial work with suggestions for improvements, which I implemented.  I also automated the process of generating the coverage run with a view to running it once or twice a week from cron.  Perhaps it could even run daily now that I have also compressed the results and configured nginx to serve those.

At the moment the most recent results are found at
"http://cpancover.com/latest/index.html":http://cpancover.com/latest/index.html which is redirected to from "http://cpancover.com":http://cpancover.com but I need to come up with a main page linking to the available versions and perhaps some way of keeping links alive for a reasonable amount of time.

It would also be nice to get something like Thomas Klausner's App::ArchiveDevelCover ("https://github.com/domm/App-ArchiveDevelCover":https://github.com/domm/App-ArchiveDevelCover) running to track coverage changes to modules.  I have got many more ideas, but they are for another day.

I also tested Devel::Cover against 5.17.0.  I had some cpantesters failure reports against 5.17.0 but I was not able to reproduce them, so I will need to look further into that.

I looked into perl #113690 and #113684 which were about perl commit 6a31dbf44ee919c340a3372c95b28d581979d165 breaking Devel::Cover 0.89. This commit removed some superfluous parentheses from B::Deparse output which caused Devel::Cover output to change and hence tests to fail.

I had planned to fix this as soon as perl 5.17.1 was released, but Father Chrysostomos realised that the fix was incomplete and needed further work to get the precedence levels correct.  So he reverted the change and suggested I wait until the complete fix is committed.

As to why I was planning on waiting until 5.17.1 was released before fixing Devel::Cover, perl is a moving target as far as Devel::Cover is concerned and I choose to only worry about stable releases since 5.6.1 and development releases with version numbers above the most recent stable release.  Arguably, that's already too many versions, but it's not causing a problem at the moment.

As expected 5.17.1 was released, and Devel::Cover passes all its tests with it.

And since I've started working on this grant a few people have approached me with ideas and suggestions, and with offers of help.  This is great and something I had hoped would happen.  So I have also spent a little time discussing these things and making future plans.

The work I have completed in the time covered by this report is:

Closed RT tickets:

* 75633 Directory names with ++ not supported
* 77598 [PATCH] Fix test failures when the build directory contains regexp metacharacters.
* 77599 [PATCH] POD errors
* 61515 Cover.pm messes with cwd
* 70162 cover fails with spaces in pwd
* 69562 cpancover won't run without Parallel::Iterator

Closed Github tickets:

* 16 cover -coverage documentation
* 17 #line directives
* 19 Add some math modules
* 21 Add links from subroutine/branch etc reports to the main report file
* 22 Database permissions
 

Github tickets worked on:

* 18 Seems to hang on C file

Merged pull requests:

* 19 Add some math modules

Fixed cpantesters reports:

* "htt