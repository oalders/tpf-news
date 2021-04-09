
---
title: "Grant Proposal: Complete YAML::PP"
author: Coke
type: post
date: 2017-08-07 00:00:00 +0000 UTC
url: "/post/grant_proposal_complete_yamlpp"
categories:
 - Grants
 - Sign in

---

<p>
The Grants Committee has received the following grant proposal for the July/August round.
Before the Committee members vote, we would like to solicit
feedback from the Perl community on the proposal.
<p>
Review the proposal below and please comment here by
August 14th, 2017. The Committee members will start the voting process following that and the conclusion will be
 announced approximately one week after public comments.

<p>

<h2> Complete YAML::PP</h2>
<dl>

<dt id="Name">Name:</dt>
<dd>

<p>Tina Müller</p>
</dl>

<dl>
</dd>
<dt id="Amount-Requested">Amount Requested:</dt>
<dd>

<p>USD 2,500</p>

</dd>
</dl>

<h2 id="Synopsis">Synopsis</h2>

<p>I have been working on a new YAML Pure Perl Parser, already on CPAN as <a href="https://metacpan.org/pod/YAML::PP">YAML::PP</a>. It aims to parse <a href="http://yaml.org/spec/1.2/spec.html">YAML 1.2</a>.</p>

<p>The existing YAML frameworks in Perl all lack important features and don&#39;t support YAML 1.2. I will continue development of YAML::PP so that it&#39;s able to parse all valid syntaxes (with some minor exceptions). I will complete the Loader to support tags. I will add a Dumper and Emitter.</p>

<p>I will add test cases to the cross framework <a href="https://github.com/yaml/yaml-test-suite">YAML Test Suite</a> and continue developing the <a href="http://matrix.yaml.io/">YAML Test Matrix</a> to compare all frameworks.</p>

<h2 id="Benefits-to-the-Perl-Community">Benefits to the Perl Community</h2>

<p>While JSON has become popular as a simple format to exchange data, there are still a lot of use cases for YAML. Imagine how verbose an Ansible Playbook would look in JSON. Comments are an important feature, and Aliases come in very handy sometimes.</p>

<p>While <a href="http://pyyaml.org/">PyYAML</a> still only supports YAML 1.1, it has at least support for safe loading (See below). <a href="https://pypi.python.org/pypi/ruamel.yaml">Python Ruamel</a> aims to support 1.2.</p>

<p>I find it very unfortunate for Perl, that there is no support for YAML 1.2, for safe loading and for booleans.</p>

<p>It would be a nice opportunity for Perl to have a framework that supports all that.</p>

<p>Since the YAML Test Suite is supposed to become the number one source to write tests in any language, it can promote the new Perl framework.</p>

<p>Since I&#39;m aiming for a portable implementation, this framework might also be easily ported to Perl 6, which currently has no full support for YAML, although there is some development going on.</p>

<h2 id="Project-Details">Project Details</h2>

<p>The current state of YAML in Perl is as follows:</p>

<dl>

<dt id="YAML.pm">YAML.pm</dt>
<dd>

<p>Based on YAML 1.0. It can&#39;t do trailing comments and has problems with a lot of valid 1.1 and 1.2 syntaxes.</p>

</dd>
<dt id="YAML::XS">YAML::XS</dt>
<dd>

<p>Based on libyaml and the most recommended module. It supports YAML 1.1. It diverges from the spec for several edge cases.</p>

</dd>
<dt id="YAML::Syck">YAML::Syck</dt>
<dd>

<p>Supports YAML 1.0. It has problems with a lot of valid YAML 1.1 and 1.2 syntaxes.</p>

</dd>
<dt id="Safe-Loading">Safe Loading</dt>
<dd>

<p>YAML.pm and YAML::XS have no possibility to disable loading into objects. That means if you load an untrusted YAML file, it can be a security hole. YAML::Syck supports disabling that via &quot;LoadBlessed&quot;.</p>

</dd>
<dt id="Booleans">Booleans</dt>
<dd>

<p>The three mentioned modules don&#39;t support booleans. If you need to dump your data into JSON or let it be validated, booleans get lost (turned into 1 or 0). Only YAML::XS provides a limited way of keeping booleans when roundtripping.</p>

</dd>
<dt id="Separate-Parser-and-Constructor">Separate Parser and Constructor</dt>
<dd>

<p>The mentioned modules can only be used as complete Loaders. There is no possibility to put your own Loader on top of a parser.</p>

<p>You can check which test cases these modules are passing or failing: <a href="http://matrix.yaml.io/">YAML Test Matrix</a></p>

</dd>
</dl>

<p>I have been going over a number of RT tickets for YAML.pm at the end of 2016, creating and merging Pull Requests from patches and writing Pull Requests myself.</p>

<p>I&#39;m working a lot with Ingy döt Net, one of the creators of YAML, and Felix Krause, developer of NimYAML, on the YAML Test Suite and on RFCs for creating YAML 1.3.</p>

<p>I created the YAML Test Matrix to show the results of the tests for a growing number of YAML frameworks, based on Ingy&#39;s Docker image for <a href="https://github.com/yaml/yaml-editor">YAML Editor</a>.</p>

<p>I started to implement my own parser YAML::PP in 2017, and it currently passes most of the tests with the exception of Flow Style. The loader can already load YAML documents that the parser can parse. It supports booleans and aliases, but no tags yet.</p>

<p>I&#39;m currently transforming it into a tokenizer which allows correct syntax highlighting, making it also easier to spot errors.</p>

<p>I want it to be able to do roundtrips including comments at some point.</p>

<p>At the Perl Toolchain Summit 2017 in Lyon I have been working together with Ingy to create a concept of a new API for YAML loading. The goal is to integrate YAML::PP into that API.</p>

<p>Ingy and I started to implement the API in <a href="https://metacpan.org/pod/YAML::Perl">YAML::Perl</a>, using YAML::PP as a backend.</p>

<p>I also started to implement the new Loader API in <a href="https://github.com/yaml/yaml-perl6">Perl 6</a>, currently using the <a href="https://github.com/yaml/yaml-libyaml-perl6">libyaml binding</a> originally written by Curt Tilmes as a backend.</p>

<h2 id="Deliverables-and-Inch-Stones">Deliverables and Inch-Stones</h2>

<dl>

<dt id="Complete-YAML::PP::Parser">Complete YAML::PP::Parser</dt>
<dd>

<p>A couple of features are still missing from the parser</p>

<dl>

<dt id="Flow-Style">Flow Style</dt>
<dd>

<p>This is the biggest part. Flow Style is not indent based, and some rules are different than in block style. (I estimate 40h.)</p>

</dd>
<dt id="Flow-Nodes-as-mapping-keys">Flow Nodes as mapping keys</dt>
<dd>

<p>This is also a major part, because stacking of parser events is necessary until the parser knows if it&#39;s a mapping key or a node. (30h)</p>

</dd>
<dt id="Line-and-Column-Numbers-for-error-messages">Line and Column Numbers for error messages</dt>
<dd>

<p>Currently no information about line and column is saved. (20h)</p>

</dd>
</dl>

</dd>
<dt id="YAML::PP::Loader-Constructor">YAML::PP::Loader/Constructor</dt>
<dd>

<dl>

<dt id="Implement-loading-of-Tags-and-blessing-into-objects">Implement loading of Tags and blessing into objects</dt>
<dd>

<p>(20h)</p>

</dd>
<dt id="Provide-a-possibility-for-safe-loading">Provide a possibility for safe loading</dt>
<dd>

<p>(10h)</p>

</dd>
<dt id="Ideally-provide-a-way-to-only-load-certain-tags">Ideally provide a way to only load certain tags</dt>
<dd>

</dd>
</dl>

</dd>
<dt id="Write-YAML::PP::Emitter">Write YAML::PP::Emitter</dt>
<dd>

<p>(20h)</p>

</dd>
<dt id="Write-YAML::PP::Dumper-Deconstructor">Write YAML::PP::Dumper/Deconstructor</dt>
<dd>

<p>(20h)</p>

</dd>
<dt id="Add-more-test-cases-to-YAML-Test-Suite">Add more test cases to YAML Test Suite</dt>
<dd>

</dd>
<dt id="Show-also-results-of-invalid-examples-in-YAML-Test-Matrix">Show also results of invalid examples in YAML Test Matrix</dt>
<dd>

<p>(10h)</p>

</dd>
<dt id="Make-the-code-integrateable-into-the-new-YAML-Loader-API">Make the code integrateable into the new YAML Loader API</dt>
<dd>

</dd>
<dt id="Keep-in-touch-with-the-development-of-YAML-1.3-specification">Keep in touch with the development of YAML 1.3 specification</dt>
<dd>

<p>Implement the current parser in a way that makes it easy to add support for YAML 1.3</p>

</dd>
<dt id="Talk-about-this-project-at-TPC-in-Amsterdam">Talk about this project at TPC in Amsterdam</dt>
<dd>

<p>My talk and my published slides will explain why YAML currently is difficult to implement. I also gave this talk at the German Perl Workshop in Hamburg.</p>

</dd>
</dl>

<h2 id="Project-Schedule">Project Schedule</h2>

<p>I can start to work on this immediately and almost full time over the next two months.</p>

<h2 id="Completeness-Criteria">Completeness Criteria</h2>

<p>I release YAML::PP with the features implemented I mentioned above. The parser shall pass most of the tests in YAML Test Suite, with the exception of edge cases. Since the spec is often not very clear, there are some cases where it is unclear what should be the correct behavior, or what behavior actually makes sense. These edge cases are usually not relevant for real use cases and are easy to avoid. I will look at other frameworks and find out the most common behavior.</p>

<p>The Emitter should be able to transform every test input into valid YAML. The style (quotes/block scalar, spaces/newlines etc.) might still differ from the test suite.</p>

<p>The Loader/Dumper API, and especially the Parser and Emitter API, might not be completely fixed at the end of this grant. Ingy can me help me out here, supposed he&#39;s got time, and I need potential user feedback.</p>

<p>Ingy also offered to review the work.</p>

<p>I appreciate new test cases, bug reports, patches and co-maintainers, and I want to keep maintaining this module in the future.</p>

<h2 id="Bio">Bio</h2>

<p>I wrote my first Perl code in 1998 and have been in touch with the Perl Community since about 2001.</p>

<p>I already have two parsing modules on CPAN.</p>

<p>One is HTML::Template::Compiled, one of the fastest (and still feature rich) pure perl templating modules that gains its speed from compiling to perl code.</p>

<p>The other is Parse::BBCode, which is unique among the Perl BBCode modules, in that it provides a parse tree, it allows addition of own tags, it tries to correct invalid BBCode instead of simply dying, and it&#39;s fast.</p>

<p>YAML is a bit more complicated to parse, because it&#39;s indentation based, but I like solving programming puzzles.</p>

<p>I do a lot of pair programming with Ingy and I&#39;m also in contact with Felix Krause, so I have two people available who know the Spec.</p>

<h2 id="YAML-Details">YAML Details</h2>

<p>If you are wondering about terminology, here is a short explanation:</p>

<p>Loading YAML can be divided into two steps.</p>

<p>The Parser parses a Stream and returns a list of parsing events. The Constructor then takes these events, decides about numbers, tags, booleans and aliases/anchors and constructs a data structure.</p>

<p>Vice versa, Dumping YAML can be divided into deconstructing and emitting. The Deconstructor creates a list of emitter events from a data structure. The Emitter creates a YAML Stream from these events.</p>

<p>If you keep these things separate, it allows you to use the language independent Test Suite to test your parser. It also makes debugging and maintaining easier. Also you can use a different parser backend, for example a libyaml based one.</p>


</body>

</html>




