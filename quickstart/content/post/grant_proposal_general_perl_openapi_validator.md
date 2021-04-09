
---
title: "Grant Proposal: General Perl OpenAPI Validator / Interpreter"
author: Jason A. Crome
type: post
date: 2020-09-24 00:00:00 +0000 UTC
url: "/post/grant_proposal_general_perl_openapi_validator"
categories:
 - Grants
 - Sign in

---

## Synopsis

Create a standalone implementation of OpenAPI 3.0 to facilitate parsing, creating, and validating OpenAPI specifications.

## Overview

It is common contemporary practice to provide a OpenAPI (previously Swagger^[Swagger was renamed to OpenAPI for version 3]) schema^[<https://swagger.io/specification/>] to communicate to clients the available functionality of a RESTful web API. It is well documented and flexible, and well known to application developers.

There exists no Perl module to facilitate working with an OpenAPI 3.0 schema. Project-specific modules exist for Swagger 2.0, but nothing generic and without dependencies.

We intend to write a reference module to work with a variety of different web frameworks, such as Dancer, POE, Catalyst, etc. It will follow the emerging standards for Perl 7 (but target Perl 5).

The project will use as few CPAN modules as possible - this is a *reference* module: the idea is people extend *from* it, and it is likely to act as a base for other applications developed.

To reduce project risk, the work will be conducted in two parts, firstly an analysis and specification stage, and, once complete, an implementation stage. This allows us to properly scope the project and collect use cases from the community rather than ploughing on ahead with code that nobody wants to use.

## The benefits of the project to the Perl Community

Swagger has been around since 2011 and has been adopted as the *de facto* standard for describing RESTful web APIs. It was renamed to OpenAPI for version 3.0 when the specification was donated to the OpenAPI initiative.

OpenAPI 3.0 creates a standardised json based API definition that can easily be consumed by a range of different clients; a website or some form of automation service. At present development and standards adherance are hindered by having so few server-side options for OpenAPI.

No work has been done on CPAN for this version of the specification. The only one presently available is a plugin for Mojo [Mojolicious::Plugin::OpenAPI](https://metacpan.org/pod/Mojolicious::Plugin::OpenAPI) which is Mojo-specific and only targets Swagger 2.

We have collected opinions from the community regarding OpenAPI 3 with the conclusion that it would be a great tool, but there is nothing available to make it easy to use in Perl. In comparison, SOAP is easy to use in Perl due to existing modules allowing developers to automate the generation and utilisation of their project specification.

A base reference module would allow for other developers to easily create applications that can consume OpenAPI schemas, allowing for validation and dynamic handler creation without spending possibly weeks validating both the API definition and that the handlers internal to the developed application are acting correctly.

It would also allow for more efficient deployment in terms of development time for developers in any industry that has applications connected to a network.

Adherence to the Perl 7 standards will allow us to maintain the code even if Perl 7 becomes the dominant form of the language, and is in any case a set of best practices for writing Perl 5 code. This should hopefully lower the barrier to contributors and, naturally, reduce errors within the module itself.

I will invite other perl developers to work with me on this project and all code and commits will be made through a public git repository.

## Deliverables

These are based on other OpenAPI implementations in other languages as well as additional requirements for our communities and developers.

### Community engagement

Contact others in the Perl community who have an interest (e.g. Mojo) and take on board their input.

The community can get involved with us through the standard perl communication networks and by raising issues here: https://github.com/altreus/OpenAPI

### Analysis

Review inputs from perl communtiy and of other implementations of OpenAPI and investigate how these relate to the perl ecosystem. 

- [OpenAPI Specification](https://github.com/OAI/OpenAPI-Specification)
- [OpenAPI tools](https://openapi.tools/)
- [npm openapi](https://www.npmjs.com/package/express-openapi)
- [PYPI openapi3](https://pypi.org/project/openapi3/)

Prior art in perl:

- [Mojolicious::Plugin::OpenAPI](https://metacpan.org/pod/Mojolicious::Plugin::OpenAPI) allows you to autogenerate a web app out of a Swagger 2 specification (a bit of a naming mistake alas, as OpenAPI is 3.0!). It will ensure you provide handlers, and validate your inputs and outputs.
- [JSON::Schema](https://metacpan.org/pod/JSON::Schema) provides the base behaviour, since OpenAPI is a JSON-Schema specification.
- [JSON::Pointer](https://metacpan.org/pod/JSON::Pointer) implements the part of the OpenAPI specification that makes use of references


### Specification

The distribution will be fully specced out before development starts (modulo any prototype code that happens to be useful). This is milestone 1. Milestones 2 and 3 are also likely to be begun here.

The specification will detail:
- the modules to be created
    - feature brief
    - module API
- dependencies identified
- definition of the test approach
- scope for documentation

A primary outcome of milestone 1, therefore, is to properly scope milestone 3. The rest of this section outlines what we anticipate to be able to deliver, but it remains hypothetical until the specification is complete.

### Functionality

The main function of the module is validation. This will be responsible for:

- Validating an OpenAPI schema against the OpenAPI spec
- Validating a payload against the schema, as a request or response
- Validating an application fulfils a complete specification
    - By means of some special data structure

Otherwise the module will be designed to facilitate the use of OpenAPI. The idea is to expose the features of the OpenAPI specification as methods and classes, so code using an OpenAPI object can be written expressively.

### JSON::Schema

OpenAPI is a JSON-schema specification. This project may require contributions to JSON::Schema as well.

### JSON::Pointer

OpenAPI uses the JSON-Pointer specification to reference itself and other documents. JSON::Pointer implements this, so contributions here may be in scope.

### Documentation

We will have a technical writer to build solid documentation from the outset. This will be throughout, including the specification, but chiefly for milestone 2.

### Testing

The module will endeavour to have 100% test coverage. Meaning:

- Unit tests for all classes, all known inputs and outputs
- Meaningful integration tests covering the broader OpenAPI specification (where different from unit tests)
- Regression tests for bugfixes
- Usage examples

There are CPAN modules that can assist with this measurement, which will be evaluted as part of the specification.

The github repository will have a continuous integration badge showing whether tests are currently passing in master.

### Registration

Update various API resource locations and directories to note that there is a perl provider implementation:
- https://github.com/OAI/OpenAPI-Specification/blob/master/IMPLEMENTATIONS.md#implementations
- https://github.com/APIs-guru/openapi-directory


## Project team

### Lead developer

- Paul G Webster / OpusVL 

I am a senior developer at OpusVL, an opensource focussed and perl supporting company that has developed software in use in industries as diverse from one another as medical/healthcare, biometrics and sales/auctions for the DVLA. 

I have been developing in perl since October of 2002, where it was present in the base of the FreeBSD operating system and required for building its userland and kernel. 

I hold a degree in computer science from the University of Lincoln and aside from my work with perl I am actively participating in technical documentation and creation of learning materials for FreeBSD as well as an ongoing effort to develop a language agnostic microservices protocol.

### Co-developer

- Alastair Douglas

Well known in the community, having joined in 2008 and remained active since. Actively used Perl throughout, even when not employed as a Perl developer.

I have extensive experience in designing APIs and telling other people off for doing them wrong. After I delved into the OpenAPI spec for a project I took control of the OpenAPI namespace on CPAN with the intention of Doing It Right, but have until now not had the opportunity to make something of it.

### Documentor / technical writer

- Laurence Bogle

A technical writer with over a decade of experience in writing user guides and technical documents for both hardware and software. An MEng in Electronic and Electrical Engineering, plus lots of time spent working with developers has ensured a strong technical background. 

## Project cost and milestones

### Specification phase

#### Milestone 1: Community engagement, analysis, specification

- Delivers: 
    - Skeleton repository
    - Spec document in the repo
    - Start of longer-term roadmap
    - Community comms channels established

#### Costs

- Total time required: 2 working weeks at 6 hours / day for two developers
- Rate: $1350 / week at $45 / hour per developer
- Time scale: 3 calendar weeks
- Total cost: $2700

### Implementation phase

#### Milestone 2: Technical documentation

- Delivers: 
    - POD Documentation and supporting documents
    - Proof of concept code
    - Testing plan

#### Milestone 3: CPAN module 

- Delivers: 
    - Module checked in to CPAN and available for others to use
    - Automated testing/CI
    - Final roadmap

## Costs

- Total time required: 4 working weeks at 6 hours / day for two developers plus 1 day for documentation
- Rate: $1350 / week at $45 per hour per person
- Time scale: 6 calendar weeks
- Total cost: $5670

