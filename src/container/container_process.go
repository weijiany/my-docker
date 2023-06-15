package container

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
	"weijiany/docker/src/subsystems"
)

func newParentProcess(tty bool, command string) *exec.Cmd {
	cmd := exec.Command("unshare", "--fork", "--pid", "--mount-proc=/proc", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | // Unix Timesharing System 用于隔离和管理主机的主机名和域名信息
			syscall.CLONE_NEWPID | // Process Identifier 用于隔离和管理进程的标识符（PID）
			syscall.CLONE_NEWNS | // Mount Namespace Group 用于隔离和管理文件系统挂载点
			syscall.CLONE_NEWNET | // Network Namespace 用于隔离和管理网络栈和网络资源
			syscall.CLONE_NEWIPC, // Inter-Process Communication 用于隔离和管理进程间通信的资源，如消息队列、信号量和共享内存等
	}
	cmd.Env = os.Environ()
	if tty {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
	return cmd
}

func Run(tty bool, command string) {
	parent := newParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err.Error())
		os.Exit(-1)
	}
	cm := subsystems.NewCgroupManager("mydocker-cgroup", &subsystems.ResourceConfig{
		CpuShare:    "512",
		MemoryLimit: "5m",
	})
	defer cm.Destroy()
	cm.Set()
	cm.Apply(parent.Process.Pid)

	parent.Wait()
}
