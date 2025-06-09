# Go Proof of Concept Projects

This repository contains a collection of Go proof of concept (POC) projects designed to explore and demonstrate various Go technologies, patterns, and frameworks.

## Project Overview

The repository is organized into the following projects:

### Redpanda POC

A proof of concept application demonstrating integration between Go and Redpanda (a Kafka API-compatible streaming platform). The project provides a simple REST API for producing messages to Redpanda topics.

- **Technologies**: Go, Gin, franz-go Kafka client, Redpanda, Docker
- **Directory**: [redpanda-poc](./redpanda-poc)
- **Features**:
  - REST API for producing messages to Redpanda
  - Docker Compose setup for Redpanda infrastructure
  - Clean architecture with separation of concerns

### Tests POC

A collection of example projects exploring Test-Driven Development (TDD) patterns in Go, following the "Learn Go with Tests" tutorial.

- **Technologies**: Go, Go testing package
- **Directory**: [tests-poc](./tests-poc)
- **Features**:
  - TDD examples for various Go concepts
  - Banking wallet implementation with TDD
  - Shapes area calculations demonstrating interfaces
  - Reference document of Go tools and tips

## Getting Started

Each project contains its own README with specific instructions for setting up and running that particular POC. Navigate to the project directory of interest and follow the instructions there.

### Prerequisites

- Go 1.18 or later
- Docker and Docker Compose (for some projects)
- Make (optional, for using provided Makefiles)

## Repository Structure

```
go-pocs/
├── README.md                  # Main repository documentation
├── redpanda-poc/              # Redpanda integration POC
│   ├── api/                   # API documentation
│   ├── build/                 # Build artifacts
│   ├── cmd/                   # Application entry points
│   ├── configs/               # Configuration files
│   ├── deployments/           # Deployment configurations
│   └── internal/              # Internal application code
└── tests-poc/                 # TDD examples and patterns
    ├── go-tools-and-tips.md   # Reference document for Go tools
    └── tdd/                   # Test-Driven Development examples
        ├── banking/           # Banking wallet implementation
        └── shapes/            # Shapes and area calculations
```

## Development Approach

These projects follow several key principles:

1. **Test-Driven Development**: Most code is written following the TDD approach, writing tests first before implementation.
2. **Clean Architecture**: Projects use a clean architecture approach with clear separation of concerns.
3. **Idiomatic Go**: Code attempts to follow Go best practices and idiomatic patterns.
4. **Containerization**: Where applicable, projects include Docker configurations for easy environment setup.

## Learning Resources

- [Go Tools and Tips](./tests-poc/go-tools-and-tips.md) - A compilation of useful Go tools and development tips
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) - Source tutorial for the tests-poc examples

## Future Projects

Planned proof of concept projects include:

- gRPC service implementation
- GraphQL API with Go
- Concurrent programming patterns

## License

This repository is licensed under the MIT License - see individual project directories for specific licensing information.