
---
title: "Grant Proposal: App::rs - The first reference counting CPAN client"
author: Coke
type: post
date: 2018-06-11 00:00:00 +0000 UTC
url: "/post/grant_proposal_apprs_-_the_fir"
categories:
 - Grants

---

The Grants Committee has received the following grant proposal for the May/June round.
Before the Committee members vote, we would like to solicit feedback from the Perl community on the proposal.

Review the proposal below and please comment here by June 16th, 2018.
The Committee members will start the voting process following that and the conclusion will be announced the
last week of June.

# App::rs - The first reference counting CPAN client

- Name:

    Yang Bo

- Amount Requested:

    USD $5,000

## Synopsis

When a CPAN client installs a module, it first installs the dependency
of this module, and potentially installs its dependency's dependency,
etc., that's a required and well known behavior, and every CPAN client
does this. However, the reverse is not true, i.e. when a CPAN client
uninstalls a module, it first uninstalls the dependency of this module,
and potentially uninstalls its dependency's dependency, etc., `App::rs`
is the first CPAN client that's known to do both installation and uninstallation
recursively. I use the term **reference counting** since that's the same
thing `perl` uses itself to do garbage collection, i.e. when a variable
goes out of scope it will not be freed if anything still refers to it, it will only
be freed when there's nothing refers to it anymore, and then the things it refers to
will be freed, and potentially what it refers to refers to, etc. `App::rs`
does exactly the same thing as `perl`, except not for Perl variables, but for
Perl modules, hence the name _The first reference counting CPAN client_.
And it comes with much more additional benefits.



## Benefits to the Perl Community

CPAN is considered by many Perl users to be its greatest feature, which naturally
makes `App::rs` a beneficial and worth doing project, since it enables one to
fully explore CPAN without any hesitation, you will never have to worry about
how to remove a module along with its dependency it pulled in; never be afraid to upgrade
a module to a new version since you always have the option to downgrade to a
previous version from a binary package instantly; never feel that inside the
perl module directory there lies a whole mess since each file is tracked, you know
exactly to which module a file belongs, why a module is installed, and whether
and when it will be gone.

In addition to that, you can also upgrade a module including its dependency
recursively, transfer the reference counting database and the binary package pool to
another machine and have a fully controllable Perl library restored with no effort.

In summary, `App::rs` is worth it by making the best of Perl even better.

## Deliverables

1. A CPAN client that does reference counting, i.e. when being used to uninstall
a module, this module will be removed if and only if there's no other module
refers to it, and when it's being removed the dependency of it will also be
removed, potentially its dependency's dependency, etc.
2. A CPAN client that's capable to upgrade a module, including its dependency,
its dependency's dependency, etc., i.e., recursively.
3. A CPAN client that's capable to serialize every module it installs as a
binary package, so that removed modules could easily be restored, multiple versions
of the same module could be switched, and it's all done instantly, since
no re-downloading, re-building, re-testing is required.
4. A CPAN client that's able to filter the binary package pool and reference
counting database according to the user's selection of a subset of the
installed modules, so that no matter how many modules are installed they
could always be partitioned, divided, isolated according to what a user
requested, and this filtered package pool and database could then be
transferred to another machine and instantly form a new Perl library that is
fully controllable.

## Project Details

`App::rs` is originally used only for my custom Linux distribution, it uses
a very different approach to do package management, it calculates the difference
between the previous and current state of the root directory it manages, and
serialize this difference as a binary package that could be easily removed,
restored, transferred, it's a general method that does not rely on any
package specific quirk and works with any installation process without exception, so needless
to say that includes Perl, naturally, I thought about extending it as a CPAN
client.

Unlike a Linux package and thanks to `metacpan-api`, there's an easy, standard,
and universal way to enumerate the dependency of a Perl module, so I thought
it really would be a waste if I were not to take full advantage of this, I
tried to build the install and uninstall process for CPAN both recursively,
with the reference counting database as an additional data structure, and
it turns out `App::rs` fits perfectly as a CPAN client, since it already provides a solid
base for single package managing, I've been using it through recent years and the
code is quite stable, and connecting each single package together through a
reference counting database is a relatively easy task, only a small amount of
code is required in the core logic (most of the code is about networking).

Right now `App::rs` is already usable, the first three of the deliverables
are basically done, the problem of it is that it's not thoroughly tested, and
there's not much meaningful error reporting, you currently have to understand
its guts to know what's going on when something didn't go smoothly, and
it definitely could use some more documentation. But given time, I would
say `App::rs` will be a guaranteed success since the first, the most fundamental,
and the most difficult step is already done.

## Inch-stones

- Right now JSON is used as the reference counting database, it's not
a problem now since it's still small, but I'm unsure when thousands of
packages are installed whether it will become too big to be efficient,
I will inspect this more and switch to an alternative as the database
format.
- Currently everything is possible in App::rs, but almost nothing is easy,
it's required to carefully read the mannual and understand what each
configuration file does, how they're chained together before taking effective
use of it, I will work towards the Perl spirit, i.e. easy things should be
easy and hard things should be possible, by making sane defaults, providing
more and friendlier warning/error messages, etc.
- By default `App::rs` uses HTTP to get information from metacpan, HTTPS is
supported through `socat`, depending on the community feedback I may
use `openssl s_client` or `gnutls_cli` as well, or switch to a more
standard way.
- I didn't use any platform specific functionalities when I wrote `App::rs`,
but that being said I also only had Linux in mind since that's the only
OS I use, I would like to test it on more platforms and ensure it's
available on as many environments as possible.
- At the moment `App::rs` only installs dependency of the `requires`
relationship on the `configure/build/test/runtime` phase, I will change
`App::rs` so that it offers choice on whether an optional dependency
should be installed, by command line switch or an interactive dialog.
- I would like to add the ability to show the dependency relationship as
a graph, either in TUI or GUI, althrough you could view the dependency details
as JSON in the reference counting database, it's not comparable to the
clarity a graph provides, suppose A depends on B which depends on C
which depends on D, you will have a hard time figuring out that D is
installed because of C which is turn installed because of B which is
in turn installed because of A, and finally you know why D is there
all time along, as you can see it's tedious and difficult to figure
this out all through a big text database, so a graphical representation
would provide much more convenience and insight.
- Currently there's no way to avoid tests, they will always be run and
installation will fail if any of the test fails, it may be desired to
avoid them in some cases, I would like to provide a command line option
to provide that behavior.
- When installation fails due to configure/build/test failure, anything installed
during this session will be not be updated in either the reference counting
database or the metadata database of the root directory, so that any changes
made could be easily reverted, this may be the desired behavior, since
a failure like this is usually fatal, like your perl is too old,
you're missing a C compiler, some vital libraries are missing, yet
sometimes the failure may just be fixable as well, I would like to provide
more choices when a failure happens, like saving the database, launching
a shell so that the user could fix things and re-launch the building
process.
- When fetching info from metacpan `App::rs` will only ask for the latest
release, I would like to provide a way for the user to specify any valid version
to be installed.
- Switching between multiple versions of a module is now done by mannually
uninstalling and then edit the reference counting database and then
reinstalling, I will write a command that specifically does this, so the
user doesn't have to know that much and follow all these steps to switch
to a previous version.
- Each of the compiled modules will be preserved as a binary package so that
the next time any of them is required it will be restored from the binary
package, but there's currently no way to throw them away automatically,
I think it would a good idea to provide a way to clean up those modules
that aren't referenced by any, both the binary packages of them and their
entries in the reference counting database.
- Change the `uninstall` command so that it could do dry uninstallation,
so that the user could have a clear view of what module will be uninstalled,
and make sure nothing will be removed by accident.

## Project Schedule

This project will take approximately six months, although it's already
usable for myself now, it's really different and still far from easily usable
by everybody, much time and patience are required for crossing platform,
user friendliness, and more detailed documentation.

Since I'm not under employment at the moment, I can begin work on `App::rs`
immediately after the approval of this proposal.

## Completeness Criteria

Since reference counting is a precise concept, all the deliverables
are thus well defined, so completeness of this grant could be evaluated
by whether a new release of `App::rs` is uploaded to CPAN with
all the deliverables correctly implemented.

## Bio

I created my own Linux distribution with `App::rs` as the package manager
at the end of 2015, it's the first package manager known to use the
packaging by difference method, which is general enough to be used
with any installation process. Due to my expertise on this field,
I extended `App::rs` to be the first reference counting CPAN client in
less than a week.

I'm the best person to work on this project since I created it from
scratch, I wrote every line of it and I know every bit of it by heart,
I use `App::rs` to manage RSLinux (my own Linux distribution) on my laptop,
desktop, KVM based VPS, OpenVZ based VPS, and package management is
exactly what a CPAN client does at its essential, I know this for a
fact since I found out how perfect a foundation `App::rs` is and how
easily a reference counting CPAN client could sit on top of it myself
when I was extending it, which fits perfectly with my experience and
justifies me as the best candidate.


