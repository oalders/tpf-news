
---
title: "(Updated)  Grant Proposal: General Perl OpenAPI Validator / Interpreter"
author: Jason A. Crome
type: post
date: 2020-10-08 00:00:00 +0000 UTC
url: "/post/updated_grant_proposal_general_perl_openapi_validator"
categories:
 - Grants
 - Sign in

---

(**Note: ** This is an updated of the [original grant request](https://news.perlfoundation.org/post/grant_proposal_general_perl_openapi_validator) - cromedome)
# Synopsis

Develop a clean and easy interface for managing OpenAPI 3.x schemas.

# Proposal

OpenAPI is a format that makes use of the JSON-Schema specification to communicate what a web API can do in a machine-readable way. The current module that is most active for JSON-Schema support in Perl is JSON::Validator.

The feeling in the community is that the JSON::Validator interface is difficult to work with. This is possibly because it covers a broader set of applications than solely OpenAPI, and thus doesn't provide what people might consider to be logical functions when their focus is an API.

The OpenAPI namespace will therefore contain functions specifically aimed at OpenAPI and the normal applications for it (for example, knowing that it will be used with HTTP and paths lets us write functions to use the module in those terms).

We propose to do this by

- Contributing to existing modules, chiefly JSON::Validator, which already has OpenAPI functionality at an experimental level but has known limitations.

- Use the OpenAPI namespace to create an interface for simple, user-friendly interactions with OpenAPI 3.1 schemas. This will include the following steps:
    - Specification and design to outline what the module will comprise
    - Producing the actual functionality for the module
    - Testing to ensure that the module works as expected

# Deliverables

- A satisfactory completion of the implementation of OpenAPI 3.x in JSON::Validator
    - At the time of writing, Karen Etheridge has the only implementation of OpenAPI 3.1 on CPAN; JSON::Validator has experimental 3.0 support and will need 3.1 support.
- A module called OpenAPI specifically geared towards creating APIs
    - A highly-discoverable entrypoint into CPAN for people wanting to work with OpenAPI and Perl
    - An extensively documented namespace that allows new users and old to quickly get working with OpenAPI
    - Signposting to other modules that implement relevant behaviour
    - Module interface designed with community input
    - Extensive automated tests

# Schedule

To be completed within 10 weeks of start of work.

# Project team

## Developers

- Paul G Webster / OpusVL 

I am a senior developer at OpusVL, an opensource focussed and perl supporting company that has developed software in use in industries as diverse from one another as medical/healthcare, biometrics and sales/auctions for the DVLA. 

I have been developing in perl since October of 2002, where it was present in the base of the FreeBSD operating system and required for building its userland and kernel. 

I hold a degree in computer science from the University of Lincoln and aside from my work with perl I am actively participating in technical documentation and creation of learning materials for FreeBSD as well as an ongoing effort to develop a language agnostic microservices protocol.

- Alastair Douglas

Well known in the community, having joined in 2008 and remained active since. Actively used Perl throughout, even when not employed as a Perl developer.

I have extensive experience in designing APIs and telling other people off for doing them wrong. After I delved into the OpenAPI spec for a project I took control of the OpenAPI namespace on CPAN with the intention of Doing It Right, but have until now not had the opportunity to make something of it.

## Documentor / technical writer

- Laurence Bogle

A technical writer with over a decade of experience in writing user guides and technical documents for both hardware and software. An MEng in Electronic and Electrical Engineering, plus lots of time spent working with developers has ensured a strong technical backg