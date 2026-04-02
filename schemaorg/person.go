package schemaorg

import (
	"encoding/json"
	"strings"
)

// Person is a subset of https://schema.org/Person for structured data (JSON-LD).
// JSON property names follow schema.org; encoding uses [Person.MarshalJSON].
type Person struct {
	GivenName      string `json:"givenName,omitempty"`
	AdditionalName string `json:"additionalName,omitempty"`
	FamilyName     string `json:"familyName,omitempty"`
	// Name overrides the default full name in JSON-LD when non-empty.
	Name        string         `json:"name,omitempty"`
	Pronouns    string         `json:"pronouns,omitempty"`
	EmailAddr   string         `json:"email,omitempty"` // use [Person.Email] as accessor (field name avoids clashing with the method)
	URL         string         `json:"url,omitempty"`
	Image       string         `json:"image,omitempty"`
	Description string         `json:"description,omitempty"`
	JobTitle    string         `json:"jobTitle,omitempty"`
	BirthPlace  string         `json:"birthPlace,omitempty"`
	Nationality string         `json:"nationality,omitempty"`
	SameAs      []string       `json:"sameAs,omitempty"`
	AlumniOf    []Organization `json:"alumniOf,omitempty"`
	// KnowsLanguage uses schema.org Language nodes (see [Language]).
	KnowsLanguage []Language `json:"knowsLanguage,omitempty"`
}

// FullName returns a display name from name parts.
func (p Person) FullName() string {
	if p.AdditionalName != "" {
		return strings.Join([]string{p.GivenName, p.AdditionalName, p.FamilyName}, " ")
	}
	if p.GivenName != "" || p.FamilyName != "" {
		return strings.TrimSpace(p.GivenName + " " + p.FamilyName)
	}
	return strings.TrimSpace(p.Name)
}

// Email returns the person's email address.
func (p Person) Email() string {
	return p.EmailAddr
}

// MarshalJSON encodes Person as schema.org JSON-LD with @context and @type.
func (p Person) MarshalJSON() ([]byte, error) {
	m := map[string]any{
		"@context": "https://schema.org",
		"@type":    "Person",
	}

	displayName := strings.TrimSpace(p.Name)
	if displayName == "" {
		displayName = p.FullName()
	}
	if displayName != "" {
		m["name"] = displayName
	}
	if p.GivenName != "" {
		m["givenName"] = p.GivenName
	}
	if p.AdditionalName != "" {
		m["additionalName"] = p.AdditionalName
	}
	if p.FamilyName != "" {
		m["familyName"] = p.FamilyName
	}
	if p.Pronouns != "" {
		m["pronouns"] = p.Pronouns
	}
	if p.EmailAddr != "" {
		m["email"] = p.EmailAddr
	}
	if p.URL != "" {
		m["url"] = p.URL
	}
	if p.Image != "" {
		m["image"] = p.Image
	}
	if p.Description != "" {
		m["description"] = p.Description
	}
	if p.JobTitle != "" {
		m["jobTitle"] = p.JobTitle
	}
	if p.BirthPlace != "" {
		m["birthPlace"] = p.BirthPlace
	}
	if p.Nationality != "" {
		m["nationality"] = p.Nationality
	}
	if len(p.SameAs) > 0 {
		m["sameAs"] = p.SameAs
	}
	if len(p.AlumniOf) > 0 {
		m["alumniOf"] = p.AlumniOf
	}
	if len(p.KnowsLanguage) > 0 {
		m["knowsLanguage"] = p.KnowsLanguage
	}
	return json.Marshal(m)
}

// JSONLD returns the JSON-LD representation of the person (same as json.Marshal).
func (p Person) JSONLD() ([]byte, error) {
	return json.Marshal(p)
}
