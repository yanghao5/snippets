package main

import (
	"fmt"
	"os"
	"os/exec"
)

func executeCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run() // 执行命令
	if err != nil {
		return fmt.Errorf("error executing command: %w", err)
	}
	return nil
}

func main() {
	// Define the commands you want to execute
	// 第一个值应该是某个可执行文件的名称，不应该是包含空格的字符串
	commands := [][]string{
		{"echo", "Hello, World!"},
		{"sleep", "10"}, // Example of a long-running command
		{"ls", "-la"},
		{"pwd"},
		// Add more commands as needed
	}

	for _, cmd := range commands {
		if len(cmd) > 0 {
			command := cmd[0]
			args := cmd[1:]
			fmt.Printf("Executing command: %s %v\n", command, args)
			err := executeCommand(command, args)
			if err != nil {
				fmt.Printf("Command failed: %v\n", err)
				break
			}
		}
	}
}
