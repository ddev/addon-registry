---
title: mxr576/ddev-pi-golang
github_url: https://github.com/mxr576/ddev-pi-golang
description: "Adds Go programming capabilities to PI Coding Agent. "
user: mxr576
repo: ddev-pi-golang
repo_id: 1279328982
default_branch: main
tag_name: 1.0.0-alpha1
ddev_version_constraint: ">= v1.25.2"
dependencies: ["mxr576/ddev-pi"]
type: contrib
created_at: 2026-06-24
updated_at: 2026-06-25
workflow_status: failure
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/mxr576/ddev-pi-golang/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/mxr576/ddev-pi-golang/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/mxr576/ddev-pi-golang)](https://github.com/mxr576/ddev-pi-golang/commits)
[![release](https://img.shields.io/github/v/release/mxr576/ddev-pi-golang)](https://github.com/mxr576/ddev-pi-golang/releases/latest)

# DDEV Pi Golang

📚 **Related Documentation**: [SUMMARY.md](https://github.com/mxr576/ddev-pi-golang/blob/main/SUMMARY.md) - Tool selection philosophy | [TOOLS.md](https://github.com/mxr576/ddev-pi-golang/blob/main/TOOLS.md) - Detailed tool reference

## Overview

This add-on extends the [ddev-pi](https://github.com/mxr576/ddev-pi) addon with comprehensive Golang development capabilities, transforming the Pi coding agent into an expert Go developer.

## Features

### Go Runtime
- **Go 1.26.4** (configurable version)
- Multi-architecture support (amd64, arm64, armv6l)
- Properly configured GOPATH and Go modules support

### Essential Go Tools (10 tools)

**Language Server:**
- **`gopls`** - Official Go language server for IDE features

**Code Formatting:**
- **`goimports`** - Format code and manage imports automatically

**Linting & Analysis:**
- **`golangci-lint`** - Comprehensive linting (50+ linters in one)
- **`staticcheck`** - Advanced static analysis for deeper insights
- **`govulncheck`** - Security vulnerability scanning for dependencies

**Testing:**
- **`gotests`** - Generate table-driven test boilerplate
- **`mockgen`** - Generate mock implementations for testing

**Development Tools:**
- **`dlv`** (Delve) - Full-featured debugger
- **`impl`** - Generate interface method stubs

### Pi Agent Enhancements

The addon includes a custom Pi extension that:
- Injects Go best practices into the agent's system prompt
- Provides context-aware guidance for common Go tasks
- Offers helpful suggestions during development
- Follows official Go guidelines and idioms

> **Extension Location**: `pi/global/agent/extensions/golang-development.ts`  
> Advanced users can customize the extension by editing this file.

## Installation

```bash
ddev add-on get mxr576/ddev-pi-golang
ddev restart
```

After installation, commit the `.ddev` directory to version control.

## Usage

### Starting the Pi Service

The Pi service uses a Docker Compose profile and must be started explicitly.

**For existing projects:**
```bash
ddev start --profiles=pi
```

**After addon installation or configuration changes:**
```bash
ddev restart && ddev start --profiles=pi
```

**Note**: The `ddev restart` ensures the addon configuration is loaded before starting the Pi service.

### Basic Go Development

```bash
# Enter the Pi container
ddev ssh -s pi

# Create a new Go project
mkdir myproject && cd myproject
go mod init github.com/username/myproject

# Create main.go
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
EOF

# Run your program
go run .

# Format code and organize imports
goimports -w .

# Run linters
golangci-lint run

# Run tests with race detection
go test -race ./...

# Check for vulnerabilities
govulncheck ./...
```

### Using the Pi Agent for Go Development

```bash
# Start an interactive Pi session
ddev pi

# Example prompts:
# "Create a new Go web service using the standard library"
# "Add unit tests for this function"
# "Refactor this code to follow Go best practices"
# "Add proper error handling to this function"
```

## Advanced Customization

### Changing Go Version

To use a different Go version, set the environment variable before installation:

```bash
echo "GOLANG_VERSION=1.22.5" >> .ddev/.env.pi
ddev restart && ddev start --profiles=pi
```

### Changing golangci-lint Version

```bash
echo "GOLANGCI_LINT_VERSION=1.61.0" >> .ddev/.env.pi
ddev restart && ddev start --profiles=pi
```

### Customization Options

| Variable | Default | Description |
| -------- | ------- | ----------- |
| `GOLANG_VERSION` | `1.26.4` | Go version to install (check latest: https://go.dev/VERSION?m=text) |
| `GOLANGCI_LINT_VERSION` | `1.62.2` | golangci-lint version |

**Note:** Changing versions requires rebuilding the Pi container. Due to [ddev/ddev#8463](https://github.com/ddev/ddev/pull/8463), use the command above rather than `ddev utility rebuild`.

## Go Best Practices Enforced by Pi Agent

The Pi agent extension automatically guides you to follow these practices:

- **Idiomatic Go**: Follows [Effective Go](https://go.dev/doc/effective_go)
- **Error Handling**: Explicit error checking with context wrapping
- **Testing**: Table-driven tests with subtests
- **Concurrency**: Proper goroutine lifecycle and channel usage
- **Project Structure**: Standard Go project layout
- **Dependencies**: Minimal, well-maintained dependencies

## Troubleshooting

### Module cache issues

Clear the Go module cache:
```bash
ddev exec -s pi go clean -modcache
```

## Available Go Commands Reference

| Command | Description |
| ------- | ----------- |
| `go run .` | Run the main package |
| `go build` | Compile packages and dependencies |
| `go test ./...` | Run all tests recursively |
| `go test -v ./...` | Run tests with verbose output |
| `go test -race ./...` | Run tests with race detection |
| `go test -cover ./...` | Run tests with coverage |
| `go mod init <module>` | Initialize a new module |
| `go mod tidy` | Add missing and remove unused modules |
| `go get <package>@<version>` | Add or update a dependency |
| `go install <package>@latest` | Install a Go binary |
| `go vet ./...` | Report likely mistakes |
| `goimports -w .` | Format code and fix imports |
| `golangci-lint run` | Run comprehensive linting |
| `staticcheck ./...` | Advanced static analysis |
| `govulncheck ./...` | Check for vulnerabilities |
| `gotests -all -w file.go` | Generate tests for file |
| `dlv debug` | Start debugger |
| `impl 'r *Type' io.Reader` | Generate interface stubs |
| `mockgen -source=file.go` | Generate mocks |

## Credits

**Contributed and maintained by [@mxr576](https://github.com/mxr576)**
