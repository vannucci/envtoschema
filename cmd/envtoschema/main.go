package main

import (
	"flag"
	"fmt"
	"os"

	"vannucci.com/envtoschema/m/internal/infer"
	"vannucci.com/envtoschema/m/internal/read"
	"vannucci.com/envtoschema/m/internal/validate"
)

func main() {

	var target string
	var output string
	var mode string

	flag.StringVar(&target, "target", "", "target file to parse into a schema")
	flag.StringVar(&output, "output", "", "output schema name (optional)")
	flag.StringVar(&mode, "mode", "", "1 for json to schema, 2 for validate schema against json")

	flag.Parse()

	if target == "" {
		fmt.Println("No target specified, ending")
		os.Exit(1)
	}

	if output == "" {
		output = "schema.json" // default
	}

	file_bytes, err := read.ReadFile(target, 1)

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
	}

	fmt.Printf("Parsed file: %v\n", parsed_file)

	// Workflow 1
	// parse
	// infer type
	// exposure selection form to user
	// POST form and generate schema

	// Workflow 2
	// read schema
	// read config file
	// validate config satisfies schema
	// if valid, output 1 for CI/CD, else output diff

	// Details
	// Readfile can be either over HTTP from AppConfig SDK, or locally for file

	fmt.Println("Schema generation complete")
	os.Exit(0)
}
