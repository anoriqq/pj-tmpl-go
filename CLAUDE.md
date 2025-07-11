# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go project template (`pj-tmpl-go`) that provides a well-structured foundation for new Go projects. It includes an HTTP server example with graceful shutdown, a CLI wrapper, and professional development tooling.

## Prerequisites

- `mise` - Tool version manager (manages all other tools via mise.toml)
- Tools managed by mise:
  - `gotest` - Test wrapper tool
  - `lefthook` - Git hooks manager
  - `gitleaks` - Secret scanning tool
  - `air` - Hot reload development server

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
- `make test` - Run all tests with race detection (3s timeout, count=2, shuffle=on)
- `make run` - Run the application with hot reloading using Air
- `make gen` - Run code generation (go generate ./...)
- `make clean` - Remove built binaries
- `RELEASE=1 make build` - Production build (strips symbols, static linking, no race detector)

### Testing
- `gotest ./...` - Run tests (wrapper around go test)
- `gotest -run TestName` - Run specific test
- `gotest ./internal/infra/cli -update` - Update golden test files

### CLI Options
- `-help` - Show help message
- `-env` - Set environment: lcl, dev, stg, prd (default: lcl)
- `-port` - Set port number (default: 8080)
- Environment variables override flags: `ENV=dev PORT=9000 ./bin/cmd`

## Architecture

### Project Structure
- `cmd/main.go` - CLI application entry point with signal handling
- `internal/infra/cli/` - CLI implementation with options parsing
- `internal/infra/log/` - Custom slog handler with pretty JSON output
- `internal/infra/server/` - HTTP server with graceful shutdown
- `internal/domain/env/` - Environment domain models with code generation
- `internal/domain/port/` - Port value object with validation (max 65535)
- `run.go` - Main application logic (coordinates CLI and server)

### Key Patterns
1. **Error Handling**: Uses `github.com/go-errors/errors` for stack traces
2. **Logging**: Environment-aware slog configuration:
   - LCL: Pretty JSON with debug level and colored output
   - DEV: Standard JSON with debug level
   - STG/PRD: Standard JSON with info level
3. **Context**: Proper context propagation with graceful shutdown (5s timeout)
4. **Testing**: Table-driven tests using maps (not slices) for randomized execution
5. **Code Generation**: Enum generation using `github.com/anoriqq/enumer`

### Testing Best Practices (from docs/general/testing_essentials.md)
- Use descriptive test names: `Test_FunctionName_Condition_ExpectedResult`
- Separate tests for different specifications
- Provide detailed error messages
- Golden files are used for testing (see `testdata/` directories)
- Tests use `github.com/tenntenn/golden` for golden file testing
- Test data generation with `github.com/brianvoe/gofakeit/v7`

## Important Notes

### Git Hooks (lefthook)
- Prevents direct commits to main/master branches
- Enforces conventional commits (Angular style)
- Runs gitleaks security scanning
- Validates Signed-off-by in commits

### Development Considerations
- This is a template meant to be used with `gonew`
- Some documentation is in Japanese (testing guide in `docs/general/testing_essentials.md`)
- Golden test files stored in `testdata/` directories
- Binaries are built to `bin/` directory which is gitignored
- Air configuration in `.air.toml` for hot reloading
- Build uses CGO_ENABLED=0 for static binaries
- Release builds include netgo tag and static linking

### Initial Setup
After cloning or creating from template:
1. Install mise if not already installed
2. Run `mise install` to install all required tools
3. Install lefthook: `lefthook install`
4. Verify git hooks are working: `lefthook run pre-commit`

### Running Tests
- `make test` - Run all tests with race detection
- `gotest ./...` - Run tests using the gotest wrapper
- `gotest -run TestName` - Run a specific test
- `gotest ./internal/infra/cli -update` - Update golden test files

## CI/CD

### GitHub Actions Workflows
- **ci** - Main CI pipeline triggered by pushes/PRs to main branch
  - Runs on Go file changes, Makefile changes, CI config updates, and test data changes
  - Uses `mise` for tool management, builds with `RELEASE=1 make build`, runs `make test`
- **claude-assistant** - AI-powered PR assistant activated by `@claude` mentions
  - Responds to issue comments, PR review comments, and issue assignments
  - 60-minute timeout with full tool access (Bash, Edit, WebSearch, etc.)
- **claude-review** - Automated PR reviews for every opened/updated PR
  - Reviews coding standards, error handling, security, test coverage, documentation
  - Uses O3 search MCP for enhanced reasoning capabilities

### Dependency Management
- **Renovate** - Automated dependency updates via `renovate.json`
  - Uses custom configuration: `local>anoriqq/renovate-config`
  - Manages Go modules and tool versions automatically

### Required Secrets
For Claude AI workflows:
- `CLAUDE_CODE_OAUTH_TOKEN` - Claude Code authentication
- `OPENAI_API_KEY` - OpenAI API access for enhanced search capabilities