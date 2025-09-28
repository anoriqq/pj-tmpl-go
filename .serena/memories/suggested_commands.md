# Suggested Commands for pj-tmpl-go

## Development Commands
- `mise tasks` - Display all available tasks with descriptions
- `mise run build` - Build all binaries to `bin/` directory (gitignored)
- `mise run test` - Run all tests with race detection (3s timeout, count=2, shuffle=on)
- `mise run run` - Run the application with hot reloading using Air
- `mise run gen` - Run code generation (go generate ./...)
- `mise run clean` - Remove built binaries
- `mise run lint` - Run all linters (actionlint, ghalint, pinact, zizmor)
- `RELEASE=1 mise run build` - Production build (strips symbols, static linking, no race detector)

## Testing Commands
- `gotest ./...` - Run tests (wrapper around go test)
- `gotest -run TestName` - Run specific test
- `gotest ./internal/infra/cli -update` - Update golden test files

## Project Setup Commands
- `mise install` - Install all required tools
- `lefthook install` - Install git hooks
- `lefthook run pre-commit` - Verify git hooks

## CLI Application Commands
- `./bin/cmd -help` - Show help message
- `./bin/cmd -env dev -port 9000` - Run with specific environment and port
- Environment variables: `ENV=dev PORT=9000 ./bin/cmd`

## System Commands (Darwin)
- `ls` - List files
- `find` - Find files
- `grep` - Search text
- `git` - Git operations