package tools

import (
	"bytes"
	"errors"
	"os/exec"
)

//	pipe several commands into each other
//	returns a byte array of the output or an error
func PipeCommands(cmds ...*exec.Cmd) ([]byte, error) {
	if len(cmds) < 1 {
		return nil, errors.New("no commands provided")
	}

	// Collect the output from the command(s)
	var output bytes.Buffer
	var stderr bytes.Buffer

	last := len(cmds) - 1
	for i, cmd := range cmds[:last] {
		var err error
		// Connect each command's stdin to the previous command's stdout
		if cmds[i+1].Stdin, err = cmd.StdoutPipe(); err != nil {
			return nil, err
		}
		// Connect each command's stderr to a buffer
		cmd.Stderr = &stderr
	}

	// Connect the output and error for the last command
	cmds[last].Stdout, cmds[last].Stderr = &output, &stderr

	// Start each command
	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return output.Bytes(), err
		}
	}

	// Wait for each command to complete
	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			return output.Bytes(), err
		}
	}

	// Return the pipeline output and the collected standard error
	return output.Bytes(), nil
}
