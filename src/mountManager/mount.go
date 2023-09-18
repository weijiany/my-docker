package mountManager

import (
	"fmt"
	"path"
	"syscall"
)

type MountInfo struct {
	source string
	target string
	fstype string
}

var mountInfos = []MountInfo{
	{
		source: "tmpfs",
		target: "/sys",
		fstype: "tmpfs",
	},
	{
		source: "proc",
		target: "/proc",
		fstype: "proc",
	},
}

func Mount() error {
	mountFlags := uintptr(syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV)
	for _, info := range mountInfos {
		if err := syscall.Mount(info.fstype, info.target, info.fstype, mountFlags, ""); err != nil {
			return fmt.Errorf("mount %v error: %v", info.target, err)
		}
	}

	return nil
}

func Umount(baseDir string) error {
	for _, info := range mountInfos {
		if err := syscall.Unmount(path.Join(baseDir, info.target), syscall.MNT_DETACH); err != nil {
			return fmt.Errorf("mount %v error: %v", info.target, err)
		}
	}
	return nil
}
