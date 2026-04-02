package schemaorg

import "encoding/json"

// Place is a subset of https://schema.org/Place (used for birthPlace, deathPlace, workLocation, etc.).
type Place struct {
	Name    string         `json:"name,omitempty"`
	Address *PostalAddress `json:"address,omitempty"`
	SameAs  string         `json:"sameAs,omitempty"`
	URL     string         `json:"url,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type Place.
func (pl Place) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "Place"}
	if pl.Name != "" {
		m["name"] = pl.Name
	}
	if pl.Address != nil {
		m["address"] = pl.Address
	}
	if pl.SameAs != "" {
		m["sameAs"] = pl.SameAs
	}
	if pl.URL != "" {
		m["url"] = pl.URL
	}
	return json.Marshal(m)
}

// Country is a subset of https://schema.org/Country (for nationality).
type Country struct {
	Name   string `json:"name,omitempty"`
	SameAs string `json:"sameAs,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type Country.
func (c Country) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "Country"}
	if c.Name != "" {
		m["name"] = c.Name
	}
	if c.SameAs != "" {
		m["sameAs"] = c.SameAs
	}
	return json.Marshal(m)
}
