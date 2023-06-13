package subsystems

import log "github.com/sirupsen/logrus"

type CgroupManager struct {
	cgroupPath string
	resources  []Subsystem
}

func NewCgroupManager(cgroupPath string) *CgroupManager {
	return &CgroupManager{
		cgroupPath: cgroupPath,
		resources: []Subsystem{
			&MemorySubSystem{},
			&CPUShareSubSystem{},
		},
	}
}

func (cm *CgroupManager) Set(res *ResourceConfig) error {
	for _, sys := range cm.resources {
		err := sys.Set(cm.cgroupPath, res)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cm *CgroupManager) Apply(pid int) error {
	for _, sys := range cm.resources {
		err := sys.Apply(cm.cgroupPath, pid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cm *CgroupManager) Destroy() error {
	for _, sys := range cm.resources {
		err := sys.Remove(cm.cgroupPath)
		if err != nil {
			log.Warnf("Remove cgroup fail, err: %v, name: %v", err, sys.Name())
		}
	}
	return nil
}
