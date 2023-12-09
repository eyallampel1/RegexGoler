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

## Contributing
Contributions to RegexGoler are warmly welcomed. Fork the repository, make your changes, and submit pull requests. Feel free to create issues for bugs or suggest new features.

## License
RegexGoler is available under the [MIT License](LICENSE).

## Acknowledgements
- Thanks to the Go community for the robust programming language and tools.
- Credit to Fatih Arslan for the `fatih/color` package used in this project.