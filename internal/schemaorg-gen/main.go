// Command schemaorg-gen generates Go structs from schema.org N-Triples (schema-dts-gen style source).
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ontology := flag.String("ontology", "", "path to schemaorg-all-https.nt (download from schema.org/version/latest/)")
	outDir := flag.String("o", "", "output directory (e.g. schemaorggen)")
	pkg := flag.String("package", "schemaorggen", "Go package name")
	flag.Parse()
	if *ontology == "" || *outDir == "" {
		fmt.Fprintln(os.Stderr, "usage: schemaorg-gen -ontology file.nt -o outdir [-package name]")
		flag.PrintDefaults()
		os.Exit(2)
	}
	f, err := os.Open(*ontology)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	tris, err := parseNT(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	m, err := buildModel(tris)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	absOut, _ := filepath.Abs(*outDir)
	if err := emit(m, absOut, *pkg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "wrote %d types to %s\n", len(m.Classes), absOut)
}
