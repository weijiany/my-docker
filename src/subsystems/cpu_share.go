package subsystems

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

type CPUShareSubSystem struct{}

func (c *CPUShareSubSystem) Name() string {
	return "cpu"
}

func (c *CPUShareSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	subSysCgroupPath, err := cgroupPathOf(c.Name(), cgroupPath)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(subSysCgroupPath, "cpu.shares"), []byte(res.CpuShare), 0644)
	if err != nil {
		return fmt.Errorf("set cgroup cpu share failed: %v", err)
	}
	return nil
}

func (c *CPUShareSubSystem) Remove(cgroupPath string) error {
	subSysCgroupPath, err := cgroupPathOf(c.Name(), cgroupPath)
	if err != nil {
		return err
	}
	return os.RemoveAll(subSysCgroupPath)
}

func (c *CPUShareSubSystem) Apply(cgroupPath string, pid int) error {
	subSysCgroupPath, err := cgroupPathOf(c.Name(), cgroupPath)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(subSysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		return fmt.Errorf("set cgroup proc failed: %v", err)
	}
	return nil
}
