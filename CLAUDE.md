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
  - `actionlint` - GitHub Actions linter
  - `ghalint` - GitHub Actions security linter
  - `pinact` - Pin GitHub Actions versions
  - `zizmor` - GitHub Actions security scanner
  - `golangci-lint` - Go linter aggregator

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
- `mise tasks` - Display all available tasks with descriptions
- `mise run build` - Build all binaries to `bin/` directory (gitignored)
- `mise run test` - Run all tests with race detection (3s timeout, count=2, shuffle=on)
- `mise run run` - Run the application with hot reloading using Air
- `mise run gen` - Run code generation (go generate ./...)
- `mise run clean` - Remove built binaries
- `mise run lint` - Run all linters (actionlint, ghalint, pinact, zizmor, golangci-lint)
- `RELEASE=1 mise run build` - Production build (strips symbols, static linking, no race detector)

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
- `cmd/main.go` - Application entry point with context and error handling
- `cmd/config.go` - Configuration loading from environment variables
- `cmd/eval.go` - Application lifecycle management with panic recovery
- `internal/infra/log/` - Custom slog handler with pretty JSON output
- `internal/infra/server/` - HTTP server with graceful shutdown
- `internal/infra/pnc/` - Panic handling utilities
- `internal/domain/env/` - Environment domain models with code generation
- `internal/domain/port/` - Port value object with validation (max 65535)

### Key Patterns
1. **Error Handling**: Uses `github.com/go-errors/errors` for stack traces
2. **Panic Recovery**: `eval()` function recovers panics and converts to errors via `pnc.Parse()`
3. **Configuration**: Environment-based config via `loadConfig()` using `sync.OnceValue` for singleton pattern
4. **Logging**: Environment-aware slog configuration:
   - LCL: Pretty JSON with debug level and colored output
   - DEV: Standard JSON with debug level
   - STG/PRD: Standard JSON with info level
5. **Context**: Proper context propagation with signal handling via `signal.NotifyContext`
6. **Testing**: Table-driven tests using maps (not slices) for randomized execution
7. **Code Generation**: Enum generation using `github.com/anoriqq/enumer`
8. **Layered Architecture**: `depguard` enforces domain layer cannot import infra layer

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
- **Linting**: `.golangci.yml` configures all linters with some disabled (wsl, varnamelen, revive, etc.)
  - `depguard` rule prevents domain layer from importing infra layer
  - Auto-formatters enabled: gci, gofumpt, goimports, gofmt, golines

### Initial Setup
After cloning or creating from template:
1. Install mise if not already installed
2. Run `mise install` to install all required tools
3. Install lefthook: `lefthook install`
4. Verify git hooks are working: `lefthook run pre-commit`

### Running Tests
- `mise run test` - Run all tests with race detection
- `gotest ./...` - Run tests using the gotest wrapper
- `gotest -run TestName` - Run a specific test
- `gotest ./internal/infra/cli -update` - Update golden test files

## CI/CD

### GitHub Actions Workflows
- **ci** - Main CI pipeline triggered by pushes/PRs to main branch
  - Uses path filtering to run only on Go file changes, mise.toml changes, CI config updates, and test data changes
  - Uses `mise` for tool management, builds with `RELEASE=1 mise run build`, runs `mise run test`
  - Composite action at `.github/actions/go/action.yml` handles setup and execution
- **cd** - Continuous deployment pipeline for infrastructure
  - Triggered on pushes/PRs to main branch
  - Runs Pulumi preview when `.pulumi/**` files change
  - Uses composite action at `.github/actions/pulumi/action.yml`
- **claude-assistant** - AI-powered PR assistant activated by `@claude` mentions
  - Responds to issue comments, PR review comments, issue assignments, and PR reviews
  - 60-minute timeout with full tool access including O3 search MCP
  - Condition checks: `@claude` mentions in comments or issue body
- **claude-review** - Automated PR reviews for every opened/updated PR
  - Only runs for PRs from user `anoriqq` (owner-specific review)
  - Reviews coding standards, error handling, security, test coverage, documentation
  - Uses O3 search MCP with medium context size and reasoning effort
  - Provides reviews in Japanese as specified in `direct_prompt`

### Path-Based CI Optimization
Both CI and CD workflows use `dorny/paths-filter` to conditionally run jobs:
- **CI triggers**: `**/*.go`, `**/go.mod`, `**/go.sum`, `**/go.work*`, `**/testdata/**`
- **CD triggers**: `.pulumi/**`
- **Force triggers**: `.github/workflows/**`, `.github/actions/**`, `mise.toml`

### Dependency Management
- **Renovate** - Automated dependency updates via `renovate.json`
  - Uses custom configuration: `local>anoriqq/renovate-config`
  - Manages Go modules and tool versions automatically

### Required Secrets
For Claude AI workflows:
- `CLAUDE_CODE_OAUTH_TOKEN` - Claude Code authentication
- `OPENAI_API_KEY` - OpenAI API access for enhanced search capabilities

For Pulumi CD:
- `PULUMI_ACCESS_TOKEN` - Pulumi cloud access
- `PULUMI_GITHUB_TOKEN` - GitHub token for Pulumi operations