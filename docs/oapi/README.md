## OpenAPIv3 specification

This directory contains both a specification for our Todo Backend, as well as an example demonstrating code generation based on the specification.

### OpenAPIv3 and Code Generation with oapi-codegen

API First (or Document Driven Design) is an engineering and architecture best practice. API First involves establishing a contract, separate from the code. This allows us to more clearly track the evolution of that contract, separate from the evolution of the implementation of that contract.

API contracts can be specified in OpenAPI Specification (previously Swagger), RAML, and API Blueprint. Of these, OpenAPI is the most popular, and OpenAPI v3 is the most current. There is a large and healthy ecosystem for OpenAPI tools in most languages.

As both API specifications and code implementing them evolves, you can need to pick one or the other as the source of truth, and rigorously test that they agree. Alternatively, with automated Code Generation tools, updating API clients/servers from specifications can be released in a matter of seconds, with no fear of breaking the new API contract, and no need for brittle, costly contract tests.

For Go, the https://github.com/deepmap/oapi-codegen tool is the most popular OpenAPI V3 client/server generator currently available. It generated a viable client with minimal deviation from our existing coding standards, and little to no reliance on third party dependencies. The codebase of the tool was relatively straightforward, and my pull request to support a new feature was quickly and gratefully accepted by the company who is actively maintaining it.

Oapi-codegen is released under the Apache license and runs as a devtool rather than in production. Because of the Apache license, the generated code should be safe (from a licensing perspective) in our system.

Another popular, older alternative is go-swagger for Swagger 2.0 code generation, but the maintainers have said they will definitely never support OpenAPI V3 as it is a hobby project.

There is a java program swagger-codegen that can generate Go (and other languages) but when I tried it, the code it generated was broken in a number of ways.

There are a number of private paid services that will codegen to Go, but it is not clear if they are just using one of the existing open source tools under the hood. I did not pursue them with any seriousness.
