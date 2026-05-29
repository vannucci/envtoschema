package infer

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ParseFlat takes in our byte stream and turns it into a slice of KV pairs that
// we will run through Infer to guess at their type

// Preconditions:
// 		- input is a byte stream and is valid JSON

// Postconditions:
// 		- returns a slice of non-empty Entries
// 		- returns error if an object is found, violating flatness condition
// 		- returns error if an empty value is found

// Invariants:
// 		- empty JSON object {} returns empty slice, nil error
// 		- caller is responsible for deciding if empty result is an error

type Entry struct {
	Key   string
	Value string
}

var ErrNonFlatJSON = errors.New("JSON is not flat")
var ErrEmptyValue = errors.New("value is empty")
var ErrNullValue = errors.New("null value is not allowed")
var ErrArrayValue = errors.New("array value is not allowed")

func ParseFlat(data []byte) ([]Entry, error) {
	var dataMap map[string]json.RawMessage
	var returnEntries []Entry
	if err := json.Unmarshal(data, &dataMap); err != nil {
		return nil, fmt.Errorf("error unmarshaling: %w", err)
	}

	for k, v := range dataMap {
		if v[0] == '{' {
			return nil, ErrNonFlatJSON
		}
		if v[0] == '[' {
			return nil, ErrArrayValue
		}
		if string(v) == "null" {
			return nil, ErrNullValue
		}
		var s string
		if err := json.Unmarshal(v, &s); err != nil {
			return nil, fmt.Errorf("error unmarshaling: %w", err)
		}
		if s == "" {
			return nil, ErrEmptyValue
		}
		returnEntries = append(returnEntries, Entry{
			Key:   k,
			Value: s,
		})
	}
	return returnEntries, nil
}

// InferType examines a raw string value from the env file and returns the most
// specific type candidate possible

// Preconditions:
// 		- input is a non-empty string (empty strings handled upstream)

// Postconditions:
// 		- returns exactly one TypeCandidate
//		- never returns an error

// Ambiguity rules:
// 		- "true" / "false" (case-insensitive) -> { Primary: TypeBool, }
// 		- valid integer string -> { Primary: TypeInt, }
//		- valid float string -> { Primary: TypeFloat, }
// 		- anything else -> { Primary: TypeString, }

// Invariants:
// 		- same input always produces same output (pure, no side effects)
//		- caller is responsible for resolving TypeAmbiguous via user input

type Type int

const (
	TypeString Type = iota
	TypeInt
	TypeFloat
	TypeBool
)

type TypeCandidate struct {
	Primary Type
}

func InferType(value string) TypeCandidate {
	lower := strings.ToLower(value)
	if lower == "true" || lower == "false" {
		return TypeCandidate{
			Primary: TypeBool,
		}
	}

	if _, err := strconv.Atoi(value); err == nil {
		return TypeCandidate{
			Primary: TypeInt,
		}
	}

	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return TypeCandidate{
			Primary: TypeFloat,
		}
	}

	return TypeCandidate{
		Primary: TypeString,
	}
}

type InferredElement struct {
	Key       string
	Candidate TypeCandidate
}

func Infer(entries []Entry) []InferredElement {
	results := make([]InferredElement, 0, len(entries))
	for _, e := range entries {
		results = append(results, InferredElement{
			Key:       e.Key,
			Candidate: InferType(e.Value),
		})
	}
	return results
}
func TypeToString(t Type) string {
	switch t {
	case TypeInt:
		return "integer"
	case TypeFloat:
		return "number"
	case TypeBool:
		return "boolean"
	default:
		return "string"
	}
}
