package schemaorg

import "encoding/json"

// DefinedTerm is a subset of https://schema.org/DefinedTerm (e.g. jobTitle, skills).
type DefinedTerm struct {
	Name             string `json:"name,omitempty"`
	TermCode         string `json:"termCode,omitempty"`
	InDefinedTermSet string `json:"inDefinedTermSet,omitempty"`
	URL              string `json:"url,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type DefinedTerm.
func (d DefinedTerm) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "DefinedTerm"}
	if d.Name != "" {
		m["name"] = d.Name
	}
	if d.TermCode != "" {
		m["termCode"] = d.TermCode
	}
	if d.InDefinedTermSet != "" {
		m["inDefinedTermSet"] = d.InDefinedTermSet
	}
	if d.URL != "" {
		m["url"] = d.URL
	}
	return json.Marshal(m)
}
