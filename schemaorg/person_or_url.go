package schemaorg

import "encoding/json"

// PersonOrURL is a colleague (or similar) value that schema.org allows as Person or URL.
type PersonOrURL struct {
	URL    string
	Person *Person
}

// MarshalJSON encodes either a URL string or a nested Person node.
func (p PersonOrURL) MarshalJSON() ([]byte, error) {
	switch {
	case p.URL != "":
		return json.Marshal(p.URL)
	case p.Person != nil:
		return json.Marshal(p.Person)
	default:
		return []byte("null"), nil
	}
}
