#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
NT="${ROOT}/internal/schemaorg-gen/cache/schemaorg-all-https.nt"
mkdir -p "$(dirname "$NT")"
curl -fsSL "https://schema.org/version/latest/schemaorg-all-https.nt" -o "$NT"
go run "${ROOT}/internal/schemaorg-gen" -ontology "$NT" -o "${ROOT}/schemaorggen" -package schemaorggen
echo "Regenerated ${ROOT}/schemaorggen"
