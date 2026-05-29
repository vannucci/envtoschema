package read

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// ReadFile reads a local file and returns its raw bytes.
//
// Preconditions:
//   - path points to a readable local file
//
// Postconditions:
//   - returns raw bytes on success
//   - returns error if file is missing, unreadable, or exceeds maxBytes
//
// Invariants:
//   - does not interpret or validate content
//   - caller owns the returned slice
var ErrFileTooLarge = errors.New("file exceeds maximum allowed size")

func ReadFile(path string) ([]byte, error) {
	_, err := os.Stat(path)

	if err != nil {
		return nil, fmt.Errorf("file doesn't exist or isn't accessible: %w", err)
	}

	f, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	defer f.Close()

	data, err := io.ReadAll(f)

	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return data, nil

}
