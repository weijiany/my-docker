package subsystems

type SubSystem struct {
	Name     string
	FileName string
	Value    string
}

type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
}

func memoryOf(res *ResourceConfig) *SubSystem {
	return &SubSystem{
		Name:     "memory",
		FileName: "memory.limit_in_bytes",
		Value:    res.MemoryLimit,
	}
}

func cpuSharesOf(res *ResourceConfig) *SubSystem {
	return &SubSystem{
		Name:     "cpu,",
		FileName: "cpu.shares",
		Value:    res.CpuShare,
	}
}

type SubSystemInput struct {
	Cpu    string
	Memory string
}
