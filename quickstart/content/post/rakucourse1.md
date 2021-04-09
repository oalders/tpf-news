
---
title: "Raku Course January 2021 Grant Report"
author: Nicolas R.
type: post
date: 2021-01-13 00:00:00 +0000 UTC
url: "/post/rakucourse1"
categories:
 - Grants
 - Perl 6 Development

---

This is a monthly report by Andrew Shitov on his grant for the Raku course.

Andrew is happy to report that the first part of the Raku course is completed and published. 
The course is available at [course.raku.org](https://course.raku.org).

The grant was approved a year and a half ago right before the PerlCon conference in Rīga. 
Andrew was the organiser of the event and had to postpone the course due to high load.
Also note that during the conference, it was proposed to rename Perl 6 which also lead to some confusion and wait before working on the grant.

After months, the name was settled, the distinction between Perl and Raku became clearer, and, more importantly, external resourses and services, e.g., [Rosettacode](http://rosettacode.org/wiki/Category:Raku) and [glot.io](https://glot.io/new/raku) started using the new name. 

It's now a perfect timing to start working on the course Andrew dreamed about a couple of years ago.
Andrew started the main work in the middle of November 2020, and by the beginning of January 2021, he had the first part ready.

The current plan includes five parts:

1. Raku essentials
1. Advanced Raku subjects
1. Object-oriented programming in Raku
1. Regexes and grammars
1. Functional, concurrent, and reactive programming

It differs a bit from the original plan published in the grant proposal. While the material stays the same, Andrew decided to split it differently. Initially, he was going to go through all the topics one after another. Now, the first sections reveal the basics of some topics, and we will return to the same topics on the next level in the second part.

For example, in the first part, Andrew only talk about the basic data types: `Int`, `Rat`, `Num`, `Str`, `Range`, `Array`, `List`, and `Hash` and basic usage of them. The rest, including other types (e.g., `Date` or `DateTime`) and the methods such as `@array.rotate` or `%hash.kv` is delayed until the second part.

Contrary, functions were a subject of the second part initially, but they are now discussed in the first part. So, we now have Part 1 “[Raku essentials](https://course.raku.org/essentials/)” and Part 2 “Advanced Raku topics”. This shuffling allowed Andrew to create a liner flow in such a way that the reader can start writing real programs already after they finish the first part of the course.

It is quite a tricky task to organise the material without backward links. In the ideal course, any topic may only be based on the previously explained information. A couple of the most challenging cases were ranges and typed variables. They both causes a few chicken-and-egg loops.

During the work on the first part, Andrew also prepared a ‘framework’ that [generates](https://github.com/ash/raku-course/blob/master/_includes/menu.html) [the navigation](https://github.com/ash/raku-course/blob/master/_includes/nav.html) [through](https://github.com/ash/raku-course/blob/master/_includes/toc.html) the site and helps with [quiz automation](https://github.com/ash/raku-course/blob/master/_includes/quiz.html). 

It is hosted as GitHub Pages and uses [Jekyll](https://jekyllrb.com) and [Liquid](https://shopify.github.io/liquid/) for generating static pages, and a couple of Raku programs to automate the process of [adding new exercises](https://github.com/ash/raku-course/blob/master/new-exercise-template.raku) and [highlighting code snippets](https://github.com/ash/raku-course/blob/master/highlight.raku). 
Syntax highlighting is done with [Pygments](https://pygments.org).

Returning the to course itself, it includes pages of a few different types:

* The theory that covers the current topic
* Interactive quizzes that accomplish the theory of the topic and/or the section
* Exercises for the material of the whole section
* Answers to the exercises

The quizzes were not part of the grant proposal, but Andrew think they help making a be