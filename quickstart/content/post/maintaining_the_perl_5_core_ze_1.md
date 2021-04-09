
---
title: "Maintaining the Perl 5 Core (Zefram): Grant Report for Dec 2017"
author: Makoto Nozaki
type: post
date: 2018-01-17 00:00:00 +0000 UTC
url: "/post/maintaining_the_perl_5_core_ze_1"
categories:
 - Grants
 - Perl 5 Development

---

This is a monthly report by Zefram on his grant under [Perl 5 Core Maintenance Fund](http://www.perlfoundation.org/perl_5_core_maintenance_fund). We thank the TPF sponsors to make this grant possible.

<pre>The hours that I have worked in 2017-12 pursuant to my TPF core
maintenance grant are as follows.

 22h40m  [perl #130851] [PATCH] socket SOCK_CLOEXEC
 21h58m  review tickets
 16h19m  review mail
 12h47m  smartmatch
</pre>


<pre>
 11h19m  [perl #125806] Perl segfaults in BEGIN, write to null
         pointer, separate bug
 10h52m  [perl #132425] Suggested warning on attempt to 'use' with
         arguments when no import() subroutine exists
  4h13m  [perl #132577] BBC Module::Install broken by
         0301e899536a22752f40481d8a1d141b7a7dda82
  3h46m  [perl #2754] [BUG] can't exit 0 from CHECK{}
  2h24m  [perl #130076] readline argument list is messed up
  2h19m  [perl #105920] Perl parser sometimes tolerates stray nulls,
         sometimes not
  2h12m  [perl #129888] null ptr deref, segfault in Perl_do_aexec5
         (doio.c:1595)
  2h12m  [perl #110056] installhtml uses absolute paths in links
  1h41m  [perl #115814] open ${\$x} leaks
  1h41m  [perl #110520] pod2html 1.12 & 1.13 broken
  1h39m  [perl #119831] Data::Dumper: Useqq should apply to glob
         names, too
  1h38m  [perl #92264] Freeing $a or $b during sort causes a double
         free
  1h38m  grammar token typing
  1h32m  perldelta
  1h32m  [perl #132234] use-of-uninitialized-value in Perl_upg_version
         (vutil.c:669)
  1h32m  [perl #132644] The 'each' function documentation is missing
         some information.
  1h25m  [perl #132142] Bleadperl v5.27.3-34-gf6107ca24b breaks
         MLEHMANN/AnyEvent-HTTP-2.23.tar.gz
  1h25m  [perl #132540] uninitialized variable and integer overrun in
         pp.c and toke.c
  1h25m  [perl #132651] commit d2e38af7 exhausts swap space on FreeBSD
  1h11m  [perl #126042] Segmentation fault in Perl_pp_multiply (and
         other functions)
  1h09m  [perl #132578] BBC List::MoreUtils::XS broken by
         16ada235c332e017667585e1a5a00ce43a31c529
  1h06m  [perl #119829] usemymalloc cannot handle long strings
  1h03m  [perl #132633] [Win32] 5.27.7 fails all tests.
  1h01m  [perl #119367] Another 32-bit residual in 64-bit perl 5.18
  1h00m  create tickets
  1h00m  [perl #74142] provide a better c wrapper example in perlsec
    59m  [perl #131894] runtime error: shift exponent -2 is negative
         (toke.c:10966:54)
    55m  [perl #4574] readpipe() broken: 2 bugs
    54m  [perl #113406] perldoc in 5.16.0 required groff upgrade but
         now misdisplays asterisks
    52m  [perl #104060] Writing to $> and friends hide errors
    49m  [perl #132648] Cwd: different return values between pure perl
         and XS variants
    45m  C++ dNOOP
    45m  [perl #130726] semicolons on *_DIAG_{IGNORE,RESTORE}
    44m  [perl #126150] Regex missing from perlref
    41m  perl_run() on Windows
    40m  [perl #132634] Strange bug in the new given/whereis/whereso
         construct (perl-5.27.7)
    38m  [perl #119623] GDBM_FILE: gdbm_open requires a blocksize to
         be a power of two
    37m  [perl #130958] inconsistent block/hash detection (again)
    37m  [perl #121105] During a system(), unquoted Perl vars are
         evaluated _after_ the fork() call
    37m  [perl #121372] warn broken when operand is overloaded
    34m  [perl #128261] Assert fail in Perl_sv_2pv_flags: 'sub
         MODIFY_HASH_ATTRIBUTES{}my(%o):s==0'
    33m  Devel::NYTProf BBC from fake-import change
    32m  [perl #113090] Data::Dumper -- undef not recognized as
         "false" in config booleans
    31m  [perl #132544] heap-buffer-overflow (WRITE of size 8) in
         Perl_pp_reverse
    30m  [perl #130578] op.c:10706: OP *Perl_ck_refassign(OP *):
         Assertion `left->op_type == OP_SREFGEN' failed
 