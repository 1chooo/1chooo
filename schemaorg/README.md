# schemaorg

Go types for a **subset** of [schema.org](https://schema.org) used to describe a person ([Person](https://schema.org/Person)) and related nodes (`Organization`, `Language`).

## Purpose

- Hold profile data in Go with a clear mapping to schema.org property names in JSON-LD.
- Emit structured data for search engines and other consumers of JSON-LD.

## Install / import

Module: `github.com/1chooo/1chooo`

```go
import "github.com/1chooo/1chooo/schemaorg"
```

This package is generic schema.org types only. For this author’s filled-in profile, import `github.com/1chooo/1chooo/profile` and call `profile.Hugo()`.

## Usage

```go
p := schemaorg.Person{
    GivenName: "Ada",
    FamilyName: "Lovelace",
    EmailAddr: "ada@example.com",
}
data, err := json.Marshal(p)
```

Or use the author profile:

```go
import "github.com/1chooo/1chooo/profile"

data, err := profile.Hugo().JSONLD()
```

`Person` implements `json.Marshaler` with `@context` and `@type`.

## HTML embedding

```html
<script type="application/ld+json">
PASTE_JSON_HERE
</script>
```

Generate `PASTE_JSON_HERE` from `json.MarshalIndent(profile.Hugo(), "", "  ")` (or any `schemaorg.Person`) if you want readable HTML source.

## Testing

```bash
go test ./schemaorg/... ./profile/...
```

## References

- [Person](https://schema.org/Person)
- [JSON-LD](https://json-ld.org/)
- [Organization](https://schema.org/Organization)
- [Language](https://schema.org/Language)
