package schemaorg

import "encoding/json"

// ContactPoint is a subset of https://schema.org/ContactPoint (Person.contactPoint).
type ContactPoint struct {
	ContactType       string `json:"contactType,omitempty"`
	Telephone         string `json:"telephone,omitempty"`
	Email             string `json:"email,omitempty"`
	URL               string `json:"url,omitempty"`
	AvailableLanguage string `json:"availableLanguage,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type ContactPoint.
func (c ContactPoint) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "ContactPoint"}
	if c.ContactType != "" {
		m["contactType"] = c.ContactType
	}
	if c.Telephone != "" {
		m["telephone"] = c.Telephone
	}
	if c.Email != "" {
		m["email"] = c.Email
	}
	if c.URL != "" {
		m["url"] = c.URL
	}
	if c.AvailableLanguage != "" {
		m["availableLanguage"] = c.AvailableLanguage
	}
	return json.Marshal(m)
}
