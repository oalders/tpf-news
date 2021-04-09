
---
title: "Rakudo Perl 6 performance analysis tooling - Grant Report #2"
author: Alan Haggai Alavi
type: post
date: 2018-07-05 00:00:00 +0000 UTC
url: "/post/rakudo_perl_6_performance_anal_1"
categories:
 - Grants
 - Perl 6 Development

---

<p>Timo is still not completely well. However, he has been able to make good progress:</p>
<ol start="1">
<li>
<p>Optional parametres are now correctly logged</p>
</li>
<li>
<p>Fixed a crash caused when the profiler is used on multithreaded code</p>
</li>
</ol>
<p>Read more at: <a href="https://wakelift.de/2018/06/15/no-major-breakthroughs/">No Major Breakthroughs</a></p>
<hr />
<p></p>
<h1 class="post-title">No Major Breakthroughs</h1>
<div class="kg-card-markdown">
<p>Sadly, the time since the last post on this blog hasn't been fruitful with regards to the profiling project. There have been slight improvements to the profiler inside MoarVM, like handling profiles with a very deep call graph better, making the first GC run show up again, capturing allocations from optional parameters properly, and hopefully finally making programs that have multiple threads running no longer crash during the profile dumping phase. A recently merged <a href="http://brrt-to-the-future.blogspot.com/2018/06/controlled-stack-hacking-for-moarvm-jit.html">branch by esteemed colleague brrt</a> will allow me to properly fix one nasty issue that remains in the profiler that relates to inlining.</p>
<p>Even though I can't show off lovely screenshots of the profiler UI (if you consider Coder's Artâ„¢ lovely), I can briefly go over the changes that have happened and what's next on the list. And of course I'm still very much interested in finishing the grant work!</p>
<p><img src="https://wakelift.de/content/images/2018/06/katzenpfote_soft.jpg" alt="katzenpfote_soft" style="width: 100%;" /></p>
<h2 id="missedoptionalparameters">Missed Optional Parameters</h2>
<p>The first change I'd like to talk about is the one that was causing allocations from boxing optional parameters to go missing from the profile. Optional parameters are implemented as an op that accesses the passed arguments to see if something was present or not. Then it either runs code to put the default value in - if no argument was present - or it skips over that code. Additionally, it handles arguments that were passed as native ints, nums, or strings.</p>
<p>If an object was expected by the code that uses the parameter, this op will also create a box for the value, for example an Int object. The crucial mistake was in the instrumentation by the profiler.</p>
<p>Finding everything that is allocated is done by putting a little "take note of this object" op after every op that may create an object. This op then checks if the object was probably allocated by the last instruction, or if it was probably already logged earlier. If it was just allocated, that allocation is recorded for the profile.</p>
<p>The problem in this case lies in the placement of the logging op: It was placed right after the instruction that grabs the argument. However, that made it land in the place that gets skipped over if an argument was present. So either no argument was passed, and the logging op was just asked to log that a null was allocated, or an argument was passed that was perhaps boxed, and the logging op was skipped over. Oops!</p>
<p>Fixing this was simply a matter of following the skip and putting the logging op in the right place.</p>
<h2 id="multithreadedprogramscrashingmysteriously">Multi-threaded Programs Crashing Mysteriously</h2>
<p>If you used the profiler on code that runs multiple threads, you may have seen very suspicious looking internal error messages like "const_iX NYI" pop up. This was caused by the instrumentation aspect of the profiler, more specifically what it did when the instrumentation was no longer needed. Allow me to explain:</p>
<p>Instrumentation in this context refers to creating a version of the program bytecode that does some extra work in the right places. For the profiler this includes putting ops in the code that record that a function was called or exited, and ops that record allocations of objects.</p>
<p>This instrumentation happens lazily, i.e. when a function is entered the first time, it runs into the "instrumentation barrier", which pauses the program and creates the instrumented code right then and there. The instrumented code then gets installed and the program continues. This is implemented by having a global "instrumentation level" that just gets increased by 1 every time functions should go through an instrumentation step. This is done when profiling starts, and it is done when profiling ends.</p>
<p>Here's where the problem lies: Profiling is turned on before user code runs, which just happens to always be in single-threaded territory. However, profiling gets turned off as soon as the main thread is finished. This is done by increasing the instrumentation level by 1 again. Every function that is entered from now on will have to go through instrumentation again, which will restore the original bytecode in this case.</p>
<p>Other threads might still continue running, though. The first example that made this problem clear was finding the 1000st prime by grepping over a hypered range from 0 to infinity. Crucially, after finding the 1000st prime, some workers were still busy with their batch of numbers.</p>
<p>Here's where the instrumentation barrier becomes a problem. One of the remaining workers calls into a function, for example is-prime, for the first time since the instrumentation level was changed. It will have its instrumented bytecode replaced by the original bytecode. However, the other threads, which may still be inside is-prime in this example, will not know about this. They keep happily interpreting the bytecode when all of a sudden the bytecode changes.</p>
<p>Since the uninstrumented bytecode is shorter than the instrumented bytecode, the worst case is that it reads code past the end of the bytecode segment, but the more common case is that the instruction pointer just suddenly points either at the wrong instruction, or in the middle of an instruction.</p>
<p>Instructions usually start with the opcode, a 16 bit number usually between 0 and 1000. The next part is often a 16 bit number holding the index of a register, which is usually a number below about 40, but quite often below 10. If the instruction pointer accidentally treats the register number as an opcode, it will therefor often land on ops with low numbers. Opcode 0 is no_op, i.e. "do nothing". The next three ops are const_i8 through const_i32, which all just throw the exception that I mentioned in the first paragraph: "const_iX NYI". Two spots ahead is the op "const_n32", which also thrown as NYI error.</p>
<p>And there you have it, mystery solved. But what's the solution to the underlying problem? In this case, I took the easy way out. All the profiling ops first check if profiling is currently turned on or not anyway, so leaving the instrumented code in after profiling has ended is not dangerous. That's why MoarVM now just keeps instrumentation the same after profiling ends. After all, the next thing is usually dumping the profile data and exiting anyway.</p>
<h2 id="thenextsteps">The Next Steps</h2>
<p>The MoarVM branch that brrt recently merged is very helpful for a very specific situation that can throw the profiler off and cause gigantic profile files: When a block has its bytecode inlined into the containing routine, and the block that was inlined had a "return" in it, it knows that it has to "skip" over the inner block, since blocks don't handle returns.</p>
<p>However, the block is still counted as a routine that gets entered and left. The long and short of it is that returning from the inner block jumps directly to the exit, but having the block inlined frees us from doing the whole "create a call frame, and tear it down afterwards" dance. That dance would have contained telling the profiler that a frame was exited "abnormally"; since the regular "prof_exit" op that would have recorded the exit will be skipped over, tearing down the frame would have contained the logging.</p>
<p>In this particular case, though, no exit would be logged! This makes the call graph - think of it like <a href="http://www.brendangregg.com/flamegraphs.html">a flame graph</a> - look very strange. Imagine a function being called in a loop, and returning from an inner block as described above. It would miss all of the exits, so every time the function is called again, it will look like the function called itself, never returning to the loop. Every time around the loop, the call will seem to be nested deeper and deeper. Since the profiler keeps around the whole call graph, the file will just keep growing with every single iteration.</p>
<p>Now, how does brrt's code change this situation? It will allow very easily to figure out how many inlines deep a "return from this routine" op is, so that the profiler can accurately log the right amount of exits.</p>
<p>On the UI side of things, I want to bring the routine overview list into a good state that will finally be worth showing. The list of GC runs will also be interesting, especially since the profiler recently learned to log how each individual thread performed its GC run, but the current HTML frontend doesn't know how to display that yet.</p>
<p>Hopefully the wait for the next post on my blog won't be as long as this time!<br />   - Timo</p>
</div>

