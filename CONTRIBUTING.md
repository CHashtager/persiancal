# Contributing to PersianCal

First off, thank you for considering contributing to PersianCal! It's people like you that make PersianCal such a great tool.

## Code of Conduct

This project and everyone participating in it is governed by a spirit of respect and collaboration. By participating, you are expected to uphold this spirit.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues to avoid duplicates. When you create a bug report, include as many details as possible:

* **Use a clear and descriptive title**
* **Describe the exact steps to reproduce the problem**
* **Provide specific examples to demonstrate the steps**
* **Describe the behavior you observed and what you expected**
* **Include your Go version and operating system**

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, include:

* **Use a clear and descriptive title**
* **Provide a detailed description of the suggested enhancement**
* **Explain why this enhancement would be useful**
* **List any similar features in other libraries**

### Pull Requests

* Fill in the required template
* Follow the Go coding style
* Include tests for new functionality
* Update documentation as needed
* End all files with a newline

## Development Setup

1. **Fork and clone the repository**
   ```bash
   git clone https://github.com/yourusername/persiancal.git
   cd persiancal
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build the project**
   ```bash
   make build
   # or
   go build ./cmd/persiancal
   ```

4. **Run tests**
   ```bash
   make test
   # or
   go test ./...
   ```

5. **Run the example**
   ```bash
   make example
   # or
   go run example/main.go
   ```

## Project Structure

```
persiancal/
├── pkg/
│   └── persiancal/          # Core library
│       ├── conversion.go    # Gregorian ↔ Jalali conversion
│       ├── time.go          # JalaliDate struct and methods
│       ├── format.go        # Formatting and parsing
│       ├── locale.go        # Month names and localization
│       ├── util.go          # Utility functions
│       └── errors.go        # Error definitions
├── cmd/
│   └── persiancal/          # CLI application
│       ├── main.go          # Entry point
│       └── cmd/             # CLI commands
└── example/                 # Usage examples
```

## Coding Style

* Follow [Effective Go](https://golang.org/doc/effective_go.html)
* Use `go fmt` for formatting
* Use `go vet` to check for common errors
* Write clear, descriptive variable and function names
* Add comments for exported functions and types
* Keep functions small and focused

## Testing

* Write tests for all new functionality
* Ensure all tests pass before submitting a PR
* Aim for high code coverage
* Test edge cases and error conditions

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Generate coverage report
make test-coverage
```

## Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
* Limit the first line to 72 characters or less
* Reference issues and pull requests liberally after the first line

Examples:
```
Add Persian digit support for parsing
Fix leap year calculation for edge cases
Update README with new examples
```

## Documentation

* Update the README.md if you change functionality
* Add examples for new features
* Document all exported functions and types
* Keep documentation clear and concise

## Release Process

Releases are handled by maintainers:

1. Update version in relevant files
2. Update CHANGELOG.md
3. Create a git tag: `git tag -a v1.0.0 -m "Release v1.0.0"`
4. Push the tag: `git push origin v1.0.0`
5. GitHub Actions will automatically build and publish the release

## Questions?

Feel free to open an issue with your question or reach out to the maintainers.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

