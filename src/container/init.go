package container

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
	"weijiany/docker/src/mountManager"
)

func RunContainerInitProcess(cmdArr []string) error {
	if err := changeRoot("/mydocker/aufs/mnt"); err != nil {
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

func changeRoot(root string) error {
	if err := syscall.Chroot(root); err != nil {
		return fmt.Errorf("chroot error: %v", err)
	}
	return syscall.Chdir("/")
}
