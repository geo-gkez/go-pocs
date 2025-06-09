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
├── cmd/                   # Application entry points
│   └── main.go            # Main application entry point
├── configs/               # Configuration files
│   └── config.yml         # Application configuration
├── deployments/           # Deployment configurations
│   └── docker-compose.yml # Docker Compose for Redpanda
├── internal/              # Private application code
│   ├── config/            # Configuration management
│   ├── model/             # Data models
│   ├── routes/            # HTTP routes
│   └── service/           # Business logic services
├── go.mod                 # Go module file
├── go.sum                 # Go module checksums
├── Makefile               # Build and utility commands
└── README.md              # Project documentation
```

## Prerequisites

- Go 1.18 or later
- Docker and Docker Compose
- Make (optional, for using the provided Makefile)

## Getting Started

### Setting Up the Infrastructure

Start the Redpanda cluster and console:

```bash
make infra
```

This will start:
- A single-node Redpanda cluster
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

The API server will start on port 8085 (or as configured).

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

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [franz-go](https://github.com/twmb/franz-go) - Kafka client for Go
- [Redpanda](https://redpanda.com/) - Kafka-compatible streaming platform
- [Docker](https://www.docker.com/) - Containerization

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Redpanda for providing a Kafka-compatible streaming platform
- The Go community for excellent libraries and tools