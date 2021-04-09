
---
title: "Grant Proposal: Revising Perl's open and perlopentut manual pages"
author: Coke
type: post
date: 2020-02-25 00:00:00 +0000 UTC
url: "/post/gp_jan_2020_open_docs"
categories:

---

## Title
Revising Perl's `open` and `perlopentut` manual pages for clarity and completeness

## Synopsis
The manual page for one of Perl's most central functions -- `open` -- has problems. While nobody doubts the page's technical accuracy, head Perl maintainer Sawyer X nonetheless holds it up as an example of documentation in sore need of reorganization and revision. Towards the end of his TPC 2019 talk "[Perl 5: The Past, The Present, and One Possible Future](https://www.youtube.com/watch?v=4wKyNYs7eZw)", Sawyer described the `open` page as an example of Perl core documentation that fails to put its most useful information in front -- making it less likely to be read at all, especially by a newcomer to Perl. "It's kind of like a Google result," he said. "You go through the first page, and that's it, you know?"

In a separate but related issue, the `open` page links early on to `perlopentut` as a "gentler introduction" to its concepts. However, that page has problems of its own -- most obviously, the fact that it ends with three sections containing only the text "To be announced" (and in one case adding "Or deleted"). Perl's public distributions have included this page, in its current form, since 2013. This marks it as another case of a core manual page that could greatly and obviously benefit from a fresh look.

I -- a writer and long-time Perl programmer -- propose to address these issues by taking a focused, critical look at these two manual pages, and then revising them for clarity, completeness, and friendliness to Perl newcomers.

## Your name
Jason McIntosh

## Benefits of the project to the Perl Community

This work would address known issues regarding the official manual pages for one of Perl's most important built-in functions. After this work is complete and delivered, Perl (and related documentations projects, such as the perldoc.org website) will be able to teach newcomers about its all-important `open` function with the clarity and immediacy they desire, while also allowing people seeking deeper knowledge to read on for a more detailed examination of the topic.

One imagines that this work, upon successful delivery, could also spark a side-effect of setting a positive example for future documentation improvement, ready for application to other areas of Perl's voluminous manual with similar problems. But we must leave that as a potential happy accident; this project shall focus only on the `open` pages.

## Deliverables
The deliverables, in their entirety, shall include documentation patches submitted to the P5P team regarding the source files for the `open` and `perlopentut` manual pages. As with any such proposal, I intend to continue working with P5P after submitting the initial drafts, responding to feedback and critiques until we arrive at a version of the revised documentation accepted for a future release of Perl.

The possibility exists that, even after discussion and revision, P5P will ultimately decline to make these changes to Perl's official documentation. Should we arrive at this (I hope unlikely) circumstance, then I would mark the project as complete but undeliverable.

## Project details and a proposed schedule
Upon acceptance of this grant proposal, I would introduce myself to P5P and make my intentions regarding this project known. I would then put my best writerly foot forward to revise and rearrange the content of `open`, keeping in mind a _primary_ audience of newcomers to the language, with awareness of only basic programming concepts assumed.

This effort would reorganize the content such that, within minutes, a reader will grasp the basics of using Perl's `open` function to read from and write to files, including the core concept of filehandles. The documentation will open with clear emphasis on the best practice of calling `open` with three arguments, and explain the importance of the common `o