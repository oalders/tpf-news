
---
title: "Rakudo Perl 6 performance analysis tooling - Grant Report #1"
author: Alan Haggai Alavi
type: post
date: 2018-03-07 00:00:00 +0000 UTC
url: "/post/rakudo_perl_6_performance_anal"
categories:
 - Grants

---

<p>Timo has made good progress and is writing about it. Please see: <a href="https://wakelift.de/2018/03/05/delays-and-delights/">Delays and Delights</a></p>
<hr />
<p></p>
<h1 class="post-title">Delays and Delights</h1>
<div class="kg-card-markdown">
<p>Hi, my name is timotimo and I'm a Perl 6 developer. I've set up this blog to write reports on my <a href="http://news.perlfoundation.org/2017/09/grant-proposal-rakudo-perl-6-p.html">TPF Grant</a>.</p>
<p>Before the actual report starts, I'd like to issue an apology. In between my grant application and the grant being accepted I developed a bit of RSI that lasted for multiple months. I already anticipated that near the end of January a move would occupy me a bit, but I had no idea how stressful it would turn out to be, and how long afterwards it would keep me occupied.</p>
<p><img src="https://images.unsplash.com/photo-1442600678244-fc0253be6acb?ixlib=rb-0.3.5&amp;q=80&amp;fm=jpg&amp;crop=entropy&amp;cs=tinysrgb&amp;w=1080&amp;fit=max&amp;ixid=eyJhcHBfaWQiOjExNzczfQ&amp;s=e43bc5f3b1529e88b4fd2c3ecddee827" alt="" style="width: 100%;" /><br /> <small>Photo by <a href="https://unsplash.com/@lyndsburk?utm_source=ghost&amp;utm_medium=referral&amp;utm_campaign=api-credit">Lyndsey B</a> / <a href="https://unsplash.com/?utm_source=ghost&amp;utm_medium=referral&amp;utm_campaign=api-credit">Unsplash</a></small></p>
<p>I regret that it took me so long to actually get started. However, I've already got little bits to show for the first report of this grant!</p>
<h1 id="delight1noncrashymultithreadedprofiles">Delight &#8470;1: Non-crashy multi-threaded profiles</h1>
<p>Until now if you've added a signal handler for <code>Ctrl-C</code> to your program, or used <code>run</code>, <code>shell</code>, or even <code>Proc::Async</code> or any async IO, rakudo or moarvm will have started an additional thread for you. Even if you don't necessarily realize it, this caused the profiler to either violently explode, or return useless or partial results.</p>
<p>Just these past few days I've made a bit of reliability work for the profiler available to everyone running rakudo from the latest source code. Now the profiler will only sometimes abort with the error message "profiler lost sequence" â€“ a bug that I'm also going to chase down as part of the grant. On top of that, the HTML profiler frontend now shows timing and allocations from every thread.</p>
<h1 id="delight2fasterheapsnapshots">Delight â„–2: Faster Heap Snapshots</h1>
<p><small>N.B.: I have done the work discussed in this section before the grant actually started, but I think it's still interesting for you.</small></p>
<p>As you can tell from the â€“ sadly mildly messed-up â€“ grant application, I also work on a second profiler: The Heap Snapshot Profiler. It takes a snapshot of the structure of all objects in memory, and their connections. With this information, you can figure out what kinds of objects account for how much of your memory usage, and for what reason any given object is not being removed by the garbage collector.</p>
<p>If you've already tried out this profiler in the past, you may have noticed that it incurs a significant memory usage increase during run-time. After that, it takes a surprising amount of time to load in the <a href="https://github.com/jnthn/p6-app-moarvm-heapanalyzer">heap snapshot analyzer</a>. This changed noticeably when I implemented a new format for heap snapshots.</p>
<p>Until then, the format had one line (i.e. <code>\n</code> delimited) per snapshot (i.e. for each GC run) and a json-encoded list of strings up front. The snapshots themselves were then a few lines with lists of integers.<br /> This caused two little issues in terms of performance:</p>
<ul>
<li>Encoding the list of strings as JSON had to be done in <abbr title="Not Quite Perl">NQP</abbr> code, because of the somewhat complex escaping rules.<sup class="footnote-ref"><a href="#fn1" id="fnref1">[1]</a></sup></li>
<li>Reading the heap snapshot requires a whole lot of string splitting and integer parsing, which is comparatively slow.</li>
</ul>
<p>The binary-based format I replaced it with addresses both of these issues, and has a few extra features for speed purposes:</p>
<ul>
<li>The list of strings is written as a 64 bit integer for the string length followed by the string itself encoded in utf-8. At the very front is a single 64 bit integer that holds the list's size.</li>
<li>Snapshot data is encoded with either 64 bit integers or a very simple variable-sized integer scheme. Each list stores its size at the beginning, too.</li>
<li>The very end of the file contains an index with the sizes of all the sections, so that skipping to any point of interest is cheap.</li>
</ul>
<p>Armed with the index at the end of the file, the binary format can be read by multiple threads at the same time, and that's exactly what the new Heap Snapshot Analyzer will do. In addition to the start positions of each section, there's also a pointer right into the middle of the one section that holds variable-sized entries. That way both halves can be read at the same time and cheaply combined to form the full result.</p>
<p>The "summary all" command loads every snapshot, measures the total size of recorded objects, how many objects there are, and how many connections exist. The result is displayed as a table. Running this against a 1.1 gigabyte big example snapshot with 44 snapshots takes about a minute, uses up 3â…“ CPU cores out of the 4 my machine has, and maxes out at 3 gigs of ram used.<sup class="footnote-ref"><a href="#fn2" id="fnref2">[2]</a></sup> That's about 19 megabytes per second. It's not blazing fast, but it's decent.</p>
<h1 id="delight3webfrontend">Delight â„–3: Web Frontend</h1>
<p>Sadly, there isn't anything to show for this yet, but I've <a href="http://cro.services/docs/intro/spa-with-cro">started a Cro application</a> that lets the user load up a heap snapshot file and load individual snapshots in it.</p>
<p>It doesn't seem like much at all, but the foundation on which the other bits will rest is laid now.</p>
<p>I intend for the next posts to have a bunch of tiny animated gif screencasts for you to enjoy!</p>
<h1 id="delight4fasterprofilefilewriting">Delight â„–4: Faster Profile File Writing</h1>
<p>The SQL output will likely be the primary data storage format after my grant is done. Therefore I've taken first steps to optimize outputting the data. Right now it already writes out data about 30% faster during the actual writing stage. The patch needs a bit of polish, but then everyone can enjoy it!</p>
<h1 id="nextsteps">Next steps</h1>
<p>The next thing I'd like to work towards is trying out the profiler with a diverse set of workloads: From hyper/race to a Cro application. That will give me a good idea of any pain points you'd encounter. For example, once the ThreadPoolScheduler's supervisor thread runs, it'll spend a majority of its time sleeping in between ticks. This shows up very prominently in the Routines list, to name one example.</p>
<p>Of course, I take inspiration from my fellow Perl 6 developers and users on the <a href="https://perl6.org/community/irc">IRC channel</a> and the <a href="https://perl6.org/community/">mailing lists</a>, who offer a steady supply of challenges :)</p>
<p>If you have questions, suggestions, or comments, please feel free to ping me on IRC, or <a href="https://www.reddit.com/domain/wakelift.de/">find this blog post on reddit</a>.</p>
<p>Thanks for reading; see you in the next one!<br /> - Timo</p>
<hr class="footnotes-sep" />
<section class="footnotes">
<ol class="footnotes-list">
<li id="fn1" class="footnote-item">
<p>It could have been done in C code, but it feels weird to have JSON string escaping code built into MoarVM. <a href="#fnref1" class="footnote-backref">â†©ï¸Ž</a></p>
</li>
<li id="fn2" class="footnote-item">
<p>Figuring out why the ram usage is so high is also on my list. Amusingly, the heap analyzer can help me improve the heap analyzer! <a href="#fnref2" class="footnote-backref">â†©ï¸Ž</a></p>
</li>
</ol>
</section>
</div>
<p></p>

