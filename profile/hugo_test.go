package profile

import (
	"encoding/json"
	"testing"
)

func TestHugo_JSONLDShape(t *testing.T) {
	data, err := Hugo().JSONLD()
	if err != nil {
		t.Fatal(err)
	}
	if !json.Valid(data) {
		t.Fatal("JSONLD output is not valid JSON")
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if got, want := m["@context"], "https://schema.org"; got != want {
		t.Errorf("@context = %v, want %v", got, want)
	}
	if got, want := m["@type"], "Person"; got != want {
		t.Errorf("@type = %v, want %v", got, want)
	}
	if got, want := m["givenName"], "Chun-Ho"; got != want {
		t.Errorf("givenName = %v, want %v", got, want)
	}
	if got, want := m["familyName"], "Lin"; got != want {
		t.Errorf("familyName = %v, want %v", got, want)
	}
	sameAs, ok := m["sameAs"].([]any)
	if !ok || len(sameAs) != 3 {
		t.Errorf("sameAs: got %v (len ok=%v), want 3 URLs", m["sameAs"], ok)
	}
}

func TestHugo_Email(t *testing.T) {
	if got, want := Hugo().Email(), "hugo@1chooo.com"; got != want {
		t.Errorf("Email() = %q, want %q", got, want)
	}
}

func TestHugo_AlumniOf_OrganizationType(t *testing.T) {
	data, err := json.Marshal(Hugo())
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	raw, ok := m["alumniOf"].([]any)
	if !ok || len(raw) < 1 {
		t.Fatalf("alumniOf: %v", m["alumniOf"])
	}
	first, ok := raw[0].(map[string]any)
	if !ok {
		t.Fatalf("alumniOf[0] type %T", raw[0])
	}
	if got, want := first["@type"], "Organization"; got != want {
		t.Errorf("alumniOf[0].@type = %v, want %v", got, want)
	}
	if got, want := first["name"], "University of Southern California"; got != want {
		t.Errorf("alumniOf[0].name = %v, want %v", got, want)
	}
}

func TestHugo_KnowsLanguage_LanguageType(t *testing.T) {
	data, err := json.Marshal(Hugo())
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	raw, ok := m["knowsLanguage"].([]any)
	if !ok || len(raw) < 1 {
		t.Fatalf("knowsLanguage: %v", m["knowsLanguage"])
	}
	first, ok := raw[0].(map[string]any)
	if !ok {
		t.Fatalf("knowsLanguage[0] type %T", raw[0])
	}
	if got, want := first["@type"], "Language"; got != want {
		t.Errorf("knowsLanguage[0].@type = %v, want %v", got, want)
	}
}
