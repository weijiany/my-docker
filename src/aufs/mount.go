package aufs

import (
	"fmt"
	cp "github.com/otiai10/copy"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

const baseDir = "/mydocker/aufs"

func NewWorkSpace(mnt string) {
	readDir := path.Join(baseDir, "read")
	createDir(readDir)

	cp.Copy("/app/busybox", readDir)

	writeDir := path.Join(baseDir, "write")
	createDir(writeDir)

	createMountPoint(readDir, writeDir, mnt)
}

func createMountPoint(readDir string, writeDir string, mnt string) {
	mntPath := path.Join(baseDir, mnt)
	createDir(mntPath)
	dirs := "dirs=" + strings.Join([]string{writeDir, readDir}, ":")
	exec.Command("mount", "-t", "aufs", "-o", dirs, "none", mntPath).Run()
}

func createDir(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 777); err != nil {
			log.Errorf("create dir err: %v", err)
			os.Exit(-1)
		}
	}
}

func DeleteWorkSpace(mnt string) {
	if err := syscall.Unmount(path.Join(baseDir, mnt), syscall.MNT_DETACH); err != nil {
		log.Errorf("unmount mnt err: %v", err)
	}

	if err := os.RemoveAll(path.Join(baseDir, mnt)); err != nil {
		log.Errorf("remove mnt err: %v", err)
	}
	if err := os.RemoveAll(path.Join(baseDir, "write")); err != nil {
		log.Errorf("remove write err: %v", err)
	}
}

func ChangeRoot(mnt string) error {
	if err := syscall.Chroot(path.Join(baseDir, mnt)); err != nil {
		return fmt.Errorf("chroot error: %v", err)
	}
	return syscall.Chdir("/")
}
