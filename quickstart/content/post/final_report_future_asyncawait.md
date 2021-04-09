
---
title: "Final Grant Report: Future::AsyncAwait"
author: domm
type: post
date: 2019-10-10 00:00:00 +0000 UTC
url: "/post/final_report_future_asyncawait"
categories:
 - Grants

---

# Future::AsyncAwait Final Report
 
Here is the final report by [Paul Evans](http://leonerds-code.blogspot.com) on his [Future::AsyncAwait Grant](http://news.perlfoundation.org/post/grant_proposal_futureasyncawai):
 
This project set out to improve the [Future::AsyncAwait](https://metacpan.org/pod/Future::AsyncAwait) module, fixing a number of known bugs and adding missing features. It also aimed to improve documentation and user-awareness of the new abilities that the module adds to Perl. I believe it has been successful in all of these areas.

## Bugs Fixed

Of particular note, the original proposal made reference to four bugs. These have all been fixed:

* https://rt.cpan.org/Ticket/Display.html?id=126037
* https://rt.cpan.org/Ticket/Display.html?id=126036
* https://rt.cpan.org/Ticket/Display.html?id=125613
* https://rt.cpan.org/Ticket/Display.html?id=123062

In addition, many others have been found and fixed along the way. The full list can be found  [in the RT queue](https://rt.cpan.org/Public/Dist/Display.html?Status=Resolved;Name=Future-AsyncAwait). A few bugs notable for over-average size are RT 128176, 128222, 128619 and 129215.

As the initial tasks were completed far quicker than originally estimated, I spent some of the time working on some missing features that were also in the RT queue.

The largest of these is documented by RT 123465; the integration of Perl's `sub` signatures with the `async sub` syntax. Fixing this involved some work on core perl itself to expose a new API function, `parse_subsignature`. This is mentioned in https://www.nntp.perl.org/group/perl.perl5.porters/2019/07/msg255512.html.

Another new feature was found to make integration with `Syntax::Keyword::Try` or other exception-catch mechanisms simpler; documented in RT 129373.

## Blog Posts

Over the course of this project, I have written a number of blog posts:

* [Awaiting The Future](http://leonerds-code.blogspot.com/2019/04/awaiting-future.html> ) - An overall introduction to `async/await` and the [Future::AsyncAwait](https://metacpan.org/pod/Future::AsyncAwait) module. Gives the state of the project at the time, listing several bugs now fixed and missing features now added.
* A series of near-identical copy-and-pasted articles, comparing the `async/await` syntax between perl and four other mainstream languages, highlighting the similarities and differences between them:
    * http://leonerds-code.blogspot.com/2019/08/asyncawait-in-perl-5-and-ecmascript-6.html,
    * http://leonerds-code.blogspot.com/2019/08/asyncawait-in-perl-5-and-python-3.html,
    * http://leonerds-code.blogspot.com/2019/08/asyncawait-in-perl-5-and-c-5.html,
    * http://leonerds-code.blogspot.com/2019/08/asyncawait-in-perl-5-and-dart.html.

I also wrote a series of posts on the [Binary.com Tech Blog](https://tech.binary.com) on the subject of how to migrate existing code structures using legacy `Future` or `Future::Utils` techniques into using `async/await` syntax:

* https://tech.binary.com/migrating-to-async-await-1/,
* https://tech.binary.com/migrating-to-async-await-2/,
* https://tech.binary.com/migrating-to-async-await-3/.

I have not been able to attend a Perl conference during this project, so my plan of delivering a talk on this subject has not been achieved. I do however aim to turn my idea into another post, on a theme of how the `async/await` syntax makes safe concurrency by whitelisting the places where context switches can occur, rather than the more traditional approaches like mutex blocks and locking, which are more comparable to blacklisting.

## Documentation

Finally, as part of the research phase several details about data structures internal to the perl interpreter were discovered, and written down in the form of some documentation that was eventually merged into the core perl tree alongside the rest of the internals documentation.

https://www.nntp.perl.org/