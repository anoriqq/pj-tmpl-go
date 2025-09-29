# Task Completion Checklist

## When a coding task is completed, run these commands:

### Testing
1. `mise run test` - Run all tests with race detection
2. `gotest ./path/to/specific/package` - Run specific package tests if needed
3. `gotest ./path/to/package -update` - Update golden files if needed

### Linting
1. `mise run lint` - Run all linters (actionlint, ghalint, pinact, zizmor)
2. Check for any linting errors and fix them

### Code Generation
1. `mise run gen` - Run code generation if domain models were modified
2. Verify generated files are correct

### Build Verification
1. `mise run build` - Verify the project builds successfully
2. `RELEASE=1 mise run build` - Verify release build works (if applicable)

### Git Hooks
1. Ensure lefthook is installed: `lefthook install`
2. Pre-commit hooks will run automatically and must pass

### Additional Checks
- Verify no sensitive information is committed
- Ensure commit messages follow conventional commit format
- Check that all new code follows project conventions
- Update documentation if public APIs changed