package schemaorg

import "encoding/json"

// Language is a subset of https://schema.org/Language for Person.knowsLanguage.
// Encoding uses [Language.MarshalJSON] for @type.
type Language struct {
	Name          string `json:"name,omitempty"`
	AlternateName string `json:"alternateName,omitempty"` // BCP 47, e.g. "en", "zh-Hant"
}

// MarshalJSON emits a JSON-LD node with @type Language.
func (l Language) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "Language"}
	if l.Name != "" {
		m["name"] = l.Name
	}
	if l.AlternateName != "" {
		m["alternateName"] = l.AlternateName
	}
	return json.Marshal(m)
}
