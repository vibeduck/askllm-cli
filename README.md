# askllm-cli

A simple yet powerful Command-Line Interface (CLI) tool that enables users to ask ad-hoc questions to an OpenAI-compatible API directly from the command line.

## Features

- Simple command-line interface for interacting with OpenAI-compatible APIs
- Secure API key handling via environment variables or command-line flags
- Cross-platform support (Linux, macOS, Raspberry Pi)
- User-friendly error messages and help documentation

## Installation

### From Source

```bash
git clone https://github.com/vibeduck/askllm-cli.git
cd askllm-cli
go build
```

### From Releases

Download the appropriate binary for your platform from the [GitHub Releases](https://github.com/vibeduck/askllm-cli/releases) page.

## Usage

```bash
# Using environment variable
OPENAI_API_KEY=your-api-key askllm "Write hello world program in Python"

# Using command-line flag
askllm --openai-api-key=your-api-key "Write hello world program in Python"
```

## Configuration

The API key can be provided in two ways:
1. Environment variable: `OPENAI_API_KEY`
2. Command-line flag: `--openai-api-key`

## Development

### Prerequisites

- Go 1.21 or later
- Make (optional, for using Makefile commands)

### Building

```bash
go build
```

### Testing

```bash
go test ./...
```

### Linting

```bash
golangci-lint run
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 