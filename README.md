# Qompass Go Template

A template for creating Go applications at Qompass AI.

## Features

- Modern Go project layout
- Post-quantum cryptography ready
- Containerization with Zig compilation for native cross-platform deployment
- CI/CD setup with GitHub Actions

## Getting Started

1. Clone this repository
2. Replace all instances of `github.com/Qompass/gotemplate` with your module path
3. Run `make build` to build the application
4. Run `make docker` to build a container

## Development

- Run tests: `make test`
- Start local server: `make run`

## Architecture

This template follows clean architecture principles with:
- API handlers in `internal/api`
- Business logic in `internal/services`
- Data models in `internal/models`
- Reusable packages in `pkg`

## Base Structure

```go
gotemplate/
├── cmd/                    # Main applications
│   └── app/                # Main application entry point
│       └── main.go         # Main executable code
├── internal/               # Private code that won't be imported
│   ├── api/                # API handlers
│   ├── models/             # Data models
│   └── services/           # Business logic
├── pkg/                    # Public library code
│   └── quantum/            # Post-quantum crypto utilities
├── configs/                # Configuration files
├── scripts/                # Build and automation scripts
├── test/                   # Additional test applications/configs
├── docs/                   # Documentation files
├── .github/                # GitHub Actions workflows
├── Containerfile.go        # Already present
├── .gitignore              # Specify files to ignore
├── go.mod                  # Module dependencies
├── go.sum                  # Checksum for module dependencies
├── Makefile                # Common commands
├── README.md               # Project documentation
└── LICENSE-*               # Dual-Licenses 
```
