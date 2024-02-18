# getheaders ![build-release-binary](https://github.com/rnemeth90/getheaders/actions/workflows/build.yaml/badge.svg) ![Go Report Card](https://goreportcard.com/badge/github.com/rnemeth90/getheaders/)

## Description
`getheaders` is a simple and efficient tool written in Go that fetches and displays the HTTP headers of a given URL. It supports output in both JSON and YAML formats and offers a pretty-print option for JSON output. This tool is designed to help developers, security researchers, and enthusiasts quickly check and analyze the headers sent by web servers, which can be crucial for debugging, security assessments, and understanding web server configurations.

## Getting Started

### Dependencies
- To build the project yourself, you must have Go version 1.13 or higher installed on your machine. This ensures compatibility with module management and recent language features.

### Installing
- You can clone the repository to your local machine using Git:
  ```
  git clone https://github.com/rnemeth90/getheaders.git
  ```
- Navigate to the cloned directory and build the project using:
  ```
  go build -o getheaders
  ```
- Or download the latest release from the [Releases](https://github.com/rnemeth90/getheaders/releases) page.

### Executing program
- To use `getheaders`, run it from the command line with the desired options. Here are some examples:
  ```
  ./getheaders --url https://example.com --json
  ./getheaders --url https://example.com --yaml
  ./getheaders --url https://example.com --json --pretty
  ```
- For more information on available flags, use the `--help` flag:
  ```
  ./getheaders --help
  ```

## Help
If you encounter any issues or have questions, please feel free to [submit an issue](https://github.com/rnemeth90/getheaders/issues) on GitHub.

## To Do
- [ ] Enhance error handling for non-successful HTTP status codes.
- [ ] Include unit tests for key functions.

## Version History
* 1.0.0
    * Initial Release

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
