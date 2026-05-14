# ADRs

## 20260514-1 Using Go

accepted

### Context

MCP servers can be written in different languages.
Following [modelcontextprotocol](https://modelcontextprotocol.io/docs/sdk), the tier 1 languages are TypeScript, Python, C#, and Go.

### Decision

The author already uses Python and Go and is less familiar with TypeScript.
Because the author wants to improve Go skills and Go is a secure and fast language, this project will use Go.

### Consequences

- The codebase and development workflow will follow Go conventions (modules, build tooling, dependency management).
- There is a learning curve, which may result in some suboptimal or non-idiomatic code early on.
- Choosing Go influences CI, packaging, and contributor expectations.

## 20260514-2 Using `testify` for assertions

accepted

### Context

The project requires tests with clear assertions. The author is more familiar with assertion-style testing from other languages and prefers concise assertion helpers for readability.

### Decision

Use the `github.com/stretchr/testify/assert` assertion helpers to make tests more expressive and easier to read.

### Consequences

- Adds a testing dependency.
- Tests may be easier to read and write, especially for those coming from other ecosystems.
- This may reduce emphasis on the minimal standard-library testing style.

## 20260514-3 Focusing on Integration Tests

accepted

### Context

Software can be tested at different levels: unit, integration, and end-to-end.
Focusing on one level while using the others selectively has different tradeoffs.
A clear testing strategy helps keep the project consistent.

### Decision

The author values integration tests more because they show that expected use cases are working and still run quickly.
Integration tests also allow more radical refactorings without breaking every test.

Integration tests will use an mcp client and start the mcp server by there own. Bitbucket server will be replaced by a docker container to simulate responses.

Unit tests can be used for complex cases or for edge conditions that are difficult to express with integration tests.

### Consequences

- Integration tests require more setup and teardown to avoid affecting each other.
- The test suite will be easier to understand for others because it focuses on behavior and use cases rather than implementation details.
