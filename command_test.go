package zfs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tobiashienzsch/zfs"
)

func TestCommandRun(t *testing.T) {
	c := zfs.Command{Command: "pwd"}
	actual, err := c.Run()
	assert.NoError(t, err, "while running command")
	assert.Contains(t, actual[0], "/zfs", "")
}

func TestCommandRunCommandDoesntExist(t *testing.T) {
	c := zfs.Command{Command: "qwertgyhuj"}
	actual, err := c.Run()
	assert.Error(t, err, "while running command")
	assert.Nil(t, actual, "")
}
