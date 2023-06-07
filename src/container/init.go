package container

import (
	log "github.com/sirupsen/logrus"
	"os"
	"syscall"
)

func RunContainerInitProcess(command string) error {
	log.Infof("command: %s", command)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("none", "/proc", "proc", uintptr(defaultMountFlags), "")
	if err := syscall.Exec(command, []string{command}, os.Environ()); err != nil {
		log.Errorf(err.Error())
	}
	return nil
}
