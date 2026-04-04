package schemaorg

import (
	"encoding/json"
	"strings"
)

// Person maps to https://schema.org/Person (broad property coverage for JSON-LD).
// Encoding uses [Person.MarshalJSON] for @context, @type, computed name, and unions (image, jobTitle, nationality, etc.).
type Person struct {
	GivenName       string
	AdditionalName  string
	FamilyName      string
	Name            string
	HonorificPrefix string
	HonorificSuffix string

	AlternateName             string
	Description               string
	DisambiguatingDescription string
	Identifier                string
	MainEntityOfPage          string
	AdditionalType            string

	EmailAddr    string
	URL          string
	Image        string
	ImageObject  *ImageObject
	Telephone    string
	FaxNumber    string
	ContactPoint []ContactPoint

	Pronouns  string
	Gender    string
	BirthDate string
	DeathDate string
	// BirthPlace is plain text when BirthPlacePlace is nil (schema prefers Place; text is widely used).
	BirthPlace         string
	BirthPlacePlace    *Place
	DeathPlace         *Place
	Nationality        string
	NationalityCountry *Country

	Address     *PostalAddress
	AddressText string

	JobTitle        string
	JobTitleDefined *DefinedTerm
	HasOccupation   []HasOccupationItem
	WorksFor        []Organization
	Affiliation     *Organization
	Skills          []string
	SkillsDefined   []DefinedTerm
	KnowsAbout      []string

	SameAs   []string
	AlumniOf []Organization
	// AlumniOfEducational is merged into JSON property alumniOf with [Person.AlumniOf].
	AlumniOfEducational []EducationalOrganization
	KnowsLanguage       []Language
	Colleague           []PersonOrURL
	Knows               []*Person
	Follows             []*Person
	Parent              *Person
	Sibling             []*Person
	Children            []*Person
	Spouse              *Person

	MemberOf []Organization
	Award    string
	Awards   []string
	CallSign string

	HomeLocation         *Place
	WorkLocation         *Place
	PublishingPrinciples string

	SponsorOrg    *Organization
	SponsorPerson *Person

	Brand        *Organization
	FunderOrg    *Organization
	FunderPerson *Person

	Height string
	Weight string

	TaxID                string
	VatID                string
	Duns                 string
	GlobalLocationNumber string
	IsicV4               string
	Naics                string

	RelatedTo []*Person

	InteractionStatistic []InteractionCounter
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
		"@context": SchemaContextURL,
		"@type":    "Person",
	}

	displayName := strings.TrimSpace(p.Name)
	if displayName == "" {
		displayName = p.FullName()
	}
	if displayName != "" {
		m["name"] = displayName
	}
	putString(m, "givenName", p.GivenName)
	putString(m, "additionalName", p.AdditionalName)
	putString(m, "familyName", p.FamilyName)
	putString(m, "honorificPrefix", p.HonorificPrefix)
	putString(m, "honorificSuffix", p.HonorificSuffix)
	putString(m, "alternateName", p.AlternateName)
	putString(m, "description", p.Description)
	putString(m, "disambiguatingDescription", p.DisambiguatingDescription)
	putString(m, "identifier", p.Identifier)
	putString(m, "mainEntityOfPage", p.MainEntityOfPage)
	putString(m, "additionalType", p.AdditionalType)
	putString(m, "email", p.EmailAddr)
	putString(m, "url", p.URL)
	putString(m, "telephone", p.Telephone)
	putString(m, "faxNumber", p.FaxNumber)
	putString(m, "pronouns", p.Pronouns)
	putString(m, "gender", p.Gender)
	putString(m, "birthDate", p.BirthDate)
	putString(m, "deathDate", p.DeathDate)
	putString(m, "callSign", p.CallSign)
	putString(m, "publishingPrinciples", p.PublishingPrinciples)
	putString(m, "height", p.Height)
	putString(m, "weight", p.Weight)
	putString(m, "taxID", p.TaxID)
	putString(m, "vatID", p.VatID)
	putString(m, "duns", p.Duns)
	putString(m, "globalLocationNumber", p.GlobalLocationNumber)
	putString(m, "isicV4", p.IsicV4)
	putString(m, "naics", p.Naics)

	if p.ImageObject != nil {
		m["image"] = p.ImageObject
	} else if p.Image != "" {
		m["image"] = p.Image
	}

	if p.BirthPlacePlace != nil {
		m["birthPlace"] = p.BirthPlacePlace
	} else if p.BirthPlace != "" {
		m["birthPlace"] = p.BirthPlace
	}
	if p.DeathPlace != nil {
		m["deathPlace"] = p.DeathPlace
	}

	if p.NationalityCountry != nil {
		m["nationality"] = p.NationalityCountry
	} else if p.Nationality != "" {
		m["nationality"] = p.Nationality
	}

	if p.Address != nil {
		m["address"] = p.Address
	} else if p.AddressText != "" {
		m["address"] = p.AddressText
	}

	if p.JobTitleDefined != nil {
		m["jobTitle"] = p.JobTitleDefined
	} else if p.JobTitle != "" {
		m["jobTitle"] = p.JobTitle
	}

	if occ := marshalHasOccupation(p.HasOccupation); len(occ) > 0 {
		m["hasOccupation"] = occ
	}

	putSlice(m, "sameAs", p.SameAs)
	putSlice(m, "knowsAbout", p.KnowsAbout)
	putAnySlice(m, "skills", marshalSkills(p.Skills, p.SkillsDefined))
	putSlice(m, "knowsLanguage", p.KnowsLanguage)
	if al := marshalAlumni(p.AlumniOf, p.AlumniOfEducational); len(al) > 0 {
		m["alumniOf"] = al
	}
	putSlice(m, "worksFor", p.WorksFor)
	putSlice(m, "memberOf", p.MemberOf)
	putSlice(m, "contactPoint", p.ContactPoint)
	putSlice(m, "interactionStatistic", p.InteractionStatistic)

	if a := marshalAwards(p.Award, p.Awards); len(a) > 0 {
		if len(a) == 1 {
			m["award"] = a[0]
		} else {
			m["award"] = a
		}
	}

	if c := marshalColleagues(p.Colleague); len(c) > 0 {
		m["colleague"] = c
	}

	putOrgPtr(m, "affiliation", p.Affiliation)
	putOrgPtr(m, "brand", p.Brand)
	if p.FunderOrg != nil {
		m["funder"] = p.FunderOrg
	} else if p.FunderPerson != nil {
		m["funder"] = p.FunderPerson
	}
	putPersonSlice(m, "relatedTo", p.RelatedTo)
	putPersonSlice(m, "knows", p.Knows)
	putPersonSlice(m, "follows", p.Follows)
	putPersonPtr(m, "parent", p.Parent)
	putPersonSlice(m, "sibling", p.Sibling)
	putPersonSlice(m, "children", p.Children)
	putPersonPtr(m, "spouse", p.Spouse)
	putPlacePtr(m, "homeLocation", p.HomeLocation)
	putPlacePtr(m, "workLocation", p.WorkLocation)

	if p.SponsorOrg != nil {
		m["sponsor"] = p.SponsorOrg
	} else if p.SponsorPerson != nil {
		m["sponsor"] = p.SponsorPerson
	}

	return json.Marshal(m)
}

// JSONLD returns the JSON-LD representation of the person (same as json.Marshal).
func (p Person) JSONLD() ([]byte, error) {
	return json.Marshal(p)
}

func putString(m map[string]any, key, s string) {
	if s == "" {
		return
	}
	m[key] = s
}

func putSlice[T any](m map[string]any, key string, s []T) {
	if len(s) == 0 {
		return
	}
	m[key] = s
}

func putAnySlice(m map[string]any, key string, s []any) {
	if len(s) == 0 {
		return
	}
	m[key] = s
}

func putOrgPtr(m map[string]any, key string, o *Organization) {
	if o == nil {
		return
	}
	m[key] = o
}

func putPersonPtr(m map[string]any, key string, p *Person) {
	if p == nil {
		return
	}
	m[key] = p
}

func putPersonSlice(m map[string]any, key string, ps []*Person) {
	if len(ps) == 0 {
		return
	}
	m[key] = ps
}

func putPlacePtr(m map[string]any, key string, pl *Place) {
	if pl == nil {
		return
	}
	m[key] = pl
}

func marshalHasOccupation(items []HasOccupationItem) []any {
	var out []any
	for _, it := range items {
		if it.Occupation == nil && it.Role == nil {
			continue
		}
		b, err := json.Marshal(it)
		if err != nil {
			continue
		}
		var v any
		if json.Unmarshal(b, &v) != nil {
			continue
		}
		out = append(out, v)
	}
	return out
}

func marshalSkills(texts []string, terms []DefinedTerm) []any {
	var out []any
	for _, s := range texts {
		out = append(out, s)
	}
	for _, t := range terms {
		out = append(out, t)
	}
	return out
}

func marshalAwards(one string, many []string) []string {
	var out []string
	if one != "" {
		out = append(out, one)
	}
	out = append(out, many...)
	return out
}

func marshalAlumni(orgs []Organization, schools []EducationalOrganization) []any {
	var out []any
	for _, o := range orgs {
		out = append(out, o)
	}
	for _, e := range schools {
		out = append(out, e)
	}
	return out
}

func marshalColleagues(cs []PersonOrURL) []any {
	var out []any
	for _, c := range cs {
		if c.URL == "" && c.Person == nil {
			continue
		}
		b, err := json.Marshal(c)
		if err != nil {
			continue
		}
		var v any
		if json.Unmarshal(b, &v) != nil {
			continue
		}
		out = append(out, v)
	}
	return out
}
