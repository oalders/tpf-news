
---
title: "Grant Proposal: Curating and improving Perl6 documentation"
author: Coke
type: post
date: 2018-02-13 00:00:00 +0000 UTC
url: "/post/grant_proposal_curating_and_im"
categories:
 - Grants
 - Sign in

---

The Grants Committee has received the following grant proposal for the January/February round.
Before the Committee members vote, we would like to solicit feedback from the Perl community on the proposal.

Review the proposal below and please comment here by February 20th, 2018.
The Committee members will start the voting process following that and the conclusion will be announced the
last week of February.

# Curating and improving Perl6 documentation

- Name:

    JJ Merelo

- Amount Requested:

    USD 5,000

## Synopsis

The documentation of Perl6 at https://github.com/perl6/doc is an
impressive piece of work, but working on it has revealed some
problems. I propose to work part time for two months on improving this
documentation by

- working on all outstanding issues and address all of the issues in 2015, 2016 and 2017,
- creating a set of rules for existing and new contributions and enforce those rules in the test suite,
- assigning outstanding issues during the grant and create a system for assigning work to volunteers,
- creating an entry page with a tutorial for complete beginners with very little or no knowledge of programming, with the intent of providing a good landing page for the language
- contributing to Perl and general open source conferences with entry-level tutorials and learning material, and
- writing articles in websites such as The Practical Dev (https://dev.to), HackerNoon or Medium with practical "30 minutes" tutorials to perform usual tasks in Perl6.

In general, the main intention of this work is to attract new users or
retain them after they make their first attempts to work with the
language. So far and unfortunately, Perl6 has failed to attract users
and it stands in the lower 10% of the top 100 languages in GitHub
according to the number of pull requests. It has around 2 questions a
week in StackOverflow, and in general they are advanced questions,
with no space for first steps. Once again, documentation is not
development (or not fully), but it is essential to attract a critical
mass of developers, that will go on to develop the next killer
application that uses Perl6. It has got to reach from the entry-level
programmers to the more advanced ones. 

So my main objective, and the stick by which it will be measured, is
the attention gathered by the language and how the user base
increases. 

# Benefits to the Perl6 community and Perl community at large

Curating documentation is an integral part of the development
process. Besides including work on little code snippets, that have to
absolutely and positively work flawlessly, there are pieces of code in
the [perl6/doc](https://github.com/perl6/doc) repository used for
testing documentation; we will propose to also work on this aspect.

On the other hand, documentation is the entry
point for newcomers to a language, and it needs to be the very best to
increase the community of Perl6 users and obviously developers. A good
documentation will Y help developers find their way to use the best
perl6ish data structures and functions. Curation of the
documentation base will help to make it sustainable on the long term,
which will eventually benefit the communities at large.

# Project details

Here are the rest of the project details, starting with the
deliverables. 

## Deliverables

These is what I will deliver during and by the end of the grant.

- Addressing 5 doc issues per day on average, minimum 2 issues per day. Reduce the amount of open issues by at least 120.
- Write guidelines for contributions and write test code for enforcing it.
- Analyze landing pages and site paths and contribute to work on
preferred landing pages. Produce report to improve usability of the
site. 
- Contribute 1 entry-level tutorial per week, to be published preferably on non-perl publications (Medium, dev.to, hackernoon).
- Monitor StackOverflow and answer and/or identify needs to be addressed in documentation. Contribute to work on those pages or create new ones.
- Hang out at the IRC channel to help newcomers in general and
identify possible problems in documentation; address them when
needed. In general, monitor all possible channels where people might
need help with Perl6 and, besides helping when needed, try to identify
possible problems in the documentation.
- Refactor documentation tests and write new ones to make new additions follow guidelines. This is already in an issue in the GitHub repo.

## Project details and schedule

[Perl6 documentation](https://doc.perl.org) is an impressive piece of
work. There are close to 8K commits, close to 700 PRs and around 800
closed issues. The way it is organized and served allows it to be
searched very fast and it can be really useful.

However, from the point of view of the newcomer, he or she might find it difficult to get started on the language. Descriptions are sometimes quite technical and examples do not correspond to real use cases and in some other cases do not cover all cases.

Internally, it is mainly a volunteer effort and lots of people are
working hard reviewing it and keeping it relevant and, over all,
coherent and correct. But since it is mainly self-organized and
self-assigned sometimes it is difficult to get a particular thing
done. This means that there are issues open since 2015 (#93 since June
13th 2015), and almost 300 of them are still open. Even "low hanging
fruit" issues are hanging low for really a long time, with no
systematic effort to pick them up from the tree.

A bigger problem is that these issues are mostly self-generated by
people that have already contributed, which is a challenge since it
seems that people coming to the documentation are not bothered enough
to actually find a particular error and create an issue. Encouraging
this behavior would help to improve the documentation, and also
attract new documentors to this, even those who are not totally
proficient in Perl6.

Another problem I have found is that there is some lack of coherency
and consistency in the documentation, which should only be expected,
since every page has been written by a different set of
developers. For instance, some examples included in the documentation
can be directly copy/pasted and run. Some other are not, with several
examples lumped together into a code block. Output is written
sometimes in different lines, sometimes in the same line, almost
always as a comment. It would be convenient to have some rules, even
more convenient to enforce those rules adding a test to the test
suite; it would be a plus to make all examples runnable so that they
are tested along with the documentation and flagged when they become
obsolete or erroneous.

In general, the objective is to improve the documentation. Getting a
detailed log of the visits to the site will allow also to analyze the
path, see whether people are finding what they are looking for, and
optimize landing pages. A bit of referrer-specific code might also
help, but in general analyzing the effectiveness of the documentation
and reacting to it would also improve the user experience of the site.
This work would be done locally from Granada. I will also attend
Glasgow's YAPC and any other during the grant (for instance, the Dutch
Perl Workshop) if it's considered necessary. I can organize locally a
documentation hackathon by the end of the grant or do an online
squashathon.

## Project schedule

- Day 7: identify core of people working on the documentation and create an open channel of communication, for instance, a Telegram channel with the GitHub bot connected. Use it for meta-discussions. Then, daily: work on that channel to add a small layer of organization, assign tasks, discuss issues.
- Daily: monitor StackOverflow, Perl6 IRC channel, Twitter and other social networks for people needing help with Perl6. Identify main sources of problem.
- As-it-happens: check out PRs for documentation, engage them, accept them if adequate. Check out new issues, engage and assign them to self.
- Weekly: produce entry-level tutorials and publication on non-perl sites.
- Daily: work on outstanding perl6/doc issues, starting with the first. Target: 5 a day, minimum 2 a day.
- Day 15: produce a report on site use and usability. Day 30: implement improvements. Day 45: analyze improvements and produce report on site visits and interactions. Day 60: implement second round of improvements.
- Day 15: RFC of guidelines for documentation and code included with it. Day 30: code implementation of those guidelines.
- Day 45: organize online or offline Perl6 documentation hackathon in Granada (Spain) to address outstanding issues, improve code examples and in general add a layer of sustainability to the documentation effort.

### Report schedule

The reports will be done according to grant rules, and, additionally...

- The grant report will be integrated with the doc GitHub repo, as a ChangeLog or grant report folder. This ChangeLog will be updated as-it-happens.
- 4 reports, 2 a month, on activity happening outside the log:
tutorials published, StackOverflow activity. They can be published
wherever the committee sees fit. I can create a GitHub repo for the
grant, for instance, and publish it as github pages.

## Requested amount for the Project Grant

Whatever amount is the usual for two months, part time, work. The
contract will be signed through my employer, the University of
Granada. 

## Bio

Born in 1965, I have been using Perl since 1993 and Perl5 since its
release by the beginning of 2016. Currently I am professor in an
university in Spain, where I have been teaching since 1988. I teach
cloud computing and other related subjects.

I have authored "Learning to program with Perl6", where I actually
used Perl6 to put the book into production. I have been contributing
to the Perl devroom since 2013, and attended several YAPC, last one in
Cluj-Napoca. I also organized YAPC::Europe in Granada, and two
workshop (next one in one week) in the same city. 

I am mainly a Perl programmer, but I use Perl6 extensively; I
contribute to the Perl6 repositories since 2016. 

Juan J. Merelo, [JJ in GitHub](https://github.com/JJ), [JMERELO in CPAN](https://metacpan.org/author/JMERELO), [jjmerelo in Medium](https://medium.com/@jjmerelo), [JJ in The Practical Dev](https://dev.to/JJ).

### Country of Residence and nationality

Spain and Spanish. 

## Suggestions for Grant Manager

Liz Mathijsen, Wendy van Dijk, Jeff Goff, Simon Proctor, Tom Browder.

## SEE ALSO

- [https://docs.perl.org](https://docs.perl.org), main documentation page and its
[GitHub repository](https://github.com/perl6/doc).
- [Monthly statistics](https://github.com/perl6/doc/pulse/monthly)
for the GitHub repo, showing 42 active issues during that time.
- [My contributions to the repo so
far](https://github.com/perl6/doc/commits?author=JJ), with position 41
in the ranking.

## Copyright

    This file is released under the GPL. See the LICENSE file included in this distribution,
    or go to http://www.fsf.org/licenses/gpl.txt


