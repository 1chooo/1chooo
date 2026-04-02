package schemaorg

import "encoding/json"

// Occupation is a subset of https://schema.org/Occupation (Person.hasOccupation).
type Occupation struct {
	Name                  string `json:"name,omitempty"`
	EducationRequirements string `json:"educationRequirements,omitempty"`
	OccupationalCategory  string `json:"occupationalCategory,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type Occupation.
func (o Occupation) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "Occupation"}
	if o.Name != "" {
		m["name"] = o.Name
	}
	if o.EducationRequirements != "" {
		m["educationRequirements"] = o.EducationRequirements
	}
	if o.OccupationalCategory != "" {
		m["occupationalCategory"] = o.OccupationalCategory
	}
	return json.Marshal(m)
}

// Role is a subset of https://schema.org/Role for dated or named roles (e.g. past jobs under hasOccupation).
type Role struct {
	RoleName      string      `json:"roleName,omitempty"`
	StartDate     string      `json:"startDate,omitempty"`
	EndDate       string      `json:"endDate,omitempty"`
	HasOccupation *Occupation `json:"hasOccupation,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type Role.
func (r Role) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "Role"}
	if r.RoleName != "" {
		m["roleName"] = r.RoleName
	}
	if r.StartDate != "" {
		m["startDate"] = r.StartDate
	}
	if r.EndDate != "" {
		m["endDate"] = r.EndDate
	}
	if r.HasOccupation != nil {
		m["hasOccupation"] = r.HasOccupation
	}
	return json.Marshal(m)
}
