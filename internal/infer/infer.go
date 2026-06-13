package infer

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	Key   string
	Value string
}

var ErrNonFlatJSON = errors.New("JSON is not flat")
var ErrEmptyValue = errors.New("value is empty")
var ErrNullValue = errors.New("null value is not allowed")
var ErrArrayValue = errors.New("array value is not allowed")
var ErrInvalidJson = errors.New("file is not valid JSON")
var ErrReadingFile = errors.New("error reading file")
var ErrFileDoesNotExist = errors.New("file does not exist")
var ErrUnmarshalFile = errors.New("failed to unmarshal file")
var ErrUnmarsalLine = errors.New("failed to unmarshal line")

func IsJSON(data []byte) error {
	if json.Valid(data) {
		return nil
	}
	return ErrInvalidJson

}

func ParseFlat(path string) ([]Entry, error) {
	var dataMap map[string]json.RawMessage
	var returnEntries []Entry

	_, err := os.Stat(path)

	if err != nil {
		return nil, ErrFileDoesNotExist
	}

	f, err := os.Open(path)

	if err != nil {
		return nil, ErrFileDoesNotExist
	}

	defer f.Close()

	data, err := io.ReadAll(f)

	if err != nil {
		return nil, ErrReadingFile
	}

	if err := json.Unmarshal(data, &dataMap); err != nil {
		return nil, ErrUnmarshalFile
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
			return nil, ErrUnmarsalLine
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
