package parser_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Mahdyaralipor/gcraft/internal/parser"
)

// writeTemp writes content to a temp .go file and returns its path.
func writeTemp(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "input.go")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("writeTemp: %v", err)
	}
	return path
}

// ── Struct tests ──────────────────────────────────────────────────────────────

func TestParse_SimpleStruct(t *testing.T) {
	src := writeTemp(t, `package myapp

type User struct {
	ID    int
	Name  string
	Email string
}
`)
	ti, err := parser.Parse(src, "User")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !ti.IsStruct {
		t.Error("expected IsStruct = true")
	}
	if ti.IsInterface {
		t.Error("expected IsInterface = false")
	}
	if ti.Package != "myapp" {
		t.Errorf("package: got %q, want %q", ti.Package, "myapp")
	}
	if len(ti.Fields) != 3 {
		t.Fatalf("fields: got %d, want 3", len(ti.Fields))
	}

	cases := []struct{ name, typ string }{
		{"ID", "int"},
		{"Name", "string"},
		{"Email", "string"},
	}
	for i, c := range cases {
		f := ti.Fields[i]
		if f.Name != c.name {
			t.Errorf("field[%d].Name: got %q, want %q", i, f.Name, c.name)
		}
		if f.Type != c.typ {
			t.Errorf("field[%d].Type: got %q, want %q", i, f.Type, c.typ)
		}
	}
}

func TestParse_StructWithPointerAndSlice(t *testing.T) {
	src := writeTemp(t, `package myapp

type Order struct {
	ID       int
	Items    []string
	Customer *User
	Meta     map[string]string
}
`)
	ti, err := parser.Parse(src, "Order")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []struct{ name, typ string }{
		{"ID", "int"},
		{"Items", "[]string"},
		{"Customer", "*User"},
		{"Meta", "map[string]string"},
	}

	if len(ti.Fields) != len(want) {
		t.Fatalf("fields: got %d, want %d", len(ti.Fields), len(want))
	}
	for i, w := range want {
		f := ti.Fields[i]
		if f.Name != w.name || f.Type != w.typ {
			t.Errorf("field[%d]: got {%s %s}, want {%s %s}", i, f.Name, f.Type, w.name, w.typ)
		}
	}
}

func TestParse_StructWithTags(t *testing.T) {
	src := writeTemp(t, `package myapp

type Product struct {
	ID   int    `+"`json:\"id\"`"+`
	Name string `+"`json:\"name\" validate:\"required\"`"+`
}
`)
	ti, err := parser.Parse(src, "Product")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ti.Fields[0].Tag != `json:"id"` {
		t.Errorf("tag[0]: got %q", ti.Fields[0].Tag)
	}
	if ti.Fields[1].Tag != `json:"name" validate:"required"` {
		t.Errorf("tag[1]: got %q", ti.Fields[1].Tag)
	}
}

// ── Interface tests ───────────────────────────────────────────────────────────

func TestParse_Interface(t *testing.T) {
	src := writeTemp(t, `package myapp

type Repository interface {
	FindByID(id int) (*User, error)
	Save(u *User) error
	Delete(id int) error
}
`)
	ti, err := parser.Parse(src, "Repository")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !ti.IsInterface {
		t.Error("expected IsInterface = true")
	}
	if len(ti.Methods) != 3 {
		t.Fatalf("methods: got %d, want 3", len(ti.Methods))
	}
	if ti.Methods[0].Name != "FindByID" {
		t.Errorf("method[0].Name: got %q, want %q", ti.Methods[0].Name, "FindByID")
	}
}

// ── Error cases ───────────────────────────────────────────────────────────────

func TestParse_TypeNotFound(t *testing.T) {
	src := writeTemp(t, `package myapp

type User struct{ ID int }
`)
	_, err := parser.Parse(src, "Ghost")
	if err == nil {
		t.Error("expected error for missing type, got nil")
	}
}

func TestParse_InvalidPath(t *testing.T) {
	_, err := parser.Parse("/nonexistent/path.go", "Foo")
	if err == nil {
		t.Error("expected error for invalid path, got nil")
	}
}
