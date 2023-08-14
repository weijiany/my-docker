package aufs

import (
	"fmt"
	"path/filepath"
	"syscall"
)

type MountInfo struct {
	paths  []string
	fstype string
}

var mountInfos = []MountInfo{
	{
		paths:  []string{"sys"},
		fstype: "sysfs",
	},
	{
		paths:  []string{"proc"},
		fstype: "proc",
	},
}

func join(root string, paths []string) string {
	return filepath.Join(append([]string{root}, paths...)...)
}

func ChangeRoot(root string) error {
	mountFlags := uintptr(syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV)
	for _, info := range mountInfos {
		dirPath := join(root, info.paths)
		if err := syscall.Mount(info.fstype, dirPath, info.fstype, mountFlags, ""); err != nil {
			return fmt.Errorf("mount %v error: %v", dirPath, err)
		}
	}

	if err := syscall.Chroot(root); err != nil {
		return fmt.Errorf("chroot error: %v", err)
	}
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("cd error: %v", err)
	}
	return nil
}

func Unmount() error {
	for _, info := range mountInfos {
		if err := syscall.Unmount(join("/", info.paths), syscall.MNT_DETACH); err != nil {
			return fmt.Errorf("mount %v error: %v", info.paths, err)
		}
	}
	return nil
}
