package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithoutProg)
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
