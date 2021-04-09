
---
title: "Grant Proposal: Create a complete course of the Perl 6 programming language"
author: Coke
type: post
date: 2019-03-16 00:00:00 +0000 UTC
url: "/post/grant_proposal_create_a_comple"
categories:
 - Grants
 - Sign in

---

The Grants Committee has received the following grant proposal for the March/April 2019 round.
Before the Committee members vote, we would like to solicit feedback from the Perl
community on the proposal.

Review the proposal below and please comment here by March 22nd, 2019.
The Committee members will start the voting process following that.

# A Complete (Interactive) Perl 6 Course with Exercises

- Name

    Andrew Shitov

- Amount Requested

  USD $9,996

_Task: Create a complete course of the Perl 6 programming language._

## Abstract 

I want to create a complete course with exercises covering all aspects of Perl 6. It will be aimed at everyone who is familiar with programming. The goal is to make a course that you can use in self-studying or use as a platform in the class.


## Target audience

Perl 6 is a language that many people find extremely attractive. The efforts of our activists during recent years show that there are people from outside of the Perl community who also want to start learning Perl 6. 

There are two groups of potential users: those with and without Perl 5 background. As Perl 6 significantly differs from Perl 5, both target groups can benefit from a single Perl 6 course.

## Unique points

How this is different from what already exists: vast online documentation, books, etc.? The proposed course is a step-by-step flow that begins from simple things, which makes it different from the documentation. Unlike the books, the main focus will be on having small lessons with many exercises. There are also a few video lectures and introductions, but again with very little homework. Neither I want to have an extended version of the documentation (see, for example, [learn-perl.org/en/Operators](http://learn-perl.org/en/Operators) as an example of the tutorial with long lists of features).

My idea is aligned towards interactive courses such as [perltuts.com](http://perltuts.com) (not completed for Perl 6) and [www.snakify.org/en](http://www.snakify.org/en/) that I used myself to learn and teach Python.

## Content

The course contains approximately 15 sections (roughly following the chapters of "Perl 6 Deep Dive"). Each section includes 15-40 lessons. Each lesson covers a single topic (such as accessing array elements or using different variants of multi-methods, or a regex quantifier, or an element of concurrent code) and includes 2-4 exercises with displaying correct solutions on request.

## Timeline and deliverable chunks

The project needs about six months to complete all 15 sections; independent sections can be published earlier, upon completion.

The work is divided in four parts. The first three parts are about content, the fourth is to make it interactive.

1. Materials and exercises for the chapters devoted to the aspects of Perl 6 as a general programming language (thus, variables, functions, file operations, object-oriented programming, etc.).

1. Chapters about regular expressions and grammars.

1. Everything else. This part covers topics such as concurrent, functional, and reactive programming, etc. In other words, everything that is beyond the general programming language from point 1.

1. Implement the exercises from the three above parts as interactive online pages.

The result of the first three parts is a GitHub repository with lessons in Markdown format.

## Interactivity

We will need to host it on a subdomain of perl6.org and/or link at the Resources page of the site. The technical implementation should use [JJ's docker image](https://github.com/JJ/alpine-perl6) to spawn a compiler for each user. I suggest postponing the question of who pays for the hosting. Potentially we can re-use perltuts's engine. In a minimal form, exercises do not have to be really interactive so that we can gain both time and money on programming the server side.

Additionally or