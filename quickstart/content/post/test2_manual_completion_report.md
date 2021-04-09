
---
title: "Test2 Manual Completion Report"
author: Alberto Simões
type: post
date: 2018-04-24 00:00:00 +0000 UTC
url: "/post/test2_manual_completion_report"
categories:
 - Grants

---

Follows the report by Chad Granum (aka EXODIST) for his grant: [Test2 Manual](https://metacpan.org/pod/release/EXODIST/Test2-Suite-0.000114/lib/Test2/Manual.pm).

# Test2 Manual grant complete

This is a completion report for the Test2 manual grant.

# Deliverables

- Test2::Manual

    Original description:

        A brief introduction and table of contents.

    Completed form:

        Jumping off point with a map of the manual layout/table of contents.

- Test2::Manual::Tooling

    Original description:

        This section will cover writing test tools. This would be a root document
        with several other page links.

    Complete form:

        The tooling page is the portal to all tooling tutorials. This section
        covers writing test tools.

- Test2::Manual::Maintenance / Test2::Manual::Anatomy

    Original Description:

        This section will cover maintaining Test2 itself. This will be a root
        document pointing to pages that explain the internals, and how they work
        together.

    Final Form:

        I renamed this section to Test2::Tools::Anatomy. This section covers docs
        about Test2 internals and implementation details. This section is more
        complete and useful than I originally intended.

# Completeness Criteria

Original criteria section from my proposal:

    The manual will be included in a Test2 release on cpan. The manual will
    have all expected sections, with no significant gaps or TODO sections.

Test2::Manual, with all the specified sections is now included in the
Test2-Suite distribution on cpan. There are no TODO's or gaps in the sections
that were part of my proposal, though there are TODO's in a third section that
went beyond the scope of my grant.

The version with the complete manual is Test2-Suite-0.000114.

# Bonus Material

I went beyond the scope of my original grant report and added the
Test2::Manual::Testing section. This section is still incomplete, but was not
part of my original grant proposal anyway, as such all the content that is
complete is a nice bonus.

I also added the [Test2::Manual::Concurrency](https://metacpan.org/pod/Test2::Manual::Concurrency) and
[Test2::Manual::Contributing](https://metacpan.org/pod/Test2::Manual::Contributing) pages.

# Revisiting the inch-stones

- Migrating from Test::Builder

    Original:

        Conversion notes and examples for people moving from Test::Builder. This
        will detail how tools might have changed, what is gone, and introduce some
        new or replacement concepts.

    This is done as [Test2::Manual::Tooling::TestBuilder](https://metacpan.org/pod/Test2::Manual::Tooling::TestBuilder).

- Simple OK tool

    Original:

        This is the most basic tool you can write, and is valuable for explaining
        the key concepts universal to all Test2 tools.

    This section is complete, [Test2::Manual::Tooling::FirstTool](https://metacpan.org/pod/Test2::Manual::Tooling::FirstTool).

- The Test2 API

    Original:

        This covers Test2::API, and all the functionality it exposes.

    This was not done as a single section. Instead there are handful of tutorials
    which show how to use some of the most common API functions.

- The 'Context' object

    Original:

        Explain what the context object is, why it is important, and how to use it
        effectively.

    This did not need its own section. The [Test2::Manual::Tooling::FirstTool](https://metacpan.org/pod/Test2::Manual::Tooling::FirstTool)
    covers all the important points about the context object from a tool writers
    perspective. Everything else is covered by the
    [Test2::Manual::Anatomy::Context](https://metacpan.org/pod/Test2::Manual::Anatomy::Context) document.

- The 'Hub' object and API

    Original:

        Explanation of the hub objects and how they work.

    This section is not necessary. Tool developers no longe