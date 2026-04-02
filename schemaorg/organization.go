package schemaorg

import "encoding/json"

// Organization is a subset of https://schema.org/Organization.
// Encoding uses [Organization.MarshalJSON] for @type.
type Organization struct {
	Name   string `json:"name,omitempty"`
	SameAs string `json:"sameAs,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type Organization.
func (o Organization) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "Organization"}
	if o.Name != "" {
		m["name"] = o.Name
	}
	if o.SameAs != "" {
		m["sameAs"] = o.SameAs
	}
	return json.Marshal(m)
}
