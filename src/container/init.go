package container

import (
	"fmt"
	"os"
	"syscall"
)

func RunContainerInitProcess(command string) error {
	defaultMountFlags := uintptr(syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV)
	if err := syscall.Mount("proc", "/proc", "proc", defaultMountFlags, ""); err != nil {
		return fmt.Errorf("mount proc error: %v", err)
	}

	if err := syscall.Exec(command, []string{}, os.Environ()); err != nil {
		return fmt.Errorf("run init command error: %v", err)
	}
	return nil
}
