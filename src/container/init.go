package container

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
	"weijiany/docker/src/aufs"
	"weijiany/docker/src/mountManager"
)

func RunContainerInitProcess(cmdArr []string) error {
	if err := aufs.ChangeRoot("mnt"); err != nil {
		return fmt.Errorf("change root err: %v", err)
	}

	if err := mountManager.Mount(); err != nil {
		return err
	}

	cmdPath, err := exec.LookPath(cmdArr[0])
	if err != nil {
		log.Errorf("Exec loop cmdPath error %v", err)
		return err
	}
	return syscall.Exec(cmdPath, cmdArr[0:], os.Environ())
}
