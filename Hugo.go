// Hugo.go

package Hugo

// see https://schema.org/Organization
type Organization struct {
	name   string
	sameAs string
}

type Language struct {
	name          string
	alternateName string
}

// Schema.org type for Person
//
// This is a simplified version of the schema.org Person type.
// It includes only the properties that are relevant for this application.
//
// @see https://schema.org/Person
type Person struct {
	givenName      string
	additionalName string
	familyName     string
	pronouns       string
	email          string
	birthPlace     string
	nationality    string
	sameAs         []string
	alumniOf       []Organization
	knowsLanguage  []Language
}

func Hugo() Person {
	return Person{
		givenName:      "Chun-Ho",
		additionalName: "Hugo",
		familyName:     "Lin",
		pronouns:       "he/him",
		email:          "hugo@1chooo.com",
		birthPlace:     "Taiwan",
		nationality:    "Taiwanese",
		sameAs: []string{
			"https://www.linkedin.com/in/1chooo/",
			"http://github.com/1chooo",
			"https://medium.com/@1chooo",
		},
		alumniOf: []Organization{
			{
				name:   "University of Southern California",
				sameAs: "https://usc.edu",
			},
			{
				name:   "National Central University",
				sameAs: "https://www.ncu.edu.tw",
			},
		},
		knowsLanguage: []Language{
			{
				name:          "English",
				alternateName: "en",
			},
			{
				name:          "Mandarin",
				alternateName: "zh",
			},
		},
	}
}

func (p Person) FullName() string {
	if p.additionalName != "" {
		return p.givenName + " " + p.additionalName + " " + p.familyName
	}
	return p.givenName + " " + p.familyName
}

func (p Person) Email() string {
	return p.email
}
