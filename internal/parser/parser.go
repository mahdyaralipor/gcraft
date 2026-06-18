// Package parser parses Go source files and extracts type information
// using Go's standard AST (Abstract Syntax Tree) libraries.
package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Field represents a single field in a struct.
type Field struct {
	Name string
	Type string
	Tag  string
}

// Method represents a method signature in an interface.
type Method struct {
	Name    string
	Params  []Field
	Returns []Field
}

// TypeInfo holds all extracted information about a parsed type.
type TypeInfo struct {
	Name       string
	Package    string
	IsStruct   bool
	IsInterface bool
	Fields     []Field  // populated for structs
	Methods    []Method // populated for interfaces
}

// Parse reads a Go source file and extracts type information
// for the given type name.
func Parse(src, typeName string) (*TypeInfo, error) {
	// resolve path — file or directory
	info, err := os.Stat(src)
	if err != nil {
		return nil, fmt.Errorf("cannot access %q: %w", src, err)
	}

	var files []string
	if info.IsDir() {
		files, err = goFilesIn(src)
		if err != nil {
			return nil, err
		}
	} else {
		files = []string{src}
	}

	for _, f := range files {
		ti, err := parseFile(f, typeName)
		if err != nil {
			return nil, err
		}
		if ti != nil {
			return ti, nil
		}
	}

	return nil, fmt.Errorf("type %q not found in %q", typeName, src)
}

// parseFile parses a single .go file looking for typeName.
func parseFile(path, typeName string) (*TypeInfo, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parse error in %s: %w", path, err)
	}

	pkgName := f.Name.Name

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok || typeSpec.Name.Name != typeName {
				continue
			}

			ti := &TypeInfo{
				Name:    typeName,
				Package: pkgName,
			}

			switch t := typeSpec.Type.(type) {
			case *ast.StructType:
				ti.IsStruct = true
				ti.Fields = extractFields(t)
			case *ast.InterfaceType:
				ti.IsInterface = true
				ti.Methods = extractMethods(t)
			default:
				return nil, fmt.Errorf("type %q is neither a struct nor an interface", typeName)
			}

			return ti, nil
		}
	}

	return nil, nil // not found in this file
}

// extractFields pulls all fields from a struct type.
func extractFields(st *ast.StructType) []Field {
	var fields []Field

	for _, f := range st.Fields.List {
		typStr := typeToString(f.Type)
		tag := ""
		if f.Tag != nil {
			tag = strings.Trim(f.Tag.Value, "`")
		}

		if len(f.Names) == 0 {
			// embedded field
			fields = append(fields, Field{
				Name: typStr,
				Type: typStr,
				Tag:  tag,
			})
			continue
		}

		for _, name := range f.Names {
			fields = append(fields, Field{
				Name: name.Name,
				Type: typStr,
				Tag:  tag,
			})
		}
	}

	return fields
}

// extractMethods pulls all method signatures from an interface type.
func extractMethods(it *ast.InterfaceType) []Method {
	var methods []Method

	for _, m := range it.Methods.List {
		if len(m.Names) == 0 {
			continue // embedded interface — skip for now
		}

		fn, ok := m.Type.(*ast.FuncType)
		if !ok {
			continue
		}

		method := Method{
			Name:    m.Names[0].Name,
			Params:  extractFieldList(fn.Params),
			Returns: extractFieldList(fn.Results),
		}
		methods = append(methods, method)
	}

	return methods
}

// extractFieldList converts an *ast.FieldList to []Field.
func extractFieldList(fl *ast.FieldList) []Field {
	if fl == nil {
		return nil
	}

	var fields []Field
	for i, f := range fl.List {
		typStr := typeToString(f.Type)

		if len(f.Names) == 0 {
			fields = append(fields, Field{
				Name: fmt.Sprintf("arg%d", i),
				Type: typStr,
			})
			continue
		}

		for _, name := range f.Names {
			fields = append(fields, Field{
				Name: name.Name,
				Type: typStr,
			})
		}
	}

	return fields
}

// typeToString converts an ast.Expr to its string representation.
func typeToString(expr ast.Expr) string {
	if expr == nil {
		return ""
	}

	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return typeToString(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + typeToString(t.X)
	case *ast.ArrayType:
		if t.Len == nil {
			return "[]" + typeToString(t.Elt)
		}
		return "[...]" + typeToString(t.Elt)
	case *ast.MapType:
		return "map[" + typeToString(t.Key) + "]" + typeToString(t.Value)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.ChanType:
		return "chan " + typeToString(t.Value)
	case *ast.Ellipsis:
		return "..." + typeToString(t.Elt)
	default:
		return fmt.Sprintf("%T", expr)
	}
}

// goFilesIn returns all .go files in a directory (non-recursive).
func goFilesIn(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("cannot read dir %q: %w", dir, err)
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".go") &&
			!strings.HasSuffix(e.Name(), "_test.go") {
			files = append(files, filepath.Join(dir, e.Name()))
		}
	}

	return files, nil
}
