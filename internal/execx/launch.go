package execx

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Launch(cmdText string) error {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return fmt.Errorf("shell not found")
	}

	cmd := exec.Command(shell, "-lc", cmdText)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
	devNull, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer devNull.Close()

	cmd.Stdin = devNull
	cmd.Stdout = devNull
	cmd.Stderr = devNull

	err = cmd.Start()
	if err != nil {
		return err
	}

	cmd.Process.Release()

	return nil
}
