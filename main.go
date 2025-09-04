package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func runComposeCmd(dir, action string) error {
	var cmd *exec.Cmd

	switch action {
	case "update":
		// update = pull + up -d
		cmd = exec.Command("docker", "compose", "pull")
		cmd.Dir = dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("[%s] pull failed: %w", dir, err)
		}

		cmd = exec.Command("docker", "compose", "up", "-d")
	case "start":
		cmd = exec.Command("docker", "compose", "start")
	case "stop":
		cmd = exec.Command("docker", "compose", "stop")
	case "restart":
		cmd = exec.Command("docker", "compose", "restart")
	default:
		return fmt.Errorf("未知操作: %s", action)
	}

	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: composeman [命令] <目录1> <目录2> ...")
		fmt.Println("命令支持: update(默认), start, stop, restart")
		os.Exit(1)
	}

	action := "update"
	args := os.Args[1:]

	// 如果第一个参数是命令
	switch args[0] {
	case "update", "start", "stop", "restart":
		action = args[0]
		args = args[1:]
	}

	if len(args) == 0 {
		fmt.Println("需要至少一个目录")
		os.Exit(1)
	}

	for _, dir := range args {
		absDir, _ := filepath.Abs(dir)
		fmt.Printf("===> [%s] 执行 %s\n", absDir, action)
		if err := runComposeCmd(absDir, action); err != nil {
			fmt.Printf("目录 %s 执行失败: %v\n", absDir, err)
		} else {
			fmt.Printf("目录 %s 执行完成\n", absDir)
		}
	}
}
