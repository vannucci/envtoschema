package main

import (
	"flag"
	"fmt"
	"os"

	"vannucci.com/envtoschema/m/internal/infer"
	"vannucci.com/envtoschema/m/internal/read"
	"vannucci.com/envtoschema/m/internal/server"
	"vannucci.com/envtoschema/m/internal/validate"
)

func main() {

	var target string
	var outputPath string
	// var mode string

	flag.StringVar(&target, "target", "", "target file to parse into a schema")
	flag.StringVar(&outputPath, "outputPath", "", "outputPath schema name (optional)")
	// flag.StringVar(&mode, "mode", "", "1 for json to schema, 2 for validate schema against json")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: envtoschema -target <file> [-output <file>] [-mode <1|2>]\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if target == "" {
		fmt.Println("No target specified, ending")
		os.Exit(1)
	}

	if outputPath == "" {
		outputPath = "schema.json" // default
	}

	file_bytes, err := read.ReadFile(target)

	if err != nil {
		fmt.Printf("Error formatting file: %v", err)
		os.Exit(1)
	}

	if err = validate.IsJSON(file_bytes); err != nil {
		fmt.Printf("Error file is not valid JSON: %v", err)
		os.Exit(1)
	}

	parsed_file, err := infer.ParseFlat(file_bytes)

	if err != nil {
		fmt.Printf("Error parsing file: %v", err)
		os.Exit(1)
	}

	inferred_types := infer.Infer(parsed_file)

	fmt.Printf("Inferred types: %v\n", inferred_types)

	fields := server.ToFieldViews(inferred_types)
	server.Start(server.PageData{Fields: fields}, outputPath)

	// Workflow 1
	// exposure selection form to user
	// POST form and generate schema

	// Workflow 2
	// read schema
	// read config file
	// validate config satisfies schema
	// if valid, outputPath 1 for CI/CD, else outputPath diff

	// Details
	// Readfile can be either over HTTP from AppConfig SDK, or locally for file

	fmt.Println("Schema generation complete")
	os.Exit(0)
}
