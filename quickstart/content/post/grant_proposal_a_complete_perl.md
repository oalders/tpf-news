
---
title: "Grant Proposal: A Complete Perl 6 Course with Exercises"
author: Coke
type: post
date: 2019-06-10 00:00:00 +0000 UTC
url: "/post/grant_proposal_a_complete_perl"
categories:
 - Grants
 - Sign in

---

The Grants Committee has received the following grant proposal for the May/June 2019 round.
Before the Committee members vote, we would like to solicit feedback from the Perl
community on the proposal.

Review the proposal below and please comment here by June 14th, 2019.
The Committee members will start the voting process following that.

# A Complete Perl 6 Course with Exercises

- Name

    Andrew Shitov

- Amount Requested

  USD $10,000

_Task: Create a complete course of the Perl 6 programming language._

## Abstract 

I want to create a complete course with exercises covering all aspects of Perl 6. It will be aimed at everyone familiar with programming. The goal is to make a course that you can use in self-studying or as a platform in the class.


## Target audience

Perl 6 is a language that many people find extremely attractive. The efforts of our activists during recent years show that there are people from outside of the Perl community who also want to start learning Perl 6. 

There are two groups of potential users: those with and without Perl 5 background. As Perl 6 significantly differs from Perl 5, both target groups can benefit from a single Perl 6 course.

## Unique points

I am willing to create a *course* (not documentation), which is primarily intended to be read in order from page 1 to the last page. It is aimed at those, who are not familiar with Perl 6 and want to go through all the features of the language from the simplest to the most advanced.

The second target group is an offline class where the course can be used as teaching material. You just follow it, and you are happy with the order of topics, and as soon as the lesson turns from theory to practice, here you are: you have a list of problems to solve in the class or at home.

How this is different from what already exists: vast online documentation, books, etc.? The proposed course is a step-by-step flow that begins from simple things, which makes it different from the documentation. Unlike the books, the main focus will be on having small lessons with many exercises. There are also a few video lectures and introductions, but again with very little homework.

My idea is aligned towards online courses such as [perltuts.com](http://perltuts.com) (not completed for Perl 6) and [www.snakify.org/en](http://www.snakify.org/en/) that I used myself to learn and teach Python.

## Content

The course contains 16 sections. Each section includes 15-30 lessons. Each lesson covers a single topic (such as accessing array elements or using different variants of multi-methods, or a regex quantifier, or an aspect of concurrent code) and includes 2-4 exercises with displaying correct solutions on request.

The work is divided into five parts.

### Part 1

1. Basic introduction to Perl 6 and its compiler
    1. How to install Rakudo
    1. How to run it from console or from a web service such as [glot.io](https://glot.io)
    1. Notes on using Unicode
    1. Simple input and output

1. Running Perl 6
    1. Command-line options
    1. Phasers
    1. Understanding error messages
    1. The MAIN subroutine

1. Variables and data types
    1. Types of variable containers in Perl 6 (scalar, arrays, etc.)
    1. Understanding Perl 6 containers    
    1. Integers
    1. Strings
    1. Date and Time built-in support
    1. Other data times
    1. Understanding a sequence

1. Operators
    1. Types of operators in Perl 6 (infix, prefix, etc.)
    1. Overview of operators in Perl 6
    1. Meta-operators
    1. User-defined operators
    1. Data type conversion (e.g., prefix operators "+" or "~", etc.)

### Part 2

1. Control flow
    1. Blocks and scoping 101
    1. Boolean data type and type casting
    1. Conditional checks
    1. Loops
    1. Block-related phasers (e.g. LEAVE)
    1. Other options (e.g., gather, given)

1. Subroutines
    1. Creating a subroutine
    1. Signature
    1. Calling a subroutine
    1. Multiple dispatch
    1. Multiple MAIN subroutines
    1. Nested subroutines
    1. Anonymous subroutines

1. Modules
    1. Creating modules
    1. Using modules
    1. Different types of importing ("import", "need", etc.)
    1. Introspections
    1. Installing modules from web

### Part 3

1. Object-oriented programming
    1. Classes in Perl 6
    1. Attributes
    1. Methods
    1. Class methods
    1. Subroutines vs methods
    1. Inheritance
    1. Roles
    1. Introspection

1. Input and output
    1. Standard input, output, and errors
    1. Working with files
    1. Working with directories
    1. File streams

1. Exceptions
    1. The "try" block
    1. What is a soft failure
    1. The CATCH phaser
    1. Exception objects
    1. Failure objects
    1. Multiple dispatch in handling exceptions
    1. Custom exceptions

### Part 4

1. Regexes
    1. Literals and character classes
    1. Regexp matching
    1. Quantifiers
    1. Captures
    1. Alternations
    1. Anchors
    1. Forward and backward assertions
    1. Adverbs (such as :g etc.)
    1. String substitution and replacement

1. Grammars
    1. What is a grammar
    1. Creating grammars
    1. Rules
    1. Tokens
    1. Grammars vs. classes and inheritance
    1. AST (Abstract syntax tree) ("make", "made")
    1. Actions
    1. Inline actions vs. action class

### Part 5

1. Functional programming
    1. Recursion
    1. Reduction
    1. Higher-order functions
    1. Lambdas
    1. Data feeds
    1. Iterators
    1. Lazy and infinite sequences

1. Concurrent programming
    1. Junctions
    1. Threads
    1. Promises
    1. Channels

1. Reactive programming
    1. Supplies
    1. Live and on-demand supplies
    1. Understanding "react"
    1. Understanding "whenever"
    1. Understanding "await"

1. Web programming
    1. Making remote connections
    1. Simple HTTP client
    1. Simple HTTP server
    1. Cro 101


## Timeline and deliverable chunks

The project needs about six months to complete all sections; independent sections are published earlier, upon completion. The content will be uploaded to a GitHub repository in Markdown format.

## Licensing

The course will be publicly available and will be published and released under the Creative Commons license. It will be allowed for everyone to freely use and modify it, both commercially and not.

The course can be used by anyone who want to prepare their own teaching class on Perl 6 and related topics (such as 101, or grammars, or concurrency).

The course will contain hints for the trainer of which parts should be used and included in a one-day, three-day or a five-day class.

## Hosting

The materials will be submitted to a GitHub repository ([github.com](https://github.com) â€” GitHub is a web-based hosting service for version control using Git) and will be hosted by the GitHub Pages service.

For the end user, the course will be accessible online. For the potential future contributors, it will be accessible via pull requests.

The Perl Foundation can offer a different place of hosting static files if that will benefit the Perl community.

## About me

I am a Perl 5 developer since around 1998 and a Perl 6 enthusiast since around 2000.

I have/had run a few Perl 6 blogs:

- [Perl6.Ru](https://perl6.ru)
- [Perl 6 Online](https://perl6.online)
- [Perl 6 One-Liners Advent Calendar 2018](https://perl6.online/category/advent-calendar/)

I wrote three Perl 6 books:

- [Perl 6 at a Glance](https://deeptext.media/perl6-at-a-glance), which is the first book in the new season of Perl 6
- [Perl 6 Deep Dive](https://www.packtpub.com/application-development/perl-6-deep-dive)
- [Using Perl 6](https://deeptext.media/using-perl6), which is a collection of exercises aimed to understand the nature and beauty of Perl 6

I gave a lot of Perl 6 talks at different events including FOSDEM 2019. A couple of the recent talks being:

- [Perl 6 as a new tool for language compilers](https://www.youtube.com/watch?v=lwIXF25KJCo) at FOSDEM 2019
- [Creating a compiler in Perl 6](https://www.youtube.com/watch?v=YWTmd4Hdfa4)

I prepared and ran a few classes on Perl and other programming languages, including:

- [Introduction to Perl 6](http://act.yapc.eu/gpw2018/workshops/perl6-intro.html) at the German Perl Workshop 2018
- [Python 3 3-day course](https://github.com/ash/python-tut)
- [Python in 5 days](https://github.com/ash/python-in-5days)
- [5-day C++ course](https://github.com/ash/cpp-tut)
- [2-day JavaScript course](http://github.com/ash/js-tutorial)

There are a lot of exercises and materials prepared specially for Perl 6, including:

- [Perl 6 challenges](https://github.com/ash/p6challenges)
- [Perl 6 tests](https://github.com/ash/perl6tests)
- [Perl 6 assorti](https://github.com/ash/perl6-assorti)
- [Perl 6 at a glance](https://github.com/ash/perl6-at-a-glance)
- [Migrating to Perl 6](https://github.com/ash/perl5to6)
- [Solving Euler problems](https://github.com/ash/projecteuler)

I have organises a number of events dedicated to the Perl programming languages. The event map covers 8 countries spanning [the distance of about 10,000 km](https://www.google.com/maps/dir/Riga,+Letland/Vladivostok,+Kraj+Primorski,+Rusland/@44.4701035,44.7978367,3z/data=!3m1!4b1!4m14!4m13!1m5!1m1!1s0x46eecfb0e5073ded:0x400cfcd68f2fe30!2m2!1d24.1051865!2d56.9496487!1m5!1m1!1s0x5fb39cba5249d485:0x186704d4dd967e35!2m2!1d131.8869243!2d43.1198091!3e0) between the most Western and the most Eastern city. Among the events, there are three YAPC::Europe conferences ([YAPC::Europe](http://yapceurope.org) conferences are the biggest annual Perl events in Europe):

- [YAPC::Europe 2011](http://act.yapc.eu/ye2011)
- [YAPC::Europe 2013](http://act.yapc.eu/ye2013)
- [PerlCon 2019](https://perlcon.eu/)


## Financial

The requested amount is US$2000 per each of the five content parts, thus US$10,000 in total.


