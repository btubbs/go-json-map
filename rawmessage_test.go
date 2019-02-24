package pqjson

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONMetaEmpty(t *testing.T) {
	tt := []struct {
		desc  string
		input RawMessage
		out   bool
	}{
		{
			desc:  "not empty",
			input: RawMessage("{}"),
			out:   false,
		},
		{
			desc:  "empty",
			input: RawMessage(""),
			out:   true,
		},
	}

	for _, tc := range tt {
		out := tc.input.Empty()
		assert.Equal(t, tc.out, out, tc.desc)
	}
}

func TestJSONMetaValue(t *testing.T) {
	tt := []struct {
		desc  string
		input RawMessage
		out   []byte
		err   error
	}{
		{
			desc:  "empty input",
			input: RawMessage(""),
			out:   nil,
			err:   nil,
		},
		{
			desc:  "non-empty input",
			input: RawMessage("{\"foo\": \"bar\"}"),
			out:   []byte("{\"foo\": \"bar\"}"),
			err:   nil,
		},
	}

	for _, tc := range tt {
		out, err := tc.input.Value()

		if tc.out == nil {
			assert.Nil(t, out, tc.desc)
		} else {
			assert.Equal(t, tc.out, out, tc.desc)
		}
		assert.Equal(t, tc.err, err, tc.desc)
	}
}

func TestJSONMetaScan(t *testing.T) {
	tt := []struct {
		desc     string
		expected RawMessage
		src      interface{}
		meta     RawMessage
		err      error
	}{
		{
			desc:     "nil source",
			expected: RawMessage(""),
			src:      nil,
			meta:     RawMessage(""),
			err:      nil,
		},
		{
			desc:     "source can't be converted to bytes",
			expected: RawMessage(""),
			src:      1123123,
			meta:     RawMessage(""),
			err:      errors.New("source failed type assertion to []byte"),
		},
		{
			desc:     "happy path",
			expected: RawMessage("{\"foo\": \"bar\"}"),
			src:      []byte("{\"foo\": \"bar\"}"),
			meta:     RawMessage(""),
			err:      nil,
		},
	}

	for _, tc := range tt {
		metaPtr := &tc.meta
		err := metaPtr.Scan(tc.src)

		assert.Equal(t, tc.expected, *metaPtr, tc.desc)
		assert.Equal(t, tc.err, err, tc.desc)
	}
}

func TestJSONMetaScanDoesntShareMemory(t *testing.T) {
	meta := RawMessage("")
	metaPtr := &meta
	src := []byte("{\"foo\": \"bar\"}")

	err := metaPtr.Scan(src)
	assert.Nil(t, err)

	assert.Equal(t, RawMessage("{\"foo\": \"bar\"}"), *metaPtr)
	src[0] = 'a'
	assert.Equal(t, RawMessage("{\"foo\": \"bar\"}"), *metaPtr)
}
