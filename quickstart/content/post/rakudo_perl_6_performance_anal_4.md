
---
title: "Rakudo Perl 6 performance analysis tooling - Grant Report #5"
author: Alan Haggai Alavi
type: post
date: 2018-11-05 00:00:00 +0000 UTC
url: "/post/rakudo_perl_6_performance_anal_4"
categories:
 - Grants
 - Perl 6 Development

---

<p>"Overview" tab is now functional but in flux. The "Routines" tab has been improved to include sorting functionality for columns, a minimal view in the "Paths" sub tab and a new "Callers" sub tab. An "Allocations" top level tab has also been added. Read more at: <a href="https://wakelift.de/2018/10/28/full-screen-ahead/" title="Full Screen Ahead!">Full Screen Ahead!</a></p>
<hr />
<p></p>
<h1 class="post-title">Full Screen Ahead!</h1>
<div class="kg-card-markdown">
<p>Whew, it's been a long time since the last report already! Let's see what's new.</p>
<p><img src="https://images.unsplash.com/photo-1505832018823-50331d70d237?ixlib=rb-0.3.5&amp;q=80&amp;fm=jpg&amp;crop=entropy&amp;cs=tinysrgb&amp;w=1080&amp;fit=max&amp;ixid=eyJhcHBfaWQiOjExNzczfQ&amp;s=b8eab5d0af2205b40d99e4f8ace6df10" alt="train on bridge surrounded with trees at daytime" style="width: 100%;" /><br /> <small>Photo by <a href="https://unsplash.com/@jack_anstey?utm_source=ghost&amp;utm_medium=referral&amp;utm_campaign=api-credit">Jack Anstey</a> / <a href="https://unsplash.com/?utm_source=ghost&amp;utm_medium=referral&amp;utm_campaign=api-credit">Unsplash</a></small></p>
<h1 id="improvementstotheroutineoverview">Improvements to the Routine Overview</h1>
<p>The first tab you'll usually go to is the routine overview. It gives you a list of every routine that was encountered in the program. In a bunch of columns it shows you the name of the routine, along with the file it comes from and the line it's defined on. It tells you how often the routine was entered and how many of these entries were into or inside of jitted code.</p>
<p>The last column shows how much time the routine accounted for. The last part is split up twice. Once into inclusive vs exclusive time, which tells you how much total time was spent from entry to exit vs how much time was spent only inside the routine itself, rather than any routines called by it. The second split is into total time vs time divided by number of entries. That lets you more easily figure out which routines are slow by themselves vs which routines take a lot of time because they are called very, very often.</p>
<p>Usually it's best to start with the routines that take the most time in total, but routines that take longer per call may have more obvious optimization opportunities in them in my experience.</p>
<p>It is for that reason that being able to sort by different columns is extra important, and that wasn't available in the new tool for a whole while. That changed just yesterday, though! It still looks quite odd, because it's just a bunch of buttons that don't even react to which sorting mode is currently active, but that will change soon enough.</p>
<p><img src="https://wakelift.de/content/images/2018/10/Selection_074.png" alt="Selection_074" style="width: 100%;" /></p>
<p>The routines overview of course lets you expand the individual routines to get at more details. The details that were already in the previous blog post, were Callees, Paths, and Allocations. All three of them have changed a bit, and one more has been added recently.</p>
<p>All that happened in the Callees tab is that it is now also sortable.</p>
<p>The Paths tab has got a little toggle button called "Minimal" that reduces every cell in the tree to a little square. On top of that, the cells are now colored by the routine's filename - in the normal view the lower border of the cells shows the color, in the minimal view the cells themselves are colored in.</p>
<p>Here's two screenshots comparing regular and minimal view:</p>
<p><img src="https://wakelift.de/content/images/2018/10/Selection_076.png" alt="Selection_076" style="width: 100%;" /></p>
<p><img src="https://wakelift.de/content/images/2018/10/Selection_077.png" alt="Selection_077" style="width: 100%;" /></p>
<p>The Allocations tab now shows the total amount of memory allocated by the given routine for the given type. Of course this doesn't directly correspond to memory usage, since the profiler can't tell you how long the given objects survive.</p>
<p><img src="https://wakelift.de/content/images/2018/10/Selection_078.png" alt="Selection_078" style="width: 100%;" /></p>
<p>The new tab that was recently added is a very useful one. It's the Callers tab. It lets you see which routines have called the given routine, how often they called it, and to what percentage the routine got inlined into the callers. Here's a screenshot:</p>
<p><img src="https://wakelift.de/content/images/2018/10/Selection_079.png" alt="Selection_079" style="width: 100%;" /></p>
<h1 id="allocationstab">Allocations Tab</h1>
<p>There's a whole new top-level tab, too. It's titled "Allocations" and does very much what you would expect.</p>
<p>The tab contains a big table of all types the profiler saw in the program. Every row shows the Type name along with potentially a link to the docs website, the size of each object and how much memory in total was allocated for it, and how often that type of object was allocated in total.</p>
<p>On the left there is a button that lets you expand the type's column to get extra details:</p>
<p><img src="https://wakelift.de/content/images/2018/10/Selection_080.png" alt="Selection_080" style="width: 100%;" /></p>
<p>None of these things are very surprising, but they are useful.</p>
<h1 id="overviewtab">Overview Tab</h1>
<p>The overview tab used to just have a placeholder text. However, it now displays at least a little bit of information. Have a look!</p>
<p><img src="https://wakelift.de/content/images/2018/10/Selection_081.png" alt="Selection_081" style="width: 100%;" /></p>
<p>It's still very fresh and I expect it to look completely different in a week or two.</p>
<h1 id="whatswiththefullscreenpun">What's with the "Full Screen" pun?</h1>
<p>This post starts with a silly little pun on "Full Steam Ahead". What does it mean for the profiler tool?</p>
<p>The answer is very simple, it lets you use the whole width of your browser window. Bootstrap â€“ which is the library of CSS styles and primitives I use so that I don't have to spend three fourths of the entire development time fiddling with CSS rules so that things look okay or even work at all â€“ is convinced that the outermost Container element of the site shouldn't be wider than a specific width, probably because things can look kind of "lost" on a page that's not filled side-to-side with stuff. If you have big tables full of data or wide graphs with tiny bars, it sure is good to be able to expand the page sideways.</p>
<p>Here's a little animation I made a few weeks ago that shows the fullscreen button in action for one example.</p>
<h1 id="youremyfavouritecustomer">You're my favourite customer!</h1>
<p>Stefan 'nine' Seifert has recently been working on replacing the "mast" stage in the Perl 6 compiler.</p>
<p>In the simplest terms, it used to take the abstract syntax tree to rewrite it into a very flat tree consisting of lists of moar instructions. Those were then fed into MoarVM as objects, which were then translated into actual moar bytecode, literally as bytes.</p>
<p>The main drawback of the earlier approach was memory use. On top of all the QAST node objects that were the input to the mast stage, a huge batch of new objects were created before it was all passed to moar's bytecode compiler as one set. The more alive objects there are, the higher the maximum memory usage gets, and the longer it takes the GC to go through everything to decide what's garbage and what needs to stick around.</p>
<p>The new approach starts with the QAST nodes again, but immediately starts writing out bytes into buffers. That has a whole bunch of implications immediately: Every Object consists of at least a header that's a couple of bytes big. Buffers (and native arrays) are just a single object with a big consecutive blob of memory they "own". No matter how big the blob grows, the garbage collector only has to look at the object header and can ignore the blob itself! Not only does it save memory by not having as many individual objects, it also saves the garbage collector time!</p>
<p>Stefan told me on IRC that the new profiler tool already helped him a lot in making the new implementation faster.</p>
<blockquote>
<p>I now managed to get a profile of the MAST stage using a patched --profile-stage, but the profile is too large even for the QT frontend :/<br /> timotimo++ # moarperf works beautifully with this huge profile :)</p>
</blockquote>
<p><a href="https://colabti.org/irclogger/irclogger_log/perl6-dev?date=2018-10-09#l56">https://colabti.org/irclogger/irclogger_log/perl6-dev?date=2018-10-09#l56</a></p>
<blockquote>
<p>timotimo++ # the profiler is indespensable in getting nqp-mbc merge worthy.</p>
</blockquote>
<p><a href="https://colabti.org/irclogger/irclogger_log/perl6-dev?date=2018-10-23#l66">https://colabti.org/irclogger/irclogger_log/perl6-dev?date=2018-10-23#l66</a></p>
<p>This obviously makes yours truly very happy :)</p>
<p>There's still lots of "paper cuts" to be ironed out from a User Experience standpoint, but I'm very glad that the tool is already making work possible that used to be impossible.</p>
<p>That's all the improvements I have to report for now, but I'll be continuing work on the UI and backend and write up another report (hopefully) soon!</p>
<p>BTW, a big conference on React.js (the tech I use to make the frontend tick) just happened and the react core dev team unveiled <a href="https://reactjs.org/docs/hooks-custom.html">a few new APIs</a> and <a href="https://reactjs.org/docs/hooks-effect.html">tricks</a> that I expect will let me make the existing code a lot simpler and write new code faster than before!</p>
<p>Also, during the conference <a href="https://twitter.com/darthtrevino">Chris Trevino</a> presented a new charting library based around the "Grammar of Graphics" philosophy(?) that is also the basis of <a href="https://vega.github.io/">vega and vega-lite</a> which as I understand it is The Big Cheese among data visualization experts. I've been having some gripes with recharts, so I'll be checking this out soon!</p>
<p>Hope to see y'all in the next one!<br />   - Timo</p>
</div>

