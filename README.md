# RegexGoler

## Overview
RegexGoler is a command-line interface (CLI) tool written in Go, designed for highlighting regex patterns in text files. Utilizing Go's robust regex package and the `fatih/color` library for terminal output, it offers a user-friendly way to visualize regex matches.

## Features
- Efficient regex pattern matching in text files.
- Highlighting of matched regex patterns with alternating colors in terminal output.
- Intuitive CLI experience.
- Supports a wide range of regex patterns.
- Fast and efficient processing, leveraging Go's performance.

## Getting Started

### Prerequisites
- Make sure you have Go installed on your system. [Go Installation Guide](https://golang.org/doc/install).

### Installation
Clone the RegexGoler repository and build the project:
```bash
git clone https://github.com/eyallampel1/RegexGoler.git
cd RegexGoler
go build
```

### Usage
To use RegexGoler, run the compiled executable with a file path and a regex pattern:
```bash
./RegexGoler <file_path> <regex_pattern>
```
Example:
```bash
./RegexGoler sample.txt "hell"
```
To format and lint your Go code, you can use `gofmt` for formatting and tools like `golint` or `staticcheck` for linting. Here's how to use them in your project:

### Code Formatting with `gofmt` or `goimports`
- **gofmt**: This tool automatically formats Go source code.
   - To format all the Go files in your current directory, run:
     ```bash
     gofmt -w .
     ```
   - The `-w` flag tells `gofmt` to write changes to the files instead of printing them.

- **goimports**: This tool not only formats your code but also adds (or removes) import lines as necessary.
   - To install `goimports`, run:
     ```bash
     go install golang.org/x/tools/cmd/goimports@latest
     ```
   - To run `goimports` on your project, use:
     ```bash
     goimports -w .
     ```

### Linting with `golint` or `staticcheck`
- **golint**: This tool prints out style mistakes in your Go code.
   - To install `golint`, run:
     ```bash
     go install golang.org/x/lint/golint@latest
     ```
   - Then, you can lint your code with:
     ```bash
     golint ./...
     ```

- **staticcheck**: This is a more advanced linter that includes checks for performance, correctness, and style issues.
   - Install `staticcheck` with:
     ```bash
     go install honnef.co/go/tools/cmd/staticcheck@latest
     ```
   - To run `staticcheck` on your project:
     ```bash
     staticcheck ./...
     ```

How to use the Makefile:

1. **Build**: To build your project and place the executable in `bin`, simply run:
   ```
   make build
   ```

2. **Clean**: To clean up any build artifacts (like removing the executable), use:
   ```
   make clean
   ```

3. **Run**: To build and run your project in one step, use:
   ```
   make run
   ```
## Contributing
Contributions to RegexGoler are warmly welcomed. Fork the repository, make your changes, and submit pull requests. Feel free to create issues for bugs or suggest new features.

## License
RegexGoler is available under the [MIT License](LICENSE).

## Acknowledgements
- Thanks to the Go community for the robust programming language and tools.
- Credit to Fatih Arslan for the `fatih/color` package used in this project.