
---
title: "Grant Proposal: Improve Perl 6 Networking Support"
author: Coke
type: post
date: 2018-12-16 00:00:00 +0000 UTC
url: "/post/grant_proposal_improve_perl_6_"
categories:
 - Grants
 - Sign in

---

The Grants Committee has received the following grant proposal for the Nov/Dec round.
Before the Committee members vote, we would like to solicit feedback from the Perl
community on the proposal.

Review the proposal below and please comment here by December 22nd, 2018.
The Committee members will start the voting process following that.

# Improve Perl 6 Networking Support

- Name:

    Ben Davies

- Amount Requested:

    USD $7200

## Synopsis

Perl 6 lacks many features when it comes to networking in comparison to other
languages. The aim of this is to give Perl 6 networking support equivalent to
that of other languages, such as Python.

## Benefits to the Perl Community

This will allow users to do more advanced networking without having to depend
on NativeCall and will make the API for IO::Socket::Async and IO::Socket::INET
more consistent. IO::Socket::UNIX will also be introduced to provide support
for IPC with UNIX sockets.

Some of the reasons I am proposing this are that:

\- SO\_OOBINLINE is required for TELNET protocol

\- Even after adding support for getsockopt/setsockopt, there's no way to use options that need to be set before connect/bind is called on both synchronous and asynchronous sockets, like SO\_REUSEADDR

\- Raw sockets are currently unusable without using NativeCall. This makes it impossible to write anything that depends on raw sockets, like ICMP, purely in Perl 6

\- IPC is possible to do with TCP/UDP sockets, but UNIX sockets would allow users to share file descriptors between processes and would bypass the networking stack for faster IPC

## Deliverables

This grant will allow Perl 6 to support:

\- getting the file descriptors of IO::Socket::Async instances

\- setsockopt/getsockopt on networking sockets

\- setting options on IO::Socket::INET and IO::Socket::Async instances before bind/connect is called

\- creating raw and UDP IO::Socket::INET sockets across all OSes

\- sending/receiving data over raw and UDP IO::Socket::INET sockets

\- creating UNIX sockets and being able to send and receive data, get/set options, etc. over them

Some of these features are breaking changes and will need to wait until 6.e.

## Project Details

I have already started working on implementing .native-descriptor for
IO::Socket::Async and getsockopt/setsockopt support.

## Inch-stones

\- Implement .native-descriptor for getting the file descriptors of IO::Socket::Async sockets

\- Implement getsockopts op to get getsockopt/setsockopt option constants

\- Implement getsockopt/setsockopt ops

\- Implement asyncsocket op

\- Refactor IO::Socket::Async to use nqp::asyncsocket

\- Add a .new method to IO::Socket::Async

\- Refactor IO::Socket::INET to allow creating sockets without having bind/connect called immediately

\- Refactor IO::Socket::Async's UDP methods to allow using getsockopt/setsockopt on UDP sockets before they are bound

\- Make initializing raw and UDP IO::Socket::INET sockets work across all OSes

\- Implement sendto/recvfrom ops

\- Implement unixsocket, unixopen, unixlisten, unixconnect ops

\- Create IO::Socket::UNIX

\- Implement unixsendfd/unixreadfd ops and support sending/receiving file descriptors over unix sockets

## Project Schedule

I imagine this would take a few months to complete. I've already begun work on
implementing .native-descriptor and getsockopt/setsockopt support, so I'm ready
to start as soon as possible. I will be working around 6 hours a day on this.

3 months \* 4 weeks \* 5 days \* 6 hours \* $20 = $7200

## Completeness Criteria

Whenever my code gets merged into MoarVM/nqp/rakudo's master branch.

## Bio

I am a 22 year old self-taught programmer with 6 years experience. I know C,
Perl 6, node.js, as well as some Python, Java, Rust, and Go. I've been working
with networking-related code for around 5 years, particularly with TELNET and
WebSockets. I've been cont