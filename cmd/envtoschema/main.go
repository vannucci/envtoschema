package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	target_flag := flag.String("target", "", "target file to parse into a schema")

	flag.Parse()

	target := *target_flag

	if target == "" {
		fmt.Println("No target specified, ending")
		os.Exit(1)
	}

	file_bytes := ReadFile(target, 100)

	// Workflow 1
	// read config file
	// check size
	// validate json
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
}
