
---
title: "Hague Grant Application"
author: Karen Pauley
type: post
date: 2013-07-24 00:00:00 +0000 UTC
url: "/post/hague_grant_application"
categories:
 - Grants
 - Sign in

---

We have received the following Hague Grant application from diakopter.
Before we vote on this proposal we would like to have a period of community consultation for 10 days. Please leave feedback in the comments or if you prefer send email with your comments to karen at perlfoundation.org.

**Applicant's Name**: diakopter

**Project Title**

Wrap up Perl 5 Interop for Rakudo Perl 6 on MoarVM (2 months)

**Synopsis**

Embed Perl 5 in Rakudo's implementation on one of its two newest backends, MoarVM, in order to enable seamless two-way interop between the Perl 6 environment and the Perl 5 environment (9 weeks).

**Benefits to Perl 6 Development**

Rakudo Perl 6 currently has limited access to Perl 5 modules and CPAN distributions via Blizkost, but only on the Parrot VM.  Tobias Leich is building rakudo-p5/v5 (use v5;), which is a reimplementation of Perl 5 in Perl 6, but that effort will be able to support only a small subset of what's on CPAN. (However, see below on how it could help in the future.)  Bringing all of CPAN to Rakudo via embedding the Perl 5 interpreter would open up entirely new worlds of opportunities and additional Perl 6 (and Perl 5!) adoption and maintenance paths.

**Deliverables & Project Details**

We expect Rakudo to be basically done being ported to the JVM by August, and then also to MoarVM not too many weeks after that.  I can get much of the coding done before Rakudo is bootstrapped on MoarVM, but nearly all of the writing tests in Perl 6 would occur in the third month.  As far as coding goes, I've already sketched out around 40% of the code for the Perl 5 interop.

Long-term, there are many potential Perl 5 Interop deliverables, but to start, it is an extension to MoarVM and Rakudo p6, either distributed separately or optionally built-in (or both), that embeds p5 libperl to enable calling p5 code (interacting with p5 objects) from p6, and vice versa, for the typical use case of using p5 CPAN distributions in p6 code.

**Feature details/schedule:**


# finalize MoarVM's dynamic extension interface protocols & standards (using DynCall - design already pushed to the extops branch). Note: dependency on importing DynCall into MoarVM (from NQP proper).  See http://goo.gl/W5mI0 for the design
# Skeleton extension "mvmp5" that installs the custom op for p5 interop and installs the custom REPRs into MoarVM/6model (for the p5 object wrappers and others)
# Custom p5 REPRs - MVMP5Interpreter, MVMP5Val, MVMP5Method, MVMP5Package
# Custom p5 op - mvmp5::code(string), which takes p5 code and creates an invokable MVMP5Val (wrapping a CV ref)
# Appropriate handling (symbol synchronization) of outermost "use" statements to mvmp5::code (not testable until Rakudo (or NQP) is bootstrapped on MoarVM, because the use case of the feature is compile-time from the HLL's parser)
# Custom MetaObject implementation (mostly in NQP) for MVMP5Val and MVMP5Package objects (mostly for find_method)
# Special handling of p6 type objects as represented in p5, p5 packages in p6
# Special handling of HLL-config-ish things of p5 objects in p6, and likewise special handling of override.pm and TIE interfaces of p6 objects in p5
# Wiring up the two VMs' GCs by an intermediate references table with refcounts for p6 pointers from p5, and *double*-increment refcount of p5 pointers from p6, so p6 can hold a reference while also being able to control exactly when the final reference is decremented
# calling p5 code (invokable MVMP5Vals that wrap CVs) from p6 (including arg flattening & coercion)
# Method calls on p5 objects from p6 (use of MVMP5Method intermediate invokables)
# Method calls on p6 objects from p5 (with the optimization of installing stubs into the stashes, a la Moose class finalization)
# Integration of CallAsOp module to avoid nested p5 runloops (enabling co-recursive cross VM callbacks in a stackless manner)
# p5 wrappers of p6 objects using blessed/tied/autoload SVrefs that all refer to the same dummy SV; indexes to the MVMObjects representing the reference table offsets of each's origin MVMP5Interpreter and referent MVMObject are embedded in another magic on the SV
# Marshaling of value types bidirectionally on args to the VMs (P6int/num MVMObjects & IV/NV, e.g.) cross-invocations
# Coercion of p5 args to mvm ints/nums/strs registers. These are not [yet] listed in implementation order.

The design of the p5 interop scheme has been thoroughly discussed with Larry Wall, Jonathan Worthington, and Nicholas Clark.  A prominent p5p-er offered to help drive any changes needed to p5 to support it, but based on my current design and understanding, no changes will be necessary to support 5.8.3 and later. See http://goo.gl/WLJUQ for one iteration of the working p5interop design doc.

Note: we also considered the JVM while designing this scheme, and nearly all of the design of the MoarVM integration can be directly translated to work on the JVM, even the stackless co-recursive callbacks, since Stefan recently added support for stack saving and resuming for implementation continuations for NQP/rakudo on the JVM.  It would probably be another couple of weeks work to port it to the JVM also.  Note also that this would provide a more complete bridge from p5 to the JVM via p6 (more complete than Larry's prior work that shipped with p5 for a long time and other Java interop projects).

**Project Details**

(also see deliverables.)  Long term, goals for the p5 interop include well- "productized" (though not commercialized!) tools to package p6 modules as p5 distributions (via MoarVM/Rakudo p6 provided as an extension to p5), the converse (p5 extensions installable to the local p5 module trees via the p6 library management tools), programmatic dual (but integrated) packaging of p6 code for deployment/distribution to both p5 and p6 installations from the CPAN system.  The potential exists to use Tobias Leich's p5 grammar & parts of his p5 re-implementation to better facilitate seamless p5 code inline with p6 code as S01 describes, but the modest goals of this project don't include that functionality; this is mainly for loading p5 CPAN modules and scripts into the p6 runtime and accessing their exposed interaction points.  The current design also leaves room for p5 ithreads (but not forks.pm) to be supported in the future, but isn't entirely in scope for this proposal.

**Project Schedule**

I expect the p5 interop to take up to 9 weeks, commencing anytime.

**Report Schedule**

Progress reports will be mentioned weekly to the #perl6 and #moarvm channels on Freenode IRC, and as such will be published in the logs at irc.perl6.org; certain major milestones of feature availability will get blog posts on the planet Perl 6 syndication feed.

**Public Repository**

MoarVM's git repository is at github.com/MoarVM, and changelogs for these contributions will mostly exist in the commit logs, as there likely won't be "releases" of MoarVM separately from NQP itself (as MoarVM is fundamentally an NQP backend).

**Grant Deliverables ownership/copyright and License Information**

MoarVM is published under the Artistic License 2.0, as are these sub-projects. The Perl Foundation will retain copyright of these initial contributions, but not the software as a whole, because patches and fixes to the software won't usually be able to have their copyright assigned to TPF.

**Bio**

I'm a co-founder of MoarVM, the Perl6-specialized virtual machine.  I designed and implemented (or will implement) the strings/IO subsystems, HLL locks and atomic operations, the bytecode validator and disassembler, the QAST->MAST compiler, both VM-internal and REPR-level hashes, the regex compiler, coordination of ending threads and portions of the GC parallelization and work-stealing system, the native and HLL extension system/interface, the threadsafe (and lock-free!) hash and array REPRs and their auto-upgrading from their speedier single-threaded editions, the NFG implementation, the Unicode Character Database character properties and case changes compilation/runtime, and finally the Perl 5 interop scheme.

**Amount Requested**

USD$10k for 2 months at $5k/month.





