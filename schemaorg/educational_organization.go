package schemaorg

import "encoding/json"

// EducationalOrganization is a subset of https://schema.org/EducationalOrganization (e.g. Person.alumniOf).
type EducationalOrganization struct {
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	SameAs      string `json:"sameAs,omitempty"`
	Description string `json:"description,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type EducationalOrganization.
func (e EducationalOrganization) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "EducationalOrganization"}
	if e.Name != "" {
		m["name"] = e.Name
	}
	if e.URL != "" {
		m["url"] = e.URL
	}
	if e.SameAs != "" {
		m["sameAs"] = e.SameAs
	}
	if e.Description != "" {
		m["description"] = e.Description
	}
	return json.Marshal(m)
}
