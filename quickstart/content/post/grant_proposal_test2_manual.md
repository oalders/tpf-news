
---
title: "Grant Proposal: Test2 Manual"
author: Makoto Nozaki
type: post
date: 2016-03-29 00:00:00 +0000 UTC
url: "/post/grant_proposal_test2_manual"
categories:
 - Grants
 - Sign in

---

The Grants Committee has received one grant proposal for the March round.  Before the Committee members vote, we would like to solicit feedback from the Perl community on the proposal.

Review the proposal below and please comment here by April 9th, 2016.  The Committee members will start the voting process following that and the conclusion will be announced by April 16th.

# Test2 Manual

- Name:

    Chad 'Exodist' Granum.

- Amount Requested:

    USD 2,000

## Synopsis

I am proposing to write a Test2 manual. The Test2 distribution already has very
complete module documentation. The manual will build off the individual module
documentation and provide very valuable integration and usage details.

## Benefits to the Perl Community

Test tool development has been held back for a very long time now. Test2 should
unblock a lot of areas related to testing, but only if people can figure out 
how to use it. This manual should make current and future test developments
accessible to all perl developers.

## Deliverables

- Test2::Manual

    A brief introduction and table of contents.

- Test2::Manual::Tooling

    This section will cover writing test tools. This would be a root document with
    several other page links.

    \* The name of the manual page, and the manual layout will need to be finalized.

- Test2::Manual::Maintenance

    This section will cover maintaining Test2 itself. This will be a root document
    pointing to pages that explain the internals, and how they work together.

    \* The name of the manual page, and the manual layout will need to be finalized.

## Project Details

Test2 is intended to be the new core of perl testing tools. Test2
itself is now on cpan, and Test::Builder will be updated to use Test2
under the hood (I have an ongoing grant for this that is near-completion).
Something this important needs serious documentation.

This manual will explain how all the bits and pieces work together. The manual
will target 3 specific audiences:

- Test Tool Authors

    This section of the manual will cover people who want to write new tools or
    that use Test2 under the hood. This is probably the most important audience as
    it will be necessary if tools are going to be compatible with each other.

- Test2 Maintainers

    This section of the manual will cover people who wish to participate in Test2
    development directly. This is important to reduce the bus factor (which is very
    close to 1 currently).

## Inch-stones

For test authors:

- Migrating from Test::Builder

    Conversion notes and examples for people moving from Test::Builder. This will
    detail how tools might have changed, what is gone, and introduce some new or
    replacement concepts.

- Simple OK tool

    This is the most basic tool you can write, and is valuable for explaining the 
    key concepts universal to all Test2 tools.

- The Test2 API

    This covers [Test2::API](https://metacpan.org/pod/Test2::API), and all the functionality it exposes.

- The 'Context' object

    Explain what the context object is, why it is important, and how to use it
    effectively.

- The 'Hub' object and API

    Explanation of the hub objects and how they work.

- Custom Hubs

    How to write a hub subclass (essential for Subtest like tools).

- Custom event types

    How to write an event.

- Custom output formatters

    How to write an output formatter.

- Writing IPC drivers

    How to write a custom IPC Driver.

For Test2 maintainers:

- Component map

    Map of all Test2 components.

- Basic building blocks

    Explanation of low-level infrastructure such as Test2::Util::HashBase.

- How the 'Context' works

    Detailed overview of the Context object, and implementation details.

- The hub stack

    What the hub stack is, and why it is important.

- The IPC system internals

    How IPC works, and more importantly how it can fa