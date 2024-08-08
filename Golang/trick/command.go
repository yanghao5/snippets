package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func executeCommand(commandStr string) error {
	parse := strings.Fields(commandStr)
	command := parse[0]
	args := parse[1:]

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
	// 获取环境变量中的密码
	password := os.Getenv("DOCKER_ALIYUN_PASSWORD")
	if password == "" {
		fmt.Println("Environment variable DOCKER_ALIYUN_PASSWORD is not set")
		return
	}
	
	login_cmd := fmt.Sprintf("echo %s | docker login --username=xxxxx --password-stdin registry.aliyuncs.com", password)
	
	commands := []string{
		//login_cmd,
		"docker build --tag registry.aliyuncs.com/xxxxx/myalpine:01 .",
		"docker push registry.aliyuncs.com/xxxxx/myalpine:01",
		"docker rmi registry.aliyuncs.com/xxxxx/myalpine:01",
	}

	for _, cmd := range commands {

		fmt.Printf("Executing command: %s\n", cmd)
		err := executeCommand(cmd)
		if err != nil {
			fmt.Printf("Command failed: %v\n", err)
			break
		}
	}
}
