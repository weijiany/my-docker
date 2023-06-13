package subsystems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func cgroupMountInfoOf(subsystem string) string {
	file, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lines := strings.Split(text, "\n")
		for _, line := range lines {
			if strings.Contains(line, subsystem) {
				compile, err := regexp.Compile("(/\\w+)+")
				if err != nil {
					return ""
				}
				return compile.FindString(line)
			}
		}
	}

	return ""
}

func cgroupPathOf(subsystem string, group string) (string, error) {
	cgroupRoot := cgroupMountInfoOf(subsystem)
	cgroupPath := path.Join(cgroupRoot, group)
	_, err := os.Stat(cgroupPath)
	if err != nil && os.IsNotExist(err) {
		if err = os.Mkdir(cgroupPath, 0755); err != nil {
			return "", fmt.Errorf("error: create cgroup: %v", err)
		}
	}
	if err != nil {
		return "", fmt.Errorf("cgroup path error: %v", err)
	}
	return cgroupPath, nil
}
