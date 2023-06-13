package subsystems

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

type MemorySubSystem struct{}

func (m *MemorySubSystem) Name() string {
	return "memory"
}

func (m *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	subSysCgroupPath, err := cgroupPathOf(m.Name(), cgroupPath)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(subSysCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644)
	if err != nil {
		return fmt.Errorf("set cgroup memory failed: %v", err)
	}
	return nil
}

func (m *MemorySubSystem) Remove(cgroupPath string) error {
	subSysCgroupPath, err := cgroupPathOf(m.Name(), cgroupPath)
	if err != nil {
		return err
	}
	return os.RemoveAll(subSysCgroupPath)
}

func (m *MemorySubSystem) Apply(cgroupPath string, pid int) error {
	subSysCgroupPath, err := cgroupPathOf(m.Name(), cgroupPath)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(subSysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		return fmt.Errorf("set cgroup proc failed: %v", err)
	}
	return nil
}
