
---
title: "Grant Proposal: MoarVM JIT Compiler Expression Backend Maturation"
author: Coke
type: post
date: 2018-10-10 00:00:00 +0000 UTC
url: "/post/grant_proposal_moarvm_jit_comp"
categories:
 - Grants
 - Sign in

---

The Grants Committee has received the following grant proposal for the Sep/Oct round.
Before the Committee members vote, we would like to solicit feedback from the Perl
community on the proposal.
<p>
Review the proposal below and please comment here by October 17th, 2018.
The Committee members will start the voting process following that.
<p>
<strong>MoarVM JIT Compiler Expression Backend Maturation
</strong>
<p>
<strong> Name:
</strong><p>
    Bart Wiegmans
<p>
<strong> Amount Requested:
</strong><p>
    USD 7,000.

<strong>Synopsis
</strong>
<p>
August 2017 saw the completion of the new ‘expression’ JIT backend. This backend is still embedded in the original JIT compiler backend. This has allowed the expression JIT to be incrementally developed rather than as an immediate replacement. Compared to the legacy JIT, the expression JIT relies less on specialized assembly language and is able to do more in C. It also allows contributors to add support for MoarVM instructions with a simple template language.
<p>
By at least one metric, contributions, the expression JIT backend is a success. 10 different contributors besides myself have now committed changes to the expression templates, compared to just 4 for the original assembly code templates. Debugging problems with the expression JIT also tends to be simpler, which I attribute to improved tool support and consistency checks in the template precompiler. Because of this, and to reduce the long-term maintenance burden, I think it is a good idea to move towards deprecating the legacy JIT compiler backend in favor of the new expression JIT.
<p>
There are several areas that require work in order to make this a reality, which is the reason for this grant request. The focus of this grant request will be infrastructure improvements, rather than adding template support for multiple opcodes. Because several contributors have already demonstrated aptitude at adding expression templates I feel that this is the best way to direct my efforts.
<p>
<strong>Benefits to the Perl Community
</strong>
<p>
Depending on the nature of the code, the MoarVM JIT compiler can enable a significant performance boost to Perl6 programs. However the JIT compiler is limited to the x86-64 backend. A primary result of this grant work will be a reduction in volume of specialized x86-64 assembly in the JIT compiler. Based on the experience of the past few years, it is my expectation that this will improve the maintainability of the JIT compiler. It is my hope that it may one day enable a port to other architectures.
<p>
<strong>Deliverables
</strong>
<p>
As result of this grant request, I plan to deliver the following improvements to the MoarVM JIT compiler:
<p>
<ul>
<li>Floating point support in the expression JIT backend.
<li>Code quality improvements (optimizations) based on an expression-tree rewriting optimizer, specifically:
<ul><li>Completing the work on the jit-expr-optimizer branch, including several common optimizations.
<li>Implement optimizations in the register allocator.</ul>
<li>Replaced specialized code in the legacy JIT with portable expression tree structures.
<li>Improved handling of irregular instructions by the register allocator.
<li>Further, I will write at least 3 reports, published on my public blog (http://brrt-to-the-future.blogspot.nl/) on the progress 
</ul>
<p>
<strong>Project details and proposed schedule
</strong>
<p>
I plan to perform this work over at most a three-month period and can start working as soon as the grant is accepted. This work will be implemented as one or more feature branches and will be considered completed when these branches have been merged into the master branch of MoarVM.
<p>
<strong>Supporting floating point operations
</strong>
<p>
The legacy JIT compiler supports floating point operations, while the expression JIT does not yet support these. Floating point operations on x86-64 require the use of a separate set of registers (SSE registers). Passing floating point arguments to functions also requires the use of SSE registers. Currently, the register allocator for the expression JIT supports only general purpose (integer and pointer) registers. 
<p>
Implementing SSE register support for the expression JIT requires work on the register allocator to support multiple register sets. Furthermore, it will require modification of the DynASM assembly library to fix a code generation bug in selecting SSE registers. Finally, there is a small number of additional expression operators and tiles that need to be added.
<p>
I expect this to be worth approximately 40 hours of work, of which 20 will be spent on the instruction encoding (DynASM) and 20 on the register allocator itself.
<p>
<strong>Improving the quality of generated code
</strong>
<p>
In some cases the JIT can compile code that is worse than the equivalent assembly code written by hand. This is not unexpected, but there are some relatively simple optimizations that can be applied to improve the quality of generated code. There is a branch in development (jit-expr-optimizer) that can rewrite expression trees to more efficient forms.
<p>
Aside from tree optimizations, there are several optimizations possible in the register allocator (such as copying constant values rather than spilling them to memory). Such optimizations are also part of the deliverables.
<p>
I estimate this to take approximately 40 hours of work as well.
<p>
Implementation of specialized structures as expression trees
Several constructs in the legacy JIT are implemented with specialized assembly language support. These can be implemented as expression trees which require no special support, which will improve portability and reduce the maintenance burden. Of special interest are:
<p>
<ul>
<li><strong>Spesh guards</strong> - currently the expression JIT has to break off whenever a guard is encountered, preventing it from reaching code with the most potential for optimization. Supporting spesh guards requires specific handling in the expression JIT, to ensure that the JIT is consistent with the interpreter in case a deoptimization is needed.
<li><strong>User subroutine invocation</strong> - The lego JIT backend contains specialized assembly code for the support of preparing and invoking subroutine. This used to be necessary because JIT compiled code needs to return control to the interpreter before entering the new code. After the jit-stack-walker branch was merged, this specialized code is no longer required. It can be replaced with a specialized expression tree, which will be (at the very least) more portable and maintainable.
<li><strong>NativeCall</strong> - on x86-64, with JIT enabled, NativeCall is implemented by the JIT compiler, and NativeCall calls can be inlined into JIT compiled code. The current implementation builds legacy JIT nodes, but it will be simpler to build this as an expression tree, and this will allow for more efficient code as well.
</ul>
Because of the focus on infrastructure, the replacement of ‘primitive’ operators (that are directly implemented as assembly code fragments in the legacy JIT) is not a priority, except insofar as to remove existing limitations and add capabilities.
<p>
I estimate this to take approximately 40 hours of work.
<p>
<strong>Improved handling of irregular instructions by the register allocator
</strong>
<p>
Most x86-64 instructions can support operands from any (pair of) register(s). However, there are many irregular instructions as well, which read (and overwrite) values from specific registers only. The register allocator for the expression backend does not yet handle the requirements for such irregular instructions. An example of irregular instructions in x86-64 is ‘div’, which implements integer division, a relatively common operation.
<p>
Satisfying such register requirements may require complex value shuffling sequences. Fortunately the code to implement this already exists in the form of C function call support, which has even more extensive requirements. I intend to generalize this code in order to support irregular instructions.
<p>
I estimate this to take 20 hours of work.
<p>
<strong>Biography</strong>
<p>
I’m known in the perl community as ‘brrt’. I joined during the perl6 community during the Google Summer of Code program of 2014, when I implemented the original JIT compiler backend for MoarVM. I’ve maintained this ever since. In 2015, I requested a grant from the perl foundation to support the development of the advanced ‘expression’ JIT compiler backend, which was successfully completed (albeit much later than initially expected). I’ve continued to support the JIT, debugging issues and supporting contributors by code reviews and advice on the #moarvm IRC channel.
<p>
I gave presentations about my work at YAPC::EU 2015 in Granada, at FOSDEM 2016 in Brussels and at TPC in Amsterdam in 2017. I maintain a blog about MoarVM development at http://brrt-to-the-future.blogspot.com/, which is also where I intend to publish progress reports.
<p>
Prior to that I had implemented an Apache HTTPD module for the Parrot VM, during the Google Summer of Code program of 2012. (Unfortunately, Apache HTTPD modules would cease to be a common deployment method soon after, and we never got Rakudo Perl 6 to work within this module).
<p>

