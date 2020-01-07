package zfs

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
)

// Command helper struct os commands
type Command struct {
	Command string
	Stdin   io.Reader
	Stdout  io.Writer
}

// Run executes the os command & returns the output
func (c *Command) Run(arg ...string) ([][]string, error) {

	cmd := exec.Command(c.Command, arg...)

	var stdout, stderr bytes.Buffer

	if c.Stdout == nil {
		cmd.Stdout = &stdout
	} else {
		cmd.Stdout = c.Stdout
	}

	if c.Stdin != nil {
		cmd.Stdin = c.Stdin

	}
	cmd.Stderr = &stderr

	joinedArgs := strings.Join(cmd.Args, " ")
	err := cmd.Run()

	if err != nil {
		return nil, &Error{
			Err:    err,
			Debug:  strings.Join([]string{cmd.Path, joinedArgs[1:]}, " "),
			Stderr: stderr.String(),
		}
	}

	// assume if you passed in something for stdout, that you know what to do with it
	if c.Stdout != nil {
		return nil, nil
	}

	lines := strings.Split(stdout.String(), "\n")

	//last line is always blank
	lines = lines[0 : len(lines)-1]
	output := make([][]string, len(lines))

	for i, l := range lines {
		output[i] = strings.Fields(l)
	}

	return output, nil
}
