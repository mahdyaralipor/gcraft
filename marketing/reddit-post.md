# r/golang Post

## Title (یکی از این دو رو انتخاب کن)

Option A: "I built gcraft — a CLI that generates Builder, Validator, Clone and Mock boilerplate from your structs/interfaces"

Option B: "gcraft: stop writing the same Builder/Mock/Validator boilerplate by hand"

## Body

I got tired of writing the same boilerplate over and over for every struct in my projects — Builder pattern, validators, Clone methods, and interface mocks for testing. So I built **gcraft**, a small CLI that parses your Go source using `go/ast` and generates that code for you.

**What it does:**

```bash
gcraft generate -type User -src ./user.go
```

Given:

```go
type User struct {
    ID    int
    Name  string
    Email string
    Tags  []string
}
```

It generates `user_gen.go` with:
- A fluent **Builder** (`NewUserBuilder().WithName("Alice").Build()`)
- A **Validator** that checks required string fields
- A **Clone** method that deep-copies slices
- For interfaces: a full **Mock** with `*Func` and `*Called` fields for testing

All output goes through `gofmt`, zero dependencies (stdlib only).

**Why not just use `go generate` + existing tools?**

Most generators I found either need a config file, only do one thing (mocks OR builders, not both), or pull in heavy dependencies. gcraft is a single binary, works with `//go:generate` out of the box, and does all four in one pass.

It's early (v0.1.1) — I'd genuinely appreciate feedback on the API, edge cases I'm missing, or features you'd actually want. Source + examples here:

👉 https://github.com/Mahdyaralipor/gcraft

Happy to answer questions in the comments.

---

## نکات مهم برای پست کردن:

1. **زمان:** سه‌شنبه یا چهارشنبه، ساعت ۹-۱۱ صبح به وقت EST (یعنی حدود ۱۶:۳۰-۱۸:۳۰ به وقت ایران)
2. **Flair:** اگه ساب‌ردیت گزینه "Show & Tell" یا "Project" داره ازش استفاده کن
3. **اولین کامنت خودت:** بلافاصله بعد از پست، یه کامنت بذار با جزئیات فنی بیشتر (مثلاً چرا AST و نه reflection) — این باعث میشه engagement بالا بره
4. **حتماً پاسخ بده:** به هر کامنت در ۲۴ ساعت اول جواب بده، حتی انتقادی
5. **Self-promotion rule:** قبلش چک کن قوانین r/golang رو بخون (معمولاً self-promo باید کمتر از ۹/۱۰ پست‌هات باشه)
