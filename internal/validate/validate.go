package validate

import (
	"encoding/json"
	"errors"
)

// IsJSON takes in a byte stream and confirms it is a valid JSON type file

// Preconditions:
// 		- input is a valid byte stream of the appropriate size

// Postconditions:
// 		- returns an error iff the input is not JSON, nil if it is

// Invariants:
// 		- Does not comment on how to fix file
// 		- Does not parse or comment on contents

var ErrInvalidJson = errors.New("file is not valid JSON")

func IsJSON(data []byte) error {
	if json.Valid(data) {
		return nil
	}
	return ErrInvalidJson

}