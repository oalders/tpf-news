
---
title: "Grant Proposal: Perl-Based CloudForFree.org Platform"
author: Coke
type: post
date: 2018-02-13 00:00:00 +0000 UTC
url: "/post/grant_proposal_perl-based_clou"
categories:
 - Grants
 - Sign in

---

The Grants Committee has received the following grant proposal for the January/February round.
Before the Committee members vote, we would like to solicit feedback from the Perl community on the proposal.

Review the proposal below and please comment here by February 20th, 2018.
The Committee members will start the voting process following that and the conclusion will be announced the
last week of February.

# Perl-Based CloudForFree.org Platform

- Name:

    John Napiorkowski

- Amount Requested:

    USD $1,500

## Synopsis

Complete development of the partially-operational [http://www.cloudforfree.org](http://www.cloudforfree.org) platform.


## Benefits to the Perl Community

Few people understand what The Cloud is, and fewer still have access.

The Perl community will benefit by having wide-spread free access to a Perl-based Cloud computing platform, which will showcase Perl as both modern and competitive against other Cloud technologies.

## Deliverables

The CloudForFree.org platform itself will be updated with new features which will automatically become available to all users.

## Project Details

The CloudForFree.org platform is a Perl-based Cloud computing software system created from scratch.

Numerous technologies are utilized in the current version of CloudForFree.org, including but not limited to:

- Perl
- Catalyst MVC
- ShinyCMS
- RPerl
- GCC C++
- Javascript
- CGI::Ajax
- Ace Editor

I have reviewed the CloudForFree.org source code and started initial planning discussions with the founder, Will Braswell.

## Inch-stones

Phase 1, LAMP Installer Script Section 51 

- Review & Comprehend LAMP\_installer.sh Script 
- Upgrade LAMP\_installer.sh Section 51 "PERL CLOUDFORFREE" To Function Properly 
- Demonstrate Section 51 Installing CloudForFree Software In Ubuntu 16.04 
- Upgrade CFF Server To RPerl v3.2 

Phase 2, IDE Code Editor 

- Create Linux User Account For Existing & New Users 
- New Copies Of Learning RPerl Exercises 
- Save All Files In Linux User Home Directories 
- Button For "Save File" Function 
- Save & Quit Functions For Ace Editor vi & emacs Keyboard Commands 
    - :w :wq ZZ etc. 
    - Ctrl-x Ctrl-s Ctrl-x s Ctrl-x Ctrl-w etc. 
- Multiple Open Files Via Tab Widget 

Phase 3, Open GitHub Issues 

- Review & Comprehend All Open GitHub Issues 
github.com/wbraswell/cloudforfree.org... 
- Collaborate With Client Personnel To Plan Solutions For Each Open Issue 
- Implement Solutions For Each Open Issue 

Phase 4, Terminal Emulation 

- Full xterm Terminal Emulation For Command-Line Job Execution 
- Includes All VT100 & VT102 & VT220 Features 
- Resize Number Of Rows & Columns By Window Resize 
- Resize Font Size By User Configuration 
- Backspace, Tab, Other Special Command Characters 
- Arbitrary Placement Of Characters At Any Row & Column 
- Color Characters & Background 
- Curses & Other Menu Support 
- Full Window (F10) & Full Screen (F11) 

Phase 5, Graphical Output & Mouse Input 

- Full X-Windows Graphics Output Using Xpra HTML5 Client 
- Generate Output Using SDL From Perl & C 
- Mouse Left-Click, Right-Click, Drag, Scroll 
- Resize Number Of Available X & Y Pixels By Window Resize 
- Full Window (F10) & Full Screen (F11) 

Phase 6, GitHub Repositories Integration 

- Import & Setup User's GitHub Keys Via Linux User Account (From Phase 2) 
- List Of All GitHub Repos With User As Owner Or Collaborator 
- Admin Mode, Display All Repos 
- Allow User To Enter Any Other Readable GitHub Repo URL 
ex. github.com/wbraswell/physicsperl 
- Buttons For Basic User Git Commands 
    - clone 
    - add 
    - commit 
    - push 
    - pull 
    - checkout 
    - status 
- Do Not Duplicate Any GitHub Web Functionality 

Phase 7, Job Queue 

- Job Scheduler & Monitor Using OAR (1st Choice), Or HTCondor Or Slurm (2nd Choices) 
- Manage User's Jobs Via Linux User Account (From Phase 2) 
- List Of 