
---
title: "Grant Proposal: Rakudo Perl 6 performance analysis tooling"
author: Coke
type: post
date: 2017-09-19 00:00:00 +0000 UTC
url: "/post/grant_proposal_rakudo_perl_6_p"
categories:
 - Grants
 - Sign in

---

<p>The Grants Committee has received the following grant proposal for the September/October round. Before the Committee members vote, we would like to solicit feedback from the Perl community on the proposal.</p>
<p>Review the proposal below and please comment here by September 26th, 2017. The Committee members will start the voting process following that and the conclusion will be announced the first week of October.</p>
<h2>Rakudo Perl 6 performance analysis tooling</h2>
<h2>Name</h2>
<p class="Standard">Timo Paulssen</p>
<h2>Amount Requested:</h2>
<p class="Standard">7,500 USD</p>
<h2>Synopsis</h2>
<p class="Standard">Rakudo on MoarVM already comes with two kinds of profilers and multiple kinds of internal logs. There's also a way to get coverage analysis from any given program or set of programs.</p>
<p class="Standard">The focus on internals that these tools have in common can make it quite daunting for someone not involved in core development to use these tools and to make sense of the data gathered.</p>
<p class="Standard">I will create a new GUI (browser-based) for the instrumented profiler (which takes timings of functions and counts allocations) and the heap snapshot profiler (which records what objects are alive and how they are interconnected) and make invoking the profiling modes simpler.</p>

<style type="text/css"><!--
p.p1 {margin: 18.0px 0.0px 6.0px 0.0px; font: 16.0px Arial}
p.p2 {margin: 0.0px 0.0px 0.0px 0.0px; text-indent: 36.0px; font: 11.0px Arial}
p.p3 {margin: 0.0px 0.0px 0.0px 0.0px; font: 11.0px Arial}
p.p4 {margin: 0.0px 0.0px 0.0px 0.0px; font: 11.0px Helvetica; min-height: 13.0px}
li.li3 {margin: 0.0px 0.0px 0.0px 0.0px; font: 11.0px Arial}
span.s1 {font: 11.0px 'Courier New'}
span.Apple-tab-span {white-space:pre}
ul.ul1 {list-style-type: disc}
--></style>
<p class="p1"></p>
<p class="p3">Even though there already is a browser-based UI for the instrumented profiler, it quickly becomes unusable as program sizes increase.</p>
<p class="p4"></p>
<p class="p3">Another shortcoming of the instrumented profiler is that it doesn't work when multiple threads are involved in the user's program.</p>
<p class="p1">Benefits to the Perl Community</p>
<p class="p3">Rewriting the Profiler UI for the instrumented profiler to understand multiple threads and our multithreaded primitives like <span class="s1">react</span>/<span class="s1">whenever</span> will allow both our users and our core developers to figure out the performance of multithreaded programs more easily and precisely than possible so far. On top of that, basing the UI on the SQL output from our profiler rather than the JSON output we had before will improve the responsiveness and memory use of the profiler drastically, even making it possible to get a profile of compiling the Rakudo Perl 6 Core Setting, a single Perl 6 source file of almost 60k lines.</p>
<p class="p4"></p>
<p class="p3">The Heap Snapshot Analyzer allows users to figure out what objects take up the most memory, how the distribution of object counts and sizes changes over the runtime of a program, and why any given object is being kept from being garbage collected. Giving our users a UI for the Heap Snapshot Analyzer will allow them to more easily navigate between objects, get graphs of queries, and to have the information more organized than the current shell-like interactive UI.</p>
<p class="p4"></p>
<p class="p3">Another task I want to complete as part of this grant is to have much more end-user oriented documentation for how to use the profilers effectively, including a few example programs.</p>
<p class="p1">Deliverables and Inchstones</p>
<ul class="ul1">
<li class="li3">A blog with progress reports.</li>
</ul>
<ul class="ul1">
<li class="li3">A web frontend for the heap snapshot analyzer</li>
<ul class="ul1">
<li class="li3">Refactor how the analyzer gives data to the shell</li>
<li class="li3">Draft a concept for how the user will interact with the analyzer</li>
<li class="li3">UI for Per-Snapshot Summary: total heap size, total object count, â€¦</li>
<li class="li3">UI for Top Lists for objects sorted by count or memory usage</li>
<li class="li3">UI for Details of individual objects: size, pointers to other objects</li>
<li class="li3">UI for the shortest path that keeps an object alive</li>
<li class="li3">UI for Across-Snapshot comparisons: object counts over time, â€¦</li>
<li class="li3">UI for Heap Exploration: Find all objects of a specific type, â€¦</li>
<li class="li3">Functionality for finding paths from one object to all roots that reach it.</li>
<li class="li3">UI for whole parts of the network, like multiple paths to a single object.</li>
<li class="li3">If an instrumented profile is also loaded
<ul>
<li class="li3">Links from types to routines allocating the type</li>
<li class="li3"><span>Links from frames (closures for example) to the call graph</span></li>
<li class="li3"><span>A new web frontend for the instrumented profiler</span></li>
</ul>
</li>
<li class="li3">Communication between the web frontend and a backend for data queries</li>
<li class="li3">Overview page with at-a-glance metrics for quick feedback for code changes</li>
<li class="li3">Call Graph explorer</li>
<ul class="ul1">
<li class="li3">Robust forwards/backwards navigation</li>
<li class="li3">Icicle Graph with acceptable performance</li>
<li class="li3">Search function for Call Graph Explorer</li>
</ul>
<li class="li3">Routines List</li>
<ul class="ul1">
<li class="li3">Sortable by every sensible column</li>
<li class="li3">Groupable by filename</li>
<li class="li3">Expand a routine to see its callers or callees</li>
<li class="li3">â€¦ with links directly to the call graph explorer</li>
</ul>
<li class="li3">Allocation explorer</li>
<ul class="ul1">
<li class="li3">Show what routines allocate a given class</li>
<li class="li3">Expand a routine to see which of its callers is responsible for what amount of allocations</li>
<li class="li3">Link to the heap snapshot explorer for any given type</li>
</ul>
<li class="li3">OSR and Deopt explorer</li>
<ul class="ul1">
<li class="li3">Show routines that have been OSR'd (On-Stack-Replaced by the optimizer)</li>
<li class="li3">Again, allowing to expand routines to see their callers</li>
<li class="li3">Expose a lot more information about deopts - which types are involved, for example</li>
<li class="li3">Explain in clear terms what OSR and Deopt mean</li>
</ul>
<li class="li3">â—‹<span class="Apple-tab-span"> </span>GC explorer</li>
<ul class="ul1">
<li class="li3">Figure out a good way to present information from multi-threaded programs here</li>
<li class="li3">Expose time-of-collection info to profiler</li>
<li class="li3">Show the time and duration of collections, how much data was kept, promoted to the old generation, or deleted, and whether the collection was a minor or major one</li>
<li class="li3">Filterable to show major collections only</li>
<li class="li3">Compactly graph collection times and kept/promoted/deleted amounts</li>
<li class="li3">User-facing documentation on using the profiling facilities and interpreting the data</li>
<li class="li3">Build a bunch of examples that show different performance issues and describe how to figure them out</li>
</ul>
<li class="li3">Using an unoptimized built-in</li>
<li class="li3">Accidentally keeping objects alive</li>
<li class="li3">Code using native ints but inadvertently causing lots of boxing</li>
<li class="li3">Using mixin on a hot path causing lots of deopt</li>
</ul>
</ul>
<p class="p1">Project details and a proposed Schedule</p>
<p class="p3">I can start working on this grant immediately. Over the last weeks I've implemented a new data format for the Heap Snapshot Profiler itself, and started improving the Heap Analyzer's Shell and Model in preparation for a new Frontend.</p>
<p class="p3"></p>
<p class="p3">This grant is likely to take at least 160 hours, which at 40 euros per hour is about 7,500 dollars. Therefore, I'd like to request 7,500 USD for this work. Any left-over hours at the end of the grant can easily be spent improving the whole result, though typically a project will take more time than estimated, not less.</p>
<p class="p3"></p>
<p class="p3">I'm planning to spend at least 15 hours per week on this, with a few weeks of break near the end of January in order to move apartments.</p>
<p class="p1">Completeness Criteria</p>
<p class="p3">There will be changes to MoarVM to expose more information to profiler output. In the case of the Instrumenting Profiler, there will be minor changes to NQP to transfer the extra information into the output formats.</p>
<p class="p3">Most of the work will live in a new repository, with a sizable chunk living in the existing Heap Analyzer in jnthn's github repository.</p>
<p class="p3">At the end of the project there will be a Perl 6 program that you can use to run your code with profilers enabled and then introspect all the information. It will be interacted with using a web browser.</p>
<p class="p1">Bio</p>
<p class="p3">I came to the Perl 6 project in March 2012 after having made a few minor contributions to PyPy. Initially I worked on error messages, then Rakudo's Optimizer, Rakudo's Grammar, and later also contributed bits and pieces to Parrot. When the Rakudo JVM Port began, I contributed little pieces to that, and finally when MoarVM became public I concentrated my efforts on that. I also participate in discussion and support on the IRC and on the perl6-users mailing list.</p>
<p class="p3"></p>
<p class="p3">In total, I've added 779 commits to Moarvm, 318 to NQP, and 359 to Rakudo, excluding commits in branches that have not been merged. Those commits range from single-line patches to bigger chunks of work.</p>
<p class="p3"></p>
<p class="p3">I've also contributed pieces to the heap snapshot analyzer, helped fellow contributor MasterDuke17 implement an SQL output mode for the profiler, and in general spent a sizable portion of my time and energy on optimization-related work.</p>
<p class="p3"></p>
<p class="p3">I also just merged a new output format and API for the heap snapshot profiler.</p>

