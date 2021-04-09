
---
title: "Perl 6 Hague Grant Application"
author: Karen Pauley
type: post
date: 2015-04-19 00:00:00 +0000 UTC
url: "/post/perl_6_hague_grant_application"
categories:
 - Grants
 - Sign in

---

We have received the following Hague Grant application from Bart Wiegmans.  Before we vote on this proposal we would like to have a period of community consultation for 10 days. Please leave feedback in the comments or if you prefer send email with your comments to karen at perlfoundation.org.

**Name**: Bart Wiegmans

**Project Title**:  Advancing the MoarVM JIT

**Synopsis:**

Implement an advanced code generation algorithm for the MoarVM JIT compiler, leading to more efficient machine code for JIT compiled frames.

**Benefits to Perl 6 Development:**

This project will enable the MoarVM JIT compiler to generate smaller and more efficient code, producing greater performance and less memory traffic during execution. This will help to make Perl 6 more competitive with other languages. (Speed is generally regarded as a feature).

As a secondary benefit this will decouple the JIT intermediate representation from the MoarVM bytecode and the x64 machine code making it easier for developers to extend or to port the JIT to architectures other than x64, such as ARM.

As an example of the potential speedup I've created the following example demonstrating a 5x speedup on tight numeric code. Although that example is highly artificial it does demonstrate the limits of the current JIT rather well.  Note also that the original hot loop uses 22 instructions whereas the new hot loop uses only 7 instructions.

**Deliverables:**

These deliverables are not ordered chronologically.


# An implementation of the code generation algorithm described below, including instructio  selection tables for the x64 architecture 
# A runtime representation of machine operations ('expressions') that will form the input to the code generator and is suitable for targeting to different architectures
# A patched version of DynASM capable of addressing the extended registers of x64
# Conversion of large parts of the current JIT to the new algorithm
# An extension API for REPR ops to insert (inline) expressions into the JIT in place of some operations
# A set of automated tests known to trigger JIT compilation to either the NQP or Rakudo Perl 6 test suite.
# Reports and documentation explaining the JIT and it's API 


**Project Details:**

Since September 2014 MoarVM implements a JIT compiler for optimised frames. This JIT compiler has proven to be reliable after it became the default option in October 2014, but it does not generate very efficient machine code. This is because it is very simple (simplistic perhaps) and generates code by pasting together preformed blocks of machine code.

Because these machine code fragments are necessarily independent, they must store and load values from memory every time they execute, even if these values are never read afterwards or if they are only temporarily useful. This results in large code size and heavy memory traffic. Moreover, because each part is independent, even very simple optimizations across fragment boundaries are impossible or very hard. Thus, the potential performance increase gained by JIT compilation is limited by the current design.

In this project I propose to implement a more advanced code generation algorithm based on the paper by Aho et al, which I will not describe in detail here. What is relevant for this proposal is that I intend to embed it into the existing JIT compiler as an additional node type. This means that the current JIT compiler can be converted to use this new algorithm in small pieces. This also means that the new algorithm must be written to use the DynASM bytecode generation library. This in turn means that DynASM must be extended to support register addressing on x64 (it currently only does so on x86).

Although this is a large project it, it can be divided into 'inchstones' relatively cleanly, each of which can be divided further:


* Implement a 'proof of concept' code generator using DynASM with limited register addressing. This is intended to find any limits to DynASM (or my understanding) beyond the register addressing limitation already discussed.
* Eliminate the register addressing limitation. In practice this comes down to conditionally adding register extension prefix bytes before the actual instructions. Last year I discussed briefly how this should be done with Mike Pall, the author of DynASM and luajit.
* Implement a machine-level expression graph data type, suitable for expressing MoarVM- evel concepts close to machine-level operations. 
* Actually implement the algorithm, including components like register selection, instruction selection, and value spills.
* Convert current JIT compiler fragments to the new expression trees. Areas of specific interest include branching, C calls, VM routine calls, GC barriers, deoptimization and exception handlers. Each of these is a goal onto itself. The current JIT fragments can function as templates for the expression graph.
* Implement an extension API for REPR's (object representations) to inline operations into JIT expressions.


Automated tests will be developed during development of features and collected over time. Given enough time and less-than-expected difficulties in implementation, I have many more ambitions, but I think this should provide enough work for now.

**Project Schedule:**

I estimate that this project will take 10 weeks of full-time work, excluding time spent studying in preparation. Around two to three weeks will be spent on creating a proof of concept code generator and extending DynASM for x64 register addressing. If this is successful I think around 2 weeks will be spent on implementing this for MoarVM, of which one week will be spent on the algorithm and another week on instruction and register tables. After this I'll be able to demonstrate a simple JIT-compiled frame using the new algorithm.

The next three weeks will be spent converting more complex parts until a large part of the current JIT has been converted. Finally in the last two weeks I will implement the extension API as well as implement one or more of such extensions, for example inlined array access.

I can begin working in the week starting Sunday June 14 2015, meaning I intend to finish Saturday August 22. The halfway milestone should then be reached Saturday July 18.

**Report Schedule:**

I will report on my work using my blog as I have done last year during GSoC. Blog entries will appear at least every two weeks and preferably more often. This blog is syndicated on the "pl6anet.org":http://pl6anet.org/ news aggegration site. Furthermore, I will report often to the #perl6 and #moarvm channel of freenode to discuss my progress. I'm also willing to provide mailing-list announcements (e.g. on perl6-announce) in case of milestones.

**Public Repository:**

The final code will be hosted in the MoarVM public git repository at github.com/MoarVM/MoarVM as well as in MoarVM's fork of DynASM. Changes to DynASM will also be offered to the luajit project from which DynASM has been extracted. Initially work will be done in a separate branch so as to not disturb 'regular' users. Any proof-of-concept code will be hosted at my personal github account.

**Grant Deliverables ownership/copyright and License Information:**

MoarVM is licensed under the Artistic 2.0 license. DynASM is licensed under MIT. Copyright on MoarVM belongs to Jonathan Worthington and others (according to the current LICENSE file). Copyright to DynASM belongs to Mike Pall. I have naturally no authority to change the copyright on these projects. The Perl Foundation may of course have copyright on the patches applied for this project (although I don't know if that works out legally).

**Bio:**

I was the author of the current MoarVM JIT compiler during the 2014 edition of GSoC. Thus, I know it very well. I consider myself to have been a member of the Perl 6 community since 2012, when I wrote the mod_parrot module for apache httpd during the 2012 edition of GSoC. I wrote my first interpreter around 10 years ago. 

In real life I've been studying physics, biology, and currently environmental science. I have, in general, lots of time available in the summer. Lastly, I obsess about speed and MoarVM / Rakudo Perl 6 not being as fast as they can bothers me.

**Amount Requested:**

I request $10.000 for 2,5 months of work. I'd suggest payment of $5000 after the first halfway milestone and $5000 after the project is finished. 

**Suggestions for Grant Manager:**

Jonathan Worthington (jnthn), who also mentored me during the 2014 edition of Google Summer of Code. He is the lead developer on MoarVM and a large contributor to the Rakudo Perl 6 compiler, on which he has given many talks at Perl conferences and elsewhere.


