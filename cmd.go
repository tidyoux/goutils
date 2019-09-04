package goutils

import (
	"os/exec"
)

// ExeCmd is a simple wrapper of exec.Cmd.
func ExeCmd(name string, args []string, setter func(*exec.Cmd)) ([]byte, error) {
	c := exec.Command(name, args...)
	if setter != nil {
		setter(c)
	}
	return c.CombinedOutput()
}
