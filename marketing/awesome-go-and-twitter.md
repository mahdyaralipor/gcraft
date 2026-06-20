# Awesome Go — Pull Request

## مرحله ۱: Fork کن
https://github.com/avelino/awesome-go رو fork کن

## مرحله ۲: خط مناسب رو پیدا کن
فایل `README.md` رو باز کن، بخش **"Generation and Generators"** یا **"Code Generators"** رو پیدا کن (لیست به ترتیب الفبایی مرتبه)

## مرحله ۳: این خط رو اضافه کن (به ترتیب الفبا)

```markdown
* [gcraft](https://github.com/Mahdyaralipor/gcraft) - Generates Builder, Validator, Clone and Mock boilerplate from struct/interface definitions using AST parsing.
```

## مرحله ۴: قوانین رو رعایت کن
- Awesome Go قوانین سخت‌گیرانه‌ای داره: پروژه باید حداقل چند ماه فعال باشه و معمولاً حداقل چندتا star/contributor بخواد
- بهتره PR رو بعد از گرفتن چند star از Reddit/HN بزنی، نه روز اول
- توضیحات باید **خیلی مختصر و خنثی** باشه — نه marketing language
- CONTRIBUTING.md پروژه‌شون رو حتما بخون قبل از PR

## مرحله ۵: PR title
```
Add gcraft to Generation and Generators
```

---

# Twitter / X Post

## نسخه کوتاه (تک توییت)

```
Built gcraft — a CLI that generates Builder, Validator, Clone, and Mock 
boilerplate for Go structs/interfaces using AST parsing.

go install github.com/Mahdyaralipor/gcraft/cmd/gcraft@latest

Zero deps, gofmt-clean output, //go:generate compatible.

https://github.com/Mahdyaralipor/gcraft

#golang
```

## نسخه Thread (۳ توییت)

**Tweet 1:**
```
I kept writing the same Builder/Validator/Mock boilerplate for every 
Go struct. So I built a tool to generate it from the AST.

Introducing gcraft 🧵
```

**Tweet 2 (با کد یا اسکرین‌شات):**
```
gcraft generate -type User -src ./user.go

→ generates a fluent builder, a validator for required fields, 
a deep-copy Clone(), and (for interfaces) a full test mock.

All gofmt-clean, zero runtime deps.
```

**Tweet 3:**
```
v0.1.1 is out — early days, feedback very welcome.

go install github.com/Mahdyaralipor/gcraft/cmd/gcraft@latest

https://github.com/Mahdyaralipor/gcraft

#golang #opensource
```

## نکات:
- بهترین زمان: ۹-۱۱ صبح یا ۱-۳ بعدازظهر به وقت PST، روزهای کاری
- اگه عکس یا GIF داری حتماً attach کن — engagement رو ۲-۳ برابر می‌کنه
- @golang رو می‌تونی mention کنی ولی امیدوار نباش retweet کنه؛ صرفاً دیده میشه
