package infer

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestParseFlat(t *testing.T) {
	t.Run("valid flat file", func(t *testing.T) {
		path := filepath.Join("testdata", "valid.json")
		file_bytes, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("expected no read errors, got: %v", err)
		}
		_, err = ParseFlat(file_bytes)
		if err != nil {
			t.Fatalf("expected no errors, got: %v", err)
		}
	})

	t.Run("invalid non-flat file", func(t *testing.T) {
		path := filepath.Join("testdata", "non_flat.json")
		file_bytes, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("expected no read errors, got: %v", err)
		}
		_, err = ParseFlat(file_bytes)
		if !errors.Is(err, ErrNonFlatJSON) {
			t.Fatalf("expected non flat error, got: %v", err)
		}
	})

	t.Run("empty but valid file", func(t *testing.T) {
		path := filepath.Join("testdata", "empty_json.json")
		file_bytes, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("expected no read errors, got: %v", err)
		}
		entries, err := ParseFlat(file_bytes)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if len(entries) != 0 {
			t.Errorf("expected empty file, got %d", len(entries))
		}
	})
	t.Run("valid file empty value", func(t *testing.T) {
		path := filepath.Join("testdata", "empty_value.json")
		file_bytes, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("expected no read errors, got: %v", err)
		}
		_, err = ParseFlat(file_bytes)
		if !errors.Is(err, ErrEmptyValue) {
			t.Fatalf("expected empty values error, got: %v", err)
		}
	})
	t.Run("file with null value", func(t *testing.T) {
		path := filepath.Join("testdata", "null_value.json")
		file_bytes, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("expected no read errors, got: %v", err)
		}
		_, err = ParseFlat(file_bytes)
		if !errors.Is(err, ErrNullValue) {
			t.Fatalf("expected ErrNullValue, got: %v", err)
		}
	})

	t.Run("file with array value", func(t *testing.T) {
		path := filepath.Join("testdata", "array_value.json")
		file_bytes, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("expected no read errors, got: %v", err)
		}
		_, err = ParseFlat(file_bytes)
		if !errors.Is(err, ErrArrayValue) {
			t.Fatalf("expected ErrArrayValue, got: %v", err)
		}
	})

}

func TestInferType(t *testing.T) {
	t.Run("TypeBool", func(t *testing.T) {
		expect := TypeCandidate{
			Primary: TypeBool,
			Options: []Type{TypeBool, TypeString},
		}
		candidate := InferType("true")
		if candidate.Primary != expect.Primary {
			t.Fatalf("expected primary TypeBool, got: %v", candidate)
		}
		if !slices.Equal(candidate.Options, expect.Options) {
			t.Fatalf("expected options TypeBool, got: %v", candidate)
		}
	})

	t.Run("TypeString", func(t *testing.T) {
		expect := TypeCandidate{
			Primary: TypeString,
			Options: []Type{TypeString},
		}
		candidate := InferType("foo")
		if candidate.Primary != expect.Primary {
			t.Fatalf("expected primary TypeString, got: %v", candidate)
		}
		if !slices.Equal(candidate.Options, expect.Options) {
			t.Fatalf("expected options TypeString, got: %v", candidate)
		}
	})
	t.Run("TypeInt", func(t *testing.T) {
		expect := TypeCandidate{
			Primary: TypeInt,
			Options: []Type{TypeInt, TypeFloat, TypeString},
		}
		candidate := InferType("100")
		if candidate.Primary != expect.Primary {
			t.Fatalf("expected primary TypeInt, got: %v", candidate)
		}
		if !slices.Equal(candidate.Options, expect.Options) {
			t.Fatalf("expected options TypeInt, got: %v", candidate)
		}
	})
	t.Run("TypeFloat", func(t *testing.T) {
		expect := TypeCandidate{
			Primary: TypeFloat,
			Options: []Type{TypeFloat, TypeInt, TypeString},
		}
		candidate := InferType("42.555")
		if candidate.Primary != expect.Primary {
			t.Fatalf("expected primary TypeFloat, got: %v", candidate)
		}
		if !slices.Equal(candidate.Options, expect.Options) {
			t.Fatalf("expected options TypeFloat, got: %v", candidate)
		}
	})
}
