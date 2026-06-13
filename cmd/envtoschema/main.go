package main

import (
	"flag"
	"fmt"
	"os"

	"vannucci.com/envtoschema/m/internal/infer"
	"vannucci.com/envtoschema/m/internal/server"
)

func main() {

	var targetPath string
	var outputPath string
	// var mode string

	flag.StringVar(&targetPath, "targetPath", "", "target file to parse into a schema")
	flag.StringVar(&outputPath, "outputPath", "", "outputPath schema name (optional)")
	// flag.StringVar(&mode, "mode", "", "1 for json to schema, 2 for validate schema against json")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: envtoschema -target <file> [-output <file>] [-mode <1|2>]\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if targetPath == "" {
		fmt.Println("No target specified, ending")
		os.Exit(1)
	}

	if outputPath == "" {
		outputPath = "schema.json" // default
	}

	parsed_file, err := infer.ParseFlat(targetPath)

	if err != nil {
		fmt.Printf("Error on parsing", err)
	}

	inferred_types := infer.Infer(parsed_file)

	fields := server.ToFieldViews(inferred_types)
	server.Start(server.PageData{Fields: fields}, outputPath)

	fmt.Println("Schema generation complete")
	os.Exit(0)
}
