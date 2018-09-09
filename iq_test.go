package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_runQuery(t *testing.T) {
	cases := []struct {
		name       string
		givenFile  []byte
		givenQuery *query
		expected   string
	}{
		{
			name:       "empty file",
			givenFile:  []byte(""),
			givenQuery: &query{Section: "foo", Key: "bar"},
			expected:   "",
		},
		{
			name:       "absent section",
			givenFile:  []byte("bar=baz\n"),
			givenQuery: &query{Section: "foo", Key: "bar"},
			expected:   "",
		},
		{
			name:       "absent key",
			givenFile:  []byte("[foo]"),
			givenQuery: &query{Section: "foo", Key: "bar"},
			expected:   "",
		},
		{
			name:       "empty value",
			givenFile:  []byte("[foo]\nbar=\n"),
			givenQuery: &query{Section: "foo", Key: "bar"},
			expected:   "",
		},
		{
			name:       "normal",
			givenFile:  []byte("[foo]\nbar=baz\n"),
			givenQuery: &query{Section: "foo", Key: "bar"},
			expected:   "baz",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := runQuery(tc.givenQuery, tc.givenFile)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
