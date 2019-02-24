package pqjson

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// RawMessage wraps json.RawMessage as a type to use for fragments of JSON metadata.   Use this type
// if you need to represent all possible JSON values (including arrays, int literals, etc.). It's
// fine for passing data through from an API to Postgres without needing to touch it in Go, but not
// great if you need to access the data inside Go code. If you need to access the data in Go, and
// you're OK with being restricted to a map[string]interface{}, consider using pqjson.StringMap
// instead.
type RawMessage json.RawMessage

// Value returns the raw byte array for JSON fragment or nil if the JSON raw message is empty.
func (rm RawMessage) Value() (driver.Value, error) {
	if rm.Empty() {
		return nil, nil
	}
	return []byte(rm), nil
}

// Scan converts raw db bytes into a RawMessage.
func (rm *RawMessage) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return errors.New("source failed type assertion to []byte")
	}

	dup := make([]byte, len(source))
	copy(dup, source)

	*rm = RawMessage(json.RawMessage(dup))
	return nil
}

// Empty returns true if the RawMessage's underlying byte slice is empty, false otherwise.
func (rm RawMessage) Empty() bool {
	return string(rm) == ""
}
