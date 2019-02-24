package pqjson

import (
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMapValue(t *testing.T) {
	tt := []struct {
		sm          StringMap
		expectedVal driver.Value
		expectedErr error
	}{
		{
			sm:          StringMap{"foo": "bar"},
			expectedVal: []byte(`{"foo":"bar"}`),
		},
		{
			sm:          StringMap{},
			expectedVal: `{}`,
		},
	}

	for _, tc := range tt {
		val, err := tc.sm.Value()
		assert.Equal(t, tc.expectedVal, val)
		assert.Equal(t, tc.expectedErr, err)
	}
}

func TestStringMapScan(t *testing.T) {
	tt := []struct {
		src         interface{}
		expectedSM  StringMap
		expectedErr error
	}{
		{
			src:        []byte("{}"),
			expectedSM: StringMap{},
		},
		{
			src:        nil,
			expectedSM: nil,
		},
		{
			src:        []byte(`{"foo": "bar"}`),
			expectedSM: StringMap{"foo": "bar"},
		},
		{
			src:         "not a byte array",
			expectedErr: errors.New("source failed type assertion to []byte"),
		},
	}

	for _, tc := range tt {
		var sm StringMap
		err := sm.Scan(tc.src)
		assert.Equal(t, tc.expectedErr, err)
		assert.Equal(t, tc.expectedSM, sm)
	}
}
