
---
title: "Maintaining the Perl 5 Core: February 2017 report"
author: Makoto Nozaki
type: post
date: 2017-03-13 00:00:00 +0000 UTC
url: "/post/maintaining_the_perl_5_core_fe"
categories:
 - Grants
 - Perl 5 Development

---

This is a monthly report by Dave Mitchell on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.

<pre>
The main things I did last month were:

Firstly, fixing various issues with scopes in regexes. In particular,
(RT #126697), code blocks sometimes failed to undo localisations when
backtracking. For example the $s below wasn't always being restored when
the B part of the match failed and it backtracked to try another A - where
A represents something complex like (\w+\s+)* which can match in multiple
ways:

    /A(?{ local $s = ...})B/

As part of that work, non-greedy matching of complex sub-expressions with
captures and repeated backtracking was made more efficient under some
circumstances; for example the following now runs about 25% faster:

    $s = ("a" x 1000);
    $s =~ /^(?:(.)(.))*?[XY]/ for 1..10_000;

Secondly, improving 't/TEST -deparse'.

The -deparse option to t/TEST causes it to run all the core's test
scripts, but after running them through the deparser first. Many of these
modified scripts are currently known to fail, and there is an exclusion
file, Porting/deparse-skips.txt, which is supposed to list the known
failures. However, over time, new failures have appeared which are not in
the exclusion list. Last August I did some work on Deparse.pm and managed
to reduce some of the expected and unexpected failures, but since then
more failures have crept in.

My recent work includes: modifying t/TEST so that it distinguishes
between expected failures and unexpected passes, and warning of unknown
files in Porting/deparse-skips.txt; purging Porting/deparse-skips.txt to
account for files that have been renamed or are no longer in the core, and
to reflect the current state of things; and fixing Deparse.pm to:
    * better handle 'BEGIN { use_ok() }';
    * better handle 'BEGIN { require expr }' (as opposed to require Bareword);
    * deparse lexical var attributes, e.g. 'my $foo :bar(baz)';
    * avoid a 'deep recursion' warning;
    * handle an escaped literal tab char in a pattern, e.g
        /.....\ ..../x where the whitespace char following the backslash is
        actually a tab; previously the deparse failed to emit the backslash;
    * handle declarations with multiple lexical vars within a pattern code
      block, e.g. /(?{ my ($x, $y) = @a; })/

Because we're currently in code freeze, this as been pushed as
smoke-me/davem/deparse and will be merged in 5.27.x.

Thirdly, reviewing and fixing tickets in the security queue. There's quite
a lot of tickets in the security queue due to fuzzing, where if the fuzzer
detects a use-after-free or buffer overrun for example, the reporter
submits it to the security queue rather than the normal queue. Once
examined, 95% of the time it will be found to be harmless or
non-exploitable, but until someone has assessed and fixed it, it lingers
as an open security ticket.

</pre>


<pre>
SUMMARY:
     10:52 RT #126697 local() in embedded code in regex not working as expected
      0:09 RT #128528 XSLoader may load relative paths
      7:39 RT #129861 heap-use-after-free S_mro_gather_and_rename
      6:05 RT #129881 heap-buffer-overflow Perl_pad_sv
      2:11 RT #130321 heap-buffer-overflow Perl_vivify_ref (pp_hot.c:4362)
      0:04 RT #130332 double-free affecting multiple Perl versions
      0:36 RT #130336 attempting free on address which was not malloc()-ed
      0:10 RT #130344 heap-use-after-free S_gv_fetchmeth_internal
      1:46 RT #130569 heap-use-after-free in S_regmatch
      0:43 RT #130624 heap-use-after-free in Perl_sv_setpvn
      1:56 RT #130650 heap-use-after-free in S_free_codeblocks
      6:04 RT #130703 heap-buffer-overflow in Perl_pp_formline
      4:22 RT #130727 S_maybe_multideref: Assertion failed
      2:22 RT #13076