
---
title: "Perl 6 IO Grant: February 2017 Report"
author: Coke
type: post
date: 2017-03-24 00:00:00 +0000 UTC
url: "/post/perl_6_io_grant_february_2017_"
categories:
 - Grants
 - Perl 6 Development

---


<p><i>Zoffix Znet provided this report on February 26, 2017</i></p>

<hr>
                            <p>This document is the February, 2017 progress report for <a href="http://news.perlfoundation.org/2017/01/grant-proposal-standardization.html">TPF Standardization,
Test Coverage, and Documentation of Perl 6 I/O Routines
grant</a></p>

<h2>Timing</h2>

<p>I'm currently running slightly behind the schedule outlined in the grant. I expect to complete the Action Plan and have it ratified by other core members by March 18th, which is the date of the 2017.03 compiler release. Then, I'll implement all of the Action Plan (and complete the grant) by the 2017.04 compiler release on April 15th. This is also the release the next Rakudo Star distribution will be based on, and so the regular end users will receive better IO there and then.</p>

<p>Some members of the Core Team voiced concerns over implementing any changes that can break users' code, even if the changes do not break 6.c-errata specification
tests. Once the full set of changes is known, they will be reviewed on a
case-by-case basis, and some of them may be implemented under 6.d.PREVIEW
pragma, to be included in 6.d language version, leaving 6.c language versions
untouched. Note that changes that are decided to be 6.d material may delay
the completion of this grant due to not fully-fleshed out infrastructure for
supporting multiple language versions. The April 15th deadline stated above
applies only to changes to 6.c language and new deadline will be ascertained
for completion of the 6.d changes.</p>

<h2>User Communications</h2>

<p>I wrote and disseminated advanced notice of the changes to be made due to this grant, to prepare the users to expect some code to break (some routines were found to be documented, despite being absent entirely from the <a href="https://github.com/perl6/roast/tree/6.c-errata">Specification</a> and not officially part of the language).</p>

<p>The notice can be seen at: <a href="http://rakudo.org/2017/02/26/advance-notice-of-significant-changes/">http://rakudo.org/2017/02/26/advance-notice-of-significant-changes/
    </a></p>

<p>It is possible the Core Team will decide to defer all breaking changes to
6.d language version, to be currently implemented under <code class="prettyprint">v6.d.PREVIEW</code> pragma.</p>

<h2>Bonus Deliverable</h2>

<p>The bonus deliverable—The Map of Perl 6 Routines—is now usable. The code is available in <a href="https://github.com/perl6/routine-map">perl6/routine-map</a> repository, and the rendered version is available on <a href="https://map.perl6.party">map.perl6.party</a>. Its current state is sufficient
to serve the intended purpose for this grant, but I'll certainly add improvements to it sometime in the future, such as linking to docs, linking to routines' source code, having an IRC bot looking stuff up in it, etc.</p>

<p>It'll also be fairy easy to use the Map to detect undocumented routines or ones that are documented under the incorrect type.</p>

<h2>Identified Issues/Deficiencies with IO Routines</h2>

<p>These points, issues, and ideas were identified this month and will be included for consideration in the Action Plan.</p>

<ul>
<li>Calling practically any method on a closed IO::Handle results in an LTA (Less Than Awesome)
error message that reads <code class="prettyprint">&lt;something&gt; requires an object with REPR MVMOSHandle</code> where <code class="prettyprint">&lt;something&gt;</code> is
sometimes the name of the method called by the user and others is some internal method
invoked indirectly. We need better errors for closed file handles; and not something that would require a
<code class="prettyprint">is-fh-closed()</code> type of conditional called in all the methods, which would be a hefty
performance hit.</li>
<li>Several routines have been identified which in other languages return useful information:
number of bytes actually written or current file position, whereas in Perl 6 they just
return a Bool (<code class="prettyprint">.print</code>, <code class="prettyprint">.say</code>, <code class="prettyprint">.write</code>) or a Mu type object (<code class="prettyprint">.seek</code>). Inconsistently,
<code class="prettyprint">.printf</code> does appear to return the number of bytes written. It should be possible
to make other routines similarly useful, although I suspect some of it may have to
wait until 6.d language release.</li>
<li>The <code class="prettyprint">.seek</code> routine takes the seek location as one of three Enum values. Not only are they
quite lengthy to type, they're globally available for no good reason and <code class="prettyprint">.seek</code> is virtually
unique in using this calling convention. I will seek to standardize this routine to take
mutually-exclusive named arguments instead, preferably with much shorter names, but those
are yet to be bikeshed.</li>
<li><code class="prettyprint">IO.umask</code> routine simply shells out to <code class="prettyprint">umask</code>. This fails terribly on OSes that don't have
that command, especially since the code still tries to decode the received input as
an octal string, even after the failure. Needs improvement.</li>
<li><code class="prettyprint">link</code>'s implementation and documentation confuses what a "target" is. Luckily (or sadly?)
there are exactly zero tests for this routine in the Perl 6 Specification, so we can
change it to match the behaviour of <code class="prettyprint">ln</code> Linux command and the <code class="prettyprint">foo $existing-thing, $new-thing</code>
argument order of <code class="prettyprint">move</code>, <code class="prettyprint">rename</code>, and other similar routines.</li>
<li>When using <code class="prettyprint">run(:out, 'some-non-existant-command').out.slurp-rest</code>
it will silently succeed and return an empty string. If possible, this
should be changed to return the failure or throw at some point.</li>
<li><code class="prettyprint">chdir</code>'s <code class="prettyprint">:test</code> parameter for directory permissions test is taken as a
single string parameter. This makes it extremely easy to mistakenly write
broken code: for example, <code class="prettyprint">"/".IO.chdir: "tmp", :test&lt;r w&gt;</code> succeeds, while
<code class="prettyprint">"/".IO.chdir: "tmp", :test&lt;w r&gt;</code> fails with a misleading error message
saying the directory is not readable/writable. I will propose for <code class="prettyprint">:test</code>
parameter to be deprecated in favour of using multiple named arguments to
indicate desired tests. By extension, similar change will be applied to
<code class="prettyprint">indir</code>, <code class="prettyprint">tmpdir</code>, and <code class="prettyprint">homedir</code> routines (if they remain in the language).</li>
<li><em>Documentation:</em> several inaccuracies in the documentation were found. I won't be identifying these in my reports/Action Plan, but will simply ensure the documentation matches the implementation once the Action Plan is fully implemented.</li>
</ul>

<h2>Discovered Bugs</h2>

<p>The hunt for 6-legged friends has these finds so far:</p>

<h4>Will (attempt to) fix as part of the grant</h4>

<ul>
<li>indir() has a race condition where the actual dir it runs in ends up being wrong.
Using <code class="prettyprint">indir '/tmp/useless', { qx/rm -fr */ }</code> in one thread and backing
up your precious data in another has the potential to result in some spectacular failurage.</li>
<li><code class="prettyprint">perl6 -ne '@ = lines'</code> crashes after first iteration, crying about <code class="prettyprint">MVMOSHandle REPR</code>. I suspect
the code is failing to follow iterator protocol somewhere and is attempting to read
on an already closed handle. I expect to be able to resolve this and the related
<a href="https://rt.perl.org/Ticket/Display.html?id=128047">RT#128047</a> as part of the grant.</li>
<li><code class="prettyprint">.tell</code> incorrectly always returns <code class="prettyprint">0</code> on files opened in append mode</li>
<li><code class="prettyprint">link</code> mixes up target and link name in its error message</li>
</ul>

<h4>Don't think I will be able to fix these as part of the grant</h4>

<ul>
<li><code class="prettyprint">seek()</code> with <code class="prettyprint">SeekFromCurrent</code> as location fails to seek correctly if called
after <code class="prettyprint">.readchars</code>, but only on MoarVM. This appears to occur due to some sort of buffering.
I filed this as <a href="https://rt.perl.org/Ticket/Display.html?id=130843">RT#130843</a>.</li>
<li>On JVM, <code class="prettyprint">.readchars</code> incorrectly assumes all chars are 2 bytes long. This appears to be
just a naive substitute for nqp::readcharsfh op. I filed this as
<a href="https://rt.perl.org/Ticket/Display.html?id=130840">RT#130840</a>.</li>
</ul>

<h4>Already Fixed</h4>

<ul>
<li>While making the Routine Map, I discovered <code class="prettyprint">.WHICH</code> and <code class="prettyprint">.Str</code> methods on <code class="prettyprint">IO::Special</code> were <code class="prettyprint">only</code>
methods defined only for the <code class="prettyprint">:D</code> subtype, resulting in a crash when using, say, <code class="prettyprint">infix:&lt;eqv&gt;</code>
operator on the type object, instead <code class="prettyprint">Mu.WHICH</code>/<code class="prettyprint">.Str</code> candidates getting invoked.
This bug was easy and I already commited fix in
<a href="https://github.com/rakudo/rakudo/commit/dd4dfb14d3ccfe50dbd4b425778a005d3303edb9">radudo/dd4dfb14d3</a>
and tests to cover it in
<a href="https://github.com/perl6/roast/commit/63370fe0546eded34cbaa695f6d928aa3db42395">roast/63370fe054</a></li>
</ul>

<h4>Auxiliary Bugs</h4>

<p>While doing the work for this grant, I also discovered some non-IO related bugs (some of which I fixed):</p>

<ul>
<li><code class="prettyprint">.Bool</code>, <code class="prettyprint">.so</code>, <code class="prettyprint">.not</code>, <code class="prettyprint">.has</code>h, and <code class="prettyprint">.elems</code> on <code class="prettyprint">Baggy:U</code> crash
(<a href="https://github.com/rakudo/rakudo/commit/e8af8553eae9abfe4f5cd02dcf4114c5c4877c51">fixed in e8af8553</a>)</li>
<li><code class="prettyprint">.sort</code> on reified empty arrays crashes (<a href="https://rt.perl.org/Ticket/Display.html?id=130866">reported as RT#130866</a>)</li>
<li>SEGV with Scalar type object in unique() (<a href="https://rt.perl.org/Ticket/Display.html?id=130852">reported as RT#130852</a>)</li>
<li><code class="prettyprint">.Bool</code>, <code class="prettyprint">.so</code>, <code class="prettyprint">.not</code> and possibly others crash on <code class="prettyprint">Seq:U</code>; need to evaluate entire codebase to see where <code class="prettyprint">only</code>
methods are used instead of <code class="prettyprint">multi</code>, preventing dispatch to <code class="prettyprint">Mu.*</code> candidates. (<a href="https://rt.perl.org/Ticket/Display.html?id=130867">reported as RT#130867
</a>)</li>
<li>Incorrect line number reported for wrong routine call when unpacking/heredocs are used (<a href="https://rt.perl.org/Ticket/Display.html?id=130862#ticket-history">reported
as RT#130862</a>)</li>
<li>Warnings produced by core code marked as "shouldn't happen" (<a href="https://rt.perl.org/Ticket/Display.html?id=130857#ticket-history">reported as
RT#130857</a>)</li>
</ul>


