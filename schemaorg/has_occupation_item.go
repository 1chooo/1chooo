package schemaorg

import "encoding/json"

// HasOccupationItem is one entry for Person.hasOccupation (Occupation and/or Role per schema.org examples).
type HasOccupationItem struct {
	Occupation *Occupation
	Role       *Role
}

// MarshalJSON encodes the non-nil occupation or role.
func (h HasOccupationItem) MarshalJSON() ([]byte, error) {
	switch {
	case h.Occupation != nil:
		return json.Marshal(h.Occupation)
	case h.Role != nil:
		return json.Marshal(h.Role)
	default:
		return []byte("null"), nil
	}
}
