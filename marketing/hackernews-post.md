# Hacker News — Show HN

## Title
Show HN: gcraft – Generate Builder/Validator/Clone/Mock boilerplate for Go structs

## URL
https://github.com/Mahdyaralipor/gcraft

## First comment (post this immediately after submitting)

Hi HN! I built gcraft because I kept writing nearly identical boilerplate — builders, validators, clone methods, interface mocks — for every struct in my Go projects.

How it works: it parses your source with `go/ast`, extracts struct fields / interface methods, and runs them through Go templates to generate idiomatic code, formatted with `gofmt`. No reflection at runtime — everything is generated as plain `.go` files you can read, diff, and commit.

```bash
gcraft generate -type User -src ./user.go
```

A few design decisions I'd love feedback on:
- Single static binary, zero runtime dependencies (stdlib only)
- Generated code lives in `<type>_gen.go`, designed to work with `//go:generate`
- Mocks use a `FuncField` + `CalledCount` pattern rather than a full mocking framework — simpler, no extra assertions library needed

It's v0.1.0, so there are rough edges — embedded structs and generics aren't fully supported yet. Happy to answer questions or hear what you'd want from a tool like this.

---

## نکات مهم:

1. **زمان:** صبح زود به وقت PST — حدود ۶-۸ صبح (یعنی ۱۷:۳۰-۱۹:۳۰ به وقت ایران، شب قبلش)
2. **عنوان:** نباید تبلیغاتی باشه — HN از عناوین marketing-y بدش میاد. کلمات ساده و دقیق بهتره
3. **مهم‌ترین نکته در HN:** اولین ۲۰-۳۰ دقیقه حیاتیه. اگه چند upvote اولیه بگیره، الگوریتم نشونش میده بقیه. از چند تا دوست بخواه که upvote کنن (نه coordinated voting واضح، فقط organic)
4. **پاسخ سریع:** در HN، مردم سوالای فنی سخت می‌پرسن. آماده باش برای edge cases بپرسن (مثل embedded structs، generics)
5. **فروتنی:** HN از ادعاهای اغراق‌آمیز بدش میاد؛ بگو "rough edges" و "early" — صداقت بهتر جواب میده
