package schemaorg

import (
	"encoding/json"
	"testing"
)

func TestPerson_FullName(t *testing.T) {
	tests := []struct {
		name string
		p    Person
		want string
	}{
		{
			name: "with additional name",
			p: Person{
				GivenName: "Chun-Ho", AdditionalName: "Hugo", FamilyName: "Lin",
			},
			want: "Chun-Ho Hugo Lin",
		},
		{
			name: "without additional name",
			p: Person{
				GivenName: "Ada", FamilyName: "Lovelace",
			},
			want: "Ada Lovelace",
		},
		{
			name: "only Name field",
			p: Person{
				Name: "Mononym",
			},
			want: "Mononym",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.FullName(); got != tt.want {
				t.Errorf("FullName() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPerson_JSONLD_minimal(t *testing.T) {
	p := Person{GivenName: "A", FamilyName: "B", EmailAddr: "a@b.c"}
	data, err := p.JSONLD()
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if m["@context"] != "https://schema.org" || m["@type"] != "Person" {
		t.Fatalf("context/type: %v", m)
	}
}

func TestPerson_JSONLD_alumniOf_mergesEducationalOrganization(t *testing.T) {
	p := Person{
		GivenName: "A", FamilyName: "B",
		AlumniOf: []Organization{{Name: "Org"}},
		AlumniOfEducational: []EducationalOrganization{{Name: "School"}},
	}
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	raw, ok := m["alumniOf"].([]any)
	if !ok || len(raw) != 2 {
		t.Fatalf("alumniOf = %v", m["alumniOf"])
	}
	first := raw[0].(map[string]any)
	second := raw[1].(map[string]any)
	if first["@type"] != "Organization" || second["@type"] != "EducationalOrganization" {
		t.Fatalf("alumniOf @types: %v, %v", first["@type"], second["@type"])
	}
}
