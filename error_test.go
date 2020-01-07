package zfs_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tobiashienzsch/zfs"
)

func TestError(t *testing.T) {
	var tests = []struct {
		name   string
		err    error
		debug  string
		stderr string
	}{
		{"empty", nil, "", ""},
		{"typical", errors.New("exit status foo"), "/sbin/foo bar qux", "command not found"},
		{"quoted", errors.New("exit status quoted"), "\"/sbin/foo\" bar qux", "\"some\" 'random' `quotes`"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Generate error from tests
			zErr := zfs.Error{
				Err:    test.err,
				Debug:  test.debug,
				Stderr: test.stderr,
			}

			// Verify output format is consistent, so that any changes to the
			// Error method must be reflected by the test
			expected := fmt.Sprintf("%s: %q => %s", test.err, test.debug, test.stderr)
			actual := zErr.Error()

			assert.Equal(t, expected, actual, "they should be equal")

		})
	}
}
