
---
title: "Rakudo Perl 6 performance analysis tooling - Grant Report #4"
author: Alan Haggai Alavi
type: post
date: 2018-09-16 00:00:00 +0000 UTC
url: "/post/rakudo_perl_6_performance_anal_3"
categories:
 - Grants
 - Perl 6 Development

---

<p>The first public release! Code is now hosted in <a href="https://github.com/timo/moarperf/">GitHub</a>. Please see the <a href="https://wakelift.de/2018/08/15/the-first-public-release/#wherecanigetit">instructions on how to install and run</a>.</p>
<p>The release features a renewed "Routines" tab. Please read Timo's blog post to know how it compares to the previous profiler: <a href="https://wakelift.de/2018/08/15/the-first-public-release/">The first public release!</a>.</p>
<hr />
<p></p>
<h1 class="post-title">The first public release!</h1>
<div class="kg-card-markdown">
<p>Hello esteemed readers, and thank you for checking in on my progress. Not a full month ago I showed off the GC tab in <a href="http://wakelift.de/2018/07/26/wow-check-out-this-garbage/">the previous post, titled "Wow check out this garbage"</a>. Near the end of that post I wrote this:</p>
<blockquote>
<p>I had already started the Routine Overview tab, but it's not quite ready to be shown off. I hope it'll be pretty by the time my next report comes out.</p>
</blockquote>
<p>Well, turns out I got a whole lot of work done since then. Not only is the Routine tab pretty, but there's also a Call Graph explorer beside it. This post has the details, and at the end I'll give you the link to the github repository where you can get the code and try it out for yourself!</p>
<h1 id="routinetab">Routine Tab</h1>
<h2 id="routineoverview">Routine Overview</h2>
<p>Now, the simplest part of the Routine tab is its overview function. Every routine that has been called while the profiler was recording shows up here. Here's a screenshot so you can understand what I'm referring to:</p>
<p><img src="https://wakelift.de/content/images/2018/08/Selection_036.png" alt="Selection_036" style="width: 100%;" /></p>
<p>For comparison, here's the old profiler's Routine Tab</p>
<p><img src="https://wakelift.de/content/images/2018/08/Selection_037.png" alt="Selection_037" style="width: 100%;" /></p>
<p>There's actually a lot of differences. Going from the old to the new, the <code>Interp / Spesh / Jit</code> column has disappeared, and with it the OSR badges. Also, the table headers are no longer clickable (to change the sorting order) and there is no filter search field. Both of these features will make a come-back in the new profiler's UI as well, though!</p>
<p>There are also additions, though: There's now a column labelled "Sites", some of the filename + line texts are clickable (they take you directly to github, though opening the files locally in an editor is a feature on my wish-list), and there's a column of mysterious buttons. On top of that, you can now see not only the <code>exclusive / inclusive</code> times for each routine, but also how much time that is when divided by the number of <code>entries</code>.</p>
<p>I wonder what happens when I click one of these!</p>
<h2 id="expanding">Expanding</h2>
<p><img src="https://wakelift.de/content/images/2018/08/Selection_038.png" alt="Selection_038" style="width: 100%;" /></p>
<p>Neat, clicking the button expands an extra section below the routine. It has three tabs: Callees, Paths, and Allocations. Let's go through them one by one.</p>
<h2 id="callees">Callees</h2>
<p>Listed here are all routines that got called by the parent routine (in this case <code>ACCEPTS</code> from the <code>Regex</code> source file). They are ordered by inclusive time, as opposed to the outer list which is ordered by exclusive time. <sup class="footnote-ref"><a href="#fn1" id="fnref1">[1]</a></sup></p>
<p>Since there is now a parent/child relationship between the routines, there's also the number of <code>entries per entry</code> in the entries column. That's simply the entries of the child divided by the entries of the parent. This number can tell you how often another routine is mentioned, or what the probability for the child being called by the parent is.</p>
<h2 id="pathsandwhatsthisaboutsites">Paths, and what's this about "Sites"?</h2>
<p><img src="https://wakelift.de/content/images/2018/08/Selection_039.png" alt="Selection_039" style="width: 100%;" /></p>
<p>The next tab you can open in the expanded view is called "Paths". Here you can see a vaguely tree-shaped table with four rows. These rows actually correspond to the <code>Sites</code> column that I have not explained yet. That's because the number of <code>Sites</code> corresponds directly to the number of rows in this table, or the number of leafs in the tree it displays.</p>
<p>The same <code>Routine</code> can behave in different ways depending on where it was called from. Normally, a difference in arguments passed is the main cause of different behaviour, but it's very likely that each unique location in the program will call the <code>Routine</code> the same way every time. Such a "location" is sometimes called a "Call Site", i.e. the site where a call resides. A <code>Site</code> in the profiler's nomenclature refers to one specific path from the outermost routine of the profile to the given <code>Routine</code>. In the screenshot above, all paths to <code>ACCEPTS</code> go through <code>to-json</code> either once, twice, or three times. And every path goes through <code>str-escape</code>.</p>
<p>The names in the table/tree (trable?) are all clickable. I can tell you right now, that they bring you right over to the <code>Call Graph</code>. More on that later, though.</p>
<h2 id="allocations">Allocations</h2>
<p><img src="https://wakelift.de/content/images/2018/08/Selection_040.png" alt="Selection_040" style="width: 100%;" /></p>
<p>This is a small one, at least for <code>ACCEPTS</code>. It has one row per type, and splits the number of objects created for each type into <code>before spesh</code> and <code>after spesh/jit</code>. Coincidentally, the <code>ACCEPTS</code> method is already a good example for having the split: <code>BOOTHash</code> is just a lower-level hash class used in some internals. Notably, <code>BOOTHash</code> is used to pass named arguments to methods. Of course, many method invocations don't actually pass named arguments at all, and many methods don't care about named arguments either. Thankfully, <code>spesh</code> is competent at spotting this situation and removes all traces of the hash ever existing. The <code>Scalar</code> on the other hand seems to be used, so it stays even after spesh optimized the code.</p>
<p>There's also a <code>Sites</code> column here. This lets you spot cases where one or two sites differ strikingly from the others.</p>
<h1 id="callgraphtab">Call Graph Tab</h1>
<p><img src="https://wakelift.de/content/images/2018/08/Selection_041.png" alt="Selection_041" style="width: 100%;" /></p>
<p>Here's the new <code>Call Graph</code> tab. It's similar to the old one with one major omission. The new version of the call graph explorer currently lacks a flame graph (or icicle graph in this case). It will return later, of course.</p>
<p>Until then, there's a few improvements to be enjoyed. One of them is that the breadcrumbs navigation now works reliably, whereas in the previous profiler it tended to lose elements near the beginning or in the middle sometimes. On top of that, your browser's back and forward buttons will work between nodes in the call graph as well as the different tabs!</p>
<p>Something that's completely new is the Allocations section at the bottom of the page, shown in the screenshot below:</p>
<p><img src="https://wakelift.de/content/images/2018/08/Selection_042.png" alt="Selection_042" style="width: 100%;" /></p>
<p>Here you can see that the <code>Routine</code> in question (<a href="https://github.com/timo/json_fast/blob/d58ca555ff1bccfb5a48daec47249fc5259856c0/lib/JSON/Fast.pm#L168-L175">it's the body of this for loop</a> allocates <code>Str</code> and <code>Scalar</code> objects. The <code>Str</code> objects seem to be optimized by spesh again.</p>
<p>There's still two unexplained buttons here, though. The one at the bottom is labelled "Load inclusive allocations". Clicking on it reveals a second table of allocations, which is quite a lot bigger:</p>
<p><img src="https://wakelift.de/content/images/2018/08/Selection_043.png" alt="Selection_043" style="width: 100%;" /></p>
<p>This view is about everything allocated by anything in the call graph from the current node downwards (towards the leaves, not towards the root. You know, because trees hang from the ceiling, right?). For something so close to the root in a rather deep call graph, you'll get quite a big list, and it's not very helpful in finding where exactly individual types are being allocated.</p>
<p>That's where the second button comes in. It says "Show allocations for all children" on it, and clicking it expands every row in the table of routines:</p>
<p><img src="https://wakelift.de/content/images/2018/08/Selection_044.png" alt="Selection_044" style="width: 100%;" /></p>
<p>This way you can drill down from the root towards nodes that interest you.</p>
<h1 id="whatsmissing">What's missing?</h1>
<p>Here's a little overview of what I'd like to include in the near and medium future:</p>
<ul>
<li>Features the old profiler had
<ul>
<li>An overview page that shows a lot of things at once</li>
<li>Sorting and filtering in the routine list, GC table, ...</li>
<li>An icicle graph or flame graph in the call graph tab</li>
<li>An allocation overview tab</li>
<li>An overview over deoptimizations and on-stack-replacements</li>
<li>A graph of cleared / retained / promoted bytes in the GC</li>
</ul>
</li>
<li>Features that would be new
<ul>
<li>Searching deeply in the call graph</li>
<li>Sorting by fields that the old profiler didn't have, like time per entry</li>
<li>Filtering the GC view to show only major or only minor collections</li>
<li>Callers tab for expanded routines in the overview</li>
<li>Perhaps: An "inlining performance" tab for every routine, since inlining is such an important part of Perl 6 performance</li>
</ul>
</li>
<li>Features requiring changes to the MoarVM and NQP parts of the profiler
<ul>
<li>Expose the actual size of types in the allocation displays</li>
<li>Split the GC's cleared / retained / promoted bytes into managed and unmanaged parts. <sup class="footnote-ref"><a href="#fn2" id="fnref2">[2]</a></sup></li>
<li>Split time spent in routines into time spent in interpreted vs speshed/jit modes (could be difficult with OSR?)</li>
<li>Record the "earliest call" time for every call graph node</li>
</ul>
</li>
</ul>
<h1 id="wherecanigetit">Where can I get it?</h1>
<p>I just uploaded all the code to my github. <a href="https://github.com/timo/moarperf/">You can find it here</a>. To use it, you will have to run <code>npm install</code> in the source folder to get all the javascript dependencies, and <code>npm run build</code> to compile all the javascript bundles. It will run a watcher process that will update as soon as any sources change, but if you're not working on the moarperf source code, you can just kill it after it has done its thing once.</p>
<p>Next, you'll have to install the dependencies of the Perl 6 program that acts as the back-end. You can do that with <code>zef --depsonly install</code> in the same folder as the <code>META6.json</code>. Please note that I haven't prepared the app for actually being installed, so don't do that yet :)</p>
<p>You can then start the backend with <code>perl6 -Ilib service.p6 /path/to/profile.sql</code>, where profile.sql is a profile generated with a commandline like <code>perl6 --profile --profile-filename=profile.sql my_script.p6</code>. Passing the filename to <code>service.p6</code> on the commandline is optional, you can also enter it in the web frontend. By default, it is reachable on <a href="http://localhost:20000">http port 20000 on localhost</a>, but you can set <code>MOARPERF_HOST</code> and <code>MOARPERF_PORT</code> in the environment using the <code>env</code> or <code>export</code> commands of your shell.</p>
<p>A word of warning, though: It still has a bunch of rough edges. In some parts, loading data isn't implemented cleanly yet, which can lead to big red blocks with frowny faces. A refresh and/or doing things more slowly can help in that case. There's places where "there is no data" looks a lot like "something broke". Little usability problems all around.</p>
<p>I would appreciate a bit of feedback either on the github issue tracker, on <code>#perl6</code> on the freenode IRC server, on reddit if someone's posted this article in <code>/r/perl6</code>, or via mail <a href="https://perl6.org/community/">to the <code>perl6-users</code> mailing list</a> or to <code>timo</code> at this website's domain name.</p>
<p>Thanks again to The Perl Foundation for funding <a href="http://news.perlfoundation.org/2017/09/grant-proposal-rakudo-perl-6-p.html">my grant</a>, and to you for reading my little report.</p>
<p>Have fun with the program!<br /> - Timo</p>
<hr class="footnotes-sep" />
<section class="footnotes">
<ol class="footnotes-list">
<li id="fn1" class="footnote-item">
<p>The reasoning behind the different default sorting modes is that in the first step you are likely more interested in routines that are expensive by themselves. Sorting by inclusive time just puts the entry frame first, and then gives you the outermost frames that quite often hardly do any work themselves. When you've found a routine that has a lot of exclusive time, the next step is - at least for me - to look at what routines below that take up the most time in total. That's why I prefer the inner routines list to be sorted by inclusive time. Once the user can change sorting in the UI i'll also offer a way to set the selected sorting as the default, i think. <a href="#fnref1" class="footnote-backref">â†©ï¸Ž</a></p>
</li>
<li id="fn2" class="footnote-item">
<p>The difference between managed and unmanaged bytes is that the managed bytes are what lands in the actual nursery, whereas unmanaged bytes refers to all kinds of extra data allocated for an Object. For example, an <code>Array</code>, <code>Hash</code>, or <code>String</code> would have a memory buffer somewhere in memory, and a header object pointing to that buffer in the nursery or gen2 memory pools. The managed size doesn't change for objects of the same type, which is why it's fine to put them in the allocations tabs. <a href="#fnref2" class="footnote-backref">â†©ï¸Ž</a></p>
</li>
</ol>
</section>
</div>

