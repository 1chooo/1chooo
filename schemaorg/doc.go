// Package schemaorg models a small subset of https://schema.org types for a personal
// profile, centered on [Person]. Values serialize to JSON-LD suitable for
// <script type="application/ld+json"> or HTTP responses.
//
// # Usage
//
// Build a [Person], then marshal to JSON:
//
//	p := schemaorg.Person{GivenName: "Ada", FamilyName: "Lovelace"}
//	data, err := json.Marshal(p)
//
// Author-specific data lives in package github.com/1chooo/1chooo/profile (e.g. profile.Hugo).
//
// See README.md in this directory for HTML embedding and references.
//
// [Person]: https://schema.org/Person
package schemaorg
