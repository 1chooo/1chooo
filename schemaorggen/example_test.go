package schemaorggen_test

import (
	"encoding/json"
	"fmt"

	"github.com/1chooo/1chooo/schemaorggen"
)

func ExamplePerson_withContext() {
	email := "ada@example.com"
	p := schemaorggen.Person{
		JSONLDType: schemaorggen.TypePerson,
		GivenName:  ptr("Ada"),
		FamilyName: ptr("Lovelace"),
		Email:      &email,
	}
	doc := struct {
		Context string `json:"@context"`
		schemaorggen.Person
	}{
		Context: schemaorggen.SchemaContextURL,
		Person:  p,
	}
	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	// Output:
	// {
	//   "@context": "https://schema.org",
	//   "@type": "Person",
	//   "email": "ada@example.com",
	//   "familyName": "Lovelace",
	//   "givenName": "Ada"
	// }
}

func ptr(s string) *string { return &s }
