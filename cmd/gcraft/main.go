// gcraft — Go boilerplate code generator
//
// Usage:
//
//	gcraft generate -type <TypeName> [flags]
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Mahdyaralipor/gcraft/internal/generator"
	"github.com/Mahdyaralipor/gcraft/internal/parser"
)

const version = "0.1.1"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate", "gen":
		if err := runGenerate(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "gcraft: %v\n", err)
			os.Exit(1)
		}
	case "version", "--version", "-v":
		fmt.Printf("gcraft v%s\n", version)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "gcraft: unknown command %q\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func runGenerate(args []string) error {
	fs := flag.NewFlagSet("generate", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprint(os.Stderr, generateUsage)
		fs.PrintDefaults()
	}

	typeName := fs.String("type", "", "type name to generate for (required)")
	src      := fs.String("src", ".", "source file or directory")
	out      := fs.String("out", "", "output file (default: <type>_gen.go in same dir as src)")
	builder  := fs.Bool("builder", true, "generate Builder pattern")
	mock     := fs.Bool("mock", true, "generate Mock (interfaces only)")
	validate := fs.Bool("validate", true, "generate Validator")
	clone    := fs.Bool("clone", true, "generate Clone method")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *typeName == "" {
		fs.Usage()
		return fmt.Errorf("-type is required")
	}

	// resolve output path
	outPath := *out
	if outPath == "" {
		outPath = resolveOutPath(*src, *typeName)
	}

	opts := generator.Options{
		Builder:  *builder,
		Mock:     *mock,
		Validate: *validate,
		Clone:    *clone,
	}

	fmt.Printf("gcraft: parsing %q for type %q...\n", *src, *typeName)

	ti, err := parser.Parse(*src, *typeName)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	fmt.Printf("gcraft: generating code (builder=%v validate=%v clone=%v mock=%v)...\n",
		opts.Builder, opts.Validate, opts.Clone, opts.Mock)

	code, err := generator.Generate(ti, opts)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	if err := os.WriteFile(outPath, []byte(code), 0644); err != nil {
		return fmt.Errorf("write %s: %w", outPath, err)
	}

	fmt.Printf("gcraft: ✓ wrote %s\n", outPath)
	return nil
}

// resolveOutPath builds the default output file path.
func resolveOutPath(src, typeName string) string {
	info, err := os.Stat(src)
	if err != nil || info.IsDir() {
		return toLower(typeName) + "_gen.go"
	}
	// same directory as the source file
	dir := dirOf(src)
	if dir == "" {
		dir = "."
	}
	return dir + "/" + toLower(typeName) + "_gen.go"
}

func dirOf(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			return path[:i]
		}
	}
	return ""
}

func toLower(s string) string {
	b := make([]byte, len(s))
	for i := range s {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		b[i] = c
	}
	return string(b)
}

const generateUsage = `Usage: gcraft generate -type <TypeName> [flags]

Flags:
`

func printUsage() {
	fmt.Print(`gcraft — Go boilerplate code generator

Usage:
  gcraft <command> [flags]

Commands:
  generate    Generate boilerplate for a struct or interface
  version     Print version
  help        Show this help

Examples:
  gcraft generate -type User -src ./user.go
  gcraft generate -type Repository -src ./repo.go -builder=false
  gcraft generate -type Order -src ./models/ -out ./models/order_gen.go

Run 'gcraft generate --help' for all flags.
`)
}
