package zfs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tobiashienzsch/zfs"
)

func TestCommand(t *testing.T) {
	c := zfs.Command{Command: "pwd"}
	actual, err := c.Run()
	assert.NoError(t, err, "while running command")
	assert.Contains(t, actual[0][0], "/zfs", "")

}
