package generator_test

import (
	"strings"
	"testing"

	"github.com/Mahdyaralipor/gcraft/internal/generator"
	"github.com/Mahdyaralipor/gcraft/internal/parser"
)

// ── helpers ───────────────────────────────────────────────────────────────────

func structType(name string, fields ...parser.Field) *parser.TypeInfo {
	return &parser.TypeInfo{
		Name:     name,
		Package:  "myapp",
		IsStruct: true,
		Fields:   fields,
	}
}

func ifaceType(name string, methods ...parser.Method) *parser.TypeInfo {
	return &parser.TypeInfo{
		Name:        name,
		Package:     "myapp",
		IsInterface: true,
		Methods:     methods,
	}
}

func field(name, typ string) parser.Field { return parser.Field{Name: name, Type: typ} }

func mustContain(t *testing.T, src, substr string) {
	t.Helper()
	if !strings.Contains(src, substr) {
		t.Errorf("expected output to contain %q\n\ngot:\n%s", substr, src)
	}
}

func mustNotContain(t *testing.T, src, substr string) {
	t.Helper()
	if strings.Contains(src, substr) {
		t.Errorf("expected output NOT to contain %q\n\ngot:\n%s", substr, src)
	}
}

// ── Builder tests ─────────────────────────────────────────────────────────────

func TestGenerate_Builder(t *testing.T) {
	ti := structType("User",
		field("ID", "int"),
		field("Name", "string"),
		field("Email", "string"),
	)

	out, err := generator.Generate(ti, generator.Options{Builder: true})
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	mustContain(t, out, "type UserBuilder struct")
	mustContain(t, out, "func NewUserBuilder() *UserBuilder")
	mustContain(t, out, "func (b *UserBuilder) WithID(v int) *UserBuilder")
	mustContain(t, out, "func (b *UserBuilder) WithName(v string) *UserBuilder")
	mustContain(t, out, "func (b *UserBuilder) WithEmail(v string) *UserBuilder")
	mustContain(t, out, "func (b *UserBuilder) Build() User")
	mustContain(t, out, "DO NOT EDIT")
}

func TestGenerate_Builder_OnlyBuilder(t *testing.T) {
	ti := structType("Product", field("Price", "float64"))
	out, err := generator.Generate(ti, generator.Options{Builder: true})
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	mustContain(t, out, "ProductBuilder")
	mustNotContain(t, out, "Validate()")
	mustNotContain(t, out, "Clone()")
}

// ── Validator tests ───────────────────────────────────────────────────────────

func TestGenerate_Validator(t *testing.T) {
	ti := structType("User",
		field("ID", "int"),
		field("Name", "string"),
		field("Email", "string"),
	)

	out, err := generator.Generate(ti, generator.Options{Validate: true})
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	mustContain(t, out, "func (v User) Validate() error")
	mustContain(t, out, `User.Name is required`)
	mustContain(t, out, `User.Email is required`)
	// int fields should NOT produce required checks
	mustNotContain(t, out, `User.ID is required`)
}

func TestGenerate_Validator_NoStringFields(t *testing.T) {
	ti := structType("Point",
		field("X", "float64"),
		field("Y", "float64"),
	)

	out, err := generator.Generate(ti, generator.Options{Validate: true})
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	mustContain(t, out, "func (v Point) Validate() error")
	mustContain(t, out, "return nil")
}

// ── Clone tests ───────────────────────────────────────────────────────────────

func TestGenerate_Clone(t *testing.T) {
	ti := structType("User",
		field("ID", "int"),
		field("Name", "string"),
		field("Tags", "[]string"),
	)

	out, err := generator.Generate(ti, generator.Options{Clone: true})
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	mustContain(t, out, "func (v User) Clone() User")
	mustContain(t, out, "copy(c.Tags, v.Tags)")
}

func TestGenerate_Clone_NoSlices(t *testing.T) {
	ti := structType("Config",
		field("Debug", "bool"),
		field("Port", "int"),
	)

	out, err := generator.Generate(ti, generator.Options{Clone: true})
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	mustContain(t, out, "func (v Config) Clone() Config")
	mustNotContain(t, out, "copy(")
}

// ── Mock tests ────────────────────────────────────────────────────────────────

func TestGenerate_Mock(t *testing.T) {
	ti := ifaceType("UserRepository",
		parser.Method{
			Name:    "FindByID",
			Params:  []parser.Field{field("id", "int")},
			Returns: []parser.Field{field("arg0", "*User"), field("arg1", "error")},
		},
		parser.Method{
			Name:    "Save",
			Params:  []parser.Field{field("u", "*User")},
			Returns: []parser.Field{field("arg0", "error")},
		},
	)

	out, err := generator.Generate(ti, generator.Options{Mock: true})
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	mustContain(t, out, "type MockUserRepository struct")
	mustContain(t, out, "FindByIDFunc")
	mustContain(t, out, "func(id int) (*User, error)")
	mustContain(t, out, "FindByIDCalled int")
	mustContain(t, out, "func (m *MockUserRepository) FindByID(id int) (*User, error)")
	mustContain(t, out, "m.FindByIDCalled++")
	mustContain(t, out, "SaveFunc")
}

// ── All-in-one test ───────────────────────────────────────────────────────────

func TestGenerate_AllOptions(t *testing.T) {
	ti := structType("Order",
		field("ID", "int"),
		field("CustomerName", "string"),
		field("Items", "[]string"),
	)

	out, err := generator.Generate(ti, generator.DefaultOptions())
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	mustContain(t, out, "OrderBuilder")
	mustContain(t, out, "Validate()")
	mustContain(t, out, "Clone()")
}

// ── Package header test ───────────────────────────────────────────────────────

func TestGenerate_PackageHeader(t *testing.T) {
	ti := structType("Foo", field("X", "int"))
	out, err := generator.Generate(ti, generator.Options{Builder: true})
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	mustContain(t, out, "package myapp")
	mustContain(t, out, "// Code generated by gcraft.")
}
