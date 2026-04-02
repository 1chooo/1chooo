package profile

import "github.com/1chooo/1chooo/schemaorg"

// Hugo returns the author's schema.org Person (see [schemaorg.Person]).
func Hugo() schemaorg.Person {
	return schemaorg.Person{
		GivenName:      "Chun-Ho",
		AdditionalName: "Hugo",
		FamilyName:     "Lin",
		URL:            "https://github.com/1chooo",
		Pronouns:       "he/him",
		EmailAddr:      "hugo@1chooo.com",
		BirthPlace:     "Taiwan",
		Nationality:    "Taiwanese",
		SameAs: []string{
			"https://www.linkedin.com/in/1chooo/",
			"https://github.com/1chooo",
			"https://medium.com/@1chooo",
		},
		AlumniOf: []schemaorg.Organization{
			{Name: "University of Southern California", SameAs: "https://usc.edu"},
			{Name: "National Central University", SameAs: "https://www.ncu.edu.tw"},
		},
		KnowsLanguage: []schemaorg.Language{
			{Name: "English", AlternateName: "en"},
			{Name: "Mandarin", AlternateName: "zh"},
		},
	}
}
