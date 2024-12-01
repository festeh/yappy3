package astal

import (
	"os"
	"os/exec"
)

type Astal struct {
}

func (a *Astal) SendMessage(message string) error {
	// Run terminal command astal message
	cmd := exec.Command("astal", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
