package subsystems

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

type CgroupManager struct {
	CgroupPath string
	SubSystems []*SubSystem
}

func NewCgroupManager(cgroupPath string, res *ResourceConfig) *CgroupManager {
	return &CgroupManager{
		CgroupPath: cgroupPath,
		SubSystems: []*SubSystem{
			memoryOf(res),
			cpuSharesOf(res),
		},
	}
}

func (cm *CgroupManager) Set() error {
	for _, system := range cm.SubSystems {
		subSysCgroupPath, err := cgroupPathOf(system.Name, cm.CgroupPath)
		if err != nil {
			return err
		}
		err = os.WriteFile(path.Join(subSysCgroupPath, system.FileName), []byte(system.Value), 0644)
		if err != nil {
			return fmt.Errorf("set cgroup %v failed: %v", system.Name, err)
		}
	}
	return nil
}

func (cm *CgroupManager) Apply(pid int) error {
	for _, system := range cm.SubSystems {
		subSysCgroupPath, err := cgroupPathOf(system.Name, cm.CgroupPath)
		if err != nil {
			return err
		}
		err = os.WriteFile(path.Join(subSysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644)
		if err != nil {
			return fmt.Errorf("set cgroup proc failed: %v", err)
		}
	}
	return nil
}

func (cm *CgroupManager) Destroy() error {
	for _, system := range cm.SubSystems {
		subSysCgroupPath, err := cgroupPathOf(system.Name, cm.CgroupPath)
		if err != nil {
			return err
		}
		return os.RemoveAll(subSysCgroupPath)
	}
	return nil
}
