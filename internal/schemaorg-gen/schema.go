package main

import (
	"fmt"
	"path"
	"sort"
	"strings"
	"unicode"
)

const (
	rdfType            = "http://www.w3.org/1999/02/22-rdf-syntax-ns#type"
	rdfsSubClassOf     = "http://www.w3.org/2000/01/rdf-schema#subClassOf"
	rdfsClass          = "http://www.w3.org/2000/01/rdf-schema#Class"
	owlClass           = "http://www.w3.org/2002/07/owl#Class"
	rdfProperty        = "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property"
	schemaDomain       = "https://schema.org/domainIncludes"
	schemaRange        = "https://schema.org/rangeIncludes"
	schemaSupersededBy = "https://schema.org/supersededBy"
	schemaDataType     = "https://schema.org/DataType"
	schemaThing        = "https://schema.org/Thing"
	schemaPrefix       = "https://schema.org/"
)

// rangeGoType maps schema range IRIs to Go pointer types used when all ranges are scalar.
var rangeGoType = map[string]string{
	"https://schema.org/Text":            "*string",
	"https://schema.org/URL":             "*string",
	"https://schema.org/Date":            "*string",
	"https://schema.org/DateTime":        "*string",
	"https://schema.org/Time":            "*string",
	"https://schema.org/Pronunciation":   "*string",
	"https://schema.org/CssSelectorType": "*string",
	"https://schema.org/XPathType":       "*string",
	"https://schema.org/Boolean":         "*bool",
	"https://schema.org/Number":          "*float64",
	"https://schema.org/Float":           "*float64",
	"https://schema.org/Integer":         "*int",
}

type schemaModel struct {
	Classes    []string // sorted schema.org class IRIs to emit
	Properties map[string]*propInfo
	Ancestors  map[string]map[string]struct{} // class -> ancestor set (incl. self)
}

type propInfo struct {
	URI         string
	JSONName    string
	Domains     map[string]struct{}
	Ranges      map[string]struct{}
	Superseded  bool
	GoField     string
	GoType      string
}

func buildModel(tris []triple) (*schemaModel, error) {
	var classes []string
	classSet := map[string]struct{}{}
	subClass := map[string][]string{} // child -> parents
	propTypes := map[string]struct{}{}
	domainOf := map[string][]string{} // prop -> domain types
	rangeOf := map[string][]string{}  // prop -> range types
	superseded := map[string]struct{}{}
	superProp := map[string]string{} // old -> new

	addClass := func(uri string) {
		if !strings.HasPrefix(uri, schemaPrefix) {
			return
		}
		if _, ok := classSet[uri]; ok {
			return
		}
		classSet[uri] = struct{}{}
		classes = append(classes, uri)
	}

	for _, t := range tris {
		if !strings.HasPrefix(t.Subj, schemaPrefix) {
			continue
		}
		switch t.Pred {
		case rdfType:
			if t.Obj == rdfsClass || t.Obj == owlClass {
				addClass(t.Subj)
			}
			if t.Obj == rdfProperty {
				propTypes[t.Subj] = struct{}{}
			}
		case rdfsSubClassOf:
			if strings.HasPrefix(t.Obj, schemaPrefix) {
				subClass[t.Subj] = append(subClass[t.Subj], t.Obj)
			}
		case schemaDomain:
			if strings.HasPrefix(t.Obj, schemaPrefix) {
				domainOf[t.Subj] = append(domainOf[t.Subj], t.Obj)
			}
		case schemaRange:
			if strings.HasPrefix(t.Obj, schemaPrefix) {
				rangeOf[t.Subj] = append(rangeOf[t.Subj], t.Obj)
			}
		case schemaSupersededBy:
			superProp[t.Subj] = t.Obj
		}
	}
	for _, t := range tris {
		switch t.Pred {
		case schemaDomain, schemaRange:
			if strings.HasPrefix(t.Obj, schemaPrefix) {
				addClass(t.Obj)
			}
		}
	}

	// Mark superseded properties (as subject of supersededBy).
	for p := range superProp {
		superseded[p] = struct{}{}
	}

	dataTypeDesc := descendantsOf(subClass, schemaDataType)
	dataTypeDesc[schemaDataType] = struct{}{}

	emitClass := map[string]struct{}{}
	for _, c := range classes {
		if _, skip := dataTypeDesc[c]; skip {
			continue
		}
		emitClass[c] = struct{}{}
	}

	ancestors := map[string]map[string]struct{}{}
	var walk func(c string) map[string]struct{}
	walk = func(c string) map[string]struct{} {
		if m, ok := ancestors[c]; ok {
			return m
		}
		m := map[string]struct{}{c: {}}
		for _, p := range subClass[c] {
			for a := range walk(p) {
				m[a] = struct{}{}
			}
		}
		ancestors[c] = m
		return m
	}
	for c := range emitClass {
		walk(c)
	}

	props := map[string]*propInfo{}
	for pURI := range propTypes {
		if _, gone := superseded[pURI]; gone {
			continue
		}
		doms := uniq(domainOf[pURI])
		rngs := uniq(rangeOf[pURI])
		if len(doms) == 0 || len(rngs) == 0 {
			continue
		}
		jsonName := path.Base(pURI)
		if jsonName == "." || jsonName == "" {
			continue
		}
		gf, err := goFieldName(jsonName)
		if err != nil {
			continue
		}
		gt := goTypeForRanges(rngs, emitClass, dataTypeDesc)
		props[pURI] = &propInfo{
			URI: pURI, JSONName: jsonName, GoField: gf, GoType: gt,
			Domains: sliceToSet(doms), Ranges: sliceToSet(rngs),
		}
	}

	var classList []string
	for c := range emitClass {
		classList = append(classList, c)
	}
	sort.Strings(classList)

	return &schemaModel{
		Classes:    classList,
		Properties: props,
		Ancestors:  ancestors,
	}, nil
}

func descendantsOf(subClass map[string][]string, root string) map[string]struct{} {
	children := map[string][]string{}
	for ch, ps := range subClass {
		for _, p := range ps {
			children[p] = append(children[p], ch)
		}
	}
	out := map[string]struct{}{}
	var dfs func(string)
	dfs = func(n string) {
		if _, ok := out[n]; ok {
			return
		}
		out[n] = struct{}{}
		for _, ch := range children[n] {
			dfs(ch)
		}
	}
	dfs(root)
	return out
}

func goTypeForRanges(rngs []string, emitClass map[string]struct{}, dataTypeDesc map[string]struct{}) string {
	if len(rngs) == 0 {
		return "any"
	}
	allScalar := true
	var single string
	for _, r := range rngs {
		gt, ok := rangeGoType[r]
		if !ok {
			if _, dt := dataTypeDesc[r]; dt {
				gt = "*string"
			} else {
				allScalar = false
				break
			}
		}
		if single == "" {
			single = gt
		} else if single != gt {
			// mixed scalars -> any
			allScalar = false
			break
		}
	}
	if allScalar && single != "" {
		return single
	}
	return "any"
}

func propertiesForClass(m *schemaModel, classURI string) []*propInfo {
	anc := m.Ancestors[classURI]
	var uris []string
	for u := range m.Properties {
		uris = append(uris, u)
	}
	sort.Strings(uris)
	var list []*propInfo
	seen := map[string]struct{}{}
	for _, u := range uris {
		p := m.Properties[u]
		for d := range p.Domains {
			if _, ok := anc[d]; ok {
				if _, dup := seen[p.JSONName]; dup {
					continue
				}
				seen[p.JSONName] = struct{}{}
				list = append(list, p)
				break
			}
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].JSONName < list[j].JSONName })
	return list
}

func goFieldName(jsonName string) (string, error) {
	if jsonName == "" {
		return "", fmt.Errorf("empty")
	}
	runes := []rune(jsonName)
	if !unicode.IsLetter(runes[0]) && runes[0] != '_' {
		return "X_" + jsonName, nil
	}
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes), nil
}

func goTypeName(classURI string) (string, error) {
	local := path.Base(classURI)
	if local == "" || local == "." {
		return "", fmt.Errorf("bad class uri %s", classURI)
	}
	runes := []rune(local)
	if !unicode.IsLetter(runes[0]) {
		return "X_" + local, nil
	}
	return local, nil
}

func sliceToSet(s []string) map[string]struct{} {
	m := map[string]struct{}{}
	for _, x := range s {
		m[x] = struct{}{}
	}
	return m
}

func uniq(s []string) []string {
	seen := map[string]struct{}{}
	var out []string
	for _, x := range s {
		if _, ok := seen[x]; ok {
			continue
		}
		seen[x] = struct{}{}
		out = append(out, x)
	}
	sort.Strings(out)
	return out
}
