package read

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("valid file", func(t *testing.T) {
		path := filepath.Join("testdata", "valid.json")
		data, err := ReadFile(path, 1024)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if len(data) == 0 {
			t.Fatal("expected data, got empty slice")
		}
	})

	t.Run("file does not exist", func(t *testing.T) {
		_, err := ReadFile("testdata/nonexistent.json", 1024)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("file too large", func(t *testing.T) {
		path := filepath.Join("testdata", "valid.json")
		_, err := ReadFile(path, 1)
		if !errors.Is(err, ErrFileTooLarge) {
			t.Fatalf("expected ErrFileTooLarge, got: %v", err)
		}
	})

	t.Run("file not readable", func(t *testing.T) {
		path := filepath.Join("testdata", "noperms.json")
		os.Chmod(path, 0000)
		defer os.Chmod(path, 0644)
		_, err := ReadFile(path, 1024)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}