package pqjson

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// A StringMap is the Go representation of a JSON object, as stored in a Postgres JSONB field.
type StringMap map[string]interface{}

// Value marshals the map to JSON. If the map is empty, it is marshaled as "{}".
func (sm StringMap) Value() (driver.Value, error) {
	if sm.Empty() {
		return "{}", nil
	}
	return json.Marshal(sm)
}

// Scan converts raw db bytes into a StringMap.
func (sm *StringMap) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return errors.New("source failed type assertion to []byte")
	}
	return json.Unmarshal(source, sm)
}

// Empty returns true if the StringMap's underlying map is empty, false otherwise.
func (sm StringMap) Empty() bool {
	return len(sm) == 0
}
