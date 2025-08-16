# Shadow Docs

A simple Go backend server.

## Project Structure

```
shadow-docs/
├── cmd/
│   └── server/          # Application entry points
├── internal/
│   ├── handlers/        # HTTP request handlers
│   ├── models/          # Data models and structs
│   └── database/        # Database operations
├── pkg/
│   ├── middleware/      # HTTP middleware
│   └── utils/           # Utility functions
├── configs/             # Configuration files
├── docs/                # Documentation
├── main.go              # Main application entry point
├── go.mod               # Go module file
└── README.md            # This file
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd shadow-docs
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

The server will start on port 8080 by default. You can change the port by setting the `PORT` environment variable.

### API Endpoints

- `GET /` - Welcome message
- `GET /health` - Health check endpoint

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o bin/server main.go
```

### Environment Variables

- `PORT` - Server port (default: 8080)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## License

This project is licensed under the MIT License. 