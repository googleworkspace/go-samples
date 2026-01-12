# AGENTS.md

## Repository Overview

This repository contains Go code samples for Google Workspace APIs.
**IMPORTANT**: This is a collection of independent samples, NOT a monolithic application. Each directory typically contains a standalone executable or package.

## Project Structure

The repository is organized by API service:

- `[api]/[sample-type]/` (e.g., `drive/quickstart/`, `sheets/snippets/`)

## Development Workflow

### Build

To build all samples:

```bash
go build -v ./...
```

_Note: If builds are cached and silent, use `go build -v -a ./...` to force output._

### Format

Ensure all code is formatted with standard `gofmt`.

```bash
# Check for unformatted files (exit code 1 if output is not empty)
test -z $(gofmt -l .)

# Fix formatting
go fmt ./...
```

### Vet

Run static analysis:

```bash
go vet ./...
```

### Test

Run unit tests (where available):

```bash
go test -v ./...
```

### Tidy

Ensure modules are clean:

```bash
go mod tidy
```

## Environment

- **Go Version**: Latest Stable (currently 1.24+)
- **CI Pipelines**:
  - `Test`: Runs build checks via `go build`.
  - `Lint`: Runs `gofmt` and `go vet`.

## Code Snippets

When writing samples that will be referenced in documentation, mark the regions using the following format:

```go
// [START unique_snippet_id]
func Example() {
        // ...
}
// [END unique_snippet_id]
```

- Ensure the ID is unique across the repository.
- Do not indent the `// [START ...]` and `// [END ...]` tags.
