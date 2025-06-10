# Go Proof of Concept Projects

This repository contains a collection of Go proof of concept (POC) projects designed to explore and demonstrate various Go technologies, patterns, and frameworks.

## Project Overview

The repository is organized into the following projects:

### Redpanda POC

A proof of concept application demonstrating integration between Go and Redpanda (a Kafka API-compatible streaming platform). The project provides a simple REST API for producing messages to Redpanda topics and includes a Kafka consumer implementation.

- **Technologies**: Go (v1.24), Gin (v1.10.1), franz-go Kafka client (v1.19.5), Redpanda (v25.1.4), Docker, Viper (v1.20.1)
- **Directory**: [redpanda-poc](./redpanda-poc)
- **Features**:
  - REST API for producing messages to Redpanda
  - Kafka consumer implementation for processing messages
  - Docker Compose setup for Redpanda infrastructure
  - Clean architecture with separation of concerns

### Tests POC

A collection of example projects exploring Test-Driven Development (TDD) patterns in Go, closely following the "Learn Go with Tests" tutorial, including concurrency patterns and best practices.

- **Technologies**: Go, Go testing package
- **Directory**: [tests-poc](./tests-poc)
- **Features**:
  - Comprehensive TDD examples covering:
    - Concurrency patterns with goroutines and channels
    - Dependency injection techniques
    - Dictionary implementation with maps
    - Mocking for effective testing
    - Banking wallet implementation
    - Shapes and interfaces
  - Reference document of Go tools and tips, including a dedicated section on Go concurrency (goroutines, channels, race conditions, and common pitfalls)

## Getting Started

Each project contains its own README with specific instructions for setting up and running that particular POC. Navigate to the project directory of interest and follow the instructions there.

### Prerequisites

- Go 1.24 or later
- Docker and Docker Compose (for some projects)
- Make (optional, for using provided Makefiles)

## Repository Structure

```
go-pocs/
├── README.md                  # Main repository documentation
├── redpanda-poc/              # Redpanda integration POC
│   ├── api/                   # API documentation
│   │   └── go-redpanda.http   # HTTP request examples
│   ├── build/                 # Build artifacts
│   │   └── redpanda-poc       # Compiled binary
│   ├── cmd/                   # Application entry points
│   │   └── main.go            # Main application entry point
│   ├── configs/               # Configuration files
│   │   └── config.yml         # Application configuration
│   ├── deployments/           # Deployment configurations
│   │   └── docker-compose.yml # Docker Compose for Redpanda
│   ├── internal/              # Internal application code
│   │   ├── config/            # Configuration management
│   │   │   ├── app_config.go  # App configuration
│   │   │   ├── kafka_config.go # Kafka configuration 
│   │   │   ├── logger/        # Logging configuration
│   │   │   │   └── logger.go  # Logger implementation
│   │   │   └── models/        # Configuration models
│   │   │       └── config_models.go # Configuration data structures
│   │   ├── model/             # Data models
│   │   │   └── http_models.go # HTTP request/response models
│   │   ├── routes/            # HTTP routes
│   │   │   └── route.go       # Route definitions
│   │   └── service/           # Business logic services
│   │       ├── kafka_consumer.go # Kafka consumer implementation
│   │       └── kafka_service.go  # Kafka service implementation
│   ├── go.mod                 # Go module file
│   ├── go.sum                 # Go module checksums
│   ├── Makefile               # Build and utility commands
│   └── README.md              # Project documentation
└── tests-poc/                 # TDD examples and patterns
    ├── go-tools-and-tips.md   # Reference document for Go tools and concurrency
    ├── go.mod                 # Go module file for tests
    └── tdd/                   # Test-Driven Development examples
        ├── banking/           # Banking wallet implementation with TDD
        ├── concurrency/       # Concurrency patterns with goroutines and channels
        ├── dependency_injection/ # Dependency injection patterns
        ├── dictionary/        # Dictionary implementation showing map usage
        ├── mocking/           # Mocking examples for testing
        └── shapes/            # Interfaces demonstrated with shapes
```

## Development Approach

These projects follow several key principles:

1. **Test-Driven Development**: Most code is written following the TDD approach, writing tests first before implementation.
2. **Clean Architecture**: Projects use a clean architecture approach with clear separation of concerns.
3. **Idiomatic Go**: Code attempts to follow Go best practices and idiomatic patterns.
4. **Containerization**: Where applicable, projects include Docker configurations for easy environment setup.

## Learning Resources

- [Go Tools and Tips](./tests-poc/go-tools-and-tips.md) - A compilation of useful Go tools, concurrency patterns, and development tips inspired by "Learn Go with Tests"
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) - Source tutorial for the tests-poc examples, including concurrency fundamentals

## Future Projects

Planned proof of concept projects include:

- gRPC service implementation
- GraphQL API with Go
- Advanced state management patterns
- Performance optimization techniques