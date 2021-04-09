
---
title: "Rakudo Perl 6 performance analysis tooling - Grant Report #6"
author: Alan Haggai Alavi
type: post
date: 2018-11-21 00:00:00 +0000 UTC
url: "/post/rakudo_perl_6_performance_anal_5"
categories:
 - Grants
 - Perl 6 Development

---

<p><em>Original article was published on November 9, 2018</em></p>
<p>The overview page now shows all data displayed in the previous profiler's page as well as adds a "Start times of threads" chart. "GC" tab has been updated with sub-tabs to customise graphs using different display modes. The routines list now features a "goto" arrow for smooth and easy navigation.</p>
<p>Read more at: <a href="https://wakelift.de/2018/11/09/where-did-i-leave-my-at-key-s/" title="Where did I leave my AT-KEYs?">Where did I leave my AT-KEYs?</a></p>
<hr />
<h1 class="post-title">Where did I leave my AT-KEYs?</h1>
<div class="kg-card-markdown">
<p>Even though it's only been a week and a half since my last report, there's already enough new stuff to report on! Let's dig right in.</p>
<p><img src="https://images.unsplash.com/photo-1520468447370-264ed2d81430?ixlib=rb-0.3.5&amp;q=80&amp;fm=jpg&amp;crop=entropy&amp;cs=tinysrgb&amp;w=1080&amp;fit=max&amp;ixid=eyJhcHBfaWQiOjExNzczfQ&amp;s=9c570217201c526fe928620f12bfdfd6" alt="shallow focus of lovelocks" style="width: 100%;" /><br /> <small>Photo by <a href="https://unsplash.com/@neonbrand?utm_source=ghost&amp;utm_medium=referral&amp;utm_campaign=api-credit">NeONBRAND</a> / <a href="https://unsplash.com/?utm_source=ghost&amp;utm_medium=referral&amp;utm_campaign=api-credit">Unsplash</a></small></p>
<h2 id="overviewpage">Overview Page</h2>
<p><img src="https://wakelift.de/content/images/2018/11/Selection_099.png" alt="Selection_099" style="width: 100%;" /></p>
<p>Is there terribly much to say about this? It now shows the same data that was already available in the old profiler's overview page. It's the go-to page when you've changed your code in the hopes of making it faster, or changed your version of rakudo in hopes of not having to change your code in order to make it faster.</p>
<p>Here's a screenshot from the other profiler for comparison:</p>
<p><img src="https://wakelift.de/content/images/2018/11/Selection_100.png" alt="Selection_100" style="width: 100%;" /></p>
<p>The main addition over the previous version is the "start times of threads" piece at the top left. In multi-threaded programs it shows you when more threads were added, for example if you use <code>start</code> blocks on the default <code>ThreadPoolScheduler</code>.</p>
<p>The GC Performance section gives you not only the average time spent doing minor and major collections, but also the minimum and maximum time.</p>
<p>The rest is pretty much the same, except the new version puts separators in numbers with more than three digits, which I find much easier on the eyes than eight-digit numbers without any hints to guide the eye.</p>
<h2 id="gcrunlist">GC Run List</h2>
<p><img src="https://wakelift.de/content/images/2018/11/Selection_101.png" alt="Selection_101" style="width: 100%;" /></p>
<p>The graphs at the top of the GC tab has changed a bit! There's now the option to show only major, only minor, or all collections in the graph, and there are three different display modes for the "Amounts of data" graphs.</p>
<p>The one shown by default gives bars split into three colors to signify how much of the nursery's content has been freed (green), kept around for another round (orange), or promoted to the old generation (red). That's the mode you can see in the screenshot above.</p>
<p><img src="https://wakelift.de/content/images/2018/11/Selection_102.png" alt="Selection_102" style="width: 100%;" /></p>
<p>The second mode is "Combined Chart", which just stacks the amounts in kilobytes on top of each other. That means when more threads get added, the bars grow. In this example screenshot, you can barely even see orange or red in the bars, but this program is very light on long-lived allocations.</p>
<p><img src="https://wakelift.de/content/images/2018/11/Selection_103.png" alt="Selection_103" style="width: 100%;" /></p>
<p>The third mode is "Split Charts", which has one chart for each "color". Since they all have their own scales, you can more easily see differences from run to run, even if some of the charts appear tiny in the "combined" or "combined relative" charts.</p>
<h2 id="routineslist">Routines List</h2>
<p>The routines overview â€“ and actually all lists of routines in the program â€“ have a new little clickable icon now. Here it is:</p>
<p><img src="https://wakelift.de/content/images/2018/11/Selection_104.png" alt="Selection_104" style="width: 100%;" /></p>
<p>The icon I'm talking about is the little up-right-arrow in a little box after a routine's name. When you click on it, the row turns blue. Huh, that doesn't sound so useful? That's because the button brings you to the routines list and scrolls to and highlights the routine you've clicked on. If you're already right there, you will not notice a lot of change, of course.</p>
<p>However, it gets more interesting in the callers or callees lists:</p>
<p><img src="https://wakelift.de/content/images/2018/11/moarperf-navigation-routineslist.gif" alt="moarperf-navigation-routineslist" style="width: 100%;" /></p>
<p>Even better, since it actually uses links to actual URLs, the browser's back/forward buttons work with this.</p>
<p>Other useful places you can find this navigation feature are the allocations list and the call graph explorer:</p>
<p><img src="https://wakelift.de/content/images/2018/11/Selection_105.png" alt="Selection_105" style="width: 100%;" /></p>
<p><img src="https://wakelift.de/content/images/2018/11/Selection_106.png" alt="Selection_106" style="width: 100%;" /></p>
<h2 id="wherearemyatkeysat">Where are my <code>AT-KEY</code>s at?</h2>
<p>If you have a very big profile, a routine you're interested in may be called in many, many places. Here's a profile of "zef list". Loading up the call graph for this routine may just explode my computer:</p>
<p><img src="https://wakelift.de/content/images/2018/11/Selection_107.png" alt="Selection_107" style="width: 100%;" /></p>
<p>Note the number of Sites: 27 thousand. Not good.</p>
<p>But what if you're already in the call graph explorer anyway, and you just want to find your way towards functions that call your routine?</p>
<p>Enter the search box:</p>
<p><img src="https://wakelift.de/content/images/2018/11/search-for-at-key.gif" alt="search-for-at-key" style="width: 100%;" /></p>
<p>As you can see, when you input something in the search bar, hand icons will point you towards your destination in the call graph.</p>
<p>I'm looking to add many more different kinds of searches, for example I can imagine it would be interesting to see at a glance "which branches will ever reach non-core code". Searching for files ought to also be interesting.</p>
<p>Another idea I've had is that when you've entered a search term, it should be possible to exclude specific results, for example if there are many routines with the same name, but some of them are not the ones you mean. For example, "identity" is in pretty much every profile, since that's what many "return"s will turn into (when there's neither a decont nor a type check needed). However, Distributions (which is what zef deals in) also have an "identity" attribute, which is about name, version, author, etc.</p>
<p>At a much later point, perhaps even after the grant has finished, there could also be search queries that depend on the call tree's shape, for example "all instances where <code>&amp;postcircumfix:{ }</code> is called by <code>&amp;postcircumfix:{ }</code>".</p>
<h2 id="thatsit">That's it?</h2>
<p>Don't worry! I've already got an extra blog post in the works which will be a full report on overall completion of the grant. There'll be a copy of the original list (well, tree) of the "deliverables and inchstones" along with screenshots and short explanations.</p>
<p>I hope you're looking forward to it! I still need to replace the section that says "search functionality is currently missing" with a short and sweet description of what you read in the previous section :)</p>
<p>With that I wish you a good day and a pleasant weekend<br /> - Timo</p>
</div>

