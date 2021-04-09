
---
title: "Grant Proposal: Introduction to Application Development with Perl 6"
author: Coke
type: post
date: 2018-06-11 00:00:00 +0000 UTC
url: "/post/grant_proposal_introduction_to"
categories:
 - Grants

---

The Grants Committee has received the following grant proposal for the May/June round.
Before the Committee members vote, we would like to solicit feedback from the Perl community on the proposal.

Review the proposal below and please comment here by June 16th, 2018.
The Committee members will start the voting process following that and the conclusion will be announced the
last week of June.

# Introduction to Application Development with Perl 6

- Name:

    Patrick Spek

- Amount Requested:

    USD $2,500

## Synopsis

A book about getting started with application development in Perl 6. It will
cover installation of Perl 6, the basics of the language, how to use the
`perl6` binary to run programs, how to create a terminal application and a GUI
application using `GTK::Simple`.



## Benefits to the Perl Community

The book, if approved, will be released in the public domain, which allows use
for everyone interested. This can help new programmers to easily gain access to
information to get started to use Perl 6 for real world applications. People
who are already familiar with Perl 6 can freely distribute the contents of the
book to other people to promote Perl 6.

## Deliverables

A book titled "Introduction to Application Development with Perl 6", released
under the Creative Commons license.

## Project Details

The book will cover introduction to terminal application and GUI application
development. Thus far, I have the following 20 chapters planned, with a short
description of what they will contain:

- **Foreword**: About the author, special thanks for contributors, what will this book teach you, what prior knowledge is required, a word about Unicode in Perl 6.
- **Getting ready to use Perl 6**: Installing Perl 6 and setting up an editor
- **Hello world!**: Run a minimal "Hello World" using `perl6 -e`, once written in a single file to be ran with `perl6 file.pl` and once in a generic Perl 6 module structure. Using App::Assixt for the last part.
- **An overview of Perl 6 constructs**: Basic operators, conditional blocks, loops, differences between sigils.
- **The `MAIN` subroutine**: What is the `MAIN` subroutine, and why should you use it?
- **Adding some plain old documentation**: Using Perl 6 POD to document your code, along with common formatting constructs from Perl 6 POD.
- **Generating usage info with `USAGE`**: Applying POD knowledge to `MAIN` and overriding the default `USAGE` to generate custom messages.
- **Accepting flags and options**: How to make `MAIN` deal with command line flags and options.
- **Variable types**: What are typed variables, why would you use them, some common types.
- **Splitting off into a module**: Moving parts of your code into a separate module and using them.
- **Using modules**: How to use modules, specifically third party modules.
- **Classes and objects**: How to make classes and instantiate them.
- **Application configuration**: How to deal with user configuration of your application? This will use the `Config` module.
- **Dealing with user input at run time**: How to prompt the user for input and use it in your application.
- **Pretty output**: How to format and color output of your application.  This will use `Terminal::ANSIColor`.
- **Reading and writing files**: How to read files and write files.
- **Making a GUI with GTK**: Using `GTK::Simple` to create a GUI.
- **Handling GUI events**: Continuing on `GTK::Simple`, how to handle events and change elements in the GUI based on these events.
- **Running as a service**: How to run a Perl 6 application as a system service. This will target systemd, OpenRC and possibly BSD's RC as service managers.
- **Addendum**
    - List of Unicode operators

The book itself will be written using LaTeX, though this can change if it is
found to be unfeasable (lack of Unicode support, for instance). Whatever the
source format will be