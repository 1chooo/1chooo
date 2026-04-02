package schemaorg

import "encoding/json"

// Organization is a subset of https://schema.org/Organization.
// Encoding uses [Organization.MarshalJSON] for @type.
type Organization struct {
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	SameAs      string `json:"sameAs,omitempty"`
	Description string `json:"description,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type Organization.
func (o Organization) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "Organization"}
	if o.Name != "" {
		m["name"] = o.Name
	}
	if o.URL != "" {
		m["url"] = o.URL
	}
	if o.SameAs != "" {
		m["sameAs"] = o.SameAs
	}
	if o.Description != "" {
		m["description"] = o.Description
	}
	return json.Marshal(m)
}
