package container

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
	"weijiany/docker/src/subsystems"
)

func newParentProcess(tty bool, command string) *exec.Cmd {
	args := []string{"init", command}
	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | // Unix Timesharing System 用于隔离和管理主机的主机名和域名信息
			syscall.CLONE_NEWPID | // Process Identifier 用于隔离和管理进程的标识符（PID）
			syscall.CLONE_NEWNS | // Mount Namespace Group 用于隔离和管理文件系统挂载点
			syscall.CLONE_NEWNET | // Network Namespace 用于隔离和管理网络栈和网络资源
			syscall.CLONE_NEWIPC, // Inter-Process Communication 用于隔离和管理进程间通信的资源，如消息队列、信号量和共享内存等
	}
	if tty {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
	return cmd
}

func Run(tty bool, command string, resourceConfig *subsystems.ResourceConfig) error {
	parent := newParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error("start error: ", err.Error())
		return err
	}

	cm := subsystems.NewCgroupManager("mydocker-cgroup", resourceConfig)
	defer cm.Destroy()
	cm.Set()
	cm.Apply(parent.Process.Pid)

	if err := parent.Wait(); err != nil {
		return err
	}
	return nil
}
