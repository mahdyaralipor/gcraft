# gcraft

[![Docs](https://img.shields.io/badge/docs-mahdyaralipor.github.io-blue)](https://mahdyaralipor.github.io/gcraft-docs/)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev)
[![Go Report Card](https://goreportcard.com/badge/github.com/Mahdyaralipor/gcraft)](https://goreportcard.com/report/github.com/Mahdyaralipor/gcraft)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://pkg.go.dev/badge/github.com/Mahdyaralipor/gcraft.svg)](https://pkg.go.dev/github.com/Mahdyaralipor/gcraft)

**gcraft** is a Go code generation tool that reads your `struct` and `interface` definitions and automatically generates idiomatic boilerplate — Builder pattern, Mocks, Validators, and Clone methods — so you can focus on business logic instead of repetitive code.

## Why gcraft?

Writing the same boilerplate for every struct is tedious and error-prone. `gcraft` analyzes your code using Go's AST and generates clean, production-ready code right next to your source files.

```bash
# Before: write 80 lines of boilerplate by hand
# After:
gcraft generate ./...
```

## Features

- 🏗️ **Builder pattern** — fluent builder for any struct
- 🧪 **Mock generation** — interface mocks for testing
- ✅ **Validator** — struct validation boilerplate
- 🔁 **Clone** — deep copy methods
- ⚙️ **go generate** compatible
- 📦 Zero dependencies (uses only Go stdlib)

## Installation

```bash
go install github.com/Mahdyaralipor/gcraft/cmd/gcraft@latest
```

## Quick Start

Given this struct:

```go
// user.go
//go:generate gcraft generate -type User

type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}
```

Run:

```bash
gcraft generate -type User -src ./user.go
```

gcraft generates `user_gen.go`:

```go
// UserBuilder — fluent builder
user := NewUserBuilder().
    WithID(1).
    WithName("Alice").
    WithEmail("alice@example.com").
    Build()

// Validate
if err := user.Validate(); err != nil {
    log.Fatal(err)
}

// Clone
copy := user.Clone()
```

## Usage

```
gcraft generate [flags]

Flags:
  -type     Type name to generate for (required)
  -src      Source file or directory (default: current directory)
  -out      Output file (default: <type>_gen.go)
  -builder  Generate Builder pattern (default: true)
  -mock     Generate Mock for interfaces (default: true)
  -validate Generate Validator (default: true)
  -clone    Generate Clone method (default: true)
```

## Examples

See the [examples/](examples/) directory for complete working examples.

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) first.

## License

MIT — see [LICENSE](LICENSE)
