
---
title: "Rakudo Perl 6 performance analysis tooling - Grant Report #3"
author: Alan Haggai Alavi
type: post
date: 2018-07-30 00:00:00 +0000 UTC
url: "/post/rakudo_perl_6_performance_anal_2"
categories:
 - Grants
 - Perl 6 Development

---

<p>Timo has developed a shiny new UI for the tool. It displays thread-level garbage collection details.</p>
<p>Read more at: <a href="https://wakelift.de/2018/07/26/wow-check-out-this-garbage/">Wow, check out this garbage</a></p>
<hr />
<p></p>
<h1 class="post-title">Wow, check out this garbage</h1>
<div class="kg-card-markdown">
<p>Hello everyone! It's been more than a month since the last report, but I've been able to put hours in and get code out. And now I'll show you what's come out of the last weeks.</p>
<h2 id="thegarbagecollector">The Garbage Collector</h2>
<p>One important aspect of dynamic languages is that memory is managed for the user. You don't usually see <code>malloc</code> and <code>free</code> in your perl code. But objects are constantly being created, often "in the background". Asking the user to take care of freeing up space again for objects that may have been created and become obsolete in the same line of code sounds like a good recipe for user dissatisfaction.</p>
<p>To take this potential full-time-job off the user's hands, MoarVM employs Garbage Collection, specifically a scheme based on "reachability". Whenever space needs to be freed up, the process begins. Starting from a set of objects that MoarVM happens to know for sure are currently needed, references from one object to another are followed until everything "reachable" has been reached and "marked" for being kept. Afterwards, everything that has been created but not "marked" will be removed in a step commonly called "sweeping". Garbage Collectors following this scheme are called "Mark &amp; Sweep Garbage Collectors". Find a section about this on <a href="https://en.wikipedia.org/wiki/Tracing_garbage_collection#Na%C3%AFve_mark-and-sweep">the "Tracing Garbage Collection" wikipedia article</a>, though MoarVM's garbage collector has a few additional tweaks for performance reasons.</p>
<p>Naturally, going through a large amount of objects consumes time. Seeing what the program spends its time doing is an important part of making it perform better, of course. That's why the MoarVM profiler records data about the Garbage Collector's activities.</p>
<h2 id="gcintheprofiler">GC in the profiler</h2>
<p>The old HTML profiler frontend already showed you how many GC runs happened during your program's run. For each time the GC ran you'll see how long the run took, whether it went through the nursery only (a "minor" collection) or everything on the heap (a "major" collection, called "full" in the tools). I recently added a column displaying when the run started in milliseconds from the start of recording. However, the old profiler frontend didn't do anything with information recorded by threads other than the main thread. Here's a screenshot from the old frontend:</p>
<p><img src="https://wakelift.de/content/images/2018/07/moar_profiler_gc_screenshot-1.png" alt="moar_profiler_gc_screenshot-1" style="width: 100%;" /></p>
<p>I'm not quite sure why it says it cleared about 4 gigabytes of data during the first two runs, but that'll go on the pile of things to check out later.</p>
<p>For now, I can use this to explain a few things before I go on to show the first version of the new interface and what it does with multithreaded programs.</p>
<p>The profile was taken from a hypered (i.e. automatically multithreaded) implementation of the fannkuch benchmark. The rightmost column shows how much data was copied over from the nursery to the next nursery, how much was taken into the "old generation" (either called "old" or "gen2"), and how much was freed.</p>
<p>There's also a count of how many references there were from objects in the old generation to objects in the nursery, the "gen2 roots". This happens when you keep adding objects to a list, for example. At some point the list becomes "old" and fresh objects that are i