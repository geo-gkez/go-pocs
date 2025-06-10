# Redpanda POC

A Proof of Concept application demonstrating how to use Redpanda (Kafka API-compatible streaming platform) with Go. This project provides a simple API for producing messages to Redpanda topics.

## Overview

This project demonstrates:

- Setting up a Redpanda cluster using Docker Compose
- Creating a Go service that interacts with Redpanda using the franz-go client
- Exposing a REST API for producing messages to Redpanda topics
- Using Gin as the web framework
- Containerized development environment

## Project Structure

```
redpanda-poc/
├── api/                   # API documentation and examples
│   └── go-redpanda.http   # HTTP request examples
├── build/                 # Build artifacts
│   └── redpanda-poc       # Compiled binary
├── cmd/                   # Application entry points
│   └── main.go            # Main application entry point
├── configs/               # Configuration files
│   └── config.yml         # Application configuration
├── deployments/           # Deployment configurations
│   └── docker-compose.yml # Docker Compose for Redpanda
├── internal/              # Private application code
│   ├── config/            # Configuration management
│   │   ├── app_config.go  # App configuration
│   │   ├── kafka_config.go # Kafka configuration
│   │   ├── logger/        # Logging configuration
│   │   │   └── logger.go  # Logger implementation
│   │   └── models/        # Configuration models
│   │       └── config_models.go # Configuration data structures
│   ├── model/             # Data models
│   │   └── http_models.go # HTTP request/response models
│   ├── routes/            # HTTP routes
│   │   └── route.go       # Route definitions
│   └── service/           # Business logic services
│       ├── kafka_consumer.go # Kafka consumer implementation
│       └── kafka_service.go  # Kafka service implementation
├── go.mod                 # Go module file
├── go.sum                 # Go module checksums
├── Makefile               # Build and utility commands
└── README.md              # Project documentation
```

## Prerequisites

- Go 1.24 or later
- Docker and Docker Compose
- Make (optional, for using the provided Makefile)

## Getting Started

### Setting Up the Infrastructure

Start the Redpanda cluster and console:

```bash
make infra
```

This will start:
- A single-node Redpanda cluster (v25.1.4)
- Redpanda Console UI accessible at http://localhost:9080

### Building and Running the Application

Build the application:

```bash
make build
```

Run the application:

```bash
make run
```

The API server will start on port 8085 (as configured in configs/config.yml).

### Tearing Down the Infrastructure

When you're done, you can stop and remove the Redpanda containers:

```bash
make infra-down
```

## Using the API

### Produce a Message

```http
POST /produce
Content-Type: application/json

{
  "key": "your-key",
  "message": "your-message"
}
```

Example using curl:

```bash
curl -X POST http://localhost:8085/produce \
  -H "Content-Type: application/json" \
  -d '{"key":"test-key","message":"Hello Redpanda!"}'
```

## Makefile Commands

- `make fmt`: Format the code
- `make vet`: Run go vet
- `make build`: Build the application
- `make run`: Run the application
- `make clean`: Clean up build artifacts
- `make infra`: Start infrastructure with Docker Compose
- `make infra-down`: Stop infrastructure and remove volumes
- `make help`: Show help message

## Redpanda Console

The Redpanda Console provides a web UI for monitoring and managing your Redpanda cluster. After starting the infrastructure, access it at:

http://localhost:9080

### Default Kafka Topics

The application is configured with the following default topics (as defined in configs/config.yml):

- Producer Topic: `test.output`
- Consumer Topic: `test.input`
- Consumer Group: `test.group`

## Technologies Used

- [Go](https://golang.org/) - Programming language (v1.24)
- [Gin](https://github.com/gin-gonic/gin) - Web framework (v1.10.1)
- [franz-go](https://github.com/twmb/franz-go) - Kafka client for Go (v1.19.5)
- [Redpanda](https://redpanda.com/) - Kafka-compatible streaming platform (v25.1.4)
- [Docker](https://www.docker.com/) - Containerization
- [Viper](https://github.com/spf13/viper) - Configuration management (v1.20.1)

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Redpanda for providing a Kafka-compatible streaming platform
- The Go community for excellent libraries and tools