package schemaorg

import "encoding/json"

// PostalAddress is a subset of https://schema.org/PostalAddress.
type PostalAddress struct {
	StreetAddress   string `json:"streetAddress,omitempty"`
	AddressLocality string `json:"addressLocality,omitempty"`
	AddressRegion   string `json:"addressRegion,omitempty"`
	PostalCode      string `json:"postalCode,omitempty"`
	AddressCountry  string `json:"addressCountry,omitempty"`
	ExtendedAddress string `json:"extendedAddress,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type PostalAddress.
func (a PostalAddress) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "PostalAddress"}
	if a.StreetAddress != "" {
		m["streetAddress"] = a.StreetAddress
	}
	if a.AddressLocality != "" {
		m["addressLocality"] = a.AddressLocality
	}
	if a.AddressRegion != "" {
		m["addressRegion"] = a.AddressRegion
	}
	if a.PostalCode != "" {
		m["postalCode"] = a.PostalCode
	}
	if a.AddressCountry != "" {
		m["addressCountry"] = a.AddressCountry
	}
	if a.ExtendedAddress != "" {
		m["extendedAddress"] = a.ExtendedAddress
	}
	return json.Marshal(m)
}
