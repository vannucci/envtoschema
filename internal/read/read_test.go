package read

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("valid file", func(t *testing.T) {
		path := filepath.Join("testdata", "valid.json")
		_, err := ReadFile(path)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
	})

	t.Run("file does not exist", func(t *testing.T) {
		_, err := ReadFile("testdata/nonexistent.json")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("file not readable", func(t *testing.T) {
		path := filepath.Join("testdata", "noperms.json")
		os.Chmod(path, 0000)
		defer os.Chmod(path, 0644)
		_, err := ReadFile(path)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
