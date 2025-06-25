# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go project template (`pj-tmpl-go`) that provides a well-structured foundation for new Go projects. It includes an HTTP server example with graceful shutdown, a CLI wrapper, and professional development tooling.

## Prerequisites

- Go 1.24.2 or later
- `gotest` - Test wrapper tool
- `lefthook` - Git hooks manager (install with `go install github.com/evilmartians/lefthook@latest`)
- `gitleaks` - Secret scanning tool (used by git hooks)

## Template Usage

To create a new project from this template:
```bash
PKG=github.com/yourname/yourpj
ghq create ${PKG}
gonew github.com/anoriqq/pj-tmpl-go $(ghq list -e ${PKG}) $(ghq list -p -e ${PKG})/tmp
mv $(ghq list -p -e ${PKG})/tmp/* $(ghq list -p -e ${PKG})
rm -r $(ghq list -p -e ${PKG})/tmp
```

## Essential Commands

### Development
- `make help` - Display all available make targets with descriptions
- `make build` - Build all binaries to `bin/` directory (gitignored)
- `make test` - Run all tests with race detection (1s timeout)
- `make clean` - Remove built binaries
- `RELEASE=1 make build` - Production build (strips symbols, static linking, no race detector)

### Testing
- `gotest ./...` - Run tests (wrapper around go test)
- `gotest -run TestName` - Run specific test

## Architecture

### Project Structure
- `cmd/server/` - CLI application entry point with signal handling
- `cmd/server/internal/cli/` - CLI implementation with options parsing
- `cmd/server/internal/log/` - Custom slog handler with pretty JSON output
- `run.go` - Main application logic (HTTP server on :8888)
- `internal/` - Shared internal packages

### Key Patterns
1. **Error Handling**: Uses `github.com/go-errors/errors` for stack traces
2. **Logging**: Custom pretty JSON slog handler with colored output
3. **Context**: Proper context propagation with graceful shutdown (5s timeout)
4. **Testing**: Table-driven tests using maps (not slices) for randomized execution

### Testing Best Practices (from docs/general/testing_essentials.md)
- Use descriptive test names: `Test_FunctionName_Condition_ExpectedResult`
- Separate tests for different specifications
- Provide detailed error messages
- Golden files are used for testing (see `testdata/` directories)

## Important Notes

### Git Hooks (lefthook)
- Prevents direct commits to main/master branches
- Enforces conventional commits (Angular style)
- Runs gitleaks security scanning
- Validates Signed-off-by in commits

### Development Considerations
- This is a template meant to be used with `gonew`
- Some documentation is in Japanese (testing guide in `docs/general/testing_essentials.md`)
- No CI/CD configuration included (add per project needs)
- Golden test files stored in `testdata/` directories
- Binaries are built to `bin/` directory which is gitignored

### Initial Setup
After cloning or creating from template:
1. Install lefthook: `lefthook install`
2. Verify git hooks are working: `lefthook run pre-commit`