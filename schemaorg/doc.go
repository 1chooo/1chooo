// Package schemaorg models schema.org types used for a personal profile, centered on [Person].
// Values serialize to JSON-LD suitable for <script type="application/ld+json"> or HTTP responses.
//
// # Relationship to Google schema-dts
//
// The [schema-dts] project (TypeScript) generates typings from the official Schema.org ontology
// via [schema-dts-gen] and the N-Triples dump (e.g. schemaorg-all-https.nt). Types such as
// PersonLeaf use "@type": "Person" and property names identical to schema.org (e.g. givenName,
// alumniOf, hasOccupation). TypeScript expresses unions with SchemaValue<T>, Role wrappers, and
// IdReference for @id stubs.
//
// This Go package does not embed schema-dts (it is npm/TypeScript-only). It follows the same
// vocabulary contract: JSON keys match schema.org and [SchemaContextURL] matches schema-dts’s
// WithContext helper. Structs here are a curated, write-oriented subset; optional unions are
// represented as multiple fields (e.g. JobTitle vs JobTitleDefined) resolved in [Person.MarshalJSON].
//
// For exhaustive compile-time coverage in TypeScript/JavaScript, use schema-dts. For Go, package
// schemaorggen (at module path github.com/1chooo/1chooo/schemaorggen) mirrors the full vocabulary
// from the official N-Triples dump; this package remains a small ergonomic layer for hand-written profiles.
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
// [schema-dts]: https://github.com/google/schema-dts
// [schema-dts-gen]: https://www.npmjs.com/package/schema-dts-gen
package schemaorg
