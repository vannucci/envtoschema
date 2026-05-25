package validate

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestIsJson(t *testing.T) {
	t.Run("valid file", func(t *testing.T) {
		path := filepath.Join("testdata", "valid.json")
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		err = IsJSON(data)
		if err != nil {
			t.Fatalf("expected nil, got: %v", err)
		}
	})
	t.Run("invalid file", func(t *testing.T) {
		path := filepath.Join("testdata", "invalid.json")
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		err = IsJSON(data)
		if !errors.Is(err, ErrInvalidJson) {
			t.Fatalf("expected ErrInvalidJson, got: %v", err)
		}
	})
	t.Run("valid array", func(t *testing.T) {
		path := filepath.Join("testdata", "valid_array.json")
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		err = IsJSON(data)
		if err != nil {
			t.Fatalf("expected nil, got: %v", err)
		}
	})
}