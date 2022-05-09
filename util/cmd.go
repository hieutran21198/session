package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
)

func bindStdOut(c *exec.Cmd) (stdout *bytes.Buffer) {
	b := bytes.Buffer{}

	if c.Stdout != nil {
		c.Stdout = io.MultiWriter(c.Stdout, &b)
	} else {
		c.Stdout = io.MultiWriter(&b)
	}

	return &b
}

// GetJSONOutputFromCMD used to get the output of a command.
func GetJSONOutputFromCMD(cmd string, output interface{}, args ...string) error {
	c := exec.Command(cmd, args...)
	fmt.Println(c.String())

	stdout := bindStdOut(c)

	if err := c.Run(); err != nil {
		return err
	}

	if err := json.NewDecoder(stdout).Decode(output); err != nil {
		return err
	}

	return nil
}
