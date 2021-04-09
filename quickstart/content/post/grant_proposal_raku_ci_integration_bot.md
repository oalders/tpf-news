
---
title: "Grant Proposal: Raku CI integration bot"
author: Jason A. Crome
type: post
date: 2021-01-28 00:00:00 +0000 UTC
url: "/post/grant_proposal_raku_ci_integration_bot"
categories:
 - Grants
 - Sign in

---

## Synopsis
Implement a software, named Raku CI Bot (RCB), which can orchestrate the testing process of the Rakudo, NQP and MoarVM projects.

## Author
Patrick Böker

## Why
The three Rakudo core projects named above have a longish history of troubles with its testing infrastructure. Public and free CI platforms often inhibit one of several problems:

* Imposing limits on CPU or wallclock time
* Not supporting all needed platforms
* Unreliability
* Usability 

The proposed software is designed to solve these problems:
* By polling GitHub and the CI platforms in addition to listening to a hook, changes a guaranteed to not get lost.
* By interfacing with both AzureCI and Open Build Service, we get access to a very wide range of platforms.
* By implementing several command words usable in PR comments, a [PR-only workflow](https://github.com/Raku/problem-solving/pull/219/files) becomes a lot easier to adhere to.
* By persisting sources, build logs and artifacts test failures are easier to diagnose.

## Milestones
- Setup
    - Set up a stack with Cro and Postgres and put it in a container.
    - Get it up and running on a publicly accessible temporary test server.
- Pan out the core model and class interfaces
- Extend [WebService::GitHub](https://github.com/fayland/perl6-WebService-GitHub) to cover the [Checks](https://docs.github.com/en/free-pro-team@latest/rest/reference/checks), [Pulls](https://docs.github.com/en/free-pro-team@latest/rest/reference/pulls) and [Commits](https://docs.github.com/en/free-pro-team@latest/rest/reference/repos#commits) APIs.
- A `Webservice::OBS` module providing access to the relevant bits of the OBS API.
- A `Webservice::AzurePipelines` module providing access to the [Build](https://docs.microsoft.com/en-us/rest/api/azure/devops/build/?view=azure-devops-rest-6.0) API.
- Implement the core model including tests and DB serialization
- Implement logic to listen to GitHub hooks and poll filling a persisted work queue.
- Azure CI:
    - A worker that pushes test jobs to the platform
    - Listen to a [platform hook](https://docs.microsoft.com/en-us/azure/devops/service-hooks/services/webhooks?view=azure-devops)
    - Poll to retrieve test results
- Implement logic to report a test status back to GitHub.
- OBS CI:
    - Set up an OBS project to build Rakudo. This can build on the previous work of niner.
    - A worker that pushes test jobs to the platform
    - Listen to a platform hook (I probably have to implement the notification part on the OBS site as well, as OBS seems to provide no hooks for build completion.)
    - Poll to retrieve test results
- A website
    - Basic setup
    - An index that lists build jobs
    - Build job page
        - Build state and result
        - Build logs
        - Source and artifact downloads
        - Link to previous tests (in case of re-runs)
        - Link to AzureCI, OBS, GitHub
    - Add some styling

## Communication
I plan to write status update posts whenever there is relevant progress. I estimate there will be 4 - 8 posts overall. The posts will be published [here](https://dev.to/patrickbkr/). I intend the posts to have a focus on explaining the software so the posts can double as documentation.

## Schedule
I estimate the project to be about 80 hours of work. I can allocate *at least* 8 hours per week to work on this. This would mean I finish the project in two and a half months. These are estimates and unforseen difficulties can push the timeline.

## Requested amount
I request $4000 USD.

82 h * 40 € = 3,280 € ~= $4000

## Biography
I have been involved in Raku development since 2018 with a focus on integration and tooling. I:

- Designed and implemented Rakudos relocatability functionality
- Set up automated building of precompiled relocatable packages for Windows, Linux and MacOS. First on CircleCI then on AzureCI.
- Migrated the core projects CI infrastructure to AzureCI (shutdown of Travis, CircleCI and AppVeyor is still to be done)
- Gave the rakudo.org website an overhaul
- Took over the rakubrew project (then named Rakudobrew) refactored it largely and implemented many new features, among them
    - Windows support
    - Utilizing the precompiled archives served on rakudo.org
    - A new website
    - Installation via a self-contained executable
- Fixed many, many papercut bugs in a wide range of modules, documentation and some in the compiler stack itself
- Have utilized Raku in my dayjob in multiple medium sized projects
- Wrote a small program in Perl which remotely monitors a MediaWiki website for changes, runs a validator and feeds the validation results back to the site. The functionality is in several aspects similar to the proposed project.

## Open questions

Will we be able to persist the sources and artifacts of every build job in the long run? Depending on how many platforms we want to test on the build artifacts for a single test run could end up in the range of half a gigabyte. Some clever compression might prove to be helpful here.

## Implementation details
RCB acts as an intermediate between changes in GitHub and the CI services. Initially supported CI services will be Azure and Open Build Service (OBS).

RCB will monitor GitHub for three event types:

- Commits to a master branch
- A new PR or a commit in a PR
- A comment in a PR which includes a command word

A change typically triggers a CI run. The following steps happen:
- RCB retrieves the sources off of GitHub and creates a source tarball compatible with our source release files.
- RCB triggers CI services to test this source tarball.
- The CIs finish their test and report back to RCB.
- RCB retrieves and saves build logs from the CI backends. This is especially important for re-tests, because on some CI platforms doing a re-test makes the original build log inaccessible.
- RCB retrieves and saves build artifacts from the CI backends.
- RCB reports the results back to GitHub as Status notifications.


## Comment triggers
In addition to commit and PR events RCB will scan PR comments for command words.
- `{merge on success}` will cause RCB to automatically merge the PR should the CI tests be successful. If the tests are unsuccessful RCB will add a comment stating that automatic merging did not happen.
  There will be a check for the permissions of the person writing the comment whether the person has merge permissions.
- `{re-test}` will cause RCB to run the CI for the respective PR again.

## Flapper detection
RCB will scan failed CI build logs for known flappers. If a flapper is identified as the only failure the test is re-run automatically once.

## Branch matching
When preparing a CI run for a PR in either the `rakudo/rakudo`, `Raku/nqp` or `MoarVM/MoarVM` repo, RCB will check if a PR with a matching name exists in any of the other repos and use that PR instead of the master branch for testing.

If no such PR is found, RCB will follow a commit matching logic.
- Commits in `rakudo/rakudo` will use the NQP version given in `tools/templates/NQP_REVISION` and MoarVM commit given in NQPs `tools/templates/MOAR_REVISION`.
- Commits in `Raku/nqp` will use the MoarVM commit given in `tools/templates/MOAR_REVISION` and Rakudo master.
- Commits in `MoarVM/MoarVM` will use NQP master and Rakudo master.

This commit matching logic is already in use in the Azure CI pipeline. (The branch matching logic isn't though.)

## Website
RCB will serve a website that provides a list of all CI runs. Each run will provide the following information:
- Link to the GitHub page that triggered the run
- The build source archive
- Links to previous runs for the same source
- For each CI backend
    - Link to the respective CI page
    - The build log
    - The build artifacts
- A button to retrigger the CI run

## Third party API triggers
RCB will interface with several third party APIs that provide push triggers. Namely GitHub, OBS and Azure. Push triggers have proven to not be entirely reliable in the past. As a counter measure RCB will not only listen to the push triggers, but also poll the external APIs periodically. This way RCB will be:
- Responsive (reaction time of a few seconds) by directly reacting to push triggers.
- Reliable by internally keeping a list of all events in the third party and keeping that list in sync by polling.
