package schemaorggen

// SchemaContextURL is the JSON-LD "@context" for this vocabulary (same contract as schema-dts WithContext).
//
// For a root document, use an outer struct with Context set to this value and embed your generated
// type (e.g. Person) so fields flatten into one JSON object. See package example_test.go.
const SchemaContextURL = "https://schema.org"
